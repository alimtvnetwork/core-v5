package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
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
