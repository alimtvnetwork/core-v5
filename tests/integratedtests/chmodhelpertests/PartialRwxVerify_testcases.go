package chmodhelpertests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var partialRwxVerifyTestCases = []coretestcases.CaseV1{
	{
		Title: "Same input returns true.",
		ArrangeInput: map[string]string{
			"partialRwx": "-rwx-*-r*x",
			"fullRwx":    "-rwx-*-r*x",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Same [-rwx---r*x] comparing " +
			"with [-rwx-*-r*x] returns false.",
		ArrangeInput: map[string]string{
			"partialRwx": "-rwx---r*x",
			"fullRwx":    "-rwx-*-r*x",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Same [-rwx-*-r*x] comparing with " +
			"[-rwx-w-r*x] returns true.",
		ArrangeInput: map[string]string{
			"partialRwx": "-rwx-*-r*x",
			"fullRwx":    "-rwx-w-r*x",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Same [-rwx-*-] or [-rwx-*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns true.",
		ArrangeInput: map[string]string{
			"partialRwx": "-rwx-*-",
			"fullRwx":    "-rwx-w--*x",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Same [-rwxr*-] or [-rwxr*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns false.",
		ArrangeInput: map[string]string{
			"partialRwx": "-rwxr*-",
			"fullRwx":    "-rwx-w--*x",
		},
		ExpectedInput: []string{
			"false",
		},
	},
}
