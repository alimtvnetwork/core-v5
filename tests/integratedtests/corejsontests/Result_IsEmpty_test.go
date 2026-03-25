package corejsontests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var resultIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmpty - empty bytes returns true",
		ArrangeInput: args.Map{
			"result": func() *corejson.Result {
				r := corejson.NewResult.UsingBytes([]byte{})
				return &r
			}(),
		},
		ExpectedInput: "true", // isEmpty
	},
	{
		Title: "IsEmpty - nil receiver returns true",
		ArrangeInput: args.Map{
			"result": (*corejson.Result)(nil),
		},
		ExpectedInput: "true", // isEmpty
	},
	{
		Title: "IsEmpty - valid bytes returns false",
		ArrangeInput: args.Map{
			"result": func() *corejson.Result {
				r := corejson.New(map[string]string{"key": "value"})
				return &r
			}(),
		},
		ExpectedInput: "false", // isEmpty
	},
}

func Test_Result_IsEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range resultIsEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		result := input["result"].(*corejson.Result)

		// Act
		actual := fmt.Sprintf("%v", result.IsEmpty())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actual)
	}
}
