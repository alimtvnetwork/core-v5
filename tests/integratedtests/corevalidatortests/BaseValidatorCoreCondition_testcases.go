package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var baseValidatorCoreConditionTestCases = []coretestcases.CaseV1{
	{
		Title: "ValidatorCoreConditionDefault nil creates default condition",
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
		Title: "ValidatorCoreConditionDefault non-nil returns existing condition",
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
