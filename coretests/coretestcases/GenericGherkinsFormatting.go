package coretestcases

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

// String returns a Gherkins-formatted representation using index 0.
func (it *GenericGherkins[TInput, TExpect]) String() string {
	return it.ToString(constants.Zero)
}

// ToString returns a Gherkins-formatted string with the given test index.
func (it *GenericGherkins[TInput, TExpect]) ToString(testIndex int) string {
	return errcore.GherkinsString(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then,
	)
}

// GetWithExpectation returns a Gherkins string that includes
// the Actual and Expected values for diagnostic output.
func (it *GenericGherkins[TInput, TExpect]) GetWithExpectation(
	testIndex int,
) string {
	return errcore.GherkinsStringWithExpectation(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then,
		it.Actual,
		it.Expected,
	)
}

// GetMessageConditional returns the Gherkins string with or without
// expectation details based on the isExpectation flag.
func (it *GenericGherkins[TInput, TExpect]) GetMessageConditional(
	isExpectation bool,
	testIndex int,
) string {
	if isExpectation {
		return it.GetWithExpectation(testIndex)
	}

	return it.ToString(testIndex)
}

// FullString returns a verbose multi-line representation of all fields
// for debugging purposes.
func (it *GenericGherkins[TInput, TExpect]) FullString() string {
	if it == nil {
		return "<nil GenericGherkins>"
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Title:      %s\n", it.Title))
	sb.WriteString(fmt.Sprintf("Feature:    %s\n", it.Feature))
	sb.WriteString(fmt.Sprintf("Given:      %s\n", it.Given))
	sb.WriteString(fmt.Sprintf("When:       %s\n", it.When))
	sb.WriteString(fmt.Sprintf("Then:       %s\n", it.Then))
	sb.WriteString(fmt.Sprintf("Input:      %v\n", it.Input))
	sb.WriteString(fmt.Sprintf("Expected:   %v\n", it.Expected))
	sb.WriteString(fmt.Sprintf("Actual:     %v\n", it.Actual))
	sb.WriteString(fmt.Sprintf("IsMatching: %v\n", it.IsMatching))

	if len(it.ExtraArgs) > 0 {
		sb.WriteString(fmt.Sprintf("ExtraArgs:  %v\n", it.ExtraArgs))
	}

	return sb.String()
}
