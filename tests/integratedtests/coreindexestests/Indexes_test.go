package coreindexestests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_HasIndex_Verification(t *testing.T) {
	for caseIndex, testCase := range hasIndexTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		indexesVal, _ := input.Get("indexes")
		indexes := indexesVal.([]int)
		current, _ := input.GetAsInt("current")

		// Act
		result := coreindexes.HasIndex(indexes, current)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_LastIndex_Verification(t *testing.T) {
	for caseIndex, testCase := range lastIndexTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		length, _ := input.GetAsInt("length")

		// Act
		result := coreindexes.LastIndex(length)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsWithinIndexRange_Verification(t *testing.T) {
	for caseIndex, testCase := range isWithinIndexRangeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		index, _ := input.GetAsInt("index")
		length, _ := input.GetAsInt("length")

		// Act
		result := coreindexes.IsWithinIndexRange(index, length)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
