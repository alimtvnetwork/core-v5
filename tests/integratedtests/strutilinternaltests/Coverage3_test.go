package strutilinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

// ── AnyToFieldNameString ──

func Test_Cov3_AnyToFieldNameString_Nil(t *testing.T) {
	actual := args.Map{"result": strutilinternal.AnyToFieldNameString(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyToFieldNameString nil", actual)
}

func Test_Cov3_AnyToFieldNameString_Value(t *testing.T) {
	result := strutilinternal.AnyToFieldNameString("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToFieldNameString value", actual)
}

// ── AnyToString ──

func Test_Cov3_AnyToString_Nil(t *testing.T) {
	actual := args.Map{"result": strutilinternal.AnyToString(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyToString nil", actual)
}

func Test_Cov3_AnyToString_Value(t *testing.T) {
	result := strutilinternal.AnyToString("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToString value", actual)
}

func Test_Cov3_AnyToString_Ptr(t *testing.T) {
	v := "hello"
	result := strutilinternal.AnyToString(&v)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToString ptr", actual)
}

// ── AnyToStringUsing ──

func Test_Cov3_AnyToStringUsing_Nil(t *testing.T) {
	actual := args.Map{"result": strutilinternal.AnyToStringUsing(true, nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyToStringUsing nil", actual)
}

func Test_Cov3_AnyToStringUsing_IncludeFields(t *testing.T) {
	result := strutilinternal.AnyToStringUsing(true, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToStringUsing include fields", actual)
}

func Test_Cov3_AnyToStringUsing_NoFields(t *testing.T) {
	result := strutilinternal.AnyToStringUsing(false, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToStringUsing no fields", actual)
}

// ── MaskLine ──

func Test_Cov3_MaskLine_Empty(t *testing.T) {
	actual := args.Map{"result": strutilinternal.MaskLine("****", "")}
	expected := args.Map{"result": "****"}
	expected.ShouldBeEqual(t, 0, "MaskLine empty line", actual)
}

func Test_Cov3_MaskLine_LongerThanMask(t *testing.T) {
	actual := args.Map{"result": strutilinternal.MaskLine("**", "hello")}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskLine longer than mask", actual)
}

func Test_Cov3_MaskLine_EmptyMask(t *testing.T) {
	actual := args.Map{"result": strutilinternal.MaskLine("", "hello")}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskLine empty mask", actual)
}

func Test_Cov3_MaskLine_Partial(t *testing.T) {
	result := strutilinternal.MaskLine("********", "hi")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 8}
	expected.ShouldBeEqual(t, 0, "MaskLine partial", actual)
}

// ── MaskTrimLine ──

func Test_Cov3_MaskTrimLine_Whitespace(t *testing.T) {
	actual := args.Map{"result": strutilinternal.MaskTrimLine("****", "   ")}
	expected := args.Map{"result": "****"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine whitespace", actual)
}

func Test_Cov3_MaskTrimLine_Longer(t *testing.T) {
	actual := args.Map{"result": strutilinternal.MaskTrimLine("**", " hello ")}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine longer than mask", actual)
}

func Test_Cov3_MaskTrimLine_Partial(t *testing.T) {
	result := strutilinternal.MaskTrimLine("********", " hi ")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 8}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine partial", actual)
}

// ── SplitLeftRight / SplitLeftRightTrim ──

func Test_Cov3_SplitLeftRight_TwoParts(t *testing.T) {
	l, r := strutilinternal.SplitLeftRight("=", "key=value")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": "key", "right": "value"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight two parts", actual)
}

func Test_Cov3_SplitLeftRight_NoParts(t *testing.T) {
	l, r := strutilinternal.SplitLeftRight("=", "noequals")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": "noequals", "right": ""}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight no parts", actual)
}

func Test_Cov3_SplitLeftRightTrim(t *testing.T) {
	l, r := strutilinternal.SplitLeftRightTrim("=", " key = value ")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": "key", "right": "value"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrim", actual)
}

// ── CurlyWrapIf ──

func Test_Cov3_CurlyWrapIf_True(t *testing.T) {
	result := strutilinternal.CurlyWrapIf(true, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf true", actual)
}

func Test_Cov3_CurlyWrapIf_False(t *testing.T) {
	result := strutilinternal.CurlyWrapIf(false, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf false", actual)
}

// ── Clone ──

func Test_Cov3_Clone_Empty(t *testing.T) {
	result := strutilinternal.Clone([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone empty", actual)
}

func Test_Cov3_Clone_Items(t *testing.T) {
	result := strutilinternal.Clone([]string{"a", "b"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "Clone items", actual)
}

// ── ReflectInterfaceVal ──

func Test_Cov3_ReflectInterfaceVal_NonPtr(t *testing.T) {
	result := strutilinternal.ReflectInterfaceVal("hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal non-ptr", actual)
}

func Test_Cov3_ReflectInterfaceVal_Ptr(t *testing.T) {
	v := "hello"
	result := strutilinternal.ReflectInterfaceVal(&v)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal ptr", actual)
}

// ── NonEmpty / NonEmptyJoin ──

func Test_Cov3_NonEmpty(t *testing.T) {
	result := strutilinternal.NonEmpty([]string{"a", "", "b", ""})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmpty", actual)
}

func Test_Cov3_NonEmptyJoin(t *testing.T) {
	result := strutilinternal.NonEmptyJoin([]string{"a", "", "b"}, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin", actual)
}

// ── NonEmptySlicePtr ──

func Test_Cov3_NonEmptySlicePtr(t *testing.T) {
	result := strutilinternal.NonEmptySlicePtr([]string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr", actual)
}

// ── NonWhitespaceSlice ──

func Test_Cov3_NonWhitespaceSlice(t *testing.T) {
	result := strutilinternal.NonWhitespaceSlice([]string{"a", " ", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceSlice", actual)
}

// ── NonWhitespaceTrimSlice ──

func Test_Cov3_NonWhitespaceTrimSlice(t *testing.T) {
	result := strutilinternal.NonWhitespaceTrimSlice([]string{" a ", " "})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "a"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceTrimSlice", actual)
}

// ── NonWhitespaceJoin ──

func Test_Cov3_NonWhitespaceJoin(t *testing.T) {
	result := strutilinternal.NonWhitespaceJoin([]string{"a", " ", "b"}, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin", actual)
}

// ── IsEmptyOrWhitespace / IsNullOrEmpty / IsNullOrEmptyOrWhitespace ──

func Test_Cov3_IsEmptyOrWhitespace(t *testing.T) {
	actual := args.Map{
		"empty": strutilinternal.IsEmptyOrWhitespace(""),
		"ws":    strutilinternal.IsEmptyOrWhitespace("  "),
		"val":   strutilinternal.IsEmptyOrWhitespace("a"),
	}
	expected := args.Map{"empty": true, "ws": true, "val": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace", actual)
}

func Test_Cov3_IsNullOrEmpty(t *testing.T) {
	actual := args.Map{
		"nil":   strutilinternal.IsNullOrEmpty(nil),
		"empty": strutilinternal.IsNullOrEmpty(ptrStr("")),
		"val":   strutilinternal.IsNullOrEmpty(ptrStr("a")),
	}
	expected := args.Map{"nil": true, "empty": true, "val": false}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmpty", actual)
}

func Test_Cov3_IsNullOrEmptyOrWhitespace(t *testing.T) {
	actual := args.Map{
		"nil": strutilinternal.IsNullOrEmptyOrWhitespace(nil),
		"ws":  strutilinternal.IsNullOrEmptyOrWhitespace(ptrStr("  ")),
		"val": strutilinternal.IsNullOrEmptyOrWhitespace(ptrStr("a")),
	}
	expected := args.Map{"nil": true, "ws": true, "val": false}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyOrWhitespace", actual)
}

func ptrStr(s string) *string { return &s }

// ── ReplaceTemplateMap curly ──

func Test_Cov3_ReplaceTemplateMap_Curly(t *testing.T) {
	result := strutilinternal.ReplaceTemplateMap(
		true,
		"Hello {name}, you are {age}",
		map[string]string{"name": "Alice", "age": "30"},
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "Hello Alice, you are 30"}
	expected.ShouldBeEqual(t, 0, "ReplaceTemplateMap curly", actual)
}
