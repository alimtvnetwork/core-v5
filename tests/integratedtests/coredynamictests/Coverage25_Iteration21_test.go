package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — constructors, accessors, navigation
// ══════════════════════════════════════════════════════════════════════════════

func Test_I21_AnyCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	actual := args.Map{"empty": ac.IsEmpty(), "len": ac.Length(), "count": ac.Count()}
	expected := args.Map{"empty": true, "len": 0, "count": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Empty", actual)
}

func Test_I21_AnyCollection_New(t *testing.T) {
	ac := coredynamic.NewAnyCollection(5)
	actual := args.Map{"empty": ac.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- New", actual)
}

func Test_I21_AnyCollection_Add(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	actual := args.Map{"len": ac.Length(), "hasAny": ac.HasAnyItem()}
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Add", actual)
}

func Test_I21_AnyCollection_AddAny(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("x", true)
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAny", actual)
}

func Test_I21_AnyCollection_AddMany(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddMany("a", "b", nil, "c")
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 3} // nil skipped
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddMany", actual)
}

func Test_I21_AnyCollection_AddMany_Nil(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddMany()
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddMany nil", actual)
}

func Test_I21_AnyCollection_AddNonNull(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddNonNull("a").AddNonNull(nil)
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddNonNull", actual)
}

func Test_I21_AnyCollection_AddNonNullDynamic(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddNonNullDynamic("a", true).AddNonNullDynamic(nil, false)
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddNonNullDynamic", actual)
}

func Test_I21_AnyCollection_AddAnyManyDynamic(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnyManyDynamic("a", "b")
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAnyManyDynamic", actual)
}

func Test_I21_AnyCollection_AddAnyManyDynamic_Nil(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnyManyDynamic()
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddAnyManyDynamic nil", actual)
}

func Test_I21_AnyCollection_At(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	actual := args.Map{"val": ac.At(1)}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- At", actual)
}

func Test_I21_AnyCollection_AtAsDynamic(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true)
	d := ac.AtAsDynamic(0)
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AtAsDynamic", actual)
}

func Test_I21_AnyCollection_Items(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	actual := args.Map{"len": len(ac.Items())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Items", actual)
}

func Test_I21_AnyCollection_Items_Nil(t *testing.T) {
	var ac *coredynamic.AnyCollection
	items := ac.Items()
	actual := args.Map{"nil": items == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- Items nil", actual)
}

func Test_I21_AnyCollection_DynamicItems(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	di := ac.DynamicItems()
	actual := args.Map{"len": len(di)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- DynamicItems", actual)
}

func Test_I21_AnyCollection_DynamicCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true)
	dc := ac.DynamicCollection()
	actual := args.Map{"notNil": dc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- DynamicCollection", actual)
}

func Test_I21_AnyCollection_First(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	actual := args.Map{"first": ac.First(), "last": ac.Last()}
	expected := args.Map{"first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- First/Last", actual)
}

func Test_I21_AnyCollection_FirstOrDefault_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	actual := args.Map{"nil": ac.FirstOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- FirstOrDefault empty", actual)
}

func Test_I21_AnyCollection_LastOrDefault_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	actual := args.Map{"nil": ac.LastOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- LastOrDefault empty", actual)
}

func Test_I21_AnyCollection_FirstOrDefault_HasItem(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("x")
	actual := args.Map{"val": ac.FirstOrDefault()}
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- FirstOrDefault has item", actual)
}

func Test_I21_AnyCollection_Skip(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	skipped := ac.Skip(1)
	actual := args.Map{"len": len(skipped)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Skip", actual)
}

func Test_I21_AnyCollection_SkipCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	sc := ac.SkipCollection(2)
	actual := args.Map{"len": sc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- SkipCollection", actual)
}

func Test_I21_AnyCollection_Take(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	taken := ac.Take(2)
	actual := args.Map{"len": len(taken)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Take", actual)
}

func Test_I21_AnyCollection_TakeCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	tc := ac.TakeCollection(2)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- TakeCollection", actual)
}

func Test_I21_AnyCollection_LimitCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	lc := ac.LimitCollection(2)
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LimitCollection", actual)
}

func Test_I21_AnyCollection_SafeLimitCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	lc := ac.SafeLimitCollection(10)
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- SafeLimitCollection", actual)
}

func Test_I21_AnyCollection_SafeLimitCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	lc := ac.SafeLimitCollection(10)
	actual := args.Map{"empty": lc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- SafeLimitCollection empty", actual)
}

func Test_I21_AnyCollection_LastIndex(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	actual := args.Map{"idx": ac.LastIndex()}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LastIndex", actual)
}

func Test_I21_AnyCollection_HasIndex(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	actual := args.Map{"has0": ac.HasIndex(0), "has1": ac.HasIndex(1)}
	expected := args.Map{"has0": true, "has1": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- HasIndex", actual)
}

func Test_I21_AnyCollection_RemoveAt(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	ok := ac.RemoveAt(1)
	actual := args.Map{"ok": ok, "len": ac.Length()}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- RemoveAt", actual)
}

func Test_I21_AnyCollection_RemoveAt_Invalid(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ok := ac.RemoveAt(5)
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns error -- RemoveAt invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — Loop, LoopDynamic
// ══════════════════════════════════════════════════════════════════════════════

func Test_I21_AnyCollection_Loop_Sync(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop sync", actual)
}

func Test_I21_AnyCollection_Loop_Break(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return i == 0
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop break", actual)
}

func Test_I21_AnyCollection_Loop_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Loop empty", actual)
}

func Test_I21_AnyCollection_Loop_Async(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	ac.Loop(true, func(i int, item any) bool {
		return false
	})
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop async", actual)
}

func Test_I21_AnyCollection_LoopDynamic_Sync(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic sync", actual)
}

func Test_I21_AnyCollection_LoopDynamic_Break(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return true
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic break", actual)
}

func Test_I21_AnyCollection_LoopDynamic_Async(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	ac.LoopDynamic(true, func(i int, item coredynamic.Dynamic) bool {
		return false
	})
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic async", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — Type validation
// ══════════════════════════════════════════════════════════════════════════════

func Test_I21_AnyCollection_AddAnyWithTypeValidation_Match(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(""), "hello")
	actual := args.Map{"noErr": err == nil, "len": ac.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyWithTypeValidation match", actual)
}

func Test_I21_AnyCollection_AddAnyWithTypeValidation_Mismatch(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(""), 42)
	actual := args.Map{"hasErr": err != nil, "len": ac.Length()}
	expected := args.Map{"hasErr": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyWithTypeValidation mismatch", actual)
}

func Test_I21_AnyCollection_AddAnyItemsWithTypeValidation_StopOnErr(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyItemsWithTypeValidation(
		false, false,
		reflect.TypeOf(""),
		"a", 42, "c",
	)
	actual := args.Map{"hasErr": err != nil, "len": ac.Length()}
	expected := args.Map{"hasErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyItemsWithTypeValidation stop", actual)
}

func Test_I21_AnyCollection_AddAnyItemsWithTypeValidation_ContinueOnErr(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyItemsWithTypeValidation(
		true, false,
		reflect.TypeOf(""),
		"a", 42, "c",
	)
	actual := args.Map{"hasErr": err != nil, "len": ac.Length()}
	expected := args.Map{"hasErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyItemsWithTypeValidation continue", actual)
}

func Test_I21_AnyCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- AddAnyItemsWithTypeValidation empty", actual)
}

func Test_I21_AnyCollection_AddAnySliceFromSingleItem(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnySliceFromSingleItem([]string{"a", "b"})
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAnySliceFromSingleItem", actual)
}

func Test_I21_AnyCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnySliceFromSingleItem(nil)
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddAnySliceFromSingleItem nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — Paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_I21_AnyCollection_GetPagesSize(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	actual := args.Map{"pages": ac.GetPagesSize(2)}
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagesSize", actual)
}

func Test_I21_AnyCollection_GetPagesSize_Zero(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	actual := args.Map{"pages": ac.GetPagesSize(0)}
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagesSize zero", actual)
}

func Test_I21_AnyCollection_GetPagedCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	for i := 0; i < 5; i++ {
		ac.Add(i)
	}
	pages := ac.GetPagedCollection(2)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagedCollection", actual)
}

func Test_I21_AnyCollection_GetPagedCollection_Small(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	pages := ac.GetPagedCollection(10)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagedCollection small", actual)
}

func Test_I21_AnyCollection_GetSinglePageCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	page := ac.GetSinglePageCollection(3, 2)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetSinglePageCollection", actual)
}

func Test_I21_AnyCollection_GetSinglePageCollection_Small(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	page := ac.GetSinglePageCollection(10, 1)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetSinglePageCollection small", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — JSON methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I21_AnyCollection_JsonString(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	s, err := ac.JsonString()
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonString", actual)
}

func Test_I21_AnyCollection_JsonStringMust(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	s := ac.JsonStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonStringMust", actual)
}

func Test_I21_AnyCollection_MarshalJSON(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	b, err := ac.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- MarshalJSON", actual)
}

func Test_I21_AnyCollection_UnmarshalJSON(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	err := ac.UnmarshalJSON([]byte(`["a","b"]`))
	actual := args.Map{"noErr": err == nil, "len": ac.Length()}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- UnmarshalJSON", actual)
}

func Test_I21_AnyCollection_UnmarshalJSON_Invalid(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	err := ac.UnmarshalJSON([]byte(`not json`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns error -- UnmarshalJSON invalid", actual)
}

func Test_I21_AnyCollection_JsonResultsCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	rc := ac.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_I21_AnyCollection_JsonResultsCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	rc := ac.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- JsonResultsCollection empty", actual)
}

func Test_I21_AnyCollection_JsonResultsPtrCollection(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	rc := ac.JsonResultsPtrCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_I21_AnyCollection_JsonModel(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	model := ac.JsonModel()
	actual := args.Map{"len": len(model)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonModel", actual)
}

func Test_I21_AnyCollection_JsonModelAny(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	actual := args.Map{"notNil": ac.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonModelAny", actual)
}

func Test_I21_AnyCollection_Json(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	jr := ac.Json()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Json", actual)
}

func Test_I21_AnyCollection_JsonPtr(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	jr := ac.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonPtr", actual)
}

func Test_I21_AnyCollection_ParseInjectUsingJson(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	jr := corejson.NewPtr([]any{"a", "b"})
	result, err := ac.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_I21_AnyCollection_JsonParseSelfInject(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	jr := corejson.NewPtr([]any{"x"})
	err := ac.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonParseSelfInject", actual)
}

func Test_I21_AnyCollection_Strings(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	strs := ac.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Strings", actual)
}

func Test_I21_AnyCollection_Strings_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	strs := ac.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Strings empty", actual)
}

func Test_I21_AnyCollection_String(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	actual := args.Map{"notEmpty": ac.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- String", actual)
}

func Test_I21_AnyCollection_ListStringsPtr(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("hello", true)
	strs := ac.ListStringsPtr(false)
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ListStringsPtr", actual)
}

func Test_I21_AnyCollection_ListStrings(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("hello", true)
	strs := ac.ListStrings(true)
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ListStrings", actual)
}

func Test_I21_AnyCollection_ReflectSetAt(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("hello")
	var target string
	err := ac.ReflectSetAt(0, &target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ReflectSetAt", actual)
}

func Test_I21_AnyCollection_GetPagingInfo(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	info := ac.GetPagingInfo(3, 2)
	actual := args.Map{"hasSkip": info.SkipItems > 0}
	expected := args.Map{"hasSkip": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagingInfo", actual)
}

func Test_I21_AnyCollection_Nil_Length(t *testing.T) {
	var ac *coredynamic.AnyCollection
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- nil Length", actual)
}

func Test_I21_AnyCollection_Nil_IsEmpty(t *testing.T) {
	var ac *coredynamic.AnyCollection
	actual := args.Map{"empty": ac.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- nil IsEmpty", actual)
}
