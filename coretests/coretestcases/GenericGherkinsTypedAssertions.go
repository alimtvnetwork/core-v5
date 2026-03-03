package coretestcases

import (
	"fmt"
	"testing"
)

// ShouldExpectedMatch asserts that the given bool result matches
// the Expected field value. Uses the typed Expected field directly
// instead of converting to string lines.
//
// This is preferred over string-based comparison for boolean tests.
func (it *GenericGherkins[TInput, TExpect]) ShouldExpectedMatch(
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
