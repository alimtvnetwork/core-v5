package strutilinternal

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func AnyToString(any any) string {
	if any == nil {
		return ""
	}

	val := ReflectInterfaceVal(any)

	return fmt.Sprintf(
		constants.SprintValueFormat,
		val)
}
