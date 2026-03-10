package coredynamictests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
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

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", col.Length()))
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

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", col.Length()))
	}
}

// ==========================================
// Test: LengthLock
// ==========================================

func Test_Collection_LengthLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionLengthLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", col.LengthLock()))
	}
}

// ==========================================
// Test: IsEmptyLock — empty
// ==========================================

func Test_Collection_IsEmptyLock_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsEmptyLockTestCases {
		// Act
		col := coredynamic.New.Collection.String.Empty()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", col.IsEmptyLock()))
	}
}

// ==========================================
// Test: IsEmptyLock — non-empty
// ==========================================

func Test_Collection_IsEmptyLock_NonEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsEmptyLockNonEmptyTestCases {
		// Act
		col := coredynamic.New.Collection.String.Items("x")

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", col.IsEmptyLock()))
	}
}

// ==========================================
// Test: ItemsLock — returns independent copy
// ==========================================

func Test_Collection_ItemsLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionItemsLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		copied := col.ItemsLock()
		copied = append(copied, "mutated")

		actual := args.Map{
			"length":        len(items),
			"first":         items[0],
			"last":          items[len(items)-1],
			"isIndependent": col.Length() != len(copied),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: ClearLock
// ==========================================

func Test_Collection_ClearLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionClearLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		col.ClearLock()

		actual := args.Map{
			"length":  col.Length(),
			"isEmpty": col.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: AddCollectionLock
// ==========================================

func Test_Collection_AddCollectionLock_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionAddCollectionLockTestCases {
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
		col1.AddCollectionLock(col2)

		actual := args.Map{
			"length": col1.Length(),
			"first":  col1.First(),
			"last":   col1.Last(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
