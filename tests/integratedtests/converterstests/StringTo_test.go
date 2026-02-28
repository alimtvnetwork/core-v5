package converterstests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_StringTo_Integer_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToIntegerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.Integer(inputStr)
		valueStr := fmt.Sprintf("%v", value)
		hasError := fmt.Sprintf("%v", err != nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, valueStr, hasError)
	}
}

func Test_BytesTo_String_Verification(t *testing.T) {
	for caseIndex, testCase := range bytesToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		inputBytes := []byte(inputStr)

		// Act
		result := converters.BytesTo.String(inputBytes)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_StringTo_IntegerWithDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToIntegerWithDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		defaultInt, _ := input.GetAsInt("defaultInt")

		// Act
		value, isSuccess := converters.StringTo.IntegerWithDefault(inputStr, defaultInt)

		// Assert
		testCase.ShouldBeEqual(
			t, caseIndex,
			fmt.Sprintf("%v", value),
			fmt.Sprintf("%v", isSuccess),
		)
	}
}

func Test_StringTo_Float64_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToFloat64TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.Float64(inputStr)
		hasError := fmt.Sprintf("%v", err != nil)

		// Assert
		testCase.ShouldBeEqual(
			t, caseIndex,
			fmt.Sprintf("%v", value),
			hasError,
		)
	}
}

func Test_StringTo_Byte_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.Byte(inputStr)

		// Assert
		testCase.ShouldBeEqual(
			t, caseIndex,
			fmt.Sprintf("%v", value),
			fmt.Sprintf("%v", err != nil),
		)
	}
}

func Test_BytesTo_PtrString_Verification(t *testing.T) {
	for caseIndex, testCase := range bytesToPtrStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var result string
		if isNil {
			result = converters.BytesTo.PtrString(nil)
		} else {
			bytes := []byte(inputStr)
			result = converters.BytesTo.PtrString(&bytes)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_StringsTo_Hashset_Verification(t *testing.T) {
	for caseIndex, testCase := range stringsToHashsetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, _ := input.Get("input")
		inputSlice := inputRaw.([]string)

		// Act
		hashset := converters.StringsTo.Hashset(inputSlice)
		length := fmt.Sprintf("%v", len(hashset))

		actLines := []string{length}
		for _, s := range inputSlice {
			if _, exists := hashset[s]; exists {
				// only add unique check
				found := false
				for _, existing := range actLines[1:] {
					if existing == fmt.Sprintf("%v", hashset[s]) {
						found = true
						break
					}
				}
				if !found {
					actLines = append(actLines, fmt.Sprintf("%v", hashset[s]))
				}
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_StringTo_IntegerDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToIntegerDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value := converters.StringTo.IntegerDefault(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", value))
	}
}

func Test_StringTo_Bool_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value := converters.StringTo.Bool(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", value))
	}
}

func Test_StringTo_UnsignedInteger_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToUnsignedIntegerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.UnsignedInteger(inputStr)

		// Assert
		testCase.ShouldBeEqual(
			t, caseIndex,
			fmt.Sprintf("%v", value),
			fmt.Sprintf("%v", err != nil),
		)
	}
}
