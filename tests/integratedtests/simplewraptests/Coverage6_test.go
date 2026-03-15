package simplewraptests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── TitleCurlyWrap ──

func Test_Cov6_TitleCurlyWrap(t *testing.T) {
	result := simplewrap.TitleCurlyWrap("title", "value")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyWrap", actual)
}

// ── TitleSquare ──

func Test_Cov6_TitleSquare(t *testing.T) {
	result := simplewrap.TitleSquare("title", "value")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquare", actual)
}

// ── TitleSquareMeta ──

func Test_Cov6_TitleSquareMeta(t *testing.T) {
	result := simplewrap.TitleSquareMeta("title", "value", "meta")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMeta", actual)
}

// ── TitleSquareMetaUsingFmt ──

type cov6Stringer struct{ val string }

func (s cov6Stringer) String() string { return s.val }

func Test_Cov6_TitleSquareMetaUsingFmt(t *testing.T) {
	result := simplewrap.TitleSquareMetaUsingFmt(
		cov6Stringer{"t"}, cov6Stringer{"v"}, cov6Stringer{"m"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMetaUsingFmt", actual)
}

// ── TitleSquareCsvMeta ──

func Test_Cov6_TitleSquareCsvMeta(t *testing.T) {
	result := simplewrap.TitleSquareCsvMeta("title", "value", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareCsvMeta", actual)
}

// ── ToJsonName ──

func Test_Cov6_ToJsonName(t *testing.T) {
	result := simplewrap.ToJsonName("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName", actual)
}

// ── MsgWrapNumber ──

func Test_Cov6_MsgWrapNumber(t *testing.T) {
	result := simplewrap.MsgWrapNumber("count", 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber", actual)
}

// ── With / WithPtr ──

func Test_Cov6_With(t *testing.T) {
	result := simplewrap.With("[", "hello", "]")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "With", actual)
}

func Test_Cov6_WithPtr(t *testing.T) {
	s, e, v := "[", "]", "hello"
	result := simplewrap.WithPtr(&s, &v, &e)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithPtr", actual)
}

func Test_Cov6_WithPtr_Nils(t *testing.T) {
	v := "hello"
	result := simplewrap.WithPtr(nil, &v, nil)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "WithPtr nils", actual)
}

func Test_Cov6_WithPtr_NilSource(t *testing.T) {
	s, e := "[", "]"
	result := simplewrap.WithPtr(&s, nil, &e)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "WithPtr nil source", actual)
}

// ── WithStartEnd / WithStartEndPtr ──

func Test_Cov6_WithStartEnd(t *testing.T) {
	result := simplewrap.WithStartEnd("'", "hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "'hello'"}
	expected.ShouldBeEqual(t, 0, "WithStartEnd", actual)
}

func Test_Cov6_WithStartEndPtr(t *testing.T) {
	w, v := "'", "hello"
	result := simplewrap.WithStartEndPtr(&w, &v)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "'hello'"}
	expected.ShouldBeEqual(t, 0, "WithStartEndPtr", actual)
}

// ── WithDoubleQuoteAny ──

func Test_Cov6_WithDoubleQuoteAny(t *testing.T) {
	result := simplewrap.WithDoubleQuoteAny(42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny", actual)
}

// ── WithSingleQuote ──

func Test_Cov6_WithSingleQuote(t *testing.T) {
	result := simplewrap.WithSingleQuote("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote", actual)
}

// ── WithDoubleQuote ──

func Test_Cov6_WithDoubleQuote(t *testing.T) {
	result := simplewrap.WithDoubleQuote("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote", actual)
}

// ── WithBrackets ──

func Test_Cov6_WithBrackets(t *testing.T) {
	result := simplewrap.WithBrackets("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBrackets", actual)
}

// ── WithCurly ──

func Test_Cov6_WithCurly(t *testing.T) {
	result := simplewrap.WithCurly("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithCurly", actual)
}

// ── WithParenthesis ──

func Test_Cov6_WithParenthesis(t *testing.T) {
	result := simplewrap.WithParenthesis("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesis", actual)
}

// ── CurlyWrap ──

func Test_Cov6_CurlyWrap(t *testing.T) {
	result := simplewrap.CurlyWrap("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrap", actual)
}

// ── CurlyWrapIf ──

func Test_Cov6_CurlyWrapIf(t *testing.T) {
	result := simplewrap.CurlyWrapIf(true, "hello")
	noWrap := simplewrap.CurlyWrapIf(false, "hello")
	actual := args.Map{"wrapped": result != "", "noWrap": noWrap != ""}
	expected := args.Map{"wrapped": true, "noWrap": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf", actual)
}

// ── ParenthesisWrap ──

func Test_Cov6_ParenthesisWrap(t *testing.T) {
	result := simplewrap.ParenthesisWrap("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap", actual)
}

// ── ParenthesisWrapIf ──

func Test_Cov6_ParenthesisWrapIf(t *testing.T) {
	result := simplewrap.ParenthesisWrapIf(true, "hello")
	noWrap := simplewrap.ParenthesisWrapIf(false, "hello")
	actual := args.Map{"wrapped": result != "", "noWrap": noWrap != ""}
	expected := args.Map{"wrapped": true, "noWrap": true}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrapIf", actual)
}

// ── SquareWrap ──

func Test_Cov6_SquareWrap(t *testing.T) {
	result := simplewrap.SquareWrap("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SquareWrap", actual)
}

// ── SquareWrapIf ──

func Test_Cov6_SquareWrapIf(t *testing.T) {
	result := simplewrap.SquareWrapIf(true, "hello")
	noWrap := simplewrap.SquareWrapIf(false, "hello")
	actual := args.Map{"wrapped": result != "", "noWrap": noWrap != ""}
	expected := args.Map{"wrapped": true, "noWrap": true}
	expected.ShouldBeEqual(t, 0, "SquareWrapIf", actual)
}

// ── TitleCurlyMeta ──

func Test_Cov6_TitleCurlyMeta(t *testing.T) {
	result := simplewrap.TitleCurlyMeta("title", "value", "meta")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyMeta", actual)
}

// ── TitleQuotationMeta ──

func Test_Cov6_TitleQuotationMeta(t *testing.T) {
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleQuotationMeta", actual)
}

// ── DoubleQuoteWrapElements — skip on existence ──

func Test_Cov6_DoubleQuoteWrapElements_SkipOnExistence(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(true, `"hello"`, "world")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements skip on existence", actual)
}

func Test_Cov6_DoubleQuoteWrapElements_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements nil", actual)
}

func Test_Cov6_DoubleQuoteWrapElements_Empty(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements empty", actual)
}

// ── DoubleQuoteWrapElementsWithIndexes ──

func Test_Cov6_DoubleQuoteWrapElementsWithIndexes_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes nil", actual)
}

func Test_Cov6_DoubleQuoteWrapElementsWithIndexes_Empty(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes empty", actual)
}

func Test_Cov6_DoubleQuoteWrapElementsWithIndexes_Items(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")
	actual := args.Map{"len": len(result), "notEmpty": result[0] != ""}
	expected := args.Map{"len": 2, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes items", actual)
}

// ── ConditionalWrapWith — missing left/right ──

func Test_Cov6_ConditionalWrapWith_MissingLeft(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "x]", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith missing left", actual)
}

func Test_Cov6_ConditionalWrapWith_MissingRight(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "[x", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith missing right", actual)
}

func Test_Cov6_ConditionalWrapWith_SingleCharPresent(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "[", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": fmt.Sprintf("[%c", ']')}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith single char present", actual)
}

// ── MsgCsvItems with items ──

func Test_Cov6_MsgCsvItems_WithItems(t *testing.T) {
	result := simplewrap.MsgCsvItems("msg", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems with items", actual)
}
