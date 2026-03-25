package strutilinternal

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func AnyToStringUsing(
	isIncludeFields bool,
	any any,
) string {
	if any == nil {
		return ""
	}

	if isIncludeFields {
		return fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			any)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any)
}
