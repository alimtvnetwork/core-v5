package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
)

// ==========================================
// Condition.IsSplitByWhitespace
// ==========================================

func Test_Condition_IsSplitByWhitespace_AllFalse(t *testing.T) {
	tc := conditionAllFalseTestCase
	c := corevalidator.Condition{}

	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_UniqueWordOnly(t *testing.T) {
	tc := conditionUniqueWordOnlyTestCase
	c := corevalidator.Condition{IsUniqueWordOnly: true}

	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_NonEmptyWhitespace(t *testing.T) {
	tc := conditionNonEmptyWhitespaceTestCase
	c := corevalidator.Condition{IsNonEmptyWhitespace: true}

	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_SortBySpace(t *testing.T) {
	tc := conditionSortBySpaceTestCase
	c := corevalidator.Condition{IsSortStringsBySpace: true}

	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_TrimOnlyNotEnough(t *testing.T) {
	tc := conditionTrimOnlyTestCase
	c := corevalidator.Condition{IsTrimCompare: true}

	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Preset Conditions
// ==========================================

func Test_DefaultDisabledCondition_NoSplit(t *testing.T) {
	tc := conditionDisabledTestCase
	c := corevalidator.DefaultDisabledCoreCondition

	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultTrimCondition_NoSplit(t *testing.T) {
	tc := conditionTrimTestCase
	c := corevalidator.DefaultTrimCoreCondition

	actual := args.Map{
		"isSplit":       c.IsSplitByWhitespace(),
		"isTrimCompare": c.IsTrimCompare,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultSortTrimCondition_Split(t *testing.T) {
	tc := conditionSortTrimTestCase
	c := corevalidator.DefaultSortTrimCoreCondition

	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultUniqueWordsCondition_Split(t *testing.T) {
	tc := conditionUniqueWordsTestCase
	c := corevalidator.DefaultUniqueWordsCoreCondition

	actual := args.Map{
		"isSplit":          c.IsSplitByWhitespace(),
		"isUniqueWordOnly": c.IsUniqueWordOnly,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}
