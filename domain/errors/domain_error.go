package errors

import "github.com/anjulapaulus/iban-api/errors"

// DomainError is the errors related to domain.
type DomainError struct {
	*errors.GenericError
}

// NewDomainError creates a new DomainError instance.
func NewDomainError(code, msg string, errs error) error {
	return &DomainError{
		GenericError: errors.NewGenericError("DomainError", code, msg, errs),
	}
}
