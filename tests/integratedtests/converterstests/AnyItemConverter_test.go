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
		isSkipRaw, skipErr := input.Get("isSkipOnNil")
		errcore.HandleErrMessage("isSkipOnNil", skipErr)
		isSkip := isSkipRaw.(bool)

		inputRaw, _ := input.Get("input")

		// Act
		result := converters.AnyItem.ToNonNullItems(isSkip, inputRaw)
		actLines := []string{fmt.Sprintf("%d", len(result))}

		for _, item := range result {
			actLines = append(actLines, fmt.Sprintf("%v", item))
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
