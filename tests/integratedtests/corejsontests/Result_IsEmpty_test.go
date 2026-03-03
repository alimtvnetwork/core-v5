package corejsontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var resultIsEmptyTestCases = []coretestcases.CaseV1{
	{Title: "IsEmpty - empty bytes returns true", ArrangeInput: args.Map{"result": corejson.NewResult.UsingBytes([]byte{})}, ExpectedInput: []string{"true"}},
	{Title: "IsEmpty - nil receiver returns true", ArrangeInput: args.Map{"result": (*corejson.Result)(nil)}, ExpectedInput: []string{"true"}},
	{Title: "IsEmpty - valid bytes returns false", ArrangeInput: args.Map{"result": func() *corejson.Result { r := corejson.New(map[string]string{"key": "value"}); return &r }()}, ExpectedInput: []string{"false"}},
}

func Test_Result_IsEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range resultIsEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		result := input["result"].(*corejson.Result)

		// Act
		actLines := []string{fmt.Sprintf("%v", result.IsEmpty())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.PrintLineDiff(caseIndex, tc.Title, actLines, expectedLines)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
