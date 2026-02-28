package coresorttests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coresort/intsort"
	"gitlab.com/auk-go/core/coresort/strsort"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_IntSort_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range intSortQuickTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")
		original := inputVal.([]int)
		// clone to avoid mutation issues
		clone := make([]int, len(original))
		copy(clone, original)

		// Act
		result := intsort.Quick(&clone)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", *result))
	}
}

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
