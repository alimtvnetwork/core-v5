package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: IsEmpty
// ==========================================

func Test_LeftRight_IsEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsEmptyTestCases {
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.IsEmpty()))
	}
}

// ==========================================
// Test: HasLeft
// ==========================================

func Test_LeftRight_HasLeft_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasLeftTestCases {
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.HasLeft()))
	}
}

// ==========================================
// Test: HasRight
// ==========================================

func Test_LeftRight_HasRight_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasRightTestCases {
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.HasRight()))
	}
}

// ==========================================
// Test: IsLeftEmpty
// ==========================================

func Test_LeftRight_IsLeftEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsLeftEmptyTestCases {
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.IsLeftEmpty()))
	}
}

// ==========================================
// Test: IsRightEmpty
// ==========================================

func Test_LeftRight_IsRightEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsRightEmptyTestCases {
		tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", tc.LR.IsRightEmpty()))
	}
}

// ==========================================
// Test: DeserializeLeft
// ==========================================

func Test_LeftRight_DeserializeLeft_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightDeserializeLeftTestCases {
		result := tc.LR.DeserializeLeft()

		if result == nil {
			tc.Case.ShouldBeEqual(t, caseIndex, "true")
		} else {
			actual := args.Map{
				"isNil":    result == nil,
				"hasError": result.HasError(),
			}

			tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

// ==========================================
// Test: DeserializeRight
// ==========================================

func Test_LeftRight_DeserializeRight_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightDeserializeRightTestCases {
		result := tc.LR.DeserializeRight()

		if result == nil {
			tc.Case.ShouldBeEqual(t, caseIndex, "true")
		} else {
			tc.Case.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result == nil))
		}
	}
}

// ==========================================
// Test: TypeStatus
// ==========================================

func Test_LeftRight_TypeStatus_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightTypeStatusTestCases {
		status := tc.LR.TypeStatus()

		actual := args.Map{
			"isSame":             fmt.Sprintf("%v", status.IsSame),
			"isLeftUnknownNull":  fmt.Sprintf("%v", status.IsLeftUnknownNull),
			"isRightUnknownNull": fmt.Sprintf("%v", status.IsRightUnknownNull),
		}

		if expectedMap, ok := tc.Case.ExpectedInput.(args.Map); ok {
			if _, has := expectedMap["isLeftPointer"]; has {
				actual["isLeftPointer"] = fmt.Sprintf("%v", status.IsLeftPointer)
			}
			if _, has := expectedMap["isRightPointer"]; has {
				actual["isRightPointer"] = fmt.Sprintf("%v", status.IsRightPointer)
			}
		}

		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
