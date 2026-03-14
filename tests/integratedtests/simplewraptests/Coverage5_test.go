package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── ConditionalWrapWith ──

func Test_Cov5_ConditionalWrapWith_True(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith(true, "[", "x", "]")}
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith true -- wrapped", actual)
}

func Test_Cov5_ConditionalWrapWith_False(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith(false, "[", "x", "]")}
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith false -- unwrapped", actual)
}

// ── MsgWrapMsg ──

func Test_Cov5_MsgWrapMsg(t *testing.T) {
	result := simplewrap.MsgWrapMsg("hello", "world")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg -- not empty", actual)
}

// ── CurlyWrapOption ──

func Test_Cov5_CurlyWrapOption_NonEmpty(t *testing.T) {
	result := simplewrap.CurlyWrapOption("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption non-empty -- wrapped", actual)
}

func Test_Cov5_CurlyWrapOption_Empty(t *testing.T) {
	result := simplewrap.CurlyWrapOption("")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption empty -- empty", actual)
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

// ── wrapDoubleQuoteOnNonExist (toString helper) ──

func Test_Cov5_MsgCsvItems_Empty(t *testing.T) {
	result := simplewrap.MsgCsvItems("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems no items -- msg only", actual)
}
