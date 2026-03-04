package isanytests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/isany"
	"gitlab.com/auk-go/core/issetter"
)

var (
	arrangeTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]any{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	twoArgsTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.TwoAny{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}
	oneFuncTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.OneFuncAny{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	interfaceArrayTypeVerification = &coretests.VerifyTypeOf{

		ArrangeInput:  reflect.TypeOf([][]any{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	someNull *args.TwoAny = nil

	nullTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "null tests - all nulls will be returned as null, don't panic.",
			ArrangeInput: []any{
					nil,
					&args.TwoAny{},
					someNull,
					1,
					2,
					args.TwoAny{},
				},
			ExpectedInput: []string{
				"0 : true (value: <nil>, type: <nil>)",
				"1 : false (value: Two {  }, type: *args.Two[interface {},interface {}])",
				"2 : true (value: <nil>, type: *args.Two[interface {},interface {}])",
				"3 : false (value: 1, type: int)",
				"4 : false (value: 2, type: int)",
				"5 : false (value: {<nil> <nil> <nil> [] false { false}}, type: args.Two[interface {},interface {}])",
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
			ArrangeInput: []any{
					nil,
					&args.TwoAny{},
					someNull,
					1,
					2,
					args.TwoAny{},
				},
			ExpectedInput: []string{
				"0 : false (<nil>, *args.Two[interface {},interface {}], *args.Two[interface {},interface {}], int, int, args.Two[interface {},interface {}])",
			},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "all are null, will return true.",
			ArrangeInput: []any{
					nil,
					someNull,
					someNull,
					nil,
				},
			ExpectedInput: []string{
				"1 : true (<nil>, *args.Two[interface {},interface {}], *args.Two[interface {},interface {}], <nil>)",
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
			ArrangeInput: []any{
					&args.TwoAny{},
					1,
					2,
					args.TwoAny{},
					someNull,
				},
			ExpectedInput: []string{
				"0 : true (*args.Two[interface {},interface {}], int, int, args.Two[interface {},interface {}], *args.Two[interface {},interface {}])",
			},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if any case is null, it will result true, because one is nil.",
			ArrangeInput: []any{
					nil,
					someNull,
					someNull,
					nil,
				},
			ExpectedInput: []string{
				"1 : true (<nil>, *args.Two[interface {},interface {}], *args.Two[interface {},interface {}], <nil>)",
			},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if any case is null, it will result false, because none is nil.",
			ArrangeInput: []any{
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
				Title: "defined items test cases - only true if defined (not null) ones will be true.",
				ArrangeInput: []any{
					&args.TwoAny{},
					1,
					nil,
					args.TwoAny{},
					someNull,
				},
			ExpectedInput: []string{
				"0 : true (value: Two {  }, type: *args.Two[interface {},interface {}])",
				"1 : true (value: 1, type: int)",
				"2 : false (value: <nil>, type: <nil>)",
				"3 : true (value: {<nil> <nil> <nil> [] false { false}}, type: args.Two[interface {},interface {}])",
				"4 : false (value: <nil>, type: *args.Two[interface {},interface {}])",
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
				ArrangeInput: []args.TwoAny{
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
						First:  &args.TwoAny{},
						Second: 2,
					},
					{
						First:  &args.TwoAny{},
						Second: args.TwoAny{},
					},
				},
			ExpectedInput: []string{
				"0 : false (<nil>, <nil>)",
				"1 : false (<nil>, *args.Two[interface {},interface {}])",
				"2 : false (*args.Two[interface {},interface {}], <nil>)",
				"3 : false (*args.Two[interface {},interface {}], *args.Two[interface {},interface {}])",
				"4 : false (int, *args.Two[interface {},interface {}])",
				"5 : false (*args.Two[interface {},interface {}], int)",
				"6 : false (int, <nil>)",
				"7 : false (<nil>, int)",
				"8 : true (int, int)",
				"9 : true (*args.Two[interface {},interface {}], int)",
				"10 : true (*args.Two[interface {},interface {}], args.Two[interface {},interface {}])",
			},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	nullBothTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both are null (not defined). " +
					"Kind of inverse of any defined.",
				ArrangeInput: []args.TwoAny{
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
						First:  &args.TwoAny{},
						Second: 2,
					},
					{
						First:  &args.TwoAny{},
						Second: args.TwoAny{},
					},
				},
			ExpectedInput: []string{
				"0 : true (<nil>, <nil>)",
				"1 : true (<nil>, *args.Two[interface {},interface {}])",
				"2 : true (*args.Two[interface {},interface {}], <nil>)",
				"3 : true (*args.Two[interface {},interface {}], *args.Two[interface {},interface {}])",
				"4 : false (int, *args.Two[interface {},interface {}])",
				"5 : false (*args.Two[interface {},interface {}], int)",
				"6 : false (int, <nil>)",
				"7 : false (<nil>, int)",
				"8 : false (int, int)",
				"9 : false (*args.Two[interface {},interface {}], int)",
				"10 : false (*args.Two[interface {},interface {}], args.Two[interface {},interface {}])",
			},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	definedAllOfTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if all are defined (not null) - DefinedAllOf.",
			ArrangeInput: [][]any{
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
				"2 : false (int, int, *args.Two[interface {},interface {}])",
				"3 : true (string, int, float64)",
			},
				VerifyTypeOf: interfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	definedAnyOfTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if any is defined (not null) - DefinedAnyOf.",
				ArrangeInput: [][]any{
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
				"2 : true (int, int, *args.Two[interface {},interface {}])",
				"3 : true (string, int, float64)",
				"4 : false (<nil>, *args.Two[interface {},interface {}], *args.Two[interface {},interface {}])",
			},
				VerifyTypeOf: interfaceArrayTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	jsonEqualTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. Here all are null comparison.",
				ArrangeInput: []args.TwoAny{
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
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting not equal comparing with null with non null.",
				ArrangeInput: []args.TwoAny{
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
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting not equal comparing with same type different values.",
				ArrangeInput: []args.TwoAny{
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
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting equal comparing with any value as long as both are equal in terms of bytes.",
				ArrangeInput: []args.TwoAny{
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
						First: &args.TwoAny{
							First:  "1",
							Second: "alim",
						},
						Second: args.TwoAny{
							First:  "1",
							Second: "alim",
						},
					},
					{
						First:  &args.TwoAny{},
						Second: args.TwoAny{},
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
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Only true if both equal in terms of json bytes. " +
					"Expecting not equal comparing with string with integer with same value.",
				ArrangeInput: []args.TwoAny{
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
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	conclusiveTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Expect all not equal and inconclusive because value same but data types are different.",
				ArrangeInput: []args.TwoAny{
					{
						First:  nil,
						Second: nil,
					},
					{
						First:  someNull,
						Second: someNull,
					},
				},
			ExpectedInput: []string{
				"0 - Equal : true - Conclusive ('<nil>', '<nil>')",
				"1 - Equal : true - Conclusive ('*args.Two[interface {},interface {}]', '*args.Two[interface {},interface {}]')",
			},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Expect all not equal and inconclusive because value same but data types are different.",
				ArrangeInput: []args.TwoAny{
					{
						First:  nil,
						Second: someNull,
					},
					{
						First:  someNull,
						Second: nil,
					},
				},
			ExpectedInput: []string{
				"0 - Equal : false - Conclusive ('<nil>', '*args.Two[interface {},interface {}]')",
				"1 - Equal : false - Conclusive ('*args.Two[interface {},interface {}]', '<nil>')",
			},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Expect all equal and conclusive because value same and data type also same.",
				ArrangeInput: []args.TwoAny{
					{
						First:  1,
						Second: 1,
					},
					{
						First:  2,
						Second: 2,
					},
					{
						First:  float64(-1),
						Second: float64(-1),
					},
					{
						First:  "some string",
						Second: "some string",
					},
				},
				ExpectedInput: []string{
					"0 - Equal : true - Conclusive ('1 - int', '1 - int')",
					"1 - Equal : true - Conclusive ('2 - int', '2 - int')",
					"2 - Equal : true - Conclusive ('-1 - float64', '-1 - float64')",
					"3 - Equal : true - Conclusive ('some string - string', 'some string - string')",
				},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Expect all not equal and inconclusive because value different and data type also same and non pointer.",
				ArrangeInput: []args.TwoAny{
					{
						First:  1,
						Second: 5,
					},
					{
						First:  2,
						Second: 3,
					},
					{
						First:  1,
						Second: 3,
					},
					{
						First:  "some string",
						Second: "some stringx",
					},
				},
				ExpectedInput: []string{
					"0 - Equal : false - Inconclusive ('1 - int', '5 - int')",
					"1 - Equal : false - Inconclusive ('2 - int', '3 - int')",
					"2 - Equal : false - Inconclusive ('1 - int', '3 - int')",
					"3 - Equal : false - Inconclusive ('some string - string', 'some stringx - string')",
				},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Same value different type results not equal and conclusive.",
				ArrangeInput: []args.TwoAny{
					{
						First:  1,
						Second: byte(2),
					},
					{
						First:  1,
						Second: float64(1),
					},
					{
						First:  "1",
						Second: 1,
					},
				},
				ExpectedInput: []string{
					"0 - Equal : false - Conclusive ('1 - int', '2 - uint8')",
					"1 - Equal : false - Conclusive ('1 - int', '1 - float64')",
					"2 - Equal : false - Conclusive ('1 - string', '1 - int')",
				},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	reflectionTypesTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Primitive types verification test.",
				ArrangeInput: []args.OneFuncAny{
					{
						First:    1,
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    args.TwoAny{},
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    "some string",
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    float32(23),
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    uint(23),
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    uint32(23),
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    uint64(23),
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    int32(23),
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    int64(23),
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    []int{1},
						WorkFunc: isany.PrimitiveType,
					},
					{
						First:    true,
						WorkFunc: isany.PrimitiveType,
					},
				},
				ExpectedInput: []string{
					"0 : true (type: int, PrimitiveType, 1)",
					"1 : false (type: args.Two[interface {},interface {}], PrimitiveType, {<nil> <nil> <nil> [] false { false}})",
					"2 : true (type: string, PrimitiveType, some string)",
					"3 : true (type: float32, PrimitiveType, 23)",
					"4 : true (type: uint, PrimitiveType, 23)",
					"5 : true (type: uint32, PrimitiveType, 23)",
					"6 : true (type: uint64, PrimitiveType, 23)",
					"7 : true (type: int32, PrimitiveType, 23)",
					"8 : true (type: int64, PrimitiveType, 23)",
					"9 : false (type: []int, PrimitiveType, [1])",
					"10 : true (type: bool, PrimitiveType, true)",
				},
				VerifyTypeOf: oneFuncTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "NumberType : verification test.",
				ArrangeInput: []args.OneFuncAny{
					{
						First:    1,
						WorkFunc: isany.NumberType,
					},
					{
						First:    float32(2),
						WorkFunc: isany.NumberType,
					},
					{
						First:    int64(1),
						WorkFunc: isany.NumberType,
					},
					{
						First:    byte(23),
						WorkFunc: isany.NumberType,
					},
					{
						First:    uint(23),
						WorkFunc: isany.NumberType,
					},
					{
						First:    uint32(23),
						WorkFunc: isany.NumberType,
					},
					{
						First:    uint64(23),
						WorkFunc: isany.NumberType,
					},
					{
						First:    int32(23),
						WorkFunc: isany.NumberType,
					},
					{
						First:    int64(23),
						WorkFunc: isany.NumberType,
					},
					{
						First:    []int{1},
						WorkFunc: isany.NumberType,
					},
					{
						First:    true,
						WorkFunc: isany.NumberType,
					},
				},
				ExpectedInput: []string{
					"0 : true (type: int, NumberType, 1)",
					"1 : true (type: float32, NumberType, 2)",
					"2 : true (type: int64, NumberType, 1)",
					"3 : true (type: uint8, NumberType, 23)",
					"4 : true (type: uint, NumberType, 23)",
					"5 : true (type: uint32, NumberType, 23)",
					"6 : true (type: uint64, NumberType, 23)",
					"7 : true (type: int32, NumberType, 23)",
					"8 : true (type: int64, NumberType, 23)",
					"9 : false (type: []int, NumberType, [1])",
					"10 : false (type: bool, NumberType, true)",
				},
				VerifyTypeOf: oneFuncTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Floating : verification test.",
				ArrangeInput: []args.OneFuncAny{
					{
						First:    1,
						WorkFunc: isany.FloatingPointType,
					},
					{
						First:    args.TwoAny{},
						WorkFunc: isany.FloatingPointType,
					},
					{
						First:    "some string",
						WorkFunc: isany.FloatingPointType,
					},
					{
						First:    float32(23),
						WorkFunc: isany.FloatingPointType,
					},
					{
						First:    1.5,
						WorkFunc: isany.FloatingPointType,
					},
					{
						First:    float64(65),
						WorkFunc: isany.FloatingPointType,
					},
					{
						First:    true,
						WorkFunc: isany.FloatingPointType,
					},
				},
				ExpectedInput: []string{
					"0 : false (type: int, FloatingPointType, 1)",
					"1 : false (type: args.Two[interface {},interface {}], FloatingPointType, {<nil> <nil> <nil> [] false { false}})",
					"2 : false (type: string, FloatingPointType, some string)",
					"3 : true (type: float32, FloatingPointType, 23)",
					"4 : true (type: float64, FloatingPointType, 1.5)",
					"5 : true (type: float64, FloatingPointType, 65)",
					"6 : false (type: bool, FloatingPointType, true)",
				},
				VerifyTypeOf: oneFuncTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "PositiveIntegerType: verification test.",
				ArrangeInput: []args.OneFuncAny{
					{
						First:    uint(64),
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    byte(255),
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    uint16(16),
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    uint32(32),
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    uint64(64),
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    1,
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    args.OneFuncAny{},
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    "some string",
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    float32(23),
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    1.5,
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    float64(65),
						WorkFunc: isany.PositiveIntegerType,
					},
					{
						First:    true,
						WorkFunc: isany.PositiveIntegerType,
					},
				},
				ExpectedInput: []string{
					"0 : true (type: uint, PositiveIntegerType, 64)",
					"1 : true (type: uint8, PositiveIntegerType, 255)",
					"2 : true (type: uint16, PositiveIntegerType, 16)",
					"3 : true (type: uint32, PositiveIntegerType, 32)",
					"4 : true (type: uint64, PositiveIntegerType, 64)",
					"5 : false (type: int, PositiveIntegerType, 1)",
					"6 : false (type: args.OneFunc[interface {}], PositiveIntegerType, OneFunc {  })",
					"7 : false (type: string, PositiveIntegerType, some string)",
					"8 : false (type: float32, PositiveIntegerType, 23)",
					"9 : false (type: float64, PositiveIntegerType, 1.5)",
					"10 : false (type: float64, PositiveIntegerType, 65)",
					"11 : false (type: bool, PositiveIntegerType, true)",
				},
				VerifyTypeOf: oneFuncTypeVerification,
				IsEnable:     issetter.True,
			},
		},
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "FuncOnly : function verification test.",
				Parameters: &args.HolderAny{
					First: "isFunc",
				},
				ArrangeInput: []args.OneFuncAny{
					{
						First:    1,
						WorkFunc: isany.FuncOnly,
					},
					{
						First:    args.OneFuncAny{},
						WorkFunc: isany.FuncOnly,
					},
					{
						First:    "some string",
						WorkFunc: isany.FuncOnly,
					},
					{
						First:    someNull,
						WorkFunc: isany.FuncOnly,
					},
					{
						First:    nil,
						WorkFunc: isany.FuncOnly,
					},
					{
						First:    isany.PrimitiveType,
						WorkFunc: isany.FuncOnly,
					},
				},
				ExpectedInput: []string{
					"0 : false (type: int, FuncOnly, FuncOnly)",
				"1 : false (type: args.OneFunc[interface {}], FuncOnly, FuncOnly)",
				"2 : false (type: string, FuncOnly, FuncOnly)",
				"3 : false (type: *args.Two[interface {},interface {}], FuncOnly, FuncOnly)",
					"4 : false (type: <nil>, FuncOnly, FuncOnly)",
					"5 : true (type: func(interface {}) bool, FuncOnly, FuncOnly)",
				},
				VerifyTypeOf: oneFuncTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
