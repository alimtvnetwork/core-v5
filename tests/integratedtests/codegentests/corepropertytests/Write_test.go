package corepropertytests

import (
	"testing"

	"gitlab.com/auk-go/core/codegen/coreproperty"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	writeTestCases = []coretestcases.CaseV1{
		{
			Title: "Some string given outputs double quoted string.",
			ArrangeInput: []args.OneAny{
				{
					First:  "some string",
					Expect: "some string",
				},
			},
			ExpectedInput: []string{
				`0 : some string -> "some string"`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.OneAny{}),
		},
		{
			Title: "Slice of string ,int, bytes, boolean outputs in similar fashion",
			ArrangeInput: []args.OneAny{
				{
					First: []string{
						"some val 1",
						"some val 2",
						"some val 3",
					},
				},
				{
					First: []int{
						-1,
						5,
						10,
						255,
						1500,
					},
				},
				{
					First: []uint8{
						0,
						5,
						10,
						255,
						100,
					},
				},
				{
					First: []bool{
						true,
						false,
						true,
					},
				},
			},
			ExpectedInput: []string{
				`0 : some val 1
some val 2
some val 3 -> []string {
	"some val 1", 
	"some val 2", 
	"some val 3",
}`,
				`1 : [-1,5,10,255,1500] -> []int {
	-1, 
	5, 
	10, 
	255, 
	1500,
}`,
				`2 : "AAUK/2Q=" -> []uint8 {
	0, 
	5, 
	10, 
	255, 
	100,
}`,
				`3 : [true,false,true] -> []bool {
	true, 
	false, 
	true,
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.OneAny{}),
		},
		{
			Title: "Struct with slice, etc outputs as it was given.",
			ArrangeInput: []args.OneAny{
				{
					First: &coretests.DraftType{
						SampleString1: "sample 1",
						SampleString2: "sample 2",
						SampleInteger: -59,
						Lines: []string{
							"hello 1",
							"hello 2",
						},
						RawBytes: []uint8{
							114,
							101,
							97,
							108,
							108,
							121,
							32,
							33,
						},
					},
				},
			},
			ExpectedInput: []string{
				`0 : {"SampleString1":"sample 1","SampleString2":"sample 2","SampleInteger":-59,"Lines":["hello 1","hello 2"],"RawBytes":"cmVhbGx5ICE="} -> &coretests.DraftType{
	SampleString1: "sample 1",
	SampleString2: "sample 2",
	SampleInteger: -59,
	Lines: []string {
	"hello 1", 
	"hello 2",
},
	RawBytes: []uint8 {
	114, 
	101, 
	97, 
	108, 
	108, 
	121, 
	32, 
	33,
},
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.OneAny{}),
		},
		{
			Title: "Pointer struct with slice, etc outputs as it was given.",
			ArrangeInput: []args.OneAny{
				{
					First: &coretests.DraftType{
						SampleString1: "sample 1 Ptr",
						SampleString2: "sample 2 Ptr",
						SampleInteger: -59,
						Lines: []string{
							"hello 1 Ptr",
							"hello 2 Ptr",
						},
						RawBytes: []uint8{
							114,
							101,
							97,
							108,
							108,
							121,
							32,
							33,
							32,
							80,
							116,
							114,
						},
					},
				},
			},
			ExpectedInput: []string{
				`0 : {"SampleString1":"sample 1 Ptr","SampleString2":"sample 2 Ptr","SampleInteger":-59,"Lines":["hello 1 Ptr","hello 2 Ptr"],"RawBytes":"cmVhbGx5ICEgUHRy"} -> &coretests.DraftType{
	SampleString1: "sample 1 Ptr",
	SampleString2: "sample 2 Ptr",
	SampleInteger: -59,
	Lines: []string {
	"hello 1 Ptr", 
	"hello 2 Ptr",
},
	RawBytes: []uint8 {
	114, 
	101, 
	97, 
	108, 
	108, 
	121, 
	32, 
	33, 
	32, 
	80, 
	116, 
	114,
},
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.OneAny{}),
		},
	}
)

func Test_Write_Verification(t *testing.T) {
	for caseIndex, testCase := range writeTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.OneAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(40)
		actFuncWrite := coreproperty.
			Writer.
			Write

		// Act
		for i, input := range inputs {
			inArgInterface := input.First.(any)

			allInArgsCompiled := converters.AnyTo.SmartStringsOf(
				inArgInterface,
			)

			result := actFuncWrite(
				inArgInterface,
			)

			allOutArgsCompiled := converters.AnyTo.SmartStringsOf(
				result,
			)

			actualSlice.AppendFmt(
				"%d : %s -> %s",
				i,
				allInArgsCompiled,
				allOutArgsCompiled,
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
