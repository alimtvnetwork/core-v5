package errcore

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func ConcatMessageWithErr(
	errMessage string,
	err error,
) error {
	if err == nil {
		return nil
	}

	return errors.New(errMessage + " " + err.Error())
}

func ConcatMessageWithErrWithStackTrace(
	errMessage string,
	err error,
) error {
	if err == nil {
		return nil
	}

	fullMessage := fmt.Sprintf(
		"%s %s\n%s",
		errMessage,
		err.Error(),
		reflectinternal.CodeStack.StacksStringDefault(2),
	)

	return errors.New(fullMessage)
}
