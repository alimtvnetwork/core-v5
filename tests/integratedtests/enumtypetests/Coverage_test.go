package enumtypetests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl/enumtype"
)

func Test_Variant_Constants(t *testing.T) {
	// Assert
	if enumtype.Invalid != 0 {
		t.Error("Invalid should be 0")
	}
	if enumtype.Boolean != 1 {
		t.Error("Boolean should be 1")
	}
	if enumtype.String != 11 {
		t.Error("String should be 11")
	}
}

func Test_Variant_Name(t *testing.T) {
	// Act & Assert
	if enumtype.Boolean.Name() != "Boolean" {
		t.Errorf("expected Boolean, got %s", enumtype.Boolean.Name())
	}
	if enumtype.Integer.String() != "Integer" {
		t.Error("Integer String mismatch")
	}
	if enumtype.Byte.NameValue() == "" {
		t.Error("NameValue should not be empty")
	}
}

func Test_Variant_TypeChecks(t *testing.T) {
	// Assert
	if !enumtype.Boolean.IsBoolean() {
		t.Error("Boolean.IsBoolean should be true")
	}
	if !enumtype.Byte.IsByte() {
		t.Error("Byte.IsByte should be true")
	}
	if !enumtype.UnsignedInteger16.IsUnsignedInteger16() {
		t.Error("UnsignedInteger16 check failed")
	}
	if !enumtype.UnsignedInteger32.IsUnsignedInteger32() {
		t.Error("UnsignedInteger32 check failed")
	}
	if !enumtype.UnsignedInteger64.IsUnsignedInteger64() {
		t.Error("UnsignedInteger64 check failed")
	}
	if !enumtype.Integer8.IsInteger8() {
		t.Error("Integer8 check failed")
	}
	if !enumtype.Integer16.IsInteger16() {
		t.Error("Integer16 check failed")
	}
	if !enumtype.Integer32.IsInteger32() {
		t.Error("Integer32 check failed")
	}
	if !enumtype.Integer64.IsInteger64() {
		t.Error("Integer64 check failed")
	}
	if !enumtype.Integer.IsInteger() {
		t.Error("Integer check failed")
	}
	if !enumtype.String.IsString() {
		t.Error("String check failed")
	}
}

func Test_Variant_IsNumber(t *testing.T) {
	if !enumtype.Integer.IsNumber() {
		t.Error("Integer should be number")
	}
	if enumtype.Boolean.IsNumber() {
		t.Error("Boolean should not be number")
	}
	if enumtype.String.IsNumber() {
		t.Error("String should not be number")
	}
}

func Test_Variant_IsAnyInteger(t *testing.T) {
	if !enumtype.Integer.IsAnyInteger() {
		t.Error("Integer should be any integer")
	}
	if enumtype.Byte.IsAnyInteger() {
		t.Error("Byte should not be any integer")
	}
}

func Test_Variant_IsAnyUnsignedNumber(t *testing.T) {
	if !enumtype.Byte.IsAnyUnsignedNumber() {
		t.Error("Byte should be unsigned")
	}
	if enumtype.Integer.IsAnyUnsignedNumber() {
		t.Error("Integer should not be unsigned")
	}
}

func Test_Variant_ValidInvalid(t *testing.T) {
	if enumtype.Invalid.IsValid() {
		t.Error("Invalid should not be valid")
	}
	if !enumtype.Invalid.IsInvalid() {
		t.Error("Invalid should be invalid")
	}
	if !enumtype.Boolean.IsValid() {
		t.Error("Boolean should be valid")
	}
	if enumtype.Boolean.IsInvalid() {
		t.Error("Boolean should not be invalid")
	}
}

func Test_Variant_ValueConversions(t *testing.T) {
	v := enumtype.Integer // 10

	if v.Value() != 10 {
		t.Error("Value mismatch")
	}
	if v.ValueByte() != 10 {
		t.Error("ValueByte mismatch")
	}
	if v.ValueInt() != 10 {
		t.Error("ValueInt mismatch")
	}
	if v.ValueInt8() != 10 {
		t.Error("ValueInt8 mismatch")
	}
	if v.ValueInt16() != 10 {
		t.Error("ValueInt16 mismatch")
	}
	if v.ValueInt32() != 10 {
		t.Error("ValueInt32 mismatch")
	}
	if v.ValueUInt16() != 10 {
		t.Error("ValueUInt16 mismatch")
	}
	if v.ValueString() == "" {
		t.Error("ValueString should not be empty")
	}
	if v.ToNumberString() == "" {
		t.Error("ToNumberString should not be empty")
	}
}

func Test_Variant_IsNameEqual(t *testing.T) {
	if !enumtype.Boolean.IsNameEqual("Boolean") {
		t.Error("IsNameEqual should be true for Boolean")
	}
	if enumtype.Boolean.IsNameEqual("String") {
		t.Error("IsNameEqual should be false")
	}
}

func Test_Variant_IsAnyNamesOf(t *testing.T) {
	if !enumtype.Boolean.IsAnyNamesOf("String", "Boolean") {
		t.Error("IsAnyNamesOf should find Boolean")
	}
	if enumtype.Boolean.IsAnyNamesOf("String", "Integer") {
		t.Error("IsAnyNamesOf should not find Boolean")
	}
}

func Test_Variant_TypeName(t *testing.T) {
	if enumtype.Boolean.TypeName() == "" {
		t.Error("TypeName should not be empty")
	}
}

func Test_Variant_RangeNamesCsv(t *testing.T) {
	if enumtype.Boolean.RangeNamesCsv() == "" {
		t.Error("RangeNamesCsv should not be empty")
	}
}

func Test_Variant_MinMaxAny(t *testing.T) {
	min, max := enumtype.Boolean.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny should not return nil")
	}
}

func Test_Variant_MinMaxStrings(t *testing.T) {
	if enumtype.Boolean.MinValueString() == "" {
		t.Error("MinValueString should not be empty")
	}
	if enumtype.Boolean.MaxValueString() == "" {
		t.Error("MaxValueString should not be empty")
	}
}

func Test_Variant_MinMaxInt(t *testing.T) {
	if enumtype.Boolean.MaxInt() != enumtype.String.ValueInt() {
		t.Error("MaxInt mismatch")
	}
	if enumtype.Boolean.MinInt() != enumtype.Invalid.ValueInt() {
		t.Error("MinInt mismatch")
	}
}

func Test_Variant_RangesDynamicMap(t *testing.T) {
	m := enumtype.Boolean.RangesDynamicMap()
	if len(m) == 0 {
		t.Error("RangesDynamicMap should not be empty")
	}
}

func Test_Variant_IntegerEnumRanges(t *testing.T) {
	ranges := enumtype.Boolean.IntegerEnumRanges()
	if len(ranges) == 0 {
		t.Error("IntegerEnumRanges should not be empty")
	}
}

func Test_Variant_EnumType(t *testing.T) {
	et := enumtype.Boolean.EnumType()
	if et == nil {
		t.Error("EnumType should not be nil")
	}
}

func Test_Variant_MarshalJSON(t *testing.T) {
	data, err := json.Marshal(enumtype.Boolean)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	if len(data) == 0 {
		t.Error("MarshalJSON should not be empty")
	}
}

func Test_Variant_UnmarshalJSON(t *testing.T) {
	var v enumtype.Variant
	err := json.Unmarshal([]byte(`"Boolean"`), &v)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}
	if v != enumtype.Boolean {
		t.Error("UnmarshalJSON should parse to Boolean")
	}
}

func Test_Variant_UnmarshalJSON_Invalid(t *testing.T) {
	var v enumtype.Variant
	err := json.Unmarshal([]byte(`""`), &v)
	if err == nil {
		t.Error("UnmarshalJSON should error on empty")
	}

	err = json.Unmarshal([]byte(`"NonExistent"`), &v)
	if err == nil {
		t.Error("UnmarshalJSON should error on nonexistent")
	}
}

func Test_Variant_Format_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Format should panic")
		}
	}()

	enumtype.Boolean.Format("{name}")
}
