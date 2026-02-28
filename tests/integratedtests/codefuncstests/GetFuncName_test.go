package codefuncstests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/corefuncs"
)

func sampleFunc() {}

func Test_GetFuncName_Verification(t *testing.T) {
	for caseIndex, testCase := range getFuncNameTestCases {
		// Act
		name := corefuncs.GetFuncName(sampleFunc)
		isNotEmpty := fmt.Sprintf("%v", name != "")

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotEmpty)
	}
}
