package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── SliceValidator ──

func Test_Cov7_SliceValidator_IsUsedAlready_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"used": sv.IsUsedAlready()}
	expected := args.Map{"used": false}
	expected.ShouldBeEqual(t, 0, "IsUsedAlready nil", actual)
}

func Test_Cov7_SliceValidator_ActualLinesLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"len": sv.ActualLinesLength()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ActualLinesLength nil", actual)
}

func Test_Cov7_SliceValidator_MethodName(t *testing.T) {
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	actual := args.Map{"notEmpty": sv.MethodName() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MethodName", actual)
}

func Test_Cov7_SliceValidator_SetActual(t *testing.T) {
	sv := &corevalidator.SliceValidator{}
	sv.SetActual([]string{"a", "b"})
	actual := args.Map{"used": sv.IsUsedAlready(), "len": sv.ActualLinesLength()}
	expected := args.Map{"used": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "SetActual", actual)
}

func Test_Cov7_SliceValidator_SetActualVsExpected(t *testing.T) {
	sv := &corevalidator.SliceValidator{}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})
	actual := args.Map{"used": sv.IsUsedAlready(), "actualLen": sv.ActualLinesLength(), "expectedLen": sv.ExpectingLinesLength()}
	expected := args.Map{"used": true, "actualLen": 1, "expectedLen": 1}
	expected.ShouldBeEqual(t, 0, "SetActualVsExpected", actual)
}

func Test_Cov7_SliceValidator_ActualLinesString(t *testing.T) {
	sv := &corevalidator.SliceValidator{ActualLines: []string{"line1", "line2"}}
	result := sv.ActualLinesString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ActualLinesString", actual)
}

func Test_Cov7_SliceValidator_ActualLinesString_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	result := sv.ActualLinesString()
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ActualLinesString nil", actual)
}

func Test_Cov7_SliceValidator_ExpectingLinesString(t *testing.T) {
	sv := &corevalidator.SliceValidator{ExpectedLines: []string{"line1"}}
	result := sv.ExpectingLinesString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingLinesString", actual)
}

func Test_Cov7_SliceValidator_ExpectingLinesString_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	result := sv.ExpectingLinesString()
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingLinesString nil", actual)
}

func Test_Cov7_SliceValidator_ExpectingLinesLength_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"len": sv.ExpectingLinesLength()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpectingLinesLength nil", actual)
}

func Test_Cov7_SliceValidator_ComparingValidators(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a", "b"},
		CompareAs:     stringcompareas.Equal,
	}
	validators := sv.ComparingValidators()
	actual := args.Map{"count": len(validators.Items)}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "ComparingValidators", actual)
}

func Test_Cov7_SliceValidator_ComparingValidators_Cached(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	v1 := sv.ComparingValidators()
	v2 := sv.ComparingValidators()
	actual := args.Map{"same": v1 == v2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "ComparingValidators cached", actual)
}

func Test_Cov7_SliceValidator_IsValid_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid nil", actual)
}

func Test_Cov7_SliceValidator_IsValid_Match(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello", "world"},
		ExpectedLines: []string{"hello", "world"},
		CompareAs:     stringcompareas.Equal,
	}
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid match", actual)
}

func Test_Cov7_SliceValidator_IsValid_Mismatch(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"world"},
		CompareAs:     stringcompareas.Equal,
	}
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "IsValid mismatch", actual)
}

func Test_Cov7_SliceValidator_IsValid_DifferentLengths(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.Equal,
	}
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "IsValid different lengths", actual)
}

func Test_Cov7_SliceValidator_IsValidOtherLines(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a", "b"},
		CompareAs:     stringcompareas.Equal,
	}
	actual := args.Map{
		"match":    sv.IsValidOtherLines(true, []string{"a", "b"}),
		"mismatch": sv.IsValidOtherLines(true, []string{"c", "d"}),
	}
	expected := args.Map{"match": true, "mismatch": false}
	expected.ShouldBeEqual(t, 0, "IsValidOtherLines", actual)
}

func Test_Cov7_SliceValidator_IsValidLines_BothNil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "isValidLines both nil", actual)
}

func Test_Cov7_SliceValidator_IsValidLines_LinesNilExpectedNil(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ExpectedLines: nil,
		CompareAs:     stringcompareas.EqualMatch,
	}
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "isValidLines both nil (non-nil receiver)", actual)
}

func Test_Cov7_SliceValidator_IsValidLines_OneNil(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.EqualMatch,
	}
	actual := args.Map{"valid": sv.IsValidOtherLines(true, nil)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "isValidLines one nil", actual)
}

func Test_Cov7_SliceValidator_Dispose(t *testing.T) {
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.EqualMatch,
	}
	_ = sv.ComparingValidators() // force lazy init
	sv.Dispose()
	actual := args.Map{"actualNil": sv.ActualLines == nil, "expectedNil": sv.ExpectedLines == nil}
	expected := args.Map{"actualNil": true, "expectedNil": true}
	expected.ShouldBeEqual(t, 0, "Dispose", actual)
}

func Test_Cov7_SliceValidator_Dispose_Nil(t *testing.T) {
	var sv *corevalidator.SliceValidator
	sv.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Dispose nil", actual)
}

// ── TextValidator ──

func Test_Cov7_TextValidator_IsMatch_EqualMatch(t *testing.T) {
	tv := corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.EqualMatch,
	}
	actual := args.Map{
		"match":    tv.IsMatch("hello", true),
		"mismatch": tv.IsMatch("world", true),
	}
	expected := args.Map{"match": true, "mismatch": false}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch equal", actual)
}

func Test_Cov7_TextValidator_IsMatch_Contains(t *testing.T) {
	tv := corevalidator.TextValidator{
		Search:   "ell",
		SearchAs: stringcompareas.Contains,
	}
	actual := args.Map{"match": tv.IsMatch("hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch contains", actual)
}

func Test_Cov7_TextValidator_IsMatch_StartsWith(t *testing.T) {
	tv := corevalidator.TextValidator{
		Search:   "hel",
		SearchAs: stringcompareas.StartsWith,
	}
	actual := args.Map{"match": tv.IsMatch("hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch startsWith", actual)
}

func Test_Cov7_TextValidator_IsMatch_EndsWith(t *testing.T) {
	tv := corevalidator.TextValidator{
		Search:   "llo",
		SearchAs: stringcompareas.EndsWith,
	}
	actual := args.Map{"match": tv.IsMatch("hello", true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidator.IsMatch endsWith", actual)
}

// ── TextValidators ──

func Test_Cov7_TextValidators_Add(t *testing.T) {
	validators := corevalidator.NewTextValidators(5)
	validators.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.EqualMatch})
	validators.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.EqualMatch})
	actual := args.Map{"count": len(validators.Items)}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "TextValidators.Add", actual)
}

func Test_Cov7_TextValidators_Dispose(t *testing.T) {
	validators := corevalidator.NewTextValidators(5)
	validators.Add(corevalidator.TextValidator{Search: "a"})
	validators.Dispose()
	actual := args.Map{"nil": validators.Items == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose", actual)
}

func Test_Cov7_TextValidators_Dispose_Nil(t *testing.T) {
	var validators *corevalidator.TextValidators
	validators.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TextValidators.Dispose nil", actual)
}

// ── Condition ──

func Test_Cov7_Condition_IsAnd(t *testing.T) {
	c := corevalidator.Condition{IsAnd: true}
	actual := args.Map{"isAnd": c.IsAnd}
	expected := args.Map{"isAnd": true}
	expected.ShouldBeEqual(t, 0, "Condition.IsAnd", actual)
}

// ── Parameter ──

func Test_Cov7_Parameter_Fields(t *testing.T) {
	p := &corevalidator.Parameter{
		IsSkipCompareOnActualEmpty: true,
	}
	actual := args.Map{"skip": p.IsSkipCompareOnActualEmpty}
	expected := args.Map{"skip": true}
	expected.ShouldBeEqual(t, 0, "Parameter fields", actual)
}

// ── SimpleSliceValidator ──

func Test_Cov7_SimpleSliceValidator_IsValid(t *testing.T) {
	sv := corevalidator.SimpleSliceValidator{
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a", "b"},
	}
	actual := args.Map{"valid": sv.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.IsValid", actual)
}

func Test_Cov7_SimpleSliceValidator_IsValid_Mismatch(t *testing.T) {
	sv := corevalidator.SimpleSliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	actual := args.Map{"valid": sv.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.IsValid mismatch", actual)
}

func Test_Cov7_SimpleSliceValidator_IsValid_DifferentLength(t *testing.T) {
	sv := corevalidator.SimpleSliceValidator{
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a"},
	}
	actual := args.Map{"valid": sv.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator.IsValid different length", actual)
}

// ── LineNumber ──

func Test_Cov7_LineNumber_Fields(t *testing.T) {
	ln := corevalidator.LineNumber{
		Index:      0,
		LineNumber: 1,
	}
	actual := args.Map{"idx": ln.Index, "num": ln.LineNumber}
	expected := args.Map{"idx": 0, "num": 1}
	expected.ShouldBeEqual(t, 0, "LineNumber fields", actual)
}

// ── HeaderSliceValidator ──

func Test_Cov7_HeaderSliceValidator_IsValid(t *testing.T) {
	hsv := corevalidator.HeaderSliceValidator{
		Header: "test-header",
		SliceValidator: corevalidator.SliceValidator{
			ActualLines:   []string{"a"},
			ExpectedLines: []string{"a"},
			CompareAs:     stringcompareas.EqualMatch,
		},
	}
	actual := args.Map{"valid": hsv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidator.IsValid", actual)
}

// ── HeaderSliceValidators ──

func Test_Cov7_HeaderSliceValidators_IsEmpty(t *testing.T) {
	hsvs := &corevalidator.HeaderSliceValidators{}
	actual := args.Map{"empty": hsvs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.IsEmpty", actual)
}

func Test_Cov7_HeaderSliceValidators_Add(t *testing.T) {
	hsvs := corevalidator.NewHeaderSliceValidators(5)
	hsvs.Add(corevalidator.HeaderSliceValidator{Header: "h1"})
	actual := args.Map{"count": hsvs.Length()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidators.Add", actual)
}

// ── SliceValidators ──

func Test_Cov7_SliceValidators_IsEmpty(t *testing.T) {
	svs := corevalidator.NewSliceValidators(5)
	actual := args.Map{"empty": svs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SliceValidators.IsEmpty", actual)
}

func Test_Cov7_SliceValidators_Add(t *testing.T) {
	svs := corevalidator.NewSliceValidators(5)
	svs.Add(corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
		CompareAs:     stringcompareas.EqualMatch,
	})
	actual := args.Map{"count": svs.Length()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "SliceValidators.Add", actual)
}

// ── RangesSegment ──

func Test_Cov7_RangesSegment(t *testing.T) {
	rs := corevalidator.RangesSegment{
		Start: 0,
		End:   10,
	}
	actual := args.Map{"start": rs.Start, "end": rs.End}
	expected := args.Map{"start": 0, "end": 10}
	expected.ShouldBeEqual(t, 0, "RangesSegment", actual)
}
