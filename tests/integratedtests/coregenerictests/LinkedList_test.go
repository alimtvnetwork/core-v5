package coregenerictests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// =============================================================================
// LinkedList — Constructors
// =============================================================================

func Test_LinkedList_Empty(t *testing.T) {
	convey.Convey("EmptyLinkedList creates empty list", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		convey.So(ll.IsEmpty(), should.BeTrue)
		convey.So(ll.Length(), should.Equal, 0)
		convey.So(ll.HasItems(), should.BeFalse)
	})
}

func Test_LinkedList_From(t *testing.T) {
	convey.Convey("LinkedListFrom creates list from slice", t, func() {
		ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})
		convey.So(ll.Length(), should.Equal, 3)
		convey.So(ll.First(), should.Equal, "a")
		convey.So(ll.Last(), should.Equal, "c")
	})
}

func Test_LinkedList_From_Empty(t *testing.T) {
	convey.Convey("LinkedListFrom with empty slice", t, func() {
		ll := coregeneric.LinkedListFrom([]int{})
		convey.So(ll.IsEmpty(), should.BeTrue)
	})
}

// =============================================================================
// LinkedList — Add / AddFront
// =============================================================================

func Test_LinkedList_Add_Single(t *testing.T) {
	convey.Convey("LinkedList.Add to empty list sets head and tail", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.Add(42)
		convey.So(ll.Length(), should.Equal, 1)
		convey.So(ll.First(), should.Equal, 42)
		convey.So(ll.Last(), should.Equal, 42)
	})
}

func Test_LinkedList_Add_Multiple(t *testing.T) {
	convey.Convey("LinkedList.Add appends to back", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.Add(1).Add(2).Add(3)
		convey.So(ll.First(), should.Equal, 1)
		convey.So(ll.Last(), should.Equal, 3)
		convey.So(ll.Length(), should.Equal, 3)
	})
}

func Test_LinkedList_AddFront(t *testing.T) {
	convey.Convey("LinkedList.AddFront prepends to front", t, func() {
		ll := coregeneric.LinkedListFrom([]int{2, 3})
		ll.AddFront(1)
		convey.So(ll.First(), should.Equal, 1)
		convey.So(ll.Last(), should.Equal, 3)
		convey.So(ll.Length(), should.Equal, 3)
	})
}

func Test_LinkedList_AddFront_Empty(t *testing.T) {
	convey.Convey("LinkedList.AddFront on empty list", t, func() {
		ll := coregeneric.EmptyLinkedList[string]()
		ll.AddFront("first")
		convey.So(ll.First(), should.Equal, "first")
		convey.So(ll.Last(), should.Equal, "first")
		convey.So(ll.Length(), should.Equal, 1)
	})
}

func Test_LinkedList_Adds(t *testing.T) {
	convey.Convey("LinkedList.Adds appends multiple", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.Adds(1, 2, 3)
		convey.So(ll.Length(), should.Equal, 3)
		convey.So(ll.Items(), should.Resemble, []int{1, 2, 3})
	})
}

func Test_LinkedList_AddSlice(t *testing.T) {
	convey.Convey("LinkedList.AddSlice appends slice", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddSlice([]int{10, 20})
		convey.So(ll.Length(), should.Equal, 2)
	})
}

func Test_LinkedList_AddIf_True(t *testing.T) {
	convey.Convey("LinkedList.AddIf adds when true", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddIf(true, 5)
		convey.So(ll.Length(), should.Equal, 1)
	})
}

func Test_LinkedList_AddIf_False(t *testing.T) {
	convey.Convey("LinkedList.AddIf skips when false", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddIf(false, 5)
		convey.So(ll.IsEmpty(), should.BeTrue)
	})
}

func Test_LinkedList_AddsIf_False(t *testing.T) {
	convey.Convey("LinkedList.AddsIf skips when false", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddsIf(false, 1, 2, 3)
		convey.So(ll.IsEmpty(), should.BeTrue)
	})
}

func Test_LinkedList_AddFunc(t *testing.T) {
	convey.Convey("LinkedList.AddFunc adds function result", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddFunc(func() int { return 99 })
		convey.So(ll.First(), should.Equal, 99)
	})
}

func Test_LinkedList_PushBack_PushFront_Push(t *testing.T) {
	convey.Convey("Push aliases work correctly", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.Push(1)
		ll.PushBack(2)
		ll.PushFront(0)
		convey.So(ll.Items(), should.Resemble, []int{0, 1, 2})
	})
}

// =============================================================================
// LinkedList — FirstOrDefault / LastOrDefault
// =============================================================================

func Test_LinkedList_FirstOrDefault_Empty(t *testing.T) {
	convey.Convey("LinkedList.FirstOrDefault returns zero on empty", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		convey.So(ll.FirstOrDefault(), should.Equal, 0)
	})
}

func Test_LinkedList_LastOrDefault_Empty(t *testing.T) {
	convey.Convey("LinkedList.LastOrDefault returns zero on empty", t, func() {
		ll := coregeneric.EmptyLinkedList[string]()
		convey.So(ll.LastOrDefault(), should.BeEmpty)
	})
}

func Test_LinkedList_FirstOrDefault_NonEmpty(t *testing.T) {
	convey.Convey("LinkedList.FirstOrDefault returns first element", t, func() {
		ll := coregeneric.LinkedListFrom([]int{10, 20})
		convey.So(ll.FirstOrDefault(), should.Equal, 10)
	})
}

func Test_LinkedList_LastOrDefault_NonEmpty(t *testing.T) {
	convey.Convey("LinkedList.LastOrDefault returns last element", t, func() {
		ll := coregeneric.LinkedListFrom([]int{10, 20})
		convey.So(ll.LastOrDefault(), should.Equal, 20)
	})
}

// =============================================================================
// LinkedList — Items / Collection / String
// =============================================================================

func Test_LinkedList_Items(t *testing.T) {
	convey.Convey("LinkedList.Items returns all elements", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
		convey.So(ll.Items(), should.Resemble, []int{1, 2, 3})
	})
}

func Test_LinkedList_Items_Empty(t *testing.T) {
	convey.Convey("LinkedList.Items returns empty slice on empty list", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		convey.So(ll.Items(), should.Resemble, []int{})
	})
}

func Test_LinkedList_Collection(t *testing.T) {
	convey.Convey("LinkedList.Collection converts to Collection", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2})
		col := ll.Collection()
		convey.So(col.Length(), should.Equal, 2)
	})
}

func Test_LinkedList_String(t *testing.T) {
	convey.Convey("LinkedList.String returns string representation", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
		convey.So(ll.String(), should.Equal, "[1 2 3]")
	})
}

// =============================================================================
// LinkedList — IndexAt
// =============================================================================

func Test_LinkedList_IndexAt_Valid(t *testing.T) {
	convey.Convey("LinkedList.IndexAt returns correct node", t, func() {
		ll := coregeneric.LinkedListFrom([]string{"a", "b", "c"})
		node := ll.IndexAt(1)
		convey.So(node, should.NotBeNil)
		convey.So(node.Element, should.Equal, "b")
	})
}

func Test_LinkedList_IndexAt_First(t *testing.T) {
	convey.Convey("LinkedList.IndexAt(0) returns head", t, func() {
		ll := coregeneric.LinkedListFrom([]int{10, 20})
		convey.So(ll.IndexAt(0).Element, should.Equal, 10)
	})
}

func Test_LinkedList_IndexAt_Last(t *testing.T) {
	convey.Convey("LinkedList.IndexAt(last) returns tail", t, func() {
		ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
		convey.So(ll.IndexAt(2).Element, should.Equal, 30)
	})
}

func Test_LinkedList_IndexAt_OutOfBounds(t *testing.T) {
	convey.Convey("LinkedList.IndexAt returns nil for out of bounds", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2})
		convey.So(ll.IndexAt(5), should.BeNil)
		convey.So(ll.IndexAt(-1), should.BeNil)
	})
}

func Test_LinkedList_IndexAt_Empty(t *testing.T) {
	convey.Convey("LinkedList.IndexAt returns nil on empty list", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		convey.So(ll.IndexAt(0), should.BeNil)
	})
}

// =============================================================================
// LinkedList — ForEach / ForEachBreak
// =============================================================================

func Test_LinkedList_ForEach(t *testing.T) {
	convey.Convey("LinkedList.ForEach visits all elements", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
		sum := 0
		ll.ForEach(func(index int, item int) {
			sum += item
		})
		convey.So(sum, should.Equal, 6)
	})
}

func Test_LinkedList_ForEach_Empty(t *testing.T) {
	convey.Convey("LinkedList.ForEach does nothing on empty", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		called := false
		ll.ForEach(func(index int, item int) { called = true })
		convey.So(called, should.BeFalse)
	})
}

func Test_LinkedList_ForEachBreak(t *testing.T) {
	convey.Convey("LinkedList.ForEachBreak stops early", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3, 4, 5})
		count := 0
		ll.ForEachBreak(func(index int, item int) bool {
			count++
			return item == 3
		})
		convey.So(count, should.Equal, 3)
	})
}

func Test_LinkedList_ForEachBreak_FirstElement(t *testing.T) {
	convey.Convey("LinkedList.ForEachBreak can break on first element", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
		count := 0
		ll.ForEachBreak(func(index int, item int) bool {
			count++
			return true
		})
		convey.So(count, should.Equal, 1)
	})
}

// =============================================================================
// LinkedList — Head / Tail nodes
// =============================================================================

func Test_LinkedList_Head_Tail_Nodes(t *testing.T) {
	convey.Convey("LinkedList.Head and Tail return correct nodes", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
		convey.So(ll.Head().Element, should.Equal, 1)
		convey.So(ll.Tail().Element, should.Equal, 3)
		convey.So(ll.Head().HasNext(), should.BeTrue)
		convey.So(ll.Tail().HasNext(), should.BeFalse)
	})
}

func Test_LinkedList_Node_Next(t *testing.T) {
	convey.Convey("LinkedListNode.Next traverses correctly", t, func() {
		ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
		node := ll.Head()
		convey.So(node.Element, should.Equal, 10)
		node = node.Next()
		convey.So(node.Element, should.Equal, 20)
		node = node.Next()
		convey.So(node.Element, should.Equal, 30)
		convey.So(node.HasNext(), should.BeFalse)
	})
}

// =============================================================================
// LinkedList — Lock variants
// =============================================================================

func Test_LinkedList_LengthLock(t *testing.T) {
	convey.Convey("LinkedList.LengthLock returns correct length", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2})
		convey.So(ll.LengthLock(), should.Equal, 2)
	})
}

func Test_LinkedList_IsEmptyLock(t *testing.T) {
	convey.Convey("LinkedList.IsEmptyLock returns true on empty", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		convey.So(ll.IsEmptyLock(), should.BeTrue)
	})
}

func Test_LinkedList_AddLock(t *testing.T) {
	convey.Convey("LinkedList.AddLock appends with mutex", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		ll.AddLock(1)
		ll.AddLock(2)
		convey.So(ll.Length(), should.Equal, 2)
	})
}

// =============================================================================
// LinkedList — Nil receiver guard
// =============================================================================

func Test_LinkedList_IsEmpty_NilReceiver(t *testing.T) {
	convey.Convey("LinkedList.IsEmpty true on nil receiver", t, func() {
		var ll *coregeneric.LinkedList[int]
		convey.So(ll.IsEmpty(), should.BeTrue)
	})
}

// =============================================================================
// LinkedList — AppendNode / AppendChainOfNodes
// =============================================================================

func Test_LinkedList_AppendNode(t *testing.T) {
	convey.Convey("LinkedList.AppendNode appends a node", t, func() {
		ll := coregeneric.LinkedListFrom([]int{1, 2})
		node := &coregeneric.LinkedListNode[int]{Element: 3}
		ll.AppendNode(node)
		convey.So(ll.Length(), should.Equal, 3)
		convey.So(ll.Last(), should.Equal, 3)
	})
}

func Test_LinkedList_AppendNode_Empty(t *testing.T) {
	convey.Convey("LinkedList.AppendNode on empty list", t, func() {
		ll := coregeneric.EmptyLinkedList[int]()
		node := &coregeneric.LinkedListNode[int]{Element: 99}
		ll.AppendNode(node)
		convey.So(ll.Length(), should.Equal, 1)
		convey.So(ll.First(), should.Equal, 99)
	})
}
