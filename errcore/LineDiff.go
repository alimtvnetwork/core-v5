package errcore

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

// LineDiffResult holds a single line comparison result
type LineDiffResult struct {
	LineNumber int
	Status     string // "  " match, "!!" mismatch, "+" extra actual, "-" missing expected
	Actual     string
	Expected   string
}

// LineDiff compares two slices line-by-line and returns a diff report.
// Each line shows the line number and whether it matched or not.
func LineDiff(
	actual []string,
	expected []string,
) []LineDiffResult {
	maxLen := len(actual)
	if len(expected) > maxLen {
		maxLen = len(expected)
	}

	results := make([]LineDiffResult, 0, maxLen)

	for i := 0; i < maxLen; i++ {
		r := LineDiffResult{LineNumber: i}

		hasActual := i < len(actual)
		hasExpected := i < len(expected)

		switch {
		case hasActual && hasExpected:
			r.Actual = actual[i]
			r.Expected = expected[i]
			if actual[i] == expected[i] {
				r.Status = "  "
			} else {
				r.Status = "!!"
			}
		case hasActual && !hasExpected:
			r.Actual = actual[i]
			r.Expected = "<missing>"
			r.Status = "+"
		case !hasActual && hasExpected:
			r.Actual = "<missing>"
			r.Expected = expected[i]
			r.Status = "-"
		}

		results = append(results, r)
	}

	return results
}

// LineDiffToString formats a LineDiff result into a readable multi-line string.
// Only mismatched lines are shown with full detail; matched lines show a summary.
func LineDiffToString(
	caseIndex int,
	header string,
	actual []string,
	expected []string,
) string {
	diffs := LineDiff(actual, expected)

	if len(diffs) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(
		"\n=== Line-by-Line Diff (Case %d: %s) ===\n",
		caseIndex,
		header,
	))
	sb.WriteString(fmt.Sprintf(
		"    Actual lines: %d, Expected lines: %d\n",
		len(actual),
		len(expected),
	))

	mismatchCount := 0
	for _, d := range diffs {
		switch d.Status {
		case "  ":
			sb.WriteString(fmt.Sprintf(
				"  Line %3d [OK]: `%s`\n",
				d.LineNumber,
				d.Actual,
			))
		case "!!":
			mismatchCount++
			sb.WriteString(fmt.Sprintf(
				"  Line %3d [MISMATCH]:\n"+
					"      actual   : `%s`\n"+
					"      expected : `%s`\n",
				d.LineNumber,
				d.Actual,
				d.Expected,
			))
		case "+":
			mismatchCount++
			sb.WriteString(fmt.Sprintf(
				"  Line %3d [EXTRA ACTUAL]: `%s`\n",
				d.LineNumber,
				d.Actual,
			))
		case "-":
			mismatchCount++
			sb.WriteString(fmt.Sprintf(
				"  Line %3d [MISSING EXPECTED]: `%s`\n",
				d.LineNumber,
				d.Expected,
			))
		}
	}

	sb.WriteString(fmt.Sprintf(
		"=== Total: %d lines, %d mismatches ===\n",
		len(diffs),
		mismatchCount,
	))

	return sb.String()
}

// PrintLineDiff prints the line-by-line diff to stdout for test failure diagnostics.
func PrintLineDiff(
	caseIndex int,
	header string,
	actual []string,
	expected []string,
) {
	msg := LineDiffToString(caseIndex, header, actual, expected)
	if len(msg) > 0 {
		fmt.Print(msg)
	}
}

// HasAnyMismatchOnLines returns true if actual and expected differ in any way.
func HasAnyMismatchOnLines(
	actual []string,
	expected []string,
) bool {
	if len(actual) != len(expected) {
		return true
	}

	for i := range actual {
		if actual[i] != expected[i] {
			return true
		}
	}

	return false
}

// PrintLineDiffOnFail prints the diff only if there are mismatches.
func PrintLineDiffOnFail(
	caseIndex int,
	header string,
	actual []string,
	expected []string,
) {
	if HasAnyMismatchOnLines(actual, expected) {
		PrintLineDiff(caseIndex, header, actual, expected)
	}
}

// ErrorToLinesLineDiff takes an error, splits it into lines,
// compares against expected lines, and returns the diff string.
func ErrorToLinesLineDiff(
	caseIndex int,
	header string,
	err error,
	expectedLines []string,
) string {
	actual := ErrorToSplitLines(err)
	if err == nil {
		actual = []string{}
	}

	return LineDiffToString(
		caseIndex,
		header,
		actual,
		expectedLines,
	)
}

// PrintErrorLineDiff prints the error-vs-expected diff with line numbers.
func PrintErrorLineDiff(
	caseIndex int,
	header string,
	err error,
	expectedLines []string,
) {
	msg := ErrorToLinesLineDiff(caseIndex, header, err, expectedLines)
	if len(msg) > 0 {
		fmt.Print(msg)
	}
}

// SliceDiffSummary returns a compact summary of which lines differ.
func SliceDiffSummary(
	actual []string,
	expected []string,
) string {
	diffs := LineDiff(actual, expected)

	var mismatches []string
	for _, d := range diffs {
		if d.Status != "  " {
			mismatches = append(mismatches, fmt.Sprintf(
				"line %d [%s]",
				d.LineNumber,
				d.Status,
			))
		}
	}

	if len(mismatches) == 0 {
		return "all lines match"
	}

	return fmt.Sprintf(
		"%d mismatches: %s",
		len(mismatches),
		strings.Join(mismatches, constants.CommaSpace),
	)
}
