package namevaluetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// region StringStringCollection tests

var stringStringCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Add items and verify length",
		ArrangeInput: args.Map{
			"when":  "given 3 StringString items",
			"count": 3,
		},
		ExpectedInput: []string{"3", "false", "true"},
	},
	{
		Title: "Positive: Empty collection",
		ArrangeInput: args.Map{
			"when":  "given no items",
			"count": 0,
		},
		ExpectedInput: []string{"0", "true", "false"},
	},
}

// endregion

// region StringIntCollection tests

var stringIntCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Add StringInt items and join",
		ArrangeInput: args.Map{
			"when":  "given 2 StringInt items",
			"count": 2,
		},
		ExpectedInput: []string{"2", "true", "true"},
	},
	{
		Title: "Negative: Single item collection",
		ArrangeInput: args.Map{
			"when":  "given 1 StringInt item",
			"count": 1,
		},
		ExpectedInput: []string{"1", "true", "false"},
	},
}

// endregion

// region Collection Prepend/Append tests

var collectionPrependAppendTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Prepend adds items to front",
		ArrangeInput: args.Map{
			"when": "prepend 1 item to 2-item collection",
			"op":   "prepend",
		},
		ExpectedInput: []string{"3", "prepended"},
	},
	{
		Title: "Positive: Append adds items to back",
		ArrangeInput: args.Map{
			"when": "append 1 item to 2-item collection",
			"op":   "append",
		},
		ExpectedInput: []string{"3", "appended"},
	},
	{
		Title: "Negative: PrependIf with false skips",
		ArrangeInput: args.Map{
			"when": "prepend with false condition",
			"op":   "prependif-false",
		},
		ExpectedInput: []string{"2", "original-0"},
	},
	{
		Title: "Negative: AppendIf with false skips",
		ArrangeInput: args.Map{
			"when": "append with false condition",
			"op":   "appendif-false",
		},
		ExpectedInput: []string{"2", "original-0"},
	},
}

// endregion

// region Collection Clone tests

var collectionCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Clone produces independent copy",
		ArrangeInput: args.Map{
			"when":  "clone a 3-item collection",
			"count": 3,
		},
		ExpectedInput: []string{"3", "true", "true"},
	},
	{
		Title: "Negative: ClonePtr on nil returns nil",
		ArrangeInput: args.Map{
			"when":  "clone nil collection",
			"count": -1,
		},
		ExpectedInput: []string{"true"},
	},
}

// endregion

// region Collection IsEqualByString tests

var collectionIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Same items are equal",
		ArrangeInput: args.Map{
			"when":   "two identical collections",
			"case":   "equal",
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "Negative: Different items are not equal",
		ArrangeInput: args.Map{
			"when":   "two different collections",
			"case":   "notequal",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "Negative: Different lengths are not equal",
		ArrangeInput: args.Map{
			"when":   "collections with different lengths",
			"case":   "difflength",
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "Negative: Both nil are equal",
		ArrangeInput: args.Map{
			"when":   "both nil collections",
			"case":   "bothnils",
		},
		ExpectedInput: []string{"true"},
	},
}

// endregion

// region Collection Error tests

var collectionErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Non-empty collection returns error",
		ArrangeInput: args.Map{
			"when":  "collection with items",
			"count": 2,
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "Negative: Empty collection returns nil error",
		ArrangeInput: args.Map{
			"when":  "empty collection",
			"count": 0,
		},
		ExpectedInput: []string{"false", "false"},
	},
}

// endregion

// region Collection Dispose tests

var collectionDisposeTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: Dispose clears all items",
		ArrangeInput: args.Map{
			"when":  "dispose a 3-item collection",
			"count": 3,
		},
		ExpectedInput: []string{"true"},
	},
}

// endregion

// region Collection ConcatNew tests

var collectionConcatNewTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: ConcatNew returns new collection with merged items",
		ArrangeInput: args.Map{
			"when":     "concat 2 items onto 2-item collection",
			"original": 2,
			"extra":    2,
		},
		ExpectedInput: []string{"4", "2"},
	},
	{
		Title: "Negative: ConcatNew with no extra items",
		ArrangeInput: args.Map{
			"when":     "concat 0 items onto 2-item collection",
			"original": 2,
			"extra":    0,
		},
		ExpectedInput: []string{"2", "2"},
	},
}

// endregion

// region StringMapAnyCollection tests

var stringMapAnyCollectionTestCases = []coretestcases.CaseV1{
	{
		Title: "Positive: StringMapAnyCollection stores map values",
		ArrangeInput: args.Map{
			"when":  "given 2 map items",
			"count": 2,
		},
		ExpectedInput: []string{"2", "true"},
	},
	{
		Title: "Negative: StringMapAnyCollection with nil map value",
		ArrangeInput: args.Map{
			"when":  "given item with nil map",
			"count": 1,
		},
		ExpectedInput: []string{"1", "true"},
	},
}

// endregion
