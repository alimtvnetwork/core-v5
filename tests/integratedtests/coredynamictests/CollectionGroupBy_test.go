package coredynamictests

import (
	"fmt"
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
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
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actLines := make([]string, 0, len(groups))
			for key, group := range groups {
				actLines = append(actLines, fmt.Sprintf("%s:%d", key, group.Length()))
			}
			sort.Strings(actLines)

			actual := args.Map{}
			for i, line := range actLines {
				actual[fmt.Sprintf("group%c", 'A'+i)] = line
			}

			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			actLines := make([]string, 0, len(groups))
			for key, group := range groups {
				actLines = append(actLines, fmt.Sprintf("%s:%d", key, group.Length()))
			}
			sort.Strings(actLines)

			testCase.ShouldBeEqual(t, caseIndex, actLines...)
		}
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
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actLines := make([]string, 0, len(counts))
			for key, count := range counts {
				actLines = append(actLines, fmt.Sprintf("%s:%d", key, count))
			}
			sort.Strings(actLines)

			actual := args.Map{}
			for i, line := range actLines {
				keys := []string{"blueCount", "greenCount", "redCount"}
				if i < len(keys) {
					actual[keys[i]] = line
				}
			}

			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", len(counts)))
		}
	}
}
