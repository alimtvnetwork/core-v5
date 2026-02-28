package defaulterrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/defaulterr"
)

func Test_DefaultErr_Verification(t *testing.T) {
	for caseIndex, testCase := range defaultErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		errorName, _ := input.GetAsString("error")

		// Act
		var err error
		switch errorName {
		case "Marshalling":
			err = defaulterr.Marshalling
		case "UnMarshalling":
			err = defaulterr.UnMarshalling
		case "OutOfRange":
			err = defaulterr.OutOfRange
		}

		isNotNil := fmt.Sprintf("%v", err != nil)
		hasMessage := fmt.Sprintf("%v", err != nil && err.Error() != "")

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotNil, hasMessage)
	}
}
