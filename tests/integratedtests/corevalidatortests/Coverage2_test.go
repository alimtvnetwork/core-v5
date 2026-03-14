package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// =============================================================================
// HeaderSliceValidator — uncovered branches
// =============================================================================

func Test_Cov2_HeaderSliceValidator_Create(t *testing.T) {
	hsv := corevalidator.HeaderSliceValidator{
		Header: "test-header",
		SliceValidator: corevalidator.SliceValidator{
			CompareAs:     stringcompareas.Equal,
			ExpectedLines: []string{"line1", "line2"},
		},
	}
	actual := args.Map{"header": hsv.Header, "linesLen": len(hsv.ExpectedLines)}
	expected := args.Map{"header": "test-header", "linesLen": 2}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidator returns expected -- valid input", actual)
}

// =============================================================================
// SliceValidator — additional branch coverage
// =============================================================================

func Test_Cov2_SliceValidator_SetActualVsExpected_Mismatch(t *testing.T) {
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	sv.SetActualVsExpected([]string{"a"}, []string{"b"})
	actual := args.Map{"isValid": sv.IsValid(true)}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator SetActualVsExpected returns invalid -- mismatch", actual)
}

func Test_Cov2_SliceValidator_VerifyFirstError_Valid(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.VerifyFirstError(params)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator VerifyFirstError returns nil -- matching", actual)
}

func Test_Cov2_SliceValidator_AllVerifyError_Valid(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.AllVerifyError(params)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator AllVerifyError returns nil -- matching", actual)
}

func Test_Cov2_SliceValidator_ActualLines_Valid(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	sv.SetActualVsExpected([]string{"a", "b"}, []string{"a", "b"})
	actual := args.Map{
		"actualLen":    sv.ActualLinesLength(),
		"expectedLen":  sv.ExpectingLinesLength(),
		"actualStr":    sv.ActualLinesString() != "",
		"expectedStr":  sv.ExpectingLinesString() != "",
		"isUsed":       sv.IsUsedAlready(),
	}
	expected := args.Map{
		"actualLen": 2, "expectedLen": 2,
		"actualStr": true, "expectedStr": true, "isUsed": false,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidator ActualLines returns expected -- 2 lines", actual)
}

// =============================================================================
// TextValidator — additional branch coverage
// =============================================================================

func Test_Cov2_TextValidator_IsMatch_Equal_True(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- equal", actual)
}

func Test_Cov2_TextValidator_IsMatch_Equal_False(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}
	actual := args.Map{"isMatch": tv.IsMatch("world", true)}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns false -- not equal", actual)
}

func Test_Cov2_TextValidator_IsMatch_StartsWith(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "hel", SearchAs: stringcompareas.StartsWith}
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- starts with", actual)
}

func Test_Cov2_TextValidator_IsMatch_EndsWith(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "llo", SearchAs: stringcompareas.EndsWith}
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- ends with", actual)
}

func Test_Cov2_TextValidator_IsMatch_Contains(t *testing.T) {
	tv := &corevalidator.TextValidator{Search: "ell", SearchAs: stringcompareas.Contains}
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- contains", actual)
}

func Test_Cov2_TextValidator_ToString_SingleLine(t *testing.T) {
	tv := corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal}
	actual := args.Map{"notEmpty": tv.ToString(true) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator ToString returns non-empty -- single line", actual)
}

// =============================================================================
// TextValidators — additional branch coverage
// =============================================================================

func Test_Cov2_TextValidators_IsMatch_WithItems(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	actual := args.Map{
		"matchTrue":  tvs.IsMatch("hello", true),
		"matchFalse": tvs.IsMatch("world", true),
	}
	expected := args.Map{"matchTrue": true, "matchFalse": false}
	expected.ShouldBeEqual(t, 0, "TextValidators IsMatch returns expected -- with validator", actual)
}

func Test_Cov2_TextValidators_VerifyFirstError_WithMatch(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.VerifyFirstError(0, "hello", true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TextValidators VerifyFirstError returns nil -- matching", actual)
}

func Test_Cov2_TextValidators_AllVerifyError_WithMatch(t *testing.T) {
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.AllVerifyError(0, "hello", true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TextValidators AllVerifyError returns nil -- matching", actual)
}

// =============================================================================
// RangeSegmentsValidator — uncovered branches
// =============================================================================

func Test_Cov2_RangeSegmentsValidator_Create(t *testing.T) {
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
	}
	actual := args.Map{"title": rsv.Title, "segLen": rsv.LengthOfVerifierSegments()}
	expected := args.Map{"title": "test-range", "segLen": 0}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns expected -- basic", actual)
}

// =============================================================================
// Condition — additional branches
// =============================================================================

func Test_Cov2_Condition_AllTrue(t *testing.T) {
	c := corevalidator.Condition{
		IsUniqueWordOnly:     true,
		IsNonEmptyWhitespace: true,
		IsSortStringsBySpace: true,
	}
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}
	expected := args.Map{"isSplit": true}
	expected.ShouldBeEqual(t, 0, "Condition IsSplitByWhitespace returns true -- all true", actual)
}

// =============================================================================
// Parameter — additional branches
// =============================================================================

func Test_Cov2_Parameter_SkipOnEmpty(t *testing.T) {
	p := corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	actual := args.Map{"isSkip": p.IsSkipCompareOnActualEmpty}
	expected := args.Map{"isSkip": true}
	expected.ShouldBeEqual(t, 0, "Parameter IsSkipCompareOnActualEmpty returns true -- set", actual)
}
