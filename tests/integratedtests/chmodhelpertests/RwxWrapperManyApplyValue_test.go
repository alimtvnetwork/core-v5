package chmodhelpertests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/errcore"
)

// Test_RwxWrapperManyApplyValue_Unix
//
//	for directory `-` will be placed not `d`
func Test_RwxWrapperManyApplyValue_Unix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	createPathInstructions := pathInstructionsV2
	chmodhelper.CreateDirFilesWithRwxPermissionsMust(
		true,
		createPathInstructions,
	)

	firstCreationIns := createPathInstructions[0]
	paths := firstCreationIns.GetPaths()
	condition := chmodins.DefaultAllTrueCondition()
	existingAppliedRwxFull := firstCreationIns.ApplyRwx.String()

	for caseIndex, testCase := range rwxWrapperManyApplyTestCases {
		// Arrange
		rwxWrapper, err := testCase.SingleRwx.ToDisabledRwxWrapper()
		errcore.SimpleHandleErr(err, "SingleRwx ToDisabledRwxWrapper failed")

		expectation := rwxWrapper.ToFullRwxValueString()

		header := fmt.Sprintf(
			"Existing [%s] Applied by [%s] should result [%s]",
			existingAppliedRwxFull,
			expectation,
			expectation,
		)

		// Act
		applyErr := rwxWrapper.ApplyLinuxChmodOnMany(condition, paths...)
		errcore.SimpleHandleErr(
			applyErr,
			"rwxWrapper.ApplyLinuxChmodOnMany failed",
		)

		fileChmodMap := firstCreationIns.GetFilesChmodMap()
		var actLines []string

		for filePath, chmodValueString := range fileChmodMap.Items() {
			isEqual := chmodValueString == expectation
			actLines = append(actLines, fmt.Sprintf(
				"%s=%v",
				filePath,
				isEqual,
			))

			if !isEqual {
				errcore.PrintDiffOnMismatch(
					caseIndex,
					testCase.Case.Title,
					[]string{chmodValueString},
					[]string{expectation},
					fmt.Sprintf("  File: %s", filePath),
				)
			}
		}

		// Assert
		assertSingleChmod(
			t,
			header,
			firstCreationIns,
			expectation,
		)
	}
}
