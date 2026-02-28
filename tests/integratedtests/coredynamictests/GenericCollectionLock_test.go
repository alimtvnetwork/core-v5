package coredynamictests

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: Generic AddLock — concurrent safety
// ==========================================

func Test_Generic_Collection_AddLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericAddLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 100)
		col := coredynamic.New.Collection.Generic.Empty()

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
// Test: Generic AddsLock — concurrent safety
// ==========================================

func Test_Generic_Collection_AddsLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericAddsLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 50)
		batch := input.GetAsIntDefault("batch", 3)
		col := coredynamic.New.Collection.Generic.Empty()

		// Act
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func(idx int) {
				items := make([]any, batch)
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
// Test: Generic LengthLock
// ==========================================

func Test_Generic_Collection_LengthLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericLengthLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsAnySlice("items")

		// Act
		col := coredynamic.New.Collection.Generic.From(items)
		actLines := []string{
			fmt.Sprintf("%d", col.LengthLock()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Generic IsEmptyLock — empty
// ==========================================

func Test_Generic_Collection_IsEmptyLock_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range genericIsEmptyLockEmptyTestCases {
		// Arrange — empty collection

		// Act
		col := coredynamic.New.Collection.Generic.Empty()
		actLines := []string{
			fmt.Sprintf("%v", col.IsEmptyLock()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Generic IsEmptyLock — non-empty
// ==========================================

func Test_Generic_Collection_IsEmptyLock_NonEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range genericIsEmptyLockNonEmptyTestCases {
		// Arrange

		// Act
		col := coredynamic.New.Collection.Generic.Items("x")
		actLines := []string{
			fmt.Sprintf("%v", col.IsEmptyLock()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Generic ItemsLock — returns independent copy
// ==========================================

func Test_Generic_Collection_ItemsLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericItemsLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsAnySlice("items")

		// Act
		col := coredynamic.New.Collection.Generic.From(items)
		copied := col.ItemsLock()
		copied = append(copied, "mutated")
		actLines := []string{
			fmt.Sprintf("%d", len(items)),
			fmt.Sprintf("%v", items[0]),
			fmt.Sprintf("%v", items[len(items)-1]),
			fmt.Sprintf("%v", col.Length() != len(copied)),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Generic ClearLock
// ==========================================

func Test_Generic_Collection_ClearLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericClearLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsAnySlice("items")

		// Act
		col := coredynamic.New.Collection.Generic.From(items)
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
// Test: Generic AddCollectionLock
// ==========================================

func Test_Generic_Collection_AddCollectionLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericAddCollectionLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first, _ := input.GetAsAnySlice("first")
		second, _ := input.GetAsAnySlice("second")

		// Act
		col1 := coredynamic.New.Collection.Generic.From(first)
		col2 := coredynamic.New.Collection.Generic.From(second)
		col1.AddCollectionLock(col2)
		actLines := []string{
			fmt.Sprintf("%d", col1.Length()),
			fmt.Sprintf("%v", col1.First()),
			fmt.Sprintf("%v", col1.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Generic FilterLock — concurrent safety
// ==========================================

func Test_Generic_Collection_FilterLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericFilterLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsAnySlice("items")
		col := coredynamic.New.Collection.Generic.From(items)

		// Act — filter strings starting with "a" or "d"
		wg := sync.WaitGroup{}
		wg.Add(5)
		for i := 0; i < 5; i++ {
			go func() {
				col.LengthLock()
				wg.Done()
			}()
		}

		filtered := col.FilterLock(func(item any) bool {
			s, ok := item.(string)
			if !ok {
				return false
			}
			return len(s) > 0 && (s[0] == 'a' || s[0] == 'd')
		})
		wg.Wait()

		actLines := []string{
			fmt.Sprintf("%d", filtered.Length()),
			fmt.Sprintf("%v", filtered.First()),
			fmt.Sprintf("%v", filtered.Last()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Generic LoopLock — concurrent safety
// ==========================================

func Test_Generic_Collection_LoopLock_Verification(t *testing.T) {
	for caseIndex, testCase := range genericLoopLockTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 50)
		col := coredynamic.New.Collection.Generic.Empty()
		for i := 0; i < count; i++ {
			col.Add(fmt.Sprintf("item-%d", i))
		}

		// Act — loop while concurrent writes happen
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func(idx int) {
				col.AddLock(fmt.Sprintf("extra-%d", idx))
				wg.Done()
			}(i)
		}

		visited := 0
		col.LoopLock(func(index int, item any) bool {
			visited++
			return false
		})
		wg.Wait()

		actLines := []string{
			fmt.Sprintf("%d", visited),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
