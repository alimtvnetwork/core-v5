package chmodhelpertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxPartialChmodLocations_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := chmodhelpertestwrappers.CreatePathInstruction2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		&createPathInstructions)
	for i, testCase := range chmodhelpertestwrappers.VerifyRwxPartialChmodLocationsTestCases {
		header := testCase.Header
		expectationMessage := testCase.ExpectationErrorMessage

		// Act
		err := chmodhelper.VerifyChmodLocationsUsingPartialRwx(
			testCase.IsContinueOnError,
			testCase.IsSkipOnInvalid,
			testCase.ExpectedPartialRwx,
			testCase.Locations)

		// Assert
		Convey(header, t, func() {
			isEqual := coretests.IsStringErrorWithoutWhitespaceSortedEqual(
				true,
				true,
				testCase.Header,
				err,
				expectationMessage,
				i)

			So(isEqual, ShouldBeTrue)
		})
	}
}
