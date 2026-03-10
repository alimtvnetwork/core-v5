package issettertests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/issetter"
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
		ExpectedInput: args.Map{"hasError": false, "name": "True"},
	},
	{
		Title:         "New with 'false' returns False",
		ArrangeInput:  "false",
		ExpectedInput: args.Map{"hasError": false, "name": "False"},
	},
	{
		Title:         "New with 'yes' returns True",
		ArrangeInput:  "yes",
		ExpectedInput: args.Map{"hasError": false, "name": "True"},
	},
	{
		Title:         "New with 'no' returns False",
		ArrangeInput:  "no",
		ExpectedInput: args.Map{"hasError": false, "name": "False"},
	},
	{
		Title:         "New with 'Set' returns Set",
		ArrangeInput:  "Set",
		ExpectedInput: args.Map{"hasError": false, "name": "Set"},
	},
	{
		Title:         "New with 'Unset' returns Unset",
		ArrangeInput:  "Unset",
		ExpectedInput: args.Map{"hasError": false, "name": "Unset"},
	},
	{
		Title:         "New with '*' returns Wildcard",
		ArrangeInput:  "*",
		ExpectedInput: args.Map{"hasError": false, "name": "Wildcard"},
	},
	{
		Title:         "New with empty string returns Uninitialized",
		ArrangeInput:  "",
		ExpectedInput: args.Map{"hasError": false, "name": "Uninitialized"},
	},
	{
		Title:         "New with invalid string returns error",
		ArrangeInput:  "invalid_value_xyz",
		ExpectedInput: args.Map{"hasError": true, "name": "Uninitialized"},
	},
}

var getBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetBool(true) returns True",
		ArrangeInput:  true,
		ExpectedInput: "True",
	},
	{
		Title:         "GetBool(false) returns False",
		ArrangeInput:  false,
		ExpectedInput: "False",
	},
}

var newBoolTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewBool(true) returns True",
		ArrangeInput:  true,
		ExpectedInput: "True",
	},
	{
		Title:         "NewBool(false) returns False",
		ArrangeInput:  false,
		ExpectedInput: "False",
	},
}

// booleanLogicTestCases
// Expected keys: isOn, isOff, isTrue, isFalse, isSet, isUnset, isValid, isWildcard
var booleanLogicTestCases = []coretestcases.CaseV1{
	{
		Title:        "Uninitialized boolean logic",
		ArrangeInput: issetter.Uninitialized,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": false, "isTrue": false, "isFalse": false,
			"isSet": false, "isUnset": false, "isValid": false, "isWildcard": false,
		},
	},
	{
		Title:        "True boolean logic",
		ArrangeInput: issetter.True,
		ExpectedInput: args.Map{
			"isOn": true, "isOff": false, "isTrue": true, "isFalse": false,
			"isSet": false, "isUnset": false, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "False boolean logic",
		ArrangeInput: issetter.False,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": true, "isTrue": false, "isFalse": true,
			"isSet": false, "isUnset": false, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "Set boolean logic",
		ArrangeInput: issetter.Set,
		ExpectedInput: args.Map{
			"isOn": true, "isOff": false, "isTrue": false, "isFalse": false,
			"isSet": true, "isUnset": false, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "Unset boolean logic",
		ArrangeInput: issetter.Unset,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": true, "isTrue": false, "isFalse": false,
			"isSet": false, "isUnset": true, "isValid": true, "isWildcard": false,
		},
	},
	{
		Title:        "Wildcard boolean logic",
		ArrangeInput: issetter.Wildcard,
		ExpectedInput: args.Map{
			"isOn": false, "isOff": false, "isTrue": false, "isFalse": false,
			"isSet": false, "isUnset": false, "isValid": true, "isWildcard": true,
		},
	},
}

var combinedBooleansTestCases = []coretestcases.CaseV1{
	{
		Title:         "All true returns True",
		ArrangeInput:  []bool{true, true, true},
		ExpectedInput: "True",
	},
	{
		Title:         "Any false returns False",
		ArrangeInput:  []bool{true, false, true},
		ExpectedInput: "False",
	},
	{
		Title:         "Empty returns True",
		ArrangeInput:  []bool{},
		ExpectedInput: "True",
	},
	{
		Title:         "Single true returns True",
		ArrangeInput:  []bool{true},
		ExpectedInput: "True",
	},
	{
		Title:         "Single false returns False",
		ArrangeInput:  []bool{false},
		ExpectedInput: "False",
	},
}

// conversionTestCases
// Expected keys: toBooleanValue, toSetUnsetValue
var conversionTestCases = []coretestcases.CaseV1{
	{
		Title:         "True converts to True/Set",
		ArrangeInput:  issetter.True,
		ExpectedInput: args.Map{"toBooleanValue": "True", "toSetUnsetValue": "Set"},
	},
	{
		Title:         "False converts to False/Unset",
		ArrangeInput:  issetter.False,
		ExpectedInput: args.Map{"toBooleanValue": "False", "toSetUnsetValue": "Unset"},
	},
	{
		Title:         "Set converts to True/Set",
		ArrangeInput:  issetter.Set,
		ExpectedInput: args.Map{"toBooleanValue": "True", "toSetUnsetValue": "Set"},
	},
	{
		Title:         "Unset converts to False/Unset",
		ArrangeInput:  issetter.Unset,
		ExpectedInput: args.Map{"toBooleanValue": "False", "toSetUnsetValue": "Unset"},
	},
	{
		Title:         "Wildcard converts to Wildcard/Wildcard",
		ArrangeInput:  issetter.Wildcard,
		ExpectedInput: args.Map{"toBooleanValue": "Wildcard", "toSetUnsetValue": "Wildcard"},
	},
	{
		Title:         "Uninitialized converts to Uninitialized/Uninitialized",
		ArrangeInput:  issetter.Uninitialized,
		ExpectedInput: args.Map{"toBooleanValue": "Uninitialized", "toSetUnsetValue": "Uninitialized"},
	},
}

var getSetTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetSet true returns trueValue",
		ArrangeInput:  getSetInput{condition: true, trueVal: issetter.True, falseVal: issetter.False},
		ExpectedInput: "True",
	},
	{
		Title:         "GetSet false returns falseValue",
		ArrangeInput:  getSetInput{condition: false, trueVal: issetter.True, falseVal: issetter.False},
		ExpectedInput: "False",
	},
	{
		Title:         "GetSet true with Set/Unset returns Set",
		ArrangeInput:  getSetInput{condition: true, trueVal: issetter.Set, falseVal: issetter.Unset},
		ExpectedInput: "Set",
	},
}

var isOutOfRangeTestCases = []coretestcases.CaseV1{
	{
		Title:         "Value 0 (Uninitialized) is in range",
		ArrangeInput:  byte(0),
		ExpectedInput: "false",
	},
	{
		Title:         "Value 5 (Wildcard/max) is in range",
		ArrangeInput:  byte(5),
		ExpectedInput: "false",
	},
	{
		Title:         "Value 6 is out of range",
		ArrangeInput:  byte(6),
		ExpectedInput: "true",
	},
	{
		Title:         "Value 255 is out of range",
		ArrangeInput:  byte(255),
		ExpectedInput: "true",
	},
}
