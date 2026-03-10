package mapdiffinternal

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var isStringTypeInternalTestCases = []coretestcases.CaseV1{
	{
		Title: "String value returns true",
		ArrangeInput: args.Map{
			"value": "hello",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Int value returns false",
		ArrangeInput: args.Map{
			"value": 42,
		},
		ExpectedInput: "false",
	},
	{
		Title: "Bool value returns false",
		ArrangeInput: args.Map{
			"value": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "Nil value returns false",
		ArrangeInput: args.Map{
			"value": nil,
		},
		ExpectedInput: "false",
	},
	{
		Title: "Empty string returns true",
		ArrangeInput: args.Map{
			"value": "",
		},
		ExpectedInput: "true",
	},
}

func Test_isStringType(t *testing.T) {
	for caseIndex, testCase := range isStringTypeInternalTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		value := input["value"]

		// Act
		result := isStringType(value)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}
