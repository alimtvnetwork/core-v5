package conditionaltests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var ifStringTestCases = []coretestcases.CaseV1{
	{
		Title: "If true returns trueValue string",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "yes",
	},
	{
		Title: "If false returns falseValue string",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  "yes",
			"falseValue": "no",
		},
		ExpectedInput: "no",
	},
	{
		Title: "If true with empty strings returns empty trueValue",
		ArrangeInput: args.Map{
			"when":       "given true with empty trueValue",
			"isTrue":     true,
			"trueValue":  "",
			"falseValue": "fallback",
		},
		ExpectedInput: "",
	},
	{
		Title: "If false with empty falseValue returns empty",
		ArrangeInput: args.Map{
			"when":       "given false with empty falseValue",
			"isTrue":     false,
			"trueValue":  "something",
			"falseValue": "",
		},
		ExpectedInput: "",
	},
	{
		Title: "If true with same values returns that value",
		ArrangeInput: args.Map{
			"when":       "given true with identical values",
			"isTrue":     true,
			"trueValue":  "same",
			"falseValue": "same",
		},
		ExpectedInput: "same",
	},
}

var ifIntTestCases = []coretestcases.CaseV1{
	{
		Title: "If true returns trueValue int",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "10",
	},
	{
		Title: "If false returns falseValue int",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  10,
			"falseValue": 20,
		},
		ExpectedInput: "20",
	},
	{
		Title: "If true with zero returns zero",
		ArrangeInput: args.Map{
			"when":       "given true with zero trueValue",
			"isTrue":     true,
			"trueValue":  0,
			"falseValue": 99,
		},
		ExpectedInput: "0",
	},
	{
		Title: "If false with negative returns negative",
		ArrangeInput: args.Map{
			"when":       "given false with negative falseValue",
			"isTrue":     false,
			"trueValue":  100,
			"falseValue": -42,
		},
		ExpectedInput: "-42",
	},
}

var nilDefTestCases = []coretestcases.CaseV1{
	{
		Title: "NilDef with nil pointer returns default",
		ArrangeInput: args.Map{
			"when":   "given nil pointer",
			"isNil":  true,
			"defVal": "default",
		},
		ExpectedInput: "default",
	},
	{
		Title: "NilDef with non-nil pointer returns value",
		ArrangeInput: args.Map{
			"when":   "given non-nil pointer",
			"isNil":  false,
			"value":  "actual",
			"defVal": "default",
		},
		ExpectedInput: "actual",
	},
	{
		Title: "NilDef with empty string pointer returns empty",
		ArrangeInput: args.Map{
			"when":   "given non-nil pointer to empty string",
			"isNil":  false,
			"value":  "",
			"defVal": "fallback",
		},
		ExpectedInput: "",
	},
}

var ifFuncStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfFunc true evaluates trueFunc only",
		ArrangeInput: args.Map{
			"when":       "given true condition with funcs",
			"isTrue":     true,
			"trueValue":  "from-true-func",
			"falseValue": "from-false-func",
		},
		ExpectedInput: "from-true-func",
	},
	{
		Title: "IfFunc false evaluates falseFunc only",
		ArrangeInput: args.Map{
			"when":       "given false condition with funcs",
			"isTrue":     false,
			"trueValue":  "from-true-func",
			"falseValue": "from-false-func",
		},
		ExpectedInput: "from-false-func",
	},
}

var ifTrueFuncStringTestCases = []coretestcases.CaseV1{
	{
		Title: "IfTrueFunc true returns trueFunc result",
		ArrangeInput: args.Map{
			"when":      "given true condition",
			"isTrue":    true,
			"trueValue": "computed",
		},
		ExpectedInput: "computed",
	},
	{
		Title: "IfTrueFunc false returns zero value",
		ArrangeInput: args.Map{
			"when":      "given false condition",
			"isTrue":    false,
			"trueValue": "computed",
		},
		ExpectedInput: "",
	},
}

var ifSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "IfSlice true returns trueSlice",
		ArrangeInput: args.Map{
			"when":       "given true condition",
			"isTrue":     true,
			"trueValue":  []string{"a", "b"},
			"falseValue": []string{"x", "y"},
		},
		ExpectedInput: []string{"2", "a"},
	},
	{
		Title: "IfSlice false returns falseSlice",
		ArrangeInput: args.Map{
			"when":       "given false condition",
			"isTrue":     false,
			"trueValue":  []string{"a", "b"},
			"falseValue": []string{"x", "y", "z"},
		},
		ExpectedInput: []string{"3", "x"},
	},
}

var nilCheckTestCases = []coretestcases.CaseV1{
	{
		Title: "NilCheck returns onNil when input is nil",
		ArrangeInput: args.Map{
			"when":     "given nil input",
			"isNil":    true,
			"onNil":    "nil-result",
			"onNonNil": "non-nil-result",
		},
		ExpectedInput: "nil-result",
	},
	{
		Title: "NilCheck returns onNonNil when input is not nil",
		ArrangeInput: args.Map{
			"when":     "given non-nil input",
			"isNil":    false,
			"onNil":    "nil-result",
			"onNonNil": "non-nil-result",
		},
		ExpectedInput: "non-nil-result",
	},
}

var defOnNilTestCases = []coretestcases.CaseV1{
	{
		Title: "DefOnNil returns default when input is nil",
		ArrangeInput: args.Map{
			"when":     "given nil input",
			"isNil":    true,
			"onNonNil": "default-val",
		},
		ExpectedInput: "default-val",
	},
	{
		Title: "DefOnNil returns input when input is not nil",
		ArrangeInput: args.Map{
			"when":     "given non-nil input",
			"isNil":    false,
			"value":    "actual-val",
			"onNonNil": "default-val",
		},
		ExpectedInput: "actual-val",
	},
}
