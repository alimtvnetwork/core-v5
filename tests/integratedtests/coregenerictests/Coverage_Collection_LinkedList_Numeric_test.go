package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coredata/coregeneric"
)

// === Collection uncovered ===

func Test_Cov_Collection_LengthLock(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2})
	if c.LengthLock() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_Collection_IsEmptyLock(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	if !c.IsEmptyLock() {
		t.Error("expected empty")
	}
}

func Test_Cov_Collection_AddLock(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	c.AddLock(1)
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_Collection_AddsLock(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	c.AddsLock(1, 2)
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_Collection_AddIfMany(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	c.AddIfMany(false, 1, 2)
	if c.Length() != 0 {
		t.Error("expected 0")
	}
	c.AddIfMany(true, 1, 2)
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_Collection_ForEachBreak(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2, 3})
	count := 0
	c.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})
	if count != 2 {
		t.Errorf("expected 2 got %d", count)
	}
}

func Test_Cov_Collection_CountFunc(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	n := c.CountFunc(func(v int) bool { return v > 2 })
	if n != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_Collection_SortFunc(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{3, 1, 2})
	c.SortFunc(func(a, b int) bool { return a < b })
	if c.First() != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_Collection_Reverse(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2, 3})
	c.Reverse()
	if c.First() != 3 {
		t.Error("expected 3")
	}
}

func Test_Cov_Collection_ConcatNew(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1})
	n := c.ConcatNew(2, 3)
	if n.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_Cov_Collection_String(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1})
	if c.String() == "" {
		t.Error("expected non-empty")
	}
}

func Test_Cov_Collection_CollectionLenCap(t *testing.T) {
	c := coregeneric.CollectionLenCap[int](3, 10)
	if c.Length() != 3 || c.Capacity() < 10 {
		t.Error("unexpected")
	}
}

// === LinkedList uncovered ===

func Test_Cov_LinkedList_LengthLock(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	if ll.LengthLock() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_LinkedList_IsEmptyLock(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	if !ll.IsEmptyLock() {
		t.Error("expected empty")
	}
}

func Test_Cov_LinkedList_AddLock(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(1)
	if ll.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_LinkedList_AddsIf(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(false, 1, 2)
	if ll.Length() != 0 {
		t.Error("expected 0")
	}
	ll.AddsIf(true, 1, 2)
	if ll.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_LinkedList_AppendChainOfNodes(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1)
	chain := coregeneric.LinkedListFrom([]int{2, 3})
	ll.AppendChainOfNodes(chain.Head())
	if ll.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_Cov_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	chain := coregeneric.LinkedListFrom([]int{1, 2})
	ll.AppendChainOfNodes(chain.Head())
	if ll.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_LinkedList_ForEachBreak(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})
	if count != 2 {
		t.Errorf("expected 2 got %d", count)
	}
}

func Test_Cov_LinkedList_ForEachBreak_FirstItem(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return true
	})
	if count != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_LinkedList_IndexAt(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	node := ll.IndexAt(2)
	if node == nil || node.Element != 30 {
		t.Error("expected 30")
	}
	if ll.IndexAt(-1) != nil {
		t.Error("expected nil for negative")
	}
	if ll.IndexAt(10) != nil {
		t.Error("expected nil for out of range")
	}
}

func Test_Cov_LinkedList_Collection(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	c := ll.Collection()
	if c.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_LinkedList_String(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1})
	if ll.String() == "" {
		t.Error("expected non-empty")
	}
}

// === Numeric funcs uncovered ===

func Test_Cov_CompareNumeric(t *testing.T) {
	if coregeneric.CompareNumeric(1, 2) != corecomparator.LeftLess {
		t.Error("expected LeftLess")
	}
	if coregeneric.CompareNumeric(2, 1) != corecomparator.LeftGreater {
		t.Error("expected LeftGreater")
	}
	if coregeneric.CompareNumeric(1, 1) != corecomparator.Equal {
		t.Error("expected Equal")
	}
}

func Test_Cov_Clamp(t *testing.T) {
	if coregeneric.Clamp(5, 1, 10) != 5 {
		t.Error("in range")
	}
	if coregeneric.Clamp(-1, 0, 10) != 0 {
		t.Error("below min")
	}
	if coregeneric.Clamp(20, 0, 10) != 10 {
		t.Error("above max")
	}
}

func Test_Cov_Abs(t *testing.T) {
	if coregeneric.Abs(-5) != 5 {
		t.Error("expected 5")
	}
}

func Test_Cov_AbsDiff(t *testing.T) {
	if coregeneric.AbsDiff(3, 5) != 2 {
		t.Error("expected 2")
	}
}

func Test_Cov_Sign(t *testing.T) {
	if coregeneric.Sign(-5) != -1 {
		t.Error("expected -1")
	}
	if coregeneric.Sign(0) != 0 {
		t.Error("expected 0")
	}
	if coregeneric.Sign(5) != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_SafeDiv(t *testing.T) {
	if coregeneric.SafeDiv(10, 0) != 0 {
		t.Error("expected 0 for div by zero")
	}
	if coregeneric.SafeDiv(10, 2) != 5 {
		t.Error("expected 5")
	}
}

func Test_Cov_SafeDivOrDefault(t *testing.T) {
	if coregeneric.SafeDivOrDefault(10, 0, -1) != -1 {
		t.Error("expected -1")
	}
}

func Test_Cov_MinOfSlice(t *testing.T) {
	if coregeneric.MinOfSlice([]int{3, 1, 2}) != 1 {
		t.Error("expected 1")
	}
}

func Test_Cov_MaxOfSlice(t *testing.T) {
	if coregeneric.MaxOfSlice([]int{3, 1, 2}) != 3 {
		t.Error("expected 3")
	}
}

func Test_Cov_InRangeExclusive(t *testing.T) {
	if !coregeneric.InRangeExclusive(5, 0, 10) {
		t.Error("expected true")
	}
	if coregeneric.InRangeExclusive(0, 0, 10) {
		t.Error("expected false for boundary")
	}
}

func Test_Cov_IsNegative(t *testing.T) {
	if !coregeneric.IsNegative(-1) {
		t.Error("expected true")
	}
}

func Test_Cov_IsNonNegative(t *testing.T) {
	if !coregeneric.IsNonNegative(0) {
		t.Error("expected true")
	}
}
