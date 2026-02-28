package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: ContainsAll
// ==========================================

func Test_ContainsAll_Verification(t *testing.T) {
	for caseIndex, testCase := range containsAllTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		var result bool
		if caseIndex == 0 {
			result = coregeneric.ContainsAll(col, 1, 3, 5)
		} else {
			result = coregeneric.ContainsAll(col, 1, 2, 99)
		}
		actLines := []string{
			fmt.Sprintf("%v", result),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ContainsAny
// ==========================================

func Test_ContainsAny_Verification(t *testing.T) {
	for caseIndex, testCase := range containsAnyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		var result bool
		if caseIndex == 0 {
			result = coregeneric.ContainsAny(col, 99, 3, 100)
		} else {
			result = coregeneric.ContainsAny(col, 88, 99, 100)
		}
		actLines := []string{
			fmt.Sprintf("%v", result),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: RemoveItem
// ==========================================

func Test_RemoveItem_Verification(t *testing.T) {
	for caseIndex, testCase := range removeItemTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		var removed bool
		if caseIndex == 0 {
			removed = coregeneric.RemoveItem(col, 2)
		} else {
			removed = coregeneric.RemoveItem(col, 99)
		}
		actLines := []string{
			fmt.Sprintf("%v", removed),
			fmt.Sprintf("%d", col.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
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
