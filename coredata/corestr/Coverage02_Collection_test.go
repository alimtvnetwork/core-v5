package corestr

import (
	"testing"
)

// ── isCollectionPrecheckEqual (unexported — must remain in source package) ──

func TestIsCollectionPrecheckEqual_C02(t *testing.T) {
	// both nil
	r, h := isCollectionPrecheckEqual(nil, nil)
	if !r || !h {
		t.Fatal("expected true, true")
	}

	// one nil
	c := New.Collection.Strings([]string{"a"})
	r, h = isCollectionPrecheckEqual(nil, c)
	if r || !h {
		t.Fatal("expected false, true")
	}

	// same ptr
	r, h = isCollectionPrecheckEqual(c, c)
	if !r || !h {
		t.Fatal("expected true, true")
	}

	// both empty
	e1 := New.Collection.Empty()
	e2 := New.Collection.Empty()
	r, h = isCollectionPrecheckEqual(e1, e2)
	if !r || !h {
		t.Fatal("expected true, true")
	}

	// one empty
	r, h = isCollectionPrecheckEqual(e1, c)
	if r || !h {
		t.Fatal("expected false, true")
	}

	// diff length
	c2 := New.Collection.Strings([]string{"a", "b"})
	r, h = isCollectionPrecheckEqual(c, c2)
	if r || !h {
		t.Fatal("expected false, true")
	}

	// same length, not handled
	c3 := New.Collection.Strings([]string{"b"})
	r, h = isCollectionPrecheckEqual(c, c3)
	if h {
		t.Fatal("expected not handled")
	}
}
