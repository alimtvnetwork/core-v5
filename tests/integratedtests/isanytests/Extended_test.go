package isanytests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

func Test_NotNull_Verification(t *testing.T) {
	for caseIndex, tc := range notNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := isany.NotNull(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_AllNull_Verification(t *testing.T) {
	for caseIndex, tc := range allNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AllNull(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_AnyNull_Verification(t *testing.T) {
	for caseIndex, tc := range anyNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AnyNull(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Zero_Verification(t *testing.T) {
	for caseIndex, tc := range zeroTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := isany.Zero(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_AllZero_Verification(t *testing.T) {
	for caseIndex, tc := range allZeroTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AllZero(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_AnyZero_Verification(t *testing.T) {
	for caseIndex, tc := range anyZeroTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AnyZero(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Pointer_Verification(t *testing.T) {
	for caseIndex, tc := range pointerTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		usePtrRaw, _ := input.Get("usePointer")
		usePtr := usePtrRaw.(bool)

		// Act
		var result bool
		if usePtr {
			val := 42
			result = isany.Pointer(&val)
		} else {
			inputVal, _ := input.Get("input")
			result = isany.Pointer(inputVal)
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Function_Verification(t *testing.T) {
	for caseIndex, tc := range functionTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		useFuncRaw, _ := input.Get("useFunc")
		useFunc := useFuncRaw.(bool)

		// Act
		var result bool
		if useFunc {
			result = isany.Function(func() {})
		} else {
			inputVal, _ := input.Get("input")
			result = isany.Function(inputVal)
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_TypeSame_Verification(t *testing.T) {
	for caseIndex, tc := range typeSameTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.Get("left")
		right, _ := input.Get("right")

		// Act
		result := isany.TypeSame(left, right)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_StringEqual_Verification(t *testing.T) {
	for caseIndex, tc := range stringEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		result := isany.StringEqual(left, right)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_DefinedAllOf_Verification(t *testing.T) {
	for caseIndex, tc := range definedAllOfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.DefinedAllOf(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_DefinedAnyOf_Verification(t *testing.T) {
	for caseIndex, tc := range definedAnyOfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.DefinedAnyOf(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_ReflectNull_Verification(t *testing.T) {
	for caseIndex, tc := range reflectNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		useNilPtrRaw, _ := input.Get("useNilPtr")
		useNilPtr := useNilPtrRaw.(bool)

		// Act
		var result bool
		if useNilPtr {
			var p *int
			rv := reflect.ValueOf(&p).Elem()
			result = isany.ReflectNull(rv)
		} else {
			val := 42
			rv := reflect.ValueOf(&val)
			result = isany.ReflectNull(rv)
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_DefinedLeftRight_Verification(t *testing.T) {
	for caseIndex, tc := range definedLeftRightTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		leftNilRaw, hasLeftNil := input.Get("leftNil")

		var left, right any
		if hasLeftNil && leftNilRaw == true {
			left = nil
		} else {
			left, _ = input.Get("left")
		}
		right, _ = input.Get("right")

		// Act
		leftDef, rightDef := isany.DefinedLeftRight(left, right)

		// Assert
		actual := args.Map{
			"leftDefined":  fmt.Sprintf("%v", leftDef),
			"rightDefined": fmt.Sprintf("%v", rightDef),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_DefinedItems_Verification(t *testing.T) {
	for caseIndex, tc := range definedItemsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.DefinedItems(inputs...)

		// Assert
		actual := args.Map{
			"count": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NotDeepEqual_Verification(t *testing.T) {
	for caseIndex, tc := range notDeepEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.Get("left")
		right, _ := input.Get("right")

		// Act
		result := isany.NotDeepEqual(left, right)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Conclusive_Verification(t *testing.T) {
	for caseIndex, tc := range conclusiveTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")
		compare := corecomparator.Compare(value)

		// Act
		result := isany.Conclusive(compare)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_FuncOnly_Verification(t *testing.T) {
	for caseIndex, tc := range funcOnlyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		useFuncRaw, _ := input.Get("useFunc")
		useFunc := useFuncRaw.(bool)

		// Act
		var result bool
		if useFunc {
			result = isany.FuncOnly(func() {})
		} else {
			result = isany.FuncOnly(42)
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
