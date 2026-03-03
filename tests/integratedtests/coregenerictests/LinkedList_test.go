package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================================================
// Test: EmptyLinkedList
// ==========================================================================

func Test_LinkedList_Empty(t *testing.T) {
	tc := linkedListEmptyTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{
		fmt.Sprintf("%v", ll.IsEmpty()),
		fmt.Sprintf("%v", ll.Length()),
		fmt.Sprintf("%v", ll.HasItems()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: LinkedListFrom
// ==========================================================================

func Test_LinkedList_From(t *testing.T) {
	// Case 0: from slice
	{
		tc := linkedListFromTestCases[0]
		ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})

		actLines := []string{
			fmt.Sprintf("%v", ll.Length()),
			ll.First(),
			ll.Last(),
		}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty slice
	{
		tc := linkedListFromTestCases[1]
		ll := coregeneric.LinkedListFrom([]int{})

		actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Add
// ==========================================================================

func Test_LinkedList_AddSingle(t *testing.T) {
	tc := linkedListAddSingleTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(42)

	actLines := []string{
		fmt.Sprintf("%v", ll.Length()),
		fmt.Sprintf("%v", ll.First()),
		fmt.Sprintf("%v", ll.Last()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_AddMultiple(t *testing.T) {
	tc := linkedListAddMultipleTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1).Add(2).Add(3)

	actLines := []string{
		fmt.Sprintf("%v", ll.First()),
		fmt.Sprintf("%v", ll.Last()),
		fmt.Sprintf("%v", ll.Length()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: AddFront
// ==========================================================================

func Test_LinkedList_AddFront(t *testing.T) {
	// Case 0: prepends
	{
		tc := linkedListAddFrontTestCases[0]
		ll := coregeneric.LinkedListFrom([]int{2, 3})
		ll.AddFront(1)

		actLines := []string{
			fmt.Sprintf("%v", ll.First()),
			fmt.Sprintf("%v", ll.Last()),
			fmt.Sprintf("%v", ll.Length()),
		}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty
	{
		tc := linkedListAddFrontTestCases[1]
		ll := coregeneric.EmptyLinkedList[string]()
		ll.AddFront("first")

		actLines := []string{
			ll.First(),
			ll.Last(),
			fmt.Sprintf("%v", ll.Length()),
		}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Adds
// ==========================================================================

func Test_LinkedList_Adds(t *testing.T) {
	tc := linkedListAddsTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Adds(1, 2, 3)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: AddSlice
// ==========================================================================

func Test_LinkedList_AddSlice(t *testing.T) {
	tc := linkedListAddSliceTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddSlice([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: AddIf
// ==========================================================================

func Test_LinkedList_AddIf(t *testing.T) {
	// Case 0: true adds
	{
		tc := linkedListAddIfTestCases[0]
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddIf(true, 5)

		actLines := []string{fmt.Sprintf("%v", ll.Length())}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: false skips
	{
		tc := linkedListAddIfTestCases[1]
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddIf(false, 5)

		actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: AddsIf
// ==========================================================================

func Test_LinkedList_AddsIf(t *testing.T) {
	tc := linkedListAddsIfTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(false, 1, 2, 3)

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: AddFunc
// ==========================================================================

func Test_LinkedList_AddFunc(t *testing.T) {
	tc := linkedListAddFuncTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddFunc(func() int { return 99 })

	actLines := []string{fmt.Sprintf("%v", ll.First())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Push
// ==========================================================================

func Test_LinkedList_Push(t *testing.T) {
	tc := linkedListPushTestCases[0]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Push(1)
	ll.PushBack(2)
	ll.PushFront(0)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: FirstOrDefault
// ==========================================================================

func Test_LinkedList_FirstOrDefault(t *testing.T) {
	// Case 0: empty returns zero
	{
		tc := linkedListFirstDefaultTestCases[0]
		ll := coregeneric.EmptyLinkedList[int]()

		actLines := []string{fmt.Sprintf("%v", ll.FirstOrDefault())}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: non-empty
	{
		tc := linkedListFirstDefaultTestCases[1]
		ll := coregeneric.LinkedListFrom([]int{10, 20})

		actLines := []string{fmt.Sprintf("%v", ll.FirstOrDefault())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: LastOrDefault
// ==========================================================================

func Test_LinkedList_LastOrDefault(t *testing.T) {
	// Case 0: empty returns zero
	{
		tc := linkedListLastDefaultTestCases[0]
		ll := coregeneric.EmptyLinkedList[string]()

		actLines := []string{ll.LastOrDefault()}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: non-empty
	{
		tc := linkedListLastDefaultTestCases[1]
		ll := coregeneric.LinkedListFrom([]int{10, 20})

		actLines := []string{fmt.Sprintf("%v", ll.LastOrDefault())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Items
// ==========================================================================

func Test_LinkedList_Items(t *testing.T) {
	// Case 0: returns all
	{
		tc := linkedListItemsTestCases[0]
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

		actLines := []string{fmt.Sprintf("%v", len(ll.Items()))}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty
	{
		tc := linkedListItemsTestCases[1]
		ll := coregeneric.EmptyLinkedList[int]()

		actLines := []string{fmt.Sprintf("%v", len(ll.Items()))}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Collection
// ==========================================================================

func Test_LinkedList_Collection(t *testing.T) {
	tc := linkedListCollectionTestCases[0]
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{fmt.Sprintf("%v", ll.Collection().Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: String
// ==========================================================================

func Test_LinkedList_String(t *testing.T) {
	tc := linkedListStringTestCases[0]
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	actLines := []string{ll.String()}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
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

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_IndexAt_First(t *testing.T) {
	tc := linkedListIndexAtFirstTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20})

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(0).Element)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_IndexAt_Last(t *testing.T) {
	tc := linkedListIndexAtLastTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(2).Element)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_IndexAt_OutOfBounds(t *testing.T) {
	tc := linkedListIndexAtOutOfBoundsTestCase
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{
		fmt.Sprintf("%v", ll.IndexAt(5) == nil),
		fmt.Sprintf("%v", ll.IndexAt(-1) == nil),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_IndexAt_Empty(t *testing.T) {
	tc := linkedListIndexAtEmptyTestCase
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.IndexAt(0) == nil)}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: ForEach
// ==========================================================================

func Test_LinkedList_ForEach(t *testing.T) {
	// Case 0: visits all
	{
		tc := linkedListForEachTestCases[0]
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
		sum := 0
		ll.ForEach(func(_ int, item int) { sum += item })

		actLines := []string{fmt.Sprintf("%v", sum)}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty noop
	{
		tc := linkedListForEachTestCases[1]
		ll := coregeneric.EmptyLinkedList[int]()
		called := false
		ll.ForEach(func(_ int, _ int) { called = true })

		actLines := []string{fmt.Sprintf("%v", called)}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: ForEachBreak
// ==========================================================================

func Test_LinkedList_ForEachBreak(t *testing.T) {
	// Case 0: stops early
	{
		tc := linkedListForEachBreakTestCases[0]
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3, 4, 5})
		count := 0
		ll.ForEachBreak(func(_ int, item int) bool { count++; return item == 3 })

		actLines := []string{fmt.Sprintf("%v", count)}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: first element
	{
		tc := linkedListForEachBreakTestCases[1]
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
		count := 0
		ll.ForEachBreak(func(_ int, _ int) bool { count++; return true })

		actLines := []string{fmt.Sprintf("%v", count)}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Head / Tail
// ==========================================================================

func Test_LinkedList_HeadTail(t *testing.T) {
	tc := linkedListHeadTailTestCases[0]
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})

	actLines := []string{
		fmt.Sprintf("%v", ll.Head().Element),
		fmt.Sprintf("%v", ll.Tail().Element),
		fmt.Sprintf("%v", ll.Head().HasNext()),
		fmt.Sprintf("%v", ll.Tail().HasNext()),
	}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_NodeNext(t *testing.T) {
	tc := linkedListNodeNextTestCases[0]
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	n := ll.Head()

	actLines := []string{fmt.Sprintf("%v", n.Element)}
	n = n.Next()
	actLines = append(actLines, fmt.Sprintf("%v", n.Element))
	n = n.Next()
	actLines = append(actLines, fmt.Sprintf("%v", n.Element), fmt.Sprintf("%v", n.HasNext()))

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Lock variants
// ==========================================================================

func Test_LinkedList_LengthLock(t *testing.T) {
	tc := linkedListLockTestCases[0]
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	actLines := []string{fmt.Sprintf("%v", ll.LengthLock())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_IsEmptyLock(t *testing.T) {
	tc := linkedListLockTestCases[1]
	ll := coregeneric.EmptyLinkedList[int]()

	actLines := []string{fmt.Sprintf("%v", ll.IsEmptyLock())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_LinkedList_AddLock(t *testing.T) {
	tc := linkedListLockTestCases[2]
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(1)
	ll.AddLock(2)

	actLines := []string{fmt.Sprintf("%v", ll.Length())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: Nil receiver
// ==========================================================================

func Test_LinkedList_NilReceiver(t *testing.T) {
	tc := linkedListNilReceiverTestCases[0]
	var ll *coregeneric.LinkedList[int]

	actLines := []string{fmt.Sprintf("%v", ll.IsEmpty())}

	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================================================
// Test: AppendNode
// ==========================================================================

func Test_LinkedList_AppendNode(t *testing.T) {
	// Case 0: appends
	{
		tc := linkedListAppendNodeTestCases[0]
		ll := coregeneric.LinkedListFrom([]int{1, 2})
		ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 3})

		actLines := []string{
			fmt.Sprintf("%v", ll.Length()),
			fmt.Sprintf("%v", ll.Last()),
		}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: empty
	{
		tc := linkedListAppendNodeTestCases[1]
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AppendNode(&coregeneric.LinkedListNode[int]{Element: 99})

		actLines := []string{
			fmt.Sprintf("%v", ll.Length()),
			fmt.Sprintf("%v", ll.First()),
		}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}
