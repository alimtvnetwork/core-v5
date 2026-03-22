package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// ResultsCollection — Length / IsEmpty / HasAnyItem / nil receiver
// =============================================================================

func Test_Cov43_ResultsCollection_NilLength(t *testing.T) {
	var c *corejson.ResultsCollection
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection nil Length", actual)
}

func Test_Cov43_ResultsCollection_IsEmpty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection IsEmpty", actual)
}

func Test_Cov43_ResultsCollection_HasAnyItem(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	actual := args.Map{"has": c.HasAnyItem()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection HasAnyItem", actual)
}

func Test_Cov43_ResultsCollection_LastIndex(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b"))
	actual := args.Map{"lastIdx": c.LastIndex()}
	expected := args.Map{"lastIdx": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection LastIndex", actual)
}

// =============================================================================
// ResultsCollection — FirstOrDefault / LastOrDefault
// =============================================================================

func Test_Cov43_ResultsCollection_FirstOrDefault_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"isNil": c.FirstOrDefault() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection FirstOrDefault empty", actual)
}

func Test_Cov43_ResultsCollection_FirstOrDefault_HasItem(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("first"))
	actual := args.Map{"notNil": c.FirstOrDefault() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection FirstOrDefault has item", actual)
}

func Test_Cov43_ResultsCollection_LastOrDefault_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"isNil": c.LastOrDefault() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection LastOrDefault empty", actual)
}

func Test_Cov43_ResultsCollection_LastOrDefault_HasItem(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("last"))
	actual := args.Map{"notNil": c.LastOrDefault() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection LastOrDefault has item", actual)
}

// =============================================================================
// ResultsCollection — Take / Limit / Skip
// =============================================================================

func Test_Cov43_ResultsCollection_Take_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := c.Take(5)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Take empty", actual)
}

func Test_Cov43_ResultsCollection_Take_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b")).Add(corejson.New("c"))
	r := c.Take(2)
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Take valid", actual)
}

func Test_Cov43_ResultsCollection_Limit_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := c.Limit(5)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Limit empty", actual)
}

func Test_Cov43_ResultsCollection_Limit_TakeAll(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b"))
	r := c.Limit(-1) // TakeAllMinusOne
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Limit takeAll", actual)
}

func Test_Cov43_ResultsCollection_Limit_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b")).Add(corejson.New("c"))
	r := c.Limit(2)
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Limit valid", actual)
}

func Test_Cov43_ResultsCollection_Skip_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := c.Skip(1)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Skip empty", actual)
}

func Test_Cov43_ResultsCollection_Skip_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b")).Add(corejson.New("c"))
	r := c.Skip(1)
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Skip valid", actual)
}

// =============================================================================
// ResultsCollection — Add variants
// =============================================================================

func Test_Cov43_ResultsCollection_AddSkipOnNil_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSkipOnNil(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddSkipOnNil nil", actual)
}

func Test_Cov43_ResultsCollection_AddSkipOnNil_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := corejson.NewPtr("hi")
	c.AddSkipOnNil(r)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddSkipOnNil valid", actual)
}

func Test_Cov43_ResultsCollection_AddNonNilNonError_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddNonNilNonError(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddNonNilNonError nil", actual)
}

func Test_Cov43_ResultsCollection_AddNonNilNonError_Error(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := &corejson.Result{Error: errors.New("fail")}
	c.AddNonNilNonError(r)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddNonNilNonError error", actual)
}

func Test_Cov43_ResultsCollection_AddNonNilNonError_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := corejson.NewPtr("hi")
	c.AddNonNilNonError(r)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddNonNilNonError valid", actual)
}

func Test_Cov43_ResultsCollection_AddPtr_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddPtr(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddPtr nil", actual)
}

func Test_Cov43_ResultsCollection_Adds_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Adds()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Adds nil", actual)
}

func Test_Cov43_ResultsCollection_AddsPtr_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddsPtr(nil, nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddsPtr nil items", actual)
}

func Test_Cov43_ResultsCollection_AddAny_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddAny(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddAny nil", actual)
}

func Test_Cov43_ResultsCollection_AddAny_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddAny("hello")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddAny valid", actual)
}

func Test_Cov43_ResultsCollection_AddAnyItems_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddAnyItems()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddAnyItems nil", actual)
}

func Test_Cov43_ResultsCollection_AddAnyItemsSlice_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddAnyItemsSlice(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddAnyItemsSlice nil", actual)
}

func Test_Cov43_ResultsCollection_AddResultsCollection_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddResultsCollection(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddResultsCollection nil", actual)
}

func Test_Cov43_ResultsCollection_AddNonNilItemsPtr_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddNonNilItemsPtr()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddNonNilItemsPtr empty", actual)
}

func Test_Cov43_ResultsCollection_AddMapResults_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	mr := corejson.NewMapResults.Empty()
	c.AddMapResults(mr)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddMapResults empty", actual)
}

func Test_Cov43_ResultsCollection_AddRawMapResults_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddRawMapResults(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddRawMapResults empty", actual)
}

// =============================================================================
// ResultsCollection — HasError / AllErrors / GetErrorsStrings
// =============================================================================

func Test_Cov43_ResultsCollection_HasError_False(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	actual := args.Map{"hasErr": c.HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection HasError false", actual)
}

func Test_Cov43_ResultsCollection_HasError_True(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.Result{Error: errors.New("fail")})
	actual := args.Map{"hasErr": c.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection HasError true", actual)
}

func Test_Cov43_ResultsCollection_AllErrors_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	errList, hasErr := c.AllErrors()
	actual := args.Map{"len": len(errList), "hasErr": hasErr}
	expected := args.Map{"len": 0, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AllErrors empty", actual)
}

func Test_Cov43_ResultsCollection_AllErrors_WithErrors(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.Result{Error: errors.New("fail")})
	c.Add(corejson.New("ok"))
	errList, hasErr := c.AllErrors()
	actual := args.Map{"len": len(errList), "hasErr": hasErr}
	expected := args.Map{"len": 1, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AllErrors with errors", actual)
}

func Test_Cov43_ResultsCollection_GetErrorsStrings_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := c.GetErrorsStrings()
	actual := args.Map{"len": len(r)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetErrorsStrings empty", actual)
}

func Test_Cov43_ResultsCollection_GetErrorsStringsPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := c.GetErrorsStringsPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetErrorsStringsPtr", actual)
}

func Test_Cov43_ResultsCollection_GetErrorsAsSingleString(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	s := c.GetErrorsAsSingleString()
	actual := args.Map{"empty": s == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetErrorsAsSingleString", actual)
}

func Test_Cov43_ResultsCollection_GetErrorsAsSingle(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	e := c.GetErrorsAsSingle()
	actual := args.Map{"notNil": e != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetErrorsAsSingle", actual)
}

// =============================================================================
// ResultsCollection — GetAt / GetAtSafe / GetAtSafeUsingLength
// =============================================================================

func Test_Cov43_ResultsCollection_GetAt(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	r := c.GetAt(0)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetAt", actual)
}

func Test_Cov43_ResultsCollection_GetAtSafe_InRange(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	r := c.GetAtSafe(0)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetAtSafe in range", actual)
}

func Test_Cov43_ResultsCollection_GetAtSafe_OutOfRange(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := c.GetAtSafe(5)
	actual := args.Map{"isNil": r == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetAtSafe out of range", actual)
}

func Test_Cov43_ResultsCollection_GetAtSafeUsingLength(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	r := c.GetAtSafeUsingLength(0, 1)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetAtSafeUsingLength", actual)
}

// =============================================================================
// ResultsCollection — Unmarshal / Inject / Paging / Json / Clone
// =============================================================================

func Test_Cov43_ResultsCollection_UnmarshalAt(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("hello"))
	var s string
	err := c.UnmarshalAt(0, &s)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "ResultsCollection UnmarshalAt", actual)
}

func Test_Cov43_ResultsCollection_UnmarshalIntoSameIndex_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	errList, hasErr := c.UnmarshalIntoSameIndex()
	actual := args.Map{"len": len(errList), "hasErr": hasErr}
	expected := args.Map{"len": 0, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection UnmarshalIntoSameIndex nil", actual)
}

func Test_Cov43_ResultsCollection_UnmarshalIntoSameIndex_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("hello"))
	var s string
	errList, hasErr := c.UnmarshalIntoSameIndex(&s)
	actual := args.Map{"len": len(errList), "hasErr": hasErr, "r": s}
	expected := args.Map{"len": 1, "hasErr": false, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "ResultsCollection UnmarshalIntoSameIndex valid", actual)
}

func Test_Cov43_ResultsCollection_UnmarshalIntoSameIndex_NilItem(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("hello"))
	errList, hasErr := c.UnmarshalIntoSameIndex(nil)
	actual := args.Map{"len": len(errList), "hasErr": hasErr}
	expected := args.Map{"len": 1, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection UnmarshalIntoSameIndex nil item", actual)
}

func Test_Cov43_ResultsCollection_UnmarshalIntoSameIndex_ErrorResult(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.Result{Error: errors.New("fail")})
	var s string
	errList, hasErr := c.UnmarshalIntoSameIndex(&s)
	actual := args.Map{"len": len(errList), "hasErr": hasErr}
	expected := args.Map{"len": 1, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection UnmarshalIntoSameIndex error result", actual)
}

func Test_Cov43_ResultsCollection_InjectIntoSameIndex_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	errList, hasErr := c.InjectIntoSameIndex()
	actual := args.Map{"len": len(errList), "hasErr": hasErr}
	expected := args.Map{"len": 0, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "ResultsCollection InjectIntoSameIndex nil", actual)
}

func Test_Cov43_ResultsCollection_GetPagesSize_Zero(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"pages": c.GetPagesSize(0)}
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetPagesSize zero", actual)
}

func Test_Cov43_ResultsCollection_GetPagesSize_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a")).Add(corejson.New("b")).Add(corejson.New("c"))
	actual := args.Map{"pages": c.GetPagesSize(2)}
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetPagesSize valid", actual)
}

func Test_Cov43_ResultsCollection_GetPagedCollection_SmallSet(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	pages := c.GetPagedCollection(10)
	actual := args.Map{"len": len(pages)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetPagedCollection small", actual)
}

func Test_Cov43_ResultsCollection_GetPagedCollection_Multi(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	for i := 0; i < 5; i++ {
		c.Add(corejson.New(i))
	}
	pages := c.GetPagedCollection(2)
	actual := args.Map{"len": len(pages)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetPagedCollection multi", actual)
}

func Test_Cov43_ResultsCollection_GetSinglePageCollection_SmallSet(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	page := c.GetSinglePageCollection(10, 1)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetSinglePageCollection small", actual)
}

func Test_Cov43_ResultsCollection_Json(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	r := c.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Json", actual)
}

func Test_Cov43_ResultsCollection_JsonPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	r := c.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection JsonPtr", actual)
}

func Test_Cov43_ResultsCollection_GetStrings(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	ss := c.GetStrings()
	actual := args.Map{"len": len(ss)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetStrings empty", actual)
}

func Test_Cov43_ResultsCollection_GetStringsPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	ss := c.GetStringsPtr()
	actual := args.Map{"notNil": ss != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection GetStringsPtr", actual)
}

func Test_Cov43_ResultsCollection_NonPtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	np := c.NonPtr()
	actual := args.Map{"len": np.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection NonPtr", actual)
}

func Test_Cov43_ResultsCollection_Ptr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	p := c.Ptr()
	actual := args.Map{"notNil": p != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Ptr", actual)
}

func Test_Cov43_ResultsCollection_Clear(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	c.Clear()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Clear", actual)
}

func Test_Cov43_ResultsCollection_Clear_Nil(t *testing.T) {
	var c *corejson.ResultsCollection
	r := c.Clear()
	actual := args.Map{"isNil": r == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Clear nil", actual)
}

func Test_Cov43_ResultsCollection_Dispose(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	c.Dispose()
	actual := args.Map{"nilItems": c.Items == nil}
	expected := args.Map{"nilItems": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Dispose", actual)
}

func Test_Cov43_ResultsCollection_Dispose_Nil(t *testing.T) {
	var c *corejson.ResultsCollection
	c.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Dispose nil", actual)
}

func Test_Cov43_ResultsCollection_ShadowClone(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	cloned := c.ShadowClone()
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection ShadowClone", actual)
}

func Test_Cov43_ResultsCollection_Clone(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	cloned := c.Clone(true)
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection Clone deep", actual)
}

func Test_Cov43_ResultsCollection_ClonePtr_Nil(t *testing.T) {
	var c *corejson.ResultsCollection
	r := c.ClonePtr(false)
	actual := args.Map{"isNil": r == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection ClonePtr nil", actual)
}

func Test_Cov43_ResultsCollection_ClonePtr(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	r := c.ClonePtr(true)
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsCollection ClonePtr", actual)
}

func Test_Cov43_ResultsCollection_AsJsonContractsBinder(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"notNil": c.AsJsonContractsBinder() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AsJsonContractsBinder", actual)
}

func Test_Cov43_ResultsCollection_AsJsoner(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"notNil": c.AsJsoner() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AsJsoner", actual)
}

func Test_Cov43_ResultsCollection_AsJsonParseSelfInjector(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"notNil": c.AsJsonParseSelfInjector() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AsJsonParseSelfInjector", actual)
}

func Test_Cov43_ResultsCollection_JsonModel(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"notNil": c.JsonModel() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection JsonModel", actual)
}

func Test_Cov43_ResultsCollection_JsonModelAny(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	actual := args.Map{"notNil": c.JsonModelAny() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection JsonModelAny", actual)
}

func Test_Cov43_ResultsCollection_ParseInjectUsingJson(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	jr := c.JsonPtr()
	c2 := corejson.NewResultsCollection.Empty()
	_, err := c2.ParseInjectUsingJson(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection ParseInjectUsingJson", actual)
}

func Test_Cov43_ResultsCollection_ParseInjectUsingJsonMust(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	jr := c.JsonPtr()
	c2 := corejson.NewResultsCollection.Empty()
	r := c2.ParseInjectUsingJsonMust(jr)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection ParseInjectUsingJsonMust", actual)
}

func Test_Cov43_ResultsCollection_JsonParseSelfInject(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	jr := c.JsonPtr()
	c2 := corejson.NewResultsCollection.Empty()
	err := c2.JsonParseSelfInject(jr)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ResultsCollection JsonParseSelfInject", actual)
}

// =============================================================================
// ResultsCollection — AddSerializer / AddSerializerFunc
// =============================================================================

func Test_Cov43_ResultsCollection_AddSerializer_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializer(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddSerializer nil", actual)
}

func Test_Cov43_ResultsCollection_AddSerializers_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializers()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddSerializers empty", actual)
}

func Test_Cov43_ResultsCollection_AddSerializerFunc_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializerFunc(nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddSerializerFunc nil", actual)
}

func Test_Cov43_ResultsCollection_AddSerializerFunctions_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddSerializerFunctions()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddSerializerFunctions empty", actual)
}

// =============================================================================
// ResultsCollection — AddJsoners
// =============================================================================

func Test_Cov43_ResultsCollection_AddJsoners_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.AddJsoners(true)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsCollection AddJsoners nil", actual)
}

// =============================================================================
// newResultsCollectionCreator
// =============================================================================

func Test_Cov43_NewResultsCollection_Default(t *testing.T) {
	c := corejson.NewResultsCollection.Default()
	actual := args.Map{"notNil": c != nil, "empty": c.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection Default", actual)
}

func Test_Cov43_NewResultsCollection_UsingCap(t *testing.T) {
	c := corejson.NewResultsCollection.UsingCap(10)
	actual := args.Map{"notNil": c != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection UsingCap", actual)
}

func Test_Cov43_NewResultsCollection_AnyItems(t *testing.T) {
	c := corejson.NewResultsCollection.AnyItems("a", "b")
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection AnyItems", actual)
}

func Test_Cov43_NewResultsCollection_AnyItemsPlusCap_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.AnyItemsPlusCap(5)
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection AnyItemsPlusCap empty", actual)
}

func Test_Cov43_NewResultsCollection_UsingResults(t *testing.T) {
	r := corejson.New("a")
	c := corejson.NewResultsCollection.UsingResults(r)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection UsingResults", actual)
}

func Test_Cov43_NewResultsCollection_UsingResultsPtr(t *testing.T) {
	r := corejson.NewPtr("a")
	c := corejson.NewResultsCollection.UsingResultsPtr(r)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection UsingResultsPtr", actual)
}

func Test_Cov43_NewResultsCollection_UsingResultsPlusCap_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.UsingResultsPlusCap(5)
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection UsingResultsPlusCap nil", actual)
}

func Test_Cov43_NewResultsCollection_UsingResultsPtrPlusCap_Nil(t *testing.T) {
	c := corejson.NewResultsCollection.UsingResultsPtrPlusCap(5)
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection UsingResultsPtrPlusCap nil", actual)
}

func Test_Cov43_NewResultsCollection_UsingJsoners(t *testing.T) {
	c := corejson.NewResultsCollection.UsingJsoners()
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection UsingJsoners empty", actual)
}

func Test_Cov43_NewResultsCollection_Serializers_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.Serializers()
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection Serializers empty", actual)
}

func Test_Cov43_NewResultsCollection_SerializerFunctions_Empty(t *testing.T) {
	c := corejson.NewResultsCollection.SerializerFunctions()
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection SerializerFunctions empty", actual)
}

func Test_Cov43_NewResultsCollection_DeserializeUsingBytes_Valid(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	b := corejson.Serialize.ToBytesMust(c)
	c2, err := corejson.NewResultsCollection.DeserializeUsingBytes(b)
	actual := args.Map{"noErr": err == nil, "notNil": c2 != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection DeserializeUsingBytes", actual)
}

func Test_Cov43_NewResultsCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := corejson.NewResultsCollection.DeserializeUsingBytes([]byte(`bad`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection DeserializeUsingBytes invalid", actual)
}

func Test_Cov43_NewResultsCollection_UnmarshalUsingBytes(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	b := corejson.Serialize.ToBytesMust(c)
	c2, err := corejson.NewResultsCollection.UnmarshalUsingBytes(b)
	actual := args.Map{"noErr": err == nil, "notNil": c2 != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection UnmarshalUsingBytes", actual)
}

func Test_Cov43_NewResultsCollection_DeserializeUsingResult(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.New("a"))
	jr := c.JsonPtr()
	c2, err := corejson.NewResultsCollection.DeserializeUsingResult(jr)
	actual := args.Map{"noErr": err == nil, "notNil": c2 != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection DeserializeUsingResult", actual)
}

func Test_Cov43_NewResultsCollection_DeserializeUsingResult_Error(t *testing.T) {
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := corejson.NewResultsCollection.DeserializeUsingResult(jr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection DeserializeUsingResult error", actual)
}
