package corejsontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var resultUnmarshalValidTestCases = []coretestcases.CaseV1{
	{
		Title: "Unmarshal - valid JSON deserializes correctly",
		ExpectedInput: []string{
			"<nil>",
			"Alice",
			"30",
		},
	},
}

var resultUnmarshalNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Unmarshal - nil receiver returns error",
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}

var resultUnmarshalInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "Unmarshal - invalid bytes returns error",
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}

var resultUnmarshalExistingErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Unmarshal - existing error propagates",
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}
