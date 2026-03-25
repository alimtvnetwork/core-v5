package casenilsafetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Test: Nil-safe pointer receiver methods
// =============================================================================

func Test_CaseNilSafe_PointerReceiverMethods(t *testing.T) {
	for caseIndex, tc := range nilSafePointerReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Test: Void methods (no return values)
// =============================================================================

func Test_CaseNilSafe_VoidMethods(t *testing.T) {
	for caseIndex, tc := range nilSafeVoidTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Test: Multi-return methods
// =============================================================================

func Test_CaseNilSafe_MultiReturnMethods(t *testing.T) {
	for caseIndex, tc := range nilSafeMultiReturnTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// =============================================================================
// Test: Unsafe methods (expected panics)
// =============================================================================

func Test_CaseNilSafe_UnsafeMethods(t *testing.T) {
	for caseIndex, tc := range nilUnsafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
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

		// Assert — MethodName is not an invocation test,
		// so we compare directly via args.Map
		expected := args.Map{
			"methodName": tc.Expected.Value,
		}

		actLines := actual.CompileToStrings()
		expLines := expected.CompileToStrings()

		if len(actLines) != len(expLines) {
			t.Errorf("Case %d: line count mismatch", caseIndex)

			continue
		}

		for i, line := range actLines {
			if line != expLines[i] {
				t.Errorf("Case %d: got %q, want %q", caseIndex, line, expLines[i])
			}
		}
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
