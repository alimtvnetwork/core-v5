package chmodhelpertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
)

func Test_PartialRwxVerify(t *testing.T) {
	for caseIndex, testCase := range partialRwxVerifyTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.(map[string]string)
		partialRwx := inputs["partialRwx"]
		fullRwx := inputs["fullRwx"]

		rwx, err := chmodhelper.NewRwxVariableWrapper(partialRwx)
		errcore.SimpleHandleErr(err, "rwxVar create failed.")

		// Act
		actual := rwx.IsEqualPartialRwxPartial(fullRwx)
		actLines := []string{fmt.Sprintf("%v", actual)}
		expectedLines := testCase.ExpectedInput.([]string)

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== PartialRwxVerify Diff (Case %d: %s) ===\n",
				caseIndex,
				testCase.Title,
			)

			fmt.Printf("  Input 1 (partial): %s\n", partialRwx)
			fmt.Printf("  Input 2 (full):    %s\n", fullRwx)

			errcore.PrintLineDiff(
				caseIndex,
				testCase.Title,
				actLines,
				expectedLines,
			)

			fmt.Println("=== End ===")
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
