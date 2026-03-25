package errcore

import (
	"fmt"

	"github.com/alimtvnetwork/core/internal/msgformats"
)

func GherkinsStringWithExpectation(
	testCaseIndex int,
	feature,
	given,
	when,
	then,
	actual,
	expectation any,
) string {
	return fmt.Sprintf(
		msgformats.SimpleGherkinsWithExpectationFormat,
		testCaseIndex,
		feature,
		testCaseIndex,
		given,
		when,
		then,
		actual,
		expectation,
	)
}
