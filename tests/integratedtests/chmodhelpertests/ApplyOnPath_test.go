package chmodhelpertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_ApplyOnPath_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	for caseIndex, testCase := range applyOnPathTestCases {
		// Arrange
		wrapper := testCase.ArrangeInput.(chmodhelpertestwrappers.RwxInstructionTestWrapper)
		
		// Replace hardcoded paths with temp dirs
		tempDir := t.TempDir()
		modifiedPaths := make([]chmodhelper.DirFilesWithRwxPermission, len(wrapper.CreatePaths))
		for i, path := range wrapper.CreatePaths {
			modifiedPaths[i] = path
			modifiedPaths[i].DirWithFiles.Dir = tempDir + "/test-cases-" + fmt.Sprintf("%d", i)
		}
		wrapper.CreatePaths = modifiedPaths
		
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
