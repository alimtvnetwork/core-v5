package errcore

import (
	"github.com/alimtvnetwork/core/constants"
)

func ToString(err error) string {
	if err == nil {
		return constants.EmptyString
	}

	return err.Error()
}
