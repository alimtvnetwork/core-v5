package coretestcasestests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

func Test_CaseV1_ExpectedLines_Verification(t *testing.T) {
	for caseIndex, tc := range expectedLinesTestCases {
		// Arrange — ExpectedInput varies by type (int, bool, []int, etc.)
		expected := expectedLinesExpectedOutputs[caseIndex]

		// Act
		actLines := tc.ExpectedLines()

		// Build actual args.Map with lineCount + indexed keys
		actual := args.Map{
			"lineCount": fmt.Sprintf("%d", len(actLines)),
		}
		for i, line := range actLines {
			actual[fmt.Sprintf("line%d", i)] = line
		}

		// Assert
		verifyCase := coretestcases.CaseV1{
			Title:         tc.Title,
			ExpectedInput: expected,
		}
		errcore.AssertMapDiffOnMismatch(t, caseIndex, verifyCase.Title, actual, expected)
	}
}
