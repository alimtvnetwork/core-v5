package coretestcasestests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

func Test_CaseV1_ExpectedLines_Verification(t *testing.T) {
	for caseIndex, tc := range expectedLinesTestCases {
		// Arrange
		expectedOutput := expectedLinesExpectedOutputs[caseIndex]

		// Act
		actLines := tc.ExpectedLines()

		// Build actual args.Map with lineCount + indexed keys
		actual := args.Map{
			"lineCount": fmt.Sprintf("%d", len(actLines)),
		}
		for i, line := range actLines {
			actual[fmt.Sprintf("line%d", i)] = line
		}

		// Assert via verification CaseV1
		verifyCase := coretestcases.CaseV1{
			Title:         tc.Title,
			ExpectedInput: expectedOutput,
		}
		verifyCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
