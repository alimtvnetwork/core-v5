package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// Collection — remove, append, filter, unique, sort, search
// ═══════════════════════════════════════════════════════════════

func Test_Cov48_Collection_RemoveItemsIndexes_Valid(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	c.RemoveItemsIndexes(true, 1)
	tc := coretestcases.CaseV1{Name: "RemoveItemsIndexes", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_RemoveItemsIndexes_NilIgnore(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	c.RemoveItemsIndexes(true)
	tc := coretestcases.CaseV1{Name: "RemoveItemsIndexes nil ignore", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendCollectionPtr(t *testing.T) {
	c1 := corestr.New.Collection.Strings([]string{"a"})
	c2 := corestr.New.Collection.Strings([]string{"b", "c"})
	c1.AppendCollectionPtr(c2)
	tc := coretestcases.CaseV1{Name: "AppendCollectionPtr", Expected: 3, Actual: c1.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendCollections(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	c2 := corestr.New.Collection.Strings([]string{"b"})
	c3 := corestr.New.Collection.Strings([]string{"c"})
	c.AppendCollections(c2, c3)
	tc := coretestcases.CaseV1{Name: "AppendCollections", Expected: 3, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendCollections_Empty(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	c.AppendCollections()
	tc := coretestcases.CaseV1{Name: "AppendCollections empty", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnys(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnys("hello", 42)
	tc := coretestcases.CaseV1{Name: "AppendAnys", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnys_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnys()
	tc := coretestcases.CaseV1{Name: "AppendAnys empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnys_NilSkip(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnys(nil, "ok")
	tc := coretestcases.CaseV1{Name: "AppendAnys nil skip", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnysLock(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnysLock("hello")
	tc := coretestcases.CaseV1{Name: "AppendAnysLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnysLock_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnysLock()
	tc := coretestcases.CaseV1{Name: "AppendAnysLock empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnysUsingFilter(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
		return s, true, false
	}, "a", "b")
	tc := coretestcases.CaseV1{Name: "AppendAnysUsingFilter", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
		return s, true, true
	}, "a", "b")
	tc := coretestcases.CaseV1{Name: "AppendAnysUsingFilter break", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendAnysUsingFilter_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendAnysUsingFilter(nil)
	tc := coretestcases.CaseV1{Name: "AppendAnysUsingFilter empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendNonEmptyAnys(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendNonEmptyAnys("hello", nil, "world")
	tc := coretestcases.CaseV1{Name: "AppendNonEmptyAnys", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AppendNonEmptyAnys(nil)
	tc := coretestcases.CaseV1{Name: "AppendNonEmptyAnys nil only", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddsNonEmpty(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddsNonEmpty("a", "", "b")
	tc := coretestcases.CaseV1{Name: "AddsNonEmpty", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddsNonEmpty_Nil(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddsNonEmpty()
	tc := coretestcases.CaseV1{Name: "AddsNonEmpty nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	s := "hello"
	c.AddsNonEmptyPtrLock(&s, nil)
	tc := coretestcases.CaseV1{Name: "AddsNonEmptyPtrLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_UniqueBoolMap(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
	m := c.UniqueBoolMap()
	tc := coretestcases.CaseV1{Name: "UniqueBoolMap", Expected: 2, Actual: len(m), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_UniqueBoolMapLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "a"})
	m := c.UniqueBoolMapLock()
	tc := coretestcases.CaseV1{Name: "UniqueBoolMapLock", Expected: 1, Actual: len(m), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_UniqueList(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
	list := c.UniqueList()
	tc := coretestcases.CaseV1{Name: "UniqueList", Expected: 2, Actual: len(list), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_UniqueListLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "b"})
	list := c.UniqueListLock()
	tc := coretestcases.CaseV1{Name: "UniqueListLock", Expected: 2, Actual: len(list), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_Filter(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"abc", "def", "ab"})
	result := c.Filter(func(s string, i int) (string, bool, bool) {
		return s, len(s) == 3, false
	})
	tc := coretestcases.CaseV1{Name: "Filter", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_Filter_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	result := c.Filter(func(s string, i int) (string, bool, bool) {
		return s, true, false
	})
	tc := coretestcases.CaseV1{Name: "Filter empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_Filter_Break(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	result := c.Filter(func(s string, i int) (string, bool, bool) {
		return s, true, i == 0
	})
	tc := coretestcases.CaseV1{Name: "Filter break", Expected: 1, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_FilteredCollection(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "bb", "c"})
	result := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
		return s, len(s) == 1, false
	})
	tc := coretestcases.CaseV1{Name: "FilteredCollection", Expected: 2, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_Has_Found(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "Has found", Expected: true, Actual: c.Has("b"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_Has_NotFound(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "Has not found", Expected: false, Actual: c.Has("z"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_Has_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "Has empty", Expected: false, Actual: c.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "HasLock", Expected: true, Actual: c.HasLock("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasPtr_Found(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	s := "a"
	tc := coretestcases.CaseV1{Name: "HasPtr found", Expected: true, Actual: c.HasPtr(&s), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasPtr_Nil(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "HasPtr nil", Expected: false, Actual: c.HasPtr(nil), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasAll(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{Name: "HasAll", Expected: true, Actual: c.HasAll("a", "c"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasAll_Missing(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "HasAll missing", Expected: false, Actual: c.HasAll("a", "z"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasAll_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "HasAll empty", Expected: false, Actual: c.HasAll("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_SortedListAsc(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
	result := c.SortedListAsc()
	tc := coretestcases.CaseV1{Name: "SortedListAsc first", Expected: "a", Actual: result[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_SortedListAsc_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	result := c.SortedListAsc()
	tc := coretestcases.CaseV1{Name: "SortedListAsc empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_SortedAsc(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
	c.SortedAsc()
	tc := coretestcases.CaseV1{Name: "SortedAsc", Expected: "a", Actual: c.First(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_SortedAsc_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	c.SortedAsc()
	tc := coretestcases.CaseV1{Name: "SortedAsc empty", Expected: true, Actual: c.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_SortedAscLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"c", "a"})
	c.SortedAscLock()
	tc := coretestcases.CaseV1{Name: "SortedAscLock", Expected: "a", Actual: c.First(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_SortedListDsc(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "c", "b"})
	result := c.SortedListDsc()
	tc := coretestcases.CaseV1{Name: "SortedListDsc first", Expected: "c", Actual: result[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasUsingSensitivity_CaseSensitive(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"Hello"})
	tc := coretestcases.CaseV1{Name: "HasUsingSensitivity sensitive", Expected: false, Actual: c.HasUsingSensitivity("hello", true), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HasUsingSensitivity_CaseInsensitive(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"Hello"})
	tc := coretestcases.CaseV1{Name: "HasUsingSensitivity insensitive", Expected: true, Actual: c.HasUsingSensitivity("hello", false), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_IsContainsPtr_Found(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	s := "b"
	tc := coretestcases.CaseV1{Name: "IsContainsPtr found", Expected: true, Actual: c.IsContainsPtr(&s), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_IsContainsPtr_Nil(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "IsContainsPtr nil", Expected: false, Actual: c.IsContainsPtr(nil), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_IsContainsAll(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	tc := coretestcases.CaseV1{Name: "IsContainsAll", Expected: true, Actual: c.IsContainsAll("a", "b"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_IsContainsAll_Missing(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "IsContainsAll missing", Expected: false, Actual: c.IsContainsAll("a", "z"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_IsContainsAllSlice_Empty(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "IsContainsAllSlice empty", Expected: false, Actual: c.IsContainsAllSlice([]string{}), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_IsContainsAllLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "IsContainsAllLock", Expected: true, Actual: c.IsContainsAllLock("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_GetHashsetPlusHasAll(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "GetHashsetPlusHasAll", Expected: true, Actual: hasAll, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "GetHashsetPlusHasAll hs", Expected: true, Actual: hs != nil, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov48_Collection_GetHashsetPlusHasAll_Nil(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	_, hasAll := c.GetHashsetPlusHasAll(nil)
	tc := coretestcases.CaseV1{Name: "GetHashsetPlusHasAll nil", Expected: false, Actual: hasAll, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_List(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "List", Expected: 2, Actual: len(c.List()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_Items(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "Items", Expected: 1, Actual: len(c.Items()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_ListPtr(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "ListPtr", Expected: 1, Actual: len(c.ListPtr()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_ListCopyPtrLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "ListCopyPtrLock", Expected: 1, Actual: len(c.ListCopyPtrLock()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_ListCopyPtrLock_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "ListCopyPtrLock empty", Expected: 0, Actual: len(c.ListCopyPtrLock()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_NonEmptyList(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "", "b"})
	result := c.NonEmptyList()
	tc := coretestcases.CaseV1{Name: "NonEmptyList", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_NonEmptyList_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	result := c.NonEmptyList()
	tc := coretestcases.CaseV1{Name: "NonEmptyList empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_NonEmptyListPtr(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", ""})
	result := c.NonEmptyListPtr()
	tc := coretestcases.CaseV1{Name: "NonEmptyListPtr", Expected: 1, Actual: len(*result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HashsetAsIs(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	hs := c.HashsetAsIs()
	tc := coretestcases.CaseV1{Name: "HashsetAsIs", Expected: true, Actual: hs.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HashsetWithDoubleLength(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	hs := c.HashsetWithDoubleLength()
	tc := coretestcases.CaseV1{Name: "HashsetWithDoubleLength", Expected: true, Actual: hs.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_HashsetLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	hs := c.HashsetLock()
	tc := coretestcases.CaseV1{Name: "HashsetLock", Expected: true, Actual: hs.Has("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_NonEmptyItems(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "", "b"})
	result := c.NonEmptyItems()
	tc := coretestcases.CaseV1{Name: "NonEmptyItems", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
	result := c.NonEmptyItemsOrNonWhitespace()
	tc := coretestcases.CaseV1{Name: "NonEmptyItemsOrNonWhitespace", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_New(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	result := c.New("a", "b")
	tc := coretestcases.CaseV1{Name: "New", Expected: 2, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_New_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	result := c.New()
	tc := coretestcases.CaseV1{Name: "New empty", Expected: 0, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddNonEmptyStrings(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddNonEmptyStrings("a", "", "b")
	tc := coretestcases.CaseV1{Name: "AddNonEmptyStrings", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddNonEmptyStrings_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddNonEmptyStrings()
	tc := coretestcases.CaseV1{Name: "AddNonEmptyStrings empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddFuncResult(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddFuncResult(func() string { return "hello" })
	tc := coretestcases.CaseV1{Name: "AddFuncResult", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddFuncResult_Nil(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddFuncResult()
	tc := coretestcases.CaseV1{Name: "AddFuncResult nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddStringsByFuncChecking(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddStringsByFuncChecking([]string{"ok", "bad", "ok2"}, func(s string) bool {
		return s != "bad"
	})
	tc := coretestcases.CaseV1{Name: "AddStringsByFuncChecking", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_MergeSlicesOfSlice(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
	tc := coretestcases.CaseV1{Name: "MergeSlicesOfSlice", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_GetAllExceptCollection(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
	except := corestr.New.Collection.Strings([]string{"b"})
	result := c.GetAllExceptCollection(except)
	tc := coretestcases.CaseV1{Name: "GetAllExceptCollection", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_GetAllExceptCollection_Nil(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	result := c.GetAllExceptCollection(nil)
	tc := coretestcases.CaseV1{Name: "GetAllExceptCollection nil", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov48_Collection_AddPointerCollectionsLock(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c2 := corestr.New.Collection.Strings([]string{"a"})
	c.AddPointerCollectionsLock(c2)
	tc := coretestcases.CaseV1{Name: "AddPointerCollectionsLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
