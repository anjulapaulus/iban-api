package errors

import "fmt"

// GenericError is the base error struct that can be used to create different types of errors.
type GenericError struct {
	Type    string `json:"type"`
	Code    string `json:"code"`
	Message string `json:"message"`
	err     error
}

// NewBaseError creates a new BaseError instance.
func NewGenericError(typ, code, msg string, cause error) *GenericError {
	return &GenericError{
		Type:    typ,
		Code:    code,
		Message: msg,
		err:     cause,
	}
}

// Error returns the error message.
func (e *GenericError) Error() string {
	return fmt.Sprintf("%s|%s|%s", e.Type, e.Code, e.Message)
}

// Unwrap returns the wrapped error.
func (e *GenericError) Unwrap() error {
	return e.err
}
