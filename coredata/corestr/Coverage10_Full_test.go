package corestr

import (
	"testing"
)

// ── Collection extended coverage ──

func TestCollection_TakeSkipLimit(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
	taken := c.Take(2)
	if taken.Length() != 2 { t.Fatal("expected 2") }
	limited := c.Limit(2)
	if limited.Length() != 2 { t.Fatal("expected 2") }
	limitAll := c.Limit(-1)
	if limitAll.Length() != 5 { t.Fatal("expected 5") }
	skipped := c.Skip(2)
	if skipped.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_AddNonEmptyStrings(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmptyStrings("a", "", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddNonEmptyStringsSlice(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmptyStringsSlice([]string{"a", "", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyList(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	list := c.NonEmptyList()
	if len(list) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyListPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	list := c.NonEmptyListPtr()
	if list == nil || len(*list) != 2 { t.Fatal("unexpected") }
}

func TestCollection_Items(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	items := c.Items()
	if len(items) != 2 { t.Fatal("expected 2") }
}

func TestCollection_ListPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	lp := c.ListPtr()
	if lp == nil { t.Fatal("expected non-nil") }
}

func TestCollection_ListCopyPtrLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	cp := c.ListCopyPtrLock()
	if cp == nil || len(*cp) != 1 { t.Fatal("unexpected") }
}

func TestCollection_Has(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.Has("b") { t.Fatal("expected true") }
	if c.Has("z") { t.Fatal("expected false") }
}

func TestCollection_HasLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if !c.HasLock("a") { t.Fatal("expected true") }
}

func TestCollection_HasPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.HasPtr("a") { t.Fatal("expected true") }
}

func TestCollection_HasAll(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.HasAll("a", "b") { t.Fatal("expected true") }
	if c.HasAll("a", "z") { t.Fatal("expected false") }
}

func TestCollection_SortedListAsc(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a", "b"})
	sorted := c.SortedListAsc()
	if sorted[0] != "a" { t.Fatal("expected a first") }
}

func TestCollection_SortedAsc(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a", "b"})
	sortedC := c.SortedAsc()
	list := sortedC.List()
	if list[0] != "a" { t.Fatal("expected a first") }
}

func TestCollection_SortedAscLock(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a"})
	sortedC := c.SortedAscLock()
	list := sortedC.List()
	if list[0] != "a" { t.Fatal("expected a first") }
}

func TestCollection_SortedListDsc(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "c", "b"})
	sorted := c.SortedListDsc()
	if sorted[0] != "c" { t.Fatal("expected c first") }
}

func TestCollection_FilterAndFilterPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"apple", "banana", "cherry"})
	filtered := c.Filter(func(s string) bool { return len(s) > 5 })
	if len(filtered) < 2 { t.Fatal("expected >= 2") }
}

func TestCollection_FilterLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb", "ccc"})
	filtered := c.FilterLock(func(s string) bool { return len(s) > 1 })
	if len(filtered) != 2 { t.Fatal("expected 2") }
}

func TestCollection_FilteredCollection(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fc := c.FilteredCollection(func(s string) bool { return len(s) > 1 })
	if fc.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_FilteredCollectionLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fc := c.FilteredCollectionLock(func(s string) bool { return len(s) > 1 })
	if fc.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_FilterPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fp := c.FilterPtr(func(s string) bool { return len(s) > 1 })
	if fp == nil || len(*fp) != 1 { t.Fatal("unexpected") }
}

func TestCollection_FilterPtrLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "bb"})
	fp := c.FilterPtrLock(func(s string) bool { return len(s) > 1 })
	if fp == nil || len(*fp) != 1 { t.Fatal("unexpected") }
}

func TestCollection_HashsetAsIs(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "a"})
	hs := c.HashsetAsIs()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestCollection_HashsetWithDoubleLength(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	hs := c.HashsetWithDoubleLength()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestCollection_HashsetLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	hs := c.HashsetLock()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestCollection_NonEmptyItems(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	items := c.NonEmptyItems()
	if len(items) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyItemsPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	ip := c.NonEmptyItemsPtr()
	if ip == nil || len(*ip) != 2 { t.Fatal("unexpected") }
}

func TestCollection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	c := New.Collection.Strings([]string{"a", " ", "b"})
	items := c.NonEmptyItemsOrNonWhitespace()
	if len(items) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyItemsOrNonWhitespacePtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a", " ", "b"})
	ip := c.NonEmptyItemsOrNonWhitespacePtr()
	if ip == nil || len(*ip) != 2 { t.Fatal("unexpected") }
}

func TestCollection_IsContainsPtr(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.IsContainsPtr("a") { t.Fatal("expected true") }
}

func TestCollection_IsContainsAll(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.IsContainsAll("a", "b") { t.Fatal("expected true") }
}

func TestCollection_IsContainsAllLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.IsContainsAllLock("a", "c") { t.Fatal("expected true") }
}

func TestCollection_IsContainsAllSlice(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.IsContainsAllSlice([]string{"a", "b"}) { t.Fatal("expected true") }
}

func TestCollection_GetHashsetPlusHasAll(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	hs, hasAll := c.GetHashsetPlusHasAll("a", "b")
	if !hasAll || hs == nil { t.Fatal("unexpected") }
}

func TestCollection_HasUsingSensitivity(t *testing.T) {
	c := New.Collection.Strings([]string{"Hello", "World"})
	if !c.HasUsingSensitivity(false, "hello") { t.Fatal("expected true case-insensitive") }
	if !c.HasUsingSensitivity(true, "Hello") { t.Fatal("expected true case-sensitive") }
}

func TestCollection_New(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	newC := c.New("c", "d")
	if newC.Length() != 4 { t.Fatal("expected 4") }
}

func TestCollection_ExpandSlicePlusAdd(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.ExpandSlicePlusAdd([]string{"b", "c"})
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_MergeSlicesOfSlice(t *testing.T) {
	c := New.Collection.Empty()
	c.MergeSlicesOfSlice([][]string{{"a", "b"}, {"c"}})
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_GetAllExceptCollection(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	except := New.Collection.Strings([]string{"b"})
	result := c.GetAllExceptCollection(except)
	if result.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_GetAllExcept(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	result := c.GetAllExcept("b")
	if result.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_Joins(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	s := c.Joins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_NonEmptyJoins(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	s := c.NonEmptyJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_NonWhitespaceJoins(t *testing.T) {
	c := New.Collection.Strings([]string{"a", " ", "b"})
	s := c.NonWhitespaceJoins(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_String(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.String() == "" { t.Fatal("expected non-empty") }
}

func TestCollection_StringLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestCollection_SummaryString(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.SummaryString() == "" { t.Fatal("expected non-empty") }
}

func TestCollection_SummaryStringWithHeader(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := c.SummaryStringWithHeader("header")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_CsvLines(t *testing.T) {
	c := New.Collection.Strings([]string{"a,b", "c,d"})
	_ = c.CsvLines()
}

func TestCollection_Csv(t *testing.T) {
	c := New.Collection.Strings([]string{"a,b", "c,d"})
	_ = c.Csv()
}

func TestCollection_AddCapacity(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.AddCapacity(10)
}

func TestCollection_Resize(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	c.Resize(5)
}

func TestCollection_CharCollectionMap(t *testing.T) {
	c := New.Collection.Strings([]string{"apple", "banana"})
	ccm := c.CharCollectionMap()
	if ccm == nil { t.Fatal("expected non-nil") }
}

func TestCollection_Join(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	s := c.Join(",")
	if s != "a,b" { t.Fatal("unexpected") }
}

func TestCollection_JoinLine(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	s := c.JoinLine()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_JsonAndInterfaces(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonMarshaller()
	_ = c.AsJsonContractsBinder()
	_ = c.Serialize()
	_ = c.Deserialize()
}

func TestCollection_ClearDispose(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	c.Clear()
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Add("x")
	c.Dispose()
}

func TestCollection_AddFuncResult(t *testing.T) {
	c := New.Collection.Empty()
	c.AddFuncResult(func() string { return "hello" })
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddStringsByFuncChecking(t *testing.T) {
	c := New.Collection.Empty()
	c.AddStringsByFuncChecking(func(s string) bool { return len(s) > 1 }, "a", "bb", "c", "dd")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_ParseInjectUsingJson(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	j := c.Serialize()
	c2 := New.Collection.Empty()
	err := c2.ParseInjectUsingJson(j.SafeBytes())
	if err != nil { t.Fatal(err) }
}

func TestCollection_JsonParseSelfInject(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	j := c.Serialize()
	c2 := New.Collection.Empty()
	err := c2.JsonParseSelfInject(j.SafeBytes())
	if err != nil { t.Fatal(err) }
}

// ── Hashmap extended coverage ──

func TestHashmap_HasItems(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasItems() { t.Fatal("expected true") }
}

func TestHashmap_Collection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := hm.Collection()
	if c == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_IsEmptyLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	if !hm.IsEmptyLock() { t.Fatal("expected true") }
}

func TestHashmap_AddOrUpdateLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateLock("a", "1")
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ContainsAndHas(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.Contains("a") { t.Fatal("expected true") }
	if !hm.ContainsLock("a") { t.Fatal("expected true") }
	if hm.IsKeyMissing("a") { t.Fatal("expected false") }
	if hm.IsKeyMissingLock("a") { t.Fatal("expected false") }
	if !hm.HasLock("a") { t.Fatal("expected true") }
}

func TestHashmap_HasAllStrings(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	if !hm.HasAllStrings("a", "b") { t.Fatal("expected true") }
	if hm.HasAllStrings("a", "c") { t.Fatal("expected false") }
}

func TestHashmap_HasAll(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasAll("a") { t.Fatal("expected true") }
}

func TestHashmap_HasAny(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasAny("a", "b") { t.Fatal("expected true") }
}

func TestHashmap_HasWithLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if !hm.HasWithLock("a") { t.Fatal("expected true") }
}

func TestHashmap_Keys(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys := hm.Keys()
	if len(keys) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysCollection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kc := hm.KeysCollection()
	if kc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_AllKeys(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys := hm.AllKeys()
	if len(keys) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	keys := hm.KeysLock()
	if len(keys) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ValuesList(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vals := hm.ValuesList()
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ValuesCollection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vc := hm.ValuesCollection()
	if vc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesHashset(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vh := hm.ValuesHashset()
	if vh == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesCollectionLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vc := hm.ValuesCollectionLock()
	if vc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesHashsetLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vh := hm.ValuesHashsetLock()
	if vh == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_KeysValuesCollection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvc := hm.KeysValuesCollection()
	if kvc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_KeysValuesList(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvl := hm.KeysValuesList()
	if len(kvl) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysValuePairs(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvp := hm.KeysValuePairs()
	if len(kvp) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysValuePairsCollection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvpc := hm.KeysValuePairsCollection()
	if kvpc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_KeysValuesListLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	kvl := hm.KeysValuesListLock()
	if len(kvl) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_LengthLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if hm.LengthLock() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_Remove(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.Remove("a")
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_RemoveWithLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.RemoveWithLock("a")
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_String(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if hm.String() == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_StringLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	if hm.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_ItemsCopyLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	cp := hm.ItemsCopyLock()
	if len(cp) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_SafeItems(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	si := hm.SafeItems()
	if si == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ValuesListCopyLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	vl := hm.ValuesListCopyLock()
	if len(vl) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ValuesToLower(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "HELLO")
	hm.ValuesToLower()
}

func TestHashmap_KeysToLower(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("ABC", "1")
	hm.KeysToLower()
}

func TestHashmap_IsEqual(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	if !hm1.IsEqual(hm2) { t.Fatal("expected equal") }
}

func TestHashmap_IsEqualPtrLock(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	if !hm1.IsEqualPtrLock(hm2) { t.Fatal("expected equal") }
}

func TestHashmap_Clone(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	cloned := hm.Clone()
	if cloned.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetValue(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	v := hm.GetValue("a")
	if v != "1" { t.Fatal("expected 1") }
}

func TestHashmap_Join(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	s := hm.Join(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_JoinKeys(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	s := hm.JoinKeys(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_Dispose(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.Dispose()
}

func TestHashmap_ToError(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("err", "something")
	e := hm.ToError(",")
	if e == nil { t.Fatal("expected error") }
}

func TestHashmap_ToDefaultError(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("err", "something")
	e := hm.ToDefaultError()
	if e == nil { t.Fatal("expected error") }
}

func TestHashmap_KeyValStringLines(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	lines := hm.KeyValStringLines()
	if len(lines) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_JsonInterfaces(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	_ = hm.JsonModelAny()
	_ = hm.Json()
	_ = hm.JsonPtr()
	_ = hm.AsJsoner()
	_ = hm.AsJsonContractsBinder()
	_ = hm.AsJsonParseSelfInjector()
	_ = hm.AsJsonMarshaller()
	_ = hm.Serialize()
	_ = hm.Deserialize()
}

func TestHashmap_ParseInjectUsingJson(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	j := hm.Serialize()
	hm2 := New.Hashmap.Empty()
	err := hm2.ParseInjectUsingJson(j.SafeBytes())
	if err != nil { t.Fatal(err) }
}

func TestHashmap_JsonParseSelfInject(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	j := hm.Serialize()
	hm2 := New.Hashmap.Empty()
	err := hm2.JsonParseSelfInject(j.SafeBytes())
	if err != nil { t.Fatal(err) }
}

func TestHashmap_AddOrUpdateHashmap(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdateHashmap(hm1)
	if hm2.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddOrUpdateMap(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateMap(map[string]string{"a": "1", "b": "2"})
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddsOrUpdates(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdates("a", "1", "b", "2")
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddOrUpdateCollection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	c := New.Collection.Strings([]string{"a"})
	hm.AddOrUpdateCollection("k", c)
}

func TestHashmap_Set(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.Set("a", "1")
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_SetTrim(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.SetTrim(" a ", " 1 ")
	if hm.GetValue("a") != "1" { t.Fatal("expected trimmed") }
}

func TestHashmap_SetBySplitter(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.SetBySplitter("key=value", "=")
	if hm.GetValue("key") != "value" { t.Fatal("expected value") }
}

func TestHashmap_DiffRaw(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm1.AddOrUpdate("b", "2")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	diff := hm1.DiffRaw(hm2)
	_ = diff
}

func TestHashmap_Diff(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("a", "1")
	diff := hm1.Diff(hm2)
	_ = diff
}

func TestHashmap_HasAllCollectionItems(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	c := New.Collection.Strings([]string{"a"})
	if !hm.HasAllCollectionItems(c) { t.Fatal("expected true") }
}

func TestHashmap_GetKeysFilteredItems(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	filtered := hm.GetKeysFilteredItems(func(k, v string) bool { return k == "a" })
	if len(filtered) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetKeysFilteredCollection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	fc := hm.GetKeysFilteredCollection(func(k, v string) bool { return k == "a" })
	if fc == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	hs := New.Hashset.Strings([]string{"a"})
	vals := hm.GetValuesExceptKeysInHashset(hs)
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetValuesKeysExcept(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	vals := hm.GetValuesKeysExcept("a")
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetAllExceptCollection(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	hm.AddOrUpdate("b", "2")
	result := hm.GetAllExceptCollection("a")
	if result == nil { t.Fatal("expected non-nil") }
}

func TestHashmap_ToStringsUsingCompiler(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	strs := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
	if len(strs) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ConcatNew(t *testing.T) {
	hm1 := New.Hashmap.Cap(5)
	hm1.AddOrUpdate("a", "1")
	hm2 := New.Hashmap.Cap(5)
	hm2.AddOrUpdate("b", "2")
	result := hm1.ConcatNew(hm2)
	if result.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_ConcatNewUsingMaps(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdate("a", "1")
	m := map[string]string{"b": "2"}
	result := hm.ConcatNewUsingMaps(m)
	if result.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesUsingFilter(func(k, v string) bool { return k != "skip" }, "a", "1", "skip", "2", "b", "3")
	if hm.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	items := map[string]any{"a": "1", "b": 2}
	hm.AddsOrUpdatesAnyUsingFilter(items, func(k string, v any) bool { return k == "a" })
}

func TestHashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	items := map[string]any{"a": "1"}
	hm.AddsOrUpdatesAnyUsingFilterLock(items, func(k string, v any) bool { return true })
}

func TestHashmap_AddOrUpdateKeyStrValInt(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValInt("a", 42)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddOrUpdateKeyStrValFloat(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat("a", 3.14)
}

func TestHashmap_AddOrUpdateKeyStrValFloat64(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValFloat64("a", 3.14)
}

func TestHashmap_AddOrUpdateKeyStrValAny(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyStrValAny("a", "hello")
}

func TestHashmap_AddOrUpdateKeyValueAny(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyValueAny("a", 42)
}

func TestHashmap_AddOrUpdateKeyVal(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyVal("a", "1")
}

func TestHashmap_AddOrUpdateWithWgLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateWithWgLock("a", "1")
}

func TestHashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	items := []string{"a", "1", "b", "2"}
	hm.AddOrUpdateStringsPtrWgLock(&items)
}

func TestHashmap_AddOrUpdateKeyAnyValues(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddOrUpdateKeyAnyValues(map[string]any{"a": 1, "b": "hello"})
}

// ── Hashset extended coverage ──

func TestHashset_AddAndHas(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.Has("a") { t.Fatal("expected true") }
	if hs.Has("z") { t.Fatal("expected false") }
	hs.Add("c")
	if hs.Length() != 3 { t.Fatal("expected 3") }
}

func TestHashset_HasAll(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b", "c"})
	if !hs.HasAll("a", "b") { t.Fatal("expected true") }
	if hs.HasAll("a", "z") { t.Fatal("expected false") }
}

func TestHashset_HasAny(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	if !hs.HasAny("a", "z") { t.Fatal("expected true") }
	if hs.HasAny("x", "z") { t.Fatal("expected false") }
}

func TestHashset_Remove(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	hs.Remove("a")
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashset_Adds(t *testing.T) {
	hs := New.Hashset.Empty()
	hs.Adds("a", "b", "c")
	if hs.Length() != 3 { t.Fatal("expected 3") }
}

func TestHashset_List(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	list := hs.List()
	if len(list) != 1 { t.Fatal("expected 1") }
}

func TestHashset_ListSortedAsc(t *testing.T) {
	hs := New.Hashset.Strings([]string{"c", "a", "b"})
	sorted := hs.ListSortedAsc()
	if sorted[0] != "a" { t.Fatal("expected a first") }
}

func TestHashset_Collection(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	c := hs.Collection()
	if c == nil { t.Fatal("expected non-nil") }
}

func TestHashset_Clone(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a", "b"})
	cloned := hs.Clone()
	if cloned.Length() != 2 { t.Fatal("expected 2") }
}

func TestHashset_String(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	if hs.String() == "" { t.Fatal("expected non-empty") }
}

func TestHashset_Join(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	s := hs.Join(",")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_JoinLine(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	s := hs.JoinLine()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestHashset_Dispose(t *testing.T) {
	hs := New.Hashset.Strings([]string{"a"})
	hs.Dispose()
}

func TestHashset_JsonInterfaces(t *testing.T) {
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

func TestNewCreator_CollectionEmpty(t *testing.T) {
	c := New.Collection.Empty()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestNewCreator_CollectionCloneStrings(t *testing.T) {
	c := New.Collection.CloneStrings([]string{"a", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_CollectionCreate(t *testing.T) {
	c := New.Collection.Create(5)
	if c == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_CollectionStringsPlusCap(t *testing.T) {
	c := New.Collection.StringsPlusCap(10, []string{"a"})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCreator_CollectionCapStrings(t *testing.T) {
	c := New.Collection.CapStrings(10, "a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_CollectionLenCap(t *testing.T) {
	c := New.Collection.LenCap(0, 10)
	if c == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_CollectionLineDefault(t *testing.T) {
	c := New.Collection.LineDefault("a\nb\nc")
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestNewCreator_CollectionLineUsingSep(t *testing.T) {
	c := New.Collection.LineUsingSep("a,b,c", ",")
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestNewCreator_CollectionStringsOptions(t *testing.T) {
	c := New.Collection.StringsOptions(true, []string{"a", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_HashmapEmpty(t *testing.T) {
	hm := New.Hashmap.Empty()
	if hm == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_HashmapMapWithCap(t *testing.T) {
	hm := New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCreator_HashsetEmpty(t *testing.T) {
	hs := New.Hashset.Empty()
	if hs == nil { t.Fatal("expected non-nil") }
}

func TestNewCreator_HashsetStringsOption(t *testing.T) {
	hs := New.Hashset.StringsOption(true, []string{"a"})
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCreator_HashsetStringSpreadItems(t *testing.T) {
	hs := New.Hashset.StringsSpreadItems("a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCreator_HashsetUsingCollection(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	hs := New.Hashset.UsingCollection(c)
	if hs.Length() != 1 { t.Fatal("expected 1") }
}

// ── emptyCreator paths ──

func TestEmptyCreator_LinkedList(t *testing.T) {
	ll := Empty.LinkedList()
	if ll == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyValuePair(t *testing.T) {
	kvp := Empty.KeyValuePair()
	if kvp == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyAnyValuePair(t *testing.T) {
	kavp := Empty.KeyAnyValuePair()
	if kavp == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyValueCollection(t *testing.T) {
	kvc := Empty.KeyValueCollection()
	if kvc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_LinkedCollections(t *testing.T) {
	lc := Empty.LinkedCollections()
	if lc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_LeftRight(t *testing.T) {
	lr := Empty.LeftRight()
	_ = lr
}

func TestEmptyCreator_SimpleStringOnce(t *testing.T) {
	s := Empty.SimpleStringOnce()
	_ = s
}

func TestEmptyCreator_SimpleStringOncePtr(t *testing.T) {
	sp := Empty.SimpleStringOncePtr()
	if sp == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_HashsetsCollection(t *testing.T) {
	hsc := Empty.HashsetsCollection()
	if hsc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_CharCollectionMap(t *testing.T) {
	ccm := Empty.CharCollectionMap()
	if ccm == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_KeyValuesCollection(t *testing.T) {
	kvc := Empty.KeyValuesCollection()
	if kvc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_CollectionsOfCollection(t *testing.T) {
	coc := Empty.CollectionsOfCollection()
	if coc == nil { t.Fatal("expected non-nil") }
}

func TestEmptyCreator_CharHashsetMap(t *testing.T) {
	chm := Empty.CharHashsetMap()
	if chm == nil { t.Fatal("expected non-nil") }
}

// ── AnyToString ──

func TestAnyToString(t *testing.T) {
	s := AnyToString(42)
	if s == "" { t.Fatal("expected non-empty") }
	s2 := AnyToString("hello")
	if s2 != "hello" { t.Fatal("unexpected") }
	s3 := AnyToString(nil)
	_ = s3
}

// ── AllIndividualStringsOfStringsLength ──

func TestAllIndividualStringsOfStringsLength(t *testing.T) {
	length := AllIndividualStringsOfStringsLength([][]string{{"a", "b"}, {"c"}})
	if length != 3 { t.Fatal("expected 3") }
}

// ── AllIndividualsLengthOfSimpleSlices ──

func TestAllIndividualsLengthOfSimpleSlices(t *testing.T) {
	ss1 := New.SimpleSlice.Lines("a", "b")
	ss2 := New.SimpleSlice.Lines("c")
	length := AllIndividualsLengthOfSimpleSlices([]*SimpleSlice{ss1, ss2})
	if length != 3 { t.Fatal("expected 3") }
}
