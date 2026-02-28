package namevaluetests

import (
	"fmt"
	"strings"
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

func Test_Instance_JsonString_Verification(t *testing.T) {
	for caseIndex, testCase := range instanceJsonStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		value, _ := input.Get("value")

		// Act
		instance := namevalue.Instance{
			Name:  name,
			Value: value,
		}
		jsonStr := instance.JsonString()
		isNotEmpty := fmt.Sprintf("%v", jsonStr != "")
		containsName := fmt.Sprintf("%v", strings.Contains(jsonStr, name))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotEmpty, containsName)
	}
}

func Test_Instance_Dispose_Verification(t *testing.T) {
	for caseIndex, testCase := range instanceDisposeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		name, _ := input.GetAsString("name")
		value, _ := input.Get("value")

		// Act
		instance := &namevalue.Instance{
			Name:  name,
			Value: value,
		}
		instance.Dispose()
		isValueNil := fmt.Sprintf("%v", instance.Value == nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, instance.Name, isValueNil)
	}
}
