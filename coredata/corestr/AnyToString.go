package corestr

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToString(
	isIncludeFieldName bool,
	any interface{},
) string {
	if any == "" {
		return constants.EmptyString
	}

	if isIncludeFieldName {
		return fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			any)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any)
}
