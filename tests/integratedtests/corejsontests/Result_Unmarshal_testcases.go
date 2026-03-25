package corejsontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var resultUnmarshalValidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns deserialized struct -- valid JSON input",
	ExpectedInput: args.Map{
		"error":            "<nil>",
		"deserializedName": "Alice",
		"deserializedAge":  "30",
	},
}

var resultUnmarshalNilTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns error -- nil receiver",
	ExpectedInput: args.Map{
		"hasError":          true,
		"errorContainsNull": true,
	},
}

var resultUnmarshalInvalidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns error -- invalid bytes input",
	ExpectedInput: args.Map{
		"hasError":               true,
		"errorContainsUnmarshal": true,
	},
}

var resultUnmarshalExistingErrorTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns propagated error -- existing error on result",
	ExpectedInput: args.Map{
		"hasError":               true,
		"errorContainsUnmarshal": true,
	},
}
