package coretestcases

import (
	"testing"

	"github.com/alimtvnetwork/core/errcore"
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
