package iserrortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// ExitError
// ==========================================================================

var extExitErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "ExitError returns false -- given standard error",
		ArrangeInput: args.Map{
			"when": "given standard error",
		},
		ExpectedInput: args.Map{
			"isExitError": false,
		},
	},
}
