package codestacktests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
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

		actual := args.Map{
			"filePath":   fwl.FullFilePath(),
			"lineNumber": fmt.Sprintf("%v", fwl.LineNumber()),
			"isValid":    fmt.Sprintf("%v", fwl.IsNotNil()),
		}

		// Assert
		expected := testCase.ExpectedInput.(args.Map)
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			expected,
			actual,
		)
	}
}
