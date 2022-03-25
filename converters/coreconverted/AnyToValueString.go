package coreconverted

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToValueString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem)
}
