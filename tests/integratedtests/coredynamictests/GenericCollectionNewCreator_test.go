package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: New.Collection.Any.Empty
// ==========================================

func Test_NewCreator_Generic_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericEmptyTestCases {
		// Arrange — no input needed

		// Act
		col := coredynamic.New.Collection.Any.Empty()
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Cap
// ==========================================

func Test_NewCreator_Generic_Cap_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericCapTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)

		// Act
		col := coredynamic.New.Collection.Any.Cap(capacity)
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
// Test: New.Collection.Any.Cap zero
// ==========================================

func Test_NewCreator_Generic_Cap_Zero_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericCapZeroTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)

		// Act
		col := coredynamic.New.Collection.Any.Cap(capacity)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.From
// ==========================================

func Test_NewCreator_Generic_From_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericFromTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]any)

		// Act
		col := coredynamic.New.Collection.Any.From(items)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
			fmt.Sprintf("%v", col.First()),
			fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.From empty
// ==========================================

func Test_NewCreator_Generic_From_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericFromEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]any)

		// Act
		col := coredynamic.New.Collection.Any.From(items)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Clone
// ==========================================

func Test_NewCreator_Generic_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]any)

		// Act
		col := coredynamic.New.Collection.Any.Clone(items)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.First()),
			fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Clone mutation independence
// ==========================================

func Test_NewCreator_Generic_Clone_Mutation_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericCloneMutationTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]any)

		// Act
		original := items
		col := coredynamic.New.Collection.Any.Clone(items)
		col.Add("mutated")
		actLines := []string{
			fmt.Sprintf("%d", len(original)),
			fmt.Sprintf("%v", len(original) != col.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Items
// ==========================================

func Test_NewCreator_Generic_Items_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]any)

		// Act
		col := coredynamic.New.Collection.Any.Items(items...)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.First()),
			fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Items single
// ==========================================

func Test_NewCreator_Generic_Items_Single_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericItemsSingleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]any)

		// Act
		col := coredynamic.New.Collection.Any.Items(items...)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.First()),
			fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.From nil slice
// ==========================================

func Test_NewCreator_Generic_From_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericFromNilTestCases {
		// Arrange — nil slice

		// Act
		col := coredynamic.New.Collection.Any.From(nil)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Cap large capacity
// ==========================================

func Test_NewCreator_Generic_Cap_Large_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericCapLargeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		capacity := input.GetAsIntDefault("capacity", 0)

		// Act
		col := coredynamic.New.Collection.Any.Cap(capacity)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
			fmt.Sprintf("%d", col.Capacity()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Items no args
// ==========================================

func Test_NewCreator_Generic_Items_NoArgs_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericItemsNoArgsTestCases {
		// Arrange — no args

		// Act
		col := coredynamic.New.Collection.Any.Items()
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: New.Collection.Any.Clone nil slice
// ==========================================

func Test_NewCreator_Generic_Clone_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericCloneNilTestCases {
		// Arrange — nil slice

		// Act
		col := coredynamic.New.Collection.Any.Clone(nil)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
