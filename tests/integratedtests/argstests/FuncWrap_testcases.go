package argstests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.OneFunc{}),
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
				"----------------------",
				"3 )  When: some title, or case when,",
				"   Actual: actual rec,",
				" Expected: expected item",
			},
			VerifyTypeOf: commonType,
		},
	}
)
