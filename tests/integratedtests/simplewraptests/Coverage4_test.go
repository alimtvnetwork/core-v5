package simplewraptests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── Basic wrap functions ──

func Test_Cov4_With(t *testing.T) {
	actual := args.Map{"result": simplewrap.With("[", "hello", "]")}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "With", actual)
}

func Test_Cov4_WithPtr_AllPresent(t *testing.T) {
	s, src, e := "[", "hello", "]"
	result := simplewrap.WithPtr(&s, &src, &e)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithPtr all present", actual)
}

func Test_Cov4_WithStartEnd(t *testing.T) {
	actual := args.Map{"result": simplewrap.WithStartEnd("|", "hello")}
	expected := args.Map{"result": "|hello|"}
	expected.ShouldBeEqual(t, 0, "WithStartEnd", actual)
}

func Test_Cov4_WithStartEndPtr(t *testing.T) {
	w, src := "|", "hello"
	result := simplewrap.WithStartEndPtr(&w, &src)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "|hello|"}
	expected.ShouldBeEqual(t, 0, "WithStartEndPtr", actual)
}

func Test_Cov4_WithDoubleQuote(t *testing.T) {
	result := simplewrap.WithDoubleQuote("hello")
	actual := args.Map{"contains": strings.Contains(result, "hello")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote", actual)
}

func Test_Cov4_WithSingleQuote(t *testing.T) {
	result := simplewrap.WithSingleQuote("hello")
	actual := args.Map{"contains": strings.Contains(result, "hello")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote", actual)
}

func Test_Cov4_WithDoubleQuoteAny(t *testing.T) {
	result := simplewrap.WithDoubleQuoteAny(42)
	actual := args.Map{"contains": strings.Contains(result, "42")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny", actual)
}

func Test_Cov4_ToJsonName(t *testing.T) {
	result := simplewrap.ToJsonName("test")
	actual := args.Map{"contains": strings.Contains(result, "test")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName", actual)
}

// ── Bracket/Curly/Parenthesis wraps ──

func Test_Cov4_WithBrackets(t *testing.T) {
	result := simplewrap.WithBrackets("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithBrackets", actual)
}

func Test_Cov4_WithCurly(t *testing.T) {
	result := simplewrap.WithCurly("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "WithCurly", actual)
}

func Test_Cov4_WithParenthesis(t *testing.T) {
	result := simplewrap.WithParenthesis("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "WithParenthesis", actual)
}

func Test_Cov4_CurlyWrap(t *testing.T) {
	result := simplewrap.CurlyWrap("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrap", actual)
}

func Test_Cov4_SquareWrap(t *testing.T) {
	result := simplewrap.SquareWrap("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "SquareWrap", actual)
}

func Test_Cov4_ParenthesisWrap(t *testing.T) {
	result := simplewrap.ParenthesisWrap("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap", actual)
}

// ── If variants ──

func Test_Cov4_CurlyWrapIf_True(t *testing.T) {
	actual := args.Map{"result": simplewrap.CurlyWrapIf(true, "x")}
	expected := args.Map{"result": "{x}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf true", actual)
}

func Test_Cov4_CurlyWrapIf_False(t *testing.T) {
	actual := args.Map{"result": simplewrap.CurlyWrapIf(false, "x")}
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf false", actual)
}

func Test_Cov4_SquareWrapIf_True(t *testing.T) {
	actual := args.Map{"result": simplewrap.SquareWrapIf(true, "x")}
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "SquareWrapIf true", actual)
}

func Test_Cov4_SquareWrapIf_False(t *testing.T) {
	actual := args.Map{"result": simplewrap.SquareWrapIf(false, "x")}
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "SquareWrapIf false", actual)
}

func Test_Cov4_ParenthesisWrapIf_True(t *testing.T) {
	actual := args.Map{"result": simplewrap.ParenthesisWrapIf(true, "x")}
	expected := args.Map{"result": "(x)"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrapIf true", actual)
}

func Test_Cov4_ParenthesisWrapIf_False(t *testing.T) {
	actual := args.Map{"result": simplewrap.ParenthesisWrapIf(false, "x")}
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrapIf false", actual)
}

// ── Title wraps ──

func Test_Cov4_TitleCurlyWrap(t *testing.T) {
	result := simplewrap.TitleCurlyWrap("title", "value")
	actual := args.Map{"containsTitle": strings.Contains(result, "title"), "containsVal": strings.Contains(result, "value")}
	expected := args.Map{"containsTitle": true, "containsVal": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyWrap", actual)
}

func Test_Cov4_TitleSquare(t *testing.T) {
	result := simplewrap.TitleSquare("title", "value")
	actual := args.Map{"containsTitle": strings.Contains(result, "title"), "containsVal": strings.Contains(result, "value")}
	expected := args.Map{"containsTitle": true, "containsVal": true}
	expected.ShouldBeEqual(t, 0, "TitleSquare", actual)
}

func Test_Cov4_TitleCurlyMeta(t *testing.T) {
	result := simplewrap.TitleCurlyMeta("title", "value", "meta")
	actual := args.Map{"containsMeta": strings.Contains(result, "meta")}
	expected := args.Map{"containsMeta": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyMeta", actual)
}

func Test_Cov4_TitleSquareMeta(t *testing.T) {
	result := simplewrap.TitleSquareMeta("title", "value", "meta")
	actual := args.Map{"containsMeta": strings.Contains(result, "meta")}
	expected := args.Map{"containsMeta": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMeta", actual)
}

func Test_Cov4_TitleSquareCsvMeta(t *testing.T) {
	result := simplewrap.TitleSquareCsvMeta("title", "value", "m1", "m2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareCsvMeta", actual)
}

type cov4Stringer struct{ val string }
func (s cov4Stringer) String() string { return s.val }

func Test_Cov4_TitleSquareMetaUsingFmt(t *testing.T) {
	result := simplewrap.TitleSquareMetaUsingFmt(
		cov4Stringer{"title"},
		cov4Stringer{"value"},
		cov4Stringer{"meta"},
	)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMetaUsingFmt", actual)
}

func Test_Cov4_TitleQuotationMeta(t *testing.T) {
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")
	actual := args.Map{"containsQuote": strings.Contains(result, `"`)}
	expected := args.Map{"containsQuote": true}
	expected.ShouldBeEqual(t, 0, "TitleQuotationMeta", actual)
}

// ── MsgWrapNumber ──

func Test_Cov4_MsgWrapNumber(t *testing.T) {
	result := simplewrap.MsgWrapNumber("count", 42)
	actual := args.Map{"contains": strings.Contains(result, "42")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber", actual)
}

// ── MsgCsvItems ──

func Test_Cov4_MsgCsvItems(t *testing.T) {
	result := simplewrap.MsgCsvItems("msg", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems", actual)
}

// ── DoubleQuoteWrapElements ──

func Test_Cov4_DoubleQuoteWrapElements_Normal(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")
	actual := args.Map{"len": len(result), "firstQuoted": strings.HasPrefix(result[0], fmt.Sprintf("%c", '"'))}
	expected := args.Map{"len": 2, "firstQuoted": true}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements normal", actual)
}

func Test_Cov4_DoubleQuoteWrapElements_Empty(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements empty", actual)
}

// ── DoubleQuoteWrapElementsWithIndexes ──

func Test_Cov4_DoubleQuoteWrapElementsWithIndexes_Normal(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")
	actual := args.Map{"len": len(result), "containsIdx": strings.Contains(result[0], "[0]")}
	expected := args.Map{"len": 2, "containsIdx": true}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes normal", actual)
}
