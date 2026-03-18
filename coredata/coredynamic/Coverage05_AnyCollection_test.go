package coredynamic

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

func TestAnyCollection_NewAndEmpty(t *testing.T) {
	ac := NewAnyCollection(5)
	if ac.Length() != 0 { t.Fatal("expected 0") }
	if !ac.IsEmpty() { t.Fatal("expected empty") }
	if ac.HasAnyItem() { t.Fatal("expected no items") }
	if ac.LastIndex() != -1 { t.Fatal("expected -1") }

	ec := EmptyAnyCollection()
	if ec.Length() != 0 { t.Fatal("expected 0") }
}

func TestAnyCollection_NilReceiver(t *testing.T) {
	var ac *AnyCollection
	if ac.Length() != 0 { t.Fatal("expected 0") }
	if !ac.IsEmpty() { t.Fatal("expected empty") }
}

func TestAnyCollection_AddAndAccess(t *testing.T) {
	ac := NewAnyCollection(5)
	ac.Add("hello")
	ac.Add(42)
	if ac.Length() != 2 { t.Fatal("expected 2") }
	if ac.At(0) != "hello" { t.Fatal("expected hello") }
	if ac.First() != "hello" { t.Fatal("expected hello") }
	if ac.Last() != 42 { t.Fatal("expected 42") }
	if ac.FirstOrDefault() != "hello" { t.Fatal("expected hello") }
	if ac.LastOrDefault() != 42 { t.Fatal("expected 42") }
	if !ac.HasIndex(1) { t.Fatal("expected true") }
	if ac.HasIndex(5) { t.Fatal("expected false") }
}

func TestAnyCollection_AtAsDynamic(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("test")
	d := ac.AtAsDynamic(0)
	if d.IsInvalid() { t.Fatal("expected valid") }
}

func TestAnyCollection_Items(t *testing.T) {
	ac := NewAnyCollection(2)
	if len(ac.Items()) != 0 { t.Fatal("expected empty") }
	ac.Add("a")
	if len(ac.Items()) != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_DynamicItems(t *testing.T) {
	ac := EmptyAnyCollection()
	if len(ac.DynamicItems()) != 0 { t.Fatal("expected empty") }
	ac.Add("x")
	items := ac.DynamicItems()
	if len(items) != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_DynamicCollection(t *testing.T) {
	ac := EmptyAnyCollection()
	dc := ac.DynamicCollection()
	if dc.Length() != 0 { t.Fatal("expected 0") }
	ac.Add("x")
	dc2 := ac.DynamicCollection()
	if dc2.Length() != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_FirstLastOrDefault_Empty(t *testing.T) {
	ac := EmptyAnyCollection()
	if ac.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if ac.LastOrDefault() != nil { t.Fatal("expected nil") }
	if ac.FirstOrDefaultDynamic() != nil { t.Fatal("expected nil") }
	if ac.LastOrDefaultDynamic() != nil { t.Fatal("expected nil") }
}

func TestAnyCollection_FirstLastDynamic(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("a")
	ac.Add("b")
	_ = ac.FirstDynamic()
	_ = ac.LastDynamic()
}

func TestAnyCollection_SkipTakeLimit(t *testing.T) {
	ac := NewAnyCollection(5)
	for i := 0; i < 5; i++ { ac.Add(i) }
	if len(ac.Skip(2)) != 3 { t.Fatal("expected 3") }
	if len(ac.Take(3)) != 3 { t.Fatal("expected 3") }
	if len(ac.Limit(3)) != 3 { t.Fatal("expected 3") }
	_ = ac.SkipDynamic(1)
	_ = ac.TakeDynamic(2)
	_ = ac.LimitDynamic(2)
	sc := ac.SkipCollection(2)
	if sc.Length() != 3 { t.Fatal("expected 3") }
	tc := ac.TakeCollection(3)
	if tc.Length() != 3 { t.Fatal("expected 3") }
	lc := ac.LimitCollection(3)
	if lc.Length() != 3 { t.Fatal("expected 3") }
	slc := ac.SafeLimitCollection(10)
	if slc.Length() != 5 { t.Fatal("expected 5") }
}

func TestAnyCollection_Count(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("x")
	if ac.Count() != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_ListStrings(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("hello")
	ac.Add(42)
	strs := ac.ListStrings(false)
	if len(strs) != 2 { t.Fatal("expected 2") }
	strs2 := ac.ListStringsPtr(true)
	if len(strs2) != 2 { t.Fatal("expected 2") }
}

func TestAnyCollection_RemoveAt(t *testing.T) {
	ac := NewAnyCollection(3)
	ac.Add("a").Add("b").Add("c")
	if !ac.RemoveAt(1) { t.Fatal("expected success") }
	if ac.Length() != 2 { t.Fatal("expected 2") }
	if ac.RemoveAt(10) { t.Fatal("expected failure") }
}

func TestAnyCollection_Loop_Sync(t *testing.T) {
	ac := NewAnyCollection(3)
	ac.Add("a").Add("b").Add("c")
	count := 0
	ac.Loop(false, func(index int, item any) bool {
		count++
		return index == 1
	})
	if count != 2 { t.Fatal("expected 2") }
}

func TestAnyCollection_Loop_Async(t *testing.T) {
	ac := NewAnyCollection(3)
	ac.Add(1).Add(2).Add(3)
	ac.Loop(true, func(index int, item any) bool { return false })
}

func TestAnyCollection_Loop_Empty(t *testing.T) {
	ac := EmptyAnyCollection()
	ac.Loop(false, func(index int, item any) bool { return false })
}

func TestAnyCollection_LoopDynamic_Sync(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("a").Add("b")
	count := 0
	ac.LoopDynamic(false, func(index int, item Dynamic) bool {
		count++
		return index == 0
	})
	if count != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_LoopDynamic_Async(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("a").Add("b")
	ac.LoopDynamic(true, func(index int, item Dynamic) bool { return false })
}

func TestAnyCollection_LoopDynamic_Empty(t *testing.T) {
	ac := EmptyAnyCollection()
	ac.LoopDynamic(false, func(index int, item Dynamic) bool { return false })
}

func TestAnyCollection_AddAny(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.AddAny("hello", true)
	if ac.Length() != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_AddAnyWithTypeValidation(t *testing.T) {
	ac := NewAnyCollection(2)
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(""), "hello")
	if err != nil { t.Fatal("unexpected error") }
	err2 := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(0), "hello")
	if err2 == nil { t.Fatal("expected error") }
}

func TestAnyCollection_AddAnyItemsWithTypeValidation(t *testing.T) {
	ac := NewAnyCollection(5)
	err := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""), "a", "b")
	if err != nil { t.Fatal("unexpected error") }
	err2 := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0), "a")
	if err2 == nil { t.Fatal("expected error") }
	err3 := ac.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(0), "a", "b")
	if err3 == nil { t.Fatal("expected error") }
	err4 := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""))
	if err4 != nil { t.Fatal("unexpected error") }
}

func TestAnyCollection_AddNonNull(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.AddNonNull(nil)
	ac.AddNonNull("x")
	if ac.Length() != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_AddNonNullDynamic(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.AddNonNullDynamic(nil, false)
	ac.AddNonNullDynamic("x", true)
	if ac.Length() != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_AddAnyManyDynamic(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.AddAnyManyDynamic(nil)
	ac.AddAnyManyDynamic("a", "b")
	if ac.Length() != 2 { t.Fatal("expected 2") }
}

func TestAnyCollection_AddAnySliceFromSingleItem(t *testing.T) {
	ac := NewAnyCollection(3)
	ac.AddAnySliceFromSingleItem(nil)
	ac.AddAnySliceFromSingleItem([]string{"a", "b"})
	if ac.Length() != 2 { t.Fatal("expected 2") }
}

func TestAnyCollection_AddMany(t *testing.T) {
	ac := NewAnyCollection(3)
	ac.AddMany(nil)
	ac.AddMany("a", nil, "b")
	if ac.Length() != 2 { t.Fatal("expected 2") }
}

func TestAnyCollection_ReflectSetAt(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("hello")
	var s string
	err := ac.ReflectSetAt(0, &s)
	if err != nil { t.Fatal("unexpected error") }
}

func TestAnyCollection_JsonString(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("a").Add(1)
	s, err := ac.JsonString()
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestAnyCollection_JsonStringMust(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("a")
	s := ac.JsonStringMust()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestAnyCollection_MarshalUnmarshalJSON(t *testing.T) {
	ac := NewAnyCollection(2)
	ac.Add("a").Add("b")
	b, err := ac.MarshalJSON()
	if err != nil { t.Fatal("unexpected error") }
	ac2 := EmptyAnyCollection()
	err2 := ac2.UnmarshalJSON(b)
	if err2 != nil { t.Fatal("unexpected error") }
}

func TestAnyCollection_JsonResultsCollection(t *testing.T) {
	ac := EmptyAnyCollection()
	rc := ac.JsonResultsCollection()
	if rc.Length() != 0 { t.Fatal("expected 0") }
	ac.Add("x")
	rc2 := ac.JsonResultsCollection()
	if rc2.Length() != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_JsonResultsPtrCollection(t *testing.T) {
	ac := EmptyAnyCollection()
	rc := ac.JsonResultsPtrCollection()
	if rc.Length() != 0 { t.Fatal("expected 0") }
	ac.Add("x")
	rc2 := ac.JsonResultsPtrCollection()
	if rc2.Length() != 1 { t.Fatal("expected 1") }
}

func TestAnyCollection_Paging(t *testing.T) {
	ac := NewAnyCollection(10)
	for i := 0; i < 10; i++ { ac.Add(i) }
	if ac.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if ac.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	if ac.GetPagesSize(-1) != 0 { t.Fatal("expected 0") }
	paged := ac.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
	single := ac.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }
	small := NewAnyCollection(1)
	small.Add("x")
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
	if small.GetSinglePageCollection(5, 1).Length() != 1 { t.Fatal("expected 1") }
	_ = ac.GetPagingInfo(3, 1)
}

func TestAnyCollection_JsonMethods(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("x")
	_ = ac.JsonModel()
	_ = ac.JsonModelAny()
	_ = ac.Json()
	_ = ac.JsonPtr()
}

func TestAnyCollection_ParseInjectUsingJson(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("x")
	jr := ac.Json()
	ac2 := EmptyAnyCollection()
	_, err := ac2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected error") }
}

func TestAnyCollection_ParseInjectUsingJsonMust(t *testing.T) {
	ac := NewAnyCollection(1)
	ac.Add("x")
	jr := ac.Json()
	ac2 := EmptyAnyCollection()
	_ = ac2.ParseInjectUsingJsonMust(&jr)
}

func TestAnyCollection_JsonParseSelfInject(t *testing.T) {
	ac := EmptyAnyCollection()
	jr := corejson.NewResult.Any([]any{"a"})
	err := ac.JsonParseSelfInject(&jr)
	if err != nil { t.Fatal("unexpected error") }
}

func TestAnyCollection_Strings(t *testing.T) {
	ac := EmptyAnyCollection()
	if len(ac.Strings()) != 0 { t.Fatal("expected 0") }
	ac.Add("hello")
	strs := ac.Strings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = ac.String()
}
