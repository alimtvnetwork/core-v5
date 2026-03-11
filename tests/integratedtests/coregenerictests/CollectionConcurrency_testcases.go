package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Collection Concurrency — test cases
// ==========================================================================

var collectionAddLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "AddLock concurrent safety -- all items added",
	ExpectedInput: args.Map{"length": 500},
}

var collectionAddsLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "AddsLock concurrent safety — all batches added",
	ExpectedInput: args.Map{"length": 1000},
}

var collectionLengthLockConcurrencyTestCase = coretestcases.CaseV1{
	Title: "LengthLock concurrent reads during writes",
	ExpectedInput: args.Map{
		"finalLength":   100,
		"noNegativeLen": true,
	},
}

var collectionIsEmptyLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "IsEmptyLock concurrent check with writes",
	ExpectedInput: args.Map{"length": 100},
}
