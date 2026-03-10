package errcore

import (
	"github.com/alimtvnetwork/core/constants"
)

func MergeErrorsToStringDefault(
	errorItems ...error,
) string {
	if errorItems == nil {
		return ""
	}

	return MergeErrorsToString(constants.Space, errorItems...)
}
