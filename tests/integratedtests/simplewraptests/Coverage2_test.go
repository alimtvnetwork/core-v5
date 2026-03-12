package simplewraptests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
)

// ── TitleCurlyMeta / TitleSquareMeta / TitleQuotationMeta ──

func Test_TitleCurlyMeta_Cov2(t *testing.T) {
	r := simplewrap.TitleCurlyMeta("title", "val", "meta")
	if r == "" || !strings.Contains(r, "title") {
		t.Error("should contain title")
	}
}

func Test_TitleSquareMeta_Cov2(t *testing.T) {
	r := simplewrap.TitleSquareMeta("title", "val", "meta")
	if r == "" || !strings.Contains(r, "title") {
		t.Error("should contain title")
	}
}

func Test_TitleQuotationMeta_Cov2(t *testing.T) {
	r := simplewrap.TitleQuotationMeta("title", "val", "meta")
	if r == "" || !strings.Contains(r, "title") {
		t.Error("should contain title")
	}
}

// ── TitleSquareCsvMeta ──

func Test_TitleSquareCsvMeta_Cov2(t *testing.T) {
	r := simplewrap.TitleSquareCsvMeta("title", "val", "a", "b")
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── TitleSquareMetaUsingFmt ──

type testStringer struct{}

func (s testStringer) String() string { return "stringer" }

func Test_TitleSquareMetaUsingFmt_Cov2(t *testing.T) {
	r := simplewrap.TitleSquareMetaUsingFmt(
		testStringer{},
		testStringer{},
		testStringer{},
	)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── WithBracketsQuotation / WithCurlyQuotation / WithParenthesisQuotation ──

func Test_WithBracketsQuotation_Cov2(t *testing.T) {
	r := simplewrap.WithBracketsQuotation("hello")
	if r == "" {
		t.Error("should not be empty")
	}
}

func Test_WithCurlyQuotation_Cov2(t *testing.T) {
	r := simplewrap.WithCurlyQuotation("hello")
	if r == "" {
		t.Error("should not be empty")
	}
}

func Test_WithParenthesisQuotation_Cov2(t *testing.T) {
	r := simplewrap.WithParenthesisQuotation("hello")
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── CurlyWrapOption ──

func Test_CurlyWrapOption_SkipIfExists_Cov2(t *testing.T) {
	r := simplewrap.CurlyWrapOption(true, "{hello}")
	if r != "{hello}" {
		t.Errorf("expected {hello}, got %s", r)
	}
}

func Test_CurlyWrapOption_NoSkip_Cov2(t *testing.T) {
	r := simplewrap.CurlyWrapOption(false, "hello")
	if r != "{hello}" {
		t.Errorf("expected {hello}, got %s", r)
	}
}

func Test_CurlyWrapOption_SkipNotPresent_Cov2(t *testing.T) {
	r := simplewrap.CurlyWrapOption(true, "hello")
	if r != "{hello}" {
		t.Errorf("expected {hello}, got %s", r)
	}
}

// ── DoubleQuoteWrapElements nil input ──

func Test_DoubleQuoteWrapElements_Nil_Cov2(t *testing.T) {
	r := simplewrap.DoubleQuoteWrapElements(false, nil...)
	if r == nil {
		t.Error("nil input should return empty slice")
	}
}

func Test_DoubleQuoteWrapElements_EmptySlice_Cov2(t *testing.T) {
	r := simplewrap.DoubleQuoteWrapElements(false)
	if len(r) != 0 {
		t.Error("empty input should return empty")
	}
}

func Test_DoubleQuoteWrapElements_SkipExistence_Cov2(t *testing.T) {
	r := simplewrap.DoubleQuoteWrapElements(true, `"already"`, "naked")
	if len(r) != 2 {
		t.Error("should have 2")
	}
}

// ── DoubleQuoteWrapElementsWithIndexes nil ──

func Test_DoubleQuoteWrapElementsWithIndexes_Nil_Cov2(t *testing.T) {
	r := simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...)
	if r == nil {
		t.Error("nil input should return empty slice")
	}
}

func Test_DoubleQuoteWrapElementsWithIndexes_Empty_Cov2(t *testing.T) {
	r := simplewrap.DoubleQuoteWrapElementsWithIndexes()
	if len(r) != 0 {
		t.Error("empty should return empty")
	}
}

func Test_DoubleQuoteWrapElementsWithIndexes_Items_Cov2(t *testing.T) {
	r := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")
	if len(r) != 2 {
		t.Error("should have 2")
	}
	if !strings.Contains(r[0], "[0]") {
		t.Error("should contain index")
	}
}

// ── WithDoubleQuote / WithDoubleQuoteAny ──

func Test_WithDoubleQuote_Empty_Cov2(t *testing.T) {
	r := simplewrap.WithDoubleQuote("")
	if r != `""` {
		t.Errorf("expected empty quotes, got %s", r)
	}
}

func Test_WithDoubleQuoteAny_Int_Cov2(t *testing.T) {
	r := simplewrap.WithDoubleQuoteAny(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── WithSingleQuote ──

func Test_WithSingleQuote_Empty_Cov2(t *testing.T) {
	r := simplewrap.WithSingleQuote("")
	if r != "''" {
		t.Errorf("expected '', got %s", r)
	}
}

// ── ToJsonName ──

func Test_ToJsonName_Int_Cov2(t *testing.T) {
	r := simplewrap.ToJsonName(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── WithCurly int ──

func Test_WithCurly_Int_Cov2(t *testing.T) {
	r := simplewrap.WithCurly(42)
	if !strings.Contains(r, "42") {
		t.Error("should contain 42")
	}
}

// ── WrapWith / WrapWithStartEnd extended ──

func Test_With_Empty_Cov2(t *testing.T) {
	r := simplewrap.With("", "", "")
	if r != "" {
		t.Error("all empty should be empty")
	}
}

func Test_WithStartEnd_Empty_Cov2(t *testing.T) {
	r := simplewrap.WithStartEnd("", "")
	if r != "" {
		t.Error("all empty should be empty")
	}
}

// ── MsgWrapNumber int64 ──

func Test_MsgWrapNumber_Int64_Cov2(t *testing.T) {
	r := simplewrap.MsgWrapNumber("total", int64(100))
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── MsgCsvItems empty ──

func Test_MsgCsvItems_Empty_Cov2(t *testing.T) {
	r := simplewrap.MsgCsvItems("msg")
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── ConditionalWrapWith extended ──

func Test_ConditionalWrapWith_BothPresent2Char_Cov2(t *testing.T) {
	r := simplewrap.ConditionalWrapWith('{', "{}", '}')
	if r != "{}" {
		t.Errorf("expected {}, got %s", r)
	}
}

// ── CurlyWrap int ──

func Test_CurlyWrap_Int_Cov2(t *testing.T) {
	r := simplewrap.CurlyWrap(42)
	if !strings.Contains(r, "42") {
		t.Error("should contain 42")
	}
}

// ── SquareWrap int ──

func Test_SquareWrap_Int_Cov2(t *testing.T) {
	r := simplewrap.SquareWrap(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── ParenthesisWrap int ──

func Test_ParenthesisWrap_Int_Cov2(t *testing.T) {
	r := simplewrap.ParenthesisWrap(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── TitleCurlyWrap int ──

func Test_TitleCurlyWrap_Int_Cov2(t *testing.T) {
	r := simplewrap.TitleCurlyWrap("t", 42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── TitleSquare int ──

func Test_TitleSquare_Int_Cov2(t *testing.T) {
	r := simplewrap.TitleSquare("t", 42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── WithBrackets int ──

func Test_WithBrackets_Int_Cov2(t *testing.T) {
	r := simplewrap.WithBrackets(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── WithParenthesis int ──

func Test_WithParenthesis_Int_Cov2(t *testing.T) {
	r := simplewrap.WithParenthesis(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

// ── toString via CurlyWrapIf with fmt.Stringer ──

func Test_CurlyWrapIf_Stringer_Cov2(t *testing.T) {
	r := simplewrap.CurlyWrapIf(true, testStringer{})
	if r == "" {
		t.Error("should not be empty")
	}
}

func Test_CurlyWrapIf_FmtStringer_Cov2(t *testing.T) {
	var s fmt.Stringer = testStringer{}
	r := simplewrap.CurlyWrapIf(true, s)
	if r == "" {
		t.Error("should not be empty")
	}
}
