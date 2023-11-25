package argstests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	commonType = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(args.Map{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	funWrapCreationTestCases = []coretestcases.CaseV1{
		{
			Title: "Quick output as gherkins format",
			ArrangeInput: args.OneFunc{
				First: someFunctionV1,
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
