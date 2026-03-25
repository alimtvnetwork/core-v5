package enumimpl

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func convAnyValToString(val any) string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		val)
}
