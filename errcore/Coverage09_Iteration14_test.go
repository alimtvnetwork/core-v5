package errcore

import (
	"errors"
	"testing"
)

// ── Unexported stackTraceEnhance — must remain in source package ──

func Test_I14_StackTraceEnhance_MsgErrorSkip_AlreadyTraced(t *testing.T) {
	ste := stackTraceEnhance{}
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
