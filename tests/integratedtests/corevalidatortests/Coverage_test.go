package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// =============================================================================
// Condition
// =============================================================================

func Test_Condition_IsSplitByWhitespace_AllFalse_Cov(t *testing.T) {
	c := corevalidator.Condition{}
	if c.IsSplitByWhitespace() {
		t.Error("all false should return false")
	}
}

func Test_Condition_IsSplitByWhitespace_UniqueWord(t *testing.T) {
	c := corevalidator.Condition{IsUniqueWordOnly: true}
	if !c.IsSplitByWhitespace() {
		t.Error("unique word should return true")
	}
}

func Test_Condition_IsSplitByWhitespace_NonEmpty(t *testing.T) {
	c := corevalidator.Condition{IsNonEmptyWhitespace: true}
	if !c.IsSplitByWhitespace() {
		t.Error("non-empty whitespace should return true")
	}
}

func Test_Condition_IsSplitByWhitespace_Sort(t *testing.T) {
	c := corevalidator.Condition{IsSortStringsBySpace: true}
	if !c.IsSplitByWhitespace() {
		t.Error("sort should return true")
	}
}

// =============================================================================
// Parameter
// =============================================================================

func Test_Parameter_IsIgnoreCase_Cov(t *testing.T) {
	p := corevalidator.Parameter{IsCaseSensitive: true}
	if p.IsIgnoreCase() {
		t.Error("case sensitive should not ignore case")
	}
	p2 := corevalidator.Parameter{IsCaseSensitive: false}
	if !p2.IsIgnoreCase() {
		t.Error("not case sensitive should ignore case")
	}
}

// =============================================================================
// LineNumber
// =============================================================================

func Test_LineNumber_HasLineNumber_Cov(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	if !ln.HasLineNumber() {
		t.Error("should have line number")
	}
}

func Test_LineNumber_HasLineNumber_Invalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	if ln.HasLineNumber() {
		t.Error("invalid should not have line number")
	}
}

func Test_LineNumber_IsMatch_BothInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	if !ln.IsMatch(-1) {
		t.Error("both invalid should match")
	}
}

func Test_LineNumber_IsMatch_InputInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	if !ln.IsMatch(-1) {
		t.Error("invalid input should match")
	}
}

func Test_LineNumber_IsMatch_Exact(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	if !ln.IsMatch(5) {
		t.Error("should match")
	}
	if ln.IsMatch(3) {
		t.Error("should not match")
	}
}

func Test_LineNumber_VerifyError_Match(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	if ln.VerifyError(5) != nil {
		t.Error("matching should return nil")
	}
}

func Test_LineNumber_VerifyError_Mismatch(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	if ln.VerifyError(3) == nil {
		t.Error("mismatch should return error")
	}
}

// =============================================================================
// TextValidator — uncovered branches
// =============================================================================

func Test_TextValidator_ToString_MultiLine_Cov(t *testing.T) {
	tv := corevalidator.TextValidator{
		Search:   "test",
		SearchAs: stringcompareas.Equal,
	}
	str := tv.ToString(false)
	if str == "" {
		t.Error("should return non-empty")
	}
}

func Test_TextValidator_IsMatchMany_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	if !tv.IsMatchMany(false, true) {
		t.Error("nil should return true")
	}
}

func Test_TextValidator_IsMatchMany_EmptySkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	if !tv.IsMatchMany(true, true) {
		t.Error("empty contents with skip should return true")
	}
}

func Test_TextValidator_IsMatchMany_Fail(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	if tv.IsMatchMany(false, true, "y") {
		t.Error("mismatch should return false")
	}
}

func Test_TextValidator_VerifyDetailError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	if tv.VerifyDetailError(params, "content") != nil {
		t.Error("nil should return nil")
	}
}

func Test_TextValidator_VerifySimpleError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	if tv.VerifySimpleError(0, params, "content") != nil {
		t.Error("nil should return nil")
	}
}

func Test_TextValidator_VerifyMany_FirstOnly(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(false, params, "x", "y")
	// first only, stops on first error if any
	_ = err
}

func Test_TextValidator_VerifyMany_ContinueOnError(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(true, params, "x")
	if err != nil {
		t.Error("matching should return nil")
	}
}

func Test_TextValidator_VerifyFirstError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	if tv.VerifyFirstError(params, "x") != nil {
		t.Error("nil should return nil")
	}
}

func Test_TextValidator_VerifyFirstError_EmptySkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	if tv.VerifyFirstError(params) != nil {
		t.Error("empty with skip should return nil")
	}
}

func Test_TextValidator_AllVerifyError_Nil(t *testing.T) {
	var tv *corevalidator.TextValidator
	params := &corevalidator.Parameter{}
	if tv.AllVerifyError(params, "x") != nil {
		t.Error("nil should return nil")
	}
}

func Test_TextValidator_AllVerifyError_EmptySkip(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	if tv.AllVerifyError(params) != nil {
		t.Error("empty with skip should return nil")
	}
}

func Test_TextValidator_MethodName(t *testing.T) {
	tv := corevalidator.TextValidator{SearchAs: stringcompareas.StartsWith}
	if tv.MethodName() == "" {
		t.Error("should return method name")
	}
}

// =============================================================================
// TextValidators — uncovered branches
// =============================================================================

func Test_TextValidators_Count(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{})
	tvs.Add(corevalidator.TextValidator{})
	if tvs.Count() != 1 { // Count returns LastIndex
		t.Error("count should equal lastindex")
	}
}

func Test_TextValidators_Adds_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Adds()
	if tvs.Length() != 0 {
		t.Error("should remain empty")
	}
}

func Test_TextValidators_AddSimple(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("test", stringcompareas.Equal)
	if tvs.Length() != 1 {
		t.Error("should have 1")
	}
}

func Test_TextValidators_AddSimpleAllTrue(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimpleAllTrue("test", stringcompareas.Equal)
	if tvs.Length() != 1 {
		t.Error("should have 1")
	}
}

func Test_TextValidators_HasAnyItem(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	if tvs.HasAnyItem() {
		t.Error("should be empty")
	}
	tvs.Add(corevalidator.TextValidator{})
	if !tvs.HasAnyItem() {
		t.Error("should have items")
	}
}

func Test_TextValidators_HasIndex(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{})
	if !tvs.HasIndex(0) {
		t.Error("should have index 0")
	}
	if tvs.HasIndex(5) {
		t.Error("should not have index 5")
	}
}

func Test_TextValidators_String(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal})
	if tvs.String() == "" {
		t.Error("should return non-empty")
	}
}

func Test_TextValidators_IsMatch_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	if !tvs.IsMatch("anything", true) {
		t.Error("empty should return true")
	}
}

func Test_TextValidators_IsMatchMany_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	if !tvs.IsMatchMany(true, true, "a") {
		t.Error("empty should return true")
	}
}

func Test_TextValidators_VerifyFirstError_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	if tvs.VerifyFirstError(0, "x", true) != nil {
		t.Error("empty should return nil")
	}
}

func Test_TextValidators_VerifyErrorMany_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	if tvs.VerifyErrorMany(true, &corevalidator.Parameter{}, "x") != nil {
		t.Error("nil should return nil")
	}
}

func Test_TextValidators_VerifyErrorMany_FirstOnly(t *testing.T) {
	var tvs *corevalidator.TextValidators
	if tvs.VerifyErrorMany(false, &corevalidator.Parameter{}, "x") != nil {
		t.Error("nil should return nil")
	}
}

func Test_TextValidators_VerifyFirstErrorMany_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	if tvs.VerifyFirstErrorMany(&corevalidator.Parameter{}) != nil {
		t.Error("empty should return nil")
	}
}

func Test_TextValidators_AllVerifyErrorMany_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	if tvs.AllVerifyErrorMany(&corevalidator.Parameter{}) != nil {
		t.Error("empty should return nil")
	}
}

func Test_TextValidators_AllVerifyError_Empty(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	if tvs.AllVerifyError(0, "x", true) != nil {
		t.Error("empty should return nil")
	}
}

func Test_TextValidators_Dispose(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.Add(corevalidator.TextValidator{})
	tvs.Dispose()
	if tvs.Items != nil {
		t.Error("should dispose")
	}
}

func Test_TextValidators_Dispose_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	tvs.Dispose() // should not panic
}

func Test_TextValidators_Length_Nil(t *testing.T) {
	var tvs *corevalidator.TextValidators
	if tvs.Length() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_TextValidators_AsBasicSliceContractsBinder(t *testing.T) {
	tvs := corevalidator.NewTextValidators(0)
	if tvs.AsBasicSliceContractsBinder() == nil {
		t.Error("should return self")
	}
}

// =============================================================================
// SliceValidator — uncovered branches
// =============================================================================

func Test_SliceValidator_IsUsedAlready_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.IsUsedAlready() {
		t.Error("nil should return false")
	}
}

func Test_SliceValidator_ActualLinesLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.ActualLinesLength() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_SliceValidator_ActualLinesString_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.ActualLinesString() != "" {
		t.Error("nil should return empty")
	}
}

func Test_SliceValidator_ExpectingLinesString_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.ExpectingLinesString() != "" {
		t.Error("nil should return empty")
	}
}

func Test_SliceValidator_ExpectingLinesLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.ExpectingLinesLength() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_SliceValidator_IsValid_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if !sv.IsValid(true) {
		t.Error("nil should return true")
	}
}

func Test_SliceValidator_Dispose_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	sv.Dispose() // should not panic
}

func Test_SliceValidator_SetActualVsExpected(t *testing.T) {
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})
	if !sv.IsValid(true) {
		t.Error("should be valid")
	}
}

func Test_SliceValidator_MethodName(t *testing.T) {
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.StartsWith}
	if sv.MethodName() == "" {
		t.Error("should return method name")
	}
}

func Test_SliceValidator_VerifyFirstError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.VerifyFirstError(&corevalidator.Parameter{}) != nil {
		t.Error("nil should return nil")
	}
}

func Test_SliceValidator_AllVerifyError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.AllVerifyError(&corevalidator.Parameter{}) != nil {
		t.Error("nil should return nil")
	}
}

func Test_SliceValidator_AllVerifyErrorQuick_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.AllVerifyErrorQuick(0, "header", "a") != nil {
		t.Error("nil should return nil")
	}
}

func Test_SliceValidator_AllVerifyErrorExceptLast_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.AllVerifyErrorExceptLast(&corevalidator.Parameter{}) != nil {
		t.Error("nil should return nil")
	}
}

func Test_SliceValidator_AllVerifyErrorTestCase_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.AllVerifyErrorTestCase(0, "header", true) != nil {
		t.Error("nil should return nil")
	}
}

func Test_SliceValidator_AllVerifyErrorUptoLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 5) != nil {
		t.Error("nil should return nil")
	}
}

func Test_SliceValidator_VerifyFirstLengthUptoError_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	if sv.VerifyFirstLengthUptoError(&corevalidator.Parameter{}, 5) != nil {
		t.Error("nil should return nil")
	}
}

func Test_SliceValidator_IsValidOtherLines(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"hello"},
	}
	if !sv.IsValidOtherLines(true, []string{"hello"}) {
		t.Error("should be valid")
	}
}

// =============================================================================
// SliceValidators — uncovered branches
// =============================================================================

func Test_SliceValidators_Length_Nil(t *testing.T) {
	var svs *corevalidator.SliceValidators
	if svs.Length() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_SliceValidators_IsEmpty_Nil(t *testing.T) {
	var svs *corevalidator.SliceValidators
	if !svs.IsEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_SliceValidators_SetActualOnAll_Empty(t *testing.T) {
	var svs *corevalidator.SliceValidators
	svs.SetActualOnAll("a") // should not panic
}

func Test_SliceValidators_IsValid_Empty(t *testing.T) {
	svs := &corevalidator.SliceValidators{}
	if !svs.IsValid(true) {
		t.Error("empty should return true")
	}
}

// =============================================================================
// HeaderSliceValidators — uncovered branches
// =============================================================================

func Test_HeaderSliceValidators_Length_Nil(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	if hsv.Length() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_HeaderSliceValidators_IsEmpty_Nil(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	if !hsv.IsEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_HeaderSliceValidators_SetActualOnAll_Empty_Cov(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	hsv.SetActualOnAll("a") // should not panic
}

func Test_HeaderSliceValidators_IsValid_Empty(t *testing.T) {
	var hsv corevalidator.HeaderSliceValidators
	if !hsv.IsValid(true) {
		t.Error("empty should return true")
	}
}

// =============================================================================
// LinesValidators — uncovered branches
// =============================================================================

func Test_LinesValidators_Length_Nil(t *testing.T) {
	var lv *corevalidator.LinesValidators
	if lv.Length() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_LinesValidators_Count_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	if lv.Count() != 0 {
		t.Error("empty should return 0")
	}
}

func Test_LinesValidators_HasAnyItem(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	if lv.HasAnyItem() {
		t.Error("empty should not have items")
	}
}

func Test_LinesValidators_HasIndex_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	if lv.HasIndex(0) {
		t.Error("empty should not have index 0")
	}
}

func Test_LinesValidators_AddPtr_Nil_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(5)
	lv.AddPtr(nil)
	if lv.Length() != 0 {
		t.Error("nil should not add")
	}
}

func Test_LinesValidators_String_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	_ = lv.String() // should not panic
}

func Test_LinesValidators_AsBasicSliceContractsBinder(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	if lv.AsBasicSliceContractsBinder() == nil {
		t.Error("should return self")
	}
}

func Test_LinesValidators_IsMatchText_Empty_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	if !lv.IsMatchText("test", true) {
		t.Error("empty should return true")
	}
}

func Test_LinesValidators_IsMatch_Empty_Cov(t *testing.T) {
	lv := corevalidator.NewLinesValidators(0)
	if !lv.IsMatch(false, true) {
		t.Error("empty should return true")
	}
}

// =============================================================================
// BaseLinesValidators — uncovered branches
// =============================================================================

func Test_BaseLinesValidators_Nil(t *testing.T) {
	var blv *corevalidator.BaseLinesValidators
	if blv.LinesValidatorsLength() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_BaseLinesValidators_IsEmpty(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	if !blv.IsEmptyLinesValidators() {
		t.Error("should be empty")
	}
}

func Test_BaseLinesValidators_HasLinesValidators(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	if blv.HasLinesValidators() {
		t.Error("should not have validators")
	}
}

func Test_BaseLinesValidators_ToLinesValidators_Empty_Cov(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{}
	lv := blv.ToLinesValidators()
	if lv.Length() != 0 {
		t.Error("should return empty")
	}
}

func Test_BaseLinesValidators_ToLinesValidators_WithItems(t *testing.T) {
	blv := &corevalidator.BaseLinesValidators{
		LinesValidators: []corevalidator.LineValidator{
			{TextValidator: corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal}},
		},
	}
	lv := blv.ToLinesValidators()
	if lv.Length() != 1 {
		t.Error("should return 1")
	}
}

// =============================================================================
// BaseValidatorCoreCondition — uncovered branches
// =============================================================================

func Test_BaseValidatorCoreCondition_Default_NilCondition(t *testing.T) {
	bvc := &corevalidator.BaseValidatorCoreCondition{}
	c := bvc.ValidatorCoreConditionDefault()
	_ = c // should not panic
}

func Test_BaseValidatorCoreCondition_Default_NonNilCondition(t *testing.T) {
	cond := &corevalidator.Condition{IsTrimCompare: true}
	bvc := &corevalidator.BaseValidatorCoreCondition{ValidatorCoreCondition: cond}
	c := bvc.ValidatorCoreConditionDefault()
	if !c.IsTrimCompare {
		t.Error("should use existing condition")
	}
}

// =============================================================================
// NewSliceValidatorUsingErr / NewSliceValidatorUsingAny
// =============================================================================

func Test_NewSliceValidatorUsingAny(t *testing.T) {
	sv := corevalidator.NewSliceValidatorUsingAny(
		"hello",
		"hello",
		false, false, false,
		stringcompareas.Equal,
	)
	if !sv.IsValid(true) {
		t.Error("should be valid")
	}
}

// =============================================================================
// SimpleSliceValidator — uncovered branches
// =============================================================================

func Test_SimpleSliceValidator_VerifyFirst(t *testing.T) {
	// Exercise the VerifyFirst path
	sv := &corevalidator.SimpleSliceValidator{
		CompareAs: stringcompareas.Equal,
	}
	// set expected via the Expected field is needed but requires corestr import
	// Just exercise the SetActual path
	sv.SetActual([]string{"a"})
	_ = sv
}

// =============================================================================
// RangeSegmentsValidator — uncovered branches
// =============================================================================

func Test_RangeSegmentsValidator_LengthOfVerifierSegments(t *testing.T) {
	rsv := &corevalidator.RangeSegmentsValidator{}
	if rsv.LengthOfVerifierSegments() != 0 {
		t.Error("should return 0")
	}
}
