package typesconvtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// StringToBool
// ==========================================================================

var stringToBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "StringToBool returns true for 'true'",
		ArrangeInput: args.Map{
			"when":  "given 'true'",
			"input": "true",
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "StringToBool returns true for 'yes'",
		ArrangeInput: args.Map{
			"when":  "given 'yes'",
			"input": "yes",
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "StringToBool returns false for empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title: "StringToBool returns false for 'no'",
		ArrangeInput: args.Map{
			"when":  "given 'no'",
			"input": "no",
		},
		ExpectedInput: args.Map{"result": false},
	},
}

// ==========================================================================
// IntPtrToSimple / IntPtrToSimpleDef
// ==========================================================================

var intPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "IntPtrToSimple returns value for non-nil",
		ArrangeInput: args.Map{
			"when":  "given non-nil int pointer",
			"isNil": false,
			"value": 42,
		},
		ExpectedInput: args.Map{"result": 42},
	},
	{
		Title: "IntPtrToSimple returns 0 for nil",
		ArrangeInput: args.Map{
			"when":  "given nil int pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{"result": 0},
	},
}

var intPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title: "IntPtrToSimpleDef returns value for non-nil",
		ArrangeInput: args.Map{
			"when":   "given non-nil int pointer",
			"isNil":  false,
			"value":  42,
			"defVal": 99,
		},
		ExpectedInput: args.Map{
			"result":      42,
			"defaultUsed": false,
		},
	},
	{
		Title: "IntPtrToSimpleDef returns default for nil",
		ArrangeInput: args.Map{
			"when":   "given nil int pointer",
			"isNil":  true,
			"defVal": 99,
		},
		ExpectedInput: args.Map{
			"result":      99,
			"defaultUsed": true,
		},
	},
}

var intPtrTestCase = coretestcases.CaseV1{
	Title: "IntPtr creates pointer to value",
	ExpectedInput: args.Map{
		"value": 42,
		"isNil": false,
	},
}

var intPtrToDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "IntPtrToDefPtr returns original for non-nil",
		ArrangeInput: args.Map{
			"when":  "given non-nil",
			"isNil": false,
			"value": 10,
		},
		ExpectedInput: args.Map{"value": 10},
	},
	{
		Title: "IntPtrToDefPtr returns default for nil",
		ArrangeInput: args.Map{
			"when":  "given nil",
			"isNil": true,
		},
		ExpectedInput: args.Map{"value": 77},
	},
}

// ==========================================================================
// BoolPtrToSimple / BoolPtrToSimpleDef
// ==========================================================================

var boolPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "BoolPtrToSimple returns value for non-nil true",
		ArrangeInput: args.Map{
			"when":  "given non-nil true",
			"isNil": false,
			"value": true,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "BoolPtrToSimple returns value for non-nil false",
		ArrangeInput: args.Map{
			"when":  "given non-nil false",
			"isNil": false,
			"value": false,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title: "BoolPtrToSimple returns false for nil",
		ArrangeInput: args.Map{
			"when":  "given nil bool pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var boolPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title: "BoolPtrToSimpleDef returns value for non-nil",
		ArrangeInput: args.Map{
			"when":   "given non-nil",
			"isNil":  false,
			"value":  true,
			"defVal": false,
		},
		ExpectedInput: args.Map{
			"result":      true,
			"defaultUsed": false,
		},
	},
	{
		Title: "BoolPtrToSimpleDef returns default for nil",
		ArrangeInput: args.Map{
			"when":   "given nil",
			"isNil":  true,
			"defVal": true,
		},
		ExpectedInput: args.Map{
			"result":      true,
			"defaultUsed": true,
		},
	},
}

var boolPtrTestCase = coretestcases.CaseV1{
	Title: "BoolPtr creates pointer to value",
	ExpectedInput: args.Map{
		"value": true,
		"isNil": false,
	},
}

var boolPtrToDefPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "BoolPtrToDefPtr returns original for non-nil",
		ArrangeInput: args.Map{
			"when":  "given non-nil",
			"isNil": false,
			"value": true,
		},
		ExpectedInput: args.Map{"value": true},
	},
	{
		Title: "BoolPtrToDefPtr returns default for nil",
		ArrangeInput: args.Map{
			"when":  "given nil",
			"isNil": true,
		},
		ExpectedInput: args.Map{"value": false},
	},
}

// ==========================================================================
// BytePtrToSimple / BytePtrToSimpleDef
// ==========================================================================

var bytePtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "BytePtrToSimple returns value for non-nil",
		ArrangeInput: args.Map{
			"when":  "given non-nil byte pointer",
			"isNil": false,
			"value": 255,
		},
		ExpectedInput: args.Map{"result": 255},
	},
	{
		Title: "BytePtrToSimple returns 0 for nil",
		ArrangeInput: args.Map{
			"when":  "given nil byte pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{"result": 0},
	},
}

var bytePtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title: "BytePtrToSimpleDef returns value for non-nil",
		ArrangeInput: args.Map{
			"when":   "given non-nil",
			"isNil":  false,
			"value":  100,
			"defVal": 50,
		},
		ExpectedInput: args.Map{
			"result":      100,
			"defaultUsed": false,
		},
	},
	{
		Title: "BytePtrToSimpleDef returns default for nil",
		ArrangeInput: args.Map{
			"when":   "given nil",
			"isNil":  true,
			"defVal": 50,
		},
		ExpectedInput: args.Map{
			"result":      50,
			"defaultUsed": true,
		},
	},
}

var bytePtrTestCase = coretestcases.CaseV1{
	Title: "BytePtr creates pointer to value",
	ExpectedInput: args.Map{
		"value": 42,
		"isNil": false,
	},
}

// ==========================================================================
// FloatPtrToSimple / FloatPtrToSimpleDef
// ==========================================================================

var floatPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "FloatPtrToSimple returns value for non-nil",
		ArrangeInput: args.Map{
			"when":  "given non-nil float pointer",
			"isNil": false,
			"value": 3.14,
		},
		ExpectedInput: args.Map{"result": float32(3.14)},
	},
	{
		Title: "FloatPtrToSimple returns 0 for nil",
		ArrangeInput: args.Map{
			"when":  "given nil float pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{"result": float32(0)},
	},
}

var floatPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title: "FloatPtrToSimpleDef returns value for non-nil",
		ArrangeInput: args.Map{
			"when":   "given non-nil",
			"isNil":  false,
			"value":  2.71,
			"defVal": 9.99,
		},
		ExpectedInput: args.Map{
			"result":      float32(2.71),
			"defaultUsed": false,
		},
	},
	{
		Title: "FloatPtrToSimpleDef returns default for nil",
		ArrangeInput: args.Map{
			"when":   "given nil",
			"isNil":  true,
			"defVal": 9.99,
		},
		ExpectedInput: args.Map{
			"result":      float32(9.99),
			"defaultUsed": true,
		},
	},
}

var floatPtrTestCase = coretestcases.CaseV1{
	Title: "FloatPtr creates pointer to value",
	ExpectedInput: args.Map{
		"value": float32(1.5),
		"isNil": false,
	},
}

// ==========================================================================
// StringPtrToSimple / StringPtrToSimpleDef / StringPointerToBool
// ==========================================================================

var stringPtrToSimpleTestCases = []coretestcases.CaseV1{
	{
		Title: "StringPtrToSimple returns value for non-nil",
		ArrangeInput: args.Map{
			"when":  "given non-nil string pointer",
			"isNil": false,
			"value": "hello",
		},
		ExpectedInput: args.Map{"result": "hello"},
	},
	{
		Title: "StringPtrToSimple returns empty for nil",
		ArrangeInput: args.Map{
			"when":  "given nil string pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{"result": ""},
	},
}

var stringPtrToSimpleDefTestCases = []coretestcases.CaseV1{
	{
		Title: "StringPtrToSimpleDef returns value for non-nil",
		ArrangeInput: args.Map{
			"when":   "given non-nil",
			"isNil":  false,
			"value":  "actual",
			"defVal": "fallback",
		},
		ExpectedInput: args.Map{
			"result":      "actual",
			"defaultUsed": false,
		},
	},
	{
		Title: "StringPtrToSimpleDef returns default for nil",
		ArrangeInput: args.Map{
			"when":   "given nil",
			"isNil":  true,
			"defVal": "fallback",
		},
		ExpectedInput: args.Map{
			"result":      "fallback",
			"defaultUsed": true,
		},
	},
}

var stringPtrTestCase = coretestcases.CaseV1{
	Title: "StringPtr creates pointer to value",
	ExpectedInput: args.Map{
		"value": "test",
		"isNil": false,
	},
}

var stringPointerToBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "StringPointerToBool returns true for 'true'",
		ArrangeInput: args.Map{
			"when":  "given pointer to 'true'",
			"isNil": false,
			"value": "true",
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "StringPointerToBool returns false for nil",
		ArrangeInput: args.Map{
			"when":  "given nil",
			"isNil": true,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title: "StringPointerToBool returns false for empty",
		ArrangeInput: args.Map{
			"when":  "given pointer to empty",
			"isNil": false,
			"value": "",
		},
		ExpectedInput: args.Map{"result": false},
	},
}
