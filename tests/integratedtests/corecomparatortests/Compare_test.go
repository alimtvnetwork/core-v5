package corecomparatortests

import (
	"fmt"
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
		name := compare.String()
		symbol := compare.OperatorSymbol()
		shortForm := compare.OperatorShortForm()
		isEqual := fmt.Sprintf("%v", compare.IsEqual())
		isValid := fmt.Sprintf("%v", compare.IsValid())

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			name,
			symbol,
			shortForm,
			isEqual,
			isValid,
		)
	}
}
