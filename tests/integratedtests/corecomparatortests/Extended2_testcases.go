package corecomparatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var compareLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title:         "Equal vs Equal -- true",
		ArrangeInput:  args.Map{"when": "both equal", "left": 0, "right": 0},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "LeftGreater vs Equal -- false",
		ArrangeInput:  args.Map{"when": "greater vs equal", "left": 1, "right": 0},
		ExpectedInput: args.Map{"result": false},
	},
	{
		Title:         "LeftLess vs LeftLess -- true",
		ArrangeInput:  args.Map{"when": "both left less", "left": 3, "right": 3},
		ExpectedInput: args.Map{"result": true},
	},
}

var compareIsAnyOfTestCases = []coretestcases.CaseV1{
	{
		Title:         "Equal is any of Equal,LeftLess -- true",
		ArrangeInput:  args.Map{"when": "Equal", "value": 0},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "LeftGreater is any of Equal,LeftLess -- false",
		ArrangeInput:  args.Map{"when": "LeftGreater", "value": 1},
		ExpectedInput: args.Map{"result": false},
	},
}

var compareIsAnyOfEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsAnyOf empty -- true",
		ArrangeInput:  args.Map{"when": "empty values"},
		ExpectedInput: args.Map{"result": true},
	},
}

var compareNameValueTestCases = []coretestcases.CaseV1{
	{
		Title:         "NameValue Equal -- not empty",
		ArrangeInput:  args.Map{"when": "Equal", "value": 0},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

var compareCsvStringsTestCases = []coretestcases.CaseV1{
	{
		Title:         "CsvStrings 2 values -- length 2",
		ArrangeInput:  args.Map{"when": "2 values"},
		ExpectedInput: args.Map{"length": 2},
	},
}

var compareCsvStringsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title:         "CsvStrings empty -- length 0",
		ArrangeInput:  args.Map{"when": "no values"},
		ExpectedInput: args.Map{"length": 0},
	},
}

var compareValueConversionsTestCases = []coretestcases.CaseV1{
	{
		Title:        "Value conversions for Equal -- correct values",
		ArrangeInput: args.Map{"when": "Equal (0)", "value": 0},
		ExpectedInput: args.Map{
			"valueByte":        0,
			"valueInt":         0,
			"toNumberString":   "0",
			"numberString":     "0",
			"numberJsonString": "\"0\"",
		},
	},
	{
		Title:        "Value conversions for LeftGreater -- correct values",
		ArrangeInput: args.Map{"when": "LeftGreater (1)", "value": 1},
		ExpectedInput: args.Map{
			"valueByte":        1,
			"valueInt":         1,
			"toNumberString":   "1",
			"numberString":     "1",
			"numberJsonString": "\"1\"",
		},
	},
}

var compareMarshalJsonTestCases = []coretestcases.CaseV1{
	{
		Title:         "MarshalJSON Equal -- no error and not empty",
		ArrangeInput:  args.Map{"when": "Equal"},
		ExpectedInput: args.Map{"hasError": false, "notEmpty": true},
	},
}

var compareOnlySupportedErrTestCases = []coretestcases.CaseV1{
	{
		Title:         "OnlySupportedErr Equal in list -- no error",
		ArrangeInput:  args.Map{"when": "Equal supported", "value": 0},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title:         "OnlySupportedErr LeftGreater not in list -- error",
		ArrangeInput:  args.Map{"when": "LeftGreater unsupported", "value": 1},
		ExpectedInput: args.Map{"hasError": true},
	},
}

var compareOnlySupportedDirectErrTestCases = []coretestcases.CaseV1{
	{
		Title:         "OnlySupportedDirectErr Equal in list -- no error",
		ArrangeInput:  args.Map{"when": "Equal supported", "value": 0},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title:         "OnlySupportedDirectErr LeftGreater not in list -- error",
		ArrangeInput:  args.Map{"when": "LeftGreater unsupported", "value": 1},
		ExpectedInput: args.Map{"hasError": true},
	},
}

var compareOnlySupportedEmptyMsgTestCases = []coretestcases.CaseV1{
	{
		Title:         "OnlySupportedErr empty message unsupported -- error",
		ArrangeInput:  args.Map{"when": "empty message unsupported"},
		ExpectedInput: args.Map{"hasError": true},
	},
}

var minLengthTestCases = []coretestcases.CaseV1{
	{
		Title:         "MinLength left smaller -- returns left",
		ArrangeInput:  args.Map{"when": "3 vs 5", "left": 3, "right": 5},
		ExpectedInput: args.Map{"result": 3},
	},
	{
		Title:         "MinLength right smaller -- returns right",
		ArrangeInput:  args.Map{"when": "7 vs 4", "left": 7, "right": 4},
		ExpectedInput: args.Map{"result": 4},
	},
	{
		Title:         "MinLength equal -- returns either",
		ArrangeInput:  args.Map{"when": "5 vs 5", "left": 5, "right": 5},
		ExpectedInput: args.Map{"result": 5},
	},
}

var compareIsAnyNamesOfTestCases = []coretestcases.CaseV1{
	{
		Title:         "IsAnyNamesOf Equal in list -- true",
		ArrangeInput:  args.Map{"when": "Equal name in list"},
		ExpectedInput: args.Map{"result": true},
	},
}

var compareIsInconclusiveOrNotEqualTestCases = []coretestcases.CaseV1{
	{
		Title:         "Inconclusive -- true",
		ArrangeInput:  args.Map{"when": "Inconclusive", "value": 6},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "NotEqual -- true",
		ArrangeInput:  args.Map{"when": "NotEqual", "value": 5},
		ExpectedInput: args.Map{"result": true},
	},
	{
		Title:         "Equal -- false",
		ArrangeInput:  args.Map{"when": "Equal", "value": 0},
		ExpectedInput: args.Map{"result": false},
	},
}
