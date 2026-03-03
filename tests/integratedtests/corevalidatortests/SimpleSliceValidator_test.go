package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// SimpleSliceValidator.SetActual
// ==========================================

func Test_SimpleSliceValidator_SetActual(t *testing.T) {
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a", "b"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	result := v.SetActual([]string{"a", "b"})
	if result != v {
		t.Error("SetActual should return same instance")
	}
}

// ==========================================
// SimpleSliceValidator.SliceValidator
// ==========================================

func Test_SimpleSliceValidator_SliceValidator_NotNil(t *testing.T) {
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual([]string{"a"})
	sv := v.SliceValidator()
	if sv == nil {
		t.Error("SliceValidator should not be nil")
	}
}

// ==========================================
// SimpleSliceValidator.VerifyAll
// ==========================================

func Test_SimpleSliceValidator_VerifyAll_Pass(t *testing.T) {
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a", "b"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual([]string{"a", "b"})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyAll([]string{"a", "b"}, params)
	if err != nil {
		t.Errorf("matching should pass: %v", err)
	}
}

func Test_SimpleSliceValidator_VerifyAll_Fail(t *testing.T) {
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a", "b"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual([]string{"x", "y"})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyAll([]string{"x", "y"}, params)
	if err == nil {
		t.Error("mismatch should return error")
	}
}

// ==========================================
// SimpleSliceValidator.VerifyFirst
// ==========================================

func Test_SimpleSliceValidator_VerifyFirst_Pass(t *testing.T) {
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual([]string{"a"})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyFirst([]string{"a"}, params)
	if err != nil {
		t.Errorf("matching should pass: %v", err)
	}
}

// ==========================================
// SimpleSliceValidator.VerifyUpto
// ==========================================

func Test_SimpleSliceValidator_VerifyUpto_Pass(t *testing.T) {
	expected := corestr.New.SimpleSlice.Direct(false, []string{"a", "b", "c"})
	v := &corevalidator.SimpleSliceValidator{
		Expected:  expected,
		Condition: corevalidator.DefaultDisabledCoreCondition,
		CompareAs: stringcompareas.Equal,
	}
	v.SetActual([]string{"a", "b", "c"})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyUpto([]string{"a", "b", "c"}, params, 2)
	if err != nil {
		t.Errorf("matching upto should pass: %v", err)
	}
}
