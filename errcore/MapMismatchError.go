package errcore

import (
	"fmt"
	"strings"
)

const mapSeparator = "============================>"

// MapMismatchError builds a diagnostic error for map assertion failures.
//
// Each map entry is shown on its own line with tab indentation in Go literal
// format, making the output directly copy-pasteable into _testcases.go.
//
// Output format:
//
//	============================>
//	Actual Received (2 entries):
//	    Case Title
//	============================>
//		"containsName": false,
//		"hasError":      false,
//	============================>
//
//	============================>
//	Expected Input (1 entries):
//	    Case Title
//	============================>
//		"hasError": false,
//	============================>
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

	// Actual block
	sb.WriteString(mapSeparator + "\n")
	sb.WriteString(fmt.Sprintf(
		"%d) Actual Received (%d entries):\n",
		caseIndex,
		len(actualGoLiteralLines),
	))
	sb.WriteString(fmt.Sprintf("    %s\n", title))
	sb.WriteString(mapSeparator + "\n")

	for _, line := range actualGoLiteralLines {
		sb.WriteString("\t")
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	sb.WriteString(mapSeparator + "\n")

	// Expected block
	sb.WriteString("\n")
	sb.WriteString(mapSeparator + "\n")
	sb.WriteString(fmt.Sprintf(
		"%d) Expected Input (%d entries):\n",
		caseIndex,
		len(expectedGoLiteralLines),
	))
	sb.WriteString(fmt.Sprintf("    %s\n", title))
	sb.WriteString(mapSeparator + "\n")

	for _, line := range expectedGoLiteralLines {
		sb.WriteString("\t")
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	sb.WriteString(mapSeparator)

	return sb.String()
}
