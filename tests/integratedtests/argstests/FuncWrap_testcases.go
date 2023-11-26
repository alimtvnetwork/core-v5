package argstests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.ThreeFunc{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	funWrapCreationTestCases = []coretestcases.CaseV1{
		{
			Title: "someFunctionV1 => Calls dynamically with valid params, outputs as it should.",
			ArrangeInput: args.ThreeFunc{
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
			Title: "someFunctionV1 => Calls dynamically with less param (null), outputs error args count mismatch.",
			ArrangeInput: args.ThreeFunc{
				First:    "f1",
				Second:   "f2",
				Third:    nil,
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"error : ",
				"    someFunctionV1 [Func] =>",
				"      arguments count doesn't match for - count:",
				"        expected : 3",
				"        given    : 2",
				"      expected types listed :",
				"        - string",
				"        - string",
				"        - string",
				"      actual given types list :",
				"        - 0. string [value: f1]",
				"        - 1. string [value: f2]",
			},
			VerifyTypeOf: commonType,
		},
		{
			Title: "someFunctionV1 => Calls dynamically with mismatch datatype for arg 2nd, it expects string but given int, outputs error",
			ArrangeInput: args.ThreeFunc{
				First:    "f1",
				Second:   1,
				Third:    "f3",
				WorkFunc: someFunctionV1,
			},
			ExpectedInput: []string{
				"error : ",
				"    someFunctionV1 =>",
				"        - Index {1}, 2nd args : Expected Type (string) != (int) Given Type",
			},
			VerifyTypeOf: commonType,
		},
	}
)
