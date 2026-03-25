package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
)

func Test_LineNumber_HasLineNumber(t *testing.T) {
	for caseIndex, tc := range lineNumberHasLineNumberTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		lineNum, _ := input.GetAsInt("lineNumber")
		ln := corevalidator.LineNumber{LineNumber: lineNum}

		// Act
		actual := args.Map{
			"result": ln.HasLineNumber(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LineNumber_IsMatch(t *testing.T) {
	for caseIndex, tc := range lineNumberIsMatchTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		lineNum, _ := input.GetAsInt("lineNumber")
		inputNum, _ := input.GetAsInt("input")
		ln := corevalidator.LineNumber{LineNumber: lineNum}

		// Act
		actual := args.Map{
			"result": ln.IsMatch(inputNum),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LineNumber_VerifyError(t *testing.T) {
	for caseIndex, tc := range lineNumberVerifyErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		lineNum, _ := input.GetAsInt("lineNumber")
		inputNum, _ := input.GetAsInt("input")
		ln := corevalidator.LineNumber{LineNumber: lineNum}

		// Act
		err := ln.VerifyError(inputNum)

		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
