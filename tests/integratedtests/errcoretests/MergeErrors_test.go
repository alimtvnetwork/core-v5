package errcoretests

import (
	"errors"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// SliceToError
// ==========================================

func Test_SliceToError_Nil(t *testing.T) {
	err := errcore.SliceToError(nil)
	if err != nil {
		t.Error("nil slice should return nil")
	}
}

func Test_SliceToError_Empty(t *testing.T) {
	err := errcore.SliceToError([]string{})
	if err != nil {
		t.Error("empty slice should return nil")
	}
}

func Test_SliceToError_Single(t *testing.T) {
	err := errcore.SliceToError([]string{"error one"})
	if err == nil {
		t.Fatal("should return error")
	}
	if err.Error() != "error one" {
		t.Errorf("expected 'error one', got '%s'", err.Error())
	}
}

func Test_SliceToError_Multiple(t *testing.T) {
	err := errcore.SliceToError([]string{"err1", "err2", "err3"})
	if err == nil {
		t.Fatal("should return error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "err1") || !strings.Contains(msg, "err2") || !strings.Contains(msg, "err3") {
		t.Errorf("should contain all errors, got '%s'", msg)
	}
	if strings.Count(msg, "\n") != 2 {
		t.Errorf("expected 2 newlines joining 3 errors, got %d", strings.Count(msg, "\n"))
	}
}

// ==========================================
// SliceToErrorPtr
// ==========================================

func Test_SliceToErrorPtr_Nil(t *testing.T) {
	err := errcore.SliceToErrorPtr(nil)
	if err != nil {
		t.Error("nil slice should return nil")
	}
}

func Test_SliceToErrorPtr_Empty(t *testing.T) {
	err := errcore.SliceToErrorPtr([]string{})
	if err != nil {
		t.Error("empty slice should return nil")
	}
}

func Test_SliceToErrorPtr_Single(t *testing.T) {
	err := errcore.SliceToErrorPtr([]string{"one"})
	if err == nil || err.Error() != "one" {
		t.Errorf("expected 'one', got %v", err)
	}
}

// ==========================================
// MergeErrors
// ==========================================

func Test_MergeErrors_NoArgs(t *testing.T) {
	err := errcore.MergeErrors()
	if err != nil {
		t.Error("no args should return nil")
	}
}

func Test_MergeErrors_AllNil(t *testing.T) {
	err := errcore.MergeErrors(nil, nil, nil)
	if err != nil {
		t.Error("all nil errors should return nil")
	}
}

func Test_MergeErrors_SingleError(t *testing.T) {
	err := errcore.MergeErrors(errors.New("fail"))
	if err == nil {
		t.Fatal("should return error")
	}
	if err.Error() != "fail" {
		t.Errorf("expected 'fail', got '%s'", err.Error())
	}
}

func Test_MergeErrors_MultipleErrors(t *testing.T) {
	err := errcore.MergeErrors(
		errors.New("a"),
		errors.New("b"),
		errors.New("c"),
	)
	if err == nil {
		t.Fatal("should return error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "a") || !strings.Contains(msg, "b") || !strings.Contains(msg, "c") {
		t.Errorf("should contain all errors, got '%s'", msg)
	}
}

func Test_MergeErrors_MixedNilAndErrors(t *testing.T) {
	err := errcore.MergeErrors(
		nil,
		errors.New("real"),
		nil,
		errors.New("also real"),
		nil,
	)
	if err == nil {
		t.Fatal("should return error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "real") || !strings.Contains(msg, "also real") {
		t.Errorf("should contain non-nil errors, got '%s'", msg)
	}
	// Should NOT contain nil representations
	if strings.Contains(msg, "<nil>") {
		t.Error("should skip nil errors")
	}
}

func Test_MergeErrors_SingleNil(t *testing.T) {
	err := errcore.MergeErrors(nil)
	if err != nil {
		t.Error("single nil should return nil")
	}
}

// ==========================================
// SliceErrorsToStrings
// ==========================================

func Test_SliceErrorsToStrings_Nil(t *testing.T) {
	result := errcore.SliceErrorsToStrings()
	if len(result) != 0 {
		t.Errorf("nil should return empty slice, got %d", len(result))
	}
}

func Test_SliceErrorsToStrings_AllNil(t *testing.T) {
	result := errcore.SliceErrorsToStrings(nil, nil)
	if len(result) != 0 {
		t.Errorf("all nil should return empty, got %d", len(result))
	}
}

func Test_SliceErrorsToStrings_Mixed(t *testing.T) {
	result := errcore.SliceErrorsToStrings(
		errors.New("a"),
		nil,
		errors.New("b"),
	)
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
	if result[0] != "a" || result[1] != "b" {
		t.Errorf("expected [a, b], got %v", result)
	}
}

// ==========================================
// MergeErrorsToString
// ==========================================

func Test_MergeErrorsToString_Nil(t *testing.T) {
	result := errcore.MergeErrorsToString(", ")
	if result != "" {
		t.Errorf("nil should return empty, got '%s'", result)
	}
}

func Test_MergeErrorsToString_WithJoiner(t *testing.T) {
	result := errcore.MergeErrorsToString(" | ",
		errors.New("x"),
		errors.New("y"),
	)
	if result != "x | y" {
		t.Errorf("expected 'x | y', got '%s'", result)
	}
}

func Test_MergeErrorsToString_SkipsNil(t *testing.T) {
	result := errcore.MergeErrorsToString(", ",
		nil,
		errors.New("only"),
		nil,
	)
	if result != "only" {
		t.Errorf("expected 'only', got '%s'", result)
	}
}

// ==========================================
// MergeErrorsToStringDefault
// ==========================================

func Test_MergeErrorsToStringDefault_Nil(t *testing.T) {
	result := errcore.MergeErrorsToStringDefault()
	if result != "" {
		t.Errorf("nil should return empty, got '%s'", result)
	}
}

func Test_MergeErrorsToStringDefault_Multiple(t *testing.T) {
	result := errcore.MergeErrorsToStringDefault(
		errors.New("a"),
		errors.New("b"),
	)
	if result != "a b" {
		t.Errorf("expected 'a b' (space-joined), got '%s'", result)
	}
}
