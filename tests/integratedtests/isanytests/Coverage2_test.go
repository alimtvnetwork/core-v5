package isanytests

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── Conclusive all branches ──

func Test_Cov2_Conclusive_BothNil(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(nil, nil)
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": true, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive_BothNil", actual)
}

func Test_Cov2_Conclusive_SamePointer(t *testing.T) {
	v := 42
	isEqual, isConcl := isany.Conclusive(&v, &v)
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": true, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive_SamePointer", actual)
}

func Test_Cov2_Conclusive_LeftNil(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(nil, 42)
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": false, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive_LeftNil", actual)
}

func Test_Cov2_Conclusive_RightNil(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(42, nil)
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": false, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive_RightNil", actual)
}

func Test_Cov2_Conclusive_BothTypedNilSameType(t *testing.T) {
	var a, b *int
	isEqual, isConcl := isany.Conclusive(a, b)
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": true, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive_BothTypedNilSameType", actual)
}

func Test_Cov2_Conclusive_OneTypedNilOtherNot(t *testing.T) {
	var a *int
	v := 42
	isEqual, isConcl := isany.Conclusive(a, &v)
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": false, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive_OneTypedNilOtherNot", actual)
}

func Test_Cov2_Conclusive_DiffTypes(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(42, "hello")
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": false, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive_DiffTypes", actual)
}

func Test_Cov2_Conclusive_Inconclusive(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(42, 43)
	actual := args.Map{"isEqual": isEqual, "isConcl": isConcl}
	expected := args.Map{"isEqual": false, "isConcl": false}
	expected.ShouldBeEqual(t, 0, "Conclusive_Inconclusive", actual)
}

// ── ReflectValueNull ──

func Test_Cov2_ReflectValueNull_NilPtr(t *testing.T) {
	var ptr *int
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(ptr))}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilPtr", actual)
}

func Test_Cov2_ReflectValueNull_NonNilPtr(t *testing.T) {
	v := 42
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(&v))}
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NonNilPtr", actual)
}

func Test_Cov2_ReflectValueNull_NilSlice(t *testing.T) {
	var s []string
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(s))}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilSlice", actual)
}

func Test_Cov2_ReflectValueNull_NilMap(t *testing.T) {
	var m map[string]int
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(m))}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilMap", actual)
}

func Test_Cov2_ReflectValueNull_NilChan(t *testing.T) {
	var ch chan int
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(ch))}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilChan", actual)
}

func Test_Cov2_ReflectValueNull_NilFunc(t *testing.T) {
	var fn func()
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(fn))}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilFunc", actual)
}

func Test_Cov2_ReflectValueNull_NilUnsafePtr(t *testing.T) {
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(unsafe.Pointer(nil)))}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_NilUnsafePtr", actual)
}

func Test_Cov2_ReflectValueNull_IntKind(t *testing.T) {
	actual := args.Map{"isNull": isany.ReflectValueNull(reflect.ValueOf(42))}
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull_IntKind", actual)
}

// ── ReflectNull extended kinds ──

func Test_Cov2_ReflectNull_NilMap(t *testing.T) {
	var m map[string]int
	actual := args.Map{"isNull": isany.ReflectNull(m)}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull_NilMap", actual)
}

func Test_Cov2_ReflectNull_NonNilMap(t *testing.T) {
	actual := args.Map{"isNull": isany.ReflectNull(map[string]int{"a": 1})}
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull_NonNilMap", actual)
}

func Test_Cov2_ReflectNull_Int(t *testing.T) {
	actual := args.Map{"isNull": isany.ReflectNull(42)}
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull_Int", actual)
}

// ── FloatingPointTypeRv ──

func Test_Cov2_FloatingPointTypeRv_Float32(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(float32(3.14)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv_Float32", actual)
}

func Test_Cov2_FloatingPointTypeRv_Int(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(42))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv_Int", actual)
}

// ── NumberTypeRv ──

func Test_NumberTypeRv_Int8_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int8(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv_Int8", actual)
}

func Test_NumberTypeRv_Uint32_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint32(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv_Uint32", actual)
}

func Test_NumberTypeRv_String_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf("hello"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv_String", actual)
}

// ── PositiveIntegerTypeRv ──

func Test_PositiveIntegerTypeRv_Uint8_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint8(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint8", actual)
}

func Test_PositiveIntegerTypeRv_Uint16_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint16(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint16", actual)
}

func Test_PositiveIntegerTypeRv_Uint32_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint32(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint32", actual)
}

func Test_PositiveIntegerTypeRv_Uint64_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint64(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Uint64", actual)
}

func Test_PositiveIntegerTypeRv_Int_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(42))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv_Int", actual)
}

// ── PrimitiveTypeRv ──

func Test_PrimitiveTypeRv_Uintptr_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Uintptr)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv_Uintptr", actual)
}

func Test_PrimitiveTypeRv_Struct_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Struct)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv_Struct", actual)
}

// ── DeepEqualAllItems edge cases ──

func Test_DeepEqualAllItems_Cov2(t *testing.T) {
	actual := args.Map{
		"empty":      isany.DeepEqualAllItems(),
		"single":     isany.DeepEqualAllItems(42),
		"twoEqual":   isany.DeepEqualAllItems(42, 42),
		"twoDiff":    isany.DeepEqualAllItems(42, 43),
		"threeMixed": isany.DeepEqualAllItems(42, 42, 43),
		"threeEqual": isany.DeepEqualAllItems(42, 42, 42),
	}
	expected := args.Map{
		"empty":      true,
		"single":     true,
		"twoEqual":   true,
		"twoDiff":    false,
		"threeMixed": false,
		"threeEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems", actual)
}

// ── DefinedItems full coverage ──

func Test_DefinedItems_Empty_Cov2(t *testing.T) {
	isAll, items := isany.DefinedItems()
	actual := args.Map{"isAll": isAll, "isNil": items == nil}
	expected := args.Map{"isAll": false, "isNil": true}
	expected.ShouldBeEqual(t, 0, "DefinedItems_Empty", actual)
}

func Test_DefinedItems_AllDefined_Cov2(t *testing.T) {
	isAll, items := isany.DefinedItems(1, "hello", 3.14)
	actual := args.Map{"isAll": isAll, "len": len(items)}
	expected := args.Map{"isAll": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "DefinedItems_AllDefined", actual)
}

func Test_DefinedItems_SomeNil_Cov2(t *testing.T) {
	isAll, items := isany.DefinedItems(nil, 42, nil, "hello")
	actual := args.Map{"isAll": isAll, "len": len(items)}
	expected := args.Map{"isAll": false, "len": 2}
	expected.ShouldBeEqual(t, 0, "DefinedItems_SomeNil", actual)
}

// ── DefinedAllOf / DefinedAnyOf ──

func Test_DefinedAllOf_Empty_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf_Empty", actual)
}

func Test_DefinedAnyOf_Cov2(t *testing.T) {
	actual := args.Map{
		"empty":      isany.DefinedAnyOf(),
		"allDefined": isany.DefinedAnyOf(42, "hello"),
		"allNil":     isany.DefinedAnyOf(nil, nil),
	}
	expected := args.Map{
		"empty":      false,
		"allDefined": true,
		"allNil":     false,
	}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf", actual)
}

// ── Null / NumberType / PositiveIntegerType / FloatingPointType / PrimitiveType ──

func Test_Null_UnsafePointer_Cov2(t *testing.T) {
	actual := args.Map{"isNull": isany.Null(unsafe.Pointer(nil))}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "Null_UnsafePointer", actual)
}

func Test_NumberType_Extended_Cov2(t *testing.T) {
	actual := args.Map{
		"uint":    isany.NumberType(uint(1)),
		"int8":    isany.NumberType(int8(1)),
		"float32": isany.NumberType(float32(1.0)),
	}
	expected := args.Map{
		"uint":    true,
		"int8":    true,
		"float32": true,
	}
	expected.ShouldBeEqual(t, 0, "NumberType_Extended", actual)
}

func Test_PositiveIntegerType_Extended_Cov2(t *testing.T) {
	actual := args.Map{
		"uint8":  isany.PositiveIntegerType(uint8(1)),
		"uint16": isany.PositiveIntegerType(uint16(1)),
		"uint32": isany.PositiveIntegerType(uint32(1)),
		"uint64": isany.PositiveIntegerType(uint64(1)),
	}
	expected := args.Map{
		"uint8":  true,
		"uint16": true,
		"uint32": true,
		"uint64": true,
	}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType_Extended", actual)
}

func Test_FloatingPointType_Float32_Cov2(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointType(float32(1.0))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointType_Float32", actual)
}

func Test_PrimitiveType_Extended_Cov2(t *testing.T) {
	actual := args.Map{
		"uint":  isany.PrimitiveType(uint(1)),
		"int8":  isany.PrimitiveType(int8(1)),
		"slice": isany.PrimitiveType([]int{}),
	}
	expected := args.Map{
		"uint":  true,
		"int8":  true,
		"slice": false,
	}
	expected.ShouldBeEqual(t, 0, "PrimitiveType_Extended", actual)
}

// keep fmt imported for any future use
var _ = fmt.Sprintf
