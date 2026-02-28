package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// Collection — New Creator String
// ==========================================

var collectionStringCapTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.String.Cap creates collection with correct capacity",
		ArrangeInput: args.Map{
			"when":     "given capacity 10",
			"capacity": 10,
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

var collectionStringEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.String.Empty creates zero-length collection",
		ArrangeInput: args.Map{
			"when": "given empty string collection",
		},
		ExpectedInput: []string{
			"0",
			"true",
		},
	},
}

var collectionStringFromTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.String.From wraps existing slice",
		ArrangeInput: args.Map{
			"when":  "given existing string slice",
			"items": []string{"alpha", "beta", "gamma"},
		},
		ExpectedInput: []string{
			"3",
			"false",
			"alpha",
			"gamma",
		},
	},
}

var collectionStringItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.String.Items creates from variadic",
		ArrangeInput: args.Map{
			"when":  "given variadic strings",
			"items": []string{"one", "two", "three"},
		},
		ExpectedInput: []string{
			"3",
			"one",
			"three",
		},
	},
}

// ==========================================
// Collection — New Creator Int
// ==========================================

var collectionIntItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Int.Items stores integers correctly",
		ArrangeInput: args.Map{
			"when":  "given int items",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"3",
			"10",
			"30",
		},
	},
}

// ==========================================
// Collection — Filter
// ==========================================

var collectionFilterTestCases = []coretestcases.CaseV1{
	{
		Title: "Filter returns only matching items",
		ArrangeInput: args.Map{
			"when":  "given ints, filter > 2",
			"items": []int{1, 2, 3, 4, 5},
		},
		ExpectedInput: []string{
			"3",
			"3",
			"5",
		},
	},
}

// ==========================================
// Collection — Clone independence
// ==========================================

var collectionCloneIndependenceTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy — mutations don't propagate",
		ArrangeInput: args.Map{
			"when":  "given collection cloned then mutated",
			"items": []string{"x", "y", "z"},
		},
		ExpectedInput: []string{
			"3",
			"true",
		},
	},
}

// ==========================================
// Hashset — Basic operations
// ==========================================

var hashsetAddHasTestCases = []coretestcases.CaseV1{
	{
		Title: "Hashset.Add then Has returns true",
		ArrangeInput: args.Map{
			"when":  "given string hashset with items added",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: []string{
			"3",
			"true",
			"true",
			"false",
		},
	},
}

var hashsetRemoveTestCases = []coretestcases.CaseV1{
	{
		Title: "Hashset.Remove removes existing key",
		ArrangeInput: args.Map{
			"when":   "given hashset with item removed",
			"items":  []string{"a", "b", "c"},
			"remove": "b",
		},
		ExpectedInput: []string{
			"2",
			"false",
			"true",
		},
	},
}

// ==========================================
// Hashmap — Basic operations
// ==========================================

var hashmapSetGetTestCases = []coretestcases.CaseV1{
	{
		Title: "Hashmap.Set then Get returns correct value",
		ArrangeInput: args.Map{
			"when": "given string-string hashmap with entries",
		},
		ExpectedInput: []string{
			"2",
			"value1",
			"true",
			"false",
		},
	},
}

// ==========================================
// SimpleSlice — Basic operations
// ==========================================

var simpleSliceAddTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice.Add appends items correctly",
		ArrangeInput: args.Map{
			"when":  "given int simple slice with items added",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: []string{
			"3",
			"10",
			"30",
		},
	},
}

// ==========================================
// LinkedList — Basic operations
// ==========================================

var linkedListAddTestCases = []coretestcases.CaseV1{
	{
		Title: "LinkedList.Add appends items and maintains order",
		ArrangeInput: args.Map{
			"when":  "given string linked list with items",
			"items": []string{"first", "second", "third"},
		},
		ExpectedInput: []string{
			"3",
			"first",
			"third",
		},
	},
}

var linkedListAddFrontTestCases = []coretestcases.CaseV1{
	{
		Title: "LinkedList.AddFront prepends to the front",
		ArrangeInput: args.Map{
			"when":    "given linked list with front-added item",
			"items":   []string{"b", "c"},
			"prepend": "a",
		},
		ExpectedInput: []string{
			"3",
			"a",
			"c",
		},
	},
}

// ==========================================
// Generic funcs — MapCollection
// ==========================================

var mapCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "MapCollection transforms int to string",
		ArrangeInput: args.Map{
			"when":  "given int collection mapped to strings",
			"items": []int{1, 2, 3},
		},
		ExpectedInput: []string{
			"3",
			"1",
			"3",
		},
	},
}

// ==========================================
// Generic funcs — Distinct
// ==========================================

var distinctTestCases = []coretestcases.CaseV1{
	{
		Title: "Distinct removes duplicates preserving order",
		ArrangeInput: args.Map{
			"when":  "given collection with duplicates",
			"items": []int{1, 2, 3, 2, 1, 4},
		},
		ExpectedInput: []string{
			"4",
			"1",
			"4",
		},
	},
}
