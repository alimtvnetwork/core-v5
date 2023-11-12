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

	arrangeArgsTwoTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]coretests.ArgTwo{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	arrangeInterfaceOfInterfaceArrayTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([][]interface{}{}),
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
				Title: "Only true if defined (not null) ones will be true.",
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
				Title: "Only true if both defined (not null) ones will be true.",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  nil,
						Second: nil,
					},
					{
						First:  nil,
						Second: someNull,
					},
					{
						First:  someNull,
						Second: nil,
					},
					{
						First:  someNull,
						Second: someNull,
					},
					{
						First:  1,
						Second: someNull,
					},
					{
						First:  someNull,
						Second: 2,
					},
					{
						First:  1,
						Second: nil,
					},
					{
						First:  nil,
						Second: 2,
					},
					{
						First:  1,
						Second: 2,
					},
					{
						First:  &coretests.ArgTwo{},
						Second: 2,
					},
					{
						First:  &coretests.ArgTwo{},
						Second: coretests.ArgTwo{},
					},
				},
				ExpectedInput: []string{
					"0 : false (<nil>, <nil>)",
					"1 : false (<nil>, *coretests.ArgTwo)",
					"2 : false (*coretests.ArgTwo, <nil>)",
					"3 : false (*coretests.ArgTwo, *coretests.ArgTwo)",
					"4 : false (int, *coretests.ArgTwo)",
					"5 : false (*coretests.ArgTwo, int)",
					"6 : false (int, <nil>)",
					"7 : false (<nil>, int)",
					"8 : true (int, int)",
					"9 : true (*coretests.ArgTwo, int)",
					"10 : true (*coretests.ArgTwo, coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	nullBothTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both are null (not defined). " +
					"Kind of inverse of any defined.",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  nil,
						Second: nil,
					},
					{
						First:  nil,
						Second: someNull,
					},
					{
						First:  someNull,
						Second: nil,
					},
					{
						First:  someNull,
						Second: someNull,
					},
					{
						First:  1,
						Second: someNull,
					},
					{
						First:  someNull,
						Second: 2,
					},
					{
						First:  1,
						Second: nil,
					},
					{
						First:  nil,
						Second: 2,
					},
					{
						First:  1,
						Second: 2,
					},
					{
						First:  &coretests.ArgTwo{},
						Second: 2,
					},
					{
						First:  &coretests.ArgTwo{},
						Second: coretests.ArgTwo{},
					},
				},
				ExpectedInput: []string{
					"0 : true (<nil>, <nil>)",
					"1 : true (<nil>, *coretests.ArgTwo)",
					"2 : true (*coretests.ArgTwo, <nil>)",
					"3 : true (*coretests.ArgTwo, *coretests.ArgTwo)",
					"4 : false (int, *coretests.ArgTwo)",
					"5 : false (*coretests.ArgTwo, int)",
					"6 : false (int, <nil>)",
					"7 : false (<nil>, int)",
					"8 : false (int, int)",
					"9 : false (*coretests.ArgTwo, int)",
					"10 : false (*coretests.ArgTwo, coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	definedAllOfTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if all are defined (not null) - DefinedAllOf.",
				ArrangeInput: [][]interface{}{
					{
						1,
						2,
						"some string",
					},
					{
						1,
						nil,
						"some string",
					},
					{
						1,
						3,
						someNull,
					},
					{
						"",
						3,
						555.3,
					},
				},
				ExpectedInput: []string{
					"0 : true (int, int, string)",
					"1 : false (int, <nil>, string)",
					"2 : false (int, int, *coretests.ArgTwo)",
					"3 : true (string, int, float64)",
				},
				VerifyTypeOf: arrangeInterfaceOfInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	definedAnyOfTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if any is defined (not null) - DefinedAnyOf.",
				ArrangeInput: [][]interface{}{
					{
						1,
						2,
						"some string",
					},
					{
						1,
						nil,
						"some string",
					},
					{
						1,
						3,
						someNull,
					},
					{
						"",
						3,
						555.3,
					},
					{
						nil,
						someNull,
						someNull,
					},
				},
				ExpectedInput: []string{
					"0 : true (int, int, string)",
					"1 : true (int, <nil>, string)",
					"2 : true (int, int, *coretests.ArgTwo)",
					"3 : true (string, int, float64)",
					"4 : false (<nil>, *coretests.ArgTwo, *coretests.ArgTwo)",
				},
				VerifyTypeOf: arrangeInterfaceOfInterfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	jsonEqualTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. Here all are null comparison.",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  nil,
						Second: nil,
					},
					{
						First:  nil,
						Second: someNull,
					},
					{
						First:  someNull,
						Second: nil,
					},
					{
						First:  someNull,
						Second: someNull,
					},
				},
				ExpectedInput: []string{
					"0 : true (null, null)",
					"1 : true (null, null)",
					"2 : true (null, null)",
					"3 : true (null, null)",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting not equal comparing with null with non null.",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  1,
						Second: someNull,
					},
					{
						First:  someNull,
						Second: 2,
					},
					{
						First:  1,
						Second: nil,
					},
					{
						First:  nil,
						Second: 2,
					},
				},
				ExpectedInput: []string{
					"0 : false (1, null)",
					"1 : false (null, 2)",
					"2 : false (1, null)",
					"3 : false (null, 2)",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting not equal comparing with same type different values.",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  1,
						Second: 2,
					},
					{
						First:  1,
						Second: 5,
					},
					{
						First:  "some alim",
						Second: "some not alim",
					},
				},
				ExpectedInput: []string{
					"0 : false (1, 2)",
					"1 : false (1, 5)",
					"2 : false (\"some alim\", \"some not alim\")",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting equal comparing with any value as long as both are equal in terms of bytes.",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  1,
						Second: "1",
					},
					{
						First:  float32(20),
						Second: 20,
					},
					{
						First:  -11,
						Second: float32(-11),
					},
					{
						First:  "alim is equal",
						Second: "alim is equal",
					},
					{
						First: &coretests.ArgTwo{
							First:  "1",
							Second: "alim",
						},
						Second: coretests.ArgTwo{
							First:  "1",
							Second: "alim",
						},
					},
					{
						First:  &coretests.ArgTwo{},
						Second: coretests.ArgTwo{},
					},
				},
				ExpectedInput: []string{
					"0 : false (1, \"1\")",
					"1 : true (20, 20)",
					"2 : true (-11, -11)",
					"3 : true (\"alim is equal\", \"alim is equal\")",
					"4 : true ({\"First\":\"1\",\"Second\":\"alim\"}, {\"First\":\"1\",\"Second\":\"alim\"})",
					"5 : true (, )",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting not equal comparing with string with integer with same value.",
				ArrangeInput: []coretests.ArgTwo{
					{
						First:  1,
						Second: "1",
					},
					{
						First:  "55",
						Second: 55,
					},
					{
						First:  -1,
						Second: "-1",
					},
				},
				ExpectedInput: []string{
					"0 : false (1, \"1\")",
					"1 : false (\"55\", 55)",
					"2 : false (-1, \"-1\")",
				},
				VerifyTypeOf: arrangeArgsTwoTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
