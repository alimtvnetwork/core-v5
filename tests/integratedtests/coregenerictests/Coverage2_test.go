package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── numericfuncs uncovered branches ──

func Test_Cov2_InRangeExclusive_OutOfRange(t *testing.T) {
	actual := args.Map{
		"atMin":    coregeneric.InRangeExclusive(1, 1, 10),
		"atMax":    coregeneric.InRangeExclusive(10, 1, 10),
		"inside":   coregeneric.InRangeExclusive(5, 1, 10),
		"below":    coregeneric.InRangeExclusive(0, 1, 10),
	}
	expected := args.Map{
		"atMin":    false,
		"atMax":    false,
		"inside":   true,
		"below":    false,
	}
	expected.ShouldBeEqual(t, 0, "InRangeExclusive", actual)
}

func Test_Cov2_SafeDivOrDefault(t *testing.T) {
	actual := args.Map{
		"normal": coregeneric.SafeDivOrDefault(10, 3, -1),
		"zero":   coregeneric.SafeDivOrDefault(10, 0, -1),
	}
	expected := args.Map{
		"normal": 3,
		"zero":   -1,
	}
	expected.ShouldBeEqual(t, 0, "SafeDivOrDefault", actual)
}

func Test_Cov2_IsNonNegative(t *testing.T) {
	actual := args.Map{
		"positive": coregeneric.IsNonNegative(5),
		"zero":     coregeneric.IsNonNegative(0),
		"negative": coregeneric.IsNonNegative(-1),
	}
	expected := args.Map{
		"positive": true,
		"zero":     true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNonNegative", actual)
}

func Test_Cov2_Sign(t *testing.T) {
	actual := args.Map{
		"negative": coregeneric.Sign(-5),
		"zero":     coregeneric.Sign(0),
		"positive": coregeneric.Sign(5),
	}
	expected := args.Map{
		"negative": -1,
		"zero":     0,
		"positive": 1,
	}
	expected.ShouldBeEqual(t, 0, "Sign", actual)
}

func Test_Cov2_IsNotEqual(t *testing.T) {
	actual := args.Map{
		"same": coregeneric.IsNotEqual(5, 5),
		"diff": coregeneric.IsNotEqual(5, 6),
	}
	expected := args.Map{
		"same": false,
		"diff": true,
	}
	expected.ShouldBeEqual(t, 0, "IsNotEqual", actual)
}

func Test_Cov2_IsNumericEqual(t *testing.T) {
	actual := args.Map{
		"same": coregeneric.IsNumericEqual(5, 5),
		"diff": coregeneric.IsNumericEqual(5, 6),
	}
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNumericEqual", actual)
}

// ── Collection uncovered branches ──

func Test_Cov2_Collection_Capacity_Nil(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	actual := args.Map{"cap": col.Capacity()}
	expected := args.Map{"cap": 0}
	expected.ShouldBeEqual(t, 0, "Collection Capacity empty", actual)
}

func Test_Cov2_Collection_HasAnyItem(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	actual := args.Map{"result": col.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Collection HasAnyItem", actual)
}

func Test_Cov2_Collection_AddIfMany_Skip(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(false, 1, 2, 3)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddIfMany skip", actual)
}

func Test_Cov2_Collection_AddFunc(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddFunc(func() int { return 42 })
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 42}
	expected.ShouldBeEqual(t, 0, "AddFunc", actual)
}

func Test_Cov2_Collection_CountFunc(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4, 5})
	count := col.CountFunc(func(v int) bool { return v > 3 })
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "CountFunc", actual)
}

func Test_Cov2_Collection_ConcatNew(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	result := col.ConcatNew(3, 4)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ConcatNew", actual)
}

func Test_Cov2_Collection_Reverse(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	col.Reverse()
	actual := args.Map{"first": col.First(), "last": col.Last()}
	expected := args.Map{"first": 3, "last": 1}
	expected.ShouldBeEqual(t, 0, "Reverse", actual)
}

// ── Hashmap uncovered branches ──

func Test_Cov2_Hashmap_Set_ReturnsBool(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("a", 1)
	isUpdate := hm.Set("a", 2)
	actual := args.Map{"isNew": isNew, "isUpdate": isUpdate}
	// Set returns true if newly added
	expected := args.Map{"isNew": true, "isUpdate": false}
	expected.ShouldBeEqual(t, 0, "Hashmap Set return", actual)
}

func Test_Cov2_Hashmap_ForEachBreak(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	hm.Set("b", 2)
	count := 0
	hm.ForEachBreak(func(k string, v int) bool {
		count++
		return true // break immediately
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap ForEachBreak", actual)
}

func Test_Cov2_Hashmap_ConcatNew_NilOther(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	result := hm.ConcatNew(nil)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap ConcatNew nil other", actual)
}

// ── Hashset uncovered branches ──

func Test_Cov2_Hashset_AddBool(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	existed1 := hs.AddBool("a")
	existed2 := hs.AddBool("a")
	actual := args.Map{"first": existed1, "second": existed2}
	expected := args.Map{"first": false, "second": true}
	expected.ShouldBeEqual(t, 0, "Hashset AddBool", actual)
}

func Test_Cov2_Hashset_AddIfMany_Skip(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIfMany(false, "a", "b")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset AddIfMany skip", actual)
}

func Test_Cov2_Hashset_AddItemsMap_FalseValue(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{"a": true, "b": false})
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset AddItemsMap false value", actual)
}

func Test_Cov2_Hashset_Resize_TooSmall(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
	hs.Resize(1) // smaller than current, should not resize
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Hashset Resize too small", actual)
}

// ── MapSimpleSlice nil ──

func Test_Cov2_MapSimpleSlice_Nil(t *testing.T) {
	result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapSimpleSlice nil", actual)
}

// ── DistinctSimpleSlice nil ──

func Test_Cov2_DistinctSimpleSlice_Nil(t *testing.T) {
	result := coregeneric.DistinctSimpleSlice[int](nil)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DistinctSimpleSlice nil", actual)
}

// ── ContainsSimpleSliceItem nil ──

func Test_Cov2_ContainsSimpleSliceItem_Nil(t *testing.T) {
	actual := args.Map{"result": coregeneric.ContainsSimpleSliceItem[int](nil, 1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ContainsSimpleSliceItem nil", actual)
}
