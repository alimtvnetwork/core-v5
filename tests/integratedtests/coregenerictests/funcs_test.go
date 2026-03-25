package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: MapCollection
// ==========================================================================

func Test_MapCollection_IntToString(t *testing.T) {
	tc := mapCollectionIntToStringTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)
	result := coregeneric.MapCollection(src, func(i int) string { return fmt.Sprintf("v%d", i) })

	actual := args.Map{
		"length": result.Length(),
		"first":  result.First(),
		"last":   result.Last(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapCollection_NilSource(t *testing.T) {
	tc := mapCollectionNilSourceTestCase
	result := coregeneric.MapCollection[int, string](nil, func(i int) string { return "" })

	actual := args.Map{"isEmpty": result.IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapCollection_EmptySource(t *testing.T) {
	tc := mapCollectionEmptySourceTestCase
	src := coregeneric.EmptyCollection[int]()
	result := coregeneric.MapCollection(src, func(i int) string { return "" })

	actual := args.Map{"isEmpty": result.IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: FlatMapCollection
// ==========================================================================

func Test_FlatMapCollection_Flattens(t *testing.T) {
	tc := flatMapCollectionFlattensTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)
	result := coregeneric.FlatMapCollection(src, func(i int) []int { return []int{i, i * 10} })

	actual := args.Map{"length": result.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FlatMapCollection_Nil(t *testing.T) {
	tc := flatMapCollectionNilTestCase
	result := coregeneric.FlatMapCollection[int, int](nil, func(i int) []int { return nil })

	actual := args.Map{"isEmpty": result.IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ReduceCollection
// ==========================================================================

func Test_ReduceCollection_Sum(t *testing.T) {
	tc := reduceCollectionSumTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4)
	sum := coregeneric.ReduceCollection(src, 0, func(a, b int) int { return a + b })

	actual := args.Map{"result": sum}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ReduceCollection_Nil(t *testing.T) {
	tc := reduceCollectionNilTestCase
	result := coregeneric.ReduceCollection[int, int](nil, 99, func(a, b int) int { return a + b })

	actual := args.Map{"result": result}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ReduceCollection_Concat(t *testing.T) {
	tc := reduceCollectionConcatTestCase
	src := coregeneric.New.Collection.String.Items("a", "b", "c")
	result := coregeneric.ReduceCollection(src, "", func(a, b string) string { return a + b })

	actual := args.Map{"result": result}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: GroupByCollection
// ==========================================================================

func Test_GroupByCollection_Groups(t *testing.T) {
	tc := groupByCollectionGroupsTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5, 6)
	groups := coregeneric.GroupByCollection(src, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})

	actual := args.Map{
		"groupCount": len(groups),
		"evenCount":  groups["even"].Length(),
		"oddCount":   groups["odd"].Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_GroupByCollection_Nil(t *testing.T) {
	tc := groupByCollectionNilTestCase
	groups := coregeneric.GroupByCollection[int, string](nil, func(i int) string { return "" })

	actual := args.Map{"groupCount": len(groups)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ContainsFunc
// ==========================================================================

func Test_ContainsFunc_Found(t *testing.T) {
	tc := containsFuncFoundTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)

	actual := args.Map{"result": coregeneric.ContainsFunc(src, func(i int) bool { return i == 2 })}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsFunc_NotFound(t *testing.T) {
	tc := containsFuncNotFoundTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)

	actual := args.Map{"result": coregeneric.ContainsFunc(src, func(i int) bool { return i == 99 })}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsFunc_Nil(t *testing.T) {
	tc := containsFuncNilTestCase

	actual := args.Map{"result": coregeneric.ContainsFunc[int](nil, func(i int) bool { return true })}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ContainsItem
// ==========================================================================

func Test_ContainsItem_Found(t *testing.T) {
	tc := containsItemFoundTestCase
	src := coregeneric.New.Collection.String.Items("a", "b", "c")

	actual := args.Map{"result": coregeneric.ContainsItem(src, "b")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsItem_NotFound(t *testing.T) {
	tc := containsItemNotFoundTestCase
	src := coregeneric.New.Collection.String.Items("a", "b")

	actual := args.Map{"result": coregeneric.ContainsItem(src, "z")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsItem_Nil(t *testing.T) {
	tc := containsItemNilTestCase

	actual := args.Map{"result": coregeneric.ContainsItem[string](nil, "x")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IndexOfFunc
// ==========================================================================

func Test_IndexOfFunc_Found(t *testing.T) {
	tc := indexOfFuncFoundTestCase
	src := coregeneric.New.Collection.Int.Items(10, 20, 30)

	actual := args.Map{"index": coregeneric.IndexOfFunc(src, func(i int) bool { return i == 20 })}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IndexOfFunc_NotFound(t *testing.T) {
	tc := indexOfFuncNotFoundTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)

	actual := args.Map{"index": coregeneric.IndexOfFunc(src, func(i int) bool { return i == 99 })}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IndexOfFunc_Nil(t *testing.T) {
	tc := indexOfFuncNilTestCase

	actual := args.Map{"index": coregeneric.IndexOfFunc[int](nil, func(i int) bool { return true })}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IndexOfItem
// ==========================================================================

func Test_IndexOfItem_Found(t *testing.T) {
	tc := indexOfItemFoundTestCase
	src := coregeneric.New.Collection.String.Items("x", "y", "z")

	actual := args.Map{"index": coregeneric.IndexOfItem(src, "z")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IndexOfItem_NotFound(t *testing.T) {
	tc := indexOfItemNotFoundTestCase
	src := coregeneric.New.Collection.String.Items("a")

	actual := args.Map{"index": coregeneric.IndexOfItem(src, "q")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Distinct
// ==========================================================================

func Test_Distinct_RemovesDuplicates(t *testing.T) {
	tc := distinctRemovesDuplicatesTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 2, 3, 1, 3)

	actual := args.Map{"length": coregeneric.Distinct(src).Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Distinct_Nil(t *testing.T) {
	tc := distinctNilTestCase

	actual := args.Map{"isEmpty": coregeneric.Distinct[int](nil).IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Distinct_NoDuplicates(t *testing.T) {
	tc := distinctNoDuplicatesTestCase
	src := coregeneric.New.Collection.String.Items("a", "b", "c")

	actual := args.Map{"length": coregeneric.Distinct(src).Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: MapSimpleSlice
// ==========================================================================

func Test_MapSimpleSlice_Transforms(t *testing.T) {
	tc := mapSimpleSliceTransformsTestCase
	src := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	result := coregeneric.MapSimpleSlice(src, func(i int) string { return fmt.Sprintf("%d", i) })

	actual := args.Map{"length": result.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapSimpleSlice_Nil(t *testing.T) {
	tc := mapSimpleSliceNilTestCase
	result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })

	actual := args.Map{"isEmpty": result.IsEmpty()}

	tc.ShouldBeEqualMapFirst(t, actual)
}
