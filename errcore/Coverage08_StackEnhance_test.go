package errcore

import (
	"errors"
	"testing"
)

func TestStackEnhance_Error_Nil(t *testing.T) {
	if StackEnhance.Error(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestStackEnhance_Error_WithErr(t *testing.T) {
	err := StackEnhance.Error(errors.New("e"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestStackEnhance_ErrorSkip_Nil(t *testing.T) {
	if StackEnhance.ErrorSkip(0, nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestStackEnhance_ErrorSkip_WithErr(t *testing.T) {
	err := StackEnhance.ErrorSkip(0, errors.New("e"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestStackEnhance_Msg_Empty(t *testing.T) {
	if StackEnhance.Msg("") != "" {
		t.Fatal("expected empty")
	}
}

func TestStackEnhance_Msg_WithMsg(t *testing.T) {
	s := StackEnhance.Msg("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackEnhance_MsgSkip_Empty(t *testing.T) {
	if StackEnhance.MsgSkip(0, "") != "" {
		t.Fatal("expected empty")
	}
}

func TestStackEnhance_MsgSkip_WithMsg(t *testing.T) {
	s := StackEnhance.MsgSkip(0, "hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackEnhance_MsgSkip_AlreadyHasStackTrace(t *testing.T) {
	s := StackEnhance.MsgSkip(0, "hello Stack Trace: existing")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackEnhance_MsgToErrSkip_Empty(t *testing.T) {
	if StackEnhance.MsgToErrSkip(0, "") != nil {
		t.Fatal("expected nil")
	}
}

func TestStackEnhance_MsgToErrSkip_WithMsg(t *testing.T) {
	err := StackEnhance.MsgToErrSkip(0, "hello")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestStackEnhance_FmtSkip_Empty(t *testing.T) {
	if StackEnhance.FmtSkip(0, "") != nil {
		t.Fatal("expected nil")
	}
}

func TestStackEnhance_FmtSkip_WithFmt(t *testing.T) {
	err := StackEnhance.FmtSkip(0, "hello %s", "world")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestStackEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	if StackEnhance.MsgErrorSkip(0, "msg", nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestStackEnhance_MsgErrorSkip_WithErr(t *testing.T) {
	s := StackEnhance.MsgErrorSkip(0, "msg", errors.New("e"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackEnhance_MsgErrorSkip_AlreadyHasStack(t *testing.T) {
	s := StackEnhance.MsgErrorSkip(0, "msg Stack Trace: existing", errors.New("e"))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestStackEnhance_MsgErrorToErrSkip_NilErr(t *testing.T) {
	if StackEnhance.MsgErrorToErrSkip(0, "msg", nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestStackEnhance_MsgErrorToErrSkip_WithErr(t *testing.T) {
	err := StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestCountStateChangeTracker(t *testing.T) {
	c := &RawErrCollection{}
	tracker := NewCountStateChangeTracker(c)
	if !tracker.IsSameState() {
		t.Fatal("should be same")
	}
	if !tracker.IsValid() {
		t.Fatal("should be valid")
	}
	if !tracker.IsSuccess() {
		t.Fatal("should be success")
	}
	if tracker.HasChanges() {
		t.Fatal("should not have changes")
	}
	if tracker.IsFailed() {
		t.Fatal("should not be failed")
	}
	if !tracker.IsSameStateUsingCount(0) {
		t.Fatal("should be same with 0")
	}

	c.Add(errors.New("e"))
	if tracker.IsSameState() {
		t.Fatal("should have changed")
	}
	if !tracker.HasChanges() {
		t.Fatal("should have changes")
	}
	if !tracker.IsFailed() {
		t.Fatal("should be failed")
	}
}
