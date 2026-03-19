package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ── Hashset ──

func Test_C31_Hashset_IsEmpty(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	if !h.IsEmpty() { t.Fatal("expected true") }
}

func Test_C31_Hashset_HasItems(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	if !h.HasItems() { t.Fatal("expected true") }
}

func Test_C31_Hashset_AddCapacities(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	h.AddCapacities(10, 5)
	h.AddCapacities()
}

func Test_C31_Hashset_AddCapacitiesLock(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	h.AddCapacitiesLock(10)
	h.AddCapacitiesLock()
}

func Test_C31_Hashset_Resize(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	h.Resize(10)
	h.Resize(0)
}

func Test_C31_Hashset_ResizeLock(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	h.ResizeLock(10)
	h.ResizeLock(0)
}

func Test_C31_Hashset_Add(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.Add("a")
	if !h.Has("a") { t.Fatal("expected true") }
}

func Test_C31_Hashset_AddBool(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.AddBool("a")
	h.AddBool("a") // second time should return true (exists)
}

func Test_C31_Hashset_AddNonEmpty(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.AddNonEmpty("")
	h.AddNonEmpty("a")
}

func Test_C31_Hashset_AddLock(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.AddLock("a")
}

func Test_C31_Hashset_AddCollection(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	c := corestr.New.Collection.Strings([]string{"a"})
	h.AddCollection(c)
	h.AddCollection(nil)
}

func Test_C31_Hashset_AddHashsetItems(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.AddHashsetItems(corestr.New.Hashset.StringsSpreadItems("a"))
	h.AddHashsetItems(nil)
}

func Test_C31_Hashset_AddHashsetWgLock(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	var wg sync.WaitGroup
	wg.Add(1)
	h.AddHashsetWgLock(corestr.New.Hashset.StringsSpreadItems("a"), &wg)
	wg.Wait()
}

func Test_C31_Hashset_Adds(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.Adds("a", "b")
}

func Test_C31_Hashset_AddStrings(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.AddStrings([]string{"a"})
	h.AddStrings(nil)
}

func Test_C31_Hashset_AddItemsMap(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	h.AddItemsMap(map[string]bool{"a": true})
}

func Test_C31_Hashset_Remove(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "b")
	h.Remove("a")
}

func Test_C31_Hashset_RemoveWithLock(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	h.RemoveWithLock("a")
}

func Test_C31_Hashset_SafeRemove(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "b")
	h.SafeRemove("a")
}

func Test_C31_Hashset_RemovesLock(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "b")
	h.RemoveWithLock("a")
	h.RemoveWithLock("b")
}

func Test_C31_Hashset_Length(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	if h.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C31_Hashset_LengthLock(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.LengthLock()
}

func Test_C31_Hashset_Has(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	if !h.Has("a") { t.Fatal("expected true") }
	if h.Has("z") { t.Fatal("expected false") }
}

func Test_C31_Hashset_HasLock(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.HasLock("a")
}

func Test_C31_Hashset_HasAll(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "b")
	if !h.HasAll("a") { t.Fatal("expected true") }
	if h.HasAll("z") { t.Fatal("expected false") }
}

func Test_C31_Hashset_HasAllStrings(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "b")
	if !h.HasAllStrings([]string{"a"}) { t.Fatal("expected true") }
}

func Test_C31_Hashset_HasAny(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	if !h.HasAny("a", "z") { t.Fatal("expected true") }
}

func Test_C31_Hashset_HasAnyOfStrings(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.HasAny("a", "z")
}

func Test_C31_Hashset_List(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.List()
}

func Test_C31_Hashset_ListLock(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.ListCopyLock()
}

func Test_C31_Hashset_SortedListAsc(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("b", "a")
	_ = h.ListPtrSortedAsc()
	_ = corestr.New.Hashset.Empty().SortedList()
}

func Test_C31_Hashset_SortedListDsc(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "b")
	_ = h.ListPtrSortedDsc()
}

func Test_C31_Hashset_Map(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.MapStringAny()
}

func Test_C31_Hashset_MapStringAnyDiff(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.MapStringAnyDiff()
}

func Test_C31_Hashset_ListPtrSorted(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.ListPtrSortedAsc()
	_ = h.ListPtrSortedDsc()
}

// CopyMapLock, Clone, CloneLock, Diff, DiffLock, SummaryString, SummaryStringLock
// do not exist on Hashset — removed.

func Test_C31_Hashset_MapStringAnyDiff_Coverage(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "b")
	_ = h.MapStringAnyDiff()
}

func Test_C31_Hashset_DistinctDiffHashset(t *testing.T) {
	a := corestr.New.Hashset.StringsSpreadItems("a", "b")
	b := corestr.New.Hashset.StringsSpreadItems("a")
	_ = a.DistinctDiffHashset(b)
}

func Test_C31_Hashset_DistinctDiffLines(t *testing.T) {
	a := corestr.New.Hashset.StringsSpreadItems("a", "b")
	_ = a.DistinctDiffLines([]string{"a"})
}

func Test_C31_Hashset_DistinctDiffLinesRaw(t *testing.T) {
	a := corestr.New.Hashset.StringsSpreadItems("a", "b")
	_ = a.DistinctDiffLinesRaw([]string{"a"})
}

func Test_C31_Hashset_JsonMethods(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_ = h.Json()
	_ = h.JsonPtr()
	_ = h.JsonModel()
	_ = h.JsonModelAny()
	_, _ = h.MarshalJSON()
	_ = h.AsJsonContractsBinder()
	_ = h.AsJsonMarshaller()
	_ = h.AsJsonParseSelfInjector()
	_ = h.AsJsoner()
}

func Test_C31_Hashset_UnmarshalJSON(t *testing.T) {
	h := &corestr.Hashset{}
	_ = h.UnmarshalJSON([]byte(`["a"]`))
}

func Test_C31_Hashset_ParseInjectUsingJson(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	r := corejson.New([]string{"a"})
	_, _ = h.ParseInjectUsingJson(&r)
}

func Test_C31_Hashset_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	h := corestr.New.Hashset.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	h.ParseInjectUsingJsonMust(bad)
}

func Test_C31_Hashset_JsonParseSelfInject(t *testing.T) {
	h := corestr.New.Hashset.Empty()
	r := corejson.New([]string{"a"})
	_ = h.JsonParseSelfInject(&r)
}

func Test_C31_Hashset_Clear(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	h.Clear()
}

// ClearLock and DisposeLock do not exist on Hashset — removed.

func Test_C31_Hashset_Dispose(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	h.Dispose()
}

func Test_C31_Hashset_Filter(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a", "bb")
	_ = h.Filter(func(s string) bool { return len(s) == 1 })
}

func Test_C31_Hashset_Serialize(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	_, _ = h.Serialize()
}

func Test_C31_Hashset_Deserialize(t *testing.T) {
	h := corestr.New.Hashset.StringsSpreadItems("a")
	var target []string
	_ = h.Deserialize(&target)
}

// ── newHashsetCreator ──

func Test_C31_NHC_PointerStrings(t *testing.T) {
	s := "a"
	_ = corestr.New.Hashset.PointerStrings([]*string{&s})
	_ = corestr.New.Hashset.PointerStrings([]*string{})
}

func Test_C31_NHC_PointerStringsPtrOption(t *testing.T) {
	s := "a"
	arr := []*string{&s}
	_ = corestr.New.Hashset.PointerStringsPtrOption(0, false, &arr)
	_ = corestr.New.Hashset.PointerStringsPtrOption(5, false, nil)
}

func Test_C31_NHC_StringsOption(t *testing.T) {
	_ = corestr.New.Hashset.StringsOption(0, false)
	_ = corestr.New.Hashset.StringsOption(5, false)
	_ = corestr.New.Hashset.StringsOption(0, false, "a")
}

func Test_C31_NHC_StringsSpreadItems(t *testing.T) {
	_ = corestr.New.Hashset.StringsSpreadItems("a")
	_ = corestr.New.Hashset.StringsSpreadItems()
}

func Test_C31_NHC_SimpleSlice(t *testing.T) {
	ss := corestr.New.SimpleSlice.Lines("a")
	_ = corestr.New.Hashset.SimpleSlice(ss)
	_ = corestr.New.Hashset.SimpleSlice(corestr.New.SimpleSlice.Empty())
}

// ── Hashmap ──

func Test_C31_Hashmap_IsEmpty(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	if !h.IsEmpty() { t.Fatal("expected true") }
}

func Test_C31_Hashmap_HasItems(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	if !h.HasItems() { t.Fatal("expected true") }
}

func Test_C31_Hashmap_IsEmptyLock(t *testing.T) {
	_ = corestr.New.Hashmap.Empty().IsEmptyLock()
}

func Test_C31_Hashmap_AddOrUpdate(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
}

func Test_C31_Hashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateKeyStrValInt("k", 42)
}

func Test_C31_Hashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateKeyStrValFloat("k", 3.14)
}

func Test_C31_Hashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateKeyStrValFloat64("k", 3.14)
}

func Test_C31_Hashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateKeyStrValAny("k", "v")
}

func Test_C31_Hashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateKeyValueAny(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
}

func Test_C31_Hashmap_AddOrUpdateCollection(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	kvc := corestr.New.KeyValues.Empty()
	kvc.Add("k", "v")
	h.AddOrUpdateCollection(kvc)
	h.AddOrUpdateCollection(nil)
}

func Test_C31_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	other := corestr.New.Hashmap.Empty()
	other.AddOrUpdate("k", "v")
	h.AddOrUpdateHashmap(other)
	h.AddOrUpdateHashmap(nil)
}

func Test_C31_Hashmap_AddOrUpdateKeyValues(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateKeyValues(corestr.KeyValuePair{Key: "k", Value: "v"})
}

func Test_C31_Hashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateKeyAnyValues(corestr.KeyAnyValuePair{Key: "k", Value: "v"})
}

func Test_C31_Hashmap_AddOrUpdateMap(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateMap(map[string]string{"k": "v"})
}

func Test_C31_Hashmap_Get(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.Get("k")
	_ = h.Get("missing")
}

func Test_C31_Hashmap_GetLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.GetLock("k")
}

func Test_C31_Hashmap_Has(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	if !h.Has("k") { t.Fatal("expected true") }
}

func Test_C31_Hashmap_HasLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.HasLock("k")
}

func Test_C31_Hashmap_Length(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	if h.Length() != 1 { t.Fatal("expected 1") }
}

func Test_C31_Hashmap_LengthLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.LengthLock()
}

func Test_C31_Hashmap_Remove(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	h.Remove("k")
}

func Test_C31_Hashmap_RemoveLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	h.RemoveLock("k")
}

func Test_C31_Hashmap_KeysList(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.KeysList()
}

func Test_C31_Hashmap_ValuesList(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.ValuesList()
}

func Test_C31_Hashmap_Collection(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.Collection()
}

func Test_C31_Hashmap_Clone(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.Clone()
}

func Test_C31_Hashmap_CloneLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.CloneLock()
}

func Test_C31_Hashmap_Map(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.Map()
}

func Test_C31_Hashmap_MapLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.MapLock()
}

func Test_C31_Hashmap_CopyMap(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.CopyMap()
}

func Test_C31_Hashmap_CopyMapLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.CopyMapLock()
}

func Test_C31_Hashmap_IsEquals(t *testing.T) {
	a := corestr.New.Hashmap.Empty()
	a.AddOrUpdate("k", "v")
	b := corestr.New.Hashmap.Empty()
	b.AddOrUpdate("k", "v")
	if !a.IsEquals(b) { t.Fatal("expected true") }
}

func Test_C31_Hashmap_IsEqualsLock(t *testing.T) {
	a := corestr.New.Hashmap.Empty()
	b := corestr.New.Hashmap.Empty()
	_ = a.IsEqualsLock(b)
}

func Test_C31_Hashmap_String(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.String()
	h.AddOrUpdate("k", "v")
	_ = h.String()
}

func Test_C31_Hashmap_StringLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.StringLock()
}

func Test_C31_Hashmap_SummaryString(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.SummaryString()
	_ = h.SummaryStringLock()
}

func Test_C31_Hashmap_JsonMethods(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.Json()
	_ = h.JsonPtr()
	_ = h.JsonModel()
	_ = h.JsonModelAny()
	_, _ = h.MarshalJSON()
	_ = h.AsJsonContractsBinder()
	_ = h.AsJsoner()
	_ = h.AsJsonMarshaller()
	_ = h.AsJsonParseSelfInjector()
	_, _ = h.Serialize()
}

func Test_C31_Hashmap_UnmarshalJSON(t *testing.T) {
	h := &corestr.Hashmap{}
	_ = h.UnmarshalJSON([]byte(`{"k":"v"}`))
}

func Test_C31_Hashmap_ParseInjectUsingJson(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	r := corejson.New(map[string]string{"k": "v"})
	_, _ = h.ParseInjectUsingJson(&r)
}

func Test_C31_Hashmap_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	h := corestr.New.Hashmap.Empty()
	bad := corejson.NewResult.UsingString(`invalid`)
	h.ParseInjectUsingJsonMust(bad)
}

func Test_C31_Hashmap_JsonParseSelfInject(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	r := corejson.New(map[string]string{"k": "v"})
	_ = h.JsonParseSelfInject(&r)
}

func Test_C31_Hashmap_Clear(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	h.Clear()
}

func Test_C31_Hashmap_ClearLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.ClearLock()
}

func Test_C31_Hashmap_Dispose(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.Dispose()
}

func Test_C31_Hashmap_DisposeLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.DisposeLock()
}

func Test_C31_Hashmap_Deserialize(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	var target map[string]string
	_ = h.Deserialize(&target)
}

func Test_C31_Hashmap_SortedKeysListAsc(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("b", "1")
	h.AddOrUpdate("a", "2")
	_ = h.SortedKeysListAsc()
}

func Test_C31_Hashmap_SortedValuesListAsc(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.SortedValuesListAsc()
}

func Test_C31_Hashmap_AllKeysSorted(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("b", "1")
	h.AddOrUpdate("a", "2")
	_ = h.AllKeysSorted()
}

func Test_C31_Hashmap_Diff(t *testing.T) {
	a := corestr.New.Hashmap.Empty()
	a.AddOrUpdate("k", "v")
	b := corestr.New.Hashmap.Empty()
	b.AddOrUpdate("k", "v2")
	_ = a.Diff(b)
}

func Test_C31_Hashmap_DiffLock(t *testing.T) {
	a := corestr.New.Hashmap.Empty()
	_ = a.DiffLock(corestr.New.Hashmap.Empty())
}

func Test_C31_Hashmap_DiffMessage(t *testing.T) {
	a := corestr.New.Hashmap.Empty()
	a.AddOrUpdate("k", "v")
	_ = a.DiffMessage(corestr.New.Hashmap.Empty())
}

func Test_C31_Hashmap_AddOrUpdateLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdateLock("k", "v")
}

func Test_C31_Hashmap_GetSafe(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_, _ = h.GetSafe("k")
	h.AddOrUpdate("k", "v")
	_, _ = h.GetSafe("k")
}

func Test_C31_Hashmap_GetSafeLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_, _ = h.GetSafeLock("k")
}

func Test_C31_Hashmap_KeysCollection(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.KeysCollection()
}

func Test_C31_Hashmap_ValuesCollection(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.ValuesCollection()
}

func Test_C31_Hashmap_Filter(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	h.AddOrUpdate("k", "v")
	_ = h.Filter(func(k, v string) bool { return true })
}

func Test_C31_Hashmap_FilterLock(t *testing.T) {
	h := corestr.New.Hashmap.Empty()
	_ = h.FilterLock(func(k, v string) bool { return true })
}

// ── newHashmapCreator ──

func Test_C31_NHM_Empty(t *testing.T)       { _ = corestr.New.Hashmap.Empty() }
func Test_C31_NHM_KeyAnyValues(t *testing.T) { _ = corestr.New.Hashmap.KeyAnyValues() }
func Test_C31_NHM_KeyValues(t *testing.T)    { _ = corestr.New.Hashmap.KeyValues() }
func Test_C31_NHM_KeyValuesCollection(t *testing.T) {
	k := corestr.New.Collection.Strings([]string{"k"})
	v := corestr.New.Collection.Strings([]string{"v"})
	_ = corestr.New.Hashmap.KeyValuesCollection(k, v)
	_ = corestr.New.Hashmap.KeyValuesCollection(nil, nil)
}
func Test_C31_NHM_KeyValuesStrings(t *testing.T) {
	_ = corestr.New.Hashmap.KeyValuesStrings([]string{"k"}, []string{"v"})
	_ = corestr.New.Hashmap.KeyValuesStrings(nil, nil)
}
func Test_C31_NHM_UsingMap(t *testing.T) {
	_ = corestr.New.Hashmap.UsingMap(map[string]string{"k": "v"})
}
func Test_C31_NHM_UsingMapOptions(t *testing.T) {
	_ = corestr.New.Hashmap.UsingMapOptions(true, 5, map[string]string{"k": "v"})
	_ = corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{})
	_ = corestr.New.Hashmap.UsingMapOptions(false, 0, map[string]string{"k": "v"})
}
func Test_C31_NHM_MapWithCap(t *testing.T) {
	_ = corestr.New.Hashmap.MapWithCap(5, map[string]string{"k": "v"})
	_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{"k": "v"})
	_ = corestr.New.Hashmap.MapWithCap(0, map[string]string{})
}
