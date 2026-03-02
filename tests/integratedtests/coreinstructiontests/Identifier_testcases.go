package coreinstructiontests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var identifierTestCases = []coretestcases.CaseV1{
	// === Positive ===
	{
		Title: "NewIdentifier sets Id correctly",
		ArrangeInput: args.Map{
			"when": "given id 'test-123'",
			"id":   "test-123",
		},
		ExpectedInput: []string{
			"test-123",
			"false",
			"false",
		},
	},
	{
		Title: "NewIdentifier with special characters",
		ArrangeInput: args.Map{
			"when": "given id with special chars",
			"id":   "user@domain.com/resource#123",
		},
		ExpectedInput: []string{
			"user@domain.com/resource#123",
			"false",
			"false",
		},
	},

	// === Negative / empty ===
	{
		Title: "NewIdentifier with empty id is empty",
		ArrangeInput: args.Map{
			"when": "given empty id",
			"id":   "",
		},
		ExpectedInput: []string{
			"",
			"true",
			"true",
		},
	},
	{
		Title: "NewIdentifier with whitespace-only id",
		ArrangeInput: args.Map{
			"when": "given whitespace-only id",
			"id":   "   ",
		},
		ExpectedInput: []string{
			"   ",
			"false",
			"true",
		},
	},
}

// ============================================================================
// Identifiers collection tests
// ============================================================================

var identifiersLengthTestCases = []coretestcases.CaseV1{
	// === Positive ===
	{
		Title: "Identifiers Length returns correct count",
		ArrangeInput: args.Map{
			"when": "given 3 ids",
			"ids":  []string{"a", "b", "c"},
		},
		ExpectedInput: []string{"3", "false", "true"},
	},

	// === Boundary: empty ===
	{
		Title: "Identifiers Length returns 0 for empty",
		ArrangeInput: args.Map{
			"when": "given no ids",
			"ids":  []string{},
		},
		ExpectedInput: []string{"0", "true", "false"},
	},

	// === Single item ===
	{
		Title: "Identifiers Length returns 1 for single id",
		ArrangeInput: args.Map{
			"when": "given single id",
			"ids":  []string{"only"},
		},
		ExpectedInput: []string{"1", "false", "true"},
	},
}

var identifiersGetByIdTestCases = []coretestcases.CaseV1{
	// === Positive: found ===
	{
		Title: "GetById returns matching identifier",
		ArrangeInput: args.Map{
			"when":     "given existing id",
			"ids":      []string{"alpha", "beta", "gamma"},
			"searchId": "beta",
		},
		ExpectedInput: []string{"true", "beta"},
	},
	{
		Title: "GetById returns first item",
		ArrangeInput: args.Map{
			"when":     "given first id in list",
			"ids":      []string{"first", "second"},
			"searchId": "first",
		},
		ExpectedInput: []string{"true", "first"},
	},
	{
		Title: "GetById returns last item",
		ArrangeInput: args.Map{
			"when":     "given last id in list",
			"ids":      []string{"first", "last"},
			"searchId": "last",
		},
		ExpectedInput: []string{"true", "last"},
	},

	// === Negative: not found ===
	{
		Title: "GetById returns nil for non-existent id",
		ArrangeInput: args.Map{
			"when":     "given non-existent id",
			"ids":      []string{"alpha", "beta"},
			"searchId": "missing",
		},
		ExpectedInput: []string{"false", ""},
	},
	{
		Title: "GetById returns nil for empty search id",
		ArrangeInput: args.Map{
			"when":     "given empty search id",
			"ids":      []string{"alpha", "beta"},
			"searchId": "",
		},
		ExpectedInput: []string{"false", ""},
	},
	{
		Title: "GetById returns nil from empty collection",
		ArrangeInput: args.Map{
			"when":     "given empty collection",
			"ids":      []string{},
			"searchId": "any",
		},
		ExpectedInput: []string{"false", ""},
	},
}

var identifiersIndexOfTestCases = []coretestcases.CaseV1{
	// === Positive ===
	{
		Title: "IndexOf returns correct index for existing id",
		ArrangeInput: args.Map{
			"when":     "given existing id at index 1",
			"ids":      []string{"a", "b", "c"},
			"searchId": "b",
		},
		ExpectedInput: []string{"1"},
	},
	{
		Title: "IndexOf returns 0 for first item",
		ArrangeInput: args.Map{
			"when":     "given first id",
			"ids":      []string{"first", "second"},
			"searchId": "first",
		},
		ExpectedInput: []string{"0"},
	},

	// === Negative ===
	{
		Title: "IndexOf returns -1 for missing id",
		ArrangeInput: args.Map{
			"when":     "given non-existent id",
			"ids":      []string{"a", "b"},
			"searchId": "missing",
		},
		ExpectedInput: []string{"-1"},
	},
	{
		Title: "IndexOf returns -1 for empty search",
		ArrangeInput: args.Map{
			"when":     "given empty string search",
			"ids":      []string{"a"},
			"searchId": "",
		},
		ExpectedInput: []string{"-1"},
	},
	{
		Title: "IndexOf returns -1 for empty collection",
		ArrangeInput: args.Map{
			"when":     "given empty collection",
			"ids":      []string{},
			"searchId": "a",
		},
		ExpectedInput: []string{"-1"},
	},
}

var identifiersCloneTestCases = []coretestcases.CaseV1{
	// === Positive ===
	{
		Title: "Clone produces equal independent copy",
		ArrangeInput: args.Map{
			"when": "given 3 ids",
			"ids":  []string{"x", "y", "z"},
		},
		ExpectedInput: []string{"3", "x", "y", "z"},
	},

	// === Boundary: empty ===
	{
		Title: "Clone of empty produces empty",
		ArrangeInput: args.Map{
			"when": "given empty identifiers",
			"ids":  []string{},
		},
		ExpectedInput: []string{"0"},
	},
}

var identifiersAddTestCases = []coretestcases.CaseV1{
	// === Positive ===
	{
		Title: "Add appends new id",
		ArrangeInput: args.Map{
			"when":  "given existing ids and new id",
			"ids":   []string{"a"},
			"addId": "b",
		},
		ExpectedInput: []string{"2", "a", "b"},
	},

	// === Negative: empty id skipped ===
	{
		Title: "Add skips empty string id",
		ArrangeInput: args.Map{
			"when":  "given empty id to add",
			"ids":   []string{"a"},
			"addId": "",
		},
		ExpectedInput: []string{"1", "a"},
	},
}

// ============================================================================
// Specification tests
// ============================================================================

var specificationCloneTestCases = []coretestcases.CaseV1{
	// === Positive ===
	{
		Title: "Specification Clone copies all fields",
		ArrangeInput: args.Map{
			"when":     "given spec with all fields",
			"id":       "spec-1",
			"display":  "My Spec",
			"typeName": "typeA",
			"tags":     []string{"tag1", "tag2"},
			"isGlobal": true,
		},
		ExpectedInput: []string{
			"spec-1",
			"My Spec",
			"typeA",
			"2",
			"tag1",
			"tag2",
			"true",
		},
	},

	// === Boundary: empty tags ===
	{
		Title: "Specification Clone with empty tags",
		ArrangeInput: args.Map{
			"when":     "given spec with no tags",
			"id":       "spec-2",
			"display":  "Display",
			"typeName": "typeB",
			"tags":     []string{},
			"isGlobal": false,
		},
		ExpectedInput: []string{
			"spec-2",
			"Display",
			"typeB",
			"0",
			"false",
		},
	},
}

// ============================================================================
// BaseTags tests
// ============================================================================

var baseTagsTestCases = []coretestcases.CaseV1{
	// === Positive ===
	{
		Title: "BaseTags HasAllTags returns true when all present",
		ArrangeInput: args.Map{
			"when":       "given matching tags",
			"tags":       []string{"a", "b", "c"},
			"searchTags": []string{"a", "c"},
		},
		ExpectedInput: []string{"3", "false", "true", "true"},
	},

	// === Negative: partial match ===
	{
		Title: "BaseTags HasAllTags returns false when partial match",
		ArrangeInput: args.Map{
			"when":       "given partially matching tags",
			"tags":       []string{"a", "b"},
			"searchTags": []string{"a", "missing"},
		},
		ExpectedInput: []string{"2", "false", "false", "true"},
	},

	// === Boundary: empty tags ===
	{
		Title: "BaseTags empty tags returns true for empty search",
		ArrangeInput: args.Map{
			"when":       "given empty tags and empty search",
			"tags":       []string{},
			"searchTags": []string{},
		},
		ExpectedInput: []string{"0", "true", "true", "true"},
	},

	// === Negative: search on empty ===
	{
		Title: "BaseTags HasAllTags false when tags empty but search non-empty",
		ArrangeInput: args.Map{
			"when":       "given empty tags with non-empty search",
			"tags":       []string{},
			"searchTags": []string{"a"},
		},
		ExpectedInput: []string{"0", "true", "false", "false"},
	},
}
