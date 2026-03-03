package issettertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/issetter"
)

// =============================================================================
// Test Cases
// =============================================================================

var valueLogicTestCases = []coretestcases.CaseV1{
	// IsOnLogically
	{Title: "IsOnLogically - Uninitialized returns false", ArrangeInput: args.Map{"method": "IsOnLogically", "value": issetter.Uninitialized}, ExpectedInput: []string{"false"}},
	{Title: "IsOnLogically - True returns true", ArrangeInput: args.Map{"method": "IsOnLogically", "value": issetter.True}, ExpectedInput: []string{"true"}},
	{Title: "IsOnLogically - False returns false", ArrangeInput: args.Map{"method": "IsOnLogically", "value": issetter.False}, ExpectedInput: []string{"false"}},
	{Title: "IsOnLogically - Unset returns false", ArrangeInput: args.Map{"method": "IsOnLogically", "value": issetter.Unset}, ExpectedInput: []string{"false"}},
	{Title: "IsOnLogically - Set returns true", ArrangeInput: args.Map{"method": "IsOnLogically", "value": issetter.Set}, ExpectedInput: []string{"true"}},
	{Title: "IsOnLogically - Wildcard returns false", ArrangeInput: args.Map{"method": "IsOnLogically", "value": issetter.Wildcard}, ExpectedInput: []string{"false"}},
	// IsOffLogically
	{Title: "IsOffLogically - Uninitialized returns false", ArrangeInput: args.Map{"method": "IsOffLogically", "value": issetter.Uninitialized}, ExpectedInput: []string{"false"}},
	{Title: "IsOffLogically - True returns false", ArrangeInput: args.Map{"method": "IsOffLogically", "value": issetter.True}, ExpectedInput: []string{"false"}},
	{Title: "IsOffLogically - False returns true", ArrangeInput: args.Map{"method": "IsOffLogically", "value": issetter.False}, ExpectedInput: []string{"true"}},
	{Title: "IsOffLogically - Unset returns true", ArrangeInput: args.Map{"method": "IsOffLogically", "value": issetter.Unset}, ExpectedInput: []string{"true"}},
	{Title: "IsOffLogically - Set returns false", ArrangeInput: args.Map{"method": "IsOffLogically", "value": issetter.Set}, ExpectedInput: []string{"false"}},
	{Title: "IsOffLogically - Wildcard returns false", ArrangeInput: args.Map{"method": "IsOffLogically", "value": issetter.Wildcard}, ExpectedInput: []string{"false"}},
	// WildcardApply
	{Title: "WildcardApply - Wildcard passes through true", ArrangeInput: args.Map{"method": "WildcardApply", "value": issetter.Wildcard, "input": true}, ExpectedInput: []string{"true"}},
	{Title: "WildcardApply - Wildcard passes through false", ArrangeInput: args.Map{"method": "WildcardApply", "value": issetter.Wildcard, "input": false}, ExpectedInput: []string{"false"}},
	{Title: "WildcardApply - Uninitialized passes through true", ArrangeInput: args.Map{"method": "WildcardApply", "value": issetter.Uninitialized, "input": true}, ExpectedInput: []string{"true"}},
	{Title: "WildcardApply - Unset passes through false", ArrangeInput: args.Map{"method": "WildcardApply", "value": issetter.Unset, "input": false}, ExpectedInput: []string{"false"}},
	{Title: "WildcardApply - True ignores input returns true", ArrangeInput: args.Map{"method": "WildcardApply", "value": issetter.True, "input": false}, ExpectedInput: []string{"true"}},
	{Title: "WildcardApply - False ignores input returns false", ArrangeInput: args.Map{"method": "WildcardApply", "value": issetter.False, "input": true}, ExpectedInput: []string{"false"}},
	{Title: "WildcardApply - Set ignores input returns false", ArrangeInput: args.Map{"method": "WildcardApply", "value": issetter.Set, "input": true}, ExpectedInput: []string{"false"}},
	// IsWildcardOrBool
	{Title: "IsWildcardOrBool - Wildcard always true", ArrangeInput: args.Map{"method": "IsWildcardOrBool", "value": issetter.Wildcard, "input": false}, ExpectedInput: []string{"true"}},
	{Title: "IsWildcardOrBool - True with true", ArrangeInput: args.Map{"method": "IsWildcardOrBool", "value": issetter.True, "input": true}, ExpectedInput: []string{"true"}},
	{Title: "IsWildcardOrBool - False with false", ArrangeInput: args.Map{"method": "IsWildcardOrBool", "value": issetter.False, "input": false}, ExpectedInput: []string{"false"}},
	// ToByteCondition
	{Title: "ToByteCondition - True returns trueVal", ArrangeInput: args.Map{"method": "ToByteCondition", "value": issetter.True, "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"10"}},
	{Title: "ToByteCondition - False returns falseVal", ArrangeInput: args.Map{"method": "ToByteCondition", "value": issetter.False, "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"20"}},
	{Title: "ToByteCondition - Uninitialized returns invalid", ArrangeInput: args.Map{"method": "ToByteCondition", "value": issetter.Uninitialized, "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"255"}},
	{Title: "ToByteCondition - Set returns invalid", ArrangeInput: args.Map{"method": "ToByteCondition", "value": issetter.Set, "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"255"}},
	{Title: "ToByteCondition - Wildcard returns invalid", ArrangeInput: args.Map{"method": "ToByteCondition", "value": issetter.Wildcard, "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"255"}},
	// ToByteConditionWithWildcard
	{Title: "ToByteConditionWithWildcard - Wildcard returns wildcard byte", ArrangeInput: args.Map{"method": "ToByteConditionWithWildcard", "value": issetter.Wildcard, "wildcardVal": byte(99), "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"99"}},
	{Title: "ToByteConditionWithWildcard - True returns trueVal", ArrangeInput: args.Map{"method": "ToByteConditionWithWildcard", "value": issetter.True, "wildcardVal": byte(99), "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"10"}},
	{Title: "ToByteConditionWithWildcard - False returns falseVal", ArrangeInput: args.Map{"method": "ToByteConditionWithWildcard", "value": issetter.False, "wildcardVal": byte(99), "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"20"}},
	{Title: "ToByteConditionWithWildcard - Uninitialized returns invalid", ArrangeInput: args.Map{"method": "ToByteConditionWithWildcard", "value": issetter.Uninitialized, "wildcardVal": byte(99), "trueVal": byte(10), "falseVal": byte(20), "invalidVal": byte(255)}, ExpectedInput: []string{"255"}},
	// IsDefinedLogically
	{Title: "IsDefinedLogically - Uninitialized false", ArrangeInput: args.Map{"method": "IsDefinedLogically", "value": issetter.Uninitialized}, ExpectedInput: []string{"false"}},
	{Title: "IsDefinedLogically - True true", ArrangeInput: args.Map{"method": "IsDefinedLogically", "value": issetter.True}, ExpectedInput: []string{"true"}},
	{Title: "IsDefinedLogically - False true", ArrangeInput: args.Map{"method": "IsDefinedLogically", "value": issetter.False}, ExpectedInput: []string{"true"}},
	{Title: "IsDefinedLogically - Unset true", ArrangeInput: args.Map{"method": "IsDefinedLogically", "value": issetter.Unset}, ExpectedInput: []string{"true"}},
	{Title: "IsDefinedLogically - Set true", ArrangeInput: args.Map{"method": "IsDefinedLogically", "value": issetter.Set}, ExpectedInput: []string{"true"}},
	{Title: "IsDefinedLogically - Wildcard false", ArrangeInput: args.Map{"method": "IsDefinedLogically", "value": issetter.Wildcard}, ExpectedInput: []string{"false"}},
	// IsUndefinedLogically
	{Title: "IsUndefinedLogically - Uninitialized true", ArrangeInput: args.Map{"method": "IsUndefinedLogically", "value": issetter.Uninitialized}, ExpectedInput: []string{"true"}},
	{Title: "IsUndefinedLogically - True false", ArrangeInput: args.Map{"method": "IsUndefinedLogically", "value": issetter.True}, ExpectedInput: []string{"false"}},
	{Title: "IsUndefinedLogically - False false", ArrangeInput: args.Map{"method": "IsUndefinedLogically", "value": issetter.False}, ExpectedInput: []string{"false"}},
	{Title: "IsUndefinedLogically - Unset false", ArrangeInput: args.Map{"method": "IsUndefinedLogically", "value": issetter.Unset}, ExpectedInput: []string{"false"}},
	{Title: "IsUndefinedLogically - Set false", ArrangeInput: args.Map{"method": "IsUndefinedLogically", "value": issetter.Set}, ExpectedInput: []string{"false"}},
	{Title: "IsUndefinedLogically - Wildcard true", ArrangeInput: args.Map{"method": "IsUndefinedLogically", "value": issetter.Wildcard}, ExpectedInput: []string{"true"}},
	// IsPositive
	{Title: "IsPositive - Uninitialized false", ArrangeInput: args.Map{"method": "IsPositive", "value": issetter.Uninitialized}, ExpectedInput: []string{"false"}},
	{Title: "IsPositive - True true", ArrangeInput: args.Map{"method": "IsPositive", "value": issetter.True}, ExpectedInput: []string{"true"}},
	{Title: "IsPositive - False false", ArrangeInput: args.Map{"method": "IsPositive", "value": issetter.False}, ExpectedInput: []string{"false"}},
	{Title: "IsPositive - Unset false", ArrangeInput: args.Map{"method": "IsPositive", "value": issetter.Unset}, ExpectedInput: []string{"false"}},
	{Title: "IsPositive - Set true", ArrangeInput: args.Map{"method": "IsPositive", "value": issetter.Set}, ExpectedInput: []string{"true"}},
	{Title: "IsPositive - Wildcard false", ArrangeInput: args.Map{"method": "IsPositive", "value": issetter.Wildcard}, ExpectedInput: []string{"false"}},
	// IsNegative
	{Title: "IsNegative - Uninitialized true", ArrangeInput: args.Map{"method": "IsNegative", "value": issetter.Uninitialized}, ExpectedInput: []string{"true"}},
	{Title: "IsNegative - True false", ArrangeInput: args.Map{"method": "IsNegative", "value": issetter.True}, ExpectedInput: []string{"false"}},
	{Title: "IsNegative - False true", ArrangeInput: args.Map{"method": "IsNegative", "value": issetter.False}, ExpectedInput: []string{"true"}},
	{Title: "IsNegative - Unset true", ArrangeInput: args.Map{"method": "IsNegative", "value": issetter.Unset}, ExpectedInput: []string{"true"}},
	{Title: "IsNegative - Set false", ArrangeInput: args.Map{"method": "IsNegative", "value": issetter.Set}, ExpectedInput: []string{"false"}},
	{Title: "IsNegative - Wildcard false", ArrangeInput: args.Map{"method": "IsNegative", "value": issetter.Wildcard}, ExpectedInput: []string{"false"}},
}

// =============================================================================
// Test Runner
// =============================================================================

func Test_Value_Logic_Verification(t *testing.T) {
	for caseIndex, tc := range valueLogicTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method := input["method"].(string)
		value := input["value"].(issetter.Value)

		var result string

		// Act
		switch method {
		case "IsOnLogically":
			result = fmt.Sprintf("%v", value.IsOnLogically())
		case "IsOffLogically":
			result = fmt.Sprintf("%v", value.IsOffLogically())
		case "WildcardApply":
			result = fmt.Sprintf("%v", value.WildcardApply(input["input"].(bool)))
		case "IsWildcardOrBool":
			result = fmt.Sprintf("%v", value.IsWildcardOrBool(input["input"].(bool)))
		case "ToByteCondition":
			result = fmt.Sprintf("%v", value.ToByteCondition(input["trueVal"].(byte), input["falseVal"].(byte), input["invalidVal"].(byte)))
		case "ToByteConditionWithWildcard":
			result = fmt.Sprintf("%v", value.ToByteConditionWithWildcard(input["wildcardVal"].(byte), input["trueVal"].(byte), input["falseVal"].(byte), input["invalidVal"].(byte)))
		case "IsDefinedLogically":
			result = fmt.Sprintf("%v", value.IsDefinedLogically())
		case "IsUndefinedLogically":
			result = fmt.Sprintf("%v", value.IsUndefinedLogically())
		case "IsPositive":
			result = fmt.Sprintf("%v", value.IsPositive())
		case "IsNegative":
			result = fmt.Sprintf("%v", value.IsNegative())
		default:
			t.Fatalf("unknown method: %s", method)
		}

		actLines := []string{result}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

// =============================================================================
// Stateful tests (GetSetBoolOnInvalid, LazyEvaluateBool, LazyEvaluateSet)
// =============================================================================

var getSetBoolTestCases = []coretestcases.CaseV1{
	{Title: "GetSetBoolOnInvalid - already True returns true ignores setter", ArrangeInput: args.Map{"initial": issetter.True, "setter": false}, ExpectedInput: []string{"true", "true"}},
	{Title: "GetSetBoolOnInvalid - already False returns false ignores setter", ArrangeInput: args.Map{"initial": issetter.False, "setter": true}, ExpectedInput: []string{"false", "true"}},
	{Title: "GetSetBoolOnInvalid - Uninitialized with true sets True", ArrangeInput: args.Map{"initial": issetter.Uninitialized, "setter": true}, ExpectedInput: []string{"true", "true"}},
	{Title: "GetSetBoolOnInvalid - Uninitialized with false sets False", ArrangeInput: args.Map{"initial": issetter.Uninitialized, "setter": false}, ExpectedInput: []string{"false", "true"}},
	{Title: "GetSetBoolOnInvalid - Set triggers setter with true", ArrangeInput: args.Map{"initial": issetter.Set, "setter": true}, ExpectedInput: []string{"true", "true"}},
}

func Test_Value_GetSetBoolOnInvalid_Verification(t *testing.T) {
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

var lazyEvaluateBoolTestCases = []coretestcases.CaseV1{
	{Title: "LazyEvaluateBool - Uninitialized calls func sets True", ArrangeInput: args.Map{"initial": issetter.Uninitialized, "expectCalled": true, "expectResult": true}, ExpectedInput: []string{"true", "true", "true"}},
	{Title: "LazyEvaluateBool - already True skips func", ArrangeInput: args.Map{"initial": issetter.True, "expectCalled": false, "expectResult": false}, ExpectedInput: []string{"false", "false"}},
	{Title: "LazyEvaluateBool - already False skips func", ArrangeInput: args.Map{"initial": issetter.False, "expectCalled": false, "expectResult": false}, ExpectedInput: []string{"false", "false"}},
}

func Test_Value_LazyEvaluateBool_Verification(t *testing.T) {
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

var lazyEvaluateSetTestCases = []coretestcases.CaseV1{
	{Title: "LazyEvaluateSet - Uninitialized calls func sets Set", ArrangeInput: args.Map{"initial": issetter.Uninitialized, "expectCalled": true, "expectResult": true}, ExpectedInput: []string{"true", "true", "true"}},
	{Title: "LazyEvaluateSet - already Set skips func", ArrangeInput: args.Map{"initial": issetter.Set, "expectCalled": false, "expectResult": false}, ExpectedInput: []string{"false", "false"}},
	{Title: "LazyEvaluateSet - already Unset skips func", ArrangeInput: args.Map{"initial": issetter.Unset, "expectCalled": false, "expectResult": false}, ExpectedInput: []string{"false", "false"}},
}

func Test_Value_LazyEvaluateSet_Verification(t *testing.T) {
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
