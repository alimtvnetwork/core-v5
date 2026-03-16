package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DiffLeftRight ──

func Test_Cov10_DiffLeftRight_NilJsonString(t *testing.T) {
	var dlr *enumimpl.DiffLeftRight
	actual := args.Map{"result": dlr.JsonString()}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- nil receiver", actual)
}

func Test_Cov10_DiffLeftRight_SameValues(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "abc", Right: "abc"}
	actual := args.Map{
		"isSame":           dlr.IsSame(),
		"isSameType":       dlr.IsSameTypeSame(),
		"isEqual":          dlr.IsEqual(false),
		"isEqualRegardless": dlr.IsEqual(true),
		"isNotEqual":       dlr.IsNotEqual(),
		"diffString":       dlr.DiffString(),
	}
	expected := args.Map{
		"isSame":           true,
		"isSameType":       true,
		"isEqual":          true,
		"isEqualRegardless": true,
		"isNotEqual":       false,
		"diffString":       "",
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight all same -- equal strings", actual)
}

func Test_Cov10_DiffLeftRight_DifferentValues(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "abc", Right: "xyz"}
	actual := args.Map{
		"isSame":       dlr.IsSame(),
		"isNotEqual":   dlr.IsNotEqual(),
		"hasMismatch":  dlr.HasMismatch(false),
		"hasMismatchR": dlr.HasMismatch(true),
		"hasMismatchRegardless": dlr.HasMismatchRegardlessOfType(),
		"diffNotEmpty": dlr.DiffString() != "",
	}
	expected := args.Map{
		"isSame":       false,
		"isNotEqual":   true,
		"hasMismatch":  true,
		"hasMismatchR": true,
		"hasMismatchRegardless": true,
		"diffNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight all different -- different strings", actual)
}

func Test_Cov10_DiffLeftRight_DifferentTypes(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: 1, Right: "1"}
	actual := args.Map{
		"sameTypeSame":       dlr.IsSameTypeSame(),
		"sameRegardless":     dlr.IsSameRegardlessOfType(),
	}
	expected := args.Map{
		"sameTypeSame":       false,
		"sameRegardless":     true,
	}
	expected.ShouldBeEqual(t, 0, "DiffLeftRight different types same value -- int vs string", actual)
}

func Test_Cov10_DiffLeftRight_SpecificFullString(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	l, r := dlr.SpecificFullString()
	actual := args.Map{"lNotEmpty": l != "", "rNotEmpty": r != ""}
	expected := args.Map{"lNotEmpty": true, "rNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "SpecificFullString returns non-empty -- both sides", actual)
}

func Test_Cov10_DiffLeftRight_Types(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: 1, Right: "x"}
	l, r := dlr.Types()
	actual := args.Map{"lNotNil": l != nil, "rNotNil": r != nil, "different": l != r}
	expected := args.Map{"lNotNil": true, "rNotNil": true, "different": true}
	expected.ShouldBeEqual(t, 0, "Types returns reflect types -- different types", actual)
}

func Test_Cov10_DiffLeftRight_StringNonNil(t *testing.T) {
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	actual := args.Map{"notEmpty": dlr.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns json -- non-nil", actual)
}

// ── DynamicMap — ConcatNew ──

func Test_Cov10_DynamicMap_ConcatNew_OverrideExisting(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	other := enumimpl.DynamicMap{"b": 99, "c": 3}
	result := dm.ConcatNew(true, other)
	actual := args.Map{"b": result["b"], "c": result["c"], "a": result["a"]}
	expected := args.Map{"b": 99, "c": 3, "a": 1}
	expected.ShouldBeEqual(t, 0, "ConcatNew overrides existing -- b becomes 99", actual)
}

func Test_Cov10_DynamicMap_ConcatNew_NoOverride(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	other := enumimpl.DynamicMap{"b": 99, "c": 3}
	result := dm.ConcatNew(false, other)
	actual := args.Map{"b": result["b"], "c": result["c"]}
	expected := args.Map{"b": 2, "c": 3}
	expected.ShouldBeEqual(t, 0, "ConcatNew no override -- b stays 2", actual)
}

func Test_Cov10_DynamicMap_ConcatNew_BothEmpty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	other := enumimpl.DynamicMap{}
	result := dm.ConcatNew(true, other)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ConcatNew both empty -- zero length", actual)
}

// ── DynamicMap — StringsUsingFmt ──

func Test_Cov10_DynamicMap_StringsUsingFmt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return key + "=" + enumimpl.NameWithValue(val)
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsUsingFmt returns formatted -- two items", actual)
}

func Test_Cov10_DynamicMap_StringsUsingFmt_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	result := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return ""
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsUsingFmt returns empty -- empty map", actual)
}

// ── DynamicMap — Serialize ──

func Test_Cov10_DynamicMap_Serialize(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	b, err := dm.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize returns json bytes -- single item", actual)
}

// ── DynamicMap — IsStringEqual ──

func Test_Cov10_DynamicMap_IsStringEqual(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{
		"selfEqual":    dm.IsStringEqual(dm.String()),
		"notEqual":     dm.IsStringEqual("nope"),
	}
	expected := args.Map{
		"selfEqual":    true,
		"notEqual":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsStringEqual matches own string -- true/false", actual)
}

// ── DynamicMap — IsKeysEqualOnly ──

func Test_Cov10_DynamicMap_IsKeysEqualOnly_Match(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 99, "b": 88}
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly true -- same keys different values", actual)
}

func Test_Cov10_DynamicMap_IsKeysEqualOnly_Mismatch(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 99, "b": 88}
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly false -- different key count", actual)
}

func Test_Cov10_DynamicMap_IsKeysEqualOnly_BothNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	actual := args.Map{"equal": dm.IsKeysEqualOnly(nil)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly true -- both nil", actual)
}

func Test_Cov10_DynamicMap_IsKeysEqualOnly_OneNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly false -- left nil right non-nil", actual)
}

func Test_Cov10_DynamicMap_IsKeysEqualOnly_MissingKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "c": 3}
	right := map[string]any{"a": 99, "b": 88}
	actual := args.Map{"equal": dm.IsKeysEqualOnly(right)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsKeysEqualOnly false -- key mismatch", actual)
}

// ── DynamicMap — KeyValue / KeyValueString ──

func Test_Cov10_DynamicMap_KeyValue_Found(t *testing.T) {
	dm := enumimpl.DynamicMap{"x": 42}
	val, found := dm.KeyValue("x")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": 42, "found": true}
	expected.ShouldBeEqual(t, 0, "KeyValue found -- existing key", actual)
}

func Test_Cov10_DynamicMap_KeyValue_NotFound(t *testing.T) {
	dm := enumimpl.DynamicMap{"x": 42}
	_, found := dm.KeyValue("z")
	actual := args.Map{"found": found}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "KeyValue not found -- missing key", actual)
}

func Test_Cov10_DynamicMap_KeyValueString_NotFound(t *testing.T) {
	dm := enumimpl.DynamicMap{"x": 42}
	val, found := dm.KeyValueString("z")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": "", "found": false}
	expected.ShouldBeEqual(t, 0, "KeyValueString returns empty -- missing key", actual)
}

// ── DynamicMap — Add ──

func Test_Cov10_DynamicMap_Add(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.Add("b", 2)
	actual := args.Map{"len": dm.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Add adds key -- length increases", actual)
}

// ── DynamicMap — HasAllKeys / HasAnyKeys ──

func Test_Cov10_DynamicMap_HasAllKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2, "c": 3}
	actual := args.Map{
		"allPresent": dm.HasAllKeys("a", "b"),
		"oneMissing": dm.HasAllKeys("a", "z"),
	}
	expected := args.Map{
		"allPresent": true,
		"oneMissing": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAllKeys correct -- present and missing", actual)
}

func Test_Cov10_DynamicMap_HasAnyKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	actual := args.Map{
		"onePresent": dm.HasAnyKeys("z", "a"),
		"nonPresent": dm.HasAnyKeys("x", "y"),
	}
	expected := args.Map{
		"onePresent": true,
		"nonPresent": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyKeys correct -- found and not found", actual)
}

// ── DynamicMap — Set / AddNewOnly ──

func Test_Cov10_DynamicMap_AddOrUpdate(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.AddOrUpdate("b", 2)
	isUpdate := dm.AddOrUpdate("a", 99)
	actual := args.Map{"isNew": isNew, "isUpdate": isUpdate, "aVal": dm["a"]}
	expected := args.Map{"isNew": true, "isUpdate": false, "aVal": 99}
	expected.ShouldBeEqual(t, 0, "AddOrUpdate new then update -- correct flags", actual)
}

// ── DynamicMap — Diff with LeftRightDiffCheckerImpl ──

func Test_Cov10_DynamicMap_DiffJsonMessageLeftRight_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "b": 2}
	result := dm.DiffJsonMessageLeftRight(false, right)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageLeftRight empty -- no diff", actual)
}

func Test_Cov10_DynamicMap_DiffJsonMessageLeftRight_HasDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}
	result := dm.DiffJsonMessageLeftRight(false, right)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageLeftRight non-empty -- has diff", actual)
}

func Test_Cov10_DynamicMap_LogShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.LogShouldDiffLeftRightMessage(false, "test", right)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LogShouldDiffLeftRightMessage empty -- no diff", actual)
}

// ── DynamicMap — ExpectingMessage ──

func Test_Cov10_DynamicMap_ExpectingMessage_Equal(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.ExpectingMessage("test", map[string]any{"a": 1})
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingMessage empty -- maps equal", actual)
}

func Test_Cov10_DynamicMap_ExpectingMessage_NotEqual(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.ExpectingMessage("test", map[string]any{"a": 2})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingMessage non-empty -- maps differ", actual)
}

func Test_Cov10_DynamicMap_LogExpectingMessage_Equal(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
	// no panic = pass
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LogExpectingMessage no panic -- equal maps", actual)
}

// ── DynamicMap — IsMismatch / IsRawMismatch ──

func Test_Cov10_DynamicMap_IsMismatch(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := enumimpl.DynamicMap{"a": 2}
	actual := args.Map{"mismatch": dm.IsMismatch(false, &right)}
	expected := args.Map{"mismatch": true}
	expected.ShouldBeEqual(t, 0, "IsMismatch true -- values differ", actual)
}

func Test_Cov10_DynamicMap_IsRawMismatch(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 2}
	actual := args.Map{"mismatch": dm.IsRawMismatch(false, right)}
	expected := args.Map{"mismatch": true}
	expected.ShouldBeEqual(t, 0, "IsRawMismatch true -- values differ", actual)
}

// ── DynamicMap — IsEqual edge cases ──

func Test_Cov10_DynamicMap_IsEqual_SamePointer(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{"equal": dm.IsEqual(false, &dm)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual true -- same pointer", actual)
}

func Test_Cov10_DynamicMap_IsEqual_Regardless(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{"equal": dm.IsEqual(true, &right)}
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual regardless true -- same values", actual)
}

func Test_Cov10_DynamicMap_IsRawEqual_MissingKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}
	actual := args.Map{"equal": dm.IsRawEqual(false, right)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsRawEqual false -- key mismatch", actual)
}

// ── DynamicMap — DiffRaw nil cases ──

func Test_Cov10_DynamicMap_DiffRawUsingDifferChecker_LeftNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}
	result := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns right -- left nil", actual)
}

func Test_Cov10_DynamicMap_DiffRawUsingDifferChecker_RightNil(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns left -- right nil", actual)
}

func Test_Cov10_DynamicMap_DiffRawLeftRight_BothNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	actual := args.Map{"lLen": l.Length(), "rLen": r.Length()}
	expected := args.Map{"lLen": 0, "rLen": 0}
	expected.ShouldBeEqual(t, 0, "DiffRawLeftRight both empty -- both nil", actual)
}

func Test_Cov10_DynamicMap_DiffRawLeftRight_LeftNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)
	actual := args.Map{"lLen": l.Length(), "rLen": r.Length()}
	expected := args.Map{"lLen": 1, "rLen": 0}
	expected.ShouldBeEqual(t, 0, "DiffRawLeftRight returns right as lDiff -- left nil", actual)
}

func Test_Cov10_DynamicMap_DiffRawLeftRight_RightNil(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	actual := args.Map{"lLen": l.Length(), "rLen": r.Length()}
	expected := args.Map{"lLen": 1, "rLen": 0}
	expected.ShouldBeEqual(t, 0, "DiffRawLeftRight returns left as lDiff -- right nil", actual)
}

// ── DynamicMap — ShouldDiff messages ──

func Test_Cov10_DynamicMap_ShouldDiffMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.ShouldDiffMessage(false, "title", right)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage empty -- no diff", actual)
}

func Test_Cov10_DynamicMap_ShouldDiffMessage_HasDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 2}
	result := dm.ShouldDiffMessage(false, "title", right)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffMessage non-empty -- has diff", actual)
}

func Test_Cov10_DynamicMap_ShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "title", right,
	)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffLeftRightMessage empty -- no diff", actual)
}

func Test_Cov10_DynamicMap_ShouldDiffLeftRightMessage_HasDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"b": 2}
	result := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "title", right,
	)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldDiffLeftRightMessage non-empty -- has diff", actual)
}

// ── DynamicMap — LogShouldDiffMessageUsingDifferChecker ──

func Test_Cov10_DynamicMap_LogShouldDiffMessageUsingDifferChecker_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", right,
	)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LogShouldDiffMessage empty -- no diff", actual)
}

// ── DynamicMap — DiffJsonMessage ──

func Test_Cov10_DynamicMap_DiffJsonMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.DiffJsonMessage(false, right)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessage empty -- no diff", actual)
}

func Test_Cov10_DynamicMap_DiffJsonMessageUsingDifferChecker_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1}
	result := dm.DiffJsonMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, right,
	)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DiffJsonMessageUsingDifferChecker empty -- no diff", actual)
}

// ── DynamicMap — AllValuesIntegers / AllValuesStrings ──

func Test_Cov10_DynamicMap_AllValuesIntegers(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.AllValuesIntegers()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllValuesIntegers returns integers -- two items", actual)
}

func Test_Cov10_DynamicMap_AllValuesStringsSorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	result := dm.AllValuesStringsSorted()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllValuesStringsSorted returns sorted -- two items", actual)
}

// ── DynamicMap — HasIndex / LastIndex / Count ──

func Test_Cov10_DynamicMap_HasIndex(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	actual := args.Map{
		"hasIdx0": dm.HasIndex(0),
		"hasIdx1": dm.HasIndex(1),
		"hasIdx5": dm.HasIndex(5),
	}
	expected := args.Map{
		"hasIdx0": true,
		"hasIdx1": true,
		"hasIdx5": false,
	}
	expected.ShouldBeEqual(t, 0, "HasIndex correct -- two items", actual)
}

func Test_Cov10_DynamicMap_Count(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	actual := args.Map{"count": dm.Count()}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Count returns 1 -- single item", actual)
}

// ── DynamicMap — SortedKeyValues with string values ──

func Test_Cov10_DynamicMap_SortedKeyAnyValues_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"b": "beta", "a": "alpha"}
	result := dm.SortedKeyAnyValues()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyAnyValues returns sorted -- string values", actual)
}

// ── DynamicMap — MapIntegerString with string values ──

func Test_Cov10_DynamicMap_MapIntegerString_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "alpha", "b": "beta"}
	rangeMap, sortedKeys := dm.MapIntegerString()
	actual := args.Map{"mapLen": len(rangeMap), "keysLen": len(sortedKeys)}
	expected := args.Map{"mapLen": 2, "keysLen": 2}
	expected.ShouldBeEqual(t, 0, "MapIntegerString handles string values -- two items", actual)
}

// ── DynamicMap — ConvMap variants ──

func Test_Cov10_DynamicMap_ConvMapStringString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "alpha", "b": "beta"}
	result := dm.ConvMapStringString()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapStringString returns map -- two items", actual)
}

func Test_Cov10_DynamicMap_ConvMapInt64String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.ConvMapInt64String()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapInt64String returns map -- two items", actual)
}

func Test_Cov10_DynamicMap_ConvMapStringInteger(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	result := dm.ConvMapStringInteger()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapStringInteger returns map -- two items", actual)
}

// ── DynamicMap — KeyValueByte edge cases ──

func Test_Cov10_DynamicMap_KeyValueByte_NotFound(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_, isFound, _ := dm.KeyValueByte("z")
	actual := args.Map{"found": isFound}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "KeyValueByte not found -- missing key", actual)
}

func Test_Cov10_DynamicMap_KeyValueByte_DirectByte(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": byte(42)}
	val, isFound, isFailed := dm.KeyValueByte("a")
	actual := args.Map{"val": val, "found": isFound, "failed": isFailed}
	expected := args.Map{"val": byte(42), "found": true, "failed": false}
	expected.ShouldBeEqual(t, 0, "KeyValueByte returns byte -- direct byte value", actual)
}

// ── DynamicMap — KeyValueInt edge cases ──

func Test_Cov10_DynamicMap_KeyValueInt_DirectInt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 42}
	val, isFound, isFailed := dm.KeyValueInt("a")
	actual := args.Map{"val": val, "found": isFound, "failed": isFailed}
	expected := args.Map{"val": 42, "found": true, "failed": false}
	expected.ShouldBeEqual(t, 0, "KeyValueInt returns int -- direct int value", actual)
}

func Test_Cov10_DynamicMap_KeyValueInt_NotFound(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_, isFound, isFailed := dm.KeyValueInt("z")
	actual := args.Map{"found": isFound, "failed": isFailed}
	expected := args.Map{"found": false, "failed": true}
	expected.ShouldBeEqual(t, 0, "KeyValueInt not found -- missing key", actual)
}

func Test_Cov10_DynamicMap_KeyValueInt_DirectByte(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": byte(5)}
	val, isFound, isFailed := dm.KeyValueInt("a")
	actual := args.Map{"val": val, "found": isFound, "failed": isFailed}
	expected := args.Map{"val": 5, "found": true, "failed": false}
	expected.ShouldBeEqual(t, 0, "KeyValueInt converts byte -- direct byte value", actual)
}

// ── DynamicMap — IsValueTypeOf ──

func Test_Cov10_DynamicMap_IsValueTypeOf(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "hello"}
	actual := args.Map{"isString": dm.IsValueString()}
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "IsValueString true -- string value", actual)
}

// ── Format / FormatUsingFmt ──

func Test_Cov10_Format(t *testing.T) {
	result := enumimpl.Format("MyType", "MyName", "42", "Enum of {type-name} - {name} - {value}")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "Enum of MyType - MyName - 42"}
	expected.ShouldBeEqual(t, 0, "Format compiles template -- all keys replaced", actual)
}

// ── KeyAnyVal ──

func Test_Cov10_KeyAnyVal_StringMethods(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "Name", AnyValue: 5}
	actual := args.Map{
		"keyString":   kav.KeyString(),
		"anyVal":      kav.AnyVal(),
		"anyValStr":   kav.AnyValString() != "",
		"wrapKey":     kav.WrapKey() != "",
		"wrapValue":   kav.WrapValue() != "",
		"isString":    kav.IsString(),
		"valInt":      kav.ValInt(),
		"stringOut":   kav.String() != "",
	}
	expected := args.Map{
		"keyString":   "Name",
		"anyVal":      5,
		"anyValStr":   true,
		"wrapKey":     true,
		"wrapValue":   true,
		"isString":    false,
		"valInt":      5,
		"stringOut":   true,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal methods -- integer value", actual)
}

func Test_Cov10_KeyAnyVal_IsString_StringValue(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "Name", AnyValue: "hello"}
	actual := args.Map{"isString": kav.IsString(), "str": kav.String() != ""}
	expected := args.Map{"isString": true, "str": true}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal IsString true -- string value", actual)
}

func Test_Cov10_KeyAnyVal_KeyValInteger(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "Name", AnyValue: 5}
	kvi := kav.KeyValInteger()
	actual := args.Map{"key": kvi.Key, "val": kvi.ValueInteger}
	expected := args.Map{"key": "Name", "val": 5}
	expected.ShouldBeEqual(t, 0, "KeyValInteger converts correctly -- int value", actual)
}

// ── KeyValInteger ──

func Test_Cov10_KeyValInteger_Methods(t *testing.T) {
	kvi := enumimpl.KeyValInteger{Key: "Name", ValueInteger: 5}
	kav := kvi.KeyAnyVal()
	actual := args.Map{
		"wrapKey":    kvi.WrapKey() != "",
		"wrapValue":  kvi.WrapValue() != "",
		"isString":   kvi.IsString(),
		"stringOut":  kvi.String() != "",
		"kavKey":     kav.Key,
	}
	expected := args.Map{
		"wrapKey":    true,
		"wrapValue":  true,
		"isString":   false,
		"stringOut":  true,
		"kavKey":     "Name",
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger methods -- integer value", actual)
}

// ── KeyAnyValues func ──

func Test_Cov10_KeyAnyValues_Empty(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{}, []byte{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns empty -- empty input", actual)
}

func Test_Cov10_KeyAnyValues_NonEmpty(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{"A", "B"}, []byte{0, 1})
	actual := args.Map{"len": len(result), "firstKey": result[0].Key}
	expected := args.Map{"len": 2, "firstKey": "A"}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns items -- two entries", actual)
}

// ── AllNameValues func ──

func Test_Cov10_AllNameValues(t *testing.T) {
	result := enumimpl.AllNameValues([]string{"A", "B"}, []byte{0, 1})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns strings -- two entries", actual)
}

// ── IntegersRangesOfAnyVal ──

func Test_Cov10_IntegersRangesOfAnyVal(t *testing.T) {
	result := enumimpl.IntegersRangesOfAnyVal([]byte{2, 0, 1})
	actual := args.Map{"first": result[0], "last": result[2]}
	expected := args.Map{"first": 0, "last": 2}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal sorted -- byte input", actual)
}

// ── PrependJoin / JoinPrependUsingDot ──

func Test_Cov10_PrependJoin(t *testing.T) {
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "prefix.a.b"}
	expected.ShouldBeEqual(t, 0, "PrependJoin joins with dot -- three parts", actual)
}

func Test_Cov10_JoinPrependUsingDot(t *testing.T) {
	result := enumimpl.JoinPrependUsingDot("prefix", "a", "b")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "prefix.a.b"}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot joins -- three parts", actual)
}

// ── NameWithValue func ──

func Test_Cov10_NameWithValue(t *testing.T) {
	result := enumimpl.NameWithValue(42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValue returns formatted -- integer", actual)
}

// ── OnlySupportedErr / UnsupportedNames ──

func Test_Cov10_OnlySupportedErr_NoUnsupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(1, []string{"A", "B"}, "A", "B")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr nil -- all supported", actual)
}

func Test_Cov10_OnlySupportedErr_HasUnsupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(1, []string{"A", "B", "C"}, "A")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr error -- B and C unsupported", actual)
}

func Test_Cov10_OnlySupportedErr_EmptyAllNames(t *testing.T) {
	err := enumimpl.OnlySupportedErr(1, []string{}, "A")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr nil -- empty allNames", actual)
}

func Test_Cov10_UnsupportedNames(t *testing.T) {
	result := enumimpl.UnsupportedNames([]string{"A", "B", "C"}, "A")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns B and C -- A supported", actual)
}

// ── BasicString — additional coverage ──

func Test_Cov10_BasicString_GetNameByIndex(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"Alpha", "Beta", "Gamma"})
	actual := args.Map{
		"idx1":      bs.GetNameByIndex(1),
		"idx0":      bs.GetNameByIndex(0),
		"idxNeg":    bs.GetNameByIndex(-1),
		"idxTooBig": bs.GetNameByIndex(99),
	}
	expected := args.Map{
		"idx1":      "Beta",
		"idx0":      "",
		"idxNeg":    "",
		"idxTooBig": "",
	}
	expected.ShouldBeEqual(t, 0, "GetNameByIndex returns name or empty -- boundary checks", actual)
}

func Test_Cov10_BasicString_GetIndexByName(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"Alpha", "Beta"})
	actual := args.Map{
		"alpha":   bs.GetIndexByName("Alpha"),
		"unknown": bs.GetIndexByName("Unknown"),
		"empty":   bs.GetIndexByName(""),
	}
	expected := args.Map{
		"alpha":   0,
		"unknown": -1,
		"empty":   -1,
	}
	expected.ShouldBeEqual(t, 0, "GetIndexByName returns index or -1 -- various inputs", actual)
}

func Test_Cov10_BasicString_RangesIntegers(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})
	result := bs.RangesIntegers()
	actual := args.Map{"len": len(result), "last": result[2]}
	expected := args.Map{"len": 3, "last": 2}
	expected.ShouldBeEqual(t, 0, "RangesIntegers returns 0-based -- three items", actual)
}

func Test_Cov10_BasicString_Hashset(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	actual := args.Map{"hasItems": len(bs.Hashset()) > 0, "ptrNotNil": bs.HashsetPtr() != nil}
	expected := args.Map{"hasItems": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset and HashsetPtr non-empty -- two items", actual)
}

func Test_Cov10_BasicString_IsAnyOf(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	actual := args.Map{
		"emptyCheck": bs.IsAnyOf("A"),
		"found":      bs.IsAnyOf("A", "A", "B"),
		"notFound":   bs.IsAnyOf("A", "X", "Y"),
	}
	expected := args.Map{
		"emptyCheck": true,
		"found":      true,
		"notFound":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyOf correct -- various inputs", actual)
}

func Test_Cov10_BasicString_IsAnyNamesOf(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	actual := args.Map{
		"found":    bs.IsAnyNamesOf("A", "A"),
		"notFound": bs.IsAnyNamesOf("A", "X"),
	}
	expected := args.Map{
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf correct -- found and not found", actual)
}

func Test_Cov10_BasicString_HasAnyItem(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A"})
	actual := args.Map{"hasAny": bs.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem true -- one item", actual)
}

func Test_Cov10_BasicString_MaxIndex(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})
	actual := args.Map{"maxIdx": bs.MaxIndex()}
	expected := args.Map{"maxIdx": 2}
	expected.ShouldBeEqual(t, 0, "MaxIndex returns 2 -- three items", actual)
}

func Test_Cov10_BasicString_NameWithIndexMap(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	m := bs.NameWithIndexMap()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NameWithIndexMap returns map -- two items", actual)
}

func Test_Cov10_BasicString_ToEnumJsonBytes_Found(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	b, err := bs.ToEnumJsonBytes("A")
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes found -- valid value", actual)
}

func Test_Cov10_BasicString_ToEnumJsonBytes_NotFound(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	_, err := bs.ToEnumJsonBytes("UNKNOWN")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes error -- unknown value", actual)
}

func Test_Cov10_BasicString_UnmarshallToValue_NilNotMapped(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	_, err := bs.UnmarshallToValue(false, nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue error -- nil not mapped", actual)
}

func Test_Cov10_BasicString_UnmarshallToValue_NilMapped(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	val, err := bs.UnmarshallToValue(true, nil)
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": "A", "noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped", actual)
}

func Test_Cov10_BasicString_UnmarshallToValue_EmptyMapped(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	val, err := bs.UnmarshallToValue(true, []byte(`""`))
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": "A", "noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty mapped", actual)
}

func Test_Cov10_BasicString_EnumType(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A"})
	actual := args.Map{"isString": bs.EnumType() != 0}
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns non-zero -- String type", actual)
}

func Test_Cov10_BasicString_IsValidRange(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	actual := args.Map{
		"valid":   bs.IsValidRange("A"),
		"invalid": bs.IsValidRange("UNKNOWN"),
	}
	expected := args.Map{
		"valid":   true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValidRange correct -- valid and invalid", actual)
}

func Test_Cov10_BasicString_AppendPrependJoinValue(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B"})
	result := bs.AppendPrependJoinValue(".", "B", "A")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinValue non-empty -- valid values", actual)
}

func Test_Cov10_BasicString_OnlySupportedErr(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})
	err := bs.OnlySupportedErr("A")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr error -- B and C unsupported", actual)
}

func Test_Cov10_BasicString_OnlySupportedMsgErr(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestStr", []string{"A", "B", "C"})
	err := bs.OnlySupportedMsgErr("custom msg", "A")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr error -- with custom message", actual)
}

// ── numberEnumBase — additional coverage ──

func Test_Cov10_NumberEnumBase_NameWithValueOption_WithQuotation(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	// Access via BasicByte which embeds numberEnumBase
	result := bb.NameWithValueOption(byte(0), true)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValueOption with quotation -- non-empty", actual)
}

func Test_Cov10_NumberEnumBase_NameWithValueOption_NoQuotation(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.NameWithValueOption(byte(0), false)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValueOption without quotation -- non-empty", actual)
}

func Test_Cov10_NumberEnumBase_Format(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.Format("Enum {type-name} - {name} - {value}", byte(0))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Format compiles template -- byte enum", actual)
}

func Test_Cov10_NumberEnumBase_OnlySupportedMsgErr(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B", "C"})
	err := bb.OnlySupportedMsgErr("custom", "A")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr error -- with message", actual)
}

func Test_Cov10_NumberEnumBase_RangesMap(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.RangesMap()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesMap returns map -- two items", actual)
}

func Test_Cov10_NumberEnumBase_NamesHashset(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.NamesHashset()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NamesHashset returns hashset -- two items", actual)
}

func Test_Cov10_NumberEnumBase_JsonString(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.JsonString(byte(0))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns non-empty -- byte value", actual)
}

func Test_Cov10_NumberEnumBase_KeyValIntegers(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.KeyValIntegers()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValIntegers returns slice -- two items", actual)
}

func Test_Cov10_NumberEnumBase_Loop(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B", "C"})
	count := 0
	bb.Loop(func(index int, name string, anyVal any) (isBreak bool) {
		count++
		return index == 1 // break after second
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Loop breaks early -- after second item", actual)
}

func Test_Cov10_NumberEnumBase_LoopInteger(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B", "C"})
	count := 0
	bb.LoopInteger(func(index int, name string, anyVal int) (isBreak bool) {
		count++
		return index == 0 // break after first
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LoopInteger breaks early -- after first item", actual)
}

func Test_Cov10_NumberEnumBase_RangesIntegerStringMap(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.RangesIntegerStringMap()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesIntegerStringMap returns map -- two items", actual)
}

// ── ConvEnumAnyValToInteger — type switch branches ──

func Test_Cov10_ConvEnumAnyValToInteger_String(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger("hello")
	actual := args.Map{"isMinInt": result < 0}
	expected := args.Map{"isMinInt": true}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns MinInt -- string input", actual)
}

func Test_Cov10_ConvEnumAnyValToInteger_Int(t *testing.T) {
	result := enumimpl.ConvEnumAnyValToInteger(42)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ConvEnumAnyValToInteger returns 42 -- int input", actual)
}

// ── BasicByte — AppendPrependJoinNamer ──

type mockNamer struct{ name string }

func (m mockNamer) Name() string { return m.name }

func Test_Cov10_BasicByte_AppendPrependJoinNamer(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.AppendPrependJoinNamer(".", mockNamer{"B"}, mockNamer{"A"})
	actual := args.Map{"result": result}
	expected := args.Map{"result": "A.B"}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinNamer joins names -- dot separator", actual)
}

// ── BasicByte — AsBasicByter ──

func Test_Cov10_BasicByte_AsBasicByter(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	byter := bb.AsBasicByter()
	actual := args.Map{"notNil": byter != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AsBasicByter returns non-nil -- valid enum", actual)
}

// ── BasicByte — ToNumberString ──

func Test_Cov10_BasicByte_ToNumberString(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("TestByte", []string{"A", "B"})
	result := bb.ToNumberString(byte(1))
	actual := args.Map{"result": result}
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "ToNumberString returns string -- byte 1", actual)
}

// ── DynamicMap — BasicByte / BasicString creation ──

func Test_Cov10_DynamicMap_BasicByte(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": byte(0), "B": byte(1)}
	bb := dm.BasicByte("TestDM")
	actual := args.Map{"notNil": bb != nil, "typeName": bb.TypeName()}
	expected := args.Map{"notNil": true, "typeName": "TestDM"}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicByte creates enum -- two items", actual)
}

func Test_Cov10_DynamicMap_BasicString(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": "alpha", "B": "beta"}
	bs := dm.BasicString("TestDM")
	actual := args.Map{"notNil": bs != nil, "typeName": bs.TypeName()}
	expected := args.Map{"notNil": true, "typeName": "TestDM"}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicString creates enum -- two items", actual)
}

func Test_Cov10_DynamicMap_BasicInt8(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicInt8("TestDM")
	actual := args.Map{"notNil": bi != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicInt8 creates enum -- two items", actual)
}

func Test_Cov10_DynamicMap_BasicInt16(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicInt16("TestDM")
	actual := args.Map{"notNil": bi != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicInt16 creates enum -- two items", actual)
}

func Test_Cov10_DynamicMap_BasicInt32(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicInt32("TestDM")
	actual := args.Map{"notNil": bi != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicInt32 creates enum -- two items", actual)
}

func Test_Cov10_DynamicMap_BasicUInt16(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": 0, "B": 1}
	bi := dm.BasicUInt16("TestDM")
	actual := args.Map{"notNil": bi != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap.BasicUInt16 creates enum -- two items", actual)
}
