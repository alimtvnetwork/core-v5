package errcore

import (
	"errors"
	"testing"
)

func TestCompiledError_NilErr(t *testing.T) {
	if CompiledError(nil, "msg") != nil {
		t.Fatal("expected nil")
	}
}

func TestCompiledError_EmptyMsg(t *testing.T) {
	err := errors.New("base")
	if CompiledError(err, "") != err {
		t.Fatal("expected same error")
	}
}

func TestCompiledError_WithMsg(t *testing.T) {
	err := errors.New("base")
	result := CompiledError(err, "context")
	if result == nil || result.Error() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCompiledErrorString_NilErr(t *testing.T) {
	if CompiledErrorString(nil, "msg") != "" {
		t.Fatal("expected empty")
	}
}

func TestCompiledErrorString_WithErr(t *testing.T) {
	s := CompiledErrorString(errors.New("base"), "ctx")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestJoinErrors(t *testing.T) {
	e1 := errors.New("a")
	e2 := errors.New("b")
	joined := JoinErrors(e1, e2)
	if joined == nil {
		t.Fatal("expected non-nil")
	}
	joined2 := JoinErrors()
	if joined2 != nil {
		t.Fatal("expected nil")
	}
}

func TestErrorWithRef_NilErr(t *testing.T) {
	if ErrorWithRef(nil, "ref") != "" {
		t.Fatal("expected empty")
	}
}

func TestErrorWithRef_NilRef(t *testing.T) {
	s := ErrorWithRef(errors.New("e"), nil)
	if s != "e" {
		t.Fatal("expected just error")
	}
}

func TestErrorWithRef_EmptyRef(t *testing.T) {
	s := ErrorWithRef(errors.New("e"), "")
	if s != "e" {
		t.Fatal("expected just error")
	}
}

func TestErrorWithRef_WithRef(t *testing.T) {
	s := ErrorWithRef(errors.New("e"), "ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestErrorWithRefToError_NilErr(t *testing.T) {
	if ErrorWithRefToError(nil, "ref") != nil {
		t.Fatal("expected nil")
	}
}

func TestErrorWithRefToError_WithErr(t *testing.T) {
	err := ErrorWithRefToError(errors.New("e"), "ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestErrorWithCompiledTraceRef_NilErr(t *testing.T) {
	if ErrorWithCompiledTraceRef(nil, "traces", "ref") != "" {
		t.Fatal("expected empty")
	}
}

func TestErrorWithCompiledTraceRef_EmptyTraces(t *testing.T) {
	s := ErrorWithCompiledTraceRef(errors.New("e"), "", "ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestErrorWithCompiledTraceRef_NilRef(t *testing.T) {
	s := ErrorWithCompiledTraceRef(errors.New("e"), "traces", nil)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestErrorWithCompiledTraceRef_All(t *testing.T) {
	s := ErrorWithCompiledTraceRef(errors.New("e"), "traces", "ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestErrorWithCompiledTraceRefToError_NilErr(t *testing.T) {
	if ErrorWithCompiledTraceRefToError(nil, "t", "r") != nil {
		t.Fatal("expected nil")
	}
}

func TestErrorWithCompiledTraceRefToError_WithErr(t *testing.T) {
	err := ErrorWithCompiledTraceRefToError(errors.New("e"), "t", "r")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestErrorWithTracesRefToError_NilErr(t *testing.T) {
	if ErrorWithTracesRefToError(nil, []string{"t"}, "r") != nil {
		t.Fatal("expected nil")
	}
}

func TestErrorWithTracesRefToError_EmptyTraces(t *testing.T) {
	err := ErrorWithTracesRefToError(errors.New("e"), []string{}, "r")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestErrorWithTracesRefToError_WithTraces(t *testing.T) {
	err := ErrorWithTracesRefToError(errors.New("e"), []string{"t1", "t2"}, "r")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestConcatMessageWithErr_NilErr(t *testing.T) {
	if ConcatMessageWithErr("msg", nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestConcatMessageWithErr_WithErr(t *testing.T) {
	err := ConcatMessageWithErr("msg", errors.New("e"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestConcatMessageWithErrWithStackTrace_NilErr(t *testing.T) {
	if ConcatMessageWithErrWithStackTrace("msg", nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestConcatMessageWithErrWithStackTrace_WithErr(t *testing.T) {
	err := ConcatMessageWithErrWithStackTrace("msg", errors.New("e"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestErrorToSplitLines_NilErr(t *testing.T) {
	lines := ErrorToSplitLines(nil)
	if len(lines) != 0 {
		t.Fatal("expected empty")
	}
}

func TestErrorToSplitLines_WithErr(t *testing.T) {
	lines := ErrorToSplitLines(errors.New("a\nb"))
	if len(lines) != 2 {
		t.Fatal("expected 2")
	}
}

func TestErrorToSplitNonEmptyLines(t *testing.T) {
	lines := ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))
	if len(lines) < 2 {
		t.Fatal("expected at least 2")
	}
}

func TestErrorToSplitNonEmptyLines_Nil(t *testing.T) {
	lines := ErrorToSplitNonEmptyLines(nil)
	_ = lines
}
