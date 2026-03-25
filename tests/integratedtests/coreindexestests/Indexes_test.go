package coreindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreindexes"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_HasIndex_Verification(t *testing.T) {
	for caseIndex, tc := range hasIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		indexesVal, _ := input.Get("indexes")
		indexes := indexesVal.([]int)
		current, _ := input.GetAsInt("current")

		// Act
		result := coreindexes.HasIndex(indexes, current)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LastIndex_Verification(t *testing.T) {
	for caseIndex, tc := range lastIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		length, _ := input.GetAsInt("length")

		// Act
		result := coreindexes.LastIndex(length)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsWithinIndexRange_Verification(t *testing.T) {
	for caseIndex, tc := range isWithinIndexRangeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		index, _ := input.GetAsInt("index")
		length, _ := input.GetAsInt("length")

		// Act
		result := coreindexes.IsWithinIndexRange(index, length)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
