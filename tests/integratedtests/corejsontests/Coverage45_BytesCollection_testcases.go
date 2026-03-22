package corejsontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// BytesCollection — Length / IsEmpty / HasAnyItem / nil receiver
// =============================================================================

var bytesCollNilLengthTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Length returns 0 -- nil receiver",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollIsEmptyTrueTestCase = coretestcases.CaseV1{
	Title: "BytesCollection IsEmpty returns true -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var bytesCollHasAnyItemTestCase = coretestcases.CaseV1{
	Title: "BytesCollection HasAnyItem returns true -- one item added",
	ExpectedInput: args.Map{
		"hasAny": true,
	},
}

var bytesCollLastIndexTestCase = coretestcases.CaseV1{
	Title: "BytesCollection LastIndex returns 1 -- two items",
	ExpectedInput: args.Map{
		"lastIndex": 1,
	},
}

// =============================================================================
// BytesCollection — FirstOrDefault / LastOrDefault
// =============================================================================

var bytesCollFirstOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection FirstOrDefault returns nil -- empty collection",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var bytesCollFirstOrDefaultHasItemTestCase = coretestcases.CaseV1{
	Title: "BytesCollection FirstOrDefault returns non-nil -- one item",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

var bytesCollLastOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection LastOrDefault returns nil -- empty collection",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var bytesCollLastOrDefaultHasItemTestCase = coretestcases.CaseV1{
	Title: "BytesCollection LastOrDefault returns non-nil -- two items",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

// =============================================================================
// BytesCollection — Take / Limit / Skip
// =============================================================================

var bytesCollTakeEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Take returns empty -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var bytesCollTakeValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Take returns 2 -- take first 2 of 3",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var bytesCollLimitTakeAllTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Limit returns all -- limit -1 means take all",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var bytesCollLimitEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Limit returns empty -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var bytesCollSkipEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Skip returns empty -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var bytesCollSkipValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Skip returns 2 -- skip 1 of 3",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

// =============================================================================
// BytesCollection — Add variants
// =============================================================================

var bytesCollAddSkipOnNilNilTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddSkipOnNil returns 0 -- nil bytes skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddSkipOnNilValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddSkipOnNil returns 1 -- valid bytes added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollAddNonEmptyEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddNonEmpty returns 0 -- empty bytes skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddNonEmptyValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddNonEmpty returns 1 -- valid bytes added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollAddResultPtrSkipTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddResultPtr returns 0 -- empty result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddResultPtrValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddResultPtr returns 1 -- valid result added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollAddResultSkipTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddResult returns 0 -- empty result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddResultValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddResult returns 1 -- valid result added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollAddPtrEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddPtr returns 0 -- empty bytes skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddPtrValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddPtr returns 1 -- valid bytes added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollAddsEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Adds returns 0 -- no items provided",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddsValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Adds returns 2 -- two valid items added",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var bytesCollAddsPtrNilTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddsPtr returns 0 -- nil results skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// =============================================================================
// BytesCollection — AddAny / AddAnyItems / AddSerializerFunc
// =============================================================================

var bytesCollAddAnyValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddAny returns no error -- valid string",
	ExpectedInput: args.Map{
		"hasError": false,
		"length":   1,
	},
}

var bytesCollAddAnyItemsEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddAnyItems returns no error -- empty items",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var bytesCollAddAnyItemsValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddAnyItems returns no error -- two valid items",
	ExpectedInput: args.Map{
		"hasError": false,
		"length":   2,
	},
}

var bytesCollAddSerializerFuncNilTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddSerializerFunc returns 0 -- nil func skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddSerializerFunctionsEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddSerializerFunctions returns 0 -- no funcs",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// =============================================================================
// BytesCollection — GetAtSafe / GetResultAtSafe / GetAtSafeUsingLength
// =============================================================================

var bytesCollGetAtSafeOutOfRangeTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetAtSafe returns nil -- index out of range",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var bytesCollGetAtSafeValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetAtSafe returns bytes -- valid index",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

var bytesCollGetResultAtSafeOutOfRangeTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetResultAtSafe returns nil -- index out of range",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var bytesCollGetAtSafeUsingLengthOutTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetAtSafeUsingLength returns nil -- index out of range",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

// =============================================================================
// BytesCollection — Clear / Dispose / Clone / ClonePtr
// =============================================================================

var bytesCollClearTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Clear returns 0 -- items cleared",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollClearNilTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Clear returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var bytesCollDisposeTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Dispose sets items nil -- items disposed",
	ExpectedInput: args.Map{
		"nilItems": true,
	},
}

var bytesCollCloneTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Clone returns 1 -- shallow clone with 1 item",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollClonePtrNilTestCase = coretestcases.CaseV1{
	Title: "BytesCollection ClonePtr returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var bytesCollClonePtrValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection ClonePtr returns 1 -- deep clone with 1 item",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

// =============================================================================
// BytesCollection — Strings / Json / Paging / UnmarshalAt
// =============================================================================

var bytesCollStringsEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Strings returns empty -- empty collection",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollStringsValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Strings returns 2 -- two items",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var bytesCollJsonTestCase = coretestcases.CaseV1{
	Title: "BytesCollection Json returns no error -- valid collection",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var bytesCollGetPagesSizeZeroTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetPagesSize returns 0 -- page size 0",
	ExpectedInput: args.Map{
		"pages": 0,
	},
}

var bytesCollGetPagesSizeValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetPagesSize returns 2 -- 3 items page size 2",
	ExpectedInput: args.Map{
		"pages": 2,
	},
}

var bytesCollUnmarshalAtTestCase = coretestcases.CaseV1{
	Title: "BytesCollection UnmarshalAt returns hello -- index 0 valid bytes",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var bytesCollGetAtTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetAt returns non-nil -- valid index",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

var bytesCollJsonResultAtTestCase = coretestcases.CaseV1{
	Title: "BytesCollection JsonResultAt returns non-nil -- valid index",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

// =============================================================================
// BytesCollection — AddBytesCollection / AddMapResults / AddRawMapResults
// =============================================================================

var bytesCollAddBytesCollectionEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddBytesCollection returns 0 -- empty source",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var bytesCollAddBytesCollectionValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddBytesCollection returns 1 -- one source item",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollAddRawMapResultsEmptyTestCase = coretestcases.CaseV1{
	Title: "BytesCollection AddRawMapResults returns 0 -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// =============================================================================
// BytesCollection — ParseInjectUsingJson / ShadowClone
// =============================================================================

var bytesCollShadowCloneTestCase = coretestcases.CaseV1{
	Title: "BytesCollection ShadowClone returns 1 -- shallow clone",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var bytesCollParseInjectUsingJsonValidTestCase = coretestcases.CaseV1{
	Title: "BytesCollection ParseInjectUsingJson returns no error -- valid result",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// BytesCollection — UnmarshalIntoSameIndex
// =============================================================================

var bytesCollUnmarshalIntoSameIndexNilTestCase = coretestcases.CaseV1{
	Title: "BytesCollection UnmarshalIntoSameIndex returns no error -- nil anys",
	ExpectedInput: args.Map{
		"hasAnyErr": false,
	},
}

var bytesCollGetAtSafePtrOutOfRangeTestCase = coretestcases.CaseV1{
	Title: "BytesCollection GetAtSafePtr returns nil -- index out of range",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}
