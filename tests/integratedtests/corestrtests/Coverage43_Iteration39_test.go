package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// newHashmapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashmapCreator_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	tc := coretestcases.CaseV1{Name: "Empty hashmap", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_Cap(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(10)
	tc := coretestcases.CaseV1{Name: "Cap hashmap empty", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Valid(t *testing.T) {
	pair := corestr.KeyAnyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyAnyValues(pair)
	tc := coretestcases.CaseV1{Name: "KeyAnyValues valid", Expected: true, Actual: hm.Has("k1"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyAnyValues()
	tc := coretestcases.CaseV1{Name: "KeyAnyValues empty", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValues_Valid(t *testing.T) {
	pair := corestr.KeyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyValues(pair)
	tc := coretestcases.CaseV1{Name: "KeyValues valid", Expected: true, Actual: hm.Has("k1"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues()
	tc := coretestcases.CaseV1{Name: "KeyValues empty", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_Valid(t *testing.T) {
	keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
	vals := corestr.New.Collection.Strings([]string{"v1", "v2"})
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := coretestcases.CaseV1{Name: "KeyValuesCollection valid", Expected: 2, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_NilKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
	tc := coretestcases.CaseV1{Name: "KeyValuesCollection nil", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_EmptyKeys(t *testing.T) {
	keys := corestr.New.Collection.Empty()
	vals := corestr.New.Collection.Empty()
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := coretestcases.CaseV1{Name: "KeyValuesCollection empty keys", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_Valid(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{"a", "b"}, []string{"1", "2"})
	tc := coretestcases.CaseV1{Name: "KeyValuesStrings valid", Expected: 2, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_EmptyKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{}, []string{})
	tc := coretestcases.CaseV1{Name: "KeyValuesStrings empty", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMap(t *testing.T) {
	m := map[string]string{"x": "y"}
	hm := corestr.New.Hashmap.UsingMap(m)
	tc := coretestcases.CaseV1{Name: "UsingMap", Expected: "y", Actual: hm.Get("x"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_Clone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, m)
	tc := coretestcases.CaseV1{Name: "UsingMapOptions clone", Expected: true, Actual: hm.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_NoClone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(false, 0, m)
	tc := coretestcases.CaseV1{Name: "UsingMapOptions no clone", Expected: true, Actual: hm.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_EmptyMap(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{})
	tc := coretestcases.CaseV1{Name: "UsingMapOptions empty", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Valid(t *testing.T) {
	m := map[string]string{"a": "1", "b": "2"}
	hm := corestr.New.Hashmap.MapWithCap(5, m)
	tc := coretestcases.CaseV1{Name: "MapWithCap valid", Expected: 2, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.MapWithCap(5, map[string]string{})
	tc := coretestcases.CaseV1{Name: "MapWithCap empty", Expected: 0, Actual: hm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_ZeroCap(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.MapWithCap(0, m)
	tc := coretestcases.CaseV1{Name: "MapWithCap zero cap", Expected: true, Actual: hm.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// newHashsetCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashsetCreator_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Empty()
	tc := coretestcases.CaseV1{Name: "Empty hashset", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_Cap(t *testing.T) {
	hs := corestr.New.Hashset.Cap(10)
	tc := coretestcases.CaseV1{Name: "Cap hashset", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_Strings_Valid(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{Name: "Strings valid", Expected: 3, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_Strings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{})
	tc := coretestcases.CaseV1{Name: "Strings empty", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x", "y")
	tc := coretestcases.CaseV1{Name: "StringsSpreadItems", Expected: 2, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems_Empty(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems()
	tc := coretestcases.CaseV1{Name: "StringsSpreadItems empty", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsOption_ValidNoClone(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false, "a", "b")
	tc := coretestcases.CaseV1{Name: "StringsOption valid", Expected: 2, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilZeroCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false)
	tc := coretestcases.CaseV1{Name: "StringsOption nil zero cap", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilWithCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(5, false)
	tc := coretestcases.CaseV1{Name: "StringsOption nil with cap", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Valid(t *testing.T) {
	a, b := "a", "b"
	hs := corestr.New.Hashset.PointerStrings([]*string{&a, &b})
	tc := coretestcases.CaseV1{Name: "PointerStrings valid", Expected: 2, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.PointerStrings([]*string{})
	tc := coretestcases.CaseV1{Name: "PointerStrings empty", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Valid(t *testing.T) {
	a := "a"
	arr := []*string{&a}
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, true, &arr)
	tc := coretestcases.CaseV1{Name: "PointerStringsPtrOption valid clone", Expected: 1, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Nil(t *testing.T) {
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, false, nil)
	tc := coretestcases.CaseV1{Name: "PointerStringsPtrOption nil", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Valid(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := coretestcases.CaseV1{Name: "UsingCollection valid", Expected: 2, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Nil(t *testing.T) {
	hs := corestr.New.Hashset.UsingCollection(nil)
	tc := coretestcases.CaseV1{Name: "UsingCollection nil", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := coretestcases.CaseV1{Name: "UsingCollection empty", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Valid(t *testing.T) {
	ss := corestr.New.SimpleSlice.Strings([]string{"x", "y"})
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := coretestcases.CaseV1{Name: "SimpleSlice valid", Expected: 2, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Empty(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := coretestcases.CaseV1{Name: "SimpleSlice empty", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMap_Valid(t *testing.T) {
	m := map[string]bool{"a": true, "b": true}
	hs := corestr.New.Hashset.UsingMap(m)
	tc := coretestcases.CaseV1{Name: "UsingMap valid", Expected: 2, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMap_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMap(map[string]bool{})
	tc := coretestcases.CaseV1{Name: "UsingMap empty", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Clone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(5, true, m)
	tc := coretestcases.CaseV1{Name: "UsingMapOption clone", Expected: 1, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_NoClone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(0, false, m)
	tc := coretestcases.CaseV1{Name: "UsingMapOption no clone", Expected: 1, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{})
	tc := coretestcases.CaseV1{Name: "UsingMapOption empty", Expected: 0, Actual: hs.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// newSimpleStringOnceCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewSimpleStringOnceCreator_Any_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(false, "hello", true)
	tc := coretestcases.CaseV1{Name: "Any init", Expected: true, Actual: sso.IsInitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Any_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(true, 42, false)
	tc := coretestcases.CaseV1{Name: "Any uninit", Expected: false, Actual: sso.IsInitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Uninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("test")
	tc := coretestcases.CaseV1{Name: "Uninitialized", Expected: false, Actual: sso.IsInitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("val")
	tc := coretestcases.CaseV1{Name: "Init", Expected: "val", Actual: sso.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_InitPtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("pval")
	tc := coretestcases.CaseV1{Name: "InitPtr", Expected: true, Actual: sso.IsInitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Create(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Create("cv", true)
	tc := coretestcases.CaseV1{Name: "Create", Expected: "cv", Actual: sso.Value(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_CreatePtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.CreatePtr("cpv", false)
	tc := coretestcases.CaseV1{Name: "CreatePtr uninit", Expected: false, Actual: sso.IsInitialized(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Empty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	tc := coretestcases.CaseV1{Name: "Empty SSO", Expected: true, Actual: sso.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// newCharHashsetMapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewCharHashsetMapCreator_Cap(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(20, 10)
	tc := coretestcases.CaseV1{Name: "Cap", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_Cap_BelowLimit(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(1, 1)
	tc := coretestcases.CaseV1{Name: "Cap below limit", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_CapItems(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
	tc := coretestcases.CaseV1{Name: "CapItems", Expected: 2, Actual: chm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Valid(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, []string{"alpha", "beta"})
	tc := coretestcases.CaseV1{Name: "Strings valid", Expected: 2, Actual: chm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Nil(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, nil)
	tc := coretestcases.CaseV1{Name: "Strings nil", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashmapDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashmapDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashmapDataModel{Items: map[string]string{"k": "v"}}
	hm := corestr.NewHashmapUsingDataModel(dm)
	tc := coretestcases.CaseV1{Name: "NewHashmapUsingDataModel", Expected: "v", Actual: hm.Get("k"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_HashmapDataModel_NewDataModelUsing(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	hm.AddOrUpdate("x", "y")
	dm := corestr.NewHashmapsDataModelUsing(hm)
	tc := coretestcases.CaseV1{Name: "NewHashmapsDataModelUsing", Expected: "y", Actual: dm.Items["x"], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
	hs := corestr.NewHashsetUsingDataModel(dm)
	tc := coretestcases.CaseV1{Name: "NewHashsetUsingDataModel", Expected: true, Actual: hs.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_HashsetDataModel_NewDataModelUsing(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x")
	dm := corestr.NewHashsetsDataModelUsing(hs)
	tc := coretestcases.CaseV1{Name: "NewHashsetsDataModelUsing", Expected: true, Actual: dm.Items["x"], Args: args.Map{}}
	tc.ShouldBeEqual(t)
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
	tc := coretestcases.CaseV1{Name: "NewCharHashsetMapUsingDataModel", Expected: true, Actual: chm.Has("apple"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CharHashsetDataModel_NewDataModelUsing(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "banana")
	dm := corestr.NewCharHashsetMapDataModelUsing(chm)
	tc := coretestcases.CaseV1{Name: "NewCharHashsetMapDataModelUsing", Expected: 10, Actual: dm.EachHashsetCapacity, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetsCollectionDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetsCollectionDataModel_NewUsingDataModel(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{hs1}}
	hsc := corestr.NewHashsetsCollectionUsingDataModel(dm)
	tc := coretestcases.CaseV1{Name: "NewHashsetsCollectionUsingDataModel", Expected: 1, Actual: hsc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_HashsetsCollectionDataModel_NewDataModelUsing(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	hsc := corestr.New.HashsetsCollection.UsingHashsets(hs1)
	dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
	tc := coretestcases.CaseV1{Name: "NewHashsetsCollectionDataModelUsing", Expected: 1, Actual: len(dm.Items), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — SimpleStringOnceModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_SimpleStringOnceModel_Fields(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "test", IsInitialize: true}
	tc := coretestcases.CaseV1{Name: "SimpleStringOnceModel Value", Expected: "test", Actual: m.Value, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_SimpleStringOnceModel_IsInit(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "v", IsInitialize: true}
	tc := coretestcases.CaseV1{Name: "SimpleStringOnceModel IsInit", Expected: true, Actual: m.IsInitialize, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — CollectionsOfCollectionModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CollectionsOfCollectionModel_Fields(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	m := corestr.CollectionsOfCollectionModel{Items: []*corestr.Collection{col}}
	tc := coretestcases.CaseV1{Name: "CollectionsOfCollectionModel", Expected: 1, Actual: len(m.Items), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AnyToString_EmptyString(t *testing.T) {
	result := corestr.AnyToString(false, "")
	tc := coretestcases.CaseV1{Name: "AnyToString empty", Expected: "", Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AnyToString_WithFieldName(t *testing.T) {
	result := corestr.AnyToString(true, "hello")
	tc := coretestcases.CaseV1{Name: "AnyToString with field name", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AnyToString_WithoutFieldName(t *testing.T) {
	result := corestr.AnyToString(false, 42)
	tc := coretestcases.CaseV1{Name: "AnyToString without field name", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AnyToString_Pointer(t *testing.T) {
	v := "ptr"
	result := corestr.AnyToString(false, &v)
	tc := coretestcases.CaseV1{Name: "AnyToString pointer", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_InvalidValueStatusNoMessage(t *testing.T) {
	vs := corestr.InvalidValueStatusNoMessage()
	tc := coretestcases.CaseV1{Name: "InvalidValueStatusNoMessage index", Expected: -1, Actual: vs.Index, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_InvalidValueStatus(t *testing.T) {
	vs := corestr.InvalidValueStatus("err msg")
	tc := coretestcases.CaseV1{Name: "InvalidValueStatus", Expected: -1, Actual: vs.Index, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_ValueStatus_Clone(t *testing.T) {
	vs := corestr.InvalidValueStatus("msg")
	cloned := vs.Clone()
	tc := coretestcases.CaseV1{Name: "ValueStatus Clone", Expected: vs.Index, Actual: cloned.Index, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_TextWithLineNumber_HasLineNumber_Valid(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
	tc := coretestcases.CaseV1{Name: "HasLineNumber valid", Expected: true, Actual: twl.HasLineNumber(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_HasLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Name: "HasLineNumber nil", Expected: false, Actual: twl.HasLineNumber(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
	tc := coretestcases.CaseV1{Name: "IsInvalidLineNumber", Expected: true, Actual: twl.IsInvalidLineNumber(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Name: "IsInvalidLineNumber nil", Expected: true, Actual: twl.IsInvalidLineNumber(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_Length(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}
	tc := coretestcases.CaseV1{Name: "Length", Expected: 3, Actual: twl.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_Length_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Name: "Length nil", Expected: 0, Actual: twl.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_True(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := coretestcases.CaseV1{Name: "IsEmpty true", Expected: true, Actual: twl.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_False(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
	tc := coretestcases.CaseV1{Name: "IsEmpty false", Expected: false, Actual: twl.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Name: "IsEmpty nil", Expected: true, Actual: twl.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmptyText(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
	tc := coretestcases.CaseV1{Name: "IsEmptyText", Expected: true, Actual: twl.IsEmptyText(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmptyText_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Name: "IsEmptyText nil", Expected: true, Actual: twl.IsEmptyText(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := coretestcases.CaseV1{Name: "IsEmptyTextLineBoth", Expected: true, Actual: twl.IsEmptyTextLineBoth(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CloneSlice_Valid(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := corestr.CloneSlice(input)
	tc := coretestcases.CaseV1{Name: "CloneSlice valid", Expected: 3, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CloneSlice_Nil(t *testing.T) {
	result := corestr.CloneSlice(nil)
	tc := coretestcases.CaseV1{Name: "CloneSlice nil", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CloneSliceIf_Clone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(true, input)
	tc := coretestcases.CaseV1{Name: "CloneSliceIf clone", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CloneSliceIf_NoClone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(false, input)
	tc := coretestcases.CaseV1{Name: "CloneSliceIf no clone same ref", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength / AllIndividualsLengthOfSimpleSlices
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AllIndividualStringsOfStringsLength_Valid(t *testing.T) {
	input := [][]string{{"a", "b"}, {"c"}}
	result := corestr.AllIndividualStringsOfStringsLength(input)
	tc := coretestcases.CaseV1{Name: "AllIndividualStringsOfStringsLength", Expected: 3, Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	result := corestr.AllIndividualStringsOfStringsLength([][]string{})
	tc := coretestcases.CaseV1{Name: "AllIndividualStringsOfStringsLength empty", Expected: 0, Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Valid(t *testing.T) {
	ss1 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
	ss2 := corestr.New.SimpleSlice.Strings([]string{"c"})
	result := corestr.AllIndividualsLengthOfSimpleSlices([]*corestr.SimpleSlice{ss1, ss2})
	tc := coretestcases.CaseV1{Name: "AllIndividualsLengthOfSimpleSlices", Expected: 3, Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Empty(t *testing.T) {
	result := corestr.AllIndividualsLengthOfSimpleSlices([]*corestr.SimpleSlice{})
	tc := coretestcases.CaseV1{Name: "AllIndividualsLengthOfSimpleSlices empty", Expected: 0, Actual: result, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Vars — StaticJsonError, LeftRightExpectingLengthMessager
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_StaticJsonError_NotNil(t *testing.T) {
	tc := coretestcases.CaseV1{Name: "StaticJsonError not nil", Expected: true, Actual: corestr.StaticJsonError != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_ExpectingLengthForLeftRight(t *testing.T) {
	tc := coretestcases.CaseV1{Name: "ExpectingLengthForLeftRight", Expected: 2, Actual: corestr.ExpectingLengthForLeftRight, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LeftRightExpectingLengthMessager_NotNil(t *testing.T) {
	tc := coretestcases.CaseV1{Name: "LeftRightExpectingLengthMessager not nil", Expected: true, Actual: corestr.LeftRightExpectingLengthMessager != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Funcs types — ReturningBool, filter types
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_ReturningBool_Fields(t *testing.T) {
	rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}
	tc := coretestcases.CaseV1{Name: "ReturningBool IsBreak", Expected: true, Actual: rb.IsBreak, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedCollectionFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedCollectionFilterResult{IsKeep: true, IsBreak: false}
	tc := coretestcases.CaseV1{Name: "LinkedCollectionFilterResult", Expected: true, Actual: r.IsKeep, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedListFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedListFilterResult{IsKeep: false, IsBreak: true}
	tc := coretestcases.CaseV1{Name: "LinkedListFilterResult", Expected: true, Actual: r.IsBreak, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedCollectionFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionFilterParameter{Index: 5}
	tc := coretestcases.CaseV1{Name: "LinkedCollectionFilterParameter", Expected: 5, Actual: p.Index, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedListFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedListFilterParameter{Index: 3}
	tc := coretestcases.CaseV1{Name: "LinkedListFilterParameter", Expected: 3, Actual: p.Index, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedListProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedListProcessorParameter{Index: 0, IsFirstIndex: true, IsEndingIndex: false}
	tc := coretestcases.CaseV1{Name: "LinkedListProcessorParameter", Expected: true, Actual: p.IsFirstIndex, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedCollectionProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionProcessorParameter{Index: 1, IsFirstIndex: false, IsEndingIndex: true}
	tc := coretestcases.CaseV1{Name: "LinkedCollectionProcessorParameter", Expected: true, Actual: p.IsEndingIndex, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Consts — RegularCollectionEfficiencyLimit, DoubleLimit, NoElements
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_RegularCollectionEfficiencyLimit(t *testing.T) {
	tc := coretestcases.CaseV1{Name: "RegularCollectionEfficiencyLimit", Expected: 1000, Actual: corestr.RegularCollectionEfficiencyLimit, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_DoubleLimit(t *testing.T) {
	tc := coretestcases.CaseV1{Name: "DoubleLimit", Expected: 3000, Actual: corestr.DoubleLimit, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NoElements(t *testing.T) {
	tc := coretestcases.CaseV1{Name: "NoElements", Expected: " {No Element}", Actual: corestr.NoElements, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
