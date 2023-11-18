package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/isany"
)

func Test_SimpleFileWriter_Unix(t *testing.T) {
	for caseIndex, testCase := range simpleFileWriterTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.Two)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			files := parameter.First.([]chmodhelper.DirFilesWithContent)

			for i2, filesWithContent := range files {

			}

			actualSlice.AppendFmt(
				defaultCaseIndexBoolStringStringFmt,
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.AssertEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
