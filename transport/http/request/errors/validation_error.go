package errors

import "github.com/anjulapaulus/iban-api/errors"

// ValidationError is the type of errors during request validation.
type ValidationError struct {
	*errors.GenericError
}

// NewValidationError creates a new ValidationError instance.
func NewValidationError(code, msg string, errs error) error {
	return &ValidationError{
		GenericError: errors.NewGenericError("ValidationError", code, msg, errs),
	}
}
