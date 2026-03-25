package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{
			"length":   col.Length(),
			"first":    col.First(),
			"last":     col.Last(),
			"isSorted": coregeneric.IsSortedCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": col.Length(),
			"first":  col.First(),
			"last":   col.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"min": coregeneric.MinCollection(col),
			"max": coregeneric.MaxCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"min": coregeneric.MinCollectionOrDefault(col, -1),
			"max": coregeneric.MaxCollectionOrDefault(col, -1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinMaxOrDefault_Empty_Collection_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxCollectionOrDefaultEmptyTestCases {
		// Arrange — empty collection

		// Act
		col := coregeneric.New.Collection.Int.Empty()
		actual := args.Map{
			"min": coregeneric.MinCollectionOrDefault(col, -1),
			"max": coregeneric.MaxCollectionOrDefault(col, -1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"isSorted": coregeneric.IsSortedCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"sum": coregeneric.SumCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{}
		for i := 0; i < col.Length(); i++ {
			actual[fmt.Sprintf("val%d", i)] = col.SafeAt(i)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": len(sorted),
			"first":  sorted[0],
			"last":   sorted[len(sorted)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": len(sorted),
			"first":  sorted[0],
			"last":   sorted[len(sorted)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": col.Length(),
			"first":  col.First(),
			"last":   col.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"min": coregeneric.MinHashset(hs),
			"max": coregeneric.MaxHashset(hs),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinMaxOrDefault_Hashset_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxHashsetOrDefaultTestCases {
		// Arrange — empty hashset

		// Act
		hs := coregeneric.New.Hashset.Int.Empty()
		actual := args.Map{
			"min": coregeneric.MinHashsetOrDefault(hs, -1),
			"max": coregeneric.MaxHashsetOrDefault(hs, -1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinMaxOrDefault_Hashset_NonEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxHashsetOrDefaultNonEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		actual := args.Map{
			"min": coregeneric.MinHashsetOrDefault(hs, -1),
			"max": coregeneric.MaxHashsetOrDefault(hs, -1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": len(sortedKeys),
			"first":  sortedKeys[0],
			"last":   sortedKeys[len(sortedKeys)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": len(sortedKeys),
			"first":  sortedKeys[0],
			"last":   sortedKeys[len(sortedKeys)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"minKey": coregeneric.MinKeyHashmap(hm),
			"maxKey": coregeneric.MaxKeyHashmap(hm),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"minKey": coregeneric.MinKeyHashmapOrDefault(hm, "none"),
			"maxKey": coregeneric.MaxKeyHashmapOrDefault(hm, "none"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"minKey": coregeneric.MinKeyHashmapOrDefault(hm, "none"),
			"maxKey": coregeneric.MaxKeyHashmapOrDefault(hm, "none"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": len(sortedVals),
			"first":  sortedVals[0],
			"last":   sortedVals[len(sortedVals)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"minValue": coregeneric.MinValueHashmap(hm),
			"maxValue": coregeneric.MaxValueHashmap(hm),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"minValue": coregeneric.MinValueHashmapOrDefault(hm, -1),
			"maxValue": coregeneric.MaxValueHashmapOrDefault(hm, -1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"minValue": coregeneric.MinValueHashmapOrDefault(hm, -1),
			"maxValue": coregeneric.MaxValueHashmapOrDefault(hm, -1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": ss.Length(),
			"first":  ss.First(),
			"last":   ss.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"min": coregeneric.MinSimpleSlice(ss),
			"max": coregeneric.MaxSimpleSlice(ss),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// EDGE CASE TESTS
// ==========================================================================

// ==========================================
// Edge: SortCollection — empty
// ==========================================

func Test_SortCollection_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range sortCollectionEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		coregeneric.SortCollection(col)
		actual := args.Map{
			"length":   col.Length(),
			"isSorted": coregeneric.IsSortedCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: SortCollection — negative numbers
// ==========================================

func Test_SortCollection_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range sortCollectionNegativeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		coregeneric.SortCollection(col)
		actual := args.Map{
			"length":   col.Length(),
			"first":    col.First(),
			"last":     col.Last(),
			"isSorted": coregeneric.IsSortedCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: MinCollection / MaxCollection — negative
// ==========================================

func Test_MinMax_Collection_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxCollectionNegativeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actual := args.Map{
			"min": coregeneric.MinCollection(col),
			"max": coregeneric.MaxCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: SumCollection — negative
// ==========================================

func Test_SumCollection_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range sumCollectionNegativeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actual := args.Map{
			"sum": coregeneric.SumCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: ClampCollection — negative range
// ==========================================

func Test_ClampCollection_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range clampCollectionNegativeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		coregeneric.ClampCollection(col, -5, -1)
		actual := args.Map{}
		for i := 0; i < col.Length(); i++ {
			actual[fmt.Sprintf("val%d", i)] = col.SafeAt(i)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: IsSortedCollection — single and empty
// ==========================================

func Test_IsSortedCollection_Edge_Verification(t *testing.T) {
	for caseIndex, testCase := range isSortedCollectionEdgeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actual := args.Map{
			"isSorted": coregeneric.IsSortedCollection(col),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: SortedListHashset — single element
// ==========================================

func Test_SortedListHashset_Single_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedListHashsetSingleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		sorted := coregeneric.SortedListHashset(hs)
		actual := args.Map{
			"length": len(sorted),
			"first":  sorted[0],
			"last":   sorted[len(sorted)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: MinHashset / MaxHashset — single element
// ==========================================

func Test_MinMax_Hashset_Single_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxHashsetSingleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		actual := args.Map{
			"min": coregeneric.MinHashset(hs),
			"max": coregeneric.MaxHashset(hs),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: MinHashset / MaxHashset — negative numbers
// ==========================================

func Test_MinMax_Hashset_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxHashsetNegativeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		actual := args.Map{
			"min": coregeneric.MinHashset(hs),
			"max": coregeneric.MaxHashset(hs),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: SortedListHashset — negative numbers
// ==========================================

func Test_SortedListHashset_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedListHashsetNegativeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		hs := coregeneric.New.Hashset.Int.From(items)
		sorted := coregeneric.SortedListHashset(hs)
		actual := args.Map{
			"length": len(sorted),
			"first":  sorted[0],
			"last":   sorted[len(sorted)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: Hashmap — single entry
// ==========================================

func Test_SortedKeysHashmap_Single_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedKeysHashmapSingleTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(1)
		hm.Set("only", 99)

		// Act
		sortedKeys := coregeneric.SortedKeysHashmap(hm)
		actual := args.Map{
			"length": len(sortedKeys),
			"first":  sortedKeys[0],
			"last":   sortedKeys[len(sortedKeys)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinMaxKey_Hashmap_Single_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxKeyHashmapSingleTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(1)
		hm.Set("only", 99)

		// Act
		actual := args.Map{
			"minKey": coregeneric.MinKeyHashmap(hm),
			"maxKey": coregeneric.MaxKeyHashmap(hm),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Edge: Hashmap — negative values
// ==========================================

func Test_MinMaxValue_Hashmap_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range minMaxValueHashmapNegativeTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("alpha", -20)
		hm.Set("beta", 5)
		hm.Set("gamma", -3)

		// Act
		actual := args.Map{
			"minValue": coregeneric.MinValueHashmap(hm),
			"maxValue": coregeneric.MaxValueHashmap(hm),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_SortedValuesHashmap_Negative_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedValuesHashmapNegativeTestCases {
		// Arrange
		hm := coregeneric.New.Hashmap.StringInt.Cap(3)
		hm.Set("alpha", -20)
		hm.Set("beta", 5)
		hm.Set("gamma", -3)

		// Act
		sortedVals := coregeneric.SortedValuesHashmap(hm)
		actual := args.Map{
			"length": len(sortedVals),
			"first":  sortedVals[0],
			"last":   sortedVals[len(sortedVals)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
