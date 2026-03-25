package stringutil

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func AnyToTypeString(any any) string {
	return fmt.Sprintf(constants.SprintTypeFormat, any)
}
