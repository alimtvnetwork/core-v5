package namevaluetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/namevalue"
)

func Test_Instance_String_Verification(t *testing.T) {
	for caseIndex, testCase := range instanceStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		value, _ := input.Get("value")

		// Act
		instance := namevalue.Instance{
			Name:  name,
			Value: value,
		}
		result := instance.String()
		isNotEmpty := fmt.Sprintf("%v", result != "")

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotEmpty)
	}
}
