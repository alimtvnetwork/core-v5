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
	if !errcore.LineDiffHasMismatch(actLines, expected) {
		return
	}

	it.printHeader()
	it.printLeftRight()
	it.printRawMap()
	it.printClone()
	it.printError()
	errcore.PrintLineDiff(it.CaseIndex, it.Title, actLines, expected)
	it.printFooter()
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

	it.printHeader()
	it.printLeftRight()
	it.printRawMap()
	it.printJsonDiff()
	it.printFooter()
}

func (it MapDiffDiagnostics) printHeader() {
	fmt.Printf(
		"\n=== MapDiff (Case %d: %s) ===\n",
		it.CaseIndex,
		it.Title,
	)
}

func (it MapDiffDiagnostics) printFooter() {
	fmt.Println("=== End ===")
}

func (it MapDiffDiagnostics) printLeftRight() {
	if it.Left == nil {
		fmt.Println("  Left:  <nil>")
	} else {
		fmt.Printf("  Left:  %s\n", it.Left.String())
		fmt.Printf("  Left keys:  %v\n", it.Left.AllKeys())
	}

	if it.Right == nil && it.RawMap == nil {
		fmt.Println("  Right: <nil>")
	} else if it.Right != nil {
		fmt.Printf("  Right: %s\n", it.Right.String())
		fmt.Printf("  Right keys: %v\n", it.Right.AllKeys())
	}
}

func (it MapDiffDiagnostics) printRawMap() {
	if it.RawMap != nil {
		fmt.Printf("  RawMap: %v\n", it.RawMap)
	}
}

func (it MapDiffDiagnostics) printJsonDiff() {
	if it.Left == nil {
		return
	}

	var rightItems map[string]any

	if it.Right != nil {
		rightItems = it.Right.Items
	} else if it.RawMap != nil {
		rightItems = it.RawMap
	}

	if rightItems == nil {
		return
	}

	diffMsg := it.Left.DiffJsonMessage(true, rightItems)
	if len(diffMsg) > 0 {
		fmt.Printf("  DiffJson: %s\n", diffMsg)
	} else {
		fmt.Println("  DiffJson: <no differences>")
	}
}

func (it MapDiffDiagnostics) printClone() {
	if it.Clone == nil {
		return
	}

	fmt.Printf("  Clone: %s\n", it.Clone.String())
}

func (it MapDiffDiagnostics) printError() {
	if it.Error == nil {
		return
	}

	fmt.Printf("  Error: %v\n", it.Error)
}
