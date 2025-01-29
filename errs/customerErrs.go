package errs

import "errors"

// Define a custom error for email already registered
var ErrEmailAlreadyRegistered = errors.New("email already registered")
