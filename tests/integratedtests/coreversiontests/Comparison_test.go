package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coreversion"
)

func Test_Comparison_Verification(t *testing.T) {
	for caseIndex, testCase := range comparisonStringTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coretests.ArgThree)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		creatorFunc := coreversion.New.Default

		// Act
		for i, input := range inputs {
			l := input.First.(string)
			r := input.Second.(string)
			expectation := input.Third.(corecomparator.Compare)

			left := creatorFunc(l)
			right := creatorFunc(r)
			isMatch := left.IsExpectedComparison(
				expectation,
				right)

			actualSlice.AppendFmt(
				comparisonFmt,
				i,
				expectation,
				l,
				r,
				isMatch)
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.AssertEqual(
			t,
			caseIndex,
			finalActLines...)
	}
}
