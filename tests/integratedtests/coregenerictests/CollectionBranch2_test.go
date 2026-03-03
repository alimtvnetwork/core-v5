package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Collection — RemoveAt edge cases
// ==========================================================================

func Test_Collection_RemoveAt_Middle(t *testing.T) {
	tc := collectionRemoveAtMiddleTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(2)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_RemoveAt_First(t *testing.T) {
	tc := collectionRemoveAtFirstTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(0)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_RemoveAt_Last(t *testing.T) {
	tc := collectionRemoveAtLastTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(4)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_RemoveAt_Negative(t *testing.T) {
	tc := collectionRemoveAtNegativeTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(-1)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
	)
}

func Test_Collection_RemoveAt_OutOfBounds(t *testing.T) {
	tc := collectionRemoveAtOutOfBoundsTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(100)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
	)
}

func Test_Collection_RemoveAt_Empty(t *testing.T) {
	tc := collectionRemoveAtEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	ok := col.RemoveAt(0)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
	)
}

// ==========================================================================
// Test: Collection — Reverse
// ==========================================================================

func Test_Collection_Reverse_Populated(t *testing.T) {
	tc := collectionReversePopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	col.Reverse()
	items := col.Items()
	actLines := make([]string, len(items))
	for i, v := range items {
		actLines[i] = fmt.Sprintf("%d", v)
	}
	tc.ShouldBeEqual(t, 0, actLines...)
}

func Test_Collection_Reverse_Single(t *testing.T) {
	tc := collectionReverseSingleTestCase
	col := coregeneric.New.Collection.Int.Items(42)
	col.Reverse()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.First()))
}

func Test_Collection_Reverse_Empty(t *testing.T) {
	tc := collectionReverseEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	col.Reverse()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.Length()))
}

// ==========================================================================
// Test: Collection — FirstOrDefault
// ==========================================================================

func Test_Collection_FirstOrDefault_Populated(t *testing.T) {
	tc := collectionFirstOrDefaultPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.FirstOrDefault()))
}

func Test_Collection_FirstOrDefault_Empty(t *testing.T) {
	tc := collectionFirstOrDefaultEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.FirstOrDefault()))
}

// ==========================================================================
// Test: Collection — LastOrDefault
// ==========================================================================

func Test_Collection_LastOrDefault_Populated(t *testing.T) {
	tc := collectionLastOrDefaultPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.LastOrDefault()))
}

func Test_Collection_LastOrDefault_Empty(t *testing.T) {
	tc := collectionLastOrDefaultEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.LastOrDefault()))
}

// ==========================================================================
// Test: Collection — SafeAt
// ==========================================================================

func Test_Collection_SafeAt_Valid(t *testing.T) {
	tc := collectionSafeAtValidTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.SafeAt(1)))
}

func Test_Collection_SafeAt_Negative(t *testing.T) {
	tc := collectionSafeAtNegativeTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.SafeAt(-1)))
}

func Test_Collection_SafeAt_OutOfBounds(t *testing.T) {
	tc := collectionSafeAtOutOfBoundsTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.SafeAt(100)))
}

func Test_Collection_SafeAt_Empty(t *testing.T) {
	tc := collectionSafeAtEmptyTestCase
	empty := coregeneric.EmptyCollection[int]()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", empty.SafeAt(0)))
}

// ==========================================================================
// Test: Collection — ConcatNew
// ==========================================================================

func Test_Collection_ConcatNew_Populated(t *testing.T) {
	tc := collectionConcatNewPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	result := col.ConcatNew(4, 5)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", result.Length()),
		fmt.Sprintf("%d", result.First()),
		fmt.Sprintf("%d", result.Last()),
		fmt.Sprintf("%d", col.Length()),
	)
}

func Test_Collection_ConcatNew_Empty(t *testing.T) {
	tc := collectionConcatNewEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	result := col.ConcatNew(10, 20)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", result.Length()),
		fmt.Sprintf("%d", result.First()),
		fmt.Sprintf("%d", result.Last()),
	)
}

// ==========================================================================
// Test: Collection — AddIf
// ==========================================================================

func Test_Collection_AddIf_True(t *testing.T) {
	tc := collectionAddIfTrueTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(true, 42)
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
	)
}

func Test_Collection_AddIf_False(t *testing.T) {
	tc := collectionAddIfFalseTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(false, 42)
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.Length()))
}

// ==========================================================================
// Test: Collection — ForEachBreak on empty
// ==========================================================================

func Test_Collection_ForEachBreak_Empty(t *testing.T) {
	tc := collectionForEachBreakEmptyTestCase
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

func Test_Collection_AddSlice_Populated(t *testing.T) {
	tc := collectionAddSlicePopulatedTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddSlice([]int{10, 20, 30})
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", col.Length()),
		fmt.Sprintf("%d", col.First()),
		fmt.Sprintf("%d", col.Last()),
	)
}

func Test_Collection_AddSlice_Empty(t *testing.T) {
	tc := collectionAddSliceEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddSlice([]int{})
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", col.Length()))
}

// ==========================================================================
// Test: Collection — Items / ItemsPtr
// ==========================================================================

func Test_Collection_Items_Slice(t *testing.T) {
	tc := collectionItemsSliceTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	items := col.Items()
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", len(items)),
		fmt.Sprintf("%d", items[0]),
	)
}

func Test_Collection_ItemsPtr(t *testing.T) {
	tc := collectionItemsPtrTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	ptr := col.ItemsPtr()
	tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ptr != nil))
}

// ==========================================================================
// Test: Collection — RemoveAtLock
// ==========================================================================

func Test_Collection_RemoveAtLock(t *testing.T) {
	tc := collectionRemoveAtLockTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	col.Lock()
	ok := col.RemoveAt(1)
	col.Unlock()
	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%v", ok),
		fmt.Sprintf("%d", col.Length()),
	)
}
