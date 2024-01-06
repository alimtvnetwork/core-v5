package samplefunctests

import (
	"testing"

	"gitlab.com/auk-go/core/cmd/main/samplefunc"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	myFuncTestCases = []coretestcases.CaseV1{
		{
			Title: "Some",
			ArrangeInput: []args.Five{
				{
					First: 1,
					Second: "alim 1",
					Third: "alim 2",
					Fourth: &samplefunc.AlimStruct{
						First: "alim 1", LeftRight: args.LeftRight{
							Left: "l", Right: "r", Expect: "e", toSlice: (*[]interface{})(nil),
							toString: corestr.SimpleStringOnce{value: "", isInitialize: false},
						}, Draft: coretests.DraftType{
							SampleString1: "alim 1", SampleString2: "", SampleInteger: 0, Lines: []string(nil),
							RawBytes: []uint8(nil), f1String: "", f2Integer: 0,
						},
					},
					Fifth: []samplefunc.AlimStruct{
						{
							First: "alim 2", LeftRight: args.LeftRight{
							Left: "a2-l", Right: "a2-r", Expect: "a2-e", toSlice: (*[]interface{})(nil),
							toString: corestr.SimpleStringOnce{value: "", isInitialize: false},
						}, Draft: coretests.DraftType{
							SampleString1: "alim 2", SampleString2: "", SampleInteger: 0, Lines: []string(nil),
							RawBytes: []uint8(nil), f1String: "", f2Integer: 0,
													},
												},
					},
					Expect: "some expect",
				},
			},
			ExpectedInput: []string{
				"0 : 1
				alim, 1
				alim 2
			{
				"First":"alim 1", "LeftRight":{
				"Left":"l", "Right":"r", "Expect":"e"
			}, "Draft":{
				"SampleString1":"alim 1", "SampleString2":"", "SampleInteger":0, "Lines":null, "RawBytes":null
			}
			}
			[{
				"First":"alim 2", "LeftRight":{
				"Left":"a2-l", "Right":"a2-r", "Expect":"a2-e"
			}, "Draft":{
				"SampleString1":"alim 2", "SampleString2":"", "SampleInteger":0, "Lines":null, "RawBytes":null
			}
			}] -> alim 1 alim 2-> Processed
				2
			{
				"First":"someName - alim 1alim 2", "LeftRight":{
				"Left":"l", "Right":"r", "Expect":"e"
			}, "Draft":{
				"SampleString1":"alim 2", "SampleString2":"", "SampleInteger":0, "Lines":null, "RawBytes":null
			}
			} | some expect",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.Five{}),
		},
	}
)

func Test_MyFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range myFuncTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.Five)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(40)
		actFuncMyFunc := samplefunc.
			MyFunc

		// Act
		for i, input := range inputs {
			inArgInt0 := input.First.(int)
			inArgString1 := input.Second.(string)
			inArgString2 := input.Third.(string)
			inArg * samplefunc.AlimStruct3 := input.Fourth.(*samplefunc.AlimStruct)
			inArg[]
			samplefunc.AlimStruct4 := input.Fifth.([]samplefunc.AlimStruct)

			result1, result2, result3 := actFuncMyFunc(
				inArgInt0,
				inArgString1,
				inArgString2,
				inArg*samplefunc.AlimStruct3,
				inArg[]
			samplefunc.AlimStruct4)

			actualSlice.AppendFmt(
				"%d : %s -> %s | %s",
				i,
				inArgInt0,
				inArgString1,
				inArgString2,
				inArg*samplefunc.AlimStruct3,
				inArg[]
			samplefunc.AlimStruct4,
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
