package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── Collection AppendPrependIf ──

func Test_Cov5_AppendPrependIf_True(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.AppendPrependIf(true,
		[]namevalue.Instance[string, string]{{Name: "a", Value: "1"}},
		[]namevalue.Instance[string, string]{{Name: "c", Value: "3"}},
	)
	actual := args.Map{"len": c.Length(), "first": c.Items[0].Name, "last": c.Items[2].Name}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf true -- both applied", actual)
}

func Test_Cov5_AppendPrependIf_False(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.AppendPrependIf(false,
		[]namevalue.Instance[string, string]{{Name: "a", Value: "1"}},
		[]namevalue.Instance[string, string]{{Name: "c", Value: "3"}},
	)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendPrependIf false -- skipped", actual)
}

// ── Collection AddsIf ──

func Test_Cov5_AddsIf_True(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(true, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsIf true -- added", actual)
}

func Test_Cov5_AddsIf_False(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsIf false -- skipped", actual)
}

// ── Collection PrependUsingFuncIf ──

func Test_Cov5_PrependUsingFuncIf_True(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	actual := args.Map{"len": c.Length(), "first": c.Items[0].Name}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf true", actual)
}

func Test_Cov5_PrependUsingFuncIf_False(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	c.PrependUsingFuncIf(false, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf false", actual)
}

func Test_Cov5_PrependUsingFuncIf_NilFunc(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependUsingFuncIf(true, nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependUsingFuncIf nil func", actual)
}

// ── Collection AppendUsingFuncIf ──

func Test_Cov5_AppendUsingFuncIf_True(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.AppendUsingFuncIf(true, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf true", actual)
}

func Test_Cov5_AppendUsingFuncIf_False(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(false, func() []namevalue.Instance[string, string] {
		return []namevalue.Instance[string, string]{{Name: "b", Value: "2"}}
	})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf false", actual)
}

func Test_Cov5_AppendUsingFuncIf_NilFunc(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendUsingFuncIf(true, nil)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendUsingFuncIf nil func", actual)
}

// ── Collection AddsPtr ──

func Test_Cov5_AddsPtr(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	item := namevalue.Instance[string, string]{Name: "a", Value: "1"}
	c.AddsPtr(&item, nil) // nil should be skipped
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddsPtr with nil skip", actual)
}

func Test_Cov5_AddsPtr_Empty(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AddsPtr()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddsPtr empty", actual)
}

// ── Collection ConcatNew / ConcatNewPtr ──

func Test_Cov5_ConcatNew(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	result := c.ConcatNew(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"origLen": c.Length(), "newLen": result.Length()}
	expected := args.Map{"origLen": 1, "newLen": 2}
	expected.ShouldBeEqual(t, 0, "ConcatNew -- immutable", actual)
}

func Test_Cov5_ConcatNewPtr(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	item := namevalue.Instance[string, string]{Name: "b", Value: "2"}
	result := c.ConcatNewPtr(&item, nil)
	actual := args.Map{"newLen": result.Length()}
	expected := args.Map{"newLen": 2}
	expected.ShouldBeEqual(t, 0, "ConcatNewPtr", actual)
}

// ── Collection IsEqualByString ──

func Test_Cov5_IsEqualByString_Equal(t *testing.T) {
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	c2.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqualByString equal", actual)
}

func Test_Cov5_IsEqualByString_DiffLen(t *testing.T) {
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqualByString diff len", actual)
}

func Test_Cov5_IsEqualByString_DiffContent(t *testing.T) {
	c1 := namevalue.NewGenericCollectionDefault[string, string]()
	c1.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c2 := namevalue.NewGenericCollectionDefault[string, string]()
	c2.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqualByString diff content", actual)
}

func Test_Cov5_IsEqualByString_BothNil(t *testing.T) {
	var c1, c2 *namevalue.Collection[string, string]
	actual := args.Map{"result": c1.IsEqualByString(c2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEqualByString both nil", actual)
}

func Test_Cov5_IsEqualByString_OneNil(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	actual := args.Map{"result": c.IsEqualByString(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEqualByString one nil", actual)
}

// ── NewGenericCollectionUsing no-clone ──

func Test_Cov5_NewGenericCollectionUsing_NoClone(t *testing.T) {
	items := []namevalue.Instance[string, string]{{Name: "a", Value: "1"}}
	c := namevalue.NewGenericCollectionUsing[string, string](false, items...)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewGenericCollectionUsing no clone", actual)
}

func Test_Cov5_NewGenericCollectionUsing_NilItems(t *testing.T) {
	c := namevalue.NewGenericCollectionUsing[string, string](true)
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewGenericCollectionUsing nil items", actual)
}

// ── Collection HasIndex / LastIndex / HasAnyItem ──

func Test_Cov5_Collection_IndexMethods(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	actual := args.Map{
		"lastIndex": c.LastIndex(),
		"hasIndex0": c.HasIndex(0),
		"hasIndex1": c.HasIndex(1),
		"hasIndex2": c.HasIndex(2),
		"hasAny":    c.HasAnyItem(),
		"count":     c.Count(),
	}
	expected := args.Map{
		"lastIndex": 1, "hasIndex0": true, "hasIndex1": true,
		"hasIndex2": false, "hasAny": true, "count": 2,
	}
	expected.ShouldBeEqual(t, 0, "Collection index methods", actual)
}

// ── Collection CompiledLazyString nil ──

func Test_Cov5_CompiledLazyString_Nil(t *testing.T) {
	var c *namevalue.Collection[string, string]
	result := c.CompiledLazyString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "CompiledLazyString nil -- empty", actual)
}

// ── Instance IsNull / Dispose ──

func Test_Cov5_Instance_IsNull(t *testing.T) {
	inst := &namevalue.Instance[string, string]{Name: "a", Value: "1"}
	var nilInst *namevalue.Instance[string, string]
	actual := args.Map{"nonNil": inst.IsNull(), "nil": nilInst.IsNull()}
	expected := args.Map{"nonNil": false, "nil": true}
	expected.ShouldBeEqual(t, 0, "Instance IsNull", actual)
}

func Test_Cov5_Instance_Dispose(t *testing.T) {
	inst := &namevalue.Instance[string, string]{Name: "a", Value: "1"}
	inst.Dispose()
	actual := args.Map{"name": inst.Name, "val": inst.Value}
	expected := args.Map{"name": "", "val": ""}
	expected.ShouldBeEqual(t, 0, "Instance Dispose -- zeroed", actual)
}

func Test_Cov5_Instance_Dispose_Nil(t *testing.T) {
	var inst *namevalue.Instance[string, string]
	inst.Dispose() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Instance Dispose nil -- no panic", actual)
}

// ── Append / Prepend / AppendIf / PrependIf empty ──

func Test_Cov5_Append_Empty(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Append()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Append empty -- no-op", actual)
}

func Test_Cov5_Prepend_Empty(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Prepend()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Prepend empty -- no-op", actual)
}

func Test_Cov5_AppendIf_FalseSkip(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.AppendIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendIf false -- skipped", actual)
}

func Test_Cov5_PrependIf_FalseSkip(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.PrependIf(false, namevalue.Instance[string, string]{Name: "a", Value: "1"})
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependIf false -- skipped", actual)
}

// ── Adds empty ──

func Test_Cov5_Adds_Empty(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Adds()
	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Adds empty -- no-op", actual)
}

// ── Collection Join ──

func Test_Cov5_Join(t *testing.T) {
	c := namevalue.NewGenericCollectionDefault[string, string]()
	c.Add(namevalue.Instance[string, string]{Name: "a", Value: "1"})
	c.Add(namevalue.Instance[string, string]{Name: "b", Value: "2"})
	result := c.Join("; ")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection Join", actual)
}
