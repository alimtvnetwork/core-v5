package coredynamictests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// DynamicCollection — nil/empty branches
// =============================================================================

func Test_Cov40_DynColl_Length_Nil(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Length nil", actual)
}

func Test_Cov40_DynColl_IsEmpty_Nil(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"r": dc.IsEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection IsEmpty nil", actual)
}

func Test_Cov40_DynColl_IsEmpty_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"r": dc.IsEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection IsEmpty empty", actual)
}

func Test_Cov40_DynColl_HasAnyItem_False(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"r": dc.HasAnyItem()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasAnyItem false", actual)
}

func Test_Cov40_DynColl_HasAnyItem_True(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("x", true)
	actual := args.Map{"r": dc.HasAnyItem()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasAnyItem true", actual)
}

func Test_Cov40_DynColl_Count(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	actual := args.Map{"r": dc.Count()}
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Count", actual)
}

func Test_Cov40_DynColl_LastIndex(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	actual := args.Map{"r": dc.LastIndex()}
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastIndex", actual)
}

func Test_Cov40_DynColl_HasIndex_True(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	actual := args.Map{"r": dc.HasIndex(0)}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasIndex true", actual)
}

func Test_Cov40_DynColl_HasIndex_False(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"r": dc.HasIndex(0)}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasIndex false", actual)
}

// =============================================================================
// DynamicCollection — Items branches
// =============================================================================

func Test_Cov40_DynColl_Items_Nil(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"len": len(dc.Items())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Items nil", actual)
}

func Test_Cov40_DynColl_Items_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	actual := args.Map{"len": len(dc.Items())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Items valid", actual)
}

// =============================================================================
// DynamicCollection — First / Last / OrDefault
// =============================================================================

func Test_Cov40_DynColl_FirstOrDefault_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"isNil": dc.FirstOrDefault() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection FirstOrDefault empty", actual)
}

func Test_Cov40_DynColl_FirstOrDefault_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("first", true)
	actual := args.Map{"notNil": dc.FirstOrDefault() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection FirstOrDefault valid", actual)
}

func Test_Cov40_DynColl_FirstOrDefaultDynamic_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"isNil": dc.FirstOrDefaultDynamic() == nil}
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection FirstOrDefaultDynamic empty", actual)
}

func Test_Cov40_DynColl_LastOrDefault_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"isNil": dc.LastOrDefault() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastOrDefault empty", actual)
}

func Test_Cov40_DynColl_LastOrDefault_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("last", true)
	actual := args.Map{"notNil": dc.LastOrDefault() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastOrDefault valid", actual)
}

func Test_Cov40_DynColl_LastOrDefaultDynamic_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"isNil": dc.LastOrDefaultDynamic() == nil}
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastOrDefaultDynamic empty", actual)
}

// =============================================================================
// DynamicCollection — Skip / Take / Limit
// =============================================================================

func Test_Cov40_DynColl_Skip(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	actual := args.Map{"len": len(dc.Skip(1))}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Skip", actual)
}

func Test_Cov40_DynColl_SkipCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	sc := dc.SkipCollection(2)
	actual := args.Map{"len": sc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection SkipCollection", actual)
}

func Test_Cov40_DynColl_Take(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	actual := args.Map{"len": len(dc.Take(2))}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Take", actual)
}

func Test_Cov40_DynColl_TakeCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	tc := dc.TakeCollection(2)
	actual := args.Map{"len": tc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection TakeCollection", actual)
}

func Test_Cov40_DynColl_LimitCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	lc := dc.LimitCollection(1)
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LimitCollection", actual)
}

func Test_Cov40_DynColl_SafeLimitCollection(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	lc := dc.SafeLimitCollection(100)
	actual := args.Map{"len": lc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection SafeLimitCollection", actual)
}

func Test_Cov40_DynColl_LimitDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	r := dc.LimitDynamic(1)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LimitDynamic", actual)
}

func Test_Cov40_DynColl_Limit(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	actual := args.Map{"len": len(dc.Limit(1))}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Limit", actual)
}

func Test_Cov40_DynColl_SkipDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	r := dc.SkipDynamic(1)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection SkipDynamic", actual)
}

func Test_Cov40_DynColl_TakeDynamic(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	r := dc.TakeDynamic(1)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection TakeDynamic", actual)
}

// =============================================================================
// DynamicCollection — RemoveAt
// =============================================================================

func Test_Cov40_DynColl_RemoveAt_Invalid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"r": dc.RemoveAt(0)}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection RemoveAt invalid", actual)
}

func Test_Cov40_DynColl_RemoveAt_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	ok := dc.RemoveAt(1)
	actual := args.Map{"ok": ok, "len": dc.Length()}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection RemoveAt valid", actual)
}

// =============================================================================
// DynamicCollection — Loop
// =============================================================================

func Test_Cov40_DynColl_Loop_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	called := false
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool { called = true; return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Loop empty", actual)
}

func Test_Cov40_DynColl_Loop_Break(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return true
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Loop break", actual)
}

func Test_Cov40_DynColl_Loop_All(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Loop all", actual)
}

// =============================================================================
// DynamicCollection — Add variants
// =============================================================================

func Test_Cov40_DynColl_AddAnyNonNull_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull(nil, false)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyNonNull nil", actual)
}

func Test_Cov40_DynColl_AddAnyNonNull_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull("a", true)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyNonNull valid", actual)
}

func Test_Cov40_DynColl_AddAnyMany_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany(nil...)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyMany nil", actual)
}

func Test_Cov40_DynColl_AddAnyMany_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyMany valid", actual)
}

func Test_Cov40_DynColl_Add(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("x", true)
	dc.Add(d)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Add", actual)
}

func Test_Cov40_DynColl_AddPtr_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddPtr(nil)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddPtr nil", actual)
}

func Test_Cov40_DynColl_AddPtr_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamicPtr("x", true)
	dc.AddPtr(d)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddPtr valid", actual)
}

func Test_Cov40_DynColl_AddManyPtr_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddManyPtr(nil...)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddManyPtr nil", actual)
}

func Test_Cov40_DynColl_AddManyPtr_WithNils(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	d1 := coredynamic.NewDynamicPtr("a", true)
	dc.AddManyPtr(d1, nil)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddManyPtr with nils", actual)
}

// =============================================================================
// DynamicCollection — AnyItems / AnyItemsCollection / AddAnySliceFromSingleItem
// =============================================================================

func Test_Cov40_DynColl_AnyItems_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"len": len(dc.AnyItems())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItems empty", actual)
}

func Test_Cov40_DynColl_AnyItems_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	actual := args.Map{"len": len(dc.AnyItems())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItems valid", actual)
}

func Test_Cov40_DynColl_AnyItemsCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	ac := dc.AnyItemsCollection()
	actual := args.Map{"empty": ac.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItemsCollection empty", actual)
}

func Test_Cov40_DynColl_AnyItemsCollection_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	ac := dc.AnyItemsCollection()
	actual := args.Map{"len": ac.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItemsCollection valid", actual)
}

func Test_Cov40_DynColl_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnySliceFromSingleItem(true, nil)
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnySliceFromSingleItem nil", actual)
}

func Test_Cov40_DynColl_AddAnySliceFromSingleItem_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnySliceFromSingleItem(true, []int{1, 2, 3})
	actual := args.Map{"len": dc.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnySliceFromSingleItem valid", actual)
}

// =============================================================================
// DynamicCollection — Type validation
// =============================================================================

func Test_Cov40_DynColl_AddAnyWithTypeValidation_Error(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(true, reflect.TypeOf(""), 42)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyWithTypeValidation error", actual)
}

func Test_Cov40_DynColl_AddAnyWithTypeValidation_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(true, reflect.TypeOf(""), "hello")
	actual := args.Map{"noErr": err == nil, "len": dc.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyWithTypeValidation valid", actual)
}

func Test_Cov40_DynColl_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, true, reflect.TypeOf(""))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyItemsWithTypeValidation empty", actual)
}

func Test_Cov40_DynColl_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(true, true, reflect.TypeOf(""), "a", 42, "b")
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyItemsWithTypeValidation continue on error", actual)
}

func Test_Cov40_DynColl_AddAnyItemsWithTypeValidation_StopOnError(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, true, reflect.TypeOf(""), "a", 42, "b")
	actual := args.Map{"hasErr": err != nil, "len": dc.Length()}
	expected := args.Map{"hasErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyItemsWithTypeValidation stop on error", actual)
}

// =============================================================================
// DynamicCollection — JSON
// =============================================================================

func Test_Cov40_DynColl_JsonString_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	s, err := dc.JsonString()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(s) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonString valid", actual)
}

func Test_Cov40_DynColl_JsonStringMust_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	s := dc.JsonStringMust()
	actual := args.Map{"nonEmpty": len(s) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonStringMust valid", actual)
}

func Test_Cov40_DynColl_MarshalJSON(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	b, err := dc.MarshalJSON()
	actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
	expected := args.Map{"noErr": true, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection MarshalJSON", actual)
}

func Test_Cov40_DynColl_UnmarshalJSON_Invalid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.UnmarshalJSON([]byte(`not json`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection UnmarshalJSON invalid", actual)
}

func Test_Cov40_DynColl_ParseInjectUsingJson_Error(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := dc.ParseInjectUsingJson(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ParseInjectUsingJson error", actual)
}

func Test_Cov40_DynColl_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		dc.ParseInjectUsingJsonMust(jr)
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ParseInjectUsingJsonMust panics", actual)
}

func Test_Cov40_DynColl_JsonParseSelfInject_Error(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	err := dc.JsonParseSelfInject(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonParseSelfInject error", actual)
}

func Test_Cov40_DynColl_Json(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Json", actual)
}

func Test_Cov40_DynColl_JsonPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.JsonPtr()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonPtr", actual)
}

// =============================================================================
// DynamicCollection — JsonResultsCollection / JsonResultsPtrCollection
// =============================================================================

func Test_Cov40_DynColl_JsonResultsCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	r := dc.JsonResultsCollection()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsCollection empty", actual)
}

func Test_Cov40_DynColl_JsonResultsCollection_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.JsonResultsCollection()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsCollection valid", actual)
}

func Test_Cov40_DynColl_JsonResultsPtrCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	r := dc.JsonResultsPtrCollection()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsPtrCollection empty", actual)
}

func Test_Cov40_DynColl_JsonResultsPtrCollection_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.JsonResultsPtrCollection()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsPtrCollection valid", actual)
}

// =============================================================================
// DynamicCollection — Paging
// =============================================================================

func Test_Cov40_DynColl_GetPagesSize_Zero(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"r": dc.GetPagesSize(0)}
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagesSize zero", actual)
}

func Test_Cov40_DynColl_GetPagesSize_Negative(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"r": dc.GetPagesSize(-1)}
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagesSize negative", actual)
}

func Test_Cov40_DynColl_GetPagesSize_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	actual := args.Map{"r": dc.GetPagesSize(2)}
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagesSize valid", actual)
}

func Test_Cov40_DynColl_GetPagedCollection_SmallData(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	pages := dc.GetPagedCollection(10)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagedCollection small", actual)
}

func Test_Cov40_DynColl_GetPagedCollection_MultiPage(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	for i := 0; i < 5; i++ {
		dc.AddAny(i, true)
	}
	pages := dc.GetPagedCollection(2)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagedCollection multi", actual)
}

func Test_Cov40_DynColl_GetSinglePageCollection_Small(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.GetSinglePageCollection(10, 1)
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetSinglePageCollection small", actual)
}

// =============================================================================
// DynamicCollection — Misc
// =============================================================================

func Test_Cov40_DynColl_Strings_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"len": len(dc.Strings())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Strings empty", actual)
}

func Test_Cov40_DynColl_Strings_Valid(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	actual := args.Map{"len": len(dc.Strings())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Strings valid", actual)
}

func Test_Cov40_DynColl_String(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)
	actual := args.Map{"nonEmpty": len(dc.String()) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection String", actual)
}

func Test_Cov40_DynColl_JsonModel(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	m := dc.JsonModel()
	actual := args.Map{"hasItems": len(m.Items) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonModel", actual)
}

func Test_Cov40_DynColl_JsonModelAny(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"notNil": dc.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonModelAny", actual)
}

func Test_Cov40_DynColl_ListStrings(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	r := dc.ListStrings()
	actual := args.Map{"nonEmpty": len(r) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ListStrings", actual)
}

func Test_Cov40_DynColl_ListStringsPtr(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	r := dc.ListStringsPtr()
	actual := args.Map{"nonEmpty": len(r) > 0}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ListStringsPtr", actual)
}

func Test_Cov40_DynColl_At(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("x", true)
	d := dc.At(0)
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection At", actual)
}
