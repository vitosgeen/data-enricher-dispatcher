package apperrors

import "net/http"

var (
	ServiceDispatcherError = &AppError{
		Message:  "Failed to dispatch service request",
		Code:     "SERVICE_DISPATCHER_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServiceDispatcherStartError = &AppError{
		Message:  "Failed to start dispatcher service",
		Code:     "SERVICE_DISPATCHER_START_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServiceDispatcherGetUsersError = &AppError{
		Message:  "Failed to get users in dispatcher service",
		Code:     "SERVICE_DISPATCHER_GET_USERS_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServiceDispatcherInvalidUserError = &AppError{
		Message:  "Invalid user data in dispatcher service",
		Code:     "SERVICE_DISPATCHER_INVALID_USER_ERROR",
		HTTPCode: http.StatusBadRequest,
	}
	ServiceDispatcherPostUserError = &AppError{
		Message:  "Failed to post user in dispatcher service",
		Code:     "SERVICE_DISPATCHER_POST_USER_ERROR",
		HTTPCode: http.StatusInternalServerError,
	}
	ServiceDispatcherSkippingUserEmailWithSpecialPostfix = &AppError{
		Message:  "Skipping user with email due to special postfix exclusion",
		Code:     "SERVICE_DISPATCHER_SKIPPING_USER_EMAIL_WITH_SPECIAL_POSTFIX",
		HTTPCode: http.StatusOK,
	}
)
