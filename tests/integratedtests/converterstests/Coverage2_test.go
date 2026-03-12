package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── StringTo additional coverage ──

func Test_Cov2_StringTo_Integer(t *testing.T) {
	result, err := converters.StringTo.Integer("42")
	actual := args.Map{"value": result, "hasError": err != nil}
	expected := args.Map{"value": 42, "hasError": false}
	expected.ShouldBeEqual(t, 0, "StringTo.Integer", actual)
}

func Test_Cov2_StringTo_Integer_Invalid(t *testing.T) {
	_, err := converters.StringTo.Integer("abc")
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Integer invalid", actual)
}

func Test_Cov2_StringTo_Float64(t *testing.T) {
	result, err := converters.StringTo.Float64("3.14")
	actual := args.Map{"gt3": result > 3.0, "hasError": err != nil}
	expected := args.Map{"gt3": true, "hasError": false}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64", actual)
}

func Test_Cov2_StringTo_Float64_Invalid(t *testing.T) {
	_, err := converters.StringTo.Float64("abc")
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64 invalid", actual)
}

func Test_Cov2_StringTo_Byte(t *testing.T) {
	result, err := converters.StringTo.Byte("42")
	actual := args.Map{"value": int(result), "hasError": err != nil}
	expected := args.Map{"value": 42, "hasError": false}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte", actual)
}

func Test_Cov2_StringTo_Byte_Invalid(t *testing.T) {
	_, err := converters.StringTo.Byte("abc")
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte invalid", actual)
}

func Test_Cov2_StringTo_Byte_Overflow(t *testing.T) {
	_, err := converters.StringTo.Byte("300")
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte overflow", actual)
}

// ── AnyTo additional coverage ──

func Test_Cov2_AnyTo_SmartStrings(t *testing.T) {
	result := converters.AnyTo.SmartStrings("a", 42, true)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyTo.SmartStrings", actual)
}

func Test_Cov2_AnyTo_SmartStrings_Empty(t *testing.T) {
	result := converters.AnyTo.SmartStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo.SmartStrings empty", actual)
}

func Test_Cov2_AnyTo_ToStringsUsingProcessor_WithBreak(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingProcessor(true, func(index int, in any) (string, bool, bool) {
		return "x", true, index >= 0 // break on first
	}, []any{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor with break", actual)
}

func Test_Cov2_AnyTo_ToStringsUsingProcessor_Skip(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingProcessor(true, func(index int, in any) (string, bool, bool) {
		return "", false, false // skip all
	}, []any{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor skip all", actual)
}

func Test_Cov2_AnyTo_ToNonNullItems_WithValues(t *testing.T) {
	result := converters.AnyTo.ToNonNullItems(true, []any{nil, 42, nil, "hello"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToNonNullItems with values", actual)
}

func Test_Cov2_AnyTo_ToStringsUsingSimpleProcessor_WithValues(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingSimpleProcessor(true, func(index int, in any) string {
		return "item"
	}, []any{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingSimpleProcessor with values", actual)
}

// ── StringsTo additional ──

func Test_Cov2_StringsTo_Csv_WithTrim(t *testing.T) {
	result := converters.StringsTo.Csv(true, "  a  ", "  b  ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsTo.Csv with trim", actual)
}

func Test_Cov2_StringsTo_HashmapOptions_NoTrim(t *testing.T) {
	result := converters.StringsTo.HashmapOptions(false, "=", "a=1")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapOptions no trim", actual)
}

func Test_Cov2_StringsTo_HashmapUsingFuncOptions_NoTrim(t *testing.T) {
	result := converters.StringsTo.HashmapUsingFuncOptions(false, func(line string) (string, string) {
		return "k", "v"
	}, "line1")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapUsingFuncOptions no trim", actual)
}

// ── StringsToMapConverter additional ──

func Test_Cov2_StringsToMapConverter_Length(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"a", "b", "c"})
	actual := args.Map{
		"length":    mc.Length(),
		"isEmpty":   mc.IsEmpty(),
		"hasAny":    mc.HasAnyItem(),
		"lastIndex": mc.LastIndex(),
	}
	expected := args.Map{
		"length":    3,
		"isEmpty":   false,
		"hasAny":    true,
		"lastIndex": 2,
	}
	expected.ShouldBeEqual(t, 0, "StringsToMapConverter collection methods", actual)
}

func Test_Cov2_StringsToMapConverter_Empty(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{})
	actual := args.Map{
		"length":  mc.Length(),
		"isEmpty": mc.IsEmpty(),
	}
	expected := args.Map{
		"length":  0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsToMapConverter empty", actual)
}

// ── BytesTo ──

func Test_Cov2_BytesTo_String(t *testing.T) {
	result := converters.BytesTo.String([]byte("hello"))
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesTo.String", actual)
}

func Test_Cov2_BytesTo_String_Empty(t *testing.T) {
	result := converters.BytesTo.String(nil)
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "BytesTo.String nil", actual)
}

func Test_Cov2_UnsafeBytesTo_String(t *testing.T) {
	result := converters.UnsafeBytesTo.String([]byte("hello"))
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesTo.String", actual)
}

func Test_Cov2_UnsafeBytesTo_String_Empty(t *testing.T) {
	result := converters.UnsafeBytesTo.String(nil)
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesTo.String nil", actual)
}

// ── MapStringAnyUsingFunc trim ──

func Test_Cov2_StringsTo_MapStringAnyUsingFunc_Trim(t *testing.T) {
	result := converters.StringsTo.MapStringAnyUsingFunc(true, func(line string) (string, any) {
		return "  k  ", "v"
	}, "line1")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringAnyUsingFunc trim", actual)
}

func Test_Cov2_StringsTo_MapStringIntegerUsingFunc_NoTrim(t *testing.T) {
	result := converters.StringsTo.MapStringIntegerUsingFunc(false, func(line string) (string, int) {
		return "k", 1
	}, "line1")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringIntegerUsingFunc no trim", actual)
}
