package bytetypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_GetSet_Verification(t *testing.T) {
	for caseIndex, testCase := range getSetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		condition, _ := input.GetAsBool("condition")
		trueVal, _ := input.GetAsInt("trueValue")
		falseVal, _ := input.GetAsInt("falseValue")

		// Act
		result := bytetype.GetSet(
			condition,
			bytetype.New(byte(trueVal)),
			bytetype.New(byte(falseVal)),
		)

		actual := args.Map{
			"result": result.ValueInt(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_GetSetVariant_Verification(t *testing.T) {
	for caseIndex, testCase := range getSetVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		condition, _ := input.GetAsBool("condition")
		trueVal, _ := input.GetAsInt("trueValue")
		falseVal, _ := input.GetAsInt("falseValue")

		// Act
		result := bytetype.GetSetVariant(
			condition,
			byte(trueVal),
			byte(falseVal),
		)

		actual := args.Map{
			"result": result.ValueInt(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_Comparisons(t *testing.T) {
	for caseIndex, testCase := range comparisonTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"isEqual3":        v.IsEqual(3),
			"isEqual5":        v.IsEqual(5),
			"isGreater3":      v.IsGreater(3),
			"isGreater7":      v.IsGreater(7),
			"isGreaterEqual5": v.IsGreaterEqual(5),
			"isLess3":         v.IsLess(3),
			"isLess7":         v.IsLess(7),
			"isLessEqual5":    v.IsLessEqual(5),
			"isBetween3and7":  v.IsBetween(3, 7),
			"isBetween6and8":  v.IsBetween(6, 8),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_String_Conversion(t *testing.T) {
	for caseIndex, testCase := range stringConversionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := bytetype.String([]byte(inputStr))

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_Methods(t *testing.T) {
	for caseIndex, testCase := range variantMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"isZero":    v.IsZero(),
			"isOne":     v.IsOne(),
			"isTwo":     v.IsTwo(),
			"isThree":  v.IsThree(),
			"isMin":     v.IsMin(),
			"isValid":   v.IsValid(),
			"isInvalid": v.IsInvalid(),
			"valueInt":  v.ValueInt(),
			"valueByte": int(v.ValueByte()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_Arithmetic(t *testing.T) {
	for caseIndex, testCase := range variantArithmeticTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		base, _ := input.GetAsInt("base")
		n, _ := input.GetAsInt("n")
		v := bytetype.New(byte(base))

		// Act
		actual := args.Map{
			"addResult":      v.Add(byte(n)).ValueInt(),
			"subtractResult": v.Subtract(byte(n)).ValueInt(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IntComparisons(t *testing.T) {
	// Arrange
	v := bytetype.New(5)

	// Act & Assert
	if !v.IsEqualInt(5) {
		t.Error("expected IsEqualInt(5) true")
	}
	if v.IsEqualInt(3) {
		t.Error("expected IsEqualInt(3) false")
	}
	if !v.IsGreaterInt(3) {
		t.Error("expected IsGreaterInt(3) true")
	}
	if !v.IsGreaterEqualInt(5) {
		t.Error("expected IsGreaterEqualInt(5) true")
	}
	if !v.IsLessInt(7) {
		t.Error("expected IsLessInt(7) true")
	}
	if !v.IsLessEqualInt(5) {
		t.Error("expected IsLessEqualInt(5) true")
	}
	if !v.IsBetweenInt(3, 7) {
		t.Error("expected IsBetweenInt(3,7) true")
	}
}

func Test_Variant_Is(t *testing.T) {
	// Arrange
	v := bytetype.New(3)

	// Act & Assert
	if !v.Is(bytetype.Three) {
		t.Error("expected Is(Three) true")
	}
	if v.Is(bytetype.One) {
		t.Error("expected Is(One) false")
	}
}

func Test_Variant_HasIndexInStrings(t *testing.T) {
	// Arrange
	v := bytetype.New(1)
	items := []string{"zero", "one", "two"}

	// Act
	val, isValid := v.HasIndexInStrings(items...)

	// Assert
	if !isValid {
		t.Error("expected isValid true")
	}
	if val != "one" {
		t.Errorf("expected 'one', got '%s'", val)
	}

	// Out of range
	v2 := bytetype.New(10)
	_, isValid2 := v2.HasIndexInStrings(items...)
	if isValid2 {
		t.Error("expected isValid false for out of range")
	}

	// Empty slice
	_, isValid3 := v.HasIndexInStrings()
	if isValid3 {
		t.Error("expected isValid false for empty slice")
	}
}

func Test_Variant_ValueConversions(t *testing.T) {
	// Arrange
	v := bytetype.New(42)

	// Act & Assert
	if v.ValueUInt16() != 42 {
		t.Errorf("expected ValueUInt16=42, got %d", v.ValueUInt16())
	}
	if v.ValueInt8() != 42 {
		t.Errorf("expected ValueInt8=42, got %d", v.ValueInt8())
	}
	if v.ValueInt16() != 42 {
		t.Errorf("expected ValueInt16=42, got %d", v.ValueInt16())
	}
	if v.ValueInt32() != 42 {
		t.Errorf("expected ValueInt32=42, got %d", v.ValueInt32())
	}
	if v.ValueString() != "42" {
		t.Errorf("expected ValueString='42', got '%s'", v.ValueString())
	}
	if v.ToNumberString() != "42" {
		t.Errorf("expected ToNumberString='42', got '%s'", v.ToNumberString())
	}
	if v.StringValue() != "42" {
		t.Errorf("expected StringValue='42', got '%s'", v.StringValue())
	}
}

func Test_Variant_IsValueEqual(t *testing.T) {
	// Arrange
	v := bytetype.New(5)

	// Act & Assert
	if !v.IsValueEqual(5) {
		t.Error("expected IsValueEqual(5) true")
	}
	if v.IsValueEqual(3) {
		t.Error("expected IsValueEqual(3) false")
	}
}

func Test_Variant_ToPtr(t *testing.T) {
	// Arrange
	v := bytetype.New(7)

	// Act
	ptr := v.ToPtr()

	// Assert
	if ptr == nil {
		t.Error("expected ToPtr to not be nil")
	}
	if ptr.ValueInt() != 7 {
		t.Errorf("expected ptr value 7, got %d", ptr.ValueInt())
	}
}

func Test_Variant_IsNameEqual(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	name := v.Name()

	// Assert
	if !v.IsNameEqual(name) {
		t.Errorf("expected IsNameEqual('%s') true", name)
	}
	if v.IsNameEqual("NonExistent") {
		t.Error("expected IsNameEqual('NonExistent') false")
	}
}

func Test_Variant_IsAnyNamesOf(t *testing.T) {
	// Arrange
	v := bytetype.One
	name := v.Name()

	// Act & Assert
	if !v.IsAnyNamesOf(name, "Other") {
		t.Error("expected IsAnyNamesOf to find match")
	}
	if v.IsAnyNamesOf("None", "Other") {
		t.Error("expected IsAnyNamesOf to not match")
	}
}
