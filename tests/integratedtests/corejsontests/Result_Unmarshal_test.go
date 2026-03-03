package corejsontests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"
)

func Test_Result_Unmarshal_Valid(t *testing.T) {
	tc := resultUnmarshalValidTestCases[0]

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
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Result_Unmarshal_NilReceiver(t *testing.T) {
	tc := resultUnmarshalNilTestCases[0]

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
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Result_Unmarshal_InvalidBytes(t *testing.T) {
	tc := resultUnmarshalInvalidTestCases[0]

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
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_Result_Unmarshal_ExistingError(t *testing.T) {
	tc := resultUnmarshalExistingErrorTestCases[0]

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
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}
