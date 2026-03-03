package corejsontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var resultUnmarshalValidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - valid JSON deserializes correctly",
	ExpectedInput: []string{
		"<nil>",
		"Alice",
		"30",
	},
}

var resultUnmarshalNilTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - nil receiver returns error",
	ExpectedInput: []string{
		"true",
		"true",
	},
}

var resultUnmarshalInvalidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - invalid bytes returns error",
	ExpectedInput: []string{
		"true",
		"true",
	},
}

var resultUnmarshalExistingErrorTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - existing error propagates",
	ExpectedInput: []string{
		"true",
		"true",
	},
}
