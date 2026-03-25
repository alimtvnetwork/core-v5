package errcore

import (
	"testing"
)

// AssertDiffOnMismatch prints a formatted diff diagnostic block
// and calls t.Errorf when actual and expected lines differ.
//
// This combines diagnostics and assertion into one call,
// eliminating the need for a separate ShouldBeEqual invocation.
//
// It delegates to PrintDiffOnMismatch for diagnostics output,
// then marks the test as failed via t.Errorf.
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

	if !HasAnyMismatchOnLines(actLines, expectedLines) {
		return
	}

	PrintDiffOnMismatch(caseIndex, title, actLines, expectedLines, contextLines...)
	t.Errorf("Case %d (%s): actual lines do not match expected", caseIndex, title)
}

// AssertErrorDiffOnMismatch splits an error into lines,
// prints a formatted diff diagnostic block via PrintDiffOnMismatch,
// and calls t.Errorf when actual and expected lines differ.
//
// If err is nil, actual lines will be empty.
//
// contextLines are printed as-is between the header and the diff.
func AssertErrorDiffOnMismatch(
	t *testing.T,
	caseIndex int,
	title string,
	err error,
	expectedLines []string,
	contextLines ...string,
) {
	t.Helper()

	actLines := ErrorToSplitLines(err)
	if err == nil {
		actLines = []string{}
	}

	if !HasAnyMismatchOnLines(actLines, expectedLines) {
		return
	}

	PrintDiffOnMismatch(caseIndex, title, actLines, expectedLines, contextLines...)
	t.Errorf("Case %d (%s): error lines do not match expected", caseIndex, title)
}
