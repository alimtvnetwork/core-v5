package reflectinternaltests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	pascalFuncNameTestCases = []coretestcases.CaseV1{
		{
			Title: "Some",
			ArrangeInput: []args.One{
				{
					First:  "someName",
					Expect: "some expect",
				},
				{
					First: "someName 2",
				},
			},
			ExpectedInput: []string{
				"0 : someName -> SomeName | nil",
				"1 : someName 2 -> SomeName 2 | nil",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
	}
)

func Test_PascalFuncName_Verification(t *testing.T) {
	for caseIndex, testCase := range pascalFuncNameTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.One)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(40)
		actFuncPascalFuncName := reflectinternal.
			GetFunc.
			PascalFuncName

		// Act
		for i, input := range inputs {
			inArgString := input.First.(string)

			result := actFuncPascalFuncName(inArgString)

			actualSlice.AppendFmt(
				"%d : %s -> %s | %s",
				i,
				inArgString,
				result,
				input.Expect,
			)
		}

		finalActLines := actualSlice.Strings()
		actualSlice.Dispose()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
