package simplewraptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── ConditionalWrapWith exhaustive branch coverage ──

func Test_Cov3_ConditionalWrapWith_BothMissing(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "hello", '}')}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith both missing", actual)
}

func Test_Cov3_ConditionalWrapWith_LeftMissing(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "hello}", '}')}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith left missing", actual)
}

func Test_Cov3_ConditionalWrapWith_RightMissing(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{hello", '}')}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith right missing", actual)
}

func Test_Cov3_ConditionalWrapWith_BothPresent(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{hello}", '}')}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith both present", actual)
}

func Test_Cov3_ConditionalWrapWith_SingleCharMatchStart(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{", '}')}
	expected := args.Map{"result": "{}" }
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith single char start", actual)
}

func Test_Cov3_ConditionalWrapWith_Empty(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "", '}')}
	expected := args.Map{"result": "{}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith empty", actual)
}

// ── MsgWrapMsg all branches ──

func Test_Cov3_MsgWrapMsg_BothEmpty(t *testing.T) {
	actual := args.Map{"result": simplewrap.MsgWrapMsg("", "")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg both empty", actual)
}

func Test_Cov3_MsgWrapMsg_MsgEmpty(t *testing.T) {
	actual := args.Map{"result": simplewrap.MsgWrapMsg("", "wrapped")}
	expected := args.Map{"result": "wrapped"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg msg empty", actual)
}

func Test_Cov3_MsgWrapMsg_WrappedEmpty(t *testing.T) {
	actual := args.Map{"result": simplewrap.MsgWrapMsg("msg", "")}
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg wrapped empty", actual)
}

func Test_Cov3_MsgWrapMsg_BothPresent(t *testing.T) {
	result := simplewrap.MsgWrapMsg("msg", "wrapped")
	actual := args.Map{"containsMsg": strings.Contains(result, "msg"), "containsWrapped": strings.Contains(result, "wrapped")}
	expected := args.Map{"containsMsg": true, "containsWrapped": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg both present", actual)
}

// ── DoubleQuoteWrapElements branches ──

func Test_Cov3_DoubleQuoteWrapElements_NilInput(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)
	actual := args.Map{"notNil": result != nil, "len": len(result)}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements nil input", actual)
}

func Test_Cov3_DoubleQuoteWrapElements_SkipExistence(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(true, `"already"`, "naked")
	actual := args.Map{"len": len(result), "alreadyWrapped": result[0] == `"already"`, "nakedWrapped": strings.HasPrefix(result[1], `"`)}
	expected := args.Map{"len": 2, "alreadyWrapped": true, "nakedWrapped": true}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements skip existence", actual)
}

// ── DoubleQuoteWrapElementsWithIndexes branches ──

func Test_Cov3_DoubleQuoteWrapElementsWithIndexes_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...)
	actual := args.Map{"notNil": result != nil, "len": len(result)}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes nil", actual)
}

func Test_Cov3_DoubleQuoteWrapElementsWithIndexes_Empty(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes empty", actual)
}

// ── WithPtr nil combinations ──

func Test_Cov3_WithPtr_AllNil(t *testing.T) {
	result := simplewrap.WithPtr(nil, nil, nil)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "WithPtr all nil", actual)
}

// ── CurlyWrapOption branches ──

func Test_Cov3_CurlyWrapOption_SkipAlreadyWrapped(t *testing.T) {
	actual := args.Map{"result": simplewrap.CurlyWrapOption(true, "{test}")}
	expected := args.Map{"result": "{test}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption skip already wrapped", actual)
}

func Test_Cov3_CurlyWrapOption_NoSkip(t *testing.T) {
	actual := args.Map{"result": simplewrap.CurlyWrapOption(false, "test")}
	expected := args.Map{"result": "{test}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption no skip", actual)
}

// ── WithBracketsQuotation / WithCurlyQuotation / WithParenthesisQuotation ──

func Test_Cov3_WithBracketsQuotation(t *testing.T) {
	r := simplewrap.WithBracketsQuotation("test")
	actual := args.Map{"containsBracket": strings.Contains(r, "["), "containsQuote": strings.Contains(r, `"`)}
	expected := args.Map{"containsBracket": true, "containsQuote": true}
	expected.ShouldBeEqual(t, 0, "WithBracketsQuotation", actual)
}

func Test_Cov3_WithCurlyQuotation(t *testing.T) {
	r := simplewrap.WithCurlyQuotation("test")
	actual := args.Map{"containsCurly": strings.Contains(r, "{"), "containsQuote": strings.Contains(r, `"`)}
	expected := args.Map{"containsCurly": true, "containsQuote": true}
	expected.ShouldBeEqual(t, 0, "WithCurlyQuotation", actual)
}

func Test_Cov3_WithParenthesisQuotation(t *testing.T) {
	r := simplewrap.WithParenthesisQuotation("test")
	actual := args.Map{"containsParen": strings.Contains(r, "("), "containsQuote": strings.Contains(r, `"`)}
	expected := args.Map{"containsParen": true, "containsQuote": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesisQuotation", actual)
}
