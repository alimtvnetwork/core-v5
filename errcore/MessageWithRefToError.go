package errcore

import (
	"errors"
	"fmt"
)

func MessageWithRefToError(msg string, reference any) error {
	return errors.New(fmt.Sprintf(
		messageMapFormat,
		msg,
		reference),
	)
}
