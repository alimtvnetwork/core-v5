package issettertests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/issetter"
)

// =============================================================================
// IsOnLogically test cases
// =============================================================================

var isOnLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsOnLogically - Uninitialized returns false",
		ArrangeInput: args.Map{
			"value": issetter.Uninitialized,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsOnLogically - True returns true",
		ArrangeInput: args.Map{
			"value": issetter.True,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsOnLogically - False returns false",
		ArrangeInput: args.Map{
			"value": issetter.False,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsOnLogically - Unset returns false",
		ArrangeInput: args.Map{
			"value": issetter.Unset,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsOnLogically - Set returns true",
		ArrangeInput: args.Map{
			"value": issetter.Set,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsOnLogically - Wildcard returns false",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// IsOffLogically test cases
// =============================================================================

var isOffLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsOffLogically - Uninitialized returns false",
		ArrangeInput: args.Map{
			"value": issetter.Uninitialized,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsOffLogically - True returns false",
		ArrangeInput: args.Map{
			"value": issetter.True,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsOffLogically - False returns true",
		ArrangeInput: args.Map{
			"value": issetter.False,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsOffLogically - Unset returns true",
		ArrangeInput: args.Map{
			"value": issetter.Unset,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsOffLogically - Set returns false",
		ArrangeInput: args.Map{
			"value": issetter.Set,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsOffLogically - Wildcard returns false",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// WildcardApply test cases
// =============================================================================

var wildcardApplyTestCases = []coretestcases.CaseV1{
	{
		Title: "WildcardApply - Wildcard passes through true",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
			"input": true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "WildcardApply - Wildcard passes through false",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
			"input": false,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "WildcardApply - Uninitialized passes through true",
		ArrangeInput: args.Map{
			"value": issetter.Uninitialized,
			"input": true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "WildcardApply - Unset passes through false",
		ArrangeInput: args.Map{
			"value": issetter.Unset,
			"input": false,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "WildcardApply - True ignores input returns true",
		ArrangeInput: args.Map{
			"value": issetter.True,
			"input": false,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "WildcardApply - False ignores input returns false",
		ArrangeInput: args.Map{
			"value": issetter.False,
			"input": true,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "WildcardApply - Set ignores input returns false",
		ArrangeInput: args.Map{
			"value": issetter.Set,
			"input": true,
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// IsWildcardOrBool test cases
// =============================================================================

var isWildcardOrBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "IsWildcardOrBool - Wildcard always true",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
			"input": false,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsWildcardOrBool - True with true",
		ArrangeInput: args.Map{
			"value": issetter.True,
			"input": true,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsWildcardOrBool - False with false",
		ArrangeInput: args.Map{
			"value": issetter.False,
			"input": false,
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// ToByteCondition test cases
// =============================================================================

var toByteConditionTestCases = []coretestcases.CaseV1{
	{
		Title: "ToByteCondition - True returns trueVal",
		ArrangeInput: args.Map{
			"value":      issetter.True,
			"trueVal":    byte(10),
			"falseVal":   byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: []string{"10"},
	},
	{
		Title: "ToByteCondition - False returns falseVal",
		ArrangeInput: args.Map{
			"value":      issetter.False,
			"trueVal":    byte(10),
			"falseVal":   byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: []string{"20"},
	},
	{
		Title: "ToByteCondition - Uninitialized returns invalid",
		ArrangeInput: args.Map{
			"value":      issetter.Uninitialized,
			"trueVal":    byte(10),
			"falseVal":   byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: []string{"255"},
	},
	{
		Title: "ToByteCondition - Set returns invalid",
		ArrangeInput: args.Map{
			"value":      issetter.Set,
			"trueVal":    byte(10),
			"falseVal":   byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: []string{"255"},
	},
	{
		Title: "ToByteCondition - Wildcard returns invalid",
		ArrangeInput: args.Map{
			"value":      issetter.Wildcard,
			"trueVal":    byte(10),
			"falseVal":   byte(20),
			"invalidVal": byte(255),
		},
		ExpectedInput: []string{"255"},
	},
}

// =============================================================================
// ToByteConditionWithWildcard test cases
// =============================================================================

var toByteConditionWithWildcardTestCases = []coretestcases.CaseV1{
	{
		Title: "ToByteConditionWithWildcard - Wildcard returns wildcard byte",
		ArrangeInput: args.Map{
			"value":       issetter.Wildcard,
			"wildcardVal": byte(99),
			"trueVal":     byte(10),
			"falseVal":    byte(20),
			"invalidVal":  byte(255),
		},
		ExpectedInput: []string{"99"},
	},
	{
		Title: "ToByteConditionWithWildcard - True returns trueVal",
		ArrangeInput: args.Map{
			"value":       issetter.True,
			"wildcardVal": byte(99),
			"trueVal":     byte(10),
			"falseVal":    byte(20),
			"invalidVal":  byte(255),
		},
		ExpectedInput: []string{"10"},
	},
	{
		Title: "ToByteConditionWithWildcard - False returns falseVal",
		ArrangeInput: args.Map{
			"value":       issetter.False,
			"wildcardVal": byte(99),
			"trueVal":     byte(10),
			"falseVal":    byte(20),
			"invalidVal":  byte(255),
		},
		ExpectedInput: []string{"20"},
	},
	{
		Title: "ToByteConditionWithWildcard - Uninitialized returns invalid",
		ArrangeInput: args.Map{
			"value":       issetter.Uninitialized,
			"wildcardVal": byte(99),
			"trueVal":     byte(10),
			"falseVal":    byte(20),
			"invalidVal":  byte(255),
		},
		ExpectedInput: []string{"255"},
	},
}

// =============================================================================
// IsDefinedLogically test cases
// =============================================================================

var isDefinedLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsDefinedLogically - Uninitialized false",
		ArrangeInput: args.Map{
			"value": issetter.Uninitialized,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsDefinedLogically - True true",
		ArrangeInput: args.Map{
			"value": issetter.True,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsDefinedLogically - False true",
		ArrangeInput: args.Map{
			"value": issetter.False,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsDefinedLogically - Unset true",
		ArrangeInput: args.Map{
			"value": issetter.Unset,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsDefinedLogically - Set true",
		ArrangeInput: args.Map{
			"value": issetter.Set,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsDefinedLogically - Wildcard false",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// IsUndefinedLogically test cases
// =============================================================================

var isUndefinedLogicallyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsUndefinedLogically - Uninitialized true",
		ArrangeInput: args.Map{
			"value": issetter.Uninitialized,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsUndefinedLogically - True false",
		ArrangeInput: args.Map{
			"value": issetter.True,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsUndefinedLogically - False false",
		ArrangeInput: args.Map{
			"value": issetter.False,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsUndefinedLogically - Unset false",
		ArrangeInput: args.Map{
			"value": issetter.Unset,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsUndefinedLogically - Set false",
		ArrangeInput: args.Map{
			"value": issetter.Set,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsUndefinedLogically - Wildcard true",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
		},
		ExpectedInput: []string{"true"},
	},
}

// =============================================================================
// IsPositive test cases
// =============================================================================

var isPositiveTestCases = []coretestcases.CaseV1{
	{
		Title: "IsPositive - Uninitialized false",
		ArrangeInput: args.Map{
			"value": issetter.Uninitialized,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsPositive - True true",
		ArrangeInput: args.Map{
			"value": issetter.True,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsPositive - False false",
		ArrangeInput: args.Map{
			"value": issetter.False,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsPositive - Unset false",
		ArrangeInput: args.Map{
			"value": issetter.Unset,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsPositive - Set true",
		ArrangeInput: args.Map{
			"value": issetter.Set,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsPositive - Wildcard false",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// IsNegative test cases
// =============================================================================

var isNegativeTestCases = []coretestcases.CaseV1{
	{
		Title: "IsNegative - Uninitialized true",
		ArrangeInput: args.Map{
			"value": issetter.Uninitialized,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNegative - True false",
		ArrangeInput: args.Map{
			"value": issetter.True,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsNegative - False true",
		ArrangeInput: args.Map{
			"value": issetter.False,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNegative - Unset true",
		ArrangeInput: args.Map{
			"value": issetter.Unset,
		},
		ExpectedInput: []string{"true"},
	},
	{
		Title: "IsNegative - Set false",
		ArrangeInput: args.Map{
			"value": issetter.Set,
		},
		ExpectedInput: []string{"false"},
	},
	{
		Title: "IsNegative - Wildcard false",
		ArrangeInput: args.Map{
			"value": issetter.Wildcard,
		},
		ExpectedInput: []string{"false"},
	},
}

// =============================================================================
// GetSetBoolOnInvalid test cases
// =============================================================================

var getSetBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "GetSetBoolOnInvalid - already True returns true ignores setter",
		ArrangeInput: args.Map{
			"initial": issetter.True,
			"setter":  false,
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "GetSetBoolOnInvalid - already False returns false ignores setter",
		ArrangeInput: args.Map{
			"initial": issetter.False,
			"setter":  true,
		},
		ExpectedInput: []string{"false", "true"},
	},
	{
		Title: "GetSetBoolOnInvalid - Uninitialized with true sets True",
		ArrangeInput: args.Map{
			"initial": issetter.Uninitialized,
			"setter":  true,
		},
		ExpectedInput: []string{"true", "true"},
	},
	{
		Title: "GetSetBoolOnInvalid - Uninitialized with false sets False",
		ArrangeInput: args.Map{
			"initial": issetter.Uninitialized,
			"setter":  false,
		},
		ExpectedInput: []string{"false", "true"},
	},
	{
		Title: "GetSetBoolOnInvalid - Set triggers setter with true",
		ArrangeInput: args.Map{
			"initial": issetter.Set,
			"setter":  true,
		},
		ExpectedInput: []string{"true", "true"},
	},
}

// =============================================================================
// LazyEvaluateBool test cases
// =============================================================================

var lazyEvaluateBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyEvaluateBool - Uninitialized calls func sets True",
		ArrangeInput: args.Map{
			"initial":      issetter.Uninitialized,
			"expectCalled": true,
			"expectResult": true,
		},
		ExpectedInput: []string{"true", "true", "true"},
	},
	{
		Title: "LazyEvaluateBool - already True skips func",
		ArrangeInput: args.Map{
			"initial":      issetter.True,
			"expectCalled": false,
			"expectResult": false,
		},
		ExpectedInput: []string{"false", "false", "true"},
	},
	{
		Title: "LazyEvaluateBool - already False skips func",
		ArrangeInput: args.Map{
			"initial":      issetter.False,
			"expectCalled": false,
			"expectResult": false,
		},
		ExpectedInput: []string{"false", "false", "false"},
	},
}

// =============================================================================
// LazyEvaluateSet test cases
// =============================================================================

var lazyEvaluateSetTestCases = []coretestcases.CaseV1{
	{
		Title: "LazyEvaluateSet - Uninitialized calls func sets Set",
		ArrangeInput: args.Map{
			"initial":      issetter.Uninitialized,
			"expectCalled": true,
			"expectResult": true,
		},
		ExpectedInput: []string{"true", "true", "true"},
	},
	{
		Title: "LazyEvaluateSet - already Set skips func",
		ArrangeInput: args.Map{
			"initial":      issetter.Set,
			"expectCalled": false,
			"expectResult": false,
		},
		ExpectedInput: []string{"false", "false", "true"},
	},
	{
		Title: "LazyEvaluateSet - already Unset skips func",
		ArrangeInput: args.Map{
			"initial":      issetter.Unset,
			"expectCalled": false,
			"expectResult": false,
		},
		ExpectedInput: []string{"false", "false", "false"},
	},
}
