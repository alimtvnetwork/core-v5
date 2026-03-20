package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CollectionTypes — factory shortcuts
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_NewStringCollection(t *testing.T) {
	col := coredynamic.NewStringCollection(5)
	col.Add("a").Add("b")
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewStringCollection", actual)
}

func Test_I22_EmptyStringCollection(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	actual := args.Map{"empty": col.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyStringCollection", actual)
}

func Test_I22_NewIntCollection(t *testing.T) {
	col := coredynamic.NewIntCollection(3)
	col.Add(1).Add(2).Add(3)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "NewIntCollection", actual)
}

func Test_I22_EmptyIntCollection(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	actual := args.Map{"empty": col.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyIntCollection", actual)
}

func Test_I22_NewInt64Collection(t *testing.T) {
	col := coredynamic.NewInt64Collection(2)
	col.Add(int64(99))
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewInt64Collection", actual)
}

func Test_I22_NewByteCollection(t *testing.T) {
	col := coredynamic.NewByteCollection(2)
	col.Add(byte(0x41))
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewByteCollection", actual)
}

func Test_I22_NewBoolCollection(t *testing.T) {
	col := coredynamic.NewBoolCollection(2)
	col.Add(true).Add(false)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewBoolCollection", actual)
}

func Test_I22_NewFloat64Collection(t *testing.T) {
	col := coredynamic.NewFloat64Collection(2)
	col.Add(3.14)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewFloat64Collection", actual)
}

func Test_I22_NewAnyMapCollection(t *testing.T) {
	col := coredynamic.NewAnyMapCollection(2)
	col.Add(map[string]any{"k": "v"})
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewAnyMapCollection", actual)
}

func Test_I22_NewStringMapCollection(t *testing.T) {
	col := coredynamic.NewStringMapCollection(2)
	col.Add(map[string]string{"k": "v"})
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewStringMapCollection", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionDistinct
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_Distinct_Duplicates(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a").Add("c").Add("b")
	result := coredynamic.Distinct(col)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Distinct duplicates", actual)
}

func Test_I22_Distinct_Empty(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	result := coredynamic.Distinct(col)
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Distinct empty", actual)
}

func Test_I22_Unique_Alias(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(1)
	result := coredynamic.Unique(col)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Unique alias", actual)
}

func Test_I22_DistinctLock(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("x").Add("x").Add("y")
	result := coredynamic.DistinctLock(col)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DistinctLock", actual)
}

func Test_I22_DistinctCount(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a")
	actual := args.Map{"count": coredynamic.DistinctCount(col)}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "DistinctCount", actual)
}

func Test_I22_DistinctCount_Empty(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	actual := args.Map{"count": coredynamic.DistinctCount(col)}
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "DistinctCount empty", actual)
}

func Test_I22_IsDistinct_True(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	actual := args.Map{"distinct": coredynamic.IsDistinct(col)}
	expected := args.Map{"distinct": true}
	expected.ShouldBeEqual(t, 0, "IsDistinct true", actual)
}

func Test_I22_IsDistinct_False(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(1)
	actual := args.Map{"distinct": coredynamic.IsDistinct(col)}
	expected := args.Map{"distinct": false}
	expected.ShouldBeEqual(t, 0, "IsDistinct false", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionMap — Map, FlatMap, Reduce
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_Map_Transform(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	result := coredynamic.Map(col, func(i int) string {
		return "x"
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Map transform", actual)
}

func Test_I22_Map_Empty(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	result := coredynamic.Map(col, func(i int) string { return "" })
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Map empty", actual)
}

func Test_I22_Map_Nil(t *testing.T) {
	result := coredynamic.Map[int, string](nil, func(i int) string { return "" })
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Map nil", actual)
}

func Test_I22_FlatMap(t *testing.T) {
	col := coredynamic.EmptyCollection[[]string]()
	col.Add([]string{"a", "b"}).Add([]string{"c"})
	result := coredynamic.FlatMap(col, func(s []string) []string { return s })
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "FlatMap", actual)
}

func Test_I22_FlatMap_Empty(t *testing.T) {
	col := coredynamic.EmptyCollection[[]string]()
	result := coredynamic.FlatMap(col, func(s []string) []string { return s })
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FlatMap empty", actual)
}

func Test_I22_FlatMap_Nil(t *testing.T) {
	result := coredynamic.FlatMap[[]string, string](nil, func(s []string) []string { return s })
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FlatMap nil", actual)
}

func Test_I22_Reduce_Sum(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	sum := coredynamic.Reduce(col, 0, func(acc int, item int) int { return acc + item })
	actual := args.Map{"sum": sum}
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "Reduce sum", actual)
}

func Test_I22_Reduce_Empty(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	sum := coredynamic.Reduce(col, 10, func(acc int, item int) int { return acc + item })
	actual := args.Map{"sum": sum}
	expected := args.Map{"sum": 10}
	expected.ShouldBeEqual(t, 0, "Reduce empty returns initial", actual)
}

func Test_I22_Reduce_Nil(t *testing.T) {
	sum := coredynamic.Reduce[int, int](nil, 42, func(acc int, item int) int { return acc + item })
	actual := args.Map{"sum": sum}
	expected := args.Map{"sum": 42}
	expected.ShouldBeEqual(t, 0, "Reduce nil returns initial", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionSearch — Contains, IndexOf, Has, HasAll, LastIndexOf, Count, Lock variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_Contains_Found(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b")
	actual := args.Map{"found": coredynamic.Contains(col, "b")}
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "Contains found", actual)
}

func Test_I22_Contains_NotFound(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a")
	actual := args.Map{"found": coredynamic.Contains(col, "z")}
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "Contains not found", actual)
}

func Test_I22_IndexOf_Found(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("x").Add("y").Add("z")
	actual := args.Map{"idx": coredynamic.IndexOf(col, "y")}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "IndexOf found", actual)
}

func Test_I22_IndexOf_NotFound(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("x")
	actual := args.Map{"idx": coredynamic.IndexOf(col, "z")}
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "IndexOf not found", actual)
}

func Test_I22_Has_Alias(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2)
	actual := args.Map{"has": coredynamic.Has(col, 2)}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Has alias", actual)
}

func Test_I22_HasAll_True(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	actual := args.Map{"all": coredynamic.HasAll(col, 1, 3)}
	expected := args.Map{"all": true}
	expected.ShouldBeEqual(t, 0, "HasAll true", actual)
}

func Test_I22_HasAll_False(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2)
	actual := args.Map{"all": coredynamic.HasAll(col, 1, 9)}
	expected := args.Map{"all": false}
	expected.ShouldBeEqual(t, 0, "HasAll false", actual)
}

func Test_I22_HasAll_Empty(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	actual := args.Map{"all": coredynamic.HasAll(col, 1)}
	expected := args.Map{"all": false}
	expected.ShouldBeEqual(t, 0, "HasAll empty", actual)
}

func Test_I22_LastIndexOf_Found(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a")
	actual := args.Map{"idx": coredynamic.LastIndexOf(col, "a")}
	expected := args.Map{"idx": 2}
	expected.ShouldBeEqual(t, 0, "LastIndexOf found", actual)
}

func Test_I22_LastIndexOf_NotFound(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a")
	actual := args.Map{"idx": coredynamic.LastIndexOf(col, "z")}
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "LastIndexOf not found", actual)
}

func Test_I22_Count_Occurrences(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a").Add("a")
	actual := args.Map{"count": coredynamic.Count(col, "a")}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Count occurrences", actual)
}

func Test_I22_ContainsLock(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("x")
	actual := args.Map{"found": coredynamic.ContainsLock(col, "x")}
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "ContainsLock", actual)
}

func Test_I22_IndexOfLock(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b")
	actual := args.Map{"idx": coredynamic.IndexOfLock(col, "b")}
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "IndexOfLock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionSort
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_SortFunc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	col.SortFunc(func(a, b int) bool { return a < b })
	actual := args.Map{"first": col.First(), "last": col.Last()}
	expected := args.Map{"first": 1, "last": 3}
	expected.ShouldBeEqual(t, 0, "SortFunc", actual)
}

func Test_I22_SortFunc_Single(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1)
	col.SortFunc(func(a, b int) bool { return a < b })
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFunc single", actual)
}

func Test_I22_SortFuncLock(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	col.SortFuncLock(func(a, b int) bool { return a < b })
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFuncLock", actual)
}

func Test_I22_SortedFunc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	sorted := col.SortedFunc(func(a, b int) bool { return a < b })
	actual := args.Map{"origFirst": col.First(), "sortedFirst": sorted.First()}
	expected := args.Map{"origFirst": 3, "sortedFirst": 1}
	expected.ShouldBeEqual(t, 0, "SortedFunc does not mutate", actual)
}

func Test_I22_SortAsc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	coredynamic.SortAsc(col)
	actual := args.Map{"first": col.First(), "last": col.Last()}
	expected := args.Map{"first": 1, "last": 3}
	expected.ShouldBeEqual(t, 0, "SortAsc", actual)
}

func Test_I22_SortDesc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(3).Add(2)
	coredynamic.SortDesc(col)
	actual := args.Map{"first": col.First(), "last": col.Last()}
	expected := args.Map{"first": 3, "last": 1}
	expected.ShouldBeEqual(t, 0, "SortDesc", actual)
}

func Test_I22_SortAscLock(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1)
	coredynamic.SortAscLock(col)
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortAscLock", actual)
}

func Test_I22_SortDescLock(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(3)
	coredynamic.SortDescLock(col)
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortDescLock", actual)
}

func Test_I22_SortedAsc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	sorted := coredynamic.SortedAsc(col)
	actual := args.Map{"origFirst": col.First(), "sortedFirst": sorted.First()}
	expected := args.Map{"origFirst": 3, "sortedFirst": 1}
	expected.ShouldBeEqual(t, 0, "SortedAsc", actual)
}

func Test_I22_SortedDesc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(3).Add(2)
	sorted := coredynamic.SortedDesc(col)
	actual := args.Map{"origFirst": col.First(), "sortedFirst": sorted.First()}
	expected := args.Map{"origFirst": 1, "sortedFirst": 3}
	expected.ShouldBeEqual(t, 0, "SortedDesc", actual)
}

func Test_I22_IsSorted_True(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	actual := args.Map{"sorted": col.IsSorted(func(a, b int) bool { return a < b })}
	expected := args.Map{"sorted": true}
	expected.ShouldBeEqual(t, 0, "IsSorted true", actual)
}

func Test_I22_IsSorted_False(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1)
	actual := args.Map{"sorted": col.IsSorted(func(a, b int) bool { return a < b })}
	expected := args.Map{"sorted": false}
	expected.ShouldBeEqual(t, 0, "IsSorted false", actual)
}

func Test_I22_IsSorted_Single(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1)
	actual := args.Map{"sorted": col.IsSorted(func(a, b int) bool { return a < b })}
	expected := args.Map{"sorted": true}
	expected.ShouldBeEqual(t, 0, "IsSorted single", actual)
}

func Test_I22_IsSortedAsc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	actual := args.Map{"asc": coredynamic.IsSortedAsc(col)}
	expected := args.Map{"asc": true}
	expected.ShouldBeEqual(t, 0, "IsSortedAsc", actual)
}

func Test_I22_IsSortedDesc(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(2).Add(1)
	actual := args.Map{"desc": coredynamic.IsSortedDesc(col)}
	expected := args.Map{"desc": true}
	expected.ShouldBeEqual(t, 0, "IsSortedDesc", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionGroupBy
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_GroupBy(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("ab").Add("ac").Add("bd")
	groups := coredynamic.GroupBy(col, func(s string) string { return string(s[0]) })
	actual := args.Map{"groups": len(groups), "aLen": groups["a"].Length()}
	expected := args.Map{"groups": 2, "aLen": 2}
	expected.ShouldBeEqual(t, 0, "GroupBy", actual)
}

func Test_I22_GroupBy_Empty(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	groups := coredynamic.GroupBy(col, func(s string) string { return s })
	actual := args.Map{"len": len(groups)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupBy empty", actual)
}

func Test_I22_GroupBy_Nil(t *testing.T) {
	groups := coredynamic.GroupBy[string, string](nil, func(s string) string { return s })
	actual := args.Map{"len": len(groups)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupBy nil", actual)
}

func Test_I22_GroupByLock(t *testing.T) {
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3).Add(4)
	groups := coredynamic.GroupByLock(col, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})
	actual := args.Map{"even": groups["even"].Length(), "odd": groups["odd"].Length()}
	expected := args.Map{"even": 2, "odd": 2}
	expected.ShouldBeEqual(t, 0, "GroupByLock", actual)
}

func Test_I22_GroupByCount(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a")
	counts := coredynamic.GroupByCount(col, func(s string) string { return s })
	actual := args.Map{"a": counts["a"], "b": counts["b"]}
	expected := args.Map{"a": 2, "b": 1}
	expected.ShouldBeEqual(t, 0, "GroupByCount", actual)
}

func Test_I22_GroupByCount_Empty(t *testing.T) {
	col := coredynamic.EmptyStringCollection()
	counts := coredynamic.GroupByCount(col, func(s string) string { return s })
	actual := args.Map{"len": len(counts)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupByCount empty", actual)
}

func Test_I22_GroupByCount_Nil(t *testing.T) {
	counts := coredynamic.GroupByCount[string, string](nil, func(s string) string { return s })
	actual := args.Map{"len": len(counts)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupByCount nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectSetFromTo — deeper paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_ReflectSetFromTo_BothNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo(nil, nil)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo both nil", actual)
}

func Test_I22_ReflectSetFromTo_SameType(t *testing.T) {
	from := "hello"
	var to string
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"noErr": err == nil, "val": to}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo same type", actual)
}

func Test_I22_ReflectSetFromTo_SamePointerType(t *testing.T) {
	from := new(string)
	*from = "hello"
	to := new(string)
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"noErr": err == nil, "val": *to}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo same ptr type", actual)
}

func Test_I22_ReflectSetFromTo_ToNonPointer(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("hello", "notpointer")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo to non-pointer", actual)
}

func Test_I22_ReflectSetFromTo_ToNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("hello", nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo to nil", actual)
}

func Test_I22_ReflectSetFromTo_FromNilPtr(t *testing.T) {
	var from *string
	var to string
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo from nil ptr", actual)
}

func Test_I22_ReflectSetFromTo_BytesToStruct(t *testing.T) {
	type Simple struct {
		Name string `json:"name"`
	}
	from := []byte(`{"name":"test"}`)
	var to Simple
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"noErr": err == nil, "name": to.Name}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo bytes to struct", actual)
}

func Test_I22_ReflectSetFromTo_StructToBytes(t *testing.T) {
	type Simple struct {
		Name string `json:"name"`
	}
	from := Simple{Name: "test"}
	var to []byte
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"noErr": err == nil, "hasBytes": len(to) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo struct to bytes", actual)
}

func Test_I22_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	from := "hello"
	var to int
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo type mismatch", actual)
}

func Test_I22_ReflectSetFromTo_IntValue(t *testing.T) {
	from := 42
	var to int
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"noErr": err == nil, "val": to}
	expected := args.Map{"noErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo int value", actual)
}

func Test_I22_ReflectSetFromTo_ReflectType(t *testing.T) {
	_ = reflect.TypeOf("")
	from := true
	var to bool
	err := coredynamic.ReflectSetFromTo(from, &to)
	actual := args.Map{"noErr": err == nil, "val": to}
	expected := args.Map{"noErr": true, "val": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo bool", actual)
}
