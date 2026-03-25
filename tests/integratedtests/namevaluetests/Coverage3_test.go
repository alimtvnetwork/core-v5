package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ============================================================================
// AppendsIf
// ============================================================================

func Test_Cov3_AppendsIf_True(t *testing.T) {
	items := []namevalue.StringAny{{Name: "a", Value: 1}}
	result := namevalue.AppendsIf(true, items, namevalue.StringAny{Name: "b", Value: 2})
	actual := args.Map{"len": len(result), "last": result[1].Name}
	expected := args.Map{"len": 2, "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendsIf appends -- isAdd true", actual)
}

func Test_Cov3_AppendsIf_False(t *testing.T) {
	items := []namevalue.StringAny{{Name: "a", Value: 1}}
	result := namevalue.AppendsIf(false, items, namevalue.StringAny{Name: "b", Value: 2})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendsIf no-op -- isAdd false", actual)
}

func Test_Cov3_AppendsIf_EmptyAppend(t *testing.T) {
	items := []namevalue.StringAny{{Name: "a", Value: 1}}
	result := namevalue.AppendsIf(true, items)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendsIf no-op -- empty appending items", actual)
}

// ============================================================================
// PrependsIf
// ============================================================================

func Test_Cov3_PrependsIf_True(t *testing.T) {
	items := []namevalue.StringAny{{Name: "b", Value: 2}}
	result := namevalue.PrependsIf(true, items, namevalue.StringAny{Name: "a", Value: 1})
	actual := args.Map{"len": len(result), "first": result[0].Name}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "PrependsIf prepends -- isAdd true", actual)
}

func Test_Cov3_PrependsIf_False(t *testing.T) {
	items := []namevalue.StringAny{{Name: "b", Value: 2}}
	result := namevalue.PrependsIf(false, items, namevalue.StringAny{Name: "a", Value: 1})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependsIf no-op -- isAdd false", actual)
}

// ============================================================================
// NewNameValuesCollection / EmptyNameValuesCollection
// ============================================================================

func Test_Cov3_NewNameValuesCollection(t *testing.T) {
	c := namevalue.NewNameValuesCollection(5)
	actual := args.Map{"len": c.Length(), "isEmpty": c.IsEmpty()}
	expected := args.Map{"len": 0, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewNameValuesCollection creates empty -- cap 5", actual)
}

func Test_Cov3_EmptyNameValuesCollection(t *testing.T) {
	c := namevalue.EmptyNameValuesCollection()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "EmptyNameValuesCollection creates empty -- cap 0", actual)
}

func Test_Cov3_NewNewNameValuesCollectionUsing(t *testing.T) {
	c := namevalue.NewNewNameValuesCollectionUsing(true,
		namevalue.StringAny{Name: "a", Value: 1},
		namevalue.StringAny{Name: "b", Value: 2},
	)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewNewNameValuesCollectionUsing creates with items -- clone", actual)
}

// ============================================================================
// Collection — Count, HasIndex
// ============================================================================

func Test_Cov3_Collection_Count(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	actual := args.Map{"count": c.Count()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Count returns item count -- two items", actual)
}

func Test_Cov3_Collection_HasIndex_True(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	actual := args.Map{"result": c.HasIndex(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasIndex returns true -- index 0 exists", actual)
}

func Test_Cov3_Collection_HasIndex_False(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	actual := args.Map{"result": c.HasIndex(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasIndex returns false -- empty collection", actual)
}

// ============================================================================
// Collection — AppendIf / PrependIf
// ============================================================================

func Test_Cov3_Collection_AppendIf_True(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendIf(true, namevalue.StringInt{Name: "a", Value: 1})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendIf appends -- isAppend true", actual)
}

func Test_Cov3_Collection_AppendIf_False(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendIf(false, namevalue.StringInt{Name: "a", Value: 1})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendIf no-op -- isAppend false", actual)
}

func Test_Cov3_Collection_PrependIf_True(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	c.PrependIf(true, namevalue.StringInt{Name: "a", Value: 1})
	actual := args.Map{"first": c.Items[0].Name, "len": c.Length()}
	expected := args.Map{"first": "a", "len": 2}
	expected.ShouldBeEqual(t, 0, "PrependIf prepends -- isPrepend true", actual)
}

func Test_Cov3_Collection_PrependIf_False(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	c.PrependIf(false, namevalue.StringInt{Name: "a", Value: 1})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependIf no-op -- isPrepend false", actual)
}

// ============================================================================
// Collection — AddsIf
// ============================================================================

func Test_Cov3_Collection_AddsIf_True(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.AddsIf(true, namevalue.StringInt{Name: "a", Value: 1}, namevalue.StringInt{Name: "b", Value: 2})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddsIf adds items -- isAdd true", actual)
}

func Test_Cov3_Collection_AddsIf_False(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.AddsIf(false, namevalue.StringInt{Name: "a", Value: 1})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsIf no-op -- isAdd false", actual)
}

// ============================================================================
// Collection — AddsPtr
// ============================================================================

func Test_Cov3_Collection_AddsPtr(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	item := namevalue.StringInt{Name: "a", Value: 1}
	c.AddsPtr(&item, nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsPtr adds non-nil skips nil -- mixed", actual)
}

// ============================================================================
// Collection — PrependUsingFuncIf / AppendUsingFuncIf
// ============================================================================

func Test_Cov3_Collection_PrependUsingFuncIf_True(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	c.PrependUsingFuncIf(true, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})
	actual := args.Map{"first": c.Items[0].Name, "len": c.Length()}
	expected := args.Map{"first": "a", "len": 2}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf prepends -- true", actual)
}

func Test_Cov3_Collection_PrependUsingFuncIf_False(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.PrependUsingFuncIf(false, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf no-op -- false", actual)
}

func Test_Cov3_Collection_PrependUsingFuncIf_NilFunc(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.PrependUsingFuncIf(true, nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf no-op -- nil func", actual)
}

func Test_Cov3_Collection_AppendUsingFuncIf_True(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendUsingFuncIf(true, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf appends -- true", actual)
}

func Test_Cov3_Collection_AppendUsingFuncIf_False(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.AppendUsingFuncIf(false, func() []namevalue.StringInt {
		return []namevalue.StringInt{{Name: "a", Value: 1}}
	})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf no-op -- false", actual)
}

// ============================================================================
// Collection — AppendPrependIf
// ============================================================================

func Test_Cov3_Collection_AppendPrependIf_True(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "mid", Value: 0})
	c.AppendPrependIf(true,
		[]namevalue.StringInt{{Name: "first", Value: 1}},
		[]namevalue.StringInt{{Name: "last", Value: 2}},
	)
	actual := args.Map{"len": c.Length(), "first": c.Items[0].Name, "last": c.Items[2].Name}
	expected := args.Map{"len": 3, "first": "first", "last": "last"}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf prepends and appends -- true", actual)
}

func Test_Cov3_Collection_AppendPrependIf_False(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "mid", Value: 0})
	c.AppendPrependIf(false,
		[]namevalue.StringInt{{Name: "first", Value: 1}},
		[]namevalue.StringInt{{Name: "last", Value: 2}},
	)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf no-op -- false", actual)
}

// ============================================================================
// Collection — CompiledLazyString
// ============================================================================

func Test_Cov3_Collection_CompiledLazyString(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	first := c.CompiledLazyString()
	second := c.CompiledLazyString()
	actual := args.Map{"same": first == second, "hasContent": len(first) > 0}
	expected := args.Map{"same": true, "hasContent": true}
	expected.ShouldBeEqual(t, 0, "CompiledLazyString caches result -- second call same", actual)
}

func Test_Cov3_Collection_CompiledLazyString_Nil(t *testing.T) {
	var c *namevalue.Collection[string, int]
	result := c.CompiledLazyString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "CompiledLazyString returns empty -- nil receiver", actual)
}

// ============================================================================
// Collection — ConcatNewPtr
// ============================================================================

func Test_Cov3_Collection_ConcatNewPtr(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	item := namevalue.StringInt{Name: "b", Value: 2}
	result := c.ConcatNewPtr(&item)
	actual := args.Map{"origLen": c.Length(), "newLen": result.Length()}
	expected := args.Map{"origLen": 1, "newLen": 2}
	expected.ShouldBeEqual(t, 0, "ConcatNewPtr creates new collection with added ptr -- clone + add", actual)
}

// ============================================================================
// Collection — JsonStrings / JoinJsonStrings / JoinCsv / JoinCsvLine
// ============================================================================

func Test_Cov3_Collection_JsonStrings(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.JsonStrings()
	actual := args.Map{"len": len(result), "hasContent": len(result[0]) > 0}
	expected := args.Map{"len": 1, "hasContent": true}
	expected.ShouldBeEqual(t, 0, "JsonStrings returns json strings -- one item", actual)
}

func Test_Cov3_Collection_JoinJsonStrings(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	result := c.JoinJsonStrings(",")
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JoinJsonStrings returns joined json -- comma joiner", actual)
}

func Test_Cov3_Collection_JoinCsv(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.JoinCsv()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JoinCsv returns csv string -- one item", actual)
}

func Test_Cov3_Collection_JoinCsvLine(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	c.Add(namevalue.StringInt{Name: "b", Value: 2})
	result := c.JoinCsvLine()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JoinCsvLine returns csv lines -- two items", actual)
}

// ============================================================================
// Collection — JsonString / CsvStrings
// ============================================================================

func Test_Cov3_Collection_JsonString(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.JsonString()
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns json -- one item", actual)
}

func Test_Cov3_Collection_JsonString_Empty(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	result := c.JsonString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JsonString returns empty -- empty collection", actual)
}

func Test_Cov3_Collection_CsvStrings_Empty(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	result := c.CsvStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns empty slice -- empty collection", actual)
}

func Test_Cov3_Collection_CsvStrings_NonEmpty(t *testing.T) {
	c := namevalue.NewGenericCollection[string, int](5)
	c.Add(namevalue.StringInt{Name: "a", Value: 1})
	result := c.CsvStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CsvStrings returns quoted strings -- one item", actual)
}

// ============================================================================
// EmptyGenericCollection
// ============================================================================

func Test_Cov3_EmptyGenericCollection(t *testing.T) {
	c := namevalue.EmptyGenericCollection[string, int]()
	actual := args.Map{"len": c.Length(), "isEmpty": c.IsEmpty()}
	expected := args.Map{"len": 0, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyGenericCollection creates empty -- zero length", actual)
}
