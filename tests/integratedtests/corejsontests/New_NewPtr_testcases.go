package corejsontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var newValidTestCases = []coretestcases.CaseV1{
	{
		Title:         "New - valid struct produces bytes no error",
		ExpectedInput: []string{"false", "false", "true", "true"},
	},
}

var newNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "New - nil input produces null bytes",
		ExpectedInput: []string{"false", "null"},
	},
}

var newChannelTestCases = []coretestcases.CaseV1{
	{
		Title:         "New - channel produces error",
		ExpectedInput: []string{"true", "true"},
	},
}

var newPtrValidTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewPtr - valid struct produces non-nil result",
		ExpectedInput: []string{"true", "false", "false", "true"},
	},
}

var newPtrNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewPtr - nil input produces null bytes",
		ExpectedInput: []string{"true", "false", "null"},
	},
}

var newPtrChannelTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewPtr - channel produces error",
		ExpectedInput: []string{"true", "true", "true"},
	},
}
