package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"errors"
	"testing"
	"time"
)

func TestMapResults_Basic(t *testing.T) {
	var nilM *MapResults
	if nilM.Length() != 0 { t.Fatal("expected 0") }
	if nilM.LastIndex() != -1 { t.Fatal("expected -1") }
	if !nilM.IsEmpty() { t.Fatal("expected empty") }
	if nilM.HasAnyItem() { t.Fatal("expected false") }
}

func TestMapResults_AddAndAccess(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("a", corejson.NewResult.Any("hello"))
	if m.Length() != 1 { t.Fatal("expected 1") }
	r := m.GetByKey("a")
	if r == nil { t.Fatal("expected non-nil") }
	if m.GetByKey("missing") != nil { t.Fatal("expected nil") }
}

func TestMapResults_AddMethods(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(5)
	m.AddSkipOnNil("a", nil)
	m.AddSkipOnNil("a", &corejson.Result{Bytes: []byte(`"x"`)})
	m.AddPtr("b", nil)
	m.AddPtr("b", &corejson.Result{Bytes: []byte(`"y"`)})
	err := m.AddAny("c", "hello")
	if err != nil { t.Fatal("unexpected error") }
	err2 := m.AddAny("d", nil)
	if err2 == nil { t.Fatal("expected error") }
	err3 := m.AddAnySkipOnNil("e", nil)
	if err3 != nil { t.Fatal("expected nil") }
	err4 := m.AddAnySkipOnNil("e", "val")
	if err4 != nil { t.Fatal("unexpected error") }
	m.AddAnyNonEmptyNonError("f", nil)
	m.AddAnyNonEmptyNonError("f", "val")
	m.AddAnyNonEmpty("g", nil)
	m.AddAnyNonEmpty("g", "val")
	m.AddNonEmptyNonErrorPtr("h", nil)
	m.AddNonEmptyNonErrorPtr("h", &corejson.Result{Error: errors.New("e")})
	m.AddNonEmptyNonErrorPtr("h", &corejson.Result{Bytes: []byte(`"z"`)})

	m.AddKeyWithResult(corejson.KeyWithResult{Key: "i", corejson.Result: corejson.NewResult.Any("v")})
	m.AddKeyWithResultPtr(nil)
	kr := &corejson.KeyWithResult{Key: "j", corejson.Result: corejson.NewResult.Any("v")}
	m.AddKeyWithResultPtr(kr)
	m.AddKeysWithResultsPtr()
	m.AddKeysWithResultsPtr(kr)
	m.AddKeysWithResults(corejson.KeyWithResult{Key: "k", corejson.Result: corejson.NewResult.Any("v")})

	m.AddKeyAnyInf(corejson.KeyAny{Key: "l", AnyInf: "val"})
	m.AddKeyAnyInfPtr(nil)
	ka := &corejson.KeyAny{Key: "m", AnyInf: "val"}
	m.AddKeyAnyInfPtr(ka)
	m.AddKeyAnyItems(corejson.KeyAny{Key: "n", AnyInf: "val"})
	m.AddKeyAnyItemsPtr(nil)
	m.AddKeyAnyItemsPtr(ka)

	m.AddMapResults(nil)
	sub := corejson.NewMapResults.Empty()
	sub.Add("sub", corejson.NewResult.Any("v"))
	m.AddMapResults(sub)
	m.AddMapAnyItems(nil)
	m.AddMapAnyItems(map[string]any{"o": "val"})
}

func TestMapResults_Errors(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(3)
	m.Add("ok", corejson.NewResult.Any("x"))
	m.Add("err", corejson.Result{Error: errors.New("e1")})
	if !m.HasError() { t.Fatal("expected error") }
	errs, has := m.AllErrors()
	if !has || len(errs) != 1 { t.Fatal("expected 1") }
	strs := m.GetErrorsStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = m.GetErrorsStringsPtr()
	_ = m.GetErrorsAsSingleString()
	_ = m.GetErrorsAsSingle()
}

func TestMapResults_AllKeys(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("b", corejson.NewResult.Any("x"))
	m.Add("a", corejson.NewResult.Any("y"))
	keys := m.AllKeys()
	if len(keys) != 2 { t.Fatal("expected 2") }
	sorted := m.AllKeysSorted()
	if sorted[0] != "a" { t.Fatal("expected a first") }
	vals := m.AllValues()
	if len(vals) != 2 { t.Fatal("expected 2") }
	_ = m.AllResults()
	_ = m.AllResultsCollection()
	empty := corejson.NewMapResults.Empty()
	if len(empty.AllKeys()) != 0 { t.Fatal("expected 0") }
	if len(empty.AllKeysSorted()) != 0 { t.Fatal("expected 0") }
	if len(empty.AllValues()) != 0 { t.Fatal("expected 0") }
}

func TestMapResults_GetStrings(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	strs := m.GetStrings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = m.GetStringsPtr()
	empty := corejson.NewMapResults.Empty()
	if len(empty.GetStrings()) != 0 { t.Fatal("expected 0") }
}

func TestMapResults_Paging(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(10)
	for i := 0; i < 10; i++ {
		m.Add(corejson.Serialize.ToString(i), corejson.NewResult.Any(i))
	}
	if m.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if m.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := m.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
	small := corejson.NewMapResults.UsingCap(1)
	small.Add("a", corejson.NewResult.Any("x"))
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
}

func TestMapResults_ClearDispose(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("x"))
	m.Clear()
	time.Sleep(10 * time.Millisecond)
	if m.Length() != 0 { t.Fatal("expected 0") }
	m.Dispose()
	var nilM *MapResults
	nilM.Clear()
	nilM.Dispose()
}

func TestMapResults_ResultCollection(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	rc := m.ResultCollection()
	if rc.Length() != 1 { t.Fatal("expected 1") }
	empty := corejson.NewMapResults.Empty()
	if empty.ResultCollection().Length() != 0 { t.Fatal("expected 0") }
}

func TestMapResults_JsonMethods(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(1)
	m.Add("a", corejson.NewResult.Any("x"))
	_ = m.JsonModel()
	_ = m.JsonModelAny()
	_ = m.Json()
	_ = m.JsonPtr()
	_ = m.AsJsonContractsBinder()
	_ = m.AsJsoner()
	_ = m.AsJsonParseSelfInjector()
}

func TestMapResults_AddMapResultsUsingCloneOption(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	items := map[string]corejson.Result{"a": corejson.NewResult.Any("x")}
	m.AddMapResultsUsingCloneOption(false, false, items)
	m2 := corejson.NewMapResults.UsingCap(2)
	m2.AddMapResultsUsingCloneOption(true, true, items)
	m3 := corejson.NewMapResults.UsingCap(2)
	m3.AddMapResultsUsingCloneOption(false, false, nil)
}

func TestMapResults_GetNewMapUsingKeys(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.Add("a", corejson.NewResult.Any("x"))
	m.Add("b", corejson.NewResult.Any("y"))
	sub := m.GetNewMapUsingKeys(false, "a")
	if sub.Length() != 1 { t.Fatal("expected 1") }
	empty := m.GetNewMapUsingKeys(false)
	if empty.Length() != 0 { t.Fatal("expected 0") }
}

func TestMapResults_Creators(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyAnyItems(0, corejson.KeyAny{Key: "a", AnyInf: "x"})
	_ = corejson.NewMapResults.UsingKeyAnyItems(5)
	_ = corejson.NewMapResults.UsingMapPlusCap(5, nil)
	_ = corejson.NewMapResults.UsingMapPlusCap(0, map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingMapPlusCapClone(5, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapClone(0, map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(5, nil)
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(0, map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingMap(nil)
	_ = corejson.NewMapResults.UsingMap(map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingMapAnyItems(nil)
	_ = corejson.NewMapResults.UsingMapAnyItems(map[string]any{"a": "x"})
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(5, nil)
	_ = corejson.NewMapResults.UsingKeyWithResults(corejson.KeyWithResult{Key: "a", corejson.Result: corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(5, corejson.KeyWithResult{Key: "a", corejson.Result: corejson.NewResult.Any("x")})
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(5)
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, map[string]corejson.Result{"a": corejson.NewResult.Any("x")})
	_, _ = corejson.NewMapResults.UnmarshalUsingBytes([]byte(`{}`))
}

func TestMapResults_AddJsoner(t *testing.T) {
	m := corejson.NewMapResults.UsingCap(2)
	m.AddJsoner("a", nil)
	_ = corejson.NewMapResults.UsingKeyJsoners()
	_ = corejson.NewMapResults.UsingKeyJsonersPlusCap(5)
}
