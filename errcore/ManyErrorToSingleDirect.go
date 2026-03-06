package errcore

import "errors"

// ManyErrorToSingleDirect is the variadic wrapper for ManyErrorToSingle.
func ManyErrorToSingleDirect(errorItems ...error) error {
	return errors.Join(errorItems...)
}
