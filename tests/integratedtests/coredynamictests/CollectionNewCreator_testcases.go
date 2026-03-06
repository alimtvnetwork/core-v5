package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// New Creator — String Collection
// ==========================================

var newCreatorStringCapTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.String.Cap creates collection with correct capacity",
		ArrangeInput: args.Map{
			"when":     "given capacity 10",
			"capacity": 10,
		},
		ExpectedInput: args.Map{
			"length":     0,
			"isEmpty":    true,
			"hasAnyItem": false,
		},
	},
}

var newCreatorStringEmptyTestCases = []coretestcases.CaseV1{
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

var newCreatorStringFromTestCases = []coretestcases.CaseV1{
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

var newCreatorStringCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.String.Clone creates independent copy",
		ArrangeInput: args.Map{
			"when":  "given cloned string slice",
			"items": []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"length": 2,
			"first":  "x",
			"last":   "y",
		},
	},
}

var newCreatorStringItemsTestCases = []coretestcases.CaseV1{
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
// New Creator — Int Collection
// ==========================================

var newCreatorIntCapTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Int.Cap creates int collection",
		ArrangeInput: args.Map{
			"when":     "given capacity 5",
			"capacity": 5,
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

var newCreatorIntItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "New.Collection.Int.Items stores integers correctly",
		ArrangeInput: args.Map{
			"when":  "given int items",
			"items": []int{10, 20, 30},
		},
		ExpectedInput: args.Map{
			"length": 3,
			"first":  "10",
			"last":   "30",
		},
	},
}

// ==========================================
// Collection Methods — AddIf
// ==========================================

var collectionAddIfTrueTestCases = []coretestcases.CaseV1{
	{
		Title: "AddIf true appends item",
		ArrangeInput: args.Map{
			"when":  "given isAdd true",
			"isAdd": true,
			"item":  "added",
		},
		ExpectedInput: args.Map{
			"length": 1,
			"first":  "added",
		},
	},
}

var collectionAddIfFalseTestCases = []coretestcases.CaseV1{
	{
		Title: "AddIf false does not append",
		ArrangeInput: args.Map{
			"when":  "given isAdd false",
			"isAdd": false,
			"item":  "skipped",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================
// Collection Methods — AddCollection
// ==========================================

var collectionAddCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCollection merges two collections",
		ArrangeInput: args.Map{
			"when":   "given two string collections",
			"first":  []string{"a", "b"},
			"second": []string{"c", "d"},
		},
		ExpectedInput: args.Map{
			"length": 4,
			"first":  "a",
			"last":   "d",
		},
	},
}

var collectionAddCollectionNilTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCollection with nil other keeps original",
		ArrangeInput: args.Map{
			"when":  "given nil other collection",
			"first": []string{"a"},
		},
		ExpectedInput: args.Map{
			"length": 1,
			"first":  "a",
		},
	},
}

// ==========================================
// Collection Methods — Clone
// ==========================================

var collectionCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy",
		ArrangeInput: args.Map{
			"when":  "given collection to clone",
			"items": []string{"x", "y", "z"},
		},
		ExpectedInput: args.Map{
			"length":        3,
			"first":         "x",
			"last":          "z",
			"isIndependent": true,
		},
	},
}

// ==========================================
// Collection Methods — Reverse
// ==========================================

var collectionReverseTestCases = []coretestcases.CaseV1{
	{
		Title: "Reverse reverses items in place",
		ArrangeInput: args.Map{
			"when":  "given 3-item collection",
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: []string{"c", "b", "a"},
	},
}

var collectionReverseEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Reverse on empty collection is safe",
		ArrangeInput: args.Map{
			"when": "given empty collection",
		},
		ExpectedInput: args.Map{
			"length":  0,
			"isEmpty": true,
		},
	},
}

// ==========================================
// Collection Methods — ConcatNew
// ==========================================

var collectionConcatNewTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatNew creates new collection without mutating original",
		ArrangeInput: args.Map{
			"when":     "given original + new items",
			"original": []string{"a", "b"},
			"adding":   []string{"c", "d"},
		},
		ExpectedInput: args.Map{
			"mergedLength":   4,
			"mergedFirst":    "a",
			"mergedLast":     "d",
			"originalLength": 2,
		},
	},
}

// ==========================================
// Collection Methods — Capacity/Resize
// ==========================================

var collectionCapacityTestCases = []coretestcases.CaseV1{
	{
		Title: "Capacity returns allocated capacity",
		ArrangeInput: args.Map{
			"when":     "given collection with capacity 20",
			"capacity": 20,
		},
		ExpectedInput: args.Map{
			"capacity": 20,
			"length":   0,
		},
	},
}

var collectionResizeTestCases = []coretestcases.CaseV1{
	{
		Title: "AddCapacity grows capacity",
		ArrangeInput: args.Map{
			"when":       "given capacity 5 then add 10",
			"capacity":   5,
			"additional": 10,
		},
		ExpectedInput: "true",
	},
}
