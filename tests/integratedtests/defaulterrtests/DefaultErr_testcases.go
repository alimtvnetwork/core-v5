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
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "UnMarshalling error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking UnMarshalling error",
			"error": "UnMarshalling",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "OutOfRange error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking OutOfRange error",
			"error": "OutOfRange",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "CannotProcessNilOrEmpty error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking CannotProcessNilOrEmpty error",
			"error": "CannotProcessNilOrEmpty",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "NegativeDataCannotProcess error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking NegativeDataCannotProcess error",
			"error": "NegativeDataCannotProcess",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "NilResult error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking NilResult error",
			"error": "NilResult",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "UnexpectedValue error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking UnexpectedValue error",
			"error": "UnexpectedValue",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "CannotRemoveFromEmptyCollection error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking CannotRemoveFromEmptyCollection error",
			"error": "CannotRemoveFromEmptyCollection",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "MarshallingFailedDueToNilOrEmpty error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking MarshallingFailedDueToNilOrEmpty error",
			"error": "MarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "UnmarshallingFailedDueToNilOrEmpty error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking UnmarshallingFailedDueToNilOrEmpty error",
			"error": "UnmarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "KeyNotExistInMap error is not nil",
		ArrangeInput: args.Map{
			"when":  "checking KeyNotExistInMap error",
			"error": "KeyNotExistInMap",
		},
		ExpectedInput: []string{"true", "true"},
	},
}
