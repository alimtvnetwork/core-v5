package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ═══════════════════════════════════════════════════════════════
// ValidValues — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_ValidValues_Add(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("hello")
	tc := caseV1Compat{Name: "Add", Expected: 1, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddFull(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddFull(false, "val", "msg")
	tc := caseV1Compat{Name: "AddFull", Expected: 1, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValueAt_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("hello")
	tc := caseV1Compat{Name: "SafeValueAt valid", Expected: "hello", Actual: vv.SafeValueAt(0)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValueAt_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := caseV1Compat{Name: "SafeValueAt empty", Expected: "", Actual: vv.SafeValueAt(0)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValueAt_OutOfRange(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("x")
	tc := caseV1Compat{Name: "SafeValueAt oob", Expected: "", Actual: vv.SafeValueAt(5)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValueAt_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("hello") // IsValid=true
	tc := caseV1Compat{Name: "SafeValidValueAt valid", Expected: "hello", Actual: vv.SafeValidValueAt(0)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValueAt_Invalid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddFull(false, "val", "msg") // IsValid=false
	tc := caseV1Compat{Name: "SafeValidValueAt invalid", Expected: "", Actual: vv.SafeValidValueAt(0)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValueAt_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := caseV1Compat{Name: "SafeValidValueAt empty", Expected: "", Actual: vv.SafeValidValueAt(0)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	result := vv.SafeValuesAtIndexes(0, 1)
	tc := caseV1Compat{Name: "SafeValuesAtIndexes", Expected: 2, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValuesAtIndexes_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	result := vv.SafeValuesAtIndexes()
	tc := caseV1Compat{Name: "SafeValuesAtIndexes empty", Expected: 0, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.AddFull(false, "b", "msg")
	result := vv.SafeValidValuesAtIndexes(0, 1)
	tc := caseV1Compat{Name: "SafeValidValuesAtIndexes", Expected: 2, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Strings(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	result := vv.Strings()
	tc := caseV1Compat{Name: "Strings", Expected: 2, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Strings_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	result := vv.Strings()
	tc := caseV1Compat{Name: "Strings empty", Expected: 0, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_FullStrings(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.FullStrings()
	tc := caseV1Compat{Name: "FullStrings", Expected: 1, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_FullStrings_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	result := vv.FullStrings()
	tc := caseV1Compat{Name: "FullStrings empty", Expected: 0, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_String(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.String()
	tc := caseV1Compat{Name: "String", Expected: true, Actual: len(result) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Length_Nil(t *testing.T) {
	var vv *corestr.ValidValues
	tc := caseV1Compat{Name: "Length nil", Expected: 0, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_IsEmpty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := caseV1Compat{Name: "IsEmpty", Expected: true, Actual: vv.IsEmpty()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Adds(t *testing.T) {
	vv := corestr.EmptyValidValues()
	v1 := corestr.ValidValue{Value: "a", IsValid: true}
	v2 := corestr.ValidValue{Value: "b", IsValid: true}
	vv.Adds(v1, v2)
	tc := caseV1Compat{Name: "Adds", Expected: 2, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Adds_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Adds()
	tc := caseV1Compat{Name: "Adds empty", Expected: 0, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddsPtr(t *testing.T) {
	vv := corestr.EmptyValidValues()
	v1 := &corestr.ValidValue{Value: "a", IsValid: true}
	vv.AddsPtr(v1)
	tc := caseV1Compat{Name: "AddsPtr", Expected: 1, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddsPtr_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddsPtr()
	tc := caseV1Compat{Name: "AddsPtr empty", Expected: 0, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddValidValues(t *testing.T) {
	vv1 := corestr.EmptyValidValues()
	vv1.Add("a")
	vv2 := corestr.EmptyValidValues()
	vv2.AddValidValues(vv1)
	tc := caseV1Compat{Name: "AddValidValues", Expected: 1, Actual: vv2.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddValidValues_Nil(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddValidValues(nil)
	tc := caseV1Compat{Name: "AddValidValues nil", Expected: 0, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.ConcatNew(true)
	tc := caseV1Compat{Name: "ConcatNew empty clone", Expected: 1, Actual: result.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.ConcatNew(false)
	tc := caseV1Compat{Name: "ConcatNew empty no clone", Expected: 1, Actual: result.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_ConcatNew_WithArgs(t *testing.T) {
	vv1 := corestr.EmptyValidValues()
	vv1.Add("a")
	vv2 := corestr.EmptyValidValues()
	vv2.Add("b")
	result := vv1.ConcatNew(true, vv2)
	tc := caseV1Compat{Name: "ConcatNew with args", Expected: 2, Actual: result.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashsetMap_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	m := map[string]bool{"a": true, "b": false}
	vv.AddHashsetMap(m)
	tc := caseV1Compat{Name: "AddHashsetMap", Expected: 2, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddHashsetMap(nil)
	tc := caseV1Compat{Name: "AddHashsetMap nil", Expected: 0, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashset_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	hs := corestr.New.Hashset.StringsSpreadItems("x", "y")
	vv.AddHashset(hs)
	tc := caseV1Compat{Name: "AddHashset", Expected: 2, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashset_Nil(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddHashset(nil)
	tc := caseV1Compat{Name: "AddHashset nil", Expected: 0, Actual: vv.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Hashmap(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("key")
	hm := vv.Hashmap()
	tc := caseV1Compat{Name: "Hashmap", Expected: true, Actual: hm.Has("key")}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Hashmap_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	hm := vv.Hashmap()
	tc := caseV1Compat{Name: "Hashmap empty", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Map(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("key")
	m := vv.Map()
	tc := caseV1Compat{Name: "Map", Expected: true, Actual: len(m) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Find_Found(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	found := vv.Find(func(index int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, v.Value == "a", false
	})
	tc := caseV1Compat{Name: "Find found", Expected: 1, Actual: len(found)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Find_Break(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	found := vv.Find(func(index int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, true, true // break on first
	})
	tc := caseV1Compat{Name: "Find break", Expected: 1, Actual: len(found)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Find_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	found := vv.Find(func(index int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, true, false
	})
	tc := caseV1Compat{Name: "Find empty", Expected: 0, Actual: len(found)}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// HashmapDiff
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_HashmapDiff_Length(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := caseV1Compat{Name: "HashmapDiff Length", Expected: 1, Actual: hd.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Length_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	tc := caseV1Compat{Name: "HashmapDiff Length nil", Expected: 0, Actual: hd.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_IsEmpty(t *testing.T) {
	hd := corestr.HashmapDiff{}
	tc := caseV1Compat{Name: "HashmapDiff IsEmpty", Expected: true, Actual: hd.IsEmpty()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HasAnyItem(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := caseV1Compat{Name: "HashmapDiff HasAnyItem", Expected: true, Actual: hd.HasAnyItem()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_LastIndex(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1", "b": "2"}
	tc := caseV1Compat{Name: "HashmapDiff LastIndex", Expected: 1, Actual: hd.LastIndex()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_AllKeysSorted(t *testing.T) {
	hd := corestr.HashmapDiff{"b": "2", "a": "1"}
	keys := hd.AllKeysSorted()
	tc := caseV1Compat{Name: "AllKeysSorted first", Expected: "a", Actual: keys[0]}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_MapAnyItems(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	m := hd.MapAnyItems()
	tc := caseV1Compat{Name: "MapAnyItems", Expected: "1", Actual: m["a"]}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_MapAnyItems_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	m := hd.MapAnyItems()
	tc := caseV1Compat{Name: "MapAnyItems nil", Expected: 0, Actual: len(m)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Raw(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	raw := hd.Raw()
	tc := caseV1Compat{Name: "Raw", Expected: "1", Actual: raw["a"]}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Raw_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	raw := hd.Raw()
	tc := caseV1Compat{Name: "Raw nil", Expected: 0, Actual: len(raw)}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_IsRawEqual_Same(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := caseV1Compat{Name: "IsRawEqual same", Expected: true, Actual: hd.IsRawEqual(map[string]string{"a": "1"})}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_IsRawEqual_Diff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := caseV1Compat{Name: "IsRawEqual diff", Expected: false, Actual: hd.IsRawEqual(map[string]string{"a": "2"})}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HasAnyChanges(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := caseV1Compat{Name: "HasAnyChanges", Expected: true, Actual: hd.HasAnyChanges(map[string]string{"a": "2"})}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})
	tc := caseV1Compat{Name: "HashmapDiffUsingRaw no diff", Expected: 0, Actual: result.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HashmapDiffUsingRaw_HasDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.HashmapDiffUsingRaw(map[string]string{"a": "2"})
	tc := caseV1Compat{Name: "HashmapDiffUsingRaw has diff", Expected: true, Actual: result.HasAnyItem()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_DiffRaw(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.DiffRaw(map[string]string{"a": "2"})
	tc := caseV1Compat{Name: "DiffRaw", Expected: true, Actual: len(result) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_DiffJsonMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.DiffJsonMessage(map[string]string{"a": "2"})
	tc := caseV1Compat{Name: "DiffJsonMessage", Expected: true, Actual: len(msg) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})
	tc := caseV1Compat{Name: "ShouldDiffMessage", Expected: true, Actual: len(msg) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_LogShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.LogShouldDiffMessage("test", map[string]string{"a": "2"})
	tc := caseV1Compat{Name: "LogShouldDiffMessage", Expected: true, Actual: len(msg) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	diffMap := map[string]string{"a": "changed"}
	result := hd.ToStringsSliceOfDiffMap(diffMap)
	tc := caseV1Compat{Name: "ToStringsSliceOfDiffMap", Expected: true, Actual: len(result) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.RawMapStringAnyDiff()
	tc := caseV1Compat{Name: "RawMapStringAnyDiff", Expected: true, Actual: len(result) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Serialize(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	data, err := hd.Serialize()
	tc := caseV1Compat{Name: "Serialize no err", Expected: true, Actual: err == nil}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "Serialize has data", Expected: true, Actual: len(data) > 0}
	tc2.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// LeftRightFromSplit factories
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_LeftRightFromSplit(t *testing.T) {
	lr := corestr.LeftRightFromSplit("key=value", "=")
	tc := caseV1Compat{Name: "LeftRightFromSplit left", Expected: "key", Actual: lr.Left}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "LeftRightFromSplit right", Expected: "value", Actual: lr.Right}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftRightFromSplitTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
	tc := caseV1Compat{Name: "LeftRightFromSplitTrimmed left", Expected: "key", Actual: lr.Left}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "LeftRightFromSplitTrimmed right", Expected: "value", Actual: lr.Right}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftRightFromSplitFull(t *testing.T) {
	lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
	tc := caseV1Compat{Name: "LeftRightFromSplitFull left", Expected: "a", Actual: lr.Left}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "LeftRightFromSplitFull right", Expected: "b:c:d", Actual: lr.Right}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftRightFromSplitFullTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
	tc := caseV1Compat{Name: "LeftRightFromSplitFullTrimmed left", Expected: "a", Actual: lr.Left}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "LeftRightFromSplitFullTrimmed right trimmed", Expected: true, Actual: len(lr.Right) > 0}
	tc2.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit factories
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_LeftMiddleRightFromSplit(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
	tc := caseV1Compat{Name: "LMR left", Expected: "a", Actual: lmr.Left}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "LMR middle", Expected: "b", Actual: lmr.Middle}
	tc2.ShouldBeEqual(t)
	tc3 := caseV1Compat{Name: "LMR right", Expected: "c", Actual: lmr.Right}
	tc3.ShouldBeEqual(t)
}

func Test_Cov45_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
	tc := caseV1Compat{Name: "LMR trimmed left", Expected: "a", Actual: lmr.Left}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "LMR trimmed middle", Expected: "b", Actual: lmr.Middle}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftMiddleRightFromSplitN(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
	tc := caseV1Compat{Name: "LMR SplitN left", Expected: "a", Actual: lmr.Left}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "LMR SplitN middle", Expected: "b", Actual: lmr.Middle}
	tc2.ShouldBeEqual(t)
	tc3 := caseV1Compat{Name: "LMR SplitN right remainder", Expected: "c:d:e", Actual: lmr.Right}
	tc3.ShouldBeEqual(t)
}

func Test_Cov45_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
	tc := caseV1Compat{Name: "LMR SplitNTrimmed left", Expected: "a", Actual: lmr.Left}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// KeyAnyValuePair — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_KeyAnyValuePair_KeyName(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "mykey", Value: "val"}
	tc := caseV1Compat{Name: "KeyName", Expected: "mykey", Actual: kv.KeyName()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_VariableName(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "var1", Value: 42}
	tc := caseV1Compat{Name: "VariableName", Expected: "var1", Actual: kv.VariableName()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_ValueAny(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: 99}
	tc := caseV1Compat{Name: "ValueAny", Expected: 99, Actual: kv.ValueAny()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k1", Value: "v"}
	tc := caseV1Compat{Name: "IsVariableNameEqual", Expected: true, Actual: kv.IsVariableNameEqual("k1")}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_SerializeMust(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	data := kv.SerializeMust()
	tc := caseV1Compat{Name: "SerializeMust", Expected: true, Actual: len(data) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Compile(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := caseV1Compat{Name: "Compile", Expected: true, Actual: len(kv.Compile()) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueNull_Nil(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := caseV1Compat{Name: "IsValueNull nil", Expected: true, Actual: kv.IsValueNull()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueNull_NilReceiver(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	tc := caseV1Compat{Name: "IsValueNull nil receiver", Expected: true, Actual: kv.IsValueNull()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_HasNonNull(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := caseV1Compat{Name: "HasNonNull", Expected: true, Actual: kv.HasNonNull()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_HasNonNull_Nil(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	tc := caseV1Compat{Name: "HasNonNull nil", Expected: false, Actual: kv.HasNonNull()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_HasValue(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: 1}
	tc := caseV1Compat{Name: "HasValue", Expected: true, Actual: kv.HasValue()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := caseV1Compat{Name: "IsValueEmptyString nil value", Expected: true, Actual: kv.IsValueEmptyString()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := caseV1Compat{Name: "IsValueWhitespace nil", Expected: true, Actual: kv.IsValueWhitespace()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_ValueString_Cached(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
	v1 := kv.ValueString()
	v2 := kv.ValueString() // cached
	tc := caseV1Compat{Name: "ValueString cached", Expected: v1, Actual: v2}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_ValueString_NullValue(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	result := kv.ValueString()
	tc := caseV1Compat{Name: "ValueString null", Expected: "", Actual: result}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_String(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := caseV1Compat{Name: "String", Expected: true, Actual: len(kv.String()) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Json(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	j := kv.Json()
	tc := caseV1Compat{Name: "Json", Expected: true, Actual: j.HasAnyItem()}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_JsonPtr(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	j := kv.JsonPtr()
	tc := caseV1Compat{Name: "JsonPtr", Expected: true, Actual: j != nil}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Serialize(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	data, err := kv.Serialize()
	tc := caseV1Compat{Name: "Serialize no err", Expected: true, Actual: err == nil}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "Serialize data", Expected: true, Actual: len(data) > 0}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := caseV1Compat{Name: "AsJsonContractsBinder", Expected: true, Actual: kv.AsJsonContractsBinder() != nil}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_AsJsoner(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := caseV1Compat{Name: "AsJsoner", Expected: true, Actual: kv.AsJsoner() != nil}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_AsJsonParseSelfInjector(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := caseV1Compat{Name: "AsJsonParseSelfInjector", Expected: true, Actual: kv.AsJsonParseSelfInjector() != nil}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Clear(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	kv.Clear()
	tc := caseV1Compat{Name: "Clear key", Expected: "", Actual: kv.Key}
	tc.ShouldBeEqual(t)
	tc2 := caseV1Compat{Name: "Clear value nil", Expected: true, Actual: kv.Value == nil}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Clear_Nil(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	kv.Clear() // should not panic
	tc := caseV1Compat{Name: "Clear nil no panic", Expected: true, Actual: true}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Dispose(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	kv.Dispose()
	tc := caseV1Compat{Name: "Dispose key", Expected: "", Actual: kv.Key}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Dispose_Nil(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	kv.Dispose() // should not panic
	tc := caseV1Compat{Name: "Dispose nil no panic", Expected: true, Actual: true}
	tc.ShouldBeEqual(t)
}
