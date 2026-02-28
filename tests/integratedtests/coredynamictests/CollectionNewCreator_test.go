package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: New.Collection.String.Cap
// ==========================================

func Test_NewCreator_String_Cap_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorStringCapTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)

		// Act
		col := coredynamic.New.Collection.String.Cap(capacity)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
			fmt.Sprintf("%v", col.HasAnyItem()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.String.Empty
// ==========================================

func Test_NewCreator_String_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorStringEmptyTestCases {
		// Arrange — no input needed

		// Act
		col := coredynamic.New.Collection.String.Empty()
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

func Test_NewCreator_String_From_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorStringFromTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
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
// Test: New.Collection.String.Clone
// ==========================================

func Test_NewCreator_String_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorStringCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.Clone(items)
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
// Test: New.Collection.String.Items
// ==========================================

func Test_NewCreator_String_Items_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorStringItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.Items(items...)
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
// Test: New.Collection.Int.Cap
// ==========================================

func Test_NewCreator_Int_Cap_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorIntCapTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)

		// Act
		col := coredynamic.New.Collection.Int.Cap(capacity)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Int.Items
// ==========================================

func Test_NewCreator_Int_Items_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorIntItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.Items(items...)
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
// Test: AddIf true
// ==========================================

func Test_Collection_AddIf_True_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddIfTrueTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		item, isValid := input.GetAsString("item")
		if !isValid {
			errcore.HandleErrMessage("GetAsString 'item' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.Empty()
		col.AddIf(true, item)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			col.First(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AddIf false
// ==========================================

func Test_Collection_AddIf_False_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddIfFalseTestCases {
		// Arrange — no special input

		// Act
		col := coredynamic.New.Collection.String.Empty()
		col.AddIf(false, "skipped")
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AddCollection
// ==========================================

func Test_Collection_AddCollection_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddCollectionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first, isValid := input.GetAsStrings("first")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'first' failed")
		}
		second, isValid := input.GetAsStrings("second")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'second' failed")
		}

		// Act
		col1 := coredynamic.New.Collection.String.From(first)
		col2 := coredynamic.New.Collection.String.From(second)
		col1.AddCollection(col2)
		actLines := []string{
			fmt.Sprintf("%d", col1.Length()),
			col1.First(),
			col1.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AddCollection nil
// ==========================================

func Test_Collection_AddCollection_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddCollectionNilTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first, isValid := input.GetAsStrings("first")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'first' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(first)
		col.AddCollection(nil)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			col.First(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Clone
// ==========================================

func Test_Collection_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		original := coredynamic.New.Collection.String.From(items)
		cloned := original.Clone()
		cloned.Add("mutated")
		actLines := []string{
			fmt.Sprintf("%d", original.Length()),
			original.First(),
			original.Last(),
			fmt.Sprintf("%v", original.Length() != cloned.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Reverse
// ==========================================

func Test_Collection_Reverse_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionReverseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.Clone(items)
		col.Reverse()
		actLines := col.Strings()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Reverse empty
// ==========================================

func Test_Collection_Reverse_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionReverseEmptyTestCases {
		// Arrange — empty collection

		// Act
		col := coredynamic.New.Collection.String.Empty()
		col.Reverse()
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ConcatNew
// ==========================================

func Test_Collection_ConcatNew_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionConcatNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		original, isValid := input.GetAsStrings("original")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'original' failed")
		}
		adding, isValid := input.GetAsStrings("adding")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'adding' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(original)
		result := col.ConcatNew(adding...)
		actLines := []string{
			fmt.Sprintf("%d", result.Length()),
			result.First(),
			result.Last(),
			fmt.Sprintf("%d", col.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Capacity
// ==========================================

func Test_Collection_Capacity_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCapacityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)

		// Act
		col := coredynamic.New.Collection.String.Cap(capacity)
		actLines := []string{
			fmt.Sprintf("%d", col.Capacity()),
			fmt.Sprintf("%d", col.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AddCapacity / Resize
// ==========================================

func Test_Collection_AddCapacity_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionResizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)
		additional := input.GetAsIntDefault("additional", 0)

		// Act
		col := coredynamic.New.Collection.String.Cap(capacity)
		col.AddCapacity(additional)
		actLines := []string{
			fmt.Sprintf("%v", col.Capacity() >= capacity+additional),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
