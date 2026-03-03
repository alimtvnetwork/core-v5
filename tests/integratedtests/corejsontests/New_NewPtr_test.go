package corejsontests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var newNewPtrTestCases = []coretestcases.CaseV1{
	{Title: "New - valid struct produces bytes no error", ArrangeInput: args.Map{"case": "new-valid"}, ExpectedInput: []string{"false", "false", "true", "true"}},
	{Title: "New - nil input produces null bytes", ArrangeInput: args.Map{"case": "new-nil"}, ExpectedInput: []string{"false", "null"}},
	{Title: "New - channel produces error", ArrangeInput: args.Map{"case": "new-channel"}, ExpectedInput: []string{"true", "true"}},
	{Title: "NewPtr - valid struct produces non-nil result", ArrangeInput: args.Map{"case": "newptr-valid"}, ExpectedInput: []string{"true", "false", "false", "true"}},
	{Title: "NewPtr - nil input produces null bytes", ArrangeInput: args.Map{"case": "newptr-nil"}, ExpectedInput: []string{"true", "false", "null"}},
	{Title: "NewPtr - channel produces error", ArrangeInput: args.Map{"case": "newptr-channel"}, ExpectedInput: []string{"true", "true", "true"}},
}

func Test_New_NewPtr_Verification(t *testing.T) {
	for caseIndex, tc := range newNewPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		caseType := input["case"].(string)

		var actLines []string

		// Act
		switch caseType {
		case "new-valid":
			result := corejson.New(struct {
				Name string
				Age  int
			}{Name: "Alice", Age: 30})
			actLines = []string{
				fmt.Sprintf("%v", result.HasError()),
				fmt.Sprintf("%v", result.IsEmpty()),
				fmt.Sprintf("%v", len(result.Bytes) > 0),
				fmt.Sprintf("%v", result.TypeName != ""),
			}
		case "new-nil":
			result := corejson.New(nil)
			actLines = []string{fmt.Sprintf("%v", result.HasError()), string(result.Bytes)}
		case "new-channel":
			result := corejson.New(make(chan int))
			actLines = []string{fmt.Sprintf("%v", result.HasError()), fmt.Sprintf("%v", strings.Contains(result.Error.Error(), "marshal"))}
		case "newptr-valid":
			result := corejson.NewPtr(struct {
				Name string
				Age  int
			}{Name: "Bob", Age: 25})
			actLines = []string{
				fmt.Sprintf("%v", result != nil),
				fmt.Sprintf("%v", result.HasError()),
				fmt.Sprintf("%v", result.IsEmpty()),
				fmt.Sprintf("%v", len(result.Bytes) > 0),
			}
		case "newptr-nil":
			result := corejson.NewPtr(nil)
			actLines = []string{fmt.Sprintf("%v", result != nil), fmt.Sprintf("%v", result.HasError()), string(result.Bytes)}
		case "newptr-channel":
			result := corejson.NewPtr(make(chan string))
			actLines = []string{
				fmt.Sprintf("%v", result != nil),
				fmt.Sprintf("%v", result.HasError()),
				fmt.Sprintf("%v", strings.Contains(result.Error.Error(), "marshal")),
			}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}
