package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
)

// ==========================================
// LineNumber.HasLineNumber
// ==========================================

func Test_LineNumber_HasLineNumber_Positive(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 5}
	if !ln.HasLineNumber() {
		t.Error("LineNumber 5 should have line number")
	}
}

func Test_LineNumber_HasLineNumber_Zero(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 0}
	if !ln.HasLineNumber() {
		t.Error("LineNumber 0 should have line number (>-1)")
	}
}

func Test_LineNumber_HasLineNumber_Invalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	if ln.HasLineNumber() {
		t.Error("LineNumber -1 should not have line number")
	}
}

func Test_LineNumber_HasLineNumber_NegativeTwo(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -2}
	if ln.HasLineNumber() {
		t.Error("LineNumber -2 should not have line number")
	}
}

// ==========================================
// LineNumber.IsMatch
// ==========================================

func Test_LineNumber_IsMatch_SameNumber(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 3}
	if !ln.IsMatch(3) {
		t.Error("same line number should match")
	}
}

func Test_LineNumber_IsMatch_DifferentNumber(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 3}
	if ln.IsMatch(5) {
		t.Error("different line number should not match")
	}
}

func Test_LineNumber_IsMatch_InputInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 3}
	if !ln.IsMatch(-1) {
		t.Error("input -1 should always match (skip check)")
	}
}

func Test_LineNumber_IsMatch_ValidatorInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	if !ln.IsMatch(5) {
		t.Error("validator -1 should always match (skip check)")
	}
}

func Test_LineNumber_IsMatch_BothInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	if !ln.IsMatch(-1) {
		t.Error("both -1 should match")
	}
}

func Test_LineNumber_IsMatch_Zero(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 0}
	if !ln.IsMatch(0) {
		t.Error("both 0 should match")
	}
}

// ==========================================
// LineNumber.VerifyError
// ==========================================

func Test_LineNumber_VerifyError_Match(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 2}
	err := ln.VerifyError(2)
	if err != nil {
		t.Errorf("matching line numbers should not error, got: %v", err)
	}
}

func Test_LineNumber_VerifyError_Mismatch(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: 2}
	err := ln.VerifyError(5)
	if err == nil {
		t.Error("mismatched line numbers should return error")
	}
}

func Test_LineNumber_VerifyError_SkipOnInvalid(t *testing.T) {
	ln := corevalidator.LineNumber{LineNumber: -1}
	err := ln.VerifyError(5)
	if err != nil {
		t.Errorf("validator -1 should skip check, got: %v", err)
	}
}
