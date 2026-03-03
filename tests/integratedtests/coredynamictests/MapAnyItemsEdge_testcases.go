package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// MapAnyItems — IsEqual
// ==========================================

var mapAnyItemsIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqual - both nil should return true",
		ArrangeInput: args.Map{
			"when":      "both left and right are nil",
			"leftNil":   true,
			"rightNil":  true,
			"rightMap":  map[string]any{},
			"leftMap":   map[string]any{},
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsEqual - left nil right non-nil should return false",
		ArrangeInput: args.Map{
			"when":     "left is nil, right has data",
			"leftNil":  true,
			"rightNil": false,
			"rightMap": map[string]any{"k": "v"},
			"leftMap":  map[string]any{},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsEqual - right nil should return false",
		ArrangeInput: args.Map{
			"when":     "left has data, right is nil",
			"leftNil":  false,
			"rightNil": true,
			"leftMap":  map[string]any{"k": "v"},
			"rightMap": map[string]any{},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsEqual - same content should return true",
		ArrangeInput: args.Map{
			"when":     "both have identical key-value pairs",
			"leftNil":  false,
			"rightNil": false,
			"leftMap":  map[string]any{"a": "1", "b": "2"},
			"rightMap": map[string]any{"a": "1", "b": "2"},
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsEqual - different values should return false",
		ArrangeInput: args.Map{
			"when":     "same keys but different values",
			"leftNil":  false,
			"rightNil": false,
			"leftMap":  map[string]any{"a": "1"},
			"rightMap": map[string]any{"a": "2"},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsEqual - different keys should return false",
		ArrangeInput: args.Map{
			"when":     "different keys same values",
			"leftNil":  false,
			"rightNil": false,
			"leftMap":  map[string]any{"a": "1"},
			"rightMap": map[string]any{"b": "1"},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsEqual - different lengths should return false",
		ArrangeInput: args.Map{
			"when":     "left has 1 item, right has 2",
			"leftNil":  false,
			"rightNil": false,
			"leftMap":  map[string]any{"a": "1"},
			"rightMap": map[string]any{"a": "1", "b": "2"},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsEqual - both empty should return true",
		ArrangeInput: args.Map{
			"when":     "both are empty maps",
			"leftNil":  false,
			"rightNil": false,
			"leftMap":  map[string]any{},
			"rightMap": map[string]any{},
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// MapAnyItems — IsEqualRaw
// ==========================================

var mapAnyItemsIsEqualRawTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqualRaw - nil receiver should return false",
		ArrangeInput: args.Map{
			"when":     "receiver is nil, raw map has data",
			"leftNil":  true,
			"rightMap": map[string]any{"k": "v"},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "IsEqualRaw - nil receiver nil map should return true",
		ArrangeInput: args.Map{
			"when":    "both receiver and raw map are nil",
			"leftNil": true,
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "IsEqualRaw - matching map should return true",
		ArrangeInput: args.Map{
			"when":     "receiver and raw map have same content",
			"leftNil":  false,
			"leftMap":  map[string]any{"x": "y"},
			"rightMap": map[string]any{"x": "y"},
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// MapAnyItems — ClonePtr
// ==========================================

var mapAnyItemsClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr - nil receiver should return nil and error",
		ArrangeInput: args.Map{
			"when":    "receiver is nil",
			"leftNil": true,
		},
		ExpectedInput: []string{
			"true",  // hasError
			"true",  // cloneIsNil
		},
	},
	{
		Title: "ClonePtr - valid data should clone successfully",
		ArrangeInput: args.Map{
			"when":    "receiver has name and age",
			"leftNil": false,
			"leftMap": map[string]any{"name": "alice", "age": float64(30)},
		},
		ExpectedInput: []string{
			"false", // hasError
			"false", // cloneIsNil
			"2",     // cloneLength
			"true",  // hasName
			"true",  // hasAge
		},
	},
	{
		Title: "ClonePtr - empty map should clone to empty",
		ArrangeInput: args.Map{
			"when":    "receiver is empty",
			"leftNil": false,
			"leftMap": map[string]any{},
		},
		ExpectedInput: []string{
			"false", // hasError
			"false", // cloneIsNil
			"0",     // cloneLength
		},
	},
	{
		Title: "ClonePtr - modifying clone should not affect original",
		ArrangeInput: args.Map{
			"when":         "clone is modified after cloning",
			"leftNil":      false,
			"leftMap":      map[string]any{"key": "original"},
			"addAfterClone": true,
		},
		ExpectedInput: []string{
			"false", // hasError
			"false", // cloneIsNil
			"false", // originalHasNewKey
			"true",  // cloneHasNewKey
		},
	},
}

// ==========================================
// MapAnyItems — Edge cases (Length, HasKey, Add)
// ==========================================

var mapAnyItemsEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "Length/IsEmpty/HasAnyItem - nil receiver safety",
		ArrangeInput: args.Map{
			"when":    "receiver is nil",
			"leftNil": true,
		},
		ExpectedInput: []string{
			"0",     // length
			"true",  // isEmpty
			"false", // hasAnyItem
		},
	},
	{
		Title: "HasKey - nil receiver should return false",
		ArrangeInput: args.Map{
			"when":    "receiver is nil, checking any key",
			"leftNil": true,
			"key":     "anything",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "HasKey - existing key should return true",
		ArrangeInput: args.Map{
			"when":    "map has the key",
			"leftNil": false,
			"leftMap": map[string]any{"key": "val"},
			"key":     "key",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "HasKey - missing key should return false",
		ArrangeInput: args.Map{
			"when":    "map does not have the key",
			"leftNil": false,
			"leftMap": map[string]any{"key": "val"},
			"key":     "nope",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Add - new key should return true",
		ArrangeInput: args.Map{
			"when":     "adding a new key to empty map",
			"leftNil":  false,
			"leftMap":  map[string]any{},
			"addKey":   "k",
			"addValue": "v",
		},
		ExpectedInput: []string{
			"true", // isNew
			"1",    // lengthAfter
		},
	},
	{
		Title: "Add - existing key should return false and overwrite",
		ArrangeInput: args.Map{
			"when":     "adding existing key",
			"leftNil":  false,
			"leftMap":  map[string]any{"k": "old"},
			"addKey":   "k",
			"addValue": "new",
		},
		ExpectedInput: []string{
			"false", // isNew
			"new",   // updatedValue
		},
	},
}
