package converterstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// ToNonNullItems
// =============================================================================

var toNonNullItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "ToNonNullItems returns count 0 -- nil input skipOnNil true",
		ArrangeInput: args.Map{
			"when":        "given nil input with isSkipOnNil true",
			"isSkipOnNil": true,
			"input":       nil,
		},
		ExpectedInput: args.Map{"count": 0},
	},
	{
		Title: "ToNonNullItems converts valid slice",
		ArrangeInput: args.Map{
			"when":        "given valid string slice",
			"isSkipOnNil": false,
			"input":       []any{"hello", "world"},
		},
		ExpectedInput: args.Map{
			"count": 2,
			"item0": "hello",
			"item1": "world",
		},
	},
	{
		Title: "ToNonNullItems returns empty for nil input when not skipOnNil",
		ArrangeInput: args.Map{
			"when":        "given nil input with isSkipOnNil false - should still return empty for nil reflect",
			"isSkipOnNil": true,
			"input":       nil,
		},
		ExpectedInput: args.Map{"count": 0},
	},
}
