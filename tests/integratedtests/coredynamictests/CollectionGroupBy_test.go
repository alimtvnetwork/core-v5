package coredynamictests

import (
	"fmt"
	"sort"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: GroupBy
// ==========================================

func Test_Collection_GroupBy_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionGroupByTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		groups := coredynamic.GroupBy(col, func(s string) string {
			if len(s) == 0 {
				return ""
			}
			return string(s[0])
		})

		// Assert
		actLines := make([]string, 0, len(groups))
		for key, group := range groups {
			actLines = append(actLines, fmt.Sprintf("%s:%d", key, group.Length()))
		}
		sort.Strings(actLines)

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: GroupByCount
// ==========================================

func Test_Collection_GroupByCount_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionGroupByCountTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		// Act
		col := coredynamic.New.Collection.String.From(items)
		counts := coredynamic.GroupByCount(col, func(s string) string {
			return s
		})

		// Assert
		actLines := make([]string, 0, len(counts))
		for key, count := range counts {
			actLines = append(actLines, fmt.Sprintf("%s:%d", key, count))
		}
		sort.Strings(actLines)

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
