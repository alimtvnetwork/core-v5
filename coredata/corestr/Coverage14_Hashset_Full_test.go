package corestr

import (
	"sync"
	"testing"
)

// ── Hashset comprehensive ──

func TestHashset_AddCapacitiesLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddCapacitiesLock(10)
	hs.Add("a")
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddCapacities_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddCapacities(10, 5)
	hs.Add("a")
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_Resize_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.Add("a")
	hs.Resize(20)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_ResizeLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.Add("a")
	hs.ResizeLock(20)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_Collection_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	c := hs.Collection()
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashset_IsEmptyLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(0)
	if !hs.IsEmptyLock() { t.Fatal("expected true") }
}

func TestHashset_ConcatNewHashsets_C14(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"b"})
	result := hs1.ConcatNewHashsets(false, hs2)
	if result.Length() < 2 { t.Fatal("expected >= 2") }
}

func TestHashset_ConcatNewHashsets_Empty_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	result := hs.ConcatNewHashsets(true)
	if result.Length() < 1 { t.Fatal("expected >= 1") }
}

func TestHashset_ConcatNewStrings_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	result := hs.ConcatNewStrings(false, []string{"b"})
	if result.Length() < 2 { t.Fatal("expected >= 2") }
}

func TestHashset_ConcatNewStrings_Empty_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	result := hs.ConcatNewStrings(true)
	if result.Length() < 1 { t.Fatal("expected >= 1") }
}

func TestHashset_AddPtr_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	s := "a"
	hs.AddPtr(&s)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddPtrLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	s := "a"
	hs.AddPtrLock(&s)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddBool_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	existed := hs.AddBool("a")
	if existed { t.Fatal("expected false") }
	existed2 := hs.AddBool("a")
	if !existed2 { t.Fatal("expected true") }
}

func TestHashset_AddNonEmpty_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddNonEmpty("")
	hs.AddNonEmpty("a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddNonEmptyWhitespace_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddNonEmptyWhitespace("  ")
	hs.AddNonEmptyWhitespace("a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddIf_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddIf(true, "a")
	hs.AddIf(false, "b")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddIfMany_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddIfMany(true, "a", "b")
	hs.AddIfMany(false, "c")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddFunc_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddFunc(func() string { return "a" })
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddWithWgLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hs.AddWithWgLock("a", &wg)
	wg.Wait()
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddStrings_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddStrings([]string{"a", "b"})
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddStringsLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddStringsLock([]string{"a"})
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddCollection_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	c := New.Collection.Strings([]string{"a"})
	hs.AddCollection(c)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddCollections_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	hs.AddCollections(c1, c2)
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddHashsetItems_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs2 := New.Hashset.Strings([]string{"a"})
	hs.AddHashsetItems(hs2)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddItemsMap_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddItemsMap(map[string]bool{"a": true, "b": false})
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddLock("a")
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_Adds_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.Adds("a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_HasAnyItem_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasAnyItem() { t.Fatal("expected true") }
}

func TestHashset_IsMissing_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.IsMissing("a") { t.Fatal("expected false") }
	if !hs.IsMissing("b") { t.Fatal("expected true") }
}

func TestHashset_IsMissingLock_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.IsMissingLock("a") { t.Fatal("expected false") }
}

func TestHashset_Contains_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.Contains("a") { t.Fatal("expected true") }
}

func TestHashset_IsEqual_C14(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"a"})
	if !hs1.IsEqual(hs2) { t.Fatal("expected true") }
}

func TestHashset_SortedList_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a", "b"})
	sl := hs.SortedList()
	if sl[0] != "a" { t.Fatal("expected sorted") }
}

func TestHashset_Filter_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"abc", "def"})
	filtered := hs.Filter(func(s string) bool { return s == "abc" })
	if filtered.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_HasLock_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasLock("a") { t.Fatal("expected true") }
}

func TestHashset_HasAllStrings_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.HasAllStrings([]string{"a", "b"}) { t.Fatal("expected true") }
}

func TestHashset_HasAllCollectionItems_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	c := New.Collection.Strings([]string{"a", "b"})
	if !hs.HasAllCollectionItems(c) { t.Fatal("expected true") }
}

func TestHashset_HasAll_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.HasAll("a", "b") { t.Fatal("expected true") }
}

func TestHashset_IsAllMissing_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.IsAllMissing("x", "y") { t.Fatal("expected true") }
	if hs.IsAllMissing("a") { t.Fatal("expected false") }
}

func TestHashset_HasAny_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasAny("a", "z") { t.Fatal("expected true") }
	if hs.HasAny("x", "y") { t.Fatal("expected false") }
}

func TestHashset_HasWithLock_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasWithLock("a") { t.Fatal("expected true") }
}

func TestHashset_OrderedList_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a", "b"})
	ol := hs.OrderedList()
	if ol[0] != "a" { t.Fatal("expected sorted") }
}

func TestHashset_SafeStrings_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.SafeStrings()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_Lines_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.Lines()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_SimpleSlice_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	ss := hs.SimpleSlice()
	if ss == nil { t.Fatal("expected non-nil") }
}

func TestHashset_GetFilteredItems_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestHashset_GetFilteredCollection_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	c := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashset_GetAllExceptHashset_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b", "c"})
	except := New.Hashset.Strings([]string{"b"})
	result := hs.GetAllExceptHashset(except)
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestHashset_GetAllExcept_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetAllExcept([]string{"a"})
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashset_GetAllExceptSpread_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetAllExceptSpread("a")
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashset_GetAllExceptCollection_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	c := New.Collection.Strings([]string{"a"})
	result := hs.GetAllExceptCollection(c)
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashset_Items_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.Items()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_List_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.List()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_MapStringAny_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	m := hs.MapStringAny()
	if len(m) != 1 { t.Fatal("expected 1") }
}

func TestHashset_MapStringAnyDiff_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.MapStringAnyDiff()
}

func TestHashset_JoinSorted_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"b", "a"})
	s := hs.JoinSorted(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_ListPtrSortedAsc_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a"})
	l := hs.ListPtrSortedAsc()
	if l[0] != "a" { t.Fatal("expected sorted") }
}

func TestHashset_ListPtrSortedDsc_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "c"})
	l := hs.ListPtrSortedDsc()
	if l[0] != "c" { t.Fatal("expected reverse sorted") }
}

func TestHashset_Clear_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Clear()
	if hs.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashset_Dispose_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Dispose()
	if hs.items != nil { t.Fatal("expected nil") }
}

func TestHashset_ListCopyLock_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	l := hs.ListCopyLock()
	if len(l) != 1 { t.Fatal("expected 1") }
}

func TestHashset_ToLowerSet_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"ABC"})
	lower := hs.ToLowerSet()
	if !lower.Has("abc") { t.Fatal("expected lowercase") }
}

func TestHashset_LengthLock_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.LengthLock() != 1 { t.Fatal("expected 1") }
}

func TestHashset_IsEquals_C14(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"a"})
	if !hs1.IsEquals(hs2) { t.Fatal("expected true") }
}

func TestHashset_IsEqualsLock_C14(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"a"})
	if !hs1.IsEqualsLock(hs2) { t.Fatal("expected true") }
}

func TestHashset_Remove_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Remove("a")
	if hs.Has("a") { t.Fatal("expected removed") }
}

func TestHashset_SafeRemove_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.SafeRemove("a")
	if hs.Has("a") { t.Fatal("expected removed") }
	hs.SafeRemove("missing") // should not panic
}

func TestHashset_RemoveWithLock_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.RemoveWithLock("a")
	if hs.Has("a") { t.Fatal("expected removed") }
}

func TestHashset_String_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.String() == "" { t.Fatal("expected non-empty") }
}

func TestHashset_StringLock_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestHashset_Join_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.Join(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_NonEmptyJoins_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonEmptyJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_NonWhitespaceJoins_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonWhitespaceJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_JsonModel_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	m := hs.JsonModel()
	if len(m) != 1 { t.Fatal("expected 1") }
}

func TestHashset_JsonModelAny_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	_ = hs.JsonModelAny()
}

func TestHashset_MarshalJSON_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	b, err := hs.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashset_UnmarshalJSON_C14(t *testing.T) {
	hs := &Hashset{}
	err := hs.UnmarshalJSON([]byte(`{"a":true}`))
	if err != nil || hs.Length() != 1 { t.Fatal("unexpected") }
}

func TestHashset_Json_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	r := hs.Json()
	if r.HasError() { t.Fatal("unexpected") }
}

func TestHashset_ParseInjectUsingJson_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	jr := hs.Json()
	hs2 := New.Hashset.Cap(5)
	_, err := hs2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
}

func TestHashset_AsJsonContractsBinder_C14(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsonContractsBinder()
}

func TestHashset_AsJsoner_C14(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsoner()
}

func TestHashset_AsJsonMarshaller_C14(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsonMarshaller()
}

func TestHashset_AsJsonParseSelfInjector_C14(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsonParseSelfInjector()
}

func TestHashset_DistinctDiffLinesRaw_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	diff := hs.DistinctDiffLinesRaw("b", "c")
	if len(diff) != 2 { t.Fatal("expected 2") }
}

func TestHashset_DistinctDiffLinesRaw_BothEmpty_C14(t *testing.T) {
	hs := New.Hashset.Cap(0)
	diff := hs.DistinctDiffLinesRaw()
	if len(diff) != 0 { t.Fatal("expected 0") }
}

func TestHashset_DistinctDiffHashset_C14(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"b"})
	diff := hs1.DistinctDiffHashset(hs2)
	if len(diff) != 2 { t.Fatal("expected 2") }
}

func TestHashset_DistinctDiffLines_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	diff := hs.DistinctDiffLines("b")
	if len(diff) != 2 { t.Fatal("expected 2") }
}

func TestHashset_Serialize_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	b, err := hs.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashset_Deserialize_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	var target map[string]bool
	err := hs.Deserialize(&target)
	if err != nil { t.Fatal("unexpected") }
}

func TestHashset_WrapDoubleQuote_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapDoubleQuote()
}

func TestHashset_WrapSingleQuote_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapSingleQuote()
}

func TestHashset_WrapDoubleQuoteIfMissing_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapDoubleQuoteIfMissing()
}

func TestHashset_WrapSingleQuoteIfMissing_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapSingleQuoteIfMissing()
}

func TestHashset_Transpile_Empty_C14(t *testing.T) {
	hs := New.Hashset.Cap(0)
	result := hs.Transpile(func(s string) string { return s })
	if result.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashset_JoinLine_C14(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.JoinLine()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_AddsUsingFilter_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) {
		return s, true, false
	}, "a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddsAnyUsingFilter_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) {
		return s, true, false
	}, "a", nil)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddsAnyUsingFilterLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddsAnyUsingFilterLock(func(s string, i int) (string, bool, bool) {
		return s, true, false
	}, "a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddFuncErr_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddFuncErr(
		func() (string, error) { return "a", nil },
		func(err error) { t.Fatal("unexpected") },
	)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddStringsPtrWgLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hs.AddStringsPtrWgLock([]string{"a"}, &wg)
	wg.Wait()
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddHashsetWgLock_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs2 := New.Hashset.Strings([]string{"a"})
	wg := sync.WaitGroup{}
	wg.Add(1)
	hs.AddHashsetWgLock(hs2, &wg)
	wg.Wait()
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddSimpleSlice_C14(t *testing.T) {
	hs := New.Hashset.Cap(5)
	ss := New.SimpleSlice.SpreadStrings("a", "b")
	hs.AddSimpleSlice(ss)
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_ListPtr_C14(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.ListPtr()) != 1 { t.Fatal("expected 1") }
}
