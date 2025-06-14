package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"data-enricher-dispatcher/apperrors"
	"data-enricher-dispatcher/config"
	"data-enricher-dispatcher/model"
)

const (
	defaultWaitTime     = 2 * time.Second
	defaultTimeout      = 10 * time.Second
	failedRetryMessage  = "failed to make request after %d attempts"
	emptyTargetURLError = "target URL cannot be empty"
	invalidUserError    = "invalid user data: %v"
)

type apiClientV2 struct {
	client      *http.Client
	getUsersUrl string
	postUserUrl string
}

func NewAPIClientV2(cfg *config.Config) APIClient {
	return &apiClientV2{
		client:      &http.Client{},
		getUsersUrl: cfg.GetUsersURL,
		postUserUrl: cfg.PostUsersURL,
	}
}

func (c *apiClientV2) GetUsers() ([]model.User, error) {
	resp, err := c.client.Get(c.getUsersUrl)
	if err != nil {
		return nil, apperrors.ApiClientGetUsersGetError.AppendMessage(err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = apperrors.ApiClientGetUsersCloseBodyError.AppendMessage(closeErr)
			return
		}
	}()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf(unexpectedStatusCodeError, resp.StatusCode)
		return nil, apperrors.ApiClientGetUsersStatusCodeNotOkError.AppendMessage(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, apperrors.ApiClientGetUsersReadAllError.AppendMessage(err)
	}
	if len(body) == 0 {
		err = fmt.Errorf(emptyResponseError, len(body))
		return nil, apperrors.ApiClientGetUsersEmptyResponseError.AppendMessage(err)
	}

	var users []model.User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, apperrors.ApiClientGetUsersUnmarshalError.AppendMessage(err)
	}

	return users, nil
}

func (c *apiClientV2) PostUser(user model.User) error {
	if !user.IsValid() {
		return apperrors.ApiClientPostUserIsValidError.AppendMessage(fmt.Errorf(invalidUserError, user))
	}
	userData, err := json.Marshal(user)
	if err != nil {
		return apperrors.ApiClientPostUserMarshalError.AppendMessage(err)
	}
	userDataBuffer := bytes.NewBuffer(userData)
	resp, err := makePostRequestWithRetry(context.Background(), http.MethodPost, c.postUserUrl, "application/json", defaultAttempts, userDataBuffer, defaultTimeout)
	if err != nil {
		return apperrors.ApiClientPostUserPostError.AppendMessage(err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = apperrors.ApiClientPostUserCloseBodyError.AppendMessage(closeErr)
			return
		}
	}()

	return nil
}

func makePostRequestWithRetry(ctx context.Context, method, targetURL, contentType string, attempts int, body *bytes.Buffer, timeout time.Duration) (*http.Response, error) {
	if attempts <= 0 {
		attempts = defaultAttempts
	}
	for i := 0; i < attempts; i++ {
		resp, err := makePostRequestWithContext(ctx, method, targetURL, contentType, body, timeout)
		if err != nil {
			return nil, apperrors.ApiClientMakePostRequestWithRetryMakeRequestError.AppendMessage(err)
		}
		if resp.StatusCode != http.StatusOK {
			time.Sleep(defaultWaitTime) // Wait before retrying
			err = fmt.Errorf(unexpectedStatusCodeAttemtsError, resp.StatusCode, i+1)
			if i == defaultAttempts-1 {
				return nil, apperrors.ApiClientMakePostRequestWithRetryStatusCodeNotOkError.AppendMessage(err)
			}
			continue
		}
		return resp, nil
	}
	return nil, apperrors.ApiClientMakePostRequestWithRetryAttemptsExceededError.AppendMessage(fmt.Errorf(failedRetryMessage, attempts))
}

func makePostRequestWithContext(ctx context.Context, method, targetURL, contentType string, body *bytes.Buffer, timeout time.Duration) (*http.Response, error) {
	if timeout <= 0 {
		timeout = defaultTimeout // Use default timeout if not specified
	}
	if ctx == nil {
		ctx = context.Background() // Use background context if nil
	}
	if method == "" {
		method = http.MethodGet // Default to GET if no method is specified
	}
	if targetURL == "" {
		return nil, apperrors.ApiClientMakeRequestWithContextTargetURLError.AppendMessage(emptyTargetURLError)
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, targetURL, body)
	if err != nil {
		return nil, apperrors.ApiClientMakeRequestWithContextNewRequestWithContextError.AppendMessage(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, apperrors.ApiClientMakeRequestWithContextDoError.AppendMessage(err)
	}

	return resp, nil
}
