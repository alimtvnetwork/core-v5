package errcore

import (
	"fmt"
	"log/slog"
)

// PrintDiffOnMismatch prints a formatted diff diagnostic block
// only when actual and expected lines differ.
//
// It prints:
//   - A header with case index and title
//   - Optional context lines (e.g., "  InitValue: hello")
//   - The standard line-by-line diff via PrintLineDiff
//   - A footer closing the block
//
// contextLines are printed as-is between the header and the diff.
// Each context line should be pre-formatted (e.g., fmt.Sprintf("  Key: %v", val)).
func PrintDiffOnMismatch(
	caseIndex int,
	title string,
	actLines []string,
	expectedLines []string,
	contextLines ...string,
) {
	if !HasAnyMismatchOnLines(actLines, expectedLines) {
		return
	}

	header := fmt.Sprintf("\n=== Diff (Case %d: %s) ===\n", caseIndex, title)
	fmt.Print(header)

	for _, cl := range contextLines {
		fmt.Println(cl)
	}

	PrintLineDiff(caseIndex, title, actLines, expectedLines)

	fmt.Printf("=== End Diff (Case %d) ===\n", caseIndex)
}
