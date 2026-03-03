package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Constructors
// ==========================================================================

var hashmapEmptyTestCase = coretestcases.CaseV1{
	Title:         "EmptyHashmap creates empty",
	ExpectedInput: []string{"true", "0", "false"},
}

var hashmapNewTestCase = coretestcases.CaseV1{
	Title:         "NewHashmap with capacity",
	ExpectedInput: []string{"true"},
}

var hashmapFromTestCase = coretestcases.CaseV1{
	Title:         "HashmapFrom wraps map",
	ExpectedInput: []string{"2", "true"},
}

var hashmapCloneFuncTestCase = coretestcases.CaseV1{
	Title:         "HashmapClone independence",
	ExpectedInput: []string{"1", "99"},
}

// ==========================================================================
// Set
// ==========================================================================

var hashmapSetNewTestCase = coretestcases.CaseV1{
	Title:         "Set new key returns true",
	ExpectedInput: []string{"true", "1"},
}

var hashmapSetExistingTestCase = coretestcases.CaseV1{
	Title:         "Set existing key returns false",
	ExpectedInput: []string{"false", "2"},
}

// ==========================================================================
// Get
// ==========================================================================

var hashmapGetFoundTestCase = coretestcases.CaseV1{
	Title:         "Get found",
	ExpectedInput: []string{"true", "42"},
}

var hashmapGetNotFoundTestCase = coretestcases.CaseV1{
	Title:         "Get not found",
	ExpectedInput: []string{"false", "0"},
}

var hashmapGetOrDefaultMissingTestCase = coretestcases.CaseV1{
	Title:         "GetOrDefault missing returns default",
	ExpectedInput: []string{"99"},
}

var hashmapGetOrDefaultFoundTestCase = coretestcases.CaseV1{
	Title:         "GetOrDefault found returns value",
	ExpectedInput: []string{"5"},
}

// ==========================================================================
// Has / Contains / IsKeyMissing
// ==========================================================================

var hashmapHasTestCase = coretestcases.CaseV1{
	Title:         "Has/Contains/IsKeyMissing",
	ExpectedInput: []string{"true", "true", "false"},
}

var hashmapIsKeyMissingTestCase = coretestcases.CaseV1{
	Title:         "IsKeyMissing true",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Remove
// ==========================================================================

var hashmapRemoveExistingTestCase = coretestcases.CaseV1{
	Title:         "Remove existing",
	ExpectedInput: []string{"true", "true"},
}

var hashmapRemoveMissingTestCase = coretestcases.CaseV1{
	Title:         "Remove missing",
	ExpectedInput: []string{"false"},
}

// ==========================================================================
// Keys
// ==========================================================================

var hashmapKeysNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "Keys returns all",
	ExpectedInput: []string{"2"},
}

var hashmapKeysEmptyTestCase = coretestcases.CaseV1{
	Title:         "Keys empty",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// Values
// ==========================================================================

var hashmapValuesNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "Values returns all",
	ExpectedInput: []string{"1", "1"},
}

var hashmapValuesEmptyTestCase = coretestcases.CaseV1{
	Title:         "Values empty",
	ExpectedInput: []string{"0"},
}

// ==========================================================================
// AddOrUpdate
// ==========================================================================

var hashmapAddOrUpdateMapMergesTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateMap merges",
	ExpectedInput: []string{"2", "10"},
}

var hashmapAddOrUpdateMapEmptyNoopTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateMap empty noop",
	ExpectedInput: []string{"1"},
}

var hashmapAddOrUpdateHashmapMergesTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateHashmap merges",
	ExpectedInput: []string{"2"},
}

var hashmapAddOrUpdateHashmapNilNoopTestCase = coretestcases.CaseV1{
	Title:         "AddOrUpdateHashmap nil noop",
	ExpectedInput: []string{"1"},
}

// ==========================================================================
// ConcatNew
// ==========================================================================

var hashmapConcatNewMergedTestCase = coretestcases.CaseV1{
	Title:         "ConcatNew merged copy",
	ExpectedInput: []string{"2", "1"},
}

var hashmapConcatNewNilTestCase = coretestcases.CaseV1{
	Title:         "ConcatNew nil",
	ExpectedInput: []string{"1"},
}

// ==========================================================================
// Clone method
// ==========================================================================

var hashmapCloneMethodTestCase = coretestcases.CaseV1{
	Title:         "Clone method independence",
	ExpectedInput: []string{"1"},
}

// ==========================================================================
// IsEquals
// ==========================================================================

var hashmapIsEqualsSameContentTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same content",
	ExpectedInput: []string{"true"},
}

var hashmapIsEqualsDifferentKeysTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different keys",
	ExpectedInput: []string{"false"},
}

var hashmapIsEqualsDifferentLengthTestCase = coretestcases.CaseV1{
	Title:         "IsEquals different length",
	ExpectedInput: []string{"false"},
}

var hashmapIsEqualsBothNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals both nil",
	ExpectedInput: []string{"true"},
}

var hashmapIsEqualsOneNilTestCase = coretestcases.CaseV1{
	Title:         "IsEquals one nil",
	ExpectedInput: []string{"false"},
}

var hashmapIsEqualsSamePointerTestCase = coretestcases.CaseV1{
	Title:         "IsEquals same pointer",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// ForEach
// ==========================================================================

var hashmapForEachTestCase = coretestcases.CaseV1{
	Title:         "ForEach visits all",
	ExpectedInput: []string{"2"},
}

var hashmapForEachBreakTestCase = coretestcases.CaseV1{
	Title:         "ForEachBreak stops early",
	ExpectedInput: []string{"2"},
}

// ==========================================================================
// String
// ==========================================================================

var hashmapStringTestCase = coretestcases.CaseV1{
	Title:         "String not empty",
	ExpectedInput: []string{"true"},
}

// ==========================================================================
// Nil receiver
// ==========================================================================

var hashmapNilReceiverIsEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsEmpty nil receiver",
	ExpectedInput: []string{"true"},
}

var hashmapNilReceiverLengthTestCase = coretestcases.CaseV1{
	Title:         "Length nil receiver",
	ExpectedInput: []string{"0"},
}

var hashmapNilReceiverHasItemsTestCase = coretestcases.CaseV1{
	Title:         "HasItems nil receiver",
	ExpectedInput: []string{"false"},
}
