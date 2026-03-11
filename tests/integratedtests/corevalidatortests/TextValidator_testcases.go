package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// TextValidator.IsMatch — Equal
// ==========================================================================

var tvIsMatchExactEqualTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- exact equal text",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchExactNotEqualTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false -- different text",
	ExpectedInput: args.Map{"isMatch": false},
}

var tvIsMatchCaseInsensitiveTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- case-insensitive match",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchCaseSensitiveFailTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false -- case-sensitive mismatch",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatch — Trim
// ==========================================================================

var tvIsMatchTrimTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- trimmed search matches content",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchTrimBothTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- trim applied to both search and content",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidator.IsMatch — Contains
// ==========================================================================

var tvIsMatchContainsTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- contains substring found",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchContainsMissingTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false -- contains substring not found",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatch — NotEqual
// ==========================================================================

var tvIsMatchNotEqualDifferentTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- NotEqual with different text",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchNotEqualSameTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false -- NotEqual with same text",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatch — UniqueWords + Sort
// ==========================================================================

var tvIsMatchUniqueWordsSortedTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- unique+sorted reordered words",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidator.IsMatch — Empty strings
// ==========================================================================

var tvIsMatchEmptyBothTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns true -- both search and content empty",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchEmptySearchNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsMatch returns false -- empty search vs non-empty content",
	ExpectedInput: args.Map{"isMatch": false},
}

// ==========================================================================
// TextValidator.IsMatchMany
// ==========================================================================

var tvIsMatchManyAllTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany returns true -- all lines identical",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvIsMatchManyOneFailsTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany returns false -- one line mismatched",
	ExpectedInput: args.Map{"isMatch": false},
}

var tvIsMatchManyEmptySkipTestCase = coretestcases.CaseV1{
	Title:         "IsMatchMany returns true -- empty contents with skip",
	ExpectedInput: args.Map{"isMatch": true},
}

// ==========================================================================
// TextValidator.VerifyDetailError
// ==========================================================================

var tvVerifyDetailMatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyDetailError returns nil -- matching text",
	ExpectedInput: args.Map{"hasError": false},
}

var tvVerifyDetailMismatchTestCase = coretestcases.CaseV1{
	Title:         "VerifyDetailError returns error -- mismatched text",
	ExpectedInput: args.Map{"hasError": true},
}

// ==========================================================================
// TextValidator.VerifyMany
// ==========================================================================

var tvVerifyManyFirstOnlyTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany returns first error -- firstOnly mode",
	ExpectedInput: args.Map{"hasError": true},
}

var tvVerifyManyAllErrorsTestCase = coretestcases.CaseV1{
	Title:         "VerifyMany returns all errors -- collect mode",
	ExpectedInput: args.Map{"hasError": true},
}

var tvVerifyFirstEmptySkipTestCase = coretestcases.CaseV1{
	Title:         "VerifyFirstError returns nil -- empty contents with skip",
	ExpectedInput: args.Map{"hasError": false},
}

// ==========================================================================
// TextValidator.SearchTextFinalized — caching
// ==========================================================================

var tvSearchTextFinalizedTestCase = coretestcases.CaseV1{
	Title: "SearchTextFinalized returns cached trimmed value -- 'hello' with whitespace",
	ExpectedInput: args.Map{
		"isCached": true,
		"value":    "hello",
	},
}

// ==========================================================================
// EmptyValidator preset
// ==========================================================================

var tvEmptyMatchesEmptyTestCase = coretestcases.CaseV1{
	Title:         "EmptyValidator returns match true -- empty string input",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvEmptyMatchesTrimmedTestCase = coretestcases.CaseV1{
	Title:         "EmptyValidator returns match true -- whitespace-only input",
	ExpectedInput: args.Map{"isMatch": true},
}

var tvEmptyNoMatchNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "EmptyValidator returns match false -- non-empty input",
	ExpectedInput: args.Map{"isMatch": false},
}
