package bytetypetests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Variant_IsCompareResult(t *testing.T) {
	compareNameToEnum := map[string]corecomparator.Compare{
		"Equal":             corecomparator.Equal,
		"LeftGreater":       corecomparator.LeftGreater,
		"LeftGreaterEqual":  corecomparator.LeftGreaterEqual,
		"LeftLess":          corecomparator.LeftLess,
		"LeftLessEqual":     corecomparator.LeftLessEqual,
		"NotEqual":          corecomparator.NotEqual,
	}

	for caseIndex, testCase := range extIsCompareResultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		n, _ := input.GetAsInt("n")
		compareName, _ := input.GetAsString("compare")
		v := bytetype.New(byte(val))
		compareEnum := compareNameToEnum[compareName]

		// Act
		result := v.IsCompareResult(byte(n), compareEnum)

		actual := args.Map{
			"result": result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsCompareResult_Panic(t *testing.T) {
	// Arrange
	v := bytetype.New(5)

	// Act & Assert
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic for out-of-range comparator")
		}
	}()

	v.IsCompareResult(3, corecomparator.Compare(99))
}

func Test_Variant_EnumMethods(t *testing.T) {
	for caseIndex, testCase := range extEnumMethodsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"name":           v.Name(),
			"nameValue":      v.NameValue(),
			"typeName":       v.TypeName(),
			"isValidRange":   v.IsValidRange(),
			"isInvalidRange": v.IsInvalidRange(),
			"stringValue":    v.StringValue(),
			"rangeNamesCsv":  v.RangeNamesCsv(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsMax(t *testing.T) {
	for caseIndex, testCase := range extIsMaxTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		v := bytetype.New(byte(val))

		// Act
		actual := args.Map{
			"isMax": v.IsMax(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsEnumEqual(t *testing.T) {
	for caseIndex, testCase := range extIsEnumEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("value")
		other, _ := input.GetAsInt("other")
		v := bytetype.New(byte(val))
		otherVariant := bytetype.New(byte(other))

		// Act
		// IsEnumEqual takes enuminf.BasicEnumer which requires pointer receiver
		// (UnmarshalJSON has pointer receiver), so pass &otherVariant.
		actual := args.Map{
			"isEnumEqual": v.IsEnumEqual(&otherVariant),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variant_IsAnyEnumsEqual(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	// IsAnyEnumsEqual takes ...enuminf.BasicEnumer which requires pointer receiver.
	// bytetype.Variant has UnmarshalJSON on pointer receiver, so use variables + &.
	two := bytetype.Two
	one := bytetype.One
	three := bytetype.Three

	if !v.IsAnyEnumsEqual(&two, &one) {
		t.Error("expected IsAnyEnumsEqual to find match")
	}

	if v.IsAnyEnumsEqual(&two, &three) {
		t.Error("expected IsAnyEnumsEqual to not match")
	}

	if v.IsAnyEnumsEqual() {
		t.Error("expected IsAnyEnumsEqual with no args to return false")
	}
}

func Test_Variant_MarshalJSON(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	data, err := json.Marshal(v)

	// Assert
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	if len(data) == 0 {
		t.Error("MarshalJSON returned empty bytes")
	}
}

func Test_Variant_UnmarshalJSON(t *testing.T) {
	// Arrange
	v := bytetype.One
	data, _ := json.Marshal(v.Value())

	// Act
	var result bytetype.Variant
	err := json.Unmarshal(data, &result)

	// Assert
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}

	if result.Value() != v.Value() {
		t.Errorf("expected value %d, got %d", v.Value(), result.Value())
	}
}

func Test_Variant_UnmarshallToValue(t *testing.T) {
	// Arrange
	v := bytetype.Two
	data, _ := json.Marshal(v)

	// Act
	val, err := v.UnmarshallToValue(data)

	// Assert
	if err != nil {
		t.Errorf("UnmarshallToValue error: %v", err)
	}

	if val != v.Value() {
		t.Errorf("expected %d, got %d", v.Value(), val)
	}
}

func Test_Variant_AllNameValues(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	names := v.AllNameValues()

	// Assert
	if len(names) == 0 {
		t.Error("AllNameValues should not be empty")
	}
}

func Test_Variant_IntegerEnumRanges(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	ranges := v.IntegerEnumRanges()

	// Assert
	if len(ranges) == 0 {
		t.Error("IntegerEnumRanges should not be empty")
	}
}

func Test_Variant_MinMaxAny(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	min, max := v.MinMaxAny()

	// Assert
	if min == nil {
		t.Error("min should not be nil")
	}

	if max == nil {
		t.Error("max should not be nil")
	}
}

func Test_Variant_MinMaxStrings(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	if v.MinValueString() == "" {
		t.Error("MinValueString should not be empty")
	}

	if v.MaxValueString() == "" {
		t.Error("MaxValueString should not be empty")
	}

	if v.MinInt() < 0 {
		t.Error("MinInt should be >= 0")
	}

	if v.MaxInt() <= 0 {
		t.Error("MaxInt should be > 0")
	}
}

func Test_Variant_RangesDynamicMap(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	m := v.RangesDynamicMap()

	// Assert
	if len(m) == 0 {
		t.Error("RangesDynamicMap should not be empty")
	}
}

func Test_Variant_Format(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	formatted := v.Format("{name}")

	// Assert
	if formatted == "" {
		t.Error("Format should return non-empty")
	}
}

func Test_Variant_StringRanges(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	if len(v.StringRanges()) == 0 {
		t.Error("StringRanges should not be empty")
	}

	if len(v.StringRangesPtr()) == 0 {
		t.Error("StringRangesPtr should not be empty")
	}
}

func Test_Variant_RangesInvalid(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act & Assert
	if v.RangesInvalidMessage() == "" {
		t.Error("RangesInvalidMessage should not be empty")
	}

	if v.RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should not be nil")
	}
}

func Test_Variant_EnumType(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	enumType := v.EnumType()

	// Assert
	if enumType == nil {
		t.Error("EnumType should not be nil")
	}
}

func Test_Variant_AsBasicEnumContractsBinder(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	binder := v.AsBasicEnumContractsBinder()

	// Assert
	if binder == nil {
		t.Error("AsBasicEnumContractsBinder should not be nil")
	}
}

func Test_Variant_JsonString(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	js := v.JsonString()

	// Assert
	if js == "" {
		t.Error("JsonString should not be empty")
	}
}

func Test_Variant_OnlySupportedErr(t *testing.T) {
	// Arrange
	v := bytetype.One
	// OnlySupportedErr compares against StringRanges() (plain names like "Zero"),
	// not AllNameValues() which returns "Name(Value)" format like "Zero(0)".
	allNames := v.StringRanges()

	// Act
	err := v.OnlySupportedErr(allNames...)

	// Assert
	if err != nil {
		t.Errorf("all names supported should not error, got: %v", err)
	}
}

func Test_Variant_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	v := bytetype.One

	// Act
	err := v.OnlySupportedMsgErr("test message", "NonExistent")

	// Assert
	if err == nil {
		t.Error("OnlySupportedMsgErr with single unsupported name should return error")
	}
}
