package coretestcases

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/results"
)

// CaseNilSafe is a test case structure for systematically testing
// nil-receiver safety of pointer receiver methods.
//
// It uses direct method references (Func field) instead of string names,
// providing compile-time safety — renaming a method causes a build error
// rather than a silent test failure.
//
// Expected is a results.ResultAny, ensuring structural alignment
// between the actual invocation result and the expected outcome.
// Assertion methods live on results.Result[T]; CaseNilSafe delegates
// to them via ShouldBeSafe for convenience.
//
// Usage:
//
//	testCases := []coretestcases.CaseNilSafe{
//	    {
//	        Title: "IsValid on nil receiver returns false",
//	        Func:  (*MyStruct).IsValid,
//	        Expected: results.ResultAny{
//	            Value:    "false",
//	            Panicked: false,
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

	// Expected holds the expected result as a typed Result struct.
	// The assertion auto-derives which fields to compare based on
	// which Expected fields are set to non-zero values.
	// Override with CompareFields for explicit control.
	Expected results.ResultAny

	// CompareFields optionally specifies which Result fields to compare.
	// When empty, fields are auto-derived from Expected:
	//   - "panicked"    — always
	//   - "value"       — if Expected.Value != nil
	//   - "hasError"    — if Expected.Error != nil
	//   - "returnCount" — if Expected.ReturnCount != 0
	CompareFields []string
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

// ShouldBeSafe is a convenience assertion that invokes with nil
// and delegates to Result.ShouldMatchResult for comparison.
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

	result.ShouldMatchResult(
		t,
		caseIndex,
		it.CaseTitle(),
		it.Expected,
		it.CompareFields...,
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
