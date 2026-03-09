package corevalidatortests

import (
	"errors"
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// SliceValidator.AllVerifyErrorExceptLast
// ==========================================

func Test_SliceValidator_AllVerifyErrorExceptLast_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "different"},
		ExpectedLines: []string{"a", "b", "c"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.AllVerifyErrorExceptLast(params)
	if err != nil {
		t.Errorf("except last should pass: %v", err)
	}
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.AllVerifyErrorQuick
// ==========================================

func Test_SliceValidator_AllVerifyErrorQuick_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	err := v.AllVerifyErrorQuick(0, "test", "a", "b")
	if err != nil {
		t.Errorf("matching should pass: %v", err)
	}
}

func Test_SliceValidator_AllVerifyErrorQuick_Fail(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	err := v.AllVerifyErrorQuick(0, "test", "a", "x")
	if err == nil {
		t.Error("mismatch should return error")
	}
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.AllVerifyErrorTestCase
// ==========================================

func Test_SliceValidator_AllVerifyErrorTestCase_Pass(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	err := v.AllVerifyErrorTestCase(0, "test", true)
	if err != nil {
		t.Errorf("should pass: %v", err)
	}
}

func Test_SliceValidator_AllVerifyErrorTestCase_Fail(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	err := v.AllVerifyErrorTestCase(0, "test", true)
	if err == nil {
		t.Error("mismatch should return error")
	}
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.ComparingValidators caching
// ==========================================

func Test_SliceValidator_ComparingValidators_Cached(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	first := v.ComparingValidators()
	second := v.ComparingValidators()
	if first != second {
		t.Error("should return same cached instance")
	}
	if first.Length() != 2 {
		t.Errorf("expected 2 validators, got %d", first.Length())
	}
}

// ==========================================
// SliceValidator.ActualLinesString / ExpectingLinesString
// ==========================================

// (nil receiver tests migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_ActualLinesString_NonEmpty(t *testing.T) {
	v := corevalidator.SliceValidator{
		ActualLines: []string{"hello", "world"},
	}
	s := v.ActualLinesString()
	if s == "" {
		t.Error("should return non-empty string")
	}
}

func Test_SliceValidator_ExpectingLinesString_NonEmpty(t *testing.T) {
	v := corevalidator.SliceValidator{
		ExpectedLines: []string{"hello", "world"},
	}
	s := v.ExpectingLinesString()
	if s == "" {
		t.Error("should return non-empty string")
	}
}

// ==========================================
// SliceValidator.IsUsedAlready — nil receiver
// ==========================================

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// NewSliceValidatorUsingErr — with actual error
// ==========================================

func Test_NewSliceValidatorUsingErr_WithError(t *testing.T) {
	err := errors.New("line1\nline2\nline3")
	v := corevalidator.NewSliceValidatorUsingErr(
		err, "line1\nline2\nline3",
		false, false, false,
		stringcompareas.Equal,
	)
	if v == nil {
		t.Error("should not be nil")
	}
	if v.ActualLinesLength() != 3 {
		t.Errorf("expected 3 actual lines, got %d", v.ActualLinesLength())
	}
	if v.ExpectingLinesLength() != 3 {
		t.Errorf("expected 3 expected lines, got %d", v.ExpectingLinesLength())
	}
}

func Test_NewSliceValidatorUsingErr_WithConditions(t *testing.T) {
	err := errors.New("  hello  \n  world  ")
	v := corevalidator.NewSliceValidatorUsingErr(
		err, "hello\nworld",
		true, true, true,
		stringcompareas.Equal,
	)
	if !v.IsTrimCompare {
		t.Error("should have IsTrimCompare true")
	}
	if !v.IsNonEmptyWhitespace {
		t.Error("should have IsNonEmptyWhitespace true")
	}
	if !v.IsSortStringsBySpace {
		t.Error("should have IsSortStringsBySpace true")
	}
}

// ==========================================
// SliceValidator.UserInputsMergeWithError
// ==========================================

func Test_SliceValidator_UserInputsMergeWithError_NoAttach(t *testing.T) {
	v := corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "test",
		IsAttachUserInputs: false,
	}
	testErr := errors.New("test error")
	result := v.UserInputsMergeWithError(params, testErr)
	if result == nil {
		t.Error("should return error")
	}
	if result.Error() != "test error" {
		t.Errorf("without attach, should return original error, got: %s", result.Error())
	}
}

func Test_SliceValidator_UserInputsMergeWithError_WithAttach(t *testing.T) {
	v := corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "test",
		IsAttachUserInputs: true,
	}
	testErr := errors.New("test error")
	result := v.UserInputsMergeWithError(params, testErr)
	if result == nil {
		t.Error("should return error")
	}
	msg := result.Error()
	if msg == "test error" {
		t.Error("with attach, should include additional context")
	}
}

// ==========================================
// SliceValidator — isEmptyIgnoreCase boundary
// ==========================================

func Test_SliceValidator_AllVerifyError_EmptyActualNoSkip(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := v.AllVerifyError(params)
	if err == nil {
		t.Error("empty actual without skip should return error")
	}
}

// ==========================================
// TextValidators.AddSimpleAllTrue
// ==========================================

func Test_TextValidators_AddSimpleAllTrue(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.AddSimpleAllTrue("hello", stringcompareas.Contains)
	if v.Length() != 1 {
		t.Error("should add one validator")
	}
	item := v.Items[0]
	if !item.IsTrimCompare {
		t.Error("should have IsTrimCompare true")
	}
	if !item.IsUniqueWordOnly {
		t.Error("should have IsUniqueWordOnly true")
	}
	if !item.IsNonEmptyWhitespace {
		t.Error("should have IsNonEmptyWhitespace true")
	}
	if !item.IsSortStringsBySpace {
		t.Error("should have IsSortStringsBySpace true")
	}
}

// ==========================================
// TextValidators.AsBasicSliceContractsBinder
// ==========================================

func Test_TextValidators_AsBasicSliceContractsBinder(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	binder := v.AsBasicSliceContractsBinder()
	if binder == nil {
		t.Error("should not be nil")
	}
}

// ==========================================
// TextValidators.Count
// ==========================================

func Test_TextValidators_Count(t *testing.T) {
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})
	if v.Count() != 1 { // Count = LastIndex = Length-1
		t.Errorf("expected Count=1 (LastIndex), got %d", v.Count())
	}
}

// ==========================================
// TextValidator.VerifySimpleError
// ==========================================

func Test_TextValidator_VerifySimpleError_Match(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifySimpleError(0, params, "hello")
	if err != nil {
		t.Errorf("match should not error: %v", err)
	}
}

func Test_TextValidator_VerifySimpleError_Mismatch(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifySimpleError(0, params, "world")
	if err == nil {
		t.Error("mismatch should return error")
	}
}

func Test_TextValidator_VerifySimpleError_NilReceiver(t *testing.T) {
	var v *corevalidator.TextValidator
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.VerifySimpleError(0, params, "anything")
	if err != nil {
		t.Error("nil receiver should return nil")
	}
}

// ==========================================
// TextValidator.MethodName
// ==========================================

func Test_TextValidator_MethodName(t *testing.T) {
	v := corevalidator.TextValidator{SearchAs: stringcompareas.Contains}
	if v.MethodName() != "IsContains" {
		t.Errorf("expected 'IsContains', got '%s'", v.MethodName())
	}
}

// ==========================================
// TextValidator.ToString
// ==========================================

func Test_TextValidator_ToString_SingleLine(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "test",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	s := v.ToString(true)
	if s == "" {
		t.Error("should return non-empty string")
	}
}

func Test_TextValidator_ToString_MultiLine(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "test",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	s := v.ToString(false)
	if s == "" {
		t.Error("should return non-empty string")
	}
}

func Test_TextValidator_String(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "test",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	s := v.String()
	if s == "" {
		t.Error("should return non-empty string")
	}
}

// ==========================================
// TextValidator.GetCompiledTermBasedOnConditions
// ==========================================

func Test_TextValidator_GetCompiledTermBasedOnConditions_NoTrim(t *testing.T) {
	v := corevalidator.TextValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	result := v.GetCompiledTermBasedOnConditions("  hello  ", true)
	if result != "  hello  " {
		t.Errorf("no trim should return original, got '%s'", result)
	}
}

func Test_TextValidator_GetCompiledTermBasedOnConditions_WithTrim(t *testing.T) {
	v := corevalidator.TextValidator{
		Condition: corevalidator.DefaultTrimCoreCondition,
	}
	result := v.GetCompiledTermBasedOnConditions("  hello  ", true)
	if result != "hello" {
		t.Errorf("trim should return 'hello', got '%s'", result)
	}
}

// ==========================================
// TextValidators.VerifyFirstErrorMany
// ==========================================

func Test_TextValidators_VerifyFirstErrorMany_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.VerifyFirstErrorMany(params, "a")
	if err != nil {
		t.Error("empty validators should return nil")
	}
}

func Test_TextValidators_VerifyFirstErrorMany_Pass(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "a",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyFirstErrorMany(params, "a")
	if err != nil {
		t.Errorf("should pass: %v", err)
	}
}

func Test_TextValidators_AllVerifyErrorMany_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.AllVerifyErrorMany(params, "a")
	if err != nil {
		t.Error("empty validators should return nil")
	}
}

// ==========================================
// TextValidators.VerifyErrorMany — routing
// ==========================================

func Test_TextValidators_VerifyErrorMany_ContinueTrue(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyErrorMany(true, params, "a", "b")
	if err == nil {
		t.Error("mismatches should return error")
	}
}

func Test_TextValidators_VerifyErrorMany_ContinueFalse(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	err := v.VerifyErrorMany(false, params, "a", "b")
	if err == nil {
		t.Error("mismatches should return error")
	}
}

// ==========================================
// TextValidators.HasAnyItem
// ==========================================

func Test_TextValidators_HasAnyItem_Empty(t *testing.T) {
	v := corevalidator.NewTextValidators(0)
	if v.HasAnyItem() {
		t.Error("empty should not have items")
	}
}

func Test_TextValidators_HasAnyItem_NonEmpty(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	if !v.HasAnyItem() {
		t.Error("should have items")
	}
}

// ==========================================
// TextValidators.String
// ==========================================

func Test_TextValidators_String(t *testing.T) {
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	s := v.String()
	if s == "" {
		t.Error("should return non-empty string")
	}
}
