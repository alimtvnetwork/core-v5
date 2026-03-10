package errcore

import "github.com/alimtvnetwork/core/constants"

func SimpleHandleErr(err error, msg string) {
	if err == nil {
		return
	}

	panic(err.Error() + constants.NewLineUnix + msg)
}
