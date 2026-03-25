package coreuniquetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// GetMap
// ==========================================================================

var extGetMapWithDuplicatesTestCase = coretestcases.CaseV1{
	Title: "GetMap returns unique map -- slice with duplicates [1,2,2,3]",
	ArrangeInput: args.Map{
		"when":  "given slice with duplicates",
		"input": []int{1, 2, 2, 3},
	},
	ExpectedInput: args.Map{
		"isNil":  false,
		"length": 3,
	},
}

var extGetMapNilTestCase = coretestcases.CaseV1{
	Title: "GetMap returns nil -- nil slice input",
	ArrangeInput: args.Map{
		"when": "given nil slice",
	},
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var extGetMapEmptyTestCase = coretestcases.CaseV1{
	Title: "GetMap returns empty map -- empty slice input",
	ArrangeInput: args.Map{
		"when":  "given empty slice",
		"input": []int{},
	},
	ExpectedInput: args.Map{
		"isNil":  false,
		"length": 0,
	},
}

// ==========================================================================
// Get with empty slice
// ==========================================================================

var extGetEmptySliceTestCase = coretestcases.CaseV1{
	Title: "Get returns same slice -- empty slice input",
	ArrangeInput: args.Map{
		"when":  "given empty slice",
		"input": []int{},
	},
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// ==========================================================================
// Get with single element
// ==========================================================================

var extGetSingleElementTestCase = coretestcases.CaseV1{
	Title: "Get returns same slice -- single element slice",
	ArrangeInput: args.Map{
		"when":  "given single element slice",
		"input": []int{42},
	},
	ExpectedInput: args.Map{
		"length": 1,
	},
}
