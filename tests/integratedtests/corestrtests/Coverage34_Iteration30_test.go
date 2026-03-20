package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — IsEmpty / HasItems / Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_IsEmpty_New(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"empty": hs.IsEmpty(), "items": hs.HasItems(), "len": hs.Length(), "hasAny": hs.HasAnyItem()}
	expected := args.Map{"empty": true, "items": false, "len": 0, "hasAny": false}
	expected.ShouldBeEqual(t, 0, "Hashset empty", actual)
}

func Test_I30_Hashset_Length_Nil(t *testing.T) {
	var hs *corestr.Hashset
	actual := args.Map{"len": hs.Length(), "empty": hs.IsEmpty()}
	expected := args.Map{"len": 0, "empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset nil length", actual)
}

func Test_I30_Hashset_IsEmptyLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"empty": hs.IsEmptyLock()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset IsEmptyLock", actual)
}

func Test_I30_Hashset_LengthLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.Add("a")
	actual := args.Map{"len": hs.LengthLock()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset LengthLock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Add variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_Add(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.Add("a")
	actual := args.Map{"has": hs.Has("a"), "len": hs.Length()}
	expected := args.Map{"has": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset Add", actual)
}

func Test_I30_Hashset_AddBool(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	existed1 := hs.AddBool("a")
	existed2 := hs.AddBool("a")
	actual := args.Map{"existed1": existed1, "existed2": existed2}
	expected := args.Map{"existed1": false, "existed2": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddBool", actual)
}

func Test_I30_Hashset_AddNonEmpty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddNonEmpty("")
	hs.AddNonEmpty("a")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset AddNonEmpty", actual)
}

func Test_I30_Hashset_AddNonEmptyWhitespace(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddNonEmptyWhitespace("   ")
	hs.AddNonEmptyWhitespace("a")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset AddNonEmptyWhitespace", actual)
}

func Test_I30_Hashset_AddIf(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddIf(true, "a")
	hs.AddIf(false, "b")
	actual := args.Map{"hasA": hs.Has("a"), "hasB": hs.Has("b")}
	expected := args.Map{"hasA": true, "hasB": false}
	expected.ShouldBeEqual(t, 0, "Hashset AddIf", actual)
}

func Test_I30_Hashset_AddIfMany(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddIfMany(true, "a", "b")
	hs.AddIfMany(false, "c")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset AddIfMany", actual)
}

func Test_I30_Hashset_AddFunc(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddFunc(func() string { return "computed" })
	actual := args.Map{"has": hs.Has("computed")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddFunc", actual)
}

func Test_I30_Hashset_AddFuncErr_NoErr(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
	actual := args.Map{"has": hs.Has("ok")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddFuncErr no err", actual)
}

func Test_I30_Hashset_AddFuncErr_Err(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	called := false
	hs.AddFuncErr(func() (string, error) { return "", fmt.Errorf("fail") }, func(e error) { called = true })
	actual := args.Map{"empty": hs.IsEmpty(), "called": called}
	expected := args.Map{"empty": true, "called": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddFuncErr err", actual)
}

func Test_I30_Hashset_AddLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddLock("a")
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddLock", actual)
}

func Test_I30_Hashset_AddPtr(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	s := "hello"
	hs.AddPtr(&s)
	actual := args.Map{"has": hs.Has("hello")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddPtr", actual)
}

func Test_I30_Hashset_AddPtrLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	s := "hello"
	hs.AddPtrLock(&s)
	actual := args.Map{"has": hs.Has("hello")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddPtrLock", actual)
}

func Test_I30_Hashset_Adds(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.Adds("a", "b")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset Adds", actual)
}

func Test_I30_Hashset_Adds_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.Adds(nil...)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset Adds nil", actual)
}

func Test_I30_Hashset_AddStrings(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddStrings([]string{"a", "b"})
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset AddStrings", actual)
}

func Test_I30_Hashset_AddStrings_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddStrings(nil)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddStrings nil", actual)
}

func Test_I30_Hashset_AddStringsLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddStringsLock([]string{"a"})
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddStringsLock", actual)
}

func Test_I30_Hashset_AddStringsLock_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddStringsLock(nil)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddStringsLock nil", actual)
}

func Test_I30_Hashset_AddHashsetItems(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	other := corestr.New.Hashset.Strings([]string{"a", "b"})
	hs.AddHashsetItems(other)
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset AddHashsetItems", actual)
}

func Test_I30_Hashset_AddHashsetItems_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddHashsetItems(nil)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddHashsetItems nil", actual)
}

func Test_I30_Hashset_AddItemsMap(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddItemsMap(map[string]bool{"a": true, "b": false})
	actual := args.Map{"hasA": hs.Has("a"), "hasB": hs.Has("b")}
	expected := args.Map{"hasA": true, "hasB": false}
	expected.ShouldBeEqual(t, 0, "Hashset AddItemsMap", actual)
}

func Test_I30_Hashset_AddItemsMap_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddItemsMap(nil)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddItemsMap nil", actual)
}

func Test_I30_Hashset_AddCollection(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	coll := corestr.New.Collection.Strings([]string{"a"})
	hs.AddCollection(coll)
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddCollection", actual)
}

func Test_I30_Hashset_AddCollection_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddCollection(nil)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddCollection nil", actual)
}

func Test_I30_Hashset_AddCollections(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	c1 := corestr.New.Collection.Strings([]string{"a"})
	c2 := corestr.New.Collection.Strings([]string{"b"})
	hs.AddCollections(c1, nil, c2)
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset AddCollections", actual)
}

func Test_I30_Hashset_AddCollections_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddCollections(nil...)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddCollections nil", actual)
}

func Test_I30_Hashset_AddWithWgLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	hs.AddWithWgLock("a", wg)
	wg.Wait()
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddWithWgLock", actual)
}

func Test_I30_Hashset_AddSimpleSlice(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	ss := corestr.SimpleSlice([]string{"a", "b"})
	hs.AddSimpleSlice(&ss)
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset AddSimpleSlice", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Has / Contains / Missing
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_Has_Contains(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"has": hs.Has("a"), "contains": hs.Contains("a"), "missing": hs.IsMissing("b")}
	expected := args.Map{"has": true, "contains": true, "missing": true}
	expected.ShouldBeEqual(t, 0, "Hashset Has/Contains/IsMissing", actual)
}

func Test_I30_Hashset_HasLock(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"hl": hs.HasLock("a"), "hwl": hs.HasWithLock("a"), "ml": hs.IsMissingLock("z")}
	expected := args.Map{"hl": true, "hwl": true, "ml": true}
	expected.ShouldBeEqual(t, 0, "Hashset lock variants", actual)
}

func Test_I30_Hashset_HasAllStrings(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	actual := args.Map{"all": hs.HasAllStrings([]string{"a", "b"}), "miss": hs.HasAllStrings([]string{"a", "c"})}
	expected := args.Map{"all": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "Hashset HasAllStrings", actual)
}

func Test_I30_Hashset_HasAll(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	actual := args.Map{"all": hs.HasAll("a", "b"), "miss": hs.HasAll("a", "c")}
	expected := args.Map{"all": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "Hashset HasAll", actual)
}

func Test_I30_Hashset_HasAny(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"any": hs.HasAny("z", "a"), "none": hs.HasAny("x", "y")}
	expected := args.Map{"any": true, "none": false}
	expected.ShouldBeEqual(t, 0, "Hashset HasAny", actual)
}

func Test_I30_Hashset_IsAllMissing(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"allMiss": hs.IsAllMissing("x", "y"), "notAll": hs.IsAllMissing("a", "x")}
	expected := args.Map{"allMiss": true, "notAll": false}
	expected.ShouldBeEqual(t, 0, "Hashset IsAllMissing", actual)
}

func Test_I30_Hashset_HasAllCollectionItems(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	coll := corestr.New.Collection.Strings([]string{"a"})
	actual := args.Map{"has": hs.HasAllCollectionItems(coll), "nil": hs.HasAllCollectionItems(nil)}
	expected := args.Map{"has": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "Hashset HasAllCollectionItems", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — List / Items / Collection / Sorted
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_List(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"len": len(hs.List())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset List", actual)
}

func Test_I30_Hashset_ListPtr(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"len": len(hs.ListPtr())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset ListPtr", actual)
}

func Test_I30_Hashset_ListCopyLock(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"len": len(hs.ListCopyLock())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset ListCopyLock", actual)
}

func Test_I30_Hashset_Items(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"len": len(hs.Items())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset Items", actual)
}

func Test_I30_Hashset_Collection(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	coll := hs.Collection()
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset Collection", actual)
}

func Test_I30_Hashset_SortedList(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"b", "a"})
	sorted := hs.SortedList()
	actual := args.Map{"first": sorted[0], "second": sorted[1]}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "Hashset SortedList", actual)
}

func Test_I30_Hashset_OrderedList(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"b", "a"})
	ol := hs.OrderedList()
	actual := args.Map{"len": len(ol)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset OrderedList", actual)
}

func Test_I30_Hashset_OrderedList_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	ol := hs.OrderedList()
	actual := args.Map{"len": len(ol)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset OrderedList empty", actual)
}

func Test_I30_Hashset_ListPtrSortedAsc(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"b", "a"})
	sorted := hs.ListPtrSortedAsc()
	actual := args.Map{"first": sorted[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Hashset ListPtrSortedAsc", actual)
}

func Test_I30_Hashset_ListPtrSortedDsc(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	sorted := hs.ListPtrSortedDsc()
	actual := args.Map{"first": sorted[0]}
	expected := args.Map{"first": "b"}
	expected.ShouldBeEqual(t, 0, "Hashset ListPtrSortedDsc", actual)
}

func Test_I30_Hashset_SafeStrings(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"len": len(hs.SafeStrings())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset SafeStrings empty", actual)
}

func Test_I30_Hashset_Lines(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"len": len(hs.Lines())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset Lines empty", actual)
}

func Test_I30_Hashset_SimpleSlice(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	ss := hs.SimpleSlice()
	actual := args.Map{"len": ss.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset SimpleSlice", actual)
}

func Test_I30_Hashset_SimpleSlice_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	ss := hs.SimpleSlice()
	actual := args.Map{"empty": ss.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset SimpleSlice empty", actual)
}

func Test_I30_Hashset_MapStringAny(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	m := hs.MapStringAny()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset MapStringAny", actual)
}

func Test_I30_Hashset_MapStringAny_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	m := hs.MapStringAny()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset MapStringAny empty", actual)
}

func Test_I30_Hashset_MapStringAnyDiff(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	d := hs.MapStringAnyDiff()
	actual := args.Map{"notNil": d != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset MapStringAnyDiff", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Resize / AddCapacities
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_Resize(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.Resize(100)
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset Resize", actual)
}

func Test_I30_Hashset_Resize_AlreadyLarger(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	hs.Resize(1)
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Hashset Resize already larger", actual)
}

func Test_I30_Hashset_ResizeLock(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.ResizeLock(100)
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset ResizeLock", actual)
}

func Test_I30_Hashset_ResizeLock_AlreadyLarger(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
	hs.ResizeLock(1)
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Hashset ResizeLock already larger", actual)
}

func Test_I30_Hashset_AddCapacities(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.AddCapacities(10, 20)
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddCapacities", actual)
}

func Test_I30_Hashset_AddCapacities_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.AddCapacities()
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddCapacities empty", actual)
}

func Test_I30_Hashset_AddCapacitiesLock(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.AddCapacitiesLock(10)
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddCapacitiesLock", actual)
}

func Test_I30_Hashset_AddCapacitiesLock_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.AddCapacitiesLock()
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddCapacitiesLock empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — ConcatNew
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_ConcatNewHashsets_NoArgs(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	c := hs.ConcatNewHashsets(true)
	actual := args.Map{"has": c.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset ConcatNewHashsets no args", actual)
}

func Test_I30_Hashset_ConcatNewHashsets_WithArgs(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	other := corestr.New.Hashset.Strings([]string{"b"})
	c := hs.ConcatNewHashsets(true, other, nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset ConcatNewHashsets with args", actual)
}

func Test_I30_Hashset_ConcatNewStrings_NoArgs(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	c := hs.ConcatNewStrings(true)
	actual := args.Map{"has": c.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset ConcatNewStrings no args", actual)
}

func Test_I30_Hashset_ConcatNewStrings_WithArgs(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	c := hs.ConcatNewStrings(true, []string{"b", "c"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Hashset ConcatNewStrings with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Filter / GetFiltered / GetAllExcept
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_Filter(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"abc", "x"})
	filtered := hs.Filter(func(s string) bool { return len(s) > 1 })
	actual := args.Map{"len": filtered.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset Filter", actual)
}

func Test_I30_Hashset_GetFilteredItems(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetFilteredItems", actual)
}

func Test_I30_Hashset_GetFilteredItems_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset GetFilteredItems empty", actual)
}

func Test_I30_Hashset_GetFilteredItems_Break(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, true })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetFilteredItems break", actual)
}

func Test_I30_Hashset_GetFilteredCollection(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetFilteredCollection", actual)
}

func Test_I30_Hashset_GetFilteredCollection_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
	actual := args.Map{"empty": coll.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset GetFilteredCollection empty", actual)
}

func Test_I30_Hashset_GetFilteredCollection_Break(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, true })
	actual := args.Map{"len": coll.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetFilteredCollection break", actual)
}

func Test_I30_Hashset_GetAllExceptHashset(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	exc := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.GetAllExceptHashset(exc)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExceptHashset", actual)
}

func Test_I30_Hashset_GetAllExceptHashset_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.GetAllExceptHashset(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExceptHashset nil", actual)
}

func Test_I30_Hashset_GetAllExcept(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetAllExcept([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExcept", actual)
}

func Test_I30_Hashset_GetAllExcept_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.GetAllExcept(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExcept nil", actual)
}

func Test_I30_Hashset_GetAllExceptSpread(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	result := hs.GetAllExceptSpread("a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExceptSpread", actual)
}

func Test_I30_Hashset_GetAllExceptSpread_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.GetAllExceptSpread(nil...)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExceptSpread nil", actual)
}

func Test_I30_Hashset_GetAllExceptCollection(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	coll := corestr.New.Collection.Strings([]string{"a"})
	result := hs.GetAllExceptCollection(coll)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExceptCollection", actual)
}

func Test_I30_Hashset_GetAllExceptCollection_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.GetAllExceptCollection(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset GetAllExceptCollection nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — AddsUsingFilter variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_AddsUsingFilter(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, "a", "b")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset AddsUsingFilter", actual)
}

func Test_I30_Hashset_AddsUsingFilter_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsUsingFilter(nil, nil...)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddsUsingFilter nil", actual)
}

func Test_I30_Hashset_AddsUsingFilter_Break(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset AddsUsingFilter break", actual)
}

func Test_I30_Hashset_AddsAnyUsingFilter_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsAnyUsingFilter(nil, nil...)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddsAnyUsingFilter nil", actual)
}

func Test_I30_Hashset_AddsAnyUsingFilter_NilItem(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, nil, "hello")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset AddsAnyUsingFilter nil item", actual)
}

func Test_I30_Hashset_AddsAnyUsingFilter_Break(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset AddsAnyUsingFilter break", actual)
}

func Test_I30_Hashset_AddsAnyUsingFilterLock_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsAnyUsingFilterLock(nil, nil...)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddsAnyUsingFilterLock nil", actual)
}

func Test_I30_Hashset_AddsAnyUsingFilterLock_Break(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	hs.AddsAnyUsingFilterLock(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset AddsAnyUsingFilterLock break", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Remove / Clear / Dispose
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_Remove(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	hs.Remove("a")
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Hashset Remove", actual)
}

func Test_I30_Hashset_SafeRemove(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.SafeRemove("a")
	hs.SafeRemove("missing")
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset SafeRemove", actual)
}

func Test_I30_Hashset_RemoveWithLock(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.RemoveWithLock("a")
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset RemoveWithLock", actual)
}

func Test_I30_Hashset_Clear(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.Clear()
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset Clear", actual)
}

func Test_I30_Hashset_Clear_Nil(t *testing.T) {
	var hs *corestr.Hashset
	result := hs.Clear()
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Hashset Clear nil", actual)
}

func Test_I30_Hashset_Dispose(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	hs.Dispose()
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset Dispose", actual)
}

func Test_I30_Hashset_Dispose_Nil(t *testing.T) {
	var hs *corestr.Hashset
	hs.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Hashset Dispose nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — IsEquals / ToLowerSet
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_IsEquals_Same(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"same": hs.IsEquals(hs)}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Hashset IsEquals same ptr", actual)
}

func Test_I30_Hashset_IsEquals_BothNil(t *testing.T) {
	var hs *corestr.Hashset
	actual := args.Map{"eq": hs.IsEquals(nil)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashset IsEquals both nil", actual)
}

func Test_I30_Hashset_IsEquals_OneNil(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"eq": hs.IsEquals(nil)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "Hashset IsEquals one nil", actual)
}

func Test_I30_Hashset_IsEquals_BothEmpty(t *testing.T) {
	hs1 := corestr.New.Hashset.Cap(5)
	hs2 := corestr.New.Hashset.Cap(5)
	actual := args.Map{"eq": hs1.IsEquals(hs2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashset IsEquals both empty", actual)
}

func Test_I30_Hashset_IsEquals_DiffLen(t *testing.T) {
	hs1 := corestr.New.Hashset.Strings([]string{"a"})
	hs2 := corestr.New.Hashset.Strings([]string{"a", "b"})
	actual := args.Map{"eq": hs1.IsEquals(hs2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "Hashset IsEquals diff len", actual)
}

func Test_I30_Hashset_IsEquals_DiffItems(t *testing.T) {
	hs1 := corestr.New.Hashset.Strings([]string{"a"})
	hs2 := corestr.New.Hashset.Strings([]string{"b"})
	actual := args.Map{"eq": hs1.IsEquals(hs2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "Hashset IsEquals diff items", actual)
}

func Test_I30_Hashset_IsEqual(t *testing.T) {
	hs1 := corestr.New.Hashset.Strings([]string{"a"})
	hs2 := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"eq": hs1.IsEqual(hs2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashset IsEqual", actual)
}

func Test_I30_Hashset_IsEqualsLock(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"eq": hs.IsEqualsLock(hs)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "Hashset IsEqualsLock", actual)
}

func Test_I30_Hashset_ToLowerSet(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"ABC"})
	lower := hs.ToLowerSet()
	actual := args.Map{"has": lower.Has("abc")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset ToLowerSet", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — String / Join
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_String_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"notEmpty": hs.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset String empty", actual)
}

func Test_I30_Hashset_String_WithItems(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"notEmpty": hs.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset String with items", actual)
}

func Test_I30_Hashset_StringLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"notEmpty": hs.StringLock() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset StringLock empty", actual)
}

func Test_I30_Hashset_StringLock_WithItems(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"notEmpty": hs.StringLock() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset StringLock with items", actual)
}

func Test_I30_Hashset_Join(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"val": hs.Join(",")}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "Hashset Join", actual)
}

func Test_I30_Hashset_JoinLine(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"val": hs.JoinLine()}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "Hashset JoinLine", actual)
}

func Test_I30_Hashset_JoinSorted_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"val": hs.JoinSorted(",")}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Hashset JoinSorted empty", actual)
}

func Test_I30_Hashset_JoinSorted(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"b", "a"})
	actual := args.Map{"val": hs.JoinSorted(",")}
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "Hashset JoinSorted", actual)
}

func Test_I30_Hashset_NonEmptyJoins(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"notEmpty": hs.NonEmptyJoins(",") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset NonEmptyJoins", actual)
}

func Test_I30_Hashset_NonWhitespaceJoins(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"notEmpty": hs.NonWhitespaceJoins(",") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset NonWhitespaceJoins", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — JSON / Serialize
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_JsonModel(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	jm := hs.JsonModel()
	actual := args.Map{"len": len(jm)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset JsonModel", actual)
}

func Test_I30_Hashset_JsonModel_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	jm := hs.JsonModel()
	actual := args.Map{"len": len(jm)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset JsonModel empty", actual)
}

func Test_I30_Hashset_JsonModelAny(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	actual := args.Map{"notNil": hs.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset JsonModelAny", actual)
}

func Test_I30_Hashset_MarshalJSON(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	b, err := hs.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Hashset MarshalJSON", actual)
}

func Test_I30_Hashset_UnmarshalJSON(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	err := hs.UnmarshalJSON([]byte(`{"a":true}`))
	actual := args.Map{"noErr": err == nil, "has": hs.Has("a")}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "Hashset UnmarshalJSON", actual)
}

func Test_I30_Hashset_UnmarshalJSON_Err(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	err := hs.UnmarshalJSON([]byte(`{invalid`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Hashset UnmarshalJSON err", actual)
}

func Test_I30_Hashset_Json(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	j := hs.Json()
	actual := args.Map{"hasBytes": j.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Hashset Json", actual)
}

func Test_I30_Hashset_JsonPtr(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	jp := hs.JsonPtr()
	actual := args.Map{"notNil": jp != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset JsonPtr", actual)
}

func Test_I30_Hashset_Serialize(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	b, err := hs.Serialize()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Hashset Serialize", actual)
}

func Test_I30_Hashset_Deserialize(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	target := map[string]bool{}
	err := hs.Deserialize(&target)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Hashset Deserialize", actual)
}

func Test_I30_Hashset_ParseInjectUsingJson(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	jr := hs.JsonPtr()
	hs2 := corestr.New.Hashset.Cap(5)
	result, err := hs2.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "has": result.Has("a")}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "Hashset ParseInjectUsingJson", actual)
}

func Test_I30_Hashset_ParseInjectUsingJson_Err(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	badJson := corejson.NewPtr(42)
	_, err := hs.ParseInjectUsingJson(badJson)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Hashset ParseInjectUsingJson err", actual)
}

func Test_I30_Hashset_JsonParseSelfInject(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	jr := hs.JsonPtr()
	hs2 := corestr.New.Hashset.Cap(5)
	err := hs2.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Hashset JsonParseSelfInject", actual)
}

func Test_I30_Hashset_AsJsoner(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"notNil": hs.AsJsoner() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset AsJsoner", actual)
}

func Test_I30_Hashset_AsJsonContractsBinder(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"notNil": hs.AsJsonContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset AsJsonContractsBinder", actual)
}

func Test_I30_Hashset_AsJsonParseSelfInjector(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"notNil": hs.AsJsonParseSelfInjector() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset AsJsonParseSelfInjector", actual)
}

func Test_I30_Hashset_AsJsonMarshaller(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	actual := args.Map{"notNil": hs.AsJsonMarshaller() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset AsJsonMarshaller", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — DistinctDiff
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	result := hs.DistinctDiffLinesRaw()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLinesRaw both empty", actual)
}

func Test_I30_Hashset_DistinctDiffLinesRaw_LeftOnly(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.DistinctDiffLinesRaw()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLinesRaw left only", actual)
}

func Test_I30_Hashset_DistinctDiffLinesRaw_RightOnly(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	result := hs.DistinctDiffLinesRaw("a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLinesRaw right only", actual)
}

func Test_I30_Hashset_DistinctDiffLinesRaw_Both(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	result := hs.DistinctDiffLinesRaw("b", "c")
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLinesRaw both", actual)
}

func Test_I30_Hashset_DistinctDiffLines_BothEmpty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	result := hs.DistinctDiffLines()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLines both empty", actual)
}

func Test_I30_Hashset_DistinctDiffLines_LeftOnly(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.DistinctDiffLines()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLines left only", actual)
}

func Test_I30_Hashset_DistinctDiffLines_RightOnly(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	result := hs.DistinctDiffLines("a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLines right only", actual)
}

func Test_I30_Hashset_DistinctDiffLines_Both(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	result := hs.DistinctDiffLines("b", "c")
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffLines both", actual)
}

func Test_I30_Hashset_DistinctDiffHashset(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a", "b"})
	other := corestr.New.Hashset.Strings([]string{"b", "c"})
	result := hs.DistinctDiffHashset(other)
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Hashset DistinctDiffHashset", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Wrap / Transpile
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_WrapDoubleQuote(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.WrapDoubleQuote()
	actual := args.Map{"hasAny": result.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "Hashset WrapDoubleQuote", actual)
}

func Test_I30_Hashset_WrapDoubleQuoteIfMissing(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.WrapDoubleQuoteIfMissing()
	actual := args.Map{"hasAny": result.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "Hashset WrapDoubleQuoteIfMissing", actual)
}

func Test_I30_Hashset_WrapSingleQuote(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.WrapSingleQuote()
	actual := args.Map{"hasAny": result.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "Hashset WrapSingleQuote", actual)
}

func Test_I30_Hashset_WrapSingleQuoteIfMissing(t *testing.T) {
	hs := corestr.New.Hashset.Strings([]string{"a"})
	result := hs.WrapSingleQuoteIfMissing()
	actual := args.Map{"hasAny": result.HasAnyItem()}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "Hashset WrapSingleQuoteIfMissing", actual)
}

func Test_I30_Hashset_Transpile_Empty(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	result := hs.Transpile(func(s string) string { return s })
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset Transpile empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — WgLock variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_AddStringsPtrWgLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(200)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
	wg.Wait()
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset AddStringsPtrWgLock", actual)
}

func Test_I30_Hashset_AddHashsetWgLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(200)
	other := corestr.New.Hashset.Strings([]string{"a"})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	hs.AddHashsetWgLock(other, wg)
	wg.Wait()
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddHashsetWgLock", actual)
}

func Test_I30_Hashset_AddHashsetWgLock_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	wg := &sync.WaitGroup{}
	hs.AddHashsetWgLock(nil, wg)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddHashsetWgLock nil", actual)
}

func Test_I30_Hashset_AddItemsMapWgLock(t *testing.T) {
	hs := corestr.New.Hashset.Cap(200)
	m := map[string]bool{"a": true, "b": false}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	hs.AddItemsMapWgLock(&m, wg)
	wg.Wait()
	actual := args.Map{"hasA": hs.Has("a"), "hasB": hs.Has("b")}
	expected := args.Map{"hasA": true, "hasB": false}
	expected.ShouldBeEqual(t, 0, "Hashset AddItemsMapWgLock", actual)
}

func Test_I30_Hashset_AddItemsMapWgLock_Nil(t *testing.T) {
	hs := corestr.New.Hashset.Cap(5)
	wg := &sync.WaitGroup{}
	hs.AddItemsMapWgLock(nil, wg)
	actual := args.Map{"empty": hs.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddItemsMapWgLock nil", actual)
}
