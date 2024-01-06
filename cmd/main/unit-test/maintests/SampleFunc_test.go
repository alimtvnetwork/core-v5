package maintests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	sampleFuncTestCases = []coretestcases.CaseV1{
		{
			Title: "Some",
			ArrangeInput: []args.One{
				,
			},
			ExpectedInput: []string{
				,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
	}
)

func Test_SampleFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range sampleFuncTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.One)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(40)
		actFuncSampleFunc := main.
			unitTestGenerator.
			SampleFunc

		// Act
		for i, input := range inputs {
			inArgInt0 := input.First.(int)
			inArgString1 := input.Second.(string)
			inArgString2 := input.Third.(string)
			inArg * main.AlimStruct3 := input.Fourth.(*main.AlimStruct)
			inArgMain.AlimStruct4 := input.Fifth.(main.AlimStruct)

			result1, result2, result3 := actFuncSampleFunc(
				inArgInt0,
				inArgString1,
				inArgString2,
				inArg*main.AlimStruct3,
				inArgMain.AlimStruct4,
			)

			actualSlice.AppendFmt(
				"%d : %s -> %s | %s",
				i,
				inArgInt0,
				inArgString1,
				inArgString2,
				inArg*main.AlimStruct3,
				inArgMain.AlimStruct4,
				result1,
				result2,
				result3,
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
