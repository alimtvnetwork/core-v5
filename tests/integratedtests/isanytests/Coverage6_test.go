package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── ReflectNull with reflect.Value input (line 15-16 branch) ──

func Test_Cov6_ReflectNull_ReflectValueInput_Invalid(t *testing.T) {
	rv := reflect.Value{} // invalid reflect.Value
	actual := args.Map{"result": isany.ReflectNull(rv)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull reflect.Value invalid", actual)
}

func Test_Cov6_ReflectNull_ReflectValueInput_Valid(t *testing.T) {
	rv := reflect.ValueOf(42)
	actual := args.Map{"result": isany.ReflectNull(rv)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull reflect.Value valid int", actual)
}

func Test_Cov6_ReflectNull_ReflectValueInput_NilPtr(t *testing.T) {
	var ptr *int
	rv := reflect.ValueOf(ptr)
	actual := args.Map{"result": isany.ReflectNull(rv)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull reflect.Value nil ptr", actual)
}

// ── JsonEqual error paths ──

func Test_Cov6_JsonEqual_BothUnmarshalable(t *testing.T) {
	// channels cannot be marshaled
	ch1 := make(chan int)
	ch2 := make(chan int)
	actual := args.Map{"result": isany.JsonEqual(ch1, ch2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual both unmarshalable", actual)
}

func Test_Cov6_JsonEqual_OneUnmarshalable(t *testing.T) {
	ch := make(chan int)
	actual := args.Map{"result": isany.JsonEqual(ch, 42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual one unmarshalable", actual)
}

func Test_Cov6_JsonEqual_OtherUnmarshalable(t *testing.T) {
	ch := make(chan int)
	actual := args.Map{"result": isany.JsonEqual(42, ch)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual other unmarshalable", actual)
}

// ── JsonMismatch with channels ──

func Test_Cov6_JsonMismatch_BothUnmarshalable(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	actual := args.Map{"result": isany.JsonMismatch(ch1, ch2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonMismatch both unmarshalable", actual)
}

// ── Null with non-nil chan ──

func Test_Cov6_Null_NonNilChan(t *testing.T) {
	ch := make(chan int)
	actual := args.Map{"result": isany.Null(ch)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null non-nil chan", actual)
}

// ── ReflectValueNull invalid ──

func Test_Cov6_ReflectValueNull_Invalid(t *testing.T) {
	rv := reflect.Value{}
	actual := args.Map{"result": isany.ReflectValueNull(rv)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull invalid", actual)
}

// ── FuncOnly with func ──

func Test_Cov6_FuncOnly_ValidFunc(t *testing.T) {
	actual := args.Map{"result": isany.FuncOnly(func() {})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FuncOnly valid func", actual)
}

// ── Conclusive same nil values but non-interface ──

func Test_Cov6_Conclusive_BothTypedNilDiffType(t *testing.T) {
	var a *int
	var b *string
	isEq, isConcl := isany.Conclusive(a, b)
	actual := args.Map{"isEqual": isEq, "isConclusive": isConcl}
	expected := args.Map{"isEqual": false, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive both typed nil diff type", actual)
}

// ── ReflectNull with nil directly ──

func Test_Cov6_ReflectNull_NilDirect(t *testing.T) {
	actual := args.Map{"result": isany.ReflectNull(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull nil direct", actual)
}

// ── ReflectNotNull with value ──

func Test_Cov6_ReflectNotNull_Value(t *testing.T) {
	actual := args.Map{"result": isany.ReflectNotNull(42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNotNull value", actual)
}

// ── PositiveIntegerType uint ──

func Test_Cov6_PositiveIntegerType_Uint(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerType(uint(42))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType uint", actual)
}

// ── FloatingPointType float64 ──

func Test_Cov6_FloatingPointType_Float64(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointType(float64(3.14))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointType float64", actual)
}

func Test_Cov6_FloatingPointType_Int(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointType(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointType int", actual)
}
