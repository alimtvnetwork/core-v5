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

type exampleStruct struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

var resultUnmarshalTestCases = []coretestcases.CaseV1{
	{Title: "Unmarshal - valid JSON deserializes correctly", ArrangeInput: args.Map{"case": "valid"}, ExpectedInput: []string{"<nil>", "Alice", "30"}},
	{Title: "Unmarshal - nil receiver returns error", ArrangeInput: args.Map{"case": "nil"}, ExpectedInput: []string{"true", "true"}},
	{Title: "Unmarshal - invalid bytes returns error", ArrangeInput: args.Map{"case": "invalid"}, ExpectedInput: []string{"true", "true"}},
	{Title: "Unmarshal - existing error propagates", ArrangeInput: args.Map{"case": "existing-error"}, ExpectedInput: []string{"true", "true"}},
}

func Test_Result_Unmarshal_Verification(t *testing.T) {
	for caseIndex, tc := range resultUnmarshalTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		caseType := input["case"].(string)

		var actLines []string

		// Act
		switch caseType {
		case "valid":
			src := exampleStruct{Name: "Alice", Age: 30}
			jsonResult := corejson.NewPtr(src)
			target := &exampleStruct{}
			err := jsonResult.Unmarshal(target)
			actLines = []string{fmt.Sprintf("%v", err), target.Name, fmt.Sprintf("%v", target.Age)}
		case "nil":
			var nilResult *corejson.Result
			target := &exampleStruct{}
			err := nilResult.Unmarshal(target)
			actLines = []string{fmt.Sprintf("%v", err != nil), fmt.Sprintf("%v", strings.Contains(err.Error(), "null"))}
		case "invalid":
			result := corejson.NewResult.UsingBytesTypePtr([]byte(`{invalid-json`), "TestType")
			target := &exampleStruct{}
			err := result.Unmarshal(target)
			actLines = []string{fmt.Sprintf("%v", err != nil), fmt.Sprintf("%v", strings.Contains(err.Error(), "unmarshal"))}
		case "existing-error":
			ch := make(chan int)
			result := corejson.NewPtr(ch)
			target := &exampleStruct{}
			err := result.Unmarshal(target)
			actLines = []string{fmt.Sprintf("%v", err != nil), fmt.Sprintf("%v", strings.Contains(err.Error(), "unmarshal"))}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.PrintDiffOnMismatch(caseIndex, tc.Title, actLines, expectedLines)
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
