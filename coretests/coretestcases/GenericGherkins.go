package coretestcases

import (
	"gitlab.com/auk-go/core/coretests/args"
)

// GenericGherkins is a typed test case representation using
// the Gherkin (Given/When/Then) pattern.
//
// Use this when you want compile-time type safety for Input/Expected
// fields instead of extracting values from args.Map at runtime.
//
// TInput  — type of the test input (e.g., string for regex pattern)
// TExpect — type of the expected result (e.g., bool for IsMatch)
type GenericGherkins[TInput, TExpect any] struct {
	// Title is the test case header / scenario name.
	Title string

	// Feature describes the feature being tested.
	Feature string

	// Given describes the precondition.
	Given string

	// When describes the scenario.
	When string

	// Then describes the expected outcome.
	Then string

	// Input is the typed input value for the test.
	Input TInput

	// Expected is the typed expected result.
	Expected TExpect

	// Actual is the typed actual result, set after the Act phase.
	Actual TExpect

	// IsMatching indicates whether a match is expected.
	// Use for validation/matching tests (e.g., regex, search).
	IsMatching bool

	// ExpectedLines holds the expected string output lines for assertion.
	// Used by ShouldBeEqualUsingExpected when comparing act lines.
	ExpectedLines []string

	// ExtraArgs provides optional dynamic key-value pairs
	// for overflow beyond the typed fields.
	ExtraArgs args.Map
}
