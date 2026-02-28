package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func NameWithValue(
	value any,
) string {
	return fmt.Sprintf(
		constants.EnumNameValueFormat,
		value,
		value)
}
