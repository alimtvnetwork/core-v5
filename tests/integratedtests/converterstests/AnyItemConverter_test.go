package converterstests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
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
