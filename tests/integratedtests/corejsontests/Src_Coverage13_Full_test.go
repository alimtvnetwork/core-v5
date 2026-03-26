package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"encoding/json"
	"errors"
	"testing"
)

// ── BytesCollection extended ──

func TestBytesCollection_TakeSkipLimit(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	c.Add([]byte(`"a"`))
	c.Add([]byte(`"b"`))
	c.Add([]byte(`"c"`))

	taken := c.Take(2)
	if taken.Length() != 2 { t.Fatal("expected 2") }

	limited := c.Limit(2)
	if limited.Length() != 2 { t.Fatal("expected 2") }

	limitAll := c.Limit(-1)
	if limitAll.Length() != 3 { t.Fatal("expected 3") }

	skipped := c.Skip(1)
	if skipped.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_AddResult(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	r := corejson.NewResult.Any("hello")
	c.AddResult(r)
	rp := corejson.NewResult.AnyPtr("world")
	c.AddResultPtr(rp)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_GetAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"x"`))
	if c.GetAt(0) == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_JsonResultAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"x"`))
	jr := c.JsonResultAt(0)
	if jr == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"hello"`))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestBytesCollection_AddSerializer(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddSerializer(func() ([]byte, error) {
		return json.Marshal("test")
	})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddSerializerFunc(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.AddSerializerFunc(func() ([]byte, error) {
		return json.Marshal(42)
	})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddSerializers(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	s1 := func() ([]byte, error) { return json.Marshal("a") }
	s2 := func() ([]byte, error) { return json.Marshal("b") }
	c.AddSerializers(s1, s2)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_AddSerializerFunctions(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	fns := []func() ([]byte, error){
		func() ([]byte, error) { return json.Marshal("x") },
	}
	c.AddSerializerFunctions(fns)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b := c.GetAtSafe(0)
	if b == nil { t.Fatal("expected non-nil") }
	b2 := c.GetAtSafe(99)
	if b2 != nil { t.Fatal("expected nil") }
}

func TestBytesCollection_GetAtSafePtr(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b := c.GetAtSafePtr(0)
	if b == nil { t.Fatal("expected non-nil") }
	b2 := c.GetAtSafePtr(99)
	if b2 != nil { t.Fatal("expected nil") }
}

func TestBytesCollection_GetResultAtSafe(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	r := c.GetResultAtSafe(0)
	if r == nil { t.Fatal("expected non-nil") }
	r2 := c.GetResultAtSafe(99)
	if r2 == nil { t.Fatal("expected non-nil empty result") }
}

func TestBytesCollection_GetAtSafeUsingLength(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b := c.GetAtSafeUsingLength(0, c.Length())
	if b == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_Adds(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	c.Adds([]byte(`"a"`), []byte(`"b"`))
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_AddsPtr(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	items := [][]byte{[]byte(`"x"`), []byte(`"y"`)}
	c.AddsPtr(&items)
	if c.Length() != 2 { t.Fatal("expected 2") }
	c.AddsPtr(nil) // should not panic
}

func TestBytesCollection_AddAnyItems(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	err := c.AddAnyItems("a", 42)
	if err != nil { t.Fatal(err) }
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_AddAny(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	err := c.AddAny("test")
	if err != nil { t.Fatal(err) }
}

func TestBytesCollection_AddBytesCollection(t *testing.T) {
	c1 := corejson.NewBytesCollection.UsingCap(2)
	c1.Add([]byte(`"a"`))
	c2 := corejson.NewBytesCollection.UsingCap(2)
	c2.Add([]byte(`"b"`))
	c1.AddBytesCollection(c2)
	if c1.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_ClearDispose(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c.Clear()
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Add([]byte(`"b"`))
	c.Dispose()
	if c.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_Strings(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c.Add([]byte(`"b"`))
	strs := c.Strings()
	if len(strs) != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_StringsPtr(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	sp := c.StringsPtr()
	if sp == nil || len(*sp) != 1 { t.Fatal("unexpected") }
}

func TestBytesCollection_InjectIntoAt(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"hello"`))
	var s string
	err := c.InjectIntoAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestBytesCollection_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c.Add([]byte(`"b"`))
	results := make([]string, 2)
	ptrs := make([]any, 2)
	for i := range results { ptrs[i] = &results[i] }
	errs := c.InjectIntoSameIndex(ptrs)
	hasErr := false
	for _, e := range errs { if e != nil { hasErr = true } }
	if hasErr { t.Fatal("unexpected errors") }
}

func TestBytesCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"x"`))
	c.Add([]byte(`"y"`))
	results := make([]string, 2)
	ptrs := make([]any, 2)
	for i := range results { ptrs[i] = &results[i] }
	errs := c.UnmarshalIntoSameIndex(ptrs)
	hasErr := false
	for _, e := range errs { if e != nil { hasErr = true } }
	if hasErr { t.Fatal("unexpected errors") }
}

func TestBytesCollection_GetPagesSize(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ { c.Add([]byte(`"x"`)) }
	pages := c.GetPagesSize(3)
	if pages < 3 { t.Fatal("expected at least 3 pages") }
}

func TestBytesCollection_GetPagedCollection(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(10)
	for i := 0; i < 5; i++ { c.Add([]byte(`"x"`)) }
	paged := c.GetPagedCollection(1, 2)
	if paged.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_GetSinglePageCollection(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(5)
	for i := 0; i < 5; i++ { c.Add([]byte(`"x"`)) }
	single := c.GetSinglePageCollection(2, 3)
	_ = single
}

func TestBytesCollection_JsonModelAndInterfaces(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_, _ = c.MarshalJSON()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func TestBytesCollection_CloneAndShadow(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	shadow := c.ShadowClone()
	if shadow.Length() != 1 { t.Fatal("expected 1") }
	cloned := c.Clone()
	if cloned.Length() != 1 { t.Fatal("expected 1") }
	clonedPtr := c.ClonePtr()
	if clonedPtr.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_ParseInjectUsingJson(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b, _ := json.Marshal(c)
	c2 := corejson.NewBytesCollection.Empty()
	err := c2.ParseInjectUsingJson(b)
	if err != nil { t.Fatal(err) }
}

func TestBytesCollection_UnmarshalJSON(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b, _ := json.Marshal(c)
	c2 := corejson.NewBytesCollection.Empty()
	err := c2.UnmarshalJSON(b)
	if err != nil { t.Fatal(err) }
}

func TestBytesCollection_JsonParseSelfInject(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b, _ := json.Marshal(c)
	c2 := corejson.NewBytesCollection.Empty()
	err := c2.JsonParseSelfInject(b)
	if err != nil { t.Fatal(err) }
}

func TestBytesCollection_AddMapResults(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	c.AddMapResults(mr)
}

func TestBytesCollection_AddRawMapResults(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	c.AddRawMapResults(m)
}

// ── ResultsCollection extended ──

func TestResultsCollection_TakeSkipLimit(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(5)
	for i := 0; i < 3; i++ {
		r := corejson.NewResult.Any(i)
		c.Items = append(c.Items, r)
	}
	if c.Take(2).Length() != 2 { t.Fatal("expected 2") }
	if c.Limit(2).Length() != 2 { t.Fatal("expected 2") }
	if c.Limit(-1).Length() != 3 { t.Fatal("expected 3") }
	if c.Skip(1).Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_AddSkipOnNil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSkipOnNil(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
	r := corejson.NewResult.AnyPtr("x")
	c.AddSkipOnNil(r)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddNonNilNonError(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddNonNilNonError(nil)
	errR := corejson.NewResult.ErrorPtr(errors.New("fail"))
	c.AddNonNilNonError(errR)
	goodR := corejson.NewResult.AnyPtr("ok")
	c.AddNonNilNonError(goodR)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AllErrors(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Error(errors.New("e1")))
	c.Items = append(c.Items, corejson.NewResult.Any("ok"))
	errs, hasErr := c.AllErrors()
	if !hasErr || len(errs) != 1 { t.Fatal("unexpected") }
}

func TestResultsCollection_GetErrorsStrings(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Error(errors.New("e1")))
	strs := c.GetErrorsStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	sp := c.GetErrorsStringsPtr()
	if sp == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_GetErrorsAsSingleString(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Error(errors.New("e1")))
	s := c.GetErrorsAsSingleString()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResultsCollection_GetErrorsAsSingle(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Error(errors.New("e1")))
	e := c.GetErrorsAsSingle()
	if e == nil { t.Fatal("expected error") }
}

func TestResultsCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestResultsCollection_InjectIntoAt(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("hello"))
	var s string
	err := c.InjectIntoAt(0, &s)
	if err != nil { t.Fatal(err) }
}

func TestResultsCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("x"))
	r := c.GetAtSafe(0)
	if r == nil { t.Fatal("expected non-nil") }
	r2 := c.GetAtSafe(99)
	if r2 == nil { t.Fatal("expected non-nil empty") }
}

func TestResultsCollection_GetAtSafeUsingLength(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("x"))
	r := c.GetAtSafeUsingLength(0, c.Length())
	if r == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_AddPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddPtr(nil)
	r := corejson.NewResult.AnyPtr("x")
	c.AddPtr(r)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_Adds(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r1 := corejson.NewResult.Any("a")
	r2 := corejson.NewResult.Any("b")
	c.Adds(r1, r2)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_AddSerializer(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializer(func() ([]byte, error) { return json.Marshal("x") })
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddSerializers(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializers(
		func() ([]byte, error) { return json.Marshal("a") },
		func() ([]byte, error) { return json.Marshal("b") },
	)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_AddSerializerFunc(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializerFunc(func() ([]byte, error) { return json.Marshal("x") })
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddSerializerFunctions(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	fns := []func() ([]byte, error){
		func() ([]byte, error) { return json.Marshal("x") },
	}
	c.AddSerializerFunctions(fns)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddsPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	items := []*corejson.Result{corejson.NewResult.AnyPtr("a"), nil, corejson.NewResult.AnyPtr("b")}
	c.AddsPtr(items)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_AddAnyItems(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddAnyItems("a", 42)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_AddAnyItemsSlice(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddAnyItemsSlice([]any{"a", 42})
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_AddResultsCollection(t *testing.T) {
	c1 := corejson.NewResultsCollection.Empty()
	c1.Items = append(c1.Items, corejson.NewResult.Any("a"))
	c2 := corejson.NewResultsCollection.Empty()
	c2.AddResultsCollection(c1)
	if c2.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_AddNonNilItemsPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	items := []*corejson.Result{corejson.NewResult.AnyPtr("a"), nil}
	c.AddNonNilItemsPtr(items)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_NonPtrAndPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("x"))
	_ = c.NonPtr()
	_ = c.Ptr()
}

func TestResultsCollection_GetStringsPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("x"))
	sp := c.GetStringsPtr()
	if sp == nil { t.Fatal("expected non-nil") }
}

func TestResultsCollection_GetPagesSize(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Items = append(c.Items, corejson.NewResult.Any(i))
	}
	pages := c.GetPagesSize(3)
	if pages < 3 { t.Fatal("expected >= 3") }
}

func TestResultsCollection_GetPagedCollection(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		c.Items = append(c.Items, corejson.NewResult.Any(i))
	}
	paged := c.GetPagedCollection(1, 2)
	if paged.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsCollection_GetSinglePageCollection(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		c.Items = append(c.Items, corejson.NewResult.Any(i))
	}
	single := c.GetSinglePageCollection(2, 3)
	_ = single
}

func TestResultsCollection_JsonInterfaces(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("a"))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func TestResultsCollection_CloneAndShadow(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("a"))
	shadow := c.ShadowClone()
	if shadow.Length() != 1 { t.Fatal("expected 1") }
	cloned := c.Clone()
	if cloned.Length() != 1 { t.Fatal("expected 1") }
	clonedPtr := c.ClonePtr()
	if clonedPtr.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsCollection_ParseInjectAndJsonParseSelfInject(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("a"))
	b, _ := json.Marshal(c)
	c2 := corejson.NewResultsCollection.Empty()
	_ = c2.ParseInjectUsingJson(b)
	c3 := corejson.NewResultsCollection.Empty()
	_ = c3.JsonParseSelfInject(b)
}

func TestResultsCollection_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("a"))
	c.Items = append(c.Items, corejson.NewResult.Any("b"))
	results := make([]string, 2)
	ptrs := make([]any, 2)
	for i := range results { ptrs[i] = &results[i] }
	errs := c.InjectIntoSameIndex(ptrs)
	hasErr := false
	for _, e := range errs { if e != nil { hasErr = true } }
	if hasErr { t.Fatal("unexpected errors") }
}

func TestResultsCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("x"))
	c.Items = append(c.Items, corejson.NewResult.Any("y"))
	results := make([]string, 2)
	ptrs := make([]any, 2)
	for i := range results { ptrs[i] = &results[i] }
	errs := c.UnmarshalIntoSameIndex(ptrs)
	hasErr := false
	for _, e := range errs { if e != nil { hasErr = true } }
	if hasErr { t.Fatal("unexpected errors") }
}

func TestResultsCollection_AddMapResults(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	c.AddMapResults(mr)
}

func TestResultsCollection_AddRawMapResults(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	c.AddRawMapResults(m)
}

// ── ResultsPtrCollection extended ──

func TestResultsPtrCollection_TakeSkipLimit(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingCap(5)
	for i := 0; i < 3; i++ {
		c.Items = append(c.Items, corejson.NewResult.AnyPtr(i))
	}
	if c.Take(2).Length() != 2 { t.Fatal("expected 2") }
	if c.Limit(2).Length() != 2 { t.Fatal("expected 2") }
	if c.Limit(-1).Length() != 3 { t.Fatal("expected 3") }
	if c.Skip(1).Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsPtrCollection_AddSkipOnNil(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Empty()
	c.AddSkipOnNil(nil)
	if c.Length() != 0 { t.Fatal("expected 0") }
	r := corejson.NewResult.AnyPtr("x")
	c.AddSkipOnNil(r)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddNonNilNonError(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Empty()
	c.AddNonNilNonError(nil)
	errR := corejson.NewResult.ErrorPtr(errors.New("fail"))
	c.AddNonNilNonError(errR)
	goodR := corejson.NewResult.AnyPtr("ok")
	c.AddNonNilNonError(goodR)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_GetAt(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("x"))
	if c.GetAt(0) == nil { t.Fatal("expected non-nil") }
}

func TestResultsPtrCollection_HasError(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.ErrorPtr(errors.New("e")))
	if !c.HasError() { t.Fatal("expected true") }
}

func TestResultsPtrCollection_AllErrors(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.ErrorPtr(errors.New("e1")))
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("ok"))
	errs, hasErr := c.AllErrors()
	if !hasErr || len(errs) != 1 { t.Fatal("unexpected") }
}

func TestResultsPtrCollection_GetErrorsStrings(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.ErrorPtr(errors.New("e1")))
	strs := c.GetErrorsStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	sp := c.GetErrorsStringsPtr()
	_ = sp
}

func TestResultsPtrCollection_GetErrorsAsSingleString(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.ErrorPtr(errors.New("e")))
	s := c.GetErrorsAsSingleString()
	if s == "" { t.Fatal("expected non-empty") }
	e := c.GetErrorsAsSingle()
	if e == nil { t.Fatal("expected error") }
}

func TestResultsPtrCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("hi"))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hi" { t.Fatal("unexpected") }
}

func TestResultsPtrCollection_InjectIntoAt(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("hi"))
	var s string
	err := c.InjectIntoAt(0, &s)
	if err != nil { t.Fatal(err) }
}

func TestResultsPtrCollection_GetAtSafe(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("x"))
	r := c.GetAtSafe(0)
	if r == nil { t.Fatal("expected non-nil") }
	r2 := c.GetAtSafe(99)
	if r2 == nil { t.Fatal("expected non-nil empty") }
}

func TestResultsPtrCollection_GetAtSafeUsingLength(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("x"))
	r := c.GetAtSafeUsingLength(0, c.Length())
	if r == nil { t.Fatal("expected non-nil") }
}

func TestResultsPtrCollection_Add(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddSerializer(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.AddSerializer(func() ([]byte, error) { return json.Marshal("x") })
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddSerializers(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.AddSerializers(
		func() ([]byte, error) { return json.Marshal("a") },
		func() ([]byte, error) { return json.Marshal("b") },
	)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsPtrCollection_AddSerializerFunc(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.AddSerializerFunc(func() ([]byte, error) { return json.Marshal("x") })
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddSerializerFunctions(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	fns := []func() ([]byte, error){
		func() ([]byte, error) { return json.Marshal("x") },
	}
	c.AddSerializerFunctions(fns)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddResult(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	r := corejson.NewResult.Any("x")
	c.AddResult(r)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_Adds(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Adds(corejson.NewResult.AnyPtr("a"), nil, corejson.NewResult.AnyPtr("b"))
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsPtrCollection_AddAnyItems(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.AddAnyItems("a", 42)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestResultsPtrCollection_AddResultsCollection(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	rc := corejson.NewResultsCollection.Empty()
	rc.Items = append(rc.Items, corejson.NewResult.Any("a"))
	c.AddResultsCollection(rc)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddNonNilItems(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.AddNonNilItems(corejson.NewResult.AnyPtr("a"), nil)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_AddNonNilItemsPtr(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	items := []*corejson.Result{corejson.NewResult.AnyPtr("a"), nil}
	c.AddNonNilItemsPtr(items)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestResultsPtrCollection_InjectIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("a"))
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("b"))
	results := make([]string, 2)
	ptrs := make([]any, 2)
	for i := range results { ptrs[i] = &results[i] }
	errs := c.InjectIntoSameIndex(ptrs)
	hasErr := false
	for _, e := range errs { if e != nil { hasErr = true } }
	if hasErr { t.Fatal("unexpected errors") }
}

func TestResultsPtrCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("x"))
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("y"))
	results := make([]string, 2)
	ptrs := make([]any, 2)
	for i := range results { ptrs[i] = &results[i] }
	errs := c.UnmarshalIntoSameIndex(ptrs)
	hasErr := false
	for _, e := range errs { if e != nil { hasErr = true } }
	if hasErr { t.Fatal("unexpected errors") }
}

// ── MapResults extended ──

func TestMapResults_GetErrorsStrings(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k1"] = corejson.NewResult.Error(errors.New("e1"))
	mr.Items["k2"] = corejson.NewResult.Any("ok")
	strs := mr.GetErrorsStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	sp := mr.GetErrorsStringsPtr()
	_ = sp
}

func TestMapResults_GetErrorsAsSingleString(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k1"] = corejson.NewResult.Error(errors.New("e1"))
	s := mr.GetErrorsAsSingleString()
	if s == "" { t.Fatal("expected non-empty") }
	e := mr.GetErrorsAsSingle()
	if e == nil { t.Fatal("expected error") }
}

func TestMapResults_Unmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["name"] = corejson.NewResult.Any("john")
	var s string
	err := mr.Unmarshal("name", &s)
	if err != nil || s != "john" { t.Fatal("unexpected") }
}

func TestMapResults_Deserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["val"] = corejson.NewResult.Any(42)
	var v int
	err := mr.Deserialize("val", &v)
	if err != nil || v != 42 { t.Fatal("unexpected") }
}

func TestMapResults_SafeUnmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["val"] = corejson.NewResult.Any("hello")
	var s string
	mr.SafeUnmarshal("val", &s)
	if s != "hello" { t.Fatal("unexpected") }
}

func TestMapResults_SafeDeserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["val"] = corejson.NewResult.Any("hello")
	var s string
	mr.SafeDeserialize("val", &s)
}

func TestMapResults_UnmarshalMany(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["a"] = corejson.NewResult.Any("x")
	mr.Items["b"] = corejson.NewResult.Any("y")
	var a, b string
	m := map[string]any{"a": &a, "b": &b}
	err := mr.UnmarshalMany(m)
	if err != nil { t.Fatal(err) }
}

func TestMapResults_UnmarshalManySafe(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["a"] = corejson.NewResult.Any("x")
	var a string
	m := map[string]any{"a": &a}
	mr.UnmarshalManySafe(m)
}

func TestMapResults_InjectIntoAt(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["key"] = corejson.NewResult.Any("val")
	var s string
	err := mr.InjectIntoAt("key", &s)
	if err != nil { t.Fatal(err) }
}

func TestMapResults_AddAnySkipOnNil(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnySkipOnNil("k1", nil)
	mr.AddAnySkipOnNil("k2", "hello")
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddAnyNonEmptyNonError(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmptyNonError("k1", "hello")
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddKeyWithResult(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	kwr := corejson.KeyWithResult{Key: "k1", corejson.Result: corejson.NewResult.Any("v")}
	mr.AddKeyWithResult(kwr)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddKeyWithResultPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	kwr := &corejson.KeyWithResult{Key: "k1", corejson.Result: corejson.NewResult.Any("v")}
	mr.AddKeyWithResultPtr(kwr)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddKeysWithResultsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	kwrs := []*corejson.KeyWithResult{{Key: "k1", corejson.Result: corejson.NewResult.Any("v")}}
	mr.AddKeysWithResultsPtr(kwrs)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddKeysWithResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResults(
		corejson.KeyWithResult{Key: "k1", corejson.Result: corejson.NewResult.Any("a")},
		corejson.KeyWithResult{Key: "k2", corejson.Result: corejson.NewResult.Any("b")},
	)
	if mr.Length() != 2 { t.Fatal("expected 2") }
}

func TestMapResults_AddKeyAnyInf(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInf(corejson.KeyAny{Key: "k", Value: "v"})
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddKeyAnyInfPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInfPtr(&corejson.KeyAny{Key: "k", Value: "v"})
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddKeyAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItems(
		corejson.KeyAny{Key: "k1", Value: "v1"},
		corejson.KeyAny{Key: "k2", Value: "v2"},
	)
	if mr.Length() != 2 { t.Fatal("expected 2") }
}

func TestMapResults_AddKeyAnyItemsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	items := []corejson.KeyAny{{Key: "k", Value: "v"}}
	mr.AddKeyAnyItemsPtr(&items)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddNonEmptyNonErrorPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.NewResult.AnyPtr("x")
	mr.AddNonEmptyNonErrorPtr("k", r)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddMapResults(t *testing.T) {
	mr1 := corejson.NewMapResults.Empty()
	mr1.Items["k1"] = corejson.NewResult.Any("v1")
	mr2 := corejson.NewMapResults.Empty()
	mr2.AddMapResults(mr1)
	if mr2.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddMapAnyItems(map[string]any{"k": "v"})
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AllKeys(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["b"] = corejson.NewResult.Any("v")
	mr.Items["a"] = corejson.NewResult.Any("v")
	keys := mr.AllKeys()
	if len(keys) != 2 { t.Fatal("expected 2") }
}

func TestMapResults_AllKeysSorted(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["b"] = corejson.NewResult.Any("v")
	mr.Items["a"] = corejson.NewResult.Any("v")
	keys := mr.AllKeysSorted()
	if len(keys) != 2 || keys[0] != "a" { t.Fatal("unexpected") }
}

func TestMapResults_AllValues(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	vals := mr.AllValues()
	if len(vals) != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AllResultsCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	rc := mr.AllResultsCollection()
	if rc.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AllResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	results := mr.AllResults()
	if len(results) != 1 { t.Fatal("expected 1") }
}

func TestMapResults_GetStrings(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("hello")
	strs := mr.GetStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	sp := mr.GetStringsPtr()
	_ = sp
}

func TestMapResults_ClearDispose(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	mr.Clear()
	if mr.Length() != 0 { t.Fatal("expected 0") }
	mr.Items["k2"] = corejson.NewResult.Any("v2")
	mr.Dispose()
}

func TestMapResults_ResultCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	rc := mr.ResultCollection()
	if rc == nil { t.Fatal("expected non-nil") }
}

func TestMapResults_JsonInterfaces(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	_ = mr.JsonModel()
	_ = mr.JsonModelAny()
	_ = mr.Json()
	_ = mr.JsonPtr()
	_ = mr.AsJsonContractsBinder()
	_ = mr.AsJsoner()
	_ = mr.AsJsonParseSelfInjector()
}

func TestMapResults_ParseInjectAndJsonParseSelfInject(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	b, _ := json.Marshal(mr)
	mr2 := corejson.NewMapResults.Empty()
	_ = mr2.ParseInjectUsingJson(b)
	mr3 := corejson.NewMapResults.Empty()
	_ = mr3.JsonParseSelfInject(b)
}

func TestMapResults_GetPagesSize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 10; i++ {
		mr.Items[string(rune('a'+i))] = corejson.NewResult.Any(i)
	}
	pages := mr.GetPagesSize(3)
	if pages < 3 { t.Fatal("expected >= 3") }
}

func TestMapResults_GetPagedCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ {
		mr.Items[string(rune('a'+i))] = corejson.NewResult.Any(i)
	}
	paged := mr.GetPagedCollection(1, 2)
	_ = paged
}

func TestMapResults_GetSinglePageCollection(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ {
		mr.Items[string(rune('a'+i))] = corejson.NewResult.Any(i)
	}
	single := mr.GetSinglePageCollection(2, 3)
	_ = single
}

func TestMapResults_GetNewMapUsingKeys(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["a"] = corejson.NewResult.Any("va")
	mr.Items["b"] = corejson.NewResult.Any("vb")
	newMr := mr.GetNewMapUsingKeys("a")
	if newMr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddMapResultsUsingCloneOption(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	source := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	mr.AddMapResultsUsingCloneOption(true, false, source)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestMapResults_AddJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	type simple struct{ X int }
	s := simple{X: 1}
	mr.AddJsoner("key", &simpleJsoner{val: s})
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

// ── Result extended ──

func TestResult_FieldsNames(t *testing.T) {
	r := corejson.NewResult.Any(map[string]int{"a": 1, "b": 2})
	names, err := r.FieldsNames()
	_ = names
	_ = err
}

func TestResult_HandleErrorWithMsg(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("test error"))
	r.HandleErrorWithMsg("context message")
}

func TestResult_DeserializeMust(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	var s string
	r.DeserializeMust(&s)
	if s != "hello" { t.Fatal("unexpected") }
}

func TestResult_UnmarshalMust(t *testing.T) {
	r := corejson.NewResult.Any(42)
	var v int
	r.UnmarshalMust(&v)
	if v != 42 { t.Fatal("unexpected") }
}

func TestResult_ParseInjectUsingJsonMust(t *testing.T) {
	r1 := corejson.NewResult.Any("hello")
	b, _ := json.Marshal(r1)
	r2 := corejson.NewResult.EmptyPtr()
	r2.ParseInjectUsingJsonMust(b)
}

// ── Serializer extended ──

func TestSerializer_FromStringer(t *testing.T) {
	r := corejson.Serialize.FromStringer(simpleStringer{"test"})
	if r.HasError() { t.Fatal(r.Error) }
}

func TestSerializer_Marshal_Func(t *testing.T) {
	b, err := corejson.Serialize.Marshal("hello")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestSerializer_ApplyMust_Valid(t *testing.T) {
	r := corejson.Serialize.ApplyMust("hello")
	if r.HasError() { t.Fatal("unexpected") }
}

func TestSerializer_ToBytesMust_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("hello")
	if len(b) == 0 { t.Fatal("expected bytes") }
}

func TestSerializer_ToStringMust_Valid(t *testing.T) {
	s := corejson.Serialize.ToStringMust("hello")
	if s == "" { t.Fatal("expected string") }
}

// ── NewResultCreator extended ──

func TestNewResultCreator_DeserializeUsingBytes(t *testing.T) {
	r1 := corejson.NewResult.Any("hello")
	b, _ := json.Marshal(r1)
	r2 := corejson.NewResult.DeserializeUsingBytes(b)
	_ = r2
}

func TestNewResultCreator_DeserializeUsingResult(t *testing.T) {
	r1 := corejson.NewResult.AnyPtr(corejson.NewResult.Any("hello"))
	r2 := corejson.NewResult.DeserializeUsingResult(r1)
	_ = r2
}

func TestNewResultCreator_Various(t *testing.T) {
	_ = corejson.NewResult.UsingBytes([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesType([]byte(`"x"`), "string")
	_ = corejson.NewResult.UsingBytesTypePtr([]byte(`"x"`), "string")
	_ = corejson.NewResult.UsingTypeBytesPtr("string", []byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtr(nil)
	_ = corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtrErrPtr(nil, nil, "t")
	_ = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr(nil, nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "t")
	s := "test"
	_ = corejson.NewResult.PtrUsingStringPtr(&s, "t")
	_ = corejson.NewResult.PtrUsingStringPtr(nil, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "t")
	_ = corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = corejson.NewResult.UsingTypePlusString("t", `"x"`)
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", &s)
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", nil)
	_ = corejson.NewResult.UsingStringWithType(`"x"`, "t")
	_ = corejson.NewResult.UsingString(`"x"`)
	_ = corejson.NewResult.UsingStringPtr(&s)
	_ = corejson.NewResult.UsingStringPtr(nil)
	_ = corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.NonPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Create([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.CastingAny("x")
	_ = corejson.NewResult.Error(errors.New("e"))
	_ = corejson.NewResult.ErrorPtr(errors.New("e"))
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
	_ = corejson.NewResult.TypeName("t")
	_ = corejson.NewResult.TypeNameBytes("t")
	_ = corejson.NewResult.Many("a", "b")
	_ = corejson.NewResult.Serialize("x")
	_ = corejson.NewResult.Marshal("x")
}

// ── NewBytesCollectionCreator ──

func TestNewBytesCollectionCreator_Deserialize(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b, _ := json.Marshal(c)
	c2, err := corejson.NewBytesCollection.DeserializeUsingBytes(b)
	if err != nil || c2 == nil { t.Fatal("unexpected") }
}

func TestNewBytesCollectionCreator_UnmarshalUsingBytes(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	b, _ := json.Marshal(c)
	c2, err := corejson.NewBytesCollection.UnmarshalUsingBytes(b)
	if err != nil || c2 == nil { t.Fatal("unexpected") }
}

func TestNewBytesCollectionCreator_DeserializeUsingResult(t *testing.T) {
	c := corejson.NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	r := corejson.NewResult.AnyPtr(c)
	c2, err := corejson.NewBytesCollection.DeserializeUsingResult(r)
	if err != nil || c2 == nil { t.Fatal("unexpected") }
}

func TestNewBytesCollectionCreator_AnyItems(t *testing.T) {
	c, err := corejson.NewBytesCollection.AnyItems("a", "b")
	if err != nil || c.Length() != 2 { t.Fatal("unexpected") }
}

func TestNewBytesCollectionCreator_Serializers(t *testing.T) {
	c := corejson.NewBytesCollection.Serializers(
		func() ([]byte, error) { return json.Marshal("x") },
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── NewResultsCollectionCreator ──

func TestNewResultsCollectionCreator_Various(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty()
	_ = corejson.NewResultsCollection.Default()
	_ = corejson.NewResultsCollection.UsingCap(5)
}

func TestNewResultsCollectionCreator_DeserializeUsingBytes(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Items = append(c.Items, corejson.NewResult.Any("a"))
	b, _ := json.Marshal(c)
	c2, err := corejson.NewResultsCollection.DeserializeUsingBytes(b)
	if err != nil || c2 == nil { t.Fatal("unexpected") }
}

func TestNewResultsCollectionCreator_AnyItems(t *testing.T) {
	c := corejson.NewResultsCollection.AnyItems("a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewResultsCollectionCreator_AnyItemsPlusCap(t *testing.T) {
	c := corejson.NewResultsCollection.AnyItemsPlusCap(5, "a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewResultsCollectionCreator_Serializers(t *testing.T) {
	c := corejson.NewResultsCollection.Serializers(
		func() ([]byte, error) { return json.Marshal("x") },
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewResultsCollectionCreator_SerializerFunctions(t *testing.T) {
	fns := []func() ([]byte, error){
		func() ([]byte, error) { return json.Marshal("x") },
	}
	c := corejson.NewResultsCollection.SerializerFunctions(fns)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── NewResultsPtrCollectionCreator ──

func TestNewResultsPtrCollectionCreator_Various(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Empty()
	_ = corejson.NewResultsPtrCollection.Default()
	_ = corejson.NewResultsPtrCollection.UsingCap(5)
}

func TestNewResultsPtrCollectionCreator_DeserializeUsingBytes(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Items = append(c.Items, corejson.NewResult.AnyPtr("a"))
	b, _ := json.Marshal(c)
	c2, err := corejson.NewResultsPtrCollection.DeserializeUsingBytes(b)
	if err != nil || c2 == nil { t.Fatal("unexpected") }
}

func TestNewResultsPtrCollectionCreator_AnyItems(t *testing.T) {
	c := corejson.NewResultsPtrCollection.AnyItems("a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewResultsPtrCollectionCreator_AnyItemsPlusCap(t *testing.T) {
	c := corejson.NewResultsPtrCollection.AnyItemsPlusCap(5, "a", "b")
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewResultsPtrCollectionCreator_UsingResults(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingResults(corejson.NewResult.AnyPtr("a"))
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewResultsPtrCollectionCreator_UsingResultsPlusCap(t *testing.T) {
	c := corejson.NewResultsPtrCollection.UsingResultsPlusCap(5, corejson.NewResult.AnyPtr("a"))
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewResultsPtrCollectionCreator_Serializers(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Serializers(
		func() ([]byte, error) { return json.Marshal("x") },
	)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── NewMapResultsCreator ──

func TestNewMapResultsCreator_Various(t *testing.T) {
	_ = corejson.NewMapResults.Empty()
	_ = corejson.NewMapResults.UsingCap(5)
}

func TestNewMapResultsCreator_DeserializeUsingBytes(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	b, _ := json.Marshal(mr)
	mr2, err := corejson.NewMapResults.DeserializeUsingBytes(b)
	if err != nil || mr2 == nil { t.Fatal("unexpected") }
}

func TestNewMapResultsCreator_DeserializeUsingResult(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Items["k"] = corejson.NewResult.Any("v")
	r := corejson.NewResult.AnyPtr(mr)
	mr2, err := corejson.NewMapResults.DeserializeUsingResult(r)
	if err != nil || mr2 == nil { t.Fatal("unexpected") }
}

func TestNewMapResultsCreator_UsingKeyAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyAnyItems(0, corejson.KeyAny{Key: "k", Value: "v"})
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewMapResultsCreator_UsingMapOptions(t *testing.T) {
	source := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	mr := corejson.NewMapResults.UsingMapOptions(true, false, 5, source)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewMapResultsCreator_UsingMapPlusCap(t *testing.T) {
	source := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	mr := corejson.NewMapResults.UsingMapPlusCap(5, source)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewMapResultsCreator_UsingMap(t *testing.T) {
	source := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	mr := corejson.NewMapResults.UsingMap(source)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewMapResultsCreator_UsingMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewMapResultsCreator_UsingKeyWithResults(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyWithResults(
		corejson.KeyWithResult{Key: "k", corejson.Result: corejson.NewResult.Any("v")},
	)
	if mr.Length() != 1 { t.Fatal("expected 1") }
}

// ── Helper types for tests ──

type simpleJsoner struct{ val any }
func (s *simpleJsoner) Json() *corejson.Result { return corejson.NewResult.AnyPtr(s.val) }

type simpleStringer struct{ s string }
func (s simpleStringer) String() string { return s.s }
