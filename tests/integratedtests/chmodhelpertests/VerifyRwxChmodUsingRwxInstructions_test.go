package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/msgtype"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxChmodUsingRwxInstructions_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := chmodhelpertestwrappers.CreatePathInstruction2
	createDefaultPaths(&createPathInstructions)
	for i, testCase := range chmodhelpertestwrappers.VerifyRwxChmodUsingRwxInstructionsTestCases {
		executor, err := chmodhelper.ParseRwxInstructionToExecutor(&testCase.RwxInstruction)

		msgtype.SimpleHandleErr(err, "")

		expectationMessage := testCase.ExpectedErrorMessage
		err2 := executor.VerifyRwxModifiersDirect(
			false,
			testCase.Locations...)

		// Act
		actualMessage := err2.Error()

		// Assert
		Convey(testCase.Header, t, func() {
			isEqual := coretests.IsStringMessageWithoutWhitespaceSortedEqual(
				true,
				true,
				testCase.Header,
				actualMessage,
				expectationMessage,
				i)

			So(isEqual, ShouldBeTrue)
		})
	}
}
