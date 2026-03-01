package coreversiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coreversion"
	"gitlab.com/auk-go/core/enums/versionindexes"
	"gitlab.com/auk-go/core/errcore"
)

func Test_ComparisonValueIndexes_Verification(t *testing.T) {
	for caseIndex, testCase := range comparisonValueIndexesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, err := input.GetAsString("left")
		errcore.HandleErrMessage(err, "left is required")
		rightStr, err := input.GetAsString("right")
		errcore.HandleErrMessage(err, "right is required")

		// Act
		leftV := coreversion.New.Create(leftStr)
		rightV := coreversion.New.Create(rightStr)
		result := leftV.ComparisonValueIndexes(
			&rightV,
			versionindexes.AllVersionIndexes...,
		)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result.Name(),
		)
	}
}

func Test_VersionSliceInteger_Verification(t *testing.T) {
	for caseIndex, testCase := range versionSliceIntegerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, err := input.GetAsString("left")
		errcore.HandleErrMessage(err, "left is required")
		rightStr, err := input.GetAsString("right")
		errcore.HandleErrMessage(err, "right is required")

		// Act
		leftV := coreversion.New.Create(leftStr)
		rightV := coreversion.New.Create(rightStr)
		leftValues := leftV.AllVersionValues()
		rightValues := rightV.AllVersionValues()
		result := corecmp.VersionSliceInteger(leftValues, rightValues)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result.Name(),
		)
	}
}

func Test_IsAtLeast_Verification(t *testing.T) {
	for caseIndex, testCase := range isAtLeastTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, err := input.GetAsString("left")
		errcore.HandleErrMessage(err, "left is required")
		rightStr, err := input.GetAsString("right")
		errcore.HandleErrMessage(err, "right is required")

		// Act
		result := coreversion.IsAtLeast(leftStr, rightStr)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}

func Test_IsLower_Verification(t *testing.T) {
	for caseIndex, testCase := range isLowerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, err := input.GetAsString("left")
		errcore.HandleErrMessage(err, "left is required")
		rightStr, err := input.GetAsString("right")
		errcore.HandleErrMessage(err, "right is required")

		// Act
		result := coreversion.IsLower(leftStr, rightStr)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}

func Test_IsExpectedVersion_Verification(t *testing.T) {
	for caseIndex, testCase := range isExpectedVersionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftStr, err := input.GetAsString("left")
		errcore.HandleErrMessage(err, "left is required")
		rightStr, err := input.GetAsString("right")
		errcore.HandleErrMessage(err, "right is required")
		expected := input["expected"].(corecomparator.Compare)

		// Act
		result := coreversion.IsExpectedVersion(expected, leftStr, rightStr)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", result),
		)
	}
}
