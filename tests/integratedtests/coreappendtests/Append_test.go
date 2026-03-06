package coreappendtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreappend"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_PrependAppendAnyItems_Verification(t *testing.T) {
	for caseIndex, testCase := range prependAppendTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		prependVal, hasPrepend := input.Get("prepend")
		appendVal, _ := input.Get("append")

		var prepend any
		if hasPrepend {
			prepend = prependVal
		}

		// Act
		result := coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
			prepend,
			appendVal,
			"middle",
		)

		actual := args.Map{
			"totalCount": fmt.Sprintf("%v", len(result)),
		}
		if len(result) > 0 {
			actual["firstItem"] = result[0]
		}
		if len(result) > 1 {
			actual["lastItem"] = result[len(result)-1]
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
