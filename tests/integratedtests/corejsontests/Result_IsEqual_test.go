package corejsontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var resultIsEqualTestCases = []coretestcases.CaseV1{
	{Title: "IsEqual - same content returns true", ArrangeInput: args.Map{"method": "IsEqual", "a": corejson.New(map[string]string{"key": "value"}), "b": corejson.New(map[string]string{"key": "value"})}, ExpectedInput: []string{"true"}},
	{Title: "IsEqual - different content returns false", ArrangeInput: args.Map{"method": "IsEqual", "a": corejson.New(map[string]string{"key": "a"}), "b": corejson.New(map[string]string{"key": "b"})}, ExpectedInput: []string{"false"}},
	{Title: "IsEqualPtr - both nil returns true", ArrangeInput: args.Map{"method": "IsEqualPtr", "aPtr": (*corejson.Result)(nil), "bPtr": (*corejson.Result)(nil)}, ExpectedInput: []string{"true"}},
	{Title: "IsEqualPtr - one nil returns false", ArrangeInput: args.Map{"method": "IsEqualPtrOneNil", "aPtr": corejson.NewPtr(map[string]string{"k": "v"}), "bPtr": (*corejson.Result)(nil)}, ExpectedInput: []string{"false"}},
}

func Test_Result_IsEqual_Verification(t *testing.T) {
	for caseIndex, tc := range resultIsEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method := input["method"].(string)

		var actLines []string

		// Act
		switch method {
		case "IsEqual":
			a := input["a"].(corejson.Result)
			b := input["b"].(corejson.Result)
			actLines = []string{fmt.Sprintf("%v", a.IsEqual(b))}
		case "IsEqualPtr":
			a := input["aPtr"].(*corejson.Result)
			b := input["bPtr"].(*corejson.Result)
			actLines = []string{fmt.Sprintf("%v", a.IsEqualPtr(b))}
		case "IsEqualPtrOneNil":
			a := input["aPtr"].(*corejson.Result)
			b := input["bPtr"].(*corejson.Result)
			actLines = []string{fmt.Sprintf("%v", a.IsEqualPtr(b))}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}
