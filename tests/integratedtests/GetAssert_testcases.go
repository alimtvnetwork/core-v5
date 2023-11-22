package integratedtests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	arrangeTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	twoArgsTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.Two{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	oneFuncTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.OneFunc{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	interfaceArrayTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([][]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	quickTestCases = []coretestcases.CaseV1{
		{
			Title:         "",
			ArrangeInput:  nil,
			ExpectedInput: nil,
			VerifyTypeOf:  nil,
		},
	}
)
