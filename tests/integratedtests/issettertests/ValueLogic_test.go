package issettertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/issetter"
)

func Test_Value_IsOnLogically(t *testing.T) {
	for caseIndex, tc := range isOnLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.IsOnLogically())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_IsOffLogically(t *testing.T) {
	for caseIndex, tc := range isOffLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.IsOffLogically())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_WildcardApply(t *testing.T) {
	for caseIndex, tc := range wildcardApplyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		boolInput := input["input"].(bool)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.WildcardApply(boolInput))}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_IsWildcardOrBool(t *testing.T) {
	for caseIndex, tc := range isWildcardOrBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		boolInput := input["input"].(bool)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.IsWildcardOrBool(boolInput))}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_ToByteCondition(t *testing.T) {
	for caseIndex, tc := range toByteConditionTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		trueVal := input["trueVal"].(byte)
		falseVal := input["falseVal"].(byte)
		invalidVal := input["invalidVal"].(byte)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.ToByteCondition(trueVal, falseVal, invalidVal))}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_ToByteConditionWithWildcard(t *testing.T) {
	for caseIndex, tc := range toByteConditionWithWildcardTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)
		wildcardVal := input["wildcardVal"].(byte)
		trueVal := input["trueVal"].(byte)
		falseVal := input["falseVal"].(byte)
		invalidVal := input["invalidVal"].(byte)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.ToByteConditionWithWildcard(wildcardVal, trueVal, falseVal, invalidVal))}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_IsDefinedLogically(t *testing.T) {
	for caseIndex, tc := range isDefinedLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.IsDefinedLogically())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_IsUndefinedLogically(t *testing.T) {
	for caseIndex, tc := range isUndefinedLogicallyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.IsUndefinedLogically())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_IsPositive(t *testing.T) {
	for caseIndex, tc := range isPositiveTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.IsPositive())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_IsNegative(t *testing.T) {
	for caseIndex, tc := range isNegativeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value := input["value"].(issetter.Value)

		// Act
		actLines := []string{fmt.Sprintf("%v", value.IsNegative())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_GetSetBoolOnInvalid(t *testing.T) {
	for caseIndex, tc := range getSetBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		setter := input["setter"].(bool)

		// Act
		result := v.GetSetBoolOnInvalid(setter)

		actLines := []string{
			fmt.Sprintf("%v", result),
			fmt.Sprintf("%v", v.IsTrue() || v.IsFalse()),
		}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_LazyEvaluateBool(t *testing.T) {
	for caseIndex, tc := range lazyEvaluateBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		called := false

		// Act
		result := v.LazyEvaluateBool(func() { called = true })

		actLines := []string{
			fmt.Sprintf("%v", called),
			fmt.Sprintf("%v", result),
		}

		if called {
			actLines = append(actLines, fmt.Sprintf("%v", v.IsTrue()))
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_Value_LazyEvaluateSet(t *testing.T) {
	for caseIndex, tc := range lazyEvaluateSetTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		v := input["initial"].(issetter.Value)
		called := false

		// Act
		result := v.LazyEvaluateSet(func() { called = true })

		actLines := []string{
			fmt.Sprintf("%v", called),
			fmt.Sprintf("%v", result),
		}

		if called {
			actLines = append(actLines, fmt.Sprintf("%v", v.IsSet()))
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}
