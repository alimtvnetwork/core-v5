package coregenerictests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: Collection — RemoveAt edge cases
// ==========================================================================

func Test_Collection_RemoveAt_Middle(t *testing.T) {
	tc := collectionRemoveAtMiddleTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(2)

	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
		"first":   col.First(),
		"last":    col.Last(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_First(t *testing.T) {
	tc := collectionRemoveAtFirstTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(0)

	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
		"first":   col.First(),
		"last":    col.Last(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_Last(t *testing.T) {
	tc := collectionRemoveAtLastTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(4)

	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
		"first":   col.First(),
		"last":    col.Last(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_Negative(t *testing.T) {
	tc := collectionRemoveAtNegativeTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(-1)

	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_OutOfBounds(t *testing.T) {
	tc := collectionRemoveAtOutOfBoundsTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(100)

	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_Empty(t *testing.T) {
	tc := collectionRemoveAtEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	ok := col.RemoveAt(0)

	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Reverse
// ==========================================================================

func Test_Collection_Reverse_Populated(t *testing.T) {
	tc := collectionReversePopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	col.Reverse()

	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
		"last":   col.Last(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_Reverse_Single(t *testing.T) {
	tc := collectionReverseSingleTestCase
	col := coregeneric.New.Collection.Int.Items(42)
	col.Reverse()

	actual := args.Map{"first": col.First()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_Reverse_Empty(t *testing.T) {
	tc := collectionReverseEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	col.Reverse()

	actual := args.Map{"length": col.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — FirstOrDefault
// ==========================================================================

func Test_Collection_FirstOrDefault_Populated(t *testing.T) {
	tc := collectionFirstOrDefaultPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	actual := args.Map{"result": col.FirstOrDefault()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_FirstOrDefault_Empty(t *testing.T) {
	tc := collectionFirstOrDefaultEmptyTestCase
	col := coregeneric.EmptyCollection[int]()

	actual := args.Map{"result": col.FirstOrDefault()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — LastOrDefault
// ==========================================================================

func Test_Collection_LastOrDefault_Populated(t *testing.T) {
	tc := collectionLastOrDefaultPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	actual := args.Map{"result": col.LastOrDefault()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_LastOrDefault_Empty(t *testing.T) {
	tc := collectionLastOrDefaultEmptyTestCase
	col := coregeneric.EmptyCollection[int]()

	actual := args.Map{"result": col.LastOrDefault()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — SafeAt
// ==========================================================================

func Test_Collection_SafeAt_Valid(t *testing.T) {
	tc := collectionSafeAtValidTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	actual := args.Map{"result": col.SafeAt(1)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SafeAt_Negative(t *testing.T) {
	tc := collectionSafeAtNegativeTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	actual := args.Map{"result": col.SafeAt(-1)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SafeAt_OutOfBounds(t *testing.T) {
	tc := collectionSafeAtOutOfBoundsTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	actual := args.Map{"result": col.SafeAt(100)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SafeAt_Empty(t *testing.T) {
	tc := collectionSafeAtEmptyTestCase
	empty := coregeneric.EmptyCollection[int]()

	actual := args.Map{"result": empty.SafeAt(0)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — ConcatNew
// ==========================================================================

func Test_Collection_ConcatNew_Populated(t *testing.T) {
	tc := collectionConcatNewPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	result := col.ConcatNew(4, 5)

	actual := args.Map{
		"resultLength": result.Length(),
		"resultFirst":  result.First(),
		"resultLast":   result.Last(),
		"origLength":   col.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_ConcatNew_Empty(t *testing.T) {
	tc := collectionConcatNewEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	result := col.ConcatNew(10, 20)

	actual := args.Map{
		"length": result.Length(),
		"first":  result.First(),
		"last":   result.Last(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddIf
// ==========================================================================

func Test_Collection_AddIf_True(t *testing.T) {
	tc := collectionAddIfTrueTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(true, 42)

	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_AddIf_False(t *testing.T) {
	tc := collectionAddIfFalseTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(false, 42)

	actual := args.Map{"length": col.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
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

	actual := args.Map{"visited": visited}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddSlice
// ==========================================================================

func Test_Collection_AddSlice_Populated(t *testing.T) {
	tc := collectionAddSlicePopulatedTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddSlice([]int{10, 20, 30})

	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
		"last":   col.Last(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_AddSlice_Empty(t *testing.T) {
	tc := collectionAddSliceEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddSlice([]int{})

	actual := args.Map{"length": col.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Items / ItemsPtr
// ==========================================================================

func Test_Collection_Items_Slice(t *testing.T) {
	tc := collectionItemsSliceTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	items := col.Items()

	actual := args.Map{
		"length": len(items),
		"first":  items[0],
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_ItemsPtr(t *testing.T) {
	tc := collectionItemsPtrTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	ptr := col.ItemsPtr()

	actual := args.Map{"isNotNil": ptr != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
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

	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}
