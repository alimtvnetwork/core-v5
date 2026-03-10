package errcore

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func VarMap(
	mappedItems map[string]any,
) string {
	if len(mappedItems) == 0 {
		return ""
	}

	items := VarMapStrings(mappedItems)

	return strings.Join(
		items,
		constants.CommaSpace)
}
