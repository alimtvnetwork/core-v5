package isanytests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/issetter"
)

var (
	arrangeTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	someNull *coretests.ArgTwo = nil

	nullTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "all nulls will be returned as null, don't panic.",
				ArrangeInput: []interface{}{
					nil,
					&coretests.ArgTwo{},
					someNull,
					1,
					2,
					coretests.ArgTwo{},
				},
				ExpectedInput: []string{
					"0 : true (<nil>)",
					"1 : false (*coretests.ArgTwo)",
					"2 : true (*coretests.ArgTwo)",
					"3 : false (int)",
					"4 : false (int)",
					"5 : false (coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	allNullTestCases2 = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "all nulls will be returned as null, don't panic.",
				ArrangeInput: []interface{}{
					nil,
					&coretests.ArgTwo{},
					someNull,
					1,
					2,
					coretests.ArgTwo{},
				},
				ExpectedInput: []string{
					"0 : Equal (<nil>, <nil>)",
					"1 : NotEqual (int, <nil>)",
					"2 : Inconclusive (int, int)",
					"3 : NotEqual (*coretests.DraftType, <nil>)",
					"4 : NotEqual (<nil>, *coretests.DraftType)",
					"5 : Inconclusive (*coretests.DraftType, *coretests.DraftType)",
					"6 : Equal (*coretests.VerifyTypeOf, *coretests.VerifyTypeOf)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
