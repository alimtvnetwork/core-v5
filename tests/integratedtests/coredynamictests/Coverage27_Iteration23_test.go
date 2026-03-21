package coredynamictests

import (
	"reflect"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — Add variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I23_DynamicCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"empty": dc.IsEmpty(), "len": dc.Length(), "count": dc.Count()}
	expected := args.Map{"empty": true, "len": 0, "count": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- Empty", actual)
}

func Test_I23_DynamicCollection_AddAny(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true).AddAny("world", true)
	actual := args.Map{"len": dc.Length(), "hasAny": dc.HasAnyItem()}
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddAny", actual)
}

func Test_I23_DynamicCollection_AddAnyNonNull(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull("a", true).AddAnyNonNull(nil, true)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddAnyNonNull", actual)
}

func Test_I23_DynamicCollection_AddAnyMany(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b", "c")
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddAnyMany", actual)
}

func Test_I23_DynamicCollection_AddAnyMany_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany()
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns nil -- AddAnyMany nil", actual)
}

func Test_I23_DynamicCollection_Add(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("hello", true)
	dc.Add(d)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Add", actual)
}

func Test_I23_DynamicCollection_AddPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("hello", true)
	dc.AddPtr(&d).AddPtr(nil)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddPtr", actual)
}

func Test_I23_DynamicCollection_AddManyPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d1 := coredynamic.NewDynamic("a", true)
	d2 := coredynamic.NewDynamic("b", true)
	dc.AddManyPtr(&d1, nil, &d2)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddManyPtr", actual)
}

func Test_I23_DynamicCollection_AddManyPtr_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddManyPtr()
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns nil -- AddManyPtr nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — Navigation
// ══════════════════════════════════════════════════════════════════════════════

func Test_I23_DynamicCollection_At(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	d := dc.At(1)
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- At", actual)
}

func Test_I23_DynamicCollection_Items(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	actual := args.Map{"len": len(dc.Items())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Items", actual)
}

func Test_I23_DynamicCollection_Items_Nil(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	items := dc.Items()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns nil -- Items nil", actual)
}

func Test_I23_DynamicCollection_First_Last(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("first", true).AddAny("last", true)
	actual := args.Map{"hasFirst": dc.First().IsValid(), "hasLast": dc.Last().IsValid()}
	expected := args.Map{"hasFirst": true, "hasLast": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- First/Last", actual)
}

func Test_I23_DynamicCollection_FirstDynamic_LastDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	actual := args.Map{"firstNotNil": dc.FirstDynamic() != nil, "lastNotNil": dc.LastDynamic() != nil}
	expected := args.Map{"firstNotNil": true, "lastNotNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- FirstDynamic/LastDynamic", actual)
}

func Test_I23_DynamicCollection_FirstOrDefault_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"nil": dc.FirstOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- FirstOrDefault empty", actual)
}

func Test_I23_DynamicCollection_FirstOrDefault_Has(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("x", true)
	actual := args.Map{"notNil": dc.FirstOrDefault() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- FirstOrDefault has", actual)
}

func Test_I23_DynamicCollection_FirstOrDefaultDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"nil": dc.FirstOrDefaultDynamic() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- FirstOrDefaultDynamic", actual)
}

func Test_I23_DynamicCollection_LastOrDefault_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"nil": dc.LastOrDefault() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- LastOrDefault empty", actual)
}

func Test_I23_DynamicCollection_LastOrDefault_Has(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("x", true)
	actual := args.Map{"notNil": dc.LastOrDefault() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- LastOrDefault has", actual)
}

func Test_I23_DynamicCollection_LastOrDefaultDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"nil": dc.LastOrDefaultDynamic() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- LastOrDefaultDynamic", actual)
}

func Test_I23_DynamicCollection_Skip(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	skipped := dc.Skip(1)
	actual := args.Map{"len": len(skipped)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Skip", actual)
}

func Test_I23_DynamicCollection_SkipDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	actual := args.Map{"notNil": dc.SkipDynamic(1) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- SkipDynamic", actual)
}

func Test_I23_DynamicCollection_SkipCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	sc := dc.SkipCollection(2)
	actual := args.Map{"len": sc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- SkipCollection", actual)
}

func Test_I23_DynamicCollection_Take(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	taken := dc.Take(2)
	actual := args.Map{"len": len(taken)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Take", actual)
}

func Test_I23_DynamicCollection_TakeDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	actual := args.Map{"notNil": dc.TakeDynamic(1) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- TakeDynamic", actual)
}

func Test_I23_DynamicCollection_TakeCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	tc := dc.TakeCollection(2)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- TakeCollection", actual)
}

func Test_I23_DynamicCollection_LimitCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	lc := dc.LimitCollection(2)
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- LimitCollection", actual)
}

func Test_I23_DynamicCollection_SafeLimitCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	lc := dc.SafeLimitCollection(10)
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- SafeLimitCollection", actual)
}

func Test_I23_DynamicCollection_LimitDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	actual := args.Map{"notNil": dc.LimitDynamic(1) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- LimitDynamic", actual)
}

func Test_I23_DynamicCollection_Limit(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	actual := args.Map{"len": len(dc.Limit(1))}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Limit", actual)
}

func Test_I23_DynamicCollection_LastIndex(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	actual := args.Map{"idx": dc.LastIndex()}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- LastIndex", actual)
}

func Test_I23_DynamicCollection_HasIndex(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	actual := args.Map{"has0": dc.HasIndex(0), "has1": dc.HasIndex(1)}
	expected := args.Map{"has0": true, "has1": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- HasIndex", actual)
}

func Test_I23_DynamicCollection_Loop(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Loop", actual)
}

func Test_I23_DynamicCollection_Loop_Break(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return i == 0
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Loop break", actual)
}

func Test_I23_DynamicCollection_Loop_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- Loop empty", actual)
}

func Test_I23_DynamicCollection_RemoveAt(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	ok := dc.RemoveAt(1)
	actual := args.Map{"ok": ok, "len": dc.Length()}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- RemoveAt", actual)
}

func Test_I23_DynamicCollection_RemoveAt_Invalid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	ok := dc.RemoveAt(5)
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns error -- RemoveAt invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — Type validation
// ══════════════════════════════════════════════════════════════════════════════

func Test_I23_DynamicCollection_AddAnyWithTypeValidation_Match(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(""), "hello")
	actual := args.Map{"noErr": err == nil, "len": dc.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns non-empty -- AddAnyWithTypeValidation match", actual)
}

func Test_I23_DynamicCollection_AddAnyWithTypeValidation_Mismatch(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(""), 42)
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns non-empty -- AddAnyWithTypeValidation mismatch", actual)
}

func Test_I23_DynamicCollection_AddAnyItemsWithTypeValidation_StopOnErr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""), "a", 42, "c")
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns non-empty -- AddAnyItemsWithTypeValidation stop", actual)
}

func Test_I23_DynamicCollection_AddAnyItemsWithTypeValidation_ContinueOnErr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(""), "a", 42, "c")
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns non-empty -- AddAnyItemsWithTypeValidation continue", actual)
}

func Test_I23_DynamicCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- AddAnyItemsWithTypeValidation empty", actual)
}

func Test_I23_DynamicCollection_AddAnySliceFromSingleItem(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnySliceFromSingleItem(true, []string{"a", "b"})
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AddAnySliceFromSingleItem", actual)
}

func Test_I23_DynamicCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnySliceFromSingleItem(true, nil)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns nil -- AddAnySliceFromSingleItem nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — AnyItems, ListStrings, Strings
// ══════════════════════════════════════════════════════════════════════════════

func Test_I23_DynamicCollection_AnyItems(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true)
	items := dc.AnyItems()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AnyItems", actual)
}

func Test_I23_DynamicCollection_AnyItems_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	items := dc.AnyItems()
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- AnyItems empty", actual)
}

func Test_I23_DynamicCollection_AnyItemsCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	ac := dc.AnyItemsCollection()
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- AnyItemsCollection", actual)
}

func Test_I23_DynamicCollection_AnyItemsCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	ac := dc.AnyItemsCollection()
	actual := args.Map{"empty": ac.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- AnyItemsCollection empty", actual)
}

func Test_I23_DynamicCollection_ListStringsPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	strs := dc.ListStringsPtr()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ListStringsPtr", actual)
}

func Test_I23_DynamicCollection_ListStrings(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	strs := dc.ListStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ListStrings", actual)
}

func Test_I23_DynamicCollection_Strings(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	strs := dc.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Strings", actual)
}

func Test_I23_DynamicCollection_Strings_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	strs := dc.Strings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- Strings empty", actual)
}

func Test_I23_DynamicCollection_String(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	s := dc.String()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- String", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — JSON
// ══════════════════════════════════════════════════════════════════════════════

func Test_I23_DynamicCollection_JsonString(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	s, err := dc.JsonString()
	actual := args.Map{"noErr": err == nil, "notEmpty": s != ""}
	expected := args.Map{"noErr": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonString", actual)
}

func Test_I23_DynamicCollection_JsonStringMust(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	s := dc.JsonStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonStringMust", actual)
}

func Test_I23_DynamicCollection_MarshalJSON(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	b, err := dc.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- MarshalJSON", actual)
}

func Test_I23_DynamicCollection_UnmarshalJSON(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("seed", true)
	b, _ := dc.MarshalJSON()
	dc2 := coredynamic.EmptyDynamicCollection()
	err := dc2.UnmarshalJSON(b)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- UnmarshalJSON", actual)
}

func Test_I23_DynamicCollection_UnmarshalJSON_Invalid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.UnmarshalJSON([]byte(`not json`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns error -- UnmarshalJSON invalid", actual)
}

func Test_I23_DynamicCollection_JsonResultsCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	rc := dc.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_I23_DynamicCollection_JsonResultsCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	rc := dc.JsonResultsCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- JsonResultsCollection empty", actual)
}

func Test_I23_DynamicCollection_JsonResultsPtrCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	rc := dc.JsonResultsPtrCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_I23_DynamicCollection_JsonResultsPtrCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	rc := dc.JsonResultsPtrCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns empty -- JsonResultsPtrCollection empty", actual)
}

func Test_I23_DynamicCollection_JsonModel(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	model := dc.JsonModel()
	actual := args.Map{"len": len(model.Items)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonModel", actual)
}

func Test_I23_DynamicCollection_JsonModelAny(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	actual := args.Map{"notNil": dc.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonModelAny", actual)
}

func Test_I23_DynamicCollection_Json(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	jr := dc.Json()
	actual := args.Map{"hasErr": jr.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- Json", actual)
}

func Test_I23_DynamicCollection_JsonPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	jr := dc.JsonPtr()
	actual := args.Map{"notNil": jr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonPtr", actual)
}

func Test_I23_DynamicCollection_ParseInjectUsingJson(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	b, _ := dc.MarshalJSON()
	jr := corejson.NewPtr(dc)
	dc2 := coredynamic.EmptyDynamicCollection()
	result, err := dc2.ParseInjectUsingJson(jr)
	_ = b
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_I23_DynamicCollection_JsonParseSelfInject(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	jr := corejson.NewPtr(dc)
	dc2 := coredynamic.EmptyDynamicCollection()
	err := dc2.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- JsonParseSelfInject", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicCollection — Paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_I23_DynamicCollection_GetPagesSize(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	actual := args.Map{"pages": dc.GetPagesSize(2)}
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- GetPagesSize", actual)
}

func Test_I23_DynamicCollection_GetPagesSize_Zero(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"pages": dc.GetPagesSize(0)}
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- GetPagesSize zero", actual)
}

func Test_I23_DynamicCollection_GetPagedCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	for i := 0; i < 5; i++ {
		dc.AddAny(i, true)
	}
	pages := dc.GetPagedCollection(2)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- GetPagedCollection", actual)
}

func Test_I23_DynamicCollection_GetPagedCollection_Small(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	pages := dc.GetPagedCollection(10)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- GetPagedCollection small", actual)
}

func Test_I23_DynamicCollection_GetSinglePageCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	for i := 0; i < 10; i++ {
		dc.AddAny(i, true)
	}
	page := dc.GetSinglePageCollection(3, 2)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- GetSinglePageCollection", actual)
}

func Test_I23_DynamicCollection_GetSinglePageCollection_Small(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	page := dc.GetSinglePageCollection(10, 1)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- GetSinglePageCollection small", actual)
}

func Test_I23_DynamicCollection_GetPagingInfo(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	for i := 0; i < 10; i++ {
		dc.AddAny(i, true)
	}
	info := dc.GetPagingInfo(3, 2)
	actual := args.Map{"hasSkip": info.SkipItems > 0}
	expected := args.Map{"hasSkip": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- GetPagingInfo", actual)
}

func Test_I23_DynamicCollection_Nil_Length(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns nil -- nil Length", actual)
}

func Test_I23_DynamicCollection_Nil_IsEmpty(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"empty": dc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns nil -- nil IsEmpty", actual)
}

func Test_I23_DynamicCollection_ParseInjectUsingJsonMust(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	jr := corejson.NewPtr(dc)
	dc2 := coredynamic.EmptyDynamicCollection()
	result := dc2.ParseInjectUsingJsonMust(jr)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection returns correct value -- ParseInjectUsingJsonMust", actual)
}

// Unused import suppressor
var _ = strings.Join
