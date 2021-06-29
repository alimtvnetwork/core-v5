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
		expectationMessage := testCase.ExpectedErrorMessage
		executor, err := chmodhelper.ParseRwxInstructionToExecutor(&testCase.RwxInstruction)

		msgtype.SimpleHandleErr(err, "")

		// Act
		actualErr := executor.VerifyRwxModifiersDirect(
			false,
			testCase.Locations...)

		// Assert
		Convey(testCase.Header, t, func() {
			isEqual := coretests.IsStringErrorWithoutWhitespaceSortedEqual(
				true,
				true,
				testCase.Header,
				actualErr,
				expectationMessage,
				i)

			So(isEqual, ShouldBeTrue)
		})
	}
}
