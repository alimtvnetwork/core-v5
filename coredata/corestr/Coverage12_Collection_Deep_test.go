package corestr

import (
	"errors"
	"sync"
	"testing"
)

// ── AddIf ──

func TestCollection_AddIf_True_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddIf(true, "a")
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddIf_False_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddIf(false, "a")
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddFuncErr ──

func TestCollection_AddFuncErr_Success_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddFuncErr(
		func() (string, error) { return "ok", nil },
		func(err error) { t.Fatal("unexpected error") },
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddFuncErr_Error_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	called := false
	c.AddFuncErr(
		func() (string, error) { return "", errors.New("fail") },
		func(err error) { called = true },
	)
	if c.Length() != 0 { t.Fatal("expected 0") }
	if !called { t.Fatal("expected error handler called") }
}

// ── AddError with value ──

func TestCollection_AddError_WithValue_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddError(errors.New("test-err"))
	if c.Length() != 1 { t.Fatal("expected 1") }
	if c.First() != "test-err" { t.Fatal("unexpected value") }
}

// ── AddHashmapsValues ──

func TestCollection_AddHashmapsValues_C12(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := New.Collection.Cap(5)
	c.AddHashmapsValues(hm)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddHashmapsValues_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddHashmapsValues(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AddHashmapsValues_NilHashmap_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddHashmapsValues(nil, nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddHashmapsKeys ──

func TestCollection_AddHashmapsKeys_C12(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := New.Collection.Cap(5)
	c.AddHashmapsKeys(hm)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddHashmapsKeys_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddHashmapsKeys(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddHashmapsKeysValues ──

func TestCollection_AddHashmapsKeysValues_C12(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := New.Collection.Cap(5)
	c.AddHashmapsKeysValues(hm)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddHashmapsKeysValues_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddHashmapsKeysValues(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddHashmapsKeysValuesUsingFilter ──

func TestCollection_AddHashmapsKeysValuesUsingFilter_C12(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := New.Collection.Cap(5)
	c.AddHashmapsKeysValuesUsingFilter(
		func(pair KeyValuePair) (string, bool, bool) {
			return pair.Key + "=" + pair.Value, true, false
		},
		hm,
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddHashmapsKeysValuesUsingFilter_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddHashmapsKeysValuesUsingFilter(nil, nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AddHashmapsKeysValuesUsingFilter_Break_C12(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	c := New.Collection.Cap(5)
	c.AddHashmapsKeysValuesUsingFilter(
		func(pair KeyValuePair) (string, bool, bool) {
			return pair.Key, true, true // break after first
		},
		hm,
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── AddWithWgLock ──

func TestCollection_AddWithWgLock_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(&wg, "x")
	wg.Wait()
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── IndexAt, SafeIndexAtUsingLength ──

func TestCollection_IndexAt_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if c.IndexAt(1) != "b" { t.Fatal("expected b") }
}

func TestCollection_SafeIndexAtUsingLength_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.SafeIndexAtUsingLength("def", 1, 0) != "a" { t.Fatal("expected a") }
	if c.SafeIndexAtUsingLength("def", 1, 5) != "def" { t.Fatal("expected def") }
}

// ── First, Last, FirstOrDefault, LastOrDefault, Single ──

func TestCollection_First_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"x", "y"})
	if c.First() != "x" { t.Fatal("expected x") }
}

func TestCollection_Last_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"x", "y"})
	if c.Last() != "y" { t.Fatal("expected y") }
}

func TestCollection_LastOrDefault_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	if c.LastOrDefault() != "" { t.Fatal("expected empty") }
}

func TestCollection_LastOrDefault_HasItems_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.LastOrDefault() != "a" { t.Fatal("expected a") }
}

func TestCollection_FirstOrDefault_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	if c.FirstOrDefault() != "" { t.Fatal("expected empty") }
}

func TestCollection_FirstOrDefault_HasItems_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"z"})
	if c.FirstOrDefault() != "z" { t.Fatal("expected z") }
}

func TestCollection_Single_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"only"})
	if c.Single() != "only" { t.Fatal("expected only") }
}

func TestCollection_Single_Panic_C12(t *testing.T) {
	defer func() {
		if r := recover(); r == nil { t.Fatal("expected panic") }
	}()
	c := New.Collection.Strings([]string{"a", "b"})
	c.Single()
}

// ── Take, Skip ──

func TestCollection_Take_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	r := c.Take(2)
	if r.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_Take_MoreThanLength_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.Take(5)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_Take_Zero_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.Take(0)
	if r.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_Skip_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	r := c.Skip(1)
	if r.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_Skip_Zero_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.Skip(0)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_Skip_Panic_C12(t *testing.T) {
	defer func() {
		if r := recover(); r == nil { t.Fatal("expected panic") }
	}()
	c := New.Collection.Strings([]string{"a"})
	c.Skip(5)
}

// ── Reverse ──

func TestCollection_Reverse_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	c.Reverse()
	if c.First() != "c" || c.Last() != "a" { t.Fatal("unexpected") }
}

func TestCollection_Reverse_Two_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	c.Reverse()
	if c.First() != "b" { t.Fatal("expected b") }
}

func TestCollection_Reverse_One_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.Reverse()
	if c.First() != "a" { t.Fatal("expected a") }
}

// ── GetPagesSize ──

func TestCollection_GetPagesSize_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
	if c.GetPagesSize(2) != 3 { t.Fatal("expected 3") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	if c.GetPagesSize(-1) != 0 { t.Fatal("expected 0") }
}

// ── GetPagedCollection ──

func TestCollection_GetPagedCollection_C12(t *testing.T) {
	items := make([]string, 10)
	for i := range items { items[i] = "x" }
	c := New.Collection.Strings(items)
	paged := c.GetPagedCollection(3)
	if paged.Length() != 4 { t.Fatal("expected 4 pages") }
}

func TestCollection_GetPagedCollection_Small_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	paged := c.GetPagedCollection(5)
	if paged.Length() != 1 { t.Fatal("expected 1") }
}

// ── GetSinglePageCollection ──

func TestCollection_GetSinglePageCollection_C12(t *testing.T) {
	items := make([]string, 10)
	for i := range items { items[i] = "x" }
	c := New.Collection.Strings(items)
	page := c.GetSinglePageCollection(3, 2)
	if page.Length() != 3 { t.Fatal("expected 3") }
}

func TestCollection_GetSinglePageCollection_LastPage_C12(t *testing.T) {
	items := make([]string, 10)
	for i := range items { items[i] = "x" }
	c := New.Collection.Strings(items)
	page := c.GetSinglePageCollection(3, 4)
	if page.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_GetSinglePageCollection_Small_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	page := c.GetSinglePageCollection(5, 1)
	if page.Length() != 1 { t.Fatal("expected 1") }
}

// ── InsertAt ──

func TestCollection_InsertAt_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.InsertAt(0, "a")
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── ChainRemoveAt ──

func TestCollection_ChainRemoveAt_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	c.ChainRemoveAt(1)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

// ── RemoveItemsIndexes ──

func TestCollection_RemoveItemsIndexes_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	c.RemoveItemsIndexes(true, 1)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_RemoveItemsIndexes_NilIgnored_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.RemoveItemsIndexes(true)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── AppendCollectionPtr ──

func TestCollection_AppendCollectionPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	c.AppendCollectionPtr(c2)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

// ── AppendCollections ──

func TestCollection_AppendCollections_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	c.AppendCollections(c1, c2)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AppendCollections_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendCollections()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AppendAnys ──

func TestCollection_AppendAnys_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnys(42, "hello", nil)
	if c.Length() != 2 { t.Fatal("expected 2, nil skipped") }
}

func TestCollection_AppendAnys_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnys()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AppendAnysLock ──

func TestCollection_AppendAnysLock_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysLock(42)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AppendAnysLock_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysLock()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AppendAnysUsingFilter ──

func TestCollection_AppendAnysUsingFilter_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilter(
		func(str string, i int) (string, bool, bool) {
			return str, true, false
		},
		"a", "b",
	)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AppendAnysUsingFilter_Skip_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilter(
		func(str string, i int) (string, bool, bool) {
			return str, false, false
		},
		"a",
	)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AppendAnysUsingFilter_Break_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilter(
		func(str string, i int) (string, bool, bool) {
			return str, true, true
		},
		"a", "b",
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AppendAnysUsingFilter_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilter(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AppendAnysUsingFilter_NilItem_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilter(
		func(str string, i int) (string, bool, bool) { return str, true, false },
		nil,
	)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AppendAnysUsingFilterLock ──

func TestCollection_AppendAnysUsingFilterLock_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilterLock(
		func(str string, i int) (string, bool, bool) { return str, true, false },
		"a",
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AppendAnysUsingFilterLock_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilterLock(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AppendAnysUsingFilterLock_Break_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilterLock(
		func(str string, i int) (string, bool, bool) { return str, true, true },
		"a", "b",
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AppendAnysUsingFilterLock_Skip_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilterLock(
		func(str string, i int) (string, bool, bool) { return str, false, false },
		"a",
	)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_AppendAnysUsingFilterLock_NilItem_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendAnysUsingFilterLock(
		func(str string, i int) (string, bool, bool) { return str, true, false },
		nil,
	)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AppendNonEmptyAnys ──

func TestCollection_AppendNonEmptyAnys_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendNonEmptyAnys(42, nil)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AppendNonEmptyAnys_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AppendNonEmptyAnys(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddsNonEmpty ──

func TestCollection_AddsNonEmpty_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddsNonEmpty("a", "", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddsNonEmpty_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddsNonEmpty()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddsNonEmptyPtrLock ──

func TestCollection_AddsNonEmptyPtrLock_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	s := "hello"
	empty := ""
	c.AddsNonEmptyPtrLock(&s, nil, &empty)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestCollection_AddsNonEmptyPtrLock_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddsNonEmptyPtrLock()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── UniqueBoolMap / UniqueList ──

func TestCollection_UniqueBoolMap_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "a"})
	m := c.UniqueBoolMap()
	if len(m) != 2 { t.Fatal("expected 2") }
}

func TestCollection_UniqueBoolMapLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "a"})
	m := c.UniqueBoolMapLock()
	if len(m) != 1 { t.Fatal("expected 1") }
}

func TestCollection_UniqueList_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "a"})
	ul := c.UniqueList()
	if len(ul) != 2 { t.Fatal("expected 2") }
}

func TestCollection_UniqueListLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "a"})
	ul := c.UniqueListLock()
	if len(ul) != 2 { t.Fatal("expected 2") }
}

// ── Filter ──

func TestCollection_Filter_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	result := c.Filter(func(str string, i int) (string, bool, bool) {
		return str, str != "b", false
	})
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestCollection_Filter_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	result := c.Filter(func(str string, i int) (string, bool, bool) { return str, true, false })
	if len(result) != 0 { t.Fatal("expected 0") }
}

func TestCollection_Filter_Break_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	result := c.Filter(func(str string, i int) (string, bool, bool) {
		return str, true, i == 0
	})
	if len(result) != 1 { t.Fatal("expected 1") }
}

// ── FilterLock ──

func TestCollection_FilterLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	result := c.FilterLock(func(str string, i int) (string, bool, bool) {
		return str, true, false
	})
	if len(result) != 2 { t.Fatal("expected 2") }
}

// ── FilteredCollection ──

func TestCollection_FilteredCollection_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	fc := c.FilteredCollection(func(str string, i int) (string, bool, bool) {
		return str, true, false
	})
	if fc.Length() != 2 { t.Fatal("expected 2") }
}

// ── FilteredCollectionLock ──

func TestCollection_FilteredCollectionLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	fc := c.FilteredCollectionLock(func(str string, i int) (string, bool, bool) {
		return str, true, false
	})
	if fc.Length() != 2 { t.Fatal("expected 2") }
}

// ── FilterPtr ──

func TestCollection_FilterPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
		return sp, true, false
	})
	if len(*result) != 2 { t.Fatal("expected 2") }
}

func TestCollection_FilterPtr_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
		return sp, true, false
	})
	if len(*result) != 0 { t.Fatal("expected 0") }
}

func TestCollection_FilterPtr_Break_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
		return sp, true, i == 0
	})
	if len(*result) != 1 { t.Fatal("expected 1") }
}

// ── FilterPtrLock ──

func TestCollection_FilterPtrLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
		return sp, true, false
	})
	if len(*result) != 2 { t.Fatal("expected 2") }
}

func TestCollection_FilterPtrLock_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
		return sp, true, false
	})
	if len(*result) != 0 { t.Fatal("expected 0") }
}

func TestCollection_FilterPtrLock_Break_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
		return sp, true, i == 0
	})
	if len(*result) != 1 { t.Fatal("expected 1") }
}

// ── NonEmptyList / NonEmptyListPtr ──

func TestCollection_NonEmptyList_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	r := c.NonEmptyList()
	if len(r) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyList_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	r := c.NonEmptyList()
	if len(r) != 0 { t.Fatal("expected 0") }
}

func TestCollection_NonEmptyListPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", ""})
	r := c.NonEmptyListPtr()
	if len(*r) != 1 { t.Fatal("expected 1") }
}

// ── Hashset helpers ──

func TestCollection_HashsetAsIs_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	hs := c.HashsetAsIs()
	if hs.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestCollection_HashsetWithDoubleLength_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	hs := c.HashsetWithDoubleLength()
	if hs.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestCollection_HashsetLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	hs := c.HashsetLock()
	if hs.IsEmpty() { t.Fatal("expected non-empty") }
}

// ── NonEmptyItems ──

func TestCollection_NonEmptyItems_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	r := c.NonEmptyItems()
	if len(r) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyItemsPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", ""})
	r := c.NonEmptyItemsPtr()
	if len(r) != 1 { t.Fatal("expected 1") }
}

func TestCollection_NonEmptyItemsOrNonWhitespace_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "  ", "b"})
	r := c.NonEmptyItemsOrNonWhitespace()
	if len(r) != 2 { t.Fatal("expected 2") }
}

func TestCollection_NonEmptyItemsOrNonWhitespacePtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "  "})
	r := c.NonEmptyItemsOrNonWhitespacePtr()
	if len(r) != 1 { t.Fatal("expected 1") }
}

// ── Has, HasPtr, HasAll, HasLock ──

func TestCollection_Has_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.Has("a") { t.Fatal("expected true") }
	if c.Has("z") { t.Fatal("expected false") }
}

func TestCollection_Has_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	if c.Has("a") { t.Fatal("expected false") }
}

func TestCollection_HasPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := "a"
	if !c.HasPtr(&s) { t.Fatal("expected true") }
	if c.HasPtr(nil) { t.Fatal("expected false") }
}

func TestCollection_HasPtr_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	s := "a"
	if c.HasPtr(&s) { t.Fatal("expected false") }
}

func TestCollection_HasAll_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.HasAll("a", "b") { t.Fatal("expected true") }
	if c.HasAll("a", "z") { t.Fatal("expected false") }
}

func TestCollection_HasAll_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	if c.HasAll("a") { t.Fatal("expected false") }
}

func TestCollection_HasLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if !c.HasLock("a") { t.Fatal("expected true") }
}

// ── HasUsingSensitivity ──

func TestCollection_HasUsingSensitivity_CaseSensitive_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"Hello"})
	if !c.HasUsingSensitivity("Hello", true) { t.Fatal("expected true") }
	if c.HasUsingSensitivity("hello", true) { t.Fatal("expected false") }
}

func TestCollection_HasUsingSensitivity_CaseInsensitive_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"Hello"})
	if !c.HasUsingSensitivity("hello", false) { t.Fatal("expected true") }
}

// ── SortedListAsc, SortedAsc, SortedListDsc ──

func TestCollection_SortedListAsc_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a", "b"})
	r := c.SortedListAsc()
	if r[0] != "a" { t.Fatal("expected a") }
}

func TestCollection_SortedListAsc_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	r := c.SortedListAsc()
	if len(r) != 0 { t.Fatal("expected 0") }
}

func TestCollection_SortedAsc_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a"})
	c.SortedAsc()
	if c.First() != "a" { t.Fatal("expected a") }
}

func TestCollection_SortedAsc_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	c.SortedAsc()
}

func TestCollection_SortedAscLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"c", "a"})
	c.SortedAscLock()
	if c.First() != "a" { t.Fatal("expected a") }
}

func TestCollection_SortedListDsc_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "c", "b"})
	r := c.SortedListDsc()
	if r[0] != "c" { t.Fatal("expected c") }
}

// ── IsContainsPtr, IsContainsAll, IsContainsAllSlice, IsContainsAllLock ──

func TestCollection_IsContainsPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := "a"
	if !c.IsContainsPtr(&s) { t.Fatal("expected true") }
	if c.IsContainsPtr(nil) { t.Fatal("expected false") }
}

func TestCollection_IsContainsAllSlice_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.IsContainsAllSlice([]string{"a", "b"}) { t.Fatal("expected true") }
	if c.IsContainsAllSlice([]string{"z"}) { t.Fatal("expected false") }
	if c.IsContainsAllSlice([]string{}) { t.Fatal("expected false") }
}

func TestCollection_IsContainsAll_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.IsContainsAll("a", "b") { t.Fatal("expected true") }
}

func TestCollection_IsContainsAllLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if !c.IsContainsAllLock("a") { t.Fatal("expected true") }
}

// ── GetHashsetPlusHasAll ──

func TestCollection_GetHashsetPlusHasAll_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	hs, ok := c.GetHashsetPlusHasAll([]string{"a", "b"})
	if !ok || hs.IsEmpty() { t.Fatal("expected true") }
}

func TestCollection_GetHashsetPlusHasAll_Nil_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	_, ok := c.GetHashsetPlusHasAll(nil)
	if ok { t.Fatal("expected false") }
}

// ── GetAllExceptCollection, GetAllExcept ──

func TestCollection_GetAllExceptCollection_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	except := New.Collection.Strings([]string{"b"})
	r := c.GetAllExceptCollection(except)
	if len(r) != 2 { t.Fatal("expected 2") }
}

func TestCollection_GetAllExceptCollection_Nil_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.GetAllExceptCollection(nil)
	if len(r) != 1 { t.Fatal("expected 1") }
}

func TestCollection_GetAllExcept_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	r := c.GetAllExcept([]string{"a"})
	if len(r) != 2 { t.Fatal("expected 2") }
}

func TestCollection_GetAllExcept_Nil_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.GetAllExcept(nil)
	if len(r) != 1 { t.Fatal("expected 1") }
}

// ── New ──

func TestCollection_New_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	nc := c.New("a", "b")
	if nc.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_New_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	nc := c.New()
	if nc.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddNonEmptyStrings, AddNonEmptyStringsSlice ──

func TestCollection_AddNonEmptyStrings_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddNonEmptyStrings("a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddNonEmptyStrings_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddNonEmptyStrings()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddFuncResult ──

func TestCollection_AddFuncResult_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddFuncResult(func() string { return "a" }, func() string { return "b" })
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestCollection_AddFuncResult_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddFuncResult()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddStringsByFuncChecking ──

func TestCollection_AddStringsByFuncChecking_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddStringsByFuncChecking(
		[]string{"ok", "bad", "ok2"},
		func(line string) bool { return line != "bad" },
	)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

// ── ExpandSlicePlusAdd ──

func TestCollection_ExpandSlicePlusAdd_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.ExpandSlicePlusAdd(
		[]string{"a,b", "c,d"},
		func(line string) []string {
			return []string{line + "_expanded"}
		},
	)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

// ── MergeSlicesOfSlice ──

func TestCollection_MergeSlicesOfSlice_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

// ── CharCollectionMap ──

func TestCollection_CharCollectionMap_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"abc", "def"})
	m := c.CharCollectionMap()
	if m == nil { t.Fatal("expected non-nil") }
}

// ── String, StringLock ──

func TestCollection_String_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.String() == "" { t.Fatal("expected non-empty") }
}

func TestCollection_String_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	s := c.String()
	if s == "" { t.Fatal("expected non-empty (NoElements)") }
}

func TestCollection_StringLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if c.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestCollection_StringLock_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	s := c.StringLock()
	if s == "" { t.Fatal("expected non-empty") }
}

// ── SummaryString ──

func TestCollection_SummaryString_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := c.SummaryString(1)
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_SummaryStringWithHeader_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := c.SummaryStringWithHeader("header:")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestCollection_SummaryStringWithHeader_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	s := c.SummaryStringWithHeader("header:")
	if s == "" { t.Fatal("expected non-empty") }
}

// ── Csv / CsvLines ──

func TestCollection_Csv_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	csv := c.Csv()
	if csv == "" { t.Fatal("expected non-empty") }
}

func TestCollection_Csv_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	csv := c.Csv()
	if csv != "" { t.Fatal("expected empty") }
}

func TestCollection_CsvOptions_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	csv := c.CsvOptions(true)
	if csv == "" { t.Fatal("expected non-empty") }
}

func TestCollection_CsvOptions_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	csv := c.CsvOptions(false)
	if csv != "" { t.Fatal("expected empty") }
}

func TestCollection_CsvLines_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	lines := c.CsvLines()
	if len(lines) != 1 { t.Fatal("expected 1") }
}

func TestCollection_CsvLinesOptions_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	lines := c.CsvLinesOptions(true)
	if len(lines) != 1 { t.Fatal("expected 1") }
}

// ── AddCapacity, Resize ──

func TestCollection_AddCapacity_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddCapacity(10)
	if c.Capacity() < 10 { t.Fatal("expected >= 10") }
}

func TestCollection_AddCapacity_Nil_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c.AddCapacity()
	if c.Capacity() == 0 { t.Fatal("expected > 0") }
}

func TestCollection_Resize_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.Resize(100)
	if c.Capacity() < 100 { t.Fatal("expected >= 100") }
}

func TestCollection_Resize_Smaller_C12(t *testing.T) {
	c := New.Collection.Cap(100)
	c.Resize(5) // should be no-op
}

// ── Joins, NonEmptyJoins, NonWhitespaceJoins ──

func TestCollection_Joins_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if c.Joins(",") == "" { t.Fatal("expected non-empty") }
}

func TestCollection_Joins_WithExtra_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.Joins(",", "b", "c")
	if r == "" { t.Fatal("expected non-empty") }
}

func TestCollection_NonEmptyJoins_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "", "b"})
	r := c.NonEmptyJoins(",")
	if r == "" { t.Fatal("expected non-empty") }
}

func TestCollection_NonWhitespaceJoins_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "  ", "b"})
	r := c.NonWhitespaceJoins(",")
	if r == "" { t.Fatal("expected non-empty") }
}

// ── Join, JoinLine ──

func TestCollection_Join_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if c.Join(",") != "a,b" { t.Fatal("unexpected") }
}

func TestCollection_Join_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	if c.Join(",") != "" { t.Fatal("expected empty") }
}

func TestCollection_JoinLine_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	r := c.JoinLine()
	if r == "" { t.Fatal("expected non-empty") }
}

func TestCollection_JoinLine_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	if c.JoinLine() != "" { t.Fatal("expected empty") }
}

// ── JSON methods ──

func TestCollection_JsonModel_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if len(c.JsonModel()) != 1 { t.Fatal("expected 1") }
}

func TestCollection_JsonModelAny_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	_ = c.JsonModelAny()
}

func TestCollection_MarshalJSON_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	b, err := c.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestCollection_UnmarshalJSON_C12(t *testing.T) {
	c := &Collection{}
	err := c.UnmarshalJSON([]byte(`["a","b"]`))
	if err != nil || c.Length() != 2 { t.Fatal("unexpected") }
}

func TestCollection_UnmarshalJSON_Error_C12(t *testing.T) {
	c := &Collection{}
	err := c.UnmarshalJSON([]byte(`invalid`))
	if err == nil { t.Fatal("expected error") }
}

func TestCollection_Json_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.Json()
	if r.HasError() { t.Fatal("unexpected error") }
}

func TestCollection_JsonPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.JsonPtr()
	if r.HasError() { t.Fatal("unexpected error") }
}

// ── ParseInjectUsingJson ──

func TestCollection_ParseInjectUsingJson_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	jr := c.Json()
	c2 := &Collection{}
	result, err := c2.ParseInjectUsingJson(&jr)
	if err != nil || result.Length() != 1 { t.Fatal("unexpected") }
}

func TestCollection_ParseInjectUsingJsonMust_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	jr := c.Json()
	c2 := &Collection{}
	result := c2.ParseInjectUsingJsonMust(&jr)
	if result.Length() != 1 { t.Fatal("unexpected") }
}

func TestCollection_JsonParseSelfInject_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	jr := c.Json()
	c2 := &Collection{}
	err := c2.JsonParseSelfInject(&jr)
	if err != nil { t.Fatal("unexpected") }
}

// ── Clear, Dispose ──

func TestCollection_Clear_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	c.Clear()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestCollection_Clear_Nil_C12(t *testing.T) {
	var c *Collection
	r := c.Clear()
	if r != nil { t.Fatal("expected nil") }
}

func TestCollection_Dispose_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	c.Dispose()
	if c.items != nil { t.Fatal("expected nil") }
}

func TestCollection_Dispose_Nil_C12(t *testing.T) {
	var c *Collection
	c.Dispose() // should not panic
}

// ── Serialize, Deserialize ──

func TestCollection_Serialize_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	b, err := c.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestCollection_Deserialize_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	var target []string
	err := c.Deserialize(&target)
	if err != nil || len(target) != 2 { t.Fatal("unexpected") }
}

// ── AsJsonMarshaller, AsJsonContractsBinder ──

func TestCollection_AsJsonMarshaller_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	_ = c.AsJsonMarshaller()
}

func TestCollection_AsJsonContractsBinder_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	_ = c.AsJsonContractsBinder()
}

// ── AddPointerCollectionsLock ──

func TestCollection_AddPointerCollectionsLock_C12(t *testing.T) {
	c := New.Collection.Cap(5)
	c2 := New.Collection.Strings([]string{"a"})
	c.AddPointerCollectionsLock(c2)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── ListCopyPtrLock ──

func TestCollection_ListCopyPtrLock_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	r := c.ListCopyPtrLock()
	if len(r) != 1 { t.Fatal("expected 1") }
}

func TestCollection_ListCopyPtrLock_Empty_C12(t *testing.T) {
	c := New.Collection.Cap(0)
	r := c.ListCopyPtrLock()
	if len(r) != 0 { t.Fatal("expected 0") }
}

// ── Items, ListPtr ──

func TestCollection_Items_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if len(c.Items()) != 1 { t.Fatal("expected 1") }
}

func TestCollection_ListPtr_C12(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if len(c.ListPtr()) != 1 { t.Fatal("expected 1") }
}
