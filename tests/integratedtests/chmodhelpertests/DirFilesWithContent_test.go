package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_DirFilesWithContent_Create_Read_Verification(t *testing.T) {
	chmodhelper.SimpleFileWriter.Lock()
	defer chmodhelper.SimpleFileWriter.Unlock()

	for caseIndex, testCase := range dirFilesWithContentCreateReadTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.OneAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			files := parameter.First.([]chmodhelper.DirFilesWithContent)

			for _, dirFiles := range files {
				err := dirFiles.Create(true)

				errcore.HandleErr(err)

				for _, file := range dirFiles.Files {
					lines, readErr := file.ReadLines(dirFiles.Dir)

					errcore.HandleErr(readErr)

					actualSlice.AppendFmt(
						"%d : %s",
						i,
						file.RelativePath,
					)

					for lineIndex, line := range lines {
						actualSlice.AppendFmt(
							"         %d. %s",
							lineIndex,
							line,
						)
					}
				}
			}
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
