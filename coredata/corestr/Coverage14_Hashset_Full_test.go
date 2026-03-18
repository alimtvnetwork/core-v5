package corestr

import (
	"sync"
	"testing"
)

// ── Hashset comprehensive ──

func TestHashset_AddCapacitiesLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddCapacitiesLock(10)
	hs.Add("a")
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddCapacities(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddCapacities(10, 5)
	hs.Add("a")
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_Resize(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.Add("a")
	hs.Resize(20)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_ResizeLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.Add("a")
	hs.ResizeLock(20)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_Collection(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	c := hs.Collection()
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashset_IsEmptyLock(t *testing.T) {
	hs := New.Hashset.Cap(0)
	if !hs.IsEmptyLock() { t.Fatal("expected true") }
}

func TestHashset_ConcatNewHashsets(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"b"})
	result := hs1.ConcatNewHashsets(false, hs2)
	if result.Length() < 2 { t.Fatal("expected >= 2") }
}

func TestHashset_ConcatNewHashsets_Empty(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	result := hs.ConcatNewHashsets(true)
	if result.Length() < 1 { t.Fatal("expected >= 1") }
}

func TestHashset_ConcatNewStrings(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	result := hs.ConcatNewStrings(false, []string{"b"})
	if result.Length() < 2 { t.Fatal("expected >= 2") }
}

func TestHashset_ConcatNewStrings_Empty(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	result := hs.ConcatNewStrings(true)
	if result.Length() < 1 { t.Fatal("expected >= 1") }
}

func TestHashset_AddPtr(t *testing.T) {
	hs := New.Hashset.Cap(5)
	s := "a"
	hs.AddPtr(&s)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddPtrLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	s := "a"
	hs.AddPtrLock(&s)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddBool(t *testing.T) {
	hs := New.Hashset.Cap(5)
	existed := hs.AddBool("a")
	if existed { t.Fatal("expected false") }
	existed2 := hs.AddBool("a")
	if !existed2 { t.Fatal("expected true") }
}

func TestHashset_AddNonEmpty(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddNonEmpty("")
	hs.AddNonEmpty("a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddNonEmptyWhitespace(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddNonEmptyWhitespace("  ")
	hs.AddNonEmptyWhitespace("a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddIf(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddIf(true, "a")
	hs.AddIf(false, "b")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddIfMany(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddIfMany(true, "a", "b")
	hs.AddIfMany(false, "c")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddFunc(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddFunc(func() string { return "a" })
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddWithWgLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hs.AddWithWgLock("a", &wg)
	wg.Wait()
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddStrings(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddStrings([]string{"a", "b"})
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddStringsLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddStringsLock([]string{"a"})
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddCollection(t *testing.T) {
	hs := New.Hashset.Cap(5)
	c := New.Collection.Strings([]string{"a"})
	hs.AddCollection(c)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddCollections(t *testing.T) {
	hs := New.Hashset.Cap(5)
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	hs.AddCollections(c1, c2)
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddHashsetItems(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs2 := New.Hashset.Strings([]string{"a"})
	hs.AddHashsetItems(hs2)
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddItemsMap(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddItemsMap(map[string]bool{"a": true, "b": false})
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddLock("a")
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_Adds(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.Adds("a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_HasAnyItem(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasAnyItem() { t.Fatal("expected true") }
}

func TestHashset_IsMissing(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.IsMissing("a") { t.Fatal("expected false") }
	if !hs.IsMissing("b") { t.Fatal("expected true") }
}

func TestHashset_IsMissingLock(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.IsMissingLock("a") { t.Fatal("expected false") }
}

func TestHashset_Contains(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.Contains("a") { t.Fatal("expected true") }
}

func TestHashset_IsEqual(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"a"})
	if !hs1.IsEqual(hs2) { t.Fatal("expected true") }
}

func TestHashset_SortedList(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a", "b"})
	sl := hs.SortedList()
	if sl[0] != "a" { t.Fatal("expected sorted") }
}

func TestHashset_Filter(t *testing.T) {
	hs := New.Hashset.Strings([]string{"abc", "def"})
	filtered := hs.Filter(func(s string) bool { return s == "abc" })
	if filtered.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_HasLock(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasLock("a") { t.Fatal("expected true") }
}

func TestHashset_HasAllStrings(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.HasAllStrings([]string{"a", "b"}) { t.Fatal("expected true") }
}

func TestHashset_HasAllCollectionItems(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	c := New.Collection.Strings([]string{"a", "b"})
	if !hs.HasAllCollectionItems(c) { t.Fatal("expected true") }
}

func TestHashset_HasAll(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.HasAll("a", "b") { t.Fatal("expected true") }
}

func TestHashset_IsAllMissing(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.IsAllMissing("x", "y") { t.Fatal("expected true") }
	if hs.IsAllMissing("a") { t.Fatal("expected false") }
}

func TestHashset_HasAny(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasAny("a", "z") { t.Fatal("expected true") }
	if hs.HasAny("x", "y") { t.Fatal("expected false") }
}

func TestHashset_HasWithLock(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasWithLock("a") { t.Fatal("expected true") }
}

func TestHashset_OrderedList(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a", "b"})
	ol := hs.OrderedList()
	if ol[0] != "a" { t.Fatal("expected sorted") }
}

func TestHashset_SafeStrings(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.SafeStrings()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_Lines(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.Lines()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_SimpleSlice(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	ss := hs.SimpleSlice()
	if ss == nil { t.Fatal("expected non-nil") }
}

func TestHashset_GetFilteredItems(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestHashset_GetFilteredCollection(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	c := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashset_GetAllExceptHashset(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b", "c"})
	except := New.Hashset.Strings([]string{"b"})
	result := hs.GetAllExceptHashset(except)
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestHashset_GetAllExcept(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetAllExcept([]string{"a"})
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashset_GetAllExceptSpread(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetAllExceptSpread("a")
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashset_GetAllExceptCollection(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	c := New.Collection.Strings([]string{"a"})
	result := hs.GetAllExceptCollection(c)
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashset_Items(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.Items()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_List(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.List()) != 1 { t.Fatal("expected 1") }
}

func TestHashset_MapStringAny(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	m := hs.MapStringAny()
	if len(m) != 1 { t.Fatal("expected 1") }
}

func TestHashset_MapStringAnyDiff(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.MapStringAnyDiff()
}

func TestHashset_JoinSorted(t *testing.T) {
	hs := New.Hashset.Strings([]string{"b", "a"})
	s := hs.JoinSorted(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_ListPtrSortedAsc(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a"})
	l := hs.ListPtrSortedAsc()
	if l[0] != "a" { t.Fatal("expected sorted") }
}

func TestHashset_ListPtrSortedDsc(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "c"})
	l := hs.ListPtrSortedDsc()
	if l[0] != "c" { t.Fatal("expected reverse sorted") }
}

func TestHashset_Clear(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Clear()
	if hs.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashset_Dispose(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Dispose()
	if hs.items != nil { t.Fatal("expected nil") }
}

func TestHashset_ListCopyLock(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	l := hs.ListCopyLock()
	if len(l) != 1 { t.Fatal("expected 1") }
}

func TestHashset_ToLowerSet(t *testing.T) {
	hs := New.Hashset.Strings([]string{"ABC"})
	lower := hs.ToLowerSet()
	if !lower.Has("abc") { t.Fatal("expected lowercase") }
}

func TestHashset_LengthLock(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.LengthLock() != 1 { t.Fatal("expected 1") }
}

func TestHashset_IsEquals(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"a"})
	if !hs1.IsEquals(hs2) { t.Fatal("expected true") }
}

func TestHashset_IsEqualsLock(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"a"})
	if !hs1.IsEqualsLock(hs2) { t.Fatal("expected true") }
}

func TestHashset_Remove(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Remove("a")
	if hs.Has("a") { t.Fatal("expected removed") }
}

func TestHashset_SafeRemove(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.SafeRemove("a")
	if hs.Has("a") { t.Fatal("expected removed") }
	hs.SafeRemove("missing") // should not panic
}

func TestHashset_RemoveWithLock(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.RemoveWithLock("a")
	if hs.Has("a") { t.Fatal("expected removed") }
}

func TestHashset_String(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.String() == "" { t.Fatal("expected non-empty") }
}

func TestHashset_StringLock(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestHashset_Join(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.Join(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_NonEmptyJoins(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonEmptyJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_NonWhitespaceJoins(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.NonWhitespaceJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_JsonModel(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	m := hs.JsonModel()
	if len(m) != 1 { t.Fatal("expected 1") }
}

func TestHashset_JsonModelAny(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	_ = hs.JsonModelAny()
}

func TestHashset_MarshalJSON(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	b, err := hs.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashset_UnmarshalJSON(t *testing.T) {
	hs := &Hashset{}
	err := hs.UnmarshalJSON([]byte(`{"a":true}`))
	if err != nil || hs.Length() != 1 { t.Fatal("unexpected") }
}

func TestHashset_Json(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	r := hs.Json()
	if r.HasError() { t.Fatal("unexpected") }
}

func TestHashset_ParseInjectUsingJson(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	jr := hs.Json()
	hs2 := New.Hashset.Cap(5)
	_, err := hs2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
}

func TestHashset_AsJsonContractsBinder(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsonContractsBinder()
}

func TestHashset_AsJsoner(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsoner()
}

func TestHashset_AsJsonMarshaller(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsonMarshaller()
}

func TestHashset_AsJsonParseSelfInjector(t *testing.T) {
	hs := New.Hashset.Cap(0)
	_ = hs.AsJsonParseSelfInjector()
}

func TestHashset_DistinctDiffLinesRaw(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	diff := hs.DistinctDiffLinesRaw("b", "c")
	if len(diff) != 2 { t.Fatal("expected 2") }
}

func TestHashset_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	hs := New.Hashset.Cap(0)
	diff := hs.DistinctDiffLinesRaw()
	if len(diff) != 0 { t.Fatal("expected 0") }
}

func TestHashset_DistinctDiffHashset(t *testing.T) {
	hs1 := New.Hashset.Strings([]string{"a"})
	hs2 := New.Hashset.Strings([]string{"b"})
	diff := hs1.DistinctDiffHashset(hs2)
	if len(diff) != 2 { t.Fatal("expected 2") }
}

func TestHashset_DistinctDiffLines(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	diff := hs.DistinctDiffLines("b")
	if len(diff) != 2 { t.Fatal("expected 2") }
}

func TestHashset_Serialize(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	b, err := hs.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashset_Deserialize(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	var target map[string]bool
	err := hs.Deserialize(&target)
	if err != nil { t.Fatal("unexpected") }
}

func TestHashset_WrapDoubleQuote(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapDoubleQuote()
}

func TestHashset_WrapSingleQuote(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapSingleQuote()
}

func TestHashset_WrapDoubleQuoteIfMissing(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapDoubleQuoteIfMissing()
}

func TestHashset_WrapSingleQuoteIfMissing(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.WrapSingleQuoteIfMissing()
}

func TestHashset_Transpile_Empty(t *testing.T) {
	hs := New.Hashset.Cap(0)
	result := hs.Transpile(func(s string) string { return s })
	if result.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashset_JoinLine(t *testing.T) {
	hs := Hashset{}
	hs.items = map[string]bool{"a": true}
	hs.hasMapUpdated = true
	s := hs.JoinLine()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_AddsUsingFilter(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) {
		return s, true, false
	}, "a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_AddsAnyUsingFilter(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) {
		return s, true, false
	}, "a", nil)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddsAnyUsingFilterLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddsAnyUsingFilterLock(func(s string, i int) (string, bool, bool) {
		return s, true, false
	}, "a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddFuncErr(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs.AddFuncErr(
		func() (string, error) { return "a", nil },
		func(err error) { t.Fatal("unexpected") },
	)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_AddStringsPtrWgLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hs.AddStringsPtrWgLock([]string{"a"}, &wg)
	wg.Wait()
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddHashsetWgLock(t *testing.T) {
	hs := New.Hashset.Cap(5)
	hs2 := New.Hashset.Strings([]string{"a"})
	wg := sync.WaitGroup{}
	wg.Add(1)
	hs.AddHashsetWgLock(hs2, &wg)
	wg.Wait()
	if !hs.Has("a") { t.Fatal("expected true") }
}

func TestHashset_AddSimpleSlice(t *testing.T) {
	hs := New.Hashset.Cap(5)
	ss := New.SimpleSlice.SpreadStrings("a", "b")
	hs.AddSimpleSlice(ss)
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_ListPtr(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if len(hs.ListPtr()) != 1 { t.Fatal("expected 1") }
}
