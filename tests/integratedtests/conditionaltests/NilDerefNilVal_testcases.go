package conditionaltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var nilDerefStringTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDeref with nil string pointer returns empty string",
		ArrangeInput: args.Map{
			"when":  "given nil string pointer",
			"isNil": true,
		},
		ExpectedInput: []string{""},
	},
	{
		Title: "NilDeref with non-nil string pointer returns value",
		ArrangeInput: args.Map{
			"when":  "given non-nil string pointer",
			"isNil": false,
			"value": "hello",
		},
		ExpectedInput: []string{"hello"},
	},
	{
		Title: "NilDeref with non-nil empty string pointer returns empty",
		ArrangeInput: args.Map{
			"when":  "given non-nil pointer to empty string",
			"isNil": false,
			"value": "",
		},
		ExpectedInput: []string{""},
	},
}

var nilDerefIntTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDeref with nil int pointer returns 0",
		ArrangeInput: args.Map{
			"when":  "given nil int pointer",
			"isNil": true,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "NilDeref with non-nil int pointer returns value",
		ArrangeInput: args.Map{
			"when":  "given non-nil int pointer",
			"isNil": false,
			"value": 42,
		},
		ExpectedInput: []string{"42"},
	},
	{
		Title: "NilDeref with non-nil zero int pointer returns 0",
		ArrangeInput: args.Map{
			"when":  "given non-nil pointer to zero",
			"isNil": false,
			"value": 0,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "NilDeref with non-nil negative int returns negative",
		ArrangeInput: args.Map{
			"when":  "given non-nil pointer to negative",
			"isNil": false,
			"value": -7,
		},
		ExpectedInput: []string{"-7"},
	},
}

var nilDerefBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDeref with nil bool pointer returns false",
		ArrangeInput: args.Map{
			"when":  "given nil bool pointer",
			"isNil": true,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "NilDeref with non-nil true bool pointer returns true",
		ArrangeInput: args.Map{
			"when":  "given non-nil true pointer",
			"isNil": false,
			"value": true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "NilDeref with non-nil false bool pointer returns false",
		ArrangeInput: args.Map{
			"when":  "given non-nil false pointer",
			"isNil": false,
			"value": false,
		},
		ExpectedInput: []string{"false"},
	},
}

var nilDerefPtrStringTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDerefPtr with nil string pointer returns pointer to empty",
		ArrangeInput: args.Map{
			"when":  "given nil string pointer",
			"isNil": true,
		},
		ExpectedInput: []string{"true", ""},
	},
	{
		Title: "NilDerefPtr with non-nil string pointer returns same value",
		ArrangeInput: args.Map{
			"when":  "given non-nil string pointer",
			"isNil": false,
			"value": "world",
		},
		ExpectedInput: []string{"true", "world"},
	},
}

var nilDerefPtrIntTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDerefPtr with nil int pointer returns pointer to 0",
		ArrangeInput: args.Map{
			"when":  "given nil int pointer",
			"isNil": true,
		},
		ExpectedInput: []string{"true", "0"},
	},
	{
		Title: "NilDerefPtr with non-nil int pointer returns same value",
		ArrangeInput: args.Map{
			"when":  "given non-nil int pointer",
			"isNil": false,
			"value": 99,
		},
		ExpectedInput: []string{"true", "99"},
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
		ExpectedInput: []string{"was-nil"},
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
		ExpectedInput: []string{"was-set"},
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
		ExpectedInput: []string{"was-set"},
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
		ExpectedInput: []string{"-1"},
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
		ExpectedInput: []string{"1"},
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
		ExpectedInput: []string{"true", "nil-label"},
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
		ExpectedInput: []string{"true", "set-label"},
	},
}
