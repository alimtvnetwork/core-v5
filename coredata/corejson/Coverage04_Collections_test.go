package corejson

import (
	"testing"
)

func TestResultsCollection_BasicOps(t *testing.T) {
	c := NewResultsCollection.Empty()
	if !c.IsEmpty() || c.HasAnyItem() || c.Length() != 0 {
		t.Fatal("should be empty")
	}
	if c.LastIndex() != -1 {
		t.Fatal("expected -1")
	}

	c.Add(NewResult.Any("a"))
	c.Add(NewResult.Any("b"))
	c.Add(NewResult.Any("c"))
	if c.Length() != 3 {
		t.Fatal("expected 3")
	}
	if c.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if !c.HasAnyItem() {
		t.Fatal("should have items")
	}
}

func TestResultsCollection_FirstLast(t *testing.T) {
	c := NewResultsCollection.Empty()
	if c.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if c.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}

	c.Add(NewResult.Any("a"))
	c.Add(NewResult.Any("b"))
	if c.FirstOrDefault() == nil || c.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResultsCollection_TakeSkipLimit(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any(1))
	c.Add(NewResult.Any(2))
	c.Add(NewResult.Any(3))

	t2 := c.Take(2)
	if t2.Length() != 2 {
		t.Fatal("expected 2")
	}

	s := c.Skip(1)
	if s.Length() != 2 {
		t.Fatal("expected 2")
	}

	l := c.Limit(2)
	if l.Length() != 2 {
		t.Fatal("expected 2")
	}

	l2 := c.Limit(-1)
	if l2.Length() != 3 {
		t.Fatal("expected all")
	}

	// empty
	empty := NewResultsCollection.Empty()
	if empty.Take(1).Length() != 0 {
		t.Fatal("expected 0")
	}
	if empty.Skip(0).Length() != 0 {
		t.Fatal("expected 0")
	}
	if empty.Limit(1).Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestResultsCollection_AddMethods(t *testing.T) {
	c := NewResultsCollection.Empty()
	r := NewResult.AnyPtr("x")
	c.AddSkipOnNil(r)
	c.AddSkipOnNil(nil)
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}

	c.AddNonNilNonError(r)
	c.AddNonNilNonError(nil)
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}

	c.AddPtr(r)
	c.AddPtr(nil)
	if c.Length() != 3 {
		t.Fatal("expected 3")
	}

	c.Adds(NewResult.Any("a"), NewResult.Any("b"))
	if c.Length() != 5 {
		t.Fatal("expected 5")
	}

	c.AddsPtr(r, nil)
	if c.Length() != 6 {
		t.Fatal("expected 6")
	}

	c.AddAny("z")
	c.AddAny(nil)
	if c.Length() != 7 {
		t.Fatal("expected 7")
	}

	c.AddAnyItems("a", nil, "b")
	if c.Length() != 9 {
		t.Fatal("expected 9")
	}
}

func TestResultsCollection_GetAt(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	r := c.GetAt(0)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResultsCollection_GetAtSafe(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	r := c.GetAtSafe(0)
	if r == nil {
		t.Fatal("expected non-nil")
	}
	r2 := c.GetAtSafe(5)
	if r2 != nil {
		t.Fatal("expected nil for out of bounds")
	}
	r3 := c.GetAtSafe(-1)
	if r3 != nil {
		t.Fatal("expected nil for negative")
	}
}

func TestResultsCollection_HasError(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	if c.HasError() {
		t.Fatal("should not have error")
	}
}

func TestResultsCollection_AllErrors(t *testing.T) {
	c := NewResultsCollection.Empty()
	errList, hasErr := c.AllErrors()
	if hasErr || len(errList) != 0 {
		t.Fatal("expected no errors")
	}
}

func TestResultsCollection_GetErrorsStrings(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	s := c.GetErrorsStrings()
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
	sp := c.GetErrorsStringsPtr()
	if sp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResultsCollection_GetErrorsAsSingleString(t *testing.T) {
	c := NewResultsCollection.Empty()
	s := c.GetErrorsAsSingleString()
	_ = s
}

func TestResultsCollection_GetErrorsAsSingle(t *testing.T) {
	c := NewResultsCollection.Empty()
	err := c.GetErrorsAsSingle()
	_ = err
}

func TestResultsCollection_GetStrings(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("a"))
	s := c.GetStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
	sp := c.GetStringsPtr()
	if sp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResultsCollection_NonPtr_Ptr(t *testing.T) {
	c := NewResultsCollection.Empty()
	np := c.NonPtr()
	_ = np
	pp := c.Ptr()
	if pp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResultsCollection_Clear_Dispose(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	c.Clear()

	c2 := NewResultsCollection.Empty()
	c2.Add(NewResult.Any("x"))
	c2.Dispose()
}

func TestResultsCollection_Json(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	j := c.Json()
	if j.HasError() {
		t.Fatal(j.Error)
	}
	jp := c.JsonPtr()
	if jp == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResultsCollection_JsonModel(t *testing.T) {
	c := NewResultsCollection.Empty()
	_ = c.JsonModel()
	_ = c.JsonModelAny()
}

func TestResultsCollection_Clone(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	cl := c.ShadowClone()
	if cl.Length() != 1 {
		t.Fatal("expected 1")
	}

	cl2 := c.Clone(true)
	if cl2.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestResultsCollection_ClonePtr(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("x"))
	cp := c.ClonePtr(false)
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilC *ResultsCollection
	if nilC.ClonePtr(false) != nil {
		t.Fatal("expected nil")
	}
}

func TestResultsCollection_Paging(t *testing.T) {
	c := NewResultsCollection.Empty()
	for i := 0; i < 25; i++ {
		c.Add(NewResult.Any(i))
	}

	pages := c.GetPagesSize(10)
	if pages != 3 {
		t.Fatal("expected 3 pages")
	}

	if c.GetPagesSize(0) != 0 {
		t.Fatal("expected 0")
	}

	paged := c.GetPagedCollection(10)
	if len(paged) != 3 {
		t.Fatal("expected 3")
	}

	single := c.GetSinglePageCollection(10, 1)
	if single.Length() != 10 {
		t.Fatal("expected 10")
	}
}

func TestResultsCollection_AsJsonContractsBinder(t *testing.T) {
	c := NewResultsCollection.Empty()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func TestResultsCollection_UnmarshalAt(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("hello"))
	var out string
	err := c.UnmarshalAt(0, &out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}
}

func TestResultsCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := NewResultsCollection.Empty()
	c.Add(NewResult.Any("a"))
	c.Add(NewResult.Any(42))

	var s string
	var i int
	errList, hasErr := c.UnmarshalIntoSameIndex(&s, &i)
	if hasErr {
		t.Fatal("expected no error")
	}
	_ = errList
}

// ================== ResultsPtrCollection ==================

func TestResultsPtrCollection_BasicOps(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	if !c.IsEmpty() || c.HasAnyItem() {
		t.Fatal("should be empty")
	}

	c.Add(NewResult.AnyPtr("a"))
	c.Add(NewResult.AnyPtr("b"))
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestResultsPtrCollection_FirstLast(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	if c.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if c.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}

	c.Add(NewResult.AnyPtr("a"))
	if c.FirstOrDefault() == nil || c.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResultsPtrCollection_TakeSkipLimit(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	c.Add(NewResult.AnyPtr(1))
	c.Add(NewResult.AnyPtr(2))
	c.Add(NewResult.AnyPtr(3))

	if c.Take(2).Length() != 2 {
		t.Fatal("expected 2")
	}
	if c.Skip(1).Length() != 2 {
		t.Fatal("expected 2")
	}
	if c.Limit(2).Length() != 2 {
		t.Fatal("expected 2")
	}
	if c.Limit(-1).Length() != 3 {
		t.Fatal("expected all")
	}
}

func TestResultsPtrCollection_AddMethods(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	r := NewResult.AnyPtr("x")
	c.AddSkipOnNil(r)
	c.AddSkipOnNil(nil)
	c.AddNonNilNonError(r)
	c.AddNonNilNonError(nil)
	c.AddResult(NewResult.Any("y"))
	c.Adds(r, nil)
	c.AddAny("z")
	c.AddAny(nil)
	c.AddAnyItems("a", nil, "b")
	c.AddNonNilItems(r, nil)
	c.AddNonNilItemsPtr(r, nil)
}

func TestResultsPtrCollection_ClearDispose(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	c.Add(NewResult.AnyPtr("x"))
	c.Clear()
	c.Dispose()
}

func TestResultsPtrCollection_Clone(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	c.Add(NewResult.AnyPtr("x"))
	cl := c.Clone(false)
	if cl == nil {
		t.Fatal("expected non-nil")
	}

	var nilC *ResultsPtrCollection
	if nilC.Clone(false) != nil {
		t.Fatal("expected nil")
	}
}

func TestResultsPtrCollection_Json(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	c.Add(NewResult.AnyPtr("x"))
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.NonPtr()
	_ = c.Ptr()
}

func TestResultsPtrCollection_Paging(t *testing.T) {
	c := NewResultsPtrCollection.Empty()
	for i := 0; i < 25; i++ {
		c.Add(NewResult.AnyPtr(i))
	}
	if c.GetPagesSize(10) != 3 {
		t.Fatal("expected 3")
	}
	paged := c.GetPagedCollection(10)
	if len(paged) != 3 {
		t.Fatal("expected 3")
	}
}

// ================== BytesCollection ==================

func TestBytesCollection_BasicOps(t *testing.T) {
	c := NewBytesCollection.Empty()
	if !c.IsEmpty() {
		t.Fatal("should be empty")
	}

	c.Add([]byte("hello"))
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
	if !c.HasAnyItem() {
		t.Fatal("should have items")
	}
	if c.FirstOrDefault() == nil {
		t.Fatal("expected non-nil")
	}
	if c.LastOrDefault() == nil {
		t.Fatal("expected non-nil")
	}

	empty := NewBytesCollection.Empty()
	if empty.FirstOrDefault() != nil {
		t.Fatal("expected nil")
	}
	if empty.LastOrDefault() != nil {
		t.Fatal("expected nil")
	}
}

func TestBytesCollection_AddMethods(t *testing.T) {
	c := NewBytesCollection.Empty()
	c.AddSkipOnNil(nil)
	c.AddSkipOnNil([]byte("x"))
	c.AddNonEmpty([]byte{})
	c.AddNonEmpty([]byte("y"))
	c.AddPtr([]byte{})
	c.AddPtr([]byte("z"))
	c.Adds([]byte("a"), []byte{}, []byte("b"))
	if c.Length() != 4 {
		t.Fatal("expected 4, got", c.Length())
	}
}

func TestBytesCollection_TakeSkipLimit(t *testing.T) {
	c := NewBytesCollection.Empty()
	c.Add([]byte("a"))
	c.Add([]byte("b"))
	c.Add([]byte("c"))

	if c.Take(2).Length() != 2 {
		t.Fatal("expected 2")
	}
	if c.Skip(1).Length() != 2 {
		t.Fatal("expected 2")
	}
	if c.Limit(2).Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestBytesCollection_GetAt_Strings(t *testing.T) {
	c := NewBytesCollection.Empty()
	c.Add([]byte("hello"))
	b := c.GetAt(0)
	if string(b) != "hello" {
		t.Fatal("unexpected")
	}

	s := c.Strings()
	if len(s) != 1 || s[0] != "hello" {
		t.Fatal("unexpected")
	}

	_ = c.StringsPtr()
}

func TestBytesCollection_GetAtSafe(t *testing.T) {
	c := NewBytesCollection.Empty()
	c.Add([]byte("x"))
	if c.GetAtSafe(0) == nil {
		t.Fatal("expected non-nil")
	}
	if c.GetAtSafe(5) != nil {
		t.Fatal("expected nil")
	}
	if c.GetAtSafePtr(0) == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBytesCollection_ClearDispose(t *testing.T) {
	c := NewBytesCollection.Empty()
	c.Add([]byte("x"))
	c.Clear()
	c.Dispose()
}

func TestBytesCollection_Clone(t *testing.T) {
	c := NewBytesCollection.Empty()
	c.Add([]byte("x"))
	cl := c.ShadowClone()
	_ = cl
	cl2 := c.Clone(true)
	_ = cl2
}

func TestBytesCollection_ClonePtr(t *testing.T) {
	c := NewBytesCollection.Empty()
	c.Add([]byte("x"))
	cp := c.ClonePtr(false)
	if cp == nil {
		t.Fatal("expected non-nil")
	}

	var nilC *BytesCollection
	if nilC.ClonePtr(false) != nil {
		t.Fatal("expected nil")
	}
}

func TestBytesCollection_Json(t *testing.T) {
	c := BytesCollection{Items: [][]byte{[]byte("x")}}
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.JsonModel()
	_ = c.JsonModelAny()
}

func TestBytesCollection_Paging(t *testing.T) {
	c := NewBytesCollection.Empty()
	for i := 0; i < 25; i++ {
		c.Add([]byte{byte(i)})
	}
	if c.GetPagesSize(10) != 3 {
		t.Fatal("expected 3")
	}
	paged := c.GetPagedCollection(10)
	if len(paged) != 3 {
		t.Fatal("expected 3")
	}
}

func TestBytesCollection_AddAny(t *testing.T) {
	c := NewBytesCollection.Empty()
	err := c.AddAny("hello")
	if err != nil {
		t.Fatal(err)
	}
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func TestBytesCollection_AddAnyItems(t *testing.T) {
	c := NewBytesCollection.Empty()
	err := c.AddAnyItems("a", "b")
	if err != nil {
		t.Fatal(err)
	}
	if c.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ================== MapResults ==================

func TestMapResults_BasicOps(t *testing.T) {
	m := NewMapResults.Empty()
	if !m.IsEmpty() || m.HasAnyItem() {
		t.Fatal("should be empty")
	}

	m.Add("k1", NewResult.Any("v1"))
	if m.Length() != 1 {
		t.Fatal("expected 1")
	}

	r := m.GetByKey("k1")
	if r == nil {
		t.Fatal("expected non-nil")
	}

	r2 := m.GetByKey("missing")
	if r2 != nil {
		t.Fatal("expected nil")
	}
}

func TestMapResults_AddMethods(t *testing.T) {
	m := NewMapResults.Empty()
	m.AddPtr("k", NewResult.AnyPtr("v"))
	m.AddPtr("k2", nil)
	m.AddSkipOnNil("k3", NewResult.AnyPtr("v3"))
	m.AddSkipOnNil("k4", nil)
	_ = m.AddAny("k5", "v5")
	_ = m.AddAny("k6", nil)
	_ = m.AddAnySkipOnNil("k7", "v7")
	_ = m.AddAnySkipOnNil("k8", nil)
	m.AddAnyNonEmpty("k9", "v9")
	m.AddAnyNonEmpty("k10", nil)
	m.AddAnyNonEmptyNonError("k11", "v11")
	m.AddAnyNonEmptyNonError("k12", nil)
}

func TestMapResults_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	m := NewMapResults.Empty()
	m.Add("b", NewResult.Any(1))
	m.Add("a", NewResult.Any(2))

	keys := m.AllKeys()
	if len(keys) != 2 {
		t.Fatal("expected 2")
	}

	sorted := m.AllKeysSorted()
	if sorted[0] != "a" {
		t.Fatal("expected sorted")
	}

	vals := m.AllValues()
	if len(vals) != 2 {
		t.Fatal("expected 2")
	}
}

func TestMapResults_HasError_AllErrors(t *testing.T) {
	m := NewMapResults.Empty()
	m.Add("k", NewResult.Any("v"))
	if m.HasError() {
		t.Fatal("should not have error")
	}
	errList, hasErr := m.AllErrors()
	if hasErr || len(errList) != 0 {
		t.Fatal("expected no errors")
	}
}

func TestMapResults_GetErrorsStrings(t *testing.T) {
	m := NewMapResults.Empty()
	s := m.GetErrorsStrings()
	if len(s) != 0 {
		t.Fatal("expected empty")
	}
	_ = m.GetErrorsStringsPtr()
	_ = m.GetErrorsAsSingleString()
	_ = m.GetErrorsAsSingle()
}

func TestMapResults_GetStrings(t *testing.T) {
	m := NewMapResults.Empty()
	m.Add("k", NewResult.Any("v"))
	s := m.GetStrings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
	_ = m.GetStringsPtr()
}

func TestMapResults_ClearDispose(t *testing.T) {
	m := NewMapResults.Empty()
	m.Add("k", NewResult.Any("v"))
	m.Clear()
	m.Dispose()
}

func TestMapResults_Json(t *testing.T) {
	m := NewMapResults.Empty()
	m.Add("k", NewResult.Any("v"))
	_ = m.Json()
	_ = m.JsonPtr()
	_ = m.JsonModel()
	_ = m.JsonModelAny()
}

func TestMapResults_Paging(t *testing.T) {
	m := NewMapResults.Empty()
	for i := 0; i < 25; i++ {
		m.Add(Serialize.ToString(i), NewResult.Any(i))
	}
	if m.GetPagesSize(10) != 3 {
		t.Fatal("expected 3")
	}
	paged := m.GetPagedCollection(10)
	if len(paged) != 3 {
		t.Fatal("expected 3")
	}
}

func TestMapResults_AllResultsCollection(t *testing.T) {
	m := NewMapResults.Empty()
	m.Add("k", NewResult.Any("v"))
	c := m.AllResultsCollection()
	if c.Length() != 1 {
		t.Fatal("expected 1")
	}

	_ = m.AllResults()
	_ = m.ResultCollection()
}
