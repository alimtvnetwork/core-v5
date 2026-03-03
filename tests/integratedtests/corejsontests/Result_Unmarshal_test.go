package corejsontests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
)

func Test_Result_Unmarshal_Valid(t *testing.T) {
	tc := resultUnmarshalValidTestCase

	// Arrange
	src := exampleStruct{Name: "Alice", Age: 30}
	jsonResult := corejson.NewPtr(src)
	target := &exampleStruct{}

	// Act
	err := jsonResult.Unmarshal(target)

	actLines := []string{
		fmt.Sprintf("%v", err),
		target.Name,
		fmt.Sprintf("%v", target.Age),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Result_Unmarshal_NilReceiver(t *testing.T) {
	tc := resultUnmarshalNilTestCase

	// Arrange
	var nilResult *corejson.Result
	target := &exampleStruct{}

	// Act
	err := nilResult.Unmarshal(target)

	actLines := []string{
		fmt.Sprintf("%v", err != nil),
		fmt.Sprintf("%v", strings.Contains(err.Error(), "null")),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Result_Unmarshal_InvalidBytes(t *testing.T) {
	tc := resultUnmarshalInvalidTestCase

	// Arrange
	result := corejson.NewResult.UsingBytesTypePtr([]byte(`{invalid-json`), "TestType")
	target := &exampleStruct{}

	// Act
	err := result.Unmarshal(target)

	actLines := []string{
		fmt.Sprintf("%v", err != nil),
		fmt.Sprintf("%v", strings.Contains(err.Error(), "unmarshal")),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Result_Unmarshal_ExistingError(t *testing.T) {
	tc := resultUnmarshalExistingErrorTestCase

	// Arrange
	ch := make(chan int)
	result := corejson.NewPtr(ch)
	target := &exampleStruct{}

	// Act
	err := result.Unmarshal(target)

	actLines := []string{
		fmt.Sprintf("%v", err != nil),
		fmt.Sprintf("%v", strings.Contains(err.Error(), "unmarshal")),
	}

	// Assert
	tc.ShouldBeEqual(t, 0, actLines...)
}
