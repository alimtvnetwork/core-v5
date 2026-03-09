package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// TextValidator.IsMatch — Equal
// ==========================================================================

var tvIsMatchExactEqualTestCase = coretestcases.CaseV1{
	Title:         "Exact equal match returns true",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchExactNotEqualTestCase = coretestcases.CaseV1{
	Title:         "Different text returns false",
	ExpectedInput: args.Map{"isMatch": false},
}

var tvIsMatchCaseInsensitiveTestCase = coretestcases.CaseV1{
	Title:         "Case-insensitive match returns true",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchCaseSensitiveFailTestCase = coretestcases.CaseV1{
	Title:         "Case-sensitive mismatch returns false",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatch — Trim
// ==========================================================================

var tvIsMatchTrimTestCase = coretestcases.CaseV1{
	Title:         "Trimmed search matches trimmed content",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchTrimBothTestCase = coretestcases.CaseV1{
	Title:         "Trim handles both search and content",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidator.IsMatch — Contains
// ==========================================================================

var tvIsMatchContainsTestCase = coretestcases.CaseV1{
	Title:         "Contains finds substring",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchContainsMissingTestCase = coretestcases.CaseV1{
	Title:         "Contains does not find missing substring",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatch — NotEqual
// ==========================================================================

var tvIsMatchNotEqualDifferentTestCase = coretestcases.CaseV1{
	Title:         "NotEqual matches when different",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchNotEqualSameTestCase = coretestcases.CaseV1{
	Title:         "NotEqual does not match when same",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatch — UniqueWords + Sort
// ==========================================================================

var tvIsMatchUniqueWordsSortedTestCase = coretestcases.CaseV1{
	Title:         "Unique+sorted matches reordered unique words",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidator.IsMatch — Empty strings
// ==========================================================================

var tvIsMatchEmptyBothTestCase = coretestcases.CaseV1{
	Title:         "Empty search vs empty content matches",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchEmptySearchNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "Empty search vs non-empty content does not match",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatchMany
// ==========================================================================

var tvIsMatchManyAllTestCase = coretestcases.CaseV1{
	Title:         "All identical match",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchManyOneFailsTestCase = coretestcases.CaseV1{
	Title:         "One mismatch fails",
	ExpectedInput: args.Map{"isMatch": false},
}

var tvIsMatchManyEmptySkipTestCase = coretestcases.CaseV1{
	Title:         "Empty contents with skip returns true",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidator.VerifyDetailError
// ==========================================================================

var tvVerifyDetailMatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyDetailError returns nil on match",
	ExpectedInput: args.Map{"hasError": false},
}

var tvVerifyDetailMismatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyDetailError returns error on mismatch",
	ExpectedInput: args.Map{"hasError": true},
}

// ==========================================================================
// TextValidator.VerifyMany
// ==========================================================================

var tvVerifyManyFirstOnlyTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany firstOnly returns first error",
	ExpectedInput: args.Map{"hasError": true},
}

var tvVerifyManyAllErrorsTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany collects all errors",
	ExpectedInput: args.Map{"hasError": true},
}

var tvVerifyFirstEmptySkipTestCase = coretestcases.CaseV1{
	Title:         "VerifyFirstError empty with skip returns nil",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// TextValidator.SearchTextFinalized — caching
// ==========================================================================

var tvSearchTextFinalizedTestCase = coretestcases.CaseV1{
	Title: "SearchTextFinalized caches and trims",
	ExpectedInput: args.Map{
		"isCached": true,
		"value":    "hello",
	},
}

// ==========================================================================
// EmptyValidator preset
// ==========================================================================

var tvEmptyMatchesEmptyTestCase = coretestcases.CaseV1{
	Title:         "EmptyValidator matches empty string",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvEmptyMatchesTrimmedTestCase = coretestcases.CaseV1{
	Title:         "EmptyValidator matches whitespace-only",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvEmptyNoMatchNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "EmptyValidator does not match non-empty",
	ExpectedInput: args.Map{"isMatch": false},
}
