package keymktests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/keymk"
)

func Test_KeyLegend_GroupIntRange_Verification(t *testing.T) {
	for caseIndex, testCase := range keyLegendGroupIntRangeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		root, _ := input.GetAsString("root")
		pkg, _ := input.GetAsString("package")
		group, _ := input.GetAsString("group")
		state, _ := input.GetAsString("state")
		startId := input.GetAsIntDefault("startId", 0)
		endId := input.GetAsIntDefault("endId", 0)

		// Act
		k := keymk.NewKeyWithLegend.All(
			keymk.JoinerOption,
			keymk.ShortLegends,
			false,
			root, pkg, group, state,
		)
		result := k.GroupIntRange(startId, endId)

		actual := args.Map{
			"count":    fmt.Sprintf("%d", len(result)),
			"firstKey": result[0],
			"lastKey":  result[len(result)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_KeyLegend_UserStringWithoutState_Verification(t *testing.T) {
	for caseIndex, testCase := range keyLegendUserStringWithoutStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		root, _ := input.GetAsString("root")
		pkg, _ := input.GetAsString("package")
		group, _ := input.GetAsString("group")
		state, _ := input.GetAsString("state")
		user, _ := input.GetAsString("user")

		// Act
		k := keymk.NewKeyWithLegend.All(
			keymk.JoinerOption,
			keymk.ShortLegends,
			false,
			root, pkg, group, state,
		)
		result := k.UserStringWithoutState(user)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_KeyLegend_UpToState_Verification(t *testing.T) {
	for caseIndex, testCase := range keyLegendUpToStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		root, _ := input.GetAsString("root")
		pkg, _ := input.GetAsString("package")
		group, _ := input.GetAsString("group")
		state, _ := input.GetAsString("state")
		user, _ := input.GetAsString("user")

		// Act
		k := keymk.NewKeyWithLegend.All(
			keymk.JoinerOption,
			keymk.ShortLegends,
			false,
			root, pkg, group, state,
		)
		result := k.UpToState(user)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
