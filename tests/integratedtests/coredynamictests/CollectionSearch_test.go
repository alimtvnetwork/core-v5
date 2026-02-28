package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: Contains
// ==========================================

func Test_Collection_Contains_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionContainsTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")
		search, _ := input.GetAsString("search")

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.Contains(col, search)),
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IndexOf
// ==========================================

func Test_Collection_IndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIndexOfTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")
		search, _ := input.GetAsString("search")

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.IndexOf(col, search)),
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: HasAll
// ==========================================

func Test_Collection_HasAll_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionHasAllTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")
		search, _ := input.GetAsStrings("search")

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.HasAll(col, search...)),
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: LastIndexOf
// ==========================================

func Test_Collection_LastIndexOf_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionLastIndexOfTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")
		search, _ := input.GetAsString("search")

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.LastIndexOf(col, search)),
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: Count
// ==========================================

func Test_Collection_Count_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionCountTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.GetAsStrings("items")
		search, _ := input.GetAsString("search")

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.Count(col, search)),
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
