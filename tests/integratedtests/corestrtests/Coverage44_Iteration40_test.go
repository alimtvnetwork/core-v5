package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/issetter"
)

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Split methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov44_SSO_SplitLeftRight_HasBoth(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("key=val")
	left, right := sso.SplitLeftRight("=")
	tc := coretestcases.CaseV1{Name: "SplitLeftRight left", Expected: "key", Actual: left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "SplitLeftRight right", Expected: "val", Actual: right, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SplitLeftRight_NoSep(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("nosep")
	left, right := sso.SplitLeftRight("=")
	tc := coretestcases.CaseV1{Name: "SplitLeftRight no sep left", Expected: "nosep", Actual: left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "SplitLeftRight no sep right", Expected: "", Actual: right, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SplitLeftRightTrim_HasBoth(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("  key = val  ")
	left, right := sso.SplitLeftRightTrim("=")
	tc := coretestcases.CaseV1{Name: "SplitLeftRightTrim left", Expected: "key", Actual: left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "SplitLeftRightTrim right", Expected: "val", Actual: right, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SplitLeftRightTrim_NoSep(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("  nosep  ")
	left, right := sso.SplitLeftRightTrim("=")
	tc := coretestcases.CaseV1{Name: "SplitLeftRightTrim no sep left", Expected: "nosep", Actual: left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "SplitLeftRightTrim no sep right", Expected: "", Actual: right, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_LinesSimpleSlice(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("a\nb\nc")
	ss := sso.LinesSimpleSlice()
	tc := coretestcases.CaseV1{Name: "LinesSimpleSlice", Expected: 3, Actual: ss.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SimpleSlice(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("a,b,c")
	ss := sso.SimpleSlice(",")
	tc := coretestcases.CaseV1{Name: "SimpleSlice", Expected: 3, Actual: ss.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Split(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("a-b-c")
	result := sso.Split("-")
	tc := coretestcases.CaseV1{Name: "Split", Expected: 3, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SplitNonEmpty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("a,,b")
	result := sso.SplitNonEmpty(",")
	tc := coretestcases.CaseV1{Name: "SplitNonEmpty", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SplitTrimNonWhitespace(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("a, ,b")
	result := sso.SplitTrimNonWhitespace(",")
	tc := coretestcases.CaseV1{Name: "SplitTrimNonWhitespace", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — String/Match methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov44_SSO_Is(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	tc := coretestcases.CaseV1{Name: "Is match", Expected: true, Actual: sso.Is("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsAnyOf_Found(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("b")
	tc := coretestcases.CaseV1{Name: "IsAnyOf found", Expected: true, Actual: sso.IsAnyOf("a", "b", "c"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsAnyOf_NotFound(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("z")
	tc := coretestcases.CaseV1{Name: "IsAnyOf not found", Expected: false, Actual: sso.IsAnyOf("a", "b"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsAnyOf_Empty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	tc := coretestcases.CaseV1{Name: "IsAnyOf empty returns true", Expected: true, Actual: sso.IsAnyOf(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsContains(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello world")
	tc := coretestcases.CaseV1{Name: "IsContains", Expected: true, Actual: sso.IsContains("world"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsAnyContains_Found(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello world")
	tc := coretestcases.CaseV1{Name: "IsAnyContains found", Expected: true, Actual: sso.IsAnyContains("xyz", "world"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsAnyContains_NotFound(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("hello")
	tc := coretestcases.CaseV1{Name: "IsAnyContains not found", Expected: false, Actual: sso.IsAnyContains("xyz", "abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsAnyContains_Empty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("x")
	tc := coretestcases.CaseV1{Name: "IsAnyContains empty", Expected: true, Actual: sso.IsAnyContains(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsEqualNonSensitive(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("Hello")
	tc := coretestcases.CaseV1{Name: "IsEqualNonSensitive", Expected: true, Actual: sso.IsEqualNonSensitive("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsRegexMatches_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc123")
	re := regexp.MustCompile(`\d+`)
	tc := coretestcases.CaseV1{Name: "IsRegexMatches valid", Expected: true, Actual: sso.IsRegexMatches(re), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsRegexMatches_Nil(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc")
	tc := coretestcases.CaseV1{Name: "IsRegexMatches nil", Expected: false, Actual: sso.IsRegexMatches(nil), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_RegexFindString_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc123def")
	re := regexp.MustCompile(`\d+`)
	tc := coretestcases.CaseV1{Name: "RegexFindString", Expected: "123", Actual: sso.RegexFindString(re), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_RegexFindString_Nil(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc")
	tc := coretestcases.CaseV1{Name: "RegexFindString nil", Expected: "", Actual: sso.RegexFindString(nil), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_RegexFindAllStrings_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("a1b2c3")
	re := regexp.MustCompile(`\d`)
	result := sso.RegexFindAllStrings(re, -1)
	tc := coretestcases.CaseV1{Name: "RegexFindAllStrings", Expected: 3, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_RegexFindAllStrings_Nil(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc")
	result := sso.RegexFindAllStrings(nil, -1)
	tc := coretestcases.CaseV1{Name: "RegexFindAllStrings nil", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_RegexFindAllStringsWithFlag_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("a1b2c3")
	re := regexp.MustCompile(`\d`)
	items, hasAny := sso.RegexFindAllStringsWithFlag(re, -1)
	tc := coretestcases.CaseV1{Name: "RegexFindAllStringsWithFlag hasAny", Expected: true, Actual: hasAny, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "RegexFindAllStringsWithFlag count", Expected: 3, Actual: len(items), Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc")
	_, hasAny := sso.RegexFindAllStringsWithFlag(nil, -1)
	tc := coretestcases.CaseV1{Name: "RegexFindAllStringsWithFlag nil", Expected: false, Actual: hasAny, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Numeric conversions
// ═══════════════════════════════════════════════════════════════

func Test_Cov44_SSO_Int16_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("100")
	tc := coretestcases.CaseV1{Name: "Int16 valid", Expected: int16(100), Actual: sso.Int16(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Int16_OutOfRange(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("99999")
	tc := coretestcases.CaseV1{Name: "Int16 out of range", Expected: int16(0), Actual: sso.Int16(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Int16_Error(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc")
	tc := coretestcases.CaseV1{Name: "Int16 error", Expected: int16(0), Actual: sso.Int16(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Int32_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("5000")
	tc := coretestcases.CaseV1{Name: "Int32 valid", Expected: int32(5000), Actual: sso.Int32(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Int32_Error(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc")
	tc := coretestcases.CaseV1{Name: "Int32 error", Expected: int32(0), Actual: sso.Int32(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Uint16_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("500")
	val, isInRange := sso.Uint16()
	tc := coretestcases.CaseV1{Name: "Uint16 valid", Expected: uint16(500), Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Uint16 inRange", Expected: true, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Uint32_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("1000")
	val, isInRange := sso.Uint32()
	tc := coretestcases.CaseV1{Name: "Uint32 valid", Expected: uint32(1000), Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Uint32 inRange", Expected: true, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_WithinRange_InRange(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("5")
	val, isInRange := sso.WithinRange(true, 0, 10)
	tc := coretestcases.CaseV1{Name: "WithinRange inRange", Expected: 5, Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "WithinRange isInRange", Expected: true, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_WithinRange_BelowMin(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("-5")
	val, isInRange := sso.WithinRange(true, 0, 10)
	tc := coretestcases.CaseV1{Name: "WithinRange below min", Expected: 0, Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "WithinRange below isInRange", Expected: false, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_WithinRange_AboveMax(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("15")
	val, isInRange := sso.WithinRange(true, 0, 10)
	tc := coretestcases.CaseV1{Name: "WithinRange above max", Expected: 10, Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "WithinRange above isInRange", Expected: false, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_WithinRange_NoBoundary(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("-5")
	val, isInRange := sso.WithinRange(false, 0, 10)
	tc := coretestcases.CaseV1{Name: "WithinRange noBoundary val", Expected: -5, Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "WithinRange noBoundary isInRange", Expected: false, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_WithinRange_ParseErr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("abc")
	val, isInRange := sso.WithinRange(true, 0, 10)
	tc := coretestcases.CaseV1{Name: "WithinRange parseErr", Expected: 0, Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "WithinRange parseErr isInRange", Expected: false, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_WithinRangeDefault(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("5")
	val, isInRange := sso.WithinRangeDefault(0, 10)
	tc := coretestcases.CaseV1{Name: "WithinRangeDefault", Expected: 5, Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "WithinRangeDefault isInRange", Expected: true, Actual: isInRange, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Boolean / IsSetter
// ═══════════════════════════════════════════════════════════════

func Test_Cov44_SSO_BooleanDefault_True(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("yes")
	tc := coretestcases.CaseV1{Name: "BooleanDefault true", Expected: true, Actual: sso.BooleanDefault(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Boolean_ConsiderInit_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("true")
	tc := coretestcases.CaseV1{Name: "Boolean considerInit uninit", Expected: false, Actual: sso.Boolean(true), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Boolean_YesValues(t *testing.T) {
	for _, val := range []string{"yes", "y", "1", "YES", "Y"} {
		sso := corestr.New.SimpleStringOnce.Init(val)
		tc := coretestcases.CaseV1{Name: "Boolean " + val, Expected: true, Actual: sso.Boolean(false), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	}
}

func Test_Cov44_SSO_Boolean_ParseTrue(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("true")
	tc := coretestcases.CaseV1{Name: "Boolean parse true", Expected: true, Actual: sso.Boolean(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Boolean_ParseErr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("notbool")
	tc := coretestcases.CaseV1{Name: "Boolean parse err", Expected: false, Actual: sso.Boolean(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsValueBool(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("true")
	tc := coretestcases.CaseV1{Name: "IsValueBool", Expected: false, Actual: sso.IsValueBool(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsSetter_True(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("yes")
	tc := coretestcases.CaseV1{Name: "IsSetter true", Expected: issetter.True, Actual: sso.IsSetter(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsSetter_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("yes")
	tc := coretestcases.CaseV1{Name: "IsSetter uninit", Expected: issetter.False, Actual: sso.IsSetter(true), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsSetter_ParseErr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("notbool")
	tc := coretestcases.CaseV1{Name: "IsSetter parseErr", Expected: issetter.Uninitialized, Actual: sso.IsSetter(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsSetter_ParseTrue(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("true")
	tc := coretestcases.CaseV1{Name: "IsSetter parse true", Expected: issetter.True, Actual: sso.IsSetter(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_IsSetter_ParseFalse(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("false")
	tc := coretestcases.CaseV1{Name: "IsSetter parse false", Expected: issetter.False, Actual: sso.IsSetter(false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Clone / String / Dispose / Json
// ═══════════════════════════════════════════════════════════════

func Test_Cov44_SSO_CloneUsingNewVal(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("old")
	cloned := sso.CloneUsingNewVal("new")
	tc := coretestcases.CaseV1{Name: "CloneUsingNewVal value", Expected: "new", Actual: cloned.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "CloneUsingNewVal isInit", Expected: true, Actual: cloned.IsInitialized(), Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Clone(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("val")
	cloned := sso.Clone()
	tc := coretestcases.CaseV1{Name: "Clone", Expected: "val", Actual: cloned.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_ClonePtr_Valid(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("val")
	cloned := sso.ClonePtr()
	tc := coretestcases.CaseV1{Name: "ClonePtr", Expected: "val", Actual: cloned.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_ClonePtr_Nil(t *testing.T) {
	var sso *corestr.SimpleStringOnce
	cloned := sso.ClonePtr()
	tc := coretestcases.CaseV1{Name: "ClonePtr nil", Expected: true, Actual: cloned == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_StringPtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("hello")
	ptr := sso.StringPtr()
	tc := coretestcases.CaseV1{Name: "StringPtr", Expected: "hello", Actual: *ptr, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_StringPtr_Nil(t *testing.T) {
	var sso *corestr.SimpleStringOnce
	ptr := sso.StringPtr()
	tc := coretestcases.CaseV1{Name: "StringPtr nil", Expected: "", Actual: *ptr, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Dispose(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("val")
	sso.Dispose()
	tc := coretestcases.CaseV1{Name: "Dispose value empty", Expected: "", Actual: sso.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Dispose_Nil(t *testing.T) {
	var sso *corestr.SimpleStringOnce
	sso.Dispose() // should not panic
	tc := coretestcases.CaseV1{Name: "Dispose nil no panic", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_NonPtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("val")
	np := sso.NonPtr()
	tc := coretestcases.CaseV1{Name: "NonPtr", Expected: "val", Actual: np.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Ptr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("val")
	p := sso.Ptr()
	tc := coretestcases.CaseV1{Name: "Ptr not nil", Expected: true, Actual: p != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SafeValue_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("test")
	tc := coretestcases.CaseV1{Name: "SafeValue init", Expected: "test", Actual: sso.SafeValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_SafeValue_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("test")
	tc := coretestcases.CaseV1{Name: "SafeValue uninit", Expected: "", Actual: sso.SafeValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Json(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("jsonval")
	jsonResult := sso.Json()
	tc := coretestcases.CaseV1{Name: "Json not empty", Expected: true, Actual: jsonResult.HasSafeNonEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_JsonPtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("jsonval")
	ptr := sso.JsonPtr()
	tc := coretestcases.CaseV1{Name: "JsonPtr not nil", Expected: true, Actual: ptr != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_MarshalUnmarshal(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("marshal")
	data, err := sso.MarshalJSON()
	tc := coretestcases.CaseV1{Name: "MarshalJSON no err", Expected: true, Actual: err == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)

	sso2 := corestr.New.SimpleStringOnce.Empty()
	err2 := sso2.UnmarshalJSON(data)
	tc2 := coretestcases.CaseV1{Name: "UnmarshalJSON no err", Expected: true, Actual: err2 == nil, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
	tc3 := coretestcases.CaseV1{Name: "UnmarshalJSON value", Expected: "marshal", Actual: sso2.Value(), Args: args.Map{}}
	tc3.ShouldBeEqual(t)
}

func Test_Cov44_SSO_Serialize(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("serval")
	data, err := sso.Serialize()
	tc := coretestcases.CaseV1{Name: "Serialize no err", Expected: true, Actual: err == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Serialize has data", Expected: true, Actual: len(data) > 0, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_SSO_AsJsonContractsBinder(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("v")
	binder := sso.AsJsonContractsBinder()
	tc := coretestcases.CaseV1{Name: "AsJsonContractsBinder", Expected: true, Actual: binder != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_AsJsoner(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("v")
	j := sso.AsJsoner()
	tc := coretestcases.CaseV1{Name: "AsJsoner", Expected: true, Actual: j != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_AsJsonParseSelfInjector(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("v")
	inj := sso.AsJsonParseSelfInjector()
	tc := coretestcases.CaseV1{Name: "AsJsonParseSelfInjector", Expected: true, Actual: inj != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_SSO_AsJsonMarshaller(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("v")
	m := sso.AsJsonMarshaller()
	tc := coretestcases.CaseV1{Name: "AsJsonMarshaller", Expected: true, Actual: m != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// ValidValue — Factory functions
// ═══════════════════════════════════════════════════════════════

func Test_Cov44_NewValidValueUsingAny_WithFieldName(t *testing.T) {
	vv := corestr.NewValidValueUsingAny(true, true, "hello")
	tc := coretestcases.CaseV1{Name: "NewValidValueUsingAny with field name", Expected: true, Actual: vv.IsValid, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "NewValidValueUsingAny value not empty", Expected: true, Actual: len(vv.Value) > 0, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_NewValidValueUsingAny_WithoutFieldName(t *testing.T) {
	vv := corestr.NewValidValueUsingAny(false, false, 42)
	tc := coretestcases.CaseV1{Name: "NewValidValueUsingAny without field name", Expected: false, Actual: vv.IsValid, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_NewValidValueUsingAnyAutoValid_NonEmpty(t *testing.T) {
	vv := corestr.NewValidValueUsingAnyAutoValid(false, "hello")
	tc := coretestcases.CaseV1{Name: "NewValidValueUsingAnyAutoValid non-empty", Expected: true, Actual: len(vv.Value) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_NewValidValueUsingAnyAutoValid_Empty(t *testing.T) {
	vv := corestr.NewValidValueUsingAnyAutoValid(false, "")
	tc := coretestcases.CaseV1{Name: "NewValidValueUsingAnyAutoValid empty", Expected: true, Actual: vv.IsValid, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_NewValidValue(t *testing.T) {
	vv := corestr.NewValidValue("test")
	tc := coretestcases.CaseV1{Name: "NewValidValue", Expected: "test", Actual: vv.Value, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "NewValidValue isValid", Expected: true, Actual: vv.IsValid, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_NewValidValueEmpty(t *testing.T) {
	vv := corestr.NewValidValueEmpty()
	tc := coretestcases.CaseV1{Name: "NewValidValueEmpty", Expected: "", Actual: vv.Value, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "NewValidValueEmpty isValid", Expected: true, Actual: vv.IsValid, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov44_InvalidValidValueNoMessage(t *testing.T) {
	vv := corestr.InvalidValidValueNoMessage()
	tc := coretestcases.CaseV1{Name: "InvalidValidValueNoMessage", Expected: false, Actual: vv.IsValid, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_ValidValue_ValueBytesOncePtr(t *testing.T) {
	vv := corestr.NewValidValue("bytes")
	result := vv.ValueBytesOncePtr()
	tc := coretestcases.CaseV1{Name: "ValueBytesOncePtr", Expected: 5, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// ValidValues — factory
// ═══════════════════════════════════════════════════════════════

func Test_Cov44_NewValidValuesUsingValues_Valid(t *testing.T) {
	v1 := corestr.ValidValue{Value: "a", IsValid: true}
	v2 := corestr.ValidValue{Value: "b", IsValid: true}
	vv := corestr.NewValidValuesUsingValues(v1, v2)
	tc := coretestcases.CaseV1{Name: "NewValidValuesUsingValues", Expected: 2, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_NewValidValuesUsingValues_Empty(t *testing.T) {
	vv := corestr.NewValidValuesUsingValues()
	tc := coretestcases.CaseV1{Name: "NewValidValuesUsingValues empty", Expected: 0, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_ValidValues_Count(t *testing.T) {
	vv := corestr.NewValidValues(5)
	tc := coretestcases.CaseV1{Name: "Count", Expected: 0, Actual: vv.Count(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_ValidValues_HasAnyItem(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := coretestcases.CaseV1{Name: "HasAnyItem empty", Expected: false, Actual: vv.HasAnyItem(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_ValidValues_LastIndex(t *testing.T) {
	v1 := corestr.ValidValue{Value: "a", IsValid: true}
	vv := corestr.NewValidValuesUsingValues(v1)
	tc := coretestcases.CaseV1{Name: "LastIndex", Expected: 0, Actual: vv.LastIndex(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_ValidValues_HasIndex_Valid(t *testing.T) {
	v1 := corestr.ValidValue{Value: "a", IsValid: true}
	vv := corestr.NewValidValuesUsingValues(v1)
	tc := coretestcases.CaseV1{Name: "HasIndex valid", Expected: true, Actual: vv.HasIndex(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov44_ValidValues_HasIndex_Invalid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := coretestcases.CaseV1{Name: "HasIndex invalid", Expected: false, Actual: vv.HasIndex(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
