package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// TextValidator.IsMatch — Equal
// ==========================================

func Test_TextValidator_IsMatch_ExactEqual(t *testing.T) {
	tc := tvIsMatchExactEqualTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_ExactNotEqual(t *testing.T) {
	tc := tvIsMatchExactNotEqualTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("world", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_CaseInsensitive(t *testing.T) {
	tc := tvIsMatchCaseInsensitiveTestCase
	v := corevalidator.TextValidator{
		Search: "Hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello", false)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_CaseSensitiveFail(t *testing.T) {
	tc := tvIsMatchCaseSensitiveFailTestCase
	v := corevalidator.TextValidator{
		Search: "Hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — with Trim
// ==========================================

func Test_TextValidator_IsMatch_TrimMatch(t *testing.T) {
	tc := tvIsMatchTrimTestCase
	v := corevalidator.TextValidator{
		Search: "  hello  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_TrimBothSides(t *testing.T) {
	tc := tvIsMatchTrimBothTestCase
	v := corevalidator.TextValidator{
		Search: "  hello  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("  hello  ", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — Contains
// ==========================================

func Test_TextValidator_IsMatch_Contains(t *testing.T) {
	tc := tvIsMatchContainsTestCase
	v := corevalidator.TextValidator{
		Search: "ell", SearchAs: stringcompareas.Contains,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_ContainsMissing(t *testing.T) {
	tc := tvIsMatchContainsMissingTestCase
	v := corevalidator.TextValidator{
		Search: "xyz", SearchAs: stringcompareas.Contains,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — NotEqual
// ==========================================

func Test_TextValidator_IsMatch_NotEqual_Different(t *testing.T) {
	tc := tvIsMatchNotEqualDifferentTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.NotEqual,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("world", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_NotEqual_Same(t *testing.T) {
	tc := tvIsMatchNotEqualSameTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.NotEqual,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — UniqueWords + Sort
// ==========================================

func Test_TextValidator_IsMatch_UniqueWordsSorted(t *testing.T) {
	tc := tvIsMatchUniqueWordsSortedTestCase
	v := corevalidator.TextValidator{
		Search: "  banana  apple  apple  cherry  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultUniqueWordsCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("cherry banana apple", false)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — Empty strings
// ==========================================

func Test_TextValidator_IsMatch_EmptySearchEmptyContent(t *testing.T) {
	tc := tvIsMatchEmptyBothTestCase
	v := corevalidator.TextValidator{
		Search: "", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_EmptySearchNonEmptyContent(t *testing.T) {
	tc := tvIsMatchEmptySearchNonEmptyTestCase
	v := corevalidator.TextValidator{
		Search: "", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatchMany
// ==========================================

func Test_TextValidator_IsMatchMany_AllMatch(t *testing.T) {
	tc := tvIsMatchManyAllTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatchMany(false, true, "hello", "hello", "hello")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatchMany_OneFails(t *testing.T) {
	tc := tvIsMatchManyOneFailsTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatchMany(false, true, "hello", "world", "hello")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatchMany_EmptySkip(t *testing.T) {
	tc := tvIsMatchManyEmptySkipTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	actual := args.Map{"isMatch": v.IsMatchMany(true, true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to TextValidator_NilReceiver_testcases.go)

// ==========================================
// TextValidator.VerifyDetailError
// ==========================================

func Test_TextValidator_VerifyDetailError_Match(t *testing.T) {
	tc := tvVerifyDetailMatchTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	actual := args.Map{"hasError": v.VerifyDetailError(params, "hello") != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_VerifyDetailError_Mismatch(t *testing.T) {
	tc := tvVerifyDetailMismatchTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	actual := args.Map{"hasError": v.VerifyDetailError(params, "world") != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to TextValidator_NilReceiver_testcases.go)

// ==========================================
// TextValidator.VerifyMany
// ==========================================

func Test_TextValidator_VerifyMany_FirstOnly(t *testing.T) {
	tc := tvVerifyManyFirstOnlyTestCase
	v := corevalidator.TextValidator{
		Search: "x", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	actual := args.Map{"hasError": v.VerifyMany(false, params, "a", "b", "c") != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_VerifyMany_AllErrors(t *testing.T) {
	tc := tvVerifyManyAllErrorsTestCase
	v := corevalidator.TextValidator{
		Search: "x", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	actual := args.Map{"hasError": v.VerifyMany(true, params, "a", "b") != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_VerifyFirstError_EmptySkip(t *testing.T) {
	tc := tvVerifyFirstEmptySkipTestCase
	v := corevalidator.TextValidator{
		Search: "x", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}

	actual := args.Map{"hasError": v.VerifyFirstError(params) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.SearchTextFinalized — caching
// ==========================================

func Test_TextValidator_SearchTextFinalized_Cached(t *testing.T) {
	tc := tvSearchTextFinalizedTestCase
	v := corevalidator.TextValidator{
		Search: "  hello  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}
	first := v.SearchTextFinalized()
	second := v.SearchTextFinalized()

	actual := args.Map{
		"isCached": first == second,
		"value":    first,
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// EmptyValidator preset
// ==========================================

func Test_EmptyValidator_MatchesEmpty(t *testing.T) {
	tc := tvEmptyMatchesEmptyTestCase

	actual := args.Map{"isMatch": corevalidator.EmptyValidator.IsMatch("", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyValidator_MatchesTrimmedEmpty(t *testing.T) {
	tc := tvEmptyMatchesTrimmedTestCase

	actual := args.Map{"isMatch": corevalidator.EmptyValidator.IsMatch("   ", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyValidator_NoMatchNonEmpty(t *testing.T) {
	tc := tvEmptyNoMatchNonEmptyTestCase

	actual := args.Map{"isMatch": corevalidator.EmptyValidator.IsMatch("hello", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}
