package bytetypetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var getSetTestCases = []coretestcases.CaseV1{
	{
		Title: "GetSet true condition -- returns trueValue",
		ArrangeInput: args.Map{
			"when":       "condition is true",
			"condition":  true,
			"trueValue":  5,
			"falseValue": 10,
		},
		ExpectedInput: args.Map{
			"result": 5,
		},
	},
	{
		Title: "GetSet false condition -- returns falseValue",
		ArrangeInput: args.Map{
			"when":       "condition is false",
			"condition":  false,
			"trueValue":  5,
			"falseValue": 10,
		},
		ExpectedInput: args.Map{
			"result": 10,
		},
	},
}

var getSetVariantTestCases = []coretestcases.CaseV1{
	{
		Title: "GetSetVariant true condition -- returns trueValue as Variant",
		ArrangeInput: args.Map{
			"when":       "condition is true",
			"condition":  true,
			"trueValue":  3,
			"falseValue": 7,
		},
		ExpectedInput: args.Map{
			"result": 3,
		},
	},
	{
		Title: "GetSetVariant false condition -- returns falseValue as Variant",
		ArrangeInput: args.Map{
			"when":       "condition is false",
			"condition":  false,
			"trueValue":  3,
			"falseValue": 7,
		},
		ExpectedInput: args.Map{
			"result": 7,
		},
	},
}

var comparisonTestCases = []coretestcases.CaseV1{
	{
		Title: "Variant comparisons -- 5 vs various values",
		ArrangeInput: args.Map{
			"when":  "variant value 5",
			"value": 5,
		},
		ExpectedInput: args.Map{
			"isEqual3":        false,
			"isEqual5":        true,
			"isGreater3":      true,
			"isGreater7":      false,
			"isGreaterEqual5": true,
			"isLess3":         false,
			"isLess7":         true,
			"isLessEqual5":    true,
			"isBetween3and7":  true,
			"isBetween6and8":  false,
		},
	},
}

var stringConversionTestCases = []coretestcases.CaseV1{
	{
		Title: "String from empty bytes -- returns empty",
		ArrangeInput: args.Map{
			"when":  "empty byte slice",
			"input": "",
		},
		ExpectedInput: args.Map{
			"result": "",
		},
	},
	{
		Title: "String from bytes -- returns string",
		ArrangeInput: args.Map{
			"when":  "non-empty byte slice",
			"input": "hello",
		},
		ExpectedInput: args.Map{
			"result": "hello",
		},
	},
}

var variantMethodsTestCases = []coretestcases.CaseV1{
	{
		Title: "Variant methods -- One value checks",
		ArrangeInput: args.Map{
			"when":  "Variant One",
			"value": 1,
		},
		ExpectedInput: args.Map{
			"isZero":    false,
			"isOne":     true,
			"isTwo":     false,
			"isThree":  false,
			"isMin":     false,
			"isValid":   true,
			"isInvalid": false,
			"valueInt":  1,
			"valueByte": 1,
		},
	},
	{
		Title: "Variant methods -- Zero value checks",
		ArrangeInput: args.Map{
			"when":  "Variant Zero",
			"value": 0,
		},
		ExpectedInput: args.Map{
			"isZero":    true,
			"isOne":     false,
			"isTwo":     false,
			"isThree":  false,
			"isMin":     true,
			"isValid":   false,
			"isInvalid": true,
			"valueInt":  0,
			"valueByte": 0,
		},
	},
}

var variantArithmeticTestCases = []coretestcases.CaseV1{
	{
		Title: "Variant Add -- 3 + 2 = 5",
		ArrangeInput: args.Map{
			"when": "add 2 to 3",
			"base": 3,
			"n":    2,
		},
		ExpectedInput: args.Map{
			"addResult":      5,
			"subtractResult": 1,
		},
	},
}
