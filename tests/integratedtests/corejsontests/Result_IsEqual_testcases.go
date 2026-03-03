package corejsontests

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// IsEqual test cases
// =============================================================================

var resultIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqual - same content returns true",
		ArrangeInput: args.Map{
			"a": corejson.New(map[string]string{"key": "value"}),
			"b": corejson.New(map[string]string{"key": "value"}),
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsEqual - different content returns false",
		ArrangeInput: args.Map{
			"a": corejson.New(map[string]string{"key": "a"}),
			"b": corejson.New(map[string]string{"key": "b"}),
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// IsEqualPtr test cases
// =============================================================================

var resultIsEqualPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqualPtr - both nil returns true",
		ArrangeInput: args.Map{
			"aPtr": (*corejson.Result)(nil),
			"bPtr": (*corejson.Result)(nil),
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsEqualPtr - one nil returns false",
		ArrangeInput: args.Map{
			"aPtr": corejson.NewPtr(map[string]string{"k": "v"}),
			"bPtr": (*corejson.Result)(nil),
		},
		ExpectedInput: []string{"false"},
	},
}
