package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Collection — ForEach
// ==========================================================================

func Test_Collection_ForEach(t *testing.T) {
	// Case 0: visits all items
	{
		tc := collectionForEachTestCases[0]
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

	// Case 1: empty collection
	{
		tc := collectionForEachTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		visited := 0

		col.ForEach(func(index int, item int) { visited++ })

		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", visited))
	}
}

// ==========================================================================
// Test: Collection — ForEachBreak
// ==========================================================================

func Test_Collection_ForEachBreak(t *testing.T) {
	// Case 0: stops at first item > 3
	{
		tc := collectionForEachBreakTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		visited := 0

		col.ForEachBreak(func(index int, item int) bool {
			visited++
			return item >= 3
		})

		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", visited))
	}

	// Case 1: no break, visits all
	{
		tc := collectionForEachBreakTestCases[1]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		visited := 0

		col.ForEachBreak(func(index int, item int) bool {
			visited++
			return false
		})

		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", visited))
	}
}

// ==========================================================================
// Test: Collection — SortFunc
// ==========================================================================

func Test_Collection_SortFunc(t *testing.T) {
	// Case 0: ascending
	{
		tc := collectionSortFuncTestCases[0]
		col := coregeneric.New.Collection.Int.Items(3, 1, 5, 2, 4)
		col.SortFunc(func(a, b int) bool { return a < b })
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.First()), fmt.Sprintf("%d", col.Last()))
	}

	// Case 1: descending
	{
		tc := collectionSortFuncTestCases[1]
		col := coregeneric.New.Collection.Int.Items(3, 1, 5, 2, 4)
		col.SortFunc(func(a, b int) bool { return a > b })
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.First()), fmt.Sprintf("%d", col.Last()))
	}

	// Case 2: single element
	{
		tc := collectionSortFuncTestCases[2]
		col := coregeneric.New.Collection.Int.Items(42)
		col.SortFunc(func(a, b int) bool { return a < b })
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", col.First()), fmt.Sprintf("%d", col.Last()))
	}
}

// ==========================================================================
// Test: Collection — AddIfMany
// ==========================================================================

func Test_Collection_AddIfMany(t *testing.T) {
	// Case 0: true
	{
		tc := collectionAddIfManyTestCases[0]
		col := coregeneric.EmptyCollection[int]()
		col.AddIfMany(true, 10, 20, 30)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		)
	}

	// Case 1: false
	{
		tc := collectionAddIfManyTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		col.AddIfMany(false, 10, 20, 30)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.Length()))
	}
}

// ==========================================================================
// Test: Collection — AddFunc
// ==========================================================================

func Test_Collection_AddFunc(t *testing.T) {
	tc := collectionAddFuncTestCases[0]
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

func Test_Collection_AddCollections(t *testing.T) {
	// Case 0: merge multiple
	{
		tc := collectionAddCollectionsTestCases[0]
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

	// Case 1: with nil/empty collection
	{
		tc := collectionAddCollectionsTestCases[1]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		empty := coregeneric.EmptyCollection[int]()
		col.AddCollections(empty)
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		)
	}
}

// ==========================================================================
// Test: Collection — Clone edge cases
// ==========================================================================

func Test_Collection_Clone_Empty(t *testing.T) {
	tc := collectionCloneEdgeTestCases[0]
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

func Test_Collection_SkipTake_Boundary(t *testing.T) {
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)

	// Case 0: Skip all
	{
		tc := collectionSkipTakeBoundaryTestCases[0]
		skipped := col.Skip(10)
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(skipped)))
	}

	// Case 1: Take more than length
	{
		tc := collectionSkipTakeBoundaryTestCases[1]
		taken := col.Take(100)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", len(taken)))
	}

	// Case 2: Skip 0 / Take 0
	{
		tc := collectionSkipTakeBoundaryTestCases[2]
		skipZero := col.Skip(0)
		takeZero := col.Take(0)
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%d", len(skipZero)),
			fmt.Sprintf("%d", len(takeZero)),
		)
	}
}

// ==========================================================================
// Test: Collection — Filter edge cases
// ==========================================================================

func Test_Collection_Filter_Edge(t *testing.T) {
	// Case 0: no match
	{
		tc := collectionFilterEdgeTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		filtered := col.Filter(func(item int) bool { return item > 100 })
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", filtered.Length()),
			fmt.Sprintf("%v", filtered.IsEmpty()),
		)
	}

	// Case 1: all match
	{
		tc := collectionFilterEdgeTestCases[1]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		filtered := col.Filter(func(item int) bool { return item > 0 })
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", filtered.Length()))
	}

	// Case 2: empty collection
	{
		tc := collectionFilterEdgeTestCases[2]
		col := coregeneric.EmptyCollection[int]()
		filtered := col.Filter(func(item int) bool { return true })
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%d", filtered.Length()),
			fmt.Sprintf("%v", filtered.IsEmpty()),
		)
	}
}

// ==========================================================================
// Test: Collection — CountFunc edge cases
// ==========================================================================

func Test_Collection_CountFunc_Edge(t *testing.T) {
	// Case 0: no match
	{
		tc := collectionCountFuncEdgeTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		count := col.CountFunc(func(item int) bool { return item > 100 })
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", count))
	}

	// Case 1: empty
	{
		tc := collectionCountFuncEdgeTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		count := col.CountFunc(func(item int) bool { return true })
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", count))
	}
}

// ==========================================================================
// Test: Collection — String output
// ==========================================================================

func Test_Collection_String(t *testing.T) {
	// Case 0: populated
	{
		tc := collectionStringTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		tc.ShouldBeEqual(t, 0, col.String())
	}

	// Case 1: empty
	{
		tc := collectionStringTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		tc.ShouldBeEqual(t, 1, col.String())
	}
}

// ==========================================================================
// Test: Collection — Lock variants
// ==========================================================================

func Test_Collection_Lock_Variants(t *testing.T) {
	tc := collectionLockTestCases[0]
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

func Test_Collection_Metadata(t *testing.T) {
	// Case 0: populated
	{
		tc := collectionMetadataTestCases[0]
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

	// Case 1: empty
	{
		tc := collectionMetadataTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", col.HasAnyItem()),
			fmt.Sprintf("%v", col.HasItems()),
			fmt.Sprintf("%v", col.HasIndex(0)),
			fmt.Sprintf("%d", col.LastIndex()),
			fmt.Sprintf("%d", col.Count()),
		)
	}
}

// ==========================================================================
// Test: Collection — RemoveAt single item
// ==========================================================================

func Test_Collection_RemoveAt_Single(t *testing.T) {
	tc := collectionRemoveAtSingleTestCases[0]
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
	tc := collectionAddCollectionNilTestCases[0]
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	empty := coregeneric.EmptyCollection[int]()
	col.AddCollection(empty)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.Length()))
}

// ==========================================================================
// Test: Hashmap — IsEquals (updated with key-checking)
// ==========================================================================

func Test_Hashmap_IsEquals_Updated(t *testing.T) {
	// Case 0: same keys → true
	{
		tc := hashmapIsEqualsUpdatedTestCases[0]
		hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm1.Set("a", 1)
		hm1.Set("b", 2)
		hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm2.Set("a", 99)
		hm2.Set("b", 100)
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hm1.IsEquals(hm2)))
	}

	// Case 1: same length different keys → false
	{
		tc := hashmapIsEqualsUpdatedTestCases[1]
		hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm1.Set("a", 1)
		hm1.Set("b", 2)
		hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm2.Set("x", 1)
		hm2.Set("y", 2)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", hm1.IsEquals(hm2)))
	}

	// Case 2: different length → false
	{
		tc := hashmapIsEqualsUpdatedTestCases[2]
		hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm1.Set("a", 1)
		hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm2.Set("a", 1)
		hm2.Set("b", 2)
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", hm1.IsEquals(hm2)))
	}

	// Case 3: both nil → true
	{
		tc := hashmapIsEqualsUpdatedTestCases[3]
		var hm1 *coregeneric.Hashmap[string, int]
		var hm2 *coregeneric.Hashmap[string, int]
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", hm1.IsEquals(hm2)))
	}

	// Case 4: nil vs non-nil → false
	{
		tc := hashmapIsEqualsUpdatedTestCases[4]
		var hm1 *coregeneric.Hashmap[string, int]
		hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
		tc.ShouldBeEqual(t, 4, fmt.Sprintf("%v", hm1.IsEquals(hm2)))
	}

	// Case 5: same pointer → true
	{
		tc := hashmapIsEqualsUpdatedTestCases[5]
		hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm1.Set("a", 1)
		tc.ShouldBeEqual(t, 5, fmt.Sprintf("%v", hm1.IsEquals(hm1)))
	}
}

// ==========================================================================
// Test: Collection — CollectionLenCap
// ==========================================================================

func Test_Collection_LenCap(t *testing.T) {
	tc := collectionLenCapTestCases[0]
	col := coregeneric.CollectionLenCap[int](3, 10)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.Capacity()),
		fmt.Sprintf("%d", col.First()),
	)
}
