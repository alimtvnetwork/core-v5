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

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: HasLeft
// ==========================================

func Test_LeftRight_HasLeft_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasLeftTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.HasLeft())}

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: HasRight
// ==========================================

func Test_LeftRight_HasRight_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightHasRightTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.HasRight())}

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsLeftEmpty
// ==========================================

func Test_LeftRight_IsLeftEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsLeftEmptyTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.IsLeftEmpty())}

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: IsRightEmpty
// ==========================================

func Test_LeftRight_IsRightEmpty_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightIsRightEmptyTestCases {
		actLines := []string{fmt.Sprintf("%v", tc.LR.IsRightEmpty())}

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
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

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
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

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypeStatus
// ==========================================

func Test_LeftRight_TypeStatus_Verification(t *testing.T) {
	for caseIndex, tc := range leftRightTypeStatusTestCases {
		status := tc.LR.TypeStatus()
		actLines := []string{fmt.Sprintf("%v", status != nil)}

		errcore.PrintDiffOnMismatch(caseIndex, tc.Case.Title, actLines, tc.Case.ExpectedInput)
		tc.Case.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
