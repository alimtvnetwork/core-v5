package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ═══════════════════════════════════════════════════════════════
// Collection — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_Collection_LengthLock(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "LengthLock", Expected: 2, Actual: c.LengthLock(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_IsEquals_Same(t *testing.T) {
	c1 := corestr.New.Collection.Strings([]string{"a", "b"})
	c2 := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "IsEquals same", Expected: true, Actual: c1.IsEquals(c2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_IsEquals_Different(t *testing.T) {
	c1 := corestr.New.Collection.Strings([]string{"a", "b"})
	c2 := corestr.New.Collection.Strings([]string{"a", "c"})
	tc := coretestcases.CaseV1{Name: "IsEquals diff", Expected: false, Actual: c1.IsEquals(c2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	c1 := corestr.New.Collection.Strings([]string{"ABC"})
	c2 := corestr.New.Collection.Strings([]string{"abc"})
	tc := coretestcases.CaseV1{Name: "IsEqualsWithSensitive insensitive", Expected: true, Actual: c1.IsEqualsWithSensitive(false, c2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_IsEqualsWithSensitive_BothNil(t *testing.T) {
	var c1, c2 *corestr.Collection
	tc := coretestcases.CaseV1{Name: "IsEquals both nil", Expected: true, Actual: c1.IsEquals(c2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_IsEqualsWithSensitive_OneNil(t *testing.T) {
	c1 := corestr.New.Collection.Strings([]string{"a"})
	var c2 *corestr.Collection
	tc := coretestcases.CaseV1{Name: "IsEquals one nil", Expected: false, Actual: c1.IsEquals(c2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_IsEqualsWithSensitive_DiffLength(t *testing.T) {
	c1 := corestr.New.Collection.Strings([]string{"a"})
	c2 := corestr.New.Collection.Strings([]string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "IsEquals diff length", Expected: false, Actual: c1.IsEquals(c2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_IsEmptyLock(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "IsEmptyLock", Expected: true, Actual: c.IsEmptyLock(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_HasItems(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	tc := coretestcases.CaseV1{Name: "HasItems", Expected: true, Actual: c.HasItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddLock(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddLock("hello")
	tc := coretestcases.CaseV1{Name: "AddLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddNonEmpty(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddNonEmpty("")
	c.AddNonEmpty("x")
	tc := coretestcases.CaseV1{Name: "AddNonEmpty", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddNonEmptyWhitespace(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddNonEmptyWhitespace("  ")
	c.AddNonEmptyWhitespace("x")
	tc := coretestcases.CaseV1{Name: "AddNonEmptyWhitespace", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddError(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddError(nil)
	tc := coretestcases.CaseV1{Name: "AddError nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AsDefaultError_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "AsDefaultError empty", Expected: true, Actual: c.AsDefaultError() == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AsError_HasItems(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"err1"})
	err := c.AsError(",")
	tc := coretestcases.CaseV1{Name: "AsError has items", Expected: true, Actual: err != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddIf_True(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddIf(true, "yes")
	tc := coretestcases.CaseV1{Name: "AddIf true", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddIf_False(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddIf(false, "no")
	tc := coretestcases.CaseV1{Name: "AddIf false", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddIfMany_True(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddIfMany(true, "a", "b")
	tc := coretestcases.CaseV1{Name: "AddIfMany true", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddIfMany_False(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddIfMany(false, "a", "b")
	tc := coretestcases.CaseV1{Name: "AddIfMany false", Expected: 0, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddFunc(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddFunc(func() string { return "hello" })
	tc := coretestcases.CaseV1{Name: "AddFunc", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddFuncErr_NoError(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
	tc := coretestcases.CaseV1{Name: "AddFuncErr no err", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_AddsLock(t *testing.T) {
	c := corestr.New.Collection.Cap(2)
	c.AddsLock("a", "b")
	tc := coretestcases.CaseV1{Name: "AddsLock", Expected: 2, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_EachItemSplitBy(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a,b", "c,d"})
	result := c.EachItemSplitBy(",")
	tc := coretestcases.CaseV1{Name: "EachItemSplitBy", Expected: 4, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_ConcatNew_NoAdding(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	result := c.ConcatNew(0)
	tc := coretestcases.CaseV1{Name: "ConcatNew no adding", Expected: 1, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_ConcatNew_WithAdding(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"a"})
	result := c.ConcatNew(0, "b", "c")
	tc := coretestcases.CaseV1{Name: "ConcatNew with adding", Expected: 3, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_ToError_Empty(t *testing.T) {
	c := corestr.New.Collection.Cap(0)
	tc := coretestcases.CaseV1{Name: "ToError empty", Expected: true, Actual: c.ToError(",") == nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_Collection_ToDefaultError(t *testing.T) {
	c := corestr.New.Collection.Strings([]string{"err"})
	tc := coretestcases.CaseV1{Name: "ToDefaultError", Expected: true, Actual: c.ToDefaultError() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// HashsetsCollection — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_HashsetsCollection_IsEmpty(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC IsEmpty", Expected: true, Actual: hc.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_HasItems(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	tc := coretestcases.CaseV1{Name: "HC HasItems", Expected: true, Actual: hc.HasItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_Add(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	tc := coretestcases.CaseV1{Name: "HC Add", Expected: 1, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AddNonNil(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.AddNonNil(nil)
	hc.AddNonNil(corestr.New.Hashset.StringsSpreadItems("a"))
	tc := coretestcases.CaseV1{Name: "HC AddNonNil", Expected: 1, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AddNonEmpty(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.AddNonEmpty(corestr.New.Hashset.Empty())
	hc.AddNonEmpty(corestr.New.Hashset.StringsSpreadItems("a"))
	tc := coretestcases.CaseV1{Name: "HC AddNonEmpty", Expected: 1, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_Adds(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Adds(corestr.New.Hashset.StringsSpreadItems("a"), corestr.New.Hashset.StringsSpreadItems("b"))
	tc := coretestcases.CaseV1{Name: "HC Adds", Expected: 2, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_Adds_NilSkip(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Adds(nil)
	tc := coretestcases.CaseV1{Name: "HC Adds nil", Expected: 0, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_LastIndex(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	hc.Add(corestr.New.Hashset.StringsSpreadItems("b"))
	tc := coretestcases.CaseV1{Name: "HC LastIndex", Expected: 1, Actual: hc.LastIndex(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_StringsList(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	result := hc.StringsList()
	tc := coretestcases.CaseV1{Name: "HC StringsList", Expected: 1, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_StringsList_Empty(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	result := hc.StringsList()
	tc := coretestcases.CaseV1{Name: "HC StringsList empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_ListPtr(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC ListPtr", Expected: true, Actual: hc.ListPtr() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_List(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC List", Expected: 0, Actual: len(hc.List()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_ListDirectPtr(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("x"))
	result := hc.ListDirectPtr()
	tc := coretestcases.CaseV1{Name: "HC ListDirectPtr", Expected: true, Actual: result != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AddHashsetsCollection(t *testing.T) {
	hc1 := corestr.Empty.HashsetsCollection()
	hc1.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	hc2 := corestr.Empty.HashsetsCollection()
	hc2.AddHashsetsCollection(hc1)
	tc := coretestcases.CaseV1{Name: "HC AddHashsetsCollection", Expected: 1, Actual: hc2.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AddHashsetsCollection_Nil(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.AddHashsetsCollection(nil)
	tc := coretestcases.CaseV1{Name: "HC AddHC nil", Expected: 0, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_ConcatNew_NoArgs(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	result := hc.ConcatNew()
	tc := coretestcases.CaseV1{Name: "HC ConcatNew no args", Expected: 1, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_ConcatNew_WithArgs(t *testing.T) {
	hc1 := corestr.Empty.HashsetsCollection()
	hc1.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	hc2 := corestr.Empty.HashsetsCollection()
	hc2.Add(corestr.New.Hashset.StringsSpreadItems("b"))
	result := hc1.ConcatNew(hc2)
	tc := coretestcases.CaseV1{Name: "HC ConcatNew with args", Expected: 2, Actual: result.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_HasAll(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a", "b"))
	tc := coretestcases.CaseV1{Name: "HC HasAll", Expected: true, Actual: hc.HasAll("a", "b"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_HasAll_Empty(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC HasAll empty", Expected: false, Actual: hc.HasAll("a"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_IsEqual_Same(t *testing.T) {
	hc1 := corestr.Empty.HashsetsCollection()
	hc1.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	hc2 := corestr.Empty.HashsetsCollection()
	hc2.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	tc := coretestcases.CaseV1{Name: "HC IsEqualPtr same", Expected: true, Actual: hc1.IsEqualPtr(hc2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_IsEqualPtr_BothNil(t *testing.T) {
	var hc1, hc2 *corestr.HashsetsCollection
	tc := coretestcases.CaseV1{Name: "HC IsEqualPtr both nil", Expected: true, Actual: hc1.IsEqualPtr(hc2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_IsEqualPtr_OneNil(t *testing.T) {
	hc1 := corestr.Empty.HashsetsCollection()
	var hc2 *corestr.HashsetsCollection
	tc := coretestcases.CaseV1{Name: "HC IsEqualPtr one nil", Expected: false, Actual: hc1.IsEqualPtr(hc2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_String(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	tc := coretestcases.CaseV1{Name: "HC String", Expected: true, Actual: len(hc.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_String_Empty(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC String empty", Expected: true, Actual: len(hc.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_Join(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	result := hc.Join(",")
	tc := coretestcases.CaseV1{Name: "HC Join", Expected: true, Actual: len(result) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_Serialize(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	hc.Add(corestr.New.Hashset.StringsSpreadItems("a"))
	data, err := hc.Serialize()
	tc := coretestcases.CaseV1{Name: "HC Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AsJsonContractsBinder(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC AsJsonContractsBinder", Expected: true, Actual: hc.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AsJsoner(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC AsJsoner", Expected: true, Actual: hc.AsJsoner() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AsJsonParseSelfInjector(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC AsJsonParseSelfInjector", Expected: true, Actual: hc.AsJsonParseSelfInjector() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_HashsetsCollection_AsJsonMarshaller(t *testing.T) {
	hc := corestr.Empty.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "HC AsJsonMarshaller", Expected: true, Actual: hc.AsJsonMarshaller() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// CollectionsOfCollection — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_CollectionsOfCollection_IsEmpty(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	tc := coretestcases.CaseV1{Name: "COC IsEmpty", Expected: true, Actual: coc.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_HasItems(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{Name: "COC HasItems", Expected: true, Actual: coc.HasItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_Length(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{Name: "COC Length", Expected: 1, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AllIndividualItemsLength(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
	coc.Add(corestr.New.Collection.Strings([]string{"c"}))
	tc := coretestcases.CaseV1{Name: "COC AllIndividualItemsLength", Expected: 3, Actual: coc.AllIndividualItemsLength(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_Items(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{Name: "COC Items", Expected: 1, Actual: len(coc.Items()), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_List(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
	result := coc.List(0)
	tc := coretestcases.CaseV1{Name: "COC List", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_List_Empty(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	result := coc.List(0)
	tc := coretestcases.CaseV1{Name: "COC List empty", Expected: 0, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_ToCollection(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	c := coc.ToCollection()
	tc := coretestcases.CaseV1{Name: "COC ToCollection", Expected: 1, Actual: c.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AddStrings(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.AddStrings(false, []string{"a", "b"})
	tc := coretestcases.CaseV1{Name: "COC AddStrings", Expected: 1, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AddStrings_Empty(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.AddStrings(false, []string{})
	tc := coretestcases.CaseV1{Name: "COC AddStrings empty", Expected: 0, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AddsStringsOfStrings(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})
	tc := coretestcases.CaseV1{Name: "COC AddsStringsOfStrings", Expected: 2, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AddsStringsOfStrings_Nil(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.AddsStringsOfStrings(false)
	tc := coretestcases.CaseV1{Name: "COC AddsStringsOfStrings nil", Expected: 0, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_Adds(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	c := *corestr.New.Collection.Strings([]string{"a"})
	coc.Adds(c)
	tc := coretestcases.CaseV1{Name: "COC Adds", Expected: 1, Actual: coc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_String(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	coc.Add(corestr.New.Collection.Strings([]string{"a"}))
	tc := coretestcases.CaseV1{Name: "COC String", Expected: true, Actual: len(coc.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AsJsonContractsBinder(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	tc := coretestcases.CaseV1{Name: "COC AsJsonContractsBinder", Expected: true, Actual: coc.AsJsonContractsBinder() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AsJsoner(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	tc := coretestcases.CaseV1{Name: "COC AsJsoner", Expected: true, Actual: coc.AsJsoner() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AsJsonParseSelfInjector(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	tc := coretestcases.CaseV1{Name: "COC AsJsonParseSelfInjector", Expected: true, Actual: coc.AsJsonParseSelfInjector() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CollectionsOfCollection_AsJsonMarshaller(t *testing.T) {
	coc := corestr.Empty.CollectionsOfCollection()
	tc := coretestcases.CaseV1{Name: "COC AsJsonMarshaller", Expected: true, Actual: coc.AsJsonMarshaller() != nil, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

// ═══════════════════════════════════════════════════════════════
// CharHashsetMap — deeper methods
// ═══════════════════════════════════════════════════════════════

func Test_Cov46_CharHashsetMap_GetChar(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM GetChar", Expected: byte('h'), Actual: chm.GetChar("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_GetChar_Empty(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM GetChar empty", Expected: byte(0), Actual: chm.GetChar(""), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_GetCharOf(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM GetCharOf", Expected: byte('a'), Actual: chm.GetCharOf("abc"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_Add(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	tc := coretestcases.CaseV1{Name: "CHM Add", Expected: true, Actual: chm.Has("hello"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_Add_SameChar(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Add("hi")
	tc := coretestcases.CaseV1{Name: "CHM Add same char", Expected: 1, Actual: chm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_Has_NotFound(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM Has not found", Expected: false, Actual: chm.Has("x"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_HasWithHashset_Found(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	found, hs := chm.HasWithHashset("hello")
	tc := coretestcases.CaseV1{Name: "CHM HasWithHashset found", Expected: true, Actual: found, Args: args.Map{}}
	tc.ShouldBeEqual(t)
	tc2 := coretestcases.CaseV1{Name: "CHM HasWithHashset hs", Expected: true, Actual: hs != nil, Args: args.Map{}}
	tc2.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_HasWithHashset_Empty(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	found, _ := chm.HasWithHashset("x")
	tc := coretestcases.CaseV1{Name: "CHM HasWithHashset empty", Expected: false, Actual: found, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_LengthOf(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Add("hi")
	tc := coretestcases.CaseV1{Name: "CHM LengthOf", Expected: 2, Actual: chm.LengthOf('h'), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_LengthOf_Empty(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM LengthOf empty", Expected: 0, Actual: chm.LengthOf('x'), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_AllLengthsSum(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Add("abc")
	tc := coretestcases.CaseV1{Name: "CHM AllLengthsSum", Expected: 2, Actual: chm.AllLengthsSum(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_AllLengthsSum_Empty(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM AllLengthsSum empty", Expected: 0, Actual: chm.AllLengthsSum(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_List(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Add("abc")
	result := chm.List()
	tc := coretestcases.CaseV1{Name: "CHM List", Expected: 2, Actual: len(result), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_SortedListAsc(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("banana")
	chm.Add("apple")
	result := chm.SortedListAsc()
	tc := coretestcases.CaseV1{Name: "CHM SortedListAsc first", Expected: "apple", Actual: result[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_SortedListDsc(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("banana")
	chm.Add("apple")
	result := chm.SortedListDsc()
	tc := coretestcases.CaseV1{Name: "CHM SortedListDsc first", Expected: "banana", Actual: result[0], Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_IsEquals_Same(t *testing.T) {
	chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
	chm1.Add("hello")
	chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
	chm2.Add("hello")
	tc := coretestcases.CaseV1{Name: "CHM IsEquals same", Expected: true, Actual: chm1.IsEquals(chm2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_IsEquals_Nil(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM IsEquals nil", Expected: false, Actual: chm.IsEquals(nil), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_IsEquals_BothEmpty(t *testing.T) {
	chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
	chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
	tc := coretestcases.CaseV1{Name: "CHM IsEquals both empty", Expected: true, Actual: chm1.IsEquals(chm2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_IsEquals_DiffLength(t *testing.T) {
	chm1 := corestr.New.CharHashsetMap.Cap(10, 5)
	chm1.Add("hello")
	chm2 := corestr.New.CharHashsetMap.Cap(10, 5)
	chm2.Add("hello")
	chm2.Add("abc")
	tc := coretestcases.CaseV1{Name: "CHM IsEquals diff length", Expected: false, Actual: chm1.IsEquals(chm2), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_GetMap(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	m := chm.GetMap()
	tc := coretestcases.CaseV1{Name: "CHM GetMap", Expected: true, Actual: len(m) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_String(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	tc := coretestcases.CaseV1{Name: "CHM String", Expected: true, Actual: len(chm.String()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_SummaryString(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	tc := coretestcases.CaseV1{Name: "CHM SummaryString", Expected: true, Actual: len(chm.SummaryString()) > 0, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_Print_False(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Print(false) // should not panic
	tc := coretestcases.CaseV1{Name: "CHM Print false", Expected: true, Actual: true, Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_LengthOfHashsetFromFirstChar(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Add("hi")
	tc := coretestcases.CaseV1{Name: "CHM LengthOfHashsetFromFirstChar", Expected: 2, Actual: chm.LengthOfHashsetFromFirstChar("h"), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_AddSameStartingCharItems(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.AddSameStartingCharItems('a', []string{"abc", "axy"})
	tc := coretestcases.CaseV1{Name: "CHM AddSameStartingCharItems", Expected: 2, Actual: chm.LengthOf('a'), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_AddSameStartingCharItems_Empty(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.AddSameStartingCharItems('a', []string{})
	tc := coretestcases.CaseV1{Name: "CHM AddSameStartingCharItems empty", Expected: 0, Actual: chm.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_HashsetsCollection(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	hc := chm.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "CHM HashsetsCollection", Expected: true, Actual: hc.HasItems(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_HashsetsCollection_Empty(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	hc := chm.HashsetsCollection()
	tc := coretestcases.CaseV1{Name: "CHM HashsetsCollection empty", Expected: true, Actual: hc.IsEmpty(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_HashsetsCollectionByChars(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	chm.Add("abc")
	hc := chm.HashsetsCollectionByChars('h')
	tc := coretestcases.CaseV1{Name: "CHM HashsetsCollectionByChars", Expected: 1, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}

func Test_Cov46_CharHashsetMap_HashsetsCollectionByStringsFirstChar(t *testing.T) {
	chm := corestr.New.CharHashsetMap.Cap(10, 5)
	chm.Add("hello")
	hc := chm.HashsetsCollectionByStringsFirstChar("hello")
	tc := coretestcases.CaseV1{Name: "CHM HC ByStringsFirstChar", Expected: 1, Actual: hc.Length(), Args: args.Map{}}
	tc.ShouldBeEqual(t)
}
