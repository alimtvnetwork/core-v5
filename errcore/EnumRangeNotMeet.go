package errcore

import "fmt"

func EnumRangeNotMeet(
	rangeStart any,
	rangeEnd any,
	wholeRange any,
) string {
	rangeStr := ""

	if wholeRange == nil {
		rangeStr = fmt.Sprintf(rangeWithoutRangeFormat, rangeStart, rangeEnd)
	} else {
		rangeStr = fmt.Sprintf(rangeWithRangeFormat, rangeStart, rangeEnd, wholeRange)
	}

	return EnumValuesOutOfRangeType.String() +
		ReferenceStart +
		rangeStr +
		ReferenceEnd
}
