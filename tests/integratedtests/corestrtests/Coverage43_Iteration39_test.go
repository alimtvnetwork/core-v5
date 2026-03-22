package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// newHashmapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashmapCreator_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	tc := coretestcases.CaseV1{Title: "Empty hashmap", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_Cap(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(10)
	tc := coretestcases.CaseV1{Title: "Cap hashmap empty", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Valid(t *testing.T) {
	pair := corestr.KeyAnyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyAnyValues(pair)
	tc := coretestcases.CaseV1{Title: "KeyAnyValues valid", ExpectedInput: true, ActualInput: hm.Has("k1")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyAnyValues()
	tc := coretestcases.CaseV1{Title: "KeyAnyValues empty", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyValues_Valid(t *testing.T) {
	pair := corestr.KeyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyValues(pair)
	tc := coretestcases.CaseV1{Title: "KeyValues valid", ExpectedInput: true, ActualInput: hm.Has("k1")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues()
	tc := coretestcases.CaseV1{Title: "KeyValues empty", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_Valid(t *testing.T) {
	keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
	vals := corestr.New.Collection.Strings([]string{"v1", "v2"})
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := coretestcases.CaseV1{Title: "KeyValuesCollection valid", ExpectedInput: 2, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_NilKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
	tc := coretestcases.CaseV1{Title: "KeyValuesCollection nil", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_EmptyKeys(t *testing.T) {
	keys := corestr.New.Collection.Empty()
	vals := corestr.New.Collection.Empty()
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := coretestcases.CaseV1{Title: "KeyValuesCollection empty keys", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_Valid(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{"a", "b"}, []string{"1", "2"})
	tc := coretestcases.CaseV1{Title: "KeyValuesStrings valid", ExpectedInput: 2, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_EmptyKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{}, []string{})
	tc := coretestcases.CaseV1{Title: "KeyValuesStrings empty", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_UsingMap(t *testing.T) {
	m := map[string]string{"x": "y"}
	hm := corestr.New.Hashmap.UsingMap(m)
	tc := coretestcases.CaseV1{Title: "UsingMap", ExpectedInput: "y", ActualInput: hm.Get("x")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_Clone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOptions clone", ExpectedInput: true, ActualInput: hm.Has("a")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_NoClone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(false, 0, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOptions no clone", ExpectedInput: true, ActualInput: hm.Has("a")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_EmptyMap(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{})
	tc := coretestcases.CaseV1{Title: "UsingMapOptions empty", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Valid(t *testing.T) {
	m := map[string]string{"a": "1", "b": "2"}
	hm := corestr.New.Hashmap.MapWithCap(5, m)
	tc := coretestcases.CaseV1{Title: "MapWithCap valid", ExpectedInput: 2, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.MapWithCap(5, map[string]string{})
	tc := coretestcases.CaseV1{Title: "MapWithCap empty", ExpectedInput: 0, ActualInput: hm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_ZeroCap(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.MapWithCap(0, m)
	tc := coretestcases.CaseV1{Title: "MapWithCap zero cap", ExpectedInput: true, ActualInput: hm.Has("a")}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// newHashsetCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashsetCreator_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Empty()
	tc := coretestcases.CaseV1{Title: "Empty hashset", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_Cap(t *testing.T) {
	hs := corestr.New.Hashset.Cap(10)
	tc := coretestcases.CaseV1{Title: "Cap hashset", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_Strings_Valid(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{Title: "Strings valid", ExpectedInput: 3, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_Strings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{})
	tc := coretestcases.CaseV1{Title: "Strings empty", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x", "y")
	tc := coretestcases.CaseV1{Title: "StringsSpreadItems", ExpectedInput: 2, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems_Empty(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems()
	tc := coretestcases.CaseV1{Title: "StringsSpreadItems empty", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_StringsOption_ValidNoClone(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false, "a", "b")
	tc := coretestcases.CaseV1{Title: "StringsOption valid", ExpectedInput: 2, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilZeroCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false)
	tc := coretestcases.CaseV1{Title: "StringsOption nil zero cap", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilWithCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(5, false)
	tc := coretestcases.CaseV1{Title: "StringsOption nil with cap", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Valid(t *testing.T) {
	a, b := "a", "b"
	hs := corestr.New.Hashset.PointerStrings([]*string{&a, &b})
	tc := coretestcases.CaseV1{Title: "PointerStrings valid", ExpectedInput: 2, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.PointerStrings([]*string{})
	tc := coretestcases.CaseV1{Title: "PointerStrings empty", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Valid(t *testing.T) {
	a := "a"
	arr := []*string{&a}
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, true, &arr)
	tc := coretestcases.CaseV1{Title: "PointerStringsPtrOption valid clone", ExpectedInput: 1, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Nil(t *testing.T) {
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, false, nil)
	tc := coretestcases.CaseV1{Title: "PointerStringsPtrOption nil", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Valid(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := coretestcases.CaseV1{Title: "UsingCollection valid", ExpectedInput: 2, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Nil(t *testing.T) {
	hs := corestr.New.Hashset.UsingCollection(nil)
	tc := coretestcases.CaseV1{Title: "UsingCollection nil", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := coretestcases.CaseV1{Title: "UsingCollection empty", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Valid(t *testing.T) {
	ss := corestr.New.SimpleSlice.Strings([]string{"x", "y"})
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := coretestcases.CaseV1{Title: "SimpleSlice valid", ExpectedInput: 2, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Empty(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := coretestcases.CaseV1{Title: "SimpleSlice empty", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingMap_Valid(t *testing.T) {
	m := map[string]bool{"a": true, "b": true}
	hs := corestr.New.Hashset.UsingMap(m)
	tc := coretestcases.CaseV1{Title: "UsingMap valid", ExpectedInput: 2, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingMap_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMap(map[string]bool{})
	tc := coretestcases.CaseV1{Title: "UsingMap empty", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Clone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(5, true, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOption clone", ExpectedInput: 1, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_NoClone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(0, false, m)
	tc := coretestcases.CaseV1{Title: "UsingMapOption no clone", ExpectedInput: 1, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{})
	tc := coretestcases.CaseV1{Title: "UsingMapOption empty", ExpectedInput: 0, ActualInput: hs.Length()}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// newSimpleStringOnceCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewSimpleStringOnceCreator_Any_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(false, "hello", true)
	tc := coretestcases.CaseV1{Title: "Any init", ExpectedInput: true, ActualInput: sso.IsInitialized()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewSimpleStringOnceCreator_Any_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(true, 42, false)
	tc := coretestcases.CaseV1{Title: "Any uninit", ExpectedInput: false, ActualInput: sso.IsInitialized()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewSimpleStringOnceCreator_Uninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("test")
	tc := coretestcases.CaseV1{Title: "Uninitialized", ExpectedInput: false, ActualInput: sso.IsInitialized()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewSimpleStringOnceCreator_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("val")
	tc := coretestcases.CaseV1{Title: "Init", ExpectedInput: "val", ActualInput: sso.Value()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewSimpleStringOnceCreator_InitPtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("pval")
	tc := coretestcases.CaseV1{Title: "InitPtr", ExpectedInput: true, ActualInput: sso.IsInitialized()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewSimpleStringOnceCreator_Create(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Create("cv", true)
	tc := coretestcases.CaseV1{Title: "Create", ExpectedInput: "cv", ActualInput: sso.Value()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewSimpleStringOnceCreator_CreatePtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.CreatePtr("cpv", false)
	tc := coretestcases.CaseV1{Title: "CreatePtr uninit", ExpectedInput: false, ActualInput: sso.IsInitialized()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewSimpleStringOnceCreator_Empty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	tc := coretestcases.CaseV1{Title: "Empty SSO", ExpectedInput: true, ActualInput: sso.IsEmpty()}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// newCharHashsetMapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewCharHashsetMapCreator_Cap(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(20, 10)
	tc := coretestcases.CaseV1{Title: "Cap", ExpectedInput: 0, ActualInput: chm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewCharHashsetMapCreator_Cap_BelowLimit(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(1, 1)
	tc := coretestcases.CaseV1{Title: "Cap below limit", ExpectedInput: 0, ActualInput: chm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewCharHashsetMapCreator_CapItems(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
	tc := coretestcases.CaseV1{Title: "CapItems", ExpectedInput: 2, ActualInput: chm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Valid(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, []string{"alpha", "beta"})
	tc := coretestcases.CaseV1{Title: "Strings valid", ExpectedInput: 2, ActualInput: chm.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Nil(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, nil)
	tc := coretestcases.CaseV1{Title: "Strings nil", ExpectedInput: 0, ActualInput: chm.Length()}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashmapDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashmapDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashmapDataModel{Items: map[string]string{"k": "v"}}
	hm := corestr.NewHashmapUsingDataModel(dm)
	tc := coretestcases.CaseV1{Title: "NewHashmapUsingDataModel", ExpectedInput: "v", ActualInput: hm.Get("k")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_HashmapDataModel_NewDataModelUsing(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	hm.AddOrUpdate("x", "y")
	dm := corestr.NewHashmapsDataModelUsing(hm)
	tc := coretestcases.CaseV1{Title: "NewHashmapsDataModelUsing", ExpectedInput: "y", ActualInput: dm.Items["x"]}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
	hs := corestr.NewHashsetUsingDataModel(dm)
	tc := coretestcases.CaseV1{Title: "NewHashsetUsingDataModel", ExpectedInput: true, ActualInput: hs.Has("a")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_HashsetDataModel_NewDataModelUsing(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x")
	dm := corestr.NewHashsetsDataModelUsing(hs)
	tc := coretestcases.CaseV1{Title: "NewHashsetsDataModelUsing", ExpectedInput: true, ActualInput: dm.Items["x"]}
	tc.ShouldBeEqual(t, 0)
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
	tc := coretestcases.CaseV1{Title: "NewCharHashsetMapUsingDataModel", ExpectedInput: true, ActualInput: chm.Has("apple")}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_CharHashsetDataModel_NewDataModelUsing(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "banana")
	dm := corestr.NewCharHashsetMapDataModelUsing(chm)
	tc := coretestcases.CaseV1{Title: "NewCharHashsetMapDataModelUsing", ExpectedInput: 10, ActualInput: dm.EachHashsetCapacity}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetsCollectionDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetsCollectionDataModel_NewUsingDataModel(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{hs1}}
	hsc := corestr.NewHashsetsCollectionUsingDataModel(dm)
	tc := coretestcases.CaseV1{Title: "NewHashsetsCollectionUsingDataModel", ExpectedInput: 1, ActualInput: hsc.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_HashsetsCollectionDataModel_NewDataModelUsing(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	hsc := corestr.New.HashsetsCollection.UsingHashsets(hs1)
	dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
	tc := coretestcases.CaseV1{Title: "NewHashsetsCollectionDataModelUsing", ExpectedInput: 1, ActualInput: len(dm.Items)}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — SimpleStringOnceModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_SimpleStringOnceModel_Fields(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "test", IsInitialize: true}
	tc := coretestcases.CaseV1{Title: "SimpleStringOnceModel Value", ExpectedInput: "test", ActualInput: m.Value}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_SimpleStringOnceModel_IsInit(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "v", IsInitialize: true}
	tc := coretestcases.CaseV1{Title: "SimpleStringOnceModel IsInit", ExpectedInput: true, ActualInput: m.IsInitialize}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — CollectionsOfCollectionModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CollectionsOfCollectionModel_Fields(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	m := corestr.CollectionsOfCollectionModel{Items: []*corestr.Collection{col}}
	tc := coretestcases.CaseV1{Title: "CollectionsOfCollectionModel", ExpectedInput: 1, ActualInput: len(m.Items)}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AnyToString_EmptyString(t *testing.T) {
	result := corestr.AnyToString(false, "")
	tc := coretestcases.CaseV1{Title: "AnyToString empty", ExpectedInput: "", ActualInput: result}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_AnyToString_WithFieldName(t *testing.T) {
	result := corestr.AnyToString(true, "hello")
	tc := coretestcases.CaseV1{Title: "AnyToString with field name", ExpectedInput: true, ActualInput: len(result) > 0}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_AnyToString_WithoutFieldName(t *testing.T) {
	result := corestr.AnyToString(false, 42)
	tc := coretestcases.CaseV1{Title: "AnyToString without field name", ExpectedInput: true, ActualInput: len(result) > 0}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_AnyToString_Pointer(t *testing.T) {
	v := "ptr"
	result := corestr.AnyToString(false, &v)
	tc := coretestcases.CaseV1{Title: "AnyToString pointer", ExpectedInput: true, ActualInput: len(result) > 0}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_InvalidValueStatusNoMessage(t *testing.T) {
	vs := corestr.InvalidValueStatusNoMessage()
	tc := coretestcases.CaseV1{Title: "InvalidValueStatusNoMessage index", ExpectedInput: -1, ActualInput: vs.Index}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_InvalidValueStatus(t *testing.T) {
	vs := corestr.InvalidValueStatus("err msg")
	tc := coretestcases.CaseV1{Title: "InvalidValueStatus", ExpectedInput: -1, ActualInput: vs.Index}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_ValueStatus_Clone(t *testing.T) {
	vs := corestr.InvalidValueStatus("msg")
	cloned := vs.Clone()
	tc := coretestcases.CaseV1{Title: "ValueStatus Clone", ExpectedInput: vs.Index, ActualInput: cloned.Index}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_TextWithLineNumber_HasLineNumber_Valid(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
	tc := coretestcases.CaseV1{Title: "HasLineNumber valid", ExpectedInput: true, ActualInput: twl.HasLineNumber()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_HasLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "HasLineNumber nil", ExpectedInput: false, ActualInput: twl.HasLineNumber()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
	tc := coretestcases.CaseV1{Title: "IsInvalidLineNumber", ExpectedInput: true, ActualInput: twl.IsInvalidLineNumber()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "IsInvalidLineNumber nil", ExpectedInput: true, ActualInput: twl.IsInvalidLineNumber()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_Length(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}
	tc := coretestcases.CaseV1{Title: "Length", ExpectedInput: 3, ActualInput: twl.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_Length_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "Length nil", ExpectedInput: 0, ActualInput: twl.Length()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_True(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := coretestcases.CaseV1{Title: "IsEmpty true", ExpectedInput: true, ActualInput: twl.IsEmpty()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_False(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
	tc := coretestcases.CaseV1{Title: "IsEmpty false", ExpectedInput: false, ActualInput: twl.IsEmpty()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "IsEmpty nil", ExpectedInput: true, ActualInput: twl.IsEmpty()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsEmptyText(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
	tc := coretestcases.CaseV1{Title: "IsEmptyText", ExpectedInput: true, ActualInput: twl.IsEmptyText()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsEmptyText_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := coretestcases.CaseV1{Title: "IsEmptyText nil", ExpectedInput: true, ActualInput: twl.IsEmptyText()}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := coretestcases.CaseV1{Title: "IsEmptyTextLineBoth", ExpectedInput: true, ActualInput: twl.IsEmptyTextLineBoth()}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CloneSlice_Valid(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := corestr.CloneSlice(input)
	tc := coretestcases.CaseV1{Title: "CloneSlice valid", ExpectedInput: 3, ActualInput: len(result)}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_CloneSlice_Nil(t *testing.T) {
	result := corestr.CloneSlice(nil)
	tc := coretestcases.CaseV1{Title: "CloneSlice nil", ExpectedInput: 0, ActualInput: len(result)}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_CloneSliceIf_Clone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(true, input)
	tc := coretestcases.CaseV1{Title: "CloneSliceIf clone", ExpectedInput: 2, ActualInput: len(result)}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_CloneSliceIf_NoClone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(false, input)
	tc := coretestcases.CaseV1{Title: "CloneSliceIf no clone same ref", ExpectedInput: 2, ActualInput: len(result)}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength / AllIndividualsLengthOfSimpleSlices
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AllIndividualStringsOfStringsLength_Valid(t *testing.T) {
	input := [][]string{{"a", "b"}, {"c"}}
	result := corestr.AllIndividualStringsOfStringsLength(input)
	tc := coretestcases.CaseV1{Title: "AllIndividualStringsOfStringsLength", ExpectedInput: 3, ActualInput: result}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	result := corestr.AllIndividualStringsOfStringsLength([][]string{})
	tc := coretestcases.CaseV1{Title: "AllIndividualStringsOfStringsLength empty", ExpectedInput: 0, ActualInput: result}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Valid(t *testing.T) {
	ss1 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
	ss2 := corestr.New.SimpleSlice.Strings([]string{"c"})
	result := corestr.AllIndividualsLengthOfSimpleSlices([]*corestr.SimpleSlice{ss1, ss2})
	tc := coretestcases.CaseV1{Title: "AllIndividualsLengthOfSimpleSlices", ExpectedInput: 3, ActualInput: result}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Empty(t *testing.T) {
	result := corestr.AllIndividualsLengthOfSimpleSlices([]*corestr.SimpleSlice{})
	tc := coretestcases.CaseV1{Title: "AllIndividualsLengthOfSimpleSlices empty", ExpectedInput: 0, ActualInput: result}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Vars — StaticJsonError, LeftRightExpectingLengthMessager
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_StaticJsonError_NotNil(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "StaticJsonError not nil", ExpectedInput: true, ActualInput: corestr.StaticJsonError != nil}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_ExpectingLengthForLeftRight(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "ExpectingLengthForLeftRight", ExpectedInput: 2, ActualInput: corestr.ExpectingLengthForLeftRight}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_LeftRightExpectingLengthMessager_NotNil(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "LeftRightExpectingLengthMessager not nil", ExpectedInput: true, ActualInput: corestr.LeftRightExpectingLengthMessager != nil}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Funcs types — ReturningBool, filter types
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_ReturningBool_Fields(t *testing.T) {
	rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}
	tc := coretestcases.CaseV1{Title: "ReturningBool IsBreak", ExpectedInput: true, ActualInput: rb.IsBreak}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_LinkedCollectionFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedCollectionFilterResult{IsKeep: true, IsBreak: false}
	tc := coretestcases.CaseV1{Title: "LinkedCollectionFilterResult", ExpectedInput: true, ActualInput: r.IsKeep}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_LinkedListFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedListFilterResult{IsKeep: false, IsBreak: true}
	tc := coretestcases.CaseV1{Title: "LinkedListFilterResult", ExpectedInput: true, ActualInput: r.IsBreak}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_LinkedCollectionFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionFilterParameter{Index: 5}
	tc := coretestcases.CaseV1{Title: "LinkedCollectionFilterParameter", ExpectedInput: 5, ActualInput: p.Index}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_LinkedListFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedListFilterParameter{Index: 3}
	tc := coretestcases.CaseV1{Title: "LinkedListFilterParameter", ExpectedInput: 3, ActualInput: p.Index}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_LinkedListProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedListProcessorParameter{Index: 0, IsFirstIndex: true, IsEndingIndex: false}
	tc := coretestcases.CaseV1{Title: "LinkedListProcessorParameter", ExpectedInput: true, ActualInput: p.IsFirstIndex}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_LinkedCollectionProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionProcessorParameter{Index: 1, IsFirstIndex: false, IsEndingIndex: true}
	tc := coretestcases.CaseV1{Title: "LinkedCollectionProcessorParameter", ExpectedInput: true, ActualInput: p.IsEndingIndex}
	tc.ShouldBeEqual(t, 0)
}

// ═══════════════════════════════════════════════════════════════
// Consts — RegularCollectionEfficiencyLimit, DoubleLimit, NoElements
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_RegularCollectionEfficiencyLimit(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "RegularCollectionEfficiencyLimit", ExpectedInput: 1000, ActualInput: corestr.RegularCollectionEfficiencyLimit}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_DoubleLimit(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "DoubleLimit", ExpectedInput: 3000, ActualInput: corestr.DoubleLimit}
	tc.ShouldBeEqual(t, 0)
}

func Test_Cov43_NoElements(t *testing.T) {
	tc := coretestcases.CaseV1{Title: "NoElements", ExpectedInput: " {No Element}", ActualInput: corestr.NoElements}
	tc.ShouldBeEqual(t, 0)
}
