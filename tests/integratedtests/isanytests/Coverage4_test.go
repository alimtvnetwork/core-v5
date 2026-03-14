package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── NumberTypeRv — extended int types ──

func Test_Cov4_NumberTypeRv_Int16(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int16(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv int16", actual)
}

func Test_Cov4_NumberTypeRv_Int32(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int32(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv int32", actual)
}

func Test_Cov4_NumberTypeRv_Int64(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(int64(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv int64", actual)
}

func Test_Cov4_NumberTypeRv_Uint8(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint8(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv uint8", actual)
}

func Test_Cov4_NumberTypeRv_Uint16(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint16(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv uint16", actual)
}

func Test_Cov4_NumberTypeRv_Uint64(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint64(1)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv uint64", actual)
}

func Test_Cov4_NumberTypeRv_Float32(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(float32(1.0)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv float32", actual)
}

func Test_Cov4_NumberTypeRv_Float64(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(float64(1.0)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv float64", actual)
}

func Test_Cov4_NumberTypeRv_Bool(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(true))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv bool false", actual)
}

// ── FloatingPointTypeRv — float64 ──

func Test_Cov4_FloatingPointTypeRv_Float64(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(float64(1.0)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv float64", actual)
}

func Test_Cov4_FloatingPointTypeRv_String(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf("nope"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv string", actual)
}

// ── PrimitiveTypeRv — more kinds ──

func Test_Cov4_PrimitiveTypeRv_Int(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Int)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv int", actual)
}

func Test_Cov4_PrimitiveTypeRv_Bool(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Bool)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv bool", actual)
}

func Test_Cov4_PrimitiveTypeRv_String(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.String)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv string", actual)
}

func Test_Cov4_PrimitiveTypeRv_Float32(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Float32)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv float32", actual)
}

func Test_Cov4_PrimitiveTypeRv_Float64(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Float64)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv float64", actual)
}

func Test_Cov4_PrimitiveTypeRv_Map(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Map)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv map", actual)
}

func Test_Cov4_PrimitiveTypeRv_Slice(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Slice)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv slice", actual)
}

// ── PositiveIntegerTypeRv — Uint ──

func Test_Cov4_PositiveIntegerTypeRv_Uint(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint(42)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv uint", actual)
}

func Test_Cov4_PositiveIntegerTypeRv_String(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf("nope"))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv string", actual)
}

// ── NumberType — int16, int32, int64, uint8/16/32/64 ──

func Test_Cov4_NumberType_Int16(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(int16(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType int16", actual)
}

func Test_Cov4_NumberType_Int32(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(int32(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType int32", actual)
}

func Test_Cov4_NumberType_Int64(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(int64(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType int64", actual)
}

func Test_Cov4_NumberType_Uint8(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(uint8(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType uint8", actual)
}

func Test_Cov4_NumberType_Uint16(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(uint16(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType uint16", actual)
}

func Test_Cov4_NumberType_Uint32(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(uint32(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType uint32", actual)
}

func Test_Cov4_NumberType_Uint64(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(uint64(1))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType uint64", actual)
}

func Test_Cov4_NumberType_Bool(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberType bool", actual)
}

// ── Pointer with nil ──

func Test_Cov4_Pointer_Nil(t *testing.T) {
	actual := args.Map{"result": isany.Pointer(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer nil", actual)
}

// ── FuncOnly with nil ──

func Test_Cov4_FuncOnly_Nil(t *testing.T) {
	actual := args.Map{"result": isany.FuncOnly(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncOnly nil", actual)
}

// ── TypeSame with nil ──

func Test_Cov4_TypeSame_NilBoth(t *testing.T) {
	actual := args.Map{"result": isany.TypeSame(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TypeSame nil nil", actual)
}

func Test_Cov4_TypeSame_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.TypeSame(nil, 42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeSame nil vs int", actual)
}

// ── Conclusive same values ──

func Test_Cov4_Conclusive_SameValues(t *testing.T) {
	isEq, isConcl := isany.Conclusive(42, 42)
	actual := args.Map{"isEqual": isEq, "isConcl": isConcl}
	expected := args.Map{"isEqual": true, "isConcl": true}
	expected.ShouldBeEqual(t, 0, "Conclusive same int values equal", actual)
}

// ── Zero struct ──

func Test_Cov4_Zero_Struct(t *testing.T) {
	type s struct{}
	actual := args.Map{"result": isany.Zero(s{})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero empty struct", actual)
}

// ── DeepEqual structs ──

func Test_Cov4_DeepEqual_Structs(t *testing.T) {
	type s struct{ A int }
	actual := args.Map{
		"same": isany.DeepEqual(s{1}, s{1}),
		"diff": isany.DeepEqual(s{1}, s{2}),
	}
	expected := args.Map{"same": true, "diff": false}
	expected.ShouldBeEqual(t, 0, "DeepEqual structs", actual)
}
