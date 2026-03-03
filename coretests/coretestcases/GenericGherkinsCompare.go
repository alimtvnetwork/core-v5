package coretestcases

import (
	"fmt"
	"strings"
)

// CompareWith performs a structural comparison of two GenericGherkins instances.
//
// Returns isEqual=true if all fields match, otherwise returns a diff string
// describing the first mismatched field.
func (it *GenericGherkins[TInput, TExpect]) CompareWith(
	other *GenericGherkins[TInput, TExpect],
) (isEqual bool, diff string) {
	if it == nil && other == nil {
		return true, ""
	}

	if it == nil || other == nil {
		return false, "one side is nil"
	}

	var diffs []string

	if it.Title != other.Title {
		diffs = append(diffs, fmt.Sprintf("Title: %q != %q", it.Title, other.Title))
	}

	if it.Feature != other.Feature {
		diffs = append(diffs, fmt.Sprintf("Feature: %q != %q", it.Feature, other.Feature))
	}

	if it.Given != other.Given {
		diffs = append(diffs, fmt.Sprintf("Given: %q != %q", it.Given, other.Given))
	}

	if it.When != other.When {
		diffs = append(diffs, fmt.Sprintf("When: %q != %q", it.When, other.When))
	}

	if it.Then != other.Then {
		diffs = append(diffs, fmt.Sprintf("Then: %q != %q", it.Then, other.Then))
	}

	inputStr := fmt.Sprintf("%v", it.Input)
	otherInputStr := fmt.Sprintf("%v", other.Input)
	if inputStr != otherInputStr {
		diffs = append(diffs, fmt.Sprintf("Input: %v != %v", it.Input, other.Input))
	}

	expectedStr := fmt.Sprintf("%v", it.Expected)
	otherExpectedStr := fmt.Sprintf("%v", other.Expected)
	if expectedStr != otherExpectedStr {
		diffs = append(diffs, fmt.Sprintf("Expected: %v != %v", it.Expected, other.Expected))
	}

	actualStr := fmt.Sprintf("%v", it.Actual)
	otherActualStr := fmt.Sprintf("%v", other.Actual)
	if actualStr != otherActualStr {
		diffs = append(diffs, fmt.Sprintf("Actual: %v != %v", it.Actual, other.Actual))
	}

	if it.IsMatching != other.IsMatching {
		diffs = append(diffs, fmt.Sprintf("IsMatching: %v != %v", it.IsMatching, other.IsMatching))
	}

	if len(diffs) == 0 {
		return true, ""
	}

	return false, strings.Join(diffs, "; ")
}
