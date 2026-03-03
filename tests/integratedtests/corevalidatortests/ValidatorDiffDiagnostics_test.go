package corevalidatortests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// LineValidator — enhanced tests with diff diagnostics
// ==========================================

func Test_LineValidator_VerifyError_LineAndTextMismatch_PrintsDiff(t *testing.T) {
	// Arrange: both line number and text are wrong
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{
			Search:    "expected-text",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "LineValidator line+text mismatch",
		IsCaseSensitive: true,
	}

	// Act
	err := v.VerifyError(params, 10, "actual-wrong-text")

	// Assert
	if err == nil {
		t.Fatal("expected error for line+text mismatch")
	}

	// Print diagnostic info with line numbers
	fmt.Printf("\n--- LineValidator Diagnostic (Case %d: %s) ---\n", 0, params.Header)
	fmt.Printf("  Expected line number: %d, Got: %d\n", 5, 10)
	fmt.Printf("  Expected text: %q\n", "expected-text")
	fmt.Printf("  Actual text:   %q\n", "actual-wrong-text")
	fmt.Printf("  Error: %v\n", err)
	fmt.Println("--- End Diagnostic ---")
}

func Test_LineValidator_AllVerifyError_MultipleContents_PrintsDiff(t *testing.T) {
	// Arrange: 5 items, 3 fail
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "target",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "AllVerify multi-content diff",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "target", LineNumber: 0},
		{Text: "wrong-1", LineNumber: 1},
		{Text: "target", LineNumber: 2},
		{Text: "wrong-2", LineNumber: 3},
		{Text: "wrong-3", LineNumber: 4},
	}

	// Act
	err := v.AllVerifyError(params, items...)

	// Assert
	if err == nil {
		t.Fatal("expected error for 3 mismatches")
	}

	// Enhanced diff printing with line numbers
	fmt.Printf("\n=== LineValidator AllVerifyError Diff (Case %d) ===\n", 0)
	for _, item := range items {
		match := "OK"
		if item.Text != "target" {
			match = "MISMATCH"
		}
		fmt.Printf("  Line %3d [%s]: actual=%q, expected=%q\n",
			item.LineNumber, match, item.Text, "target")
	}
	fmt.Printf("  Error details: %v\n", err)
	fmt.Println("=== End Diff ===")

	errMsg := err.Error()
	for _, expected := range []string{"wrong-1", "wrong-2", "wrong-3"} {
		if !strings.Contains(errMsg, expected) {
			t.Errorf("error should contain '%s', got:\n%s", expected, errMsg)
		}
	}
}

func Test_LineValidator_VerifyMany_CollectAll_PrintsDiff(t *testing.T) {
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "VerifyMany collectAll diff",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "is ok here", LineNumber: 0},
		{Text: "no match", LineNumber: 1},
		{Text: "also ok", LineNumber: 2},
		{Text: "missing", LineNumber: 3},
	}

	// isContinueOnError = true -> collect all
	err := v.VerifyMany(true, params, items...)
	if err == nil {
		t.Fatal("expected errors for lines 1 and 3")
	}

	// Print line-by-line result with line numbers
	fmt.Printf("\n=== VerifyMany (collectAll) Diff ===\n")
	for _, item := range items {
		containsOk := strings.Contains(item.Text, "ok")
		status := "OK"
		if !containsOk {
			status = "FAIL"
		}
		fmt.Printf("  Line %3d [%s]: %q (searching for Contains 'ok')\n",
			item.LineNumber, status, item.Text)
	}
	fmt.Println("=== End ===")

	errMsg := err.Error()
	if !strings.Contains(errMsg, "no match") {
		t.Errorf("should mention 'no match': %s", errMsg)
	}
}

// ==========================================
// LineValidator — VerifyFirstError with line number specifics
// ==========================================

func Test_LineValidator_VerifyFirstError_SpecificLineNumber_PrintsDiff(t *testing.T) {
	// Arrange: validator expects line 2 specifically
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 2},
		TextValidator: corevalidator.TextValidator{
			Search:    "expected",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "VerifyFirst specific line number",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "expected", LineNumber: 0}, // wrong line number
		{Text: "expected", LineNumber: 1}, // wrong line number
		{Text: "expected", LineNumber: 2}, // correct!
	}

	err := v.VerifyFirstError(params, items...)
	if err == nil {
		// All items are checked; line 0 and 1 don't match line number 2
		// so there should be an error on the first item
		fmt.Printf("\n  Line number check: validator expects line 2\n")
		for _, item := range items {
			fmt.Printf("  Line %d: text=%q, lineMatch=%v\n",
				item.LineNumber, item.Text, item.LineNumber == 2)
		}
	}
	// The first item has LineNumber=0 but validator expects 2, so error
	if err == nil {
		t.Fatal("line 0 doesn't match expected line 2, should error")
	}

	fmt.Printf("\n--- Line number mismatch diagnostic ---\n")
	fmt.Printf("  Validator expects line: 2\n")
	fmt.Printf("  First content line: 0 (mismatch!)\n")
	fmt.Printf("  Error: %v\n", err)
}

// ==========================================
// LinesValidators — enhanced multi-validator diff
// ==========================================

func Test_LinesValidators_AllVerifyError_MultiValidator_PrintsDiff(t *testing.T) {
	// Arrange: 2 validators, each checking different search terms
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "alpha",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "beta",
			SearchAs:  stringcompareas.Contains,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})

	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "Multi-validator AllVerifyError",
		IsCaseSensitive: true,
	}

	items := []corestr.TextWithLineNumber{
		{Text: "contains alpha here", LineNumber: 0},
		{Text: "no match at all", LineNumber: 1},
		{Text: "alpha and beta together", LineNumber: 2},
	}

	err := lv.AllVerifyError(params, items...)
	if err == nil {
		t.Fatal("expected errors: 'beta' not in line 0, 'alpha' not in line 1")
	}

	// Enhanced diagnostics
	fmt.Printf("\n=== LinesValidators Multi-Validator Diff ===\n")
	searches := []string{"alpha", "beta"}
	for si, search := range searches {
		fmt.Printf("  Validator %d (Contains '%s'):\n", si, search)
		for _, item := range items {
			match := strings.Contains(item.Text, search)
			status := "OK"
			if !match {
				status = "FAIL"
			}
			fmt.Printf("    Line %3d [%s]: %q\n", item.LineNumber, status, item.Text)
		}
	}
	fmt.Println("=== End ===")
}

func Test_LinesValidators_IsMatchText_Multiple_PrintsDiff(t *testing.T) {
	lv := corevalidator.NewLinesValidators(3)
	lv.Adds(
		corevalidator.LineValidator{
			LineNumber: corevalidator.LineNumber{LineNumber: -1},
			TextValidator: corevalidator.TextValidator{
				Search:    "hello",
				SearchAs:  stringcompareas.Contains,
				Condition: corevalidator.DefaultDisabledCoreCondition,
			},
		},
		corevalidator.LineValidator{
			LineNumber: corevalidator.LineNumber{LineNumber: -1},
			TextValidator: corevalidator.TextValidator{
				Search:    "world",
				SearchAs:  stringcompareas.Contains,
				Condition: corevalidator.DefaultDisabledCoreCondition,
			},
		},
	)

	text := "hello universe"
	result := lv.IsMatchText(text, true)

	if result {
		t.Error("'world' is not in 'hello universe', should return false")
	}

	// Print which validators matched
	fmt.Printf("\n--- IsMatchText Diagnostic ---\n")
	fmt.Printf("  Text: %q\n", text)
	fmt.Printf("  Validator 'hello' Contains: %v\n", strings.Contains(text, "hello"))
	fmt.Printf("  Validator 'world' Contains: %v\n", strings.Contains(text, "world"))
	fmt.Println("--- End ---")
}

// ==========================================
// errcore.ErrorToLinesLineDiff tests
// ==========================================

func Test_ErrorToLinesLineDiff_NilError(t *testing.T) {
	expected := []string{"line1", "line2"}
	result := errcore.ErrorToLinesLineDiff(0, "nil error test", nil, expected)

	if !strings.Contains(result, "MISSING EXPECTED") {
		t.Errorf("nil error vs expected lines should show missing, got:\n%s", result)
	}

	fmt.Print(result)
}

func Test_ErrorToLinesLineDiff_WithError(t *testing.T) {
	err := fmt.Errorf("error line 1\nerror line 2\nerror line 3")
	expected := []string{"error line 1", "error line 2", "DIFFERENT"}

	result := errcore.ErrorToLinesLineDiff(0, "error diff test", err, expected)

	if !strings.Contains(result, "MISMATCH") {
		t.Errorf("line 2 should be mismatch, got:\n%s", result)
	}
	if !strings.Contains(result, "Line") {
		t.Error("should contain line number labels")
	}

	fmt.Print(result)
}
