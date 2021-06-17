package stringutil

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToTypeString(any interface{}) string {
	return fmt.Sprintf(constants.SprintTypeFormat, any)
}
