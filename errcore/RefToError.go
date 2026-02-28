package errcore

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func RefToError(reference any) error {
	if reference == nil {
		return nil
	}

	return errors.New(fmt.Sprintf(
		constants.ReferenceWrapFormat,
		reference),
	)
}
