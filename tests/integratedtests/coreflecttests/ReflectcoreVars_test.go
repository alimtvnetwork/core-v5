package coreflecttests

import (
	"testing"

	"gitlab.com/auk-go/core/reflectcore"
)

// ==========================================
// reflectcore vars — facade re-exports are non-nil
// ==========================================

func Test_Reflectcore_Converter_NotNil(t *testing.T) {
	if reflectcore.Converter == nil {
		t.Error("Converter should not be nil")
	}
}

func Test_Reflectcore_Utils_NotNil(t *testing.T) {
	if reflectcore.Utils == nil {
		t.Error("Utils should not be nil")
	}
}

func Test_Reflectcore_Looper_NotNil(t *testing.T) {
	if reflectcore.Looper == nil {
		t.Error("Looper should not be nil")
	}
}

func Test_Reflectcore_CodeStack_NotNil(t *testing.T) {
	if reflectcore.CodeStack == nil {
		t.Error("CodeStack should not be nil")
	}
}

func Test_Reflectcore_GetFunc_NotNil(t *testing.T) {
	if reflectcore.GetFunc == nil {
		t.Error("GetFunc should not be nil")
	}
}

func Test_Reflectcore_Is_NotNil(t *testing.T) {
	if reflectcore.Is == nil {
		t.Error("Is should not be nil")
	}
}

func Test_Reflectcore_TypeName_NotNil(t *testing.T) {
	if reflectcore.TypeName == nil {
		t.Error("TypeName should not be nil")
	}
}

func Test_Reflectcore_TypeNames_NotNil(t *testing.T) {
	if reflectcore.TypeNames == nil {
		t.Error("TypeNames should not be nil")
	}
}

func Test_Reflectcore_ReflectType_NotNil(t *testing.T) {
	if reflectcore.ReflectType == nil {
		t.Error("ReflectType should not be nil")
	}
}

func Test_Reflectcore_ReflectGetter_NotNil(t *testing.T) {
	if reflectcore.ReflectGetter == nil {
		t.Error("ReflectGetter should not be nil")
	}
}

func Test_Reflectcore_SliceConverter_NotNil(t *testing.T) {
	if reflectcore.SliceConverter == nil {
		t.Error("SliceConverter should not be nil")
	}
}

func Test_Reflectcore_MapConverter_NotNil(t *testing.T) {
	if reflectcore.MapConverter == nil {
		t.Error("MapConverter should not be nil")
	}
}
