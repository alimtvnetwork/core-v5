package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/corevalidator"
)

func Test_BaseValidatorCoreCondition(t *testing.T) {
	for caseIndex, tc := range baseValidatorCoreConditionTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		presetCondition, _ := input.GetAsBool("presetCondition")

		base := &corevalidator.BaseValidatorCoreCondition{}
		if presetCondition {
			base.ValidatorCoreCondition = &corevalidator.Condition{
				IsTrimCompare:        true,
				IsNonEmptyWhitespace: true,
			}
		}

		// Act
		condition := base.ValidatorCoreConditionDefault()

		actual := args.Map{
			"isTrimCompare":        condition.IsTrimCompare,
			"isUniqueWordOnly":     condition.IsUniqueWordOnly,
			"isNonEmptyWhitespace": condition.IsNonEmptyWhitespace,
			"isSortStringsBySpace": condition.IsSortStringsBySpace,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
