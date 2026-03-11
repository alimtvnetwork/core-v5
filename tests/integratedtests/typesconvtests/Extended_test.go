package typesconvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/typesconv"
)

func Test_BoolPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range boolPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input.GetAsBoolDefault("value", false)

		// Act
		result := typesconv.BoolPtr(val)

		actual := args.Map{"notNil": result != nil, "deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BoolPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range boolPtrToSimpleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrToSimple(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BoolPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range boolPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)
		defVal := input.GetAsBoolDefault("defVal", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrToSimpleDef(ptr, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BoolPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range boolPtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)
		defVal := input.GetAsBoolDefault("defVal", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrToDefPtr(ptr, defVal)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BoolPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range boolPtrDefValFuncTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrDefValFunc(ptr, func() bool { return true })

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range intPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input.GetAsIntDefault("value", 0)

		// Act
		result := typesconv.IntPtr(val)

		actual := args.Map{"notNil": result != nil, "deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range intPtrToSimpleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *int
		if !isNil {
			val := input.GetAsIntDefault("value", 0)
			ptr = &val
		}

		// Act
		result := typesconv.IntPtrToSimple(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range intPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal := input.GetAsIntDefault("defVal", 0)

		// Act
		result := typesconv.IntPtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range intPtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal := input.GetAsIntDefault("defVal", 0)

		// Act
		result := typesconv.IntPtrToDefPtr(nil, defVal)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range intPtrDefValFuncTestCases {
		// Arrange
		_ = testCase.ArrangeInput

		// Act
		result := typesconv.IntPtrDefValFunc(nil, func() int { return 55 })

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range stringPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")

		// Act
		result := typesconv.StringPtr(val)

		actual := args.Map{"notNil": result != nil, "deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range stringPtrToSimpleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := typesconv.StringPtrToSimple(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range stringPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal, _ := input.GetAsString("defVal")

		// Act
		result := typesconv.StringPtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range stringPtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal, _ := input.GetAsString("defVal")

		// Act
		result := typesconv.StringPtrToDefPtr(nil, defVal)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range stringPtrDefValFuncTestCases {
		// Arrange
		_ = testCase.ArrangeInput

		// Act
		result := typesconv.StringPtrDefValFunc(nil, func() string { return "generated" })

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringToBool_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")

		// Act
		result := typesconv.StringToBool(val)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPointerToBool_Verification(t *testing.T) {
	for caseIndex, testCase := range stringPointerToBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := typesconv.StringPointerToBool(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPointerToBoolPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range stringPointerToBoolPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := typesconv.StringPointerToBoolPtr(ptr)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringToBoolPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range stringToBoolPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")

		// Act
		result := typesconv.StringToBoolPtr(val)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range bytePtrTestCases {
		// Arrange
		result := typesconv.BytePtr(5)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytePtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range bytePtrToSimpleTestCases {
		// Arrange - nil

		// Act
		result := typesconv.BytePtrToSimple(nil)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytePtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range bytePtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawDef, _ := input.Get("defVal")
		defVal := rawDef.(byte)

		// Act
		result := typesconv.BytePtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytePtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range bytePtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawDef, _ := input.Get("defVal")
		defVal := rawDef.(byte)

		// Act
		result := typesconv.BytePtrToDefPtr(nil, defVal)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytePtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range bytePtrDefValFuncTestCases {
		// Arrange - nil

		// Act
		result := typesconv.BytePtrDefValFunc(nil, func() byte { return 7 })

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FloatPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range floatPtrTestCases {
		// Arrange

		// Act
		result := typesconv.FloatPtr(3.14)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FloatPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range floatPtrToSimpleTestCases {
		// Arrange - nil

		// Act
		result := typesconv.FloatPtrToSimple(nil)

		actual := args.Map{"isZero": result == 0}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FloatPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range floatPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawDef, _ := input.Get("defVal")
		defVal := rawDef.(float32)

		// Act
		result := typesconv.FloatPtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FloatPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range floatPtrToDefPtrTestCases {
		// Arrange

		// Act
		result := typesconv.FloatPtrToDefPtr(nil, 2.5)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FloatPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range floatPtrDefValFuncTestCases {
		// Arrange - nil

		// Act
		result := typesconv.FloatPtrDefValFunc(nil, func() float32 { return 9.9 })

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
