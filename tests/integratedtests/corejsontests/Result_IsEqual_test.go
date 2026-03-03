package corejsontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_Result_IsEqual(t *testing.T) {
	for caseIndex, tc := range resultIsEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a := input["a"].(corejson.Result)
		b := input["b"].(corejson.Result)

		// Act
		result := fmt.Sprintf("%v", a.IsEqual(b))

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_Result_IsEqualPtr(t *testing.T) {
	for caseIndex, tc := range resultIsEqualPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a := input["aPtr"].(*corejson.Result)
		b := input["bPtr"].(*corejson.Result)

		// Act
		result := fmt.Sprintf("%v", a.IsEqualPtr(b))

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}
