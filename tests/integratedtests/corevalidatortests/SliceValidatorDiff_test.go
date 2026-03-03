package corevalidatortests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// SliceValidator — AllVerifyError with diff
// ==========================================

func Test_SliceValidator_AllVerifyError_MultiLineMismatch_WithDiff(t *testing.T) {
	// Arrange: 5 lines, 2 mismatches at lines 1 and 3
	actual := []string{"alpha", "bravo-wrong", "charlie", "delta-wrong", "echo"}
	expected := []string{"alpha", "bravo", "charlie", "delta", "echo"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Multi-line mismatch with diff output",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	// Act
	err := v.AllVerifyError(params)

	// Assert: must fail
	if err == nil {
		t.Fatal("expected error for multi-line mismatch")
	}

	// Print line-by-line diff for diagnostics
	errcore.PrintLineDiff(0, params.Header, actual, expected)

	errMsg := err.Error()
	if !strings.Contains(errMsg, "bravo") {
		t.Errorf("error should mention 'bravo' mismatch, got:\n%s", errMsg)
	}
	if !strings.Contains(errMsg, "delta") {
		t.Errorf("error should mention 'delta' mismatch, got:\n%s", errMsg)
	}
}

func Test_SliceValidator_AllVerifyError_ExtraActualLines_WithDiff(t *testing.T) {
	actual := []string{"line1", "line2", "line3", "extra-line"}
	expected := []string{"line1", "line2", "line3"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Extra actual lines diff",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	if err == nil {
		t.Fatal("expected error for length mismatch")
	}

	// Print diff showing extra line
	errcore.PrintLineDiff(0, params.Header, actual, expected)
	summary := errcore.SliceDiffSummary(actual, expected)
	t.Logf("Diff summary: %s", summary)
}

func Test_SliceValidator_AllVerifyError_MissingActualLines_WithDiff(t *testing.T) {
	actual := []string{"line1"}
	expected := []string{"line1", "line2", "line3"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Missing actual lines diff",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	if err == nil {
		t.Fatal("expected error for missing actual lines")
	}

	errcore.PrintLineDiff(0, params.Header, actual, expected)
}

// ==========================================
// SliceValidator — VerifyFirstError with diff
// ==========================================

func Test_SliceValidator_VerifyFirstError_StopsAtFirst_WithDiff(t *testing.T) {
	actual := []string{"a", "WRONG1", "WRONG2"}
	expected := []string{"a", "b", "c"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "VerifyFirst stops at first mismatch",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.VerifyFirstError(params)
	if err == nil {
		t.Fatal("expected error")
	}

	errcore.PrintLineDiff(0, params.Header, actual, expected)

	// VerifyFirst should mention line 1 mismatch
	errMsg := err.Error()
	if !strings.Contains(errMsg, "WRONG1") {
		t.Errorf("should mention first mismatch 'WRONG1', got:\n%s", errMsg)
	}
}

// ==========================================
// SliceValidator — AllVerifyErrorTestCase with diff
// ==========================================

func Test_SliceValidator_AllVerifyErrorTestCase_WithDiff(t *testing.T) {
	actual := []string{"hello", "world-different"}
	expected := []string{"hello", "world"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	err := v.AllVerifyErrorTestCase(0, "TestCase with diff", true)
	if err == nil {
		t.Fatal("expected error")
	}

	// Also print our enhanced diff
	errcore.PrintLineDiff(0, "TestCase with diff", actual, expected)
}

// ==========================================
// SliceValidator — Contains with multiple mismatches
// ==========================================

func Test_SliceValidator_AllVerifyError_Contains_MultiMismatch(t *testing.T) {
	actual := []string{
		"path/to/file.go:10",
		"some other text",
		"path/to/other.go:20",
	}
	expected := []string{
		"file.go",
		"expected-missing",
		"other.go",
	}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Contains,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Contains multi-mismatch",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	if err == nil {
		t.Fatal("expected error for line 1 mismatch")
	}

	errcore.PrintLineDiff(0, params.Header, actual, expected)

	errMsg := err.Error()
	if !strings.Contains(errMsg, "expected-missing") {
		t.Errorf("error should reference missing substring, got:\n%s", errMsg)
	}
}

// ==========================================
// SliceValidator — Trim + diff
// ==========================================

func Test_SliceValidator_AllVerifyError_Trim_WithDiff(t *testing.T) {
	actual := []string{"  hello  ", "  world  "}
	expected := []string{"hello", "universe"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultTrimCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "Trim with diff - line 1 mismatch",
		IsCaseSensitive: true,
	}

	err := v.AllVerifyError(params)
	if err == nil {
		t.Fatal("expected error: world != universe after trim")
	}

	errcore.PrintLineDiff(0, params.Header, actual, expected)
}

// ==========================================
// SliceValidator — Glob pattern with diff
// ==========================================

func Test_SliceValidator_AllVerifyError_Glob_WithDiff(t *testing.T) {
	actual := []string{
		"build-20260303/result.json",
		"build-20260303/output.txt",
		"build-20260303/data.csv",
	}
	expected := []string{
		"build-*/result.json",
		"build-*/output.txt",
		"build-*/WRONG.csv",
	}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Glob,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Glob pattern - line 2 mismatch",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	if err == nil {
		t.Fatal("expected error: data.csv doesn't match WRONG.csv glob")
	}

	errcore.PrintLineDiff(0, params.Header, actual, expected)
}

// ==========================================
// SliceValidator — AllVerifyErrorExceptLast with diff
// ==========================================

func Test_SliceValidator_AllVerifyErrorExceptLast_WithDiff(t *testing.T) {
	actual := []string{"a", "b", "INTENTIONALLY-DIFFERENT"}
	expected := []string{"a", "b", "c"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actual,
		ExpectedLines: expected,
	}

	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "ExceptLast should skip last line",
		IsCaseSensitive: true,
	}

	err := v.AllVerifyErrorExceptLast(params)
	if err != nil {
		errcore.PrintLineDiff(0, params.Header, actual, expected)
		t.Errorf("should pass when skipping last line: %v", err)
	}
}

// ==========================================
// SliceValidator — Dispose then verify
// ==========================================

func Test_SliceValidator_Dispose_ThenAllVerifyError(t *testing.T) {
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}

	v.Dispose()

	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.AllVerifyError(params)

	// After dispose, both are nil, so nil receiver-like behavior
	if err != nil {
		t.Errorf("disposed validator with nil lines should not error: %v", err)
	}
}

// ==========================================
// errcore.LineDiff utility direct tests
// ==========================================

func Test_LineDiff_BothEmpty(t *testing.T) {
	diffs := errcore.LineDiff([]string{}, []string{})
	if len(diffs) != 0 {
		t.Errorf("both empty should produce 0 diffs, got %d", len(diffs))
	}
}

func Test_LineDiff_ExactMatch(t *testing.T) {
	actual := []string{"a", "b", "c"}
	expected := []string{"a", "b", "c"}
	diffs := errcore.LineDiff(actual, expected)

	for i, d := range diffs {
		if d.Status != "  " {
			t.Errorf("line %d should match, got status %q", i, d.Status)
		}
		if d.LineNumber != i {
			t.Errorf("line number should be %d, got %d", i, d.LineNumber)
		}
	}
}

func Test_LineDiff_Mismatches(t *testing.T) {
	actual := []string{"a", "WRONG", "c"}
	expected := []string{"a", "b", "c"}
	diffs := errcore.LineDiff(actual, expected)

	if diffs[0].Status != "  " {
		t.Error("line 0 should match")
	}
	if diffs[1].Status != "!!" {
		t.Errorf("line 1 should be mismatch, got %q", diffs[1].Status)
	}
	if diffs[1].LineNumber != 1 {
		t.Errorf("mismatch line number should be 1, got %d", diffs[1].LineNumber)
	}
	if diffs[2].Status != "  " {
		t.Error("line 2 should match")
	}
}

func Test_LineDiff_ExtraActual(t *testing.T) {
	actual := []string{"a", "b", "extra"}
	expected := []string{"a", "b"}
	diffs := errcore.LineDiff(actual, expected)

	if len(diffs) != 3 {
		t.Fatalf("expected 3 diffs, got %d", len(diffs))
	}
	if diffs[2].Status != "+" {
		t.Errorf("extra line should have '+' status, got %q", diffs[2].Status)
	}
	if diffs[2].LineNumber != 2 {
		t.Errorf("extra line number should be 2, got %d", diffs[2].LineNumber)
	}
}

func Test_LineDiff_MissingActual(t *testing.T) {
	actual := []string{"a"}
	expected := []string{"a", "b", "c"}
	diffs := errcore.LineDiff(actual, expected)

	if len(diffs) != 3 {
		t.Fatalf("expected 3 diffs, got %d", len(diffs))
	}
	if diffs[1].Status != "-" {
		t.Errorf("missing line should have '-' status, got %q", diffs[1].Status)
	}
	if diffs[2].Status != "-" {
		t.Errorf("missing line should have '-' status, got %q", diffs[2].Status)
	}
}

func Test_LineDiffToString_ContainsLineNumbers(t *testing.T) {
	actual := []string{"a", "WRONG"}
	expected := []string{"a", "b"}

	result := errcore.LineDiffToString(0, "test header", actual, expected)

	if !strings.Contains(result, "Line") {
		t.Error("diff output should contain 'Line' labels")
	}
	if !strings.Contains(result, "MISMATCH") {
		t.Error("diff output should contain 'MISMATCH' for differing lines")
	}
	if !strings.Contains(result, "test header") {
		t.Error("diff output should contain the header")
	}
	if !strings.Contains(result, "Case 0") {
		t.Error("diff output should contain the case index")
	}

	// Print for visual inspection during test runs
	fmt.Print(result)
}

func Test_LineDiffHasMismatch_True(t *testing.T) {
	if !errcore.LineDiffHasMismatch([]string{"a"}, []string{"b"}) {
		t.Error("different content should be mismatch")
	}
}

func Test_LineDiffHasMismatch_DifferentLength(t *testing.T) {
	if !errcore.LineDiffHasMismatch([]string{"a"}, []string{"a", "b"}) {
		t.Error("different length should be mismatch")
	}
}

func Test_LineDiffHasMismatch_False(t *testing.T) {
	if errcore.LineDiffHasMismatch([]string{"a", "b"}, []string{"a", "b"}) {
		t.Error("same content should not be mismatch")
	}
}

func Test_SliceDiffSummary_AllMatch(t *testing.T) {
	result := errcore.SliceDiffSummary([]string{"a", "b"}, []string{"a", "b"})
	if result != "all lines match" {
		t.Errorf("expected 'all lines match', got %q", result)
	}
}

func Test_SliceDiffSummary_HasMismatches(t *testing.T) {
	result := errcore.SliceDiffSummary(
		[]string{"a", "WRONG", "c"},
		[]string{"a", "b", "c"},
	)
	if !strings.Contains(result, "1 mismatches") {
		t.Errorf("summary should show mismatch count, got %q", result)
	}
	if !strings.Contains(result, "line 1") {
		t.Errorf("summary should show line number, got %q", result)
	}
}
