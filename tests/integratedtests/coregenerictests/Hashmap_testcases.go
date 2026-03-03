package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Constructors
// ==========================================================================

var hashmapEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "EmptyHashmap creates empty",
		ExpectedInput: []string{"true", "0", "false"},
	},
}

var hashmapNewTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewHashmap with capacity",
		ExpectedInput: []string{"true"},
	},
}

var hashmapFromTestCases = []coretestcases.CaseV1{
	{
		Title:         "HashmapFrom wraps map",
		ExpectedInput: []string{"2", "true"},
	},
}

var hashmapCloneFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "HashmapClone independence",
		ExpectedInput: []string{"1", "99"},
	},
}

// ==========================================================================
// Set
// ==========================================================================

var hashmapSetNewTestCases = []coretestcases.CaseV1{
	{
		Title:         "Set new key returns true",
		ExpectedInput: []string{"true", "1"},
	},
}

var hashmapSetExistingTestCases = []coretestcases.CaseV1{
	{
		Title:         "Set existing key returns false",
		ExpectedInput: []string{"false", "2"},
	},
}

// ==========================================================================
// Get
// ==========================================================================

var hashmapGetFoundTestCases = []coretestcases.CaseV1{
	{
		Title:         "Get found",
		ExpectedInput: []string{"true", "42"},
	},
}

var hashmapGetNotFoundTestCases = []coretestcases.CaseV1{
	{
		Title:         "Get not found",
		ExpectedInput: []string{"false", "0"},
	},
}

var hashmapGetOrDefaultMissingTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetOrDefault missing returns default",
		ExpectedInput: []string{"99"},
	},
}

var hashmapGetOrDefaultFoundTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetOrDefault found returns value",
		ExpectedInput: []string{"5"},
	},
}

// ==========================================================================
// Has / Contains / IsKeyMissing
// ==========================================================================

var hashmapHasTestCases = []coretestcases.CaseV1{
	{
		Title:         "Has/Contains/IsKeyMissing",
		ExpectedInput: []string{"true", "true", "false"},
	},
}

var hashmapIsKeyMissingTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsKeyMissing true",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// Remove
// ==========================================================================

var hashmapRemoveExistingTestCases = []coretestcases.CaseV1{
	{
		Title:         "Remove existing",
		ExpectedInput: []string{"true", "true"},
	},
}

var hashmapRemoveMissingTestCases = []coretestcases.CaseV1{
	{
		Title:         "Remove missing",
		ExpectedInput: []string{"false"},
	},
}

// ==========================================================================
// Keys / Values
// ==========================================================================

var hashmapKeysTestCases = []coretestcases.CaseV1{
	{
		Title:         "Keys returns all",
		ExpectedInput: []string{"2"},
	},
	{
		Title:         "Keys empty",
		ExpectedInput: []string{"0"},
	},
}

var hashmapValuesTestCases = []coretestcases.CaseV1{
	{
		Title:         "Values returns all",
		ExpectedInput: []string{"1", "1"},
	},
	{
		Title:         "Values empty",
		ExpectedInput: []string{"0"},
	},
}

// ==========================================================================
// AddOrUpdate
// ==========================================================================

var hashmapAddOrUpdateMapTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddOrUpdateMap merges",
		ExpectedInput: []string{"2", "10"},
	},
	{
		Title:         "AddOrUpdateMap empty noop",
		ExpectedInput: []string{"1"},
	},
}

var hashmapAddOrUpdateHashmapTestCases = []coretestcases.CaseV1{
	{
		Title:         "AddOrUpdateHashmap merges",
		ExpectedInput: []string{"2"},
	},
	{
		Title:         "AddOrUpdateHashmap nil noop",
		ExpectedInput: []string{"1"},
	},
}

// ==========================================================================
// ConcatNew
// ==========================================================================

var hashmapConcatNewTestCases = []coretestcases.CaseV1{
	{
		Title:         "ConcatNew merged copy",
		ExpectedInput: []string{"2", "1"},
	},
	{
		Title:         "ConcatNew nil",
		ExpectedInput: []string{"1"},
	},
}

// ==========================================================================
// Clone method
// ==========================================================================

var hashmapCloneMethodTestCases = []coretestcases.CaseV1{
	{
		Title:         "Clone method independence",
		ExpectedInput: []string{"1"},
	},
}

// ==========================================================================
// IsEquals
// ==========================================================================

var hashmapIsEqualsTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsEquals same content",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "IsEquals different keys",
		ExpectedInput: []string{"false"},
	},
	{
		Title:         "IsEquals different length",
		ExpectedInput: []string{"false"},
	},
	{
		Title:         "IsEquals both nil",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "IsEquals one nil",
		ExpectedInput: []string{"false"},
	},
	{
		Title:         "IsEquals same pointer",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// ForEach
// ==========================================================================

var hashmapForEachTestCases = []coretestcases.CaseV1{
	{
		Title:         "ForEach visits all",
		ExpectedInput: []string{"2"},
	},
}

var hashmapForEachBreakTestCases = []coretestcases.CaseV1{
	{
		Title:         "ForEachBreak stops early",
		ExpectedInput: []string{"2"},
	},
}

// ==========================================================================
// String
// ==========================================================================

var hashmapStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "String not empty",
		ExpectedInput: []string{"true"},
	},
}

// ==========================================================================
// Nil receiver
// ==========================================================================

var hashmapNilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsEmpty nil receiver",
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "Length nil receiver",
		ExpectedInput: []string{"0"},
	},
	{
		Title:         "HasItems nil receiver",
		ExpectedInput: []string{"false"},
	},
}
