package coregenerictests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: PairFromSplit
// ==========================================

func Test_PairFromSplit(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplit(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSplitTrimmed
// ==========================================

func Test_PairFromSplitTrimmed(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplitTrimmed(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSplitFull
// ==========================================

func Test_PairFromSplitFull(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitFullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplitFull(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSplitFullTrimmed
// ==========================================

func Test_PairFromSplitFullTrimmed(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitFullTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, ok := input.GetAsString("input")
		if !ok {
			errcore.HandleErrMessage("input is required")
		}
		sep, ok := input.GetAsString("sep")
		if !ok {
			errcore.HandleErrMessage("sep is required")
		}

		// Act
		pair := coregeneric.PairFromSplitFullTrimmed(inputStr, sep)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: PairFromSlice
// ==========================================

func Test_PairFromSlice(t *testing.T) {
	for caseIndex, testCase := range pairFromSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		parts, ok := input.GetAsStrings("parts")
		if !ok {
			errcore.HandleErrMessage("parts is required")
		}

		// Act
		pair := coregeneric.PairFromSlice(parts)
		actual := args.Map{
			"left":    pair.Left,
			"right":   pair.Right,
			"isValid": pair.IsValid,
			"message": pair.Message,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
