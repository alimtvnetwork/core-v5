package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/isany"
)

func Test_SliceValidator(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorTestCases {
		// Arrange
		inputs := testCase.
			Case.
			ArrangeInput.([]args.TwoAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				"%d : %t (%s, %s)",
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
			)
		}

		actLines := actualSlice.Strings()
		actualError := testCase.Case.VerifyAllEqual(
			caseIndex,
			actLines...,
		)
		validator := testCase.Validator
		errLines := errcore.ErrorToSplitLines(actualError)

		// Assert
		validator.AssertAllQuick(
			t,
			caseIndex,
			testCase.Case.Title,
			errLines...,
		)
	}
}

func Test_SliceValidator_FirstError(t *testing.T) {
	for caseIndex, testCase := range sliceValidatorFirstErrorTestCases {
		// Arrange
		inputs := testCase.
			Case.
			ArrangeInput.([]args.TwoAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				"%d : %t (%s, %s)",
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
			)
		}

		actLines := actualSlice.Strings()
		actualError := testCase.Case.VerifyFirst(
			caseIndex,
			stringcompareas.Equal,
			actLines,
		)
		validator := testCase.Validator
		errLines := errcore.ErrorToSplitLines(actualError)

		// Assert
		validator.AssertAllQuick(
			t,
			caseIndex,
			testCase.Case.Title,
			errLines...,
		)
	}
}
