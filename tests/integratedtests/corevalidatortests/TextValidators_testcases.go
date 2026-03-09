package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// TextValidators — collection operations
// ==========================================================================

var textValidatorsNewEmptyTestCase = coretestcases.CaseV1{
	Title: "New TextValidators is empty with length 0",
	ExpectedInput: args.Map{
		"isEmpty": true,
		"length":  0,
	},
}

var textValidatorsAddTestCase = coretestcases.CaseV1{
	Title:         "Add increases length",
	ExpectedInput: args.Map{"length": 2},
}

var textValidatorsAddsTestCase = coretestcases.CaseV1{
	Title:         "Adds variadic adds all",
	ExpectedInput: args.Map{"length": 2},
}

var textValidatorsAddsEmptyTestCase = coretestcases.CaseV1{
	Title:         "Adds with nothing stays empty",
	ExpectedInput: args.Map{"length": 0},
}

var textValidatorsAddSimpleTestCase = coretestcases.CaseV1{
	Title:         "AddSimple adds one",
	ExpectedInput: args.Map{"length": 1},
}

var textValidatorsHasIndexTestCase = coretestcases.CaseV1{
	Title: "HasIndex returns true for valid, false for invalid",
	ExpectedInput: args.Map{
		"hasIndex0": true,
		"hasIndex1": false,
	},
}

var textValidatorsLastIndexTestCase = coretestcases.CaseV1{
	Title:         "LastIndex returns correct value",
	ExpectedInput: args.Map{"lastIndex": 1},
}

// ==========================================================================
// TextValidators.IsMatch
// ==========================================================================

var textValidatorsIsMatchEmptyTestCase = coretestcases.CaseV1{
	Title:         "Empty validators match anything",
	ExpectedInput: args.Map{"isMatch": true},
}

var textValidatorsIsMatchAllPassTestCase = coretestcases.CaseV1{
	Title:         "All validators pass → match",
	ExpectedInput: args.Map{"isMatch": true},
}

var textValidatorsIsMatchOneFailsTestCase = coretestcases.CaseV1{
	Title:         "One validator fails → no match",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidators.IsMatchMany
// ==========================================================================

var textValidatorsIsMatchManyEmptyTestCase = coretestcases.CaseV1{
	Title:         "Empty validators match many",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidators.VerifyFirstError
// ==========================================================================

var textValidatorsVerifyFirstAllPassTestCase = coretestcases.CaseV1{
	Title:         "VerifyFirstError passes on match",
	ExpectedInput: args.Map{"hasError": false},
}

var textValidatorsVerifyFirstFailTestCase = coretestcases.CaseV1{
	Title:         "VerifyFirstError returns error on mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

var textValidatorsVerifyFirstEmptyTestCase = coretestcases.CaseV1{
	Title:         "VerifyFirstError empty validators returns nil",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// TextValidators.AllVerifyError
// ==========================================================================

var textValidatorsAllVerifyPassTestCase = coretestcases.CaseV1{
	Title:         "AllVerifyError passes on match",
	ExpectedInput: args.Map{"hasError": false},
}

var textValidatorsAllVerifyFailTestCase = coretestcases.CaseV1{
	Title:         "AllVerifyError returns error on multiple mismatches",
	ExpectedInput: args.Map{"hasError": true},
}

// ==========================================================================
// TextValidators.Dispose
// ==========================================================================

var textValidatorsDisposeTestCase = coretestcases.CaseV1{
	Title:         "Dispose nils out Items",
	ExpectedInput: args.Map{"isNil": true},
}
