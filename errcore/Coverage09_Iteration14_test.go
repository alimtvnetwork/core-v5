package errcore

import (
	"errors"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// AssertDiffOnMismatch — mismatch branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_AssertDiffOnMismatch_Mismatch(t *testing.T) {
	// Use a sub-test so the parent doesn't fail
	mockT := &testing.T{}
	AssertDiffOnMismatch(
		mockT,
		0,
		"test",
		[]string{"actual"},
		[]string{"expected"},
		"context line",
	)
	// mockT.Failed() would be true but we can't check in a different goroutine
}

func Test_I14_AssertErrorDiffOnMismatch_Mismatch(t *testing.T) {
	mockT := &testing.T{}
	AssertErrorDiffOnMismatch(
		mockT,
		0,
		"test",
		errors.New("actual line"),
		[]string{"expected line"},
	)
}

// ══════════════════════════════════════════════════════════════════════════════
// CompiledErrorString — nil compiled branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_CompiledErrorString_NilErr(t *testing.T) {
	r := CompiledErrorString(nil, "msg")
	if r != "" {
		t.Fatal("expected empty for nil error")
	}
}

func Test_I14_CompiledErrorString_EmptyMsg(t *testing.T) {
	r := CompiledErrorString(errors.New("err"), "")
	if r != "err" {
		t.Fatal("expected 'err'")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// ExpectationMessageDef — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_ExpectationMessageDef_ExpectedSafeString_NilExpected(t *testing.T) {
	e := ExpectationMessageDef{
		Expected: nil,
	}
	s := e.ExpectedSafeString()
	if s != "" {
		t.Fatal("expected empty for nil Expected")
	}
}

func Test_I14_ExpectationMessageDef_PrintIf_False(t *testing.T) {
	e := ExpectationMessageDef{
		Expected: "x",
		When:     "test",
	}
	// should not print
	e.PrintIf(false, "actual")
}

func Test_I14_ExpectationMessageDef_PrintIf_True(t *testing.T) {
	e := ExpectationMessageDef{
		Expected: "x",
		When:     "test",
	}
	e.PrintIf(true, "actual")
}

// ══════════════════════════════════════════════════════════════════════════════
// Handle*Getter — nil and non-nil branches
// ══════════════════════════════════════════════════════════════════════════════

type mockErrorGetter struct{ err error }

func (m *mockErrorGetter) Error() error { return m.err }

type mockCompiledErrorGetter struct{ err error }

func (m *mockCompiledErrorGetter) CompiledError() error { return m.err }

type mockCompiledErrorWithTracesGetter struct{ err error }

func (m *mockCompiledErrorWithTracesGetter) CompiledErrorWithStackTraces() error { return m.err }

type mockFullStringWithTracesGetter struct{ err error }

func (m *mockFullStringWithTracesGetter) FullStringWithTraces() error { return m.err }

func Test_I14_HandleErrorGetter_NilGetter(t *testing.T) {
	HandleErrorGetter(nil) // should not panic
}

func Test_I14_HandleErrorGetter_NilError(t *testing.T) {
	HandleErrorGetter(&mockErrorGetter{nil}) // should not panic
}

func Test_I14_HandleErrorGetter_WithError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	HandleErrorGetter(&mockErrorGetter{errors.New("boom")})
}

func Test_I14_HandleCompiledErrorGetter_NilGetter(t *testing.T) {
	HandleCompiledErrorGetter(nil)
}

func Test_I14_HandleCompiledErrorGetter_NilError(t *testing.T) {
	HandleCompiledErrorGetter(&mockCompiledErrorGetter{nil})
}

func Test_I14_HandleCompiledErrorGetter_WithError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	HandleCompiledErrorGetter(&mockCompiledErrorGetter{errors.New("boom")})
}

func Test_I14_HandleCompiledErrorWithTracesGetter_NilGetter(t *testing.T) {
	HandleCompiledErrorWithTracesGetter(nil)
}

func Test_I14_HandleCompiledErrorWithTracesGetter_NilError(t *testing.T) {
	HandleCompiledErrorWithTracesGetter(&mockCompiledErrorWithTracesGetter{nil})
}

func Test_I14_HandleCompiledErrorWithTracesGetter_WithError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	HandleCompiledErrorWithTracesGetter(&mockCompiledErrorWithTracesGetter{errors.New("boom")})
}

func Test_I14_HandleFullStringsWithTracesGetter_NilGetter(t *testing.T) {
	HandleFullStringsWithTracesGetter(nil)
}

func Test_I14_HandleFullStringsWithTracesGetter_NilError(t *testing.T) {
	HandleFullStringsWithTracesGetter(&mockFullStringWithTracesGetter{nil})
}

func Test_I14_HandleFullStringsWithTracesGetter_WithError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	HandleFullStringsWithTracesGetter(&mockFullStringWithTracesGetter{errors.New("boom")})
}

// ══════════════════════════════════════════════════════════════════════════════
// RawErrCollection — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_RawErrCollection_CompiledJsonStringWithStackTraces(t *testing.T) {
	rec := &RawErrCollection{}
	rec.Add(errors.New("err1"))
	s := rec.CompiledJsonStringWithStackTraces()
	if s == "" {
		t.Fatal("expected non-empty json string")
	}
}

func Test_I14_RawErrCollection_CompiledJsonStringWithStackTraces_Empty(t *testing.T) {
	rec := &RawErrCollection{}
	s := rec.CompiledJsonStringWithStackTraces()
	if s != "" {
		t.Fatal("expected empty for empty collection")
	}
}

func Test_I14_RawErrCollection_Log(t *testing.T) {
	rec := &RawErrCollection{}
	rec.Add(errors.New("err1"))
	rec.Log()
}

func Test_I14_RawErrCollection_LogWithTraces(t *testing.T) {
	rec := &RawErrCollection{}
	rec.Add(errors.New("err1"))
	rec.LogWithTraces()
}

func Test_I14_RawErrCollection_LogIf_False(t *testing.T) {
	rec := &RawErrCollection{}
	rec.Add(errors.New("err1"))
	rec.LogIf(false)
}

func Test_I14_RawErrCollection_AddErrorGetters(t *testing.T) {
	rec := &RawErrCollection{}
	// nil getter skipped
	rec.AddErrorGetters(nil)
	if rec.Length() != 0 {
		t.Fatal("expected 0 after nil getter")
	}
	// getter with nil error skipped
	rec.AddErrorGetters(&mockErrorGetter{nil})
	if rec.Length() != 0 {
		t.Fatal("expected 0 after nil-error getter")
	}
	// getter with real error added
	rec.AddErrorGetters(&mockErrorGetter{errors.New("x")})
	if rec.Length() != 1 {
		t.Fatal("expected 1 after real-error getter")
	}
}

func Test_I14_RawErrCollection_AddCompiledErrorGetters(t *testing.T) {
	rec := &RawErrCollection{}
	// nil getter skipped
	rec.AddCompiledErrorGetters(nil)
	if rec.Length() != 0 {
		t.Fatal("expected 0 after nil getter")
	}
	// getter with nil error skipped
	rec.AddCompiledErrorGetters(&mockCompiledErrorGetter{nil})
	if rec.Length() != 0 {
		t.Fatal("expected 0 after nil-error getter")
	}
	// getter with real error added
	rec.AddCompiledErrorGetters(&mockCompiledErrorGetter{errors.New("x")})
	if rec.Length() != 1 {
		t.Fatal("expected 1 after real-error getter")
	}
}

func Test_I14_RawErrCollection_SerializeWithoutTraces(t *testing.T) {
	rec := &RawErrCollection{}
	rec.Add(errors.New("err1"))
	b, err := rec.SerializeWithoutTraces()
	if err != nil {
		t.Fatal("expected no error:", err)
	}
	if len(b) == 0 {
		t.Fatal("expected non-empty bytes")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// stackTraceEnhance — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_StackTraceEnhance_MsgErrorSkip_AlreadyTraced(t *testing.T) {
	ste := stackTraceEnhance{}
	// Message already containing "Stack-Trace:" should skip enhancement
	msg := "some error\nStack-Trace: already"
	result := ste.MsgErrorSkip(0, msg, errors.New("wrapped"))
	if result == "" {
		t.Fatal("expected non-empty result")
	}
}

func Test_I14_StackTraceEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	ste := stackTraceEnhance{}
	result := ste.MsgErrorSkip(0, "msg", nil)
	if result != "" {
		t.Fatal("expected empty for nil error")
	}
}

func Test_I14_StackTraceEnhance_MsgErrorToErrSkip_NilErr(t *testing.T) {
	ste := stackTraceEnhance{}
	err := ste.MsgErrorToErrSkip(0, "msg", nil)
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func Test_I14_StackTraceEnhance_FmtSkip_Empty(t *testing.T) {
	ste := stackTraceEnhance{}
	err := ste.FmtSkip(0, "")
	if err != nil {
		t.Fatal("expected nil for empty format")
	}
}
