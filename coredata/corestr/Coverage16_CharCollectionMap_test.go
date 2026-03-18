package corestr

import (
	"encoding/json"
	"testing"
)

// ── Creators ──────────────────────────────────────────────

func Test_CharCollectionMap_NewEmpty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	if cm == nil {
		t.Fatal("expected non-nil")
	}
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
	if cm.Length() != 0 {
		t.Fatal("expected length 0")
	}
}

func Test_CharCollectionMap_NewCapSelfCap(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(20, 5)
	if cm == nil || cm.HasItems() {
		t.Fatal("expected empty map")
	}
}

func Test_CharCollectionMap_NewItems(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"apple", "avocado", "banana"})
	if cm.Length() != 2 {
		t.Fatalf("expected 2 char groups, got %d", cm.Length())
	}
	if !cm.Has("apple") || !cm.Has("avocado") || !cm.Has("banana") {
		t.Fatal("missing items")
	}
}

func Test_CharCollectionMap_NewItems_Empty(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{})
	if !cm.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_CharCollectionMap_NewItemsPtrWithCap(t *testing.T) {
	cm := New.CharCollectionMap.ItemsPtrWithCap(5, 3, []string{"cat", "car", "dog"})
	if cm.Length() != 2 {
		t.Fatalf("expected 2, got %d", cm.Length())
	}
}

func Test_CharCollectionMap_NewItemsPtrWithCap_Empty(t *testing.T) {
	cm := New.CharCollectionMap.ItemsPtrWithCap(5, 3, []string{})
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── GetChar ───────────────────────────────────────────────

func Test_CharCollectionMap_GetChar(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	if cm.GetChar("hello") != 'h' {
		t.Fatal("expected 'h'")
	}
	if cm.GetChar("") != emptyChar {
		t.Fatal("expected emptyChar for empty string")
	}
}

// ── GetCharsGroups ────────────────────────────────────────

func Test_CharCollectionMap_GetCharsGroups(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	result := cm.GetCharsGroups([]string{"abc", "axy", "bcd"})
	if result.Length() != 2 {
		t.Fatalf("expected 2 groups, got %d", result.Length())
	}
}

func Test_CharCollectionMap_GetCharsGroups_Empty(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	result := cm.GetCharsGroups([]string{})
	if result != cm {
		t.Fatal("expected same reference on empty input")
	}
}

// ── Add / AddStrings ──────────────────────────────────────

func Test_CharCollectionMap_Add(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.Add("alpha")
	cm.Add("avocado")
	cm.Add("beta")
	if cm.Length() != 2 {
		t.Fatalf("expected 2, got %d", cm.Length())
	}
	if cm.AllLengthsSum() != 3 {
		t.Fatalf("expected sum 3, got %d", cm.AllLengthsSum())
	}
}

func Test_CharCollectionMap_AddStrings(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddStrings("x1", "x2", "y1")
	if cm.LengthOf('x') != 2 {
		t.Fatal("expected 2 for 'x'")
	}
}

func Test_CharCollectionMap_AddStrings_Empty(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddStrings()
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_CharCollectionMap_AddLock(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddLock("hello")
	cm.AddLock("help")
	if cm.LengthLock() != 1 {
		t.Fatal("expected 1 char group")
	}
}

// ── AddSameStartingCharItems ──────────────────────────────

func Test_CharCollectionMap_AddSameStartingCharItems(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddSameStartingCharItems('a', []string{"abc", "axy"}, false)
	if cm.LengthOfCollectionFromFirstChar("a") != 2 {
		t.Fatal("expected 2")
	}
	// Add more to existing char
	cm.AddSameStartingCharItems('a', []string{"azz"}, false)
	if cm.LengthOfCollectionFromFirstChar("a") != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CharCollectionMap_AddSameStartingCharItems_Empty(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddSameStartingCharItems('a', []string{}, false)
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── Has / HasWithCollection ───────────────────────────────

func Test_CharCollectionMap_Has(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"foo", "far", "bar"})
	if !cm.Has("foo") {
		t.Fatal("expected has foo")
	}
	if cm.Has("baz") {
		t.Fatal("expected no baz")
	}
	if cm.Has("zzz") {
		t.Fatal("expected no zzz")
	}
}

func Test_CharCollectionMap_Has_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	if cm.Has("anything") {
		t.Fatal("expected false on empty")
	}
}

func Test_CharCollectionMap_HasWithCollection(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"foo", "far"})
	has, col := cm.HasWithCollection("foo")
	if !has || col.IsEmpty() {
		t.Fatal("expected found with collection")
	}
	has2, col2 := cm.HasWithCollection("zzz")
	if has2 || col2 == nil {
		t.Fatal("expected not found")
	}
}

func Test_CharCollectionMap_HasWithCollection_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	has, col := cm.HasWithCollection("foo")
	if has || col == nil {
		t.Fatal("expected false, non-nil collection")
	}
}

func Test_CharCollectionMap_HasWithCollectionLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"foo", "far"})
	has, col := cm.HasWithCollectionLock("foo")
	if !has || col.IsEmpty() {
		t.Fatal("expected found")
	}
	has2, _ := cm.HasWithCollectionLock("zzz")
	if has2 {
		t.Fatal("expected not found")
	}
}

func Test_CharCollectionMap_HasWithCollectionLock_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	has, col := cm.HasWithCollectionLock("x")
	if has || col == nil {
		t.Fatal("expected false")
	}
}

// ── LengthOf / LengthOfLock ──────────────────────────────

func Test_CharCollectionMap_LengthOf(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "axy"})
	if cm.LengthOf('a') != 2 {
		t.Fatal("expected 2")
	}
	if cm.LengthOf('z') != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CharCollectionMap_LengthOf_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	if cm.LengthOf('a') != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CharCollectionMap_LengthOfLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	if cm.LengthOfLock('a') != 1 {
		t.Fatal("expected 1")
	}
	if cm.LengthOfLock('z') != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CharCollectionMap_LengthOfLock_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	if cm.LengthOfLock('a') != 0 {
		t.Fatal("expected 0")
	}
}

// ── LengthOfCollectionFromFirstChar ───────────────────────

func Test_CharCollectionMap_LengthOfCollectionFromFirstChar(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "axy"})
	if cm.LengthOfCollectionFromFirstChar("a") != 2 {
		t.Fatal("expected 2")
	}
	if cm.LengthOfCollectionFromFirstChar("z") != 0 {
		t.Fatal("expected 0")
	}
}

// ── AllLengthsSum / AllLengthsSumLock ─────────────────────

func Test_CharCollectionMap_AllLengthsSum(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"a1", "a2", "b1"})
	if cm.AllLengthsSum() != 3 {
		t.Fatalf("expected 3, got %d", cm.AllLengthsSum())
	}
}

func Test_CharCollectionMap_AllLengthsSumLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"a1", "b1"})
	if cm.AllLengthsSumLock() != 2 {
		t.Fatal("expected 2")
	}
}

// ── IsEmpty / HasItems / IsEmptyLock ──────────────────────

func Test_CharCollectionMap_IsEmpty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	if !cm.IsEmpty() || cm.HasItems() {
		t.Fatal("expected empty")
	}
	cm.Add("x")
	if cm.IsEmpty() || !cm.HasItems() {
		t.Fatal("expected has items")
	}
}

func Test_CharCollectionMap_IsEmptyLock(t *testing.T) {
	cm := Empty.CharCollectionMap()
	if !cm.IsEmptyLock() {
		t.Fatal("expected empty")
	}
}

// ── IsEquals / IsEqualsCaseSensitive ──────────────────────

func Test_CharCollectionMap_IsEquals(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	cm2 := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	if !cm1.IsEquals(cm2) {
		t.Fatal("expected equal")
	}
}

func Test_CharCollectionMap_IsEquals_Nil(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc"})
	if cm1.IsEquals(nil) {
		t.Fatal("expected not equal to nil")
	}
}

func Test_CharCollectionMap_IsEquals_SameRef(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc"})
	if !cm1.IsEqualsCaseSensitive(true, cm1) {
		t.Fatal("expected same ref equal")
	}
}

func Test_CharCollectionMap_IsEquals_BothEmpty(t *testing.T) {
	cm1 := Empty.CharCollectionMap()
	cm2 := Empty.CharCollectionMap()
	if !cm1.IsEquals(cm2) {
		t.Fatal("expected both empty equal")
	}
}

func Test_CharCollectionMap_IsEquals_OneEmpty(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc"})
	cm2 := Empty.CharCollectionMap()
	if cm1.IsEquals(cm2) {
		t.Fatal("expected not equal")
	}
}

func Test_CharCollectionMap_IsEquals_DiffLength(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	cm2 := New.CharCollectionMap.Items([]string{"abc"})
	if cm1.IsEquals(cm2) {
		t.Fatal("expected not equal")
	}
}

func Test_CharCollectionMap_IsEquals_DiffContent(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc"})
	cm2 := New.CharCollectionMap.Items([]string{"axy"})
	// same char key 'a', different content within
	if cm1.IsEquals(cm2) {
		t.Fatal("expected not equal")
	}
}

func Test_CharCollectionMap_IsEqualsLock(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc"})
	cm2 := New.CharCollectionMap.Items([]string{"abc"})
	if !cm1.IsEqualsLock(cm2) {
		t.Fatal("expected equal")
	}
}

func Test_CharCollectionMap_IsEqualsCaseSensitiveLock(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc"})
	cm2 := New.CharCollectionMap.Items([]string{"abc"})
	if !cm1.IsEqualsCaseSensitiveLock(true, cm2) {
		t.Fatal("expected equal")
	}
}

func Test_CharCollectionMap_IsEqualsMissingKey(t *testing.T) {
	cm1 := New.CharCollectionMap.Items([]string{"abc"})
	cm2 := New.CharCollectionMap.Items([]string{"xyz"})
	if cm1.IsEqualsCaseSensitive(true, cm2) {
		t.Fatal("expected not equal")
	}
}

// ── GetCollection ─────────────────────────────────────────

func Test_CharCollectionMap_GetCollection(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "axy"})
	col := cm.GetCollection("a", false)
	if col == nil || col.Length() != 2 {
		t.Fatal("expected collection with 2 items")
	}
	// missing char, no create
	col2 := cm.GetCollection("z", false)
	if col2 != nil {
		t.Fatal("expected nil")
	}
	// missing char, create
	col3 := cm.GetCollection("z", true)
	if col3 == nil {
		t.Fatal("expected new collection")
	}
}

func Test_CharCollectionMap_GetCollectionLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	col := cm.GetCollectionLock("a", false)
	if col == nil {
		t.Fatal("expected collection")
	}
}

func Test_CharCollectionMap_GetCollectionByChar(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	col := cm.GetCollectionByChar('a')
	if col == nil {
		t.Fatal("expected collection")
	}
}

// ── AddSameCharsCollection ────────────────────────────────

func Test_CharCollectionMap_AddSameCharsCollection(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	col := New.Collection.Strings([]string{"abc", "axy"})
	result := cm.AddSameCharsCollection("a", col)
	if result == nil || result.Length() != 2 {
		t.Fatal("expected 2 items")
	}
	// Add more to existing
	col2 := New.Collection.Strings([]string{"azz"})
	result2 := cm.AddSameCharsCollection("a", col2)
	if result2.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CharCollectionMap_AddSameCharsCollection_NilCol(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	result := cm.AddSameCharsCollection("a", nil)
	if result == nil {
		t.Fatal("expected new empty collection created")
	}
}

func Test_CharCollectionMap_AddSameCharsCollection_ExistingButNilAdd(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	result := cm.AddSameCharsCollection("a", nil)
	if result == nil {
		t.Fatal("expected existing collection returned")
	}
}

func Test_CharCollectionMap_AddSameCharsCollectionLock(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	col := New.Collection.Strings([]string{"abc"})
	result := cm.AddSameCharsCollectionLock("a", col)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_AddSameCharsCollectionLock_NilCol(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	result := cm.AddSameCharsCollectionLock("a", nil)
	if result == nil {
		t.Fatal("expected new collection")
	}
}

func Test_CharCollectionMap_AddSameCharsCollectionLock_ExistingButNilAdd(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	result := cm.AddSameCharsCollectionLock("a", nil)
	if result == nil {
		t.Fatal("expected existing")
	}
}

func Test_CharCollectionMap_AddSameCharsCollectionLock_AddToExisting(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	col := New.Collection.Strings([]string{"axy"})
	result := cm.AddSameCharsCollectionLock("a", col)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

// ── AddCollectionItems ────────────────────────────────────

func Test_CharCollectionMap_AddCollectionItems(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	col := New.Collection.Strings([]string{"abc", "xyz"})
	cm.AddCollectionItems(col)
	if cm.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharCollectionMap_AddCollectionItems_Nil(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddCollectionItems(nil)
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddHashmapsValues ─────────────────────────────────────

func Test_CharCollectionMap_AddHashmapsValues(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	hm := New.Hashmap.StringsKeyValue("k1", "alpha", "k2", "beta")
	cm.AddHashmapsValues(hm)
	if !cm.Has("alpha") || !cm.Has("beta") {
		t.Fatal("expected values added")
	}
}

func Test_CharCollectionMap_AddHashmapsValues_Nil(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddHashmapsValues(nil)
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddHashmapsKeysValuesBoth ─────────────────────────────

func Test_CharCollectionMap_AddHashmapsKeysValuesBoth(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	hm := New.Hashmap.StringsKeyValue("k1", "v1")
	cm.AddHashmapsKeysValuesBoth(hm)
	if !cm.Has("k1") || !cm.Has("v1") {
		t.Fatal("expected both keys and values")
	}
}

func Test_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddHashmapsKeysValuesBoth(nil)
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddHashmapsKeysOrValuesBothUsingFilter ────────────────

func Test_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	hm := New.Hashmap.StringsKeyValue("k1", "v1", "k2", "v2")
	cm.AddHashmapsKeysOrValuesBothUsingFilter(
		func(pair KeyValuePair) (string, bool, bool) {
			return pair.Key, true, false
		},
		hm,
	)
	if !cm.Has("k1") {
		t.Fatal("expected k1")
	}
}

func Test_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	cm.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Break(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	hm := New.Hashmap.StringsKeyValue("k1", "v1", "k2", "v2")
	cm.AddHashmapsKeysOrValuesBothUsingFilter(
		func(pair KeyValuePair) (string, bool, bool) {
			return pair.Key, true, true // break after first
		},
		hm,
	)
	if cm.AllLengthsSum() != 1 {
		t.Fatalf("expected 1 item, got %d", cm.AllLengthsSum())
	}
}

// ── AddCharHashsetMap ─────────────────────────────────────

func Test_CharCollectionMap_AddCharHashsetMap(t *testing.T) {
	cm := New.CharCollectionMap.CapSelfCap(10, 5)
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "axy")
	cm.AddCharHashsetMap(hsm)
	if cm.AllLengthsSum() != 2 {
		t.Fatal("expected 2")
	}
}

// ── Resize / AddLength ────────────────────────────────────

func Test_CharCollectionMap_Resize(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"a1"})
	cm.Resize(100)
	if !cm.Has("a1") {
		t.Fatal("items should be preserved")
	}
}

func Test_CharCollectionMap_Resize_NoShrink(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"a1", "b1"})
	cm.Resize(1) // smaller, should not shrink
	if cm.Length() != 2 {
		t.Fatal("should not shrink")
	}
}

func Test_CharCollectionMap_AddLength(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"a1"})
	cm.AddLength(10, 20)
	if !cm.Has("a1") {
		t.Fatal("items preserved")
	}
}

func Test_CharCollectionMap_AddLength_Empty(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"a1"})
	cm.AddLength()
	if !cm.Has("a1") {
		t.Fatal("items preserved")
	}
}

// ── List / ListLock / SortedListAsc ───────────────────────

func Test_CharCollectionMap_List(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	list := cm.List()
	if len(list) != 2 {
		t.Fatalf("expected 2, got %d", len(list))
	}
}

func Test_CharCollectionMap_List_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	list := cm.List()
	if len(list) != 0 {
		t.Fatal("expected empty list")
	}
}

func Test_CharCollectionMap_ListLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	list := cm.ListLock()
	if len(list) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CharCollectionMap_SortedListAsc(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"cherry", "apple", "banana"})
	sorted := cm.SortedListAsc()
	if len(sorted) < 3 {
		t.Fatal("expected 3")
	}
	if sorted[0] != "apple" || sorted[1] != "banana" || sorted[2] != "cherry" {
		t.Fatalf("unexpected order: %v", sorted)
	}
}

func Test_CharCollectionMap_SortedListAsc_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	sorted := cm.SortedListAsc()
	if len(sorted) != 0 {
		t.Fatal("expected empty")
	}
}

// ── GetMap / GetCopyMapLock ───────────────────────────────

func Test_CharCollectionMap_GetMap(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	m := cm.GetMap()
	if m == nil {
		t.Fatal("expected non-nil map")
	}
}

func Test_CharCollectionMap_GetCopyMapLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	m := cm.GetCopyMapLock()
	if m == nil || len(m) != 1 {
		t.Fatal("expected map with 1 entry")
	}
}

func Test_CharCollectionMap_GetCopyMapLock_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	m := cm.GetCopyMapLock()
	if len(m) != 0 {
		t.Fatal("expected empty map")
	}
}

// ── String / SummaryString ────────────────────────────────

func Test_CharCollectionMap_String(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	s := cm.String()
	if s == "" {
		t.Fatal("expected non-empty string")
	}
}

func Test_CharCollectionMap_SummaryString(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	s := cm.SummaryString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CharCollectionMap_StringLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	s := cm.StringLock()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CharCollectionMap_SummaryStringLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	s := cm.SummaryStringLock()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ── Print ─────────────────────────────────────────────────

func Test_CharCollectionMap_Print_Skip(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	cm.Print(false) // should not print
}

func Test_CharCollectionMap_PrintLock_Skip(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	cm.PrintLock(false)
}

// ── Hashset conversions ───────────────────────────────────

func Test_CharCollectionMap_HashsetByChar(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "axy"})
	hs := cm.HashsetByChar('a')
	if hs == nil || hs.Length() != 2 {
		t.Fatal("expected hashset with 2")
	}
	// missing
	hs2 := cm.HashsetByChar('z')
	if hs2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CharCollectionMap_HashsetByCharLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	hs := cm.HashsetByCharLock('a')
	if hs == nil {
		t.Fatal("expected non-nil")
	}
	hs2 := cm.HashsetByCharLock('z')
	if hs2 == nil || !hs2.IsEmpty() {
		t.Fatal("expected empty hashset")
	}
}

func Test_CharCollectionMap_HashsetByStringFirstChar(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	hs := cm.HashsetByStringFirstChar("abc")
	if hs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_HashsetByStringFirstCharLock(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	hs := cm.HashsetByStringFirstCharLock("abc")
	if hs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_HashsetsCollection(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	hsc := cm.HashsetsCollection()
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_HashsetsCollection_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	hsc := cm.HashsetsCollection()
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_HashsetsCollectionByChars(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	hsc := cm.HashsetsCollectionByChars('a', 'x')
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_HashsetsCollectionByChars_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	hsc := cm.HashsetsCollectionByChars('a')
	if hsc == nil {
		t.Fatal("expected non-nil empty")
	}
}

func Test_CharCollectionMap_HashsetsCollectionByStringFirstChar(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	hsc := cm.HashsetsCollectionByStringFirstChar("abc", "xyz")
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_HashsetsCollectionByStringFirstChar_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	hsc := cm.HashsetsCollectionByStringFirstChar("abc")
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

// ── JSON ──────────────────────────────────────────────────

func Test_CharCollectionMap_JsonModel(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	model := cm.JsonModel()
	if model == nil {
		t.Fatal("expected non-nil model")
	}
}

func Test_CharCollectionMap_JsonModelAny(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	any := cm.JsonModelAny()
	if any == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_MarshalUnmarshalJSON(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "axy", "xyz"})
	data, err := json.Marshal(cm)
	if err != nil {
		t.Fatal(err)
	}
	cm2 := Empty.CharCollectionMap()
	err = json.Unmarshal(data, cm2)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CharCollectionMap_Json(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	result := cm.Json()
	if result.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CharCollectionMap_JsonPtr(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	result := cm.JsonPtr()
	if result == nil || result.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CharCollectionMap_ParseInjectUsingJson(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	jsonResult := cm.JsonPtr()
	cm2 := New.CharCollectionMap.CapSelfCap(10, 5)
	_, err := cm2.ParseInjectUsingJson(jsonResult)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CharCollectionMap_ParseInjectUsingJsonMust(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	jsonResult := cm.JsonPtr()
	cm2 := New.CharCollectionMap.CapSelfCap(10, 5)
	result := cm2.ParseInjectUsingJsonMust(jsonResult)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_JsonParseSelfInject(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	jsonResult := cm.JsonPtr()
	cm2 := New.CharCollectionMap.CapSelfCap(10, 5)
	err := cm2.JsonParseSelfInject(jsonResult)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CharCollectionMap_AsJsonContractsBinder(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	if cm.AsJsonContractsBinder() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_AsJsoner(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	if cm.AsJsoner() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_AsJsonMarshaller(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	if cm.AsJsonMarshaller() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharCollectionMap_AsJsonParseSelfInjector(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	if cm.AsJsonParseSelfInjector() == nil {
		t.Fatal("expected non-nil")
	}
}

// ── Clear / Dispose ───────────────────────────────────────

func Test_CharCollectionMap_Clear(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	cm.Clear()
	if cm.HasItems() {
		t.Fatal("expected empty after clear")
	}
}

func Test_CharCollectionMap_Clear_Empty(t *testing.T) {
	cm := Empty.CharCollectionMap()
	cm.Clear()
	if cm.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_CharCollectionMap_Dispose(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	cm.Dispose()
	if cm.items != nil {
		t.Fatal("expected nil items after dispose")
	}
}

// ── DataModel ─────────────────────────────────────────────

func Test_CharCollectionDataModel(t *testing.T) {
	cm := New.CharCollectionMap.Items([]string{"abc"})
	model := NewCharCollectionMapDataModelUsing(cm)
	cm2 := NewCharCollectionMapUsingDataModel(model)
	if cm2 == nil {
		t.Fatal("expected non-nil")
	}
}
