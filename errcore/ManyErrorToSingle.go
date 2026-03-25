package errcore

import "errors"

// ManyErrorToSingle combines a slice of errors into a single error using errors.Join.
// Nil errors are automatically filtered. The resulting error supports
// errors.Is and errors.As for all constituent errors.
func ManyErrorToSingle(errorItems []error) error {
	return errors.Join(errorItems...)
}
