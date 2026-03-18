package corestr

import (
	"sync"
	"testing"
)

func TestHashset_Basic(t *testing.T) {
	h := New.Hashset.Empty()
	if !h.IsEmpty() || h.HasItems() || h.HasAnyItem() { t.Fatal("expected empty") }
	if h.Length() != 0 { t.Fatal("expected 0") }
	var nilH *Hashset
	if nilH.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashset_Add(t *testing.T) {
	h := New.Hashset.Cap(5)
	h.Add("a").Add("b")
	if h.Length() != 2 { t.Fatal("expected 2") }
	h.AddLock("c")
	if h.Length() != 3 { t.Fatal("expected 3") }
	h.AddNonEmpty("")
	h.AddNonEmpty("d")
	h.AddNonEmptyWhitespace("   ")
	h.AddNonEmptyWhitespace("e")
	h.AddIf(false, "skip")
	h.AddIf(true, "f")
	h.AddIfMany(false, "x", "y")
	h.AddIfMany(true, "g", "h")
	h.AddFunc(func() string { return "i" })
	h.AddFuncErr(func() (string, error) { return "j", nil }, func(e error) {})
}

func TestHashset_AddBool(t *testing.T) {
	h := New.Hashset.Empty()
	existed := h.AddBool("a")
	if existed { t.Fatal("expected new") }
	existed2 := h.AddBool("a")
	if !existed2 { t.Fatal("expected existed") }
}

func TestHashset_AddPtr(t *testing.T) {
	h := New.Hashset.Empty()
	s := "hello"
	h.AddPtr(&s)
	h.AddPtrLock(&s)
	if h.Length() != 1 { t.Fatal("expected 1 (deduplicated)") }
}

func TestHashset_Adds(t *testing.T) {
	h := New.Hashset.Empty()
	h.Adds("a", "b")
	h.Adds()
	h.AddStrings([]string{"c"})
	h.AddStrings(nil)
	h.AddStringsLock([]string{"d"})
	h.AddStringsLock(nil)
	h.AddCollection(New.Collection.Strings([]string{"e"}))
	h.AddCollection(nil)
	h.AddCollections(New.Collection.Strings([]string{"f"}))
	h.AddCollections()
	ss := New.SimpleSlice.Lines("g")
	h.AddSimpleSlice(ss)
}

func TestHashset_Has(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a", "b", "c")
	if !h.Has("a") || !h.Contains("a") { t.Fatal("expected") }
	if !h.HasLock("a") || !h.HasWithLock("a") { t.Fatal("expected") }
	if h.IsMissing("a") || !h.IsMissing("z") { t.Fatal("unexpected") }
	if h.IsMissingLock("a") { t.Fatal("unexpected") }
	if !h.HasAll("a", "b") || !h.HasAllStrings([]string{"a", "b"}) { t.Fatal("expected") }
	if !h.HasAny("a", "z") || h.HasAny("x", "z") { t.Fatal("unexpected") }
	if !h.IsAllMissing("x", "z") || h.IsAllMissing("a") { t.Fatal("unexpected") }
	if !h.HasAllCollectionItems(New.Collection.Strings([]string{"a"})) { t.Fatal("expected") }
	if h.HasAllCollectionItems(nil) { t.Fatal("expected false") }
}

func TestHashset_IsEquals(t *testing.T) {
	h1 := New.Hashset.StringsSpreadItems("a", "b")
	h2 := New.Hashset.StringsSpreadItems("a", "b")
	if !h1.IsEquals(h2) || !h1.IsEqual(h2) { t.Fatal("expected equal") }
	if !h1.IsEqualsLock(h2) { t.Fatal("expected equal") }
}

func TestHashset_Remove(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a", "b")
	h.Remove("a")
	h.SafeRemove("b")
	h.SafeRemove("z")
	h.RemoveWithLock("z")
}

func TestHashset_List(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a")
	_ = h.List()
	_ = h.ListPtr()
	_ = h.Lines()
	_ = h.SafeStrings()
	_ = h.ListPtrSortedAsc()
	_ = h.ListPtrSortedDsc()
	_ = h.OrderedList()
	_ = h.SortedList()
	_ = h.ListCopyLock()
	_ = h.SimpleSlice()
	_ = h.Items()
	_ = h.Collection()
	_ = h.MapStringAny()
	_ = h.MapStringAnyDiff()
}

func TestHashset_Filter(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("abc", "def")
	f := func(s string) bool { return s == "abc" }
	r := h.Filter(f)
	if r.Length() != 1 { t.Fatal("expected 1") }
	sf := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }
	items := h.GetFilteredItems(sf)
	if len(items) != 1 { t.Fatal("expected 1") }
	col := h.GetFilteredCollection(sf)
	if col.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_GetAllExcept(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a", "b", "c")
	r := h.GetAllExceptHashset(New.Hashset.StringsSpreadItems("a"))
	if len(r) != 2 { t.Fatal("expected 2") }
	r2 := h.GetAllExcept([]string{"a"})
	if len(r2) != 2 { t.Fatal("expected 2") }
	r3 := h.GetAllExceptSpread("a")
	if len(r3) != 2 { t.Fatal("expected 2") }
	r4 := h.GetAllExceptCollection(New.Collection.Strings([]string{"a"}))
	if len(r4) != 2 { t.Fatal("expected 2") }
	_ = h.GetAllExceptHashset(nil)
	_ = h.GetAllExcept(nil)
	_ = h.GetAllExceptSpread()
	_ = h.GetAllExceptCollection(nil)
}

func TestHashset_Resize(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a")
	h.Resize(100)
	h.ResizeLock(200)
	h.AddCapacities(10, 20)
	h.AddCapacitiesLock(10)
	h.AddCapacities()
	h.AddCapacitiesLock()
}

func TestHashset_ConcatNew(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a")
	h2 := New.Hashset.StringsSpreadItems("b")
	r := h.ConcatNewHashsets(false, h2)
	if r.Length() < 2 { t.Fatal("expected >= 2") }
	r2 := h.ConcatNewHashsets(true)
	_ = r2
	r3 := h.ConcatNewStrings(false, []string{"c"})
	_ = r3
	r4 := h.ConcatNewStrings(true)
	_ = r4
}

func TestHashset_StringAndJson(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a")
	if h.String() == "" { t.Fatal("expected non-empty") }
	if h.StringLock() == "" { t.Fatal("expected non-empty") }
	_ = h.Join(",")
	_ = h.NonEmptyJoins(",")
	_ = h.NonWhitespaceJoins(",")
	_ = h.JoinSorted(",")
	_ = h.JsonModel()
	_ = h.JsonModelAny()
	_, _ = h.MarshalJSON()
	_ = h.AsJsoner()
	_ = h.AsJsonContractsBinder()
	_ = h.AsJsonParseSelfInjector()
	_ = h.AsJsonMarshaller()
}

func TestHashset_ToLowerSet(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("ABC")
	lower := h.ToLowerSet()
	if !lower.Has("abc") { t.Fatal("expected lowercase") }
}

func TestHashset_ClearDispose(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a")
	h.Clear()
	h.Dispose()
	var nilH *Hashset
	nilH.Dispose()
}

func TestHashset_DistinctDiff(t *testing.T) {
	h := New.Hashset.StringsSpreadItems("a", "b")
	r := h.DistinctDiffLinesRaw("b", "c")
	if len(r) != 2 { t.Fatal("expected 2 (a and c)") }
	r2 := h.DistinctDiffLines("b", "c")
	_ = r2
	r3 := h.DistinctDiffHashset(New.Hashset.StringsSpreadItems("b", "c"))
	_ = r3
	// edge cases
	empty := New.Hashset.Empty()
	_ = empty.DistinctDiffLinesRaw()
	_ = empty.DistinctDiffLinesRaw("a")
	_ = h.DistinctDiffLinesRaw()
	_ = empty.DistinctDiffLines()
	_ = empty.DistinctDiffLines("a")
}

func TestHashset_WgLock(t *testing.T) {
	h := New.Hashset.Cap(10)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	h.AddWithWgLock("a", wg)
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	h.AddStringsPtrWgLock([]string{"b", "c"}, wg2)
}

func TestHashset_AddItemsMap(t *testing.T) {
	h := New.Hashset.Empty()
	h.AddItemsMap(map[string]bool{"a": true, "b": false})
	if h.Length() != 1 { t.Fatal("expected 1") }
	h.AddItemsMap(nil)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	m := map[string]bool{"c": true}
	h.AddItemsMapWgLock(&m, wg)
	h.AddItemsMapWgLock(nil, nil)
}

func TestHashset_AddHashset(t *testing.T) {
	h := New.Hashset.Empty()
	h2 := New.Hashset.StringsSpreadItems("a", "b")
	h.AddHashsetItems(h2)
	h.AddHashsetItems(nil)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	h3 := New.Hashset.StringsSpreadItems("c")
	h.AddHashsetWgLock(h3, wg)
	h.AddHashsetWgLock(nil, nil)
}

func TestHashset_AddsUsingFilter(t *testing.T) {
	h := New.Hashset.Empty()
	f := func(s string, i int) (string, bool, bool) { return s, true, false }
	h.AddsUsingFilter(f, "a", "b")
	h.AddsUsingFilter(f)
	h.AddsAnyUsingFilter(f, "c")
	h.AddsAnyUsingFilter(f)
	h.AddsAnyUsingFilterLock(f, "d")
	h.AddsAnyUsingFilterLock(f)
}

func TestHashset_EmptyString(t *testing.T) {
	h := New.Hashset.Empty()
	if h.String() == "" { t.Fatal("expected non-empty") }
	if h.StringLock() == "" { t.Fatal("expected non-empty") }
}
