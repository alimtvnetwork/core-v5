package corecsv

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

// AnyToValuesTypeString
//
// Output : 'value - type', 'value - type'...
func AnyToValuesTypeString(
	references ...any,
) string {
	if len(references) == 0 {
		return ""
	}

	toSlice := AnyToValuesTypeStrings(references...)

	return strings.Join(
		toSlice,
		constants.CommaSpace,
	)
}
