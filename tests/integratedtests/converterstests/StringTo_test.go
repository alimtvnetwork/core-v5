package converterstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_StringTo_Integer_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToIntegerTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.Integer(inputStr)
		actual := args.Map{
			"value":    fmt.Sprintf("%v", value),
			"hasError": fmt.Sprintf("%v", err != nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"value":     fmt.Sprintf("%v", value),
			"isSuccess": fmt.Sprintf("%v", isSuccess),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringTo_Float64_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToFloat64TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.Float64(inputStr)
		actual := args.Map{
			"value":    fmt.Sprintf("%v", value),
			"hasError": fmt.Sprintf("%v", err != nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringTo_Byte_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		value, err := converters.StringTo.Byte(inputStr)
		actual := args.Map{
			"value":    fmt.Sprintf("%v", value),
			"hasError": fmt.Sprintf("%v", err != nil),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
			result = converters.BytesTo.PtrString(bytes)
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
		actual := args.Map{
			"count": len(hashset),
		}

		// Check all values are true
		if len(hashset) > 0 {
			allTrue := true
			for _, v := range hashset {
				if !v {
					allTrue = false
					break
				}
			}
			actual["allTrue"] = allTrue
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
