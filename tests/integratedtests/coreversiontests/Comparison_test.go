package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coreversion"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func Test_Comparison_Verification(t *testing.T) {
	for caseIndex, testCase := range comparisonStringTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coretests.LeftRightExpect)
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
				&right)

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
		finalCase.AssertEqual(
			t,
			caseIndex,
			finalActLines...)
	}
}

func Test_Method_Verification(t *testing.T) {
	for caseIndex, testCase := range versionMethodsVerificationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coretests.ArgThree)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for index, input := range inputs {
			f := input.First.(int)
			s := input.Second.(int)
			theFunc := input.Third.(func(major, x int) bool)
			funcName := reflectinternal.GetFuncName(theFunc)

			isMatch := theFunc(f, s)

			actualSlice.AppendFmt(
				comparisonMethodFmt,
				index,
				funcName,
				f,
				s,
				isMatch,
			)
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
