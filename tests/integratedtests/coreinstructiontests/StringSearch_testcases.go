package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// IsMatch test cases
// =============================================================================

var stringSearchIsMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatch - equal match returns true",
		ArrangeInput: args.Map{
			"when":    "given matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "hello",
		},
		ExpectedInput: args.Map{
			"isMatch":       true,
			"isMatchFailed": false,
		},
	},
	{
		Title: "IsMatch - equal no match returns false",
		ArrangeInput: args.Map{
			"when":    "given non-matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "world",
		},
		ExpectedInput: args.Map{
			"isMatch":       false,
			"isMatchFailed": true,
		},
	},
	{
		Title: "IsMatch - contains match returns true",
		ArrangeInput: args.Map{
			"when":    "given content containing search",
			"method":  "contains",
			"search":  "world",
			"content": "hello world",
		},
		ExpectedInput: args.Map{
			"isMatch":       true,
			"isMatchFailed": false,
		},
	},
	{
		Title: "IsMatch - contains no match returns false",
		ArrangeInput: args.Map{
			"when":    "given content not containing search",
			"method":  "contains",
			"search":  "xyz",
			"content": "hello world",
		},
		ExpectedInput: args.Map{
			"isMatch":       false,
			"isMatchFailed": true,
		},
	},
}

// =============================================================================
// IsAllMatch test cases
// =============================================================================

var stringSearchIsAllMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "IsAllMatch - all contents match returns true",
		ArrangeInput: args.Map{
			"when":     "given all contents containing search",
			"method":   "contains",
			"search":   "o",
			"contents": []string{"hello", "world", "foo"},
		},
		ExpectedInput: args.Map{
			"isAllMatch":       true,
			"isAnyMatchFailed": false,
		},
	},
	{
		Title: "IsAllMatch - one content fails returns false",
		ArrangeInput: args.Map{
			"when":     "given one content not containing search",
			"method":   "contains",
			"search":   "z",
			"contents": []string{"hello", "buzz", "world"},
		},
		ExpectedInput: args.Map{
			"isAllMatch":       false,
			"isAnyMatchFailed": true,
		},
	},
	{
		Title: "IsAllMatch - empty contents returns true",
		ArrangeInput: args.Map{
			"when":     "given empty contents",
			"method":   "equal",
			"search":   "hello",
			"contents": []string{},
		},
		ExpectedInput: args.Map{
			"isAllMatch":       true,
			"isAnyMatchFailed": false,
		},
	},
}

// =============================================================================
// IsEmpty / IsExist / Has test cases
// =============================================================================

var stringSearchStateTestCases = []coretestcases.CaseV1{
	{
		Title: "Non-nil - IsEmpty false, IsExist true, Has true",
		ArrangeInput: args.Map{
			"when":   "given non-nil StringSearch",
			"method": "equal",
			"search": "test",
			"isNil":  false,
		},
		ExpectedInput: args.Map{
			"isEmpty": false,
			"isExist": true,
			"has":     true,
		},
	},
	{
		Title: "Nil - IsEmpty true, IsExist false, Has false",
		ArrangeInput: args.Map{
			"when":  "given nil StringSearch",
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isEmpty": true,
			"isExist": false,
			"has":     false,
		},
	},
}

// =============================================================================
// VerifyError test cases
// =============================================================================

var stringSearchVerifyErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "VerifyError - match returns nil",
		ArrangeInput: args.Map{
			"when":    "given matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "hello",
			"isNil":   false,
		},
		ExpectedInput: "false",
	},
	{
		Title: "VerifyError - no match returns error",
		ArrangeInput: args.Map{
			"when":    "given non-matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "world",
			"isNil":   false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "VerifyError - nil receiver returns nil",
		ArrangeInput: args.Map{
			"when":    "given nil StringSearch",
			"content": "anything",
			"isNil":   true,
		},
		ExpectedInput: "false",
	},
}
