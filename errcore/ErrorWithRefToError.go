package errcore

import (
	"errors"
)

func ErrorWithRefToError(err error, reference any) error {
	if err == nil {
		return nil
	}

	return errors.New(ErrorWithRef(err, reference))
}
