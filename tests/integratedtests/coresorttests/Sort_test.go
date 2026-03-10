package coresorttests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coresort/intsort"
	"github.com/alimtvnetwork/core/coresort/strsort"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// intsort.Quick — ascending
// =============================================================================

func Test_IntSort_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		clone := make([]int, len(original))
		copy(clone, original)

		// Act
		result := intsort.Quick(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// intsort.QuickDsc — descending
// =============================================================================

func Test_IntSort_QuickDsc_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickDscTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		clone := make([]int, len(original))
		copy(clone, original)

		// Act
		result := intsort.QuickDsc(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// intsort.QuickPtr — ascending pointer sort
// =============================================================================

func Test_IntSort_QuickPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		ptrs := toIntPtrs(original)

		// Act
		result := intsort.QuickPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatIntPtrs(*result))
	}
}

// =============================================================================
// intsort.QuickDscPtr — descending pointer sort
// =============================================================================

func Test_IntSort_QuickDscPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickDscPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		ptrs := toIntPtrs(original)

		// Act
		result := intsort.QuickDscPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatIntPtrs(*result))
	}
}

// =============================================================================
// strsort.Quick — ascending
// =============================================================================

func Test_StrSort_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		clone := make([]string, len(inputVal))
		copy(clone, inputVal)

		// Act
		result := strsort.Quick(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// strsort.QuickDsc — descending
// =============================================================================

func Test_StrSort_QuickDsc_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickDscTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		clone := make([]string, len(inputVal))
		copy(clone, inputVal)

		// Act
		result := strsort.QuickDsc(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

// =============================================================================
// strsort.QuickPtr — ascending pointer sort
// =============================================================================

func Test_StrSort_QuickPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		ptrs := toStrPtrs(inputVal)

		// Act
		result := strsort.QuickPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatStrPtrs(*result))
	}
}

// =============================================================================
// strsort.QuickDscPtr — descending pointer sort
// =============================================================================

func Test_StrSort_QuickDscPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range strSortQuickDscPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsStrings("input")
		ptrs := toStrPtrs(inputVal)

		// Act
		result := strsort.QuickDscPtr(&ptrs)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, formatStrPtrs(*result))
	}
}
