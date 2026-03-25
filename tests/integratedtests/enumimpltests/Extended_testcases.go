package enumimpltests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// DynamicMap
// ==========================================

var extDynMapBasicTestCases = []coretestcases.CaseV1{
	{
		Title: "DynamicMap Length and IsEmpty",
		ArrangeInput: args.Map{
			"items": map[string]any{"A": 1, "B": 2},
		},
		ExpectedInput: args.Map{
			"length":  "2",
			"isEmpty": "false",
		},
	},
	{
		Title: "DynamicMap empty",
		ArrangeInput: args.Map{
			"items": map[string]any{},
		},
		ExpectedInput: args.Map{
			"length":  "0",
			"isEmpty": "true",
		},
	},
}

// ==========================================
// DiffLeftRight
// ==========================================

var extDiffLeftRightTestCases = []coretestcases.CaseV1{
	{
		Title: "DiffLeftRight same values",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "hello",
		},
		ExpectedInput: args.Map{
			"isSame":     "true",
			"isNotEqual": "false",
		},
	},
	{
		Title: "DiffLeftRight different values",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
		ExpectedInput: args.Map{
			"isSame":     "false",
			"isNotEqual": "true",
		},
	},
}

// ==========================================
// KeyAnyVal
// ==========================================

var extKeyAnyValTestCases = []coretestcases.CaseV1{
	{
		Title: "KeyAnyVal with int value",
		ArrangeInput: args.Map{
			"key":   "Invalid",
			"value": 0,
		},
		ExpectedInput: args.Map{
			"key":      "Invalid",
			"valInt":   "0",
			"isString": "false",
		},
	},
	{
		Title: "KeyAnyVal with string value",
		ArrangeInput: args.Map{
			"key":   "Name",
			"value": "hello",
		},
		ExpectedInput: args.Map{
			"key":      "Name",
			"isString": "true",
		},
	},
}

// ==========================================
// KeyValInteger
// ==========================================

var extKeyValIntegerTestCases = []coretestcases.CaseV1{
	{
		Title: "KeyValInteger normal",
		ArrangeInput: args.Map{
			"key":   "Variant",
			"value": 5,
		},
		ExpectedInput: args.Map{
			"key":      "Variant",
			"isString": "false",
		},
	},
}
