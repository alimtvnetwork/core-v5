package coregenerictests

import (
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
		actual := args.Map{
			"length":         unique.Length(),
			"isEmpty":        unique.IsEmpty(),
			"firstOrDefault": unique.FirstOrDefault(),
			"lastOrDefault":  unique.LastOrDefault(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"removed": removed,
			"length":  col.Length(),
			"first":   col.First(),
			"last":    col.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"removedCount": removedCount,
			"length":       col.Length(),
			"isEmpty":      col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"containsAllSingle":    coregeneric.ContainsAll(col, 5),
			"containsAllWithOther": coregeneric.ContainsAll(col, 5, 6),
			"containsAnyWithMatch": coregeneric.ContainsAny(col, 5, 99),
			"containsAnyNoMatch":   coregeneric.ContainsAny(col, 88, 99),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": hs.Length(),
			"has9":   hs.Has(9),
			"has99":  hs.Has(99),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": hs.Length(),
			"hasX":   hs.Has("x"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"add1":   add1,
			"add2":   add2,
			"add3":   add3,
			"add4":   add4,
			"length": hs.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: DistinctSimpleSlice — all same (non-empty)
// ==========================================

func Test_DistinctSimpleSlice_AllSame_NonEmpty(t *testing.T) {
	tc := distinctSimpleSliceAllSameNonEmptyTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)

	// Act
	ss := coregeneric.New.SimpleSlice.Int.Items(items...)
	unique := coregeneric.DistinctSimpleSlice(ss)
	actual := args.Map{
		"length": unique.Length(),
		"first":  unique.First(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: DistinctSimpleSlice — all same (empty)
// ==========================================

func Test_DistinctSimpleSlice_AllSame_Empty(t *testing.T) {
	tc := distinctSimpleSliceAllSameEmptyTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	items := input["items"].([]int)

	// Act
	ss := coregeneric.New.SimpleSlice.Int.Items(items...)
	unique := coregeneric.DistinctSimpleSlice(ss)
	actual := args.Map{
		"length":  unique.Length(),
		"isEmpty": unique.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
