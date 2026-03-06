package coretestcasestests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// --- IsFailedToMatch ---

var isFailedToMatchWhenMatchingTestCase = coretestcases.StringBoolGherkins{
	Title:      "IsFailedToMatch returns false when IsMatching is true",
	When:       "IsMatching is true",
	IsMatching: true,
	Expected:   false,
}

var isFailedToMatchWhenNotMatchingTestCase = coretestcases.StringBoolGherkins{
	Title:      "IsFailedToMatch returns true when IsMatching is false",
	When:       "IsMatching is false",
	IsMatching: false,
	Expected:   true,
}

// --- CompareWith: equal ---

var compareWithEqualTestCase = coretestcases.StringBoolGherkins{
	Title:    "CompareWith returns true for identical structs",
	When:     "two structs are identical",
	Expected: true,
	ExtraArgs: args.Map{
		"a": &coretestcases.StringBoolGherkins{
			Title: "same",
			Input: "hello",
		},
		"b": &coretestcases.StringBoolGherkins{
			Title: "same",
			Input: "hello",
		},
		"expectedDiff": "",
	},
}

// --- CompareWith: different title ---

var compareWithDiffTitleTestCase = coretestcases.StringBoolGherkins{
	Title:    "CompareWith returns false for different Title",
	When:     "titles differ",
	Expected: false,
	ExtraArgs: args.Map{
		"a": &coretestcases.StringBoolGherkins{
			Title: "A",
			Input: "hello",
		},
		"b": &coretestcases.StringBoolGherkins{
			Title: "B",
			Input: "hello",
		},
		"expectedDiff": `Title: "A" != "B"`,
	},
}

// --- CompareWith: different input ---

var compareWithDiffInputTestCase = coretestcases.StringBoolGherkins{
	Title:    "CompareWith returns false for different Input",
	When:     "inputs differ",
	Expected: false,
	ExtraArgs: args.Map{
		"a": &coretestcases.StringBoolGherkins{
			Title: "same",
			Input: "alpha",
		},
		"b": &coretestcases.StringBoolGherkins{
			Title: "same",
			Input: "beta",
		},
		"expectedDiff": "Input: alpha != beta",
	},
}

// --- CompareWith: nil both ---

var compareWithBothNilTestCase = coretestcases.StringBoolGherkins{
	Title:    "CompareWith returns true when both nil",
	When:     "both pointers are nil",
	Expected: true,
	ExtraArgs: args.Map{
		"bothNil":      true,
		"expectedDiff": "",
	},
}

// --- CompareWith: one nil ---

var compareWithOneNilTestCase = coretestcases.StringBoolGherkins{
	Title:    "CompareWith returns false when one is nil",
	When:     "only one pointer is nil",
	Expected: false,
	ExtraArgs: args.Map{
		"a":            &coretestcases.StringBoolGherkins{Title: "exists"},
		"expectedDiff": "one side is nil",
	},
}

// --- FullString ---

var fullStringBasicTestCase = coretestcases.StringBoolGherkins{
	Title:      "FullString includes all fields",
	When:       "struct has all fields populated",
	IsMatching: true,
	ExtraArgs: args.Map{
		"subject": &coretestcases.StringBoolGherkins{
			Title:      "FullString includes all fields",
			Feature:    "regex",
			Given:      "a valid pattern",
			When:       "struct has all fields populated",
			Then:       "output is formatted",
			Input:      "test-pattern",
			Expected:   true,
			Actual:     false,
			IsMatching: true,
		},
		"expectedLineCount": "9",
		"line0":             "Title:      FullString includes all fields",
		"line1":             "Feature:    regex",
		"line2":             "Given:      a valid pattern",
		"line3":             "When:       struct has all fields populated",
		"line4":             "Then:       output is formatted",
		"line5":             "Input:      test-pattern",
		"line6":             "Expected:   true",
		"line7":             "Actual:     false",
		"line8":             "IsMatching: true",
	},
}

var fullStringNilTestCase = coretestcases.StringBoolGherkins{
	Title: "FullString handles nil receiver",
	When:  "receiver is nil",
	ExtraArgs: args.Map{
		"expectedResult": "<nil GenericGherkins>",
	},
}

// --- ShouldBeEqual (via ShouldBeEqualArgs) ---

var shouldBeEqualPassingTestCase = coretestcases.StringBoolGherkins{
	Title: "ShouldBeEqualArgs passes when lines match",
	When:  "actual lines match expected lines",
	ExpectedLines: []string{
		"line-a",
		"line-b",
	},
}

// --- CaseTitle ---

var caseTitleUseTitleTestCase = coretestcases.StringBoolGherkins{
	Title: "CaseTitle returns Title when set",
	When:  "when-fallback",
	ExtraArgs: args.Map{
		"expectedResult": "CaseTitle returns Title when set",
	},
}

var caseTitleFallbackToWhenTestCase = coretestcases.StringBoolGherkins{
	Title: "",
	When:  "when-fallback-value",
	ExtraArgs: args.Map{
		"expectedResult": "when-fallback-value",
	},
}
