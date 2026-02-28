package stringutil

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func AnyToString(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(constants.SprintValueFormat, anyItem)
}
