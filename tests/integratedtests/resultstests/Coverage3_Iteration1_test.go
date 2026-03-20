package resultstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/results"
)

// Test_Cov3_MethodName_NilInput tests MethodName with nil funcRef.
func Test_Cov3_MethodName_NilInput(t *testing.T) {
	// Arrange / Act
	actual := results.MethodName(nil)

	// Assert
	if actual != "" {
		t.Fatalf("MethodName(nil) should return empty string, got %q", actual)
	}
}

// Test_Cov3_MethodName_NonFuncInput tests MethodName with a non-function value.
func Test_Cov3_MethodName_NonFuncInput(t *testing.T) {
	// Arrange / Act
	actual := results.MethodName("not-a-func")

	// Assert
	if actual != "" {
		t.Fatalf("MethodName(string) should return empty string, got %q", actual)
	}
}

// Test_Cov3_MethodName_SimpleFuncNoDot tests MethodName with a simple function.
func Test_Cov3_MethodName_SimpleFuncNoDot(t *testing.T) {
	// Arrange
	myFunc := func() {}

	// Act
	actual := results.MethodName(myFunc)

	// Assert
	if actual == "" {
		t.Fatalf("MethodName(func) should return non-empty, got empty")
	}
}

// Test_Cov3_InvokeWithPanicRecovery_VoidFunc tests InvokeWithPanicRecovery on a void func.
func Test_Cov3_InvokeWithPanicRecovery_VoidFunc(t *testing.T) {
	// Arrange
	voidFunc := func() {}

	// Act — signature is InvokeWithPanicRecovery(funcRef any, receiver any, args ...any)
	result := results.InvokeWithPanicRecovery(voidFunc, nil)

	// Assert
	if result.Panicked {
		t.Fatalf("void func should not panic, got panicked")
	}
	if result.ReturnCount != 0 {
		t.Fatalf("void func return count: got %d, want 0", result.ReturnCount)
	}
}

// Test_Cov3_InvokeWithPanicRecovery_NilPtrError tests extractErrorFromValue with nil ptr implementing error.
func Test_Cov3_InvokeWithPanicRecovery_NilPtrError(t *testing.T) {
	// Arrange
	funcReturningNilPtrError := func() error {
		var e *customError
		return e
	}

	// Act
	result := results.InvokeWithPanicRecovery(funcReturningNilPtrError, nil)

	// Assert
	if result.Error != nil {
		t.Fatalf("expected nil error for nil ptr error, got %v", result.Error)
	}
}

// Test_Cov3_InvokeWithPanicRecovery_FuncReturning42 tests with a func returning int.
func Test_Cov3_InvokeWithPanicRecovery_FuncReturning42(t *testing.T) {
	// Arrange
	funcReturning42 := func() int { return 42 }

	// Act
	result := results.InvokeWithPanicRecovery(funcReturning42, nil)

	// Assert
	if result.Panicked {
		t.Fatalf("expected no panic")
	}
	if result.ReturnCount != 1 {
		t.Fatalf("expected return count 1, got %d", result.ReturnCount)
	}
}

// Test_Cov3_FilterByFields_MissingKey tests filterByFields with a key not present in the map.
func Test_Cov3_FilterByFields_MissingKey(t *testing.T) {
	// Arrange
	funcReturning42 := func() int { return 42 }
	expected := results.ResultAny{
		Value:       42,
		ReturnCount: 1,
	}

	// Act
	result := results.InvokeWithPanicRecovery(funcReturning42, nil)

	// Assert — request a field "isSafe" that won't be in the map
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
