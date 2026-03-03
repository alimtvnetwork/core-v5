package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_LinuxApplyRecursiveOnPath_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	for caseIndex, testCase := range []coretestcases.CaseV1{linuxApplyRecursiveOnPathTestCase} {
		// Arrange
		wrapper := testCase.ArrangeInput.(chmodhelpertestwrappers.RwxInstructionTestWrapper)
		chmodhelper.CreateDirFilesWithRwxPermissionsMust(
			true,
			wrapper.CreatePaths,
		)

		// Act
		actLine := applyAndCollectResult(&wrapper, linuxApplyRecursivePathInstructions)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLine)
	}
}
