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

// expectedLinesExpectedOutputs maps each test case index to its expected
// output lines expressed as args.Map with lineCount + indexed keys.
var expectedLinesExpectedOutputs = []args.Map{
	// int → ["42"]
	{"lineCount": "1", "line0": "42"},
	// bool true → ["true"]
	{"lineCount": "1", "line0": "true"},
	// bool false → ["false"]
	{"lineCount": "1", "line0": "false"},
	// []int → ["10","20","30"]
	{"lineCount": "3", "line0": "10", "line1": "20", "line2": "30"},
	// []bool → ["true","false","true"]
	{"lineCount": "3", "line0": "true", "line1": "false", "line2": "true"},
	// string → ["hello"]
	{"lineCount": "1", "line0": "hello"},
	// []string → ["a","b","c"]
	{"lineCount": "3", "line0": "a", "line1": "b", "line2": "c"},
	// map[string]int → sorted key : value lines
	{"lineCount": "2", "line0": "age : 30", "line1": "count : 5"},
}
