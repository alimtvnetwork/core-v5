package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_ApplyOnPath_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	for caseIndex, testCase := range applyOnPathTestCases {
		// Arrange
		wrapper := testCase.ArrangeInput.(chmodhelpertestwrappers.RwxInstructionTestWrapper)
		chmodhelper.CreateDirFilesWithRwxPermissionsMust(
			true,
			wrapper.CreatePaths,
		)

		// Act
		actLine := applyAndCollectResult(&wrapper, applyPathInstructions)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLine)
	}
}
