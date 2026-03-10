package coregenerictests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Hashmap Concurrency — test cases
// ==========================================================================

var hashmapSetLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "SetLock concurrent safety — all entries added",
	ExpectedInput: args.Map{"length": 500},
}

var hashmapGetLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "GetLock concurrent reads with writes",
	ExpectedInput: args.Map{"finalLength": 200},
}

var hashmapContainsLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "ContainsLock concurrent reads",
	ExpectedInput: args.Map{"finalLength": 200},
}

var hashmapRemoveLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "RemoveLock concurrent safety — all entries removed",
	ExpectedInput: args.Map{"length": 0},
}

var hashmapLengthLockConcurrencyTestCase = coretestcases.CaseV1{
	Title: "LengthLock concurrent reads during mutations",
	ExpectedInput: args.Map{
		"finalLength":   100,
		"noNegativeLen": true,
	},
}

var hashmapIsEmptyLockConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "IsEmptyLock concurrent check",
	ExpectedInput: args.Map{"length": 100},
}

var hashmapMixedOpsConcurrencyTestCase = coretestcases.CaseV1{
	Title:         "Mixed SetLock+GetLock+RemoveLock concurrent safety",
	ExpectedInput: args.Map{"finalLength": 300},
}
