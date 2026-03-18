package corestr

import (
	"encoding/json"
	"testing"
	"time"
)

// ── Creators ──────────────────────────────────────────────

func Test_CharHashsetMap_NewCap(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(20, 10)
	if hsm == nil || hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_CharHashsetMap_NewCapItems(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(20, 10, "abc", "axy", "xyz")
	if hsm.Length() != 2 {
		t.Fatalf("expected 2, got %d", hsm.Length())
	}
}

func Test_CharHashsetMap_NewStrings(t *testing.T) {
	hsm := New.CharHashsetMap.Strings(10, []string{"abc", "xyz"})
	if hsm.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharHashsetMap_NewStrings_Nil(t *testing.T) {
	hsm := New.CharHashsetMap.Strings(10, nil)
	if hsm == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if hsm == nil {
		t.Fatal("expected non-nil")
	}
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── GetChar / GetCharOf ───────────────────────────────────

func Test_CharHashsetMap_GetChar(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	if hsm.GetChar("hello") != 'h' {
		t.Fatal("expected h")
	}
	if hsm.GetChar("") != emptyChar {
		t.Fatal("expected emptyChar")
	}
}

func Test_CharHashsetMap_GetCharOf(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	if hsm.GetCharOf("abc") != 'a' {
		t.Fatal("expected a")
	}
	if hsm.GetCharOf("") != emptyChar {
		t.Fatal("expected emptyChar")
	}
}

// ── GetCharsGroups ────────────────────────────────────────

func Test_CharHashsetMap_GetCharsGroups(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	result := hsm.GetCharsGroups("abc", "axy", "xyz")
	if result.Length() != 2 {
		t.Fatal("expected 2 groups")
	}
}

func Test_CharHashsetMap_GetCharsGroups_Empty(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	result := hsm.GetCharsGroups()
	if result != hsm {
		t.Fatal("expected same ref on empty")
	}
}

// ── Add / AddStrings ──────────────────────────────────────

func Test_CharHashsetMap_Add(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.Add("alpha")
	hsm.Add("avocado")
	hsm.Add("alpha") // duplicate
	hsm.Add("beta")
	if hsm.Length() != 2 {
		t.Fatalf("expected 2, got %d", hsm.Length())
	}
	if hsm.AllLengthsSum() != 3 {
		t.Fatalf("expected sum 3, got %d", hsm.AllLengthsSum())
	}
}

func Test_CharHashsetMap_AddStrings(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddStrings("x1", "x2", "y1")
	if hsm.LengthOf('x') != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharHashsetMap_AddStrings_Empty(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddStrings()
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_CharHashsetMap_AddLock(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddLock("hello")
	hsm.AddLock("help")
	if hsm.LengthLock() != 1 {
		t.Fatal("expected 1 group")
	}
}

func Test_CharHashsetMap_AddStringsLock(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddStringsLock("abc", "axy")
	if hsm.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CharHashsetMap_AddStringsLock_Empty(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddStringsLock()
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddSameStartingCharItems ──────────────────────────────

func Test_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddSameStartingCharItems('a', []string{"abc", "axy"})
	if hsm.LengthOfHashsetFromFirstChar("a") != 2 {
		t.Fatal("expected 2")
	}
	hsm.AddSameStartingCharItems('a', []string{"azz"})
	if hsm.LengthOfHashsetFromFirstChar("a") != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CharHashsetMap_AddSameStartingCharItems_Empty(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddSameStartingCharItems('a', []string{})
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddCollectionItems ────────────────────────────────────

func Test_CharHashsetMap_AddCollectionItems(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	col := New.Collection.Strings([]string{"abc", "xyz"})
	hsm.AddCollectionItems(col)
	if hsm.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharHashsetMap_AddCollectionItems_Nil(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddCollectionItems(nil)
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddCharCollectionMapItems ─────────────────────────────

func Test_CharHashsetMap_AddCharCollectionMapItems(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	cm := New.CharCollectionMap.Items([]string{"abc", "xyz"})
	hsm.AddCharCollectionMapItems(cm)
	if hsm.AllLengthsSum() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharHashsetMap_AddCharCollectionMapItems_Nil(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddCharCollectionMapItems(nil)
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddCollectionItemsAsyncLock ───────────────────────────

func Test_CharHashsetMap_AddCollectionItemsAsyncLock(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	col := New.Collection.Strings([]string{"abc", "xyz"})
	done := make(chan bool)
	hsm.AddCollectionItemsAsyncLock(col, func(chm *CharHashsetMap) {
		done <- true
	})
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("timeout")
	}
}

func Test_CharHashsetMap_AddCollectionItemsAsyncLock_Nil(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddCollectionItemsAsyncLock(nil, nil)
	// should return immediately
}

// ── Has / HasWithHashset ──────────────────────────────────

func Test_CharHashsetMap_Has(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "foo", "far", "bar")
	if !hsm.Has("foo") {
		t.Fatal("expected has foo")
	}
	if hsm.Has("baz") {
		t.Fatal("expected no baz")
	}
	if hsm.Has("zzz") {
		t.Fatal("expected no zzz")
	}
}

func Test_CharHashsetMap_Has_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if hsm.Has("x") {
		t.Fatal("expected false")
	}
}

func Test_CharHashsetMap_HasWithHashset(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "foo", "far")
	has, hs := hsm.HasWithHashset("foo")
	if !has || hs.IsEmpty() {
		t.Fatal("expected found")
	}
	has2, _ := hsm.HasWithHashset("zzz")
	if has2 {
		t.Fatal("expected not found")
	}
}

func Test_CharHashsetMap_HasWithHashset_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	has, hs := hsm.HasWithHashset("x")
	if has || hs == nil {
		t.Fatal("expected false, non-nil")
	}
}

func Test_CharHashsetMap_HasWithHashsetLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "foo")
	has, hs := hsm.HasWithHashsetLock("foo")
	if !has || hs == nil {
		t.Fatal("expected found")
	}
	has2, _ := hsm.HasWithHashsetLock("zzz")
	if has2 {
		t.Fatal("expected not found")
	}
}

func Test_CharHashsetMap_HasWithHashsetLock_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	has, _ := hsm.HasWithHashsetLock("x")
	if has {
		t.Fatal("expected false")
	}
}

// ── LengthOf / LengthOfLock ──────────────────────────────

func Test_CharHashsetMap_LengthOf(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "axy")
	if hsm.LengthOf('a') != 2 {
		t.Fatal("expected 2")
	}
	if hsm.LengthOf('z') != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CharHashsetMap_LengthOf_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if hsm.LengthOf('a') != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CharHashsetMap_LengthOfLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.LengthOfLock('a') != 1 {
		t.Fatal("expected 1")
	}
	if hsm.LengthOfLock('z') != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CharHashsetMap_LengthOfLock_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if hsm.LengthOfLock('a') != 0 {
		t.Fatal("expected 0")
	}
}

// ── AllLengthsSum / AllLengthsSumLock ─────────────────────

func Test_CharHashsetMap_AllLengthsSum(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "a1", "a2", "b1")
	if hsm.AllLengthsSum() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CharHashsetMap_AllLengthsSum_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if hsm.AllLengthsSum() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CharHashsetMap_AllLengthsSumLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "a1", "b1")
	if hsm.AllLengthsSumLock() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharHashsetMap_AllLengthsSumLock_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if hsm.AllLengthsSumLock() != 0 {
		t.Fatal("expected 0")
	}
}

// ── IsEmpty / HasItems / IsEmptyLock ──────────────────────

func Test_CharHashsetMap_IsEmpty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if !hsm.IsEmpty() || hsm.HasItems() {
		t.Fatal("expected empty")
	}
	hsm.Add("x")
	if hsm.IsEmpty() || !hsm.HasItems() {
		t.Fatal("expected has items")
	}
}

func Test_CharHashsetMap_IsEmptyLock(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	if !hsm.IsEmptyLock() {
		t.Fatal("expected empty")
	}
}

// ── IsEquals / IsEqualsLock ───────────────────────────────

func Test_CharHashsetMap_IsEquals(t *testing.T) {
	hsm1 := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	hsm2 := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	if !hsm1.IsEquals(hsm2) {
		t.Fatal("expected equal")
	}
}

func Test_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.IsEquals(nil) {
		t.Fatal("expected not equal")
	}
}

func Test_CharHashsetMap_IsEquals_SameRef(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if !hsm.IsEquals(hsm) {
		t.Fatal("expected same ref equal")
	}
}

func Test_CharHashsetMap_IsEquals_BothEmpty(t *testing.T) {
	hsm1 := Empty.CharHashsetMap()
	hsm2 := Empty.CharHashsetMap()
	if !hsm1.IsEquals(hsm2) {
		t.Fatal("expected equal")
	}
}

func Test_CharHashsetMap_IsEquals_OneEmpty(t *testing.T) {
	hsm1 := New.CharHashsetMap.CapItems(10, 5, "abc")
	hsm2 := Empty.CharHashsetMap()
	if hsm1.IsEquals(hsm2) {
		t.Fatal("expected not equal")
	}
}

func Test_CharHashsetMap_IsEquals_DiffLength(t *testing.T) {
	hsm1 := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	hsm2 := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm1.IsEquals(hsm2) {
		t.Fatal("expected not equal")
	}
}

func Test_CharHashsetMap_IsEquals_DiffContent(t *testing.T) {
	hsm1 := New.CharHashsetMap.CapItems(10, 5, "abc")
	hsm2 := New.CharHashsetMap.CapItems(10, 5, "axy")
	if hsm1.IsEquals(hsm2) {
		t.Fatal("expected not equal")
	}
}

func Test_CharHashsetMap_IsEquals_MissingKey(t *testing.T) {
	hsm1 := New.CharHashsetMap.CapItems(10, 5, "abc")
	hsm2 := New.CharHashsetMap.CapItems(10, 5, "xyz")
	if hsm1.IsEquals(hsm2) {
		t.Fatal("expected not equal")
	}
}

func Test_CharHashsetMap_IsEqualsLock(t *testing.T) {
	hsm1 := New.CharHashsetMap.CapItems(10, 5, "abc")
	hsm2 := New.CharHashsetMap.CapItems(10, 5, "abc")
	if !hsm1.IsEqualsLock(hsm2) {
		t.Fatal("expected equal")
	}
}

// ── GetHashset / GetHashsetLock ───────────────────────────

func Test_CharHashsetMap_GetHashset(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := hsm.GetHashset("a", false)
	if hs == nil {
		t.Fatal("expected non-nil")
	}
	// missing, no create
	hs2 := hsm.GetHashset("z", false)
	if hs2 != nil {
		t.Fatal("expected nil")
	}
	// missing, create
	hs3 := hsm.GetHashset("z", true)
	if hs3 == nil {
		t.Fatal("expected new hashset")
	}
}

func Test_CharHashsetMap_GetHashsetLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := hsm.GetHashsetLock(false, "a")
	if hs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_GetHashsetByChar(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := hsm.GetHashsetByChar('a')
	if hs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_HashsetByChar(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := hsm.HashsetByChar('a')
	if hs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_HashsetByCharLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := hsm.HashsetByCharLock('a')
	if hs == nil {
		t.Fatal("expected non-nil")
	}
	hs2 := hsm.HashsetByCharLock('z')
	if hs2 == nil || !hs2.IsEmpty() {
		t.Fatal("expected empty hashset")
	}
}

func Test_CharHashsetMap_HashsetByStringFirstChar(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := hsm.HashsetByStringFirstChar("abc")
	if hs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_HashsetByStringFirstCharLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := hsm.HashsetByStringFirstCharLock("abc")
	if hs == nil {
		t.Fatal("expected non-nil")
	}
}

// ── AddSameCharsCollection ────────────────────────────────

func Test_CharHashsetMap_AddSameCharsCollection(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	col := New.Collection.Strings([]string{"abc", "axy"})
	result := hsm.AddSameCharsCollection("a", col)
	if result == nil || result.Length() != 2 {
		t.Fatal("expected 2")
	}
	// existing + more
	col2 := New.Collection.Strings([]string{"azz"})
	result2 := hsm.AddSameCharsCollection("a", col2)
	if result2.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CharHashsetMap_AddSameCharsCollection_NilCol(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	result := hsm.AddSameCharsCollection("a", nil)
	if result == nil {
		t.Fatal("expected new hashset")
	}
}

func Test_CharHashsetMap_AddSameCharsCollection_ExistingNilAdd(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	result := hsm.AddSameCharsCollection("a", nil)
	if result == nil {
		t.Fatal("expected existing")
	}
}

// ── AddSameCharsHashset ───────────────────────────────────

func Test_CharHashsetMap_AddSameCharsHashset(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hs := New.Hashset.Strings([]string{"abc", "axy"})
	result := hsm.AddSameCharsHashset("a", hs)
	if result == nil || result.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharHashsetMap_AddSameCharsHashset_NilHashset(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	result := hsm.AddSameCharsHashset("a", nil)
	if result == nil {
		t.Fatal("expected new hashset")
	}
}

func Test_CharHashsetMap_AddSameCharsHashset_ExistingNilAdd(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	result := hsm.AddSameCharsHashset("a", nil)
	if result == nil {
		t.Fatal("expected existing")
	}
}

func Test_CharHashsetMap_AddSameCharsHashset_AddToExisting(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := New.Hashset.Strings([]string{"axy"})
	result := hsm.AddSameCharsHashset("a", hs)
	if result.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ── AddSameCharsCollectionLock ────────────────────────────

func Test_CharHashsetMap_AddSameCharsCollectionLock(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	col := New.Collection.Strings([]string{"abc"})
	result := hsm.AddSameCharsCollectionLock("a", col)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_AddSameCharsCollectionLock_NilCol(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	result := hsm.AddSameCharsCollectionLock("a", nil)
	if result == nil {
		t.Fatal("expected new hashset")
	}
}

func Test_CharHashsetMap_AddSameCharsCollectionLock_ExistingNilAdd(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	result := hsm.AddSameCharsCollectionLock("a", nil)
	if result == nil {
		t.Fatal("expected existing")
	}
}

func Test_CharHashsetMap_AddSameCharsCollectionLock_AddToExisting(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	col := New.Collection.Strings([]string{"axy"})
	result := hsm.AddSameCharsCollectionLock("a", col)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

// ── AddHashsetLock ────────────────────────────────────────

func Test_CharHashsetMap_AddHashsetLock(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hs := New.Hashset.Strings([]string{"abc"})
	result := hsm.AddHashsetLock("a", hs)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_AddHashsetLock_NilHashset(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	result := hsm.AddHashsetLock("a", nil)
	if result == nil {
		t.Fatal("expected new hashset")
	}
}

func Test_CharHashsetMap_AddHashsetLock_ExistingNilAdd(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	result := hsm.AddHashsetLock("a", nil)
	if result == nil {
		t.Fatal("expected existing")
	}
}

func Test_CharHashsetMap_AddHashsetLock_AddToExisting(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hs := New.Hashset.Strings([]string{"axy"})
	result := hsm.AddHashsetLock("a", hs)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

// ── AddHashsetItems ───────────────────────────────────────

func Test_CharHashsetMap_AddHashsetItems(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hs := New.Hashset.Strings([]string{"abc", "xyz"})
	hsm.AddHashsetItems(hs)
	if hsm.AllLengthsSum() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CharHashsetMap_AddHashsetItems_Empty(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hs := New.Hashset.Empty()
	hsm.AddHashsetItems(hs)
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

// ── AddHashsetItemsAsyncLock ──────────────────────────────

func Test_CharHashsetMap_AddHashsetItemsAsyncLock_Nil(t *testing.T) {
	hsm := New.CharHashsetMap.Cap(10, 5)
	hsm.AddHashsetItemsAsyncLock(nil, nil)
}

// ── HashsetsCollection ────────────────────────────────────

func Test_CharHashsetMap_HashsetsCollection(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	hsc := hsm.HashsetsCollection()
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_HashsetsCollection_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	hsc := hsm.HashsetsCollection()
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	hsc := hsm.HashsetsCollectionByChars('a', 'x')
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_HashsetsCollectionByChars_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	hsc := hsm.HashsetsCollectionByChars('a')
	if hsc == nil {
		t.Fatal("expected non-nil empty")
	}
}

func Test_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	hsc := hsm.HashsetsCollectionByStringsFirstChar("abc", "xyz")
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_HashsetsCollectionByStringsFirstChar_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	hsc := hsm.HashsetsCollectionByStringsFirstChar("abc")
	if hsc == nil {
		t.Fatal("expected non-nil")
	}
}

// ── List / SortedListAsc / SortedListDsc ──────────────────

func Test_CharHashsetMap_List(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	list := hsm.List()
	if len(list) != 2 {
		t.Fatalf("expected 2, got %d", len(list))
	}
}

func Test_CharHashsetMap_SortedListAsc(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "cherry", "apple", "banana")
	sorted := hsm.SortedListAsc()
	if len(sorted) != 3 {
		t.Fatal("expected 3")
	}
	if sorted[0] != "apple" {
		t.Fatal("expected apple first")
	}
}

func Test_CharHashsetMap_SortedListDsc(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "cherry", "apple", "banana")
	sorted := hsm.SortedListDsc()
	if len(sorted) != 3 {
		t.Fatal("expected 3")
	}
	if sorted[0] != "cherry" {
		t.Fatal("expected cherry first")
	}
}

// ── GetMap / GetCopyMapLock ───────────────────────────────

func Test_CharHashsetMap_GetMap(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.GetMap() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_GetCopyMapLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	m := hsm.GetCopyMapLock()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CharHashsetMap_GetCopyMapLock_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	m := hsm.GetCopyMapLock()
	if len(m) != 0 {
		t.Fatal("expected empty")
	}
}

// ── String / SummaryString ────────────────────────────────

func Test_CharHashsetMap_String(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CharHashsetMap_SummaryString(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.SummaryString() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CharHashsetMap_StringLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.StringLock() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CharHashsetMap_SummaryStringLock(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.SummaryStringLock() == "" {
		t.Fatal("expected non-empty")
	}
}

// ── Print ─────────────────────────────────────────────────

func Test_CharHashsetMap_Print_Skip(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hsm.Print(false)
}

func Test_CharHashsetMap_PrintLock_Skip(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hsm.PrintLock(false)
}

// ── JSON ──────────────────────────────────────────────────

func Test_CharHashsetMap_JsonModel(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.JsonModel() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_JsonModelAny(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.JsonModelAny() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_MarshalUnmarshalJSON(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	data, err := json.Marshal(hsm)
	if err != nil {
		t.Fatal(err)
	}
	hsm2 := New.CharHashsetMap.Cap(10, 5)
	err = json.Unmarshal(data, hsm2)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CharHashsetMap_Json(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	result := hsm.Json()
	if result.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CharHashsetMap_JsonPtr(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	result := hsm.JsonPtr()
	if result == nil || result.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CharHashsetMap_ParseInjectUsingJson(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	jr := hsm.JsonPtr()
	hsm2 := New.CharHashsetMap.Cap(10, 5)
	_, err := hsm2.ParseInjectUsingJson(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CharHashsetMap_ParseInjectUsingJsonMust(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	jr := hsm.JsonPtr()
	hsm2 := New.CharHashsetMap.Cap(10, 5)
	result := hsm2.ParseInjectUsingJsonMust(jr)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_JsonParseSelfInject(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	jr := hsm.JsonPtr()
	hsm2 := New.CharHashsetMap.Cap(10, 5)
	err := hsm2.JsonParseSelfInject(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CharHashsetMap_AsJsonContractsBinder(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.AsJsonContractsBinder() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_AsJsoner(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.AsJsoner() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_AsJsonMarshaller(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.AsJsonMarshaller() == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CharHashsetMap_AsJsonParseSelfInjector(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	if hsm.AsJsonParseSelfInjector() == nil {
		t.Fatal("expected non-nil")
	}
}

// ── RemoveAll / Clear ─────────────────────────────────────

func Test_CharHashsetMap_RemoveAll(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc", "xyz")
	hsm.RemoveAll()
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_CharHashsetMap_Clear(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	hsm.Clear()
	if hsm.HasItems() {
		t.Fatal("expected empty")
	}
}

func Test_CharHashsetMap_Clear_Empty(t *testing.T) {
	hsm := Empty.CharHashsetMap()
	hsm.Clear()
}

// ── DataModel ─────────────────────────────────────────────

func Test_CharHashsetDataModel(t *testing.T) {
	hsm := New.CharHashsetMap.CapItems(10, 5, "abc")
	model := NewCharHashsetMapDataModelUsing(hsm)
	hsm2 := NewCharHashsetMapUsingDataModel(model)
	if hsm2 == nil {
		t.Fatal("expected non-nil")
	}
}
