package casenilsafetests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
)

// =============================================================================
// Test: Nil-safe pointer receiver methods
// =============================================================================

func Test_CaseNilSafe_PointerReceiverMethods(t *testing.T) {
	for caseIndex, tc := range nilSafePointerReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act
		result := tc.InvokeNil()

		// Assert
		tc.ShouldBeSafe(t, caseIndex)
		_ = result
	}
}

// =============================================================================
// Test: Void methods (no return values)
// =============================================================================

func Test_CaseNilSafe_VoidMethods(t *testing.T) {
	for caseIndex, tc := range nilSafeVoidTestCases {
		// Arrange (implicit — nil receiver)

		// Act
		result := tc.InvokeNil()
		actual := filterToExpected(result.ToMap(), tc.Expected)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Test: Multi-return methods
// =============================================================================

func Test_CaseNilSafe_MultiReturnMethods(t *testing.T) {
	for caseIndex, tc := range nilSafeMultiReturnTestCases {
		// Arrange (implicit — nil receiver)

		// Act
		result := tc.InvokeNil()
		actual := filterToExpected(result.ToMap(), tc.Expected)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Test: Unsafe methods (expected panics)
// =============================================================================

func Test_CaseNilSafe_UnsafeMethods(t *testing.T) {
	for caseIndex, tc := range nilUnsafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act
		result := tc.InvokeNil()
		actual := filterToExpected(result.ToMap(), tc.Expected)

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Test: MethodName extraction
// =============================================================================

func Test_CaseNilSafe_MethodName(t *testing.T) {
	for caseIndex, tc := range methodNameTestCases {
		// Arrange
		name := tc.MethodName()

		// Act
		actual := args.Map{
			"methodName": name,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// Test: CaseTitle fallback
// =============================================================================

func Test_CaseNilSafe_CaseTitleFallback(t *testing.T) {
	// Arrange
	tc := nilSafePointerReceiverTestCases[0]
	tcNoTitle := tc
	tcNoTitle.Title = ""

	// Act
	titleWithExplicit := tc.CaseTitle()
	titleFromMethod := tcNoTitle.CaseTitle()

	// Assert
	if titleWithExplicit != "IsValid on nil returns false" {
		t.Errorf("expected explicit title, got %q", titleWithExplicit)
	}

	if titleFromMethod != "IsValid" {
		t.Errorf("expected method name fallback, got %q", titleFromMethod)
	}
}

// =============================================================================
// Test: Invoke with non-nil receiver
// =============================================================================

func Test_CaseNilSafe_InvokeWithReceiver(t *testing.T) {
	// Arrange
	tc := nilSafePointerReceiverTestCases[0] // IsValid
	receiver := &sampleStruct{Name: "hello", Value: 42}

	// Act
	result := tc.Invoke(receiver)

	// Assert
	if result.HasPanicked() {
		t.Error("should not panic with valid receiver")
	}

	if result.ValueString() != "true" {
		t.Errorf("expected true, got %s", result.ValueString())
	}
}

// filterToExpected returns a subset of actual containing only keys in expected.
func filterToExpected(actual args.Map, expected args.Map) args.Map {
	filtered := args.Map{}

	for key := range expected {
		if val, exists := actual[key]; exists {
			filtered[key] = val
		}
	}

	return filtered
}
