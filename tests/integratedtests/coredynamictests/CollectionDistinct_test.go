package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: Distinct
// ==========================================

func Test_Collection_Distinct_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDistinctTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		result := coredynamic.Distinct(col)

		// Handle mixed ExpectedInput types
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actual := args.Map{
				"distinctCount": result.Length(),
			}
			for i := 0; i < result.Length(); i++ {
				actual[fmt.Sprintf("item%d", i)] = result.SafeAt(i)
			}

			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result.Length()))
		}
	}
}

// ==========================================
// Test: DistinctCount
// ==========================================

func Test_Collection_DistinctCount_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDistinctCountTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", coredynamic.DistinctCount(col)))
	}
}

// ==========================================
// Test: IsDistinct
// ==========================================

func Test_Collection_IsDistinct_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsDistinctTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		if !isValid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", coredynamic.IsDistinct(col)))
	}
}
