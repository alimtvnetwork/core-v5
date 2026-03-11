package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// IsEqual test cases
// =============================================================================

var pagingInfoIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "PagingInfo.IsEqual returns true -- both nil",
		ArrangeInput: args.Map{
			"when":       "given both nil",
			"isLeftNil":  true,
			"isRightNil": true,
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "PagingInfo.IsEqual returns false -- left nil right non-nil",
		ArrangeInput: args.Map{
			"when":                  "given left nil right non-nil",
			"isLeftNil":             true,
			"isRightNil":            false,
			"rightTotalPages":       5,
			"rightCurrentPageIndex": 1,
			"rightPerPageItems":     10,
			"rightTotalItems":       50,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "PagingInfo.IsEqual returns false -- left non-nil right nil",
		ArrangeInput: args.Map{
			"when":                 "given left non-nil right nil",
			"isLeftNil":            false,
			"isRightNil":           true,
			"leftTotalPages":       5,
			"leftCurrentPageIndex": 1,
			"leftPerPageItems":     10,
			"leftTotalItems":       50,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "PagingInfo.IsEqual returns true -- identical values",
		ArrangeInput: args.Map{
			"when":                  "given identical values",
			"isLeftNil":             false,
			"isRightNil":            false,
			"leftTotalPages":        3,
			"leftCurrentPageIndex":  2,
			"leftPerPageItems":      10,
			"leftTotalItems":        25,
			"rightTotalPages":       3,
			"rightCurrentPageIndex": 2,
			"rightPerPageItems":     10,
			"rightTotalItems":       25,
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "PagingInfo.IsEqual returns false -- different TotalPages",
		ArrangeInput: args.Map{
			"when":                  "given different TotalPages",
			"isLeftNil":             false,
			"isRightNil":            false,
			"leftTotalPages":        3,
			"leftCurrentPageIndex":  2,
			"leftPerPageItems":      10,
			"leftTotalItems":        25,
			"rightTotalPages":       5,
			"rightCurrentPageIndex": 2,
			"rightPerPageItems":     10,
			"rightTotalItems":       25,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "PagingInfo.IsEqual returns false -- different CurrentPageIndex",
		ArrangeInput: args.Map{
			"when":                  "given different CurrentPageIndex",
			"isLeftNil":             false,
			"isRightNil":            false,
			"leftTotalPages":        3,
			"leftCurrentPageIndex":  1,
			"leftPerPageItems":      10,
			"leftTotalItems":        25,
			"rightTotalPages":       3,
			"rightCurrentPageIndex": 2,
			"rightPerPageItems":     10,
			"rightTotalItems":       25,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "PagingInfo.IsEqual returns false -- different PerPageItems",
		ArrangeInput: args.Map{
			"when":                  "given different PerPageItems",
			"isLeftNil":             false,
			"isRightNil":            false,
			"leftTotalPages":        3,
			"leftCurrentPageIndex":  2,
			"leftPerPageItems":      10,
			"leftTotalItems":        25,
			"rightTotalPages":       3,
			"rightCurrentPageIndex": 2,
			"rightPerPageItems":     20,
			"rightTotalItems":       25,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "PagingInfo.IsEqual returns false -- different TotalItems",
		ArrangeInput: args.Map{
			"when":                  "given different TotalItems",
			"isLeftNil":             false,
			"isRightNil":            false,
			"leftTotalPages":        3,
			"leftCurrentPageIndex":  2,
			"leftPerPageItems":      10,
			"leftTotalItems":        25,
			"rightTotalPages":       3,
			"rightCurrentPageIndex": 2,
			"rightPerPageItems":     10,
			"rightTotalItems":       30,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}

// =============================================================================
// State check test cases (IsEmpty, Has*, IsInvalid*)
// =============================================================================

var pagingInfoStateTestCases = []coretestcases.CaseV1{
	// --- nil receiver ---
	{
		Title: "PagingInfo returns isEmpty true and all invalid -- nil receiver",
		ArrangeInput: args.Map{
			"when":  "given nil PagingInfo",
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isEmpty":                   true,
			"hasTotalPages":             false,
			"hasCurrentPageIndex":       false,
			"hasPerPageItems":           false,
			"hasTotalItems":             false,
			"isInvalidTotalPages":       true,
			"isInvalidCurrentPageIndex": true,
			"isInvalidPerPageItems":     true,
			"isInvalidTotalItems":       true,
		},
	},
	// --- zero values ---
	{
		Title: "PagingInfo returns isEmpty true and all invalid -- zero values",
		ArrangeInput: args.Map{
			"when":             "given zero-value PagingInfo",
			"isNil":            false,
			"totalPages":       0,
			"currentPageIndex": 0,
			"perPageItems":     0,
			"totalItems":       0,
		},
		ExpectedInput: args.Map{
			"isEmpty":                   true,
			"hasTotalPages":             false,
			"hasCurrentPageIndex":       false,
			"hasPerPageItems":           false,
			"hasTotalItems":             false,
			"isInvalidTotalPages":       true,
			"isInvalidCurrentPageIndex": true,
			"isInvalidPerPageItems":     true,
			"isInvalidTotalItems":       true,
		},
	},
	// --- all positive ---
	{
		Title: "PagingInfo returns isEmpty false and all valid -- all positive values",
		ArrangeInput: args.Map{
			"when":             "given fully populated PagingInfo",
			"isNil":            false,
			"totalPages":       5,
			"currentPageIndex": 3,
			"perPageItems":     10,
			"totalItems":       50,
		},
		ExpectedInput: args.Map{
			"isEmpty":                   false,
			"hasTotalPages":             true,
			"hasCurrentPageIndex":       true,
			"hasPerPageItems":           true,
			"hasTotalItems":             true,
			"isInvalidTotalPages":       false,
			"isInvalidCurrentPageIndex": false,
			"isInvalidPerPageItems":     false,
			"isInvalidTotalItems":       false,
		},
	},
	// --- negative TotalPages ---
	{
		Title: "Negative TotalPages - IsInvalidTotalPages true",
		ArrangeInput: args.Map{
			"when":             "given negative TotalPages",
			"isNil":            false,
			"totalPages":       -1,
			"currentPageIndex": 1,
			"perPageItems":     10,
			"totalItems":       50,
		},
		ExpectedInput: args.Map{
			"isEmpty":                   false,
			"hasTotalPages":             false,
			"hasCurrentPageIndex":       true,
			"hasPerPageItems":           true,
			"hasTotalItems":             true,
			"isInvalidTotalPages":       true,
			"isInvalidCurrentPageIndex": false,
			"isInvalidPerPageItems":     false,
			"isInvalidTotalItems":       false,
		},
	},
	// --- partial: only TotalPages set ---
	{
		Title: "Only TotalPages set - HasTotalPages true others false",
		ArrangeInput: args.Map{
			"when":             "given only TotalPages populated",
			"isNil":            false,
			"totalPages":       3,
			"currentPageIndex": 0,
			"perPageItems":     0,
			"totalItems":       0,
		},
		ExpectedInput: args.Map{
			"isEmpty":                   false,
			"hasTotalPages":             true,
			"hasCurrentPageIndex":       false,
			"hasPerPageItems":           false,
			"hasTotalItems":             false,
			"isInvalidTotalPages":       false,
			"isInvalidCurrentPageIndex": true,
			"isInvalidPerPageItems":     true,
			"isInvalidTotalItems":       true,
		},
	},
}

// =============================================================================
// Clone test cases
// =============================================================================

var pagingInfoCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone copies all fields correctly",
		ArrangeInput: args.Map{
			"when":             "given fully populated PagingInfo",
			"totalPages":       5,
			"currentPageIndex": 3,
			"perPageItems":     10,
			"totalItems":       50,
		},
		ExpectedInput: args.Map{
			"totalPages":       5,
			"currentPageIndex": 3,
			"perPageItems":     10,
			"totalItems":       50,
		},
	},
	{
		Title: "Clone of zero-value copies zeros",
		ArrangeInput: args.Map{
			"when":             "given zero-value PagingInfo",
			"totalPages":       0,
			"currentPageIndex": 0,
			"perPageItems":     0,
			"totalItems":       0,
		},
		ExpectedInput: args.Map{
			"totalPages":       0,
			"currentPageIndex": 0,
			"perPageItems":     0,
			"totalItems":       0,
		},
	},
}

// =============================================================================
// ClonePtr test cases
// =============================================================================

var pagingInfoClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr nil receiver returns nil",
		ArrangeInput: args.Map{
			"when":  "given nil PagingInfo pointer",
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "ClonePtr copies all fields",
		ArrangeInput: args.Map{
			"when":             "given populated PagingInfo pointer",
			"isNil":            false,
			"totalPages":       5,
			"currentPageIndex": 3,
			"perPageItems":     10,
			"totalItems":       50,
		},
		ExpectedInput: args.Map{
			"isNil":            false,
			"totalPages":       5,
			"currentPageIndex": 3,
			"perPageItems":     10,
			"totalItems":       50,
		},
	},
}

// =============================================================================
// Helper: build PagingInfo from args.Map
// =============================================================================

func buildPagingInfoFromMap(input args.Map) *corepayload.PagingInfo {
	totalPages, _ := input.GetAsInt("totalPages")
	currentPageIndex, _ := input.GetAsInt("currentPageIndex")
	perPageItems, _ := input.GetAsInt("perPageItems")
	totalItems, _ := input.GetAsInt("totalItems")

	return &corepayload.PagingInfo{
		TotalPages:       totalPages,
		CurrentPageIndex: currentPageIndex,
		PerPageItems:     perPageItems,
		TotalItems:       totalItems,
	}
}

func buildPagingInfoPrefixed(input args.Map, prefix string) *corepayload.PagingInfo {
	totalPages, _ := input.GetAsInt(prefix + "TotalPages")
	currentPageIndex, _ := input.GetAsInt(prefix + "CurrentPageIndex")
	perPageItems, _ := input.GetAsInt(prefix + "PerPageItems")
	totalItems, _ := input.GetAsInt(prefix + "TotalItems")

	return &corepayload.PagingInfo{
		TotalPages:       totalPages,
		CurrentPageIndex: currentPageIndex,
		PerPageItems:     perPageItems,
		TotalItems:       totalItems,
	}
}

// ==========================================================================
// PagingInfo — Clone independence
// ==========================================================================

var pagingInfoClonePtrIndependenceTestCase = coretestcases.CaseV1{
	Title: "ClonePtr creates independent copy",
	ExpectedInput: args.Map{
		"originalTotalPages":  5,
		"originalCurrentPage": 3,
	},
}

var pagingInfoCloneIndependenceTestCase = coretestcases.CaseV1{
	Title:         "Clone creates independent value copy",
	ExpectedInput: args.Map{"originalTotalPages": 5},
}
