package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreversion"
)

func Test_Comparison_Verification(t *testing.T) {
	for caseIndex, testCase := range comparisonStringTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.LeftRightAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		creatorFunc := coreversion.New.Default

		// Act
		for index, input := range inputs {
			l := input.Left.(string)
			r := input.Right.(string)
			expectation := input.Expect.(corecomparator.Compare)

			left := creatorFunc(l)
			right := creatorFunc(r)
			isMatch := left.IsExpectedComparison(
				expectation,
				&right,
			)

			actualSlice.AppendFmt(
				comparisonFmt,
				index,
				l,
				expectation.OperatorSymbol(),
				r,
				expectation,
				isMatch,
			)
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
