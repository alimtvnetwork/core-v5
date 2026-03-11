package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// Condition.IsSplitByWhitespace
// ==========================================================================

var conditionAllFalseTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns false -- all flags false",
	ExpectedInput: args.Map{"isSplit": false},
}

var conditionUniqueWordOnlyTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns true -- IsUniqueWordOnly enabled",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionNonEmptyWhitespaceTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns true -- IsNonEmptyWhitespace enabled",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionSortBySpaceTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns true -- IsSortStringsBySpace enabled",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionTrimOnlyTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns false -- IsTrimCompare only",
	ExpectedInput: args.Map{"isSplit": false},
}

// ==========================================================================
// Preset Conditions
// ==========================================================================

var conditionDisabledTestCase = coretestcases.CaseV1{
	Title:         "DefaultDisabled returns isSplit false -- preset disabled",
	ExpectedInput: args.Map{"isSplit": false},
}

var conditionTrimTestCase = coretestcases.CaseV1{
	Title: "DefaultTrim returns isSplit false, isTrimCompare true -- preset trim",
	ExpectedInput: args.Map{
		"isSplit":       false,
		"isTrimCompare": true,
	},
}

var conditionSortTrimTestCase = coretestcases.CaseV1{
	Title:         "DefaultSortTrim returns isSplit true -- preset sort-trim",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionUniqueWordsTestCase = coretestcases.CaseV1{
	Title: "DefaultUniqueWords returns isSplit true, isUniqueWordOnly true -- preset unique-words",
	ExpectedInput: args.Map{
		"isSplit":          true,
		"isUniqueWordOnly": true,
	},
}
