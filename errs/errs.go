package errs

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
}

func NewInternalError(message string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewCannotBeProcessedError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewConflictError(message string) error {
	return AppError{
		Code:    http.StatusConflict,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) error {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}
