package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
)

// ── WithDoubleQuote / WithDoubleQuoteAny / WithSingleQuote ──

func Test_WithDoubleQuote_Coverage(t *testing.T) {
	result := simplewrap.WithDoubleQuote("hello")
	if result != `"hello"` {
		t.Errorf("expected \"hello\", got %s", result)
	}
}

func Test_WithDoubleQuoteAny_Coverage(t *testing.T) {
	result := simplewrap.WithDoubleQuoteAny(42)
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_WithSingleQuote_Coverage(t *testing.T) {
	result := simplewrap.WithSingleQuote("hello")
	if result != `'hello'` {
		t.Errorf("expected 'hello', got %s", result)
	}
}

// ── CurlyWrap / CurlyWrapIf ──

func Test_CurlyWrap_Coverage(t *testing.T) {
	result := simplewrap.CurlyWrap("hello")
	if result != "{hello}" {
		t.Errorf("expected {hello}, got %s", result)
	}
}

func Test_CurlyWrapIf_Coverage(t *testing.T) {
	wrapped := simplewrap.CurlyWrapIf(true, "hello")
	if wrapped != "{hello}" {
		t.Errorf("expected {hello}, got %s", wrapped)
	}

	notWrapped := simplewrap.CurlyWrapIf(false, "hello")
	if notWrapped != "hello" {
		t.Errorf("expected hello, got %s", notWrapped)
	}
}

// ── SquareWrap / SquareWrapIf ──

func Test_SquareWrap_Coverage(t *testing.T) {
	result := simplewrap.SquareWrap("hello")
	if result != "[hello]" {
		t.Errorf("expected [hello], got %s", result)
	}
}

func Test_SquareWrapIf_Coverage(t *testing.T) {
	wrapped := simplewrap.SquareWrapIf(true, "hello")
	if wrapped != "[hello]" {
		t.Errorf("expected [hello], got %s", wrapped)
	}

	notWrapped := simplewrap.SquareWrapIf(false, "hello")
	if notWrapped != "hello" {
		t.Errorf("expected hello, got %s", notWrapped)
	}
}

// ── ParenthesisWrap / ParenthesisWrapIf ──

func Test_ParenthesisWrap_Coverage(t *testing.T) {
	result := simplewrap.ParenthesisWrap("hello")
	if result != "(hello)" {
		t.Errorf("expected (hello), got %s", result)
	}
}

func Test_ParenthesisWrapIf_Coverage(t *testing.T) {
	wrapped := simplewrap.ParenthesisWrapIf(true, "hello")
	if wrapped != "(hello)" {
		t.Errorf("expected (hello), got %s", wrapped)
	}

	notWrapped := simplewrap.ParenthesisWrapIf(false, "hello")
	if notWrapped != "hello" {
		t.Errorf("expected hello, got %s", notWrapped)
	}
}

// ── With / WithPtr / WithStartEnd / WithStartEndPtr ──

func Test_With_Coverage(t *testing.T) {
	result := simplewrap.With("[", "hello", "]")
	if result != "[hello]" {
		t.Errorf("expected [hello], got %s", result)
	}
}

func Test_WithPtr_Coverage(t *testing.T) {
	start, source, end := "[", "hello", "]"
	result := simplewrap.WithPtr(&start, &source, &end)
	if *result != "[hello]" {
		t.Errorf("expected [hello], got %s", *result)
	}

	// Nil cases
	resultNil := simplewrap.WithPtr(nil, &source, nil)
	if *resultNil != "hello" {
		t.Errorf("expected hello, got %s", *resultNil)
	}

	resultNilSrc := simplewrap.WithPtr(&start, nil, &end)
	if *resultNilSrc != "[]" {
		t.Errorf("expected [], got %s", *resultNilSrc)
	}
}

func Test_WithStartEnd_Coverage(t *testing.T) {
	result := simplewrap.WithStartEnd("'", "hello")
	if result != "'hello'" {
		t.Errorf("expected 'hello', got %s", result)
	}
}

func Test_WithStartEndPtr_Coverage(t *testing.T) {
	wrapper, source := "'", "hello"
	result := simplewrap.WithStartEndPtr(&wrapper, &source)
	if *result != "'hello'" {
		t.Errorf("expected 'hello', got %s", *result)
	}
}

// ── WithBrackets / WithCurly / WithParenthesis ──

func Test_WithBrackets_Coverage(t *testing.T) {
	result := simplewrap.WithBrackets("hello")
	if result != "[hello]" {
		t.Errorf("expected [hello], got %s", result)
	}
}

func Test_WithCurly_Coverage(t *testing.T) {
	result := simplewrap.WithCurly("hello")
	if result != "{hello}" {
		t.Errorf("expected {hello}, got %s", result)
	}
}

func Test_WithParenthesis_Coverage(t *testing.T) {
	result := simplewrap.WithParenthesis("hello")
	if result != "(hello)" {
		t.Errorf("expected (hello), got %s", result)
	}
}

// ── TitleCurlyWrap / TitleSquare ──

func Test_TitleCurlyWrap_Coverage(t *testing.T) {
	result := simplewrap.TitleCurlyWrap("title", "value")
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_TitleSquare_Coverage(t *testing.T) {
	result := simplewrap.TitleSquare("title", "value")
	if result == "" {
		t.Error("should not be empty")
	}
}

// ── MsgWrapMsg / MsgWrapNumber / MsgCsvItems ──

func Test_MsgWrapMsg_Coverage(t *testing.T) {
	if simplewrap.MsgWrapMsg("", "") != "" {
		t.Error("both empty should be empty")
	}
	if simplewrap.MsgWrapMsg("", "wrapped") != "wrapped" {
		t.Error("empty msg should return wrapped")
	}
	if simplewrap.MsgWrapMsg("msg", "") != "msg" {
		t.Error("empty wrapped should return msg")
	}
	result := simplewrap.MsgWrapMsg("msg", "wrapped")
	if result == "" {
		t.Error("both non-empty should not be empty")
	}
}

func Test_MsgWrapNumber_Coverage(t *testing.T) {
	result := simplewrap.MsgWrapNumber("count", 42)
	if result == "" {
		t.Error("should not be empty")
	}
}

func Test_MsgCsvItems_Coverage(t *testing.T) {
	result := simplewrap.MsgCsvItems("items", "a", "b", "c")
	if result == "" {
		t.Error("should not be empty")
	}
}

// ── ToJsonName ──

func Test_ToJsonName_Coverage(t *testing.T) {
	result := simplewrap.ToJsonName("hello")
	if result == "" {
		t.Error("should not be empty")
	}
}

// ── ConditionalWrapWith ──

func Test_ConditionalWrapWith_Coverage(t *testing.T) {
	// Both present — return as-is
	result := simplewrap.ConditionalWrapWith('[', "[hello]", ']')
	if result != "[hello]" {
		t.Errorf("both present: expected [hello], got %s", result)
	}

	// Empty input — wrap
	result = simplewrap.ConditionalWrapWith('[', "", ']')
	if result != "[]" {
		t.Errorf("empty: expected [], got %s", result)
	}

	// Both missing — add both
	result = simplewrap.ConditionalWrapWith('[', "hello", ']')
	if result != "[hello]" {
		t.Errorf("both missing: expected [hello], got %s", result)
	}

	// Right missing
	result = simplewrap.ConditionalWrapWith('[', "[hello", ']')
	if result != "[hello]" {
		t.Errorf("right missing: expected [hello], got %s", result)
	}

	// Left missing
	result = simplewrap.ConditionalWrapWith('[', "hello]", ']')
	if result != "[hello]" {
		t.Errorf("left missing: expected [hello], got %s", result)
	}

	// Single char that matches start
	result = simplewrap.ConditionalWrapWith('[', "[", ']')
	if result != "[]" {
		t.Errorf("single char: expected [], got %s", result)
	}
}

// ── DoubleQuoteWrapElements ──

func Test_DoubleQuoteWrapElements_Coverage(t *testing.T) {
	// Normal
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")
	if len(result) != 2 {
		t.Error("should have 2 items")
	}
	if result[0] != `"a"` {
		t.Errorf("expected \"a\", got %s", result[0])
	}

	// Nil input
	result = simplewrap.DoubleQuoteWrapElements(false, )
	if len(result) != 0 {
		t.Error("empty input should return empty")
	}

	// With skip
	result = simplewrap.DoubleQuoteWrapElements(true, "a", "b")
	if len(result) != 2 {
		t.Error("should have 2 items")
	}
}

// ── DoubleQuoteWrapElementsWithIndexes ──

func Test_DoubleQuoteWrapElementsWithIndexes_Coverage(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(
		false, "a", "b")
	if len(result) != 2 {
		t.Error("should have 2 items")
	}
}
