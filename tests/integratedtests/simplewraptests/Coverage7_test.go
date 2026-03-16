package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── WithDoubleQuote ──

func Test_Cov7_WithDoubleQuote(t *testing.T) {
	result := simplewrap.WithDoubleQuote("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": `"hello"`}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote", actual)
}

// ── WithSingleQuote ──

func Test_Cov7_WithSingleQuote(t *testing.T) {
	result := simplewrap.WithSingleQuote("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "'hello'"}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote", actual)
}

// ── With ──

func Test_Cov7_With(t *testing.T) {
	result := simplewrap.With("[", "hello", "]")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "With", actual)
}

// ── WithPtr ──

func Test_Cov7_WithPtr_AllNonNil(t *testing.T) {
	s, e, src := "[", "]", "hello"
	result := simplewrap.WithPtr(&s, &src, &e)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithPtr all non-nil", actual)
}

func Test_Cov7_WithPtr_NilSource(t *testing.T) {
	s, e := "[", "]"
	result := simplewrap.WithPtr(&s, nil, &e)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "WithPtr nil source", actual)
}

func Test_Cov7_WithPtr_NilStartEnd(t *testing.T) {
	src := "hello"
	result := simplewrap.WithPtr(nil, &src, nil)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "WithPtr nil start/end", actual)
}

// ── WithStartEnd ──

func Test_Cov7_WithStartEnd(t *testing.T) {
	result := simplewrap.WithStartEnd("*", "hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "*hello*"}
	expected.ShouldBeEqual(t, 0, "WithStartEnd", actual)
}

// ── WithBrackets ──

func Test_Cov7_WithBrackets(t *testing.T) {
	result := simplewrap.WithBrackets("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithBrackets", actual)
}

// ── WithCurly ──

func Test_Cov7_WithCurly(t *testing.T) {
	result := simplewrap.WithCurly("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "WithCurly", actual)
}

// ── WithParenthesis ──

func Test_Cov7_WithParenthesis(t *testing.T) {
	result := simplewrap.WithParenthesis("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "WithParenthesis", actual)
}

// ── WithDoubleQuoteAny ──

func Test_Cov7_WithDoubleQuoteAny(t *testing.T) {
	result := simplewrap.WithDoubleQuoteAny(42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny", actual)
}

// ── CurlyWrap ──

func Test_Cov7_CurlyWrap(t *testing.T) {
	result := simplewrap.CurlyWrap("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrap", actual)
}

// ── SquareWrap ──

func Test_Cov7_SquareWrap(t *testing.T) {
	result := simplewrap.SquareWrap("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "SquareWrap", actual)
}

// ── ParenthesisWrap ──

func Test_Cov7_ParenthesisWrap(t *testing.T) {
	result := simplewrap.ParenthesisWrap("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap", actual)
}

// ── ToJsonName ──

func Test_Cov7_ToJsonName(t *testing.T) {
	result := simplewrap.ToJsonName("name")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName", actual)
}

// ── MsgWrapMsg ──

func Test_Cov7_MsgWrapMsg_BothEmpty(t *testing.T) {
	result := simplewrap.MsgWrapMsg("", "")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg both empty", actual)
}

func Test_Cov7_MsgWrapMsg_EmptyMsg(t *testing.T) {
	result := simplewrap.MsgWrapMsg("", "wrapped")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "wrapped"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg empty msg", actual)
}

func Test_Cov7_MsgWrapMsg_EmptyWrapped(t *testing.T) {
	result := simplewrap.MsgWrapMsg("msg", "")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg empty wrapped", actual)
}

func Test_Cov7_MsgWrapMsg_Both(t *testing.T) {
	result := simplewrap.MsgWrapMsg("msg", "wrapped")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg both", actual)
}

// ── MsgWrapNumber ──

func Test_Cov7_MsgWrapNumber(t *testing.T) {
	result := simplewrap.MsgWrapNumber("count", 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber", actual)
}

// ── MsgCsvItems ──

func Test_Cov7_MsgCsvItems(t *testing.T) {
	result := simplewrap.MsgCsvItems("msg", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems", actual)
}

// ── ConditionalWrapWith ──

func Test_Cov7_ConditionalWrapWith_Empty(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith empty", actual)
}

func Test_Cov7_ConditionalWrapWith_AlreadyWrapped(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "[hello]", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith already wrapped", actual)
}

func Test_Cov7_ConditionalWrapWith_MissingRight(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "[hello", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith missing right", actual)
}

func Test_Cov7_ConditionalWrapWith_MissingLeft(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "hello]", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith missing left", actual)
}

func Test_Cov7_ConditionalWrapWith_BothMissing(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "hello", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith both missing", actual)
}

func Test_Cov7_ConditionalWrapWith_SingleCharMatch(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "[", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith single char match", actual)
}

// ── DoubleQuoteWrapElements ──

func Test_Cov7_DoubleQuoteWrapElements_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements nil", actual)
}

func Test_Cov7_DoubleQuoteWrapElements_Empty(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements empty", actual)
}

func Test_Cov7_DoubleQuoteWrapElements_NonEmpty(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": `"a"`}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements non-empty", actual)
}

func Test_Cov7_DoubleQuoteWrapElements_SkipOnExistence(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(true, "a", `"b"`)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements skip on existence", actual)
}
