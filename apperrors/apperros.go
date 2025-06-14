package apperrors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message  string
	Code     string
	HTTPCode int
}

var (
	EnvConfigLoadError = AppError{
		Message:  "Failed to load env file",
		Code:     "ENV_INIT_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigParseError = AppError{
		Message:  "Failed to parse env file",
		Code:     "ENV_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	EnvConfigPostgresParseError = AppError{
		Message:  "Failed to parse pastgres env file",
		Code:     "ENV_POSTGRES_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message:  fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:     appError.Code,
		HTTPCode: appError.HTTPCode,
	}
}

func Is(err1 error, err2 *AppError) bool {
	err, ok := err1.(*AppError)
	if !ok {
		return false
	}

	return err.Code == err2.Code
}
