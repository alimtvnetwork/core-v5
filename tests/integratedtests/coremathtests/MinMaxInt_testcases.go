package coremathtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var maxIntTestCases = []coretestcases.CaseV1{
	{
		Title: "MaxInt returns larger of two positives",
		ArrangeInput: args.Map{
			"when": "given 3 and 7",
			"a":    3,
			"b":    7,
		},
		ExpectedInput: args.Map{"result": 7},
	},
	{
		Title: "MaxInt returns equal when same",
		ArrangeInput: args.Map{
			"when": "given 5 and 5",
			"a":    5,
			"b":    5,
		},
		ExpectedInput: args.Map{"result": 5},
	},
	{
		Title: "MaxInt handles negatives",
		ArrangeInput: args.Map{
			"when": "given -3 and -7",
			"a":    -3,
			"b":    -7,
		},
		ExpectedInput: args.Map{"result": -3},
	},
	{
		Title: "MaxInt with zero and positive",
		ArrangeInput: args.Map{
			"when": "given 0 and 10",
			"a":    0,
			"b":    10,
		},
		ExpectedInput: args.Map{"result": 10},
	},
	{
		Title: "MaxInt with zero and negative",
		ArrangeInput: args.Map{
			"when": "given 0 and -5",
			"a":    0,
			"b":    -5,
		},
		ExpectedInput: args.Map{"result": 0},
	},
}

var minIntTestCases = []coretestcases.CaseV1{
	{
		Title: "MinInt returns smaller of two positives",
		ArrangeInput: args.Map{
			"when": "given 3 and 7",
			"a":    3,
			"b":    7,
		},
		ExpectedInput: args.Map{"result": 3},
	},
	{
		Title: "MinInt returns equal when same",
		ArrangeInput: args.Map{
			"when": "given 5 and 5",
			"a":    5,
			"b":    5,
		},
		ExpectedInput: args.Map{"result": 5},
	},
	{
		Title: "MinInt with zero and negative",
		ArrangeInput: args.Map{
			"when": "given 0 and -3",
			"a":    0,
			"b":    -3,
		},
		ExpectedInput: args.Map{"result": -3},
	},
}

var maxByteTestCases = []coretestcases.CaseV1{
	{
		Title: "MaxByte returns larger byte",
		ArrangeInput: args.Map{
			"when": "given 10 and 200",
			"a":    10,
			"b":    200,
		},
		ExpectedInput: args.Map{"result": 200},
	},
	{
		Title: "MaxByte returns equal when same",
		ArrangeInput: args.Map{
			"when": "given 128 and 128",
			"a":    128,
			"b":    128,
		},
		ExpectedInput: args.Map{"result": 128},
	},
	{
		Title: "MaxByte with zero",
		ArrangeInput: args.Map{
			"when": "given 0 and 255",
			"a":    0,
			"b":    255,
		},
		ExpectedInput: args.Map{"result": 255},
	},
}

var minByteTestCases = []coretestcases.CaseV1{
	{
		Title: "MinByte returns smaller byte",
		ArrangeInput: args.Map{
			"when": "given 10 and 200",
			"a":    10,
			"b":    200,
		},
		ExpectedInput: args.Map{"result": 10},
	},
	{
		Title: "MinByte with zero",
		ArrangeInput: args.Map{
			"when": "given 0 and 100",
			"a":    0,
			"b":    100,
		},
		ExpectedInput: args.Map{"result": 0},
	},
}

var integerWithinToByteTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegerWithin.ToByte true for 0",
		ArrangeInput: args.Map{
			"when":  "given 0",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerWithin.ToByte true for 255",
		ArrangeInput: args.Map{
			"when":  "given 255",
			"value": 255,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerWithin.ToByte false for 256",
		ArrangeInput: args.Map{
			"when":  "given 256",
			"value": 256,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title: "IntegerWithin.ToByte false for -1",
		ArrangeInput: args.Map{
			"when":  "given -1",
			"value": -1,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var integerWithinToInt8TestCases = []coretestcases.CaseV1{
	{
		Title: "IntegerWithin.ToInt8 true for 0",
		ArrangeInput: args.Map{
			"when":  "given 0",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerWithin.ToInt8 true for 127",
		ArrangeInput: args.Map{
			"when":  "given 127",
			"value": 127,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerWithin.ToInt8 true for -128",
		ArrangeInput: args.Map{
			"when":  "given -128",
			"value": -128,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerWithin.ToInt8 false for 128",
		ArrangeInput: args.Map{
			"when":  "given 128",
			"value": 128,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title: "IntegerWithin.ToInt8 false for -129",
		ArrangeInput: args.Map{
			"when":  "given -129",
			"value": -129,
		},
		ExpectedInput: args.Map{"result": false},
	},
}

var integerOutOfRangeToByteTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegerOutOfRange.ToByte false for 0 (in range)",
		ArrangeInput: args.Map{
			"when":  "given 0",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title: "IntegerOutOfRange.ToByte true for 256 (out of range)",
		ArrangeInput: args.Map{
			"when":  "given 256",
			"value": 256,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerOutOfRange.ToByte true for -1 (out of range)",
		ArrangeInput: args.Map{
			"when":  "given -1",
			"value": -1,
		},
		ExpectedInput: args.Map{"result": true},
	},
}

var integerWithinToInt16TestCases = []coretestcases.CaseV1{
	{
		Title: "IntegerWithin.ToInt16 true for 0",
		ArrangeInput: args.Map{
			"when":  "given 0",
			"value": 0,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerWithin.ToInt16 true for 32767",
		ArrangeInput: args.Map{
			"when":  "given max int16",
			"value": 32767,
		},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title: "IntegerWithin.ToInt16 false for 32768",
		ArrangeInput: args.Map{
			"when":  "given max int16 + 1",
			"value": 32768,
		},
		ExpectedInput: args.Map{"result": false},
	},
}
