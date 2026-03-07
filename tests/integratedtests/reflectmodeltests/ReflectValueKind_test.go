package reflectmodeltests

import (
	"testing"

	"gitlab.com/auk-go/core/reflectcore/reflectmodel"
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

func Test_ReflectValueKind_IsInvalid_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	if !rvk.IsInvalid() {
		t.Error("expected IsInvalid() = true on nil receiver")
	}
}

func Test_ReflectValueKind_IsInvalid_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	if !rvk.IsInvalid() {
		t.Error("expected IsInvalid() = true when IsValid=false")
	}
}

func Test_ReflectValueKind_HasError_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	if rvk.HasError() {
		t.Error("expected HasError() = false on nil receiver")
	}
}

func Test_ReflectValueKind_IsEmptyError_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	if !rvk.IsEmptyError() {
		t.Error("expected IsEmptyError() = true on nil receiver")
	}
}

func Test_ReflectValueKind_IsEmptyError_NoError(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{}

	if !rvk.IsEmptyError() {
		t.Error("expected IsEmptyError() = true when no error")
	}
}

func Test_ReflectValueKind_ActualInstance_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	got := rvk.ActualInstance()
	if got != nil {
		t.Errorf("expected ActualInstance() = nil on nil receiver, got %v", got)
	}
}

func Test_ReflectValueKind_PkgPath_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	got := rvk.PkgPath()
	if got != "" {
		t.Errorf("expected PkgPath() = empty on nil receiver, got %q", got)
	}
}

func Test_ReflectValueKind_PkgPath_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.PkgPath()
	if got != "" {
		t.Errorf("expected PkgPath() = empty when IsValid=false, got %q", got)
	}
}

func Test_ReflectValueKind_TypeName_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	got := rvk.TypeName()
	if got != "" {
		t.Errorf("expected TypeName() = empty on nil receiver, got %q", got)
	}
}

func Test_ReflectValueKind_TypeName_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.TypeName()
	if got != "" {
		t.Errorf("expected TypeName() = empty when IsValid=false, got %q", got)
	}
}

func Test_ReflectValueKind_PointerRv_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	got := rvk.PointerRv()
	if got != nil {
		t.Error("expected PointerRv() = nil on nil receiver")
	}
}

func Test_ReflectValueKind_PointerInterface_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind

	got := rvk.PointerInterface()
	if got != nil {
		t.Error("expected PointerInterface() = nil on nil receiver")
	}
}
