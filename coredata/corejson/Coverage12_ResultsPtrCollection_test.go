package corejson

import (
	"errors"
	"testing"
	"time"
)

func TestResultsPtrCollection_Basic(t *testing.T) {
	var nilC *ResultsPtrCollection
	if nilC.Length() != 0 { t.Fatal("expected 0") }
	if nilC.LastIndex() != -1 { t.Fatal("expected -1") }
	if !nilC.IsEmpty() { t.Fatal("expected empty") }
	if nilC.HasAnyItem() { t.Fatal("expected false") }
	if nilC.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if nilC.LastOrDefault() != nil { t.Fatal("expected nil") }
}

func TestResultsPtrCollection_AddAndAccess(t *testing.T) {
	c := NewResultsPtrCollection.Default()
	r := NewResult.AnyPtr("hello")
	c.Add(r)
	if c.Length() != 1 { t.Fatal("expected 1") }
	if c.FirstOrDefault() == nil { t.Fatal("expected non-nil") }
	if c.LastOrDefault() == nil { t.Fatal("expected non-nil") }
	if c.GetAt(0) == nil { t.Fatal("expected non-nil") }
}

func TestResultsPtrCollection_TakeSkipLimit(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		c.Add(NewResult.AnyPtr(i))
	}
	if c.Take(3).Length() != 3 { t.Fatal("expected 3") }
	if c.Skip(2).Length() != 3 { t.Fatal("expected 3") }
	if c.Limit(3).Length() != 3 { t.Fatal("expected 3") }
	if c.Limit(-2).Length() != 5 { t.Fatal("expected all") }

	empty := NewResultsPtrCollection.Empty()
	if empty.Take(1).Length() != 0 { t.Fatal("expected 0") }
	if empty.Skip(0).Length() != 0 { t.Fatal("expected 0") }
	if empty.Limit(1).Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_AddMethods(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(10)
	c.AddSkipOnNil(nil)
	c.AddSkipOnNil(NewResult.AnyPtr("x"))
	c.AddNonNilNonError(nil)
	c.AddNonNilNonError(&Result{Error: errors.New("e")})
	c.AddNonNilNonError(NewResult.AnyPtr("x"))
	c.AddResult(NewResult.Any("x"))
	c.Adds(nil, NewResult.AnyPtr("x"))
	c.AddAny(nil)
	c.AddAny("x")
	c.AddAnyItems(nil, "y")
	c.AddResultsCollection(nil)
	sub := NewResultsPtrCollection.UsingResults(NewResult.AnyPtr("sub"))
	c.AddResultsCollection(sub)
	c.AddNonNilItemsPtr(nil)
	c.AddNonNilItemsPtr(nil, NewResult.AnyPtr("x"))
	c.AddNonNilItems(nil, NewResult.AnyPtr("x"))
}

func TestResultsPtrCollection_Errors(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(2)
	c.Add(NewResult.AnyPtr("ok"))
	c.Add(&Result{Error: errors.New("e")})
	if !c.HasError() { t.Fatal("expected error") }
	errs, has := c.AllErrors()
	if !has || len(errs) != 1 { t.Fatal("expected 1") }
	strs := c.GetErrorsStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func TestResultsPtrCollection_UnmarshalAt(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(2)
	c.Add(NewResult.AnyPtr("hello"))
	c.Add(nil)
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestResultsPtrCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(3)
	c.Add(NewResult.AnyPtr("hello"))
	c.Add(NewResult.AnyPtr(42))
	c.Add(&Result{Error: errors.New("e")})
	var s string
	var i int
	errs, _ := c.UnmarshalIntoSameIndex(&s, &i, nil)
	if len(errs) != 3 { t.Fatal("expected 3") }
	c2 := NewResultsPtrCollection.Empty()
	c2.UnmarshalIntoSameIndex(nil)
}

func TestResultsPtrCollection_GetAtSafe(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(1)
	c.Add(NewResult.AnyPtr("x"))
	if c.GetAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if c.GetAtSafe(-1) != nil { t.Fatal("expected nil") }
	if c.GetAtSafe(5) != nil { t.Fatal("expected nil") }
	if c.GetAtSafeUsingLength(0, 1) == nil { t.Fatal("expected non-nil") }
}

func TestResultsPtrCollection_Serializers(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(2)
	c.AddSerializer(nil)
	c.AddSerializers()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunctions()
}

func TestResultsPtrCollection_PtrNonPtr(t *testing.T) {
	c := NewResultsPtrCollection.Default()
	_ = c.NonPtr()
	_ = c.Ptr()
}

func TestResultsPtrCollection_ClearDispose(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(2)
	c.Add(NewResult.AnyPtr("x"))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Dispose()
	var nilC *ResultsPtrCollection
	nilC.Clear()
	nilC.Dispose()
}

func TestResultsPtrCollection_GetStrings(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(2)
	c.Add(NewResult.AnyPtr("a"))
	strs := c.GetStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = c.GetStringsPtr()
	empty := NewResultsPtrCollection.Empty()
	if len(empty.GetStrings()) != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_Paging(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Add(NewResult.AnyPtr(i))
	}
	if c.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := c.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
	single := c.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }

	small := NewResultsPtrCollection.UsingResults(NewResult.AnyPtr("x"))
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_JsonMethods(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(1)
	c.Add(NewResult.AnyPtr("x"))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func TestResultsPtrCollection_Clone(t *testing.T) {
	c := NewResultsPtrCollection.UsingCap(2)
	c.Add(NewResult.AnyPtr("x"))
	cp := c.Clone(true)
	if cp == nil || cp.Length() != 1 { t.Fatal("expected 1") }
	var nilC *ResultsPtrCollection
	if nilC.Clone(true) != nil { t.Fatal("expected nil") }

	empty := NewResultsPtrCollection.Empty()
	ec := empty.Clone(true)
	if ec.Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_Creators(t *testing.T) {
	_ = NewResultsPtrCollection.AnyItems("a", "b")
	_ = NewResultsPtrCollection.AnyItemsPlusCap(5, "a")
	_ = NewResultsPtrCollection.AnyItemsPlusCap(5)
	_ = NewResultsPtrCollection.UsingResults(NewResult.AnyPtr("x"))
	_ = NewResultsPtrCollection.UsingResultsPlusCap(5, NewResult.AnyPtr("x"))
	_ = NewResultsPtrCollection.UsingResultsPlusCap(5)
	_ = NewResultsPtrCollection.Serializers()
	_, _ = NewResultsPtrCollection.UnmarshalUsingBytes([]byte(`{}`))
}
