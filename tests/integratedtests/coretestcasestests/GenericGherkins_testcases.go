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
	},
	ExpectedLines: []string{
		"",
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
	},
	ExpectedLines: []string{
		`Title: "A" != "B"`,
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
	},
	ExpectedLines: []string{
		"Input: alpha != beta",
	},
}

// --- CompareWith: nil both ---

var compareWithBothNilTestCase = coretestcases.StringBoolGherkins{
	Title:    "CompareWith returns true when both nil",
	When:     "both pointers are nil",
	Expected: true,
	ExtraArgs: args.Map{
		"bothNil": true,
	},
	ExpectedLines: []string{
		"",
	},
}

// --- CompareWith: one nil ---

var compareWithOneNilTestCase = coretestcases.StringBoolGherkins{
	Title:    "CompareWith returns false when one is nil",
	When:     "only one pointer is nil",
	Expected: false,
	ExtraArgs: args.Map{
		"a": &coretestcases.StringBoolGherkins{Title: "exists"},
	},
	ExpectedLines: []string{
		"one side is nil",
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
	},
	ExpectedLines: []string{
		"Title:      FullString includes all fields",
		"Feature:    regex",
		"Given:      a valid pattern",
		"When:       struct has all fields populated",
		"Then:       output is formatted",
		"Input:      test-pattern",
		"Expected:   true",
		"Actual:     false",
		"IsMatching: true",
	},
}

var fullStringNilTestCase = coretestcases.StringBoolGherkins{
	Title: "FullString handles nil receiver",
	When:  "receiver is nil",
	ExpectedLines: []string{
		"<nil GenericGherkins>",
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
	ExpectedLines: []string{
		"CaseTitle returns Title when set",
	},
}

var caseTitleFallbackToWhenTestCase = coretestcases.StringBoolGherkins{
	Title: "",
	When:  "when-fallback-value",
	ExpectedLines: []string{
		"when-fallback-value",
	},
}
