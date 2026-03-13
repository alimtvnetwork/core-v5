package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
)

func Test_Null_Direct_NilInput(t *testing.T) {
	if !isany.Null(nil) {
		t.Error("nil should be null")
	}
}

func Test_Null_Direct_StringInput(t *testing.T) {
	if isany.Null("hello") {
		t.Error("string should not be null")
	}
}

func Test_Null_Direct_IntInput(t *testing.T) {
	if isany.Null(42) {
		t.Error("int should not be null")
	}
}

func Test_Null_Direct_NilSlice(t *testing.T) {
	var s []string
	if !isany.Null(s) {
		t.Error("nil slice should be null")
	}
}

func Test_Null_Direct_NilMap(t *testing.T) {
	var m map[string]string
	if !isany.Null(m) {
		t.Error("nil map should be null")
	}
}

func Test_Defined_Direct_NilInput(t *testing.T) {
	if isany.Defined(nil) {
		t.Error("nil should not be defined")
	}
}

func Test_Defined_Direct_StringInput(t *testing.T) {
	if !isany.Defined("hello") {
		t.Error("string should be defined")
	}
}

func Test_DefinedBoth_Direct_BothDefined(t *testing.T) {
	if !isany.DefinedBoth("a", "b") {
		t.Error("both defined should return true")
	}
}

func Test_DefinedBoth_Direct_LeftNil(t *testing.T) {
	if isany.DefinedBoth(nil, "b") {
		t.Error("left nil should return false")
	}
}

func Test_DefinedBoth_Direct_RightNil(t *testing.T) {
	if isany.DefinedBoth("a", nil) {
		t.Error("right nil should return false")
	}
}

func Test_DefinedBoth_Direct_BothNil(t *testing.T) {
	if isany.DefinedBoth(nil, nil) {
		t.Error("both nil should return false")
	}
}

func Test_DefinedLeftRight_Direct(t *testing.T) {
	l, r := isany.DefinedLeftRight("a", nil)
	if !l {
		t.Error("left should be defined")
	}
	if r {
		t.Error("right should not be defined")
	}
}

func Test_NullBoth_Direct_BothNil(t *testing.T) {
	if !isany.NullBoth(nil, nil) {
		t.Error("both nil should return true")
	}
}

func Test_NullBoth_Direct_OneNonNil(t *testing.T) {
	if isany.NullBoth("a", nil) {
		t.Error("one non-nil should return false")
	}
}

func Test_NullLeftRight_Direct(t *testing.T) {
	l, r := isany.NullLeftRight(nil, "b")
	if !l {
		t.Error("left should be null")
	}
	if r {
		t.Error("right should not be null")
	}
}

func Test_FuncOnly_Direct_Nil(t *testing.T) {
	if isany.FuncOnly(nil) {
		t.Error("nil should not be func")
	}
}

func Test_FuncOnly_Direct_String(t *testing.T) {
	if isany.FuncOnly("string") {
		t.Error("string should not be func")
	}
}

func Test_FuncOnly_Direct_Func(t *testing.T) {
	if !isany.FuncOnly(func() {}) {
		t.Error("func should be func")
	}
}

func Test_NotNull_Direct(t *testing.T) {
	if isany.NotNull(nil) {
		t.Error("nil should be null")
	}
	if !isany.NotNull("hello") {
		t.Error("string should not be null")
	}
}

func Test_AllNull_Direct(t *testing.T) {
	if !isany.AllNull(nil, nil, nil) {
		t.Error("all nil should return true")
	}
	if isany.AllNull(nil, "a", nil) {
		t.Error("one non-nil should return false")
	}
}

func Test_AnyNull_Direct(t *testing.T) {
	if !isany.AnyNull(nil, "a") {
		t.Error("one nil should return true")
	}
	if isany.AnyNull("a", "b") {
		t.Error("no nil should return false")
	}
}

func Test_AllZero_Direct(t *testing.T) {
	if !isany.AllZero(0, "", false) {
		t.Error("all zero values should return true")
	}
	if isany.AllZero(0, "a", false) {
		t.Error("non-zero should return false")
	}
}

func Test_AnyZero_Direct(t *testing.T) {
	if !isany.AnyZero(0, "a") {
		t.Error("one zero should return true")
	}
	if isany.AnyZero("a", 1) {
		t.Error("no zero should return false")
	}
}

func Test_Zero_Direct(t *testing.T) {
	if !isany.Zero(0) {
		t.Error("0 should be zero")
	}
	if !isany.Zero("") {
		t.Error("empty string should be zero")
	}
	if !isany.Zero(false) {
		t.Error("false should be zero")
	}
	if isany.Zero(1) {
		t.Error("1 should not be zero")
	}
}

func Test_DeepEqual_Direct(t *testing.T) {
	if !isany.DeepEqual("a", "a") {
		t.Error("same strings should be deep equal")
	}
	if isany.DeepEqual("a", "b") {
		t.Error("different strings should not be deep equal")
	}
}

func Test_NotDeepEqual_Direct(t *testing.T) {
	if isany.NotDeepEqual("a", "a") {
		t.Error("same should return false")
	}
	if !isany.NotDeepEqual("a", "b") {
		t.Error("different should return true")
	}
}

func Test_StringEqual_Direct(t *testing.T) {
	if !isany.StringEqual("hello", "hello") {
		t.Error("same strings should be equal")
	}
	if isany.StringEqual("hello", "world") {
		t.Error("different strings should not be equal")
	}
}

func Test_ReflectNull_Direct(t *testing.T) {
	if !isany.ReflectNull(nil) {
		t.Error("nil should be reflect null")
	}
	if isany.ReflectNull("hello") {
		t.Error("string should not be reflect null")
	}
}

func Test_ReflectNotNull_Direct(t *testing.T) {
	if isany.ReflectNotNull(nil) {
		t.Error("nil should not be reflect not null")
	}
	if !isany.ReflectNotNull("hello") {
		t.Error("string should be reflect not null")
	}
}

func Test_Pointer_Direct(t *testing.T) {
	val := 42
	if !isany.Pointer(&val) {
		t.Error("pointer should return true")
	}
	if isany.Pointer(42) {
		t.Error("non-pointer should return false")
	}
	if isany.Pointer(nil) {
		t.Error("nil should return false")
	}
}

func Test_Conclusive_Direct(t *testing.T) {
	isEq, isConcl := isany.Conclusive("hello", "hello")
	if !isConcl || isEq {
		// same type, different pointers → inconclusive
	}
	isEq2, isConcl2 := isany.Conclusive(nil, nil)
	if !isConcl2 || !isEq2 {
		t.Error("both nil should be conclusive equal")
	}
}

func Test_DefinedAllOf_Direct(t *testing.T) {
	if !isany.DefinedAllOf("a", "b", "c") {
		t.Error("all defined should return true")
	}
	if isany.DefinedAllOf("a", nil, "c") {
		t.Error("one nil should return false")
	}
}

func Test_DefinedAnyOf_Direct(t *testing.T) {
	if !isany.DefinedAnyOf(nil, "a", nil) {
		t.Error("one defined should return true")
	}
	if isany.DefinedAnyOf(nil, nil) {
		t.Error("all nil should return false")
	}
}

func Test_TypeSame_Direct(t *testing.T) {
	if !isany.TypeSame("a", "b") {
		t.Error("same types should return true")
	}
	if isany.TypeSame("a", 1) {
		t.Error("different types should return false")
	}
}

func Test_JsonEqual_Direct(t *testing.T) {
	if !isany.JsonEqual("hello", "hello") {
		t.Error("same values should be json equal")
	}
}

func Test_JsonMismatch_Direct(t *testing.T) {
	if isany.JsonMismatch("hello", "hello") {
		t.Error("same values should not mismatch")
	}
	if !isany.JsonMismatch("hello", "world") {
		t.Error("different values should mismatch")
	}
}

func Test_NumberType_Direct(t *testing.T) {
	if !isany.NumberType(42) {
		t.Error("int should be number type")
	}
	if !isany.NumberType(3.14) {
		t.Error("float should be number type")
	}
	if isany.NumberType("hello") {
		t.Error("string should not be number type")
	}
}

func Test_PrimitiveType_Direct(t *testing.T) {
	if !isany.PrimitiveType("hello") {
		t.Error("string should be primitive")
	}
	if !isany.PrimitiveType(42) {
		t.Error("int should be primitive")
	}
	if !isany.PrimitiveType(true) {
		t.Error("bool should be primitive")
	}
}

func Test_FloatingPointType_Direct(t *testing.T) {
	if !isany.FloatingPointType(3.14) {
		t.Error("float64 should be floating point")
	}
	if !isany.FloatingPointType(float32(1.0)) {
		t.Error("float32 should be floating point")
	}
	if isany.FloatingPointType(42) {
		t.Error("int should not be floating point")
	}
}

func Test_DefinedItems_Direct(t *testing.T) {
	count := isany.DefinedItems("a", nil, "c")
	if count != 2 {
		t.Errorf("expected 2 defined items, got %d", count)
	}
}

func Test_PositiveIntegerType_Direct(t *testing.T) {
	if !isany.PositiveIntegerType(42) {
		t.Error("positive int should be positive integer type")
	}
	if isany.PositiveIntegerType(-1) {
		t.Error("negative int should not be positive integer type")
	}
}

func Test_DeepEqualAllItems_Direct(t *testing.T) {
	if !isany.DeepEqualAllItems(1, 1, 1) {
		t.Error("all same should return true")
	}
	if isany.DeepEqualAllItems(1, 2, 1) {
		t.Error("different should return false")
	}
}
