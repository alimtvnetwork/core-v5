package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_SimpleFileWriter(t *testing.T) {
	for caseIndex, testCase := range simpleFileWriterTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.One)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			files := parameter.First.([]chmodhelper.DirFilesWithContent)

			for _, dirFiles := range files {
				err := dirFiles.CreateUsingFileMode(
					true,
				)

				errcore.HandleErr(err)

				for _, file := range dirFiles.Files {
					lines, err2 := file.ReadLines(dirFiles.Dir)

					errcore.HandleErr(err2)

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
		testCase.AssertEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
