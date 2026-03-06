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
		// Act
		col := coredynamic.New.Collection.Any.Empty()

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length":     col.Length(),
			"isEmpty":    col.IsEmpty(),
			"hasAnyItem": col.HasAnyItem(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
			"first":   fmt.Sprintf("%v", col.First()),
			"last":    fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length": col.Length(),
			"first":  fmt.Sprintf("%v", col.First()),
			"last":   fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"originalLength": len(original),
			"isIndependent":  len(original) != col.Length(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length": col.Length(),
			"first":  fmt.Sprintf("%v", col.First()),
			"last":   fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length": col.Length(),
			"first":  fmt.Sprintf("%v", col.First()),
			"last":   fmt.Sprintf("%v", col.Last()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: New.Collection.Any.From nil slice
// ==========================================

func Test_NewCreator_Generic_From_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericFromNilTestCases {
		// Act
		col := coredynamic.New.Collection.Any.From(nil)

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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

		actual := args.Map{
			"length":   col.Length(),
			"isEmpty":  col.IsEmpty(),
			"capacity": col.Capacity(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: New.Collection.Any.Items no args
// ==========================================

func Test_NewCreator_Generic_Items_NoArgs_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericItemsNoArgsTestCases {
		// Act
		col := coredynamic.New.Collection.Any.Items()

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: New.Collection.Any.Clone nil slice
// ==========================================

func Test_NewCreator_Generic_Clone_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range newCreatorGenericCloneNilTestCases {
		// Act
		col := coredynamic.New.Collection.Any.Clone(nil)

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
