package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: TripleFromSplit
// ==========================================

func Test_TripleFromSplit(t *testing.T) {
	for caseIndex, testCase := range tripleFromSplitTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

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
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

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
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

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
		inputStr, inputErr := input.GetAsString("input")
		errcore.HandleErrMessage("input", inputErr)
		sep, sepErr := input.GetAsString("sep")
		errcore.HandleErrMessage("sep", sepErr)

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
		parts, partsErr := input.GetAsStringSlice("parts")
		errcore.HandleErrMessage("parts", partsErr)

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
