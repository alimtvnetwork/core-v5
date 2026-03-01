package issettertests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/issetter"
)

type getSetInput struct {
	condition bool
	trueVal   issetter.Value
	falseVal  issetter.Value
}

var valueNewTestCases = []coretestcases.CaseV1{
	{
		Title:         "New with 'true' returns True",
		ArrangeInput:  "true",
		ExpectedInput: []string{"false", "True"},
	},
	{
		Title:         "New with 'false' returns False",
		ArrangeInput:  "false",
		ExpectedInput: []string{"false", "False"},
	},
	{
		Title:         "New with 'yes' returns True",
		ArrangeInput:  "yes",
		ExpectedInput: []string{"false", "True"},
	},
	{
		Title:         "New with 'no' returns False",
		ArrangeInput:  "no",
		ExpectedInput: []string{"false", "False"},
	},
	{
		Title:         "New with 'Set' returns Set",
		ArrangeInput:  "Set",
		ExpectedInput: []string{"false", "Set"},
	},
	{
		Title:         "New with 'Unset' returns Unset",
		ArrangeInput:  "Unset",
		ExpectedInput: []string{"false", "Unset"},
	},
	{
		Title:         "New with '*' returns Wildcard",
		ArrangeInput:  "*",
		ExpectedInput: []string{"false", "Wildcard"},
	},
	{
		Title:         "New with empty string returns Uninitialized",
		ArrangeInput:  "",
		ExpectedInput: []string{"false", "Uninitialized"},
	},
	{
		Title:         "New with invalid string returns error",
		ArrangeInput:  "invalid_value_xyz",
		ExpectedInput: []string{"true", "Uninitialized"},
	},
}

var getBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetBool(true) returns True",
		ArrangeInput:  true,
		ExpectedInput: []string{"True"},
	},
	{
		Title:         "GetBool(false) returns False",
		ArrangeInput:  false,
		ExpectedInput: []string{"False"},
	},
}

var newBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewBool(true) returns True",
		ArrangeInput:  true,
		ExpectedInput: []string{"True"},
	},
	{
		Title:         "NewBool(false) returns False",
		ArrangeInput:  false,
		ExpectedInput: []string{"False"},
	},
}

// booleanLogicTestCases
// Expected order: isOn, isOff, isTrue, isFalse, isSet, isUnset, isValid, isWildcard
var booleanLogicTestCases = []coretestcases.CaseV1{
	{
		Title:        "Uninitialized boolean logic",
		ArrangeInput: issetter.Uninitialized,
		ExpectedInput: []string{
			"false", "false", "false", "false",
			"false", "false", "false", "false",
		},
	},
	{
		Title:        "True boolean logic",
		ArrangeInput: issetter.True,
		ExpectedInput: []string{
			"true", "false", "true", "false",
			"false", "false", "true", "false",
		},
	},
	{
		Title:        "False boolean logic",
		ArrangeInput: issetter.False,
		ExpectedInput: []string{
			"false", "true", "false", "true",
			"false", "false", "true", "false",
		},
	},
	{
		Title:        "Set boolean logic",
		ArrangeInput: issetter.Set,
		ExpectedInput: []string{
			"true", "false", "false", "false",
			"true", "false", "true", "false",
		},
	},
	{
		Title:        "Unset boolean logic",
		ArrangeInput: issetter.Unset,
		ExpectedInput: []string{
			"false", "true", "false", "false",
			"false", "true", "true", "false",
		},
	},
	{
		Title:        "Wildcard boolean logic",
		ArrangeInput: issetter.Wildcard,
		ExpectedInput: []string{
			"false", "false", "false", "false",
			"false", "false", "true", "true",
		},
	},
}

var combinedBooleansTestCases = []coretestcases.CaseV1{
	{
		Title:         "All true returns True",
		ArrangeInput:  []bool{true, true, true},
		ExpectedInput: []string{"True"},
	},
	{
		Title:         "Any false returns False",
		ArrangeInput:  []bool{true, false, true},
		ExpectedInput: []string{"False"},
	},
	{
		Title:         "Empty returns True",
		ArrangeInput:  []bool{},
		ExpectedInput: []string{"True"},
	},
	{
		Title:         "Single true returns True",
		ArrangeInput:  []bool{true},
		ExpectedInput: []string{"True"},
	},
	{
		Title:         "Single false returns False",
		ArrangeInput:  []bool{false},
		ExpectedInput: []string{"False"},
	},
}

// conversionTestCases
// Expected order: toBooleanValue.Name(), toSetUnsetValue.Name()
var conversionTestCases = []coretestcases.CaseV1{
	{
		Title:         "True converts to True/Set",
		ArrangeInput:  issetter.True,
		ExpectedInput: []string{"True", "Set"},
	},
	{
		Title:         "False converts to False/Unset",
		ArrangeInput:  issetter.False,
		ExpectedInput: []string{"False", "Unset"},
	},
	{
		Title:         "Set converts to True/Set",
		ArrangeInput:  issetter.Set,
		ExpectedInput: []string{"True", "Set"},
	},
	{
		Title:         "Unset converts to False/Unset",
		ArrangeInput:  issetter.Unset,
		ExpectedInput: []string{"False", "Unset"},
	},
	{
		Title:         "Wildcard converts to Wildcard/Wildcard",
		ArrangeInput:  issetter.Wildcard,
		ExpectedInput: []string{"Wildcard", "Wildcard"},
	},
	{
		Title:         "Uninitialized converts to Uninitialized/Uninitialized",
		ArrangeInput:  issetter.Uninitialized,
		ExpectedInput: []string{"Uninitialized", "Uninitialized"},
	},
}

var getSetTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetSet true returns trueValue",
		ArrangeInput:  getSetInput{condition: true, trueVal: issetter.True, falseVal: issetter.False},
		ExpectedInput: []string{"True"},
	},
	{
		Title:         "GetSet false returns falseValue",
		ArrangeInput:  getSetInput{condition: false, trueVal: issetter.True, falseVal: issetter.False},
		ExpectedInput: []string{"False"},
	},
	{
		Title:         "GetSet true with Set/Unset returns Set",
		ArrangeInput:  getSetInput{condition: true, trueVal: issetter.Set, falseVal: issetter.Unset},
		ExpectedInput: []string{"Set"},
	},
}

var isOutOfRangeTestCases = []coretestcases.CaseV1{
	{
		Title:         "Value 0 (Uninitialized) is in range",
		ArrangeInput:  byte(0),
		ExpectedInput: []string{"false"},
	},
	{
		Title:         "Value 5 (Wildcard/max) is in range",
		ArrangeInput:  byte(5),
		ExpectedInput: []string{"false"},
	},
	{
		Title:         "Value 6 is out of range",
		ArrangeInput:  byte(6),
		ExpectedInput: []string{"true"},
	},
	{
		Title:         "Value 255 is out of range",
		ArrangeInput:  byte(255),
		ExpectedInput: []string{"true"},
	},
}
