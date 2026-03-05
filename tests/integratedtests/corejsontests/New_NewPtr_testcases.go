package corejsontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var newValidTestCase = coretestcases.CaseV1{
	Title: "New - valid struct produces bytes no error",
	ExpectedInput: args.Four[string, string, string, string]{
		First:  "false", // hasError
		Second: "false", // isEmpty
		Third:  "true",  // hasBytes
		Fourth: "true",  // hasTypeName
	},
}

var newNilTestCase = coretestcases.CaseV1{
	Title: "New - nil input produces null bytes",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // hasError
		Second: "null",  // bytesContent
	},
}

var newChannelTestCase = coretestcases.CaseV1{
	Title: "New - channel produces error",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // hasError
		Second: "true", // errorContainsMarshal
	},
}

var newPtrValidTestCase = coretestcases.CaseV1{
	Title: "NewPtr - valid struct produces non-nil result",
	ExpectedInput: args.Four[string, string, string, string]{
		First:  "true",  // isNonNil
		Second: "false", // hasError
		Third:  "false", // isEmpty
		Fourth: "true",  // hasBytes
	},
}

var newPtrNilTestCase = coretestcases.CaseV1{
	Title: "NewPtr - nil input produces null bytes",
	ExpectedInput: args.Three[string, string, string]{
		First:  "true",  // isNonNil
		Second: "false", // hasError
		Third:  "null",  // bytesContent
	},
}

var newPtrChannelTestCase = coretestcases.CaseV1{
	Title: "NewPtr - channel produces error",
	ExpectedInput: args.Three[string, string, string]{
		First:  "true", // isNonNil
		Second: "true", // hasError
		Third:  "true", // errorContainsMarshal
	},
}
