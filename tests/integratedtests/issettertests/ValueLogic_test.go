package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

func Test_Value_IsOnLogically(t *testing.T) {
	for caseIndex, tc := range isOnLogicallyTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		actual := args.Map{"result": value.IsOnLogically()}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsOffLogically(t *testing.T) {
	for caseIndex, tc := range isOffLogicallyTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		actual := args.Map{"result": value.IsOffLogically()}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_WildcardApply(t *testing.T) {
	for caseIndex, tc := range wildcardApplyTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		boolInput := input["input"].(bool)

		actual := args.Map{"result": value.WildcardApply(boolInput)}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsWildcardOrBool(t *testing.T) {
	for caseIndex, tc := range isWildcardOrBoolTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		boolInput := input["input"].(bool)

		actual := args.Map{"result": value.IsWildcardOrBool(boolInput)}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_ToByteCondition(t *testing.T) {
	for caseIndex, tc := range toByteConditionTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		trueVal := input["trueVal"].(byte)
		falseVal := input["falseVal"].(byte)
		invalidVal := input["invalidVal"].(byte)

		actual := args.Map{"result": int(value.ToByteCondition(trueVal, falseVal, invalidVal))}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_ToByteConditionWithWildcard(t *testing.T) {
	for caseIndex, tc := range toByteConditionWithWildcardTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		wildcardVal := input["wildcardVal"].(byte)
		trueVal := input["trueVal"].(byte)
		falseVal := input["falseVal"].(byte)
		invalidVal := input["invalidVal"].(byte)

		actual := args.Map{"result": int(value.ToByteConditionWithWildcard(wildcardVal, trueVal, falseVal, invalidVal))}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsDefinedLogically(t *testing.T) {
	for caseIndex, tc := range isDefinedLogicallyTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		actual := args.Map{"result": value.IsDefinedLogically()}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsUndefinedLogically(t *testing.T) {
	for caseIndex, tc := range isUndefinedLogicallyTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		actual := args.Map{"result": value.IsUndefinedLogically()}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsPositive(t *testing.T) {
	for caseIndex, tc := range isPositiveTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		actual := args.Map{"result": value.IsPositive()}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_IsNegative(t *testing.T) {
	for caseIndex, tc := range isNegativeTestCases {
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		actual := args.Map{"result": value.IsNegative()}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_GetSetBoolOnInvalid(t *testing.T) {
	for caseIndex, tc := range getSetBoolTestCases {
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		setter := input["setter"].(bool)

		result := v.GetSetBoolOnInvalid(setter)

		actual := args.Map{
			"result":        result,
			"isTrueOrFalse": v.IsTrue() || v.IsFalse(),
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_LazyEvaluateBool(t *testing.T) {
	for caseIndex, tc := range lazyEvaluateBoolTestCases {
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		called := false

		result := v.LazyEvaluateBool(func() { called = true })

		actual := args.Map{
			"called":       called,
			"returnedTrue": result,
			"isTrue":       v.IsTrue(),
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Value_LazyEvaluateSet(t *testing.T) {
	for caseIndex, tc := range lazyEvaluateSetTestCases {
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		called := false

		result := v.LazyEvaluateSet(func() { called = true })

		actual := args.Map{
			"called":       called,
			"returnedTrue": result,
			"isSet":        v.IsSet(),
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
