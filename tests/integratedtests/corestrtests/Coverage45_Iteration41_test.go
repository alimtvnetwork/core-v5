package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// ValidValues — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_ValidValues_Add(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("hello")
	tc := coretestcases.CaseV1{Name: "Add", Expected: 1, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddFull(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddFull(false, "val", "msg")
	tc := coretestcases.CaseV1{Name: "AddFull", Expected: 1, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValueAt_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("hello")
	tc := coretestcases.CaseV1{Name: "SafeValueAt valid", Expected: "hello", Actual: vv.SafeValueAt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValueAt_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := coretestcases.CaseV1{Name: "SafeValueAt empty", Expected: "", Actual: vv.SafeValueAt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValueAt_OutOfRange(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("x")
	tc := coretestcases.CaseV1{Name: "SafeValueAt oob", Expected: "", Actual: vv.SafeValueAt(5), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValueAt_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("hello") // IsValid=true
	tc := coretestcases.CaseV1{Name: "SafeValidValueAt valid", Expected: "hello", Actual: vv.SafeValidValueAt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValueAt_Invalid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddFull(false, "val", "msg") // IsValid=false
	tc := coretestcases.CaseV1{Name: "SafeValidValueAt invalid", Expected: "", Actual: vv.SafeValidValueAt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValueAt_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := coretestcases.CaseV1{Name: "SafeValidValueAt empty", Expected: "", Actual: vv.SafeValidValueAt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	result := vv.SafeValuesAtIndexes(0, 1)
	tc := coretestcases.CaseV1{Name: "SafeValuesAtIndexes", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValuesAtIndexes_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	result := vv.SafeValuesAtIndexes()
	tc := coretestcases.CaseV1{Name: "SafeValuesAtIndexes empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.AddFull(false, "b", "msg")
	result := vv.SafeValidValuesAtIndexes(0, 1)
	tc := coretestcases.CaseV1{Name: "SafeValidValuesAtIndexes", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Strings(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	result := vv.Strings()
	tc := coretestcases.CaseV1{Name: "Strings", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Strings_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	result := vv.Strings()
	tc := coretestcases.CaseV1{Name: "Strings empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_FullStrings(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.FullStrings()
	tc := coretestcases.CaseV1{Name: "FullStrings", Expected: 1, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_FullStrings_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	result := vv.FullStrings()
	tc := coretestcases.CaseV1{Name: "FullStrings empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_String(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.String()
	tc := coretestcases.CaseV1{Name: "String", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Length_Nil(t *testing.T) {
	var vv *corestr.ValidValues
	tc := coretestcases.CaseV1{Name: "Length nil", Expected: 0, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_IsEmpty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	tc := coretestcases.CaseV1{Name: "IsEmpty", Expected: true, Actual: vv.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Adds(t *testing.T) {
	vv := corestr.EmptyValidValues()
	v1 := corestr.ValidValue{Value: "a", IsValid: true}
	v2 := corestr.ValidValue{Value: "b", IsValid: true}
	vv.Adds(v1, v2)
	tc := coretestcases.CaseV1{Name: "Adds", Expected: 2, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Adds_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Adds()
	tc := coretestcases.CaseV1{Name: "Adds empty", Expected: 0, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddsPtr(t *testing.T) {
	vv := corestr.EmptyValidValues()
	v1 := &corestr.ValidValue{Value: "a", IsValid: true}
	vv.AddsPtr(v1)
	tc := coretestcases.CaseV1{Name: "AddsPtr", Expected: 1, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddsPtr_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddsPtr()
	tc := coretestcases.CaseV1{Name: "AddsPtr empty", Expected: 0, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddValidValues(t *testing.T) {
	vv1 := corestr.EmptyValidValues()
	vv1.Add("a")
	vv2 := corestr.EmptyValidValues()
	vv2.AddValidValues(vv1)
	tc := coretestcases.CaseV1{Name: "AddValidValues", Expected: 1, Actual: vv2.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddValidValues_Nil(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddValidValues(nil)
	tc := coretestcases.CaseV1{Name: "AddValidValues nil", Expected: 0, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.ConcatNew(true)
	tc := coretestcases.CaseV1{Name: "ConcatNew empty clone", Expected: 1, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	result := vv.ConcatNew(false)
	tc := coretestcases.CaseV1{Name: "ConcatNew empty no clone", Expected: 1, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_ConcatNew_WithArgs(t *testing.T) {
	vv1 := corestr.EmptyValidValues()
	vv1.Add("a")
	vv2 := corestr.EmptyValidValues()
	vv2.Add("b")
	result := vv1.ConcatNew(true, vv2)
	tc := coretestcases.CaseV1{Name: "ConcatNew with args", Expected: 2, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashsetMap_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	m := map[string]bool{"a": true, "b": false}
	vv.AddHashsetMap(m)
	tc := coretestcases.CaseV1{Name: "AddHashsetMap", Expected: 2, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddHashsetMap(nil)
	tc := coretestcases.CaseV1{Name: "AddHashsetMap nil", Expected: 0, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashset_Valid(t *testing.T) {
	vv := corestr.EmptyValidValues()
	hs := corestr.New.Hashset.StringsSpreadItems("x", "y")
	vv.AddHashset(hs)
	tc := coretestcases.CaseV1{Name: "AddHashset", Expected: 2, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_AddHashset_Nil(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.AddHashset(nil)
	tc := coretestcases.CaseV1{Name: "AddHashset nil", Expected: 0, Actual: vv.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Hashmap(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("key")
	hm := vv.Hashmap()
	tc := coretestcases.CaseV1{Name: "Hashmap", Expected: true, Actual: hm.Has("key"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Hashmap_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	hm := vv.Hashmap()
	tc := coretestcases.CaseV1{Name: "Hashmap empty", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Map(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("key")
	m := vv.Map()
	tc := coretestcases.CaseV1{Name: "Map", Expected: true, Actual: len(m) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Find_Found(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	found := vv.Find(func(index int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, v.Value == "a", false
	})
	tc := coretestcases.CaseV1{Name: "Find found", Expected: 1, Actual: len(found), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Find_Break(t *testing.T) {
	vv := corestr.EmptyValidValues()
	vv.Add("a")
	vv.Add("b")
	found := vv.Find(func(index int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, true, true // break on first
	})
	tc := coretestcases.CaseV1{Name: "Find break", Expected: 1, Actual: len(found), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_ValidValues_Find_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	found := vv.Find(func(index int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, true, false
	})
	tc := coretestcases.CaseV1{Name: "Find empty", Expected: 0, Actual: len(found), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// HashmapDiff
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_HashmapDiff_Length(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := coretestcases.CaseV1{Name: "HashmapDiff Length", Expected: 1, Actual: hd.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Length_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	tc := coretestcases.CaseV1{Name: "HashmapDiff Length nil", Expected: 0, Actual: hd.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_IsEmpty(t *testing.T) {
	hd := corestr.HashmapDiff{}
	tc := coretestcases.CaseV1{Name: "HashmapDiff IsEmpty", Expected: true, Actual: hd.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HasAnyItem(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := coretestcases.CaseV1{Name: "HashmapDiff HasAnyItem", Expected: true, Actual: hd.HasAnyItem(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_LastIndex(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1", "b": "2"}
	tc := coretestcases.CaseV1{Name: "HashmapDiff LastIndex", Expected: 1, Actual: hd.LastIndex(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_AllKeysSorted(t *testing.T) {
	hd := corestr.HashmapDiff{"b": "2", "a": "1"}
	keys := hd.AllKeysSorted()
	tc := coretestcases.CaseV1{Name: "AllKeysSorted first", Expected: "a", Actual: keys[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_MapAnyItems(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	m := hd.MapAnyItems()
	tc := coretestcases.CaseV1{Name: "MapAnyItems", Expected: "1", Actual: m["a"], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_MapAnyItems_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	m := hd.MapAnyItems()
	tc := coretestcases.CaseV1{Name: "MapAnyItems nil", Expected: 0, Actual: len(m), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Raw(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	raw := hd.Raw()
	tc := coretestcases.CaseV1{Name: "Raw", Expected: "1", Actual: raw["a"], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Raw_Nil(t *testing.T) {
	var hd *corestr.HashmapDiff
	raw := hd.Raw()
	tc := coretestcases.CaseV1{Name: "Raw nil", Expected: 0, Actual: len(raw), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_IsRawEqual_Same(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := coretestcases.CaseV1{Name: "IsRawEqual same", Expected: true, Actual: hd.IsRawEqual(map[string]string{"a": "1"}), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_IsRawEqual_Diff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := coretestcases.CaseV1{Name: "IsRawEqual diff", Expected: false, Actual: hd.IsRawEqual(map[string]string{"a": "2"}), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HasAnyChanges(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	tc := coretestcases.CaseV1{Name: "HasAnyChanges", Expected: true, Actual: hd.HasAnyChanges(map[string]string{"a": "2"}), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "HashmapDiffUsingRaw no diff", Expected: 0, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_HashmapDiffUsingRaw_HasDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.HashmapDiffUsingRaw(map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "HashmapDiffUsingRaw has diff", Expected: true, Actual: result.HasAnyItem(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_DiffRaw(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.DiffRaw(map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "DiffRaw", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_DiffJsonMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.DiffJsonMessage(map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "DiffJsonMessage", Expected: true, Actual: len(msg) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.ShouldDiffMessage("test", map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "ShouldDiffMessage", Expected: true, Actual: len(msg) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_LogShouldDiffMessage(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	msg := hd.LogShouldDiffMessage("test", map[string]string{"a": "2"})
	tc := coretestcases.CaseV1{Name: "LogShouldDiffMessage", Expected: true, Actual: len(msg) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	diffMap := map[string]string{"a": "changed"}
	result := hd.ToStringsSliceOfDiffMap(diffMap)
	tc := coretestcases.CaseV1{Name: "ToStringsSliceOfDiffMap", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	result := hd.RawMapStringAnyDiff()
	tc := coretestcases.CaseV1{Name: "RawMapStringAnyDiff", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_HashmapDiff_Serialize(t *testing.T) {
	hd := corestr.HashmapDiff{"a": "1"}
	data, err := hd.Serialize()
	tc := coretestcases.CaseV1{Name: "Serialize no err", Expected: true, Actual: err == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Serialize has data", Expected: true, Actual: len(data) > 0, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// LeftRightFromSplit factories
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_LeftRightFromSplit(t *testing.T) {
	lr := corestr.LeftRightFromSplit("key=value", "=")
	tc := coretestcases.CaseV1{Name: "LeftRightFromSplit left", Expected: "key", Actual: lr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "LeftRightFromSplit right", Expected: "value", Actual: lr.Right, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftRightFromSplitTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
	tc := coretestcases.CaseV1{Name: "LeftRightFromSplitTrimmed left", Expected: "key", Actual: lr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "LeftRightFromSplitTrimmed right", Expected: "value", Actual: lr.Right, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftRightFromSplitFull(t *testing.T) {
	lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
	tc := coretestcases.CaseV1{Name: "LeftRightFromSplitFull left", Expected: "a", Actual: lr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "LeftRightFromSplitFull right", Expected: "b:c:d", Actual: lr.Right, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftRightFromSplitFullTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
	tc := coretestcases.CaseV1{Name: "LeftRightFromSplitFullTrimmed left", Expected: "a", Actual: lr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "LeftRightFromSplitFullTrimmed right trimmed", Expected: true, Actual: len(lr.Right) > 0, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit factories
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_LeftMiddleRightFromSplit(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
	tc := coretestcases.CaseV1{Name: "LMR left", Expected: "a", Actual: lmr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "LMR middle", Expected: "b", Actual: lmr.Middle, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
	tc3 := coretestcases.CaseV1{Name: "LMR right", Expected: "c", Actual: lmr.Right, Args: args.Map{}}
	tc3.ShouldBeEqual(t)
}

func Test_Cov45_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
	tc := coretestcases.CaseV1{Name: "LMR trimmed left", Expected: "a", Actual: lmr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "LMR trimmed middle", Expected: "b", Actual: lmr.Middle, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_LeftMiddleRightFromSplitN(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
	tc := coretestcases.CaseV1{Name: "LMR SplitN left", Expected: "a", Actual: lmr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "LMR SplitN middle", Expected: "b", Actual: lmr.Middle, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
	tc3 := coretestcases.CaseV1{Name: "LMR SplitN right remainder", Expected: "c:d:e", Actual: lmr.Right, Args: args.Map{}}
	tc3.ShouldBeEqual(t)
}

func Test_Cov45_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
	tc := coretestcases.CaseV1{Name: "LMR SplitNTrimmed left", Expected: "a", Actual: lmr.Left, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// KeyAnyValuePair — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov45_KeyAnyValuePair_KeyName(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "mykey", Value: "val"}
	tc := coretestcases.CaseV1{Name: "KeyName", Expected: "mykey", Actual: kv.KeyName(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_VariableName(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "var1", Value: 42}
	tc := coretestcases.CaseV1{Name: "VariableName", Expected: "var1", Actual: kv.VariableName(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_ValueAny(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: 99}
	tc := coretestcases.CaseV1{Name: "ValueAny", Expected: 99, Actual: kv.ValueAny(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsVariableNameEqual(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k1", Value: "v"}
	tc := coretestcases.CaseV1{Name: "IsVariableNameEqual", Expected: true, Actual: kv.IsVariableNameEqual("k1"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_SerializeMust(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	data := kv.SerializeMust()
	tc := coretestcases.CaseV1{Name: "SerializeMust", Expected: true, Actual: len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Compile(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "Compile", Expected: true, Actual: len(kv.Compile()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueNull_Nil(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := coretestcases.CaseV1{Name: "IsValueNull nil", Expected: true, Actual: kv.IsValueNull(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueNull_NilReceiver(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	tc := coretestcases.CaseV1{Name: "IsValueNull nil receiver", Expected: true, Actual: kv.IsValueNull(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_HasNonNull(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "HasNonNull", Expected: true, Actual: kv.HasNonNull(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_HasNonNull_Nil(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	tc := coretestcases.CaseV1{Name: "HasNonNull nil", Expected: false, Actual: kv.HasNonNull(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_HasValue(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: 1}
	tc := coretestcases.CaseV1{Name: "HasValue", Expected: true, Actual: kv.HasValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueEmptyString(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := coretestcases.CaseV1{Name: "IsValueEmptyString nil value", Expected: true, Actual: kv.IsValueEmptyString(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_IsValueWhitespace(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	tc := coretestcases.CaseV1{Name: "IsValueWhitespace nil", Expected: true, Actual: kv.IsValueWhitespace(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_ValueString_Cached(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
	v1 := kv.ValueString()
	v2 := kv.ValueString() // cached
	tc := coretestcases.CaseV1{Name: "ValueString cached", Expected: v1, Actual: v2, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_ValueString_NullValue(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: nil}
	result := kv.ValueString()
	tc := coretestcases.CaseV1{Name: "ValueString null", Expected: "", Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_String(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "String", Expected: true, Actual: len(kv.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Json(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	j := kv.Json()
	tc := coretestcases.CaseV1{Name: "Json", Expected: true, Actual: j.HasSafeNonEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_JsonPtr(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	j := kv.JsonPtr()
	tc := coretestcases.CaseV1{Name: "JsonPtr", Expected: true, Actual: j != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Serialize(t *testing.T) {
	kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	data, err := kv.Serialize()
	tc := coretestcases.CaseV1{Name: "Serialize no err", Expected: true, Actual: err == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Serialize data", Expected: true, Actual: len(data) > 0, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_AsJsonContractsBinder(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "AsJsonContractsBinder", Expected: true, Actual: kv.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_AsJsoner(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "AsJsoner", Expected: true, Actual: kv.AsJsoner() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_AsJsonParseSelfInjector(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "AsJsonParseSelfInjector", Expected: true, Actual: kv.AsJsonParseSelfInjector() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Clear(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	kv.Clear()
	tc := coretestcases.CaseV1{Name: "Clear key", Expected: "", Actual: kv.Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "Clear value nil", Expected: true, Actual: kv.Value == nil, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Clear_Nil(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	kv.Clear() // should not panic
	tc := coretestcases.CaseV1{Name: "Clear nil no panic", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Dispose(t *testing.T) {
	kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
	kv.Dispose()
	tc := coretestcases.CaseV1{Name: "Dispose key", Expected: "", Actual: kv.Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov45_KeyAnyValuePair_Dispose_Nil(t *testing.T) {
	var kv *corestr.KeyAnyValuePair
	kv.Dispose() // should not panic
	tc := coretestcases.CaseV1{Name: "Dispose nil no panic", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
