package corejsontests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var newValidTestCase = coretestcases.CaseV1{
	Title:         "New - valid struct produces bytes no error",
	ExpectedInput: []string{"false", "false", "true", "true"},
}

var newNilTestCase = coretestcases.CaseV1{
	Title:         "New - nil input produces null bytes",
	ExpectedInput: []string{"false", "null"},
}

var newChannelTestCase = coretestcases.CaseV1{
	Title:         "New - channel produces error",
	ExpectedInput: []string{"true", "true"},
}

var newPtrValidTestCase = coretestcases.CaseV1{
	Title:         "NewPtr - valid struct produces non-nil result",
	ExpectedInput: []string{"true", "false", "false", "true"},
}

var newPtrNilTestCase = coretestcases.CaseV1{
	Title:         "NewPtr - nil input produces null bytes",
	ExpectedInput: []string{"true", "false", "null"},
}

var newPtrChannelTestCase = coretestcases.CaseV1{
	Title:         "NewPtr - channel produces error",
	ExpectedInput: []string{"true", "true", "true"},
}
