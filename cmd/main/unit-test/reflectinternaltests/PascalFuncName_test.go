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
				"0 : someName -> SomeName",
				"1 : someName -> SomeName",
				"2 : someName 2 -> SomeName 2",
				"3 : someName 2 -> SomeName 2",
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
		for _, input := range inputs {
			inArgString := input.First.(string)

			result := actFuncPascalFuncName(inArgString)

			actualSlice.AppendFmt(
				"%d : %s -> %s",
				caseIndex,
				inArgString,
				result,
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
