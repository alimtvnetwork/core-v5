package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

func Test_Reflection_Types_Verification(t *testing.T) {
	toStringFunc := convertinternal.AnyTo.SmartString
	for caseIndex, testCase := range reflectionTypesTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.OneFuncAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			first := input.First
			isFunc := testCase.FirstParam()
			checkerFunc := convertFuncType(input.WorkFunc)
			funcName := input.GetFuncName()
			value := conditional.IfString(
				isFunc == "isFunc",
				funcName,
				toStringFunc(first),
			)

			actualSlice.AppendFmt(
				booleanTypeStringStringFormat,
				i,
				checkerFunc(first),
				first,
				funcName,
				value,
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
