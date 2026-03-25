package errcore

import (
	"errors"
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func SliceToError(errorSlice []string) error {
	if len(errorSlice) == 0 {
		return nil
	}

	fullError := strings.Join(
		errorSlice,
		constants.NewLineUnix)

	return errors.New(fullError)
}
