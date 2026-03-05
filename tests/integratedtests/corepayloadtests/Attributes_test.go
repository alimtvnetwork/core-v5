package corepayloadtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/errcore"
)

// getBoolDefault extracts a bool from args.Map with a default value.
func getBoolDefault(input args.Map, key string, defaultVal bool) bool {
	raw, found := input.Get(key)
	if !found {
		return defaultVal
	}
	val, ok := raw.(bool)
	if !ok {
		return defaultVal
	}
	return val
}

// =============================================================================
// Attributes.IsEqual — Regression: logic inversion bug in IsSafeValid/HasIssuesOrEmpty
// =============================================================================

func Test_Attributes_IsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := getBoolDefault(input, "left_nil", false)
		rightNil := getBoolDefault(input, "right_nil", false)
		samePointer := getBoolDefault(input, "same_pointer", false)

		var left, right *corepayload.Attributes

		if !leftNil {
			leftPayload, _ := input.GetAsString("left_payload")
			payload, _ := input.GetAsString("payload")

			if leftPayload == "" {
				leftPayload = payload
			}

			left = &corepayload.Attributes{
				DynamicPayloads: []byte(leftPayload),
			}
		}

		if samePointer {
			right = left
		} else if !rightNil {
			rightPayload, _ := input.GetAsString("right_payload")
			payload, _ := input.GetAsString("payload")

			if rightPayload == "" {
				rightPayload = payload
			}

			right = &corepayload.Attributes{
				DynamicPayloads: []byte(rightPayload),
			}
		}

		// Act
		result := left.IsEqual(right)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}

// =============================================================================
// Attributes.Clone — Regression: deep clone independence
// =============================================================================

func Test_Attributes_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilAttr := getBoolDefault(input, "nil_attr", false)
		deep := getBoolDefault(input, "deep", false)

		if nilAttr {
			// Act
			var attr *corepayload.Attributes
			clonedPtr, err := attr.ClonePtr(deep)

			// Assert
			testCase.ShouldBeEqual(t, caseIndex,
				fmt.Sprintf("%v", clonedPtr == nil),
				fmt.Sprintf("%v", err != nil),
			)

			continue
		}

		payload, _ := input.GetAsString("payload")
		attr := &corepayload.Attributes{
			DynamicPayloads: []byte(payload),
		}

		// Act
		cloned, err := attr.Clone(deep)
		errcore.HandleErr(err)

		isEqual := attr.IsEqual(&cloned)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			string(cloned.DynamicPayloads),
			fmt.Sprintf("%v", isEqual),
		)
	}
}

// =============================================================================
// Attributes.IsSafeValid — Regression: was returning HasIssuesOrEmpty() without negation
// =============================================================================

func Test_Attributes_IsSafeValid_Verification(t *testing.T) {
	for caseIndex, testCase := range attributesIsSafeValidTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilAttr := getBoolDefault(input, "nil_attr", false)
		empty := getBoolDefault(input, "empty", false)

		var attr *corepayload.Attributes

		if !nilAttr && !empty {
			payload, _ := input.GetAsString("payload")
			attr = &corepayload.Attributes{
				DynamicPayloads: []byte(payload),
			}
		} else if !nilAttr && empty {
			attr = &corepayload.Attributes{}
		}

		// Act
		result := attr.IsSafeValid()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}

// =============================================================================
// AuthInfo.Clone — Regression: was missing Identifier field in clone
// =============================================================================

func Test_AuthInfo_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range authInfoCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilAuth := getBoolDefault(input, "nil_auth", false)

		if nilAuth {
			// Act
			var auth *corepayload.AuthInfo
			cloned := auth.ClonePtr()

			// Assert
			testCase.ShouldBeEqual(t, caseIndex,
				fmt.Sprintf("%v", cloned == nil),
			)

			continue
		}

		identifier, _ := input.GetAsString("identifier")
		actionType, _ := input.GetAsString("action_type")
		resourceName, _ := input.GetAsString("resource_name")
		newActionType, _ := input.GetAsString("new_action_type")

		auth := &corepayload.AuthInfo{
			Identifier:   identifier,
			ActionType:   actionType,
			ResourceName: resourceName,
		}

		// Act
		cloned := auth.ClonePtr()

		if newActionType != "" {
			// Test independence: mutate clone
			cloned.ActionType = newActionType

			// Assert — original unchanged, clone mutated
			testCase.ShouldBeEqual(t, caseIndex,
				auth.ActionType,
				cloned.ActionType,
			)

			continue
		}

		// Assert — all fields including Identifier are preserved
		testCase.ShouldBeEqual(t, caseIndex,
			cloned.Identifier,
			cloned.ActionType,
			cloned.ResourceName,
		)
	}
}
