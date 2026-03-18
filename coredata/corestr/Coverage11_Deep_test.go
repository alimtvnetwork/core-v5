package corestr

import (
	"testing"
)

// ── Collection extended ──

func TestCollection_JsonString(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	s := c.JsonString()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_JsonStringMust(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := c.JsonStringMust()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_HasAnyItem(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if !c.HasAnyItem() { t.Fatal("expected true") }
}

func TestCollection_LastIndex(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if c.LastIndex() != 1 { t.Fatal("expected 1") }
}

func TestCollection_HasIndex(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.HasIndex(1) { t.Fatal("expected true") }
	if c.HasIndex(5) { t.Fatal("expected false") }
	if c.HasIndex(-1) { t.Fatal("expected false") }
}

func TestCollection_ListStringsPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if len(c.ListStringsPtr()) != 1 { t.Fatal("expected 1") }
}

func TestCollection_ListStrings(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if len(c.ListStrings()) != 1 { t.Fatal("expected 1") }
}

func TestCollection_StringJSON(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := c.StringJSON()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_RemoveAt(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	ok := c.RemoveAt(1)
	if !ok || c.Length() != 2 { t.Fatal("unexpected") }
	fail := c.RemoveAt(100)
	if fail { t.Fatal("expected false") }
	fail2 := c.RemoveAt(-1)
	if fail2 { t.Fatal("expected false") }
}

func TestCollection_Count(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.Count() != 1 { t.Fatal("expected 1") }
}

func TestCollection_Capacity(t *testing.T) {
	c := New.Collection.Cap(10)
	if c.Capacity() == 0 { t.Fatal("expected > 0") }
}

func TestCollection_Capacity_Nil(t *testing.T) {
	c := &Collection{}
	if c.Capacity() != 0 { t.Fatal("expected 0") }
}

func TestCollection_IsEquals(t *testing.T) {
	a := New.Collection.Strings([]string{"a", "b"})
	b := New.Collection.Strings([]string{"a", "b"})
	if !a.IsEquals(b) { t.Fatal("expected true") }
}

func TestCollection_IsEquals_Different(t *testing.T) {
	a := New.Collection.Strings([]string{"a"})
	b := New.Collection.Strings([]string{"b"})
	if a.IsEquals(b) { t.Fatal("expected false") }
}

func TestCollection_IsEquals_BothNil(t *testing.T) {
	var a, b *Collection
	result, handled := isCollectionPrecheckEqual(a, b)
	if !handled || !result { t.Fatal("expected true") }
}

func TestCollection_IsEquals_OneNil(t *testing.T) {
	a := New.Collection.Strings([]string{"a"})
	result, handled := isCollectionPrecheckEqual(a, nil)
	if !handled || result { t.Fatal("expected false") }
}

func TestCollection_IsEquals_BothEmpty(t *testing.T) {
	a := New.Collection.Cap(0)
	b := New.Collection.Cap(0)
	if !a.IsEquals(b) { t.Fatal("expected true") }
}

func TestCollection_IsEquals_DiffLength(t *testing.T) {
	a := New.Collection.Strings([]string{"a"})
	b := New.Collection.Strings([]string{"a", "b"})
	if a.IsEquals(b) { t.Fatal("expected false") }
}

func TestCollection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	a := New.Collection.Strings([]string{"Hello"})
	b := New.Collection.Strings([]string{"hello"})
	if !a.IsEqualsWithSensitive(false, b) { t.Fatal("expected true") }
}

func TestCollection_IsEmptyLock(t *testing.T) {
	c := New.Collection.Cap(0)
	if !c.IsEmptyLock() { t.Fatal("expected true") }
}

func TestCollection_HasItems(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if !c.HasItems() { t.Fatal("expected true") }
	var nilC *Collection
	if nilC.HasItems() { t.Fatal("expected false") }
}

func TestCollection_AddLock(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddLock("a")
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddNonEmpty(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddNonEmpty("")
	c.AddNonEmpty("a")
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddNonEmptyWhitespace(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddNonEmptyWhitespace("  ")
	c.AddNonEmptyWhitespace("a")
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddError(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddError(nil) // skip
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AsDefaultError(t *testing.T) {
	c := New.Collection.Strings([]string{"err1"})
	e := c.AsDefaultError()
	if e == nil { t.Fatal("expected error") }
}

func TestCollection_AsError_Empty(t *testing.T) {
	c := New.Collection.Cap(0)
	e := c.AsError("\n")
	if e != nil { t.Fatal("expected nil") }
}

func TestCollection_EachItemSplitBy(t *testing.T) {
	c := New.Collection.Strings([]string{"a,b", "c,d"})
	items := c.EachItemSplitBy(",")
	if len(items) != 4 { t.Fatal("expected 4") }
}

func TestCollection_ConcatNew(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c2 := c.ConcatNew(0, "b", "c")
	if c2.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_ConcatNew_Empty(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c2 := c.ConcatNew(0)
	if c2.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_ToError(t *testing.T) {
	c := New.Collection.Strings([]string{"e1"})
	_ = c.ToError("\n")
}

func TestCollection_ToDefaultError(t *testing.T) {
	c := New.Collection.Strings([]string{"e1"})
	_ = c.ToDefaultError()
}

func TestCollection_AddIfMany(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddIfMany(true, "a", "b")
	c.AddIfMany(false, "c")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddFunc(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddFunc(func() string { return "x" })
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddsLock(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddsLock("a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddStrings(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddStrings([]string{"a", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddCollection(t *testing.T) {
	c := New.Collection.Cap(5)
	c2 := New.Collection.Strings([]string{"a"})
	c.AddCollection(c2)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddCollection_Empty(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddCollection(New.Collection.Cap(0))
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AddCollections(t *testing.T) {
	c := New.Collection.Cap(5)
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	c.AddCollections(c1, c2)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_LengthLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.LengthLock() != 1 { t.Fatal("expected 1") }
}

// ── Hashmap extended ──

func TestHashmap_HasItems(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	if !hm.HasItems() { t.Fatal("expected true") }
}

func TestHashmap_Collection(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	c := hm.Collection()
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashmap_IsEmptyLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	if !hm.IsEmptyLock() { t.Fatal("expected true") }
}

func TestHashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValInt("k", 42)
	v, _ := hm.Get("k")
	if v != "42" { t.Fatal("expected 42") }
}

func TestHashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat("k", 3.14)
	_, found := hm.Get("k")
	if !found { t.Fatal("expected found") }
}

func TestHashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat64("k", 3.14)
	_, found := hm.Get("k")
	if !found { t.Fatal("expected found") }
}

func TestHashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValAny("k", 42)
	if !hm.Has("k") { t.Fatal("expected found") }
}

func TestHashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyValueAny(KeyAnyValuePair{Key: "k", Value: 42})
	if !hm.Has("k") { t.Fatal("expected found") }
}

func TestHashmap_AddOrUpdateKeyVal(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	isNew := hm.AddOrUpdateKeyVal(KeyValuePair{Key: "k", Value: "v"})
	if !isNew { t.Fatal("expected new") }
	isNew2 := hm.AddOrUpdateKeyVal(KeyValuePair{Key: "k", Value: "v2"})
	if isNew2 { t.Fatal("expected update") }
}

func TestHashmap_Set(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.Set("k", "v")
	if !hm.Has("k") { t.Fatal("expected found") }
}

func TestHashmap_SetTrim(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.SetTrim("  k  ", "  v  ")
	if !hm.Has("k") { t.Fatal("expected found") }
}

func TestHashmap_SetBySplitter(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.SetBySplitter("=", "key=value")
	v, _ := hm.Get("key")
	if v != "value" { t.Fatal("expected value") }
}

func TestHashmap_SetBySplitter_NoSplit(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.SetBySplitter("=", "keyonly")
	v, _ := hm.Get("keyonly")
	if v != "" { t.Fatal("expected empty") }
}

func TestHashmap_Contains(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	if !hm.Contains("k") { t.Fatal("expected true") }
}

func TestHashmap_ContainsLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	if !hm.ContainsLock("k") { t.Fatal("expected true") }
}

func TestHashmap_IsKeyMissing(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	if hm.IsKeyMissing("k") { t.Fatal("expected false") }
	if !hm.IsKeyMissing("x") { t.Fatal("expected true") }
}

func TestHashmap_IsKeyMissingLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	if hm.IsKeyMissingLock("k") { t.Fatal("expected false") }
}

func TestHashmap_HasLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	if !hm.HasLock("k") { t.Fatal("expected true") }
}

func TestHashmap_HasAllStrings(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	if !hm.HasAllStrings("a", "b") { t.Fatal("expected true") }
	if hm.HasAllStrings("a", "c") { t.Fatal("expected false") }
}

func TestHashmap_AddOrUpdateMap(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateMap(map[string]string{"a": "1"})
	if !hm.Has("a") { t.Fatal("expected found") }
}

func TestHashmap_AddsOrUpdates(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdates(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddOrUpdateLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateLock("k", "v")
	if !hm.Has("k") { t.Fatal("expected found") }
}

func TestHashmap_AddOrUpdateKeyValues(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyValues(KeyValuePair{Key: "a", Value: "1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyAnyValues(KeyAnyValuePair{Key: "a", Value: 1})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddOrUpdateHashmap(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdateHashmap(hm1)
	if !hm2.Has("a") { t.Fatal("expected found") }
}

func TestHashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateHashmap(nil)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_AddOrUpdateCollection(t *testing.T) {
	keys := New.Collection.Strings([]string{"a", "b"})
	vals := New.Collection.Strings([]string{"1", "2"})
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateCollection(keys, vals)
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddOrUpdateCollection_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateCollection(nil, nil)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_AddOrUpdateCollection_LenMismatch(t *testing.T) {
	keys := New.Collection.Strings([]string{"a"})
	vals := New.Collection.Strings([]string{"1", "2"})
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateCollection(keys, vals)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_ConcatNew(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := New.Hashmap.KeyValues(KeyValuePair{Key: "b", Value: "2"})
	result := hm1.ConcatNew(false, hm2)
	if result.Length() < 2 { t.Fatal("expected >= 2") }
}

func TestHashmap_ConcatNew_Empty(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	result := hm.ConcatNew(true)
	if result.Length() < 1 { t.Fatal("expected >= 1") }
}

func TestHashmap_ConcatNewUsingMaps(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	result := hm.ConcatNewUsingMaps(false, map[string]string{"b": "2"})
	if result.Length() < 2 { t.Fatal("expected >= 2") }
}

func TestHashmap_ConcatNewUsingMaps_Empty(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	result := hm.ConcatNewUsingMaps(true)
	if result.Length() < 1 { t.Fatal("expected >= 1") }
}

// ── Hashset extended ──

func TestHashset_HasItems(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if !hs.HasItems() { t.Fatal("expected true") }
}

func TestHashset_HasItems_Nil(t *testing.T) {
	var hs *Hashset
	if hs.HasItems() { t.Fatal("expected false") }
}

// ── SimpleSlice extended ──

func TestSimpleSlice_HasItems(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if !ss.HasAnyItem() { t.Fatal("expected true") }
}

// ── ValidValue / ValidValues ──

func TestValidValue_NewValidValue(t *testing.T) {
	vv := NewValidValue("hello")
	if !vv.IsValid || vv.Value != "hello" { t.Fatal("unexpected") }
}

func TestValidValue_NewInvalidValue(t *testing.T) {
	vv := NewInvalidValue("x")
	if vv.IsValid { t.Fatal("expected invalid") }
}

// ── LeftRightFromSplit ──

func TestLeftRightFromSplit_Valid(t *testing.T) {
	lr := NewLeftRightFromSplit("key=val", "=")
	if lr.Left != "key" || lr.Right != "val" { t.Fatal("unexpected") }
}

func TestLeftRightFromSplit_NoSplit(t *testing.T) {
	lr := NewLeftRightFromSplit("nosplit", "=")
	if lr.Left != "nosplit" { t.Fatal("unexpected") }
}

// ── LeftMiddleRightFromSplit ──

func TestLeftMiddleRightFromSplit_Valid(t *testing.T) {
	lmr := NewLeftMiddleRightFromSplit("a:b:c", ":")
	if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" { t.Fatal("unexpected") }
}

// ── CollectionsOfCollection ──

func TestCollectionsOfCollection_HasItems(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	if coc.HasItems() { t.Fatal("expected false") }
	c := New.Collection.Strings([]string{"a"})
	coc.Add(c)
	if !coc.HasItems() { t.Fatal("expected true") }
}

func TestCollectionsOfCollection_AllIndividualItemsLength(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	c := New.Collection.Strings([]string{"a", "b"})
	coc.Add(c)
	if coc.AllIndividualItemsLength() != 2 { t.Fatal("expected 2") }
}

func TestCollectionsOfCollection_Items(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	c := New.Collection.Strings([]string{"a"})
	coc.Add(c)
	if len(coc.Items()) != 1 { t.Fatal("expected 1") }
}

func TestCollectionsOfCollection_List(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	c := New.Collection.Strings([]string{"a", "b"})
	coc.Add(c)
	list := coc.List(0)
	if len(list) != 2 { t.Fatal("expected 2") }
}

func TestCollectionsOfCollection_ToCollection(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	c := New.Collection.Strings([]string{"a"})
	coc.Add(c)
	col := coc.ToCollection()
	if col.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollectionsOfCollection_AddStrings(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	coc.AddStrings(false, []string{"a", "b"})
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollectionsOfCollection_AddStrings_Empty(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	coc.AddStrings(false, nil)
	if coc.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollectionsOfCollection_AddsStringsOfStrings(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})
	if coc.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollectionsOfCollection_Adds(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	c := *New.Collection.Strings([]string{"a"})
	coc.Adds(c)
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollectionsOfCollection_AddCollections(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	c := *New.Collection.Strings([]string{"a"})
	coc.AddCollections(c)
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollectionsOfCollection_String(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	c := New.Collection.Strings([]string{"a"})
	coc.Add(c)
	s := coc.String()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollectionsOfCollection_JsonModel(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	m := coc.JsonModel()
	_ = m
}

func TestCollectionsOfCollection_JsonModelAny(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	a := coc.JsonModelAny()
	_ = a
}

func TestCollectionsOfCollection_AsJsoner(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	_ = coc.AsJsoner()
}

func TestCollectionsOfCollection_AsJsonParseSelfInjector(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	_ = coc.AsJsonParseSelfInjector()
}

func TestCollectionsOfCollection_AsJsonMarshaller(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	_ = coc.AsJsonMarshaller()
}

func TestCollectionsOfCollection_AsJsonContractsBinder(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	_ = coc.AsJsonContractsBinder()
}

// ── CloneSlice ──

func TestCloneSlice_Valid(t *testing.T) {
	orig := []string{"a", "b"}
	cloned := CloneSlice(orig)
	orig[0] = "X"
	if cloned[0] != "a" { t.Fatal("expected deep clone") }
}

func TestCloneSlice_Nil(t *testing.T) {
	cloned := CloneSlice(nil)
	if len(cloned) != 0 { t.Fatal("expected empty") }
}

func TestCloneSliceIf_True(t *testing.T) {
	cloned := CloneSliceIf(true, "a", "b")
	if len(cloned) != 2 { t.Fatal("expected 2") }
}

func TestCloneSliceIf_False(t *testing.T) {
	cloned := CloneSliceIf(false, "a")
	if len(cloned) != 1 { t.Fatal("expected 1") }
}

// ── KeyAnyValuePair ──

func TestKeyAnyValuePair_ValueString(t *testing.T) {
	kv := KeyAnyValuePair{Key: "k", Value: 42}
	s := kv.ValueString()
	if s == "" { t.Fatal("expected non-empty") }
}

// ── AllIndividualStringsOfStringsLength ──

func TestAllIndividualStringsOfStringsLength(t *testing.T) {
	input := [][]string{{"a", "b"}, {"c"}}
	l := AllIndividualStringsOfStringsLength(input)
	if l != 3 { t.Fatal("expected 3") }
}

func TestAllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	l := AllIndividualStringsOfStringsLength(nil)
	if l != 0 { t.Fatal("expected 0") }
}
