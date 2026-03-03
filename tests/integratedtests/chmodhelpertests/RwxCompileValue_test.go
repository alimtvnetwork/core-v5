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
		errcore.PrintDiffOnMismatch(
			caseIndex,
			testCase.Case.Title,
			actLines,
			expectedLines,
			fmt.Sprintf("  Existing: %s", testCase.Existing.ToString(false)),
			fmt.Sprintf("  Input:    %s", testCase.Input.ToString(false)),
			fmt.Sprintf("  Expected: %s", testCase.Expected.ToString(false)),
		)

		// Assert
		caseV1 := testCase.Case
		caseV1.ExpectedInput = expectedLines

		caseV1.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
