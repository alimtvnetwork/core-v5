package corepayload

import (
	"testing"
)

func TestPayloadsCollection_BasicOps(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	if !c.IsEmpty() {
		t.Fatal("should be empty")
	}
	if c.HasAnyItem() {
		t.Fatal("should not have items")
	}
	if c.Length() != 0 {
		t.Fatal("expected 0")
	}
	if c.Count() != 0 {
		t.Fatal("expected 0")
	}
	if c.LastIndex() != -1 {
		t.Fatal("expected -1")
	}
}

func TestPayloadsCollection_AddAndAccess(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	pw1 := *New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`))
	pw2 := *New.PayloadWrapper.UsingBytes("b", "2", "t", "c", "e", []byte(`"y"`))

	c.Add(pw1)
	c.Add(pw2)

	if c.Length() != 2 {
		t.Fatal("expected 2")
	}

	first := c.First()
	if first == nil || first.Name != "a" {
		t.Fatal("first mismatch")
	}

	last := c.Last()
	if last == nil || last.Name != "b" {
		t.Fatal("last mismatch")
	}

	firstOrDefault := c.FirstOrDefault()
	if firstOrDefault == nil {
		t.Fatal("expected non-nil")
	}

	lastOrDefault := c.LastOrDefault()
	if lastOrDefault == nil {
		t.Fatal("expected non-nil")
	}

	_ = c.FirstDynamic()
	_ = c.LastDynamic()
	_ = c.FirstOrDefaultDynamic()
	_ = c.LastOrDefaultDynamic()
}

func TestPayloadsCollection_EmptyAccess(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	if c.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if c.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func TestPayloadsCollection_SkipTakeLimit(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	for i := 0; i < 5; i++ {
		c.Add(*New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`)))
	}

	s := c.Skip(2)
	if len(s) != 3 {
		t.Fatal("expected 3")
	}
	_ = c.SkipDynamic(1)
	_ = c.SkipCollection(1)

	tk := c.Take(3)
	if len(tk) != 3 {
		t.Fatal("expected 3")
	}
	_ = c.TakeDynamic(2)
	_ = c.TakeCollection(2)
	_ = c.LimitCollection(3)
	_ = c.SafeLimitCollection(3)
	_ = c.LimitDynamic(2)
	_ = c.Limit(3)
}

func TestPayloadsCollection_Adds(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	pw := *New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	c.Adds(pw, pw)
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}

	c.AddsPtr(New.PayloadWrapper.UsingBytes("n", "2", "t", "c", "e", []byte(`"y"`)))
	if c.Length() != 3 {
		t.Fatal("expected 3")
	}

	c.AddsIf(true, pw)
	if c.Length() != 4 {
		t.Fatal("expected 4")
	}
	c.AddsIf(false, pw)
	if c.Length() != 4 {
		t.Fatal("expected 4 still")
	}
}

func TestPayloadsCollection_Filter(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "cat1", "e", []byte(`"x"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("b", "2", "t", "cat2", "e", []byte(`"y"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("c", "3", "t", "cat1", "e", []byte(`"z"`)))

	filtered := c.Filter(func(pw *PayloadWrapper) (bool, bool) {
		return pw.CategoryName == "cat1", false
	})
	if len(filtered) != 2 {
		t.Fatal("expected 2")
	}

	fc := c.FilterCollection(func(pw *PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})
	if fc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestPayloadsCollection_FirstByFilter(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "task1", "cat1", "ent1", []byte(`"x"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("b", "2", "task2", "cat2", "ent2", []byte(`"y"`)))

	r := c.FirstById("1")
	if r == nil || r.Name != "a" {
		t.Fatal("unexpected")
	}

	r2 := c.FirstByCategory("cat2")
	if r2 == nil || r2.Name != "b" {
		t.Fatal("unexpected")
	}

	r3 := c.FirstByTaskType("task1")
	if r3 == nil || r3.Name != "a" {
		t.Fatal("unexpected")
	}

	r4 := c.FirstByEntityType("ent2")
	if r4 == nil || r4.Name != "b" {
		t.Fatal("unexpected")
	}

	r5 := c.FirstById("missing")
	if r5 != nil {
		t.Fatal("expected nil")
	}
}

func TestPayloadsCollection_FilterCollections(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "task1", "cat1", "ent1", []byte(`"x"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("b", "2", "task2", "cat2", "ent2", []byte(`"y"`)))

	_ = c.FilterCollectionByIds("1")
	_ = c.FilterNameCollection("a")
	_ = c.FilterCategoryCollection("cat1")
	_ = c.FilterEntityTypeCollection("ent1")
	_ = c.FilterTaskTypeCollection("task1")
}

func TestPayloadsCollection_SkipFilterCollection(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("b", "2", "t", "c", "e", []byte(`"y"`)))

	fc := c.SkipFilterCollection(func(pw *PayloadWrapper) (bool, bool) {
		return pw.Name == "a", false
	})
	if fc.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestPayloadsCollection_Strings(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	s := c.Strings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func TestPayloadsCollection_Json(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))

	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.JsonString()
	_ = c.String()
	_ = c.PrettyJsonString()
	_ = c.JsonStrings()
	_ = c.CsvStrings()
	_ = c.Join(",")
	_ = c.JoinJsonStrings(",")
	_ = c.JoinCsv()
	_ = c.JoinCsvLine()
}

func TestPayloadsCollection_IsEqual(t *testing.T) {
	c1 := New.PayloadsCollection.Empty()
	c2 := New.PayloadsCollection.Empty()
	pw := *New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`))
	c1.Add(pw)
	c2.Add(pw)

	if !c1.IsEqual(c2) {
		t.Fatal("should be equal")
	}

	var nilC *PayloadsCollection
	if !nilC.IsEqual(nil) {
		t.Fatal("both nil should be equal")
	}
}

func TestPayloadsCollection_Clone(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	cl := c.Clone()
	if cl.Length() != 1 {
		t.Fatal("expected 1")
	}

	cp := c.ClonePtr()
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilC *PayloadsCollection
	if nilC.ClonePtr() != nil {
		t.Fatal("expected nil")
	}
}

func TestPayloadsCollection_Reverse(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("b", "2", "t", "c", "e", []byte(`"y"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("c", "3", "t", "c", "e", []byte(`"z"`)))
	c.Reverse()
	if c.First().Name != "c" {
		t.Fatal("expected c first after reverse")
	}

	// single element
	c2 := New.PayloadsCollection.Empty()
	c2.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c2.Reverse()

	// two elements
	c3 := New.PayloadsCollection.Empty()
	c3.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c3.Add(*New.PayloadWrapper.UsingBytes("b", "2", "t", "c", "e", []byte(`"y"`)))
	c3.Reverse()
	if c3.First().Name != "b" {
		t.Fatal("expected b first")
	}
}

func TestPayloadsCollection_ClearDispose(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c.Clear()
	c.Dispose()
}

func TestPayloadsCollection_Paging(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	for i := 0; i < 25; i++ {
		c.Add(*New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`)))
	}
	if c.GetPagesSize(10) != 3 {
		t.Fatal("expected 3")
	}
	if c.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}
	paged := c.GetPagedCollection(10)
	if len(paged) != 3 {
		t.Fatal("expected 3")
	}
}

func TestPayloadsCollection_ConcatNew(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c2 := c.ConcatNew(*New.PayloadWrapper.UsingBytes("b", "2", "t", "c", "e", []byte(`"y"`)))
	if c2.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestPayloadsCollection_HasIndex(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	if !c.HasIndex(0) {
		t.Fatal("should have index 0")
	}
	if c.HasIndex(1) {
		t.Fatal("should not have index 1")
	}
}

func TestPayloadsCollection_StringsUsingFmt(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	s := c.StringsUsingFmt(func(pw *PayloadWrapper) string {
		return pw.Name
	})
	if len(s) != 1 || s[0] != "a" {
		t.Fatal("unexpected")
	}
}

func TestPayloadsCollection_JoinUsingFmt(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("b", "2", "t", "c", "e", []byte(`"y"`)))
	s := c.JoinUsingFmt(func(pw *PayloadWrapper) string {
		return pw.Name
	}, ",")
	if s != "a,b" {
		t.Fatal("unexpected:", s)
	}
}

func TestNewPayloadsCollectionCreator(t *testing.T) {
	_ = New.PayloadsCollection.Empty()
	_ = New.PayloadsCollection.UsingCap(5)
	pw := New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`))
	_ = New.PayloadsCollection.UsingWrappers(pw)
	_ = New.PayloadsCollection.UsingWrappers()
}

func TestNewPayloadsCollectionCreator_Deserialize(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"hello"`)))
	b, _ := c.Json().Raw()
	c2, err := New.PayloadsCollection.Deserialize(b)
	if err != nil || c2 == nil {
		t.Fatal("unexpected", err)
	}

	_, err2 := New.PayloadsCollection.Deserialize([]byte("invalid"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestPayloadsCollection_FilterWithLimit(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	for i := 0; i < 10; i++ {
		c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	}
	filtered := c.FilterWithLimit(3, func(pw *PayloadWrapper) (bool, bool) {
		return true, false
	})
	if len(filtered) != 3 {
		t.Fatal("expected 3")
	}
}

func TestPayloadsCollection_EmptyJson(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	if c.JsonString() != "" {
		t.Fatal("expected empty string")
	}
	if c.String() != "" {
		t.Fatal("expected empty string")
	}
	if c.PrettyJsonString() != "" {
		t.Fatal("expected empty string")
	}
}

func TestPayloadsCollection_AsInterfaces(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func TestPayloadsCollection_InsertAt(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	c.Add(*New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`)))
	c.Add(*New.PayloadWrapper.UsingBytes("c", "3", "t", "c", "e", []byte(`"z"`)))
	c.InsertAt(1, *New.PayloadWrapper.UsingBytes("b", "2", "t", "c", "e", []byte(`"y"`)))
	if c.Length() != 3 {
		t.Fatal("expected 3")
	}
}

func TestPayloadsCollection_AddsPtrOptions(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	pw := New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`))
	c.AddsPtrOptions(false, pw)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestPayloadsCollection_AddsOptions(t *testing.T) {
	c := New.PayloadsCollection.Empty()
	pw := *New.PayloadWrapper.UsingBytes("a", "1", "t", "c", "e", []byte(`"x"`))
	c.AddsOptions(false, pw)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}
