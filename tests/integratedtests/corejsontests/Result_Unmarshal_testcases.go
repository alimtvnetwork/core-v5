package corejsontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var resultUnmarshalValidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - valid JSON deserializes correctly",
	ExpectedInput: args.Map{
		"error":            "<nil>",
		"deserializedName": "Alice",
		"deserializedAge":  "30",
	},
}

var resultUnmarshalNilTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - nil receiver returns error",
	ExpectedInput: args.Map{
		"hasError":          true,
		"errorContainsNull": true,
	},
}

var resultUnmarshalInvalidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - invalid bytes returns error",
	ExpectedInput: args.Map{
		"hasError":               true,
		"errorContainsUnmarshal": true,
	},
}

var resultUnmarshalExistingErrorTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - existing error propagates",
	ExpectedInput: args.Map{
		"hasError":               true,
		"errorContainsUnmarshal": true,
	},
}
