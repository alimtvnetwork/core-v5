package errcore

import (
	"fmt"
	"strings"
)

const mapMismatchSeparator = "============================>"

// MapMismatchError builds a diagnostic error for map assertion failures.
//
// Each map entry is shown line-by-line with aligned actual/expected labels
// and the standard header separators. No trailing commas on entries.
//
// Output format:
//
//	============================>
//	Map Mismatch (Case 0: title)
//	Actual lines: 2, Expected lines: 1
//	============================>
//		actual   : containsName : false
//		expected : hasError : false
//	============================>
//		actual   : hasError : false
//		expected : <missing>
//	============================>
func MapMismatchError(
	caseIndex int,
	title string,
	actualLines []string,
	expectedLines []string,
) string {
	var sb strings.Builder

	maxLen := len(actualLines)

	if len(expectedLines) > maxLen {
		maxLen = len(expectedLines)
	}

	sb.WriteString("\n")
	sb.WriteString(mapMismatchSeparator)
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf(
		"%d ) Map Mismatch:\n    %s",
		caseIndex,
		title,
	))
	sb.WriteString("\n")
	sb.WriteString(mapMismatchSeparator)
	sb.WriteString(fmt.Sprintf(
		"\n    Actual lines: %d, Expected lines: %d\n",
		len(actualLines),
		len(expectedLines),
	))

	for i := 0; i < maxLen; i++ {
		sb.WriteString(mapMismatchSeparator)
		sb.WriteString("\n")

		actualVal := "<missing>"
		expectedVal := "<missing>"

		if i < len(actualLines) {
			actualVal = actualLines[i]
		}

		if i < len(expectedLines) {
			expectedVal = expectedLines[i]
		}

		sb.WriteString(fmt.Sprintf(
			"\tactual   : %s\n"+
				"\texpected : %s\n",
			actualVal,
			expectedVal,
		))
	}

	sb.WriteString(mapMismatchSeparator)
	sb.WriteString("\n")

	return sb.String()
}
