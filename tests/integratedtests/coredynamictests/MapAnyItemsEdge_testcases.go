package coredynamictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// Note: MapAnyItems nil receiver test cases (Length, IsEmpty, HasAnyItem) migrated to
// MapAnyItemsEdge_NilReceiver_testcases.go using CaseNilSafe pattern.

// ==========================================
// MapAnyItems — IsEqual
// ==========================================

var mapAnyItemsIsEqualBothNilTestCase = coretestcases.CaseV1{
	Title: "IsEqual - both nil should return true",
	ArrangeInput: args.Map{
		"when":     "both left and right are nil",
		"leftNil":  true,
		"rightNil": true,
		"rightMap": map[string]any{},
		"leftMap":  map[string]any{},
	},
	ExpectedInput: args.Map{
		"isEqual": true,
	},
}

var mapAnyItemsIsEqualLeftNilTestCase = coretestcases.CaseV1{
	Title: "IsEqual - left nil right non-nil should return false",
	ArrangeInput: args.Map{
		"when":     "left is nil, right has data",
		"leftNil":  true,
		"rightNil": false,
		"rightMap": map[string]any{"k": "v"},
		"leftMap":  map[string]any{},
	},
	ExpectedInput: args.Map{
		"isEqual": false,
	},
}

var mapAnyItemsIsEqualRightNilTestCase = coretestcases.CaseV1{
	Title: "IsEqual - right nil should return false",
	ArrangeInput: args.Map{
		"when":     "left has data, right is nil",
		"leftNil":  false,
		"rightNil": true,
		"leftMap":  map[string]any{"k": "v"},
		"rightMap": map[string]any{},
	},
	ExpectedInput: args.Map{
		"isEqual": false,
	},
}

var mapAnyItemsIsEqualSameContentTestCase = coretestcases.CaseV1{
	Title: "IsEqual - same content should return true",
	ArrangeInput: args.Map{
		"when":     "both have identical key-value pairs",
		"leftNil":  false,
		"rightNil": false,
		"leftMap":  map[string]any{"a": "1", "b": "2"},
		"rightMap": map[string]any{"a": "1", "b": "2"},
	},
	ExpectedInput: args.Map{
		"isEqual": true,
	},
}

var mapAnyItemsIsEqualDiffValuesTestCase = coretestcases.CaseV1{
	Title: "IsEqual - different values should return false",
	ArrangeInput: args.Map{
		"when":     "same keys but different values",
		"leftNil":  false,
		"rightNil": false,
		"leftMap":  map[string]any{"a": "1"},
		"rightMap": map[string]any{"a": "2"},
	},
	ExpectedInput: args.Map{
		"isEqual": false,
	},
}

var mapAnyItemsIsEqualDiffKeysTestCase = coretestcases.CaseV1{
	Title: "IsEqual - different keys should return false",
	ArrangeInput: args.Map{
		"when":     "different keys same values",
		"leftNil":  false,
		"rightNil": false,
		"leftMap":  map[string]any{"a": "1"},
		"rightMap": map[string]any{"b": "1"},
	},
	ExpectedInput: args.Map{
		"isEqual": false,
	},
}

var mapAnyItemsIsEqualDiffLengthsTestCase = coretestcases.CaseV1{
	Title: "IsEqual - different lengths should return false",
	ArrangeInput: args.Map{
		"when":     "left has 1 item, right has 2",
		"leftNil":  false,
		"rightNil": false,
		"leftMap":  map[string]any{"a": "1"},
		"rightMap": map[string]any{"a": "1", "b": "2"},
	},
	ExpectedInput: args.Map{
		"isEqual": false,
	},
}

var mapAnyItemsIsEqualBothEmptyTestCase = coretestcases.CaseV1{
	Title: "IsEqual - both empty should return true",
	ArrangeInput: args.Map{
		"when":     "both are empty maps",
		"leftNil":  false,
		"rightNil": false,
		"leftMap":  map[string]any{},
		"rightMap": map[string]any{},
	},
	ExpectedInput: args.Map{
		"isEqual": true,
	},
}

// ==========================================
// MapAnyItems — IsEqualRaw
// ==========================================

var mapAnyItemsIsEqualRawNilReceiverTestCase = coretestcases.CaseV1{
	Title: "IsEqualRaw - nil receiver should return false",
	ArrangeInput: args.Map{
		"when":     "receiver is nil, raw map has data",
		"leftNil":  true,
		"rightMap": map[string]any{"k": "v"},
	},
	ExpectedInput: args.Map{
		"isEqualRaw": false,
	},
}

var mapAnyItemsIsEqualRawBothNilTestCase = coretestcases.CaseV1{
	Title: "IsEqualRaw - nil receiver nil map should return true",
	ArrangeInput: args.Map{
		"when":    "both receiver and raw map are nil",
		"leftNil": true,
	},
	ExpectedInput: args.Map{
		"isEqualRaw": true,
	},
}

var mapAnyItemsIsEqualRawMatchingTestCase = coretestcases.CaseV1{
	Title: "IsEqualRaw - matching map should return true",
	ArrangeInput: args.Map{
		"when":     "receiver and raw map have same content",
		"leftNil":  false,
		"leftMap":  map[string]any{"x": "y"},
		"rightMap": map[string]any{"x": "y"},
	},
	ExpectedInput: args.Map{
		"isEqualRaw": true,
	},
}

// ==========================================
// MapAnyItems — ClonePtr
// ==========================================

var mapAnyItemsClonePtrNilTestCase = coretestcases.CaseV1{
	Title: "ClonePtr - nil receiver should return nil and error",
	ArrangeInput: args.Map{
		"when":    "receiver is nil",
		"leftNil": true,
	},
	ExpectedInput: args.Map{
		"hasError":   true,
		"cloneIsNil": true,
	},
}

var mapAnyItemsClonePtrValidTestCase = coretestcases.CaseV1{
	Title: "ClonePtr - valid data should clone successfully",
	ArrangeInput: args.Map{
		"when":    "receiver has name and age",
		"leftNil": false,
		"leftMap": map[string]any{"name": "alice", "age": float64(30)},
	},
	ExpectedInput: args.Map{
		"hasError":    false,
		"cloneIsNil":  false,
		"cloneLength": 2,
		"hasName":     true,
		"hasAge":      true,
	},
}

var mapAnyItemsClonePtrEmptyTestCase = coretestcases.CaseV1{
	Title: "ClonePtr - empty map should clone to empty",
	ArrangeInput: args.Map{
		"when":    "receiver is empty",
		"leftNil": false,
		"leftMap": map[string]any{},
	},
	ExpectedInput: args.Map{
		"hasError":    false,
		"cloneIsNil":  false,
		"cloneLength": 0,
	},
}

var mapAnyItemsClonePtrIndependenceTestCase = coretestcases.CaseV1{
	Title: "ClonePtr - modifying clone should not affect original",
	ArrangeInput: args.Map{
		"when":         "clone is modified after cloning",
		"leftNil":      false,
		"leftMap":      map[string]any{"key": "original"},
		"addAfterClone": true,
	},
	ExpectedInput: args.Map{
		"hasError":         false,
		"cloneIsNil":       false,
		"originalHasNewKey": false,
		"cloneHasNewKey":    true,
	},
}

// ==========================================
// MapAnyItems — Edge cases (Length, HasKey, Add)
// ==========================================

var mapAnyItemsNilLengthTestCase = coretestcases.CaseV1{
	Title: "Length/IsEmpty/HasAnyItem - nil receiver safety",
	ArrangeInput: args.Map{
		"when":    "receiver is nil",
		"leftNil": true,
	},
	ExpectedInput: args.Map{
		"length":     0,
		"isEmpty":    true,
		"hasAnyItem": false,
	},
}

var mapAnyItemsHasKeyNilTestCase = coretestcases.CaseV1{
	Title: "HasKey - nil receiver should return false",
	ArrangeInput: args.Map{
		"when":    "receiver is nil, checking any key",
		"leftNil": true,
		"key":     "anything",
	},
	ExpectedInput: args.Map{
		"hasKey": false,
	},
}

var mapAnyItemsHasKeyExistsTestCase = coretestcases.CaseV1{
	Title: "HasKey - existing key should return true",
	ArrangeInput: args.Map{
		"when":    "map has the key",
		"leftNil": false,
		"leftMap": map[string]any{"key": "val"},
		"key":     "key",
	},
	ExpectedInput: args.Map{
		"hasKey": true,
	},
}

var mapAnyItemsHasKeyMissingTestCase = coretestcases.CaseV1{
	Title: "HasKey - missing key should return false",
	ArrangeInput: args.Map{
		"when":    "map does not have the key",
		"leftNil": false,
		"leftMap": map[string]any{"key": "val"},
		"key":     "nope",
	},
	ExpectedInput: args.Map{
		"hasKey": false,
	},
}

var mapAnyItemsAddNewKeyTestCase = coretestcases.CaseV1{
	Title: "Add - new key should return true",
	ArrangeInput: args.Map{
		"when":     "adding a new key to empty map",
		"leftNil":  false,
		"leftMap":  map[string]any{},
		"addKey":   "k",
		"addValue": "v",
	},
	ExpectedInput: args.Map{
		"isNew":       true,
		"lengthAfter": 1,
	},
}

var mapAnyItemsAddExistingKeyTestCase = coretestcases.CaseV1{
	Title: "Add - existing key should return false and overwrite",
	ArrangeInput: args.Map{
		"when":     "adding existing key",
		"leftNil":  false,
		"leftMap":  map[string]any{"k": "old"},
		"addKey":   "k",
		"addValue": "new",
	},
	ExpectedInput: args.Map{
		"isNew":        false,
		"updatedValue": "new",
	},
}
