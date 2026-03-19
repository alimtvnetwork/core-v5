package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
)

func Test_QW_Collection_Length_NilItems(t *testing.T) {
	var c *coregeneric.Collection[string]
	if c.Length() != 0 {
		t.Fatal("expected 0 for nil collection")
	}
}

func Test_QW_LinkedList_IndexAt_EndOfList(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	ll.Add("a")
	// Access index beyond list length — covers the out-of-range early return
	node := ll.IndexAt(5)
	if node != nil {
		t.Fatal("expected nil for out-of-range index")
	}
}

func Test_QW_MinOf_SecondSmaller(t *testing.T) {
	// Cover the else branch (a >= b)
	result := coregeneric.MinOf(5, 3)
	if result != 3 {
		t.Fatal("expected 3")
	}
}

func Test_QW_MaxOf_SecondLarger(t *testing.T) {
	// Cover the else branch (a <= b)
	result := coregeneric.MaxOf(3, 5)
	if result != 5 {
		t.Fatal("expected 5")
	}
}

func Test_QW_MinOfSlice_NonMinElements(t *testing.T) {
	// Cover the case where v < result is false
	result := coregeneric.MinOfSlice([]int{3, 5, 1, 4})
	if result != 1 {
		t.Fatal("expected 1")
	}
}
