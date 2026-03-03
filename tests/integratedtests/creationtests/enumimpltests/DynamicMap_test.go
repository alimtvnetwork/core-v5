package enumimpltests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_DynamicMapDiff1(t *testing.T) {
	for caseIndex, tc := range dynamicMapSimpleDiffCaseV1TestCases {
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
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
