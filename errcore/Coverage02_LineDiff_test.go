package errcore

import (
	"errors"
	"testing"
)

func TestLineDiff_Equal(t *testing.T) {
	diffs := LineDiff([]string{"a", "b"}, []string{"a", "b"})
	if len(diffs) != 2 {
		t.Fatal("expected 2")
	}
	if diffs[0].Status != "  " {
		t.Fatal("expected match")
	}
}

func TestLineDiff_Mismatch(t *testing.T) {
	diffs := LineDiff([]string{"a"}, []string{"b"})
	if diffs[0].Status != "!!" {
		t.Fatal("expected mismatch")
	}
}

func TestLineDiff_ExtraActual(t *testing.T) {
	diffs := LineDiff([]string{"a", "b"}, []string{"a"})
	if diffs[1].Status != "+" {
		t.Fatal("expected extra actual")
	}
}

func TestLineDiff_MissingExpected(t *testing.T) {
	diffs := LineDiff([]string{"a"}, []string{"a", "b"})
	if diffs[1].Status != "-" {
		t.Fatal("expected missing expected")
	}
}

func TestLineDiffToString_Empty(t *testing.T) {
	s := LineDiffToString(0, "test", []string{}, []string{})
	if s != "" {
		t.Fatal("expected empty")
	}
}

func TestLineDiffToString_WithDiffs(t *testing.T) {
	s := LineDiffToString(0, "test", []string{"a"}, []string{"b"})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestLineDiffToString_AllMatch(t *testing.T) {
	s := LineDiffToString(0, "test", []string{"a"}, []string{"a"})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestPrintLineDiff_NoOutput(t *testing.T) {
	PrintLineDiff(0, "test", []string{}, []string{})
}

func TestPrintLineDiff_WithOutput(t *testing.T) {
	PrintLineDiff(0, "test", []string{"a"}, []string{"b"})
}

func TestHasAnyMismatchOnLines_Match(t *testing.T) {
	if HasAnyMismatchOnLines([]string{"a"}, []string{"a"}) {
		t.Fatal("should match")
	}
}

func TestHasAnyMismatchOnLines_DiffLen(t *testing.T) {
	if !HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"}) {
		t.Fatal("should differ")
	}
}

func TestHasAnyMismatchOnLines_DiffContent(t *testing.T) {
	if !HasAnyMismatchOnLines([]string{"a"}, []string{"b"}) {
		t.Fatal("should differ")
	}
}

func TestPrintLineDiffOnFail_NoMismatch(t *testing.T) {
	PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"a"})
}

func TestPrintLineDiffOnFail_WithMismatch(t *testing.T) {
	PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"b"})
}

func TestErrorToLinesLineDiff_NilErr(t *testing.T) {
	s := ErrorToLinesLineDiff(0, "test", nil, []string{"a"})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestErrorToLinesLineDiff_WithErr(t *testing.T) {
	s := ErrorToLinesLineDiff(0, "test", errors.New("a"), []string{"a"})
	_ = s
}

func TestPrintErrorLineDiff(t *testing.T) {
	PrintErrorLineDiff(0, "test", errors.New("a"), []string{"b"})
}

func TestSliceDiffSummary_Match(t *testing.T) {
	s := SliceDiffSummary([]string{"a"}, []string{"a"})
	if s != "all lines match" {
		t.Fatal("expected all match")
	}
}

func TestSliceDiffSummary_Mismatch(t *testing.T) {
	s := SliceDiffSummary([]string{"a"}, []string{"b"})
	if s == "all lines match" {
		t.Fatal("expected mismatch")
	}
}

func TestPrintDiffOnMismatch_NoMismatch(t *testing.T) {
	PrintDiffOnMismatch(0, "test", []string{"a"}, []string{"a"})
}

func TestPrintDiffOnMismatch_WithMismatch(t *testing.T) {
	PrintDiffOnMismatch(0, "test", []string{"a"}, []string{"b"}, "ctx1")
}

func TestMapMismatchError(t *testing.T) {
	s := MapMismatchError("TestFunc", 1, "title",
		[]string{`"a": 1,`}, []string{`"a": 2,`})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAssertDiffOnMismatch_NoMismatch(t *testing.T) {
	AssertDiffOnMismatch(t, 0, "test", []string{"a"}, []string{"a"})
}

func TestAssertErrorDiffOnMismatch_NilErr(t *testing.T) {
	AssertErrorDiffOnMismatch(t, 0, "test", nil, []string{})
}

func TestAssertErrorDiffOnMismatch_NoMismatch(t *testing.T) {
	AssertErrorDiffOnMismatch(t, 0, "test", errors.New("a"), []string{"a"})
}
