package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// =============================================================================
// MapCollection
// =============================================================================

func Test_MapCollection_IntToString(t *testing.T) {
	convey.Convey("MapCollection transforms int to string", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)
		result := coregeneric.MapCollection(src, func(i int) string {
			return fmt.Sprintf("v%d", i)
		})
		convey.So(result.Length(), should.Equal, 3)
		convey.So(result.First(), should.Equal, "v1")
		convey.So(result.Last(), should.Equal, "v3")
	})
}

func Test_MapCollection_NilSource(t *testing.T) {
	convey.Convey("MapCollection returns empty on nil source", t, func() {
		result := coregeneric.MapCollection[int, string](nil, func(i int) string { return "" })
		convey.So(result.IsEmpty(), should.BeTrue)
	})
}

func Test_MapCollection_EmptySource(t *testing.T) {
	convey.Convey("MapCollection returns empty on empty source", t, func() {
		src := coregeneric.EmptyCollection[int]()
		result := coregeneric.MapCollection(src, func(i int) string { return "" })
		convey.So(result.IsEmpty(), should.BeTrue)
	})
}

// =============================================================================
// FlatMapCollection
// =============================================================================

func Test_FlatMapCollection(t *testing.T) {
	convey.Convey("FlatMapCollection flattens results", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)
		result := coregeneric.FlatMapCollection(src, func(i int) []int {
			return []int{i, i * 10}
		})
		convey.So(result.Length(), should.Equal, 6)
	})
}

func Test_FlatMapCollection_NilSource(t *testing.T) {
	convey.Convey("FlatMapCollection returns empty on nil", t, func() {
		result := coregeneric.FlatMapCollection[int, int](nil, func(i int) []int { return nil })
		convey.So(result.IsEmpty(), should.BeTrue)
	})
}

// =============================================================================
// ReduceCollection
// =============================================================================

func Test_ReduceCollection_Sum(t *testing.T) {
	convey.Convey("ReduceCollection sums integers", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4)
		sum := coregeneric.ReduceCollection(src, 0, func(acc int, item int) int {
			return acc + item
		})
		convey.So(sum, should.Equal, 10)
	})
}

func Test_ReduceCollection_NilSource(t *testing.T) {
	convey.Convey("ReduceCollection returns initial on nil source", t, func() {
		result := coregeneric.ReduceCollection[int, int](nil, 99, func(a, b int) int { return a + b })
		convey.So(result, should.Equal, 99)
	})
}

func Test_ReduceCollection_StringConcat(t *testing.T) {
	convey.Convey("ReduceCollection concatenates strings", t, func() {
		src := coregeneric.New.Collection.String.Items("a", "b", "c")
		result := coregeneric.ReduceCollection(src, "", func(acc string, item string) string {
			return acc + item
		})
		convey.So(result, should.Equal, "abc")
	})
}

// =============================================================================
// GroupByCollection
// =============================================================================

func Test_GroupByCollection(t *testing.T) {
	convey.Convey("GroupByCollection groups by key", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5, 6)
		groups := coregeneric.GroupByCollection(src, func(i int) string {
			if i%2 == 0 {
				return "even"
			}
			return "odd"
		})
		convey.So(len(groups), should.Equal, 2)
		convey.So(groups["even"].Length(), should.Equal, 3)
		convey.So(groups["odd"].Length(), should.Equal, 3)
	})
}

func Test_GroupByCollection_NilSource(t *testing.T) {
	convey.Convey("GroupByCollection returns empty map on nil", t, func() {
		groups := coregeneric.GroupByCollection[int, string](nil, func(i int) string { return "" })
		convey.So(len(groups), should.Equal, 0)
	})
}

// =============================================================================
// ContainsFunc / ContainsItem
// =============================================================================

func Test_ContainsFunc_Found(t *testing.T) {
	convey.Convey("ContainsFunc returns true when match found", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)
		convey.So(coregeneric.ContainsFunc(src, func(i int) bool { return i == 2 }), should.BeTrue)
	})
}

func Test_ContainsFunc_NotFound(t *testing.T) {
	convey.Convey("ContainsFunc returns false when no match", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)
		convey.So(coregeneric.ContainsFunc(src, func(i int) bool { return i == 99 }), should.BeFalse)
	})
}

func Test_ContainsFunc_NilSource(t *testing.T) {
	convey.Convey("ContainsFunc returns false on nil", t, func() {
		convey.So(coregeneric.ContainsFunc[int](nil, func(i int) bool { return true }), should.BeFalse)
	})
}

func Test_ContainsItem_Found(t *testing.T) {
	convey.Convey("ContainsItem returns true for existing item", t, func() {
		src := coregeneric.New.Collection.String.Items("a", "b", "c")
		convey.So(coregeneric.ContainsItem(src, "b"), should.BeTrue)
	})
}

func Test_ContainsItem_NotFound(t *testing.T) {
	convey.Convey("ContainsItem returns false for missing item", t, func() {
		src := coregeneric.New.Collection.String.Items("a", "b")
		convey.So(coregeneric.ContainsItem(src, "z"), should.BeFalse)
	})
}

func Test_ContainsItem_NilSource(t *testing.T) {
	convey.Convey("ContainsItem returns false on nil", t, func() {
		convey.So(coregeneric.ContainsItem[string](nil, "x"), should.BeFalse)
	})
}

// =============================================================================
// IndexOfFunc / IndexOfItem
// =============================================================================

func Test_IndexOfFunc_Found(t *testing.T) {
	convey.Convey("IndexOfFunc returns correct index", t, func() {
		src := coregeneric.New.Collection.Int.Items(10, 20, 30)
		convey.So(coregeneric.IndexOfFunc(src, func(i int) bool { return i == 20 }), should.Equal, 1)
	})
}

func Test_IndexOfFunc_NotFound(t *testing.T) {
	convey.Convey("IndexOfFunc returns -1 when not found", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)
		convey.So(coregeneric.IndexOfFunc(src, func(i int) bool { return i == 99 }), should.Equal, -1)
	})
}

func Test_IndexOfFunc_NilSource(t *testing.T) {
	convey.Convey("IndexOfFunc returns -1 on nil", t, func() {
		convey.So(coregeneric.IndexOfFunc[int](nil, func(i int) bool { return true }), should.Equal, -1)
	})
}

func Test_IndexOfItem_Found(t *testing.T) {
	convey.Convey("IndexOfItem returns correct index", t, func() {
		src := coregeneric.New.Collection.String.Items("x", "y", "z")
		convey.So(coregeneric.IndexOfItem(src, "z"), should.Equal, 2)
	})
}

func Test_IndexOfItem_NotFound(t *testing.T) {
	convey.Convey("IndexOfItem returns -1 when not found", t, func() {
		src := coregeneric.New.Collection.String.Items("a")
		convey.So(coregeneric.IndexOfItem(src, "q"), should.Equal, -1)
	})
}

// =============================================================================
// Distinct
// =============================================================================

func Test_Distinct(t *testing.T) {
	convey.Convey("Distinct removes duplicates", t, func() {
		src := coregeneric.New.Collection.Int.Items(1, 2, 2, 3, 1, 3)
		result := coregeneric.Distinct(src)
		convey.So(result.Length(), should.Equal, 3)
	})
}

func Test_Distinct_NilSource(t *testing.T) {
	convey.Convey("Distinct returns empty on nil", t, func() {
		result := coregeneric.Distinct[int](nil)
		convey.So(result.IsEmpty(), should.BeTrue)
	})
}

func Test_Distinct_NoDuplicates(t *testing.T) {
	convey.Convey("Distinct preserves unique items", t, func() {
		src := coregeneric.New.Collection.String.Items("a", "b", "c")
		result := coregeneric.Distinct(src)
		convey.So(result.Length(), should.Equal, 3)
	})
}

// =============================================================================
// MapSimpleSlice
// =============================================================================

func Test_MapSimpleSlice(t *testing.T) {
	convey.Convey("MapSimpleSlice transforms elements", t, func() {
		src := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
		result := coregeneric.MapSimpleSlice(src, func(i int) string {
			return fmt.Sprintf("%d", i)
		})
		convey.So(result.Length(), should.Equal, 3)
	})
}

func Test_MapSimpleSlice_NilSource(t *testing.T) {
	convey.Convey("MapSimpleSlice returns empty on nil", t, func() {
		result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })
		convey.So(result.IsEmpty(), should.BeTrue)
	})
}
