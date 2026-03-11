package converterstests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var stringToIntegerTestCases = []coretestcases.CaseV1{
	{
		Title: "StringTo.Integer returns 42 -- valid integer string '42'",
		ArrangeInput: args.Map{
			"when":  "given valid integer string",
			"input": "42",
		},
		ExpectedInput: args.Map{"value": "42", "hasError": "false"},
	},
	{
		Title: "StringTo.Integer returns error -- non-numeric string 'abc'",
		ArrangeInput: args.Map{
			"when":  "given non-numeric string",
			"input": "abc",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "true"},
	},
	{
		Title: "StringTo.Integer returns -5 -- negative integer string '-5'",
		ArrangeInput: args.Map{
			"when":  "given negative integer string",
			"input": "-5",
		},
		ExpectedInput: args.Map{"value": "-5", "hasError": "false"},
	},
	{
		Title: "StringTo.Integer returns 0 -- zero string '0'",
		ArrangeInput: args.Map{
			"when":  "given zero string",
			"input": "0",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "false"},
	},
	{
		Title: "StringTo.Integer returns error -- empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "true"},
	},
	{
		Title: "StringTo.Integer returns error -- float string '3.14'",
		ArrangeInput: args.Map{
			"when":  "given float string",
			"input": "3.14",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "true"},
	},
}

var bytesToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "BytesTo.String returns 'hello' -- valid byte slice",
		ArrangeInput: args.Map{
			"when":  "given valid byte slice",
			"input": "hello",
		},
		ExpectedInput: "hello",
	},
	{
		Title: "BytesTo.String returns empty -- empty byte slice",
		ArrangeInput: args.Map{
			"when":  "given empty byte slice",
			"input": "",
		},
		ExpectedInput: "",
	},
}

var stringToIntegerWithDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegerWithDefault returns 100 -- valid integer '100'",
		ArrangeInput: args.Map{
			"when":       "given valid integer",
			"input":      "100",
			"defaultInt": -1,
		},
		ExpectedInput: args.Map{"value": "100", "isSuccess": "true"},
	},
	{
		Title: "IntegerWithDefault returns default on invalid input",
		ArrangeInput: args.Map{
			"when":       "given non-numeric",
			"input":      "xyz",
			"defaultInt": -1,
		},
		ExpectedInput: args.Map{"value": "-1", "isSuccess": "false"},
	},
	{
		Title: "IntegerWithDefault returns default on empty string",
		ArrangeInput: args.Map{
			"when":       "given empty string",
			"input":      "",
			"defaultInt": 42,
		},
		ExpectedInput: args.Map{"value": "42", "isSuccess": "false"},
	},
}

var stringToFloat64TestCases = []coretestcases.CaseV1{
	{
		Title: "Float64 parses valid float",
		ArrangeInput: args.Map{
			"when":  "given valid float string",
			"input": "3.14",
		},
		ExpectedInput: args.Map{"value": "3.14", "hasError": "false"},
	},
	{
		Title: "Float64 parses integer as float",
		ArrangeInput: args.Map{
			"when":  "given integer string",
			"input": "42",
		},
		ExpectedInput: args.Map{"value": "42", "hasError": "false"},
	},
	{
		Title: "Float64 fails on non-numeric",
		ArrangeInput: args.Map{
			"when":  "given non-numeric string",
			"input": "abc",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "true"},
	},
	{
		Title: "Float64 parses negative float",
		ArrangeInput: args.Map{
			"when":  "given negative float",
			"input": "-2.5",
		},
		ExpectedInput: args.Map{"value": "-2.5", "hasError": "false"},
	},
}

var stringToByteTestCases = []coretestcases.CaseV1{
	{
		Title: "Byte parses valid byte value",
		ArrangeInput: args.Map{
			"when":  "given valid byte string",
			"input": "255",
		},
		ExpectedInput: args.Map{"value": "255", "hasError": "false"},
	},
	{
		Title: "Byte parses zero",
		ArrangeInput: args.Map{
			"when":  "given zero string",
			"input": "0",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "false"},
	},
	{
		Title: "Byte parses one",
		ArrangeInput: args.Map{
			"when":  "given one string",
			"input": "1",
		},
		ExpectedInput: args.Map{"value": "1", "hasError": "false"},
	},
	{
		Title: "Byte fails on empty string",
		ArrangeInput: args.Map{
			"when":  "given empty string",
			"input": "",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "true"},
	},
	{
		Title: "Byte fails on value > 255",
		ArrangeInput: args.Map{
			"when":  "given value exceeding byte range",
			"input": "256",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "true"},
	},
	{
		Title: "Byte fails on negative value",
		ArrangeInput: args.Map{
			"when":  "given negative value",
			"input": "-1",
		},
		ExpectedInput: args.Map{"value": "0", "hasError": "true"},
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
		ExpectedInput: args.Map{
			"count":   3,
			"allTrue": true,
		},
	},
	{
		Title: "Hashset handles duplicates",
		ArrangeInput: args.Map{
			"when":  "given duplicate strings",
			"input": []string{"a", "a", "b"},
		},
		ExpectedInput: args.Map{
			"count":   2,
			"allTrue": true,
		},
	},
	{
		Title: "Hashset handles empty slice",
		ArrangeInput: args.Map{
			"when":  "given empty slice",
			"input": []string{},
		},
		ExpectedInput: args.Map{
			"count": 0,
		},
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
