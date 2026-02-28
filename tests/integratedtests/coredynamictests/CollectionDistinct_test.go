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
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		result := coredynamic.Distinct(col)

		actLines := []string{
			fmt.Sprintf("%d", result.Length()),
		}
		for i := 0; i < result.Length(); i++ {
			actLines = append(actLines, result.SafeAt(i))
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: DistinctCount
// ==========================================

func Test_Collection_DistinctCount_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionDistinctCountTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%d", coredynamic.DistinctCount(col)),
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsDistinct
// ==========================================

func Test_Collection_IsDistinct_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionIsDistinctTestCases {
		input := testCase.ArrangeInput.(args.Map)
		items, isValid := input.GetAsStrings("items")
		isInvalid := !isValid

		if isInvalid {
			errcore.HandleErrMessage("GetAsStrings 'items' failed")
		}

		col := coredynamic.New.Collection.String.From(items)
		actLines := []string{
			fmt.Sprintf("%v", coredynamic.IsDistinct(col)),
		}

		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
