package msgtype

import (
	"gitlab.com/evatix-go/core/constants"
)

func ErrorToString(err error) string {
	if err == nil {
		return constants.EmptyString
	}

	return err.Error()
}
