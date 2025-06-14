package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"data-enricher-dispatcher/apperrors"
	"data-enricher-dispatcher/config"
	"data-enricher-dispatcher/model"
)

func Test_apiClient_GetUsers(t *testing.T) {
	testCases := []struct {
		name        string
		mockUsers   []model.User
		statusCode  int
		expectedErr string
	}{
		{
			name:        "successful get users",
			mockUsers:   []model.User{{Name: "John Doe", Email: "email1@email.com"}},
			statusCode:  http.StatusOK,
			expectedErr: "",
		},
		{
			name:        "empty response body",
			mockUsers:   []model.User{},
			statusCode:  http.StatusOK,
			expectedErr: "empty response body length: 0",
		},
		{
			name:       "unexpected status code",
			mockUsers:  []model.User{{Name: "John Doe", Email: "email2@email.com"}},
			statusCode: http.StatusInternalServerError,
			expectedErr: apperrors.ApiClientGetUsersStatusCodeNotOkError.AppendMessage(
				"unexpected status code: 500, attempts: 3").Error(),
		},
		{
			name:        "invalid json response",
			mockUsers:   []model.User{{Name: "John Doe", Email: "email3@email.com"}},
			statusCode:  http.StatusOK,
			expectedErr: "invalid character '}' looking for beginning of value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, tc.mockUsers, tc.statusCode)
			defer server.Close()

			client := NewAPIClient(&config.Config{
				GetUsersURL: server.URL,
			})

			users, err := client.GetUsers()
			if err != nil {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error %q, got %q", tc.expectedErr, err.Error())
				}
				return
			}

			if len(users) != len(tc.mockUsers) {
				t.Errorf("expected %d users, got %d", len(tc.mockUsers), len(users))
			}

			for i, user := range users {
				if !user.IsEqual(&tc.mockUsers[i]) {
					t.Errorf("expected user %v, got %v", tc.mockUsers[i], user)
				}
			}
		})
	}
}

func Test_apiClient_PostUser(t *testing.T) {
	testCases := []struct {
		name        string
		mockUsers   model.User
		statusCode  int
		expectedErr string
	}{
		{
			name:        "successful post user",
			mockUsers:   model.User{Name: "John Doe", Email: "email1@email.com"},
			statusCode:  http.StatusOK,
			expectedErr: "",
		},
		{
			name:       "post user with invalid email",
			mockUsers:  model.User{Name: "Invalid User", Email: ""},
			statusCode: http.StatusOK,
			expectedErr: apperrors.ApiClientPostUserIsValidError.AppendMessage(
				fmt.Errorf("invalid user: %v", model.User{Name: "Invalid User", Email: ""})).Error(),
		},
		{
			name:       "unexpected status code on post user",
			mockUsers:  model.User{Name: "John Doe", Email: "email2@email.com"},
			statusCode: http.StatusInternalServerError,
			expectedErr: apperrors.ApiClientPostUserStatusCodeNotOkError.AppendMessage(
				"unexpected status code: 500, attempts: 3").Error(),
		},
		{
			name:        "post user with empty response body",
			mockUsers:   model.User{Name: "John Doe", Email: "email3@email.com"},
			statusCode:  http.StatusOK,
			expectedErr: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, []model.User{tc.mockUsers}, tc.statusCode)
			defer server.Close()

			client := NewAPIClient(&config.Config{
				PostUsersURL: server.URL,
			})

			err := client.PostUser(tc.mockUsers)
			if err != nil {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error %q, got %q", tc.expectedErr, err.Error())
				}
				return
			}

			if tc.expectedErr != "" {
				t.Errorf("expected error %q, got nil", tc.expectedErr)
			}
		})
	}
}

func newTestServer(t *testing.T, mockUsers []model.User, statusCode int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(mockUsers)
		w.WriteHeader(statusCode)
		n, err := w.Write(data)
		if err != nil {
			t.Errorf("failed to write response: %v", err)
		}
		if n != len(data) {
			t.Errorf("expected to write %d bytes, but wrote %d", len(data), n)
		}
	}))

	t.Cleanup(server.Close)
	return server
}
