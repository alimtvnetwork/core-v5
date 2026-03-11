package corecomparatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var baseIsCaseSensitiveTestCases = []coretestcases.CaseV1{
	{
		Title:         "CaseSensitive true -- IsIgnoreCase false",
		ArrangeInput:  args.Map{"isCaseSensitive": true},
		ExpectedInput: args.Map{"isIgnoreCase": false, "cloneMatch": true},
	},
	{
		Title:         "CaseSensitive false -- IsIgnoreCase true",
		ArrangeInput:  args.Map{"isCaseSensitive": false},
		ExpectedInput: args.Map{"isIgnoreCase": true, "cloneMatch": true},
	},
}

var baseIsIgnoreCaseTestCases = []coretestcases.CaseV1{
	{
		Title:         "IgnoreCase true -- IsCaseSensitive false",
		ArrangeInput:  args.Map{"isIgnoreCase": true},
		ExpectedInput: args.Map{"isCaseSensitive": false, "cloneMatch": true},
	},
	{
		Title:         "IgnoreCase false -- IsCaseSensitive true",
		ArrangeInput:  args.Map{"isIgnoreCase": false},
		ExpectedInput: args.Map{"isCaseSensitive": true, "cloneMatch": true},
	},
}

var compareIsMethodTestCases = []coretestcases.CaseV1{
	{
		Title:        "Is -- Equal vs Equal is true",
		ArrangeInput: args.Map{"value": 0, "other": 0},
		ExpectedInput: args.Map{
			"is":                  true,
			"isInvalid":           false,
			"isValueEqual":        true,
			"isLeftGreater":       false,
			"isLeftGreaterEqual":  false,
			"isLeftLessEqual":     false,
			"isLeftLessOrLeOrEq":  true,
			"isDefinedPlus":       true,
			"isNotInconclusive":   true,
			"rangeNamesCsvNotEmpty": true,
			"sqlOpNotEmpty":       true,
			"stringValueNotEmpty": true,
			"valueInt8":           int8(0),
			"valueInt16":          int16(0),
			"valueInt32":          int32(0),
			"valueString":        "0",
			"formatPanic":        true,
		},
	},
	{
		Title:        "Is -- LeftGreater vs Equal is false",
		ArrangeInput: args.Map{"value": 1, "other": 0},
		ExpectedInput: args.Map{
			"is":                  false,
			"isInvalid":           false,
			"isValueEqual":        false,
			"isLeftGreater":       true,
			"isLeftGreaterEqual":  false,
			"isLeftLessEqual":     false,
			"isLeftLessOrLeOrEq":  false,
			"isDefinedPlus":       false,
			"isNotInconclusive":   true,
			"rangeNamesCsvNotEmpty": true,
			"sqlOpNotEmpty":       true,
			"stringValueNotEmpty": true,
			"valueInt8":           int8(1),
			"valueInt16":          int16(1),
			"valueInt32":          int32(1),
			"valueString":        "1",
			"formatPanic":        true,
		},
	},
	{
		Title:        "Is -- Inconclusive check",
		ArrangeInput: args.Map{"value": 6, "other": 6},
		ExpectedInput: args.Map{
			"is":                  true,
			"isInvalid":           true,
			"isValueEqual":        true,
			"isLeftGreater":       false,
			"isLeftGreaterEqual":  false,
			"isLeftLessEqual":     false,
			"isLeftLessOrLeOrEq":  false,
			"isDefinedPlus":       false,
			"isNotInconclusive":   false,
			"rangeNamesCsvNotEmpty": true,
			"sqlOpNotEmpty":       true,
			"stringValueNotEmpty": true,
			"valueInt8":           int8(6),
			"valueInt16":          int16(6),
			"valueInt32":          int32(6),
			"valueString":        "6",
			"formatPanic":        true,
		},
	},
}

var baseIsCaseSensitiveNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "ClonePtr nil returns nil",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var baseIsIgnoreCaseNilTestCases = []coretestcases.CaseV1{
	{
		Title:         "ClonePtr nil returns nil",
		ArrangeInput:  args.Map{"isNil": true},
		ExpectedInput: args.Map{"isNil": true},
	},
}

var compareUnmarshalJsonTestCases = []coretestcases.CaseV1{
	{
		Title:         "Unmarshal valid name -- no error",
		ArrangeInput:  args.Map{"data": "Equal"},
		ExpectedInput: args.Map{"hasError": false},
	},
	{
		Title:         "Unmarshal invalid name -- error",
		ArrangeInput:  args.Map{"data": "InvalidXyz"},
		ExpectedInput: args.Map{"hasError": true},
	},
	{
		Title:         "Unmarshal nil data -- error",
		ArrangeInput:  args.Map{"isNilData": true},
		ExpectedInput: args.Map{"hasError": true},
	},
}
