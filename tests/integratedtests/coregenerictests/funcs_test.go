package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================================================
// Test: MapCollection
// ==========================================================================

func Test_MapCollection(t *testing.T) {
	// Case 0: int to string
	{
		tc := mapCollectionTestCases[0]
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)
		result := coregeneric.MapCollection(src, func(i int) string { return fmt.Sprintf("v%d", i) })

		actLines := []string{
			fmt.Sprintf("%v", result.Length()),
			result.First(),
			result.Last(),
		}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: nil source
	{
		tc := mapCollectionTestCases[1]
		result := coregeneric.MapCollection[int, string](nil, func(i int) string { return "" })

		actLines := []string{fmt.Sprintf("%v", result.IsEmpty())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: empty source
	{
		tc := mapCollectionTestCases[2]
		src := coregeneric.EmptyCollection[int]()
		result := coregeneric.MapCollection(src, func(i int) string { return "" })

		actLines := []string{fmt.Sprintf("%v", result.IsEmpty())}

		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: FlatMapCollection
// ==========================================================================

func Test_FlatMapCollection(t *testing.T) {
	// Case 0: flattens
	{
		tc := flatMapCollectionTestCases[0]
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)
		result := coregeneric.FlatMapCollection(src, func(i int) []int { return []int{i, i * 10} })

		actLines := []string{fmt.Sprintf("%v", result.Length())}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: nil
	{
		tc := flatMapCollectionTestCases[1]
		result := coregeneric.FlatMapCollection[int, int](nil, func(i int) []int { return nil })

		actLines := []string{fmt.Sprintf("%v", result.IsEmpty())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: ReduceCollection
// ==========================================================================

func Test_ReduceCollection(t *testing.T) {
	// Case 0: sum
	{
		tc := reduceCollectionTestCases[0]
		src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4)
		sum := coregeneric.ReduceCollection(src, 0, func(a, b int) int { return a + b })

		actLines := []string{fmt.Sprintf("%v", sum)}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: nil returns initial
	{
		tc := reduceCollectionTestCases[1]
		result := coregeneric.ReduceCollection[int, int](nil, 99, func(a, b int) int { return a + b })

		actLines := []string{fmt.Sprintf("%v", result)}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: string concat
	{
		tc := reduceCollectionTestCases[2]
		src := coregeneric.New.Collection.String.Items("a", "b", "c")
		result := coregeneric.ReduceCollection(src, "", func(a, b string) string { return a + b })

		actLines := []string{result}

		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: GroupByCollection
// ==========================================================================

func Test_GroupByCollection(t *testing.T) {
	// Case 0: groups
	{
		tc := groupByCollectionTestCases[0]
		src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5, 6)
		groups := coregeneric.GroupByCollection(src, func(i int) string {
			if i%2 == 0 {
				return "even"
			}

			return "odd"
		})

		actLines := []string{
			fmt.Sprintf("%v", len(groups)),
			fmt.Sprintf("%v", groups["even"].Length()),
			fmt.Sprintf("%v", groups["odd"].Length()),
		}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: nil
	{
		tc := groupByCollectionTestCases[1]
		groups := coregeneric.GroupByCollection[int, string](nil, func(i int) string { return "" })

		actLines := []string{fmt.Sprintf("%v", len(groups))}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: ContainsFunc
// ==========================================================================

func Test_ContainsFunc(t *testing.T) {
	// Case 0: found
	{
		tc := containsFuncTestCases[0]
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)

		actLines := []string{fmt.Sprintf("%v", coregeneric.ContainsFunc(src, func(i int) bool { return i == 2 }))}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: not found
	{
		tc := containsFuncTestCases[1]
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)

		actLines := []string{fmt.Sprintf("%v", coregeneric.ContainsFunc(src, func(i int) bool { return i == 99 }))}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: nil
	{
		tc := containsFuncTestCases[2]

		actLines := []string{fmt.Sprintf("%v", coregeneric.ContainsFunc[int](nil, func(i int) bool { return true }))}

		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: ContainsItem
// ==========================================================================

func Test_ContainsItem(t *testing.T) {
	// Case 0: found
	{
		tc := containsItemTestCases[0]
		src := coregeneric.New.Collection.String.Items("a", "b", "c")

		actLines := []string{fmt.Sprintf("%v", coregeneric.ContainsItem(src, "b"))}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: not found
	{
		tc := containsItemTestCases[1]
		src := coregeneric.New.Collection.String.Items("a", "b")

		actLines := []string{fmt.Sprintf("%v", coregeneric.ContainsItem(src, "z"))}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: nil
	{
		tc := containsItemTestCases[2]

		actLines := []string{fmt.Sprintf("%v", coregeneric.ContainsItem[string](nil, "x"))}

		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: IndexOfFunc
// ==========================================================================

func Test_IndexOfFunc(t *testing.T) {
	// Case 0: found
	{
		tc := indexOfFuncTestCases[0]
		src := coregeneric.New.Collection.Int.Items(10, 20, 30)

		actLines := []string{fmt.Sprintf("%v", coregeneric.IndexOfFunc(src, func(i int) bool { return i == 20 }))}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: not found
	{
		tc := indexOfFuncTestCases[1]
		src := coregeneric.New.Collection.Int.Items(1, 2, 3)

		actLines := []string{fmt.Sprintf("%v", coregeneric.IndexOfFunc(src, func(i int) bool { return i == 99 }))}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: nil
	{
		tc := indexOfFuncTestCases[2]

		actLines := []string{fmt.Sprintf("%v", coregeneric.IndexOfFunc[int](nil, func(i int) bool { return true }))}

		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: IndexOfItem
// ==========================================================================

func Test_IndexOfItem(t *testing.T) {
	// Case 0: found
	{
		tc := indexOfItemTestCases[0]
		src := coregeneric.New.Collection.String.Items("x", "y", "z")

		actLines := []string{fmt.Sprintf("%v", coregeneric.IndexOfItem(src, "z"))}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: not found
	{
		tc := indexOfItemTestCases[1]
		src := coregeneric.New.Collection.String.Items("a")

		actLines := []string{fmt.Sprintf("%v", coregeneric.IndexOfItem(src, "q"))}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: Distinct
// ==========================================================================

func Test_Distinct(t *testing.T) {
	// Case 0: removes duplicates
	{
		tc := distinctTestCases[0]
		src := coregeneric.New.Collection.Int.Items(1, 2, 2, 3, 1, 3)

		actLines := []string{fmt.Sprintf("%v", coregeneric.Distinct(src).Length())}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: nil
	{
		tc := distinctTestCases[1]

		actLines := []string{fmt.Sprintf("%v", coregeneric.Distinct[int](nil).IsEmpty())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 2: no duplicates
	{
		tc := distinctTestCases[2]
		src := coregeneric.New.Collection.String.Items("a", "b", "c")

		actLines := []string{fmt.Sprintf("%v", coregeneric.Distinct(src).Length())}

		errcore.AssertDiffOnMismatch(t, 2, tc.Title, actLines, tc.ExpectedInput)
	}
}

// ==========================================================================
// Test: MapSimpleSlice
// ==========================================================================

func Test_MapSimpleSlice(t *testing.T) {
	// Case 0: transforms
	{
		tc := mapSimpleSliceTestCases[0]
		src := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
		result := coregeneric.MapSimpleSlice(src, func(i int) string { return fmt.Sprintf("%d", i) })

		actLines := []string{fmt.Sprintf("%v", result.Length())}

		errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
	}

	// Case 1: nil
	{
		tc := mapSimpleSliceTestCases[1]
		result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })

		actLines := []string{fmt.Sprintf("%v", result.IsEmpty())}

		errcore.AssertDiffOnMismatch(t, 1, tc.Title, actLines, tc.ExpectedInput)
	}
}
