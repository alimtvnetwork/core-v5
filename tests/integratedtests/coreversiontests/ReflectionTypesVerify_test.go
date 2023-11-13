package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corefuncs"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

func Test_Reflection_Types_Verification(t *testing.T) {
	toStringFunc := convertinteranl.AnyTo.SmartString
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
			isFunc := testCase.FirstParam()
			checkerFunc := convertFuncType(input.Second)
			funcName := corefuncs.GetFuncName(input.Second)
			value := conditional.String(
				isFunc == "isFunc",
				funcName,
				toStringFunc(first))

			actualSlice.AppendFmt(
				booleanTypeStringStringFormat,
				i,
				checkerFunc(first),
				first,
				funcName,
				value)
		}

		finalActual := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActual...)
	}
}
