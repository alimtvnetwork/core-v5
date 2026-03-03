package errcore

import (
	"fmt"
	"testing"
)

// AssertDiffOnMismatch prints a formatted diff diagnostic block
// and calls t.Errorf when actual and expected lines differ.
//
// This combines diagnostics and assertion into one call,
// eliminating the need for a separate ShouldBeEqual invocation.
//
// It prints:
//   - A header with case index and title
//   - Optional context lines (e.g., "  InitValue: hello")
//   - The standard line-by-line diff via PrintLineDiff
//   - A footer closing the block
//
// contextLines are printed as-is between the header and the diff.
// Each context line should be pre-formatted (e.g., fmt.Sprintf("  Key: %v", val)).
func AssertDiffOnMismatch(
	t *testing.T,
	caseIndex int,
	title string,
	actLines []string,
	expectedLines []string,
	contextLines ...string,
) {
	t.Helper()

	if !LineDiffHasMismatch(actLines, expectedLines) {
		return
	}

	fmt.Printf(
		"\n=== Diff (Case %d: %s) ===\n",
		caseIndex,
		title,
	)

	for _, line := range contextLines {
		fmt.Println(line)
	}

	PrintLineDiff(caseIndex, title, actLines, expectedLines)
	fmt.Println("=== End ===")

	t.Errorf("Case %d (%s): actual lines do not match expected", caseIndex, title)
}
