package errcore

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func ErrorToSplitLines(err error) []string {
	if err == nil {
		return []string{}
	}

	return strings.Split(
		err.Error(),
		constants.NewLineUnix)
}
