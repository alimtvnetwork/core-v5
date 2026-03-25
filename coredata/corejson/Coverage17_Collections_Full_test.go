package corejson

import (
	"errors"
	"testing"
)

// ── BytesCollection ──

func TestBytesCollection_Length(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_LastIndex(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	if bc.LastIndex() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_IsEmpty(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	if !bc.IsEmpty() { t.Fatal("expected true") }
}

func TestBytesCollection_HasAnyItem(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	if !bc.HasAnyItem() { t.Fatal("expected true") }
}

func TestBytesCollection_FirstOrDefault(t *testing.T) {
	bc := Empty.BytesCollectionPtr()
	if bc.FirstOrDefault() != nil { t.Fatal("expected nil") }
	bc.Add([]byte(`"a"`))
	if bc.FirstOrDefault() == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_LastOrDefault(t *testing.T) {
	bc := Empty.BytesCollectionPtr()
	if bc.LastOrDefault() != nil { t.Fatal("expected nil") }
	bc.Add([]byte(`"a"`))
	if bc.LastOrDefault() == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_Take(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	r := bc.Take(1)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_Take_Empty(t *testing.T) {
	bc := Empty.BytesCollectionPtr()
	r := bc.Take(1)
	if r.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_Skip(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	r := bc.Skip(1)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_Skip_Empty(t *testing.T) {
	bc := Empty.BytesCollectionPtr()
	r := bc.Skip(0)
	if r.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_Limit(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`)).Add([]byte(`"b"`))
	r := bc.Limit(1)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_Limit_TakeAll(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	r := bc.Limit(-1)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddSkipOnNil(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddSkipOnNil(nil)
	if bc.Length() != 0 { t.Fatal("expected 0") }
	bc.AddSkipOnNil([]byte(`"a"`))
	if bc.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddNonEmpty(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddNonEmpty([]byte{})
	if bc.Length() != 0 { t.Fatal("expected 0") }
	bc.AddNonEmpty([]byte(`"a"`))
	if bc.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddResultPtr(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	r := NewResult.Any("hello")
	bc.AddResultPtr(&r)
	if bc.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddResult(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	r := NewResult.Any("hello")
	bc.AddResult(r)
	if bc.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_GetAt(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	if len(bc.GetAt(0)) == 0 { t.Fatal("expected non-empty") }
}

func TestBytesCollection_JsonResultAt(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	r := bc.JsonResultAt(0)
	if r == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_UnmarshalAt(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"hello"`))
	var target string
	err := bc.UnmarshalAt(0, &target)
	if err != nil || target != "hello" { t.Fatal("unexpected") }
}

func TestBytesCollection_AddSerializer(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddSerializer(nil)
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_AddSerializers(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddSerializers()
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_AddSerializerFunc(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddSerializerFunc(nil)
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_AddSerializerFunctions(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddSerializerFunctions()
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_InjectIntoSameIndex_Nil(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	errs, hasErr := bc.InjectIntoSameIndex(nil)
	if hasErr || len(errs) != 0 { t.Fatal("unexpected") }
}

func TestBytesCollection_UnmarshalIntoSameIndex_Nil(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	errs, hasErr := bc.UnmarshalIntoSameIndex(nil)
	if hasErr || len(errs) != 0 { t.Fatal("unexpected") }
}

func TestBytesCollection_GetAtSafe(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	if bc.GetAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if bc.GetAtSafe(5) != nil { t.Fatal("expected nil") }
}

func TestBytesCollection_GetAtSafePtr(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	if bc.GetAtSafePtr(0) == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_GetResultAtSafe(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	if bc.GetResultAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if bc.GetResultAtSafe(5) != nil { t.Fatal("expected nil") }
}

func TestBytesCollection_GetAtSafeUsingLength(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	if bc.GetAtSafeUsingLength(0, 1) == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_AddPtr(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddPtr([]byte(`"a"`))
	if bc.Length() != 1 { t.Fatal("expected 1") }
	bc.AddPtr(nil)
	if bc.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_Adds(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Adds([]byte(`"a"`), []byte(`"b"`))
	if bc.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_Adds_Empty(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Adds()
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_AddAnyItems(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	err := bc.AddAnyItems("hello")
	if err != nil || bc.Length() != 1 { t.Fatal("unexpected") }
}

func TestBytesCollection_AddAnyItems_Empty(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	err := bc.AddAnyItems()
	if err != nil { t.Fatal("unexpected") }
}

func TestBytesCollection_AddAny(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	err := bc.AddAny("hello")
	if err != nil || bc.Length() != 1 { t.Fatal("unexpected") }
}

func TestBytesCollection_AddMapResults(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	mr := Empty.MapResults()
	bc.AddMapResults(mr)
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_AddRawMapResults(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddRawMapResults(nil)
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_AddsPtr(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	r := NewResult.Any("hello")
	bc.AddsPtr(&r)
	if bc.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddsPtr_Nil(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddsPtr(nil)
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_AddBytesCollection(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc2 := NewBytesCollection.UsingCap(5)
	bc2.Add([]byte(`"a"`))
	bc.AddBytesCollection(bc2)
	if bc.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_Clear(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	bc.Clear()
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_Dispose(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	bc.Dispose()
	if bc.Items != nil { t.Fatal("expected nil") }
}

func TestBytesCollection_Strings(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	s := bc.Strings()
	if len(s) != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_StringsPtr(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	s := bc.StringsPtr()
	if len(s) != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddJsoners(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.AddJsoners(true)
	if bc.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_GetPagesSize(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	for i := 0; i < 10; i++ { bc.Add([]byte(`"a"`)) }
	if bc.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if bc.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_GetPagedCollection(t *testing.T) {
	bc := NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ { bc.Add([]byte(`"a"`)) }
	pages := bc.GetPagedCollection(3)
	if len(pages) != 4 { t.Fatal("expected 4") }
}

func TestBytesCollection_GetPagedCollection_Small(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	bc.Add([]byte(`"a"`))
	pages := bc.GetPagedCollection(5)
	if len(pages) != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_GetSinglePageCollection(t *testing.T) {
	bc := NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ { bc.Add([]byte(`"a"`)) }
	page := bc.GetSinglePageCollection(3, 2)
	if page.Length() != 3 { t.Fatal("expected 3") }
}

func TestBytesCollection_JsonModel(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	_ = bc.JsonModel()
}

func TestBytesCollection_JsonModelAny(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	_ = bc.JsonModelAny()
}

func TestBytesCollection_Json(t *testing.T) {
	bc := BytesCollection{Items: [][]byte{[]byte(`"a"`)}}
	r := bc.Json()
	_ = r
}

func TestBytesCollection_JsonPtr(t *testing.T) {
	bc := BytesCollection{Items: [][]byte{[]byte(`"a"`)}}
	r := bc.JsonPtr()
	_ = r
}

func TestBytesCollection_AsJsonContractsBinder(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	_ = bc.AsJsonContractsBinder()
}

func TestBytesCollection_AsJsoner(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	_ = bc.AsJsoner()
}

func TestBytesCollection_AsJsonParseSelfInjector(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	_ = bc.AsJsonParseSelfInjector()
}

func TestBytesCollection_ShadowClone(t *testing.T) {
	bc := &BytesCollection{Items: [][]byte{[]byte(`"a"`)}}
	cloned := bc.ShadowClone()
	_ = cloned
}

func TestBytesCollection_ClonePtr(t *testing.T) {
	bc := &BytesCollection{Items: [][]byte{[]byte(`"a"`)}}
	c := bc.ClonePtr(true)
	if c == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_ClonePtr_Nil(t *testing.T) {
	var bc *BytesCollection
	if bc.ClonePtr(true) != nil { t.Fatal("expected nil") }
}

// ── ResultsCollection ──

func TestResultsCollection_Basic(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	r := NewResult.Any("hello")
	rc.Add(r)
	if rc.Length() != 1 { t.Fatal("expected 1") }
	if rc.IsEmpty() { t.Fatal("expected false") }
	if !rc.HasAnyItem() { t.Fatal("expected true") }
}

func TestResultsCollection_FirstOrDefault(t *testing.T) {
	rc := Empty.ResultsCollection()
	if rc.FirstOrDefault() != nil { t.Fatal("expected nil") }
	rc.Add(NewResult.Any("a"))
	if rc.FirstOrDefault() == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_LastOrDefault(t *testing.T) {
	rc := Empty.ResultsCollection()
	if rc.LastOrDefault() != nil { t.Fatal("expected nil") }
	rc.Add(NewResult.Any("a"))
	if rc.LastOrDefault() == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_Take(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a")).Add(NewResult.Any("b"))
	r := rc.Take(1)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_Limit(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a")).Add(NewResult.Any("b"))
	r := rc.Limit(1)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_Skip(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a")).Add(NewResult.Any("b"))
	r := rc.Skip(1)
	if r.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddSkipOnNil(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.AddSkipOnNil(nil)
	if rc.Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsCollection_AddNonNilNonError(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.AddNonNilNonError(nil)
	errResult := &Result{Error: errors.New("fail")}
	rc.AddNonNilNonError(errResult)
	if rc.Length() != 0 { t.Fatal("expected 0") }
	good := NewResult.Any("hello")
	rc.AddNonNilNonError(&good)
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_GetAt(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	r := rc.GetAt(0)
	if r == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_HasError(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	if rc.HasError() { t.Fatal("expected false") }
}

func TestResultsCollection_AllErrors(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	errs, has := rc.AllErrors()
	if has || len(errs) != 0 { t.Fatal("unexpected") }
}

func TestResultsCollection_GetErrorsStrings(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	s := rc.GetErrorsStrings()
	if len(s) != 0 { t.Fatal("expected 0") }
}

func TestResultsCollection_GetErrorsStringsPtr(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.GetErrorsStringsPtr()
}

func TestResultsCollection_GetErrorsAsSingleString(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.GetErrorsAsSingleString()
}

func TestResultsCollection_GetErrorsAsSingle(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.GetErrorsAsSingle()
}

func TestResultsCollection_Adds(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Adds(NewResult.Any("a"), NewResult.Any("b"))
	if rc.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_AddPtr(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	r := NewResult.Any("a")
	rc.AddPtr(&r)
	if rc.Length() != 1 { t.Fatal("expected 1") }
	rc.AddPtr(nil)
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddsPtr(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	r := NewResult.Any("a")
	rc.AddsPtr(&r, nil)
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddAnyItems(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.AddAnyItems("a", nil)
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddAnyItemsSlice(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.AddAnyItemsSlice([]any{"a"})
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddResultsCollection(t *testing.T) {
	rc1 := NewResultsCollection.UsingCap(5)
	rc1.Add(NewResult.Any("a"))
	rc2 := NewResultsCollection.UsingCap(5)
	rc2.AddResultsCollection(rc1)
	if rc2.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddNonNilItemsPtr(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	r := NewResult.Any("a")
	rc.AddNonNilItemsPtr(&r, nil)
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_NonPtr(t *testing.T) {
	rc := ResultsCollection{Items: []Result{NewResult.Any("a")}}
	_ = rc.NonPtr()
}

func TestResultsCollection_Ptr(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.Ptr()
}

func TestResultsCollection_Clear(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	rc.Clear()
	if rc.Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsCollection_Dispose(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	rc.Dispose()
	if rc.Items != nil { t.Fatal("expected nil") }
}

func TestResultsCollection_GetStrings(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	s := rc.GetStrings()
	if len(s) != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_GetStringsPtr(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.GetStringsPtr()
}

func TestResultsCollection_GetAtSafe(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	if rc.GetAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if rc.GetAtSafe(5) != nil { t.Fatal("expected nil") }
}

func TestResultsCollection_GetAtSafeUsingLength(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	if rc.GetAtSafeUsingLength(0, 1) == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_GetPagesSize(t *testing.T) {
	rc := NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ { rc.Add(NewResult.Any("a")) }
	if rc.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if rc.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
}

func TestResultsCollection_GetPagedCollection(t *testing.T) {
	rc := NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ { rc.Add(NewResult.Any("a")) }
	pages := rc.GetPagedCollection(3)
	if len(pages) != 4 { t.Fatal("expected 4") }
}

func TestResultsCollection_GetPagedCollection_Small(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	rc.Add(NewResult.Any("a"))
	pages := rc.GetPagedCollection(5)
	if len(pages) != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_JsonModel(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.JsonModel()
}

func TestResultsCollection_JsonModelAny(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.JsonModelAny()
}

func TestResultsCollection_AsJsonContractsBinder(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.AsJsonContractsBinder()
}

func TestResultsCollection_AsJsoner(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.AsJsoner()
}

func TestResultsCollection_AsJsonParseSelfInjector(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	_ = rc.AsJsonParseSelfInjector()
}

func TestResultsCollection_ShadowClone(t *testing.T) {
	rc := ResultsCollection{Items: []Result{NewResult.Any("a")}}
	_ = rc.ShadowClone()
}

func TestResultsCollection_ClonePtr(t *testing.T) {
	rc := &ResultsCollection{Items: []Result{NewResult.Any("a")}}
	c := rc.ClonePtr(true)
	if c == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_ClonePtr_Nil(t *testing.T) {
	var rc *ResultsCollection
	if rc.ClonePtr(true) != nil { t.Fatal("expected nil") }
}

// ── MapResults ──

func TestMapResults_Basic(t *testing.T) {
	mr := Empty.MapResults()
	if !mr.IsEmpty() { t.Fatal("expected true") }
	mr.Add("k", NewResult.Any("v"))
	if mr.Length() != 1 { t.Fatal("expected 1") }
	if !mr.HasAnyItem() { t.Fatal("expected true") }
}

func TestMapResults_AddSkipOnNil(t *testing.T) {
	mr := Empty.MapResults()
	mr.AddSkipOnNil("k", nil)
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func TestMapResults_GetByKey(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("k", NewResult.Any("v"))
	r := mr.GetByKey("k")
	if r == nil { t.Fatal("expected non-nil") }
	if mr.GetByKey("missing") != nil { t.Fatal("expected nil") }
}

func TestMapResults_HasError(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("k", NewResult.Any("v"))
	if mr.HasError() { t.Fatal("expected false") }
}

func TestMapResults_AllErrors(t *testing.T) {
	mr := Empty.MapResults()
	errs, has := mr.AllErrors()
	if has || len(errs) != 0 { t.Fatal("unexpected") }
}

func TestMapResults_GetErrorsStrings(t *testing.T) {
	mr := Empty.MapResults()
	s := mr.GetErrorsStrings()
	if len(s) != 0 { t.Fatal("expected 0") }
}

func TestMapResults_GetErrorsStringsPtr(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.GetErrorsStringsPtr()
}

func TestMapResults_GetErrorsAsSingleString(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.GetErrorsAsSingleString()
}

func TestMapResults_GetErrorsAsSingle(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.GetErrorsAsSingle()
}

func TestMapResults_AddPtr(t *testing.T) {
	mr := Empty.MapResults()
	r := NewResult.Any("v")
	mr.AddPtr("k", &r)
	if mr.Length() != 1 { t.Fatal("expected 1") }
	mr.AddPtr("k2", nil)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddAny(t *testing.T) {
	mr := Empty.MapResults()
	err := mr.AddAny("k", "hello")
	if err != nil || mr.Length() != 1 { t.Fatal("unexpected") }
}

func TestMapResults_AddAny_Nil(t *testing.T) {
	mr := Empty.MapResults()
	err := mr.AddAny("k", nil)
	if err == nil { t.Fatal("expected error") }
}

func TestMapResults_AddAnySkipOnNil(t *testing.T) {
	mr := Empty.MapResults()
	err := mr.AddAnySkipOnNil("k", nil)
	if err != nil { t.Fatal("expected nil") }
}

func TestMapResults_AddAnyNonEmptyNonError(t *testing.T) {
	mr := Empty.MapResults()
	mr.AddAnyNonEmptyNonError("k", "hello")
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddAnyNonEmpty(t *testing.T) {
	mr := Empty.MapResults()
	mr.AddAnyNonEmpty("k", "hello")
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddNonEmptyNonErrorPtr(t *testing.T) {
	mr := Empty.MapResults()
	r := NewResult.Any("v")
	mr.AddNonEmptyNonErrorPtr("k", &r)
	if mr.Length() != 1 { t.Fatal("expected 1") }
	mr.AddNonEmptyNonErrorPtr("k2", nil)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddMapResults(t *testing.T) {
	mr1 := Empty.MapResults()
	mr1.Add("a", NewResult.Any("1"))
	mr2 := Empty.MapResults()
	mr2.AddMapResults(mr1)
	if mr2.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddMapAnyItems(t *testing.T) {
	mr := Empty.MapResults()
	mr.AddMapAnyItems(map[string]any{"k": "v"})
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AllKeys(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("a", NewResult.Any("1"))
	keys := mr.AllKeys()
	if len(keys) != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AllKeysSorted(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("b", NewResult.Any("2"))
	mr.Add("a", NewResult.Any("1"))
	keys := mr.AllKeysSorted()
	if keys[0] != "a" { t.Fatal("expected sorted") }
}

func TestMapResults_AllValues(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("a", NewResult.Any("1"))
	vals := mr.AllValues()
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AllResultsCollection(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("a", NewResult.Any("1"))
	rc := mr.AllResultsCollection()
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_GetStrings(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("a", NewResult.Any("1"))
	s := mr.GetStrings()
	if len(s) != 1 { t.Fatal("expected 1") }
}

func TestMapResults_ResultCollection(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("a", NewResult.Any("1"))
	rc := mr.ResultCollection()
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_JsonModel(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.JsonModel()
}

func TestMapResults_JsonModelAny(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.JsonModelAny()
}

func TestMapResults_Clear(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("a", NewResult.Any("1"))
	mr.Clear()
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

func TestMapResults_Dispose(t *testing.T) {
	mr := Empty.MapResults()
	mr.Add("a", NewResult.Any("1"))
	mr.Dispose()
	if mr.Items != nil { t.Fatal("expected nil") }
}

func TestMapResults_AsJsonContractsBinder(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.AsJsonContractsBinder()
}

func TestMapResults_AsJsoner(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.AsJsoner()
}

func TestMapResults_AsJsonParseSelfInjector(t *testing.T) {
	mr := Empty.MapResults()
	_ = mr.AsJsonParseSelfInjector()
}

func TestMapResults_AddMapResultsUsingCloneOption(t *testing.T) {
	mr := Empty.MapResults()
	items := map[string]Result{"a": NewResult.Any("1")}
	mr.AddMapResultsUsingCloneOption(false, false, items)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddMapResultsUsingCloneOption_Clone(t *testing.T) {
	mr := Empty.MapResults()
	items := map[string]Result{"a": NewResult.Any("1")}
	mr.AddMapResultsUsingCloneOption(true, true, items)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddMapResultsUsingCloneOption_Empty(t *testing.T) {
	mr := Empty.MapResults()
	mr.AddMapResultsUsingCloneOption(false, false, nil)
	if mr.Length() != 0 { t.Fatal("expected 0") }
}

// ── ResultsPtrCollection ──

func TestResultsPtrCollection_Basic(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("hello")
	rpc.Add(&r)
	if rpc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_FirstOrDefault(t *testing.T) {
	rpc := Empty.ResultsPtrCollection()
	if rpc.FirstOrDefault() != nil { t.Fatal("expected nil") }
}

func TestResultsPtrCollection_LastOrDefault(t *testing.T) {
	rpc := Empty.ResultsPtrCollection()
	if rpc.LastOrDefault() != nil { t.Fatal("expected nil") }
}

func TestResultsPtrCollection_Take(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Add(&r)
	taken := rpc.Take(1)
	if taken.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_Skip(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r1 := NewResult.Any("a")
	r2 := NewResult.Any("b")
	rpc.Add(&r1).Add(&r2)
	skipped := rpc.Skip(1)
	if skipped.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_Limit(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Add(&r)
	limited := rpc.Limit(-1)
	if limited.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddSkipOnNil(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	rpc.AddSkipOnNil(nil)
	if rpc.Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_AddNonNilNonError(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	rpc.AddNonNilNonError(nil)
	if rpc.Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_HasError(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	if rpc.HasError() { t.Fatal("expected false") }
}

func TestResultsPtrCollection_AllErrors(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	errs, has := rpc.AllErrors()
	if has || len(errs) != 0 { t.Fatal("unexpected") }
}

func TestResultsPtrCollection_GetErrorsStrings(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	s := rpc.GetErrorsStrings()
	if len(s) != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_Clear(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Add(&r)
	rpc.Clear()
	if rpc.Length() != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_Dispose(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Add(&r)
	rpc.Dispose()
	if rpc.Items != nil { t.Fatal("expected nil") }
}

func TestResultsPtrCollection_GetStrings(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Add(&r)
	s := rpc.GetStrings()
	if len(s) != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_NonPtr(t *testing.T) {
	rpc := ResultsPtrCollection{}
	_ = rpc.NonPtr()
}

func TestResultsPtrCollection_Ptr(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	_ = rpc.Ptr()
}

func TestResultsPtrCollection_AsJsonContractsBinder(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	_ = rpc.AsJsonContractsBinder()
}

func TestResultsPtrCollection_AsJsoner(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	_ = rpc.AsJsoner()
}

func TestResultsPtrCollection_AsJsonParseSelfInjector(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	_ = rpc.AsJsonParseSelfInjector()
}

func TestResultsPtrCollection_Clone(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Add(&r)
	c := rpc.Clone(true)
	if c == nil { t.Fatal("expected non-nil") }
}

func TestResultsPtrCollection_Clone_Nil(t *testing.T) {
	var rpc *ResultsPtrCollection
	if rpc.Clone(true) != nil { t.Fatal("expected nil") }
}

func TestResultsPtrCollection_AddResult(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	rpc.AddResult(NewResult.Any("a"))
	if rpc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_Adds(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Adds(&r)
	if rpc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddAny(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	rpc.AddAny("hello")
	if rpc.Length() != 1 { t.Fatal("expected 1") }
	rpc.AddAny(nil)
	if rpc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddAnyItems(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	rpc.AddAnyItems("a", nil)
	if rpc.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddResultsCollection(t *testing.T) {
	rpc1 := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc1.Add(&r)
	rpc2 := NewResultsPtrCollection.UsingCap(5)
	rpc2.AddResultsCollection(rpc1)
	if rpc2.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_GetPagesSize(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		r := NewResult.Any("a")
		rpc.Add(&r)
	}
	if rpc.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if rpc.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
}

func TestResultsPtrCollection_GetPagedCollection(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		r := NewResult.Any("a")
		rpc.Add(&r)
	}
	pages := rpc.GetPagedCollection(3)
	if len(pages) != 4 { t.Fatal("expected 4") }
}

func TestResultsPtrCollection_GetPagedCollection_Small(t *testing.T) {
	rpc := NewResultsPtrCollection.UsingCap(5)
	r := NewResult.Any("a")
	rpc.Add(&r)
	pages := rpc.GetPagedCollection(5)
	if len(pages) != 1 { t.Fatal("expected 1") }
}
