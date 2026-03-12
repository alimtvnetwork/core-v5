package isanytests

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/isany"
)

// ── Conclusive all branches ──

func Test_Conclusive_BothNil_Cov2(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(nil, nil)
	if !isEqual || !isConcl {
		t.Error("both nil should be equal+conclusive")
	}
}

func Test_Conclusive_SamePointer_Cov2(t *testing.T) {
	v := 42
	isEqual, isConcl := isany.Conclusive(&v, &v)
	if !isEqual || !isConcl {
		t.Error("same pointer should be equal+conclusive")
	}
}

func Test_Conclusive_LeftNil_Cov2(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(nil, 42)
	if isEqual || !isConcl {
		t.Error("left nil should be notEqual+conclusive")
	}
}

func Test_Conclusive_RightNil_Cov2(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(42, nil)
	if isEqual || !isConcl {
		t.Error("right nil should be notEqual+conclusive")
	}
}

func Test_Conclusive_BothTypedNilSameType_Cov2(t *testing.T) {
	var a, b *int
	isEqual, isConcl := isany.Conclusive(a, b)
	if !isEqual || !isConcl {
		t.Error("both typed nil same type should be equal+conclusive")
	}
}

func Test_Conclusive_OneTypedNilOtherNot_Cov2(t *testing.T) {
	var a *int
	v := 42
	isEqual, isConcl := isany.Conclusive(a, &v)
	if isEqual || !isConcl {
		t.Error("one typed nil should be notEqual+conclusive")
	}
}

func Test_Conclusive_DiffTypes_Cov2(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(42, "hello")
	if isEqual || !isConcl {
		t.Error("different types should be notEqual+conclusive")
	}
}

func Test_Conclusive_Inconclusive_Cov2(t *testing.T) {
	isEqual, isConcl := isany.Conclusive(42, 43)
	if isEqual || isConcl {
		t.Error("same types diff values should be inconclusive")
	}
}

// ── ReflectValueNull ──

func Test_ReflectValueNull_NilPtr_Cov2(t *testing.T) {
	var ptr *int
	rv := reflect.ValueOf(ptr)
	if !isany.ReflectValueNull(rv) {
		t.Error("nil ptr should be null")
	}
}

func Test_ReflectValueNull_NonNilPtr_Cov2(t *testing.T) {
	v := 42
	rv := reflect.ValueOf(&v)
	if isany.ReflectValueNull(rv) {
		t.Error("non-nil ptr should not be null")
	}
}

func Test_ReflectValueNull_NilSlice_Cov2(t *testing.T) {
	var s []string
	rv := reflect.ValueOf(s)
	if !isany.ReflectValueNull(rv) {
		t.Error("nil slice should be null")
	}
}

func Test_ReflectValueNull_NilMap_Cov2(t *testing.T) {
	var m map[string]int
	rv := reflect.ValueOf(m)
	if !isany.ReflectValueNull(rv) {
		t.Error("nil map should be null")
	}
}

func Test_ReflectValueNull_NilChan_Cov2(t *testing.T) {
	var ch chan int
	rv := reflect.ValueOf(ch)
	if !isany.ReflectValueNull(rv) {
		t.Error("nil chan should be null")
	}
}

func Test_ReflectValueNull_NilFunc_Cov2(t *testing.T) {
	var fn func()
	rv := reflect.ValueOf(fn)
	if !isany.ReflectValueNull(rv) {
		t.Error("nil func should be null")
	}
}

func Test_ReflectValueNull_NilUnsafePtr_Cov2(t *testing.T) {
	rv := reflect.ValueOf(unsafe.Pointer(nil))
	if !isany.ReflectValueNull(rv) {
		t.Error("nil unsafe ptr should be null")
	}
}

func Test_ReflectValueNull_IntKind_Cov2(t *testing.T) {
	rv := reflect.ValueOf(42)
	if isany.ReflectValueNull(rv) {
		t.Error("int should not be null")
	}
}

// ── ReflectNull extended kinds ──

func Test_ReflectNull_NilMap_Cov2(t *testing.T) {
	var m map[string]int
	if !isany.ReflectNull(m) {
		t.Error("nil map should be null")
	}
}

func Test_ReflectNull_NonNilMap_Cov2(t *testing.T) {
	m := map[string]int{"a": 1}
	if isany.ReflectNull(m) {
		t.Error("non-nil map should not be null")
	}
}

func Test_ReflectNull_Int_Cov2(t *testing.T) {
	if isany.ReflectNull(42) {
		t.Error("int should not be null")
	}
}

// ── FloatingPointTypeRv ──

func Test_FloatingPointTypeRv_Float32_Cov2(t *testing.T) {
	rv := reflect.ValueOf(float32(3.14))
	if !isany.FloatingPointTypeRv(rv) {
		t.Error("float32 should be floating point")
	}
}

func Test_FloatingPointTypeRv_Int_Cov2(t *testing.T) {
	rv := reflect.ValueOf(42)
	if isany.FloatingPointTypeRv(rv) {
		t.Error("int should not be floating point")
	}
}

// ── NumberTypeRv ──

func Test_NumberTypeRv_Int8_Cov2(t *testing.T) {
	rv := reflect.ValueOf(int8(1))
	if !isany.NumberTypeRv(rv) {
		t.Error("int8 should be number")
	}
}

func Test_NumberTypeRv_Uint32_Cov2(t *testing.T) {
	rv := reflect.ValueOf(uint32(1))
	if !isany.NumberTypeRv(rv) {
		t.Error("uint32 should be number")
	}
}

func Test_NumberTypeRv_String_Cov2(t *testing.T) {
	rv := reflect.ValueOf("hello")
	if isany.NumberTypeRv(rv) {
		t.Error("string should not be number")
	}
}

// ── PositiveIntegerTypeRv ──

func Test_PositiveIntegerTypeRv_Uint8_Cov2(t *testing.T) {
	rv := reflect.ValueOf(uint8(1))
	if !isany.PositiveIntegerTypeRv(rv) {
		t.Error("uint8 should be positive int")
	}
}

func Test_PositiveIntegerTypeRv_Uint16_Cov2(t *testing.T) {
	rv := reflect.ValueOf(uint16(1))
	if !isany.PositiveIntegerTypeRv(rv) {
		t.Error("uint16 should be positive int")
	}
}

func Test_PositiveIntegerTypeRv_Uint32_Cov2(t *testing.T) {
	rv := reflect.ValueOf(uint32(1))
	if !isany.PositiveIntegerTypeRv(rv) {
		t.Error("uint32 should be positive int")
	}
}

func Test_PositiveIntegerTypeRv_Uint64_Cov2(t *testing.T) {
	rv := reflect.ValueOf(uint64(1))
	if !isany.PositiveIntegerTypeRv(rv) {
		t.Error("uint64 should be positive int")
	}
}

func Test_PositiveIntegerTypeRv_Int_Cov2(t *testing.T) {
	rv := reflect.ValueOf(42)
	if isany.PositiveIntegerTypeRv(rv) {
		t.Error("int should not be positive int type")
	}
}

// ── PrimitiveTypeRv ──

func Test_PrimitiveTypeRv_Uintptr_Cov2(t *testing.T) {
	if !isany.PrimitiveTypeRv(reflect.Uintptr) {
		t.Error("Uintptr should be primitive")
	}
}

func Test_PrimitiveTypeRv_Struct_Cov2(t *testing.T) {
	if isany.PrimitiveTypeRv(reflect.Struct) {
		t.Error("Struct should not be primitive")
	}
}

// ── DeepEqualAllItems edge cases ──

func Test_DeepEqualAllItems_Empty_Cov2(t *testing.T) {
	if !isany.DeepEqualAllItems() {
		t.Error("empty should be true")
	}
}

func Test_DeepEqualAllItems_Single_Cov2(t *testing.T) {
	if !isany.DeepEqualAllItems(42) {
		t.Error("single should be true")
	}
}

func Test_DeepEqualAllItems_TwoEqual_Cov2(t *testing.T) {
	if !isany.DeepEqualAllItems(42, 42) {
		t.Error("two equal should be true")
	}
}

func Test_DeepEqualAllItems_TwoDifferent_Cov2(t *testing.T) {
	if isany.DeepEqualAllItems(42, 43) {
		t.Error("two different should be false")
	}
}

func Test_DeepEqualAllItems_ThreeMixed_Cov2(t *testing.T) {
	if isany.DeepEqualAllItems(42, 42, 43) {
		t.Error("three mixed should be false")
	}
}

func Test_DeepEqualAllItems_ThreeEqual_Cov2(t *testing.T) {
	if !isany.DeepEqualAllItems(42, 42, 42) {
		t.Error("three equal should be true")
	}
}

// ── DefinedItems full coverage ──

func Test_DefinedItems_Empty_Cov2(t *testing.T) {
	isAll, items := isany.DefinedItems()
	if isAll {
		t.Error("empty should not be all defined")
	}
	if items != nil {
		t.Error("empty should return nil")
	}
}

func Test_DefinedItems_AllDefined_Cov2(t *testing.T) {
	isAll, items := isany.DefinedItems(1, "hello", 3.14)
	if !isAll {
		t.Error("all defined should be true")
	}
	if len(items) != 3 {
		t.Error("expected 3 items")
	}
}

func Test_DefinedItems_SomeNil_Cov2(t *testing.T) {
	isAll, items := isany.DefinedItems(nil, 42, nil, "hello")
	if isAll {
		t.Error("should not be all defined")
	}
	if len(items) != 2 {
		t.Errorf("expected 2 defined items, got %d", len(items))
	}
}

// ── DefinedAllOf empty ──

func Test_DefinedAllOf_Empty_Cov2(t *testing.T) {
	if isany.DefinedAllOf() {
		t.Error("empty should be false")
	}
}

// ── DefinedAnyOf empty ──

func Test_DefinedAnyOf_Empty_Cov2(t *testing.T) {
	if isany.DefinedAnyOf() {
		t.Error("empty should be false")
	}
}

func Test_DefinedAnyOf_AllDefined_Cov2(t *testing.T) {
	if !isany.DefinedAnyOf(42, "hello") {
		t.Error("all defined should be true")
	}
}

func Test_DefinedAnyOf_AllNil_Cov2(t *testing.T) {
	if isany.DefinedAnyOf(nil, nil) {
		t.Error("all nil should be false")
	}
}

// ── Null extended types ──

func Test_Null_UnsafePointer_Cov2(t *testing.T) {
	if !isany.Null(unsafe.Pointer(nil)) {
		t.Error("nil unsafe pointer should be null")
	}
}

// ── NumberType extended ──

func Test_NumberType_Uint_Cov2(t *testing.T) {
	if !isany.NumberType(uint(1)) {
		t.Error("uint should be number")
	}
}

func Test_NumberType_Int8_Cov2(t *testing.T) {
	if !isany.NumberType(int8(1)) {
		t.Error("int8 should be number")
	}
}

func Test_NumberType_Float32_Cov2(t *testing.T) {
	if !isany.NumberType(float32(1.0)) {
		t.Error("float32 should be number")
	}
}

// ── PositiveIntegerType extended ──

func Test_PositiveIntegerType_Uint8_Cov2(t *testing.T) {
	if !isany.PositiveIntegerType(uint8(1)) {
		t.Error("uint8 should be positive integer")
	}
}

func Test_PositiveIntegerType_Uint16_Cov2(t *testing.T) {
	if !isany.PositiveIntegerType(uint16(1)) {
		t.Error("uint16 should be positive integer")
	}
}

func Test_PositiveIntegerType_Uint32_Cov2(t *testing.T) {
	if !isany.PositiveIntegerType(uint32(1)) {
		t.Error("uint32 should be positive integer")
	}
}

func Test_PositiveIntegerType_Uint64_Cov2(t *testing.T) {
	if !isany.PositiveIntegerType(uint64(1)) {
		t.Error("uint64 should be positive integer")
	}
}

// ── FloatingPointType float32 ──

func Test_FloatingPointType_Float32_Cov2(t *testing.T) {
	if !isany.FloatingPointType(float32(1.0)) {
		t.Error("float32 should be floating point")
	}
}

// ── PrimitiveType extended ──

func Test_PrimitiveType_Uint_Cov2(t *testing.T) {
	if !isany.PrimitiveType(uint(1)) {
		t.Error("uint should be primitive")
	}
}

func Test_PrimitiveType_Int8_Cov2(t *testing.T) {
	if !isany.PrimitiveType(int8(1)) {
		t.Error("int8 should be primitive")
	}
}

func Test_PrimitiveType_Slice_Cov2(t *testing.T) {
	if isany.PrimitiveType([]int{}) {
		t.Error("slice should not be primitive")
	}
}
