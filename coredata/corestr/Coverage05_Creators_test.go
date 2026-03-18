package corestr

import (
	"testing"
)

// ── newCollectionCreator ──

func TestNewCollectionCreator(t *testing.T) {
	c := New.Collection.Empty()
	if !c.IsEmpty() {
		t.Fatal("expected empty")
	}
	c2 := New.Collection.Cap(10)
	if c2.Capacity() < 10 {
		t.Fatal("expected cap >= 10")
	}
	c3 := New.Collection.Strings([]string{"a", "b"})
	if c3.Length() != 2 {
		t.Fatal("expected 2")
	}
	c4 := New.Collection.Create([]string{"a"})
	if c4.Length() != 1 {
		t.Fatal("expected 1")
	}
	c5 := New.Collection.CloneStrings([]string{"a"})
	if c5.Length() != 1 {
		t.Fatal("expected 1")
	}
	c6 := New.Collection.StringsOptions(true, []string{"a"})
	if c6.Length() != 1 {
		t.Fatal("expected 1")
	}
	c7 := New.Collection.StringsOptions(false, []string{"a"})
	if c7.Length() != 1 {
		t.Fatal("expected 1")
	}
	c8 := New.Collection.StringsOptions(false, nil)
	if !c8.IsEmpty() {
		t.Fatal("expected empty")
	}
	c9 := New.Collection.LineUsingSep(",", "a,b")
	if c9.Length() != 2 {
		t.Fatal("expected 2")
	}
	c10 := New.Collection.LineDefault("a\nb")
	_ = c10
	c11 := New.Collection.StringsPlusCap(5, []string{"a"})
	if c11.Length() != 1 {
		t.Fatal("expected 1")
	}
	c12 := New.Collection.StringsPlusCap(0, []string{"a"})
	_ = c12
	c13 := New.Collection.CapStrings(5, []string{"a"})
	_ = c13
	c14 := New.Collection.CapStrings(0, []string{"a"})
	_ = c14
	c15 := New.Collection.LenCap(2, 5)
	if c15.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ── newSimpleSliceCreator ──

func TestNewSimpleSliceCreator(t *testing.T) {
	s := New.SimpleSlice.Empty()
	if !s.IsEmpty() {
		t.Fatal("expected empty")
	}
	s2 := New.SimpleSlice.Cap(10)
	if s2.Length() != 0 {
		t.Fatal("expected 0")
	}
	s3 := New.SimpleSlice.Cap(-1)
	_ = s3
	s4 := New.SimpleSlice.Default()
	_ = s4
	s5 := New.SimpleSlice.Lines("a", "b")
	if s5.Length() != 2 {
		t.Fatal("expected 2")
	}
	s6 := New.SimpleSlice.SpreadStrings("a")
	_ = s6
	s7 := New.SimpleSlice.Strings([]string{"a"})
	_ = s7
	s8 := New.SimpleSlice.Create([]string{"a"})
	_ = s8
	s9 := New.SimpleSlice.StringsPtr([]string{"a"})
	_ = s9
	s10 := New.SimpleSlice.StringsPtr(nil)
	_ = s10
	s11 := New.SimpleSlice.StringsOptions(true, []string{"a"})
	_ = s11
	s12 := New.SimpleSlice.StringsOptions(false, []string{"a"})
	_ = s12
	s13 := New.SimpleSlice.StringsOptions(false, nil)
	_ = s13
	s14 := New.SimpleSlice.StringsClone([]string{"a"})
	_ = s14
	s15 := New.SimpleSlice.StringsClone(nil)
	_ = s15
	s16 := New.SimpleSlice.Direct(true, []string{"a"})
	_ = s16
	s17 := New.SimpleSlice.Direct(false, []string{"a"})
	_ = s17
	s18 := New.SimpleSlice.Direct(true, nil)
	_ = s18
	s19 := New.SimpleSlice.UsingLines(true, "a")
	_ = s19
	s20 := New.SimpleSlice.UsingLines(false, "a")
	_ = s20
	s21 := New.SimpleSlice.UsingLines(true)
	_ = s21
	s22 := New.SimpleSlice.Split("a,b", ",")
	_ = s22
	s23 := New.SimpleSlice.SplitLines("a\nb")
	_ = s23
	s24 := New.SimpleSlice.UsingSeparatorLine(",", "a,b")
	_ = s24
	s25 := New.SimpleSlice.UsingLine("a\nb")
	_ = s25
	s26 := New.SimpleSlice.ByLen([]string{"a", "b"})
	_ = s26
	hs := New.Hashset.StringsSpreadItems("a", "b")
	s27 := New.SimpleSlice.Hashset(hs)
	_ = s27
	s28 := New.SimpleSlice.Map(map[string]int{"a": 1})
	_ = s28
	s29 := New.SimpleSlice.Map(nil)
	_ = s29
}

// ── newSimpleStringOnceCreator ──

func TestNewSimpleStringOnceCreator(t *testing.T) {
	s := New.SimpleStringOnce.Init("hello")
	if s.Value() != "hello" || !s.IsInitialized() {
		t.Fatal("unexpected")
	}
	s2 := New.SimpleStringOnce.InitPtr("hello")
	if s2.Value() != "hello" {
		t.Fatal("unexpected")
	}
	s3 := New.SimpleStringOnce.Uninitialized("val")
	if s3.IsInitialized() {
		t.Fatal("expected uninitialized")
	}
	s4 := New.SimpleStringOnce.Create("val", true)
	_ = s4
	s5 := New.SimpleStringOnce.CreatePtr("val", false)
	_ = s5
	s6 := New.SimpleStringOnce.Empty()
	_ = s6
	s7 := New.SimpleStringOnce.Any(false, "hello", true)
	_ = s7
}

// ── newHashsetCreator ──

func TestNewHashsetCreator(t *testing.T) {
	h := New.Hashset.Empty()
	if !h.IsEmpty() {
		t.Fatal("expected empty")
	}
	h2 := New.Hashset.Cap(10)
	_ = h2
	h3 := New.Hashset.Strings([]string{"a", "b"})
	if h3.Length() != 2 {
		t.Fatal("expected 2")
	}
	h4 := New.Hashset.Strings(nil)
	_ = h4
	h5 := New.Hashset.StringsSpreadItems("a")
	_ = h5
	h6 := New.Hashset.StringsSpreadItems()
	_ = h6
	h7 := New.Hashset.UsingMap(map[string]bool{"a": true})
	_ = h7
	h8 := New.Hashset.UsingMap(nil)
	_ = h8
	h9 := New.Hashset.UsingMapOption(5, true, map[string]bool{"a": true})
	_ = h9
	h10 := New.Hashset.UsingMapOption(5, false, map[string]bool{"a": true})
	_ = h10
	h11 := New.Hashset.UsingMapOption(5, false, nil)
	_ = h11
	h12 := New.Hashset.StringsOption(5, true, "a")
	_ = h12
	h13 := New.Hashset.StringsOption(0, false)
	_ = h13
	h14 := New.Hashset.StringsOption(5, false)
	_ = h14
	h15 := New.Hashset.UsingCollection(New.Collection.Strings([]string{"a"}))
	_ = h15
	h16 := New.Hashset.UsingCollection(nil)
	_ = h16
	ss := New.SimpleSlice.Lines("a")
	h17 := New.Hashset.SimpleSlice(ss)
	_ = h17
	emptySlice := New.SimpleSlice.Empty()
	h18 := New.Hashset.SimpleSlice(emptySlice)
	_ = h18
}

// ── newHashmapCreator ──

func TestNewHashmapCreator(t *testing.T) {
	h := New.Hashmap.Empty()
	if !h.IsEmpty() {
		t.Fatal("expected empty")
	}
	h2 := New.Hashmap.Cap(10)
	_ = h2
	h3 := New.Hashmap.UsingMap(map[string]string{"a": "b"})
	if h3.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	h4 := New.Hashmap.UsingMapOptions(true, 5, map[string]string{"a": "b"})
	_ = h4
	h5 := New.Hashmap.UsingMapOptions(false, 5, map[string]string{"a": "b"})
	_ = h5
	h6 := New.Hashmap.UsingMapOptions(false, 5, nil)
	_ = h6
	h7 := New.Hashmap.MapWithCap(5, map[string]string{"a": "b"})
	_ = h7
	h8 := New.Hashmap.MapWithCap(0, map[string]string{"a": "b"})
	_ = h8
	h9 := New.Hashmap.MapWithCap(5, nil)
	_ = h9
	h10 := New.Hashmap.KeyValues(KeyValuePair{Key: "k", Value: "v"})
	_ = h10
	h11 := New.Hashmap.KeyValues()
	_ = h11
	h12 := New.Hashmap.KeyAnyValues(KeyAnyValuePair{Key: "k", Value: "v"})
	_ = h12
	h13 := New.Hashmap.KeyAnyValues()
	_ = h13
	h14 := New.Hashmap.KeyValuesStrings([]string{"k"}, []string{"v"})
	_ = h14
	h15 := New.Hashmap.KeyValuesStrings(nil, nil)
	_ = h15
	keys := New.Collection.Strings([]string{"k"})
	vals := New.Collection.Strings([]string{"v"})
	h16 := New.Hashmap.KeyValuesCollection(keys, vals)
	_ = h16
	h17 := New.Hashmap.KeyValuesCollection(nil, nil)
	_ = h17
}

// ── newLinkedListCreator ──

func TestNewLinkedListCreator(t *testing.T) {
	ll := New.LinkedList.Create()
	if ll.Length() != 0 {
		t.Fatal("expected 0")
	}
	ll2 := New.LinkedList.Empty()
	_ = ll2
	ll3 := New.LinkedList.Strings([]string{"a", "b"})
	if ll3.Length() != 2 {
		t.Fatal("expected 2")
	}
	ll4 := New.LinkedList.Strings(nil)
	_ = ll4
	ll5 := New.LinkedList.SpreadStrings("a")
	_ = ll5
	ll6 := New.LinkedList.SpreadStrings()
	_ = ll6
	ll7 := New.LinkedList.UsingMap(map[string]bool{"a": true})
	_ = ll7
	ll8 := New.LinkedList.UsingMap(nil)
	_ = ll8
}

// ── newLinkedListCollectionsCreator ──

func TestNewLinkedCollectionCreator(t *testing.T) {
	lc := New.LinkedCollection.Create()
	if lc.Length() != 0 {
		t.Fatal("expected 0")
	}
	lc2 := New.LinkedCollection.Empty()
	_ = lc2
	lc3 := New.LinkedCollection.Strings("a", "b")
	if lc3.Length() != 1 { // adds as single collection
		t.Fatal("expected 1")
	}
	lc4 := New.LinkedCollection.Strings()
	_ = lc4
	lc5 := New.LinkedCollection.UsingCollections(New.Collection.Strings([]string{"a"}))
	_ = lc5
	lc6 := New.LinkedCollection.UsingCollections()
	_ = lc6
}

// ── other creators ──

func TestNewKeyValuesCreator(t *testing.T) {
	kv := New.KeyValues.Empty()
	if !kv.IsEmpty() {
		t.Fatal("expected empty")
	}
	kv2 := New.KeyValues.Cap(10)
	_ = kv2
	kv3 := New.KeyValues.UsingMap(map[string]string{"k": "v"})
	if kv3.Length() != 1 {
		t.Fatal("expected 1")
	}
	kv4 := New.KeyValues.UsingMap(nil)
	_ = kv4
	kv5 := New.KeyValues.UsingKeyValuePairs(KeyValuePair{Key: "k", Value: "v"})
	_ = kv5
	kv6 := New.KeyValues.UsingKeyValuePairs()
	_ = kv6
	kv7 := New.KeyValues.UsingKeyValueStrings([]string{"k"}, []string{"v"})
	_ = kv7
	kv8 := New.KeyValues.UsingKeyValueStrings(nil, nil)
	_ = kv8
}

func TestNewCollectionsOfCollectionCreator(t *testing.T) {
	c := New.CollectionsOfCollection.Empty()
	if !c.IsEmpty() {
		t.Fatal("expected empty")
	}
	c2 := New.CollectionsOfCollection.Cap(5)
	_ = c2
	c3 := New.CollectionsOfCollection.Strings([]string{"a"})
	_ = c3
	c4 := New.CollectionsOfCollection.CloneStrings([]string{"a"})
	_ = c4
	c5 := New.CollectionsOfCollection.StringsOption(true, 5, []string{"a"})
	_ = c5
	c6 := New.CollectionsOfCollection.StringsOptions(false, 0, []string{"a"})
	_ = c6
	c7 := New.CollectionsOfCollection.SpreadStrings(true, "a", "b")
	_ = c7
	c8 := New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
	_ = c8
	c9 := New.CollectionsOfCollection.LenCap(0, 5)
	_ = c9
}

func TestNewHashsetsCollectionCreator(t *testing.T) {
	hc := New.HashsetsCollection.Empty()
	if !hc.IsEmpty() {
		t.Fatal("expected empty")
	}
	hc2 := New.HashsetsCollection.Cap(5)
	_ = hc2
	hc3 := New.HashsetsCollection.LenCap(0, 5)
	_ = hc3
	hs := New.Hashset.StringsSpreadItems("a")
	hc4 := New.HashsetsCollection.UsingHashsetsPointers(hs)
	if hc4.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	hc5 := New.HashsetsCollection.UsingHashsetsPointers()
	_ = hc5
}

func TestNewCharCollectionMapCreator(t *testing.T) {
	ccm := New.CharCollectionMap.Empty()
	if !ccm.IsEmpty() {
		t.Fatal("expected empty")
	}
	ccm2 := New.CharCollectionMap.CapSelfCap(20, 20)
	_ = ccm2
	ccm3 := New.CharCollectionMap.CapSelfCap(1, 1) // below limit
	_ = ccm3
	ccm4 := New.CharCollectionMap.Items([]string{"a", "b"})
	_ = ccm4
	ccm5 := New.CharCollectionMap.Items(nil)
	_ = ccm5
	ccm6 := New.CharCollectionMap.ItemsPtrWithCap(5, 5, []string{"a"})
	_ = ccm6
	ccm7 := New.CharCollectionMap.ItemsPtrWithCap(5, 5, nil)
	_ = ccm7
}

func TestNewCharHashsetMapCreator(t *testing.T) {
	chm := New.CharHashsetMap.Cap(20, 20)
	if chm.IsEmpty() {
		// empty because no items added, but initialized
	}
	chm2 := New.CharHashsetMap.Cap(1, 1) // below limit
	_ = chm2
	chm3 := New.CharHashsetMap.CapItems(20, 20, "a", "b")
	_ = chm3
	chm4 := New.CharHashsetMap.Strings(20, []string{"a"})
	_ = chm4
	chm5 := New.CharHashsetMap.Strings(20, nil)
	_ = chm5
}
