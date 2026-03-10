package corecsv

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

// AnyToTypesCsvDefault
//
// if references empty or len 0 then empty string returned.
//
// Formats :
//   - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//   - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//   - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func AnyToTypesCsvDefault(
	references ...any,
) string {
	toSlice := AnyToTypesCsvStrings(
		false,
		false,
		references...)

	return strings.Join(
		toSlice,
		constants.CommaSpace,
	)
}
