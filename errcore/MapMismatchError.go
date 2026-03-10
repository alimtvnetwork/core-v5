package errcore

import (
	"fmt"
	"strings"
)

// MapMismatchError builds a diagnostic error for map assertion failures.
//
// Each map entry is shown on its own line with tab indentation in Go literal
// format, making the output directly copy-pasteable into _testcases.go.
//
// Output format:
//
//	Map Mismatch (Case 0: title)
//
//	Actual Received (2 entries):
//	  "containsName": false,
//	  "hasError":      false,
//
//	Expected Input (1 entries):
//	  "hasError": false,
func MapMismatchError(
	caseIndex int,
	title string,
	actualGoLiteralLines []string,
	expectedGoLiteralLines []string,
) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(
		"Map Mismatch (Case %d: %s)\n\n",
		caseIndex,
		title,
	))

	sb.WriteString(fmt.Sprintf(
		"Actual Received (%d entries):\n",
		len(actualGoLiteralLines),
	))

	for _, line := range actualGoLiteralLines {
		sb.WriteString("\t")
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf(
		"\nExpected Input (%d entries):\n",
		len(expectedGoLiteralLines),
	))

	for _, line := range expectedGoLiteralLines {
		sb.WriteString("\t")
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	return sb.String()
}
