package corestr

import (
	"errors"
	"testing"
)

func TestCollection_Basic_C02(t *testing.T) {
	c := New.Collection.Empty()
	if !c.IsEmpty() || c.HasItems() || c.HasAnyItem() {
		t.Fatal("should be empty")
	}
	if c.Length() != 0 || c.Count() != 0 || c.LastIndex() != -1 {
		t.Fatal("wrong length")
	}
	if c.Capacity() != 0 {
		t.Fatal("expected 0 cap")
	}
	if c.HasIndex(0) {
		t.Fatal("should not have index")
	}
}

func TestCollection_NilReceiver_Length_C02(t *testing.T) {
	var c *Collection
	if c.Length() != 0 {
		t.Fatal("expected 0")
	}
	if !c.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func TestCollection_Add_C02(t *testing.T) {
	c := New.Collection.Cap(5)
	c.Add("a").Add("b")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
	if c.Capacity() < 2 {
		t.Fatal("expected cap >= 2")
	}
}

func TestCollection_AddNonEmpty_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmpty("")
	c.AddNonEmpty("a")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddNonEmptyWhitespace_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmptyWhitespace("   ")
	c.AddNonEmptyWhitespace("a")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddError_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddError(nil)
	c.AddError(errors.New("e"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddIf_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddIf(false, "skip")
	c.AddIf(true, "add")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddIfMany_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddIfMany(false, "a", "b")
	c.AddIfMany(true, "c", "d")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_Adds_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.Adds("a", "b", "c")
	if c.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func TestCollection_AddStrings_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddStrings([]string{"x", "y"})
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_AddFunc_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddFunc(func() string { return "hello" })
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddFuncErr_NoErr_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddFuncErr(
		func() (string, error) { return "ok", nil },
		func(err error) { t.Fatal("should not be called") },
	)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddFuncErr_WithErr_C02(t *testing.T) {
	c := New.Collection.Empty()
	called := false
	c.AddFuncErr(
		func() (string, error) { return "", errors.New("e") },
		func(err error) { called = true },
	)
	if c.Length() != 0 || !called {
		t.Fatal("expected 0 and handler called")
	}
}

func TestCollection_AddLock_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddLock("a")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddsLock_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddsLock("a", "b")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_AddCollection_C02(t *testing.T) {
	c1 := New.Collection.Strings([]string{"a"})
	c2 := New.Collection.Strings([]string{"b"})
	c1.AddCollection(c2)
	if c1.Length() != 2 {
		t.Fatal("expected 2")
	}
	// empty collection
	c1.AddCollection(New.Collection.Empty())
	if c1.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_AddCollections_C02(t *testing.T) {
	c := New.Collection.Empty()
	c.AddCollections(
		New.Collection.Strings([]string{"a"}),
		New.Collection.Empty(),
		New.Collection.Strings([]string{"b"}),
	)
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_RemoveAt_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b", "c"})
	if !c.RemoveAt(1) {
		t.Fatal("expected success")
	}
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
	// out of range
	if c.RemoveAt(-1) || c.RemoveAt(100) {
		t.Fatal("expected failure")
	}
}

func TestCollection_ListStrings_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if len(c.ListStrings()) != 1 || len(c.ListStringsPtr()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_LengthLock_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if c.LengthLock() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_IsEmptyLock_C02(t *testing.T) {
	c := New.Collection.Empty()
	if !c.IsEmptyLock() {
		t.Fatal("expected empty")
	}
}

func TestCollection_AsError_C02(t *testing.T) {
	c := New.Collection.Empty()
	if c.AsDefaultError() != nil {
		t.Fatal("expected nil")
	}
	if c.AsError(",") != nil {
		t.Fatal("expected nil")
	}
	c.Add("e1")
	if c.AsDefaultError() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestCollection_ToError_C02(t *testing.T) {
	c := New.Collection.Empty()
	if c.ToError(",") != nil {
		t.Fatal("expected nil")
	}
	if c.ToDefaultError() != nil {
		t.Fatal("expected nil")
	}
	c.Add("e")
	if c.ToError(",") == nil {
		t.Fatal("expected non-nil")
	}
}

func TestCollection_EachItemSplitBy_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a,b", "c"})
	result := c.EachItemSplitBy(",")
	if len(result) != 3 {
		t.Fatal("expected 3")
	}
}

func TestCollection_ConcatNew_Empty_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	newC := c.ConcatNew(0)
	if newC.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_ConcatNew_WithItems_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	newC := c.ConcatNew(0, "b", "c")
	if newC.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func TestCollection_IsEquals_C02(t *testing.T) {
	a := New.Collection.Strings([]string{"a", "b"})
	b := New.Collection.Strings([]string{"a", "b"})
	if !a.IsEquals(b) {
		t.Fatal("expected equal")
	}
}

func TestCollection_IsEqualsWithSensitive_CaseInsensitive_C02(t *testing.T) {
	a := New.Collection.Strings([]string{"A"})
	b := New.Collection.Strings([]string{"a"})
	if !a.IsEqualsWithSensitive(false, b) {
		t.Fatal("expected equal")
	}
	if a.IsEqualsWithSensitive(true, b) {
		t.Fatal("expected not equal")
	}
}

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

func TestCollection_JsonString_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	s := c.JsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := c.JsonStringMust()
	if s2 == "" {
		t.Fatal("expected non-empty")
	}
	s3 := c.StringJSON()
	if s3 == "" {
		t.Fatal("expected non-empty")
	}
}

func TestCollection_AddHashmapsValues_C02(t *testing.T) {
	c := New.Collection.Empty()
	hm := New.Hashmap.Empty()
	hm.AddOrUpdate("k", "v")
	c.AddHashmapsValues(hm)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
	c.AddHashmapsValues(nil)
	c.AddHashmapsValues(New.Hashmap.Empty())
}

func TestCollection_AddHashmapsKeys_C02(t *testing.T) {
	c := New.Collection.Empty()
	hm := New.Hashmap.Empty()
	hm.AddOrUpdate("k", "v")
	c.AddHashmapsKeys(hm)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
	c.AddHashmapsKeys(nil)
}

func TestCollection_AddPointerCollectionsLock_C02(t *testing.T) {
	c := New.Collection.Empty()
	c2 := New.Collection.Strings([]string{"a"})
	c.AddPointerCollectionsLock(c2)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_HasIndex_C02(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.HasIndex(0) || !c.HasIndex(1) || c.HasIndex(2) || c.HasIndex(-1) {
		t.Fatal("unexpected")
	}
}
