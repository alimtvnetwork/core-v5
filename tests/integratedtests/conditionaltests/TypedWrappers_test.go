package conditionaltests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/coretests/args"
)

// ============================================================================
// If<Type> wrappers
// ============================================================================

func Test_IfBool_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifBoolTypedTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfBool(isTrue, trueVal.(bool), falseVal.(bool))

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IfInt_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifIntTypedTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsInt("trueValue")
		falseVal, _ := input.GetAsInt("falseValue")

		result := conditional.IfInt(isTrue, trueVal, falseVal)

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IfByte_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifByteTypedTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfByte(isTrue, trueVal.(byte), falseVal.(byte))

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IfFloat64_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifFloat64TypedTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfFloat64(isTrue, trueVal.(float64), falseVal.(float64))

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IfAny_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifAnyTypedTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfAny(isTrue, trueVal, falseVal)

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ============================================================================
// IfFunc<Type> wrappers
// ============================================================================

func Test_IfFuncBool_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifFuncBoolTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfFuncBool(
			isTrue,
			func() bool { return trueVal.(bool) },
			func() bool { return falseVal.(bool) },
		)

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IfFuncInt_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifFuncIntTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsInt("trueValue")
		falseVal, _ := input.GetAsInt("falseValue")

		result := conditional.IfFuncInt(
			isTrue,
			func() int { return trueVal },
			func() int { return falseVal },
		)

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IfFuncString_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifFuncStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsString("trueValue")
		falseVal, _ := input.GetAsString("falseValue")

		// Act
		result := conditional.IfFuncString(
			isTrue,
			func() string { return trueVal },
			func() string { return falseVal },
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_IfFuncAny_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifFuncAnyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		// Act
		result := conditional.IfFuncAny(
			isTrue,
			func() any { return trueVal },
			func() any { return falseVal },
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ============================================================================
// IfTrueFunc<Type> wrappers
// ============================================================================

func Test_IfTrueFuncBool_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifTrueFuncBoolTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")

		result := conditional.IfTrueFuncBool(
			isTrue,
			func() bool { return trueVal.(bool) },
		)

		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IfTrueFuncString_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifTrueFuncStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsString("trueValue")

		// Act
		result := conditional.IfTrueFuncString(
			isTrue,
			func() string { return trueVal },
		)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_IfTrueFuncStrings_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifTrueFuncStringsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")

		// Act
		result := conditional.IfTrueFuncStrings(
			isTrue,
			func() []string { return trueVal.([]string) },
		)

		// Assert
		if result == nil {
			actual := args.Map{
				"length": fmt.Sprintf("%v", 0),
				"isNil":  "true",
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			actual := args.Map{
				"length": fmt.Sprintf("%v", len(result)),
				"first":  result[0],
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

func Test_IfTrueFuncBytes_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifTrueFuncBytesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")

		// Act
		result := conditional.IfTrueFuncBytes(
			isTrue,
			func() []byte { return trueVal.([]byte) },
		)

		// Assert
		if result == nil {
			actual := args.Map{
				"length": fmt.Sprintf("%v", 0),
				"isNil":  "true",
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			actual := args.Map{
				"length": fmt.Sprintf("%v", len(result)),
				"first":  fmt.Sprintf("%v", result[0]),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

// ============================================================================
// IfSlice<Type> wrappers
// ============================================================================

func Test_IfSliceBool_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifSliceBoolTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfSliceBool(isTrue, trueVal.([]bool), falseVal.([]bool))
		actual := args.Map{
			"length": fmt.Sprintf("%v", len(result)),
			"first":  fmt.Sprintf("%v", result[0]),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfSliceInt_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifSliceIntTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfSliceInt(isTrue, trueVal.([]int), falseVal.([]int))
		actual := args.Map{
			"length": fmt.Sprintf("%v", len(result)),
			"first":  fmt.Sprintf("%v", result[0]),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfSliceString_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifSliceStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		// Act
		result := conditional.IfSliceString(isTrue, trueVal.([]string), falseVal.([]string))

		// Assert
		actual := args.Map{
			"length": fmt.Sprintf("%v", len(result)),
			"first":  result[0],
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfSliceByte_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifSliceByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		// Act
		result := conditional.IfSliceByte(isTrue, trueVal.([]byte), falseVal.([]byte))

		// Assert
		actual := args.Map{
			"length": fmt.Sprintf("%v", len(result)),
			"first":  fmt.Sprintf("%v", result[0]),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfSliceAny_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifSliceAnyTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")

		result := conditional.IfSliceAny(isTrue, trueVal.([]any), falseVal.([]any))
		actual := args.Map{
			"length": fmt.Sprintf("%v", len(result)),
			"first":  fmt.Sprintf("%v", result[0]),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// IfPtr<Type> wrappers
// ============================================================================

func Test_IfPtrString_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifPtrStringTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueStr, _ := input.GetAsString("trueValue")
		falseStr, _ := input.GetAsString("falseValue")

		result := conditional.IfPtrString(isTrue, &trueStr, &falseStr)
		actual := args.Map{
			"isNotNil": fmt.Sprintf("%v", result != nil),
			"value":    *result,
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfPtrInt_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifPtrIntTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsInt("trueValue")
		falseVal, _ := input.GetAsInt("falseValue")

		result := conditional.IfPtrInt(isTrue, &trueVal, &falseVal)
		actual := args.Map{
			"isNotNil": fmt.Sprintf("%v", result != nil),
			"value":    fmt.Sprintf("%v", *result),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfPtrBool_Typed_Verification(t *testing.T) {
	for caseIndex, testCase := range ifPtrBoolTestCases {
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.Get("trueValue")
		falseVal, _ := input.Get("falseValue")
		tv := trueVal.(bool)
		fv := falseVal.(bool)

		result := conditional.IfPtrBool(isTrue, &tv, &fv)
		actual := args.Map{
			"isNotNil": fmt.Sprintf("%v", result != nil),
			"value":    fmt.Sprintf("%v", *result),
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
