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
			Title: "When nil / null is given, nil returned as is hardcoded.",
			ArrangeInput: []args.One{
				{
					First: nil,
				},
			},
			ExpectedInput: []string{
				`0 : [{}] -> []args.One {
	args.One{
	First: nil,
	Expect: nil,
},
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
		{
			Title: "Some string given outputs double quoted string.",
			ArrangeInput: []args.One{
				{
					First:  "some string",
					Expect: "some string",
				},
			},
			ExpectedInput: []string{
				`0 : [{"First":"some string","Expect":"some string"}] -> []args.One {
	args.One{
	First: "some string",
	Expect: "some string",
},
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
		{
			Title: "Slice of string ,int, bytes, boolean outputs in similar fashion",
			ArrangeInput: []args.One{
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
				`0 : [{"First":["some val 1","some val 2","some val 3"]},{"First":[-1,5,10,255,1500]},{"First":"AAUK/2Q="},{"First":[true,false,true]}] -> []args.One {
	args.One{
	First: []string {
	"some val 1", 
	"some val 2", 
	"some val 3",
},
	Expect: nil,
}, 
	args.One{
	First: []int {
	-1, 
	5, 
	10, 
	255, 
	1500,
},
	Expect: nil,
}, 
	args.One{
	First: []uint8 {
	0, 
	5, 
	10, 
	255, 
	100,
},
	Expect: nil,
}, 
	args.One{
	First: []bool {
	true, 
	false, 
	true,
},
	Expect: nil,
},
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
		{
			Title: "Struct with slice, etc outputs as it was given.",
			ArrangeInput: []args.One{
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
				`0 : [{"First":{"SampleString1":"sample 1","SampleString2":"sample 2","SampleInteger":-59,"Lines":["hello 1","hello 2"],"RawBytes":"cmVhbGx5ICE="}}] -> []args.One {
	args.One{
	First: &coretests.DraftType{
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
},
	Expect: nil,
},
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
		{
			Title: "Pointer struct with slice, etc outputs as it was given.",
			ArrangeInput: []args.One{
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
				`0 : [{"First":{"SampleString1":"sample 1 Ptr","SampleString2":"sample 2 Ptr","SampleInteger":-59,"Lines":["hello 1 Ptr","hello 2 Ptr"],"RawBytes":"cmVhbGx5ICEgUHRy"}}] -> []args.One {
	args.One{
	First: &coretests.DraftType{
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
},
	Expect: nil,
},
}`,
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.One{}),
		},
	}
)

func Test_Write_Verification(t *testing.T) {
	for caseIndex, testCase := range writeTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.One)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(40)
		actFuncWrite := coreproperty.
			Writer.
			Write

		// Act
		for i, input := range inputs {
			inArgInterface := input.First.(interface{})

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
				"%d : %s -> %s | %s",
				i,
				allInArgsCompiled,
				allOutArgsCompiled,
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
