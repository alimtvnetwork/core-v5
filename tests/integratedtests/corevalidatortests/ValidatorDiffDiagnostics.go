package corevalidatortests

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/coredata/corestr"
)

// ValidatorDiffDiagnostics provides reusable diff-printing
// diagnostics for validator test failures.
type ValidatorDiffDiagnostics struct {
	CaseIndex int
	Header    string
	Error     error
}

// PrintLineMatchDiagnostics prints per-item match/mismatch status
// for a slice of TextWithLineNumber items checked against a search term.
//
// Each item is labeled OK or FAIL based on whether its Text satisfies
// the matchFn predicate.
func (it ValidatorDiffDiagnostics) PrintLineMatchDiagnostics(
	items []corestr.TextWithLineNumber,
	matchFn func(text string) bool,
	searchLabel string,
) {
	fmt.Printf("\n=== %s (Case %d) ===\n", it.Header, it.CaseIndex)
	for _, item := range items {
		status := "OK"
		if !matchFn(item.Text) {
			status = "FAIL"
		}
		fmt.Printf("  Line %3d [%s]: actual=%q, search=%q\n",
			item.LineNumber, status, item.Text, searchLabel)
	}
	it.printError()
	fmt.Println("=== End ===")
}

// PrintMultiValidatorDiagnostics prints per-validator, per-item match
// status for multi-validator scenarios (LinesValidators).
func (it ValidatorDiffDiagnostics) PrintMultiValidatorDiagnostics(
	items []corestr.TextWithLineNumber,
	searches []string,
) {
	fmt.Printf("\n=== %s (Case %d) ===\n", it.Header, it.CaseIndex)
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
	it.printError()
	fmt.Println("=== End ===")
}

// PrintLineNumberMismatch prints diagnostics for line-number
// and text mismatches in LineValidator scenarios.
func (it ValidatorDiffDiagnostics) PrintLineNumberMismatch(
	expectedLine int,
	actualLine int,
	expectedText string,
	actualText string,
) {
	fmt.Printf("\n=== %s (Case %d) ===\n", it.Header, it.CaseIndex)
	fmt.Printf("  Expected line: %d, Got: %d\n", expectedLine, actualLine)
	fmt.Printf("  Expected text: %q\n", expectedText)
	fmt.Printf("  Actual text:   %q\n", actualText)
	it.printError()
	fmt.Println("=== End ===")
}

// PrintTextMatchDiagnostics prints which search terms match
// against a single text value.
func (it ValidatorDiffDiagnostics) PrintTextMatchDiagnostics(
	text string,
	searchResults map[string]bool,
) {
	fmt.Printf("\n=== %s (Case %d) ===\n", it.Header, it.CaseIndex)
	fmt.Printf("  Text: %q\n", text)
	for search, matched := range searchResults {
		fmt.Printf("  Validator '%s' Contains: %v\n", search, matched)
	}
	fmt.Println("=== End ===")
}

func (it ValidatorDiffDiagnostics) printError() {
	if it.Error != nil {
		fmt.Printf("  Error: %v\n", it.Error)
	}
}
