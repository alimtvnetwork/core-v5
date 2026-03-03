package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// SliceValidator.IsValid
// ==========================================

func Test_SliceValidator_IsValid_ExactMatch(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "c"},
		ExpectedLines: []string{"a", "b", "c"},
	}
	if !v.IsValid(true) {
		t.Error("exact match should be valid")
	}
}

func Test_SliceValidator_IsValid_Mismatch(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a", "x"},
	}
	if v.IsValid(true) {
		t.Error("content mismatch should be invalid")
	}
}

func Test_SliceValidator_IsValid_LengthMismatch(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a", "b"},
	}
	if v.IsValid(true) {
		t.Error("length mismatch should be invalid")
	}
}

func Test_SliceValidator_IsValid_BothNil(t *testing.T) {
	v := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   nil,
		ExpectedLines: nil,
	}
	if !v.IsValid(true) {
		t.Error("both nil should be valid")
	}
}

func Test_SliceValidator_IsValid_OneNil(t *testing.T) {
	v := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   nil,
		ExpectedLines: []string{"a"},
	}
	if v.IsValid(true) {
		t.Error("one nil should be invalid")
	}
}

func Test_SliceValidator_IsValid_NilReceiver(t *testing.T) {
	var v *corevalidator.SliceValidator
	if !v.IsValid(true) {
		t.Error("nil receiver should return true")
	}
}

func Test_SliceValidator_IsValid_BothEmpty(t *testing.T) {
	v := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{},
	}
	if !v.IsValid(true) {
		t.Error("both empty should be valid")
	}
}

// ==========================================
// SliceValidator.IsValid — with Trim
// ==========================================

func Test_SliceValidator_IsValid_TrimMatch(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultTrimCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"  hello  ", " world "},
		ExpectedLines: []string{"hello", "world"},
	}
	if !v.IsValid(true) {
		t.Error("trimmed should match")
	}
}

// ==========================================
// SliceValidator.IsValid — Contains
// ==========================================

func Test_SliceValidator_IsValid_Contains(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Contains,
		ActualLines:   []string{"hello world", "foo bar"},
		ExpectedLines: []string{"ello", "bar"},
	}
	if !v.IsValid(true) {
		t.Error("contains should match substrings")
	}
}

// ==========================================
// SliceValidator — helper methods
// ==========================================

func Test_SliceValidator_ActualLinesLength(t *testing.T) {
	v := corevalidator.SliceValidator{
		ActualLines: []string{"a", "b"},
	}
	if v.ActualLinesLength() != 2 {
		t.Errorf("expected 2, got %d", v.ActualLinesLength())
	}
}

func Test_SliceValidator_ActualLinesLength_Nil(t *testing.T) {
	var v *corevalidator.SliceValidator
	if v.ActualLinesLength() != 0 {
		t.Error("nil receiver should return 0")
	}
}

func Test_SliceValidator_ExpectingLinesLength(t *testing.T) {
	v := corevalidator.SliceValidator{
		ExpectedLines: []string{"a", "b", "c"},
	}
	if v.ExpectingLinesLength() != 3 {
		t.Errorf("expected 3, got %d", v.ExpectingLinesLength())
	}
}

func Test_SliceValidator_IsUsedAlready_False(t *testing.T) {
	v := corevalidator.SliceValidator{
		ExpectedLines: []string{"a"},
	}
	if v.IsUsedAlready() {
		t.Error("fresh validator should not be used already")
	}
}

func Test_SliceValidator_IsUsedAlready_TrueAfterComparing(t *testing.T) {
	v := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
		ActualLines:   []string{"a"},
	}
	_ = v.ComparingValidators()
	if !v.IsUsedAlready() {
		t.Error("after ComparingValidators should be used")
	}
}

func Test_SliceValidator_MethodName(t *testing.T) {
	v := corevalidator.SliceValidator{CompareAs: stringcompareas.Contains}
	name := v.MethodName()
	if name != "IsContains" {
		t.Errorf("expected 'IsContains', got '%s'", name)
	}
}

// ==========================================
// SliceValidator.SetActual / SetActualVsExpected
// ==========================================

func Test_SliceValidator_SetActual(t *testing.T) {
	v := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}
	v.SetActual([]string{"a"})
	if v.ActualLinesLength() != 1 {
		t.Error("SetActual should set actual lines")
	}
}

func Test_SliceValidator_SetActualVsExpected(t *testing.T) {
	v := corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	v.SetActualVsExpected([]string{"a"}, []string{"b"})
	if v.ActualLinesLength() != 1 || v.ExpectingLinesLength() != 1 {
		t.Error("should set both actual and expected")
	}
}

// ==========================================
// SliceValidator.IsValidOtherLines
// ==========================================

func Test_SliceValidator_IsValidOtherLines_Match(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	if !v.IsValidOtherLines(true, []string{"a", "b"}) {
		t.Error("matching other lines should return true")
	}
}

func Test_SliceValidator_IsValidOtherLines_Mismatch(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	if v.IsValidOtherLines(true, []string{"a", "x"}) {
		t.Error("mismatching other lines should return false")
	}
}

// ==========================================
// SliceValidator.AllVerifyError
// ==========================================

func Test_SliceValidator_AllVerifyError_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.AllVerifyError(params)
	if err != nil {
		t.Errorf("matching should pass: %v", err)
	}
}

func Test_SliceValidator_AllVerifyError_Fail(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "x"},
		ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.AllVerifyError(params)
	if err == nil {
		t.Error("mismatch should return error")
	}
}

func Test_SliceValidator_AllVerifyError_NilReceiver(t *testing.T) {
	var v *corevalidator.SliceValidator
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.AllVerifyError(params)
	if err != nil {
		t.Error("nil receiver should return nil")
	}
}

func Test_SliceValidator_AllVerifyError_SkipEmpty(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := v.AllVerifyError(params)
	if err != nil {
		t.Errorf("skip empty should not error: %v", err)
	}
}

// ==========================================
// SliceValidator.VerifyFirstError
// ==========================================

func Test_SliceValidator_VerifyFirstError_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyFirstError(params)
	if err != nil {
		t.Errorf("matching should pass: %v", err)
	}
}

func Test_SliceValidator_VerifyFirstError_NilReceiver(t *testing.T) {
	var v *corevalidator.SliceValidator
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.VerifyFirstError(params)
	if err != nil {
		t.Error("nil receiver should return nil")
	}
}

// ==========================================
// SliceValidator.Dispose
// ==========================================

func Test_SliceValidator_Dispose(t *testing.T) {
	v := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	v.Dispose()
	if v.ActualLines != nil || v.ExpectedLines != nil {
		t.Error("Dispose should nil out lines")
	}
}

func Test_SliceValidator_Dispose_NilReceiver(t *testing.T) {
	var v *corevalidator.SliceValidator
	// should not panic
	v.Dispose()
}

// ==========================================
// SliceValidator — case insensitive
// ==========================================

func Test_SliceValidator_IsValid_CaseInsensitive(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"Hello", "WORLD"},
		ExpectedLines: []string{"hello", "world"},
	}
	if !v.IsValid(false) {
		t.Error("case-insensitive should match")
	}
}

func Test_SliceValidator_IsValid_CaseSensitiveFail(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"Hello"},
		ExpectedLines: []string{"hello"},
	}
	if v.IsValid(true) {
		t.Error("case-sensitive different case should not match")
	}
}

// ==========================================
// SliceValidator — NewSliceValidatorUsingErr
// ==========================================

func Test_NewSliceValidatorUsingErr_NilError(t *testing.T) {
	v := corevalidator.NewSliceValidatorUsingErr(
		nil, "expected\nlines",
		true, false, false,
		stringcompareas.Equal,
	)
	if v == nil {
		t.Error("should not be nil")
	}
	// nil error produces empty actual
	if v.ActualLinesLength() != 0 {
		t.Errorf("nil error should produce 0 actual lines, got %d", v.ActualLinesLength())
	}
}
