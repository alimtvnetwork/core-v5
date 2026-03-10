package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

var collectionStringEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.String.Empty creates zero-length collection",
		ArrangeInput: args.Map{
			"when": "given empty string collection",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
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
		ExpectedInput: args.Map{
			"length":  3,
			"isEmpty": false,
			"first":   "alpha",
			"last":    "gamma",
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "one",
			"last":   "three",
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  10,
			"last":   30,
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  3,
			"last":   5,
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
		ExpectedInput: args.Map{
			"originalLength": 3,
			"isIndependent":  true,
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
		ExpectedInput: args.Map{
			"length": 3,
			"hasA":   true,
			"hasC":   true,
			"hasZ":   false,
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
		ExpectedInput: args.Map{
			"length":     2,
			"hasRemoved": false,
			"hasA":       true,
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
		ExpectedInput: args.Map{
			"length":   2,
			"value":    "value1",
			"found":    true,
			"notFound": false,
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  10,
			"last":   30,
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "first",
			"last":   "third",
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "a",
			"last":   "c",
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
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "1",
			"last":   "3",
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
		ExpectedInput: args.Map{
			"length": 4,
			"first":  1,
			"last":   4,
		},
	},
}
