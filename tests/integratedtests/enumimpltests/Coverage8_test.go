package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BasicByte — uncovered branches ──

func Test_Cov8_BasicByte_IsAnyOf_Empty(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B", "C"})
	actual := args.Map{"anyOfEmpty": bb.IsAnyOf(0)}
	expected := args.Map{"anyOfEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns true -- empty checkingItems", actual)
}

func Test_Cov8_BasicByte_IsAnyOf_NotFound(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B", "C"})
	actual := args.Map{"found": bb.IsAnyOf(0, 5, 6)}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns false -- value not in list", actual)
}

func Test_Cov8_BasicByte_IsAnyNamesOf_NotFound(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B", "C"})
	actual := args.Map{"found": bb.IsAnyNamesOf(0, "X", "Y")}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns false -- name not in list", actual)
}

func Test_Cov8_BasicByte_GetValueByName_NotFound(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	_, err := bb.GetValueByName("UNKNOWN")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns error -- unknown name", actual)
}

func Test_Cov8_BasicByte_ToEnumJsonBytes_NotFound(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	_, err := bb.ToEnumJsonBytes(99)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes returns error -- value not in map", actual)
}

func Test_Cov8_BasicByte_ExpectingEnumValueError_Mismatch(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	err := bb.ExpectingEnumValueError("B", byte(0))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- mismatch value", actual)
}

func Test_Cov8_BasicByte_ExpectingEnumValueError_UnknownInput(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	err := bb.ExpectingEnumValueError("UNKNOWN", byte(0))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- unknown input", actual)
}

func Test_Cov8_BasicByte_UnmarshallToValue_NilNotMapped(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	_, err := bb.UnmarshallToValue(false, nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns error -- nil not mapped", actual)
}

func Test_Cov8_BasicByte_UnmarshallToValue_NilMapped(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	val, err := bb.UnmarshallToValue(true, nil)
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": byte(0), "noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped to first", actual)
}

func Test_Cov8_BasicByte_UnmarshallToValue_EmptyMapped(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	val, err := bb.UnmarshallToValue(true, []byte(`""`))
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": byte(0), "noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty string mapped", actual)
}

// ── BasicString — uncovered branches ──

func Test_Cov8_BasicString_IsAnyOf_Empty(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	actual := args.Map{"anyOfEmpty": bs.IsAnyOf("X")}
	expected := args.Map{"anyOfEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns true -- empty checkingItems", actual)
}

func Test_Cov8_BasicString_IsAnyNamesOf_NotFound(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	actual := args.Map{"found": bs.IsAnyNamesOf("X", "Z", "W")}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns false -- name not in list", actual)
}

func Test_Cov8_BasicString_GetValueByName_NotFound(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	_, err := bs.GetValueByName("UNKNOWN")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns error -- unknown name", actual)
}

func Test_Cov8_BasicString_ToEnumJsonBytes_NotFound(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	_, err := bs.ToEnumJsonBytes("UNKNOWN")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes returns error -- unknown value", actual)
}

func Test_Cov8_BasicString_UnmarshallToValue_NilNotMapped(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	_, err := bs.UnmarshallToValue(false, nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns error -- nil not mapped", actual)
}

func Test_Cov8_BasicString_UnmarshallToValue_NilMapped(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	val, err := bs.UnmarshallToValue(true, nil)
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": "X", "noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped to first", actual)
}

func Test_Cov8_BasicString_UnmarshallToValue_EmptyMapped(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	val, err := bs.UnmarshallToValue(true, []byte(`""`))
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": "X", "noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty string mapped", actual)
}

func Test_Cov8_BasicString_GetNameByIndex_OutOfRange(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	actual := args.Map{"name": bs.GetNameByIndex(99)}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "GetNameByIndex returns empty -- out of range", actual)
}

func Test_Cov8_BasicString_GetIndexByName_Empty(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	actual := args.Map{"idx": bs.GetIndexByName("")}
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "GetIndexByName returns InvalidValue -- empty name", actual)
}

func Test_Cov8_BasicString_GetIndexByName_NotFound(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{"X", "Y"})
	actual := args.Map{"idx": bs.GetIndexByName("ZZZ")}
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "GetIndexByName returns InvalidValue -- unknown name", actual)
}

// ── newBasicStringCreator — uncovered branches ──

func Test_Cov8_BasicStringCreator_CreateUsingStringersSpread(t *testing.T) {
	type testStringer struct{ val string }
	bs := enumimpl.New.BasicString.CreateUsingNamesSpread("TestEnum", "Alpha", "Beta", "Gamma")
	actual := args.Map{"len": bs.Length(), "hasAlpha": bs.IsValidRange("Alpha")}
	expected := args.Map{"len": 3, "hasAlpha": true}
	expected.ShouldBeEqual(t, 0, "CreateUsingNamesSpread returns valid enum -- three names", actual)
}

func Test_Cov8_BasicStringCreator_UsingFirstItemSliceAllCases(t *testing.T) {
	type testEnum string
	bs := enumimpl.New.BasicString.UsingFirstItemSliceAllCases(testEnum("A"), []string{"A", "B"})
	actual := args.Map{"len": bs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UsingFirstItemSliceAllCases returns enum -- with case aliases", actual)
}

func Test_Cov8_BasicStringCreator_CreateUsingSlicePlusAliasMapOptions(t *testing.T) {
	type testEnum string
	aliases := map[string]string{"alpha": "A"}
	bs := enumimpl.New.BasicString.CreateUsingSlicePlusAliasMapOptions(true, testEnum("A"), []string{"A", "B"}, aliases)
	actual := args.Map{"len": bs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingSlicePlusAliasMapOptions returns enum -- with aliases", actual)
}

// ── newBasicByteCreator — uncovered branches ──

func Test_Cov8_BasicByteCreator_CreateUsingMap(t *testing.T) {
	m := map[byte]string{0: "Off", 1: "On"}
	bb := enumimpl.New.BasicByte.CreateUsingMap("TestEnum", m)
	actual := args.Map{"len": bb.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingMap returns enum -- two entries", actual)
}

func Test_Cov8_BasicByteCreator_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	type testEnum byte
	m := map[byte]string{0: "Off", 1: "On"}
	aliases := map[string]byte{"off": 0}
	bb := enumimpl.New.BasicByte.CreateUsingMapPlusAliasMapOptions(true, testEnum(0), m, aliases)
	actual := args.Map{"len": bb.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingMapPlusAliasMapOptions returns enum -- with aliases", actual)
}

func Test_Cov8_BasicByteCreator_DefaultAllCases(t *testing.T) {
	type testEnum byte
	bb := enumimpl.New.BasicByte.DefaultAllCases(testEnum(0), []string{"Off", "On"})
	actual := args.Map{"len": bb.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultAllCases returns enum -- two entries all cases", actual)
}

func Test_Cov8_BasicByteCreator_DefaultWithAliasMapAllCases(t *testing.T) {
	type testEnum byte
	aliases := map[string]byte{"off": 0}
	bb := enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(testEnum(0), []string{"Off", "On"}, aliases)
	actual := args.Map{"len": bb.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultWithAliasMapAllCases returns enum -- with alias all cases", actual)
}

func Test_Cov8_BasicByteCreator_UsingFirstItemSliceAllCases(t *testing.T) {
	type testEnum byte
	bb := enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(testEnum(0), []string{"Off", "On"})
	actual := args.Map{"len": bb.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UsingFirstItemSliceAllCases returns enum -- all cases", actual)
}

// ── DynamicMap — uncovered branches ──

func Test_Cov8_DynamicMap_AddNewOnly_Existing(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	added := dm.AddNewOnly("a", 2)
	actual := args.Map{"added": added, "val": dm["a"]}
	expected := args.Map{"added": false, "val": 1}
	expected.ShouldBeEqual(t, 0, "AddNewOnly returns false -- key exists", actual)
}

func Test_Cov8_DynamicMap_AddNewOnly_New(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	added := dm.AddNewOnly("b", 2)
	actual := args.Map{"added": added}
	expected := args.Map{"added": true}
	expected.ShouldBeEqual(t, 0, "AddNewOnly returns true -- new key", actual)
}

func Test_Cov8_DynamicMap_HasAllKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	actual := args.Map{
		"all":     dm.HasAllKeys("a", "b"),
		"missing": dm.HasAllKeys("a", "c"),
	}
	expected := args.Map{
		"all":     true,
		"missing": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAllKeys returns correct -- present and missing", actual)
}

func Test_Cov8_DynamicMap_HasAnyKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{
		"any":  dm.HasAnyKeys("a", "b"),
		"none": dm.HasAnyKeys("c", "d"),
	}
	expected := args.Map{
		"any":  true,
		"none": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyKeys returns correct -- has and none", actual)
}

func Test_Cov8_DynamicMap_IsEqual_SamePointer(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{"equal": dm.IsEqual(false, &dm)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- same pointer", actual)
}

func Test_Cov8_DynamicMap_IsRawEqual_DiffLength(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{"equal": dm.IsRawEqual(false, map[string]any{"a": 1, "b": 2})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns false -- different length", actual)
}

func Test_Cov8_DynamicMap_IsRawEqual_MissingKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{"equal": dm.IsRawEqual(false, map[string]any{"b": 1})}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsRawEqual returns false -- missing key", actual)
}

// ── DiffLeftRight — uncovered branches ──

func Test_Cov8_DiffLeftRight_HasMismatch_Regardless(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: 1, Right: int64(1)}
	actual := args.Map{
		"mismatch":  dlr.HasMismatch(false),
		"noMismatch": !dlr.HasMismatch(true),
	}
	expected := args.Map{
		"mismatch":  true,
		"noMismatch": true,
	}
	expected.ShouldBeEqual(t, 0, "HasMismatch returns correct -- regardless type and strict", actual)
}

func Test_Cov8_DiffLeftRight_DiffString_Same(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}
	actual := args.Map{"diffStr": dlr.DiffString()}
	expected := args.Map{"diffStr": ""}
	expected.ShouldBeEqual(t, 0, "DiffString returns empty -- same values", actual)
}

func Test_Cov8_DiffLeftRight_JsonString_Nil(t *testing.T) {
	var dlr *enumimpl.DiffLeftRight
	actual := args.Map{"jsonStr": dlr.JsonString()}
	expected := args.Map{"jsonStr": ""}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- nil pointer", actual)
}

// ── toStringPrintableDynamicMap — covered via DynamicMap.String() ──

func Test_Cov8_DynamicMap_AllValuesIntegers(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	vals := dm.AllValuesIntegers()
	actual := args.Map{"len": len(vals)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllValuesIntegers returns correct length -- two int values", actual)
}

func Test_Cov8_DynamicMap_AllValuesIntegers_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	vals := dm.AllValuesIntegers()
	actual := args.Map{"len": len(vals)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AllValuesIntegers returns empty -- empty map", actual)
}

// ── numberEnumBase — uncovered branches ──

func Test_Cov8_NumberEnumBase_MinValueString(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	min1 := bb.MinValueString()
	min2 := bb.MinValueString() // cached
	actual := args.Map{"same": min1 == min2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MinValueString returns cached -- second call", actual)
}

func Test_Cov8_NumberEnumBase_MaxValueString(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A", "B"})
	max1 := bb.MaxValueString()
	max2 := bb.MaxValueString() // cached
	actual := args.Map{"same": max1 == max2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "MaxValueString returns cached -- second call", actual)
}

func Test_Cov8_NumberEnumBase_NamesHashset_Empty(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{})
	m := bs.NamesHashset()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NamesHashset returns empty -- no ranges", actual)
}

func Test_Cov8_NumberEnumBase_NameWithValueOption_WithQuotation(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestEnum", []string{"A"})
	withQ := bb.NameWithValueOption(byte(0), true)
	withoutQ := bb.NameWithValueOption(byte(0), false)
	actual := args.Map{
		"hasQuote":   len(withQ) > 0,
		"hasNoQuote": len(withoutQ) > 0,
	}
	expected := args.Map{
		"hasQuote":   true,
		"hasNoQuote": true,
	}
	expected.ShouldBeEqual(t, 0, "NameWithValueOption returns non-empty -- both modes", actual)
}

// ── toHashset — empty case ──

func Test_Cov8_ToHashset_Empty(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestEnum", []string{})
	m := bs.NamesHashset()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NamesHashset returns empty map -- empty input", actual)
}

// ── ConvEnumAnyValToInteger — fallback sprintf path ──

func Test_Cov8_ConvEnumAnyValToInteger_FallbackAtoi(t *testing.T) {
	val := enumimpl.ConvEnumAnyValToInteger(float64(42))
	// float64 falls to Sprintf path → "42" → Atoi → 42
	actual := args.Map{"val": val}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns 42 -- float64 fallback", actual)
}

func Test_Cov8_ConvEnumAnyValToInteger_FallbackNonNumeric(t *testing.T) {
	val := enumimpl.ConvEnumAnyValToInteger(struct{}{})
	actual := args.Map{"isMinInt": val < 0}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns MinInt -- non-numeric struct", actual)
}
