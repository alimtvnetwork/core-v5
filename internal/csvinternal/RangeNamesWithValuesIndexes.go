package csvinternal

import "strconv"

// RangeNamesWithValuesIndexes
//
//  Returns a new slice where
//  format
//   - `name[ValueIndex]`
//  example
//   - `SomeName[1]`
func RangeNamesWithValuesIndexes(
	rangedItems ...string,
) []string {
	if len(rangedItems) == 0 {
		return []string{}
	}

	compiledRanges := make([]string, len(rangedItems))

	for i, item := range rangedItems {
		compiledRanges[i] = item + "[" + strconv.Itoa(i) + "]"
	}

	return compiledRanges
}
