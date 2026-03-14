package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── ConditionalWrapWith (startChar byte, input string, endChar byte) ──

func Test_Cov5_ConditionalWrapWith_Wrapped(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('[', "x", ']')}
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith wraps -- not already wrapped", actual)
}

func Test_Cov5_ConditionalWrapWith_AlreadyWrapped(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('[', "[x]", ']')}
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith no-op -- already wrapped", actual)
}

func Test_Cov5_ConditionalWrapWith_Empty(t *testing.T) {
	result := simplewrap.ConditionalWrapWith('[', "", ']')
	actual := args.Map{"result": result}
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith empty -- just brackets", actual)
}

// ── MsgWrapMsg ──

func Test_Cov5_MsgWrapMsg(t *testing.T) {
	result := simplewrap.MsgWrapMsg("hello", "world")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg -- not empty", actual)
}

// ── CurlyWrapOption (isSkipIfExists bool, source any) ──

func Test_Cov5_CurlyWrapOption_NonEmpty(t *testing.T) {
	result := simplewrap.CurlyWrapOption(false, "hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption non-empty -- wrapped", actual)
}

func Test_Cov5_CurlyWrapOption_SkipIfExists(t *testing.T) {
	result := simplewrap.CurlyWrapOption(true, "{hello}")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption skip if exists -- no double wrap", actual)
}

func Test_Cov5_CurlyWrapOption_Empty(t *testing.T) {
	result := simplewrap.CurlyWrapOption(false, "")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption empty -- just curlies", actual)
}

// ── WithBracketsQuotation ──

func Test_Cov5_WithBracketsQuotation(t *testing.T) {
	result := simplewrap.WithBracketsQuotation("x")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBracketsQuotation -- wrapped", actual)
}

// ── WithCurlyQuotation ──

func Test_Cov5_WithCurlyQuotation(t *testing.T) {
	result := simplewrap.WithCurlyQuotation("x")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithCurlyQuotation -- wrapped", actual)
}

// ── WithParenthesisQuotation ──

func Test_Cov5_WithParenthesisQuotation(t *testing.T) {
	result := simplewrap.WithParenthesisQuotation("x")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesisQuotation -- wrapped", actual)
}

// ── MsgCsvItems ──

func Test_Cov5_MsgCsvItems_Empty(t *testing.T) {
	result := simplewrap.MsgCsvItems("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems no items -- msg only", actual)
}
