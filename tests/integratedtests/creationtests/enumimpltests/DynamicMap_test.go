package enumimpltests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func Test_DynamicMapDiff1(t *testing.T) {
	for caseIndex, testCase := range dynamicMapSimpleDiffTestCases {
		// Arrange
		arrangeInput := testCase.ArrangeAsLeftRightDynamicMapWithDefaultChecker()

		// Act
		diffJsonMessage := arrangeInput.Left.ShouldDiffLeftRightMessageUsingDifferChecker(
			arrangeInput.DifferChecker,
			true,
			testCase.CaseTitle(),
			arrangeInput.Right,
		)

		actualLines := strings.Split(
			diffJsonMessage,
			constants.NewLineUnix,
		)

		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.PrintLineDiff(caseIndex, testCase.Title, actualLines, expectedLines)

		if len(actualLines) != len(expectedLines) {
			t.Errorf("[case %d] %s: line count mismatch got %d, want %d",
				caseIndex, testCase.Title, len(actualLines), len(expectedLines))
			continue
		}

		for i, act := range actualLines {
			if act != expectedLines[i] {
				t.Errorf("[case %d] %s: line %d got %q, want %q",
					caseIndex, testCase.Title, i, act, expectedLines[i])
			}
		}
	}
}
