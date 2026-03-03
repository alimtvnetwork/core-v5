package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================================================
// Test: EmptyLinkedList
// ==========================================================================

func Test_LinkedList_Empty(t *testing.T) {
	tc := linkedListEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{
		fmt.Sprintf("%v", ll.IsEmpty()),
		fmt.Sprintf("%v", ll.Length()),
		fmt.Sprintf("%v", ll.HasItems()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: LinkedListFrom
// ==========================================================================

func Test_LinkedList_FromSlice(t *testing.T) {
	tc := linkedListFromSliceTestCase
	ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})

	actLines := []string{
		fmt.Sprintf("%v", ll.Length()),
		ll.First(),
		ll.Last(),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_FromEmptySlice(t *testing.T) {
	tc := linkedListFromEmptySliceTestCase
	ll := coregeneric.LinkedListFrom([]int{})

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Add
// ==========================================================================

func Test_LinkedList_AddSingle(t *testing.T) {
	tc := linkedListAddSingleTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(42)

	actLines := []string{
		fmt.Sprintf("%v", ll.Length()),
		fmt.Sprintf("%v", ll.First()),
		fmt.Sprintf("%v", ll.Last()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_AddMultiple(t *testing.T) {
	tc := linkedListAddMultipleTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1).Add(2).Add(3)

	actLines := []string{
		fmt.Sprintf("%v", ll.First()),
		fmt.Sprintf("%v", ll.Last()),
		fmt.Sprintf("%v", ll.Length()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddFront
// ==========================================================================

func Test_LinkedList_AddFrontPrepends(t *testing.T) {
	tc := linkedListAddFrontPrependsTestCase
	ll := coregeneric.LinkedListFrom([]int{2, 3})
	ll.AddFront(1)

	actLines := []string{
		fmt.Sprintf("%v", ll.First()),
		fmt.Sprintf("%v", ll.Last()),
		fmt.Sprintf("%v", ll.Length()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_AddFrontEmpty(t *testing.T) {
	tc := linkedListAddFrontEmptyTestCase
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddFront("first")

	actLines := []string{
		ll.First(),
		ll.Last(),
		fmt.Sprintf("%v", ll.Length()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Adds
// ==========================================================================

func Test_LinkedList_Adds(t *testing.T) {
	tc := linkedListAddsTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Adds(1, 2, 3)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddSlice
// ==========================================================================

func Test_LinkedList_AddSlice(t *testing.T) {
	tc := linkedListAddSliceTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddSlice([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddIf
// ==========================================================================

func Test_LinkedList_AddIfTrue(t *testing.T) {
	tc := linkedListAddIfTrueTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddIf(true, 5)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_AddIfFalse(t *testing.T) {
	tc := linkedListAddIfFalseTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddIf(false, 5)

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddsIf
// ==========================================================================

func Test_LinkedList_AddsIf(t *testing.T) {
	tc := linkedListAddsIfFalseTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(false, 1, 2, 3)

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AddFunc
// ==========================================================================

func Test_LinkedList_AddFunc(t *testing.T) {
	tc := linkedListAddFuncTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddFunc(func() int { return 99 })

	actLines := []string{fmt.Sprintf("%v", ll.First())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Push
// ==========================================================================

func Test_LinkedList_Push(t *testing.T) {
	tc := linkedListPushTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Push(1)
	ll.PushBack(2)
	ll.PushFront(0)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: FirstOrDefault
// ==========================================================================

func Test_LinkedList_FirstOrDefaultEmpty(t *testing.T) {
	tc := linkedListFirstDefaultEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.FirstOrDefault())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_FirstOrDefaultNonEmpty(t *testing.T) {
	tc := linkedListFirstDefaultNonEmptyTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.FirstOrDefault())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: LastOrDefault
// ==========================================================================

func Test_LinkedList_LastOrDefaultEmpty(t *testing.T) {
	tc := linkedListLastDefaultEmptyTestCase
	ll := coregeneric.EmptyLinkedList[string]()

	actLines := []string{ll.LastOrDefault()}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_LastOrDefaultNonEmpty(t *testing.T) {
	tc := linkedListLastDefaultNonEmptyTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.LastOrDefault())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Items
// ==========================================================================

func Test_LinkedList_ItemsAll(t *testing.T) {
	tc := linkedListItemsAllTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	actLines := []string{fmt.Sprintf("%v", len(ll.Items()))}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_ItemsEmpty(t *testing.T) {
	tc := linkedListItemsEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", len(ll.Items()))}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Collection
// ==========================================================================

func Test_LinkedList_Collection(t *testing.T) {
	tc := linkedListCollectionTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{fmt.Sprintf("%v", ll.Collection().Length())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: String
// ==========================================================================

func Test_LinkedList_String(t *testing.T) {
	tc := linkedListStringTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	actLines := []string{ll.String()}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: IndexAt
// ==========================================================================

func Test_LinkedList_IndexAt_Valid(t *testing.T) {
	tc := linkedListIndexAtValidTestCase
	ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})
	node := ll.IndexAt(1)

	actLines := []string{
		fmt.Sprintf("%v", node != nil),
		node.Element,
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IndexAt_First(t *testing.T) {
	tc := linkedListIndexAtFirstTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(0).Element)}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IndexAt_Last(t *testing.T) {
	tc := linkedListIndexAtLastTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(2).Element)}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IndexAt_OutOfBounds(t *testing.T) {
	tc := linkedListIndexAtOutOfBoundsTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{
		fmt.Sprintf("%v", ll.IndexAt(5) == nil),
		fmt.Sprintf("%v", ll.IndexAt(-1) == nil),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IndexAt_Empty(t *testing.T) {
	tc := linkedListIndexAtEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(0) == nil)}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: ForEach
// ==========================================================================

func Test_LinkedList_ForEachVisitsAll(t *testing.T) {
	tc := linkedListForEachVisitsAllTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	sum := 0
	ll.ForEach(func(_ int, item int) { sum += item })

	actLines := []string{fmt.Sprintf("%v", sum)}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_ForEachEmpty(t *testing.T) {
	tc := linkedListForEachEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	called := false
	ll.ForEach(func(_ int, _ int) { called = true })

	actLines := []string{fmt.Sprintf("%v", called)}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: ForEachBreak
// ==========================================================================

func Test_LinkedList_ForEachBreakStopsEarly(t *testing.T) {
	tc := linkedListForEachBreakStopsEarlyTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3, 4, 5})
	count := 0
	ll.ForEachBreak(func(_ int, item int) bool { count++; return item == 3 })

	actLines := []string{fmt.Sprintf("%v", count)}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_ForEachBreakFirst(t *testing.T) {
	tc := linkedListForEachBreakFirstTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(_ int, _ int) bool { count++; return true })

	actLines := []string{fmt.Sprintf("%v", count)}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Head / Tail
// ==========================================================================

func Test_LinkedList_HeadTail(t *testing.T) {
	tc := linkedListHeadTailTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	actLines := []string{
		fmt.Sprintf("%v", ll.Head().Element),
		fmt.Sprintf("%v", ll.Tail().Element),
		fmt.Sprintf("%v", ll.Head().HasNext()),
		fmt.Sprintf("%v", ll.Tail().HasNext()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_NodeNext(t *testing.T) {
	tc := linkedListNodeNextTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	n := ll.Head()

	actLines := []string{fmt.Sprintf("%v", n.Element)}
	n = n.Next()
	actLines = append(actLines, fmt.Sprintf("%v", n.Element))
	n = n.Next()
	actLines = append(actLines, fmt.Sprintf("%v", n.Element), fmt.Sprintf("%v", n.HasNext()))

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Lock variants
// ==========================================================================

func Test_LinkedList_LengthLock(t *testing.T) {
	tc := linkedListLengthLockTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{fmt.Sprintf("%v", ll.LengthLock())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_IsEmptyLock(t *testing.T) {
	tc := linkedListIsEmptyLockTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.IsEmptyLock())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_AddLock(t *testing.T) {
	tc := linkedListAddLockTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(1)
	ll.AddLock(2)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: Nil receiver
// ==========================================================================

func Test_LinkedList_NilReceiver(t *testing.T) {
	tc := linkedListNilReceiverTestCase
	var ll *coregeneric.LinkedList[int]

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	tc.ShouldBeEqualFirst(t, actLines...)
}

// ==========================================================================
// Test: AppendNode
// ==========================================================================

func Test_LinkedList_AppendNodeAppends(t *testing.T) {
	tc := linkedListAppendNodeAppendsTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 3})

	actLines := []string{
		fmt.Sprintf("%v", ll.Length()),
		fmt.Sprintf("%v", ll.Last()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}

func Test_LinkedList_AppendNodeEmpty(t *testing.T) {
	tc := linkedListAppendNodeEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 99})

	actLines := []string{
		fmt.Sprintf("%v", ll.Length()),
		fmt.Sprintf("%v", ll.First()),
	}

	tc.ShouldBeEqualFirst(t, actLines...)
}
