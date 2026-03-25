package reflectmodeltests

import (
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ===== ReflectValueKind Tests =====

func Test_InvalidReflectValueKindModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")

	if rvk.IsValid {
		t.Error("expected IsValid = false")
	}

	if !rvk.HasError() {
		t.Error("expected HasError() = true")
	}

	if rvk.Error.Error() != "test error" {
		t.Errorf("Error = %q, want %q", rvk.Error.Error(), "test error")
	}
}

// Note: All nil receiver tests migrated to ReflectValueKind_NilReceiver_testcases.go

func Test_ReflectValueKind_NilReceiver(t *testing.T) {
	for caseIndex, tc := range reflectValueKindNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_ReflectValueKind_IsInvalid_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	if !rvk.IsInvalid() {
		t.Error("expected IsInvalid() = true when IsValid=false")
	}
}

func Test_ReflectValueKind_IsEmptyError_NoError(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{}

	if !rvk.IsEmptyError() {
		t.Error("expected IsEmptyError() = true when no error")
	}
}

func Test_ReflectValueKind_PkgPath_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.PkgPath()
	if got != "" {
		t.Errorf("expected PkgPath() = empty when IsValid=false, got %q", got)
	}
}

func Test_ReflectValueKind_TypeName_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.TypeName()
	if got != "" {
		t.Errorf("expected TypeName() = empty when IsValid=false, got %q", got)
	}
}
