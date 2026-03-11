package coreflecttests

import (
	"reflect"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var (
	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.ThreeFuncAny{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	funWrapCreationTestCases = []coretestcases.CaseV1{
		{
			Title: "FuncWrap returns correct output -- someFunctionV1 with valid params",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				Third:    "f3",
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"someFunctionV1 => called with (f1, f2, f3) - some new stuff",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "FuncWrap returns args count mismatch error -- someFunctionV1 with nil third param",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				Third:    nil,
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"error : ",
				"  someFunctionV1 [Func] =>",
				"    arguments count doesn't match for - count:",
				"      expected : 3",
				"      given    : 2",
				"    expected types listed :",
				"      - string",
				"      - string",
				"      - string",
				"    actual given types list :",
				"      - 0. string [value: f1]",
				"      - 1. string [value: f2]",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "FuncWrap returns type mismatch error -- someFunctionV1 with int instead of string at arg 2",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   1,
				Third:    "f3",
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"error : ",
				"  someFunctionV1 =>",
				"      - Index {1}, 2nd args : Expected Type (string) != (int) Given Type",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "FuncWrap returns invalid error -- nil work func given",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   1,
				Third:    "f3",
				WorkFunc: nil,
			},
			ExpectedInput: []string{
				"error : ",
				"  func-wrap is invalid:",
				"      given type: <nil>",
				"      name: ",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "FuncWrap returns invalid error -- int given as work func",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   1,
				Third:    "f3",
				WorkFunc: 1,
			},
			ExpectedInput: []string{
				"error : ",
				"  func-wrap is invalid:",
				"      given type: int",
				"      name: ",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "FuncWrap returns string and error output -- someFunctionV2 with valid params",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				WorkFunc: someFunctionV2,
			},
			ExpectedInput: []string{
				"someFunctionV2 => called with (f1, f2) - (string, error)",
				"some err v2",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "FuncWrap returns int, string, error output -- someFunctionV3 with valid params",
			ArrangeInput: args.ThreeFuncAny{
				First:    "f1",
				Second:   "f2",
				WorkFunc: someFunctionV3,
			},
			ExpectedInput: []string{
				"5",
				"someFunctionV3 => called with (f1, f2) - (int, string, error)",
				"some err of v3",
			},
			VerifyTypeOf: commonType,
		},
	}
)
