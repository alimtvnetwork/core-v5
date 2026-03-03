package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/errcore"
)

// ==========================================
// Test: ContainsAll
// ==========================================

func Test_ContainsAll_True(t *testing.T) {
	tc := containsAllTrueTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)
	searchItems := input["searchItems"].([]int)

	// Act
	col := coregeneric.New.Collection.Int.Items(items...)
	result := coregeneric.ContainsAll(col, searchItems...)
	actLines := []string{
		fmt.Sprintf("%v", result),
	}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_ContainsAll_False(t *testing.T) {
	tc := containsAllFalseTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)
	searchItems := input["searchItems"].([]int)

	// Act
	col := coregeneric.New.Collection.Int.Items(items...)
	result := coregeneric.ContainsAll(col, searchItems...)
	actLines := []string{
		fmt.Sprintf("%v", result),
	}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: ContainsAny
// ==========================================

func Test_ContainsAny_True(t *testing.T) {
	tc := containsAnyTrueTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)
	searchItems := input["searchItems"].([]int)

	// Act
	col := coregeneric.New.Collection.Int.Items(items...)
	result := coregeneric.ContainsAny(col, searchItems...)
	actLines := []string{
		fmt.Sprintf("%v", result),
	}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_ContainsAny_False(t *testing.T) {
	tc := containsAnyFalseTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)
	searchItems := input["searchItems"].([]int)

	// Act
	col := coregeneric.New.Collection.Int.Items(items...)
	result := coregeneric.ContainsAny(col, searchItems...)
	actLines := []string{
		fmt.Sprintf("%v", result),
	}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: RemoveItem
// ==========================================

func Test_RemoveItem_Found(t *testing.T) {
	tc := removeItemFoundTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)
	removeItem := input["removeItem"].(int)

	// Act
	col := coregeneric.New.Collection.Int.Items(items...)
	removed := coregeneric.RemoveItem(col, removeItem)
	actLines := []string{
		fmt.Sprintf("%v", removed),
		fmt.Sprintf("%d", col.Length()),
	}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

func Test_RemoveItem_Missing(t *testing.T) {
	tc := removeItemMissingTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)
	removeItem := input["removeItem"].(int)

	// Act
	col := coregeneric.New.Collection.Int.Items(items...)
	removed := coregeneric.RemoveItem(col, removeItem)
	actLines := []string{
		fmt.Sprintf("%v", removed),
		fmt.Sprintf("%d", col.Length()),
	}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, tc.Title, actLines, tc.ExpectedInput)
}

// ==========================================
// Test: RemoveAllItems
// ==========================================

func Test_RemoveAllItems_Verification(t *testing.T) {
	for caseIndex, testCase := range removeAllItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		removedCount := coregeneric.RemoveAllItems(col, 2)
		actLines := []string{
			fmt.Sprintf("%d", removedCount),
			fmt.Sprintf("%d", col.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ToHashset
// ==========================================

func Test_ToHashset_Verification(t *testing.T) {
	for caseIndex, testCase := range toHashsetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		hs := coregeneric.ToHashset(col)
		actLines := []string{
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(1)),
			fmt.Sprintf("%v", hs.Has(2)),
			fmt.Sprintf("%v", hs.Has(3)),
			fmt.Sprintf("%v", hs.Has(99)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: DistinctSimpleSlice
// ==========================================

func Test_DistinctSimpleSlice_Verification(t *testing.T) {
	for caseIndex, testCase := range distinctSimpleSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		unique := coregeneric.DistinctSimpleSlice(ss)
		actLines := []string{
			fmt.Sprintf("%d", unique.Length()),
			fmt.Sprintf("%d", unique.First()),
			fmt.Sprintf("%d", unique.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ContainsSimpleSliceItem
// ==========================================

func Test_ContainsSimpleSliceItem_Verification(t *testing.T) {
	for caseIndex, testCase := range containsSimpleSliceItemTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%v", coregeneric.ContainsSimpleSliceItem(ss, 20)),
			fmt.Sprintf("%v", coregeneric.ContainsSimpleSliceItem(ss, 99)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
