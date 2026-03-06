package errcore

import (
	"errors"
	"fmt"
)

// CompiledError wraps a main error with an additional message, preserving the error chain.
func CompiledError(mainErr error, additionalMessage string) error {
	if mainErr == nil {
		return nil
	}

	if additionalMessage == "" {
		return mainErr
	}

	return fmt.Errorf("%s: %w",
		additionalMessage, mainErr)
}

// CompiledErrorString returns the compiled error string with main error and additional message.
// Kept for backward compatibility with string-based consumers.
func CompiledErrorString(mainErr error, additionalMessage string) string {
	if mainErr == nil {
		return ""
	}

	compiled := CompiledError(mainErr, additionalMessage)
	if compiled == nil {
		return ""
	}

	return compiled.Error()
}

// JoinErrors is a convenience alias for errors.Join, provided for discoverability
// within the errcore package.
func JoinErrors(errs ...error) error {
	return errors.Join(errs...)
}
