package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxChmodUsingRwxInstructions_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Setup
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		pathInstructionsV2,
	)

	for caseIndex, testCase := range verifyRwxChmodUsingRwxInstructionsTestCases {
		// Arrange
		wrapper := testCase.ArrangeInput.(chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsWrapper)
		executor, err := chmodhelper.ParseRwxInstructionToExecutor(&wrapper.RwxInstruction)
		errcore.SimpleHandleErr(err, "")

		// Act
		actualErr := executor.VerifyRwxModifiersDirect(
			false,
			wrapper.Locations...,
		)

		// Assert
		assertNonWhiteSortedEqual(t, testCase, caseIndex, actualErr)
	}
}
