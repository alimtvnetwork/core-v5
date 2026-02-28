package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func convAnyValToString(val any) string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		val)
}
