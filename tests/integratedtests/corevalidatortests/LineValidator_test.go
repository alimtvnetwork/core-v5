package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// LineValidator.IsMatch
// ==========================================

func Test_LineValidator_IsMatch_BothMatch(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	if !v.IsMatch(0, "hello", true) {
		t.Error("line+text match should return true")
	}
}

func Test_LineValidator_IsMatch_LineNumberMismatch(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	if v.IsMatch(0, "hello", true) {
		t.Error("line mismatch should return false")
	}
}

func Test_LineValidator_IsMatch_TextMismatch(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	if v.IsMatch(0, "world", true) {
		t.Error("text mismatch should return false")
	}
}

func Test_LineValidator_IsMatch_SkipLineNumber(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	if !v.IsMatch(99, "hello", true) {
		t.Error("skip line number should pass with text match")
	}
}

// ==========================================
// LineValidator.IsMatchMany
// ==========================================

func Test_LineValidator_IsMatchMany_AllMatch(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
	}
	if !v.IsMatchMany(false, true, items...) {
		t.Error("all matching should return true")
	}
}

func Test_LineValidator_IsMatchMany_OneFails(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "nope", LineNumber: 1},
	}
	if v.IsMatchMany(false, true, items...) {
		t.Error("one failing should return false")
	}
}

func Test_LineValidator_IsMatchMany_EmptySkip(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	if !v.IsMatchMany(true, true) {
		t.Error("empty with skip should return true")
	}
}

func Test_LineValidator_IsMatchMany_NilReceiver(t *testing.T) {
	var v *corevalidator.LineValidator
	if !v.IsMatchMany(false, true, corestr.TextWithLineNumber{Text: "x"}) {
		t.Error("nil receiver should return true")
	}
}

// ==========================================
// LineValidator.VerifyError
// ==========================================

func Test_LineValidator_VerifyError_Match(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyError(params, 0, "hello")
	if err != nil {
		t.Errorf("match should not error: %v", err)
	}
}

func Test_LineValidator_VerifyError_LineNumberMismatch(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyError(params, 0, "hello")
	if err == nil {
		t.Error("line number mismatch should return error")
	}
}

func Test_LineValidator_VerifyError_TextMismatch(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyError(params, 0, "world")
	if err == nil {
		t.Error("text mismatch should return error")
	}
}

// ==========================================
// LineValidator.VerifyMany
// ==========================================

func Test_LineValidator_VerifyMany_ContinueOnError(t *testing.T) {
	v := corevalidator.LineValidator{
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
	}
	err := v.VerifyMany(true, params, items...)
	if err == nil {
		t.Error("should collect errors")
	}
}

func Test_LineValidator_VerifyMany_FirstOnly(t *testing.T) {
	v := corevalidator.LineValidator{
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
		{Text: "also bad", LineNumber: 1},
	}
	err := v.VerifyMany(false, params, items...)
	if err == nil {
		t.Error("should return first error")
	}
}
