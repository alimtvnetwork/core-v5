package corepayloadtests

import (
	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// IsEqual test cases
// =============================================================================

var pagingInfoIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqual - both nil should return true",
		ArrangeInput: args.Map{
			"when":      "given both nil",
			"isLeftNil": true,
			"isRightNil": true,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEqual - left nil right non-nil should return false",
		ArrangeInput: args.Map{
			"when":      "given left nil right non-nil",
			"isLeftNil": true,
			"isRightNil": false,
			"rightTotalPages": 5, "rightCurrentPageIndex": 1, "rightPerPageItems": 10, "rightTotalItems": 50,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEqual - left non-nil right nil should return false",
		ArrangeInput: args.Map{
			"when":      "given left non-nil right nil",
			"isLeftNil": false,
			"isRightNil": true,
			"leftTotalPages": 5, "leftCurrentPageIndex": 1, "leftPerPageItems": 10, "leftTotalItems": 50,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEqual - identical values should return true",
		ArrangeInput: args.Map{
			"when":      "given identical values",
			"isLeftNil": false, "isRightNil": false,
			"leftTotalPages": 3, "leftCurrentPageIndex": 2, "leftPerPageItems": 10, "leftTotalItems": 25,
			"rightTotalPages": 3, "rightCurrentPageIndex": 2, "rightPerPageItems": 10, "rightTotalItems": 25,
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEqual - different TotalPages should return false",
		ArrangeInput: args.Map{
			"when":      "given different TotalPages",
			"isLeftNil": false, "isRightNil": false,
			"leftTotalPages": 3, "leftCurrentPageIndex": 2, "leftPerPageItems": 10, "leftTotalItems": 25,
			"rightTotalPages": 5, "rightCurrentPageIndex": 2, "rightPerPageItems": 10, "rightTotalItems": 25,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEqual - different CurrentPageIndex should return false",
		ArrangeInput: args.Map{
			"when":      "given different CurrentPageIndex",
			"isLeftNil": false, "isRightNil": false,
			"leftTotalPages": 3, "leftCurrentPageIndex": 1, "leftPerPageItems": 10, "leftTotalItems": 25,
			"rightTotalPages": 3, "rightCurrentPageIndex": 2, "rightPerPageItems": 10, "rightTotalItems": 25,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEqual - different PerPageItems should return false",
		ArrangeInput: args.Map{
			"when":      "given different PerPageItems",
			"isLeftNil": false, "isRightNil": false,
			"leftTotalPages": 3, "leftCurrentPageIndex": 2, "leftPerPageItems": 10, "leftTotalItems": 25,
			"rightTotalPages": 3, "rightCurrentPageIndex": 2, "rightPerPageItems": 20, "rightTotalItems": 25,
		},
		ExpectedInput: "false",
	},
	{
		Title: "IsEqual - different TotalItems should return false",
		ArrangeInput: args.Map{
			"when":      "given different TotalItems",
			"isLeftNil": false, "isRightNil": false,
			"leftTotalPages": 3, "leftCurrentPageIndex": 2, "leftPerPageItems": 10, "leftTotalItems": 25,
			"rightTotalPages": 3, "rightCurrentPageIndex": 2, "rightPerPageItems": 10, "rightTotalItems": 30,
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// State check test cases (IsEmpty, Has*, IsInvalid*)
// =============================================================================

var pagingInfoStateTestCases = []coretestcases.CaseV1{
	// --- nil receiver ---
	{
		Title: "Nil receiver - all state checks",
		ArrangeInput: args.Map{
			"when":  "given nil PagingInfo",
			"isNil": true,
		},
		// IsEmpty, HasTotalPages, HasCurrentPageIndex, HasPerPageItems, HasTotalItems,
		// IsInvalidTotalPages, IsInvalidCurrentPageIndex, IsInvalidPerPageItems, IsInvalidTotalItems
		ExpectedInput: []string{
			"true",
			"false", "false", "false", "false",
			"true", "true", "true", "true",
		},
	},
	// --- zero values ---
	{
		Title: "Zero values - IsEmpty true, all Has false, all IsInvalid true",
		ArrangeInput: args.Map{
			"when":             "given zero-value PagingInfo",
			"isNil":            false,
			"totalPages":       0,
			"currentPageIndex": 0,
			"perPageItems":     0,
			"totalItems":       0,
		},
		ExpectedInput: []string{
			"true",
			"false", "false", "false", "false",
			"true", "true", "true", "true",
		},
	},
	// --- all positive ---
	{
		Title: "All positive - IsEmpty false, all Has true, all IsInvalid false",
		ArrangeInput: args.Map{
			"when":             "given fully populated PagingInfo",
			"isNil":            false,
			"totalPages":       5,
			"currentPageIndex": 3,
			"perPageItems":     10,
			"totalItems":       50,
		},
		ExpectedInput: []string{
			"false",
			"true", "true", "true", "true",
			"false", "false", "false", "false",
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
		ExpectedInput: []string{
			"false",
			"false", "true", "true", "true",
			"true", "false", "false", "false",
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
		ExpectedInput: []string{
			"false",
			"true", "false", "false", "false",
			"false", "true", "true", "true",
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
		ExpectedInput: []string{"5", "3", "10", "50"},
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
		ExpectedInput: []string{"0", "0", "0", "0"},
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
		ExpectedInput: "true",
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
		ExpectedInput: []string{"false", "5", "3", "10", "50"},
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
