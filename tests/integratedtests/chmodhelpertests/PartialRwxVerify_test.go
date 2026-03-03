package chmodhelpertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
)

func Test_PartialRwxVerify(t *testing.T) {
	for caseIndex, testCase := range partialRwxVerifyTestCases {
		// Arrange
		inputs := testCase.ArrangeInput.(map[string]string)
		partialRwx := inputs["partialRwx"]
		fullRwx := inputs["fullRwx"]

		rwx, err := chmodhelper.NewRwxVariableWrapper(partialRwx)
		errcore.SimpleHandleErr(err, "rwxVar create failed.")

		// Act
		actual := rwx.IsEqualPartialRwxPartial(fullRwx)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", actual))
	}
}
