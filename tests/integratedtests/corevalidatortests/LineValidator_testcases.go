package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// LineValidator.IsMatch
// ==========================================================================

var lineValidatorIsMatchBothTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true when line and text match",
	ExpectedInput: args.Map{"isMatch": true},
}

var lineValidatorIsMatchLineMismatchTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false on line number mismatch",
	ExpectedInput: args.Map{"isMatch": false},
}

var lineValidatorIsMatchTextMismatchTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false on text mismatch",
	ExpectedInput: args.Map{"isMatch": false},
}

var lineValidatorIsMatchSkipLineTestCase = coretestcases.CaseV1{
	Title:         "IsMatch with skip line number passes on text match",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// LineValidator.IsMatchMany
// ==========================================================================

var lineValidatorIsMatchManyAllTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany returns true when all match",
	ExpectedInput: args.Map{"isMatch": true},
}

var lineValidatorIsMatchManyOneFailsTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany returns false when one fails",
	ExpectedInput: args.Map{"isMatch": false},
}

var lineValidatorIsMatchManyEmptySkipTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany empty with skip returns true",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// LineValidator.VerifyError
// ==========================================================================

var lineValidatorVerifyErrorMatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyError returns nil on match",
	ExpectedInput: args.Map{"hasError": false},
}

var lineValidatorVerifyErrorLineMismatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyError returns error on line mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

var lineValidatorVerifyErrorTextMismatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyError returns error on text mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

// ==========================================================================
// LineValidator.VerifyMany
// ==========================================================================

var lineValidatorVerifyManyContinueTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany continueOnError collects errors",
	ExpectedInput: args.Map{"hasError": true},
}

var lineValidatorVerifyManyFirstOnlyTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany firstOnly returns first error",
	ExpectedInput: args.Map{"hasError": true},
}
