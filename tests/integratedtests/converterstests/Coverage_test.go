package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── stringsTo coverage ──

func Test_Cov_StringsTo_HashmapTrimColon(t *testing.T) {
	result := converters.StringsTo.HashmapTrimColon("a:1", "b:2")
	actual := args.Map{
		"len":    len(result),
		"hasA":   result["a"] == "1",
	}
	expected := args.Map{
		"len":    2,
		"hasA":   true,
	}
	expected.ShouldBeEqual(t, 0, "HashmapTrimColon", actual)
}

func Test_Cov_StringsTo_HashmapTrimHyphen(t *testing.T) {
	result := converters.StringsTo.HashmapTrimHyphen("a-1", "b-2")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "HashmapTrimHyphen", actual)
}

func Test_Cov_StringsTo_HashmapOptions(t *testing.T) {
	result := converters.StringsTo.HashmapOptions(true, "=", "a = 1", "b = 2")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "HashmapOptions", actual)
}

func Test_Cov_StringsTo_HashmapTrim(t *testing.T) {
	result := converters.StringsTo.HashmapTrim(":", []string{"a:1"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapTrim", actual)
}

func Test_Cov_StringsTo_HashmapUsingFuncOptions(t *testing.T) {
	result := converters.StringsTo.HashmapUsingFuncOptions(true, func(line string) (string, string) {
		return "k", "v"
	}, "line1")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapUsingFuncOptions", actual)
}

func Test_Cov_StringsTo_HashmapUsingFuncTrim(t *testing.T) {
	result := converters.StringsTo.HashmapUsingFuncTrim(func(line string) (string, string) {
		return "k", "v"
	}, "line1")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapUsingFuncTrim", actual)
}

func Test_Cov_StringsTo_MapStringIntegerUsingFunc(t *testing.T) {
	result := converters.StringsTo.MapStringIntegerUsingFunc(true, func(line string) (string, int) {
		return "k", 42
	}, "line1")
	actual := args.Map{"len": len(result), "val": result["k"]}
	expected := args.Map{"len": 1, "val": 42}
	expected.ShouldBeEqual(t, 0, "MapStringIntegerUsingFunc", actual)
}

func Test_Cov_StringsTo_MapStringAnyUsingFunc(t *testing.T) {
	result := converters.StringsTo.MapStringAnyUsingFunc(false, func(line string) (string, any) {
		return "k", "val"
	}, "line1")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringAnyUsingFunc", actual)
}

func Test_Cov_StringsTo_MapConverter(t *testing.T) {
	mc := converters.StringsTo.MapConverter("a:1", "b:2")
	actual := args.Map{"len": mc.Length(), "hasAny": mc.HasAnyItem()}
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "MapConverter", actual)
}

func Test_Cov_StringsTo_IntegersSkipMapAndDefaultValue(t *testing.T) {
	skip := map[string]bool{"skip": true}
	result := converters.StringsTo.IntegersSkipMapAndDefaultValue(-1, skip, "1", "skip", "abc")
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 3, "first": 1, "second": 0}
	expected.ShouldBeEqual(t, 0, "IntegersSkipMapAndDefaultValue", actual)
}

func Test_Cov_StringsTo_IntegersSkipAndDefaultValue(t *testing.T) {
	result := converters.StringsTo.IntegersSkipAndDefaultValue(-1, "skip", "1", "skip", "abc")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": 1}
	expected.ShouldBeEqual(t, 0, "IntegersSkipAndDefaultValue", actual)
}

func Test_Cov_StringsTo_BytesConditional(t *testing.T) {
	result := converters.StringsTo.BytesConditional(func(in string) (byte, bool, bool) {
		return in[0], true, false
	}, []string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BytesConditional", actual)
}

func Test_Cov_StringsTo_BytesWithDefaults(t *testing.T) {
	result := converters.StringsTo.BytesWithDefaults(0, "1", "abc", "300")
	actual := args.Map{
		"len":      result.Length(),
		"hasError": result.HasError(),
		"first":    int(result.Values[0]),
	}
	expected := args.Map{
		"len":      3,
		"hasError": true,
		"first":    1,
	}
	expected.ShouldBeEqual(t, 0, "BytesWithDefaults", actual)
}

func Test_Cov_StringsTo_Csv(t *testing.T) {
	result := converters.StringsTo.Csv(false, "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Csv", actual)
}

func Test_Cov_StringsTo_CsvUsingPtrStrings(t *testing.T) {
	nilResult := converters.StringsTo.CsvUsingPtrStrings(false, nil)
	items := []string{"a", "b"}
	result := converters.StringsTo.CsvUsingPtrStrings(false, &items)
	actual := args.Map{"nilEmpty": nilResult == "", "notEmpty": result != ""}
	expected := args.Map{"nilEmpty": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CsvUsingPtrStrings", actual)
}

func Test_Cov_StringsTo_CsvWithIndexes(t *testing.T) {
	result := converters.StringsTo.CsvWithIndexes([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CsvWithIndexes", actual)
}

func Test_Cov_StringsTo_Float64sMust(t *testing.T) {
	result := converters.StringsTo.Float64sMust("1.5", "2.5")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Float64sMust", actual)
}

func Test_Cov_StringsTo_Float64sConditional(t *testing.T) {
	result := converters.StringsTo.Float64sConditional(func(in string) (float64, bool, bool) {
		return 1.0, true, false
	}, []string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Float64sConditional", actual)
}

func Test_Cov_StringsTo_PtrOfPtrToPtrStrings(t *testing.T) {
	nilResult := converters.StringsTo.PtrOfPtrToPtrStrings(nil)
	actual := args.Map{"nilEmpty": len(*nilResult) == 0}
	expected := args.Map{"nilEmpty": true}
	expected.ShouldBeEqual(t, 0, "PtrOfPtrToPtrStrings_nil", actual)

	s1 := "a"
	s2 := "b"
	var nilStr *string
	items := []*string{&s1, &s2, nilStr}
	result := converters.StringsTo.PtrOfPtrToPtrStrings(&items)
	actual2 := args.Map{"len": len(*result), "third": (*result)[2]}
	expected2 := args.Map{"len": 3, "third": ""}
	expected2.ShouldBeEqual(t, 1, "PtrOfPtrToPtrStrings_values", actual2)
}

func Test_Cov_StringsTo_PtrOfPtrToMapStringBool(t *testing.T) {
	nilResult := converters.StringsTo.PtrOfPtrToMapStringBool(nil)
	actual := args.Map{"nilLen": len(nilResult)}
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "PtrOfPtrToMapStringBool_nil", actual)

	s1 := "a"
	var nilStr *string
	items := []*string{&s1, nilStr}
	result := converters.StringsTo.PtrOfPtrToMapStringBool(&items)
	actual2 := args.Map{"len": len(result), "hasA": result["a"]}
	expected2 := args.Map{"len": 1, "hasA": true}
	expected2.ShouldBeEqual(t, 1, "PtrOfPtrToMapStringBool_values", actual2)
}

func Test_Cov_StringsTo_CloneIf(t *testing.T) {
	original := []string{"a", "b"}
	cloned := converters.StringsTo.CloneIf(true, original...)
	notCloned := converters.StringsTo.CloneIf(false, original...)
	emptyClone := converters.StringsTo.CloneIf(true)
	actual := args.Map{
		"clonedLen": len(cloned),
		"notClonedLen": len(notCloned),
		"emptyLen": len(emptyClone),
	}
	expected := args.Map{
		"clonedLen": 2,
		"notClonedLen": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "CloneIf", actual)
}

// ── StringsToMapConverter coverage ──

func Test_Cov_StringsToMapConverter_SafeStrings(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"a:1"})
	actual := args.Map{"len": len(mc.SafeStrings())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeStrings", actual)

	var emptyMc converters.StringsToMapConverter
	actual2 := args.Map{"len": len(emptyMc.SafeStrings())}
	expected2 := args.Map{"len": 0}
	expected2.ShouldBeEqual(t, 1, "SafeStrings_empty", actual2)
}

func Test_Cov_StringsToMapConverter_LineSplitMapOptions(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"a:1"})
	trimResult := mc.LineSplitMapOptions(true, ":")
	noTrimResult := mc.LineSplitMapOptions(false, ":")
	actual := args.Map{
		"trimLen":   len(trimResult),
		"noTrimLen": len(noTrimResult),
	}
	expected := args.Map{
		"trimLen":   1,
		"noTrimLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "LineSplitMapOptions", actual)
}

func Test_Cov_StringsToMapConverter_LineProcessorMapOptions(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapOptions(true, func(line string) (string, string) {
		return "k", "v"
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapOptions", actual)
}

func Test_Cov_StringsToMapConverter_LineProcessorMapStringIntegerTrim(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapStringIntegerTrim(func(line string) (string, int) {
		return "k", 42
	})
	actual := args.Map{"len": len(result), "val": result["k"]}
	expected := args.Map{"len": 1, "val": 42}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringIntegerTrim", actual)
}

func Test_Cov_StringsToMapConverter_LineProcessorMapStringAnyTrim(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapStringAnyTrim(func(line string) (string, any) {
		return "k", "val"
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringAnyTrim", actual)
}

func Test_Cov_StringsToMapConverter_NilReceiver(t *testing.T) {
	var mc *converters.StringsToMapConverter
	actual := args.Map{
		"length":    mc.Length(),
		"isEmpty":   mc.IsEmpty(),
		"hasAny":    mc.HasAnyItem(),
		"lastIndex": mc.LastIndex(),
	}
	expected := args.Map{
		"length":    0,
		"isEmpty":   true,
		"hasAny":    false,
		"lastIndex": -1,
	}
	expected.ShouldBeEqual(t, 0, "NilReceiver", actual)
}

// ── anyItemConverter extra coverage ──

func Test_Cov_AnyTo_ToStringsUsingProcessor(t *testing.T) {
	nilResult := converters.AnyTo.ToStringsUsingProcessor(true, func(index int, in any) (string, bool, bool) {
		return "x", true, false
	}, nil)
	actual := args.Map{"nilLen": len(nilResult)}
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor_nil", actual)
}

func Test_Cov_AnyTo_ToStringsUsingSimpleProcessor(t *testing.T) {
	nilResult := converters.AnyTo.ToStringsUsingSimpleProcessor(true, func(index int, in any) string {
		return "x"
	}, nil)
	actual := args.Map{"nilLen": len(nilResult)}
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingSimpleProcessor_nil", actual)
}

func Test_Cov_AnyTo_ToNonNullItems(t *testing.T) {
	nilResult := converters.AnyTo.ToNonNullItems(true, nil)
	actual := args.Map{"nilLen": len(nilResult)}
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToNonNullItems_nil", actual)
}

func Test_Cov_AnyTo_SmartStringsJoiner(t *testing.T) {
	result := converters.AnyTo.SmartStringsJoiner(",", "a", "b")
	emptyResult := converters.AnyTo.SmartStringsJoiner(",")
	actual := args.Map{"notEmpty": result != "", "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "SmartStringsJoiner", actual)
}

// ── StringTo IntegerMust ──

func Test_Cov_StringTo_IntegerMust(t *testing.T) {
	result := converters.StringTo.IntegerMust("42")
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "IntegerMust", actual)
}

func Test_Cov_StringTo_Float64Must(t *testing.T) {
	result := converters.StringTo.Float64Must("3.14")
	actual := args.Map{"gt3": result > 3.0}
	expected := args.Map{"gt3": true}
	expected.ShouldBeEqual(t, 0, "Float64Must", actual)
}
