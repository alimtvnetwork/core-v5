package coregenerictests

import (
	"fmt"
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
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

		// Act
		pair := coregeneric.PairFromSplit(inputStr, sep)
		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: PairFromSplitTrimmed
// ==========================================

func Test_PairFromSplitTrimmed(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

		// Act
		pair := coregeneric.PairFromSplitTrimmed(inputStr, sep)
		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: PairFromSplitFull
// ==========================================

func Test_PairFromSplitFull(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitFullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

		// Act
		pair := coregeneric.PairFromSplitFull(inputStr, sep)
		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: PairFromSplitFullTrimmed
// ==========================================

func Test_PairFromSplitFullTrimmed(t *testing.T) {
	for caseIndex, testCase := range pairFromSplitFullTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

		// Act
		pair := coregeneric.PairFromSplitFullTrimmed(inputStr, sep)
		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: PairFromSlice
// ==========================================

func Test_PairFromSlice(t *testing.T) {
	for caseIndex, testCase := range pairFromSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		parts, partsErr := input.GetAsStringSlice("parts")
		errcore.HandleErrMessage("parts", partsErr)

		// Act
		pair := coregeneric.PairFromSlice(parts)
		actLines := []string{
			pair.Left,
			pair.Right,
			fmt.Sprintf("%v", pair.IsValid),
			pair.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
