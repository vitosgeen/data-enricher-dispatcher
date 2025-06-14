package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"data-enricher-dispatcher/apperrors"
	"data-enricher-dispatcher/config"
	"data-enricher-dispatcher/model"
)

const (
	defaultAttempts                  = 3
	emptyResponseError               = "empty response body length: %d"
	unexpectedStatusCodeError        = "unexpected status code: %d"
	unexpectedStatusCodeAttemtsError = "unexpected status code: %d, attempts: %d"
	failedGetUsersError              = "failed to get users after %d attempts"
)

type APIClient interface {
	GetUsers() ([]model.User, error)
	PostUser(user model.User) error
}
type apiClient struct {
	client      *http.Client
	getUsersUrl string
	postUserUrl string
}

func NewAPIClient(cfg *config.Config) APIClient {
	return &apiClient{
		client:      &http.Client{},
		getUsersUrl: cfg.GetUsersURL,
		postUserUrl: cfg.PostUsersURL,
	}
}

func (c *apiClient) GetUsers() ([]model.User, error) {
	var users []model.User
	for i := 0; i < defaultAttempts; i++ {
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
			err = fmt.Errorf(unexpectedStatusCodeAttemtsError, resp.StatusCode, i+1)
			if i == defaultAttempts-1 {
				return nil, apperrors.ApiClientGetUsersStatusCodeNotOkError.AppendMessage(err)
			}
			//nolint:errcheck
			apperrors.ApiClientGetUsersStatusCodeNotOkError.AppendMessage(err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, apperrors.ApiClientGetUsersReadAllError.AppendMessage(err)
		}
		if len(body) == 0 {
			err = fmt.Errorf(emptyResponseError, len(body))
			return nil, apperrors.ApiClientGetUsersEmptyResponseError.AppendMessage(err)
		}

		err = json.Unmarshal(body, &users)
		if err != nil {
			return nil, apperrors.ApiClientGetUsersUnmarshalError.AppendMessage(err)
		}

		return users, nil
	}

	return nil, apperrors.ApiClientGetUsersAttemptsExceededError.AppendMessage(fmt.Errorf(failedGetUsersError, defaultAttempts))
}

func (c *apiClient) PostUser(user model.User) error {
	if !user.IsValid() {
		return apperrors.ApiClientPostUserIsValidError.AppendMessage(fmt.Errorf("invalid user: %v", user))
	}
	for i := 0; i < defaultAttempts; i++ {
		userData, err := json.Marshal(user)
		if err != nil {
			return apperrors.ApiClientPostUserMarshalError.AppendMessage(err)
		}
		resp, err := c.client.Post(c.postUserUrl, "application/json", bytes.NewBuffer(userData))
		if err != nil {
			return apperrors.ApiClientPostUserPostError.AppendMessage(err)
		}
		defer func() {
			if closeErr := resp.Body.Close(); closeErr != nil {
				err = apperrors.ApiClientPostUserCloseBodyError.AppendMessage(closeErr)
				return
			}
		}()
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf(unexpectedStatusCodeAttemtsError, resp.StatusCode, i+1)
			if i == defaultAttempts-1 {
				return apperrors.ApiClientPostUserStatusCodeNotOkError.AppendMessage(err)
			}
			//nolint:errcheck
			apperrors.ApiClientPostUserStatusCodeNotOkError.AppendMessage(err)
			continue
		} else {
			return nil
		}
	}
	return nil
}
