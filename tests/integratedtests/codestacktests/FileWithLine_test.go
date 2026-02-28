package codestacktests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_FileWithLine_Verification(t *testing.T) {
	for caseIndex, testCase := range fileWithLineTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		file, _ := input.GetAsString("file")
		line, _ := input.GetAsInt("line")

		// Act
		fwl := &codestack.FileWithLine{
			FilePath: file,
			Line:     line,
		}

		fullPath := fwl.FullFilePath()
		lineNum := fmt.Sprintf("%v", fwl.LineNumber())
		isNotNil := fmt.Sprintf("%v", fwl.IsNotNil())

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fullPath,
			lineNum,
			isNotNil,
		)
	}
}
