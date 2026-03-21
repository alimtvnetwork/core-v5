package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// KeyValuePair
// ═══════════════════════════════════════════════════════════════

func Test_Cov50_KeyValuePair_KeyName(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP KeyName", Expected: "k", Actual: kv.KeyName(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_VariableName(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP VariableName", Expected: "k", Actual: kv.VariableName(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueString(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP ValueString", Expected: "v", Actual: kv.ValueString(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP IsVarNameEqual", Expected: true, Actual: kv.IsVariableNameEqual("k"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsValueEqual(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP IsValueEqual", Expected: true, Actual: kv.IsValueEqual("v"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsKeyEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP IsKeyEmpty", Expected: true, Actual: kv.IsKeyEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsValueEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: ""}
	tc := coretestcases.CaseV1{Name: "KVP IsValueEmpty", Expected: true, Actual: kv.IsValueEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_HasKey(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP HasKey", Expected: true, Actual: kv.HasKey(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_HasValue(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP HasValue", Expected: true, Actual: kv.HasValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsKeyValueEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{}
	tc := coretestcases.CaseV1{Name: "KVP IsKeyValueEmpty", Expected: true, Actual: kv.IsKeyValueEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_TrimKey(t *testing.T) {
	kv := corestr.KeyValuePair{Key: " k ", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP TrimKey", Expected: "k", Actual: kv.TrimKey(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_TrimValue(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: " v "}
	tc := coretestcases.CaseV1{Name: "KVP TrimValue", Expected: "v", Actual: kv.TrimValue(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueBool_True(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "true"}
	tc := coretestcases.CaseV1{Name: "KVP ValueBool true", Expected: true, Actual: kv.ValueBool(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueBool_Empty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: ""}
	tc := coretestcases.CaseV1{Name: "KVP ValueBool empty", Expected: false, Actual: kv.ValueBool(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueBool_Invalid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "notbool"}
	tc := coretestcases.CaseV1{Name: "KVP ValueBool invalid", Expected: false, Actual: kv.ValueBool(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueInt(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "42"}
	tc := coretestcases.CaseV1{Name: "KVP ValueInt", Expected: 42, Actual: kv.ValueInt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueInt_Invalid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
	tc := coretestcases.CaseV1{Name: "KVP ValueInt invalid", Expected: 99, Actual: kv.ValueInt(99), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueDefInt(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "10"}
	tc := coretestcases.CaseV1{Name: "KVP ValueDefInt", Expected: 10, Actual: kv.ValueDefInt(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueDefInt_Invalid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
	tc := coretestcases.CaseV1{Name: "KVP ValueDefInt invalid", Expected: 0, Actual: kv.ValueDefInt(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueByte(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "65"}
	tc := coretestcases.CaseV1{Name: "KVP ValueByte", Expected: byte(65), Actual: kv.ValueByte(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueByte_Invalid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
	tc := coretestcases.CaseV1{Name: "KVP ValueByte invalid", Expected: byte(5), Actual: kv.ValueByte(5), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueDefByte(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "10"}
	tc := coretestcases.CaseV1{Name: "KVP ValueDefByte", Expected: byte(10), Actual: kv.ValueDefByte(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueFloat64(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
	tc := coretestcases.CaseV1{Name: "KVP ValueFloat64", Expected: 3.14, Actual: kv.ValueFloat64(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueFloat64_Invalid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "abc"}
	tc := coretestcases.CaseV1{Name: "KVP ValueFloat64 invalid", Expected: 1.5, Actual: kv.ValueFloat64(1.5), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueDefFloat64(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "2.5"}
	tc := coretestcases.CaseV1{Name: "KVP ValueDefFloat64", Expected: 2.5, Actual: kv.ValueDefFloat64(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueValid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	vv := kv.ValueValid()
	tc := coretestcases.CaseV1{Name: "KVP ValueValid", Expected: true, Actual: vv.IsValid, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_ValueValidOptions(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	vv := kv.ValueValidOptions(false, "msg")
	tc := coretestcases.CaseV1{Name: "KVP ValueValidOptions", Expected: false, Actual: vv.IsValid, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_Is(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP Is", Expected: true, Actual: kv.Is("k", "v"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsKey(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP IsKey", Expected: true, Actual: kv.IsKey("k"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsVal(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP IsVal", Expected: true, Actual: kv.IsVal("v"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsKeyValueAnyEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: ""}
	tc := coretestcases.CaseV1{Name: "KVP IsKeyValueAnyEmpty", Expected: true, Actual: kv.IsKeyValueAnyEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_IsKeyValueAnyEmpty_Nil(t *testing.T) {
	var kv *corestr.KeyValuePair
	tc := coretestcases.CaseV1{Name: "KVP IsKeyValueAnyEmpty nil", Expected: true, Actual: kv.IsKeyValueAnyEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_FormatString(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP FormatString", Expected: "k=v", Actual: kv.FormatString("%v=%v"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_String(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP String", Expected: true, Actual: len(kv.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_Serialize(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	data, err := kv.Serialize()
	tc := coretestcases.CaseV1{Name: "KVP Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_SerializeMust(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	data := kv.SerializeMust()
	tc := coretestcases.CaseV1{Name: "KVP SerializeMust", Expected: true, Actual: len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_Compile(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	tc := coretestcases.CaseV1{Name: "KVP Compile", Expected: true, Actual: len(kv.Compile()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_Clear(t *testing.T) {
	kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
	kv.Clear()
	tc := coretestcases.CaseV1{Name: "KVP Clear", Expected: "", Actual: kv.Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_Clear_Nil(t *testing.T) {
	var kv *corestr.KeyValuePair
	kv.Clear()
	tc := coretestcases.CaseV1{Name: "KVP Clear nil", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_Dispose(t *testing.T) {
	kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
	kv.Dispose()
	tc := coretestcases.CaseV1{Name: "KVP Dispose", Expected: "", Actual: kv.Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValuePair_Dispose_Nil(t *testing.T) {
	var kv *corestr.KeyValuePair
	kv.Dispose()
	tc := coretestcases.CaseV1{Name: "KVP Dispose nil", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// KeyValueCollection
// ═══════════════════════════════════════════════════════════════

func Test_Cov50_KeyValueCollection_Add(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k1", "v1")
	tc := coretestcases.CaseV1{Name: "KVC Add", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddIf_True(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddIf(true, "k", "v")
	tc := coretestcases.CaseV1{Name: "KVC AddIf true", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddIf_False(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddIf(false, "k", "v")
	tc := coretestcases.CaseV1{Name: "KVC AddIf false", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_IsEmpty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	tc := coretestcases.CaseV1{Name: "KVC IsEmpty", Expected: true, Actual: kvc.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_HasAnyItem(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{Name: "KVC HasAnyItem", Expected: true, Actual: kvc.HasAnyItem(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Count(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{Name: "KVC Count", Expected: 1, Actual: kvc.Count(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_LastIndex(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	kvc.Add("b", "2")
	tc := coretestcases.CaseV1{Name: "KVC LastIndex", Expected: 1, Actual: kvc.LastIndex(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_HasIndex(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	tc := coretestcases.CaseV1{Name: "KVC HasIndex", Expected: true, Actual: kvc.HasIndex(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_First(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	tc := coretestcases.CaseV1{Name: "KVC First", Expected: "a", Actual: kvc.First().Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_FirstOrDefault_Has(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	tc := coretestcases.CaseV1{Name: "KVC FirstOrDefault", Expected: "a", Actual: kvc.FirstOrDefault().Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_FirstOrDefault_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	tc := coretestcases.CaseV1{Name: "KVC FirstOrDefault empty", Expected: true, Actual: kvc.FirstOrDefault() == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Last(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	kvc.Add("b", "2")
	tc := coretestcases.CaseV1{Name: "KVC Last", Expected: "b", Actual: kvc.Last().Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_LastOrDefault_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	tc := coretestcases.CaseV1{Name: "KVC LastOrDefault empty", Expected: true, Actual: kvc.LastOrDefault() == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_HasKey(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	tc := coretestcases.CaseV1{Name: "KVC HasKey", Expected: true, Actual: kvc.HasKey("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_IsContains(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	tc := coretestcases.CaseV1{Name: "KVC IsContains", Expected: true, Actual: kvc.IsContains("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Get_Found(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	val, found := kvc.Get("a")
	tc := coretestcases.CaseV1{Name: "KVC Get found", Expected: "1", Actual: val, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "KVC Get found bool", Expected: true, Actual: found, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Get_NotFound(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	_, found := kvc.Get("z")
	tc := coretestcases.CaseV1{Name: "KVC Get not found", Expected: false, Actual: found, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AllKeys(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	kvc.Add("b", "2")
	tc := coretestcases.CaseV1{Name: "KVC AllKeys", Expected: 2, Actual: len(kvc.AllKeys()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AllValues(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	tc := coretestcases.CaseV1{Name: "KVC AllValues", Expected: 1, Actual: len(kvc.AllValues()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AllKeysSorted(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("b", "2")
	kvc.Add("a", "1")
	keys := kvc.AllKeysSorted()
	tc := coretestcases.CaseV1{Name: "KVC AllKeysSorted", Expected: "a", Actual: keys[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Adds(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Adds(corestr.KeyValuePair{Key: "a", Value: "1"}, corestr.KeyValuePair{Key: "b", Value: "2"})
	tc := coretestcases.CaseV1{Name: "KVC Adds", Expected: 2, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Adds_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Adds()
	tc := coretestcases.CaseV1{Name: "KVC Adds empty", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddMap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddMap(map[string]string{"a": "1"})
	tc := coretestcases.CaseV1{Name: "KVC AddMap", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddMap_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddMap(nil)
	tc := coretestcases.CaseV1{Name: "KVC AddMap nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddHashsetMap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddHashsetMap(map[string]bool{"a": true})
	tc := coretestcases.CaseV1{Name: "KVC AddHashsetMap", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddHashsetMap_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddHashsetMap(nil)
	tc := coretestcases.CaseV1{Name: "KVC AddHashsetMap nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddHashset(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddHashset(corestr.New.Hashset.StringsSpreadItems("a"))
	tc := coretestcases.CaseV1{Name: "KVC AddHashset", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddHashset_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddHashset(nil)
	tc := coretestcases.CaseV1{Name: "KVC AddHashset nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddsHashmap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	hm := corestr.New.Hashmap.Cap(2)
	hm.Add("k", "v")
	kvc.AddsHashmap(hm)
	tc := coretestcases.CaseV1{Name: "KVC AddsHashmap", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddsHashmap_Nil(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddsHashmap(nil)
	tc := coretestcases.CaseV1{Name: "KVC AddsHashmap nil", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Hashmap(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	hm := kvc.Hashmap()
	tc := coretestcases.CaseV1{Name: "KVC Hashmap", Expected: true, Actual: hm.Has("k"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Map(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	m := kvc.Map()
	tc := coretestcases.CaseV1{Name: "KVC Map", Expected: "v", Actual: m["k"], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Join(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{Name: "KVC Join", Expected: true, Actual: len(kvc.Join(",")) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_JoinKeys(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	kvc.Add("b", "2")
	tc := coretestcases.CaseV1{Name: "KVC JoinKeys", Expected: "a,b", Actual: kvc.JoinKeys(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_JoinValues(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	tc := coretestcases.CaseV1{Name: "KVC JoinValues", Expected: "1", Actual: kvc.JoinValues(","), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Strings(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{Name: "KVC Strings", Expected: 1, Actual: len(kvc.Strings()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Strings_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	tc := coretestcases.CaseV1{Name: "KVC Strings empty", Expected: 0, Actual: len(kvc.Strings()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_StringsUsingFormat(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	result := kvc.StringsUsingFormat("%v=%v")
	tc := coretestcases.CaseV1{Name: "KVC StringsUsingFormat", Expected: "k=v", Actual: result[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_String(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{Name: "KVC String", Expected: true, Actual: len(kvc.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_SafeValueAt(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	tc := coretestcases.CaseV1{Name: "KVC SafeValueAt", Expected: "v", Actual: kvc.SafeValueAt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_SafeValueAt_OOB(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	tc := coretestcases.CaseV1{Name: "KVC SafeValueAt oob", Expected: "", Actual: kvc.SafeValueAt(0), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_SafeValuesAtIndexes(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	kvc.Add("b", "2")
	result := kvc.SafeValuesAtIndexes(0, 1)
	tc := coretestcases.CaseV1{Name: "KVC SafeValuesAtIndexes", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Serialize(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	data, err := kvc.Serialize()
	tc := coretestcases.CaseV1{Name: "KVC Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_SerializeMust(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("k", "v")
	data := kvc.SerializeMust()
	tc := coretestcases.CaseV1{Name: "KVC SerializeMust", Expected: true, Actual: len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddStringBySplit(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddStringBySplit("=", "key=value")
	tc := coretestcases.CaseV1{Name: "KVC AddStringBySplit", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddStringBySplitTrim(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddStringBySplitTrim("=", " key = value ")
	tc := coretestcases.CaseV1{Name: "KVC AddStringBySplitTrim", Expected: "key", Actual: kvc.First().Key, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Find(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.Add("a", "1")
	kvc.Add("b", "2")
	result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
		return kv, kv.Key == "a", false
	})
	tc := coretestcases.CaseV1{Name: "KVC Find", Expected: 1, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_Find_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
		return kv, true, false
	})
	tc := coretestcases.CaseV1{Name: "KVC Find empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddsHashmaps(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	hm := corestr.New.Hashmap.Cap(2)
	hm.Add("k", "v")
	kvc.AddsHashmaps(hm)
	tc := coretestcases.CaseV1{Name: "KVC AddsHashmaps", Expected: 1, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov50_KeyValueCollection_AddsHashmaps_Empty(t *testing.T) {
	kvc := &corestr.KeyValueCollection{}
	kvc.AddsHashmaps()
	tc := coretestcases.CaseV1{Name: "KVC AddsHashmaps empty", Expected: 0, Actual: kvc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
