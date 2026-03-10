package coretestcases

import (
	"log/slog"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/errcore"
)

// ExpectedAsMap retrieves ExpectedInput as args.Map.
//
// Panics if ExpectedInput is not args.Map.
func (it CaseV1) ExpectedAsMap() args.Map {
	m, ok := it.ExpectedInput.(args.Map)

	if !ok {
		panic("ExpectedInput is not args.Map")
	}

	return m
}

// ShouldBeEqualMap compares actual args.Map against ExpectedInput args.Map.
//
// Both maps are compiled to sorted "key : value" string lines using
// CompileToStrings(), then compared using the standard ShouldBeEqual
// assertion pipeline.
//
// On mismatch, a copy-pasteable Go literal block of the expected map
// is printed via slog for easy test case correction.
//
// Example:
//
//	// In _testcases.go:
//	ExpectedInput: args.Map{
//	    "value":   5,
//	    "isZero":  false,
//	    "isValid": true,
//	}
//
//	// In _test.go:
//	actual := args.Map{
//	    "value":   v.ValueInt(),
//	    "isZero":  v.IsZero(),
//	    "isValid": v.IsValid(),
//	}
//	testCase.ShouldBeEqualMap(t, caseIndex, actual)
func (it CaseV1) ShouldBeEqualMap(
	t *testing.T,
	caseIndex int,
	actual args.Map,
) {
	t.Helper()

	expectedMap := it.ExpectedAsMap()
	actualLines := actual.CompileToStrings()
	expectedLines := expectedMap.CompileToStrings()

	// Check for mismatch and print copy-pasteable expected output
	if errcore.HasAnyMismatchOnLines(actualLines, expectedLines) {
		slog.Warn("copy-pasteable expected (from actual):",
			"caseIndex", caseIndex,
			"title", it.Title,
			"actualGoLiteral", "\n"+actual.GoLiteralString(),
		)
	}

	it.ExpectedInput = expectedLines

	it.ShouldBe(
		t,
		caseIndex,
		stringcompareas.Equal,
		actualLines...,
	)
}

// ShouldBeEqualMapFirst asserts using ShouldBeEqualMap with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeEqualMapFirst(
	t *testing.T,
	actual args.Map,
) {
	t.Helper()

	it.ShouldBeEqualMap(
		t,
		0,
		actual,
	)
}
