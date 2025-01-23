package errs

import "github.com/gofiber/fiber/v2"

type FiberError struct {
	Code    int
	Message string
}

func (e FiberError) Error() string {
	return e.Message
}

func NewFiberError(code int, message string) error {
	return FiberError{
		Code:    code,
		Message: message,
	}
}

func NewFiberNotFoundError(message string) error {
	return FiberError{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}

func NewFiberUnexpectedError() error {
	return FiberError{
		Code:    fiber.StatusInternalServerError,
		Message: "something went wrong",
	}
}

func NewFiberValidationError(message string) error {
	return FiberError{
		Code:    fiber.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewFiberBadRequestError(message string) error {
	return FiberError{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}

func NewFiberUnauthorizedError(message string) error {
	return FiberError{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	}
}

func NewFiberForbiddenError(message string) error {
	return FiberError{
		Code:    fiber.StatusForbidden,
		Message: message,
	}
}

func NewFiberConflictError(message string) error {
	return FiberError{
		Code:    fiber.StatusConflict,
		Message: message,
	}
}
