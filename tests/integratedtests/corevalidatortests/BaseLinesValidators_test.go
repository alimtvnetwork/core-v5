package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// BaseLinesValidators
// ==========================================

func Test_BaseLinesValidators_Empty(t *testing.T) {
	b := corevalidator.BaseLinesValidators{}
	if !b.IsEmptyLinesValidators() {
		t.Error("empty should be empty")
	}
	if b.HasLinesValidators() {
		t.Error("empty should not have validators")
	}
	if b.LinesValidatorsLength() != 0 {
		t.Errorf("expected 0, got %d", b.LinesValidatorsLength())
	}
}

func Test_BaseLinesValidators_WithItems(t *testing.T) {
	b := corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{
				LineNumber: corevalidator.LineNumber{LineNumber: -1},
				TextValidator: corevalidator.TextValidator{
					Search:    "a",
					SearchAs:  stringcompareas.Equal,
					Condition: corevalidator.DefaultDisabledCoreCondition,
				},
			},
		},
	}
	if b.IsEmptyLinesValidators() {
		t.Error("should not be empty")
	}
	if !b.HasLinesValidators() {
		t.Error("should have validators")
	}
	if b.LinesValidatorsLength() != 1 {
		t.Errorf("expected 1, got %d", b.LinesValidatorsLength())
	}
}

func Test_BaseLinesValidators_ToLinesValidators_Empty(t *testing.T) {
	b := corevalidator.BaseLinesValidators{}
	lv := b.ToLinesValidators()
	if lv == nil {
		t.Error("should not be nil")
	}
	if !lv.IsEmpty() {
		t.Error("should be empty")
	}
}

func Test_BaseLinesValidators_ToLinesValidators_NonEmpty(t *testing.T) {
	b := corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{
				LineNumber: corevalidator.LineNumber{LineNumber: 0},
				TextValidator: corevalidator.TextValidator{
					Search:    "test",
					SearchAs:  stringcompareas.Equal,
					Condition: corevalidator.DefaultDisabledCoreCondition,
				},
			},
			{
				LineNumber: corevalidator.LineNumber{LineNumber: 1},
				TextValidator: corevalidator.TextValidator{
					Search:    "test2",
					SearchAs:  stringcompareas.Equal,
					Condition: corevalidator.DefaultDisabledCoreCondition,
				},
			},
		},
	}
	lv := b.ToLinesValidators()
	if lv.Length() != 2 {
		t.Errorf("expected 2, got %d", lv.Length())
	}
}

// (nil receiver tests migrated to BaseLinesValidators_NilReceiver_testcases.go)

// ==========================================
// LinesValidators — collection
// ==========================================

func Test_LinesValidators_New(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	if lv == nil {
		t.Error("should not be nil")
	}
	if !lv.IsEmpty() {
		t.Error("new should be empty")
	}
}

func Test_LinesValidators_Add(t *testing.T) {
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	if lv.Length() != 1 {
		t.Errorf("expected 1, got %d", lv.Length())
	}
	if !lv.HasAnyItem() {
		t.Error("should have items")
	}
}

func Test_LinesValidators_AddPtr_Nil(t *testing.T) {
	lv := corevalidator.NewLinesValidators(2)
	lv.AddPtr(nil)
	if lv.Length() != 0 {
		t.Error("nil add should not increase length")
	}
}

func Test_LinesValidators_HasIndex(t *testing.T) {
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})
	if !lv.HasIndex(0) {
		t.Error("should have index 0")
	}
	if lv.HasIndex(1) {
		t.Error("should not have index 1")
	}
}

// (nil receiver tests migrated to BaseLinesValidators_NilReceiver_testcases.go)

// ==========================================
// LinesValidators.IsMatchText
// ==========================================

func Test_LinesValidators_IsMatchText_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	if !lv.IsMatchText("anything", true) {
		t.Error("empty validators should match any text")
	}
}

func Test_LinesValidators_IsMatchText_Match(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	if !lv.IsMatchText("hello world", true) {
		t.Error("contains should match")
	}
}

func Test_LinesValidators_IsMatchText_NoMatch(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "xyz",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	if lv.IsMatchText("hello world", true) {
		t.Error("missing substring should not match")
	}
}

// ==========================================
// BaseValidatorCoreCondition
// ==========================================

func Test_BaseValidatorCoreCondition_Default_NilPtr(t *testing.T) {
	b := corevalidator.BaseValidatorCoreCondition{}
	c := b.ValidatorCoreConditionDefault()
	if c.IsTrimCompare || c.IsUniqueWordOnly {
		t.Error("default condition should have all false")
	}
	// should set the ptr
	if b.ValidatorCoreCondition == nil {
		t.Error("should have set the pointer")
	}
}

func Test_BaseValidatorCoreCondition_Default_ExistingPtr(t *testing.T) {
	cond := corevalidator.Condition{IsTrimCompare: true}
	b := corevalidator.BaseValidatorCoreCondition{
		ValidatorCoreCondition: &cond,
	}
	c := b.ValidatorCoreConditionDefault()
	if !c.IsTrimCompare {
		t.Error("should return existing condition")
	}
}
