package errcore

import (
	"errors"
	"fmt"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// AssertDiffOnMismatch — mismatch branch
// Covers AssertDiffOnMismatch.go L32-33
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov10_AssertDiffOnMismatch_Mismatch(t *testing.T) {
	// Arrange
	sub := &testing.T{}
	act := []string{"actual line"}
	exp := []string{"expected line"}

	// Act — exercises the mismatch path (PrintDiff + Errorf)
	AssertDiffOnMismatch(sub, 0, "mismatch test", act, exp)

	// Assert — sub.Failed() should be true
	if !sub.Failed() {
		t.Error("expected sub test to fail on mismatch")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// AssertErrorDiffOnMismatch — error mismatch branch
// Covers AssertDiffOnMismatch.go L62-63
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov10_AssertErrorDiffOnMismatch_Mismatch(t *testing.T) {
	// Arrange
	sub := &testing.T{}
	err := errors.New("some error text")
	exp := []string{"different expected"}

	// Act
	AssertErrorDiffOnMismatch(sub, 0, "error mismatch test", err, exp)

	// Assert
	if !sub.Failed() {
		t.Error("expected sub test to fail on error mismatch")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ExpectationMessageDef — ExpectedSafeString cached + Print
// Covers ExpectationMessageDef.go L24-26, L67
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov10_ExpectationMessageDef_ExpectedSafeString_Cached(t *testing.T) {
	// Arrange
	cached := "pre-cached"
	def := ExpectationMessageDef{
		Expected:       "test",
		expectedString: &cached,
	}

	// Act
	result := def.ExpectedSafeString()

	// Assert
	if result != "pre-cached" {
		t.Errorf("got %q, want %q", result, "pre-cached")
	}
}

func Test_Cov10_ExpectationMessageDef_Print(t *testing.T) {
	// Arrange
	def := ExpectationMessageDef{
		Expected: "expected-val",
		When:     "test scenario",
	}

	// Act — exercises Print path (slog.Warn)
	def.Print("actual-val")

	// Assert — no panic = success
}

// ══════════════════════════════════════════════════════════════════════════════
// HandleCompiledErrorGetter — non-nil getter with error → panic
// Covers HandleCompiledErrorGetter.go L8-14
// ══════════════════════════════════════════════════════════════════════════════

type mockCompiledErrorGetter struct{ err error }

func (m *mockCompiledErrorGetter) CompiledError() error { return m.err }

func Test_Cov10_HandleCompiledErrorGetter_WithError(t *testing.T) {
	// Arrange
	getter := &mockCompiledErrorGetter{err: errors.New("compiled err")}

	// Act & Assert — should panic
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic from HandleCompiledErrorGetter")
		}
	}()

	HandleCompiledErrorGetter(getter)
}

func Test_Cov10_HandleCompiledErrorGetter_NilError(t *testing.T) {
	// Arrange
	getter := &mockCompiledErrorGetter{err: nil}

	// Act — should NOT panic
	HandleCompiledErrorGetter(getter)
}

// ══════════════════════════════════════════════════════════════════════════════
// HandleCompiledErrorWithTracesGetter — non-nil getter with error → panic
// Covers HandleCompiledErrorWithTracesGetter.go L8-14
// ══════════════════════════════════════════════════════════════════════════════

type mockCompiledErrorWithTracesGetter struct{ err error }

func (m *mockCompiledErrorWithTracesGetter) CompiledErrorWithStackTraces() error { return m.err }

func Test_Cov10_HandleCompiledErrorWithTracesGetter_WithError(t *testing.T) {
	getter := &mockCompiledErrorWithTracesGetter{err: errors.New("traces err")}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic")
		}
	}()

	HandleCompiledErrorWithTracesGetter(getter)
}

func Test_Cov10_HandleCompiledErrorWithTracesGetter_NilError(t *testing.T) {
	getter := &mockCompiledErrorWithTracesGetter{err: nil}
	HandleCompiledErrorWithTracesGetter(getter)
}

// ══════════════════════════════════════════════════════════════════════════════
// HandleErrorGetter — non-nil getter with error → panic
// Covers HandleErrorGetter.go L8-14
// ══════════════════════════════════════════════════════════════════════════════

type mockErrorGetter struct{ err error }

func (m *mockErrorGetter) Error() error { return m.err }

func Test_Cov10_HandleErrorGetter_WithError(t *testing.T) {
	getter := &mockErrorGetter{err: errors.New("getter err")}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic")
		}
	}()

	HandleErrorGetter(getter)
}

func Test_Cov10_HandleErrorGetter_NilError(t *testing.T) {
	getter := &mockErrorGetter{err: nil}
	HandleErrorGetter(getter)
}

// ══════════════════════════════════════════════════════════════════════════════
// CompiledErrorString — nil compiled branch (dead but test it)
// Covers CompiledError.go L30-32
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov10_CompiledErrorString_EmptyAdditional(t *testing.T) {
	// Arrange — mainErr non-nil, additional empty → returns mainErr directly
	err := errors.New("main error")

	// Act
	result := CompiledErrorString(err, "")

	// Assert
	if result != "main error" {
		t.Errorf("got %q, want %q", result, "main error")
	}
}

func Test_Cov10_CompiledErrorString_WithAdditional(t *testing.T) {
	err := errors.New("main error")
	result := CompiledErrorString(err, "context")
	expected := fmt.Sprintf("context: %s", "main error")
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}
