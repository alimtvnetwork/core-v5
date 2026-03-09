package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// TextValidator.IsMatch — Equal
// ==========================================

func Test_TextValidator_IsMatch_ExactEqual(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if !v.IsMatch("hello", true) {
		t.Error("exact match should return true")
	}
}

func Test_TextValidator_IsMatch_ExactNotEqual(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if v.IsMatch("world", true) {
		t.Error("different text should not match")
	}
}

func Test_TextValidator_IsMatch_CaseInsensitive(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "Hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if !v.IsMatch("hello", false) {
		t.Error("case-insensitive should match")
	}
}

func Test_TextValidator_IsMatch_CaseSensitiveFail(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "Hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if v.IsMatch("hello", true) {
		t.Error("case-sensitive should not match different cases")
	}
}

// ==========================================
// TextValidator.IsMatch — with Trim
// ==========================================

func Test_TextValidator_IsMatch_TrimMatch(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "  hello  ",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}
	if !v.IsMatch("hello", true) {
		t.Error("trimmed search should match trimmed content")
	}
}

func Test_TextValidator_IsMatch_TrimBothSides(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "  hello  ",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}
	if !v.IsMatch("  hello  ", true) {
		t.Error("trim should handle both search and content")
	}
}

// ==========================================
// TextValidator.IsMatch — Contains
// ==========================================

func Test_TextValidator_IsMatch_Contains(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "ell",
		SearchAs:  stringcompareas.Contains,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if !v.IsMatch("hello world", true) {
		t.Error("contains should find substring")
	}
}

func Test_TextValidator_IsMatch_ContainsMissing(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "xyz",
		SearchAs:  stringcompareas.Contains,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if v.IsMatch("hello world", true) {
		t.Error("contains should not find missing substring")
	}
}

// ==========================================
// TextValidator.IsMatch — NotEqual
// ==========================================

func Test_TextValidator_IsMatch_NotEqual_Different(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.NotEqual,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if !v.IsMatch("world", true) {
		t.Error("NotEqual should match when different")
	}
}

func Test_TextValidator_IsMatch_NotEqual_Same(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.NotEqual,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if v.IsMatch("hello", true) {
		t.Error("NotEqual should not match when same")
	}
}

// ==========================================
// TextValidator.IsMatch — UniqueWords + Sort
// ==========================================

func Test_TextValidator_IsMatch_UniqueWordsSorted(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "  banana  apple  apple  cherry  ",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultUniqueWordsCoreCondition,
	}
	// unique words sorted: apple banana cherry
	if !v.IsMatch("cherry banana apple", false) {
		t.Error("unique+sorted should match reordered unique words")
	}
}

// ==========================================
// TextValidator.IsMatch — Empty strings
// ==========================================

func Test_TextValidator_IsMatch_EmptySearchEmptyContent(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if !v.IsMatch("", true) {
		t.Error("empty search vs empty content should match")
	}
}

func Test_TextValidator_IsMatch_EmptySearchNonEmptyContent(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if v.IsMatch("hello", true) {
		t.Error("empty search vs non-empty content should not match")
	}
}

// ==========================================
// TextValidator.IsMatchMany
// ==========================================

func Test_TextValidator_IsMatchMany_AllMatch(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if !v.IsMatchMany(false, true, "hello", "hello", "hello") {
		t.Error("all identical should match")
	}
}

func Test_TextValidator_IsMatchMany_OneFails(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if v.IsMatchMany(false, true, "hello", "world", "hello") {
		t.Error("one mismatch should fail")
	}
}

func Test_TextValidator_IsMatchMany_EmptySkip(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	if !v.IsMatchMany(true, true) {
		t.Error("empty contents with skip should return true")
	}
}

// (nil receiver test migrated to TextValidator_NilReceiver_testcases.go)

// ==========================================
// TextValidator.VerifyDetailError
// ==========================================

func Test_TextValidator_VerifyDetailError_Match(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyDetailError(params, "hello")
	if err != nil {
		t.Errorf("match should not error: %v", err)
	}
}

func Test_TextValidator_VerifyDetailError_Mismatch(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyDetailError(params, "world")
	if err == nil {
		t.Error("mismatch should return error")
	}
}

// (nil receiver test migrated to TextValidator_NilReceiver_testcases.go)

// ==========================================
// TextValidator.VerifyMany
// ==========================================

func Test_TextValidator_VerifyMany_FirstOnly(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyMany(false, params, "a", "b", "c")
	if err == nil {
		t.Error("should return first error")
	}
}

func Test_TextValidator_VerifyMany_AllErrors(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}
	err := v.VerifyMany(true, params, "a", "b")
	if err == nil {
		t.Error("should return all errors")
	}
}

func Test_TextValidator_VerifyFirstError_EmptySkip(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := v.VerifyFirstError(params)
	if err != nil {
		t.Errorf("empty contents with skip should not error: %v", err)
	}
}

// ==========================================
// TextValidator.SearchTextFinalized — caching
// ==========================================

func Test_TextValidator_SearchTextFinalized_Cached(t *testing.T) {
	v := corevalidator.TextValidator{
		Search:    "  hello  ",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}
	first := v.SearchTextFinalized()
	second := v.SearchTextFinalized()
	if first != second {
		t.Error("cached result should be identical")
	}
	if first != "hello" {
		t.Errorf("expected 'hello', got '%s'", first)
	}
}

// ==========================================
// EmptyValidator preset
// ==========================================

func Test_EmptyValidator_MatchesEmpty(t *testing.T) {
	v := corevalidator.EmptyValidator
	if !v.IsMatch("", true) {
		t.Error("EmptyValidator should match empty string")
	}
}

func Test_EmptyValidator_MatchesTrimmedEmpty(t *testing.T) {
	v := corevalidator.EmptyValidator
	if !v.IsMatch("   ", true) {
		t.Error("EmptyValidator with trim should match whitespace-only")
	}
}

func Test_EmptyValidator_NoMatchNonEmpty(t *testing.T) {
	v := corevalidator.EmptyValidator
	if v.IsMatch("hello", true) {
		t.Error("EmptyValidator should not match non-empty string")
	}
}
