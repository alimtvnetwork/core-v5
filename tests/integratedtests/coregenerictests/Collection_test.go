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
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
			col.First(),
			col.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			col.First(),
			col.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", original.Length()),
			fmt.Sprintf("%v", original.Length() != cloned.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has("a")),
			fmt.Sprintf("%v", hs.Has("c")),
			fmt.Sprintf("%v", hs.Has("z")),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", hs.Length()),
			fmt.Sprintf("%v", hs.Has(removeKey)),
			fmt.Sprintf("%v", hs.Has("a")),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", hm.Length()),
			val,
			fmt.Sprintf("%v", found),
			fmt.Sprintf("%v", notFound),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", ll.Length()),
			ll.First(),
			ll.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", ll.Length()),
			ll.First(),
			ll.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", mapped.Length()),
			mapped.First(),
			mapped.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
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
		actLines := []string{
			fmt.Sprintf("%d", unique.Length()),
			fmt.Sprintf("%d", unique.First()),
			fmt.Sprintf("%d", unique.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
