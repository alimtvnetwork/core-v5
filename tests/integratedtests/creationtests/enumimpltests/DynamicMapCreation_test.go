package enumimpltests

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_DynamicMapCreationDiff(t *testing.T) {
	for caseIndex, tc := range dynamicMapDiffCaseV1TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left := enumimpl.DynamicMap(input["left"].(enumimpl.DynamicMap))
		right := enumimpl.DynamicMap(input["right"].(enumimpl.DynamicMap))

		// Act
		diffMap := left.DiffRaw(true, right)
		mapAnyDiffer := coredynamic.MapAnyItemDiff(left)
		anotherDiff := mapAnyDiffer.DiffRaw(true, right)

		// Assert - verify both diffs produce sorted key:value lines
		actLines := dynamicMapToSortedLines(diffMap)
		expectedLines := tc.ExpectedInput.([]string)

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)

		// Assert - verify both diff methods produce equal raw maps
		anotherLines := dynamicMapToSortedLines(enumimpl.DynamicMap(anotherDiff))
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title+" (both diff equal)", anotherLines, actLines)

		if len(actLines) != len(anotherLines) {
			t.Errorf("[case %d] %s: both diff methods line count mismatch got %d, want %d",
				caseIndex, tc.Title, len(anotherLines), len(actLines))
		}
	}
}

func Test_DynamicMapCreationDiffMessage(t *testing.T) {
	for caseIndex, tc := range dynamicMapDiffMessageCaseV1TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left := enumimpl.DynamicMap(input["left"].(map[string]any))
		right := input["right"].(map[string]any)

		// Act
		diffJsonMessage := left.ShouldDiffMessage(
			true,
			tc.Title,
			right,
		)
		actLines := coretests.GetAssert.ToStrings(diffJsonMessage)

		// Assert
		expectedLines := tc.ExpectedInput.([]string)
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_DynamicMapCreationDiffMessageV2(t *testing.T) {
	for caseIndex, tc := range dynamicMapDiffMessageV2CaseV1TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left := enumimpl.DynamicMap(input["left"].(map[string]any))
		right := input["right"].(map[string]any)
		checker := input["checker"].(enumimpl.DifferChecker)

		// Act
		diffJsonMessage := left.ShouldDiffLeftRightMessageUsingDifferChecker(
			checker,
			true,
			tc.Title,
			right,
		)
		actLines := strings.Split(
			diffJsonMessage,
			constants.NewLineUnix,
		)

		// Assert
		expectedLines := tc.ExpectedInput.([]string)
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func dynamicMapToSortedLines(dm enumimpl.DynamicMap) []string {
	if dm.IsEmpty() {
		return []string{}
	}

	keys := dm.AllKeysSorted()
	sort.Strings(keys)

	lines := make([]string, len(keys))
	for i, k := range keys {
		lines[i] = fmt.Sprintf("%s : %v", k, dm[k])
	}

	return lines
}
