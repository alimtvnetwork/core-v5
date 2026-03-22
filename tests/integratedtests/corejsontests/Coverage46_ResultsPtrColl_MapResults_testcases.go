package corejsontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// ResultsPtrCollection — Length / IsEmpty / HasAnyItem
// =============================================================================

var ptrCollNilLengthTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Length returns 0 -- nil receiver",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollIsEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection IsEmpty returns true -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var ptrCollHasAnyItemTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection HasAnyItem returns true -- one item",
	ExpectedInput: args.Map{
		"hasAny": true,
	},
}

// =============================================================================
// ResultsPtrCollection — FirstOrDefault / LastOrDefault
// =============================================================================

var ptrCollFirstOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection FirstOrDefault returns nil -- empty",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var ptrCollFirstOrDefaultValidTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection FirstOrDefault returns non-nil -- one item",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

var ptrCollLastOrDefaultEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection LastOrDefault returns nil -- empty",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var ptrCollLastOrDefaultValidTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection LastOrDefault returns non-nil -- one item",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

// =============================================================================
// ResultsPtrCollection — Take / Limit / Skip
// =============================================================================

var ptrCollTakeEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Take returns empty -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var ptrCollTakeValidTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Take returns 2 -- take first 2 of 3",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var ptrCollLimitTakeAllTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Limit returns all -- limit -1",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

var ptrCollSkipEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Skip returns empty -- empty collection",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var ptrCollSkipValidTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Skip returns 2 -- skip 1 of 3",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

// =============================================================================
// ResultsPtrCollection — Add variants
// =============================================================================

var ptrCollAddSkipOnNilNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddSkipOnNil returns 0 -- nil skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollAddSkipOnNilValidTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddSkipOnNil returns 1 -- valid added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var ptrCollAddNonNilNonErrorNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddNonNilNonError returns 0 -- nil skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollAddNonNilNonErrorValidTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddNonNilNonError returns 1 -- valid added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var ptrCollAddResultTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddResult returns 1 -- value result added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var ptrCollAddsNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Adds returns 1 -- nil element appended via variadic",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var ptrCollAddAnyNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddAny returns 0 -- nil skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollAddAnyValidTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddAny returns 1 -- valid item",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var ptrCollAddAnyItemsNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddAnyItems returns 0 -- nil items",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollAddResultsCollectionNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AddResultsCollection returns 0 -- nil collection",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// =============================================================================
// ResultsPtrCollection — HasError / AllErrors / GetErrorsStrings
// =============================================================================

var ptrCollHasErrorFalseTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection HasError returns false -- no errors",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var ptrCollHasErrorTrueTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection HasError returns true -- one error item",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var ptrCollAllErrorsEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection AllErrors returns empty -- empty collection",
	ExpectedInput: args.Map{
		"errorCount": 0,
		"hasAnyErr":  false,
	},
}

var ptrCollGetErrorsStringsEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection GetErrorsStrings returns empty -- empty collection",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollGetErrorsAsSingleStringTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection GetErrorsAsSingleString returns non-empty -- errors present",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

// =============================================================================
// ResultsPtrCollection — GetAtSafe / Clear / Dispose / Clone / GetStrings
// =============================================================================

var ptrCollGetAtSafeOutOfRangeTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection GetAtSafe returns nil -- index out of range",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var ptrCollGetAtSafeUsingLengthOutTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection GetAtSafeUsingLength returns nil -- out of range",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var ptrCollClearTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Clear returns 0 -- items cleared",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollClearNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Clear returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var ptrCollDisposeTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Dispose sets items nil -- disposed",
	ExpectedInput: args.Map{
		"nilItems": true,
	},
}

var ptrCollGetStringsEmptyTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection GetStrings returns empty -- empty collection",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var ptrCollGetPagesSizeZeroTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection GetPagesSize returns 0 -- page size 0",
	ExpectedInput: args.Map{
		"pages": 0,
	},
}

var ptrCollCloneNilTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Clone returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var ptrCollJsonTestCase = coretestcases.CaseV1{
	Title: "ResultsPtrCollection Json returns no error -- valid collection",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// MapResults — Length / IsEmpty / HasAnyItem
// =============================================================================

var mapResultsNilLengthTestCase = coretestcases.CaseV1{
	Title: "MapResults Length returns 0 -- nil receiver",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsIsEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults IsEmpty returns true -- empty map",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var mapResultsHasAnyItemTestCase = coretestcases.CaseV1{
	Title: "MapResults HasAnyItem returns true -- one item",
	ExpectedInput: args.Map{
		"hasAny": true,
	},
}

// =============================================================================
// MapResults — GetByKey / AddSkipOnNil / HasError / AllErrors
// =============================================================================

var mapResultsGetByKeyMissingTestCase = coretestcases.CaseV1{
	Title: "MapResults GetByKey returns nil -- key not found",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var mapResultsGetByKeyFoundTestCase = coretestcases.CaseV1{
	Title: "MapResults GetByKey returns non-nil -- key found",
	ExpectedInput: args.Map{
		"isNil": false,
	},
}

var mapResultsAddSkipOnNilNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddSkipOnNil returns 0 -- nil result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAddSkipOnNilValidTestCase = coretestcases.CaseV1{
	Title: "MapResults AddSkipOnNil returns 1 -- valid result added",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var mapResultsHasErrorFalseTestCase = coretestcases.CaseV1{
	Title: "MapResults HasError returns false -- no errors",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var mapResultsAllErrorsEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults AllErrors returns empty -- empty map",
	ExpectedInput: args.Map{
		"errorCount": 0,
		"hasAnyErr":  false,
	},
}

var mapResultsGetErrorsStringsEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults GetErrorsStrings returns empty -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// =============================================================================
// MapResults — Add / AddPtr / AddAny / AddAnySkipOnNil
// =============================================================================

var mapResultsAddTestCase = coretestcases.CaseV1{
	Title: "MapResults Add returns 1 -- result added with key",
	ExpectedInput: args.Map{
		"length": 1,
	},
}

var mapResultsAddPtrNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddPtr returns 0 -- nil result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAddAnyNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddAny returns error -- nil item",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var mapResultsAddAnyValidTestCase = coretestcases.CaseV1{
	Title: "MapResults AddAny returns no error -- valid item",
	ExpectedInput: args.Map{
		"hasError": false,
		"length":   1,
	},
}

var mapResultsAddAnySkipOnNilNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddAnySkipOnNil returns no error -- nil skipped",
	ExpectedInput: args.Map{
		"hasError": false,
		"length":   0,
	},
}

// =============================================================================
// MapResults — AllKeys / AllKeysSorted / AllValues / Clear / Dispose
// =============================================================================

var mapResultsAllKeysEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults AllKeys returns empty -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAllKeysSortedEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults AllKeysSorted returns empty -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAllValuesEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults AllValues returns empty -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsClearTestCase = coretestcases.CaseV1{
	Title: "MapResults Clear returns 0 -- items cleared",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsClearNilTestCase = coretestcases.CaseV1{
	Title: "MapResults Clear returns nil -- nil receiver",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var mapResultsDisposeTestCase = coretestcases.CaseV1{
	Title: "MapResults Dispose sets items nil -- disposed",
	ExpectedInput: args.Map{
		"nilItems": true,
	},
}

var mapResultsJsonTestCase = coretestcases.CaseV1{
	Title: "MapResults Json returns no error -- valid map",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var mapResultsGetPagesSizeZeroTestCase = coretestcases.CaseV1{
	Title: "MapResults GetPagesSize returns 0 -- page size 0",
	ExpectedInput: args.Map{
		"pages": 0,
	},
}

var mapResultsAddMapResultsNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddMapResults returns 0 -- nil source",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAddMapAnyItemsEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults AddMapAnyItems returns 0 -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsResultCollectionEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults ResultCollection returns empty -- empty map",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var mapResultsAllResultsCollectionEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults AllResultsCollection returns empty -- empty map",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var mapResultsGetStringsEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults GetStrings returns empty -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAddJsonerNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddJsoner returns 0 -- nil jsoner skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAddNonEmptyNonErrorPtrNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddNonEmptyNonErrorPtr returns 0 -- nil result skipped",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsGetNewMapUsingKeysEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults GetNewMapUsingKeys returns empty -- no keys",
	ExpectedInput: args.Map{
		"isEmpty": true,
	},
}

var mapResultsAddMapResultsUsingCloneOptionEmptyTestCase = coretestcases.CaseV1{
	Title: "MapResults AddMapResultsUsingCloneOption returns 0 -- empty map",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

var mapResultsAddKeysWithJsonersNilTestCase = coretestcases.CaseV1{
	Title: "MapResults AddKeysWithJsoners returns nil -- nil keysWithJsoners",
	ExpectedInput: args.Map{
		"isNil": true,
	},
}
