package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// IsEqual test cases
// =============================================================================

var resultIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqual returns true -- same content",
		ArrangeInput: args.Map{
			"a": corejson.New(map[string]string{"key": "value"}),
			"b": corejson.New(map[string]string{"key": "value"}),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEqual returns false -- different content",
		ArrangeInput: args.Map{
			"a": corejson.New(map[string]string{"key": "a"}),
			"b": corejson.New(map[string]string{"key": "b"}),
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// IsEqualPtr test cases
// =============================================================================

var resultIsEqualPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqualPtr returns true -- both nil",
		ArrangeInput: args.Map{
			"aPtr": (*corejson.Result)(nil),
			"bPtr": (*corejson.Result)(nil),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEqualPtr returns false -- one nil",
		ArrangeInput: args.Map{
			"aPtr": corejson.NewPtr(map[string]string{"k": "v"}),
			"bPtr": (*corejson.Result)(nil),
		},
		ExpectedInput: "false",
	},
}
