package csvinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// RangeNamesWithValuesIndexesCsvString
//
//  Returns a new slice to joined
//  string using RangeNamesWithValuesIndexes
//
//  format
//   - `name[ValueIndex]` + ", "
//  example
//   - `SomeName[1]` + ", "
func RangeNamesWithValuesIndexesCsvString(
	rangedItems ...string,
) string {
	return strings.Join(
		RangeNamesWithValuesIndexes(rangedItems...),
		constants.CommaSpace)
}
