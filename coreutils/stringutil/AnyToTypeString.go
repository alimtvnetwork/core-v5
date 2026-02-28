package stringutil

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func AnyToTypeString(any any) string {
	return fmt.Sprintf(constants.SprintTypeFormat, any)
}
