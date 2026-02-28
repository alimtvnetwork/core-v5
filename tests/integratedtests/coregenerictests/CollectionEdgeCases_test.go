package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: Collection RemoveAt
// ==========================================

func Test_Collection_RemoveAt_Verification(t *testing.T) {
	removeIndices := []int{1, 10, -1}

	for caseIndex, testCase := range collectionRemoveAtTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		removed := col.RemoveAt(removeIndices[caseIndex])
		actLines := []string{
			fmt.Sprintf("%v", removed),
			fmt.Sprintf("%d", col.Length()),
		}

		if removed {
			actLines = append(actLines,
				fmt.Sprintf("%d", col.First()),
				fmt.Sprintf("%d", col.Last()),
			)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Collection Reverse
// ==========================================

func Test_Collection_Reverse_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionReverseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		col.Reverse()
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
// Test: Collection Skip / Take
// ==========================================

func Test_Collection_SkipTake_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionSkipTakeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		skipped := col.Skip(2)
		taken := col.Take(2)
		actLines := []string{
			fmt.Sprintf("%d", len(skipped)),
			fmt.Sprintf("%d", skipped[0]),
			fmt.Sprintf("%d", len(taken)),
			fmt.Sprintf("%d", taken[0]),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Collection AddIf
// ==========================================

func Test_Collection_AddIf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddIfTestCases {
		// Arrange — no special input

		// Act
		col := coregeneric.EmptyCollection[int]()
		col.AddIf(true, 100)
		col.AddIf(false, 200)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%d", col.First()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Collection FirstOrDefault / LastOrDefault on empty
// ==========================================

func Test_Collection_DefaultsEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDefaultsEmptyTestCases {
		// Arrange — empty collection

		// Act
		col := coregeneric.EmptyCollection[int]()
		actLines := []string{
			fmt.Sprintf("%d", col.FirstOrDefault()),
			fmt.Sprintf("%d", col.LastOrDefault()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Collection SafeAt
// ==========================================

func Test_Collection_SafeAt_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionSafeAtTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		actLines := []string{
			fmt.Sprintf("%d", col.SafeAt(1)),
			fmt.Sprintf("%d", col.SafeAt(10)),
			fmt.Sprintf("%d", col.SafeAt(-1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Collection ConcatNew
// ==========================================

func Test_Collection_ConcatNew_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionConcatNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		original := coregeneric.New.Collection.Int.Items(items...)
		concatenated := original.ConcatNew(4, 5)
		actLines := []string{
			fmt.Sprintf("%d", concatenated.Length()),
			fmt.Sprintf("%d", original.Length()),
			fmt.Sprintf("%d", concatenated.First()),
			fmt.Sprintf("%d", concatenated.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Collection CountFunc
// ==========================================

func Test_Collection_CountFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCountFuncTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		count := col.CountFunc(func(item int) bool { return item%2 == 0 })
		actLines := []string{
			fmt.Sprintf("%d", count),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Collection AddCollection
// ==========================================

func Test_Collection_AddCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddCollectionTestCases {
		// Arrange — create two collections

		// Act
		col1 := coregeneric.New.Collection.Int.Items(1, 2, 3)
		col2 := coregeneric.New.Collection.Int.Items(4, 5)
		col1.AddCollection(col2)
		actLines := []string{
			fmt.Sprintf("%d", col1.Length()),
			fmt.Sprintf("%d", col1.First()),
			fmt.Sprintf("%d", col1.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashset HasAll / HasAny
// ==========================================

func Test_Hashset_HasAll_HasAny_Verification(t *testing.T) {
	for caseIndex, testCase := range hashsetHasAllHasAnyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		hs := coregeneric.New.Hashset.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", hs.HasAll("a", "b", "c")),
			fmt.Sprintf("%v", hs.HasAll("a", "b", "d")),
			fmt.Sprintf("%v", hs.HasAny("x", "y", "a")),
			fmt.Sprintf("%v", hs.HasAny("x", "y", "z")),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashset IsEquals
// ==========================================

func Test_Hashset_IsEquals_Verification(t *testing.T) {
	for caseIndex, testCase := range hashsetIsEqualsTestCases {
		// Arrange — create two equal and one different hashsets

		// Act
		hs1 := coregeneric.New.Hashset.Int.From([]int{1, 2, 3})
		hs2 := coregeneric.New.Hashset.Int.From([]int{3, 2, 1})
		hs3 := coregeneric.New.Hashset.Int.From([]int{1, 2, 4})
		actLines := []string{
			fmt.Sprintf("%v", hs1.IsEquals(hs2)),
			fmt.Sprintf("%v", hs1.IsEquals(hs3)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashset AddBool
// ==========================================

func Test_Hashset_AddBool_Verification(t *testing.T) {
	for caseIndex, testCase := range hashsetAddBoolTestCases {
		// Arrange

		// Act
		hs := coregeneric.New.Hashset.String.Empty()
		firstAdd := hs.AddBool("key1")
		secondAdd := hs.AddBool("key1")
		actLines := []string{
			fmt.Sprintf("%v", firstAdd),
			fmt.Sprintf("%v", secondAdd),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashmap Remove
// ==========================================

func Test_Hashmap_Remove_Verification(t *testing.T) {
	for caseIndex, testCase := range hashmapRemoveTestCases {
		// Arrange

		// Act
		hm := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm.Set("a", 1)
		hm.Set("b", 2)
		existed := hm.Remove("a")
		actLines := []string{
			fmt.Sprintf("%v", existed),
			fmt.Sprintf("%d", hm.Length()),
			fmt.Sprintf("%v", hm.Has("a")),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashmap GetOrDefault
// ==========================================

func Test_Hashmap_GetOrDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range hashmapGetOrDefaultTestCases {
		// Arrange

		// Act
		hm := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm.Set("key1", 100)
		actLines := []string{
			fmt.Sprintf("%d", hm.GetOrDefault("key1", -1)),
			fmt.Sprintf("%d", hm.GetOrDefault("missing", -1)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashmap Clone
// ==========================================

func Test_Hashmap_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range hashmapCloneTestCases {
		// Arrange

		// Act
		hm := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm.Set("a", 1)
		hm.Set("b", 2)
		cloned := hm.Clone()
		cloned.Set("c", 3)
		actLines := []string{
			fmt.Sprintf("%d", hm.Length()),
			fmt.Sprintf("%d", cloned.Length()-1),
			fmt.Sprintf("%v", hm.Length() != cloned.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashmap Keys / Values
// ==========================================

func Test_Hashmap_KeysValues_Verification(t *testing.T) {
	for caseIndex, testCase := range hashmapKeysValuesTestCases {
		// Arrange

		// Act
		hm := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm.Set("a", 1)
		hm.Set("b", 2)
		hm.Set("c", 3)
		actLines := []string{
			fmt.Sprintf("%d", len(hm.Keys())),
			fmt.Sprintf("%d", len(hm.Values())),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Hashmap IsEquals
// ==========================================

func Test_Hashmap_IsEquals_Verification(t *testing.T) {
	for caseIndex, testCase := range hashmapIsEqualsTestCases {
		// Arrange

		// Act
		hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm1.Set("a", 1)
		hm1.Set("b", 2)
		hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm2.Set("x", 10)
		hm2.Set("y", 20)
		hm3 := coregeneric.New.Hashmap.StringInt.Cap(5)
		hm3.Set("a", 1)
		actLines := []string{
			fmt.Sprintf("%v", hm1.IsEquals(hm2)),
			fmt.Sprintf("%v", hm1.IsEquals(hm3)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: LinkedList Items
// ==========================================

func Test_LinkedList_Items_Verification(t *testing.T) {
	for caseIndex, testCase := range linkedListItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		ll := coregeneric.New.LinkedList.String.From(items)
		allItems := ll.Items()
		actLines := []string{
			fmt.Sprintf("%d", len(allItems)),
			allItems[0],
			allItems[len(allItems)-1],
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: LinkedList IndexAt
// ==========================================

func Test_LinkedList_IndexAt_Verification(t *testing.T) {
	for caseIndex, testCase := range linkedListIndexAtTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		ll := coregeneric.New.LinkedList.String.From(items)
		node := ll.IndexAt(1)
		nilNode := ll.IndexAt(10)
		actLines := []string{
			node.Element,
			fmt.Sprintf("%v", nilNode == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: LinkedList empty edge cases
// ==========================================

func Test_LinkedList_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range linkedListEmptyTestCases {
		// Arrange — empty linked list

		// Act
		ll := coregeneric.New.LinkedList.String.Empty()
		actLines := []string{
			fmt.Sprintf("%d", ll.Length()),
			fmt.Sprintf("%v", ll.IsEmpty()),
			fmt.Sprintf("%v", ll.HasItems()),
			ll.FirstOrDefault(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SimpleSlice Filter
// ==========================================

func Test_SimpleSlice_Filter_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceFilterTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		filtered := ss.Filter(func(item int) bool { return item > 2 })
		actLines := []string{
			fmt.Sprintf("%d", filtered.Length()),
			fmt.Sprintf("%d", filtered.First()),
			fmt.Sprintf("%d", filtered.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SimpleSlice Clone
// ==========================================

func Test_SimpleSlice_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		cloned := ss.Clone()
		cloned.Add(40)
		actLines := []string{
			fmt.Sprintf("%d", ss.Length()),
			fmt.Sprintf("%v", ss.Length() != cloned.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SimpleSlice Skip / Take
// ==========================================

func Test_SimpleSlice_SkipTake_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceSkipTakeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Items(items...)
		skipped := ss.Skip(2)
		taken := ss.Take(2)
		actLines := []string{
			fmt.Sprintf("%d", len(skipped)),
			fmt.Sprintf("%d", len(taken)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: FlatMapCollection
// ==========================================

func Test_FlatMapCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range flatMapCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		flatMapped := coregeneric.FlatMapCollection(col, func(item int) []string {
			return []string{
				fmt.Sprintf("%d", item),
				fmt.Sprintf("%d", item),
			}
		})
		actLines := []string{
			fmt.Sprintf("%d", flatMapped.Length()),
			flatMapped.First(),
			flatMapped.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ReduceCollection
// ==========================================

func Test_ReduceCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range reduceCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		sum := coregeneric.ReduceCollection(col, 0, func(acc int, item int) int {
			return acc + item
		})
		actLines := []string{
			fmt.Sprintf("%d", sum),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: GroupByCollection
// ==========================================

func Test_GroupByCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range groupByCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		groups := coregeneric.GroupByCollection(col, func(item int) string {
			if item%2 == 0 {
				return "even"
			}
			return "odd"
		})
		actLines := []string{
			fmt.Sprintf("%d", len(groups)),
			fmt.Sprintf("%d", groups["even"].Length()),
			fmt.Sprintf("%d", groups["odd"].Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
