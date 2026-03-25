package corestr

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Collection extended coverage ──

func TestCollection_TakeSkipLimit_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
	taken := c.Take(2)
	if taken.Length() != 2 { t.Fatal("expected 2") }
	skipped := c.Skip(2)
	if skipped.Length() != 3 { t.Fatal("expected 3") }

	s := New.SimpleSlice.Lines("a", "b", "c", "d", "e")
	limited := s.Limit(2)
	if len(limited) != 2 { t.Fatal("expected 2") }
	limitAll := s.Limit(-1)
	if len(limitAll) != 5 { t.Fatal("expected 5") }
}

func TestCollection_AddNonEmptyStrings_C10(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmptyStrings("a", "", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddNonEmptyStringsSlice_C10(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmptyStringsSlice([]string{"a", "", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyList_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	list := c.NonEmptyList()
	if len(list) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyListPtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	list := c.NonEmptyListPtr()
	if list == nil || len(*list) != 2 { t.Fatal("unexpected") }
}

func TestCollection_Items_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	items := c.Items()
	if len(items) != 2 { t.Fatal("expected 2") }
}

func TestCollection_ListPtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	lp := c.ListPtr()
	if lp == nil { t.Fatal("expected non-nil") }
}

func TestCollection_ListCopyPtrLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	cp := c.ListCopyPtrLock()
	if cp == nil || len(cp) != 1 { t.Fatal("unexpected") }
}

func TestCollection_Has_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.Has("b") { t.Fatal("expected true") }
	if c.Has("z") { t.Fatal("expected false") }
}

func TestCollection_HasLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if !c.HasLock("a") { t.Fatal("expected true") }
}

func TestCollection_HasPtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.Has("a") { t.Fatal("expected true") }
}

func TestCollection_HasAll_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.HasAll("a", "b") { t.Fatal("expected true") }
	if c.HasAll("a", "z") { t.Fatal("expected false") }
}

func TestCollection_SortedListAsc_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a", "b"})
	sorted := c.SortedListAsc()
	if sorted[0] != "a" { t.Fatal("expected a first") }
}

func TestCollection_SortedAsc_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a", "b"})
	sortedC := c.SortedAsc()
	list := sortedC.List()
	if list[0] != "a" { t.Fatal("expected a first") }
}

func TestCollection_SortedAscLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a"})
	sortedC := c.SortedAscLock()
	list := sortedC.List()
	if list[0] != "a" { t.Fatal("expected a first") }
}

func TestCollection_SortedListDsc_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "c", "b"})
	sorted := c.SortedListDsc()
	if sorted[0] != "c" { t.Fatal("expected c first") }
}

func TestCollection_FilterAndFilterPtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"apple", "banana", "cherry"})
	filtered := c.Filter(func(s string, i int) (string, bool, bool) { return s, len(s) > 5, false })
	if len(filtered) < 2 { t.Fatal("expected >= 2") }
}

func TestCollection_FilterLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb", "ccc"})
	filtered := c.FilterLock(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
	if len(filtered) != 2 { t.Fatal("expected 2") }
}

func TestCollection_FilteredCollection_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
	if fc.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_FilteredCollectionLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false })
	if fc.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_FilterPtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fp := c.FilterPtr(func(s *string, i int) (*string, bool, bool) { return s, len(*s) > 1, false })
	if fp == nil || len(*fp) != 1 { t.Fatal("unexpected") }
}

func TestCollection_FilterPtrLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fp := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) { return s, len(*s) > 1, false })
	if fp == nil || len(*fp) != 1 { t.Fatal("unexpected") }
}

func TestCollection_HashsetAsIs_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "a"})
	hs := c.HashsetAsIs()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestCollection_HashsetWithDoubleLength_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	hs := c.HashsetWithDoubleLength()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestCollection_HashsetLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	hs := c.HashsetLock()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestCollection_NonEmptyItems_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	items := c.NonEmptyItems()
	if len(items) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyItemsPtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	ip := c.NonEmptyItemsPtr()
	if ip == nil || len(ip) != 2 { t.Fatal("unexpected") }
}

func TestCollection_NonEmptyItemsOrNonWhitespace_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", " ", "b"})
	items := c.NonEmptyItemsOrNonWhitespace()
	if len(items) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyItemsOrNonWhitespacePtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", " ", "b"})
	ip := c.NonEmptyItemsOrNonWhitespacePtr()
	if ip == nil || len(ip) != 2 { t.Fatal("unexpected") }
}

func TestCollection_IsContainsPtr_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	v := "a"
	if !c.IsContainsPtr(&v) { t.Fatal("expected true") }
}

func TestCollection_IsContainsAll_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.IsContainsAll("a", "b") { t.Fatal("expected true") }
}

func TestCollection_IsContainsAllLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.IsContainsAllLock("a", "c") { t.Fatal("expected true") }
}

func TestCollection_IsContainsAllSlice_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.IsContainsAllSlice([]string{"a", "b"}) { t.Fatal("expected true") }
}

func TestCollection_GetHashsetPlusHasAll_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})
	if !hasAll || hs == nil { t.Fatal("unexpected") }
}

func TestCollection_HasUsingSensitivity_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"Hello", "World"})
	if !c.HasUsingSensitivity("hello", false) { t.Fatal("expected true case-insensitive") }
	if !c.HasUsingSensitivity("Hello", true) { t.Fatal("expected true case-sensitive") }
}

func TestCollection_New_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	newC := c.New("c", "d")
	if newC.Length() != 4 { t.Fatal("expected 4") }
}

func TestCollection_ExpandSlicePlusAdd_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.ExpandSlicePlusAdd([]string{"b", "c"}, func(line string) []string { return []string{line} })
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_MergeSlicesOfSlice_C10(t *testing.T) {
	c := New.Collection.Empty()
	c.MergeSlicesOfSlice([]string{"a", "b"}, []string{"c"})
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_GetAllExceptCollection_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	except := New.Collection.Strings([]string{"b"})
	result := c.GetAllExceptCollection(except)
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestCollection_GetAllExcept_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	result := c.GetAllExcept([]string{"b"})
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestCollection_Joins_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	s := c.Joins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_NonEmptyJoins_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	s := c.NonEmptyJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_NonWhitespaceJoins_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", " ", "b"})
	s := c.NonWhitespaceJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_String_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.String() == "" { t.Fatal("expected non-empty") }
}

func TestCollection_StringLock_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestCollection_SummaryString_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.SummaryString(1) == "" { t.Fatal("expected non-empty") }
}

func TestCollection_SummaryStringWithHeader_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := c.SummaryStringWithHeader("header")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_CsvLines_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a,b", "c,d"})
	_ = c.CsvLines()
}

func TestCollection_Csv_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a,b", "c,d"})
	_ = c.Csv()
}

func TestCollection_AddCapacity_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.AddCapacity(10)
}

func TestCollection_Resize_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	c.Resize(5)
}

func TestCollection_CharCollectionMap_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"apple", "banana"})
	ccm := c.CharCollectionMap()
	if ccm == nil { t.Fatal("expected non-nil") }
}

func TestCollection_Join_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	s := c.Join(",")
	if s != "a,b" { t.Fatal("unexpected") }
}

func TestCollection_JoinLine_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	s := c.JoinLine()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_JsonAndInterfaces_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonMarshaller()
	_ = c.AsJsonContractsBinder()
	_, err := c.Serialize()
	if err != nil { t.Fatal(err) }
	var out Collection
	if err := out.Deserialize(c.JsonPtr()); err != nil { t.Fatal(err) }
}

func TestCollection_ClearDispose_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	c.Clear()
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Add("x")
	c.Dispose()
}

func TestCollection_AddFuncResult_C10(t *testing.T) {
	c := New.Collection.Empty()
	c.AddFuncResult(func() string { return "hello" })
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddStringsByFuncChecking_C10(t *testing.T) {
	c := New.Collection.Empty()
	c.AddStringsByFuncChecking([]string{"a", "bb", "c", "dd"}, func(s string) bool { return len(s) > 1 })
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_ParseInjectUsingJson_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	jsonResult := c.JsonPtr()
	c2 := New.Collection.Empty()
	_, err := c2.ParseInjectUsingJson(jsonResult)
	if err != nil { t.Fatal(err) }
}

func TestCollection_JsonParseSelfInject_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Empty()
	err := c2.JsonParseSelfInject(c.JsonPtr())
	if err != nil { t.Fatal(err) }
}

// ── Hashmap extended coverage ──

func TestHashmap_HasItems_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasItems() { t.Fatal("expected true") }
}

func TestHashmap_Collection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := hm.Collection()
	if c == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_IsEmptyLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	if !hm.IsEmptyLock() { t.Fatal("expected true") }
}

func TestHashmap_AddOrUpdateLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateLock("a", "1")
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ContainsAndHas_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.Contains("a") { t.Fatal("expected true") }
	if !hm.ContainsLock("a") { t.Fatal("expected true") }
	if hm.IsKeyMissing("a") { t.Fatal("expected false") }
	if hm.IsKeyMissingLock("a") { t.Fatal("expected false") }
	if !hm.HasLock("a") { t.Fatal("expected true") }
}

func TestHashmap_HasAllStrings_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	if !hm.HasAllStrings("a", "b") { t.Fatal("expected true") }
	if hm.HasAllStrings("a", "c") { t.Fatal("expected false") }
}

func TestHashmap_HasAll_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasAll("a") { t.Fatal("expected true") }
}

func TestHashmap_HasAny_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasAny("a", "b") { t.Fatal("expected true") }
}

func TestHashmap_HasWithLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasWithLock("a") { t.Fatal("expected true") }
}

func TestHashmap_Keys_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys := hm.Keys()
	if len(keys) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysCollection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kc := hm.KeysCollection()
	if kc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_AllKeys_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys := hm.AllKeys()
	if len(keys) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys := hm.KeysLock()
	if len(keys) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ValuesList_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vals := hm.ValuesList()
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ValuesCollection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vc := hm.ValuesCollection()
	if vc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesHashset_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vh := hm.ValuesHashset()
	if vh == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesCollectionLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vc := hm.ValuesCollectionLock()
	if vc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesHashsetLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vh := hm.ValuesHashsetLock()
	if vh == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_KeysValuesCollection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys, vals := hm.KeysValuesCollection()
	if keys == nil || vals == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_KeysValuesList_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	k, v := hm.KeysValuesList()
	if len(k) != 1 || len(v) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysValuePairs_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvp := hm.KeysValuePairs()
	if len(kvp) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysValuePairsCollection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvpc := hm.KeysValuePairsCollection()
	if kvpc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_KeysValuesListLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	k, v := hm.KeysValuesListLock()
	if len(k) != 1 || len(v) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_LengthLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if hm.LengthLock() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_Remove_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.Remove("a")
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_RemoveWithLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.RemoveWithLock("a")
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_String_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if hm.String() == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_StringLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if hm.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_ItemsCopyLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	cp := hm.ItemsCopyLock()
	if cp == nil || len(*cp) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_SafeItems_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	si := hm.SafeItems()
	if si == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesListCopyLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vl := hm.ValuesListCopyLock()
	if len(vl) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ValuesToLower_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "HELLO")
	hm.ValuesToLower()
}

func TestHashmap_KeysToLower_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("ABC", "1")
	hm.KeysToLower()
}

func TestHashmap_IsEqual_C10(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	if !hm1.IsEqual(*hm2) { t.Fatal("expected equal") }
}

func TestHashmap_IsEqualPtrLock_C10(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	if !hm1.IsEqualPtrLock(hm2) { t.Fatal("expected equal") }
}

func TestHashmap_Clone_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	cloned := hm.Clone()
	if cloned.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetValue_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	v, found := hm.GetValue("a")
	if !found || v != "1" { t.Fatal("expected 1") }
}

func TestHashmap_Join_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	s := hm.Join(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_JoinKeys_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	s := hm.JoinKeys(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_Dispose_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.Dispose()
}

func TestHashmap_ToError_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("err", "something")
	e := hm.ToError(",")
	if e == nil { t.Fatal("expected error") }
}

func TestHashmap_ToDefaultError_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("err", "something")
	e := hm.ToDefaultError()
	if e == nil { t.Fatal("expected error") }
}

func TestHashmap_KeyValStringLines_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	lines := hm.KeyValStringLines()
	if len(lines) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_JsonInterfaces_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	_ = hm.JsonModelAny()
	_ = hm.Json()
	_ = hm.JsonPtr()
	_ = hm.AsJsoner()
	_ = hm.AsJsonContractsBinder()
	_ = hm.AsJsonParseSelfInjector()
	_ = hm.AsJsonMarshaller()
	_, err := hm.Serialize()
	if err != nil { t.Fatal(err) }
	var hm2 Hashmap
	if err := hm2.Deserialize(hm.JsonPtr()); err != nil { t.Fatal(err) }
}

func TestHashmap_ParseInjectUsingJson_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	json := hm.Json()
	hm2 := New.Hashmap.Empty()
	_, err := hm2.ParseInjectUsingJson(json.Ptr())
	if err != nil { t.Fatal(err) }
}

func TestHashmap_JsonParseSelfInject_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Empty()
	if err := hm2.JsonParseSelfInject(hm.JsonPtr()); err != nil { t.Fatal(err) }
}

func TestHashmap_AddOrUpdateHashmap_C10(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdateHashmap(hm1)
	if hm2.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddOrUpdateMap_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddsOrUpdates_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdates(KeyValuePair{Key: "a", Value: "1"}, KeyValuePair{Key: "b", Value: "2"})
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddOrUpdateCollection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	keys := New.Collection.Strings([]string{"a"})
	values := New.Collection.Strings([]string{"1"})
	hm.AddOrUpdateCollection(keys, values)
}

func TestHashmap_Set_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.Set("a", "1")
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_SetTrim_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.SetTrim(" a ", " 1 ")
	v, ok := hm.GetValue("a")
	if !ok || v != "1" { t.Fatal("expected trimmed") }
}

func TestHashmap_SetBySplitter_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.SetBySplitter("key=value", "=")
	v, ok := hm.GetValue("key")
	if !ok || v != "value" { t.Fatal("expected value") }
}

func TestHashmap_DiffRaw_C10(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm1.AddOrUpdate("b", "2")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	diff := hm1.DiffRaw(hm2.Items())
	_ = diff
}

func TestHashmap_Diff_C10(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	diff := hm1.Diff(hm2)
	_ = diff
}

func TestHashmap_HasAllCollectionItems_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := New.Collection.Strings([]string{"a"})
	if !hm.HasAllCollectionItems(c) { t.Fatal("expected true") }
}

func TestHashmap_GetKeysFilteredItems_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	filtered := hm.GetKeysFilteredItems(func(k string, _ int) (string, bool, bool) {
		return k, k == "a", false
	})
	if len(filtered) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetKeysFilteredCollection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	fc := hm.GetKeysFilteredCollection(func(k string, _ int) (string, bool, bool) {
		return k, k == "a", false
	})
	if fc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_GetValuesExceptKeysInHashset_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	hs := New.Hashset.Strings([]string{"a"})
	vals := hm.GetValuesExceptKeysInHashset(hs)
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetValuesKeysExcept_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	vals := hm.GetValuesKeysExcept([]string{"a"})
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetAllExceptCollection_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	c := New.Collection.Strings([]string{"a"})
	result := hm.GetAllExceptCollection(c)
	if result == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ToStringsUsingCompiler_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	strs := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
	if len(strs) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ConcatNew_C10(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("b", "2")
	result := hm1.ConcatNew(hm2)
	if result.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_ConcatNewUsingMaps_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	m := map[string]string{"b": "2"}
	result := hm.ConcatNewUsingMaps(m)
	if result.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddsOrUpdatesUsingFilter_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesUsingFilter(func(k, v string) bool { return k != "skip" }, "a", "1", "skip", "2", "b", "3")
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddsOrUpdatesAnyUsingFilter_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	items := map[string]any{"a": "1", "b": 2}
	hm.AddsOrUpdatesAnyUsingFilter(items, func(k string, v any) bool { return k == "a" })
}

func TestHashmap_AddsOrUpdatesAnyUsingFilterLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	items := map[string]any{"a": "1"}
	hm.AddsOrUpdatesAnyUsingFilterLock(items, func(k string, v any) bool { return true })
}

func TestHashmap_AddOrUpdateKeyStrValInt_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValInt("a", 42)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddOrUpdateKeyStrValFloat_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat("a", 3.14)
}

func TestHashmap_AddOrUpdateKeyStrValFloat64_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat64("a", 3.14)
}

func TestHashmap_AddOrUpdateKeyStrValAny_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValAny("a", "hello")
}

func TestHashmap_AddOrUpdateKeyValueAny_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyValueAny("a", 42)
}

func TestHashmap_AddOrUpdateKeyVal_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyVal("a", "1")
}

func TestHashmap_AddOrUpdateWithWgLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateWithWgLock("a", "1")
}

func TestHashmap_AddOrUpdateStringsPtrWgLock_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	items := []string{"a", "1", "b", "2"}
	hm.AddOrUpdateStringsPtrWgLock(&items)
}

func TestHashmap_AddOrUpdateKeyAnyValues_C10(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyAnyValues(map[string]any{"a": 1, "b": "hello"})
}

// ── Hashset extended coverage ──

func TestHashset_AddAndHas_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.Has("a") { t.Fatal("expected true") }
	if hs.Has("z") { t.Fatal("expected false") }
	hs.Add("c")
	if hs.Length() != 3 { t.Fatal("expected 3") }
}

func TestHashset_HasAll_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b", "c"})
	if !hs.HasAll("a", "b") { t.Fatal("expected true") }
	if hs.HasAll("a", "z") { t.Fatal("expected false") }
}

func TestHashset_HasAny_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.HasAny("a", "z") { t.Fatal("expected true") }
	if hs.HasAny("x", "z") { t.Fatal("expected false") }
}

func TestHashset_Remove_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	hs.Remove("a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_Adds_C10(t *testing.T) {
	hs := New.Hashset.Empty()
	hs.Adds("a", "b", "c")
	if hs.Length() != 3 { t.Fatal("expected 3") }
}

func TestHashset_List_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	list := hs.List()
	if len(list) != 1 { t.Fatal("expected 1") }
}

func TestHashset_ListSortedAsc_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a", "b"})
	sorted := hs.ListSortedAsc()
	if sorted[0] != "a" { t.Fatal("expected a first") }
}

func TestHashset_Collection_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	c := hs.Collection()
	if c == nil { t.Fatal("expected non-nil") }
}

func TestHashset_Clone_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	cloned := hs.Clone()
	if cloned.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_String_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.String() == "" { t.Fatal("expected non-empty") }
}

func TestHashset_Join_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	s := hs.Join(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_JoinLine_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	s := hs.JoinLine()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_Dispose_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Dispose()
}

func TestHashset_JsonInterfaces_C10(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	_ = hs.Json()
	_ = hs.JsonPtr()
	_ = hs.AsJsoner()
	_ = hs.AsJsonContractsBinder()
	_ = hs.AsJsonParseSelfInjector()
	_ = hs.AsJsonMarshaller()
	_ = hs.Serialize()
	_ = hs.Deserialize()
}

// ── newCreator paths ──

func TestNewCreator_CollectionEmpty_C10(t *testing.T) {
	c := New.Collection.Empty()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestNewCreator_CollectionCloneStrings_C10(t *testing.T) {
	c := New.Collection.CloneStrings([]string{"a", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_CollectionCreate_C10(t *testing.T) {
	c := New.Collection.Create(5)
	if c == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_CollectionStringsPlusCap_C10(t *testing.T) {
	c := New.Collection.StringsPlusCap(10, []string{"a"})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCreator_CollectionCapStrings_C10(t *testing.T) {
	c := New.Collection.CapStrings(10, "a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_CollectionLenCap_C10(t *testing.T) {
	c := New.Collection.LenCap(0, 10)
	if c == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_CollectionLineDefault_C10(t *testing.T) {
	c := New.Collection.LineDefault("a\nb\nc")
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestNewCreator_CollectionLineUsingSep_C10(t *testing.T) {
	c := New.Collection.LineUsingSep("a,b,c", ",")
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestNewCreator_CollectionStringsOptions_C10(t *testing.T) {
	c := New.Collection.StringsOptions(true, []string{"a", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_HashmapEmpty_C10(t *testing.T) {
	hm := New.Hashmap.Empty()
	if hm == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_HashmapMapWithCap_C10(t *testing.T) {
	hm := New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCreator_HashsetEmpty_C10(t *testing.T) {
	hs := New.Hashset.Empty()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_HashsetStringsOption_C10(t *testing.T) {
	hs := New.Hashset.StringsOption(true, []string{"a"})
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCreator_HashsetStringSpreadItems_C10(t *testing.T) {
	hs := New.Hashset.StringsSpreadItems("a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_HashsetUsingCollection_C10(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	hs := New.Hashset.UsingCollection(c)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

// ── emptyCreator paths ──

func TestEmptyCreator_LinkedList_C10(t *testing.T) {
	ll := Empty.LinkedList()
	if ll == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyValuePair_C10(t *testing.T) {
	kvp := Empty.KeyValuePair()
	if kvp == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyAnyValuePair_C10(t *testing.T) {
	kavp := Empty.KeyAnyValuePair()
	if kavp == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyValueCollection_C10(t *testing.T) {
	kvc := Empty.KeyValueCollection()
	if kvc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_LinkedCollections_C10(t *testing.T) {
	lc := Empty.LinkedCollections()
	if lc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_LeftRight_C10(t *testing.T) {
	lr := Empty.LeftRight()
	_ = lr
}

func TestEmptyCreator_SimpleStringOnce_C10(t *testing.T) {
	s := Empty.SimpleStringOnce()
	_ = s
}

func TestEmptyCreator_SimpleStringOncePtr_C10(t *testing.T) {
	sp := Empty.SimpleStringOncePtr()
	if sp == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_HashsetsCollection_C10(t *testing.T) {
	hsc := Empty.HashsetsCollection()
	if hsc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_CharCollectionMap_C10(t *testing.T) {
	ccm := Empty.CharCollectionMap()
	if ccm == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyValuesCollection_C10(t *testing.T) {
	kvc := Empty.KeyValuesCollection()
	if kvc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_CollectionsOfCollection_C10(t *testing.T) {
	coc := Empty.CollectionsOfCollection()
	if coc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_CharHashsetMap_C10(t *testing.T) {
	chm := Empty.CharHashsetMap()
	if chm == nil { t.Fatal("expected non-nil") }
}

// ── AnyToString ──

func TestAnyToString_C10(t *testing.T) {
	s := AnyToString(42)
	if s == "" { t.Fatal("expected non-empty") }
	s2 := AnyToString("hello")
	if s2 != "hello" { t.Fatal("unexpected") }
	s3 := AnyToString(nil)
	_ = s3
}

// ── AllIndividualStringsOfStringsLength ──

func TestAllIndividualStringsOfStringsLength_C10(t *testing.T) {
	length := AllIndividualStringsOfStringsLength([][]string{{"a", "b"}, {"c"}})
	if length != 3 { t.Fatal("expected 3") }
}

// ── AllIndividualsLengthOfSimpleSlices ──

func TestAllIndividualsLengthOfSimpleSlices_C10(t *testing.T) {
	ss1 := New.SimpleSlice.Lines("a", "b")
	ss2 := New.SimpleSlice.Lines("c")
	length := AllIndividualsLengthOfSimpleSlices([]*SimpleSlice{ss1, ss2})
	if length != 3 { t.Fatal("expected 3") }
}
