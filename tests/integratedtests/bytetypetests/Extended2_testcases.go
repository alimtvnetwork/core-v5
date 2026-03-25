package bytetypetests

import (
	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var extIsCompareResultTestCases = []coretestcases.CaseV1{
	{
		Title: "IsCompareResult Equal -- 5 == 5 returns true",
		ArrangeInput: args.Map{
			"when":    "Equal compare 5 to 5",
			"value":   5,
			"n":       5,
			"compare": "Equal",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsCompareResult Equal -- 5 == 3 returns false",
		ArrangeInput: args.Map{
			"when":    "Equal compare 5 to 3",
			"value":   5,
			"n":       3,
			"compare": "Equal",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "IsCompareResult LeftGreater -- 5 > 3 returns true",
		ArrangeInput: args.Map{
			"when":    "LeftGreater compare 5 to 3",
			"value":   5,
			"n":       3,
			"compare": "LeftGreater",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsCompareResult LeftGreater -- 3 > 5 returns false",
		ArrangeInput: args.Map{
			"when":    "LeftGreater compare 3 to 5",
			"value":   3,
			"n":       5,
			"compare": "LeftGreater",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
	{
		Title: "IsCompareResult LeftGreaterEqual -- 5 >= 5 returns true",
		ArrangeInput: args.Map{
			"when":    "LeftGreaterEqual compare 5 to 5",
			"value":   5,
			"n":       5,
			"compare": "LeftGreaterEqual",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsCompareResult LeftLess -- 3 < 5 returns true",
		ArrangeInput: args.Map{
			"when":    "LeftLess compare 3 to 5",
			"value":   3,
			"n":       5,
			"compare": "LeftLess",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsCompareResult LeftLessEqual -- 5 <= 5 returns true",
		ArrangeInput: args.Map{
			"when":    "LeftLessEqual compare 5 to 5",
			"value":   5,
			"n":       5,
			"compare": "LeftLessEqual",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsCompareResult NotEqual -- 5 != 3 returns true",
		ArrangeInput: args.Map{
			"when":    "NotEqual compare 5 to 3",
			"value":   5,
			"n":       3,
			"compare": "NotEqual",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
}

var extEnumMethodsTestCases = []coretestcases.CaseV1{
	{
		Title: "Variant enum methods -- One",
		ArrangeInput: args.Map{
			"when":  "Variant One",
			"value": 1,
		},
		ExpectedInput: args.Map{
			"name":           bytetype.One.Name(),
			"nameValue":      bytetype.One.NameValue(),
			"typeName":       bytetype.One.TypeName(),
			"isValidRange":   true,
			"isInvalidRange": false,
			"stringValue":    "1",
			"rangeNamesCsv":  bytetype.One.RangeNamesCsv(),
		},
	},
	{
		Title: "Variant enum methods -- Zero",
		ArrangeInput: args.Map{
			"when":  "Variant Zero",
			"value": 0,
		},
		ExpectedInput: args.Map{
			"name":           bytetype.Zero.Name(),
			"nameValue":      bytetype.Zero.NameValue(),
			"typeName":       bytetype.Zero.TypeName(),
			"isValidRange":   true,
			"isInvalidRange": false,
			"stringValue":    "0",
			"rangeNamesCsv":  bytetype.Zero.RangeNamesCsv(),
		},
	},
}

var extIsMaxTestCases = []coretestcases.CaseV1{
	{
		Title: "Variant IsMax -- Max returns true",
		ArrangeInput: args.Map{
			"when":  "Variant Max (255)",
			"value": 255,
		},
		ExpectedInput: args.Map{
			"isMax": true,
		},
	},
	{
		Title: "Variant IsMax -- One returns false",
		ArrangeInput: args.Map{
			"when":  "Variant One",
			"value": 1,
		},
		ExpectedInput: args.Map{
			"isMax": false,
		},
	},
}

var extIsEnumEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEnumEqual -- One equals One",
		ArrangeInput: args.Map{
			"when":  "comparing One to One",
			"value": 1,
			"other": 1,
		},
		ExpectedInput: args.Map{
			"isEnumEqual": true,
		},
	},
	{
		Title: "IsEnumEqual -- One not equals Two",
		ArrangeInput: args.Map{
			"when":  "comparing One to Two",
			"value": 1,
			"other": 2,
		},
		ExpectedInput: args.Map{
			"isEnumEqual": false,
		},
	},
}
