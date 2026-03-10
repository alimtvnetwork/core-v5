package mapdiffinternaltests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/mapdiffinternal"
)

func Test_MapStringAnyDiff_Length_IsEmpty_HasAnyItem_LastIndex(t *testing.T) {
	for caseIndex, testCase := range mapStringAnyDiffLengthTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		raw := input["map"].(map[string]any)
		diff := mapdiffinternal.MapStringAnyDiff(raw)

		// Act
		actual := args.Map{
			"length":     diff.Length(),
			"isEmpty":    diff.IsEmpty(),
			"hasAnyItem": diff.HasAnyItem(),
			"lastIndex":  diff.LastIndex(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringAnyDiff_AllKeysSorted(t *testing.T) {
	for caseIndex, testCase := range mapStringAnyDiffAllKeysSortedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		raw := input["map"].(map[string]any)
		diff := mapdiffinternal.MapStringAnyDiff(raw)

		// Act
		keys := diff.AllKeysSorted()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, keys...)
	}
}

func Test_MapStringAnyDiff_IsRawEqual(t *testing.T) {
	for caseIndex, testCase := range mapStringAnyDiffIsRawEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isRegardless := input["isRegardlessType"].(bool)
		left := input["left"].(map[string]any)
		right := input["right"].(map[string]any)
		diff := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		isEqual := diff.IsRawEqual(isRegardless, right)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", isEqual),
		)
	}
}

func Test_MapStringAnyDiff_HasAnyChanges(t *testing.T) {
	for caseIndex, testCase := range mapStringAnyDiffHasAnyChangesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isRegardless := input["isRegardlessType"].(bool)
		left := input["left"].(map[string]any)
		right := input["right"].(map[string]any)
		diff := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		hasChanges := diff.HasAnyChanges(isRegardless, right)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", hasChanges),
		)
	}
}

func Test_MapStringAnyDiff_DiffRaw(t *testing.T) {
	for caseIndex, testCase := range mapStringAnyDiffDiffRawTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isRegardless := input["isRegardlessType"].(bool)

		var diff *mapdiffinternal.MapStringAnyDiff
		var right map[string]any

		leftNil, hasLeftNil := input["leftNil"]
		isLeftNil := hasLeftNil && leftNil.(bool)

		rightNil, hasRightNil := input["rightNil"]
		isRightNil := hasRightNil && rightNil.(bool)

		if !isLeftNil {
			left := input["left"].(map[string]any)
			msad := mapdiffinternal.MapStringAnyDiff(left)
			diff = &msad
		}

		if !isRightNil {
			right = input["right"].(map[string]any)
		}

		// Act
		diffMap := diff.DiffRaw(isRegardless, right)

		// Assert
		actual := args.Map{
			"diffLength": len(diffMap),
		}

		for key := range diffMap {
			actual["hasKey-"+key] = true
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MapStringAnyDiff_ShouldDiffMessage(t *testing.T) {
	for caseIndex, testCase := range mapStringAnyDiffShouldDiffMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isRegardless := input["isRegardlessType"].(bool)
		title := input["title"].(string)
		left := input["left"].(map[string]any)
		right := input["right"].(map[string]any)
		diff := mapdiffinternal.MapStringAnyDiff(left)

		// Act
		message := diff.ShouldDiffMessage(isRegardless, title, right)

		// Assert
		switch testCase.ExpectedInput.(type) {
		case string:
			testCase.ShouldBeEqual(
				t,
				caseIndex,
				message,
			)

		case args.Map:
			actual := args.Map{
				"containsTitle": strings.Contains(message, title),
				"isNotEmpty":    message != "",
			}

			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

func Test_MapStringAnyDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	for caseIndex, testCase := range mapStringAnyDiffToStringsSliceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		diffMap := input["diffMap"].(map[string]any)
		base := mapdiffinternal.MapStringAnyDiff(map[string]any{})

		// Act
		slice := base.ToStringsSliceOfDiffMap(diffMap)

		// Assert
		hasQuotation := false
		if len(slice) > 0 {
			// string values get double-quoted value: "key":"val"
			// non-string values get raw: "key":42
			hasQuotation = strings.Count(slice[0], "\"") >= 4
		}

		actual := args.Map{
			"length":       len(slice),
			"hasQuotation": hasQuotation,
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
