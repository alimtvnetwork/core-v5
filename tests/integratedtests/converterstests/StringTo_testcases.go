package converterstests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var stringToIntegerTestCases = []coretestcases.CaseV1{
	{
		Title: "StringTo.Integer parses valid integer",
		ArrangeInput: args.Map{
			"when":  "given valid integer string",
			"input": "42",
		},
		ExpectedInput: []string{"42", "false"},
	},
	{
		Title: "StringTo.Integer fails on non-numeric string",
		ArrangeInput: args.Map{
			"when":  "given non-numeric string",
			"input": "abc",
		},
		ExpectedInput: []string{"0", "true"},
	},
	{
		Title: "StringTo.Integer parses negative integer",
		ArrangeInput: args.Map{
			"when":  "given negative integer string",
			"input": "-5",
		},
		ExpectedInput: []string{"-5", "false"},
	},
	{
		Title: "StringTo.Integer parses zero",
		ArrangeInput: args.Map{
			"when":  "given zero string",
			"input": "0",
		},
		ExpectedInput: []string{"0", "false"},
	},
	{
		Title: "StringTo.Integer fails on empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: []string{"0", "true"},
	},
	{
		Title: "StringTo.Integer fails on float string",
		ArrangeInput: args.Map{
			"when":  "given float string",
			"input": "3.14",
		},
		ExpectedInput: []string{"0", "true"},
	},
}

var bytesToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "BytesTo.String converts bytes to string",
		ArrangeInput: args.Map{
			"when":  "given valid byte slice",
			"input": "hello",
		},
		ExpectedInput: "hello",
	},
	{
		Title: "BytesTo.String returns empty for empty bytes",
		ArrangeInput: args.Map{
			"when":  "given empty byte slice",
			"input": "",
		},
		ExpectedInput: "",
	},
}

var stringToIntegerWithDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegerWithDefault returns parsed value on valid input",
		ArrangeInput: args.Map{
			"when":       "given valid integer",
			"input":      "100",
			"defaultInt": -1,
		},
		ExpectedInput: []string{"100", "true"},
	},
	{
		Title: "IntegerWithDefault returns default on invalid input",
		ArrangeInput: args.Map{
			"when":       "given non-numeric",
			"input":      "xyz",
			"defaultInt": -1,
		},
		ExpectedInput: []string{"-1", "false"},
	},
	{
		Title: "IntegerWithDefault returns default on empty string",
		ArrangeInput: args.Map{
			"when":       "given empty string",
			"input":      "",
			"defaultInt": 42,
		},
		ExpectedInput: []string{"42", "false"},
	},
}

var stringToFloat64TestCases = []coretestcases.CaseV1{
	{
		Title: "Float64 parses valid float",
		ArrangeInput: args.Map{
			"when":  "given valid float string",
			"input": "3.14",
		},
		ExpectedInput: []string{"3.14", "false"},
	},
	{
		Title: "Float64 parses integer as float",
		ArrangeInput: args.Map{
			"when":  "given integer string",
			"input": "42",
		},
		ExpectedInput: []string{"42", "false"},
	},
	{
		Title: "Float64 fails on non-numeric",
		ArrangeInput: args.Map{
			"when":  "given non-numeric string",
			"input": "abc",
		},
		ExpectedInput: []string{"0", "true"},
	},
	{
		Title: "Float64 parses negative float",
		ArrangeInput: args.Map{
			"when":  "given negative float",
			"input": "-2.5",
		},
		ExpectedInput: []string{"-2.5", "false"},
	},
}

var stringToByteTestCases = []coretestcases.CaseV1{
	{
		Title: "Byte parses valid byte value",
		ArrangeInput: args.Map{
			"when":  "given valid byte string",
			"input": "255",
		},
		ExpectedInput: []string{"255", "false"},
	},
	{
		Title: "Byte parses zero",
		ArrangeInput: args.Map{
			"when":  "given zero string",
			"input": "0",
		},
		ExpectedInput: []string{"0", "false"},
	},
	{
		Title: "Byte parses one",
		ArrangeInput: args.Map{
			"when":  "given one string",
			"input": "1",
		},
		ExpectedInput: []string{"1", "false"},
	},
	{
		Title: "Byte fails on empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: []string{"0", "true"},
	},
	{
		Title: "Byte fails on value > 255",
		ArrangeInput: args.Map{
			"when":  "given value exceeding byte range",
			"input": "256",
		},
		ExpectedInput: []string{"0", "true"},
	},
	{
		Title: "Byte fails on negative value",
		ArrangeInput: args.Map{
			"when":  "given negative value",
			"input": "-1",
		},
		ExpectedInput: []string{"0", "true"},
	},
}

var bytesToPtrStringTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrString returns string from valid pointer",
		ArrangeInput: args.Map{
			"when":  "given valid byte slice pointer",
			"input": "test-data",
			"isNil": false,
		},
		ExpectedInput: "test-data",
	},
	{
		Title: "PtrString returns empty for nil pointer",
		ArrangeInput: args.Map{
			"when":  "given nil pointer",
			"input": "",
			"isNil": true,
		},
		ExpectedInput: "",
	},
}

var stringsToHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "Hashset creates map with all unique entries true",
		ArrangeInput: args.Map{
			"when":  "given distinct strings",
			"input": []string{"a", "b", "c"},
		},
		ExpectedInput: []string{"3", "true", "true", "true"},
	},
	{
		Title: "Hashset handles duplicates",
		ArrangeInput: args.Map{
			"when":  "given duplicate strings",
			"input": []string{"a", "a", "b"},
		},
		ExpectedInput: []string{"2", "true", "true"},
	},
	{
		Title: "Hashset handles empty slice",
		ArrangeInput: args.Map{
			"when":  "given empty slice",
			"input": []string{},
		},
		ExpectedInput: "0",
	},
}

var stringToIntegerDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegerDefault returns parsed value",
		ArrangeInput: args.Map{
			"when":  "given valid number",
			"input": "77",
		},
		ExpectedInput: "77",
	},
	{
		Title: "IntegerDefault returns 0 on invalid",
		ArrangeInput: args.Map{
			"when":  "given non-number",
			"input": "nope",
		},
		ExpectedInput: "0",
	},
}

