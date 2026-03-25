package corejsontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var newValidTestCase = coretestcases.CaseV1{
	Title: "New returns bytes without error -- valid struct input",
	ExpectedInput: args.Map{
		"hasError":    false,
		"isEmpty":     false,
		"hasBytes":    true,
		"hasTypeName": true,
	},
}

var newNilTestCase = coretestcases.CaseV1{
	Title: "New returns null bytes -- nil input",
	ExpectedInput: args.Map{
		"hasError":     false,
		"bytesContent": "null",
	},
}

var newChannelTestCase = coretestcases.CaseV1{
	Title: "New returns error -- channel input",
	ExpectedInput: args.Map{
		"hasError":             true,
		"errorContainsMarshal": true,
	},
}

var newPtrValidTestCase = coretestcases.CaseV1{
	Title: "NewPtr returns non-nil result -- valid struct input",
	ExpectedInput: args.Map{
		"isNonNil": true,
		"hasError": false,
		"isEmpty":  false,
		"hasBytes": true,
	},
}

var newPtrNilTestCase = coretestcases.CaseV1{
	Title: "NewPtr returns null bytes -- nil input",
	ExpectedInput: args.Map{
		"isNonNil":     true,
		"hasError":     false,
		"bytesContent": "null",
	},
}

var newPtrChannelTestCase = coretestcases.CaseV1{
	Title: "NewPtr returns error -- channel input",
	ExpectedInput: args.Map{
		"isNonNil":             true,
		"hasError":             true,
		"errorContainsMarshal": true,
	},
}
