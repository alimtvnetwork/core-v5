package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: TripleFromSplit
// ==========================================

func Test_TripleFromSplit(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplit(inputStr, sep)
		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TripleFromSplitTrimmed
// ==========================================

func Test_TripleFromSplitTrimmed(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplitTrimmed(inputStr, sep)
		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TripleFromSplitN
// ==========================================

func Test_TripleFromSplitN(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitNTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplitN(inputStr, sep)
		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TripleFromSplitNTrimmed
// ==========================================

func Test_TripleFromSplitNTrimmed(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitNTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplitNTrimmed(inputStr, sep)
		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TripleFromSlice
// ==========================================

func Test_TripleFromSlice(t *testing.T) {
	for caseIndex, testCase := range tripleFromSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		parts, _ := input.GetAsStringSlice("parts")

		// Act
		triple := coregeneric.TripleFromSlice(parts)
		actLines := []string{
			triple.Left,
			triple.Middle,
			triple.Right,
			fmt.Sprintf("%v", triple.IsValid),
			triple.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
