package regexnewtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// Empty pattern edge cases
// =============================================================================

var emptyPatternEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty pattern LazyRegex is undefined",
		ExpectedInput: args.Map{
			"isUndefined": true,
		},
	},
	{
		Title: "Empty pattern LazyRegex is not applicable",
		ExpectedInput: args.Map{
			"isApplicable": false,
		},
	},
	{
		Title: "Empty pattern IsMatch returns false",
		ExpectedInput: args.Map{
			"isMatch": false,
		},
	},
	{
		Title: "Empty pattern IsFailedMatch returns true",
		ExpectedInput: args.Map{
			"isFailedMatch": true,
		},
	},
	{
		Title: "Empty pattern Compile returns error and nil regex",
		ExpectedInput: args.Map{
			"hasError":    true,
			"regexIsNil": true,
		},
	},
}
