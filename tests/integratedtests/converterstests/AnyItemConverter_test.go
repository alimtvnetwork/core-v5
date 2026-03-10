package converterstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// =============================================================================
// Tests: ToNonNullItems
// =============================================================================

func Test_AnyItemConverter_ToNonNullItems_SkipNil(t *testing.T) {
	for caseIndex, testCase := range toNonNullItemsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isSkipRaw, skipFound := input.Get("isSkipOnNil")
		if !skipFound {
			errcore.HandleErrMessage("isSkipOnNil is required")
		}
		isSkip := isSkipRaw.(bool)

		inputRaw, _ := input.Get("input")

		// Act
		result := converters.AnyItem.ToNonNullItems(isSkip, inputRaw)
		actual := args.Map{
			"count": len(result),
		}
		for i, item := range result {
			actual[fmt.Sprintf("item%d", i)] = fmt.Sprintf("%v", item)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
