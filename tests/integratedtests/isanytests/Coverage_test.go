package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
)

// ── Null / NotNull / Defined ──

func Test_Null_Coverage(t *testing.T) {
	if !isany.Null(nil) {
		t.Error("nil should be null")
	}
	if isany.Null(42) {
		t.Error("42 should not be null")
	}
	if isany.Null("hello") {
		t.Error("string should not be null")
	}

	var s []string
	if !isany.Null(s) {
		t.Error("nil slice should be null")
	}

	var m map[string]int
	if !isany.Null(m) {
		t.Error("nil map should be null")
	}

	var fn func()
	if !isany.Null(fn) {
		t.Error("nil func should be null")
	}

	var ch chan int
	if !isany.Null(ch) {
		t.Error("nil chan should be null")
	}

	var ptr *int
	if !isany.Null(ptr) {
		t.Error("nil pointer should be null")
	}

	val := 42
	if isany.Null(&val) {
		t.Error("non-nil pointer should not be null")
	}
}

func Test_NotNull_Coverage(t *testing.T) {
	if !isany.NotNull(42) {
		t.Error("42 should be not null")
	}
	if isany.NotNull(nil) {
		t.Error("nil should not be not null")
	}
}

func Test_Defined_Coverage(t *testing.T) {
	if !isany.Defined(42) {
		t.Error("42 should be defined")
	}
	if isany.Defined(nil) {
		t.Error("nil should not be defined")
	}
}

// ── AllNull / AnyNull ──

func Test_AllNull_Coverage(t *testing.T) {
	if !isany.AllNull(nil, nil, nil) {
		t.Error("all nils should be AllNull")
	}
	if isany.AllNull(nil, 42, nil) {
		t.Error("mixed should not be AllNull")
	}
	if !isany.AllNull() {
		t.Error("empty should be AllNull (vacuous truth)")
	}
}

func Test_AnyNull_Coverage(t *testing.T) {
	if !isany.AnyNull(nil, 42) {
		t.Error("should be AnyNull with nil present")
	}
	if isany.AnyNull() {
		t.Error("empty should not be AnyNull")
	}
	if isany.AnyNull(42, "hello") {
		t.Error("no nils should not be AnyNull")
	}
}

// ── Zero / AllZero / AnyZero ──

func Test_Zero_Coverage(t *testing.T) {
	if !isany.Zero(0) {
		t.Error("0 should be zero")
	}
	if !isany.Zero("") {
		t.Error("empty string should be zero")
	}
	if !isany.Zero(nil) {
		t.Error("nil should be zero")
	}
	if isany.Zero(42) {
		t.Error("42 should not be zero")
	}
}

func Test_AllZero_Coverage(t *testing.T) {
	if !isany.AllZero(0, "", nil) {
		t.Error("all zeros should be AllZero")
	}
	if isany.AllZero(0, 1) {
		t.Error("mixed should not be AllZero")
	}
	if !isany.AllZero() {
		t.Error("empty should be AllZero")
	}
}

func Test_AnyZero_Coverage(t *testing.T) {
	if !isany.AnyZero(0, 42) {
		t.Error("should be AnyZero")
	}
	if isany.AnyZero(42, "hello") {
		t.Error("no zeros should not be AnyZero")
	}
	if !isany.AnyZero() {
		t.Error("empty should be AnyZero")
	}
}

// ── DeepEqual / NotDeepEqual ──

func Test_DeepEqual_Coverage(t *testing.T) {
	if !isany.DeepEqual(42, 42) {
		t.Error("42 == 42")
	}
	if isany.DeepEqual(42, 43) {
		t.Error("42 != 43")
	}
	if !isany.DeepEqual(nil, nil) {
		t.Error("nil == nil")
	}
}

func Test_NotDeepEqual_Coverage(t *testing.T) {
	if !isany.NotDeepEqual(42, 43) {
		t.Error("42 != 43")
	}
	if isany.NotDeepEqual(42, 42) {
		t.Error("42 == 42")
	}
}

// ── DeepEqualAllItems ──

func Test_DeepEqualAllItems_Coverage(t *testing.T) {
	if !isany.DeepEqualAllItems(42, 42, 42) {
		t.Error("all 42 should be equal")
	}
	if isany.DeepEqualAllItems(42, 43, 42) {
		t.Error("mixed should not be equal")
	}
}

// ── DefinedBoth / NullBoth / DefinedAllOf / DefinedAnyOf ──

func Test_DefinedBoth_Coverage(t *testing.T) {
	if !isany.DefinedBoth(42, "hello") {
		t.Error("both defined should be true")
	}
	if isany.DefinedBoth(nil, "hello") {
		t.Error("one nil should be false")
	}
	if isany.DefinedBoth(nil, nil) {
		t.Error("both nil should be false")
	}
}

func Test_NullBoth_Coverage(t *testing.T) {
	if !isany.NullBoth(nil, nil) {
		t.Error("both nil should be true")
	}
	if isany.NullBoth(nil, 42) {
		t.Error("one non-nil should be false")
	}
}

func Test_DefinedAllOf_Coverage(t *testing.T) {
	if !isany.DefinedAllOf(42, "hello") {
		t.Error("all defined should be true")
	}
	if isany.DefinedAllOf(42, nil) {
		t.Error("one nil should be false")
	}
}

func Test_DefinedAnyOf_Coverage(t *testing.T) {
	if !isany.DefinedAnyOf(nil, 42) {
		t.Error("one defined should be true")
	}
	if isany.DefinedAnyOf(nil, nil) {
		t.Error("all nil should be false")
	}
}

// ── DefinedItems ──

func Test_DefinedItems_Coverage(t *testing.T) {
	_, items := isany.DefinedItems(nil, 42, nil, "hello")
	if len(items) != 2 {
		t.Errorf("expected 2 defined items, got %d", len(items))
	}
}

// ── DefinedLeftRight ──

func Test_DefinedLeftRight_Coverage(t *testing.T) {
	leftDef, rightDef := isany.DefinedLeftRight(42, nil)
	if !leftDef {
		t.Error("left should be defined")
	}
	if rightDef {
		t.Error("right should not be defined")
	}
}

// ── NullLeftRight ──

func Test_NullLeftRight_Coverage(t *testing.T) {
	leftNull, rightNull := isany.NullLeftRight(nil, 42)
	if !leftNull {
		t.Error("left should be null")
	}
	if rightNull {
		t.Error("right should not be null")
	}
}

// ── StringEqual ──

func Test_StringEqual_Coverage(t *testing.T) {
	if !isany.StringEqual(42, 42) {
		t.Error("same values should be string equal")
	}
	if isany.StringEqual(42, 43) {
		t.Error("different values should not be string equal")
	}
}

// ── JsonEqual / JsonMismatch ──

func Test_JsonEqual_Coverage(t *testing.T) {
	if !isany.JsonEqual("hello", "hello") {
		t.Error("same strings should be json equal")
	}
	if isany.JsonEqual("hello", "world") {
		t.Error("different strings should not be json equal")
	}
	if !isany.JsonEqual(42, 42) {
		t.Error("same ints should be json equal")
	}
	if isany.JsonEqual(42, 43) {
		t.Error("different ints should not be json equal")
	}

	type s struct{ A int }
	if !isany.JsonEqual(s{1}, s{1}) {
		t.Error("same structs should be json equal")
	}
	if isany.JsonEqual(s{1}, s{2}) {
		t.Error("different structs should not be json equal")
	}
}

func Test_JsonMismatch_Coverage(t *testing.T) {
	if !isany.JsonMismatch("hello", "world") {
		t.Error("different should mismatch")
	}
	if isany.JsonMismatch("hello", "hello") {
		t.Error("same should not mismatch")
	}
}

// ── TypeSame ──

func Test_TypeSame_Coverage(t *testing.T) {
	if !isany.TypeSame(42, 43) {
		t.Error("both int should be same type")
	}
	if isany.TypeSame(42, "hello") {
		t.Error("int vs string should not be same type")
	}
}

// ── Pointer / Function / FuncOnly ──

func Test_Pointer_Coverage(t *testing.T) {
	val := 42
	if !isany.Pointer(&val) {
		t.Error("pointer should be pointer")
	}
	if isany.Pointer(42) {
		t.Error("non-pointer should not be pointer")
	}
}

func Test_Function_Coverage(t *testing.T) {
	fn := func() {}
	isFunc, name := isany.Function(fn)
	if !isFunc {
		t.Error("func should be function")
	}
	if name == "" {
		t.Error("func name should not be empty")
	}

	isFunc2, _ := isany.Function(nil)
	if isFunc2 {
		t.Error("nil should not be function")
	}

	isFunc3, _ := isany.Function(42)
	if isFunc3 {
		t.Error("int should not be function")
	}
}

func Test_FuncOnly_Coverage(t *testing.T) {
	fn := func() {}
	if !isany.FuncOnly(fn) {
		t.Error("func should return true")
	}
	if isany.FuncOnly(42) {
		t.Error("int should return false")
	}
}

// ── PrimitiveType / NumberType / FloatingPointType / PositiveIntegerType ──

func Test_PrimitiveType_Coverage(t *testing.T) {
	if !isany.PrimitiveType(42) {
		t.Error("int should be primitive")
	}
	if !isany.PrimitiveType("hello") {
		t.Error("string should be primitive")
	}
	if !isany.PrimitiveType(true) {
		t.Error("bool should be primitive")
	}
	if !isany.PrimitiveType(3.14) {
		t.Error("float should be primitive")
	}

	type s struct{}
	if isany.PrimitiveType(s{}) {
		t.Error("struct should not be primitive")
	}
}

func Test_NumberType_Coverage(t *testing.T) {
	if !isany.NumberType(42) {
		t.Error("int should be number")
	}
	if !isany.NumberType(3.14) {
		t.Error("float should be number")
	}
	if isany.NumberType("hello") {
		t.Error("string should not be number")
	}
}

func Test_FloatingPointType_Coverage(t *testing.T) {
	if !isany.FloatingPointType(3.14) {
		t.Error("float64 should be floating point")
	}
	if isany.FloatingPointType(42) {
		t.Error("int should not be floating point")
	}
}

func Test_PositiveIntegerType_Coverage(t *testing.T) {
	if !isany.PositiveIntegerType(uint(42)) {
		t.Error("uint should be positive integer")
	}
	if isany.PositiveIntegerType(42) {
		t.Error("signed int should not be positive integer type")
	}
}

// ── Conclusive ──

func Test_Conclusive_Coverage(t *testing.T) {
	// Arrange — tests with two ints
	_, isConclusive := isany.Conclusive(42, 43)
	if isConclusive {
		t.Error("two same-type non-nil ints should be inconclusive")
	}
}

// ── ReflectNull / ReflectNotNull / ReflectValueNull ──

func Test_ReflectNull_Coverage(t *testing.T) {
	var ptr *int
	if !isany.ReflectNull(ptr) {
		t.Error("nil ptr should be reflect null")
	}

	val := 42
	if isany.ReflectNull(&val) {
		t.Error("non-nil ptr should not be reflect null")
	}
}

func Test_ReflectNotNull_Coverage(t *testing.T) {
	val := 42
	if !isany.ReflectNotNull(&val) {
		t.Error("non-nil should be reflect not null")
	}
}
