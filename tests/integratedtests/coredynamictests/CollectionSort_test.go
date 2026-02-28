package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: SortAsc — strings
// ==========================================

func Test_SortAsc_String_Verification(t *testing.T) {
	for caseIndex, testCase := range sortAscStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coredynamic.New.Collection.String.Clone(items)
		coredynamic.SortAsc(col)
		actLines := col.Items()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortDesc — strings
// ==========================================

func Test_SortDesc_String_Verification(t *testing.T) {
	for caseIndex, testCase := range sortDescStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coredynamic.New.Collection.String.Clone(items)
		coredynamic.SortDesc(col)
		actLines := col.Items()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortAsc — ints
// ==========================================

func Test_SortAsc_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range sortAscIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.Clone(items)
		coredynamic.SortAsc(col)
		actLines := make([]string, col.Length())
		for i, v := range col.Items() {
			actLines[i] = fmt.Sprintf("%d", v)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortDesc — ints
// ==========================================

func Test_SortDesc_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range sortDescIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.Clone(items)
		coredynamic.SortDesc(col)
		actLines := make([]string, col.Length())
		for i, v := range col.Items() {
			actLines[i] = fmt.Sprintf("%d", v)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortedAsc — non-mutating
// ==========================================

func Test_SortedAsc_NonMutating_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedAscNonMutatingTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		original := coredynamic.New.Collection.String.Clone(items)
		sorted := coredynamic.SortedAsc(original)
		actLines := append(sorted.Items(), original.Items()...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortAsc — empty
// ==========================================

func Test_SortAsc_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range sortEmptyTestCases {
		// Arrange

		// Act
		col := coredynamic.New.Collection.String.Empty()
		coredynamic.SortAsc(col)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortAsc — single element
// ==========================================

func Test_SortAsc_Single_Verification(t *testing.T) {
	for caseIndex, testCase := range sortSingleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coredynamic.New.Collection.String.From(items)
		coredynamic.SortAsc(col)
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			col.First(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsSortedAsc — true
// ==========================================

func Test_IsSortedAsc_True_Verification(t *testing.T) {
	for caseIndex, testCase := range isSortedAscTrueTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.IsSortedAsc(col)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsSortedAsc — false
// ==========================================

func Test_IsSortedAsc_False_Verification(t *testing.T) {
	for caseIndex, testCase := range isSortedAscFalseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.IsSortedAsc(col)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortFunc — custom comparator (by string length)
// ==========================================

func Test_SortFunc_Custom_Verification(t *testing.T) {
	for caseIndex, testCase := range sortFuncCustomTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coredynamic.New.Collection.String.Clone(items)
		col.SortFunc(func(a, b string) bool {
			return len(a) < len(b)
		})
		actLines := col.Items()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SortAsc — float64
// ==========================================

func Test_SortAsc_Float64_Verification(t *testing.T) {
	for caseIndex, testCase := range sortAscFloat64TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]float64)

		// Act
		col := coredynamic.New.Collection.Float64.Clone(items)
		coredynamic.SortAsc(col)
		actLines := make([]string, col.Length())
		for i, v := range col.Items() {
			actLines[i] = fmt.Sprintf("%g", v)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
