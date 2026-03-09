package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// SliceValidators — collection basics
// ==========================================

func Test_SliceValidators_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	if !v.IsEmpty() {
		t.Error("empty should be empty")
	}
	if v.Length() != 0 {
		t.Errorf("expected 0, got %d", v.Length())
	}
}

// (nil receiver tests migrated to SliceValidators_NilReceiver_testcases.go)

func Test_SliceValidators_WithItems(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	if v.IsEmpty() {
		t.Error("should not be empty")
	}
	if v.Length() != 1 {
		t.Errorf("expected 1, got %d", v.Length())
	}
}

// ==========================================
// SliceValidators.IsMatch / IsValid
// ==========================================

func Test_SliceValidators_IsMatch_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	if !v.IsMatch(true) {
		t.Error("empty should match")
	}
}

func Test_SliceValidators_IsValid_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	if !v.IsValid(true) {
		t.Error("empty IsValid should be true")
	}
}

func Test_SliceValidators_IsMatch_AllPass(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a", "b"},
				ExpectedLines: []string{"a", "b"},
			},
		},
	}
	if !v.IsMatch(true) {
		t.Error("matching validators should return true")
	}
}

func Test_SliceValidators_IsMatch_OneFails(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"x"},
				ExpectedLines: []string{"y"},
			},
		},
	}
	if v.IsMatch(true) {
		t.Error("one failing validator should return false")
	}
}

// (nil receiver test migrated to SliceValidators_NilReceiver_testcases.go)

// ==========================================
// SliceValidators.VerifyAll
// ==========================================

func Test_SliceValidators_VerifyAll_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyAll("header", params, false)
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_SliceValidators_VerifyAll_Pass(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}
	err := v.VerifyAll("header", params, false)
	if err != nil {
		t.Errorf("matching should pass: %v", err)
	}
}

// ==========================================
// SliceValidators.VerifyAllError
// ==========================================

func Test_SliceValidators_VerifyAllError_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyAllError(params)
	if err != nil {
		t.Error("empty should return nil")
	}
}

// ==========================================
// SliceValidators.VerifyFirst
// ==========================================

func Test_SliceValidators_VerifyFirst_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyFirst(params, false)
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_SliceValidators_VerifyFirst_Pass(t *testing.T) {
	v := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				Condition:     corevalidator.DefaultDisabledCoreCondition,
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}
	err := v.VerifyFirst(params, false)
	if err != nil {
		t.Errorf("matching should pass: %v", err)
	}
}

// ==========================================
// SliceValidators.VerifyUpto
// ==========================================

func Test_SliceValidators_VerifyUpto_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyUpto(false, false, 1, params)
	if err != nil {
		t.Error("empty should return nil")
	}
}

// ==========================================
// SliceValidators.SetActualOnAll
// ==========================================

func Test_SliceValidators_SetActualOnAll_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	// should not panic
	v.SetActualOnAll("a", "b")
}

// ==========================================
// SliceValidators.VerifyAllErrorUsingActual
// ==========================================

func Test_SliceValidators_VerifyAllErrorUsingActual_Empty(t *testing.T) {
	v := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test"}
	err := v.VerifyAllErrorUsingActual(params, "a")
	if err != nil {
		t.Error("empty should return nil")
	}
}
