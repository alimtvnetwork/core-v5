package corejson

import (
	"errors"
	"testing"
	"time"
)

func TestResultsCollection_Basic(t *testing.T) {
	var nilC *ResultsCollection
	if nilC.Length() != 0 { t.Fatal("expected 0") }
	if nilC.LastIndex() != -1 { t.Fatal("expected -1") }
	if !nilC.IsEmpty() { t.Fatal("expected empty") }
	if nilC.HasAnyItem() { t.Fatal("expected false") }
	if nilC.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if nilC.LastOrDefault() != nil { t.Fatal("expected nil") }
}

func TestResultsCollection_AddAndAccess(t *testing.T) {
	c := NewResultsCollection.Default()
	r := NewResult.Any("hello")
	c.Add(r)
	if c.Length() != 1 { t.Fatal("expected 1") }
	if c.FirstOrDefault() == nil { t.Fatal("expected non-nil") }
	if c.LastOrDefault() == nil { t.Fatal("expected non-nil") }
	g := c.GetAt(0)
	if g == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_TakeSkipLimit(t *testing.T) {
	c := NewResultsCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		c.Add(NewResult.Any(i))
	}
	tk := c.Take(3)
	if tk.Length() != 3 { t.Fatal("expected 3") }
	sk := c.Skip(2)
	if sk.Length() != 3 { t.Fatal("expected 3") }
	lm := c.Limit(3)
	if lm.Length() != 3 { t.Fatal("expected 3") }
	lm2 := c.Limit(-2) // TakeAllMinusOne
	if lm2.Length() != 5 { t.Fatal("expected all") }

	empty := NewResultsCollection.Empty()
	if empty.Take(1).Length() != 0 { t.Fatal("expected 0") }
	if empty.Skip(0).Length() != 0 { t.Fatal("expected 0") }
	if empty.Limit(1).Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsCollection_AddMethods(t *testing.T) {
	c := NewResultsCollection.UsingCap(5)
	c.AddSkipOnNil(nil)
	c.AddSkipOnNil(&Result{Bytes: []byte(`1`)})
	if c.Length() != 1 { t.Fatal("expected 1") }

	c.AddNonNilNonError(nil)
	c.AddNonNilNonError(&Result{Error: errors.New("e")})
	c.AddNonNilNonError(&Result{Bytes: []byte(`2`)})
	if c.Length() != 2 { t.Fatal("expected 2") }

	c.AddPtr(nil)
	c.AddPtr(&Result{Bytes: []byte(`3`)})
	if c.Length() != 3 { t.Fatal("expected 3") }

	c.Adds(NewResult.Any("a"), NewResult.Any("b"))
	if c.Length() != 5 { t.Fatal("expected 5") }
	c.Adds()

	c.AddsPtr(nil, &Result{Bytes: []byte(`4`)})
	if c.Length() != 6 { t.Fatal("expected 6") }

	c.AddAny(nil)
	c.AddAny("x")
	if c.Length() != 7 { t.Fatal("expected 7") }

	c.AddAnyItems(nil, "y")
	if c.Length() != 8 { t.Fatal("expected 8") }

	c.AddAnyItemsSlice(nil)
	c.AddAnyItemsSlice([]any{nil, "z"})
	if c.Length() != 9 { t.Fatal("expected 9") }

	c.AddResultsCollection(nil)
	sub := NewResultsCollection.UsingResults(NewResult.Any("sub"))
	c.AddResultsCollection(sub)
	if c.Length() != 10 { t.Fatal("expected 10") }

	c.AddNonNilItemsPtr(nil)
	c.AddNonNilItemsPtr(nil, &Result{Bytes: []byte(`5`)})
}

func TestResultsCollection_Errors(t *testing.T) {
	c := NewResultsCollection.UsingCap(3)
	c.Add(NewResult.Any("ok"))
	c.Add(Result{Error: errors.New("e1")})
	if !c.HasError() { t.Fatal("expected error") }

	errs, has := c.AllErrors()
	if !has || len(errs) != 1 { t.Fatal("expected 1 error") }

	strs := c.GetErrorsStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = c.GetErrorsStringsPtr()
	_ = c.GetErrorsAsSingleString()
	_ = c.GetErrorsAsSingle()
}

func TestResultsCollection_UnmarshalAt(t *testing.T) {
	c := NewResultsCollection.UsingCap(1)
	c.Add(NewResult.Any("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestResultsCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := NewResultsCollection.UsingCap(3)
	c.Add(NewResult.Any("hello"))
	c.Add(NewResult.Any(42))
	c.Add(Result{Error: errors.New("e")})
	var s string
	var i int
	errs, has := c.UnmarshalIntoSameIndex(&s, &i, nil)
	if len(errs) != 3 { t.Fatal("expected 3") }
	_ = has
	c2 := NewResultsCollection.Empty()
	c2.UnmarshalIntoSameIndex(nil)
}

func TestResultsCollection_GetAtSafe(t *testing.T) {
	c := NewResultsCollection.UsingCap(1)
	c.Add(NewResult.Any("x"))
	if c.GetAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if c.GetAtSafe(-1) != nil { t.Fatal("expected nil") }
	if c.GetAtSafe(5) != nil { t.Fatal("expected nil") }
	if c.GetAtSafeUsingLength(0, 1) == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_Serializers(t *testing.T) {
	c := NewResultsCollection.UsingCap(2)
	c.AddSerializerFunc(nil)
	c.AddSerializerFunctions(nil)
	c.AddSerializer(nil)
	c.AddSerializers(nil)
}

func TestResultsCollection_MapResults(t *testing.T) {
	c := NewResultsCollection.UsingCap(2)
	mr := NewMapResults.Empty()
	c.AddMapResults(mr)
	c.AddRawMapResults(nil)
	c.AddRawMapResults(map[string]Result{"a": NewResult.Any("x")})
}

func TestResultsCollection_PtrNonPtr(t *testing.T) {
	c := NewResultsCollection.Default()
	_ = c.NonPtr()
	_ = c.Ptr()
}

func TestResultsCollection_ClearDispose(t *testing.T) {
	c := NewResultsCollection.UsingCap(2)
	c.Add(NewResult.Any("x"))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Dispose()
	var nilC *ResultsCollection
	nilC.Clear()
	nilC.Dispose()
}

func TestResultsCollection_GetStrings(t *testing.T) {
	c := NewResultsCollection.UsingCap(2)
	c.Add(NewResult.Any("a"))
	strs := c.GetStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = c.GetStringsPtr()
	empty := NewResultsCollection.Empty()
	if len(empty.GetStrings()) != 0 { t.Fatal("expected 0") }
}

func TestResultsCollection_Paging(t *testing.T) {
	c := NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Add(NewResult.Any(i))
	}
	pages := c.GetPagesSize(3)
	if pages != 4 { t.Fatal("expected 4") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	if c.GetPagesSize(-1) != 0 { t.Fatal("expected 0") }

	paged := c.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4 pages") }

	single := c.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }
	single2 := c.GetSinglePageCollection(3, 4)
	if single2.Length() != 1 { t.Fatal("expected 1") }

	small := NewResultsCollection.UsingResults(NewResult.Any("x"))
	smallPaged := small.GetPagedCollection(5)
	if len(smallPaged) != 1 { t.Fatal("expected 1") }
	smallSingle := small.GetSinglePageCollection(5, 1)
	if smallSingle.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_JsonMethods(t *testing.T) {
	c := NewResultsCollection.UsingCap(2)
	c.Add(NewResult.Any("x"))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func TestResultsCollection_Clone(t *testing.T) {
	c := NewResultsCollection.UsingCap(2)
	c.Add(NewResult.Any("x"))
	sc := c.ShadowClone()
	if sc.Length() != 1 { t.Fatal("expected 1") }
	dc := c.Clone(true)
	if dc.Length() != 1 { t.Fatal("expected 1") }
	cp := c.ClonePtr(true)
	if cp == nil || cp.Length() != 1 { t.Fatal("expected 1") }
	var nilC *ResultsCollection
	if nilC.ClonePtr(true) != nil { t.Fatal("expected nil") }

	empty := NewResultsCollection.Empty()
	ec := empty.Clone(true)
	if ec.Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsCollection_Creators(t *testing.T) {
	_ = NewResultsCollection.AnyItems("a", "b")
	_ = NewResultsCollection.AnyItemsPlusCap(5, "a")
	_ = NewResultsCollection.AnyItemsPlusCap(5)
	_ = NewResultsCollection.UsingResultsPtr(nil, &Result{Bytes: []byte(`"x"`)})
	_ = NewResultsCollection.UsingResultsPtrPlusCap(5, nil)
	_ = NewResultsCollection.UsingResultsPlusCap(5, NewResult.Any("x"))
	_ = NewResultsCollection.UsingResultsPlusCap(5)
	_ = NewResultsCollection.UsingResults(NewResult.Any("x"))
	_ = NewResultsCollection.Serializers()
	_ = NewResultsCollection.SerializerFunctions()
	_ = NewResultsCollection.UsingJsoners()
	_ = NewResultsCollection.UsingJsonersNonNull(5)
}
