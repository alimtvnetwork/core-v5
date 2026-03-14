package isanytests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ============================================================================
// NotNull
// ============================================================================

var extNotNullTestCases = []coretestcases.CaseV1{
	{
		Title: "NotNull returns false for nil",
		ArrangeInput: args.Map{
			"input": nil,
		},
		ExpectedInput: "false",
	},
	{
		Title: "NotNull returns true for non-nil",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "true",
	},
}

// ============================================================================
// AllNull
// ============================================================================

var extAllNullTestCases = []coretestcases.CaseV1{
	{
		Title: "AllNull returns true for all nil",
		ArrangeInput: args.Map{
			"inputs": []any{nil, nil, nil},
		},
		ExpectedInput: "true",
	},
	{
		Title: "AllNull returns false for any non-nil",
		ArrangeInput: args.Map{
			"inputs": []any{nil, "something", nil},
		},
		ExpectedInput: "false",
	},
	{
		Title: "AllNull returns true for empty",
		ArrangeInput: args.Map{
			"inputs": []any{},
		},
		ExpectedInput: "true",
	},
}

// ============================================================================
// AnyNull
// ============================================================================

var extAnyNullTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyNull returns true when any is nil",
		ArrangeInput: args.Map{
			"inputs": []any{"hello", nil, "world"},
		},
		ExpectedInput: "true",
	},
	{
		Title: "AnyNull returns false when none is nil",
		ArrangeInput: args.Map{
			"inputs": []any{"hello", "world"},
		},
		ExpectedInput: "false",
	},
	{
		Title: "AnyNull returns false for empty",
		ArrangeInput: args.Map{
			"inputs": []any{},
		},
		ExpectedInput: "false",
	},
}

// ============================================================================
// Zero / AllZero / AnyZero
// ============================================================================

var extZeroTestCases = []coretestcases.CaseV1{
	{
		Title: "Zero returns true for zero int",
		ArrangeInput: args.Map{
			"input": 0,
		},
		ExpectedInput: "true",
	},
	{
		Title: "Zero returns false for non-zero int",
		ArrangeInput: args.Map{
			"input": 42,
		},
		ExpectedInput: "false",
	},
	{
		Title: "Zero returns true for empty string",
		ArrangeInput: args.Map{
			"input": "",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Zero returns false for non-empty string",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: "false",
	},
	{
		Title: "Zero returns true for false bool",
		ArrangeInput: args.Map{
			"input": false,
		},
		ExpectedInput: "true",
	},
}

var extAllZeroTestCases = []coretestcases.CaseV1{
	{
		Title: "AllZero returns true for all zero values",
		ArrangeInput: args.Map{
			"inputs": []any{0, "", false},
		},
		ExpectedInput: "true",
	},
	{
		Title: "AllZero returns false when any non-zero",
		ArrangeInput: args.Map{
			"inputs": []any{0, "hello", false},
		},
		ExpectedInput: "false",
	},
}

var extAnyZeroTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyZero returns true when any is zero",
		ArrangeInput: args.Map{
			"inputs": []any{42, "", true},
		},
		ExpectedInput: "true",
	},
	{
		Title: "AnyZero returns false when none is zero",
		ArrangeInput: args.Map{
			"inputs": []any{42, "hello", true},
		},
		ExpectedInput: "false",
	},
}

// ============================================================================
// Pointer / Function / TypeSame
// ============================================================================

var extPointerTestCases = []coretestcases.CaseV1{
	{
		Title: "Pointer returns true for pointer type",
		ArrangeInput: args.Map{
			"usePointer": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "Pointer returns false for non-pointer",
		ArrangeInput: args.Map{
			"usePointer": false,
			"input":      42,
		},
		ExpectedInput: "false",
	},
}

var extFunctionTestCases = []coretestcases.CaseV1{
	{
		Title: "Function returns true for func type",
		ArrangeInput: args.Map{
			"useFunc": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "Function returns false for non-func",
		ArrangeInput: args.Map{
			"useFunc": false,
			"input":   42,
		},
		ExpectedInput: "false",
	},
}

var extTypeSameTestCases = []coretestcases.CaseV1{
	{
		Title: "TypeSame returns true for same types",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
		ExpectedInput: "true",
	},
	{
		Title: "TypeSame returns false for different types",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": 42,
		},
		ExpectedInput: "false",
	},
}

// ============================================================================
// StringEqual
// ============================================================================

var extStringEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "StringEqual returns true for same strings",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "hello",
		},
		ExpectedInput: "true",
	},
	{
		Title: "StringEqual returns false for different strings",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
		ExpectedInput: "false",
	},
}

// ============================================================================
// Defined / DefinedAllOf / DefinedAnyOf
// ============================================================================

var extDefinedAllOfTestCases = []coretestcases.CaseV1{
	{
		Title: "DefinedAllOf returns true when all non-nil",
		ArrangeInput: args.Map{
			"inputs": []any{"a", "b", 1},
		},
		ExpectedInput: "true",
	},
	{
		Title: "DefinedAllOf returns false when any nil",
		ArrangeInput: args.Map{
			"inputs": []any{"a", nil, 1},
		},
		ExpectedInput: "false",
	},
}

var extDefinedAnyOfTestCases = []coretestcases.CaseV1{
	{
		Title: "DefinedAnyOf returns true when any non-nil",
		ArrangeInput: args.Map{
			"inputs": []any{nil, "something", nil},
		},
		ExpectedInput: "true",
	},
	{
		Title: "DefinedAnyOf returns false when all nil",
		ArrangeInput: args.Map{
			"inputs": []any{nil, nil, nil},
		},
		ExpectedInput: "false",
	},
}

// ============================================================================
// ReflectNull / ReflectNotNull
// ============================================================================

var extReflectNullTestCases = []coretestcases.CaseV1{
	{
		Title: "ReflectNull returns true for nil pointer reflect value",
		ArrangeInput: args.Map{
			"useNilPtr": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "ReflectNull returns false for non-nil reflect value",
		ArrangeInput: args.Map{
			"useNilPtr": false,
		},
		ExpectedInput: "false",
	},
}

// ============================================================================
// DefinedLeftRight
// ============================================================================

var extDefinedLeftRightTestCases = []coretestcases.CaseV1{
	{
		Title: "DefinedLeftRight both defined",
		ArrangeInput: args.Map{
			"left":  "a",
			"right": "b",
		},
		ExpectedInput: args.Map{
			"leftDefined":  "true",
			"rightDefined": "true",
		},
	},
	{
		Title: "DefinedLeftRight left nil",
		ArrangeInput: args.Map{
			"leftNil": true,
			"right":   "b",
		},
		ExpectedInput: args.Map{
			"leftDefined":  "false",
			"rightDefined": "true",
		},
	},
}

// ============================================================================
// DefinedItems
// ============================================================================

var extDefinedItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "DefinedItems returns only non-nil items",
		ArrangeInput: args.Map{
			"inputs": []any{"a", nil, "c", nil, "e"},
		},
		ExpectedInput: args.Map{
			"count": 3,
		},
	},
	{
		Title: "DefinedItems returns empty for all nil",
		ArrangeInput: args.Map{
			"inputs": []any{nil, nil},
		},
		ExpectedInput: args.Map{
			"count": 0,
		},
	},
}

// ============================================================================
// NotDeepEqual
// ============================================================================

var extNotDeepEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "NotDeepEqual returns true for different values",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
		ExpectedInput: "true",
	},
	{
		Title: "NotDeepEqual returns false for same values",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "hello",
		},
		ExpectedInput: "false",
	},
}

// ============================================================================
// Conclusive
// ============================================================================

var extConclusiveTestCases = []coretestcases.CaseV1{
	{
		Title: "Conclusive returns true for Equal",
		ArrangeInput: args.Map{
			"left":  42,
			"right": 42,
		},
		ExpectedInput: "true true",
	},
	{
		Title: "Conclusive returns false for Inconclusive",
		ArrangeInput: args.Map{
			"left":  42,
			"right": 43,
		},
		ExpectedInput: "false false",
	},
}

// ============================================================================
// FuncOnly
// ============================================================================

var extFuncOnlyTestCases = []coretestcases.CaseV1{
	{
		Title: "FuncOnly returns true for function",
		ArrangeInput: args.Map{
			"useFunc": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "FuncOnly returns false for non-function",
		ArrangeInput: args.Map{
			"useFunc": false,
		},
		ExpectedInput: "false",
	},
}
