package coredynamic

import (
	"testing"
)

func TestDynamicCollection_Basic(t *testing.T) {
	dc := EmptyDynamicCollection()
	if !dc.IsEmpty() || dc.HasAnyItem() { t.Fatal("expected empty") }
	if dc.Length() != 0 || dc.Count() != 0 { t.Fatal("expected 0") }
	var nilDc *DynamicCollection
	if nilDc.Length() != 0 || !nilDc.IsEmpty() { t.Fatal("expected 0/empty") }
	dc.Add(NewDynamic("a", true)).Add(NewDynamic("b", true))
	if dc.Length() != 2 { t.Fatal("expected 2") }
	_ = dc.At(0)
	_ = dc.Items()
	_ = dc.First()
	_ = dc.Last()
	_ = dc.FirstDynamic()
	_ = dc.LastDynamic()
	_ = dc.FirstOrDefault()
	_ = dc.LastOrDefault()
	_ = dc.FirstOrDefaultDynamic()
	_ = dc.LastOrDefaultDynamic()
	_ = dc.HasIndex(0)
	_ = dc.LastIndex()
	_ = dc.Skip(1)
	_ = dc.Take(1)
	_ = dc.SkipDynamic(1)
	_ = dc.TakeDynamic(1)
	_ = dc.Limit(1)
	_ = dc.LimitDynamic(1)
	_ = dc.SkipCollection(1)
	_ = dc.TakeCollection(1)
	_ = dc.LimitCollection(1)
	_ = dc.SafeLimitCollection(1)
	_ = dc.ListStrings()
	_ = dc.ListStringsPtr()
	_ = dc.AnyItems()
	_ = dc.AnyItemsCollection()
}

func TestDynamicCollection_Mutators(t *testing.T) {
	dc := NewDynamicCollection(5)
	dc.AddAny("a", true)
	dc.AddAnyNonNull("b", true)
	dc.AddAnyNonNull(nil, true)
	dc.AddAnyMany("c", "d")
	dc.AddAnyMany()
	dc.AddPtr(NewDynamicPtr("e", true))
	dc.AddPtr(nil)
	dc.AddManyPtr(NewDynamicPtr("f", true), nil)
	dc.AddManyPtr()
	dc.RemoveAt(0)
	dc.RemoveAt(999)
}

func TestDynamicCollection_Json(t *testing.T) {
	dc := NewDynamicCollection(2)
	dc.AddAny("a", true)
	_, _ = dc.JsonString()
	_ = dc.JsonStringMust()
	_, _ = dc.MarshalJSON()
	_ = dc.JsonResultsCollection()
	_ = dc.JsonResultsPtrCollection()
	_ = dc.GetPagesSize(1)
	_ = dc.GetPagesSize(0)
}

func TestAnyCollection_Basic(t *testing.T) {
	ac := EmptyAnyCollection()
	if !ac.IsEmpty() || ac.HasAnyItem() { t.Fatal("expected empty") }
	ac.Add("a").Add("b")
	if ac.Length() != 2 { t.Fatal("expected 2") }
	_ = ac.At(0)
	_ = ac.Items()
	_ = ac.First()
	_ = ac.Last()
	_ = ac.FirstDynamic()
	_ = ac.LastDynamic()
	_ = ac.FirstOrDefault()
	_ = ac.LastOrDefault()
	_ = ac.FirstOrDefaultDynamic()
	_ = ac.LastOrDefaultDynamic()
	_ = ac.HasIndex(0)
	_ = ac.LastIndex()
	_ = ac.Count()
	_ = ac.Skip(1)
	_ = ac.Take(1)
	_ = ac.SkipDynamic(1)
	_ = ac.TakeDynamic(1)
	_ = ac.Limit(1)
	_ = ac.LimitDynamic(1)
	_ = ac.SkipCollection(1)
	_ = ac.TakeCollection(1)
	_ = ac.LimitCollection(1)
	_ = ac.SafeLimitCollection(1)
	_ = ac.DynamicItems()
	_ = ac.DynamicCollection()
	_ = ac.ListStrings(false)
	_ = ac.ListStringsPtr(false)
	ac.RemoveAt(0)
	ac.RemoveAt(999)
}

func TestAnyCollection_Mutators(t *testing.T) {
	ac := NewAnyCollection(5)
	ac.AddAny("a", true)
	ac.AddNonNull("b")
	ac.AddNonNull(nil)
	ac.AddNonNullDynamic("c", true)
	ac.AddNonNullDynamic(nil, true)
	ac.AddAnyManyDynamic("d", "e")
	ac.AddAnyManyDynamic()
	ac.AddMany("f", nil, "g")
	ac.AddMany()
}

func TestAnyCollection_Loop(t *testing.T) {
	ac := NewAnyCollection(3)
	ac.Add("a").Add("b")
	ac.Loop(false, func(i int, item any) bool { return false })
	ac.Loop(true, func(i int, item any) bool { return false })
	ac.LoopDynamic(false, func(i int, item Dynamic) bool { return false })
}

func TestAnyCollection_Json(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("a")
	_, _ = ac.JsonString()
	_ = ac.JsonStringMust()
	_ = ac.GetPagesSize(1)
}

// ── Generic Collection ──

func TestGenericCollection_Basic(t *testing.T) {
	c := NewCollection[string](5)
	if !c.IsEmpty() || c.HasAnyItem() { t.Fatal("expected empty") }
	c.Add("a").Add("b").AddMany("c", "d")
	if c.Length() != 4 || c.Count() != 4 { t.Fatal("expected 4") }
	_ = c.At(0)
	_ = c.First()
	_ = c.Last()
	_, _ = c.FirstOrDefault()
	_, _ = c.LastOrDefault()
	_ = c.Items()
	_ = c.HasIndex(0)
	_ = c.LastIndex()
	_ = c.Skip(1)
	_ = c.Take(2)
	_ = c.Limit(2)
	_ = c.SkipCollection(1)
	_ = c.TakeCollection(2)
	_ = c.LimitCollection(2)
	_ = c.SafeLimitCollection(2)
	c.RemoveAt(0)
	c.RemoveAt(999)
	_ = c.Strings()
	_ = c.String()
	_, _ = c.JsonString()
	_ = c.JsonStringMust()
	_, _ = c.MarshalJSON()
	_ = c.GetPagesSize(2)
	_ = c.GetPagesSize(0)
}

func TestGenericCollection_LoopFilter(t *testing.T) {
	c := CollectionFrom([]int{1, 2, 3, 4, 5})
	c.Loop(func(i int, item int) bool { return false })
	c.LoopAsync(func(i int, item int) {})
	f := c.Filter(func(item int) bool { return item > 2 })
	if f.Length() != 3 { t.Fatal("expected 3") }
}

func TestGenericCollection_NonNil(t *testing.T) {
	c := NewCollection[int](5)
	val := 42
	c.AddNonNil(&val)
	c.AddNonNil(nil)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestGenericCollection_Clone(t *testing.T) {
	c := CollectionClone([]string{"a", "b"})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestGenericCollection_Empty(t *testing.T) {
	c := EmptyCollection[string]()
	if !c.IsEmpty() { t.Fatal("expected empty") }
	var nilC *Collection[string]
	if nilC.Length() != 0 || !nilC.IsEmpty() { t.Fatal("expected 0/empty") }
}

func TestGenericCollection_ClearDispose(t *testing.T) {
	c := CollectionFrom([]int{1, 2})
	c.Clear()
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Dispose()
}
