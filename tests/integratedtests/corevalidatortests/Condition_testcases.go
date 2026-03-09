package corevalidatortests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================================================
// Condition.IsSplitByWhitespace
// ==========================================================================

var conditionAllFalseTestCase = coretestcases.CaseV1{
	Title:         "All false should not split by whitespace",
	ExpectedInput: args.Map{"isSplit": false},
}

var conditionUniqueWordOnlyTestCase = coretestcases.CaseV1{
	Title:         "IsUniqueWordOnly triggers split",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionNonEmptyWhitespaceTestCase = coretestcases.CaseV1{
	Title:         "IsNonEmptyWhitespace triggers split",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionSortBySpaceTestCase = coretestcases.CaseV1{
	Title:         "IsSortStringsBySpace triggers split",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionTrimOnlyTestCase = coretestcases.CaseV1{
	Title:         "IsTrimCompare alone does not trigger split",
	ExpectedInput: args.Map{"isSplit": false},
}

// ==========================================================================
// Preset Conditions
// ==========================================================================

var conditionDisabledTestCase = coretestcases.CaseV1{
	Title:         "DefaultDisabled does not split",
	ExpectedInput: args.Map{"isSplit": false},
}

var conditionTrimTestCase = coretestcases.CaseV1{
	Title: "DefaultTrim does not split but has IsTrimCompare",
	ExpectedInput: args.Map{
		"isSplit":       false,
		"isTrimCompare": true,
	},
}

var conditionSortTrimTestCase = coretestcases.CaseV1{
	Title:         "DefaultSortTrim splits",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionUniqueWordsTestCase = coretestcases.CaseV1{
	Title: "DefaultUniqueWords splits and has IsUniqueWordOnly",
	ExpectedInput: args.Map{
		"isSplit":          true,
		"isUniqueWordOnly": true,
	},
}
