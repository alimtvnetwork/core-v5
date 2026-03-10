package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
)

func Test_VerifyRwxPartialChmodLocations_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Setup
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		pathInstructionsV2,
	)

	for caseIndex, testCase := range verifyPartialRwxLocationsTestCases {
		// Arrange
		wrapper := testCase.ArrangeInput.(chmodhelpertestwrappers.VerifyRwxPartialChmodLocationsWrapper)

		// Act
		err := chmodhelper.ChmodVerify.PathsUsingPartialRwxOptions(
			wrapper.IsContinueOnError,
			wrapper.IsSkipOnInvalid,
			wrapper.ExpectedPartialRwx,
			wrapper.Locations...,
		)

		// Assert
		assertNonWhiteSortedEqual(t, testCase, caseIndex, err)
	}
}
