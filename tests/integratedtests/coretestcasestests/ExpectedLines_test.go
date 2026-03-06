package coretestcasestests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
)

func Test_CaseV1_ExpectedLines_Verification(t *testing.T) {
	for caseIndex, tc := range expectedLinesTestCases {
		// Arrange — ExpectedInput varies by type (int, bool, []int, etc.)
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

		// Assert — use a verification CaseV1 with expected output as ExpectedInput
		verifyCaseV1 := expectedLinesVerificationCases[caseIndex]
		verifyCaseV1.ShouldBeEqualMap(t, caseIndex, actual)

		_ = expectedOutput // referenced via verifyCaseV1
	}
}
