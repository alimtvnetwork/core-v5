package coretestcases

import (
	"testing"

	"gitlab.com/auk-go/core/errcore"
)

// ShouldBeEqual asserts that actLines match ExpectedLines using
// the struct's Title as the test title.
//
// This delegates to errcore.AssertDiffOnMismatch for diff-based
// assertion output on failure.
func (it *GenericGherkins[TInput, TExpect]) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	actLines []string,
	expectedLines []string,
) {
	t.Helper()

	title := it.Title
	if title == "" {
		title = it.When
	}

	errcore.AssertDiffOnMismatch(
		t,
		caseIndex,
		title,
		actLines,
		expectedLines,
	)
}

// ShouldBeEqualArgs asserts that individual string arguments match
// ExpectedLines using the struct's Title as the test title.
//
// Each argument is treated as one line. This allows callers to
// pass each value on its own line for readability:
//
//	tc.ShouldBeEqualArgs(
//	    t,
//	    caseIndex,
//	    fmt.Sprintf("%v", called),
//	    fmt.Sprintf("%v", result),
//	)
func (it *GenericGherkins[TInput, TExpect]) ShouldBeEqualArgs(
	t *testing.T,
	caseIndex int,
	actLines ...string,
) {
	t.Helper()

	it.ShouldBeEqual(
		t,
		caseIndex,
		actLines,
		it.ExpectedLines,
	)
}

// ShouldBeEqualUsingExpected asserts that actLines match the struct's
// own ExpectedLines field. Useful when expected values are defined
// in the test case data itself.
func (it *GenericGherkins[TInput, TExpect]) ShouldBeEqualUsingExpected(
	t *testing.T,
	caseIndex int,
	actLines []string,
) {
	t.Helper()

	it.ShouldBeEqual(
		t,
		caseIndex,
		actLines,
		it.ExpectedLines,
	)
}
