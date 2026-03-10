package coretestcases

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
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
// This allows test cases to define expectations with raw typed values
// (int, bool, etc.) instead of pre-formatted strings, making test data
// self-documenting and eliminating manual fmt.Sprintf calls in test bodies.
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

	actualLines := actual.CompileToStrings()
	it.ExpectedInput = it.ExpectedAsMap().CompileToStrings()

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
