package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/filemode"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

func Test_SimpleFileWriter_CreateDir_Verification(t *testing.T) {
	for caseIndex, testCase := range createDirTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]chmodhelper.DirWithFiles)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))
		createDir := chmodhelper.
			SimpleFileWriter.
			CreateDir

		// Act
		for i, input := range inputs {
			dir := input.Dir
			for fileIndex, file := range input.Files {
				finalPath := pathinternal.Join(dir, file)
				parentDir := pathinternal.ParentDir(finalPath)

				err := createDir.If(
					true,
					filemode.DirDefault,
					parentDir,
				)

				errcore.HandleErr(err)

				actualSlice.AppendFmt(
					"%d - %d : %s - isCreated : %t, err: %s",
					i,
					fileIndex,
					parentDir,
					chmodhelper.IsPathExists(parentDir),
					errcore.ToString(err),
				)
			}
		}

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.AssertEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
