package coretestcases

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/results"
)

// CaseNilSafe is a test case structure for systematically testing
// nil-receiver safety of pointer receiver methods.
//
// It uses direct method references (Func field) instead of string names,
// providing compile-time safety — renaming a method causes a build error
// rather than a silent test failure.
//
// Usage:
//
//	testCases := []coretestcases.CaseNilSafe{
//	    {
//	        Title: "IsValid on nil receiver returns false",
//	        Func:  (*MyStruct).IsValid,
//	        Expected: args.Map{
//	            "value":    "false",
//	            "panicked": false,
//	        },
//	    },
//	}
type CaseNilSafe struct {
	// Title is the test case header / scenario name.
	Title string

	// Func is the direct method reference.
	// Use method expressions: (*Type).Method
	Func any

	// Args holds optional input arguments for the method call.
	Args []any

	// Expected holds the expected results as a typed map.
	// Common keys: "value", "panicked", "isSafe", "hasError", "returnCount".
	Expected args.Map
}

// MethodName returns the reflected name of the Func reference.
func (it CaseNilSafe) MethodName() string {
	return results.MethodName(it.Func)
}

// CaseTitle returns the Title, falling back to MethodName if empty.
func (it CaseNilSafe) CaseTitle() string {
	if it.Title != "" {
		return it.Title
	}

	return it.MethodName()
}

// Invoke calls the method with the given receiver and Args,
// recovering from any panic. Returns a ResultAny.
func (it CaseNilSafe) Invoke(receiver any) results.ResultAny {
	return results.InvokeWithPanicRecovery(
		it.Func,
		receiver,
		it.Args...,
	)
}

// InvokeNil calls the method with a nil receiver.
// This is the primary use case for nil-safety testing.
func (it CaseNilSafe) InvokeNil() results.ResultAny {
	return it.Invoke(nil)
}

// ShouldBeEqualMap asserts that the actual result map matches Expected
// using the standard map-based assertion pattern.
func (it CaseNilSafe) ShouldBeEqualMap(
	t *testing.T,
	caseIndex int,
	actual args.Map,
) {
	t.Helper()

	title := it.CaseTitle()

	actLines := actual.CompileToStrings()
	expectedLines := it.Expected.CompileToStrings()

	assertDiffOnMismatch(
		t,
		caseIndex,
		title,
		actLines,
		expectedLines,
	)
}

// ShouldBeEqualMapFirst is a convenience for non-loop tests (caseIndex=0).
func (it CaseNilSafe) ShouldBeEqualMapFirst(
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

// ShouldBeSafe is a convenience assertion that invokes with nil
// and asserts using the Result's ToMap output.
//
// This is the most concise assertion for standard nil-safety tests:
//
//	tc.ShouldBeSafe(t, caseIndex)
func (it CaseNilSafe) ShouldBeSafe(
	t *testing.T,
	caseIndex int,
) {
	t.Helper()

	result := it.InvokeNil()
	actual := filterMapByExpected(result.ToMap(), it.Expected)

	it.ShouldBeEqualMap(
		t,
		caseIndex,
		actual,
	)
}

// ShouldBeSafeFirst is a convenience for non-loop tests (caseIndex=0).
func (it CaseNilSafe) ShouldBeSafeFirst(
	t *testing.T,
) {
	t.Helper()

	it.ShouldBeSafe(
		t,
		0,
	)
}

// filterMapByExpected returns a new args.Map containing only the keys
// present in the expected map. This allows Expected to be a subset of
// the full ToMap output — tests only assert what they care about.
func filterMapByExpected(actual args.Map, expected args.Map) args.Map {
	filtered := args.Map{}

	for key := range expected {
		if val, exists := actual[key]; exists {
			filtered[key] = val
		} else {
			filtered[key] = fmt.Sprintf("<missing key: %s>", key)
		}
	}

	return filtered
}
