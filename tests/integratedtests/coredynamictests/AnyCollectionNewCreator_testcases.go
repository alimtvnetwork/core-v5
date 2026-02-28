package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// New Creator — Any Collection: Empty
// ==========================================

var newCreatorGenericEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Empty creates zero-length collection",
		ArrangeInput: args.Map{
			"when": "given empty any collection",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// New Creator — Any Collection: Cap
// ==========================================

var newCreatorGenericCapTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Cap creates collection with correct capacity",
		ArrangeInput: args.Map{
			"when":     "given capacity 10",
			"capacity": 10,
		},
		ExpectedInput: []string{
			"0",
			"true",
			"false",
		},
	},
}

var newCreatorGenericCapZeroTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Cap with zero capacity creates empty collection",
		ArrangeInput: args.Map{
			"when":     "given capacity 0",
			"capacity": 0,
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// New Creator — Any Collection: From
// ==========================================

var newCreatorGenericFromTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.From wraps existing slice",
		ArrangeInput: args.Map{
			"when":  "given existing any slice",
			"items": []any{"alpha", 42, true},
		},
		ExpectedInput: []string{
			"3",
			"false",
			"alpha",
			"true",
		},
	},
}

var newCreatorGenericFromEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.From with empty slice creates empty collection",
		ArrangeInput: args.Map{
			"when":  "given empty any slice",
			"items": []any{},
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// New Creator — Any Collection: Clone
// ==========================================

var newCreatorGenericCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Clone creates independent copy",
		ArrangeInput: args.Map{
			"when":  "given cloned any slice",
			"items": []any{"x", "y"},
		},
		ExpectedInput: []string{
			"2",
			"x",
			"y",
		},
	},
}

var newCreatorGenericCloneMutationTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Clone mutation does not affect original slice",
		ArrangeInput: args.Map{
			"when":  "given cloned collection then mutated",
			"items": []any{"a", "b", "c"},
		},
		ExpectedInput: []string{
			"3",
			"true",
		},
	},
}

// ==========================================
// New Creator — Any Collection: Items
// ==========================================

var newCreatorGenericItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Items creates from variadic",
		ArrangeInput: args.Map{
			"when":  "given variadic any items",
			"items": []any{"one", 2, 3.0},
		},
		ExpectedInput: []string{
			"3",
			"one",
			"3",
		},
	},
}

var newCreatorGenericItemsSingleTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Items with single item",
		ArrangeInput: args.Map{
			"when":  "given single any item",
			"items": []any{"solo"},
		},
		ExpectedInput: []string{
			"1",
			"solo",
			"solo",
		},
	},
}

// ==========================================
// Negative/Edge — Any Collection: From nil
// ==========================================

var newCreatorGenericFromNilTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.From with nil slice creates empty collection",
		ArrangeInput: args.Map{
			"when": "given nil slice",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// Negative/Edge — Any Collection: Cap large
// ==========================================

var newCreatorGenericCapLargeTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Cap with large capacity allocates without error",
		ArrangeInput: args.Map{
			"when":     "given capacity 1000000",
			"capacity": 1000000,
		},
		ExpectedInput: []string{
			"0",
			"true",
			"1000000",
		},
	},
}

// ==========================================
// Negative/Edge — Any Collection: Items no args
// ==========================================

var newCreatorGenericItemsNoArgsTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Items with no arguments creates empty collection",
		ArrangeInput: args.Map{
			"when": "given zero variadic args",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

// ==========================================
// Negative/Edge — Any Collection: Clone nil
// ==========================================

var newCreatorGenericCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Any.Clone with nil slice creates empty collection",
		ArrangeInput: args.Map{
			"when": "given nil slice to clone",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}
