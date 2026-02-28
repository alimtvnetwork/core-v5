package coredynamictests

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: AddLock — concurrent safety
// ==========================================

func Test_Collection_AddLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 100)
		col := coredynamic.New.Collection.String.Empty()

		// Act
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func(idx int) {
				col.AddLock(fmt.Sprintf("item-%d", idx))
				wg.Done()
			}(i)
		}
		wg.Wait()

		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AddsLock — concurrent safety
// ==========================================

func Test_Collection_AddsLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddsLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 50)
		batch := input.GetAsIntDefault("batch", 2)
		col := coredynamic.New.Collection.String.Empty()

		// Act
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func(idx int) {
				items := make([]string, batch)
				for b := 0; b < batch; b++ {
					items[b] = fmt.Sprintf("item-%d-%d", idx, b)
				}
				col.AddsLock(items...)
				wg.Done()
			}(i)
		}
		wg.Wait()

		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: LengthLock
// ==========================================

func Test_Collection_LengthLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionLengthLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", col.LengthLock()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsEmptyLock — empty
// ==========================================

func Test_Collection_IsEmptyLock_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsEmptyLockTestCases {
		// Arrange — empty collection

		// Act
		col := coredynamic.New.Collection.String.Empty()
		actLines := []string{
			fmt.Sprintf("%v", col.IsEmptyLock()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsEmptyLock — non-empty
// ==========================================

func Test_Collection_IsEmptyLock_NonEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsEmptyLockNonEmptyTestCases {
		// Arrange

		// Act
		col := coredynamic.New.Collection.String.Items("x")
		actLines := []string{
			fmt.Sprintf("%v", col.IsEmptyLock()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ItemsLock — returns independent copy
// ==========================================

func Test_Collection_ItemsLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionItemsLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coredynamic.New.Collection.String.From(items)
		copied := col.ItemsLock()
		copied = append(copied, "mutated")
		actLines := []string{
			fmt.Sprintf("%d", len(items)),
			items[0],
			items[len(items)-1],
			fmt.Sprintf("%v", col.Length() != len(copied)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: ClearLock
// ==========================================

func Test_Collection_ClearLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionClearLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")

		// Act
		col := coredynamic.New.Collection.String.From(items)
		col.ClearLock()
		actLines := []string{
			fmt.Sprintf("%d", col.Length()),
			fmt.Sprintf("%v", col.IsEmpty()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: AddCollectionLock
// ==========================================

func Test_Collection_AddCollectionLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddCollectionLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first, _ := input.GetAsStrings("first")
		second, _ := input.GetAsStrings("second")

		// Act
		col1 := coredynamic.New.Collection.String.From(first)
		col2 := coredynamic.New.Collection.String.From(second)
		col1.AddCollectionLock(col2)
		actLines := []string{
			fmt.Sprintf("%d", col1.Length()),
			col1.First(),
			col1.Last(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
