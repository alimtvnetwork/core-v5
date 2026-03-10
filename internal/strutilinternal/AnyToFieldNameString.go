package strutilinternal

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func AnyToFieldNameString(any any) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		any)
}
