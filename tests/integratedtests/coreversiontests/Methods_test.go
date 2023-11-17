package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func Test_TwoParams_Method_Verification(t *testing.T) {
	for caseIndex, testCase := range versionTwoParamsVerificationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coretests.ArgFour)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		actualSlice.AppendFmt(
			"Testing for -> Version(%s)",
			someVersionV5,
		)

		// Act
		for index, input := range inputs {
			// "v5.8.1.5"
			f := input.First.(int)
			s := input.Second.(int)
			third := input.Third.(bool)
			theFunc := input.Fourth.(func(major, x int) bool)
			funcName := reflectinternal.GetFuncName(theFunc)

			isMatch := theFunc(f, s)

			actualSlice.AppendFmt(
				comparisonMethodFmt,
				index,
				funcName,
				f,
				s,
				isMatch,
				third,
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

func Test_ThreeParams_Method_Verification(t *testing.T) {
	for caseIndex, testCase := range versionThreeParamsVerificationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coretests.ArgFive)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		actualSlice.AppendFmt(
			"Testing for -> Version(%s)",
			someVersionV5,
		)

		// Act
		for index, input := range inputs {
			// "v5.8.1.5"
			f := input.First.(int)
			s := input.Second.(int)
			third := input.Third.(int)
			fourth := input.Fourth.(bool)
			theFunc := input.Fifth.(func(major, x, y int) bool)
			funcName := reflectinternal.GetFuncName(theFunc)

			isMatch := theFunc(f, s, third)

			actualSlice.AppendFmt(
				comparisonMethod3Fmt,
				index,
				funcName,
				f,
				s,
				third,
				isMatch,
				fourth,
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
