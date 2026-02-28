package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_BaseIdentifier_Verification(t *testing.T) {
	for caseIndex, testCase := range identifierTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		id, _ := input.GetAsString("id")

		// Act
		identifier := coreinstruction.NewIdentifier(id)
		idStr := identifier.IdString()
		isEmpty := fmt.Sprintf("%v", identifier.IsIdEmpty())
		isWhitespace := fmt.Sprintf("%v", identifier.IsIdWhitespace())

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			idStr,
			isEmpty,
			isWhitespace,
		)
	}
}
