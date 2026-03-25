package enumimpl

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func NameWithValue(
	value any,
) string {
	return fmt.Sprintf(
		constants.EnumNameValueFormat,
		value,
		value)
}
