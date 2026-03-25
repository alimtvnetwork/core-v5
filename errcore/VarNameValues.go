package errcore

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/namevalue"
)

func VarNameValues(
	nameValues ...namevalue.StringAny,
) string {
	if len(nameValues) == 0 {
		return ""
	}

	items := VarNameValuesStrings(nameValues...)

	return strings.Join(
		items,
		constants.CommaSpace)
}
