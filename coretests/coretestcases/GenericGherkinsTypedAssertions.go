package coretestcases

import (
	"fmt"
	"testing"
)

// ShouldMatchExpected asserts that the given result matches
// the Expected field value. Uses the typed Expected field directly
// instead of converting to string lines.
//
// This is preferred over string-based comparison for boolean tests.
func (it *GenericGherkins[TInput, TExpect]) ShouldMatchExpected(
	t *testing.T,
	caseIndex int,
	result any,
) {
	t.Helper()

	title := it.CaseTitle()
	expected := fmt.Sprintf("%v", it.Expected)
	actual := fmt.Sprintf("%v", result)

	if actual == expected {
		return
	}

	t.Errorf(
		"Case %d (%s): got %s, want %s",
		caseIndex,
		title,
		actual,
		expected,
	)
}

// ShouldMatchExpectedFirst asserts that the given result matches
// the Expected field value using caseIndex=0. Use for named single
// test cases (non-loop).
func (it *GenericGherkins[TInput, TExpect]) ShouldMatchExpectedFirst(
	t *testing.T,
	result any,
) {
	t.Helper()

	it.ShouldMatchExpected(
		t,
		0,
		result,
	)
}
