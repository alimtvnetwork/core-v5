package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Hashset Concurrency — test cases
// ==========================================================================

var hashsetAddLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "AddLock concurrent safety -- all items added",
	ExpectedInput: args.Map{"length": 500},
}

var hashsetAddSliceLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "AddSliceLock concurrent safety — all batches added",
	ExpectedInput: args.Map{"length": 1000},
}

var hashsetContainsLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "ContainsLock concurrent reads during writes",
	ExpectedInput: args.Map{"finalLength": 200},
}

var hashsetRemoveLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "RemoveLock concurrent safety — all items removed",
	ExpectedInput: args.Map{"length": 0},
}

var hashsetLengthLockConcurrencyTestCase = coretestcases.CaseV1{
	Title: "LengthLock concurrent reads during mutations",
	ExpectedInput: args.Map{
		"finalLength":   100,
		"noNegativeLen": true,
	},
}

var hashsetIsEmptyLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "IsEmptyLock concurrent check with writes",
	ExpectedInput: args.Map{"length": 100},
}
