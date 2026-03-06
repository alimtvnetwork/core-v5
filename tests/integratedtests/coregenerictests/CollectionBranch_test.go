package coregenerictests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Collection — ForEach
// ==========================================================================

func Test_Collection_ForEach_VisitsAll(t *testing.T) {
	tc := collectionForEachVisitsAllTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	visited := 0
	var firstEntry, lastEntry string

	col.ForEach(func(index int, item int) {
		if visited == 0 {
			firstEntry = fmt.Sprintf("%d:%d", index, item)
		}
		lastEntry = fmt.Sprintf("%d:%d", index, item)
		visited++
	})

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", visited),
		firstEntry,
		lastEntry,
	)
}

func Test_Collection_ForEach_Empty(t *testing.T) {
	tc := collectionForEachEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	visited := 0
	col.ForEach(func(index int, item int) { visited++ })
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", visited))
}

// ==========================================================================
// Test: Collection — ForEachBreak
// ==========================================================================

func Test_Collection_ForEachBreak_Stops(t *testing.T) {
	tc := collectionForEachBreakStopsTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	visited := 0
	col.ForEachBreak(func(index int, item int) bool {
		visited++
		return item >= 3
	})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", visited))
}

func Test_Collection_ForEachBreak_VisitsAll(t *testing.T) {
	tc := collectionForEachBreakVisitsAllTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	visited := 0
	col.ForEachBreak(func(index int, item int) bool {
		visited++
		return false
	})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", visited))
}

// ==========================================================================
// Test: Collection — SortFunc
// ==========================================================================

func Test_Collection_SortFunc_Asc(t *testing.T) {
	tc := collectionSortFuncAscTestCase
	col := coregeneric.New.Collection.Int.Items(3, 1, 5, 2, 4)
	col.SortFunc(func(a, b int) bool { return a < b })
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_SortFunc_Desc(t *testing.T) {
	tc := collectionSortFuncDescTestCase
	col := coregeneric.New.Collection.Int.Items(3, 1, 5, 2, 4)
	col.SortFunc(func(a, b int) bool { return a > b })
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_SortFunc_Single(t *testing.T) {
	tc := collectionSortFuncSingleTestCase
	col := coregeneric.New.Collection.Int.Items(42)
	col.SortFunc(func(a, b int) bool { return a < b })
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

// ==========================================================================
// Test: Collection — AddIfMany
// ==========================================================================

func Test_Collection_AddIfMany_True(t *testing.T) {
	tc := collectionAddIfManyTrueTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(true, 10, 20, 30)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_AddIfMany_False(t *testing.T) {
	tc := collectionAddIfManyFalseTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(false, 10, 20, 30)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.Length()))
}

// ==========================================================================
// Test: Collection — AddFunc
// ==========================================================================

func Test_Collection_AddFunc(t *testing.T) {
	tc := collectionAddFuncTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddFunc(func() int { return 42 })
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
	)
}

// ==========================================================================
// Test: Collection — AddCollections
// ==========================================================================

func Test_Collection_AddCollections_Merge(t *testing.T) {
	tc := collectionAddCollectionsMergeTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	c2 := coregeneric.New.Collection.Int.Items(4, 5)
	c3 := coregeneric.New.Collection.Int.Items(6)
	col.AddCollections(c2, c3)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_AddCollections_WithNil(t *testing.T) {
	tc := collectionAddCollectionsNilTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	empty := coregeneric.EmptyCollection[int]()
	col.AddCollections(empty)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

// ==========================================================================
// Test: Collection — Clone edge cases
// ==========================================================================

func Test_Collection_Clone_Empty(t *testing.T) {
	tc := collectionCloneEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	cloned := col.Clone()
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", cloned.Length()),
		fmt.Sprintf("%v", cloned.IsEmpty()),
	)
}

// ==========================================================================
// Test: Collection — Skip/Take boundary
// ==========================================================================

func Test_Collection_SkipAll(t *testing.T) {
	tc := collectionSkipAllTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	skipped := col.Skip(10)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(skipped)))
}

func Test_Collection_TakeMore(t *testing.T) {
	tc := collectionTakeMoreTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	taken := col.Take(100)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(taken)))
}

func Test_Collection_SkipZeroTakeZero(t *testing.T) {
	tc := collectionSkipZeroTakeZeroTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	skipZero := col.Skip(0)
	takeZero := col.Take(0)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", len(skipZero)),
		fmt.Sprintf("%d", len(takeZero)),
	)
}

// ==========================================================================
// Test: Collection — Filter edge cases
// ==========================================================================

func Test_Collection_Filter_NoMatch(t *testing.T) {
	tc := collectionFilterNoMatchTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	filtered := col.Filter(func(item int) bool { return item > 100 })
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", filtered.Length()),
		fmt.Sprintf("%v", filtered.IsEmpty()),
	)
}

func Test_Collection_Filter_AllMatch(t *testing.T) {
	tc := collectionFilterAllMatchTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	filtered := col.Filter(func(item int) bool { return item > 0 })
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", filtered.Length()))
}

func Test_Collection_Filter_Empty(t *testing.T) {
	tc := collectionFilterEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	filtered := col.Filter(func(item int) bool { return true })
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", filtered.Length()),
		fmt.Sprintf("%v", filtered.IsEmpty()),
	)
}

// ==========================================================================
// Test: Collection — CountFunc edge cases
// ==========================================================================

func Test_Collection_CountFunc_NoMatch(t *testing.T) {
	tc := collectionCountFuncNoMatchTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	count := col.CountFunc(func(item int) bool { return item > 100 })
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", count))
}

func Test_Collection_CountFunc_Empty(t *testing.T) {
	tc := collectionCountFuncEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	count := col.CountFunc(func(item int) bool { return true })
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", count))
}

// ==========================================================================
// Test: Collection — String output
// ==========================================================================

func Test_Collection_String_Populated(t *testing.T) {
	tc := collectionStringPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	tc.ShouldBeEqual(t, 0, col.String())
}

func Test_Collection_String_Empty(t *testing.T) {
	tc := collectionStringEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	tc.ShouldBeEqual(t, 0, col.String())
}

// ==========================================================================
// Test: Collection — Lock variants
// ==========================================================================

func Test_Collection_Lock_Variants(t *testing.T) {
	tc := collectionLockVariantsTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddLock(1)
	col.AddsLock(2, 3)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.LengthLock()),
		fmt.Sprintf("%v", col.IsEmptyLock()),
		fmt.Sprintf("%d", col.Length()),
	)
}

// ==========================================================================
// Test: Collection — Metadata methods
// ==========================================================================

func Test_Collection_Metadata_Populated(t *testing.T) {
	tc := collectionMetadataPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", col.HasAnyItem()),
		fmt.Sprintf("%v", col.HasItems()),
		fmt.Sprintf("%v", col.HasIndex(2)),
		fmt.Sprintf("%v", col.HasIndex(5)),
		fmt.Sprintf("%d", col.LastIndex()),
		fmt.Sprintf("%d", col.Count()),
	)
}

func Test_Collection_Metadata_Empty(t *testing.T) {
	tc := collectionMetadataEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", col.HasAnyItem()),
		fmt.Sprintf("%v", col.HasItems()),
		fmt.Sprintf("%v", col.HasIndex(0)),
		fmt.Sprintf("%d", col.LastIndex()),
		fmt.Sprintf("%d", col.Count()),
	)
}

// ==========================================================================
// Test: Collection — RemoveAt single item
// ==========================================================================

func Test_Collection_RemoveAt_Single(t *testing.T) {
	tc := collectionRemoveAtSingleTestCase
	col := coregeneric.New.Collection.Int.Items(42)
	removed := col.RemoveAt(0)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", removed),
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%v", col.IsEmpty()),
	)
}

// ==========================================================================
// Test: Collection — AddCollection with empty
// ==========================================================================

func Test_Collection_AddCollection_Empty(t *testing.T) {
	tc := collectionAddCollectionEmptyTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	empty := coregeneric.EmptyCollection[int]()
	col.AddCollection(empty)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.Length()))
}

// ==========================================================================
// Test: Hashmap — IsEquals
// ==========================================================================

func Test_Hashmap_IsEquals_SameKeys(t *testing.T) {
	tc := hashmapIsEqualsSameKeysTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)
	hm1.Set("b", 2)
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm2.Set("a", 99)
	hm2.Set("b", 100)

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DiffKeys(t *testing.T) {
	tc := hashmapIsEqualsDiffKeysTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)
	hm1.Set("b", 2)
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm2.Set("x", 1)
	hm2.Set("y", 2)

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DiffLength(t *testing.T) {
	tc := hashmapIsEqualsDiffLengthTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm2.Set("a", 1)
	hm2.Set("b", 2)

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_BothNil(t *testing.T) {
	tc := hashmapIsEqualsBothNilTestCase
	var hm1 *coregeneric.Hashmap[string, int]
	var hm2 *coregeneric.Hashmap[string, int]

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_NilVsNonNil(t *testing.T) {
	tc := hashmapIsEqualsNilVsNonNilTestCase
	var hm1 *coregeneric.Hashmap[string, int]
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)

	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_SamePtr(t *testing.T) {
	tc := hashmapIsEqualsSamePtrTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)

	actual := args.Map{"isEquals": hm1.IsEquals(hm1)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — CollectionLenCap
// ==========================================================================

func Test_Collection_LenCap(t *testing.T) {
	tc := collectionLenCapTestCase
	col := coregeneric.CollectionLenCap[int](3, 10)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.Capacity()),
		fmt.Sprintf("%d", col.First()),
	)
}
