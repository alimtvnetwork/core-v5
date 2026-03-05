package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// IsMatch test cases
// =============================================================================

var stringCompareIsMatchTestCases = []coretestcases.CaseV1{
	// === Equal ===
	{
		Title: "Equal - identical strings should match",
		ArrangeInput: args.Map{
			"when":         "given equal strings",
			"method":       "equal",
			"search":       "hello",
			"content":      "hello",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isMatch
			Second: "false", // isMatchFailed
		},
	},
	{
		Title: "Equal - different strings should not match",
		ArrangeInput: args.Map{
			"when":         "given different strings",
			"method":       "equal",
			"search":       "hello",
			"content":      "world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isMatch
			Second: "true",  // isMatchFailed
		},
	},
	{
		Title: "Equal - case-sensitive should not match different cases",
		ArrangeInput: args.Map{
			"when":         "given different casing without ignore",
			"method":       "equal",
			"search":       "Hello",
			"content":      "hello",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isMatch
			Second: "true",  // isMatchFailed
		},
	},

	// === Contains ===
	{
		Title: "Contains - substring should match",
		ArrangeInput: args.Map{
			"when":         "given content containing search",
			"method":       "contains",
			"search":       "world",
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isMatch
			Second: "false", // isMatchFailed
		},
	},
	{
		Title: "Contains - missing substring should not match",
		ArrangeInput: args.Map{
			"when":         "given content not containing search",
			"method":       "contains",
			"search":       "xyz",
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isMatch
			Second: "true",  // isMatchFailed
		},
	},

	// === StartsWith ===
	{
		Title: "StartsWith - matching prefix should match",
		ArrangeInput: args.Map{
			"when":         "given content starting with search",
			"method":       "startsWith",
			"search":       "hello",
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isMatch
			Second: "false", // isMatchFailed
		},
	},
	{
		Title: "StartsWith - non-prefix should not match",
		ArrangeInput: args.Map{
			"when":         "given content not starting with search",
			"method":       "startsWith",
			"search":       "world",
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isMatch
			Second: "true",  // isMatchFailed
		},
	},
	{
		Title: "StartsWith - ignore case should match",
		ArrangeInput: args.Map{
			"when":         "given different casing with ignore case",
			"method":       "startsWith",
			"search":       "HELLO",
			"content":      "hello world",
			"isIgnoreCase": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isMatch
			Second: "false", // isMatchFailed
		},
	},

	// === EndsWith ===
	{
		Title: "EndsWith - matching suffix should match",
		ArrangeInput: args.Map{
			"when":         "given content ending with search",
			"method":       "endsWith",
			"search":       "world",
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isMatch
			Second: "false", // isMatchFailed
		},
	},
	{
		Title: "EndsWith - non-suffix should not match",
		ArrangeInput: args.Map{
			"when":         "given content not ending with search",
			"method":       "endsWith",
			"search":       "hello",
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isMatch
			Second: "true",  // isMatchFailed
		},
	},
	{
		Title: "EndsWith - ignore case should match",
		ArrangeInput: args.Map{
			"when":         "given different casing with ignore case",
			"method":       "endsWith",
			"search":       "WORLD",
			"content":      "hello world",
			"isIgnoreCase": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isMatch
			Second: "false", // isMatchFailed
		},
	},

	// === Regex ===
	{
		Title: "Regex - matching pattern should match",
		ArrangeInput: args.Map{
			"when":         "given content matching regex",
			"method":       "regex",
			"search":       `^hello\s\w+$`,
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isMatch
			Second: "false", // isMatchFailed
		},
	},
	{
		Title: "Regex - non-matching pattern should not match",
		ArrangeInput: args.Map{
			"when":         "given content not matching regex",
			"method":       "regex",
			"search":       `^\d+$`,
			"content":      "hello",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isMatch
			Second: "true",  // isMatchFailed
		},
	},
}

// =============================================================================
// VerifyError test cases
// =============================================================================

var stringCompareVerifyErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "VerifyError - equal match returns nil",
		ArrangeInput: args.Map{
			"when":         "given matching equal strings",
			"method":       "equal",
			"search":       "hello",
			"content":      "hello",
			"isIgnoreCase": false,
		},
		ExpectedInput: "false",
	},
	{
		Title: "VerifyError - equal mismatch returns error",
		ArrangeInput: args.Map{
			"when":         "given non-matching equal strings",
			"method":       "equal",
			"search":       "hello",
			"content":      "world",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "VerifyError - contains mismatch returns error",
		ArrangeInput: args.Map{
			"when":         "given content not containing search",
			"method":       "contains",
			"search":       "xyz",
			"content":      "hello world",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "VerifyError - valid regex match returns nil",
		ArrangeInput: args.Map{
			"when":         "given content matching regex",
			"method":       "regex",
			"search":       `^\d+$`,
			"content":      "12345",
			"isIgnoreCase": false,
		},
		ExpectedInput: "false",
	},
	{
		Title: "VerifyError - valid regex no match returns error",
		ArrangeInput: args.Map{
			"when":         "given content not matching regex",
			"method":       "regex",
			"search":       `^\d+$`,
			"content":      "hello",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "VerifyError - invalid regex returns error",
		ArrangeInput: args.Map{
			"when":         "given invalid regex pattern",
			"method":       "regex",
			"search":       `[invalid(`,
			"content":      "content",
			"isIgnoreCase": false,
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// Nil receiver test cases
// =============================================================================

var stringCompareNilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil receiver - IsMatch returns true (vacuous truth)",
		ArrangeInput: args.Map{
			"when":   "given nil StringCompare",
			"method": "IsMatch",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Nil receiver - IsMatchFailed returns false",
		ArrangeInput: args.Map{
			"when":   "given nil StringCompare",
			"method": "IsMatchFailed",
		},
		ExpectedInput: "false",
	},
	{
		Title: "Nil receiver - IsInvalid returns true",
		ArrangeInput: args.Map{
			"when":   "given nil StringCompare",
			"method": "IsInvalid",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Nil receiver - IsDefined returns false",
		ArrangeInput: args.Map{
			"when":   "given nil StringCompare",
			"method": "IsDefined",
		},
		ExpectedInput: "false",
	},
	{
		Title: "Nil receiver - VerifyError returns nil",
		ArrangeInput: args.Map{
			"when":   "given nil StringCompare",
			"method": "VerifyError",
		},
		ExpectedInput: "false",
	},
}
