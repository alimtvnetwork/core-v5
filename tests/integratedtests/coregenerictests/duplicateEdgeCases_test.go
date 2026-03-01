package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: Distinct — all same value
// ==========================================

func Test_Distinct_AllSameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range distinctAllSameTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		unique := coregeneric.Distinct(col)

		var actLines []string
		if unique.IsEmpty() {
			actLines = []string{
				fmt.Sprintf("%d", unique.Length()),
				fmt.Sprintf("%v", unique.IsEmpty()),
			}
		} else {
			actLines = []string{
				fmt.Sprintf("%d", unique.Length()),
				fmt.Sprintf("%d", unique.First()),
				fmt.Sprintf("%d", unique.Last()),
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: RemoveItem — all same value
// ==========================================

func Test_RemoveItem_AllSameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range removeItemAllSameTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		removed := coregeneric.RemoveItem(col, 3)
		actLines := []string{
			fmt.Sprintf("%v", removed),
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: RemoveAllItems — all same value
// ==========================================

func Test_RemoveAllItems_AllSameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range removeAllItemsAllSameTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		removedCount := coregeneric.RemoveAllItems(col, 3)
		actLines := []string{
			fmt.Sprintf("%d", removedCount),
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ContainsAll / ContainsAny — all same value
// ==========================================

func Test_ContainsAllAny_AllSameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range containsAllSameTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%v", coregeneric.ContainsAll(col, 5)),
			fmt.Sprintf("%v", coregeneric.ContainsAll(col, 5, 6)),
			fmt.Sprintf("%v", coregeneric.ContainsAny(col, 5, 99)),
			fmt.Sprintf("%v", coregeneric.ContainsAny(col, 88, 99)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ToHashset — all same value
// ==========================================

func Test_ToHashset_AllSameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range toHashsetAllSameTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		hs := coregeneric.ToHashset(col)
		actLines := []string{
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(9)),
			fmt.Sprintf("%v", hs.Has(99)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashset.From — all duplicates
// ==========================================

func Test_Hashset_FromAllDuplicates_Verification(t *testing.T) {
	for caseIndex, testCase := range hashsetAddDuplicatesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		hs := coregeneric.New.Hashset.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("x")),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashset.AddBool — repeated adds
// ==========================================

func Test_Hashset_AddBoolRepeated_Verification(t *testing.T) {
	for caseIndex, testCase := range hashsetAddBoolDuplicatesTestCases {
		// Arrange

		// Act
		hs := coregeneric.New.Hashset.Int.Empty()
		add1 := hs.AddBool(42)
		add2 := hs.AddBool(42)
		add3 := hs.AddBool(42)
		add4 := hs.AddBool(42)
		actLines := []string{
			fmt.Sprintf("%v", add1),
			fmt.Sprintf("%v", add2),
			fmt.Sprintf("%v", add3),
			fmt.Sprintf("%v", add4),
			fmt.Sprintf("%d", hs.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: DistinctSimpleSlice — all same value
// ==========================================

func Test_DistinctSimpleSlice_AllSameValue_Verification(t *testing.T) {
	for caseIndex, testCase := range distinctSimpleSliceAllSameTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		unique := coregeneric.DistinctSimpleSlice(ss)

		var actLines []string
		if unique.IsEmpty() {
			actLines = []string{
				fmt.Sprintf("%d", unique.Length()),
				fmt.Sprintf("%v", unique.IsEmpty()),
			}
		} else {
			actLines = []string{
				fmt.Sprintf("%d", unique.Length()),
				fmt.Sprintf("%d", unique.First()),
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
