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
		ExpectedInput: args.Three[string, string, string]{
			First:  "3",     // length
			Second: "false", // isEmpty
			Third:  "true",  // hasItems
		},
	},
	{
		Title: "Positive: Empty collection",
		ArrangeInput: args.Map{
			"when":  "given no items",
			"count": 0,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "0",     // length
			Second: "true",  // isEmpty
			Third:  "false", // hasItems
		},
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "2",    // length
			Second: "true", // hasFirstItem
			Third:  "true", // joinContainsAll
		},
	},
	{
		Title: "Negative: Single item collection",
		ArrangeInput: args.Map{
			"when":  "given 1 StringInt item",
			"count": 1,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "1",     // length
			Second: "true",  // hasFirstItem
			Third:  "false", // joinContainsAll
		},
	},
}

// endregion

// region Collection Prepend/Append tests

var collectionPrependTestCase = coretestcases.CaseV1{
	Title: "Positive: Prepend adds items to front",
	ArrangeInput: args.Map{
		"when": "prepend 1 item to 2-item collection",
		"op":   "prepend",
	},
	ExpectedInput: args.Two[string, string]{
		First:  "3",         // length
		Second: "prepended", // firstItem
	},
}

var collectionAppendTestCase = coretestcases.CaseV1{
	Title: "Positive: Append adds items to back",
	ArrangeInput: args.Map{
		"when": "append 1 item to 2-item collection",
		"op":   "append",
	},
	ExpectedInput: args.Two[string, string]{
		First:  "3",        // length
		Second: "appended", // lastItem
	},
}

var collectionPrependIfFalseTestCase = coretestcases.CaseV1{
	Title: "Negative: PrependIf with false skips",
	ArrangeInput: args.Map{
		"when": "prepend with false condition",
		"op":   "prependif-false",
	},
	ExpectedInput: args.Two[string, string]{
		First:  "2",          // length
		Second: "original-0", // firstItem
	},
}

var collectionAppendIfFalseTestCase = coretestcases.CaseV1{
	Title: "Negative: AppendIf with false skips",
	ArrangeInput: args.Map{
		"when": "append with false condition",
		"op":   "appendif-false",
	},
	ExpectedInput: args.Two[string, string]{
		First:  "2",          // length
		Second: "original-0", // firstItem
	},
}

// endregion

// region Collection Clone tests

var collectionCloneValidTestCase = coretestcases.CaseV1{
	Title: "Positive: Clone produces independent copy",
	ArrangeInput: args.Map{
		"when":  "clone a 3-item collection",
		"count": 3,
	},
	ExpectedInput: args.Three[string, string, string]{
		First:  "3",    // length
		Second: "true", // sameContent
		Third:  "true", // isIndependent
	},
}

var collectionCloneNilTestCase = coretestcases.CaseV1{
	Title: "Negative: ClonePtr on nil returns nil",
	ArrangeInput: args.Map{
		"when": "clone nil collection",
	},
	ExpectedInput: "true", // isNil
}

// endregion

// region Collection IsEqualByString tests

var collectionIsEqualEqualTestCase = coretestcases.CaseV1{
	Title: "Positive: Same items are equal",
	ArrangeInput: args.Map{
		"when": "two identical collections",
		"case": "equal",
	},
	ExpectedInput: "true", // isEqual
}

var collectionIsEqualNotEqualTestCase = coretestcases.CaseV1{
	Title: "Negative: Different items are not equal",
	ArrangeInput: args.Map{
		"when": "two different collections",
		"case": "notequal",
	},
	ExpectedInput: "false", // isEqual
}

var collectionIsEqualDiffLengthTestCase = coretestcases.CaseV1{
	Title: "Negative: Different lengths are not equal",
	ArrangeInput: args.Map{
		"when": "collections with different lengths",
		"case": "difflength",
	},
	ExpectedInput: "false", // isEqual
}

var collectionIsEqualBothNilsTestCase = coretestcases.CaseV1{
	Title: "Negative: Both nil are equal",
	ArrangeInput: args.Map{
		"when": "both nil collections",
		"case": "bothnils",
	},
	ExpectedInput: "true", // isEqual
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
		ExpectedInput: args.Two[string, string]{
			First:  "true", // hasError
			Second: "true", // errorContainsItems
		},
	},
	{
		Title: "Negative: Empty collection returns nil error",
		ArrangeInput: args.Map{
			"when":  "empty collection",
			"count": 0,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // hasError
			Second: "false", // errorContainsItems
		},
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
		ExpectedInput: "true", // isEmptyAfterDispose
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
		ExpectedInput: args.Two[string, string]{
			First:  "4", // mergedLength
			Second: "2", // originalLength
		},
	},
	{
		Title: "Negative: ConcatNew with no extra items",
		ArrangeInput: args.Map{
			"when":     "concat 0 items onto 2-item collection",
			"original": 2,
			"extra":    0,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "2", // mergedLength
			Second: "2", // originalLength
		},
	},
}

// endregion

// region StringMapAnyCollection tests

var stringMapAnyCollectionWithValuesTestCase = coretestcases.CaseV1{
	Title: "Positive: StringMapAnyCollection stores map values",
	ArrangeInput: args.Map{
		"when": "given 2 map items",
		"mapValues": []map[string]any{
			{"key": 0},
			{"key": 1},
		},
	},
	ExpectedInput: args.Two[string, string]{
		First:  "2",    // length
		Second: "true", // hasValues
	},
}

var stringMapAnyCollectionNilValueTestCase = coretestcases.CaseV1{
	Title: "Negative: StringMapAnyCollection with nil map value",
	ArrangeInput: args.Map{
		"when": "given item with nil map",
		"mapValues": []map[string]any{
			nil,
		},
	},
	ExpectedInput: args.Two[string, string]{
		First:  "1",    // length
		Second: "true", // isNilMap
	},
}

// endregion
