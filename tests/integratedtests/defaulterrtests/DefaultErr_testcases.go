package defaulterrtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var defaultErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Marshalling error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking Marshalling error",
			"error": "Marshalling",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "UnMarshalling error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking UnMarshalling error",
			"error": "UnMarshalling",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "OutOfRange error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking OutOfRange error",
			"error": "OutOfRange",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "CannotProcessNilOrEmpty error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking CannotProcessNilOrEmpty error",
			"error": "CannotProcessNilOrEmpty",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "NegativeDataCannotProcess error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking NegativeDataCannotProcess error",
			"error": "NegativeDataCannotProcess",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "NilResult error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking NilResult error",
			"error": "NilResult",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "UnexpectedValue error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking UnexpectedValue error",
			"error": "UnexpectedValue",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "CannotRemoveFromEmptyCollection error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking CannotRemoveFromEmptyCollection error",
			"error": "CannotRemoveFromEmptyCollection",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "MarshallingFailedDueToNilOrEmpty error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking MarshallingFailedDueToNilOrEmpty error",
			"error": "MarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "UnmarshallingFailedDueToNilOrEmpty error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking UnmarshallingFailedDueToNilOrEmpty error",
			"error": "UnmarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
	{
		Title: "KeyNotExistInMap error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking KeyNotExistInMap error",
			"error": "KeyNotExistInMap",
		},
		ExpectedInput: args.Map{"isNotNil": true, "hasMessage": true},
	},
}
