package reflectinternal

import (
	"reflect"
	"testing"
)

func TestIsChecker_Null(t *testing.T) {
	if !Is.Null(nil) {
		t.Fatal("nil should be null")
	}
	if Is.Null("hello") {
		t.Fatal("string should not be null")
	}
	if Is.Null(42) {
		t.Fatal("int should not be null")
	}

	var s *string
	if !Is.Null(s) {
		t.Fatal("nil pointer should be null")
	}

	var m map[string]int
	if !Is.Null(m) {
		t.Fatal("nil map should be null")
	}
}

func TestIsChecker_NotNull(t *testing.T) {
	if Is.NotNull(nil) {
		t.Fatal("nil should be null")
	}
	if !Is.NotNull("hello") {
		t.Fatal("string should not be null")
	}
}

func TestIsChecker_Defined(t *testing.T) {
	if Is.Defined(nil) {
		t.Fatal("nil should not be defined")
	}
	if !Is.Defined("hello") {
		t.Fatal("string should be defined")
	}
}

func TestIsChecker_NullRv(t *testing.T) {
	if !Is.NullRv(reflect.Value{}) {
		t.Fatal("invalid rv should be null")
	}
	if Is.NullRv(reflect.ValueOf(42)) {
		t.Fatal("int rv should not be null")
	}
}

func TestIsChecker_AnyEqual(t *testing.T) {
	if !Is.AnyEqual(1, 1) {
		t.Fatal("should be equal")
	}
	if Is.AnyEqual(1, 2) {
		t.Fatal("should not be equal")
	}
	if !Is.AnyEqual(nil, nil) {
		t.Fatal("nil should be equal")
	}
	if Is.AnyEqual(nil, 1) {
		t.Fatal("nil vs 1 should not be equal")
	}

	// slice - uncomparable
	if Is.AnyEqual([]int{1}, []int{2}) {
		t.Fatal("should not be equal")
	}
	if !Is.AnyEqual([]int{1}, []int{1}) {
		t.Fatal("should be equal via DeepEqual")
	}
}

func TestIsChecker_Conclusive(t *testing.T) {
	eq, con := Is.Conclusive(1, 1)
	if !eq || !con {
		t.Fatal("same value should be conclusive equal")
	}

	eq2, con2 := Is.Conclusive(nil, nil)
	if !eq2 || !con2 {
		t.Fatal("nil nil should be conclusive equal")
	}

	eq3, con3 := Is.Conclusive(nil, 1)
	if eq3 || !con3 {
		t.Fatal("nil vs value should be conclusive not equal")
	}

	eq4, con4 := Is.Conclusive(1, "1")
	if eq4 || !con4 {
		t.Fatal("different types should be conclusive not equal")
	}

	// both null pointers
	var a *string
	var b *string
	eq5, con5 := Is.Conclusive(a, b)
	if !eq5 || !con5 {
		t.Fatal("both null ptrs should be equal")
	}
}

func TestIsChecker_Func(t *testing.T) {
	fn := func() {}
	if !Is.Func(fn) {
		t.Fatal("should be func")
	}
	if !Is.Func(nil) {
		t.Fatal("nil should return true")
	}
	if Is.NotFunc(fn) {
		t.Fatal("should not be not func")
	}
}

func TestIsChecker_FuncTypeOf(t *testing.T) {
	fn := func() {}
	if !Is.FuncTypeOf(reflect.TypeOf(fn)) {
		t.Fatal("should be func")
	}
	if Is.FuncTypeOf(reflect.TypeOf(42)) {
		t.Fatal("int should not be func")
	}
}

func TestIsChecker_SliceOrArray(t *testing.T) {
	_ = Is.SliceOrArray(nil)
	_ = Is.SliceOrArray([]int{1})
}

func TestIsChecker_SliceOrArrayOf(t *testing.T) {
	if !Is.SliceOrArrayOf(reflect.TypeOf([]int{})) {
		t.Fatal("should be slice")
	}
	if Is.SliceOrArrayOf(reflect.TypeOf(42)) {
		t.Fatal("int should not be slice")
	}
}

func TestIsChecker_Number(t *testing.T) {
	if !Is.Number(42) {
		t.Fatal("should be number")
	}
	if Is.Number("hello") {
		t.Fatal("string should not be number")
	}
}

func TestIsChecker_NumberKind(t *testing.T) {
	if !Is.NumberKind(reflect.Int) {
		t.Fatal("int should be number")
	}
	if Is.NumberKind(reflect.String) {
		t.Fatal("string should not be number")
	}
}

func TestIsChecker_String(t *testing.T) {
	if !Is.String("hello") {
		t.Fatal("should be string")
	}
	if Is.String(42) {
		t.Fatal("int should not be string")
	}
}

func TestIsChecker_Pointer(t *testing.T) {
	s := "hello"
	if !Is.Pointer(&s) {
		t.Fatal("should be pointer")
	}
	if Is.Pointer(s) {
		t.Fatal("string should not be pointer")
	}
}

func TestIsChecker_Function(t *testing.T) {
	fn := func() {}
	if !Is.Function(fn) {
		t.Fatal("should be func")
	}
}

func TestIsChecker_Boolean(t *testing.T) {
	if !Is.Boolean(true) {
		t.Fatal("should be bool")
	}
	if Is.Boolean(42) {
		t.Fatal("int should not be bool")
	}
}

func TestIsChecker_Primitive(t *testing.T) {
	if !Is.Primitive(42) {
		t.Fatal("int should be primitive")
	}
	if !Is.Primitive("hello") {
		t.Fatal("string should be primitive")
	}
}

func TestIsChecker_PrimitiveKind(t *testing.T) {
	if !Is.PrimitiveKind(reflect.Int) {
		t.Fatal("int should be primitive")
	}
	if Is.PrimitiveKind(reflect.Map) {
		t.Fatal("map should not be primitive")
	}
}

func TestIsChecker_Zero(t *testing.T) {
	if !Is.Zero(nil) {
		t.Fatal("nil should be zero")
	}
	if !Is.Zero(0) {
		t.Fatal("0 should be zero")
	}
	if Is.Zero(42) {
		t.Fatal("42 should not be zero")
	}
	if !Is.Zero("") {
		t.Fatal("empty string should be zero")
	}
}

func TestIsChecker_ZeroRv(t *testing.T) {
	// struct
	type s struct{ A int }
	if !Is.ZeroRv(reflect.ValueOf(s{})) {
		t.Fatal("zero struct should be zero")
	}
	if Is.ZeroRv(reflect.ValueOf(s{A: 1})) {
		t.Fatal("non-zero struct should not be zero")
	}

	// array
	if !Is.ZeroRv(reflect.ValueOf([2]int{})) {
		t.Fatal("zero array should be zero")
	}
}

func TestIsChecker_Struct(t *testing.T) {
	type s struct{ A int }
	if !Is.Struct(s{}) {
		t.Fatal("should be struct")
	}
	if !Is.Struct(&s{}) {
		t.Fatal("ptr to struct should be struct")
	}
	if Is.Struct(42) {
		t.Fatal("int should not be struct")
	}
}

func TestIsChecker_StructRv(t *testing.T) {
	type s struct{ A int }
	if !Is.StructRv(reflect.ValueOf(s{})) {
		t.Fatal("should be struct")
	}
}

func TestIsChecker_Interface(t *testing.T) {
	_ = Is.Interface(42)
}

func TestIsChecker_InterfaceRv(t *testing.T) {
	_ = Is.InterfaceRv(reflect.ValueOf(42))
}
