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

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			fmt.Sprintf("%d", len(result)),
			result[0],
			result[len(result)-1],
		)
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
