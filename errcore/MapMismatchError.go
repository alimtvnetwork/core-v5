package errcore

import (
	"fmt"
	"strings"
)

// MapMismatchError builds a diagnostic error for map assertion failures.
//
// Each map entry is shown on its own indexed line in Go literal format,
// making the output directly copy-pasteable into _testcases.go.
//
// Output format:
//
//	Map Mismatch (Case 0: title)
//
//	Actual Received (2 entries):
//	  0: "containsName": false,
//	  1: "hasError":      false,
//
//	Expected Input (1 entries):
//	  0: "hasError": false,
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
	for i, line := range actualGoLiteralLines {
		sb.WriteString(fmt.Sprintf("  %d: %s\n", i, line))
	}

	sb.WriteString(fmt.Sprintf(
		"\nExpected Input (%d entries):\n",
		len(expectedGoLiteralLines),
	))
	for i, line := range expectedGoLiteralLines {
		sb.WriteString(fmt.Sprintf("  %d: %s\n", i, line))
	}

	return sb.String()
}
