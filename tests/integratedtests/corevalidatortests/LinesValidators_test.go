package corevalidatortests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// LinesValidators — collection basics
// ==========================================

func Test_LinesValidators_Count(t *testing.T) {
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})
	if lv.Count() != 1 {
		t.Errorf("expected 1, got %d", lv.Count())
	}
}

func Test_LinesValidators_LastIndex(t *testing.T) {
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})
	lv.Add(corevalidator.LineValidator{})
	if lv.LastIndex() != 1 {
		t.Errorf("expected 1, got %d", lv.LastIndex())
	}
}

func Test_LinesValidators_Adds(t *testing.T) {
	lv := corevalidator.NewLinesValidators(3)
	lv.Adds(
		corevalidator.LineValidator{},
		corevalidator.LineValidator{},
		corevalidator.LineValidator{},
	)
	if lv.Length() != 3 {
		t.Errorf("expected 3, got %d", lv.Length())
	}
}

func Test_LinesValidators_String(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:   "test",
			SearchAs: stringcompareas.Equal,
		},
	})
	s := lv.String()
	if s == "" {
		t.Error("String should not be empty")
	}
}

// ==========================================
// LinesValidators.IsMatch (with contents)
// ==========================================

func Test_LinesValidators_IsMatch_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	items := []corestr.TextWithLineNumber{
		{Text: "hello", LineNumber: 0},
	}
	if !lv.IsMatch(false, true, items...) {
		t.Error("empty validators should match")
	}
}

func Test_LinesValidators_IsMatch_NoContentsSkip(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	if !lv.IsMatch(true, true) {
		t.Error("no contents with skip should match")
	}
}

func Test_LinesValidators_IsMatch_NoContentsNoSkip(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	if lv.IsMatch(false, true) {
		t.Error("no contents without skip should not match")
	}
}

func Test_LinesValidators_IsMatch_AllMatch(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
	}
	if !lv.IsMatch(false, true, items...) {
		t.Error("all matching should return true")
	}
}

func Test_LinesValidators_IsMatch_OneFails(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "nope", LineNumber: 1},
	}
	if lv.IsMatch(false, true, items...) {
		t.Error("one failing should return false")
	}
}

// ==========================================
// LinesValidators.VerifyFirstDefaultLineNumberError
// ==========================================

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := lv.VerifyFirstDefaultLineNumberError(params)
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_SkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := lv.VerifyFirstDefaultLineNumberError(params)
	if err != nil {
		t.Errorf("skip empty should return nil: %v", err)
	}
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_NoSkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := lv.VerifyFirstDefaultLineNumberError(params)
	if err == nil {
		t.Error("empty contents without skip should return error")
	}
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Pass(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
	}
	err := lv.VerifyFirstDefaultLineNumberError(params, items...)
	if err != nil {
		t.Errorf("match should pass: %v", err)
	}
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Fail(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
	}
	err := lv.VerifyFirstDefaultLineNumberError(params, items...)
	if err == nil {
		t.Error("mismatch should return error")
	}
}

// ==========================================
// LinesValidators.AllVerifyError
// ==========================================

func Test_LinesValidators_AllVerifyError_Empty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := lv.AllVerifyError(params)
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_LinesValidators_AllVerifyError_SkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := lv.AllVerifyError(params)
	if err != nil {
		t.Errorf("skip empty should return nil: %v", err)
	}
}

func Test_LinesValidators_AllVerifyError_NoSkipEmpty(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := lv.AllVerifyError(params)
	if err == nil {
		t.Error("empty contents without skip should return error")
	}
}

func Test_LinesValidators_AllVerifyError_Pass(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
	}
	err := lv.AllVerifyError(params, items...)
	if err != nil {
		t.Errorf("match should pass: %v", err)
	}
}

func Test_LinesValidators_AllVerifyError_Fail(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
	}
	err := lv.AllVerifyError(params, items...)
	if err == nil {
		t.Error("mismatch should return error")
	}
}

func Test_LineValidator_AllVerifyError_CollectsMultipleErrors(t *testing.T) {
	// Arrange: validator expects "ok", but all 3 inputs are wrong
	lv := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad1", LineNumber: 0},
		{Text: "bad2", LineNumber: 1},
		{Text: "bad3", LineNumber: 2},
	}

	// Act
	err := lv.AllVerifyError(params, items...)

	// Assert: error should be non-nil and contain all 3 failures
	if err == nil {
		t.Fatal("AllVerifyError should return error when all items fail")
	}

	errMsg := err.Error()
	for _, expected := range []string{"bad1", "bad2", "bad3"} {
		if !strings.Contains(errMsg, expected) {
			t.Errorf("AllVerifyError should collect all errors, missing '%s' in:\n%s", expected, errMsg)
		}
	}
}

func Test_LineValidator_AllVerifyError_FirstFailOthersPass(t *testing.T) {
	// Arrange: validator expects "ok", first fails, rest pass
	lv := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
		{Text: "ok", LineNumber: 2},
	}

	// Act
	err := lv.AllVerifyError(params, items...)

	// Assert: should still report the one failure
	if err == nil {
		t.Fatal("AllVerifyError should return error when any item fails")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "bad") {
		t.Errorf("error should mention the failed content 'bad', got:\n%s", errMsg)
	}
}

// ==========================================
// LinesValidators.AsBasicSliceContractsBinder
// ==========================================

func Test_LinesValidators_AsBasicSliceContractsBinder(t *testing.T) {
	lv := corevalidator.NewLinesValidators(1)
	binder := lv.AsBasicSliceContractsBinder()
	if binder == nil {
		t.Error("should not be nil")
	}
}
