package corecomparatortests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// Test_Compare_JsonRoundtrip verifies MarshalJSON produces the correct name
// string and UnmarshalJSON restores identity including Name() and NumberString().
// Migrated from cmd/main/enumTesting.go.
func Test_Compare_JsonRoundtrip(t *testing.T) {
	for caseIndex, testCase := range compareJsonRoundtripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		value, err := input.GetAsInt("value")
		errcore.HandleErrMessage(
			err,
			"value is required for compare JSON roundtrip test",
		)

		unmarshalInput, err := input.GetAsString("unmarshalInput")
		errcore.HandleErrMessage(
			err,
			"unmarshalInput is required for compare JSON roundtrip test",
		)

		compare := corecomparator.Compare(value)

		// Act — marshal
		marshaledBytes, marshalErr := compare.MarshalJSON()
		errcore.HandleErrMessage(
			marshalErr,
			"MarshalJSON should not fail for valid compare",
		)

		marshaledString := string(marshaledBytes)

		// Act — unmarshal into a separate variable
		var target corecomparator.Compare

		unmarshalErr := target.UnmarshalJSON([]byte(unmarshalInput))
		errcore.HandleErrMessage(
			unmarshalErr,
			"UnmarshalJSON should not fail for valid input",
		)

		name := target.Name()
		numberString := target.NumberString()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			marshaledString,
			name,
			numberString,
		)
	}
}

// Test_Compare_OnlySupportedErr verifies that OnlySupportedErr returns an error
// when the compare value is not in the supported list, and nil when it is.
// Migrated from cmd/main/compareEnumTesting02.go.
func Test_Compare_OnlySupportedErr(t *testing.T) {
	for caseIndex, testCase := range onlySupportedErrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		value, err := input.GetAsInt("value")
		errcore.HandleErrMessage(
			err,
			"value is required for OnlySupportedErr test",
		)

		message, err := input.GetAsString("message")
		errcore.HandleErrMessage(
			err,
			"message is required for OnlySupportedErr test",
		)

		supportedRaw, hasSupported := input["supported"]
		if !hasSupported {
			errcore.HandleErrMessage(
				fmt.Errorf("supported key missing"),
				"supported is required for OnlySupportedErr test",
			)
		}

		supportedInts := supportedRaw.([]int)
		supportedCompares := make(
			[]corecomparator.Compare,
			len(supportedInts),
		)

		for i, s := range supportedInts {
			supportedCompares[i] = corecomparator.Compare(s)
		}

		compare := corecomparator.Compare(value)

		// Act
		resultErr := compare.OnlySupportedErr(
			message,
			supportedCompares...,
		)

		hasError := fmt.Sprintf("%v", resultErr != nil)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			hasError,
		)
	}
}
