package coredatatests

import (
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_StringsDsc_Len(t *testing.T) {
	for caseIndex, tc := range stringsDscLenTestCases {
		// Arrange
		var strings coredata.StringsDsc
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			if vals, ok := input["values"]; ok {
				strings = coredata.StringsDsc(vals.([]string))
			}
		}

		// Act
		actual := args.Map{
			"length": strings.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringsDsc_Sort(t *testing.T) {
	for caseIndex, tc := range stringsDscSortTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		src := input["values"].([]string)
		strs := make(coredata.StringsDsc, len(src))
		copy(strs, src)

		// Act
		sort.Sort(strs)

		actual := args.Map{
			"first": strs[0],
			"last":  strs[len(strs)-1],
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
