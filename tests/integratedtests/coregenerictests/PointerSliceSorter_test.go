package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

func intPtr(v int) *int       { return &v }
func strPtr(v string) *string { return &v }

func ptrStr[T any](p *T) string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", *p)
}

func ptrSliceToStrings[T any](items []*T) []string {
	result := make([]string, len(items))
	for i, p := range items {
		result[i] = ptrStr(p)
	}
	return result
}

// ==========================================================================
// Test: Ascending sort
// ==========================================================================

func Test_PointerSliceSorter_Asc(t *testing.T) {
	// Case 0: int pointers
	{
		tc := ptrSorterAscTestCases[0]
		items := []*int{intPtr(3), intPtr(1), intPtr(5), intPtr(2), intPtr(4)}
		sorter := coregeneric.NewPointerSliceSorterAsc(items)
		sorter.Sort()
		tc.ShouldBeEqual(t, 0, ptrSliceToStrings(sorter.Items())...)
	}

	// Case 1: string pointers
	{
		tc := ptrSorterAscTestCases[1]
		items := []*string{strPtr("cherry"), strPtr("apple"), strPtr("banana")}
		sorter := coregeneric.NewPointerSliceSorterAsc(items)
		sorter.Sort()
		tc.ShouldBeEqual(t, 1, ptrSliceToStrings(sorter.Items())...)
	}
}

// ==========================================================================
// Test: Descending sort
// ==========================================================================

func Test_PointerSliceSorter_Desc(t *testing.T) {
	tc := ptrSorterDescTestCases[0]
	items := []*int{intPtr(3), intPtr(1), intPtr(5), intPtr(2), intPtr(4)}
	sorter := coregeneric.NewPointerSliceSorterDesc(items)
	sorter.Sort()
	tc.ShouldBeEqual(t, 0, ptrSliceToStrings(sorter.Items())...)
}

// ==========================================================================
// Test: Nil handling
// ==========================================================================

func Test_PointerSliceSorter_NilHandling(t *testing.T) {
	// Case 0: nils to end (default asc)
	{
		tc := ptrSorterNilHandlingTestCases[0]
		items := []*int{nil, intPtr(3), intPtr(1), nil, intPtr(5)}
		sorter := coregeneric.NewPointerSliceSorterAsc(items)
		sorter.Sort()
		tc.ShouldBeEqual(t, 0, ptrSliceToStrings(sorter.Items())...)
	}

	// Case 1: nilFirst=true
	{
		tc := ptrSorterNilHandlingTestCases[1]
		items := []*int{intPtr(3), nil, intPtr(1), nil, intPtr(5)}
		sorter := coregeneric.NewPointerSliceSorterFunc(items, func(a, b int) bool {
			return a < b
		}, true)
		sorter.Sort()
		tc.ShouldBeEqual(t, 1, ptrSliceToStrings(sorter.Items())...)
	}

	// Case 2: all nils
	{
		tc := ptrSorterNilHandlingTestCases[2]
		items := []*int{nil, nil, nil}
		sorter := coregeneric.NewPointerSliceSorterAsc(items)
		sorter.Sort()
		tc.ShouldBeEqual(t, 2, ptrSliceToStrings(sorter.Items())...)
	}
}

// ==========================================================================
// Test: Custom Less function
// ==========================================================================

func Test_PointerSliceSorter_CustomLess(t *testing.T) {
	tc := ptrSorterCustomLessTestCases[0]
	items := []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5)}

	// Sort by distance from 3 (ascending)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	sorter := coregeneric.NewPointerSliceSorterFunc(items, func(a, b int) bool {
		return abs(a-3) < abs(b-3)
	}, false)
	sorter.Sort()
	tc.ShouldBeEqual(t, 0, ptrSliceToStrings(sorter.Items())...)
}

// ==========================================================================
// Test: SetAsc / SetDesc switching
// ==========================================================================

func Test_PointerSliceSorter_Switch(t *testing.T) {
	tc := ptrSorterSwitchTestCases[0]
	items := []*int{intPtr(3), intPtr(1), intPtr(5), intPtr(2), intPtr(4)}

	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()
	firstAfterAsc := ptrStr(sorter.Items()[0])
	lastAfterAsc := ptrStr(sorter.Items()[4])

	sorter.SetDesc().Sort()
	firstAfterDesc := ptrStr(sorter.Items()[0])
	lastAfterDesc := ptrStr(sorter.Items()[4])

	tc.ShouldBeEqual(t, 0, firstAfterAsc, lastAfterAsc, firstAfterDesc, lastAfterDesc)
}

// ==========================================================================
// Test: IsSorted
// ==========================================================================

func Test_PointerSliceSorter_IsSorted(t *testing.T) {
	tc := ptrSorterIsSortedTestCases[0]
	items := []*int{intPtr(3), intPtr(1), intPtr(5)}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)

	beforeSort := fmt.Sprintf("%v", sorter.IsSorted())
	sorter.Sort()
	afterSort := fmt.Sprintf("%v", sorter.IsSorted())

	tc.ShouldBeEqual(t, 0, beforeSort, afterSort)
}

// ==========================================================================
// Test: Edge cases (empty, single, nil slice)
// ==========================================================================

func Test_PointerSliceSorter_Edge(t *testing.T) {
	// Case 0: empty slice
	{
		tc := ptrSorterEdgeTestCases[0]
		items := []*int{}
		sorter := coregeneric.NewPointerSliceSorterAsc(items)
		sorter.Sort()
		tc.ShouldBeEqual(t, 0,
			fmt.Sprintf("%d", sorter.Len()),
			fmt.Sprintf("%v", sorter.IsSorted()),
		)
	}

	// Case 1: single element
	{
		tc := ptrSorterEdgeTestCases[1]
		items := []*int{intPtr(42)}
		sorter := coregeneric.NewPointerSliceSorterAsc(items)
		sorter.Sort()
		tc.ShouldBeEqual(t, 1,
			fmt.Sprintf("%d", sorter.Len()),
			fmt.Sprintf("%v", sorter.IsSorted()),
			ptrStr(sorter.Items()[0]),
		)
	}

	// Case 2: nil slice
	{
		tc := ptrSorterEdgeTestCases[2]
		sorter := coregeneric.NewPointerSliceSorterAsc[int](nil)
		tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", sorter.Len()))
	}
}

// ==========================================================================
// Test: SetItems / Items
// ==========================================================================

func Test_PointerSliceSorter_SetItems(t *testing.T) {
	tc := ptrSorterSetItemsTestCases[0]
	sorter := coregeneric.NewPointerSliceSorterAsc([]*int{intPtr(5), intPtr(1)})
	sorter.Sort()

	// Replace items and re-sort
	newItems := []*int{intPtr(30), intPtr(10), intPtr(20)}
	sorter.SetItems(newItems).Sort()

	tc.ShouldBeEqual(t, 0,
		fmt.Sprintf("%d", sorter.Len()),
		ptrSliceToStrings(sorter.Items())[0],
		ptrSliceToStrings(sorter.Items())[1],
		ptrSliceToStrings(sorter.Items())[2],
	)
}

// ==========================================================================
// Test: Chaining
// ==========================================================================

func Test_PointerSliceSorter_Chaining(t *testing.T) {
	tc := ptrSorterChainingTestCases[0]
	items := []*int{intPtr(3), nil, intPtr(1), intPtr(5)}

	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.SetDesc().SetNilFirst(true).Sort()

	tc.ShouldBeEqual(t, 0, ptrSliceToStrings(sorter.Items())...)
}
