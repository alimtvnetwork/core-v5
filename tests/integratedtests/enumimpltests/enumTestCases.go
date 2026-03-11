package enumimpltests

import (
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var enumByteTestCases = []coretestcases.CaseV1{
	{
		Title: "EnumByte returns min 0 and max 10 -- DynamicMap input",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   0,
				"A":         -2,
				"B":         8,
				"C":         5,
				"Something": 10,
			},
		},
		ExpectedInput: args.Map{
			"min": 0,
			"max": 10,
		},
	},
}

var enumInt8TestCases = []coretestcases.CaseV1{
	{
		Title: "Integer8 enum min -2, max 12 -- generates min, max from given map[string]any",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   -2,
				"A":         -2,
				"B":         8,
				"C":         5,
				"Something": 12,
			},
		},
		ExpectedInput: args.Map{
			"min": -2,
			"max": 12,
		},
	},
}

var enumInt16TestCases = []coretestcases.CaseV1{
	{
		Title: "Integer16 enum min -3, max 14 -- generates min, max from given map[string]any",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   -3,
				"A":         -2,
				"B":         -3,
				"C":         5,
				"Something": 14,
			},
		},
		ExpectedInput: args.Map{
			"min": -3,
			"max": 14,
		},
	},
}

var enumInt32TestCases = []coretestcases.CaseV1{
	{
		Title: "Integer32 enum min -4, max 15 -- generates min, max from given map[string]any",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   -4,
				"A":         -2,
				"B":         -3,
				"C":         5,
				"Something": 15,
			},
		},
		ExpectedInput: args.Map{
			"min": -4,
			"max": 15,
		},
	},
}

var enumUInt16TestCases = []coretestcases.CaseV1{
	{
		Title: "UnsignedInteger16 enum min 0, max 20 -- generates min, max from given map[string]any",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":    0,
				"Something2": 15,
				"B":          15,
				"Something":  20,
			},
		},
		ExpectedInput: args.Map{
			"min": 0,
			"max": 20,
		},
	},
}

var enumStringTestCases = []coretestcases.CaseV1{
	{
		Title: "String enum min empty, max Something2 -- string max is lexicographic",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":    0,
				"Something2": 15,
				"B":          15,
				"Something":  20,
			},
		},
		ExpectedInput: args.Map{
			"min": "",
			"max": "Something2",
		},
	},
}
