package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// SliceValidator.IsValid
// ==========================================================================

var svIsValidExactMatchTestCase = coretestcases.CaseV1{
	Title:         "Exact match is valid",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidMismatchTestCase = coretestcases.CaseV1{
	Title:         "Content mismatch is invalid",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidLengthMismatchTestCase = coretestcases.CaseV1{
	Title:         "Length mismatch is invalid",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidBothNilTestCase = coretestcases.CaseV1{
	Title:         "Both nil is valid",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidOneNilTestCase = coretestcases.CaseV1{
	Title:         "One nil is invalid",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "Both empty is valid",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidTrimMatchTestCase = coretestcases.CaseV1{
	Title:         "Trimmed match is valid",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidContainsTestCase = coretestcases.CaseV1{
	Title:         "Contains matches substrings",
	ExpectedInput: args.Map{"isValid": true},
}

// ==========================================================================
// SliceValidator — helper methods
// ==========================================================================

var svActualLinesLengthTestCase = coretestcases.CaseV1{
	Title:         "ActualLinesLength returns correct count",
	ExpectedInput: args.Map{"length": 2},
}

var svExpectingLinesLengthTestCase = coretestcases.CaseV1{
	Title:         "ExpectingLinesLength returns correct count",
	ExpectedInput: args.Map{"length": 3},
}

var svIsUsedAlreadyFalseTestCase = coretestcases.CaseV1{
	Title:         "Fresh validator is not used already",
	ExpectedInput: args.Map{"isUsed": false},
}

var svIsUsedAlreadyTrueTestCase = coretestcases.CaseV1{
	Title:         "After ComparingValidators is used already",
	ExpectedInput: args.Map{"isUsed": true},
}

var svMethodNameTestCase = coretestcases.CaseV1{
	Title:         "MethodName returns correct name",
	ExpectedInput: args.Map{"name": "IsContains"},
}

// ==========================================================================
// SliceValidator.SetActual / SetActualVsExpected
// ==========================================================================

var svSetActualTestCase = coretestcases.CaseV1{
	Title:         "SetActual sets actual lines",
	ExpectedInput: args.Map{"length": 1},
}

var svSetActualVsExpectedTestCase = coretestcases.CaseV1{
	Title: "SetActualVsExpected sets both",
	ExpectedInput: args.Map{
		"actualLen":   1,
		"expectedLen": 1,
	},
}

// ==========================================================================
// SliceValidator.IsValidOtherLines
// ==========================================================================

var svIsValidOtherLinesMatchTestCase = coretestcases.CaseV1{
	Title:         "Matching other lines returns true",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidOtherLinesMismatchTestCase = coretestcases.CaseV1{
	Title:         "Mismatching other lines returns false",
	ExpectedInput: args.Map{"isValid": false},
}

// ==========================================================================
// SliceValidator.AllVerifyError
// ==========================================================================

var svAllVerifyErrorPassTestCase = coretestcases.CaseV1{
	Title:         "AllVerifyError passes on match",
	ExpectedInput: args.Map{"hasError": false},
}

var svAllVerifyErrorFailTestCase = coretestcases.CaseV1{
	Title:         "AllVerifyError returns error on mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

var svAllVerifyErrorSkipEmptyTestCase = coretestcases.CaseV1{
	Title:         "AllVerifyError skips when actual empty",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// SliceValidator.VerifyFirstError
// ==========================================================================

var svVerifyFirstErrorPassTestCase = coretestcases.CaseV1{
	Title:         "VerifyFirstError passes on match",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// SliceValidator.Dispose
// ==========================================================================

var svDisposeTestCase = coretestcases.CaseV1{
	Title: "Dispose nils out lines",
	ExpectedInput: args.Map{
		"actualNil":   true,
		"expectedNil": true,
	},
}

// ==========================================================================
// SliceValidator — case sensitivity
// ==========================================================================

var svCaseInsensitiveTestCase = coretestcases.CaseV1{
	Title:         "Case-insensitive match is valid",
	ExpectedInput: args.Map{"isValid": true},
}

var svCaseSensitiveFailTestCase = coretestcases.CaseV1{
	Title:         "Case-sensitive different case is invalid",
	ExpectedInput: args.Map{"isValid": false},
}

// ==========================================================================
// NewSliceValidatorUsingErr
// ==========================================================================

var svNewUsingErrNilTestCase = coretestcases.CaseV1{
	Title: "NewSliceValidatorUsingErr with nil error",
	ExpectedInput: args.Map{
		"isNotNil":  true,
		"actualLen": 0,
	},
}
