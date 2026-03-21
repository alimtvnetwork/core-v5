package stringcompareastests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── VerifyMessage ignore-case positive mismatch ──

func Test_Cov2_VerifyMessage_IgnoreCase_Positive(t *testing.T) {
	msg := stringcompareas.StartsWith.VerifyMessage(true, "Hello", "world")
	actual := args.Map{"nonEmpty": msg != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage returns correct value -- ignore case positive mismatch", actual)
}

// ── VerifyMessage ignore-case negative mismatch ──

func Test_Cov2_VerifyMessage_IgnoreCase_Negative(t *testing.T) {
	msg := stringcompareas.NotStartsWith.VerifyMessage(true, "Hello World", "hello")
	actual := args.Map{"nonEmpty": msg != "", "isNeg": strings.Contains(msg, "negative")}
	expected := args.Map{"nonEmpty": true, "isNeg": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage returns correct value -- ignore case negative mismatch", actual)
}

// ── VerifyError case sensitive mismatch ──

func Test_Cov2_VerifyErrorCaseSensitive_Mismatch(t *testing.T) {
	err := stringcompareas.Equal.VerifyErrorCaseSensitive("hello", "world")
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorCaseSensitive returns error -- mismatch", actual)
}

// ── VerifyMessageCaseSensitive mismatch ──

func Test_Cov2_VerifyMessageCaseSensitive_Mismatch(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessageCaseSensitive("hello", "world")
	actual := args.Map{"nonEmpty": msg != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessageCaseSensitive returns correct value -- mismatch", actual)
}

// ── NotAnyChars IsNegativeCondition ──

func Test_Cov2_NotAnyChars_IsNegativeCondition(t *testing.T) {
	actual := args.Map{"result": stringcompareas.NotAnyChars.IsNegativeCondition()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotAnyChars returns correct value -- is negative", actual)
}

// ── NotEndsWith IsNegativeCondition ──

func Test_Cov2_NotEndsWith_IsNegativeCondition(t *testing.T) {
	actual := args.Map{"result": stringcompareas.NotEndsWith.IsNegativeCondition()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEndsWith returns non-empty -- is negative", actual)
}

// ── NotContains IsNegativeCondition ──

func Test_Cov2_NotContains_IsNegativeCondition(t *testing.T) {
	actual := args.Map{"result": stringcompareas.NotContains.IsNegativeCondition()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotContains returns correct value -- is negative", actual)
}

// ── NotMatchRegex IsNegativeCondition ──

func Test_Cov2_NotMatchRegex_IsNegativeCondition(t *testing.T) {
	actual := args.Map{"result": stringcompareas.NotMatchRegex.IsNegativeCondition()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotMatchRegex returns correct value -- is negative", actual)
}

// ── IsAnyEnumsEqual match path ──

func Test_Cov2_IsAnyEnumsEqual_Match(t *testing.T) {
	a := stringcompareas.Equal
	b := stringcompareas.Equal
	actual := args.Map{"result": a.IsAnyEnumsEqual(&b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyEnumsEqual returns correct value -- match", actual)
}

// ── VerifyError ignore case match ──

func Test_Cov2_VerifyError_IgnoreCase_Match(t *testing.T) {
	err := stringcompareas.Equal.VerifyError(true, "Hello", "hello")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "VerifyError returns error -- ignore case match", actual)
}
