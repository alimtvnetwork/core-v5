package corejsontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var resultUnmarshalValidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - valid JSON deserializes correctly",
	ExpectedInput: args.Three[string, string, string]{
		First:  "<nil>",  // error
		Second: "Alice",  // deserializedName
		Third:  "30",     // deserializedAge
	},
}

var resultUnmarshalNilTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - nil receiver returns error",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // hasError
		Second: "true", // errorContainsNull
	},
}

var resultUnmarshalInvalidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - invalid bytes returns error",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // hasError
		Second: "true", // errorContainsUnmarshal
	},
}

var resultUnmarshalExistingErrorTestCase = coretestcases.CaseV1{
	Title: "Unmarshal - existing error propagates",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // hasError
		Second: "true", // errorContainsUnmarshal
	},
}
