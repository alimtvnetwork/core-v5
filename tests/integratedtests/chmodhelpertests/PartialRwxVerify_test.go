package chmodhelpertests

import (
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_PartialRwxVerify(t *testing.T) {
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
