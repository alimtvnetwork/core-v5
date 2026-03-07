package coredatatests

import (
	"sort"
	"testing"

	"gitlab.com/auk-go/core/coredata"
)

// ===== PointerStrings Tests =====

func Test_PointerStrings_Len_NilSlice(t *testing.T) {
	var ps coredata.PointerStrings

	got := ps.Len()
	if got != 0 {
		t.Errorf("PointerStrings.Len() on nil = %d, want 0", got)
	}
}

func Test_PointerStrings_Len_WithElements(t *testing.T) {
	a, b := "alpha", "beta"
	ps := coredata.PointerStrings{&a, &b}

	got := ps.Len()
	if got != 2 {
		t.Errorf("PointerStrings.Len() = %d, want 2", got)
	}
}

func Test_PointerStrings_Less_BothNonNil(t *testing.T) {
	a, b := "alpha", "beta"
	ps := coredata.PointerStrings{&a, &b}

	if !ps.Less(0, 1) {
		t.Error("expected Less(0,1) = true for alpha < beta")
	}

	if ps.Less(1, 0) {
		t.Error("expected Less(1,0) = false for beta < alpha")
	}
}

func Test_PointerStrings_Less_NilFirst(t *testing.T) {
	b := "beta"
	ps := coredata.PointerStrings{nil, &b}

	if !ps.Less(0, 1) {
		t.Error("expected nil < non-nil to be true")
	}
}

func Test_PointerStrings_Less_NilSecond(t *testing.T) {
	a := "alpha"
	ps := coredata.PointerStrings{&a, nil}

	if ps.Less(0, 1) {
		t.Error("expected non-nil < nil to be false")
	}
}

func Test_PointerStrings_Less_BothNil(t *testing.T) {
	ps := coredata.PointerStrings{nil, nil}

	// nil is treated as less, so Less(0,1) = true (first nil < second nil)
	// This is the actual behavior since the first check returns true
	if !ps.Less(0, 1) {
		t.Error("expected Less(0,1) = true when both nil (first nil returns true)")
	}
}

func Test_PointerStrings_Swap(t *testing.T) {
	a, b := "alpha", "beta"
	ps := coredata.PointerStrings{&a, &b}
	ps.Swap(0, 1)

	if *ps[0] != "beta" || *ps[1] != "alpha" {
		t.Errorf("after Swap got [%s,%s], want [beta,alpha]", *ps[0], *ps[1])
	}
}

func Test_PointerStrings_SortInterface(t *testing.T) {
	c, a, b := "charlie", "alpha", "beta"
	ps := coredata.PointerStrings{&c, nil, &a, &b}
	sort.Sort(ps)

	// nil sorts first, then alphabetical
	if ps[0] != nil {
		t.Error("expected nil at index 0 after sort")
	}

	expected := []string{"alpha", "beta", "charlie"}
	for i, exp := range expected {
		if ps[i+1] == nil || *ps[i+1] != exp {
			got := "<nil>"
			if ps[i+1] != nil {
				got = *ps[i+1]
			}
			t.Errorf("sorted[%d] = %s, want %s", i+1, got, exp)
		}
	}
}
