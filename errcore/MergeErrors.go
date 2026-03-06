package errcore

import (
	"errors"
)

// MergeErrors combines multiple errors into a single error using errors.Join.
// Unlike the previous string-concatenation approach, this preserves the original
// error chain, enabling errors.Is and errors.As to work on constituent errors.
func MergeErrors(errorItems ...error) error {
	return errors.Join(errorItems...)
}
