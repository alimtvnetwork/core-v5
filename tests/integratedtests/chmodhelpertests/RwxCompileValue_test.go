package chmodhelpertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
)

func Test_RwxCompileValue(t *testing.T) {
	for caseIndex, testCase := range rwxCompileValueTestCases {
		// Arrange
		existingRwxWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Existing,
			)
		expectedVariableWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Expected,
			)

		expectedFullRwx := expectedVariableWrapper.
			ToCompileFixedPtr().
			ToFullRwxValueString()

		// Act
		actualVarWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Input,
			)
		actualRwxWrapper := actualVarWrapper.
			ToCompileWrapper(existingRwxWrapper.ToCompileFixedPtr())
		actualFullRwx := actualRwxWrapper.ToFullRwxValueString()

		actLines := []string{actualFullRwx}
		expectedLines := []string{expectedFullRwx}

		// Print diff on failure
		if errcore.LineDiffHasMismatch(actLines, expectedLines) {
			fmt.Printf(
				"\n=== RwxCompileValue Diff (Case %d: %s) ===\n",
				caseIndex,
				testCase.Case.Title,
			)

			existing := testCase.Existing.ToString(false)
			input := testCase.Input.ToString(false)
			expected := testCase.Expected.ToString(false)

			fmt.Printf("  Existing: %s\n", existing)
			fmt.Printf("  Input:    %s\n", input)
			fmt.Printf("  Expected: %s\n", expected)

			errcore.PrintLineDiff(
				caseIndex,
				testCase.Case.Title,
				actLines,
				expectedLines,
			)

			fmt.Println("=== End ===")
		}

		// Assert
		caseV1 := testCase.Case
		caseV1.ExpectedInput = expectedLines

		caseV1.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
