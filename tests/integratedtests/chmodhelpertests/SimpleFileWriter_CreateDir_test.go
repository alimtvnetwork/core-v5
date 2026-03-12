package chmodhelpertests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/filemode"
	"github.com/alimtvnetwork/core/internal/pathinternal"
	"github.com/alimtvnetwork/core/iserror"
)

func Test_SimpleFileWriter_CreateDir_If_Verification(t *testing.T) {
	temp := pathinternal.GetTemp()
	chmodhelper.SimpleFileWriter.Lock()
	defer chmodhelper.SimpleFileWriter.Unlock()

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

			pathinternal.RemoveDirMust(
				dir,
				"Test_SimpleFileWriter_CreateDir_Verification",
			)

			for fileIndex, file := range input.Files {
				finalPath := pathinternal.Join(dir, file)
				parentDir := pathinternal.ParentDir(finalPath)

				err := createDir.If(
					true,
					filemode.DirDefault,
					parentDir,
				)

				errcore.HandleErr(err)
				relPath := pathinternal.Relative(temp, parentDir)

				if iserror.Defined(err) {
					actualSlice.AppendFmt(
						"%d - %d : %s - isCreated : %t, err: %s",
						i,
						fileIndex,
						relPath,
						chmodhelper.IsPathExists(parentDir),
						errcore.ToString(err),
					)
				} else {
					actualSlice.AppendFmt(
						"%d - %d : %s - isCreated : %t",
						i,
						fileIndex,
						relPath,
						chmodhelper.IsPathExists(parentDir),
					)
				}
			}

			pathinternal.RemoveDirMustSimple(dir)
		}

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_SimpleFileWriter_CreateDir_IfMissing_Verification(t *testing.T) {
	temp := pathinternal.GetTemp()
	chmodhelper.SimpleFileWriter.Lock()
	defer chmodhelper.SimpleFileWriter.Unlock()

	for caseIndex, testCase := range createDirIfMissingTestCases {
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

			pathinternal.RemoveDirMust(
				dir,
				"Test_SimpleFileWriter_CreateDir_IfMissing_Verification",
			)

			for fileIndex, file := range input.Files {
				finalPath := pathinternal.Join(dir, file)
				parentDir := pathinternal.ParentDir(finalPath)

				err := createDir.IfMissing(
					filemode.DirDefault,
					parentDir,
				)

				errcore.HandleErr(err)
				relPath := pathinternal.Relative(temp, parentDir)

				if iserror.Defined(err) {
					actualSlice.AppendFmt(
						"%d - %d : %s - isCreated : %t, err: %s",
						i,
						fileIndex,
						relPath,
						chmodhelper.IsPathExists(parentDir),
						errcore.ToString(err),
					)
				} else {
					actualSlice.AppendFmt(
						"%d - %d : %s - isCreated : %t",
						i,
						fileIndex,
						relPath,
						chmodhelper.IsPathExists(parentDir),
					)
				}
			}

			pathinternal.RemoveDirMustSimple(dir)
		}

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_SimpleFileWriter_CreateDir_ExistingFileFails_Verification(t *testing.T) {
	temp := pathinternal.GetTemp()
	chmodhelper.SimpleFileWriter.Lock()
	defer chmodhelper.SimpleFileWriter.Unlock()

	for caseIndex, testCase := range createDirDirectTestCases {
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
		fileWriter := chmodhelper.
			SimpleFileWriter.
			FileWriter

		// Act
		for i, input := range inputs {
			dir := input.Dir

			pathinternal.RemoveDirMust(
				dir,
				"Test_SimpleFileWriter_CreateDir_ExistingFileFails_Verification",
			)

			for fileIndex, file := range input.Files {
				finalPath := pathinternal.Join(dir, file)

				err := fileWriter.String.Default(
					false,
					finalPath,
					"",
				)

				errcore.HandleErr(err)

				finalErr := createDir.ByChecking(
					filemode.DirDefault,
					finalPath,
				)

				errorString := errcore.ToString(finalErr)
				errorString = strings.ReplaceAll(
					errorString,
					finalPath,
					"",
				)

				relPath := pathinternal.Relative(temp, finalPath)

				if iserror.Defined(finalErr) {
					actualSlice.AppendFmt(
						"%d - %d : %s - already exist as file, err: %s",
						i,
						fileIndex,
						relPath,
						errorString,
					)
				} else {
					actualSlice.AppendFmt(
						"%d - %d : %s - no error during 2nd invoke of createDir.Direct",
						i,
						fileIndex,
						relPath,
					)
				}
			}

			pathinternal.RemoveDirMustSimple(dir)
		}

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_SimpleFileWriter_CreateDir_ByCheckingFails_Verification(t *testing.T) {
	temp := pathinternal.GetTemp()
	chmodhelper.SimpleFileWriter.Lock()
	defer chmodhelper.SimpleFileWriter.Unlock()

	for caseIndex, testCase := range createDirByCheckingTestCases {
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
		fileWriter := chmodhelper.
			SimpleFileWriter.
			FileWriter

		// Act
		for i, input := range inputs {
			dir := input.Dir

			pathinternal.RemoveDirMust(
			"Test_SimpleFileWriter_CreateDir_ByCheckingFails_Verification",
			)

			for fileIndex, file := range input.Files {
				finalPath := pathinternal.Join(dir, file)
				parentDir := pathinternal.ParentDir(finalPath)

				err := fileWriter.String.Chmod(
					false,
					filemode.X200,
					filemode.X300,
					finalPath,
					"some thing",
				)

				errcore.HandleErr(err)

				finalErr := createDir.ByChecking(
					filemode.X400,
					parentDir,
				)

				errorString := errcore.ToString(finalErr)
				errorString = strings.ReplaceAll(
					errorString,
					finalPath,
					"",
				)

				relPath := pathinternal.Relative(temp, finalPath)

				if iserror.Defined(finalErr) {
					actualSlice.AppendFmt(
						"%d - %d : %s - already exist as file, err: %s",
						i,
						fileIndex,
						relPath,
						errorString,
					)
				} else {
					actualSlice.AppendFmt(
						"%d - %d : %s - no error during 2nd invoke of createDir.Direct",
						i,
						fileIndex,
						relPath,
					)
				}
			}

			pathinternal.RemoveDirMustSimple(dir)
		}

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
