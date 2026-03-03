package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
)

// ==========================================
// Condition.IsSplitByWhitespace
// ==========================================

func Test_Condition_IsSplitByWhitespace_AllFalse(t *testing.T) {
	c := corevalidator.Condition{}
	if c.IsSplitByWhitespace() {
		t.Error("all false should not split by whitespace")
	}
}

func Test_Condition_IsSplitByWhitespace_UniqueWordOnly(t *testing.T) {
	c := corevalidator.Condition{IsUniqueWordOnly: true}
	if !c.IsSplitByWhitespace() {
		t.Error("IsUniqueWordOnly true should trigger split")
	}
}

func Test_Condition_IsSplitByWhitespace_NonEmptyWhitespace(t *testing.T) {
	c := corevalidator.Condition{IsNonEmptyWhitespace: true}
	if !c.IsSplitByWhitespace() {
		t.Error("IsNonEmptyWhitespace true should trigger split")
	}
}

func Test_Condition_IsSplitByWhitespace_SortBySpace(t *testing.T) {
	c := corevalidator.Condition{IsSortStringsBySpace: true}
	if !c.IsSplitByWhitespace() {
		t.Error("IsSortStringsBySpace true should trigger split")
	}
}

func Test_Condition_IsSplitByWhitespace_TrimOnlyNotEnough(t *testing.T) {
	c := corevalidator.Condition{IsTrimCompare: true}
	if c.IsSplitByWhitespace() {
		t.Error("only IsTrimCompare should not trigger split")
	}
}

// ==========================================
// Preset Conditions
// ==========================================

func Test_DefaultDisabledCondition_NoSplit(t *testing.T) {
	c := corevalidator.DefaultDisabledCoreCondition
	if c.IsSplitByWhitespace() {
		t.Error("disabled condition should not split")
	}
}

func Test_DefaultTrimCondition_NoSplit(t *testing.T) {
	c := corevalidator.DefaultTrimCoreCondition
	if c.IsSplitByWhitespace() {
		t.Error("trim-only condition should not split")
	}
	if !c.IsTrimCompare {
		t.Error("trim condition should have IsTrimCompare true")
	}
}

func Test_DefaultSortTrimCondition_Split(t *testing.T) {
	c := corevalidator.DefaultSortTrimCoreCondition
	if !c.IsSplitByWhitespace() {
		t.Error("sort+trim condition should split")
	}
}

func Test_DefaultUniqueWordsCondition_Split(t *testing.T) {
	c := corevalidator.DefaultUniqueWordsCoreCondition
	if !c.IsSplitByWhitespace() {
		t.Error("unique words condition should split")
	}
	if !c.IsUniqueWordOnly {
		t.Error("should have IsUniqueWordOnly true")
	}
}
