package mapdiffinternaltests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/internal/mapdiffinternal"
)

func Test_HashmapDiff_Length_IsEmpty_HasAnyItem_LastIndex(t *testing.T) {
	for caseIndex, testCase := range hashmapDiffLengthTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		raw := input["map"].(map[string]string)
		diff := mapdiffinternal.HashmapDiff(raw)

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

func Test_HashmapDiff_AllKeysSorted(t *testing.T) {
	for caseIndex, testCase := range hashmapDiffAllKeysSortedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		raw := input["map"].(map[string]string)
		diff := mapdiffinternal.HashmapDiff(raw)

		// Act
		keys := diff.AllKeysSorted()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, keys...)
	}
}

func Test_HashmapDiff_IsRawEqual(t *testing.T) {
	for caseIndex, testCase := range hashmapDiffIsRawEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		left := input["left"].(map[string]string)
		right := input["right"].(map[string]string)
		diff := mapdiffinternal.HashmapDiff(left)

		// Act
		isEqual := diff.IsRawEqual(right)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", isEqual),
		)
	}
}

func Test_HashmapDiff_DiffRaw(t *testing.T) {
	for caseIndex, testCase := range hashmapDiffDiffRawTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		var diff *mapdiffinternal.HashmapDiff
		var right map[string]string

		leftNil, hasLeftNil := input["leftNil"]
		isLeftNil := hasLeftNil && leftNil.(bool)

		rightNil, hasRightNil := input["rightNil"]
		isRightNil := hasRightNil && rightNil.(bool)

		if !isLeftNil {
			left := input["left"].(map[string]string)
			hd := mapdiffinternal.HashmapDiff(left)
			diff = &hd
		}

		if !isRightNil {
			right = input["right"].(map[string]string)
		}

		// Act
		diffMap := diff.DiffRaw(right)

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

func Test_HashmapDiff_ShouldDiffMessage(t *testing.T) {
	for caseIndex, testCase := range hashmapDiffShouldDiffMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		title := input["title"].(string)
		left := input["left"].(map[string]string)
		right := input["right"].(map[string]string)
		diff := mapdiffinternal.HashmapDiff(left)

		// Act
		message := diff.ShouldDiffMessage(title, right)

		// Assert
		expected := testCase.ExpectedInput

		switch expected.(type) {
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
