package corecomparatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_Compare_Verification(t *testing.T) {
	for caseIndex, testCase := range compareStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(value)

		// Act
		actual := args.Map{
			"name":      compare.String(),
			"symbol":    compare.OperatorSymbol(),
			"shortName": compare.OperatorShortForm(),
			"isEqual":   compare.IsEqual(),
			"isValid":   compare.IsValid(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
