package corevalidatortests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/issetter"
	"gitlab.com/auk-go/core/tests/testwrappers/corevalidatortestwrappers"
)

var (
	arrangeArgsTwoTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.TwoAny{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	sliceValidatorTestCases = []corevalidatortestwrappers.SliceValidatorWrapper{
		{
			Case: coretestcases.CaseV1{
				Title: "Diff check against invalid comparisons, it will contain all the diff Index 0 - 2",
				ArrangeInput: []args.TwoAny{
					{
						First:  1,
						Second: byte(2),
					},
					{
						First:  1,
						Second: float64(5),
					},
					{
						First:  "1",
						Second: 1,
					},
				},
				ExpectedInput: []string{
					"Wrong expectation 1",
					"Wrong expectation 2",
					"Wrong expectation 3",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
			Validator: corevalidator.SliceValidator{
				Condition: corevalidator.DefaultTrimCoreCondition,
			ExpectedLines: []string{
					"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [0]",
					"     Actual-Processed: `\"0 : false (1, 2)\"`",
					"   Expected-Processed: `\"Wrong expectation 1\"`",
					"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [1]",
					"     Actual-Processed: `\"1 : false (1, 5)\"`",
					"   Expected-Processed: `\"Wrong expectation 2\"`",
					"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [2]",
					"     Actual-Processed: `\"2 : false (\\\"1\\\", 1)\"`",
					"   Expected-Processed: `\"Wrong expectation 3\"`",
					"",
					"=== Line-by-Line Diff (Case 0: Diff check against invalid comparisons, it will contain all the diff Index 0 - 2) ===",
					"    Actual lines: 3, Expected lines: 3",
					"  Line   0 [MISMATCH]:",
					"           actual : \"0 : false (1, 2)\"",
					"         expected : \"Wrong expectation 1\"",
					"  Line   1 [MISMATCH]:",
					"           actual : \"1 : false (1, 5)\"",
					"         expected : \"Wrong expectation 2\"",
					"  Line   2 [MISMATCH]:",
					"           actual : \"2 : false (\\\"1\\\", 1)\"",
					"         expected : \"Wrong expectation 3\"",
					"=== Total: 3 lines, 3 mismatches ===",
					"",
					"",
					"============================>",
					"0 ) Actual Received:",
					"    Diff check against invalid comparisons, it will contain all the diff Index 0 - 2",
					"============================>",
					"\"0 : false (1, 2)\",",
					"\"1 : false (1, 5)\",",
					"\"2 : false (\\\"1\\\", 1)\",",
					"============================>",
					"",
					"============================>",
					"0 )  Expected Input:",
					"     Diff check against invalid comparisons, it will contain all the diff Index 0 - 2",
					"============================>",
					"\"Wrong expectation 1\",",
					"\"Wrong expectation 2\",",
					"\"Wrong expectation 3\",",
					"============================>",
				},
			},
		},
	}

	sliceValidatorFirstErrorTestCases = []corevalidatortestwrappers.SliceValidatorWrapper{
		{
			Case: coretestcases.CaseV1{
				Title: "Diff check against invalid comparisons, it will only contain the first diff Index 0.",
				ArrangeInput: []args.TwoAny{
					{
						First:  1,
						Second: byte(2),
					},
					{
						First:  1,
						Second: float64(5),
					},
					{
						First:  "1",
						Second: 1,
					},
				},
				ExpectedInput: []string{
					"Wrong expectation 1",
					"Wrong expectation 2",
					"Wrong expectation 3",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
			Validator: corevalidator.SliceValidator{
				Condition: corevalidator.DefaultTrimCoreCondition,
			ExpectedLines: []string{
					"0 )\tExpectationLines failed: Failed match method [\"Equal\"], Index : [0]",
					"     Actual-Processed: `\"0 : false (1, 2)\"`",
					"   Expected-Processed: `\"Wrong expectation 1\"`",
					"",
					"=== Line-by-Line Diff (Case 0: Diff check against invalid comparisons, it will only contain the first diff Index 0.) ===",
					"    Actual lines: 3, Expected lines: 3",
					"  Line   0 [MISMATCH]:",
					"           actual : \"0 : false (1, 2)\"",
					"         expected : \"Wrong expectation 1\"",
					"  Line   1 [MISMATCH]:",
					"           actual : \"1 : false (1, 5)\"",
					"         expected : \"Wrong expectation 2\"",
					"  Line   2 [MISMATCH]:",
					"           actual : \"2 : false (\\\"1\\\", 1)\"",
					"         expected : \"Wrong expectation 3\"",
					"=== Total: 3 lines, 3 mismatches ===",
					"",
					"",
					"============================>",
					"0 ) Actual Received:",
					"    Diff check against invalid comparisons, it will only contain the first diff Index 0.",
					"============================>",
					"\"0 : false (1, 2)\",",
					"\"1 : false (1, 5)\",",
					"\"2 : false (\\\"1\\\", 1)\",",
					"============================>",
					"",
					"============================>",
					"0 )  Expected Input:",
					"     Diff check against invalid comparisons, it will only contain the first diff Index 0.",
					"============================>",
					"\"Wrong expectation 1\",",
					"\"Wrong expectation 2\",",
					"\"Wrong expectation 3\",",
					"============================>",
				},
			},
		},
	}
)
