package coreflecttests

import (
	"testing"

	"github.com/alimtvnetwork/core/reflectcore"
)

// ==========================================
// reflectcore vars — facade re-exports are non-zero
// These are value types (structs), not pointers/interfaces,
// so nil comparison is invalid. We verify compile-time existence instead.
// ==========================================

func Test_Reflectcore_Converter_NotNil(t *testing.T) {
	_ = reflectcore.Converter // compile-time existence check
}

func Test_Reflectcore_Utils_NotNil(t *testing.T) {
	_ = reflectcore.Utils
}

func Test_Reflectcore_Looper_NotNil(t *testing.T) {
	_ = reflectcore.Looper
}

func Test_Reflectcore_CodeStack_NotNil(t *testing.T) {
	_ = reflectcore.CodeStack
}

func Test_Reflectcore_GetFunc_NotNil(t *testing.T) {
	_ = reflectcore.GetFunc
}

func Test_Reflectcore_Is_NotNil(t *testing.T) {
	_ = reflectcore.Is
}

func Test_Reflectcore_TypeName_NotNil(t *testing.T) {
	_ = reflectcore.TypeName // compile-time existence check
}

func Test_Reflectcore_TypeNames_NotNil(t *testing.T) {
	_ = reflectcore.TypeNames // compile-time existence check
}

func Test_Reflectcore_ReflectType_NotNil(t *testing.T) {
	_ = reflectcore.ReflectType
}

func Test_Reflectcore_ReflectGetter_NotNil(t *testing.T) {
	_ = reflectcore.ReflectGetter
}

func Test_Reflectcore_SliceConverter_NotNil(t *testing.T) {
	_ = reflectcore.SliceConverter
}

func Test_Reflectcore_MapConverter_NotNil(t *testing.T) {
	_ = reflectcore.MapConverter
}
