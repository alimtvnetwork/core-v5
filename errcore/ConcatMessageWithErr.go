package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// ConcatMessageWithErr wraps an error with an additional message using fmt.Errorf and %w.
// This preserves the original error chain, enabling errors.Is and errors.As.
func ConcatMessageWithErr(
	errMessage string,
	err error,
) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s %w", errMessage, err)
}

func ConcatMessageWithErrWithStackTrace(
	errMessage string,
	err error,
) error {
	if err == nil {
		return nil
	}

	fullMessage := fmt.Sprintf(
		"%s - %s %s\n\n%s",
		reflectinternal.CodeStack.MethodName(1),
		errMessage,
		err.Error(),
		reflectinternal.CodeStack.StacksStringDefault(2),
	)

	return fmt.Errorf("%s: %w", fullMessage, err)
}
