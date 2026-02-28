package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: SortCollection ascending
// ==========================================

func Test_SortCollection_Asc_Verification(t *testing.T) {
	for caseIndex, testCase := range sortCollectionAscTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		coregeneric.SortCollection(col)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
			fmt.Sprintf("%v", coregeneric.IsSortedCollection(col)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortCollectionDesc
// ==========================================

func Test_SortCollection_Desc_Verification(t *testing.T) {
	for caseIndex, testCase := range sortCollectionDescTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		coregeneric.SortCollectionDesc(col)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinCollection / MaxCollection
// ==========================================

func Test_MinMax_Collection_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinCollection(col)),
			fmt.Sprintf("%d", coregeneric.MaxCollection(col)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinCollectionOrDefault / MaxCollectionOrDefault
// ==========================================

func Test_MinMaxOrDefault_Collection_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxCollectionOrDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinCollectionOrDefault(col, -1)),
			fmt.Sprintf("%d", coregeneric.MaxCollectionOrDefault(col, -1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_MinMaxOrDefault_Empty_Collection_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxCollectionOrDefaultEmptyTestCases {
		// Arrange — empty collection

		// Act
		col := coregeneric.New.Collection.Int.Empty()
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinCollectionOrDefault(col, -1)),
			fmt.Sprintf("%d", coregeneric.MaxCollectionOrDefault(col, -1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsSortedCollection
// ==========================================

func Test_IsSortedCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range isSortedCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%v", coregeneric.IsSortedCollection(col)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SumCollection
// ==========================================

func Test_SumCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range sumCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.SumCollection(col)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ClampCollection
// ==========================================

func Test_ClampCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range clampCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		coregeneric.ClampCollection(col, 2, 4)
		actLines := make([]string, col.Length())
		for i := 0; i < col.Length(); i++ {
			actLines[i] = fmt.Sprintf("%d", col.GetAt(i))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortedListHashset
// ==========================================

func Test_SortedListHashset_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedListHashsetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		sorted := coregeneric.SortedListHashset(hs)
		actLines := []string{
			fmt.Sprintf("%d", len(sorted)),
			fmt.Sprintf("%d", sorted[0]),
			fmt.Sprintf("%d", sorted[len(sorted)-1]),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortedListDescHashset
// ==========================================

func Test_SortedListDescHashset_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedListDescHashsetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		sorted := coregeneric.SortedListDescHashset(hs)
		actLines := []string{
			fmt.Sprintf("%d", len(sorted)),
			fmt.Sprintf("%d", sorted[0]),
			fmt.Sprintf("%d", sorted[len(sorted)-1]),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortedCollectionHashset
// ==========================================

func Test_SortedCollectionHashset_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedCollectionHashsetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		col := coregeneric.SortedCollectionHashset(hs)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
			fmt.Sprintf("%d", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinHashset / MaxHashset
// ==========================================

func Test_MinMax_Hashset_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxHashsetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinHashset(hs)),
			fmt.Sprintf("%d", coregeneric.MaxHashset(hs)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_MinMaxOrDefault_Hashset_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxHashsetOrDefaultTestCases {
		// Arrange — empty hashset

		// Act
		hs := coregeneric.New.Hashset.Int.Empty()
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinHashsetOrDefault(hs, -1)),
			fmt.Sprintf("%d", coregeneric.MaxHashsetOrDefault(hs, -1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_MinMaxOrDefault_Hashset_NonEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxHashsetOrDefaultNonEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinHashsetOrDefault(hs, -1)),
			fmt.Sprintf("%d", coregeneric.MaxHashsetOrDefault(hs, -1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortedKeysHashmap
// ==========================================

func Test_SortedKeysHashmap_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedKeysHashmapTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 10)
		hm.Set("beta", 20)

		// Act
		sortedKeys := coregeneric.SortedKeysHashmap(hm)
		actLines := []string{
			fmt.Sprintf("%d", len(sortedKeys)),
			sortedKeys[0],
			sortedKeys[len(sortedKeys)-1],
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortedKeysDescHashmap
// ==========================================

func Test_SortedKeysDescHashmap_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedKeysDescHashmapTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 10)
		hm.Set("beta", 20)

		// Act
		sortedKeys := coregeneric.SortedKeysDescHashmap(hm)
		actLines := []string{
			fmt.Sprintf("%d", len(sortedKeys)),
			sortedKeys[0],
			sortedKeys[len(sortedKeys)-1],
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinKeyHashmap / MaxKeyHashmap
// ==========================================

func Test_MinMaxKey_Hashmap_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxKeyHashmapTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 10)
		hm.Set("beta", 20)

		// Act
		actLines := []string{
			coregeneric.MinKeyHashmap(hm),
			coregeneric.MaxKeyHashmap(hm),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinKeyHashmapOrDefault / MaxKeyHashmapOrDefault
// ==========================================

func Test_MinMaxKeyOrDefault_Hashmap_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxKeyHashmapOrDefaultEmptyTestCases {
		// Arrange — empty hashmap

		// Act
		hm := coregeneric.New.Hashmap.StringInt.Cap(0)
		actLines := []string{
			coregeneric.MinKeyHashmapOrDefault(hm, "none"),
			coregeneric.MaxKeyHashmapOrDefault(hm, "none"),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_MinMaxKeyOrDefault_Hashmap_NonEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxKeyHashmapOrDefaultNonEmptyTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 10)
		hm.Set("beta", 20)

		// Act
		actLines := []string{
			coregeneric.MinKeyHashmapOrDefault(hm, "none"),
			coregeneric.MaxKeyHashmapOrDefault(hm, "none"),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortedValuesHashmap
// ==========================================

func Test_SortedValuesHashmap_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedValuesHashmapTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 1)
		hm.Set("beta", 20)

		// Act
		sortedVals := coregeneric.SortedValuesHashmap(hm)
		actLines := []string{
			fmt.Sprintf("%d", len(sortedVals)),
			fmt.Sprintf("%d", sortedVals[0]),
			fmt.Sprintf("%d", sortedVals[len(sortedVals)-1]),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinValueHashmap / MaxValueHashmap
// ==========================================

func Test_MinMaxValue_Hashmap_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxValueHashmapTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 1)
		hm.Set("beta", 20)

		// Act
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinValueHashmap(hm)),
			fmt.Sprintf("%d", coregeneric.MaxValueHashmap(hm)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinValueHashmapOrDefault / MaxValueHashmapOrDefault
// ==========================================

func Test_MinMaxValueOrDefault_Hashmap_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxValueHashmapOrDefaultEmptyTestCases {
		// Arrange — empty hashmap

		// Act
		hm := coregeneric.New.Hashmap.StringInt.Cap(0)
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinValueHashmapOrDefault(hm, -1)),
			fmt.Sprintf("%d", coregeneric.MaxValueHashmapOrDefault(hm, -1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_MinMaxValueOrDefault_Hashmap_NonEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxValueHashmapOrDefaultNonEmptyTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 1)
		hm.Set("beta", 20)

		// Act
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinValueHashmapOrDefault(hm, -1)),
			fmt.Sprintf("%d", coregeneric.MaxValueHashmapOrDefault(hm, -1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinValueHashmap / MaxValueHashmap
// ==========================================

func Test_MinMaxValue_Hashmap_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxValueHashmapTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("gamma", 30)
		hm.Set("alpha", 1)
		hm.Set("beta", 20)

		// Act
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinValueHashmap(hm)),
			fmt.Sprintf("%d", coregeneric.MaxValueHashmap(hm)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortSimpleSlice
// ==========================================

func Test_SortSimpleSlice_Verification(t *testing.T) {
	for caseIndex, testCase := range sortSimpleSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		coregeneric.SortSimpleSlice(ss)
		actLines := []string{
			fmt.Sprintf("%d", ss.Length()),
			fmt.Sprintf("%d", ss.First()),
			fmt.Sprintf("%d", ss.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: MinSimpleSlice / MaxSimpleSlice
// ==========================================

func Test_MinMax_SimpleSlice_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxSimpleSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%d", coregeneric.MinSimpleSlice(ss)),
			fmt.Sprintf("%d", coregeneric.MaxSimpleSlice(ss)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
