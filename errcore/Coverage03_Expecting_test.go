package errcore

import (
	"testing"
)

func TestExpecting(t *testing.T) {
	s := Expecting("title", "exp", "act")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpectingSimple(t *testing.T) {
	s := ExpectingSimple("title", "exp", "act")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpectingSimpleNoType(t *testing.T) {
	s := ExpectingSimpleNoType("title", "exp", "act")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpectingNotEqualSimpleNoType(t *testing.T) {
	s := ExpectingNotEqualSimpleNoType("title", "exp", "act")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpectingSimpleNoTypeError(t *testing.T) {
	err := ExpectingSimpleNoTypeError("title", "exp", "act")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpectingErrorSimpleNoType(t *testing.T) {
	err := ExpectingErrorSimpleNoType("title", "exp", "act")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpectingErrorSimpleNoTypeNewLineEnds(t *testing.T) {
	err := ExpectingErrorSimpleNoTypeNewLineEnds("title", "exp", "act")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestWasExpectingErrorF(t *testing.T) {
	err := WasExpectingErrorF("exp", "act", "title %s", "val")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpectingFuture(t *testing.T) {
	rec := ExpectingFuture("title", "exp")
	if rec == nil {
		t.Fatal("expected non-nil")
	}
	if rec.Message("act") == "" {
		t.Fatal("expected non-empty")
	}
	if rec.MessageSimple("act") == "" {
		t.Fatal("expected non-empty")
	}
	if rec.MessageSimpleNoType("act") == "" {
		t.Fatal("expected non-empty")
	}
	if rec.Error("act") == nil {
		t.Fatal("expected non-nil")
	}
	if rec.ErrorSimple("act") == nil {
		t.Fatal("expected non-nil")
	}
	if rec.ErrorSimpleNoType("act") == nil {
		t.Fatal("expected non-nil")
	}
}

func TestExpectationMessageDef_ExpectedSafeString(t *testing.T) {
	def := ExpectationMessageDef{Expected: "hello"}
	s := def.ExpectedSafeString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	// call again for cache
	s2 := def.ExpectedSafeString()
	if s2 != s {
		t.Fatal("cache miss")
	}
}

func TestExpectationMessageDef_ExpectedSafeString_Nil(t *testing.T) {
	def := ExpectationMessageDef{}
	s := def.ExpectedSafeString()
	if s != "" {
		t.Fatal("expected empty for nil Expected")
	}
}

func TestExpectationMessageDef_ExpectedStringTrim(t *testing.T) {
	def := ExpectationMessageDef{Expected: " hello "}
	s := def.ExpectedStringTrim()
	if s != "hello" {
		t.Fatal("expected trimmed")
	}
}

func TestExpectationMessageDef_ToString(t *testing.T) {
	def := ExpectationMessageDef{Expected: "e", When: "w", FuncName: "f"}
	s := def.ToString("act")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestExpectationMessageDef_PrintIf(t *testing.T) {
	def := ExpectationMessageDef{Expected: "e", When: "w"}
	def.PrintIf(false, "act") // should not print
	def.PrintIf(true, "act")  // should print
}

func TestExpectationMessageDef_PrintIfFailed(t *testing.T) {
	def := ExpectationMessageDef{Expected: "e", When: "w"}
	def.PrintIfFailed(true, false, "act") // not failed
	def.PrintIfFailed(false, true, "act") // not print on fail
	def.PrintIfFailed(true, true, "act")  // should print
}

func TestExpectationMessageDef_Print(t *testing.T) {
	def := ExpectationMessageDef{Expected: "e", When: "w"}
	def.Print("act")
}
