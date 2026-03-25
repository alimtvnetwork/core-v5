package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
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

	tc.ShouldBeEqualFirst(t, actLines...)
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

	tc.ShouldBeEqualFirst(t, actLines...)
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

	tc.ShouldBeEqualFirst(t, actLines...)
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

	tc.ShouldBeEqualFirst(t, actLines...)
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
	actual := args.Map{
		"removed":   removed,
		"newLength": col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
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
	actual := args.Map{
		"removed":   removed,
		"newLength": col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
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
		actual := args.Map{
			"removedCount": removedCount,
			"newLength":    col.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"uniqueCount": hs.Length(),
			"has1":        hs.Has(1),
			"has2":        hs.Has(2),
			"has3":        hs.Has(3),
			"has99":       hs.Has(99),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"uniqueCount":  unique.Length(),
			"firstElement": unique.First(),
			"lastElement":  unique.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"containsExisting": coregeneric.ContainsSimpleSliceItem(ss, 20),
			"containsMissing":  coregeneric.ContainsSimpleSliceItem(ss, 99),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
