package corejsontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var newValidTestCase = coretestcases.CaseV1{
	Title: "New - valid struct produces bytes no error",
	ExpectedInput: args.Map{
		"hasError":    false,
		"isEmpty":     false,
		"hasBytes":    true,
		"hasTypeName": true,
	},
}

var newNilTestCase = coretestcases.CaseV1{
	Title: "New - nil input produces null bytes",
	ExpectedInput: args.Map{
		"hasError":     false,
		"bytesContent": "null",
	},
}

var newChannelTestCase = coretestcases.CaseV1{
	Title: "New - channel produces error",
	ExpectedInput: args.Map{
		"hasError":             true,
		"errorContainsMarshal": true,
	},
}

var newPtrValidTestCase = coretestcases.CaseV1{
	Title: "NewPtr - valid struct produces non-nil result",
	ExpectedInput: args.Map{
		"isNonNil": true,
		"hasError": false,
		"isEmpty":  false,
		"hasBytes": true,
	},
}

var newPtrNilTestCase = coretestcases.CaseV1{
	Title: "NewPtr - nil input produces null bytes",
	ExpectedInput: args.Map{
		"isNonNil":     true,
		"hasError":     false,
		"bytesContent": "null",
	},
}

var newPtrChannelTestCase = coretestcases.CaseV1{
	Title: "NewPtr - channel produces error",
	ExpectedInput: args.Map{
		"isNonNil":             true,
		"hasError":             true,
		"errorContainsMarshal": true,
	},
}
