package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Constructors
// ==========================================================================

var hashmapEmptyTestCase = coretestcases.CaseV1{
	Title: "EmptyHashmap creates empty",
	ExpectedInput: args.Map{
		"isEmpty":  true,
		"length":   0,
		"hasItems": false,
	},
}

var hashmapNewTestCase = coretestcases.CaseV1{
	Title:         "NewHashmap with capacity",
	ExpectedInput: args.Map{"isEmpty": true},
}

var hashmapFromTestCase = coretestcases.CaseV1{
	Title: "HashmapFrom wraps map",
	ExpectedInput: args.Map{
		"length": 2,
		"hasKey": true,
	},
}

var hashmapCloneFuncTestCase = coretestcases.CaseV1{
	Title: "HashmapClone independence",
	ExpectedInput: args.Map{
		"origValue":   1,
		"clonedValue": 99,
	},
}

// ==========================================================================
// Set
// ==========================================================================

var hashmapSetNewTestCase = coretestcases.CaseV1{
	Title: "Set new key returns true",
	ExpectedInput: args.Map{
		"isNew":  true,
		"length": 1,
	},
}

var hashmapSetExistingTestCase = coretestcases.CaseV1{
	Title: "Set existing key returns false",
	ExpectedInput: args.Map{
		"isNew":        false,
		"updatedValue": 2,
	},
}

// ==========================================================================
// Get
// ==========================================================================

var hashmapGetFoundTestCase = coretestcases.CaseV1{
	Title: "Get found",
	ExpectedInput: args.Map{
		"found": true,
		"value": 42,
	},
}

var hashmapGetNotFoundTestCase = coretestcases.CaseV1{
	Title: "Get not found",
	ExpectedInput: args.Map{
		"found": false,
		"value": 0,
	},
}

var hashmapGetOrDefaultMissingTestCase = coretestcases.CaseV1{
	Title:         "GetOrDefault missing returns default",
	ExpectedInput: args.Map{"value": 99},
}

var hashmapGetOrDefaultFoundTestCase = coretestcases.CaseV1{
	Title:         "GetOrDefault found returns value",
	ExpectedInput: args.Map{"value": 5},
}

// ==========================================================================
// Has / Contains / IsKeyMissing
// ==========================================================================

var hashmapHasTestCase = coretestcases.CaseV1{
	Title: "Has/Contains/IsKeyMissing",
	ExpectedInput: args.Map{
		"has":          true,
		"contains":     true,
		"isKeyMissing": false,
	},
}

var hashmapIsKeyMissingTestCase = coretestcases.CaseV1{
	Title:         "IsKeyMissing true",
	ExpectedInput: args.Map{"isKeyMissing": true},
}

// ==========================================================================
// Remove
// ==========================================================================

var hashmapRemoveExistingTestCase = coretestcases.CaseV1{
	Title: "Remove existing",
	ExpectedInput: args.Map{
		"removed": true,
		"isGone":  true,
	},
}

var hashmapRemoveMissingTestCase = coretestcases.CaseV1{
	Title:         "Remove missing",
	ExpectedInput: args.Map{"removed": false},
}

// ==========================================================================
// Keys
// ==========================================================================

var hashmapKeysNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "Keys returns all",
	ExpectedInput: args.Map{"keyCount": 2},
}

var hashmapKeysEmptyTestCase = coretestcases.CaseV1{
	Title:         "Keys empty",
	ExpectedInput: args.Map{"keyCount": 0},
}

// ==========================================================================
// Values
// ==========================================================================

var hashmapValuesNonEmptyTestCase = coretestcases.CaseV1{
	Title: "Values returns all",
	ExpectedInput: args.Map{
		"valueCount":       1,
		"containsExpected": 1,
	},
}

var hashmapValuesEmptyTestCase = coretestcases.CaseV1{
	Title:         "Values empty",
	ExpectedInput: args.Map{"valueCount": 0},
}

// ==========================================================================
// AddOrUpdate
// ==========================================================================

var hashmapAddOrUpdateMapMergesTestCase = coretestcases.CaseV1{
	Title: "AddOrUpdateMap merges",
	ExpectedInput: args.Map{
		"length":      2,
		"mergedValue": 10,
	},
}

var hashmapAddOrUpdateMapEmptyNoopTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateMap empty noop",
	ExpectedInput: args.Map{"length": 1},
}

var hashmapAddOrUpdateHashmapMergesTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateHashmap merges",
	ExpectedInput: args.Map{"length": 2},
}

var hashmapAddOrUpdateHashmapNilNoopTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateHashmap nil noop",
	ExpectedInput: args.Map{"length": 1},
}

// ==========================================================================
// ConcatNew
// ==========================================================================

var hashmapConcatNewMergedTestCase = coretestcases.CaseV1{
	Title: "ConcatNew merged copy",
	ExpectedInput: args.Map{
		"mergedLength":   2,
		"originalLength": 1,
	},
}

var hashmapConcatNewNilTestCase = coretestcases.CaseV1{
	Title:         "ConcatNew nil",
	ExpectedInput: args.Map{"length": 1},
}

// ==========================================================================
// Clone method
// ==========================================================================

var hashmapCloneMethodTestCase = coretestcases.CaseV1{
	Title:         "Clone method independence",
	ExpectedInput: args.Map{"origValue": 1},
}

// ==========================================================================
// IsEquals
// ==========================================================================

var hashmapIsEqualsSameContentTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same content",
	ExpectedInput: args.Map{"isEquals": true},
}

var hashmapIsEqualsDifferentKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different keys",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashmapIsEqualsDifferentLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different length",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashmapIsEqualsBothNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals both nil",
	ExpectedInput: args.Map{"isEquals": true},
}

var hashmapIsEqualsOneNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals one nil",
	ExpectedInput: args.Map{"isEquals": false},
}

var hashmapIsEqualsSamePointerTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same pointer",
	ExpectedInput: args.Map{"isEquals": true},
}

// ==========================================================================
// ForEach
// ==========================================================================

var hashmapForEachTestCase = coretestcases.CaseV1{
	Title:         "ForEach visits all",
	ExpectedInput: args.Map{"visitCount": 2},
}

var hashmapForEachBreakTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak stops early",
	ExpectedInput: args.Map{"visitCount": 2},
}

// ==========================================================================
// String
// ==========================================================================

var hashmapStringTestCase = coretestcases.CaseV1{
	Title:         "String not empty",
	ExpectedInput: args.Map{"isNonEmpty": true},
}

// ==========================================================================
// Nil receiver
// ==========================================================================

var hashmapNilReceiverIsEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEmpty nil receiver",
	ExpectedInput: args.Map{"isEmpty": true},
}

var hashmapNilReceiverLengthTestCase = coretestcases.CaseV1{
	Title:         "Length nil receiver",
	ExpectedInput: args.Map{"length": 0},
}

var hashmapNilReceiverHasItemsTestCase = coretestcases.CaseV1{
	Title:         "HasItems nil receiver",
	ExpectedInput: args.Map{"hasItems": false},
}
