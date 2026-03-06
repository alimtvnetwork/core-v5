package conditionaltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var valueOrZeroStringTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueOrZero with nil string pointer returns empty string",
		ArrangeInput: args.Map{
			"when":  "given nil string pointer",
			"isNil": true,
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "ValueOrZero with non-nil string pointer returns value",
		ArrangeInput: args.Map{
			"when":  "given non-nil string pointer",
			"isNil": false,
			"value": "hello",
		},
		ExpectedInput: "hello",
	},
	{
		Title: "ValueOrZero with non-nil empty string pointer returns empty",
		ArrangeInput: args.Map{
			"when":  "given non-nil pointer to empty string",
			"isNil": false,
			"value": "",
		},
		ExpectedInput: []string{""},
	},
}

var valueOrZeroIntTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueOrZero with nil int pointer returns 0",
		ArrangeInput: args.Map{
			"when":  "given nil int pointer",
			"isNil": true,
		},
		ExpectedInput: "0",
	},
	{
		Title: "ValueOrZero with non-nil int pointer returns value",
		ArrangeInput: args.Map{
			"when":  "given non-nil int pointer",
			"isNil": false,
			"value": 42,
		},
		ExpectedInput: "42",
	},
	{
		Title: "ValueOrZero with non-nil zero int pointer returns 0",
		ArrangeInput: args.Map{
			"when":  "given non-nil pointer to zero",
			"isNil": false,
			"value": 0,
		},
		ExpectedInput: "0",
	},
	{
		Title: "ValueOrZero with non-nil negative int returns negative",
		ArrangeInput: args.Map{
			"when":  "given non-nil pointer to negative",
			"isNil": false,
			"value": -7,
		},
		ExpectedInput: "-7",
	},
}

var valueOrZeroBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "ValueOrZero with nil bool pointer returns false",
		ArrangeInput: args.Map{
			"when":  "given nil bool pointer",
			"isNil": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "ValueOrZero with non-nil true bool pointer returns true",
		ArrangeInput: args.Map{
			"when":  "given non-nil true pointer",
			"isNil": false,
			"value": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "ValueOrZero with non-nil false bool pointer returns false",
		ArrangeInput: args.Map{
			"when":  "given non-nil false pointer",
			"isNil": false,
			"value": false,
		},
		ExpectedInput: "false",
	},
}

var ptrOrZeroStringTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrOrZero with nil string pointer returns pointer to empty",
		ArrangeInput: args.Map{
			"when":  "given nil string pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "",
		},
	},
	{
		Title: "PtrOrZero with non-nil string pointer returns same value",
		ArrangeInput: args.Map{
			"when":  "given non-nil string pointer",
			"isNil": false,
			"value": "world",
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "world",
		},
	},
}

var ptrOrZeroIntTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrOrZero with nil int pointer returns pointer to 0",
		ArrangeInput: args.Map{
			"when":  "given nil int pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "0",
		},
	},
	{
		Title: "PtrOrZero with non-nil int pointer returns same value",
		ArrangeInput: args.Map{
			"when":  "given non-nil int pointer",
			"isNil": false,
			"value": 99,
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "99",
		},
	},
}

var nilValStringTestCases = []coretestcases.CaseV1{
	{
		Title: "NilVal with nil pointer returns onNil",
		ArrangeInput: args.Map{
			"when":     "given nil string pointer",
			"isNil":    true,
			"onNil":    "was-nil",
			"onNonNil": "was-set",
		},
		ExpectedInput: "was-nil",
	},
	{
		Title: "NilVal with non-nil pointer returns onNonNil",
		ArrangeInput: args.Map{
			"when":     "given non-nil string pointer",
			"isNil":    false,
			"value":    "anything",
			"onNil":    "was-nil",
			"onNonNil": "was-set",
		},
		ExpectedInput: "was-set",
	},
	{
		Title: "NilVal with non-nil empty string pointer returns onNonNil",
		ArrangeInput: args.Map{
			"when":     "given non-nil pointer to empty string",
			"isNil":    false,
			"value":    "",
			"onNil":    "was-nil",
			"onNonNil": "was-set",
		},
		ExpectedInput: "was-set",
	},
}

var nilValIntTestCases = []coretestcases.CaseV1{
	{
		Title: "NilVal int with nil pointer returns onNil",
		ArrangeInput: args.Map{
			"when":     "given nil int pointer",
			"isNil":    true,
			"onNil":    -1,
			"onNonNil": 1,
		},
		ExpectedInput: "-1",
	},
	{
		Title: "NilVal int with non-nil pointer returns onNonNil",
		ArrangeInput: args.Map{
			"when":     "given non-nil int pointer",
			"isNil":    false,
			"value":    50,
			"onNil":    -1,
			"onNonNil": 1,
		},
		ExpectedInput: "1",
	},
}

var nilValPtrStringTestCases = []coretestcases.CaseV1{
	{
		Title: "NilValPtr with nil pointer returns pointer to onNil",
		ArrangeInput: args.Map{
			"when":     "given nil string pointer",
			"isNil":    true,
			"onNil":    "nil-label",
			"onNonNil": "set-label",
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "nil-label",
		},
	},
	{
		Title: "NilValPtr with non-nil pointer returns pointer to onNonNil",
		ArrangeInput: args.Map{
			"when":     "given non-nil string pointer",
			"isNil":    false,
			"value":    "something",
			"onNil":    "nil-label",
			"onNonNil": "set-label",
		},
		ExpectedInput: args.Map{
			"isNotNil": "true",
			"value":    "set-label",
		},
	},
}
