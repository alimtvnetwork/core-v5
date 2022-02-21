package enumimpl

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func convAnyValToString(val interface{}) string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		val)
}
