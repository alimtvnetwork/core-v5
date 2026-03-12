package simplewraptests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

type testStringer struct{}

func (s testStringer) String() string { return "stringer" }

func Test_Cov2_TitleCurlyMeta(t *testing.T) {
	r := simplewrap.TitleCurlyMeta("title", "val", "meta")
	actual := args.Map{"notEmpty": r != "", "containsTitle": strings.Contains(r, "title")}
	expected := args.Map{"notEmpty": true, "containsTitle": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyMeta", actual)
}

func Test_Cov2_TitleSquareMeta(t *testing.T) {
	r := simplewrap.TitleSquareMeta("title", "val", "meta")
	actual := args.Map{"notEmpty": r != "", "containsTitle": strings.Contains(r, "title")}
	expected := args.Map{"notEmpty": true, "containsTitle": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMeta", actual)
}

func Test_Cov2_TitleQuotationMeta(t *testing.T) {
	r := simplewrap.TitleQuotationMeta("title", "val", "meta")
	actual := args.Map{"notEmpty": r != "", "containsTitle": strings.Contains(r, "title")}
	expected := args.Map{"notEmpty": true, "containsTitle": true}
	expected.ShouldBeEqual(t, 0, "TitleQuotationMeta", actual)
}

func Test_Cov2_TitleSquareCsvMeta(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.TitleSquareCsvMeta("title", "val", "a", "b") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareCsvMeta", actual)
}

func Test_Cov2_TitleSquareMetaUsingFmt(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.TitleSquareMetaUsingFmt(testStringer{}, testStringer{}, testStringer{}) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMetaUsingFmt", actual)
}

func Test_Cov2_WithBracketsQuotation(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.WithBracketsQuotation("hello") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBracketsQuotation", actual)
}

func Test_Cov2_WithCurlyQuotation(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.WithCurlyQuotation("hello") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithCurlyQuotation", actual)
}

func Test_Cov2_WithParenthesisQuotation(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.WithParenthesisQuotation("hello") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesisQuotation", actual)
}

func Test_Cov2_CurlyWrapOption(t *testing.T) {
	actual := args.Map{
		"skipIfExists":    simplewrap.CurlyWrapOption(true, "{hello}"),
		"noSkip":          simplewrap.CurlyWrapOption(false, "hello"),
		"skipNotPresent":  simplewrap.CurlyWrapOption(true, "hello"),
	}
	expected := args.Map{
		"skipIfExists":    "{hello}",
		"noSkip":          "{hello}",
		"skipNotPresent":  "{hello}",
	}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption", actual)
}

func Test_Cov2_DoubleQuoteWrapElements_Nil(t *testing.T) {
	r := simplewrap.DoubleQuoteWrapElements(false, nil...)
	actual := args.Map{"isNotNil": r != nil}
	expected := args.Map{"isNotNil": true}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements_Nil", actual)
}

func Test_Cov2_DoubleQuoteWrapElements_EmptySlice(t *testing.T) {
	actual := args.Map{"len": len(simplewrap.DoubleQuoteWrapElements(false))}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements_EmptySlice", actual)
}

func Test_Cov2_DoubleQuoteWrapElements_SkipExistence(t *testing.T) {
	actual := args.Map{"len": len(simplewrap.DoubleQuoteWrapElements(true, `"already"`, "naked"))}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements_SkipExistence", actual)
}

func Test_Cov2_DoubleQuoteWrapElementsWithIndexes(t *testing.T) {
	actual := args.Map{
		"nilNotNil":     simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...) != nil,
		"emptyLen":      len(simplewrap.DoubleQuoteWrapElementsWithIndexes()),
		"itemsLen":      len(simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")),
		"containsIndex": strings.Contains(simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")[0], "[0]"),
	}
	expected := args.Map{
		"nilNotNil":     true,
		"emptyLen":      0,
		"itemsLen":      2,
		"containsIndex": true,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes", actual)
}

func Test_Cov2_WithDoubleQuote_Empty(t *testing.T) {
	actual := args.Map{"result": simplewrap.WithDoubleQuote("")}
	expected := args.Map{"result": `""`}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote_Empty", actual)
}

func Test_Cov2_WithDoubleQuoteAny_Int(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.WithDoubleQuoteAny(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny_Int", actual)
}

func Test_Cov2_WithSingleQuote_Empty(t *testing.T) {
	actual := args.Map{"result": simplewrap.WithSingleQuote("")}
	expected := args.Map{"result": "''"}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote_Empty", actual)
}

func Test_Cov2_ToJsonName_Int(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.ToJsonName(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName_Int", actual)
}

func Test_Cov2_WithCurly_Int(t *testing.T) {
	actual := args.Map{"contains42": strings.Contains(simplewrap.WithCurly(42), "42")}
	expected := args.Map{"contains42": true}
	expected.ShouldBeEqual(t, 0, "WithCurly_Int", actual)
}

func Test_Cov2_With_Empty(t *testing.T) {
	actual := args.Map{"result": simplewrap.With("", "", "")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "With_Empty", actual)
}

func Test_Cov2_WithStartEnd_Empty(t *testing.T) {
	actual := args.Map{"result": simplewrap.WithStartEnd("", "")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "WithStartEnd_Empty", actual)
}

func Test_Cov2_MsgWrapNumber_Int64(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.MsgWrapNumber("total", int64(100)) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber_Int64", actual)
}

func Test_Cov2_MsgCsvItems_Empty(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.MsgCsvItems("msg") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems_Empty", actual)
}

func Test_Cov2_ConditionalWrapWith_BothPresent2Char(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{}", '}')}
	expected := args.Map{"result": "{}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith_BothPresent2Char", actual)
}

func Test_Cov2_CurlyWrap_Int(t *testing.T) {
	actual := args.Map{"contains42": strings.Contains(simplewrap.CurlyWrap(42), "42")}
	expected := args.Map{"contains42": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrap_Int", actual)
}

func Test_Cov2_SquareWrap_Int(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.SquareWrap(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SquareWrap_Int", actual)
}

func Test_Cov2_ParenthesisWrap_Int(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.ParenthesisWrap(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap_Int", actual)
}

func Test_Cov2_TitleCurlyWrap_Int(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.TitleCurlyWrap("t", 42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyWrap_Int", actual)
}

func Test_TitleSquare_Int_Cov2(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.TitleSquare("t", 42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquare_Int", actual)
}

func Test_WithBrackets_Int_Cov2(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.WithBrackets(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBrackets_Int", actual)
}

func Test_WithParenthesis_Int_Cov2(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.WithParenthesis(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesis_Int", actual)
}

func Test_CurlyWrapIf_Stringer_Cov2(t *testing.T) {
	actual := args.Map{"notEmpty": simplewrap.CurlyWrapIf(true, testStringer{}) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf_Stringer", actual)
}

func Test_CurlyWrapIf_FmtStringer_Cov2(t *testing.T) {
	var s fmt.Stringer = testStringer{}
	actual := args.Map{"notEmpty": simplewrap.CurlyWrapIf(true, s) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf_FmtStringer", actual)
}
