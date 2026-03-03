package coredynamictests

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/errcore"
)

// MapDiffDiagnostics provides reusable diff-printing
// diagnostics for MapAnyItems test failures.
type MapDiffDiagnostics struct {
	CaseIndex int
	Title     string
	Left      *coredynamic.MapAnyItems
	Right     *coredynamic.MapAnyItems
	RawMap    map[string]any
	Error     error
	Clone     *coredynamic.MapAnyItems
}

// PrintIfMismatch prints map diff diagnostics
// only when actual lines differ from expected lines.
func (it MapDiffDiagnostics) PrintIfMismatch(
	actLines []string,
	expected []string,
) {
	errcore.PrintDiffOnMismatch(
		it.CaseIndex,
		it.Title,
		actLines,
		expected,
		it.contextLines()...,
	)
}

// PrintIfResultMismatch prints map diff diagnostics
// for single-result comparisons (e.g., IsEqual, IsEqualRaw).
func (it MapDiffDiagnostics) PrintIfResultMismatch(
	resultStr string,
	expected []string,
) {
	if len(expected) == 0 || resultStr == expected[0] {
		return
	}

	errcore.PrintDiffOnMismatch(
		it.CaseIndex,
		it.Title,
		[]string{resultStr},
		expected,
		it.contextLines()...,
	)
}

// contextLines builds the diagnostic context for map comparisons.
func (it MapDiffDiagnostics) contextLines() []string {
	var lines []string

	// Left map
	if it.Left == nil {
		lines = append(lines, "  Left:  <nil>")
	} else {
		lines = append(lines,
			fmt.Sprintf("  Left:  %s", it.Left.String()),
			fmt.Sprintf("  Left keys:  %v", it.Left.AllKeys()),
		)
	}

	// Right map
	if it.Right != nil {
		lines = append(lines,
			fmt.Sprintf("  Right: %s", it.Right.String()),
			fmt.Sprintf("  Right keys: %v", it.Right.AllKeys()),
		)
	} else if it.RawMap != nil {
		lines = append(lines, fmt.Sprintf("  RawMap: %v", it.RawMap))
	}

	// JSON diff
	lines = append(lines, it.jsonDiffLine()...)

	// Clone
	if it.Clone != nil {
		lines = append(lines, fmt.Sprintf("  Clone: %s", it.Clone.String()))
	}

	// Error
	if it.Error != nil {
		lines = append(lines, fmt.Sprintf("  Error: %v", it.Error))
	}

	return lines
}

func (it MapDiffDiagnostics) jsonDiffLine() []string {
	if it.Left == nil {
		return nil
	}

	var rightItems map[string]any

	if it.Right != nil {
		rightItems = it.Right.Items
	} else if it.RawMap != nil {
		rightItems = it.RawMap
	}

	if rightItems == nil {
		return nil
	}

	diffMsg := it.Left.DiffJsonMessage(true, rightItems)
	if len(diffMsg) > 0 {
		return []string{fmt.Sprintf("  DiffJson: %s", diffMsg)}
	}

	return []string{"  DiffJson: <no differences>"}
}
