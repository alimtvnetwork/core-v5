package regexnewtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// Empty pattern edge cases
// =============================================================================

var emptyPatternEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyRegex returns undefined -- empty pattern",
		ExpectedInput: args.Map{
			"isUndefined": true,
		},
	},
	{
		Title: "LazyRegex returns not applicable -- empty pattern",
		ExpectedInput: args.Map{
			"isApplicable": false,
		},
	},
	{
		Title: "IsMatch returns false -- empty pattern",
		ExpectedInput: args.Map{
			"isMatch": false,
		},
	},
	{
		Title: "IsFailedMatch returns true -- empty pattern",
		ExpectedInput: args.Map{
			"isFailedMatch": true,
		},
	},
	{
		Title: "Compile returns error and nil regex -- empty pattern",
		ExpectedInput: args.Map{
			"hasError":   true,
			"regexIsNil": true,
		},
	},
}
