package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// Reverted from broken ShouldBeEqualMapFirst pattern to working caseV1Compat pattern.
// See issues/full-126-failures-root-cause-analysis.md (Category 1)

// ═══════════════════════════════════════════════════════════════
// newHashmapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashmapCreator_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	tc := caseV1Compat{Name: "Empty hashmap", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_Cap(t *testing.T) {
	hm := corestr.New.Hashmap.Cap(10)
	tc := caseV1Compat{Name: "Cap hashmap empty", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Valid(t *testing.T) {
	pair := corestr.KeyAnyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyAnyValues(pair)
	tc := caseV1Compat{Name: "KeyAnyValues valid", Expected: true, Actual: hm.Has("k1")}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyAnyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyAnyValues()
	tc := caseV1Compat{Name: "KeyAnyValues empty", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValues_Valid(t *testing.T) {
	pair := corestr.KeyValuePair{Key: "k1", Value: "v1"}
	hm := corestr.New.Hashmap.KeyValues(pair)
	tc := caseV1Compat{Name: "KeyValues valid", Expected: true, Actual: hm.Has("k1")}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValues_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValues()
	tc := caseV1Compat{Name: "KeyValues empty", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_Valid(t *testing.T) {
	keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
	vals := corestr.New.Collection.Strings([]string{"v1", "v2"})
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := caseV1Compat{Name: "KeyValuesCollection valid", Expected: 2, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_NilKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesCollection(nil, nil)
	tc := caseV1Compat{Name: "KeyValuesCollection nil", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesCollection_EmptyKeys(t *testing.T) {
	keys := corestr.New.Collection.Empty()
	vals := corestr.New.Collection.Empty()
	hm := corestr.New.Hashmap.KeyValuesCollection(keys, vals)
	tc := caseV1Compat{Name: "KeyValuesCollection empty keys", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_Valid(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{"a", "b"}, []string{"1", "2"})
	tc := caseV1Compat{Name: "KeyValuesStrings valid", Expected: 2, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_KeyValuesStrings_EmptyKeys(t *testing.T) {
	hm := corestr.New.Hashmap.KeyValuesStrings([]string{}, []string{})
	tc := caseV1Compat{Name: "KeyValuesStrings empty", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMap(t *testing.T) {
	m := map[string]string{"x": "y"}
	hm := corestr.New.Hashmap.UsingMap(m)
	val, _ := hm.Get("x")
	tc := caseV1Compat{Name: "UsingMap", Expected: "y", Actual: val}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_Clone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, m)
	tc := caseV1Compat{Name: "UsingMapOptions clone", Expected: true, Actual: hm.Has("a")}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_NoClone(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.UsingMapOptions(false, 0, m)
	tc := caseV1Compat{Name: "UsingMapOptions no clone", Expected: true, Actual: hm.Has("a")}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_UsingMapOptions_EmptyMap(t *testing.T) {
	hm := corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{})
	tc := caseV1Compat{Name: "UsingMapOptions empty", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Valid(t *testing.T) {
	m := map[string]string{"a": "1", "b": "2"}
	hm := corestr.New.Hashmap.MapWithCap(5, m)
	tc := caseV1Compat{Name: "MapWithCap valid", Expected: 2, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_Empty(t *testing.T) {
	hm := corestr.New.Hashmap.MapWithCap(5, map[string]string{})
	tc := caseV1Compat{Name: "MapWithCap empty", Expected: 0, Actual: hm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashmapCreator_MapWithCap_ZeroCap(t *testing.T) {
	m := map[string]string{"a": "1"}
	hm := corestr.New.Hashmap.MapWithCap(0, m)
	tc := caseV1Compat{Name: "MapWithCap zero cap", Expected: true, Actual: hm.Has("a")}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// newHashsetCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewHashsetCreator_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Empty()
	tc := caseV1Compat{Name: "Empty hashset", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_Cap(t *testing.T) {
	hs := corestr.New.Hashset.Cap(10)
	tc := caseV1Compat{Name: "Cap hashset", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_Strings_Valid(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	tc := caseV1Compat{Name: "Strings valid", Expected: 3, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_Strings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{})
	tc := caseV1Compat{Name: "Strings empty", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x", "y")
	tc := caseV1Compat{Name: "StringsSpreadItems", Expected: 2, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsSpreadItems_Empty(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems()
	tc := caseV1Compat{Name: "StringsSpreadItems empty", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsOption_ValidNoClone(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false, "a", "b")
	tc := caseV1Compat{Name: "StringsOption valid", Expected: 2, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilZeroCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(0, false)
	tc := caseV1Compat{Name: "StringsOption nil zero cap", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_StringsOption_NilWithCap(t *testing.T) {
	hs := corestr.New.Hashset.StringsOption(5, false)
	tc := caseV1Compat{Name: "StringsOption nil with cap", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Valid(t *testing.T) {
	a, b := "a", "b"
	hs := corestr.New.Hashset.PointerStrings([]*string{&a, &b})
	tc := caseV1Compat{Name: "PointerStrings valid", Expected: 2, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStrings_Empty(t *testing.T) {
	hs := corestr.New.Hashset.PointerStrings([]*string{})
	tc := caseV1Compat{Name: "PointerStrings empty", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Valid(t *testing.T) {
	a := "a"
	arr := []*string{&a}
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, true, &arr)
	tc := caseV1Compat{Name: "PointerStringsPtrOption valid clone", Expected: 1, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_PointerStringsPtrOption_Nil(t *testing.T) {
	hs := corestr.New.Hashset.PointerStringsPtrOption(5, false, nil)
	tc := caseV1Compat{Name: "PointerStringsPtrOption nil", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Valid(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a", "b"})
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := caseV1Compat{Name: "UsingCollection valid", Expected: 2, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Nil(t *testing.T) {
	hs := corestr.New.Hashset.UsingCollection(nil)
	tc := caseV1Compat{Name: "UsingCollection nil", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingCollection_Empty(t *testing.T) {
	col := corestr.New.Collection.Empty()
	hs := corestr.New.Hashset.UsingCollection(col)
	tc := caseV1Compat{Name: "UsingCollection empty", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Valid(t *testing.T) {
	ss := corestr.New.SimpleSlice.Strings([]string{"x", "y"})
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := caseV1Compat{Name: "SimpleSlice valid", Expected: 2, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_SimpleSlice_Empty(t *testing.T) {
	ss := corestr.New.SimpleSlice.Empty()
	hs := corestr.New.Hashset.SimpleSlice(ss)
	tc := caseV1Compat{Name: "SimpleSlice empty", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMap_Valid(t *testing.T) {
	m := map[string]bool{"a": true, "b": true}
	hs := corestr.New.Hashset.UsingMap(m)
	tc := caseV1Compat{Name: "UsingMap valid", Expected: 2, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMap_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMap(map[string]bool{})
	tc := caseV1Compat{Name: "UsingMap empty", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Clone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(5, true, m)
	tc := caseV1Compat{Name: "UsingMapOption clone", Expected: 1, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_NoClone(t *testing.T) {
	m := map[string]bool{"a": true}
	hs := corestr.New.Hashset.UsingMapOption(0, false, m)
	tc := caseV1Compat{Name: "UsingMapOption no clone", Expected: 1, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewHashsetCreator_UsingMapOption_Empty(t *testing.T) {
	hs := corestr.New.Hashset.UsingMapOption(5, true, map[string]bool{})
	tc := caseV1Compat{Name: "UsingMapOption empty", Expected: 0, Actual: hs.Length()}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// newSimpleStringOnceCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewSimpleStringOnceCreator_Any_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(false, "hello", true)
	tc := caseV1Compat{Name: "Any init", Expected: true, Actual: sso.IsInitialized()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Any_Uninit(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Any(true, 42, false)
	tc := caseV1Compat{Name: "Any uninit", Expected: false, Actual: sso.IsInitialized()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Uninitialized(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Uninitialized("test")
	tc := caseV1Compat{Name: "Uninitialized", Expected: false, Actual: sso.IsInitialized()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Init(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Init("val")
	tc := caseV1Compat{Name: "Init", Expected: "val", Actual: sso.Value()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_InitPtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.InitPtr("pval")
	tc := caseV1Compat{Name: "InitPtr", Expected: true, Actual: sso.IsInitialized()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Create(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Create("cv", true)
	tc := caseV1Compat{Name: "Create", Expected: "cv", Actual: sso.Value()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_CreatePtr(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.CreatePtr("cpv", false)
	tc := caseV1Compat{Name: "CreatePtr uninit", Expected: false, Actual: sso.IsInitialized()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewSimpleStringOnceCreator_Empty(t *testing.T) {
	sso := corestr.New.SimpleStringOnce.Empty()
	tc := caseV1Compat{Name: "Empty SSO", Expected: true, Actual: sso.IsEmpty()}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// newCharHashsetMapCreator
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_NewCharHashsetMapCreator_Cap(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(20, 10)
	tc := caseV1Compat{Name: "Cap", Expected: 0, Actual: chm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_Cap_BelowLimit(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(1, 1)
	tc := caseV1Compat{Name: "Cap below limit", Expected: 0, Actual: chm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_CapItems(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "apple", "banana")
	tc := caseV1Compat{Name: "CapItems", Expected: 2, Actual: chm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Valid(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, []string{"alpha", "beta"})
	tc := caseV1Compat{Name: "Strings valid", Expected: 2, Actual: chm.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NewCharHashsetMapCreator_Strings_Nil(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Strings(10, nil)
	tc := caseV1Compat{Name: "Strings nil", Expected: 0, Actual: chm.Length()}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashmapDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashmapDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashmapDataModel{Items: map[string]string{"k": "v"}}
	hm := corestr.NewHashmapUsingDataModel(dm)
	val, _ := hm.Get("k")
	tc := caseV1Compat{Name: "NewHashmapUsingDataModel", Expected: "v", Actual: val}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_HashmapDataModel_NewDataModelUsing(t *testing.T) {
	hm := corestr.New.Hashmap.Empty()
	hm.AddOrUpdate("x", "y")
	dm := corestr.NewHashmapsDataModelUsing(hm)
	tc := caseV1Compat{Name: "NewHashmapsDataModelUsing", Expected: "y", Actual: dm.Items["x"]}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetDataModel_NewUsingDataModel(t *testing.T) {
	dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
	hs := corestr.NewHashsetUsingDataModel(dm)
	tc := caseV1Compat{Name: "NewHashsetUsingDataModel", Expected: true, Actual: hs.Has("a")}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_HashsetDataModel_NewDataModelUsing(t *testing.T) {
	hs := corestr.New.Hashset.StringsSpreadItems("x")
	dm := corestr.NewHashsetsDataModelUsing(hs)
	tc := caseV1Compat{Name: "NewHashsetsDataModelUsing", Expected: true, Actual: dm.Items["x"]}
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
	tc := caseV1Compat{Name: "NewCharHashsetMapUsingDataModel", Expected: true, Actual: chm.Has("apple")}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CharHashsetDataModel_NewDataModelUsing(t *testing.T) {
	chm := corestr.New.CharHashsetMap.CapItems(10, 10, "banana")
	dm := corestr.NewCharHashsetMapDataModelUsing(chm)
	tc := caseV1Compat{Name: "NewCharHashsetMapDataModelUsing", Expected: 10, Actual: dm.EachHashsetCapacity}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — HashsetsCollectionDataModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_HashsetsCollectionDataModel_NewUsingDataModel(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{hs1}}
	hsc := corestr.NewHashsetsCollectionUsingDataModel(dm)
	tc := caseV1Compat{Name: "NewHashsetsCollectionUsingDataModel", Expected: 1, Actual: hsc.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_HashsetsCollectionDataModel_NewDataModelUsing(t *testing.T) {
	hs1 := corestr.New.Hashset.StringsSpreadItems("a")
	hsc := corestr.New.HashsetsCollection.UsingHashsetsPointers(hs1)
	dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
	tc := caseV1Compat{Name: "NewHashsetsCollectionDataModelUsing", Expected: 1, Actual: len(dm.Items)}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — SimpleStringOnceModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_SimpleStringOnceModel_Fields(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "test", IsInitialize: true}
	tc := caseV1Compat{Name: "SimpleStringOnceModel Value", Expected: "test", Actual: m.Value}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_SimpleStringOnceModel_IsInit(t *testing.T) {
	m := corestr.SimpleStringOnceModel{Value: "v", IsInitialize: true}
	tc := caseV1Compat{Name: "SimpleStringOnceModel IsInit", Expected: true, Actual: m.IsInitialize}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Data Models — CollectionsOfCollectionModel
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CollectionsOfCollectionModel_Fields(t *testing.T) {
	col := corestr.New.Collection.Strings([]string{"a"})
	m := corestr.CollectionsOfCollectionModel{Items: []*corestr.Collection{col}}
	tc := caseV1Compat{Name: "CollectionsOfCollectionModel", Expected: 1, Actual: len(m.Items)}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AnyToString_EmptyString(t *testing.T) {
	result := corestr.AnyToString(false, "")
	tc := caseV1Compat{Name: "AnyToString empty", Expected: "", Actual: result}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AnyToString_WithFieldName(t *testing.T) {
	result := corestr.AnyToString(true, "hello")
	tc := caseV1Compat{Name: "AnyToString with field name", Expected: true, Actual: len(result) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AnyToString_WithoutFieldName(t *testing.T) {
	result := corestr.AnyToString(false, 42)
	tc := caseV1Compat{Name: "AnyToString without field name", Expected: true, Actual: len(result) > 0}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AnyToString_Pointer(t *testing.T) {
	v := "ptr"
	result := corestr.AnyToString(false, &v)
	tc := caseV1Compat{Name: "AnyToString pointer", Expected: true, Actual: len(result) > 0}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_InvalidValueStatusNoMessage(t *testing.T) {
	vs := corestr.InvalidValueStatusNoMessage()
	tc := caseV1Compat{Name: "InvalidValueStatusNoMessage index", Expected: -1, Actual: vs.Index}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_InvalidValueStatus(t *testing.T) {
	vs := corestr.InvalidValueStatus("err msg")
	tc := caseV1Compat{Name: "InvalidValueStatus", Expected: -1, Actual: vs.Index}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_ValueStatus_Clone(t *testing.T) {
	vs := corestr.InvalidValueStatus("msg")
	cloned := vs.Clone()
	tc := caseV1Compat{Name: "ValueStatus Clone", Expected: vs.Index, Actual: cloned.Index}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_TextWithLineNumber_HasLineNumber_Valid(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
	tc := caseV1Compat{Name: "HasLineNumber valid", Expected: true, Actual: twl.HasLineNumber()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_HasLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := caseV1Compat{Name: "HasLineNumber nil", Expected: false, Actual: twl.HasLineNumber()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
	tc := caseV1Compat{Name: "IsInvalidLineNumber", Expected: true, Actual: twl.IsInvalidLineNumber()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := caseV1Compat{Name: "IsInvalidLineNumber nil", Expected: true, Actual: twl.IsInvalidLineNumber()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_Length(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}
	tc := caseV1Compat{Name: "Length", Expected: 3, Actual: twl.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_Length_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := caseV1Compat{Name: "Length nil", Expected: 0, Actual: twl.Length()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_True(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := caseV1Compat{Name: "IsEmpty true", Expected: true, Actual: twl.IsEmpty()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_False(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
	tc := caseV1Compat{Name: "IsEmpty false", Expected: false, Actual: twl.IsEmpty()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := caseV1Compat{Name: "IsEmpty nil", Expected: true, Actual: twl.IsEmpty()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmptyText(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
	tc := caseV1Compat{Name: "IsEmptyText", Expected: true, Actual: twl.IsEmptyText()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmptyText_Nil(t *testing.T) {
	var twl *corestr.TextWithLineNumber
	tc := caseV1Compat{Name: "IsEmptyText nil", Expected: true, Actual: twl.IsEmptyText()}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	tc := caseV1Compat{Name: "IsEmptyTextLineBoth", Expected: true, Actual: twl.IsEmptyTextLineBoth()}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_CloneSlice_Valid(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := corestr.CloneSlice(input)
	tc := caseV1Compat{Name: "CloneSlice valid", Expected: 3, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CloneSlice_Nil(t *testing.T) {
	result := corestr.CloneSlice(nil)
	tc := caseV1Compat{Name: "CloneSlice nil", Expected: 0, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CloneSliceIf_Clone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(true, input...)
	tc := caseV1Compat{Name: "CloneSliceIf clone", Expected: 2, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_CloneSliceIf_NoClone(t *testing.T) {
	input := []string{"a", "b"}
	result := corestr.CloneSliceIf(false, input...)
	tc := caseV1Compat{Name: "CloneSliceIf no clone same ref", Expected: 2, Actual: len(result)}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength / AllIndividualsLengthOfSimpleSlices
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_AllIndividualStringsOfStringsLength_Valid(t *testing.T) {
	input := [][]string{{"a", "b"}, {"c"}}
	result := corestr.AllIndividualStringsOfStringsLength(&input)
	tc := caseV1Compat{Name: "AllIndividualStringsOfStringsLength", Expected: 3, Actual: result}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	input := [][]string{}
	result := corestr.AllIndividualStringsOfStringsLength(&input)
	tc := caseV1Compat{Name: "AllIndividualStringsOfStringsLength empty", Expected: 0, Actual: result}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Valid(t *testing.T) {
	ss1 := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
	ss2 := corestr.New.SimpleSlice.Strings([]string{"c"})
	result := corestr.AllIndividualsLengthOfSimpleSlices(ss1, ss2)
	tc := caseV1Compat{Name: "AllIndividualsLengthOfSimpleSlices", Expected: 3, Actual: result}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_AllIndividualsLengthOfSimpleSlices_Empty(t *testing.T) {
	result := corestr.AllIndividualsLengthOfSimpleSlices()
	tc := caseV1Compat{Name: "AllIndividualsLengthOfSimpleSlices empty", Expected: 0, Actual: result}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Vars — StaticJsonError, LeftRightExpectingLengthMessager
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_StaticJsonError_NotNil(t *testing.T) {
	tc := caseV1Compat{Name: "StaticJsonError not nil", Expected: true, Actual: corestr.StaticJsonError != nil}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_ExpectingLengthForLeftRight(t *testing.T) {
	tc := caseV1Compat{Name: "ExpectingLengthForLeftRight", Expected: 2, Actual: corestr.ExpectingLengthForLeftRight}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LeftRightExpectingLengthMessager_NotNil(t *testing.T) {
	tc := caseV1Compat{Name: "LeftRightExpectingLengthMessager not nil", Expected: true, Actual: corestr.LeftRightExpectingLengthMessager != nil}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Funcs types — ReturningBool, filter types
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_ReturningBool_Fields(t *testing.T) {
	rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}
	tc := caseV1Compat{Name: "ReturningBool IsBreak", Expected: true, Actual: rb.IsBreak}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedCollectionFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedCollectionFilterResult{IsKeep: true, IsBreak: false}
	tc := caseV1Compat{Name: "LinkedCollectionFilterResult", Expected: true, Actual: r.IsKeep}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedListFilterResult_Fields(t *testing.T) {
	r := corestr.LinkedListFilterResult{IsKeep: false, IsBreak: true}
	tc := caseV1Compat{Name: "LinkedListFilterResult", Expected: true, Actual: r.IsBreak}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedCollectionFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionFilterParameter{Index: 5}
	tc := caseV1Compat{Name: "LinkedCollectionFilterParameter", Expected: 5, Actual: p.Index}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedListFilterParameter_Fields(t *testing.T) {
	p := corestr.LinkedListFilterParameter{Index: 3}
	tc := caseV1Compat{Name: "LinkedListFilterParameter", Expected: 3, Actual: p.Index}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedListProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedListProcessorParameter{Index: 0, IsFirstIndex: true, IsEndingIndex: false}
	tc := caseV1Compat{Name: "LinkedListProcessorParameter", Expected: true, Actual: p.IsFirstIndex}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_LinkedCollectionProcessorParameter_Fields(t *testing.T) {
	p := corestr.LinkedCollectionProcessorParameter{Index: 1, IsFirstIndex: false, IsEndingIndex: true}
	tc := caseV1Compat{Name: "LinkedCollectionProcessorParameter", Expected: true, Actual: p.IsEndingIndex}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// Consts — RegularCollectionEfficiencyLimit, DoubleLimit, NoElements
// ═══════════════════════════════════════════════════════════════

func Test_Cov43_RegularCollectionEfficiencyLimit(t *testing.T) {
	tc := caseV1Compat{Name: "RegularCollectionEfficiencyLimit", Expected: 1000, Actual: corestr.RegularCollectionEfficiencyLimit}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_DoubleLimit(t *testing.T) {
	tc := caseV1Compat{Name: "DoubleLimit", Expected: 3000, Actual: corestr.DoubleLimit}
	tc.ShouldBeEqual(t)
}

func Test_Cov43_NoElements(t *testing.T) {
	tc := caseV1Compat{Name: "NoElements", Expected: 0, Actual: corestr.NoElements}
	tc.ShouldBeEqual(t)
}
