package coredatatests

import (
	"sort"
	"testing"

	"gitlab.com/auk-go/core/coredata"
)

// ===== Integers Tests =====

func Test_Integers_Len_NilSlice(t *testing.T) {
	var integers coredata.Integers

	got := integers.Len()
	if got != 0 {
		t.Errorf("Integers.Len() on nil = %d, want 0", got)
	}
}

func Test_Integers_Len_EmptySlice(t *testing.T) {
	integers := coredata.Integers{}

	got := integers.Len()
	if got != 0 {
		t.Errorf("Integers.Len() on empty = %d, want 0", got)
	}
}

func Test_Integers_Len_WithElements(t *testing.T) {
	integers := coredata.Integers{3, 1, 2}

	got := integers.Len()
	if got != 3 {
		t.Errorf("Integers.Len() = %d, want 3", got)
	}
}

func Test_Integers_Less(t *testing.T) {
	integers := coredata.Integers{5, 3, 8}

	if !integers.Less(1, 0) {
		t.Error("expected Less(1,0) = true for 3 < 5")
	}

	if integers.Less(0, 1) {
		t.Error("expected Less(0,1) = false for 5 < 3")
	}

	if integers.Less(0, 0) {
		t.Error("expected Less(0,0) = false for equal indices")
	}
}

func Test_Integers_Swap(t *testing.T) {
	integers := coredata.Integers{1, 2, 3}
	integers.Swap(0, 2)

	if integers[0] != 3 || integers[2] != 1 {
		t.Errorf("after Swap(0,2) got [%d,%d,%d], want [3,2,1]",
			integers[0], integers[1], integers[2])
	}
}

func Test_Integers_SortInterface(t *testing.T) {
	integers := coredata.Integers{5, 1, 4, 2, 3}
	sort.Sort(integers)

	expected := coredata.Integers{1, 2, 3, 4, 5}
	for i, v := range integers {
		if v != expected[i] {
			t.Errorf("sorted[%d] = %d, want %d", i, v, expected[i])
		}
	}
}
