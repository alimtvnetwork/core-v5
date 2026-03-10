package coreinstructiontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// IsMatch test cases
// =============================================================================

var stringSearchIsMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatch returns true -- equal match, search 'hello' in content 'hello'",
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
		Title: "IsMatch returns false -- equal no match, search 'hello' in content 'world'",
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
		Title: "IsMatch returns true -- contains match, search 'world' in 'hello world'",
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
		Title: "IsMatch returns false -- contains no match, search 'xyz' in 'hello world'",
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
		Title: "IsAllMatch returns true -- all contents contain 'o' in ['hello','world','foo']",
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
		Title: "IsAllMatch returns false -- 'world' missing 'z' in ['hello','buzz','world']",
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
		Title: "IsAllMatch returns true -- empty contents slice",
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
		Title: "StringSearch returns isEmpty false, isExist true -- non-nil instance",
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
		Title: "StringSearch returns isEmpty true, isExist false -- nil instance",
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
		Title: "VerifyError returns nil -- equal match, search 'hello' in 'hello'",
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
		Title: "VerifyError returns error -- equal no match, search 'hello' in 'world'",
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
		Title: "VerifyError returns nil -- nil receiver, content 'anything'",
		ArrangeInput: args.Map{
			"when":    "given nil StringSearch",
			"content": "anything",
			"isNil":   true,
		},
		ExpectedInput: "false",
	},
}
