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
					"0 : true (value: <nil>, type: <nil>)",
					"1 : false (value: &{<nil> <nil>}, type: *coretests.ArgTwo)",
					"2 : true (value: <nil>, type: *coretests.ArgTwo)",
					"3 : false (value: 1, type: int)",
					"4 : false (value: 2, type: int)",
					"5 : false (value: {<nil> <nil>}, type: coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	allNullTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if all cases are null, will return false.",
				ArrangeInput: []interface{}{
					nil,
					&coretests.ArgTwo{},
					someNull,
					1,
					2,
					coretests.ArgTwo{},
				},
				ExpectedInput: []string{
					"0 : false (<nil>, *coretests.ArgTwo, *coretests.ArgTwo, int, int, coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if all cases are null, will return true.",
				ArrangeInput: []interface{}{
					nil,
					someNull,
					someNull,
					nil,
				},
				ExpectedInput: []string{
					"1 : true (<nil>, *coretests.ArgTwo, *coretests.ArgTwo, <nil>)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	anyNullTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if any case is null, it will result true, because one is nil.",
				ArrangeInput: []interface{}{
					&coretests.ArgTwo{},
					1,
					2,
					coretests.ArgTwo{},
					someNull,
				},
				ExpectedInput: []string{
					"0 : true (*coretests.ArgTwo, int, int, coretests.ArgTwo, *coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if any case is null, it will result true, because one is nil.",
				ArrangeInput: []interface{}{
					nil,
					someNull,
					someNull,
					nil,
				},
				ExpectedInput: []string{
					"1 : true (<nil>, *coretests.ArgTwo, *coretests.ArgTwo, <nil>)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if any case is null, it will result false, because none is nil.",
				ArrangeInput: []interface{}{
					1,
					2,
					"",
					[]string{},
				},
				ExpectedInput: []string{
					"2 : false (int, int, string, []string)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	definedTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only defined (not null) ones will be true.",
				ArrangeInput: []interface{}{
					&coretests.ArgTwo{},
					1,
					nil,
					coretests.ArgTwo{},
					someNull,
				},
				ExpectedInput: []string{
					"0 : true (value: &{<nil> <nil>}, type: *coretests.ArgTwo)",
					"1 : true (value: 1, type: int)",
					"2 : false (value: <nil>, type: <nil>)",
					"3 : true (value: {<nil> <nil>}, type: coretests.ArgTwo)",
					"4 : false (value: <nil>, type: *coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	bothDefinedTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only defined (not null) ones will be true.",
				ArrangeInput: []interface{}{
					&coretests.ArgTwo{},
					1,
					nil,
					coretests.ArgTwo{},
					someNull,
				},
				ExpectedInput: []string{
					"0 : true (value: &{<nil> <nil>}, type: *coretests.ArgTwo)",
					"1 : true (value: 1, type: int)",
					"2 : false (value: <nil>, type: <nil>)",
					"3 : true (value: {<nil> <nil>}, type: coretests.ArgTwo)",
					"4 : false (value: <nil>, type: *coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
