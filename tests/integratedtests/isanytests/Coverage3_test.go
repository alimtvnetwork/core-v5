package isanytests

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── ReflectNull all kinds ──

func Test_Cov3_ReflectNull_NilFunc(t *testing.T) {
	var fn func()
	actual := args.Map{"result": isany.ReflectNull(fn)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil func", actual)
}

func Test_Cov3_ReflectNull_NilChan(t *testing.T) {
	var ch chan int
	actual := args.Map{"result": isany.ReflectNull(ch)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil chan", actual)
}

func Test_Cov3_ReflectNull_NilSlice(t *testing.T) {
	var s []string
	actual := args.Map{"result": isany.ReflectNull(s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil slice", actual)
}

func Test_Cov3_ReflectNull_NilUnsafePtr(t *testing.T) {
	actual := args.Map{"result": isany.ReflectNull(unsafe.Pointer(nil))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns nil -- nil unsafe pointer", actual)
}

func Test_Cov3_ReflectNull_String(t *testing.T) {
	actual := args.Map{"result": isany.ReflectNull("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull returns correct value -- string default kind", actual)
}

// ── ReflectNotNull ──

func Test_Cov3_ReflectNotNull_NilPtr(t *testing.T) {
	var ptr *int
	actual := args.Map{"result": isany.ReflectNotNull(ptr)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNotNull returns nil -- nil ptr", actual)
}

// ── ReflectValueNull all kinds ──

func Test_Cov3_ReflectValueNull_NilInterface(t *testing.T) {
	var iface interface{}
	rv := reflect.ValueOf(&iface).Elem()
	actual := args.Map{"result": isany.ReflectValueNull(rv)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns nil -- nil interface", actual)
}

func Test_Cov3_ReflectValueNull_NonNilSlice(t *testing.T) {
	s := []string{"a"}
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(s))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns nil -- non-nil slice", actual)
}

func Test_Cov3_ReflectValueNull_NonNilMap(t *testing.T) {
	m := map[string]int{"a": 1}
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(m))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns nil -- non-nil map", actual)
}

func Test_Cov3_ReflectValueNull_String(t *testing.T) {
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf("hello"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull returns correct value -- string kind", actual)
}

// ── AllZero / AnyZero edge cases ──

func Test_Cov3_AllZero_SingleZero(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- single zero", actual)
}

func Test_Cov3_AnyZero_SingleNonZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero returns non-empty -- single non-zero", actual)
}

// ── AllNull edge cases ──

func Test_Cov3_AllNull_SingleNil(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns nil -- single nil", actual)
}

func Test_Cov3_AllNull_SingleNonNil(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull returns nil -- single non-nil", actual)
}

// ── AnyNull edge cases ──

func Test_Cov3_AnyNull_SingleNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- single nil", actual)
}

func Test_Cov3_AnyNull_SingleNonNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- single non-nil", actual)
}

// ── Zero with various types ──

func Test_Cov3_Zero_Bool(t *testing.T) {
	actual := args.Map{"result": isany.Zero(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero returns non-empty -- false bool", actual)
}

func Test_Cov3_Zero_TrueBool(t *testing.T) {
	actual := args.Map{"result": isany.Zero(true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Zero returns non-empty -- true bool", actual)
}

func Test_Cov3_Zero_Float64(t *testing.T) {
	actual := args.Map{"result": isany.Zero(0.0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero returns correct value -- float64 0.0", actual)
}

// ── StringEqual edge cases ──

func Test_Cov3_StringEqual_DiffTypes(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual(42, "42")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual returns correct value -- different types", actual)
}

// ── Conclusive with two typed nils of different types ──

func Test_Cov3_Conclusive_BothReflectable(t *testing.T) {
	a := 42
	b := 42
	isEq, isConcl := isany.Conclusive(&a, &b)
	actual := args.Map{"isEqual": isEq, "isConclusive": isConcl}
	// Different pointers, same type, so inconclusive (needs deep equal)
	expected := args.Map{"isEqual": false, "isConclusive": false}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- diff pointers same type", actual)
}

// ── PositiveIntegerType with signed ──

func Test_Cov3_PositiveIntegerType_Int64(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerType(int64(42))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType returns correct value -- signed int64", actual)
}

// ── NumberType byte ──

func Test_Cov3_NumberType_Byte(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(byte(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType returns correct value -- byte", actual)
}

// ── PrimitiveType byte ──

func Test_Cov3_PrimitiveType_Byte(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveType(byte(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveType returns correct value -- byte", actual)
}

// ── NullBoth edge ──

func Test_Cov3_NullBoth_BothNonNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(42, "hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- both non-nil", actual)
}

// ── DefinedBoth edge ──

func Test_Cov3_DefinedBoth_LeftNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth(nil, 42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- left nil", actual)
}

// ── NullLeftRight ──

func Test_Cov3_NullLeftRight_BothNull(t *testing.T) {
	l, r := isany.NullLeftRight(nil, nil)
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": true, "right": true}
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns correct value -- both null", actual)
}

func Test_Cov3_NullLeftRight_BothDefined(t *testing.T) {
	l, r := isany.NullLeftRight(42, "hello")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": false, "right": false}
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns correct value -- both defined", actual)
}

// ── DefinedLeftRight ──

func Test_Cov3_DefinedLeftRight_BothDefined(t *testing.T) {
	l, r := isany.DefinedLeftRight(42, "hello")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": true, "right": true}
	expected.ShouldBeEqual(t, 0, "DefinedLeftRight returns correct value -- both defined", actual)
}

func Test_Cov3_DefinedLeftRight_BothNil(t *testing.T) {
	l, r := isany.DefinedLeftRight(nil, nil)
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": false, "right": false}
	expected.ShouldBeEqual(t, 0, "DefinedLeftRight returns nil -- both nil", actual)
}

// ── Function with method ref ──

func Test_Cov3_Function_NilFunc(t *testing.T) {
	var fn func()
	isFunc, _ := isany.Function(fn)
	actual := args.Map{"isFunc": isFunc}
	expected := args.Map{"isFunc": true}
	expected.ShouldBeEqual(t, 0, "Function returns nil -- nil func still detects kind", actual)
}
