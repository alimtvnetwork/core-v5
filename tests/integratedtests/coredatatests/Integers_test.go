package coredatatests

import (
	"sort"
	"testing"

	"gitlab.com/auk-go/core/coredata"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_Integers_Len(t *testing.T) {
	for caseIndex, tc := range integersLenTestCases {
		// Arrange
		var integers coredata.Integers
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			if vals, ok := input["values"]; ok {
				integers = coredata.Integers(vals.([]int))
			}
		}

		// Act
		actual := args.Map{
			"length": integers.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Integers_Less(t *testing.T) {
	tc := integersLessTestCases[0]

	// Arrange
	integers := coredata.Integers{5, 3, 8}

	// Act
	actual := args.Map{
		"less10": integers.Less(1, 0),
		"less01": integers.Less(0, 1),
		"less00": integers.Less(0, 0),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Integers_Sort(t *testing.T) {
	for caseIndex, tc := range integersSortTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		src := input["values"].([]int)
		integers := make(coredata.Integers, len(src))
		copy(integers, src)

		// Act
		sort.Sort(integers)

		actual := args.Map{
			"first": integers[0],
			"last":  integers[len(integers)-1],
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
