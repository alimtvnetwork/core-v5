package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
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
