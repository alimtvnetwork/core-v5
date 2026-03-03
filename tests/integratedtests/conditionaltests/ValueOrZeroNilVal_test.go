package conditionaltests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_ValueOrZero_String_Verification(t *testing.T) {
	for caseIndex, testCase := range valueOrZeroStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.ValueOrZero[string](ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_ValueOrZero_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range valueOrZeroIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		result := conditional.ValueOrZero[int](ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_ValueOrZero_Bool_Verification(t *testing.T) {
	for caseIndex, testCase := range valueOrZeroBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *bool
		if !isNil {
			val, _ := input.Get("value")
			boolVal := val == true
			ptr = &boolVal
		}

		result := conditional.ValueOrZero[bool](ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_PtrOrZero_String_Verification(t *testing.T) {
	for caseIndex, testCase := range ptrOrZeroStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.PtrOrZero[string](ptr)
		isNotNil := fmt.Sprintf("%v", result != nil)
		derefVal := *result

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			isNotNil,
			derefVal,
		)
	}
}

func Test_PtrOrZero_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range ptrOrZeroIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		result := conditional.PtrOrZero[int](ptr)
		isNotNil := fmt.Sprintf("%v", result != nil)
		derefVal := fmt.Sprintf("%v", *result)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			isNotNil,
			derefVal,
		)
	}
}

func Test_NilVal_String_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.NilVal[string](ptr, onNil, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_NilVal_Int_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsInt("onNil")
		onNonNil, _ := input.GetAsInt("onNonNil")

		// Act
		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		result := conditional.NilVal[int](ptr, onNil, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_NilValPtr_String_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValPtrStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		// Act
		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		result := conditional.NilValPtr[string](ptr, onNil, onNonNil)
		isNotNil := fmt.Sprintf("%v", result != nil)
		derefVal := *result

		// Assert
		testCase.ShouldBeEqual(t, caseIndex,
			isNotNil,
			derefVal,
		)
	}
}
