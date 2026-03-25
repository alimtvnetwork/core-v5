package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_PartialRwxVerify_Verification(t *testing.T) {
	for caseIndex, testCase := range partialRwxVerifyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		partialRwx, _ := input.GetAsString("partialRwx")
		fullRwx, _ := input.GetAsString("fullRwx")

		rwx, err := chmodhelper.NewRwxVariableWrapper(partialRwx)
		errcore.SimpleHandleErr(err, "rwxVar create failed.")

		// Act
		actual := rwx.IsEqualPartialRwxPartial(fullRwx)

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, args.Map{
			"isEqual": actual,
		})
	}
}
