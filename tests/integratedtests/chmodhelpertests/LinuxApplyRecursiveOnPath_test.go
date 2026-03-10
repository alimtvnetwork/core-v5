package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
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
