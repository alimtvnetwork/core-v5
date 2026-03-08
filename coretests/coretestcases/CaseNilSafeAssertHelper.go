package coretestcases

import (
	"testing"

	"gitlab.com/auk-go/core/errcore"
)

// assertDiffOnMismatch delegates to errcore.AssertDiffOnMismatch
// for consistent diff-based failure output.
func assertDiffOnMismatch(
	t *testing.T,
	caseIndex int,
	title string,
	actLines []string,
	expectedLines []string,
) {
	t.Helper()

	errcore.AssertDiffOnMismatch(
		t,
		caseIndex,
		title,
		actLines,
		expectedLines,
	)
}
