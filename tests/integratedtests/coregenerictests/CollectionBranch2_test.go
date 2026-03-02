package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Collection — RemoveAt edge cases
// ==========================================================================

func Test_Collection_RemoveAt_Edge(t *testing.T) {
	// Case 0: remove middle (index 2 from [1,2,3,4,5])
	{
		tc := collectionRemoveAtEdgeTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		ok := col.RemoveAt(2)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		)
	}

	// Case 1: remove first (index 0)
	{
		tc := collectionRemoveAtEdgeTestCases[1]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		ok := col.RemoveAt(0)
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		)
	}

	// Case 2: remove last (index 4)
	{
		tc := collectionRemoveAtEdgeTestCases[2]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		ok := col.RemoveAt(4)
		tc.ShouldBeEqual(t, 2,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		)
	}

	// Case 3: negative index
	{
		tc := collectionRemoveAtEdgeTestCases[3]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		ok := col.RemoveAt(-1)
		tc.ShouldBeEqual(t, 3,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
		)
	}

	// Case 4: out-of-bounds index
	{
		tc := collectionRemoveAtEdgeTestCases[4]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		ok := col.RemoveAt(100)
		tc.ShouldBeEqual(t, 4,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
		)
	}

	// Case 5: empty collection
	{
		tc := collectionRemoveAtEdgeTestCases[5]
		col := coregeneric.EmptyCollection[int]()
		ok := col.RemoveAt(0)
		tc.ShouldBeEqual(t, 5,
			fmt.Sprintf("%v", ok),
			fmt.Sprintf("%d", col.Length()),
		)
	}
}

// ==========================================================================
// Test: Collection — Reverse
// ==========================================================================

func Test_Collection_Reverse(t *testing.T) {
	// Case 0: populated
	{
		tc := collectionReverseTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
		col.Reverse()
		items := col.Items()
		actLines := make([]string, len(items))
		for i, v := range items {
			actLines[i] = fmt.Sprintf("%d", v)
		}
		tc.ShouldBeEqual(t, 0, actLines...)
	}

	// Case 1: single element
	{
		tc := collectionReverseTestCases[1]
		col := coregeneric.New.Collection.Int.Items(42)
		col.Reverse()
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.First()))
	}

	// Case 2: empty
	{
		tc := collectionReverseTestCases[2]
		col := coregeneric.EmptyCollection[int]()
		col.Reverse()
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", col.Length()))
	}
}

// ==========================================================================
// Test: Collection — FirstOrDefault
// ==========================================================================

func Test_Collection_FirstOrDefault(t *testing.T) {
	// Case 0: populated
	{
		tc := collectionFirstOrDefaultTestCases[0]
		col := coregeneric.New.Collection.Int.Items(10, 20, 30)
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.FirstOrDefault()))
	}

	// Case 1: empty
	{
		tc := collectionFirstOrDefaultTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.FirstOrDefault()))
	}
}

// ==========================================================================
// Test: Collection — LastOrDefault
// ==========================================================================

func Test_Collection_LastOrDefault(t *testing.T) {
	// Case 0: populated
	{
		tc := collectionLastOrDefaultTestCases[0]
		col := coregeneric.New.Collection.Int.Items(10, 20, 30)
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.LastOrDefault()))
	}

	// Case 1: empty
	{
		tc := collectionLastOrDefaultTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.LastOrDefault()))
	}
}

// ==========================================================================
// Test: Collection — SafeAt
// ==========================================================================

func Test_Collection_SafeAt(t *testing.T) {
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Case 0: valid index
	{
		tc := collectionSafeAtTestCases[0]
		tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.SafeAt(1)))
	}

	// Case 1: negative index
	{
		tc := collectionSafeAtTestCases[1]
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.SafeAt(-1)))
	}

	// Case 2: out-of-bounds
	{
		tc := collectionSafeAtTestCases[2]
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", col.SafeAt(100)))
	}

	// Case 3: empty collection
	{
		tc := collectionSafeAtTestCases[3]
		empty := coregeneric.EmptyCollection[int]()
		tc.ShouldBeEqual(t, 3, fmt.Sprintf("%d", empty.SafeAt(0)))
	}
}

// ==========================================================================
// Test: Collection — ConcatNew
// ==========================================================================

func Test_Collection_ConcatNew(t *testing.T) {
	// Case 0: concat onto populated
	{
		tc := collectionConcatNewTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		result := col.ConcatNew(4, 5)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", result.Length()),
			fmt.Sprintf("%d", result.First()),
			fmt.Sprintf("%d", result.Last()),
			fmt.Sprintf("%d", col.Length()), // original unchanged
		)
	}

	// Case 1: concat onto empty
	{
		tc := collectionConcatNewTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		result := col.ConcatNew(10, 20)
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%d", result.Length()),
			fmt.Sprintf("%d", result.First()),
			fmt.Sprintf("%d", result.Last()),
		)
	}
}

// ==========================================================================
// Test: Collection — AddIf
// ==========================================================================

func Test_Collection_AddIf(t *testing.T) {
	// Case 0: true
	{
		tc := collectionAddIfTestCases[0]
		col := coregeneric.EmptyCollection[int]()
		col.AddIf(true, 42)
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
		)
	}

	// Case 1: false
	{
		tc := collectionAddIfTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		col.AddIf(false, 42)
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.Length()))
	}
}

// ==========================================================================
// Test: Collection — ForEachBreak on empty
// ==========================================================================

func Test_Collection_ForEachBreak_Empty(t *testing.T) {
	tc := collectionForEachBreakEmptyTestCases[0]
	col := coregeneric.EmptyCollection[int]()
	visited := 0

	col.ForEachBreak(func(index int, item int) bool {
		visited++
		return false
	})

	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", visited))
}

// ==========================================================================
// Test: Collection — AddSlice
// ==========================================================================

func Test_Collection_AddSlice(t *testing.T) {
	// Case 0: add slice
	{
		tc := collectionAddSliceTestCases[0]
		col := coregeneric.EmptyCollection[int]()
		col.AddSlice([]int{10, 20, 30})
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		)
	}

	// Case 1: empty slice
	{
		tc := collectionAddSliceTestCases[1]
		col := coregeneric.EmptyCollection[int]()
		col.AddSlice([]int{})
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", col.Length()))
	}
}

// ==========================================================================
// Test: Collection — Items / ItemsPtr
// ==========================================================================

func Test_Collection_Items(t *testing.T) {
	// Case 0: Items returns slice
	{
		tc := collectionItemsTestCases[0]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		items := col.Items()
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", len(items)),
			fmt.Sprintf("%d", items[0]),
		)
	}

	// Case 1: ItemsPtr non-nil
	{
		tc := collectionItemsTestCases[1]
		col := coregeneric.New.Collection.Int.Items(1, 2, 3)
		ptr := col.ItemsPtr()
		tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", ptr != nil))
	}
}

// ==========================================================================
// Test: Collection — RemoveAtLock
// ==========================================================================

func Test_Collection_RemoveAtLock(t *testing.T) {
	tc := collectionRemoveAtLockTestCases[0]
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)

	col.Lock()
	ok := col.RemoveAt(1)
	col.Unlock()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
	)
}
