package corecomparatortests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// Test_Compare_JsonRoundtrip verifies MarshalJSON produces the correct name
// string and UnmarshalJSON restores identity including Name() and NumberString().
// Migrated from cmd/main/enumTesting.go.
func Test_Compare_JsonRoundtrip(t *testing.T) {
	for caseIndex, testCase := range compareJsonRoundtripTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		value, ok := input.GetAsInt("value")
		if !ok {
			errcore.HandleErrMessage("value is required for compare JSON roundtrip test")
		}

		unmarshalInput, ok := input.GetAsString("unmarshalInput")
		if !ok {
			errcore.HandleErrMessage("unmarshalInput is required for compare JSON roundtrip test")
		}

		compare := corecomparator.Compare(value)

		// Act — marshal
		marshaledBytes, marshalErr := compare.MarshalJSON()
		errcore.SimpleHandleErr(
			marshalErr,
			"MarshalJSON should not fail for valid compare",
		)

		marshaledString := string(marshaledBytes)

		// Act — unmarshal into a separate variable
		var target corecomparator.Compare

		unmarshalErr := target.UnmarshalJSON([]byte(unmarshalInput))
		errcore.SimpleHandleErr(
			unmarshalErr,
			"UnmarshalJSON should not fail for valid input",
		)

		name := target.Name()
		numberString := target.NumberString()

		// Assert
		actual := args.Map{
			"marshaledJson":    marshaledString,
			"unmarshaledName":  name,
			"unmarshaledValue": numberString,
		}

		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
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

		value, ok := input.GetAsInt("value")
		if !ok {
			errcore.HandleErrMessage("value is required for OnlySupportedErr test")
		}

		message, ok := input.GetAsString("message")
		if !ok {
			errcore.HandleErrMessage("message is required for OnlySupportedErr test")
		}

		supportedRaw, hasSupported := input["supported"]
		if !hasSupported {
			errcore.HandleErrMessage("supported is required for OnlySupportedErr test")
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
