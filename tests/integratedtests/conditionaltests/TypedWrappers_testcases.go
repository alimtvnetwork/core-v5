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
