package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── AllZero ──

func Test_Cov7_AllZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AllZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero empty -- true", actual)
}

func Test_Cov7_AllZero_AllZeros(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, "", false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero all zeros -- true", actual)
}

func Test_Cov7_AllZero_OneNonZero(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, 1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllZero one non-zero -- false", actual)
}

// ── AnyZero ──

func Test_Cov7_AnyZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero empty -- true", actual)
}

func Test_Cov7_AnyZero_OneZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero(1, 0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero one zero -- true", actual)
}

func Test_Cov7_AnyZero_NoneZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero(1, "a", true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero none zero -- false", actual)
}

// ── Conclusive ──

func Test_Cov7_Conclusive_BothNil(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(nil, nil)
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": true, "conclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive both nil", actual)
}

func Test_Cov7_Conclusive_LeftNil(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(nil, "x")
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": false, "conclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive left nil", actual)
}

func Test_Cov7_Conclusive_RightNil(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive("x", nil)
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": false, "conclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive right nil", actual)
}

func Test_Cov7_Conclusive_SameRef(t *testing.T) {
	s := "hello"
	isEqual, isConclusive := isany.Conclusive(s, s)
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": true, "conclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive same ref", actual)
}

func Test_Cov7_Conclusive_DiffType(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(1, "1")
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": false, "conclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive diff type", actual)
}

func Test_Cov7_Conclusive_BothNilPtr(t *testing.T) {
	var p1, p2 *int
	isEqual, isConclusive := isany.Conclusive(p1, p2)
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": true, "conclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive both nil ptr", actual)
}

func Test_Cov7_Conclusive_OneNilPtr(t *testing.T) {
	var p1 *int
	v := 5
	isEqual, isConclusive := isany.Conclusive(p1, &v)
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": false, "conclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive one nil ptr", actual)
}

func Test_Cov7_Conclusive_SameTypeDiffValue(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(1, 2)
	actual := args.Map{"equal": isEqual, "conclusive": isConclusive}
	expected := args.Map{"equal": false, "conclusive": false}
	expected.ShouldBeEqual(t, 0, "Conclusive same type diff value -- inconclusive", actual)
}

// ── Defined ──

func Test_Cov7_Defined_Nil(t *testing.T) {
	actual := args.Map{"result": isany.Defined(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Defined nil -- false", actual)
}

func Test_Cov7_Defined_Value(t *testing.T) {
	actual := args.Map{"result": isany.Defined(42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Defined value -- true", actual)
}

// ── DefinedItems ──

func Test_Cov7_DefinedItems_Empty(t *testing.T) {
	allDefined, items := isany.DefinedItems()
	actual := args.Map{"allDefined": allDefined, "len": len(items)}
	expected := args.Map{"allDefined": false, "len": 0}
	expected.ShouldBeEqual(t, 0, "DefinedItems empty", actual)
}

func Test_Cov7_DefinedItems_AllDefined(t *testing.T) {
	allDefined, items := isany.DefinedItems("a", 1, true)
	actual := args.Map{"allDefined": allDefined, "len": len(items)}
	expected := args.Map{"allDefined": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "DefinedItems all defined", actual)
}

func Test_Cov7_DefinedItems_SomeNil(t *testing.T) {
	allDefined, items := isany.DefinedItems("a", nil, "b")
	actual := args.Map{"allDefined": allDefined, "len": len(items)}
	expected := args.Map{"allDefined": false, "len": 2}
	expected.ShouldBeEqual(t, 0, "DefinedItems some nil", actual)
}

// ── DefinedLeftRight ──

func Test_Cov7_DefinedLeftRight_BothDefined(t *testing.T) {
	l, r := isany.DefinedLeftRight("a", "b")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": true, "right": true}
	expected.ShouldBeEqual(t, 0, "DefinedLeftRight both defined", actual)
}

func Test_Cov7_DefinedLeftRight_LeftNil(t *testing.T) {
	l, r := isany.DefinedLeftRight(nil, "b")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": false, "right": true}
	expected.ShouldBeEqual(t, 0, "DefinedLeftRight left nil", actual)
}

// ── NullLeftRight ──

func Test_Cov7_NullLeftRight_BothDefined(t *testing.T) {
	l, r := isany.NullLeftRight("a", "b")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": false, "right": false}
	expected.ShouldBeEqual(t, 0, "NullLeftRight both defined", actual)
}

// ── FloatingPointType / FloatingPointTypeRv ──

func Test_Cov7_FloatingPointType_Float64(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointType(3.14)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointType float64 -- true", actual)
}

func Test_Cov7_FloatingPointType_Int(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointType(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FloatingPointType int -- false", actual)
}

func Test_Cov7_FloatingPointTypeRv_Float32(t *testing.T) {
	actual := args.Map{"result": isany.FloatingPointTypeRv(reflect.ValueOf(float32(1.0)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FloatingPointTypeRv float32 -- true", actual)
}

// ── NumberType / NumberTypeRv ──

func Test_Cov7_NumberType_Int(t *testing.T) {
	actual := args.Map{"result": isany.NumberType(42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberType int -- true", actual)
}

func Test_Cov7_NumberType_String(t *testing.T) {
	actual := args.Map{"result": isany.NumberType("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NumberType string -- false", actual)
}

func Test_Cov7_NumberTypeRv_Uint(t *testing.T) {
	actual := args.Map{"result": isany.NumberTypeRv(reflect.ValueOf(uint(5)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NumberTypeRv uint -- true", actual)
}

// ── PositiveIntegerType ──

func Test_Cov7_PositiveIntegerType_Uint(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerType(uint(5))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType uint -- true", actual)
}

func Test_Cov7_PositiveIntegerType_Int(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerType(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerType int -- false", actual)
}

func Test_Cov7_PositiveIntegerTypeRv_Uint16(t *testing.T) {
	actual := args.Map{"result": isany.PositiveIntegerTypeRv(reflect.ValueOf(uint16(5)))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PositiveIntegerTypeRv uint16 -- true", actual)
}

// ── PrimitiveType / PrimitiveTypeRv ──

func Test_Cov7_PrimitiveType_String(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveType("hello")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveType string -- true", actual)
}

func Test_Cov7_PrimitiveType_Slice(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveType([]int{1})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveType slice -- false", actual)
}

func Test_Cov7_PrimitiveTypeRv_Bool(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Bool)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv bool -- true", actual)
}

func Test_Cov7_PrimitiveTypeRv_Map(t *testing.T) {
	actual := args.Map{"result": isany.PrimitiveTypeRv(reflect.Map)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PrimitiveTypeRv map -- false", actual)
}

// ── Pointer ──

func Test_Cov7_Pointer_Ptr(t *testing.T) {
	v := 5
	actual := args.Map{"result": isany.Pointer(&v)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Pointer ptr -- true", actual)
}

func Test_Cov7_Pointer_NonPtr(t *testing.T) {
	actual := args.Map{"result": isany.Pointer(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer non-ptr -- false", actual)
}

// ── FuncOnly / Function ──

func Test_Cov7_FuncOnly_Func(t *testing.T) {
	actual := args.Map{"result": isany.FuncOnly(func() {})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "FuncOnly func -- true", actual)
}

func Test_Cov7_FuncOnly_Nil(t *testing.T) {
	actual := args.Map{"result": isany.FuncOnly(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncOnly nil -- false", actual)
}

func Test_Cov7_FuncOnly_NonFunc(t *testing.T) {
	actual := args.Map{"result": isany.FuncOnly(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "FuncOnly non-func -- false", actual)
}

func Test_Cov7_Function_NilFunc(t *testing.T) {
	var fn func()
	isFunc, name := isany.Function(fn)
	actual := args.Map{"isFunc": isFunc, "name": name}
	expected := args.Map{"isFunc": true, "name": ""}
	expected.ShouldBeEqual(t, 0, "Function nil func -- isFunc true name empty", actual)
}

func Test_Cov7_Function_ValidFunc(t *testing.T) {
	isFunc, name := isany.Function(isany.Null)
	actual := args.Map{"isFunc": isFunc, "hasName": name != ""}
	expected := args.Map{"isFunc": true, "hasName": true}
	expected.ShouldBeEqual(t, 0, "Function valid -- isFunc true has name", actual)
}

// ── JsonEqual / JsonMismatch ──

func Test_Cov7_JsonEqual_Strings(t *testing.T) {
	actual := args.Map{"result": isany.JsonEqual("abc", "abc")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonEqual strings match", actual)
}

func Test_Cov7_JsonEqual_Ints(t *testing.T) {
	actual := args.Map{"result": isany.JsonEqual(1, 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonEqual ints match", actual)
}

func Test_Cov7_JsonEqual_Structs(t *testing.T) {
	type s struct{ A int }
	actual := args.Map{"result": isany.JsonEqual(s{1}, s{1})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonEqual structs match", actual)
}

func Test_Cov7_JsonEqual_Different(t *testing.T) {
	actual := args.Map{"result": isany.JsonEqual(1, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "JsonEqual different -- false", actual)
}

func Test_Cov7_JsonMismatch(t *testing.T) {
	actual := args.Map{"result": isany.JsonMismatch(1, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "JsonMismatch different -- true", actual)
}

// ── NotNull / ReflectNotNull ──

func Test_Cov7_NotNull_Value(t *testing.T) {
	actual := args.Map{"result": isany.NotNull(42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotNull value -- true", actual)
}

func Test_Cov7_ReflectNotNull_Value(t *testing.T) {
	actual := args.Map{"result": isany.ReflectNotNull(42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNotNull value -- true", actual)
}

// ── StringEqual ──

func Test_Cov7_StringEqual_Same(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual("a", "a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual same -- true", actual)
}

func Test_Cov7_StringEqual_Diff(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringEqual diff -- false", actual)
}

// ── TypeSame ──

func Test_Cov7_TypeSame_Same(t *testing.T) {
	actual := args.Map{"result": isany.TypeSame(1, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "TypeSame same type -- true", actual)
}

func Test_Cov7_TypeSame_Diff(t *testing.T) {
	actual := args.Map{"result": isany.TypeSame(1, "2")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeSame diff type -- false", actual)
}

// ── DeepEqualAllItems ──

func Test_Cov7_DeepEqualAllItems_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems empty -- true", actual)
}

func Test_Cov7_DeepEqualAllItems_Single(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems single -- true", actual)
}

func Test_Cov7_DeepEqualAllItems_TwoEqual(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems two equal -- true", actual)
}

func Test_Cov7_DeepEqualAllItems_TwoDiff(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems two diff -- false", actual)
}

func Test_Cov7_DeepEqualAllItems_ThreeEqual(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems three equal -- true", actual)
}

func Test_Cov7_DeepEqualAllItems_ThreeOneDiff(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems three one diff -- false", actual)
}

// ── Zero ──

func Test_Cov7_Zero_ZeroInt(t *testing.T) {
	actual := args.Map{"result": isany.Zero(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero zero int -- true", actual)
}

func Test_Cov7_Zero_NonZero(t *testing.T) {
	actual := args.Map{"result": isany.Zero(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Zero non-zero -- false", actual)
}
