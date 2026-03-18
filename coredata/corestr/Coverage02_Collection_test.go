package corestr

import (
	"errors"
	"testing"
)

func TestCollection_Basic(t *testing.T) {
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

func TestCollection_NilReceiver_Length(t *testing.T) {
	var c *Collection
	if c.Length() != 0 {
		t.Fatal("expected 0")
	}
	if !c.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func TestCollection_Add(t *testing.T) {
	c := New.Collection.Cap(5)
	c.Add("a").Add("b")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
	if c.Capacity() < 2 {
		t.Fatal("expected cap >= 2")
	}
}

func TestCollection_AddNonEmpty(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmpty("")
	c.AddNonEmpty("a")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddNonEmptyWhitespace(t *testing.T) {
	c := New.Collection.Empty()
	c.AddNonEmptyWhitespace("   ")
	c.AddNonEmptyWhitespace("a")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddError(t *testing.T) {
	c := New.Collection.Empty()
	c.AddError(nil)
	c.AddError(errors.New("e"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddIf(t *testing.T) {
	c := New.Collection.Empty()
	c.AddIf(false, "skip")
	c.AddIf(true, "add")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddIfMany(t *testing.T) {
	c := New.Collection.Empty()
	c.AddIfMany(false, "a", "b")
	c.AddIfMany(true, "c", "d")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_Adds(t *testing.T) {
	c := New.Collection.Empty()
	c.Adds("a", "b", "c")
	if c.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func TestCollection_AddStrings(t *testing.T) {
	c := New.Collection.Empty()
	c.AddStrings([]string{"x", "y"})
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_AddFunc(t *testing.T) {
	c := New.Collection.Empty()
	c.AddFunc(func() string { return "hello" })
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddFuncErr_NoErr(t *testing.T) {
	c := New.Collection.Empty()
	c.AddFuncErr(
		func() (string, error) { return "ok", nil },
		func(err error) { t.Fatal("should not be called") },
	)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddFuncErr_WithErr(t *testing.T) {
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

func TestCollection_AddLock(t *testing.T) {
	c := New.Collection.Empty()
	c.AddLock("a")
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_AddsLock(t *testing.T) {
	c := New.Collection.Empty()
	c.AddsLock("a", "b")
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_AddCollection(t *testing.T) {
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

func TestCollection_AddCollections(t *testing.T) {
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

func TestCollection_RemoveAt(t *testing.T) {
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

func TestCollection_ListStrings(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	if len(c.ListStrings()) != 1 || len(c.ListStringsPtr()) != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_LengthLock(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if c.LengthLock() != 2 {
		t.Fatal("expected 2")
	}
}

func TestCollection_IsEmptyLock(t *testing.T) {
	c := New.Collection.Empty()
	if !c.IsEmptyLock() {
		t.Fatal("expected empty")
	}
}

func TestCollection_AsError(t *testing.T) {
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

func TestCollection_ToError(t *testing.T) {
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

func TestCollection_EachItemSplitBy(t *testing.T) {
	c := New.Collection.Strings([]string{"a,b", "c"})
	result := c.EachItemSplitBy(",")
	if len(result) != 3 {
		t.Fatal("expected 3")
	}
}

func TestCollection_ConcatNew_Empty(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	newC := c.ConcatNew(0)
	if newC.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_ConcatNew_WithItems(t *testing.T) {
	c := New.Collection.Strings([]string{"a"})
	newC := c.ConcatNew(0, "b", "c")
	if newC.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func TestCollection_IsEquals(t *testing.T) {
	a := New.Collection.Strings([]string{"a", "b"})
	b := New.Collection.Strings([]string{"a", "b"})
	if !a.IsEquals(b) {
		t.Fatal("expected equal")
	}
}

func TestCollection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	a := New.Collection.Strings([]string{"A"})
	b := New.Collection.Strings([]string{"a"})
	if !a.IsEqualsWithSensitive(false, b) {
		t.Fatal("expected equal")
	}
	if a.IsEqualsWithSensitive(true, b) {
		t.Fatal("expected not equal")
	}
}

func TestIsCollectionPrecheckEqual(t *testing.T) {
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

func TestCollection_JsonString(t *testing.T) {
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

func TestCollection_AddHashmapsValues(t *testing.T) {
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

func TestCollection_AddHashmapsKeys(t *testing.T) {
	c := New.Collection.Empty()
	hm := New.Hashmap.Empty()
	hm.AddOrUpdate("k", "v")
	c.AddHashmapsKeys(hm)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
	c.AddHashmapsKeys(nil)
}

func TestCollection_AddPointerCollectionsLock(t *testing.T) {
	c := New.Collection.Empty()
	c2 := New.Collection.Strings([]string{"a"})
	c.AddPointerCollectionsLock(c2)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestCollection_HasIndex(t *testing.T) {
	c := New.Collection.Strings([]string{"a", "b"})
	if !c.HasIndex(0) || !c.HasIndex(1) || c.HasIndex(2) || c.HasIndex(-1) {
		t.Fatal("unexpected")
	}
}
