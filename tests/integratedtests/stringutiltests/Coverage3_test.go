package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ── AnyToStringNameField / AnyToTypeString ──

func Test_Cov3_AnyToStringNameField(t *testing.T) {
	actual := args.Map{
		"string":  stringutil.AnyToStringNameField("hello") != "",
		"int":     stringutil.AnyToStringNameField(42) != "",
		"nil":     stringutil.AnyToStringNameField(nil),
	}
	expected := args.Map{"string": true, "int": true, "nil": ""}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField", actual)
}

func Test_Cov3_AnyToTypeString(t *testing.T) {
	actual := args.Map{
		"string": stringutil.AnyToTypeString("hello") != "",
		"nil":    stringutil.AnyToTypeString(nil),
	}
	expected := args.Map{"string": true, "nil": ""}
	expected.ShouldBeEqual(t, 0, "AnyToTypeString", actual)
}

// ── IsBlankPtr / IsDefinedPtr / IsEmptyPtr ──

func Test_Cov3_IsBlankPtr(t *testing.T) {
	empty := ""
	space := "   "
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsBlankPtr(nil),
		"empty": stringutil.IsBlankPtr(&empty),
		"space": stringutil.IsBlankPtr(&space),
		"text":  stringutil.IsBlankPtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "space": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr", actual)
}

func Test_Cov3_IsDefinedPtr(t *testing.T) {
	empty := ""
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsDefinedPtr(nil),
		"empty": stringutil.IsDefinedPtr(&empty),
		"text":  stringutil.IsDefinedPtr(&text),
	}
	expected := args.Map{"nil": false, "empty": false, "text": true}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr", actual)
}

func Test_Cov3_IsEmptyOrWhitespacePtr(t *testing.T) {
	empty := ""
	space := "  "
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsEmptyOrWhitespacePtr(nil),
		"empty": stringutil.IsEmptyOrWhitespacePtr(&empty),
		"space": stringutil.IsEmptyOrWhitespacePtr(&space),
		"text":  stringutil.IsEmptyOrWhitespacePtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "space": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespacePtr", actual)
}

func Test_Cov3_IsContainsPtr(t *testing.T) {
	lines := []string{"hello", "world"}
	find1 := "world"
	find2 := "foo"
	find3 := "x"
	actual := args.Map{
		"found":    stringutil.IsContainsPtr(&lines, &find1, 0, true),
		"notFound": stringutil.IsContainsPtr(&lines, &find2, 0, true),
		"nil":      stringutil.IsContainsPtr(nil, &find3, 0, true),
	}
	expected := args.Map{"found": true, "notFound": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr", actual)
}

func Test_Cov3_IsContainsPtrSimple(t *testing.T) {
	lines := []string{"hello", "world"}
	actual := args.Map{
		"found":    stringutil.IsContainsPtrSimple(&lines, "hello", 0, true),
		"notFound": stringutil.IsContainsPtrSimple(&lines, "foo", 0, true),
		"nil":      stringutil.IsContainsPtrSimple(nil, "x", 0, true),
	}
	expected := args.Map{"found": true, "notFound": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple", actual)
}

// ── IsStarts/IsEnds ──

func Test_Cov3_IsStarts(t *testing.T) {
	actual := args.Map{
		"starts":  stringutil.IsStarts("hello world", "hello"),
		"notStarts": stringutil.IsStarts("hello world", "world"),
	}
	expected := args.Map{"starts": true, "notStarts": false}
	expected.ShouldBeEqual(t, 0, "IsStarts", actual)
}

func Test_Cov3_IsEnds(t *testing.T) {
	actual := args.Map{
		"ends":    stringutil.IsEnds("hello world", "world"),
		"notEnds": stringutil.IsEnds("hello world", "hello"),
	}
	expected := args.Map{"ends": true, "notEnds": false}
	expected.ShouldBeEqual(t, 0, "IsEnds", actual)
}

func Test_Cov3_IsStartsWith(t *testing.T) {
	actual := args.Map{
		"starts":    stringutil.IsStartsWith("hello", "hel"),
		"notStarts": stringutil.IsStartsWith("hello", "wor"),
	}
	expected := args.Map{"starts": true, "notStarts": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith", actual)
}

func Test_Cov3_IsEndsWith(t *testing.T) {
	actual := args.Map{
		"ends":    stringutil.IsEndsWith("hello", "llo"),
		"notEnds": stringutil.IsEndsWith("hello", "hel"),
	}
	expected := args.Map{"ends": true, "notEnds": false}
	expected.ShouldBeEqual(t, 0, "IsEndsWith", actual)
}

func Test_Cov3_IsStartsChar(t *testing.T) {
	actual := args.Map{
		"starts":    stringutil.IsStartsChar("hello", 'h'),
		"notStarts": stringutil.IsStartsChar("hello", 'z'),
		"empty":     stringutil.IsStartsChar("", 'h'),
	}
	expected := args.Map{"starts": true, "notStarts": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsStartsChar", actual)
}

func Test_Cov3_IsEndsChar(t *testing.T) {
	actual := args.Map{
		"ends":    stringutil.IsEndsChar("hello", 'o'),
		"notEnds": stringutil.IsEndsChar("hello", 'z'),
		"empty":   stringutil.IsEndsChar("", 'o'),
	}
	expected := args.Map{"ends": true, "notEnds": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsEndsChar", actual)
}

func Test_Cov3_IsStartsRune(t *testing.T) {
	actual := args.Map{
		"starts": stringutil.IsStartsRune("hello", 'h'),
		"empty":  stringutil.IsStartsRune("", 'h'),
	}
	expected := args.Map{"starts": true, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsStartsRune", actual)
}

func Test_Cov3_IsEndsRune(t *testing.T) {
	actual := args.Map{
		"ends":  stringutil.IsEndsRune("hello", 'o'),
		"empty": stringutil.IsEndsRune("", 'o'),
	}
	expected := args.Map{"ends": true, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsEndsRune", actual)
}

func Test_Cov3_IsStartsAndEndsChar(t *testing.T) {
	actual := args.Map{
		"match": stringutil.IsStartsAndEndsChar("hello", 'h', 'o'),
		"noMatch": stringutil.IsStartsAndEndsChar("hello", 'h', 'x'),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar", actual)
}

func Test_Cov3_IsStartsAndEndsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStartsAndEndsWith("hello world", "hello", "world"),
		"noMatch": stringutil.IsStartsAndEndsWith("hello world", "hello", "xyz"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsWith", actual)
}

func Test_Cov3_IsAnyStartsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsAnyStartsWith("hello", []string{"he", "wo"}),
		"noMatch": stringutil.IsAnyStartsWith("hello", []string{"wo", "xy"}),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith", actual)
}

func Test_Cov3_IsAnyEndsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsAnyEndsWith("hello", []string{"lo", "xy"}),
		"noMatch": stringutil.IsAnyEndsWith("hello", []string{"ab", "cd"}),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith", actual)
}

// ── FirstChar / ClonePtr / SafeClonePtr ──

func Test_Cov3_FirstChar(t *testing.T) {
	actual := args.Map{
		"first": stringutil.FirstChar("hello"),
		"empty": stringutil.FirstChar(""),
	}
	expected := args.Map{"first": "h", "empty": ""}
	expected.ShouldBeEqual(t, 0, "FirstChar", actual)
}

func Test_Cov3_ClonePtr(t *testing.T) {
	text := "hello"
	result := stringutil.ClonePtr(&text)
	nilResult := stringutil.ClonePtr(nil)
	actual := args.Map{"notNil": result != nil, "val": *result, "nilIsNil": nilResult == nil}
	expected := args.Map{"notNil": true, "val": "hello", "nilIsNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr", actual)
}

func Test_Cov3_SafeClonePtr(t *testing.T) {
	text := "hello"
	result := stringutil.SafeClonePtr(&text)
	nilResult := stringutil.SafeClonePtr(nil)
	actual := args.Map{"val": *result, "nilNotNil": nilResult != nil, "nilVal": *nilResult}
	expected := args.Map{"val": "hello", "nilNotNil": true, "nilVal": ""}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr", actual)
}

// ── SafeSubstring variants ──

func Test_Cov3_SafeSubstring(t *testing.T) {
	actual := args.Map{
		"normal":  stringutil.SafeSubstring("hello", 1, 3),
		"outOfRange": stringutil.SafeSubstring("hi", 0, 10),
	}
	expected := args.Map{"normal": "el", "outOfRange": "hi"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring", actual)
}

func Test_Cov3_SafeSubstringStarts(t *testing.T) {
	actual := args.Map{
		"normal": stringutil.SafeSubstringStarts("hello", 2),
		"outOfRange": stringutil.SafeSubstringStarts("hi", 10),
	}
	expected := args.Map{"normal": "llo", "outOfRange": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts", actual)
}

func Test_Cov3_SafeSubstringEnds(t *testing.T) {
	actual := args.Map{
		"normal": stringutil.SafeSubstringEnds("hello", 3),
		"outOfRange": stringutil.SafeSubstringEnds("hi", 10),
	}
	expected := args.Map{"normal": "hel", "outOfRange": "hi"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds", actual)
}

// ── MaskLine / MaskLines / MaskTrimLine / MaskTrimLines ──

func Test_Cov3_MaskLine(t *testing.T) {
	actual := args.Map{
		"result": stringutil.MaskLine("XXXXXXXXXX", "abc"),
	}
	expected := args.Map{"result": "abcXXXXXXX"}
	expected.ShouldBeEqual(t, 0, "MaskLine", actual)
}

func Test_Cov3_MaskLines(t *testing.T) {
	result := stringutil.MaskLines("XXXXX", []string{"ab", "cde"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MaskLines", actual)
}

func Test_Cov3_MaskTrimLine(t *testing.T) {
	actual := args.Map{"result": stringutil.MaskTrimLine("XXXXXXXXXX", "  abc  ")}
	expected := args.Map{"result": "abcXXXXXXX"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine", actual)
}

func Test_Cov3_MaskTrimLines(t *testing.T) {
	result := stringutil.MaskTrimLines("XXXXX", []string{"  ab  ", "cde"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines", actual)
}

// ── RemoveMany / RemoveManyBySplitting ──

func Test_Cov3_RemoveMany(t *testing.T) {
	result := stringutil.RemoveMany("hello world foo", "world", "foo")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello  "}
	expected.ShouldBeEqual(t, 0, "RemoveMany", actual)
}

func Test_Cov3_RemoveManyBySplitting(t *testing.T) {
	result := stringutil.RemoveManyBySplitting("hello world foo", "world", "foo")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RemoveManyBySplitting", actual)
}

// ── TrimKeepSingleSpaceOnly ──

func Test_Cov3_TrimKeepSingleSpaceOnly(t *testing.T) {
	result := stringutil.TrimKeepSingleSpaceOnly("  hello   world  ")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello world"}
	expected.ShouldBeEqual(t, 0, "TrimKeepSingleSpaceOnly", actual)
}

// ── ToBool / ToByte / ToInt variants ──

func Test_Cov3_ToBool(t *testing.T) {
	actual := args.Map{
		"true":    stringutil.ToBool("true"),
		"false":   stringutil.ToBool("false"),
		"invalid": stringutil.ToBool("abc"),
	}
	expected := args.Map{"true": true, "false": false, "invalid": false}
	expected.ShouldBeEqual(t, 0, "ToBool", actual)
}

func Test_Cov3_ToByte(t *testing.T) {
	val, err := stringutil.ToByte("42")
	_, errInvalid := stringutil.ToByte("abc")
	actual := args.Map{"val": val, "noErr": err == nil, "invalidErr": errInvalid != nil}
	expected := args.Map{"val": byte(42), "noErr": true, "invalidErr": true}
	expected.ShouldBeEqual(t, 0, "ToByte", actual)
}

func Test_Cov3_ToByteDefault(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToByteDefault("42", 0),
		"invalid": stringutil.ToByteDefault("abc", 99),
	}
	expected := args.Map{"valid": byte(42), "invalid": byte(99)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault", actual)
}

func Test_Cov3_ToInt(t *testing.T) {
	val, err := stringutil.ToInt("42")
	_, errInvalid := stringutil.ToInt("abc")
	actual := args.Map{"val": val, "noErr": err == nil, "invalidErr": errInvalid != nil}
	expected := args.Map{"val": 42, "noErr": true, "invalidErr": true}
	expected.ShouldBeEqual(t, 0, "ToInt", actual)
}

func Test_Cov3_ToIntDef(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToIntDef("42", 0),
		"invalid": stringutil.ToIntDef("abc", 99),
	}
	expected := args.Map{"valid": 42, "invalid": 99}
	expected.ShouldBeEqual(t, 0, "ToIntDef", actual)
}

func Test_Cov3_ToIntDefault(t *testing.T) {
	val, isValid := stringutil.ToIntDefault("42", 0)
	_, isValidFail := stringutil.ToIntDefault("abc", 99)
	actual := args.Map{"val": val, "isValid": isValid, "isValidFail": isValidFail}
	expected := args.Map{"val": 42, "isValid": true, "isValidFail": false}
	expected.ShouldBeEqual(t, 0, "ToIntDefault", actual)
}

func Test_Cov3_ToInt8(t *testing.T) {
	val, err := stringutil.ToInt8("42")
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": int8(42), "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToInt8", actual)
}

func Test_Cov3_ToInt8Def(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt8Def("42", 0),
		"invalid": stringutil.ToInt8Def("abc", 99),
	}
	expected := args.Map{"valid": int8(42), "invalid": int8(99)}
	expected.ShouldBeEqual(t, 0, "ToInt8Def", actual)
}

func Test_Cov3_ToInt16(t *testing.T) {
	val, err := stringutil.ToInt16("42")
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": int16(42), "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToInt16", actual)
}

func Test_Cov3_ToInt16Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt16Default("42", 0),
		"invalid": stringutil.ToInt16Default("abc", 99),
	}
	expected := args.Map{"valid": int16(42), "invalid": int16(99)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default", actual)
}

func Test_Cov3_ToInt32(t *testing.T) {
	val, err := stringutil.ToInt32("42")
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": int32(42), "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToInt32", actual)
}

func Test_Cov3_ToInt32Def(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt32Def("42", 0),
		"invalid": stringutil.ToInt32Def("abc", 99),
	}
	expected := args.Map{"valid": int32(42), "invalid": int32(99)}
	expected.ShouldBeEqual(t, 0, "ToInt32Def", actual)
}

func Test_Cov3_ToUint16Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToUint16Default("42", 0),
		"invalid": stringutil.ToUint16Default("abc", 99),
	}
	expected := args.Map{"valid": uint16(42), "invalid": uint16(99)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default", actual)
}

func Test_Cov3_ToUint32Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToUint32Default("42", 0),
		"invalid": stringutil.ToUint32Default("abc", 99),
	}
	expected := args.Map{"valid": uint32(42), "invalid": uint32(99)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default", actual)
}

// ── SplitLeftRight / SplitFirstLast ──

func Test_Cov3_SplitLeftRight(t *testing.T) {
	l, r := stringutil.SplitLeftRight("=", "key=value")
	actual := args.Map{"l": l, "r": r}
	expected := args.Map{"l": "key", "r": "value"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight", actual)
}

func Test_Cov3_SplitLeftRightTrimmed(t *testing.T) {
	l, r := stringutil.SplitLeftRightTrimmed("=", "  key  =  value  ")
	actual := args.Map{"l": l, "r": r}
	expected := args.Map{"l": "key", "r": "value"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed", actual)
}

func Test_Cov3_SplitFirstLast(t *testing.T) {
	first, last := stringutil.SplitFirstLast(".", "a.b.c")
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast", actual)
}

// ── KeyValReplacer ──

func Test_Cov3_KeyValReplacer(t *testing.T) {
	result := stringutil.KeyValReplacer(
		"Hello {name} at {place}",
		true,
		map[string]string{"name": "Alice", "place": "home"},
	)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "Hello Alice at home"}
	expected.ShouldBeEqual(t, 0, "KeyValReplacer", actual)
}
