package stringutil

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// AnyToStringNameField
//
//	Returns string names and values using (%+v)
func AnyToStringNameField(any any) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		any)
}
