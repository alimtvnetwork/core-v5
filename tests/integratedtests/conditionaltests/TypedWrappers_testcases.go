package conditionaltests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// IfBool -- typed bool wrapper
// ==========================================================================

var ifBoolTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfBool returns true -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IfBool returns false -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "false",
	},
}

// ==========================================================================
// IfInt -- typed int wrapper
// ==========================================================================

var ifIntTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfInt returns trueValue -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "10",
	},
	{
		Title: "IfInt returns falseValue -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "20",
	},
}

// ==========================================================================
// IfByte -- typed byte wrapper
// ==========================================================================

var ifByteTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfByte returns trueValue -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  byte(1),
			"falseValue": byte(0),
		},
		ExpectedInput: "1",
	},
	{
		Title: "IfByte returns falseValue -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  byte(1),
			"falseValue": byte(0),
		},
		ExpectedInput: "0",
	},
}

// ==========================================================================
// IfFloat64 -- typed float64 wrapper
// ==========================================================================

var ifFloat64TypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFloat64 returns trueValue -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  3.14,
			"falseValue": 2.71,
		},
		ExpectedInput: "3.14",
	},
	{
		Title: "IfFloat64 returns falseValue -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  3.14,
			"falseValue": 2.71,
		},
		ExpectedInput: "2.71",
	},
}

// ==========================================================================
// IfAny -- typed any wrapper
// ==========================================================================

var ifAnyTypedTestCases = []coretestcases.CaseV1{
	{
		Title: "IfAny returns trueValue -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "yes",
	},
	{
		Title: "IfAny returns falseValue -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "no",
	},
}

// ==========================================================================
// IfFuncBool -- func bool wrapper
// ==========================================================================

var ifFuncBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFuncBool returns trueFunc result -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IfFuncBool returns falseFunc result -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  true,
			"falseValue": false,
		},
		ExpectedInput: "false",
	},
}

// ==========================================================================
// IfFuncInt -- func int wrapper
// ==========================================================================

var ifFuncIntTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFuncInt returns trueFunc result -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  42,
			"falseValue": 0,
		},
		ExpectedInput: "42",
	},
	{
		Title: "IfFuncInt returns falseFunc result -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  42,
			"falseValue": 0,
		},
		ExpectedInput: "0",
	},
}

// ==========================================================================
// IfFuncAny -- func any wrapper
// ==========================================================================

var ifFuncAnyTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFuncAny returns trueFunc result -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  "from-true",
			"falseValue": "from-false",
		},
		ExpectedInput: "from-true",
	},
	{
		Title: "IfFuncAny returns falseFunc result -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  "from-true",
			"falseValue": "from-false",
		},
		ExpectedInput: "from-false",
	},
}

// ==========================================================================
// IfTrueFuncBool -- true-only func bool wrapper
// ==========================================================================

var ifTrueFuncBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFuncBool returns trueFunc result -- condition true",
		ArrangeInput: args.Map{
			"isTrue":    true,
			"trueValue": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IfTrueFuncBool returns zero -- condition false",
		ArrangeInput: args.Map{
			"isTrue":    false,
			"trueValue": true,
		},
		ExpectedInput: "false",
	},
}

// ==========================================================================
// IfTrueFuncStrings -- true-only func []string wrapper
// ==========================================================================

var ifTrueFuncStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFuncStrings returns slice -- condition true",
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
		Title: "IfTrueFuncStrings returns nil -- condition false",
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

// ==========================================================================
// IfTrueFuncBytes -- true-only func []byte wrapper
// ==========================================================================

var ifTrueFuncBytesTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFuncBytes returns slice -- condition true",
		ArrangeInput: args.Map{
			"isTrue":    true,
			"trueValue": []byte{65, 66},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "65",
		},
	},
	{
		Title: "IfTrueFuncBytes returns nil -- condition false",
		ArrangeInput: args.Map{
			"isTrue":    false,
			"trueValue": []byte{65, 66},
		},
		ExpectedInput: args.Map{
			"length": "0",
			"isNil":  "true",
		},
	},
}

// ==========================================================================
// IfSliceBool -- slice bool wrapper
// ==========================================================================

var ifSliceBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceBool returns trueSlice -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []bool{true, false},
			"falseValue": []bool{false, true},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "true",
		},
	},
	{
		Title: "IfSliceBool returns falseSlice -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []bool{true, false},
			"falseValue": []bool{false, true},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "false",
		},
	},
}

// ==========================================================================
// IfSliceInt -- slice int wrapper
// ==========================================================================

var ifSliceIntTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceInt returns trueSlice -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []int{10, 20},
			"falseValue": []int{30, 40},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "10",
		},
	},
	{
		Title: "IfSliceInt returns falseSlice -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []int{10, 20},
			"falseValue": []int{30, 40},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "30",
		},
	},
}

// ==========================================================================
// IfSliceString -- slice string wrapper
// ==========================================================================

var ifSliceStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceString returns trueSlice -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []string{"a", "b"},
			"falseValue": []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "a",
		},
	},
	{
		Title: "IfSliceString returns falseSlice -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []string{"a", "b"},
			"falseValue": []string{"x", "y", "z"},
		},
		ExpectedInput: args.Map{
			"length": "3",
			"first":  "x",
		},
	},
}

// ==========================================================================
// IfSliceByte -- slice byte wrapper
// ==========================================================================

var ifSliceByteTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceByte returns trueSlice -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []byte{1, 2},
			"falseValue": []byte{3, 4},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "1",
		},
	},
	{
		Title: "IfSliceByte returns falseSlice -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []byte{1, 2},
			"falseValue": []byte{3, 4},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "3",
		},
	},
}

// ==========================================================================
// IfSliceAny -- slice any wrapper
// ==========================================================================

var ifSliceAnyTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSliceAny returns trueSlice -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  []any{"a", 1},
			"falseValue": []any{"b", 2},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "a",
		},
	},
	{
		Title: "IfSliceAny returns falseSlice -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  []any{"a", 1},
			"falseValue": []any{"b", 2},
		},
		ExpectedInput: args.Map{
			"length": "2",
			"first":  "b",
		},
	},
}

// ==========================================================================
// IfPtrString -- pointer string wrapper
// ==========================================================================

var ifPtrStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfPtrString returns truePtr -- condition true",
		ArrangeInput: args.Map{
			"isTrue":     true,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "yes",
		},
	},
	{
		Title: "IfPtrString returns falsePtr -- condition false",
		ArrangeInput: args.Map{
			"isTrue":     false,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "no",
		},
	},
}

// ==========================================================================
// IfPtrInt -- pointer int wrapper
// ==========================================================================

var ifPtrIntTestCases = []coretestcases.CaseV1{
	{
		Title: "IfPtrInt returns truePtr -- condition true",
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
		Title: "IfPtrInt returns falsePtr -- condition false",
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

// ==========================================================================
// IfPtrBool -- pointer bool wrapper
// ==========================================================================

var ifPtrBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IfPtrBool returns truePtr -- condition true",
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
		Title: "IfPtrBool returns falsePtr -- condition false",
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
