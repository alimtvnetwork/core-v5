package resultstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// Test_Cov3_MethodName_NilInput tests MethodName with nil funcRef.
func Test_Cov3_MethodName_NilInput(t *testing.T) {
	// Arrange
	expected := ""

	// Act
	actual := results.MethodName(nil)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"MethodName(nil) should return empty string",
		actual,
		expected,
	)
}

// Test_Cov3_MethodName_NonFuncInput tests MethodName with a non-function value.
func Test_Cov3_MethodName_NonFuncInput(t *testing.T) {
	// Arrange
	expected := ""

	// Act
	actual := results.MethodName("not-a-func")

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"MethodName(string) should return empty string",
		actual,
		expected,
	)
}

// Test_Cov3_MethodName_SimpleFuncNoDot tests MethodName with a simple function name.
func Test_Cov3_MethodName_SimpleFuncNoDot(t *testing.T) {
	// Arrange
	// A plain function — name has dots in the full path, so lastDot >= 0.
	// This just verifies normal extraction works.
	myFunc := func() {}

	// Act
	actual := results.MethodName(myFunc)

	// Assert
	// The result should be non-empty (the Go runtime-assigned name)
	if actual == "" {
		t.Errorf("MethodName(func) should return non-empty, got empty")
	}
}

// Test_Cov3_InvokeWithPanicRecovery_VoidFunc tests InvokeWithPanicRecovery on a void func.
func Test_Cov3_InvokeWithPanicRecovery_VoidFunc(t *testing.T) {
	// Arrange
	voidFunc := func() {}
	expected := args.Map{
		"panicked":    false,
		"returnCount": 0,
	}

	// Act
	result := results.InvokeWithPanicRecovery(voidFunc)

	// Assert
	actual := args.Map{
		"panicked":    result.Panicked,
		"returnCount": result.ReturnCount,
	}
	coretests.GetAssert.ShouldBeEqualMap(
		t,
		0,
		"InvokeWithPanicRecovery void func",
		actual,
		expected,
	)
}

// Test_Cov3_InvokeWithPanicRecovery_NilPtrError tests extractErrorFromValue with nil ptr implementing error.
func Test_Cov3_InvokeWithPanicRecovery_NilPtrError(t *testing.T) {
	// Arrange
	// A function returning a nil *customError (which implements error)
	funcReturningNilPtrError := func() error {
		var e *customError
		return e
	}

	// Act
	result := results.InvokeWithPanicRecovery(funcReturningNilPtrError)

	// Assert
	// Error should be nil since the pointer is nil
	if result.Error != nil {
		t.Errorf("expected nil error for nil ptr error, got %v", result.Error)
	}
}

// Test_Cov3_FilterByFields_MissingKey tests filterByFields with a key not present in the map.
// This is tested indirectly through ShouldMatchResult with explicit compareFields.
func Test_Cov3_FilterByFields_MissingKey(t *testing.T) {
	// Arrange
	funcReturning42 := func() int { return 42 }
	expected := results.ResultAny{
		Value:       42,
		ReturnCount: 1,
	}

	// Act
	result := results.InvokeWithPanicRecovery(funcReturning42)

	// Assert — request a field "isSafe" that won't be in the map
	// This exercises the filterByFields missing-key branch
	result.ShouldMatchResult(
		t,
		0,
		"filterByFields missing key exercise",
		expected,
		"panicked", "value", "isSafe",
	)
}

// customError is a test helper to create a nil-pointer error type.
type customError struct {
	msg string
}

func (e *customError) Error() string {
	return fmt.Sprintf("custom: %s", e.msg)
}
