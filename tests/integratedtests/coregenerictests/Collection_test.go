package coregenerictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: New.Collection.String.Cap
// ==========================================

func Test_Collection_String_Cap_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionStringCapTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)

		// Act
		col := coregeneric.New.Collection.String.Cap(capacity)
		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: New.Collection.String.Empty
// ==========================================

func Test_Collection_String_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionStringEmptyTestCases {
		// Arrange — no input needed

		// Act
		col := coregeneric.New.Collection.String.Empty()
		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: New.Collection.String.From
// ==========================================

func Test_Collection_String_From_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionStringFromTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coregeneric.New.Collection.String.From(items)
		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
			"first":   col.First(),
			"last":    col.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: New.Collection.String.Items
// ==========================================

func Test_Collection_String_Items_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionStringItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coregeneric.New.Collection.String.Items(items...)
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
// Test: New.Collection.Int.Items
// ==========================================

func Test_Collection_Int_Items_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIntItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
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
// Test: Collection.Filter
// ==========================================

func Test_Collection_Filter_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionFilterTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		filtered := col.Filter(func(item int) bool { return item > 2 })
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
// Test: Collection.Clone independence
// ==========================================

func Test_Collection_Clone_Independence_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCloneIndependenceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		original := coregeneric.New.Collection.String.From(items)
		cloned := original.Clone()
		cloned.Add("mutated")
		actual := args.Map{
			"originalLength": original.Length(),
			"isIndependent":  original.Length() != cloned.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Hashset Add/Has
// ==========================================

func Test_Hashset_AddHas_Verification(t *testing.T) {
	for caseIndex, testCase := range hashsetAddHasTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		hs := coregeneric.New.Hashset.String.Empty()
		hs.AddSlice(items)
		actual := args.Map{
			"length": hs.Length(),
			"hasA":   hs.Has("a"),
			"hasC":   hs.Has("c"),
			"hasZ":   hs.Has("z"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Hashset Remove
// ==========================================

func Test_Hashset_Remove_Verification(t *testing.T) {
	for caseIndex, testCase := range hashsetRemoveTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")
		removeKey, _ := input.GetAsString("remove")

		// Act
		hs := coregeneric.New.Hashset.String.From(items)
		hs.Remove(removeKey)
		actual := args.Map{
			"length":     hs.Length(),
			"hasRemoved": hs.Has(removeKey),
			"hasA":       hs.Has("a"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Hashmap Set/Get
// ==========================================

func Test_Hashmap_SetGet_Verification(t *testing.T) {
	for caseIndex, testCase := range hashmapSetGetTestCases {
		// Arrange — no special input

		// Act
		hm := coregeneric.New.Hashmap.StringString.Cap(5)
		hm.Set("key1", "value1")
		hm.Set("key2", "value2")
		val, found := hm.Get("key1")
		_, notFound := hm.Get("missing")
		actual := args.Map{
			"length":   hm.Length(),
			"value":    val,
			"found":    found,
			"notFound": notFound,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: SimpleSlice Add
// ==========================================

func Test_SimpleSlice_Add_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleSliceAddTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		ss := coregeneric.New.SimpleSlice.Int.Empty()
		for _, item := range items {
			ss.Add(item)
		}
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
// Test: LinkedList Add
// ==========================================

func Test_LinkedList_Add_Verification(t *testing.T) {
	for caseIndex, testCase := range linkedListAddTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		ll := coregeneric.New.LinkedList.String.Empty()
		for _, item := range items {
			ll.Add(item)
		}
		actual := args.Map{
			"length": ll.Length(),
			"first":  ll.First(),
			"last":   ll.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: LinkedList AddFront
// ==========================================

func Test_LinkedList_AddFront_Verification(t *testing.T) {
	for caseIndex, testCase := range linkedListAddFrontTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")
		prepend, _ := input.GetAsString("prepend")

		// Act
		ll := coregeneric.New.LinkedList.String.From(items)
		ll.AddFront(prepend)
		actual := args.Map{
			"length": ll.Length(),
			"first":  ll.First(),
			"last":   ll.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: MapCollection
// ==========================================

func Test_MapCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range mapCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		mapped := coregeneric.MapCollection(col, func(item int) string {
			return fmt.Sprintf("%d", item)
		})
		actual := args.Map{
			"length": mapped.Length(),
			"first":  mapped.First(),
			"last":   mapped.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Distinct
// ==========================================

func Test_Distinct_Verification(t *testing.T) {
	for caseIndex, testCase := range distinctTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coregeneric.New.Collection.Int.Items(items...)
		unique := coregeneric.Distinct(col)
		actual := args.Map{
			"length": unique.Length(),
			"first":  unique.First(),
			"last":   unique.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
