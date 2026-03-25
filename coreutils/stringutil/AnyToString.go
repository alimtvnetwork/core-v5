package stringutil

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func AnyToString(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(constants.SprintValueFormat, anyItem)
}
