package errcore

import (
	"strings"

	"github.com/alimtvnetwork/core/namevalue"
)

func VarNameValuesJoiner(
	joiner string,
	nameValues ...namevalue.StringAny,
) string {
	if len(nameValues) == 0 {
		return ""
	}

	items := VarNameValuesStrings(nameValues...)

	return strings.Join(
		items,
		joiner)
}
