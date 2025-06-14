package apperrors

import "net/http"

var (
	ApiClientGetUsersGetError = &AppError{
		Message:  "Failed to get users from API",
		Code:     "API_CLIENT_GET_USERS_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientGetUsersCloseBodyError = &AppError{
		Message:  "Failed to close response body from API",
		Code:     "API_CLIENT_GET_USERS_CLOSE_BODY_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientGetUsersStatusCodeNotOkError = &AppError{
		Message:  "API response status code is not OK",
		Code:     "API_CLIENT_GET_USERS_STATUS_CODE_NOT_OK_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientGetUsersReadAllError = &AppError{
		Message:  "Failed to read response body from API",
		Code:     "API_CLIENT_GET_USERS_READ_ALL_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientGetUsersEmptyResponseError = &AppError{
		Message:  "API response body is empty",
		Code:     "API_CLIENT_GET_USERS_EMPTY_RESPONSE_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientGetUsersUnmarshalError = &AppError{
		Message:  "Failed to unmarshal API response",
		Code:     "API_CLIENT_GET_USERS_UNMARSHAL_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientGetUsersAttemptsExceededError = &AppError{
		Message:  "Maximum number of attempts exceeded",
		Code:     "ATTEMPTS_EXCEEDED_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientPostUserMarshalError = &AppError{
		Message:  "Failed to marshal user data for API",
		Code:     "API_CLIENT_POST_USER_MARSHAL_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientPostUserPostError = &AppError{
		Message:  "Failed to post user to API",
		Code:     "API_CLIENT_POST_USER_POST_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientPostUserCloseBodyError = &AppError{
		Message:  "Failed to close response body after posting user to API",
		Code:     "API_CLIENT_POST_USER_CLOSE_BODY_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientPostUserStatusCodeNotOkError = &AppError{
		Message:  "API response status code for post user is not OK",
		Code:     "API_CLIENT_POST_USER_STATUS_CODE_NOT_OK_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientPostUserIsValidError = &AppError{
		Message:  "User data is not valid for API",
		Code:     "API_CLIENT_POST_USER_IS_VALID_ERROR",
		HTTPCode: http.StatusBadRequest,
	}
	ApiClientRetryableMakeRequestError = &AppError{
		Message:  "Retryable make request failed",
		Code:     "API_CLIENT_RETRYABLE_MAKE_REQUEST_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientMakeRequestWithContextTargetURLError = &AppError{
		Message:  "Target URL cannot be empty for make request with context",
		Code:     "API_CLIENT_MAKE_REQUEST_WITH_CONTEXT_TARGET_URL_ERROR",
		HTTPCode: http.StatusBadRequest,
	}
	ApiClientMakePostRequestWithRetryMakeRequestError = &AppError{
		Message:  "Make post request with retry failed",
		Code:     "API_CLIENT_MAKE_POST_REQUEST_WITH_RETRY_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientMakePostRequestWithRetryStatusCodeNotOkError = &AppError{
		Message:  "Make post request with retry status code is not OK",
		Code:     "API_CLIENT_MAKE_POST_REQUEST_WITH_RETRY_STATUS_CODE_NOT_OK_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientMakePostRequestWithRetryAttemptsExceededError = &AppError{
		Message:  "Make post request with retry attempts exceeded",
		Code:     "API_CLIENT_MAKE_POST_REQUEST_WITH_RETRY_ATTEMPTS_EXCEEDED_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientMakeRequestWithContextNewRequestWithContextError = &AppError{
		Message:  "Failed to create new request with context",
		Code:     "API_CLIENT_MAKE_REQUEST_WITH_CONTEXT_NEW_REQUEST_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ApiClientMakeRequestWithContextDoError = &AppError{
		Message:  "Failed to execute request with context",
		Code:     "API_CLIENT_MAKE_REQUEST_WITH_CONTEXT_DO_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
)
