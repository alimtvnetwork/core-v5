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
	ExpectedInput: args.Three[string, string, string]{
		First:  "true",  // isEmpty
		Second: "0",     // length
		Third:  "false", // hasItems
	},
}

var hashmapNewTestCase = coretestcases.CaseV1{
	Title:         "NewHashmap with capacity",
	ExpectedInput: "true", // isNotNil
}

var hashmapFromTestCase = coretestcases.CaseV1{
	Title: "HashmapFrom wraps map",
	ExpectedInput: args.Two[string, string]{
		First:  "2",    // length
		Second: "true", // hasKey
	},
}

var hashmapCloneFuncTestCase = coretestcases.CaseV1{
	Title: "HashmapClone independence",
	ExpectedInput: args.Two[string, string]{
		First:  "1",  // length
		Second: "99", // originalValue
	},
}

// ==========================================================================
// Set
// ==========================================================================

var hashmapSetNewTestCase = coretestcases.CaseV1{
	Title: "Set new key returns true",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // isNew
		Second: "1",    // length
	},
}

var hashmapSetExistingTestCase = coretestcases.CaseV1{
	Title: "Set existing key returns false",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // isNew
		Second: "2",     // updatedValue
	},
}

// ==========================================================================
// Get
// ==========================================================================

var hashmapGetFoundTestCase = coretestcases.CaseV1{
	Title: "Get found",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // found
		Second: "42",   // value
	},
}

var hashmapGetNotFoundTestCase = coretestcases.CaseV1{
	Title: "Get not found",
	ExpectedInput: args.Two[string, string]{
		First:  "false", // found
		Second: "0",     // zeroValue
	},
}

var hashmapGetOrDefaultMissingTestCase = coretestcases.CaseV1{
	Title:         "GetOrDefault missing returns default",
	ExpectedInput: "99", // defaultValue
}

var hashmapGetOrDefaultFoundTestCase = coretestcases.CaseV1{
	Title:         "GetOrDefault found returns value",
	ExpectedInput: "5", // value
}

// ==========================================================================
// Has / Contains / IsKeyMissing
// ==========================================================================

var hashmapHasTestCase = coretestcases.CaseV1{
	Title: "Has/Contains/IsKeyMissing",
	ExpectedInput: args.Three[string, string, string]{
		First:  "true",  // has
		Second: "true",  // contains
		Third:  "false", // isKeyMissing
	},
}

var hashmapIsKeyMissingTestCase = coretestcases.CaseV1{
	Title:         "IsKeyMissing true",
	ExpectedInput: "true",
}

// ==========================================================================
// Remove
// ==========================================================================

var hashmapRemoveExistingTestCase = coretestcases.CaseV1{
	Title: "Remove existing",
	ExpectedInput: args.Two[string, string]{
		First:  "true", // removed
		Second: "true", // isGone
	},
}

var hashmapRemoveMissingTestCase = coretestcases.CaseV1{
	Title:         "Remove missing",
	ExpectedInput: "false", // removed
}

// ==========================================================================
// Keys
// ==========================================================================

var hashmapKeysNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "Keys returns all",
	ExpectedInput: "2", // keyCount
}

var hashmapKeysEmptyTestCase = coretestcases.CaseV1{
	Title:         "Keys empty",
	ExpectedInput: "0", // keyCount
}

// ==========================================================================
// Values
// ==========================================================================

var hashmapValuesNonEmptyTestCase = coretestcases.CaseV1{
	Title: "Values returns all",
	ExpectedInput: args.Two[string, string]{
		First:  "1", // value1
		Second: "1", // containsExpected
	},
}

var hashmapValuesEmptyTestCase = coretestcases.CaseV1{
	Title:         "Values empty",
	ExpectedInput: "0", // valueCount
}

// ==========================================================================
// AddOrUpdate
// ==========================================================================

var hashmapAddOrUpdateMapMergesTestCase = coretestcases.CaseV1{
	Title: "AddOrUpdateMap merges",
	ExpectedInput: args.Two[string, string]{
		First:  "2",  // length
		Second: "10", // mergedValue
	},
}

var hashmapAddOrUpdateMapEmptyNoopTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateMap empty noop",
	ExpectedInput: "1", // length
}

var hashmapAddOrUpdateHashmapMergesTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateHashmap merges",
	ExpectedInput: "2", // length
}

var hashmapAddOrUpdateHashmapNilNoopTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateHashmap nil noop",
	ExpectedInput: "1", // length
}

// ==========================================================================
// ConcatNew
// ==========================================================================

var hashmapConcatNewMergedTestCase = coretestcases.CaseV1{
	Title: "ConcatNew merged copy",
	ExpectedInput: args.Two[string, string]{
		First:  "2", // mergedLength
		Second: "1", // originalLength
	},
}

var hashmapConcatNewNilTestCase = coretestcases.CaseV1{
	Title:         "ConcatNew nil",
	ExpectedInput: "1", // length
}

// ==========================================================================
// Clone method
// ==========================================================================

var hashmapCloneMethodTestCase = coretestcases.CaseV1{
	Title:         "Clone method independence",
	ExpectedInput: "1", // originalLength
}

// ==========================================================================
// IsEquals
// ==========================================================================

var hashmapIsEqualsSameContentTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same content",
	ExpectedInput: "true",
}

var hashmapIsEqualsDifferentKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different keys",
	ExpectedInput: "false",
}

var hashmapIsEqualsDifferentLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different length",
	ExpectedInput: "false",
}

var hashmapIsEqualsBothNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals both nil",
	ExpectedInput: "true",
}

var hashmapIsEqualsOneNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals one nil",
	ExpectedInput: "false",
}

var hashmapIsEqualsSamePointerTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same pointer",
	ExpectedInput: "true",
}

// ==========================================================================
// ForEach
// ==========================================================================

var hashmapForEachTestCase = coretestcases.CaseV1{
	Title:         "ForEach visits all",
	ExpectedInput: "2", // visitCount
}

var hashmapForEachBreakTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak stops early",
	ExpectedInput: "2", // visitCount
}

// ==========================================================================
// String
// ==========================================================================

var hashmapStringTestCase = coretestcases.CaseV1{
	Title:         "String not empty",
	ExpectedInput: "true", // isNonEmpty
}

// ==========================================================================
// Nil receiver
// ==========================================================================

var hashmapNilReceiverIsEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEmpty nil receiver",
	ExpectedInput: "true",
}

var hashmapNilReceiverLengthTestCase = coretestcases.CaseV1{
	Title:         "Length nil receiver",
	ExpectedInput: "0",
}

var hashmapNilReceiverHasItemsTestCase = coretestcases.CaseV1{
	Title:         "HasItems nil receiver",
	ExpectedInput: "false",
}
