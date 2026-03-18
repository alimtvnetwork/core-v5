package corestr

import (
	"sync"
	"testing"
)

func TestHashmap_Basic(t *testing.T) {
	h := New.Hashmap.Empty()
	if !h.IsEmpty() || h.HasItems() || h.HasAnyItem() { t.Fatal("expected empty") }
	if h.Length() != 0 { t.Fatal("expected 0") }
	var nilH *Hashmap
	if nilH.Length() != 0 { t.Fatal("expected 0") }
	if nilH.SafeItems() != nil { t.Fatal("expected nil") }
}

func TestHashmap_AddAndGet(t *testing.T) {
	h := New.Hashmap.Cap(5)
	isNew := h.AddOrUpdate("k1", "v1")
	if !isNew { t.Fatal("expected new") }
	isNew2 := h.AddOrUpdate("k1", "v2")
	if isNew2 { t.Fatal("expected not new") }
	v, found := h.Get("k1")
	if !found || v != "v2" { t.Fatal("unexpected") }
	v2, found2 := h.GetValue("k1")
	if !found2 || v2 != "v2" { t.Fatal("unexpected") }
}

func TestHashmap_Set(t *testing.T) {
	h := New.Hashmap.Empty()
	h.Set("a", "1")
	h.SetTrim(" b ", " 2 ")
	if h.Length() != 2 { t.Fatal("expected 2") }
	h.SetBySplitter("=", "c=3")
	h.SetBySplitter("=", "d")
	if h.Length() != 4 { t.Fatal("expected 4") }
}

func TestHashmap_Has(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
	if !h.Has("a") || !h.Contains("a") { t.Fatal("expected has") }
	if h.IsKeyMissing("a") { t.Fatal("unexpected") }
	if !h.IsKeyMissing("z") { t.Fatal("expected missing") }
	if !h.HasAll("a", "b") { t.Fatal("expected has all") }
	if !h.HasAllStrings("a", "b") { t.Fatal("expected has all") }
	if !h.HasAny("a", "z") { t.Fatal("expected has any") }
	if h.HasAny("x", "z") { t.Fatal("unexpected") }
}

func TestHashmap_HasLock(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	if !h.HasLock("a") { t.Fatal("expected") }
	if !h.HasWithLock("a") { t.Fatal("expected") }
	if !h.ContainsLock("a") { t.Fatal("expected") }
	if h.IsKeyMissingLock("a") { t.Fatal("unexpected") }
	if !h.IsEmptyLock() == true { /* not empty */ }
}

func TestHashmap_AddVariants(t *testing.T) {
	h := New.Hashmap.Empty()
	h.AddOrUpdateKeyStrValInt("n", 42)
	h.AddOrUpdateKeyStrValFloat("f", 3.14)
	h.AddOrUpdateKeyStrValFloat64("f64", 3.14)
	h.AddOrUpdateKeyStrValAny("any", "val")
	h.AddOrUpdateKeyValueAny(KeyAnyValuePair{Key: "kav", Value: 1})
	h.AddOrUpdateKeyVal(KeyValuePair{Key: "kv", Value: "vv"})
	h.AddOrUpdateLock("lk", "lv")
	if h.Length() != 7 { t.Fatal("expected 7, got", h.Length()) }
}

func TestHashmap_AddOrUpdateHashmap(t *testing.T) {
	h1 := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	h2 := New.Hashmap.UsingMap(map[string]string{"b": "2"})
	h1.AddOrUpdateHashmap(h2)
	h1.AddOrUpdateHashmap(nil)
	if h1.Length() != 2 { t.Fatal("expected 2") }
	h1.AddOrUpdateMap(map[string]string{"c": "3"})
	h1.AddOrUpdateMap(nil)
	if h1.Length() != 3 { t.Fatal("expected 3") }
}

func TestHashmap_AddsOrUpdates(t *testing.T) {
	h := New.Hashmap.Empty()
	h.AddsOrUpdates(KeyValuePair{Key: "a", Value: "1"})
	h.AddsOrUpdates()
	h.AddOrUpdateKeyAnyValues(KeyAnyValuePair{Key: "b", Value: 2})
	h.AddOrUpdateKeyAnyValues()
	h.AddOrUpdateKeyValues(KeyValuePair{Key: "c", Value: "3"})
	h.AddOrUpdateKeyValues()
	if h.Length() != 3 { t.Fatal("expected 3") }
}

func TestHashmap_AddOrUpdateCollection(t *testing.T) {
	h := New.Hashmap.Empty()
	keys := New.Collection.Strings([]string{"k1", "k2"})
	vals := New.Collection.Strings([]string{"v1", "v2"})
	h.AddOrUpdateCollection(keys, vals)
	if h.Length() != 2 { t.Fatal("expected 2") }
	h.AddOrUpdateCollection(nil, nil)
	// length mismatch
	h.AddOrUpdateCollection(New.Collection.Strings([]string{"a"}), New.Collection.Strings([]string{"b", "c"}))
}

func TestHashmap_WgLock(t *testing.T) {
	h := New.Hashmap.Empty()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	h.AddOrUpdateWithWgLock("k", "v", wg)
	if h.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_Filter(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"abc": "1", "def": "2"})
	filter := func(s string, i int) (string, bool, bool) { return s, s == "abc", false }
	items := h.GetKeysFilteredItems(filter)
	if len(items) != 1 { t.Fatal("expected 1") }
	col := h.GetKeysFilteredCollection(filter)
	if col.Length() != 1 { t.Fatal("expected 1") }
	// empty
	empty := New.Hashmap.Empty()
	if len(empty.GetKeysFilteredItems(filter)) != 0 { t.Fatal("expected 0") }
}

func TestHashmap_FilterWithBreak(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
	filter := func(s string, i int) (string, bool, bool) { return s, true, true }
	items := h.GetKeysFilteredItems(filter)
	if len(items) != 1 { t.Fatal("expected 1") }
	col := h.GetKeysFilteredCollection(filter)
	if col.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	h := New.Hashmap.Empty()
	f := func(p KeyValuePair) (string, bool, bool) { return p.Value, true, false }
	h.AddsOrUpdatesUsingFilter(f, KeyValuePair{Key: "a", Value: "1"})
	h.AddsOrUpdatesUsingFilter(f)
	af := func(p KeyAnyValuePair) (string, bool, bool) { return "v", true, false }
	h.AddsOrUpdatesAnyUsingFilter(af, KeyAnyValuePair{Key: "b", Value: 2})
	h.AddsOrUpdatesAnyUsingFilter(af)
	h.AddsOrUpdatesAnyUsingFilterLock(af, KeyAnyValuePair{Key: "c", Value: 3})
	h.AddsOrUpdatesAnyUsingFilterLock(af)
}

func TestHashmap_Keys(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
	if len(h.AllKeys()) != 2 || len(h.Keys()) != 2 { t.Fatal("expected 2") }
	if h.KeysCollection().Length() != 2 { t.Fatal("expected 2") }
	_ = h.KeysLock()
	_ = h.ValuesListCopyLock()
	_ = h.KeysValuesListLock()
	_ = h.ItemsCopyLock()
}

func TestHashmap_Values(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	if len(h.ValuesList()) != 1 { t.Fatal("expected 1") }
	_ = h.ValuesCollection()
	_ = h.ValuesHashset()
	_ = h.ValuesCollectionLock()
	_ = h.ValuesHashsetLock()
	_ = h.Collection()
	k, v := h.KeysValuesCollection()
	if k.Length() != 1 || v.Length() != 1 { t.Fatal("expected 1") }
	k2, v2 := h.KeysValuesList()
	if len(k2) != 1 || len(v2) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeyValuePairs(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	pairs := h.KeysValuePairs()
	if len(pairs) != 1 { t.Fatal("expected 1") }
	pairsCol := h.KeysValuePairsCollection()
	if pairsCol.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_Remove(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
	h.Remove("a")
	if h.Length() != 1 { t.Fatal("expected 1") }
	h.RemoveWithLock("b")
	if h.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_Diff(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	_ = h.DiffRaw(map[string]string{"a": "2"})
	h2 := New.Hashmap.UsingMap(map[string]string{"a": "2"})
	_ = h.Diff(h2)
}

func TestHashmap_IsEqual(t *testing.T) {
	h1 := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	h2 := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	if !h1.IsEqualPtr(h2) { t.Fatal("expected equal") }
	if !h1.IsEqualPtrLock(h2) { t.Fatal("expected equal") }
	h3 := New.Hashmap.UsingMap(map[string]string{"a": "2"})
	if h1.IsEqualPtr(h3) { t.Fatal("expected not equal") }
}

func TestHashmap_ConcatNew(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	h2 := New.Hashmap.UsingMap(map[string]string{"b": "2"})
	concat := h.ConcatNew(false, h2)
	if concat.Length() < 2 { t.Fatal("expected >= 2") }
	concat2 := h.ConcatNew(true)
	_ = concat2
	concat3 := h.ConcatNewUsingMaps(false, map[string]string{"c": "3"})
	_ = concat3
	concat4 := h.ConcatNewUsingMaps(true)
	_ = concat4
}

func TestHashmap_StringAndJson(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	if h.String() == "" { t.Fatal("expected non-empty") }
	if h.StringLock() == "" { t.Fatal("expected non-empty") }
	_ = h.Join(",")
	_ = h.JoinKeys(",")
	_ = h.JsonModel()
	_ = h.JsonModelAny()
	_, _ = h.MarshalJSON()
	_, _ = h.Serialize()
	_ = h.AsJsoner()
	_ = h.AsJsonContractsBinder()
	_ = h.AsJsonParseSelfInjector()
	_ = h.AsJsonMarshaller()
}

func TestHashmap_KeysToLower(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"ABC": "1"})
	lower := h.KeysToLower()
	if !lower.Has("abc") { t.Fatal("expected lowercase") }
	_ = h.ValuesToLower() // deprecated alias
}

func TestHashmap_GetExcept(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
	hs := New.Hashset.StringsSpreadItems("a")
	r := h.GetValuesExceptKeysInHashset(hs)
	if len(r) != 1 { t.Fatal("expected 1") }
	r2 := h.GetValuesKeysExcept([]string{"a"})
	if len(r2) != 1 { t.Fatal("expected 1") }
	r3 := h.GetAllExceptCollection(New.Collection.Strings([]string{"a"}))
	if len(r3) != 1 { t.Fatal("expected 1") }
	_ = h.GetValuesExceptKeysInHashset(nil)
	_ = h.GetValuesKeysExcept(nil)
	_ = h.GetAllExceptCollection(nil)
}

func TestHashmap_HasAllCollectionItems(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1", "b": "2"})
	c := New.Collection.Strings([]string{"a", "b"})
	if !h.HasAllCollectionItems(c) { t.Fatal("expected true") }
	if h.HasAllCollectionItems(nil) { t.Fatal("expected false") }
}

func TestHashmap_ToError(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	_ = h.ToError(",")
	_ = h.ToDefaultError()
	_ = h.KeyValStringLines()
}

func TestHashmap_ClearDispose(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	h.Clear()
	if h.Length() != 0 { t.Fatal("expected 0") }
	h2 := New.Hashmap.UsingMap(map[string]string{"b": "2"})
	h2.Dispose()
}

func TestHashmap_Clone(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	c := h.Clone()
	if c.Length() != 1 { t.Fatal("expected 1") }
	cp := h.ClonePtr()
	if cp.Length() != 1 { t.Fatal("expected 1") }
	var nilH *Hashmap
	if nilH.ClonePtr() != nil { t.Fatal("expected nil") }
}

func TestHashmap_ToStringsUsingCompiler(t *testing.T) {
	h := New.Hashmap.UsingMap(map[string]string{"a": "1"})
	s := h.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
	if len(s) != 1 { t.Fatal("expected 1") }
	empty := New.Hashmap.Empty()
	s2 := empty.ToStringsUsingCompiler(func(k, v string) string { return k })
	if len(s2) != 0 { t.Fatal("expected 0") }
}

func TestHashmap_EmptyString(t *testing.T) {
	h := New.Hashmap.Empty()
	if h.String() == "" { t.Fatal("expected non-empty (NoElements)") }
	if h.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	h := New.Hashmap.Empty()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	h.AddOrUpdateStringsPtrWgLock(wg, []string{"a"}, []string{"1"})
	if h.Length() != 1 { t.Fatal("expected 1") }
	// empty
	wg2 := &sync.WaitGroup{}
	h.AddOrUpdateStringsPtrWgLock(wg2, nil, nil)
}
