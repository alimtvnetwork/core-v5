package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Fix: All tests in this file were using CaseV1.ShouldBeEqual(t, caseIndex)
// without passing actual values via the variadic ...string parameter.
// ShouldBeEqual ignores the ActualInput struct field and only compares
// the variadic actualElements with ExpectedInput.
// Fixed to use ShouldBeEqualMapFirst(t, args.Map{...}) pattern.
// See issues/corestrtests-cov43-wrong-assertion-pattern.md

// ═══════════════════════════════════════════════════════════════
// newHashmapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashmapCreator_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	tc := coretestcases.CaseV1{Title: "Empty hashmap", ExpectedInput: args.Map{"value": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_Cap(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(10)
	tc := coretestcases.CaseV1{Title: "Cap hashmap empty", ExpectedInput: args.Map{"value": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Valid(t *testing.T) {
	pair := corestr.KeyAnyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyAnyValues(pair)
	tc := coretestcases.CaseV1{Title: "KeyAnyValues valid", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": hm.Has("k1")})
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyAnyValues()
	tc := coretestcases.CaseV1{Title: "KeyAnyValues empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_KeyValues_Valid(t *testing.T) {
	pair := corestr.KeyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyValues(pair)
	tc := coretestcases.CaseV1{Title: "KeyValues valid", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": hm.Has("k1")})
}

func Test_Cov43_NewHashmapCreator_KeyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues()
	tc := coretestcases.CaseV1{Title: "KeyValues empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_Valid(t *testing.T) {
	keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
	vals := corestr.New.Collection.Strings([]string{"v1", "v2"})
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := coretestcases.CaseV1{Title: "KeyValuesCollection valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_NilKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
	tc := coretestcases.CaseV1{Title: "KeyValuesCollection nil", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_EmptyKeys(t *testing.T) {
	keys := corestr.New.Collection.Empty()
	vals := corestr.New.Collection.Empty()
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := coretestcases.CaseV1{Title: "KeyValuesCollection empty keys", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_Valid(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{"a", "b"}, []string{"1", "2"})
	tc := coretestcases.CaseV1{Title: "KeyValuesStrings valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_EmptyKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{}, []string{})
	tc := coretestcases.CaseV1{Title: "KeyValuesStrings empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_UsingMap(t *testing.T) {
	m := map[string]string{"x": "y"}
	hm := corestr.New.Hashmap.UsingMap(m)
	val, _ := hm.Get("x")
	tc := coretestcases.CaseV1{Title: "UsingMap", ExpectedInput: args.Map{"val": "y"}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"val": val})
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_Clone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOptions clone", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": hm.Has("a")})
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_NoClone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(false, 0, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOptions no clone", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": hm.Has("a")})
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_EmptyMap(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{})
	tc := coretestcases.CaseV1{Title: "UsingMapOptions empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Valid(t *testing.T) {
	m := map[string]string{"a": "1", "b": "2"}
	hm := corestr.New.Hashmap.MapWithCap(5, m)
	tc := coretestcases.CaseV1{Title: "MapWithCap valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.MapWithCap(5, map[string]string{})
	tc := coretestcases.CaseV1{Title: "MapWithCap empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hm.Length()})
}

func Test_Cov43_NewHashmapCreator_MapWithCap_ZeroCap(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.MapWithCap(0, m)
	tc := coretestcases.CaseV1{Title: "MapWithCap zero cap", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": hm.Has("a")})
}

// ═══════════════════════════════════════════════════════════════
// newHashsetCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashsetCreator_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Empty()
	tc := coretestcases.CaseV1{Title: "Empty hashset", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_Cap(t *testing.T) {
	hs := corestr.New.Hashset.Cap(10)
	tc := coretestcases.CaseV1{Title: "Cap hashset", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_Strings_Valid(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{Title: "Strings valid", ExpectedInput: args.Map{"len": 3}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_Strings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{})
	tc := coretestcases.CaseV1{Title: "Strings empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x", "y")
	tc := coretestcases.CaseV1{Title: "StringsSpreadItems", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems_Empty(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems()
	tc := coretestcases.CaseV1{Title: "StringsSpreadItems empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_StringsOption_ValidNoClone(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false, "a", "b")
	tc := coretestcases.CaseV1{Title: "StringsOption valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilZeroCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false)
	tc := coretestcases.CaseV1{Title: "StringsOption nil zero cap", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilWithCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(5, false)
	tc := coretestcases.CaseV1{Title: "StringsOption nil with cap", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Valid(t *testing.T) {
	a, b := "a", "b"
	hs := corestr.New.Hashset.PointerStrings([]*string{&a, &b})
	tc := coretestcases.CaseV1{Title: "PointerStrings valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.PointerStrings([]*string{})
	tc := coretestcases.CaseV1{Title: "PointerStrings empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Valid(t *testing.T) {
	a := "a"
	arr := []*string{&a}
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, true, &arr)
	tc := coretestcases.CaseV1{Title: "PointerStringsPtrOption valid clone", ExpectedInput: args.Map{"len": 1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Nil(t *testing.T) {
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, false, nil)
	tc := coretestcases.CaseV1{Title: "PointerStringsPtrOption nil", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Valid(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := coretestcases.CaseV1{Title: "UsingCollection valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Nil(t *testing.T) {
	hs := corestr.New.Hashset.UsingCollection(nil)
	tc := coretestcases.CaseV1{Title: "UsingCollection nil", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := coretestcases.CaseV1{Title: "UsingCollection empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Valid(t *testing.T) {
	ss := corestr.New.SimpleSlice.Strings([]string{"x", "y"})
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := coretestcases.CaseV1{Title: "SimpleSlice valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Empty(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := coretestcases.CaseV1{Title: "SimpleSlice empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingMap_Valid(t *testing.T) {
	m := map[string]bool{"a": true, "b": true}
	hs := corestr.New.Hashset.UsingMap(m)
	tc := coretestcases.CaseV1{Title: "UsingMap valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingMap_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMap(map[string]bool{})
	tc := coretestcases.CaseV1{Title: "UsingMap empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Clone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(5, true, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOption clone", ExpectedInput: args.Map{"len": 1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_NoClone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(0, false, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOption no clone", ExpectedInput: args.Map{"len": 1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{})
	tc := coretestcases.CaseV1{Title: "UsingMapOption empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hs.Length()})
}

// ═══════════════════════════════════════════════════════════════
// newSimpleStringOnceCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewSimpleStringOnceCreator_Any_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(false, "hello", true)
	tc := coretestcases.CaseV1{Title: "Any init", ExpectedInput: args.Map{"isInit": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isInit": sso.IsInitialized()})
}

func Test_Cov43_NewSimpleStringOnceCreator_Any_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(true, 42, false)
	tc := coretestcases.CaseV1{Title: "Any uninit", ExpectedInput: args.Map{"isInit": false}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isInit": sso.IsInitialized()})
}

func Test_Cov43_NewSimpleStringOnceCreator_Uninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("test")
	tc := coretestcases.CaseV1{Title: "Uninitialized", ExpectedInput: args.Map{"isInit": false}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isInit": sso.IsInitialized()})
}

func Test_Cov43_NewSimpleStringOnceCreator_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("val")
	tc := coretestcases.CaseV1{Title: "Init", ExpectedInput: args.Map{"value": "val"}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": sso.Value()})
}

func Test_Cov43_NewSimpleStringOnceCreator_InitPtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("pval")
	tc := coretestcases.CaseV1{Title: "InitPtr", ExpectedInput: args.Map{"isInit": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isInit": sso.IsInitialized()})
}

func Test_Cov43_NewSimpleStringOnceCreator_Create(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Create("cv", true)
	tc := coretestcases.CaseV1{Title: "Create", ExpectedInput: args.Map{"value": "cv"}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": sso.Value()})
}

func Test_Cov43_NewSimpleStringOnceCreator_CreatePtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.CreatePtr("cpv", false)
	tc := coretestcases.CaseV1{Title: "CreatePtr uninit", ExpectedInput: args.Map{"isInit": false}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isInit": sso.IsInitialized()})
}

func Test_Cov43_NewSimpleStringOnceCreator_Empty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	tc := coretestcases.CaseV1{Title: "Empty SSO", ExpectedInput: args.Map{"isEmpty": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": sso.IsEmpty()})
}

// ═══════════════════════════════════════════════════════════════
// newCharHashsetMapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewCharHashsetMapCreator_Cap(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(20, 10)
	tc := coretestcases.CaseV1{Title: "Cap", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": chm.Length()})
}

func Test_Cov43_NewCharHashsetMapCreator_Cap_BelowLimit(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(1, 1)
	tc := coretestcases.CaseV1{Title: "Cap below limit", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": chm.Length()})
}

func Test_Cov43_NewCharHashsetMapCreator_CapItems(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
	tc := coretestcases.CaseV1{Title: "CapItems", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": chm.Length()})
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Valid(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, []string{"alpha", "beta"})
	tc := coretestcases.CaseV1{Title: "Strings valid", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": chm.Length()})
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Nil(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, nil)
	tc := coretestcases.CaseV1{Title: "Strings nil", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": chm.Length()})
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashmapDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashmapDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashmapDataModel{Items: map[string]string{"k": "v"}}
	hm := corestr.NewHashmapUsingDataModel(dm)
	val, _ := hm.Get("k")
	tc := coretestcases.CaseV1{Title: "NewHashmapUsingDataModel", ExpectedInput: args.Map{"val": "v"}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"val": val})
}

func Test_Cov43_HashmapDataModel_NewDataModelUsing(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	hm.AddOrUpdate("x", "y")
	dm := corestr.NewHashmapsDataModelUsing(hm)
	tc := coretestcases.CaseV1{Title: "NewHashmapsDataModelUsing", ExpectedInput: args.Map{"val": "y"}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"val": dm.Items["x"]})
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
	hs := corestr.NewHashsetUsingDataModel(dm)
	tc := coretestcases.CaseV1{Title: "NewHashsetUsingDataModel", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": hs.Has("a")})
}

func Test_Cov43_HashsetDataModel_NewDataModelUsing(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x")
	dm := corestr.NewHashsetsDataModelUsing(hs)
	tc := coretestcases.CaseV1{Title: "NewHashsetsDataModelUsing", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": dm.Items["x"]})
}

// ═══════════════════════════════════════════════════════════════
// Data Models — CharHashsetDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CharHashsetDataModel_NewUsingDataModel(t *testing.T) {
	innerHs := corestr.New.Hashset.StringsSpreadItems("apple")
	dm := &corestr.CharHashsetDataModel{
		Items:               map[byte]*corestr.Hashset{'a': innerHs},
		EachHashsetCapacity: 10,
	}
	chm := corestr.NewCharHashsetMapUsingDataModel(dm)
	tc := coretestcases.CaseV1{Title: "NewCharHashsetMapUsingDataModel", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": chm.Has("apple")})
}

func Test_Cov43_CharHashsetDataModel_NewDataModelUsing(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "banana")
	dm := corestr.NewCharHashsetMapDataModelUsing(chm)
	tc := coretestcases.CaseV1{Title: "NewCharHashsetMapDataModelUsing", ExpectedInput: args.Map{"cap": 10}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"cap": dm.EachHashsetCapacity})
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetsCollectionDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetsCollectionDataModel_NewUsingDataModel(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{hs1}}
	hsc := corestr.NewHashsetsCollectionUsingDataModel(dm)
	tc := coretestcases.CaseV1{Title: "NewHashsetsCollectionUsingDataModel", ExpectedInput: args.Map{"len": 1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": hsc.Length()})
}

func Test_Cov43_HashsetsCollectionDataModel_NewDataModelUsing(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	hsc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs1)
	dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
	tc := coretestcases.CaseV1{Title: "NewHashsetsCollectionDataModelUsing", ExpectedInput: args.Map{"len": 1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(dm.Items)})
}

// ═══════════════════════════════════════════════════════════════
// Data Models — SimpleStringOnceModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_SimpleStringOnceModel_Fields(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "test", IsInitialize: true}
	tc := coretestcases.CaseV1{Title: "SimpleStringOnceModel Value", ExpectedInput: args.Map{"value": "test"}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": m.Value})
}

func Test_Cov43_SimpleStringOnceModel_IsInit(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "v", IsInitialize: true}
	tc := coretestcases.CaseV1{Title: "SimpleStringOnceModel IsInit", ExpectedInput: args.Map{"isInit": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isInit": m.IsInitialize})
}

// ═══════════════════════════════════════════════════════════════
// Data Models — CollectionsOfCollectionModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CollectionsOfCollectionModel_Fields(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	m := corestr.CollectionsOfCollectionModel{Items: []*corestr.Collection{col}}
	tc := coretestcases.CaseV1{Title: "CollectionsOfCollectionModel", ExpectedInput: args.Map{"len": 1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(m.Items)})
}

// ═══════════════════════════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AnyToString_EmptyString(t *testing.T) {
	result := corestr.AnyToString(false, "")
	tc := coretestcases.CaseV1{Title: "AnyToString empty", ExpectedInput: args.Map{"value": ""}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": result})
}

func Test_Cov43_AnyToString_WithFieldName(t *testing.T) {
	result := corestr.AnyToString(true, "hello")
	tc := coretestcases.CaseV1{Title: "AnyToString with field name", ExpectedInput: args.Map{"hasContent": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"hasContent": len(result) > 0})
}

func Test_Cov43_AnyToString_WithoutFieldName(t *testing.T) {
	result := corestr.AnyToString(false, 42)
	tc := coretestcases.CaseV1{Title: "AnyToString without field name", ExpectedInput: args.Map{"hasContent": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"hasContent": len(result) > 0})
}

func Test_Cov43_AnyToString_Pointer(t *testing.T) {
	v := "ptr"
	result := corestr.AnyToString(false, &v)
	tc := coretestcases.CaseV1{Title: "AnyToString pointer", ExpectedInput: args.Map{"hasContent": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"hasContent": len(result) > 0})
}

// ═══════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_InvalidValueStatusNoMessage(t *testing.T) {
	vs := corestr.InvalidValueStatusNoMessage()
	tc := coretestcases.CaseV1{Title: "InvalidValueStatusNoMessage index", ExpectedInput: args.Map{"index": -1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"index": vs.Index})
}

func Test_Cov43_InvalidValueStatus(t *testing.T) {
	vs := corestr.InvalidValueStatus("err msg")
	tc := coretestcases.CaseV1{Title: "InvalidValueStatus", ExpectedInput: args.Map{"index": -1}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"index": vs.Index})
}

func Test_Cov43_ValueStatus_Clone(t *testing.T) {
	vs := corestr.InvalidValueStatus("msg")
	cloned := vs.Clone()
	tc := coretestcases.CaseV1{Title: "ValueStatus Clone", ExpectedInput: args.Map{"index": vs.Index}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"index": cloned.Index})
}

// ═══════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_TextWithLineNumber_HasLineNumber_Valid(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
	tc := coretestcases.CaseV1{Title: "HasLineNumber valid", ExpectedInput: args.Map{"has": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": twl.HasLineNumber()})
}

func Test_Cov43_TextWithLineNumber_HasLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "HasLineNumber nil", ExpectedInput: args.Map{"has": false}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"has": twl.HasLineNumber()})
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
	tc := coretestcases.CaseV1{Title: "IsInvalidLineNumber", ExpectedInput: args.Map{"invalid": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"invalid": twl.IsInvalidLineNumber()})
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "IsInvalidLineNumber nil", ExpectedInput: args.Map{"invalid": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"invalid": twl.IsInvalidLineNumber()})
}

func Test_Cov43_TextWithLineNumber_Length(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}
	tc := coretestcases.CaseV1{Title: "Length", ExpectedInput: args.Map{"len": 3}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": twl.Length()})
}

func Test_Cov43_TextWithLineNumber_Length_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "Length nil", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": twl.Length()})
}

func Test_Cov43_TextWithLineNumber_IsEmpty_True(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := coretestcases.CaseV1{Title: "IsEmpty true", ExpectedInput: args.Map{"isEmpty": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": twl.IsEmpty()})
}

func Test_Cov43_TextWithLineNumber_IsEmpty_False(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
	tc := coretestcases.CaseV1{Title: "IsEmpty false", ExpectedInput: args.Map{"isEmpty": false}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": twl.IsEmpty()})
}

func Test_Cov43_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "IsEmpty nil", ExpectedInput: args.Map{"isEmpty": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": twl.IsEmpty()})
}

func Test_Cov43_TextWithLineNumber_IsEmptyText(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
	tc := coretestcases.CaseV1{Title: "IsEmptyText", ExpectedInput: args.Map{"isEmpty": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": twl.IsEmptyText()})
}

func Test_Cov43_TextWithLineNumber_IsEmptyText_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "IsEmptyText nil", ExpectedInput: args.Map{"isEmpty": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": twl.IsEmptyText()})
}

func Test_Cov43_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := coretestcases.CaseV1{Title: "IsEmptyTextLineBoth", ExpectedInput: args.Map{"isEmpty": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": twl.IsEmptyTextLineBoth()})
}

// ═══════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CloneSlice_Valid(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := corestr.CloneSlice(input)
	tc := coretestcases.CaseV1{Title: "CloneSlice valid", ExpectedInput: args.Map{"len": 3}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(result)})
}

func Test_Cov43_CloneSlice_Nil(t *testing.T) {
	result := corestr.CloneSlice(nil)
	tc := coretestcases.CaseV1{Title: "CloneSlice nil", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(result)})
}

func Test_Cov43_CloneSliceIf_Clone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(true, input...)
	tc := coretestcases.CaseV1{Title: "CloneSliceIf clone", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(result)})
}

func Test_Cov43_CloneSliceIf_NoClone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(false, input...)
	tc := coretestcases.CaseV1{Title: "CloneSliceIf no clone same ref", ExpectedInput: args.Map{"len": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(result)})
}

// ═══════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength / AllIndividualsLengthOfSimpleSlices
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AllIndividualStringsOfStringsLength_Valid(t *testing.T) {
	input := [][]string{{"a", "b"}, {"c"}}
	result := corestr.AllIndividualStringsOfStringsLength(&input)
	tc := coretestcases.CaseV1{Title: "AllIndividualStringsOfStringsLength", ExpectedInput: args.Map{"len": 3}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": result})
}

func Test_Cov43_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	input := [][]string{}
	result := corestr.AllIndividualStringsOfStringsLength(&input)
	tc := coretestcases.CaseV1{Title: "AllIndividualStringsOfStringsLength empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": result})
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Valid(t *testing.T) {
	ss1 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
	ss2 := corestr.New.SimpleSlice.Strings([]string{"c"})
	result := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)
	tc := coretestcases.CaseV1{Title: "AllIndividualsLengthOfSimpleSlices", ExpectedInput: args.Map{"len": 3}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": result})
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Empty(t *testing.T) {
	result := corestr.AllIndividualsLengthOfSimpleSlices()
	tc := coretestcases.CaseV1{Title: "AllIndividualsLengthOfSimpleSlices empty", ExpectedInput: args.Map{"len": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": result})
}

// ═══════════════════════════════════════════════════════════════
// Vars — StaticJsonError, LeftRightExpectingLengthMessager
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_StaticJsonError_NotNil(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "StaticJsonError not nil", ExpectedInput: args.Map{"notNil": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"notNil": corestr.StaticJsonError != nil})
}

func Test_Cov43_ExpectingLengthForLeftRight(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "ExpectingLengthForLeftRight", ExpectedInput: args.Map{"value": 2}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": corestr.ExpectingLengthForLeftRight})
}

func Test_Cov43_LeftRightExpectingLengthMessager_NotNil(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "LeftRightExpectingLengthMessager not nil", ExpectedInput: args.Map{"notNil": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"notNil": corestr.LeftRightExpectingLengthMessager != nil})
}

// ═══════════════════════════════════════════════════════════════
// Funcs types — ReturningBool, filter types
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_ReturningBool_Fields(t *testing.T) {
	rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}
	tc := coretestcases.CaseV1{Title: "ReturningBool IsBreak", ExpectedInput: args.Map{"isBreak": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isBreak": rb.IsBreak})
}

func Test_Cov43_LinkedCollectionFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedCollectionFilterResult{IsKeep: true, IsBreak: false}
	tc := coretestcases.CaseV1{Title: "LinkedCollectionFilterResult", ExpectedInput: args.Map{"isKeep": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isKeep": r.IsKeep})
}

func Test_Cov43_LinkedListFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedListFilterResult{IsKeep: false, IsBreak: true}
	tc := coretestcases.CaseV1{Title: "LinkedListFilterResult", ExpectedInput: args.Map{"isBreak": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isBreak": r.IsBreak})
}

func Test_Cov43_LinkedCollectionFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionFilterParameter{Index: 5}
	tc := coretestcases.CaseV1{Title: "LinkedCollectionFilterParameter", ExpectedInput: args.Map{"index": 5}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"index": p.Index})
}

func Test_Cov43_LinkedListFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedListFilterParameter{Index: 3}
	tc := coretestcases.CaseV1{Title: "LinkedListFilterParameter", ExpectedInput: args.Map{"index": 3}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"index": p.Index})
}

func Test_Cov43_LinkedListProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedListProcessorParameter{Index: 0, IsFirstIndex: true, IsEndingIndex: false}
	tc := coretestcases.CaseV1{Title: "LinkedListProcessorParameter", ExpectedInput: args.Map{"isFirst": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isFirst": p.IsFirstIndex})
}

func Test_Cov43_LinkedCollectionProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionProcessorParameter{Index: 1, IsFirstIndex: false, IsEndingIndex: true}
	tc := coretestcases.CaseV1{Title: "LinkedCollectionProcessorParameter", ExpectedInput: args.Map{"isEnding": true}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEnding": p.IsEndingIndex})
}

// ═══════════════════════════════════════════════════════════════
// Consts — RegularCollectionEfficiencyLimit, DoubleLimit, NoElements
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_RegularCollectionEfficiencyLimit(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "RegularCollectionEfficiencyLimit", ExpectedInput: args.Map{"value": 1000}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": corestr.RegularCollectionEfficiencyLimit})
}

func Test_Cov43_DoubleLimit(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "DoubleLimit", ExpectedInput: args.Map{"value": 3000}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": corestr.DoubleLimit})
}

func Test_Cov43_NoElements(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "NoElements", ExpectedInput: args.Map{"value": 0}}
	tc.ShouldBeEqualMapFirst(t, args.Map{"value": corestr.NoElements})
}

// Suppress unused import warning
var _ = fmt.Sprintf
