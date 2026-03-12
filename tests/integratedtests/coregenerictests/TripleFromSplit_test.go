package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Test: TripleFromSplit
// ==========================================

func Test_TripleFromSplit_Verification(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplit(inputStr, sep)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TripleFromSplitTrimmed
// ==========================================

func Test_TripleFromSplitTrimmed_Verification(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		sep, _ := input.GetAsString("sep")

		// Act
		triple := coregeneric.TripleFromSplitTrimmed(inputStr, sep)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: TripleFromSlice
// ==========================================

func Test_TripleFromSlice(t *testing.T) {
	for caseIndex, testCase := range tripleFromSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		parts, _ := input.GetAsStrings("parts")

		// Act
		triple := coregeneric.TripleFromSlice(parts)
		actual := args.Map{
			"left":    triple.Left,
			"middle":  triple.Middle,
			"right":   triple.Right,
			"isValid": triple.IsValid,
			"message": triple.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
