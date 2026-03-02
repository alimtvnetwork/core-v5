package corestrtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
)

// ==========================================================================
// Test: Hashset.AddNonEmpty — regression for no-op bug
// ==========================================================================

func Test_Hashset_AddNonEmpty_Regression(t *testing.T) {
	// Case 0: non-empty adds
	{
		tc := hashsetAddNonEmptyTestCases[0]
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("hello")
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("hello")),
		)
	}

	// Case 1: empty string skipped
	{
		tc := hashsetAddNonEmptyTestCases[1]
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("")
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", hs.Length()))
	}

	// Case 2: chained
	{
		tc := hashsetAddNonEmptyTestCases[2]
		hs := corestr.New.Hashset.Empty()
		hs.AddNonEmpty("a").AddNonEmpty("b").AddNonEmpty("c")
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("a")),
			fmt.Sprintf("%v", hs.Has("b")),
			fmt.Sprintf("%v", hs.Has("c")),
		)
	}
}

// ==========================================================================
// Test: SimpleSlice.InsertAt — regression for not-persisting + no bounds
// ==========================================================================

func Test_SimpleSlice_InsertAt_Regression(t *testing.T) {
	// Case 0: middle insert
	{
		tc := simpleSliceInsertAtTestCases[0]
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.InsertAt(1, "X")
		items := ss.ListCopyPtrLock()
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", ss.Length()),
			(*items)[0], (*items)[1], (*items)[2], (*items)[3],
		)
	}

	// Case 1: prepend
	{
		tc := simpleSliceInsertAtTestCases[1]
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.InsertAt(0, "X")
		items := ss.ListCopyPtrLock()
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%d", ss.Length()),
			(*items)[0], (*items)[1], (*items)[2], (*items)[3],
		)
	}

	// Case 2: append at end
	{
		tc := simpleSliceInsertAtTestCases[2]
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.InsertAt(3, "X")
		items := ss.ListCopyPtrLock()
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%d", ss.Length()),
			(*items)[0], (*items)[1], (*items)[2], (*items)[3],
		)
	}

	// Case 3: negative index — no change
	{
		tc := simpleSliceInsertAtTestCases[3]
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.InsertAt(-1, "X")
		items := ss.ListCopyPtrLock()
		tc.ShouldBeEqual(t, 3,
			fmt.Sprintf("%d", ss.Length()),
			(*items)[0], (*items)[1], (*items)[2],
		)
	}

	// Case 4: out-of-bounds — no change
	{
		tc := simpleSliceInsertAtTestCases[4]
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.InsertAt(100, "X")
		items := ss.ListCopyPtrLock()
		tc.ShouldBeEqual(t, 4,
			fmt.Sprintf("%d", ss.Length()),
			(*items)[0], (*items)[1], (*items)[2],
		)
	}
}

// ==========================================================================
// Test: Collection.RemoveAt — regression for inverted guard
// ==========================================================================

func Test_Collection_RemoveAt_Regression(t *testing.T) {
	// Case 0: middle
	{
		tc := collectionRemoveAtTestCases[0]
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := col.RemoveAt(1)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
		)
	}

	// Case 1: first
	{
		tc := collectionRemoveAtTestCases[1]
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := col.RemoveAt(0)
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
			col.First(),
		)
	}

	// Case 2: last
	{
		tc := collectionRemoveAtTestCases[2]
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := col.RemoveAt(2)
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
			col.Last(),
		)
	}

	// Case 3: negative
	{
		tc := collectionRemoveAtTestCases[3]
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := col.RemoveAt(-1)
		tc.ShouldBeEqual(t, 3,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
		)
	}

	// Case 4: out-of-bounds
	{
		tc := collectionRemoveAtTestCases[4]
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := col.RemoveAt(100)
		tc.ShouldBeEqual(t, 4,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
		)
	}

	// Case 5: empty
	{
		tc := collectionRemoveAtTestCases[5]
		col := corestr.New.Collection.Empty()
		ok := col.RemoveAt(0)
		tc.ShouldBeEqual(t, 5,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
		)
	}
}

// ==========================================================================
// Test: Hashmap.IsEqualPtr — regression for inverted comparison
// ==========================================================================

func Test_Hashmap_IsEqualPtr_Regression(t *testing.T) {
	// Case 0: same keys same values
	{
		tc := hashmapIsEqualPtrTestCases[0]
		hm1 := corestr.New.Hashmap.Empty()
		hm1.Set("a", "1")
		hm1.Set("b", "2")
		hm2 := corestr.New.Hashmap.Empty()
		hm2.Set("a", "1")
		hm2.Set("b", "2")
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
	}

	// Case 1: same keys different values
	{
		tc := hashmapIsEqualPtrTestCases[1]
		hm1 := corestr.New.Hashmap.Empty()
		hm1.Set("a", "1")
		hm2 := corestr.New.Hashmap.Empty()
		hm2.Set("a", "DIFFERENT")
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
	}

	// Case 2: different keys
	{
		tc := hashmapIsEqualPtrTestCases[2]
		hm1 := corestr.New.Hashmap.Empty()
		hm1.Set("a", "1")
		hm2 := corestr.New.Hashmap.Empty()
		hm2.Set("z", "1")
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
	}

	// Case 3: both empty
	{
		tc := hashmapIsEqualPtrTestCases[3]
		hm1 := corestr.New.Hashmap.Empty()
		hm2 := corestr.New.Hashmap.Empty()
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
	}

	// Case 4: nil vs non-nil
	{
		tc := hashmapIsEqualPtrTestCases[4]
		var hm1 *corestr.Hashmap
		hm2 := corestr.New.Hashmap.Empty()
		tc.ShouldBeEqual(t, 4, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
	}
}

// ==========================================================================
// Test: Caching removal — IsEmpty/Length on fresh instances
// ==========================================================================

func Test_Caching_Removal_Regression(t *testing.T) {
	// Case 0: fresh Hashset
	{
		tc := cachingRemovalTestCases[0]
		hs := corestr.New.Hashset.Empty()
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%v", hs.IsEmpty()),
			fmt.Sprintf("%d", hs.Length()),
		)
	}

	// Case 1: Hashset after Add
	{
		tc := cachingRemovalTestCases[1]
		hs := corestr.New.Hashset.Empty()
		hs.Add("a").Add("b")
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", hs.IsEmpty()),
			fmt.Sprintf("%d", hs.Length()),
		)
	}

	// Case 2: fresh Hashmap
	{
		tc := cachingRemovalTestCases[2]
		hm := corestr.New.Hashmap.Empty()
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%v", hm.IsEmpty()),
			fmt.Sprintf("%d", hm.Length()),
		)
	}

	// Case 3: Hashmap after Set
	{
		tc := cachingRemovalTestCases[3]
		hm := corestr.New.Hashmap.Empty()
		hm.Set("x", "1")
		hm.Set("y", "2")
		tc.ShouldBeEqual(t, 3,
			fmt.Sprintf("%v", hm.IsEmpty()),
			fmt.Sprintf("%d", hm.Length()),
		)
	}
}

// ==========================================================================
// Test: SimpleSlice.Skip/Take — regression for bounds protection
// ==========================================================================

func Test_SimpleSlice_SkipTake_Regression(t *testing.T) {
	ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

	// Case 0: Skip beyond length
	{
		tc := simpleSliceSkipTakeTestCases[0]
		result := ss.Skip(100)
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(result)))
	}

	// Case 1: Take beyond length
	{
		tc := simpleSliceSkipTakeTestCases[1]
		result := ss.Take(100)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", len(result)))
	}

	// Case 2: Skip 0
	{
		tc := simpleSliceSkipTakeTestCases[2]
		result := ss.Skip(0)
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", len(result)))
	}

	// Case 3: Take 0
	{
		tc := simpleSliceSkipTakeTestCases[3]
		result := ss.Take(0)
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%d", len(result)))
	}
}

// ==========================================================================
// Test: HasIndex — regression for negative index guard
// ==========================================================================

func Test_HasIndex_Negative_Regression(t *testing.T) {
	// Case 0: SimpleSlice
	{
		tc := hasIndexNegativeTestCases[0]
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ss.HasIndex(-1)))
	}

	// Case 1: Collection
	{
		tc := hasIndexNegativeTestCases[1]
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", col.HasIndex(-1)))
	}
}

// ==========================================================================
// Test: Hashmap.Clear nil safety — regression for nil panic
// ==========================================================================

func Test_Hashmap_Clear_NilSafety_Regression(t *testing.T) {
	// Case 0: nil receiver
	{
		tc := hashmapClearNilTestCases[0]
		var hm *corestr.Hashmap
		result := hm.Clear()
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", result == nil))
	}

	// Case 1: populated hashmap clears to empty
	{
		tc := hashmapClearNilTestCases[1]
		hm := corestr.New.Hashmap.Empty()
		hm.Set("a", "1")
		hm.Set("b", "2")
		hm.Clear()
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%d", hm.Length()),
			fmt.Sprintf("%v", hm.IsEmpty()),
		)
	}

	// Case 2: chainability after Clear
	{
		tc := hashmapClearNilTestCases[2]
		hm := corestr.New.Hashmap.Empty()
		hm.Set("x", "old")
		hm.Clear().Set("y", "new")
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%d", hm.Length()),
			fmt.Sprintf("%d", len(hm.ValuesList())),
		)
	}
}

// ==========================================================================
// Test: Hashset.AddBool cache invalidation — regression for stale cache
// ==========================================================================

func Test_Hashset_AddBool_CacheInvalidation_Regression(t *testing.T) {
	// Case 0: new item invalidates cache
	{
		tc := hashsetAddBoolCacheTestCases[0]
		hs := corestr.New.Hashset.Empty()
		isExist := hs.AddBool("hello")
		// Force cache rebuild by calling Items
		items := hs.Items()
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%v", isExist),
			fmt.Sprintf("%d", len(items)),
			fmt.Sprintf("%v", hs.Has("hello")),
		)
	}

	// Case 1: existing item returns true, no length change
	{
		tc := hashsetAddBoolCacheTestCases[1]
		hs := corestr.New.Hashset.Empty()
		hs.Add("hello")
		isExist := hs.AddBool("hello")
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", isExist),
			fmt.Sprintf("%d", hs.Length()),
		)
	}

	// Case 2: multiple new items all reflected in Items()
	{
		tc := hashsetAddBoolCacheTestCases[2]
		hs := corestr.New.Hashset.Empty()
		hs.AddBool("a")
		hs.AddBool("b")
		hs.AddBool("c")
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("a")),
			fmt.Sprintf("%v", hs.Has("b")),
			fmt.Sprintf("%v", hs.Has("c")),
		)
	}
}

// ==========================================================================
// Test: Hashmap.AddOrUpdateCollection length mismatch — regression
// ==========================================================================

func Test_Hashmap_AddOrUpdateCollection_LengthMismatch_Regression(t *testing.T) {
	// Case 0: mismatched lengths
	{
		tc := hashmapAddOrUpdateCollectionMismatchTestCases[0]
		hm := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"k1", "k2", "k3"})
		values := corestr.New.Collection.Strings([]string{"v1", "v2"})
		hm.AddOrUpdateCollection(keys, values)
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hm.Length()))
	}

	// Case 1: equal lengths adds all
	{
		tc := hashmapAddOrUpdateCollectionMismatchTestCases[1]
		hm := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
		values := corestr.New.Collection.Strings([]string{"v1", "v2"})
		hm.AddOrUpdateCollection(keys, values)
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%d", hm.Length()),
			hm.Get("k1"),
			hm.Get("k2"),
		)
	}

	// Case 2: nil keys
	{
		tc := hashmapAddOrUpdateCollectionMismatchTestCases[2]
		hm := corestr.New.Hashmap.Empty()
		values := corestr.New.Collection.Strings([]string{"v1"})
		hm.AddOrUpdateCollection(nil, values)
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", hm.Length()))
	}

	// Case 3: empty keys
	{
		tc := hashmapAddOrUpdateCollectionMismatchTestCases[3]
		hm := corestr.New.Hashmap.Empty()
		keys := corestr.New.Collection.Empty()
		values := corestr.New.Collection.Strings([]string{"v1"})
		hm.AddOrUpdateCollection(keys, values)
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%d", hm.Length()))
	}
}
