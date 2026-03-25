package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var partialRwxVerifyTestCases = []coretestcases.CaseV1{
	{
		Title: "Same input returns true.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx-*-r*x",
			"fullRwx":    "-rwx-*-r*x",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Same [-rwx---r*x] comparing " +
			"with [-rwx-*-r*x] returns false.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx---r*x",
			"fullRwx":    "-rwx-*-r*x",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "Same [-rwx-*-r*x] comparing with " +
			"[-rwx-w-r*x] returns true.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx-*-r*x",
			"fullRwx":    "-rwx-w-r*x",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Same [-rwx-*-] or [-rwx-*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns true.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx-*-",
			"fullRwx":    "-rwx-w--*x",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Same [-rwxr*-] or [-rwxr*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns false.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwxr*-",
			"fullRwx":    "-rwx-w--*x",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}
