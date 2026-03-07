package mapdiffinternaltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Length / IsEmpty / HasAnyItem / LastIndex / Raw
// ==========================================================================

var mapStringAnyDiffLengthTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty map has length 0",
		ArrangeInput: args.Map{
			"map": map[string]any{},
		},
		ExpectedInput: args.Map{
			"length":     0,
			"isEmpty":    true,
			"hasAnyItem": false,
			"lastIndex":  -1,
		},
	},
	{
		Title: "Map with 3 items has length 3",
		ArrangeInput: args.Map{
			"map": map[string]any{
				"a": 1,
				"b": "hello",
				"c": true,
			},
		},
		ExpectedInput: args.Map{
			"length":     3,
			"isEmpty":    false,
			"hasAnyItem": true,
			"lastIndex":  2,
		},
	},
}

// ==========================================================================
// AllKeysSorted
// ==========================================================================

var mapStringAnyDiffAllKeysSortedTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty map returns empty keys",
		ArrangeInput: args.Map{
			"map": map[string]any{},
		},
		ExpectedInput: []string{},
	},
	{
		Title: "Keys returned in sorted order",
		ArrangeInput: args.Map{
			"map": map[string]any{
				"zebra": 1,
				"apple": 2,
				"mango": 3,
			},
		},
		ExpectedInput: []string{"apple", "mango", "zebra"},
	},
}

// ==========================================================================
// IsRawEqual
// ==========================================================================

var mapStringAnyDiffIsRawEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal maps return true (strict type)",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
				"b": "hello",
			},
			"right": map[string]any{
				"a": 1,
				"b": "hello",
			},
		},
		ExpectedInput: "true",
	},
	{
		Title: "Different types return false when strict",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": "1",
			},
		},
		ExpectedInput: "false",
	},
	{
		Title: "Different types return true when regardless",
		ArrangeInput: args.Map{
			"isRegardlessType": true,
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": "1",
			},
		},
		ExpectedInput: "true",
	},
	{
		Title: "Missing key in right returns false",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
				"b": 2,
			},
			"right": map[string]any{
				"a": 1,
			},
		},
		ExpectedInput: "false",
	},
}

// ==========================================================================
// HasAnyChanges
// ==========================================================================

var mapStringAnyDiffHasAnyChangesTestCases = []coretestcases.CaseV1{
	{
		Title: "Same maps have no changes",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": 1,
			},
		},
		ExpectedInput: "false",
	},
	{
		Title: "Different maps have changes",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": 2,
			},
		},
		ExpectedInput: "true",
	},
}

// ==========================================================================
// DiffRaw
// ==========================================================================

var mapStringAnyDiffDiffRawTestCases = []coretestcases.CaseV1{
	{
		Title: "Identical maps yield empty diff",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": 1,
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 0,
		},
	},
	{
		Title: "Key in left missing in right appears in diff",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
				"b": 2,
			},
			"right": map[string]any{
				"a": 1,
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 1,
			"hasKey-b":   true,
		},
	},
	{
		Title: "Key in right missing in left appears in diff",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": 1,
				"c": 3,
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 1,
			"hasKey-c":   true,
		},
	},
	{
		Title: "Different values appear in diff",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": 99,
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 1,
			"hasKey-a":   true,
		},
	},
	{
		Title: "Both nil yield empty diff",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"leftNil":          true,
			"rightNil":         true,
		},
		ExpectedInput: args.Map{
			"diffLength": 0,
		},
	},
	{
		Title: "Left nil right non-nil returns right",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"leftNil":          true,
			"right": map[string]any{
				"x": 10,
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 1,
			"hasKey-x":   true,
		},
	},
}

// ==========================================================================
// ShouldDiffMessage
// ==========================================================================

var mapStringAnyDiffShouldDiffMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal maps produce empty message",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"title":            "test-title",
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": 1,
			},
		},
		ExpectedInput: "",
	},
	{
		Title: "Different maps produce non-empty message containing title",
		ArrangeInput: args.Map{
			"isRegardlessType": false,
			"title":            "my-diff",
			"left": map[string]any{
				"a": 1,
			},
			"right": map[string]any{
				"a": 2,
			},
		},
		ExpectedInput: args.Map{
			"containsTitle": true,
			"isNotEmpty":    true,
		},
	},
}

// ==========================================================================
// ToStringsSliceOfDiffMap — string vs non-string formatting
// ==========================================================================

var mapStringAnyDiffToStringsSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "String values get quotation-wrapped format",
		ArrangeInput: args.Map{
			"diffMap": map[string]any{
				"name": "alice",
			},
		},
		ExpectedInput: args.Map{
			"length":       1,
			"hasQuotation": true,
		},
	},
	{
		Title: "Non-string values get raw format",
		ArrangeInput: args.Map{
			"diffMap": map[string]any{
				"count": 42,
			},
		},
		ExpectedInput: args.Map{
			"length":       1,
			"hasQuotation": false,
		},
	},
}
