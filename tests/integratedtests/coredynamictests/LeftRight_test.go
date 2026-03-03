package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Test: IsEmpty
// ==========================================

func Test_LeftRight_IsEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsEmptyTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.IsEmpty())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: HasLeft
// ==========================================

func Test_LeftRight_HasLeft_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasLeftTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.HasLeft())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: HasRight
// ==========================================

func Test_LeftRight_HasRight_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasRightTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.HasRight())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: IsLeftEmpty
// ==========================================

func Test_LeftRight_IsLeftEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsLeftEmptyTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.IsLeftEmpty())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: IsRightEmpty
// ==========================================

func Test_LeftRight_IsRightEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsRightEmptyTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.IsRightEmpty())}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: DeserializeLeft
// ==========================================

func Test_LeftRight_DeserializeLeft_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightDeserializeLeftTestCases {
		result := tc.LR.DeserializeLeft()

		var actLines []string
		if result == nil {
			actLines = []string{"true"}
		} else {
			actLines = []string{
				fmt.Sprintf("%v", result == nil),
				fmt.Sprintf("%v", result.HasError()),
			}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: DeserializeRight
// ==========================================

func Test_LeftRight_DeserializeRight_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightDeserializeRightTestCases {
		result := tc.LR.DeserializeRight()

		var actLines []string
		if result == nil {
			actLines = []string{"true"}
		} else {
			actLines = []string{fmt.Sprintf("%v", result == nil)}
		}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}

// ==========================================
// Test: TypeStatus
// ==========================================

func Test_LeftRight_TypeStatus_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightTypeStatusTestCases {
		status := tc.LR.TypeStatus()
		actLines := []string{fmt.Sprintf("%v", status != nil)}

		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
	}
}
