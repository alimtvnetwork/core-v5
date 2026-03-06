package coretestcasestests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var expectedLinesTestCases = []coretestcases.CaseV1{
	{
		Title:         "ExpectedLines converts int to string",
		ArrangeInput:  args.Map{"inputType": "int"},
		ExpectedInput: 42,
	},
	{
		Title:         "ExpectedLines converts bool true",
		ArrangeInput:  args.Map{"inputType": "boolTrue"},
		ExpectedInput: true,
	},
	{
		Title:         "ExpectedLines converts bool false",
		ArrangeInput:  args.Map{"inputType": "boolFalse"},
		ExpectedInput: false,
	},
	{
		Title:         "ExpectedLines converts []int slice",
		ArrangeInput:  args.Map{"inputType": "intSlice"},
		ExpectedInput: []int{10, 20, 30},
	},
	{
		Title:         "ExpectedLines converts []bool slice",
		ArrangeInput:  args.Map{"inputType": "boolSlice"},
		ExpectedInput: []bool{true, false, true},
	},
	{
		Title:         "ExpectedLines wraps string into slice",
		ArrangeInput:  args.Map{"inputType": "string"},
		ExpectedInput: "hello",
	},
	{
		Title:         "ExpectedLines returns []string as-is",
		ArrangeInput:  args.Map{"inputType": "stringSlice"},
		ExpectedInput: []string{"a", "b", "c"},
	},
	{
		Title:        "ExpectedLines converts map[string]int sorted",
		ArrangeInput: args.Map{"inputType": "mapStringInt"},
		ExpectedInput: map[string]int{
			"age":   30,
			"count": 5,
		},
	},
}

// expectedLinesExpectedOutputs holds the expected output for each test case
// as args.Map with lineCount + indexed line keys.
var expectedLinesExpectedOutputs = []args.Map{
	{"lineCount": "1", "line0": "42"},
	{"lineCount": "1", "line0": "true"},
	{"lineCount": "1", "line0": "false"},
	{"lineCount": "3", "line0": "10", "line1": "20", "line2": "30"},
	{"lineCount": "3", "line0": "true", "line1": "false", "line2": "true"},
	{"lineCount": "1", "line0": "hello"},
	{"lineCount": "3", "line0": "a", "line1": "b", "line2": "c"},
	{"lineCount": "2", "line0": "age : 30", "line1": "count : 5"},
}

// expectedLinesVerificationCases are CaseV1 instances used for ShouldBeEqualMap assertion.
var expectedLinesVerificationCases = []coretestcases.CaseV1{
	{Title: "ExpectedLines converts int to string", ExpectedInput: expectedLinesExpectedOutputs[0]},
	{Title: "ExpectedLines converts bool true", ExpectedInput: expectedLinesExpectedOutputs[1]},
	{Title: "ExpectedLines converts bool false", ExpectedInput: expectedLinesExpectedOutputs[2]},
	{Title: "ExpectedLines converts []int slice", ExpectedInput: expectedLinesExpectedOutputs[3]},
	{Title: "ExpectedLines converts []bool slice", ExpectedInput: expectedLinesExpectedOutputs[4]},
	{Title: "ExpectedLines wraps string into slice", ExpectedInput: expectedLinesExpectedOutputs[5]},
	{Title: "ExpectedLines returns []string as-is", ExpectedInput: expectedLinesExpectedOutputs[6]},
	{Title: "ExpectedLines converts map[string]int sorted", ExpectedInput: expectedLinesExpectedOutputs[7]},
}
