package errcore

import "github.com/alimtvnetwork/core/constants"

func SliceErrorDefault(slice []string) error {
	return SliceError(constants.NewLineUnix, slice)
}
