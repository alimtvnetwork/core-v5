package corestr

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func AnyToString(
	isIncludeFieldName bool,
	any any,
) string {
	if any == "" {
		return constants.EmptyString
	}

	val := reflectInterfaceVal(any)

	if isIncludeFieldName {
		return fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			val)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		val)
}
