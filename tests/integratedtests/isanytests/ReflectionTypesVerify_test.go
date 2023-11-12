package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corefuncs"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

func Test_Reflection_Types_Verification(t *testing.T) {
	for caseIndex, testCase := range reflectionTypesTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coretests.ArgTwo)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			first := input.First
			checkerFunc := convertFuncType(input.Second)
			funcName := corefuncs.GetFuncName(input.Second)

			actualSlice.AppendFmt(
				defaultCaseIndexBoolStringStringFmt,
				i,
				checkerFunc(first),
				funcName,
				convertinteranl.AnyTo.SmartString(first))
		}

		finalActual := actualSlice.Strings()
		finalTestCase := coretestcases.
			TestCaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual...)
	}
}
