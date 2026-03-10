package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Test: Collection RemoveAt
// ==========================================

func Test_Collection_RemoveAt_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionRemoveAtTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)
		removeIndex := input.GetAsIntDefault("removeIndex", 0)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		removed := col.RemoveAt(removeIndex)
		actual := args.Map{
			"removed": removed,
			"length":  col.Length(),
			"first":   col.FirstOrDefault(),
			"last":    col.LastOrDefault(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"skipLength": len(skipped),
			"skipFirst":  skipped[0],
			"takeLength": len(taken),
			"takeFirst":  taken[0],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": col.Length(),
			"first":  col.First(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"firstOrDefault": col.FirstOrDefault(),
			"lastOrDefault":  col.LastOrDefault(),
			"isEmpty":        col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"safeAt1":   col.SafeAt(1),
			"safeAt10":  col.SafeAt(10),
			"safeAtNeg": col.SafeAt(-1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"concatLength":   concatenated.Length(),
			"originalLength": original.Length(),
			"concatFirst":    concatenated.First(),
			"concatLast":     concatenated.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{"count": count}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": col1.Length(),
			"first":  col1.First(),
			"last":   col1.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"hasAllPresent":     hs.HasAll("a", "b", "c"),
			"hasAllWithMissing": hs.HasAll("a", "b", "d"),
			"hasAnyWithMatch":   hs.HasAny("x", "y", "a"),
			"hasAnyNoMatch":     hs.HasAny("x", "y", "z"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"equalsSame": hs1.IsEquals(hs2),
			"equalsDiff": hs1.IsEquals(hs3),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"firstAdd":  firstAdd,
			"secondAdd": secondAdd,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"existed": existed,
			"length":  hm.Length(),
			"hasA":    hm.Has("a"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"existing": hm.GetOrDefault("key1", -1),
			"missing":  hm.GetOrDefault("missing", -1),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"originalLength": hm.Length(),
			"clonedLength":   cloned.Length(),
			"isIndependent":  hm.Length() != cloned.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"keysCount":   len(hm.Keys()),
			"valuesCount": len(hm.Values()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"equalsSameLength": hm1.IsEquals(hm2),
			"equalsDiffLength": hm1.IsEquals(hm3),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": len(allItems),
			"first":  allItems[0],
			"last":   allItems[len(allItems)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"elementAt1": node.Element,
			"nilAt10":    nilNode == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length":         ll.Length(),
			"isEmpty":        ll.IsEmpty(),
			"hasItems":       ll.HasItems(),
			"firstOrDefault": ll.FirstOrDefault(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": filtered.Length(),
			"first":  filtered.First(),
			"last":   filtered.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"originalLength": ss.Length(),
			"isIndependent":  ss.Length() != cloned.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"skipLength": len(skipped),
			"takeLength": len(taken),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"length": flatMapped.Length(),
			"first":  flatMapped.First(),
			"last":   flatMapped.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{"sum": sum}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"groupCount": len(groups),
			"evenCount":  groups["even"].Length(),
			"oddCount":   groups["odd"].Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
