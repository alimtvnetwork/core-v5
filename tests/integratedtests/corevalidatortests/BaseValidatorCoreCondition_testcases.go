package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var baseValidatorCoreConditionTestCases = []coretestcases.CaseV1{
	{
		Title: "ValidatorCoreConditionDefault returns all-false condition -- nil preset",
		ArrangeInput: args.Map{
			"presetCondition": false,
		},
		ExpectedInput: args.Map{
			"isTrimCompare":        false,
			"isUniqueWordOnly":     false,
			"isNonEmptyWhitespace": false,
			"isSortStringsBySpace": false,
		},
	},
	{
		Title: "ValidatorCoreConditionDefault returns existing condition -- non-nil preset",
		ArrangeInput: args.Map{
			"presetCondition": true,
		},
		ExpectedInput: args.Map{
			"isTrimCompare":        true,
			"isUniqueWordOnly":     false,
			"isNonEmptyWhitespace": true,
			"isSortStringsBySpace": false,
		},
	},
}
