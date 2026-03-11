package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// SliceValidator.IsValid
// ==========================================================================

var svIsValidExactMatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- exact match",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidMismatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- content mismatch",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidLengthMismatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- length mismatch",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidBothNilTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- both nil",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidOneNilTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- one nil",
	ExpectedInput: args.Map{"isValid": false},
}

var svIsValidBothEmptyTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- both empty",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidTrimMatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- trimmed match",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidContainsTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- contains substrings",
	ExpectedInput: args.Map{"isValid": true},
}

// ==========================================================================
// SliceValidator — helper methods
// ==========================================================================

var svActualLinesLengthTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.ActualLinesLength returns 2 -- two actual lines",
	ExpectedInput: args.Map{"length": 2},
}

var svExpectingLinesLengthTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.ExpectingLinesLength returns 3 -- three expected lines",
	ExpectedInput: args.Map{"length": 3},
}

var svIsUsedAlreadyFalseTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsUsedAlready returns false -- fresh instance",
	ExpectedInput: args.Map{"isUsed": false},
}

var svIsUsedAlreadyTrueTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsUsedAlready returns true -- after ComparingValidators",
	ExpectedInput: args.Map{"isUsed": true},
}

var svMethodNameTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.MethodName returns 'IsContains' -- Contains compare mode",
	ExpectedInput: args.Map{"name": "IsContains"},
}

// ==========================================================================
// SliceValidator.SetActual / SetActualVsExpected
// ==========================================================================

var svSetActualTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.SetActual returns length 1 -- one line set",
	ExpectedInput: args.Map{"length": 1},
}

var svSetActualVsExpectedTestCase = coretestcases.CaseV1{
	Title: "SliceValidator.SetActualVsExpected returns both set -- one actual one expected",
	ExpectedInput: args.Map{
		"actualLen":   1,
		"expectedLen": 1,
	},
}

// ==========================================================================
// SliceValidator.IsValidOtherLines
// ==========================================================================

var svIsValidOtherLinesMatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValidOtherLines returns true -- matching lines",
	ExpectedInput: args.Map{"isValid": true},
}

var svIsValidOtherLinesMismatchTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValidOtherLines returns false -- mismatching lines",
	ExpectedInput: args.Map{"isValid": false},
}

// ==========================================================================
// SliceValidator.AllVerifyError
// ==========================================================================

var svAllVerifyErrorPassTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.AllVerifyError returns nil -- matching lines",
	ExpectedInput: args.Map{"hasError": false},
}

var svAllVerifyErrorFailTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.AllVerifyError returns error -- mismatched lines",
	ExpectedInput: args.Map{"hasError": true},
}

var svAllVerifyErrorSkipEmptyTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.AllVerifyError returns nil -- skip when actual empty",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// SliceValidator.VerifyFirstError
// ==========================================================================

var svVerifyFirstErrorPassTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.VerifyFirstError returns nil -- matching lines",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// SliceValidator.Dispose
// ==========================================================================

var svDisposeTestCase = coretestcases.CaseV1{
	Title: "SliceValidator.Dispose returns nil lines -- after dispose",
	ExpectedInput: args.Map{
		"actualNil":   true,
		"expectedNil": true,
	},
}

// ==========================================================================
// SliceValidator — case sensitivity
// ==========================================================================

var svCaseInsensitiveTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns true -- case-insensitive match",
	ExpectedInput: args.Map{"isValid": true},
}

var svCaseSensitiveFailTestCase = coretestcases.CaseV1{
	Title:         "SliceValidator.IsValid returns false -- case-sensitive different case",
	ExpectedInput: args.Map{"isValid": false},
}

// ==========================================================================
// NewSliceValidatorUsingErr
// ==========================================================================

var svNewUsingErrNilTestCase = coretestcases.CaseV1{
	Title: "NewSliceValidatorUsingErr returns non-nil with 0 actual -- nil error input",
	ExpectedInput: args.Map{
		"isNotNil":  true,
		"actualLen": 0,
	},
}
