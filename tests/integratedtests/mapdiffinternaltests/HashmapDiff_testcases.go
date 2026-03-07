package mapdiffinternaltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Length / IsEmpty / HasAnyItem / LastIndex
// ==========================================================================

var hashmapDiffLengthTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty map has length 0",
		ArrangeInput: args.Map{
			"map": map[string]string{},
		},
		ExpectedInput: args.Map{
			"length":     0,
			"isEmpty":    true,
			"hasAnyItem": false,
			"lastIndex":  -1,
		},
	},
	{
		Title: "Map with 2 items has length 2",
		ArrangeInput: args.Map{
			"map": map[string]string{
				"a": "1",
				"b": "2",
			},
		},
		ExpectedInput: args.Map{
			"length":     2,
			"isEmpty":    false,
			"hasAnyItem": true,
			"lastIndex":  1,
		},
	},
}

// ==========================================================================
// AllKeysSorted
// ==========================================================================

var hashmapDiffAllKeysSortedTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty map returns empty keys",
		ArrangeInput: args.Map{
			"map": map[string]string{},
		},
		ExpectedInput: []string{},
	},
	{
		Title: "Keys returned in sorted order",
		ArrangeInput: args.Map{
			"map": map[string]string{
				"charlie": "3",
				"alpha":   "1",
				"bravo":   "2",
			},
		},
		ExpectedInput: []string{"alpha", "bravo", "charlie"},
	},
}

// ==========================================================================
// IsRawEqual
// ==========================================================================

var hashmapDiffIsRawEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal maps return true",
		ArrangeInput: args.Map{
			"left": map[string]string{
				"a": "1",
				"b": "2",
			},
			"right": map[string]string{
				"a": "1",
				"b": "2",
			},
		},
		ExpectedInput: "true",
	},
	{
		Title: "Different values return false",
		ArrangeInput: args.Map{
			"left": map[string]string{
				"a": "1",
			},
			"right": map[string]string{
				"a": "99",
			},
		},
		ExpectedInput: "false",
	},
	{
		Title: "Missing key in right returns false",
		ArrangeInput: args.Map{
			"left": map[string]string{
				"a": "1",
				"b": "2",
			},
			"right": map[string]string{
				"a": "1",
			},
		},
		ExpectedInput: "false",
	},
}

// ==========================================================================
// DiffRaw
// ==========================================================================

var hashmapDiffDiffRawTestCases = []coretestcases.CaseV1{
	{
		Title: "Identical maps yield empty diff",
		ArrangeInput: args.Map{
			"left": map[string]string{
				"a": "1",
			},
			"right": map[string]string{
				"a": "1",
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 0,
		},
	},
	{
		Title: "Key missing in right appears in diff",
		ArrangeInput: args.Map{
			"left": map[string]string{
				"a": "1",
				"b": "2",
			},
			"right": map[string]string{
				"a": "1",
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 1,
			"hasKey-b":   true,
		},
	},
	{
		Title: "Key missing in left appears in diff",
		ArrangeInput: args.Map{
			"left": map[string]string{
				"a": "1",
			},
			"right": map[string]string{
				"a": "1",
				"c": "3",
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
			"left": map[string]string{
				"a": "1",
			},
			"right": map[string]string{
				"a": "99",
			},
		},
		ExpectedInput: args.Map{
			"diffLength": 1,
			"hasKey-a":   true,
		},
	},
	{
		Title: "Both nil maps yield empty diff",
		ArrangeInput: args.Map{
			"leftNil":  true,
			"rightNil": true,
		},
		ExpectedInput: args.Map{
			"diffLength": 0,
		},
	},
	{
		Title: "Left nil right non-nil returns right",
		ArrangeInput: args.Map{
			"leftNil": true,
			"right": map[string]string{
				"x": "10",
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

var hashmapDiffShouldDiffMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal maps produce empty message",
		ArrangeInput: args.Map{
			"title": "test-title",
			"left": map[string]string{
				"a": "1",
			},
			"right": map[string]string{
				"a": "1",
			},
		},
		ExpectedInput: "",
	},
	{
		Title: "Different maps produce non-empty message with title",
		ArrangeInput: args.Map{
			"title": "my-diff",
			"left": map[string]string{
				"a": "1",
			},
			"right": map[string]string{
				"a": "2",
			},
		},
		ExpectedInput: args.Map{
			"containsTitle": true,
			"isNotEmpty":    true,
		},
	},
}
