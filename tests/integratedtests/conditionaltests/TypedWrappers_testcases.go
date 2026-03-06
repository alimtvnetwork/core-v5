package conditionaltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// --- IfBool ---

var ifBoolTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfBool true returns trueValue",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IfBool false returns falseValue",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "false",
	},
}

// --- IfInt ---

var ifIntTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfInt true returns trueValue",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  42,
			"falseValue": 99,
		},
		ExpectedInput: "42",
	},
	{
		Title: "IfInt false returns falseValue",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  42,
			"falseValue": 99,
		},
		ExpectedInput: "99",
	},
}

// --- IfByte ---

var ifByteTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfByte true returns trueValue",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  byte(65),
			"falseValue": byte(90),
		},
		ExpectedInput: "65",
	},
	{
		Title: "IfByte false returns falseValue",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  byte(65),
			"falseValue": byte(90),
		},
		ExpectedInput: "90",
	},
}

// --- IfFloat64 ---

var ifFloat64TypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFloat64 true returns trueValue",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  3.14,
			"falseValue": 2.71,
		},
		ExpectedInput: "3.14",
	},
	{
		Title: "IfFloat64 false returns falseValue",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  3.14,
			"falseValue": 2.71,
		},
		ExpectedInput: "2.71",
	},
}

// --- IfAny ---

var ifAnyTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfAny true returns trueValue",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  "any-true",
			"falseValue": "any-false",
		},
		ExpectedInput: "any-true",
	},
	{
		Title: "IfAny false returns falseValue",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  "any-true",
			"falseValue": "any-false",
		},
		ExpectedInput: "any-false",
	},
}

// --- IfFuncBool ---

var ifFuncBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFuncBool true evaluates trueFunc",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IfFuncBool false evaluates falseFunc",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "false",
	},
}

// --- IfFuncInt ---

var ifFuncIntTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFuncInt true evaluates trueFunc",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "10",
	},
	{
		Title: "IfFuncInt false evaluates falseFunc",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "20",
	},
}

// --- IfFuncString ---

var ifFuncStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFuncString true evaluates trueFunc",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "yes",
	},
	{
		Title: "IfFuncString false evaluates falseFunc",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "no",
	},
}

// --- IfFuncAny ---

var ifFuncAnyTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFuncAny true evaluates trueFunc",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  42,
			"falseValue": "fallback",
		},
		ExpectedInput: "42",
	},
	{
		Title: "IfFuncAny false evaluates falseFunc",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  42,
			"falseValue": "fallback",
		},
		ExpectedInput: "fallback",
	},
}

// --- IfTrueFuncBool ---

var ifTrueFuncBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFuncBool true returns trueFunc result",
		ArrangeInput: args.Map{
			"isTrue":    true,
			"trueValue": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IfTrueFuncBool false returns zero value (false)",
		ArrangeInput: args.Map{
			"isTrue":    false,
			"trueValue": true,
		},
		ExpectedInput: "false",
	},
}

// --- IfTrueFuncString ---

var ifTrueFuncStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFuncString true returns trueFunc result",
		ArrangeInput: args.Map{
			"isTrue":    true,
			"trueValue": "hello",
		},
		ExpectedInput: "hello",
	},
	{
		Title: "IfTrueFuncString false returns zero value (empty)",
		ArrangeInput: args.Map{
			"isTrue":    false,
			"trueValue": "hello",
		},
		ExpectedInput: "",
	},
}

// --- IfTrueFuncStrings ---

var ifTrueFuncStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFuncStrings true returns trueFunc result",
		ArrangeInput: args.Map{
			"isTrue":    true,
			"trueValue": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "a",
		},
	},
	{
		Title: "IfTrueFuncStrings false returns nil slice",
		ArrangeInput: args.Map{
			"isTrue":    false,
			"trueValue": []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"length": "0",
			"isNil":  "true",
		},
	},
}

// --- IfTrueFuncBytes ---

var ifTrueFuncBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFuncBytes true returns trueFunc result",
		ArrangeInput: args.Map{
			"isTrue":    true,
			"trueValue": []byte{1, 2, 3},
		},
		ExpectedInput: args.Map{
			"length": "3",
			"first":  "1",
		},
	},
	{
		Title: "IfTrueFuncBytes false returns nil slice",
		ArrangeInput: args.Map{
			"isTrue":    false,
			"trueValue": []byte{1, 2, 3},
		},
		ExpectedInput: args.Map{
			"length": "0",
			"isNil":  "true",
		},
	},
}

// --- IfSliceBool ---

var ifSliceBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceBool true returns trueSlice",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []bool{true, false},
			"falseValue": []bool{false},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "true",
		},
	},
	{
		Title: "IfSliceBool false returns falseSlice",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []bool{true, false},
			"falseValue": []bool{false},
		},
		ExpectedInput: args.Map{
			"length": "1",
			"first":  "false",
		},
	},
}

// --- IfSliceInt ---

var ifSliceIntTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceInt true returns trueSlice",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []int{1, 2, 3},
			"falseValue": []int{9},
		},
		ExpectedInput: args.Map{
			"length": "3",
			"first":  "1",
		},
	},
	{
		Title: "IfSliceInt false returns falseSlice",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []int{1, 2, 3},
			"falseValue": []int{9},
		},
		ExpectedInput: args.Map{
			"length": "1",
			"first":  "9",
		},
	},
}

// --- IfSliceAny ---

var ifSliceAnyTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceAny true returns trueSlice",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []any{"a", 1},
			"falseValue": []any{"x"},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "a",
		},
	},
	{
		Title: "IfSliceAny false returns falseSlice",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []any{"a", 1},
			"falseValue": []any{"x"},
		},
		ExpectedInput: args.Map{
			"length": "1",
			"first":  "x",
		},
	},
}

// --- IfPtrString ---

var ifPtrStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfPtrString true returns trueValue pointer",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  "alpha",
			"falseValue": "beta",
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "alpha",
		},
	},
	{
		Title: "IfPtrString false returns falseValue pointer",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  "alpha",
			"falseValue": "beta",
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "beta",
		},
	},
}

// --- IfPtrInt ---

var ifPtrIntTestCases = []coretestcases.CaseV1{
	{
		Title: "IfPtrInt true returns trueValue pointer",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "10",
		},
	},
	{
		Title: "IfPtrInt false returns falseValue pointer",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "20",
		},
	},
}

// --- IfPtrBool ---

var ifPtrBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfPtrBool true returns trueValue pointer",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "true",
		},
	},
	{
		Title: "IfPtrBool false returns falseValue pointer",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "false",
		},
	},
}
