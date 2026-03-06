package coredynamictests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: Map — int to string
// ==========================================

func Test_Map_IntToString_Verification(t *testing.T) {
	for caseIndex, testCase := range mapIntToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.From(items)
		result := coredynamic.Map(col, func(i int) string {
			return fmt.Sprintf("#%d", i)
		})
		actLines := []string{fmt.Sprintf("%d", result.Length())}
		actLines = append(actLines, result.Items()...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Map — empty collection
// ==========================================

func Test_Map_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range mapEmptyTestCases {
		// Act
		col := coredynamic.New.Collection.Int.Empty()
		result := coredynamic.Map(col, func(i int) string {
			return fmt.Sprintf("%d", i)
		})

		actual := args.Map{
			"length":  result.Length(),
			"isEmpty": result.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Map — nil collection
// ==========================================

func Test_Map_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range mapNilTestCases {
		// Arrange
		var col *coredynamic.IntCollection

		// Act
		result := coredynamic.Map(col, func(i int) string {
			return fmt.Sprintf("%d", i)
		})

		actual := args.Map{
			"length":  result.Length(),
			"isEmpty": result.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Map — string to int (length)
// ==========================================

func Test_Map_StringToInt_Verification(t *testing.T) {
	for caseIndex, testCase := range mapStringToIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		result := coredynamic.Map(col, func(s string) int {
			return len(s)
		})
		actLines := []string{fmt.Sprintf("%d", result.Length())}
		for _, item := range result.Items() {
			actLines = append(actLines, fmt.Sprintf("%d", item))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: FlatMap — string to chars
// ==========================================

func Test_FlatMap_Verification(t *testing.T) {
	for caseIndex, testCase := range flatMapTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		result := coredynamic.FlatMap(col, func(s string) []string {
			chars := make([]string, len(s))
			for i, c := range s {
				chars[i] = string(c)
			}
			return chars
		})
		actLines := []string{fmt.Sprintf("%d", result.Length())}
		actLines = append(actLines, result.Items()...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: FlatMap — empty
// ==========================================

func Test_FlatMap_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range flatMapEmptyTestCases {
		// Act
		col := coredynamic.New.Collection.String.Empty()
		result := coredynamic.FlatMap(col, func(s string) []string {
			return strings.Split(s, "")
		})

		actual := args.Map{
			"length":  result.Length(),
			"isEmpty": result.IsEmpty(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// Test: Reduce — sum
// ==========================================

func Test_Reduce_Sum_Verification(t *testing.T) {
	for caseIndex, testCase := range reduceSumTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.From(items)
		result := coredynamic.Reduce(col, 0, func(acc int, item int) int {
			return acc + item
		})

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// Test: Reduce — empty returns initial
// ==========================================

func Test_Reduce_Empty_Verification(t *testing.T) {
	for caseIndex, testCase := range reduceEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		initial := input.GetAsIntDefault("initial", 0)

		// Act
		col := coredynamic.New.Collection.Int.Empty()
		result := coredynamic.Reduce(col, initial, func(acc int, item int) int {
			return acc + item
		})

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// Test: Reduce — string concat
// ==========================================

func Test_Reduce_Concat_Verification(t *testing.T) {
	for caseIndex, testCase := range reduceConcatTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		result := coredynamic.Reduce(col, "", func(acc string, item string) string {
			if acc == "" {
				return item
			}
			return acc + "-" + item
		})

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// Test: Map then Filter — chained operations
// ==========================================

func Test_Map_Then_Filter_Verification(t *testing.T) {
	for caseIndex, testCase := range mapThenFilterTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items := input["items"].([]int)

		// Act
		col := coredynamic.New.Collection.Int.From(items)
		doubled := coredynamic.Map(col, func(i int) int {
			return i * 2
		})
		filtered := doubled.Filter(func(i int) bool {
			return i > 5
		})
		actLines := []string{fmt.Sprintf("%d", filtered.Length())}
		for _, item := range filtered.Items() {
			actLines = append(actLines, fmt.Sprintf("%d", item))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
