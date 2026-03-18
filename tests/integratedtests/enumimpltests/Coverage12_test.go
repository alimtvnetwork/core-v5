package enumimpltests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coreimpl/enumimpl/enumtype"
)

// ══════════════════════════════════════════════════════════════════════════════
// Helper: create test enums for reuse
// ══════════════════════════════════════════════════════════════════════════════

type testByte byte

const testByteVal testByte = 0

func byteEnum() *enumimpl.BasicByte {
	return enumimpl.New.BasicByte.Default(
		testByteVal,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testInt8 int8

const testInt8Val testInt8 = 0

func int8Enum() *enumimpl.BasicInt8 {
	return enumimpl.New.BasicInt8.Default(
		testInt8Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testInt16 int16

const testInt16Val testInt16 = 0

func int16Enum() *enumimpl.BasicInt16 {
	return enumimpl.New.BasicInt16.Default(
		testInt16Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testInt32 int32

const testInt32Val testInt32 = 0

func int32Enum() *enumimpl.BasicInt32 {
	return enumimpl.New.BasicInt32.Default(
		testInt32Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testUInt16 uint16

const testUInt16Val testUInt16 = 0

func uint16Enum() *enumimpl.BasicUInt16 {
	return enumimpl.New.BasicUInt16.Default(
		testUInt16Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

func stringEnum() *enumimpl.BasicString {
	return enumimpl.New.BasicString.Create(
		"TestStringEnum",
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testNamer struct{ name string }

func (n testNamer) Name() string { return n.name }

// ══════════════════════════════════════════════════════════════════════════════
// BasicByte — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_BasicByte_IsAnyOf(t *testing.T) {
	e := byteEnum()

	if !e.IsAnyOf(0) {
		t.Fatal("empty variadic should return true")
	}

	if !e.IsAnyOf(1, 0, 1, 2) {
		t.Fatal("expected true")
	}

	if e.IsAnyOf(1, 0, 2) {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicByte_IsAnyNamesOf(t *testing.T) {
	e := byteEnum()

	if !e.IsAnyNamesOf(0, "Invalid", "Active") {
		t.Fatal("expected true")
	}

	if e.IsAnyNamesOf(0, "Active") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicByte_GetValueByString(t *testing.T) {
	e := byteEnum()
	_ = e.GetValueByString("Active")
}

func Test_Cov12_BasicByte_GetValueByName_AllBranches(t *testing.T) {
	e := byteEnum()

	// Direct key
	v, err := e.GetValueByName("Active")
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}

	// Wrapped key
	v, err = e.GetValueByName(`"Active"`)
	if err != nil {
		t.Fatal("expected no error for wrapped")
	}

	// Not found
	_, err = e.GetValueByName("NotExist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicByte_GetStringValue(t *testing.T) {
	e := byteEnum()
	s := e.GetStringValue(0)

	if s != "Invalid" {
		t.Fatal("expected Invalid")
	}
}

func Test_Cov12_BasicByte_ExpectingEnumValueError(t *testing.T) {
	e := byteEnum()

	// Matching
	err := e.ExpectingEnumValueError("Active", byte(1))
	if err != nil {
		t.Fatal("expected nil")
	}

	// Not matching
	err = e.ExpectingEnumValueError("Active", byte(0))
	if err == nil {
		t.Fatal("expected error")
	}

	// Invalid raw string
	err = e.ExpectingEnumValueError("NotExist", byte(0))
	if err == nil {
		t.Fatal("expected error for invalid raw")
	}
}

func Test_Cov12_BasicByte_Ranges(t *testing.T) {
	e := byteEnum()

	if len(e.Ranges()) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov12_BasicByte_Hashmap_HashmapPtr(t *testing.T) {
	e := byteEnum()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_Cov12_BasicByte_IsValidRange(t *testing.T) {
	e := byteEnum()

	if !e.IsValidRange(1) {
		t.Fatal("expected valid")
	}

	if e.IsValidRange(100) {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_BasicByte_ToEnumJsonBytes(t *testing.T) {
	e := byteEnum()

	b, err := e.ToEnumJsonBytes(0)
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}

	// Not found
	_, err = e.ToEnumJsonBytes(99)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicByte_ToEnumString(t *testing.T) {
	e := byteEnum()

	if e.ToEnumString(0) != "Invalid" {
		t.Fatal("expected Invalid")
	}
}

func Test_Cov12_BasicByte_AppendPrependJoinValue(t *testing.T) {
	e := byteEnum()
	r := e.AppendPrependJoinValue(".", 1, 0)

	if r != "Invalid.Active" {
		t.Fatal("expected Invalid.Active")
	}
}

func Test_Cov12_BasicByte_AppendPrependJoinNamer(t *testing.T) {
	e := byteEnum()
	r := e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})

	if r != "A.B" {
		t.Fatal("expected A.B")
	}
}

func Test_Cov12_BasicByte_ToNumberString(t *testing.T) {
	e := byteEnum()
	s := e.ToNumberString(byte(42))

	if s != "42" {
		t.Fatal("expected 42")
	}
}

func Test_Cov12_BasicByte_JsonMap(t *testing.T) {
	e := byteEnum()
	_ = e.JsonMap()
}

func Test_Cov12_BasicByte_UnmarshallToValue_AllBranches(t *testing.T) {
	e := byteEnum()

	// nil + not mapped
	_, err := e.UnmarshallToValue(false, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	// nil + mapped
	v, err := e.UnmarshallToValue(true, nil)
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	// empty + mapped
	v, err = e.UnmarshallToValue(true, []byte(""))
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	// double quote empty + mapped
	v, err = e.UnmarshallToValue(true, []byte(`""`))
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	// Valid value
	v, err = e.UnmarshallToValue(false, []byte("Active"))
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov12_BasicByte_EnumType(t *testing.T) {
	e := byteEnum()

	if e.EnumType() != enumtype.Byte {
		t.Fatal("expected Byte")
	}
}

func Test_Cov12_BasicByte_AsBasicByter(t *testing.T) {
	e := byteEnum()
	byter := e.AsBasicByter()

	if byter == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicInt8 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_BasicInt8_IsAnyOf(t *testing.T) {
	e := int8Enum()

	if !e.IsAnyOf(0) {
		t.Fatal("empty should return true")
	}

	if !e.IsAnyOf(1, 0, 1) {
		t.Fatal("expected true")
	}

	if e.IsAnyOf(1, 0, 2) {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicInt8_IsAnyNamesOf(t *testing.T) {
	e := int8Enum()

	if !e.IsAnyNamesOf(0, "Invalid") {
		t.Fatal("expected true")
	}

	if e.IsAnyNamesOf(0, "Active") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicInt8_GetValueByName_AllBranches(t *testing.T) {
	e := int8Enum()

	v, err := e.GetValueByName("Active")
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}

	_, err = e.GetValueByName("NotExist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt8_GetValueByString(t *testing.T) {
	e := int8Enum()
	_ = e.GetValueByString("Active")
}

func Test_Cov12_BasicInt8_GetStringValue(t *testing.T) {
	e := int8Enum()
	_ = e.GetStringValue(0)
}

func Test_Cov12_BasicInt8_ExpectingEnumValueError(t *testing.T) {
	e := int8Enum()

	err := e.ExpectingEnumValueError("Active", int8(1))
	if err != nil {
		t.Fatal("expected nil")
	}

	err = e.ExpectingEnumValueError("Active", int8(0))
	if err == nil {
		t.Fatal("expected error")
	}

	err = e.ExpectingEnumValueError("NotExist", int8(0))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt8_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := int8Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_Cov12_BasicInt8_IsValidRange(t *testing.T) {
	e := int8Enum()

	if !e.IsValidRange(1) {
		t.Fatal("expected valid")
	}

	if e.IsValidRange(100) {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_BasicInt8_ToEnumJsonBytes(t *testing.T) {
	e := int8Enum()

	_, err := e.ToEnumJsonBytes(0)
	if err != nil {
		t.Fatal("expected no error")
	}

	_, err = e.ToEnumJsonBytes(99)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt8_ToEnumString(t *testing.T) {
	e := int8Enum()
	_ = e.ToEnumString(0)
}

func Test_Cov12_BasicInt8_AppendPrependJoinValue(t *testing.T) {
	e := int8Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_Cov12_BasicInt8_AppendPrependJoinNamer(t *testing.T) {
	e := int8Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_Cov12_BasicInt8_ToNumberString(t *testing.T) {
	e := int8Enum()
	_ = e.ToNumberString(int8(42))
}

func Test_Cov12_BasicInt8_UnmarshallToValue_AllBranches(t *testing.T) {
	e := int8Enum()

	_, err := e.UnmarshallToValue(false, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	v, err := e.UnmarshallToValue(true, nil)
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	v, err = e.UnmarshallToValue(true, []byte(""))
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov12_BasicInt8_EnumType(t *testing.T) {
	e := int8Enum()

	if e.EnumType() != enumtype.Integer8 {
		t.Fatal("expected Integer8")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicInt16 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_BasicInt16_IsAnyOf(t *testing.T) {
	e := int16Enum()

	if !e.IsAnyOf(0) {
		t.Fatal("empty should return true")
	}

	if !e.IsAnyOf(1, 0, 1) {
		t.Fatal("expected true")
	}

	if e.IsAnyOf(1, 0, 2) {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicInt16_IsAnyNamesOf(t *testing.T) {
	e := int16Enum()

	if !e.IsAnyNamesOf(0, "Invalid") {
		t.Fatal("expected true")
	}

	if e.IsAnyNamesOf(0, "Active") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicInt16_GetValueByName_AllBranches(t *testing.T) {
	e := int16Enum()

	v, err := e.GetValueByName("Active")
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}

	_, err = e.GetValueByName("NotExist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt16_GetValueByString(t *testing.T) {
	e := int16Enum()
	_ = e.GetValueByString("Active")
}

func Test_Cov12_BasicInt16_GetStringValue(t *testing.T) {
	e := int16Enum()
	_ = e.GetStringValue(0)
}

func Test_Cov12_BasicInt16_ExpectingEnumValueError(t *testing.T) {
	e := int16Enum()

	err := e.ExpectingEnumValueError("Active", int16(1))
	if err != nil {
		t.Fatal("expected nil")
	}

	err = e.ExpectingEnumValueError("Active", int16(0))
	if err == nil {
		t.Fatal("expected error")
	}

	err = e.ExpectingEnumValueError("NotExist", int16(0))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt16_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := int16Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_Cov12_BasicInt16_IsValidRange(t *testing.T) {
	e := int16Enum()

	if !e.IsValidRange(1) {
		t.Fatal("expected valid")
	}

	if e.IsValidRange(100) {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_BasicInt16_ToEnumJsonBytes(t *testing.T) {
	e := int16Enum()

	_, err := e.ToEnumJsonBytes(0)
	if err != nil {
		t.Fatal("expected no error")
	}

	_, err = e.ToEnumJsonBytes(99)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt16_ToEnumString(t *testing.T) {
	e := int16Enum()
	_ = e.ToEnumString(0)
}

func Test_Cov12_BasicInt16_AppendPrependJoinValue(t *testing.T) {
	e := int16Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_Cov12_BasicInt16_AppendPrependJoinNamer(t *testing.T) {
	e := int16Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_Cov12_BasicInt16_ToNumberString(t *testing.T) {
	e := int16Enum()
	_ = e.ToNumberString(int16(42))
}

func Test_Cov12_BasicInt16_UnmarshallToValue_AllBranches(t *testing.T) {
	e := int16Enum()

	_, err := e.UnmarshallToValue(false, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	v, err := e.UnmarshallToValue(true, nil)
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	v, err = e.UnmarshallToValue(true, []byte(""))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov12_BasicInt16_EnumType(t *testing.T) {
	e := int16Enum()

	if e.EnumType() != enumtype.Integer16 {
		t.Fatal("expected Integer16")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicInt32 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_BasicInt32_IsAnyOf(t *testing.T) {
	e := int32Enum()

	if !e.IsAnyOf(0) {
		t.Fatal("empty should return true")
	}

	if !e.IsAnyOf(1, 0, 1) {
		t.Fatal("expected true")
	}

	if e.IsAnyOf(1, 0, 2) {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicInt32_IsAnyNamesOf(t *testing.T) {
	e := int32Enum()

	if !e.IsAnyNamesOf(0, "Invalid") {
		t.Fatal("expected true")
	}

	if e.IsAnyNamesOf(0, "Active") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicInt32_GetValueByName_AllBranches(t *testing.T) {
	e := int32Enum()

	v, err := e.GetValueByName("Active")
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}

	_, err = e.GetValueByName("NotExist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt32_GetValueByString(t *testing.T) {
	e := int32Enum()
	_ = e.GetValueByString("Active")
}

func Test_Cov12_BasicInt32_GetStringValue(t *testing.T) {
	e := int32Enum()
	_ = e.GetStringValue(0)
}

func Test_Cov12_BasicInt32_ExpectingEnumValueError(t *testing.T) {
	e := int32Enum()

	err := e.ExpectingEnumValueError("Active", int32(1))
	if err != nil {
		t.Fatal("expected nil")
	}

	err = e.ExpectingEnumValueError("Active", int32(0))
	if err == nil {
		t.Fatal("expected error")
	}

	err = e.ExpectingEnumValueError("NotExist", int32(0))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt32_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := int32Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_Cov12_BasicInt32_IsValidRange(t *testing.T) {
	e := int32Enum()

	if !e.IsValidRange(1) {
		t.Fatal("expected valid")
	}

	if e.IsValidRange(100) {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_BasicInt32_ToEnumJsonBytes(t *testing.T) {
	e := int32Enum()

	_, err := e.ToEnumJsonBytes(0)
	if err != nil {
		t.Fatal("expected no error")
	}

	_, err = e.ToEnumJsonBytes(99)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicInt32_ToEnumString(t *testing.T) {
	e := int32Enum()
	_ = e.ToEnumString(0)
}

func Test_Cov12_BasicInt32_AppendPrependJoinValue(t *testing.T) {
	e := int32Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_Cov12_BasicInt32_AppendPrependJoinNamer(t *testing.T) {
	e := int32Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_Cov12_BasicInt32_ToNumberString(t *testing.T) {
	e := int32Enum()
	_ = e.ToNumberString(int32(42))
}

func Test_Cov12_BasicInt32_UnmarshallToValue_AllBranches(t *testing.T) {
	e := int32Enum()

	_, err := e.UnmarshallToValue(false, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	v, err := e.UnmarshallToValue(true, nil)
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	v, err = e.UnmarshallToValue(true, []byte(""))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov12_BasicInt32_EnumType(t *testing.T) {
	e := int32Enum()

	if e.EnumType() != enumtype.Integer32 {
		t.Fatal("expected Integer32")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicUInt16 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_BasicUInt16_IsAnyOf(t *testing.T) {
	e := uint16Enum()

	if !e.IsAnyOf(0) {
		t.Fatal("empty should return true")
	}

	if !e.IsAnyOf(1, 0, 1) {
		t.Fatal("expected true")
	}

	if e.IsAnyOf(1, 0, 2) {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicUInt16_IsAnyNamesOf(t *testing.T) {
	e := uint16Enum()

	if !e.IsAnyNamesOf(0, "Invalid") {
		t.Fatal("expected true")
	}

	if e.IsAnyNamesOf(0, "Active") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicUInt16_GetValueByName_AllBranches(t *testing.T) {
	e := uint16Enum()

	v, err := e.GetValueByName("Active")
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}

	_, err = e.GetValueByName("NotExist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicUInt16_GetValueByString(t *testing.T) {
	e := uint16Enum()
	_ = e.GetValueByString("Active")
}

func Test_Cov12_BasicUInt16_GetStringValue(t *testing.T) {
	e := uint16Enum()
	_ = e.GetStringValue(0)
}

func Test_Cov12_BasicUInt16_ExpectingEnumValueError(t *testing.T) {
	e := uint16Enum()

	err := e.ExpectingEnumValueError("Active", uint16(1))
	if err != nil {
		t.Fatal("expected nil")
	}

	err = e.ExpectingEnumValueError("Active", uint16(0))
	if err == nil {
		t.Fatal("expected error")
	}

	err = e.ExpectingEnumValueError("NotExist", uint16(0))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicUInt16_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := uint16Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_Cov12_BasicUInt16_IsValidRange(t *testing.T) {
	e := uint16Enum()

	if !e.IsValidRange(1) {
		t.Fatal("expected valid")
	}

	if e.IsValidRange(100) {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_BasicUInt16_ToEnumJsonBytes(t *testing.T) {
	e := uint16Enum()

	_, err := e.ToEnumJsonBytes(0)
	if err != nil {
		t.Fatal("expected no error")
	}

	_, err = e.ToEnumJsonBytes(99)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicUInt16_ToEnumString(t *testing.T) {
	e := uint16Enum()
	_ = e.ToEnumString(0)
}

func Test_Cov12_BasicUInt16_AppendPrependJoinValue(t *testing.T) {
	e := uint16Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_Cov12_BasicUInt16_AppendPrependJoinNamer(t *testing.T) {
	e := uint16Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_Cov12_BasicUInt16_ToNumberString(t *testing.T) {
	e := uint16Enum()
	_ = e.ToNumberString(uint16(42))
}

func Test_Cov12_BasicUInt16_UnmarshallToValue_AllBranches(t *testing.T) {
	e := uint16Enum()

	_, err := e.UnmarshallToValue(false, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	v, err := e.UnmarshallToValue(true, nil)
	if err != nil || v != 0 {
		t.Fatal("expected min")
	}

	v, err = e.UnmarshallToValue(true, []byte(""))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	if err != nil || v != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov12_BasicUInt16_EnumType(t *testing.T) {
	e := uint16Enum()

	if e.EnumType() != enumtype.UnsignedInteger16 {
		t.Fatal("expected UnsignedInteger16")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicString — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_BasicString_IsAnyNamesOf(t *testing.T) {
	e := stringEnum()

	if !e.IsAnyNamesOf("Invalid", "Invalid", "Active") {
		t.Fatal("expected true")
	}

	if e.IsAnyNamesOf("Invalid", "Active") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicString_IsAnyOf(t *testing.T) {
	e := stringEnum()

	if !e.IsAnyOf("x") {
		t.Fatal("empty should return true")
	}

	if !e.IsAnyOf("Active", "Invalid", "Active") {
		t.Fatal("expected true")
	}

	if e.IsAnyOf("Active", "Invalid") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_BasicString_MinMax(t *testing.T) {
	e := stringEnum()
	_ = e.Min()
	_ = e.Max()
}

func Test_Cov12_BasicString_Ranges(t *testing.T) {
	e := stringEnum()

	if len(e.Ranges()) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov12_BasicString_HasAnyItem(t *testing.T) {
	e := stringEnum()

	if !e.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func Test_Cov12_BasicString_MaxIndex(t *testing.T) {
	e := stringEnum()

	if e.MaxIndex() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov12_BasicString_GetNameByIndex(t *testing.T) {
	e := stringEnum()

	if e.GetNameByIndex(1) != "Active" {
		t.Fatal("expected Active")
	}

	if e.GetNameByIndex(100) != "" {
		t.Fatal("expected empty")
	}

	if e.GetNameByIndex(0) != "" {
		t.Fatal("expected empty for index 0 (condition > 0)")
	}
}

func Test_Cov12_BasicString_GetIndexByName(t *testing.T) {
	e := stringEnum()

	idx := e.GetIndexByName("Active")
	if idx < 0 {
		t.Fatal("expected valid index")
	}

	idx = e.GetIndexByName("")
	if idx >= 0 {
		t.Fatal("expected invalid for empty")
	}

	idx = e.GetIndexByName("NotExist")
	if idx >= 0 {
		t.Fatal("expected invalid for not exist")
	}
}

func Test_Cov12_BasicString_NameWithIndexMap(t *testing.T) {
	e := stringEnum()
	m := e.NameWithIndexMap()

	if len(m) == 0 {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_BasicString_RangesIntegers(t *testing.T) {
	e := stringEnum()
	r := e.RangesIntegers()

	if len(r) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov12_BasicString_Hashset_HashsetPtr(t *testing.T) {
	e := stringEnum()
	_ = e.Hashset()
	_ = e.HashsetPtr()
}

func Test_Cov12_BasicString_GetValueByName_AllBranches(t *testing.T) {
	e := stringEnum()

	v, err := e.GetValueByName("Active")
	if err != nil || v != "Active" {
		t.Fatal("expected Active")
	}

	_, err = e.GetValueByName("NotExist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicString_IsValidRange(t *testing.T) {
	e := stringEnum()

	if !e.IsValidRange("Active") {
		t.Fatal("expected valid")
	}

	if e.IsValidRange("NotExist") {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_BasicString_OnlySupportedErr(t *testing.T) {
	e := stringEnum()
	err := e.OnlySupportedErr("Active")

	if err == nil {
		t.Fatal("expected error (unsupported exist)")
	}

	err = e.OnlySupportedErr("Invalid", "Active", "Inactive")
	if err != nil {
		t.Fatal("expected nil (all supported)")
	}
}

func Test_Cov12_BasicString_OnlySupportedMsgErr(t *testing.T) {
	e := stringEnum()
	err := e.OnlySupportedMsgErr("test msg", "Active")

	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicString_AppendPrependJoinValue(t *testing.T) {
	e := stringEnum()
	_ = e.AppendPrependJoinValue(".", "Active", "Invalid")
}

func Test_Cov12_BasicString_AppendPrependJoinNamer(t *testing.T) {
	e := stringEnum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_Cov12_BasicString_ToEnumJsonBytes(t *testing.T) {
	e := stringEnum()

	_, err := e.ToEnumJsonBytes("Active")
	if err != nil {
		t.Fatal("expected no error")
	}

	_, err = e.ToEnumJsonBytes("NotExist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_BasicString_UnmarshallToValue_AllBranches(t *testing.T) {
	e := stringEnum()

	_, err := e.UnmarshallToValue(false, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	v, err := e.UnmarshallToValue(true, nil)
	if err != nil {
		t.Fatal("expected no error")
	}
	_ = v

	v, err = e.UnmarshallToValue(true, []byte(""))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	if err != nil {
		t.Fatal("expected no error")
	}

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	if err != nil || v != "Active" {
		t.Fatal("expected Active")
	}
}

func Test_Cov12_BasicString_EnumType(t *testing.T) {
	e := stringEnum()

	if e.EnumType() != enumtype.String {
		t.Fatal("expected String")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// numberEnumBase — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NumberEnumBase_MinMaxAny(t *testing.T) {
	e := byteEnum()
	min, max := e.MinMaxAny()
	_ = min
	_ = max
}

func Test_Cov12_NumberEnumBase_MinValueString_MaxValueString(t *testing.T) {
	e := byteEnum()
	s := e.MinValueString()
	if s == "" {
		t.Fatal("expected non-empty")
	}

	s = e.MaxValueString()
	if s == "" {
		t.Fatal("expected non-empty")
	}

	// Call again to test cached path
	s = e.MinValueString()
	s = e.MaxValueString()
	_ = s
}

func Test_Cov12_NumberEnumBase_MinInt_MaxInt(t *testing.T) {
	e := byteEnum()
	_ = e.MinInt()
	_ = e.MaxInt()
}

func Test_Cov12_NumberEnumBase_AllNameValues(t *testing.T) {
	e := byteEnum()
	nvs := e.AllNameValues()

	if len(nvs) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov12_NumberEnumBase_RangesMap(t *testing.T) {
	e := byteEnum()
	_ = e.RangesMap()
}

func Test_Cov12_NumberEnumBase_OnlySupportedErr(t *testing.T) {
	e := byteEnum()
	err := e.OnlySupportedErr("Active")

	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_NumberEnumBase_OnlySupportedMsgErr(t *testing.T) {
	e := byteEnum()
	err := e.OnlySupportedMsgErr("msg", "Active")

	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_NumberEnumBase_IntegerEnumRanges(t *testing.T) {
	e := byteEnum()
	r := e.IntegerEnumRanges()

	if len(r) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov12_NumberEnumBase_Length_Count(t *testing.T) {
	e := byteEnum()

	if e.Length() != 3 {
		t.Fatal("expected 3")
	}

	if e.Count() != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov12_NumberEnumBase_RangesDynamicMap_DynamicMap(t *testing.T) {
	e := byteEnum()
	_ = e.RangesDynamicMap()
	_ = e.DynamicMap()

	// Call again to test cached path
	_ = e.RangesDynamicMap()
}

func Test_Cov12_NumberEnumBase_RangesIntegerStringMap(t *testing.T) {
	e := byteEnum()
	_ = e.RangesIntegerStringMap()
}

func Test_Cov12_NumberEnumBase_KeyAnyValues(t *testing.T) {
	e := byteEnum()
	_ = e.KeyAnyValues()

	// Call again for cached
	_ = e.KeyAnyValues()
}

func Test_Cov12_NumberEnumBase_KeyValIntegers(t *testing.T) {
	e := byteEnum()
	kvs := e.KeyValIntegers()

	if len(kvs) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov12_NumberEnumBase_Loop(t *testing.T) {
	e := byteEnum()
	count := 0

	e.Loop(func(index int, name string, anyVal any) bool {
		count++
		return false
	})

	if count != 3 {
		t.Fatal("expected 3")
	}

	// Test break
	count = 0

	e.Loop(func(index int, name string, anyVal any) bool {
		count++
		return true
	})

	if count != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov12_NumberEnumBase_LoopInteger(t *testing.T) {
	e := byteEnum()
	count := 0

	e.LoopInteger(func(index int, name string, anyVal int) bool {
		count++
		return false
	})

	if count != 3 {
		t.Fatal("expected 3")
	}

	// Test break
	count = 0

	e.LoopInteger(func(index int, name string, anyVal int) bool {
		count++
		return true
	})

	if count != 1 {
		t.Fatal("expected 1")
	}
}

func Test_Cov12_NumberEnumBase_TypeName(t *testing.T) {
	e := byteEnum()

	if e.TypeName() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_NumberEnumBase_NameWithValueOption(t *testing.T) {
	e := byteEnum()
	_ = e.NameWithValueOption(byte(1), true)
	_ = e.NameWithValueOption(byte(1), false)
}

func Test_Cov12_NumberEnumBase_NameWithValue(t *testing.T) {
	e := byteEnum()
	_ = e.NameWithValue(byte(1))
}

func Test_Cov12_NumberEnumBase_ValueString(t *testing.T) {
	e := byteEnum()
	_ = e.ValueString(byte(1))
}

func Test_Cov12_NumberEnumBase_Format(t *testing.T) {
	e := byteEnum()
	r := e.Format("Enum of {type-name} - {name} - {value}", byte(1))

	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_NumberEnumBase_RangeNamesCsv(t *testing.T) {
	e := byteEnum()
	csv := e.RangeNamesCsv()

	if csv == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_NumberEnumBase_RangesInvalidMessage(t *testing.T) {
	e := byteEnum()
	msg := e.RangesInvalidMessage()

	if msg == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_NumberEnumBase_RangesInvalidErr(t *testing.T) {
	e := byteEnum()
	err := e.RangesInvalidErr()

	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NumberEnumBase_StringRangesPtr_StringRanges(t *testing.T) {
	e := byteEnum()
	_ = e.StringRangesPtr()
	_ = e.StringRanges()
}

func Test_Cov12_NumberEnumBase_NamesHashset(t *testing.T) {
	e := byteEnum()
	h := e.NamesHashset()

	if !h["Active"] {
		t.Fatal("expected Active in hashset")
	}
}

func Test_Cov12_NumberEnumBase_JsonString(t *testing.T) {
	e := byteEnum()
	_ = e.JsonString(byte(1))
}

func Test_Cov12_NumberEnumBase_ToEnumString_ToName(t *testing.T) {
	e := byteEnum()
	_ = e.ToEnumString(byte(1))
	_ = e.ToName(byte(1))
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_DynamicMap_AddOrUpdate(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.AddOrUpdate("b", 2)

	if !isNew {
		t.Fatal("expected new")
	}

	isNew = dm.AddOrUpdate("a", 3)

	if isNew {
		t.Fatal("expected not new")
	}
}

func Test_Cov12_DynamicMap_Set(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.Set("b", 2)

	if !isNew {
		t.Fatal("expected new")
	}
}

func Test_Cov12_DynamicMap_AddNewOnly(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	isAdded := dm.AddNewOnly("b", 2)

	if !isAdded {
		t.Fatal("expected added")
	}

	isAdded = dm.AddNewOnly("a", 3)

	if isAdded {
		t.Fatal("expected not added")
	}
}

func Test_Cov12_DynamicMap_AllKeys_AllKeysSorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeys()

	if len(keys) != 2 {
		t.Fatal("expected 2")
	}

	sorted := dm.AllKeysSorted()

	if sorted[0] != "a" || sorted[1] != "b" {
		t.Fatal("expected sorted")
	}

	empty := enumimpl.DynamicMap{}
	_ = empty.AllKeys()
	_ = empty.AllKeysSorted()
}

func Test_Cov12_DynamicMap_AllValuesStrings_Sorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	_ = dm.AllValuesStrings()
	_ = dm.AllValuesStringsSorted()

	empty := enumimpl.DynamicMap{}
	_ = empty.AllValuesStrings()
	_ = empty.AllValuesStringsSorted()
}

func Test_Cov12_DynamicMap_AllValuesIntegers(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	vals := dm.AllValuesIntegers()

	if len(vals) != 2 {
		t.Fatal("expected 2")
	}

	empty := enumimpl.DynamicMap{}
	_ = empty.AllValuesIntegers()
}

func Test_Cov12_DynamicMap_MapIntegerString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	m, keys := dm.MapIntegerString()
	_ = m
	_ = keys

	empty := enumimpl.DynamicMap{}
	_, _ = empty.MapIntegerString()

	// String values
	strDm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	_, _ = strDm.MapIntegerString()
}

func Test_Cov12_DynamicMap_SortedKeyValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	_ = dm.SortedKeyValues()

	empty := enumimpl.DynamicMap{}
	_ = empty.SortedKeyValues()
}

func Test_Cov12_DynamicMap_SortedKeyAnyValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	_ = dm.SortedKeyAnyValues()

	empty := enumimpl.DynamicMap{}
	_ = empty.SortedKeyAnyValues()

	// String values
	strDm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	_ = strDm.SortedKeyAnyValues()
}

func Test_Cov12_DynamicMap_First(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	k, v := dm.First()
	_ = k
	_ = v

	empty := enumimpl.DynamicMap{}
	k, v = empty.First()

	if k != "" || v != nil {
		t.Fatal("expected empty")
	}
}

func Test_Cov12_DynamicMap_IsValueTypeOf(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	isInt := dm.IsValueTypeOf(reflect.TypeOf(1))

	if !isInt {
		t.Fatal("expected true")
	}
}

func Test_Cov12_DynamicMap_IsValueString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}

	if dm.IsValueString() {
		t.Fatal("expected false")
	}

	strDm := enumimpl.DynamicMap{"a": "x"}

	if !strDm.IsValueString() {
		t.Fatal("expected true")
	}
}

func Test_Cov12_DynamicMap_Length_Count_IsEmpty_HasAnyItem(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}

	if dm.Length() != 1 || dm.Count() != 1 {
		t.Fatal("expected 1")
	}

	if dm.IsEmpty() || !dm.HasAnyItem() {
		t.Fatal("expected not empty")
	}

	var nilDm *enumimpl.DynamicMap

	if nilDm.Length() != 0 {
		t.Fatal("nil Length should be 0")
	}
}

func Test_Cov12_DynamicMap_LastIndex_HasIndex(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	if dm.LastIndex() != 1 {
		t.Fatal("expected 1")
	}

	if !dm.HasIndex(1) || dm.HasIndex(2) {
		t.Fatal("index check failed")
	}
}

func Test_Cov12_DynamicMap_HasKey_IsMissingKey_HasAllKeys_HasAnyKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	if !dm.HasKey("a") || dm.HasKey("c") {
		t.Fatal("HasKey failed")
	}

	if dm.IsMissingKey("a") || !dm.IsMissingKey("c") {
		t.Fatal("IsMissingKey failed")
	}

	if !dm.HasAllKeys("a", "b") || dm.HasAllKeys("a", "c") {
		t.Fatal("HasAllKeys failed")
	}

	if !dm.HasAnyKeys("a", "c") || dm.HasAnyKeys("c", "d") {
		t.Fatal("HasAnyKeys failed")
	}
}

func Test_Cov12_DynamicMap_IsEqual_AllBranches(t *testing.T) {
	var nilA, nilB *enumimpl.DynamicMap

	if !nilA.IsEqual(false, nilB) {
		t.Fatal("both nil should be equal")
	}

	dm := enumimpl.DynamicMap{"a": 1}

	if nilA.IsEqual(false, &dm) {
		t.Fatal("nil vs non-nil should not be equal")
	}

	if dm.IsEqual(false, nilA) {
		t.Fatal("non-nil vs nil should not be equal")
	}

	dm2 := enumimpl.DynamicMap{"a": 1}

	if !dm.IsEqual(false, &dm2) {
		t.Fatal("same maps should be equal")
	}

	if !dm.IsEqual(false, &dm) {
		t.Fatal("same pointer should be equal")
	}
}

func Test_Cov12_DynamicMap_IsRawEqual(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	raw := map[string]any{"a": 1}

	if !dm.IsRawEqual(false, raw) {
		t.Fatal("expected equal")
	}

	// Different length
	if dm.IsRawEqual(false, map[string]any{"a": 1, "b": 2}) {
		t.Fatal("expected not equal")
	}

	// Missing key
	if dm.IsRawEqual(false, map[string]any{"b": 1}) {
		t.Fatal("expected not equal")
	}

	// nil checks
	var nilDm *enumimpl.DynamicMap

	if !nilDm.IsRawEqual(false, nil) {
		t.Fatal("both nil should be equal")
	}

	if nilDm.IsRawEqual(false, raw) {
		t.Fatal("nil vs non-nil should not be equal")
	}
}

func Test_Cov12_DynamicMap_IsMismatch_IsRawMismatch(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 2}

	if !dm.IsMismatch(false, &dm2) {
		t.Fatal("expected mismatch")
	}

	if !dm.IsRawMismatch(false, map[string]any{"a": 2}) {
		t.Fatal("expected mismatch")
	}
}

func Test_Cov12_DynamicMap_Raw(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.Raw()
}

func Test_Cov12_DynamicMap_DiffRaw(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}

	diff := dm.DiffRaw(false, right)

	if diff.IsEmpty() {
		t.Fatal("expected diff")
	}
}

func Test_Cov12_DynamicMap_DiffRawUsingDifferChecker_AllBranches(t *testing.T) {
	// nil left, nil right
	var nilDm *enumimpl.DynamicMap
	diff := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	if diff.HasAnyItem() {
		t.Fatal("expected empty")
	}

	// nil left, non-nil right
	right := map[string]any{"a": 1}
	diff = nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)

	if diff.IsEmpty() {
		t.Fatal("expected non-empty")
	}

	// non-nil left, nil right
	dm := enumimpl.DynamicMap{"a": 1}
	diff = dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	if diff.IsEmpty() {
		t.Fatal("expected non-empty")
	}

	// Equal maps
	diff = dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})

	if diff.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_Cov12_DynamicMap_DiffRawLeftRightUsingDifferChecker(t *testing.T) {
	var nilDm *enumimpl.DynamicMap
	l, r := nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	_, _ = l, r

	l, r = nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})
	_, _ = l, r

	dm := enumimpl.DynamicMap{"a": 1}
	l, r = dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	_, _ = l, r

	dm2 := enumimpl.DynamicMap{"a": 1, "b": 2}
	l, r = dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, dm2)
	_, _ = l, r

	// Equal maps
	l, r = dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})
	_, _ = l, r
}

func Test_Cov12_DynamicMap_DiffJsonMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessage(false, map[string]any{"a": 1})

	if msg != "" {
		t.Fatal("expected empty for equal")
	}

	msg = dm.DiffJsonMessage(false, map[string]any{"a": 2})

	if msg == "" {
		t.Fatal("expected non-empty for different")
	}
}

func Test_Cov12_DynamicMap_DiffJsonMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.DiffJsonMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 2})
}

func Test_Cov12_DynamicMap_DiffJsonMessageLeftRight(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageLeftRight(false, map[string]any{"a": 1})

	if msg != "" {
		t.Fatal("expected empty")
	}

	msg = dm.DiffJsonMessageLeftRight(false, map[string]any{"a": 2, "b": 3})

	if msg == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_DynamicMap_ShouldDiffMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 1})

	if msg != "" {
		t.Fatal("expected empty")
	}

	msg = dm.ShouldDiffMessage(false, "test", map[string]any{"a": 2})

	if msg == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_DynamicMap_ShouldDiffMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.ShouldDiffMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
	_ = dm.ShouldDiffMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 2})
}

func Test_Cov12_DynamicMap_ShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.ShouldDiffLeftRightMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
	_ = dm.ShouldDiffLeftRightMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 2})
}

func Test_Cov12_DynamicMap_ExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 1})

	if msg != "" {
		t.Fatal("expected empty for equal")
	}

	msg = dm.ExpectingMessage("test", map[string]any{"a": 2})

	if msg == "" {
		t.Fatal("expected non-empty for different")
	}
}

func Test_Cov12_DynamicMap_IsKeysEqualOnly(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	if !dm.IsKeysEqualOnly(map[string]any{"a": 99, "b": 88}) {
		t.Fatal("expected true")
	}

	if dm.IsKeysEqualOnly(map[string]any{"a": 1}) {
		t.Fatal("expected false")
	}

	if dm.IsKeysEqualOnly(map[string]any{"a": 1, "c": 3}) {
		t.Fatal("expected false")
	}

	var nilDm *enumimpl.DynamicMap

	if !nilDm.IsKeysEqualOnly(nil) {
		t.Fatal("both nil should be equal")
	}

	if nilDm.IsKeysEqualOnly(map[string]any{"a": 1}) {
		t.Fatal("nil vs non-nil should not be equal")
	}
}

func Test_Cov12_DynamicMap_KeyValue_KeyValueString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	v, found := dm.KeyValue("a")

	if !found || v != 1 {
		t.Fatal("expected 1")
	}

	_, found = dm.KeyValue("z")

	if found {
		t.Fatal("expected not found")
	}

	s, found := dm.KeyValueString("a")

	if !found || s != "1" {
		t.Fatal("expected 1")
	}

	_, found = dm.KeyValueString("z")

	if found {
		t.Fatal("expected not found")
	}
}

func Test_Cov12_DynamicMap_KeyValueIntDefault(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 42, "b": "notint"}
	v := dm.KeyValueIntDefault("a")

	if v != 42 {
		t.Fatal("expected 42")
	}

	v = dm.KeyValueIntDefault("z")

	if v >= 0 {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_DynamicMap_KeyValueByte(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": byte(42), "b": 100}
	v, found, failed := dm.KeyValueByte("a")

	if !found || failed || v != 42 {
		t.Fatal("expected 42")
	}

	_, found, _ = dm.KeyValueByte("z")

	if found {
		t.Fatal("expected not found")
	}
}

func Test_Cov12_DynamicMap_KeyValueInt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 42, "b": byte(10)}
	v, found, failed := dm.KeyValueInt("a")

	if !found || failed || v != 42 {
		t.Fatal("expected 42")
	}

	v, found, failed = dm.KeyValueInt("b")
	_ = v
	_ = found
	_ = failed

	_, found, _ = dm.KeyValueInt("z")

	if found {
		t.Fatal("expected not found")
	}
}

func Test_Cov12_DynamicMap_Add(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	result := dm.Add("key", "val")

	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_DynamicMap_ConvMap_Methods(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	_ = dm.ConvMapStringInteger()
	_ = dm.ConvMapIntegerString()
	_ = dm.ConvMapByteString()
	_ = dm.ConvMapInt8String()
	_ = dm.ConvMapInt16String()
	_ = dm.ConvMapInt32String()
	_ = dm.ConvMapUInt16String()
	_ = dm.ConvMapStringString()
	_ = dm.ConvMapInt64String()

	empty := enumimpl.DynamicMap{}
	_ = empty.ConvMapStringInteger()
	_ = empty.ConvMapIntegerString()
	_ = empty.ConvMapByteString()
	_ = empty.ConvMapInt8String()
	_ = empty.ConvMapInt16String()
	_ = empty.ConvMapInt32String()
	_ = empty.ConvMapUInt16String()
	_ = empty.ConvMapStringString()
	_ = empty.ConvMapInt64String()
}

func Test_Cov12_DynamicMap_BasicFactories(t *testing.T) {
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}

	_ = dm.BasicByte("test")
	_ = dm.BasicByteUsingAliasMap("test", nil)
	_ = dm.BasicInt8("test")
	_ = dm.BasicInt8UsingAliasMap("test", nil)
	_ = dm.BasicInt16("test")
	_ = dm.BasicInt16UsingAliasMap("test", nil)
	_ = dm.BasicInt32("test")
	_ = dm.BasicInt32UsingAliasMap("test", nil)
	_ = dm.BasicString("test")
	_ = dm.BasicStringUsingAliasMap("test", nil)
	_ = dm.BasicUInt16("test")
	_ = dm.BasicUInt16UsingAliasMap("test", nil)
}

func Test_Cov12_DynamicMap_ConcatNew(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	another := enumimpl.DynamicMap{"b": 2}

	result := dm.ConcatNew(true, another)
	if result.Length() != 2 {
		t.Fatal("expected 2")
	}

	result = dm.ConcatNew(false, another)
	if result.Length() != 2 {
		t.Fatal("expected 2")
	}

	empty := enumimpl.DynamicMap{}
	result = empty.ConcatNew(true, enumimpl.DynamicMap{})

	if result.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_Cov12_DynamicMap_Strings_String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.Strings()
	_ = dm.String()

	empty := enumimpl.DynamicMap{}
	_ = empty.Strings()
	_ = empty.String()
}

func Test_Cov12_DynamicMap_StringsUsingFmt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	r := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return key
	})

	if len(r) != 1 {
		t.Fatal("expected 1")
	}

	empty := enumimpl.DynamicMap{}
	_ = empty.StringsUsingFmt(func(index int, key string, val any) string { return "" })
}

func Test_Cov12_DynamicMap_IsStringEqual(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	s := dm.String()

	if !dm.IsStringEqual(s) {
		t.Fatal("expected true")
	}
}

func Test_Cov12_DynamicMap_Serialize(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	b, err := dm.Serialize()

	if err != nil || len(b) == 0 {
		t.Fatal("expected serialized")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// DiffLeftRight — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_DiffLeftRight_AllMethods(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}

	l, r := d.Types()
	_ = l
	_ = r

	if d.IsSameTypeSame() != true {
		t.Fatal("expected same type")
	}

	if d.IsSame() {
		t.Fatal("expected not same")
	}

	if d.IsSameRegardlessOfType() {
		t.Fatal("expected not same")
	}

	if d.IsEqual(false) {
		t.Fatal("expected not equal")
	}

	if d.IsEqual(true) {
		t.Fatal("expected not equal")
	}

	if !d.HasMismatch(false) {
		t.Fatal("expected mismatch")
	}

	if !d.HasMismatch(true) {
		t.Fatal("expected mismatch")
	}

	if !d.IsNotEqual() {
		t.Fatal("expected not equal")
	}

	if !d.HasMismatchRegardlessOfType() {
		t.Fatal("expected mismatch")
	}

	_ = d.String()
	_ = d.JsonString()

	ls, rs := d.SpecificFullString()
	_ = ls
	_ = rs

	ds := d.DiffString()
	if ds == "" {
		t.Fatal("expected non-empty")
	}

	// Same values
	same := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	ds = same.DiffString()

	if ds != "" {
		t.Fatal("expected empty for same")
	}

	// nil
	var nilD *enumimpl.DiffLeftRight

	if nilD.JsonString() != "" {
		t.Fatal("nil should return empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// differCheckerImpl — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_DifferCheckerImpl(t *testing.T) {
	d := enumimpl.DefaultDiffCheckerImpl

	l := d.GetSingleDiffResult(true, 1, 2)
	if l != 1 {
		t.Fatal("expected left")
	}

	r := d.GetSingleDiffResult(false, 1, 2)
	if r != 2 {
		t.Fatal("expected right")
	}

	val := d.GetResultOnKeyMissingInRightExistInLeft("k", 42)
	if val != 42 {
		t.Fatal("expected 42")
	}

	if !d.IsEqual(false, 1, 1) {
		t.Fatal("expected equal")
	}

	if d.IsEqual(false, 1, 2) {
		t.Fatal("expected not equal")
	}

	if !d.IsEqual(true, 1, 1) {
		t.Fatal("expected equal regardless")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// leftRightDiffCheckerImpl — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_LeftRightDiffCheckerImpl(t *testing.T) {
	d := enumimpl.LeftRightDiffCheckerImpl

	_ = d.GetSingleDiffResult(true, 1, 2)
	_ = d.GetResultOnKeyMissingInRightExistInLeft("k", 42)

	if !d.IsEqual(false, 1, 1) {
		t.Fatal("expected equal")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Format / FormatUsingFmt — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_Format(t *testing.T) {
	result := enumimpl.Format("MyType", "Active", "1", "Enum of {type-name} - {name} - {value}")

	if result == "" {
		t.Fatal("expected non-empty")
	}
}

type mockFormatter struct{}

func (m mockFormatter) TypeName() string   { return "MyType" }
func (m mockFormatter) Name() string       { return "Active" }
func (m mockFormatter) ValueString() string { return "1" }

func Test_Cov12_FormatUsingFmt(t *testing.T) {
	result := enumimpl.FormatUsingFmt(mockFormatter{}, "Enum of {type-name} - {name} - {value}")

	if result == "" {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Standalone functions — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NameWithValue(t *testing.T) {
	r := enumimpl.NameWithValue(42)

	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov12_PrependJoin(t *testing.T) {
	r := enumimpl.PrependJoin(".", "a", "b", "c")

	if r != "a.b.c" {
		t.Fatal("expected a.b.c")
	}
}

func Test_Cov12_JoinPrependUsingDot(t *testing.T) {
	r := enumimpl.JoinPrependUsingDot("a", "b", "c")

	if r != "a.b.c" {
		t.Fatal("expected a.b.c")
	}
}

func Test_Cov12_OnlySupportedErr(t *testing.T) {
	err := enumimpl.OnlySupportedErr(4, []string{"a", "b", "c"}, "a")

	if err == nil {
		t.Fatal("expected error")
	}

	err = enumimpl.OnlySupportedErr(4, []string{"a", "b"}, "a", "b")

	if err != nil {
		t.Fatal("expected nil")
	}

	err = enumimpl.OnlySupportedErr(4, []string{}, "a")

	if err != nil {
		t.Fatal("expected nil for empty allNames")
	}
}

func Test_Cov12_UnsupportedNames(t *testing.T) {
	unsupported := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")

	if len(unsupported) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov12_AllNameValues(t *testing.T) {
	result := enumimpl.AllNameValues([]string{"a", "b"}, []byte{0, 1})

	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_Cov12_IntegersRangesOfAnyVal(t *testing.T) {
	result := enumimpl.IntegersRangesOfAnyVal([]byte{2, 0, 1})

	if len(result) != 3 || result[0] != 0 {
		t.Fatal("expected sorted [0,1,2]")
	}
}

func Test_Cov12_ConvEnumAnyValToInteger_AllBranches(t *testing.T) {
	// int
	if enumimpl.ConvEnumAnyValToInteger(42) != 42 {
		t.Fatal("expected 42")
	}

	// string
	r := enumimpl.ConvEnumAnyValToInteger("hello")
	_ = r

	// byte
	if enumimpl.ConvEnumAnyValToInteger(byte(5)) != 5 {
		t.Fatal("expected 5")
	}

	// non-convertible
	r = enumimpl.ConvEnumAnyValToInteger(struct{}{})
	_ = r
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyVal — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_KeyAnyVal_AllMethods(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "test", AnyValue: 42}

	if kav.KeyString() != "test" {
		t.Fatal("expected test")
	}

	if kav.AnyVal() != 42 {
		t.Fatal("expected 42")
	}

	_ = kav.AnyValString()
	_ = kav.WrapKey()
	_ = kav.WrapValue()

	if kav.IsString() {
		t.Fatal("expected false")
	}

	if kav.ValInt() != 42 {
		t.Fatal("expected 42")
	}

	kvi := kav.KeyValInteger()

	if kvi.Key != "test" {
		t.Fatal("expected test")
	}

	_ = kav.String()

	// String type
	kavStr := enumimpl.KeyAnyVal{Key: "test", AnyValue: "hello"}

	if !kavStr.IsString() {
		t.Fatal("expected true")
	}

	_ = kavStr.String()
}

func Test_Cov12_KeyAnyValues(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []int{1, 2})

	if len(result) != 2 {
		t.Fatal("expected 2")
	}

	result = enumimpl.KeyAnyValues([]string{}, []int{})

	if len(result) != 0 {
		t.Fatal("expected 0")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValInteger — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_KeyValInteger_AllMethods(t *testing.T) {
	kvi := enumimpl.KeyValInteger{Key: "test", ValueInteger: 42}

	_ = kvi.WrapKey()
	_ = kvi.WrapValue()
	_ = kvi.String()

	kav := kvi.KeyAnyVal()

	if kav.Key != "test" {
		t.Fatal("expected test")
	}

	if kvi.IsString() {
		t.Fatal("expected false")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicByteCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NewBasicByte_CreateUsingMap(t *testing.T) {
	e := enumimpl.New.BasicByte.CreateUsingMap("test", map[byte]string{0: "Invalid", 1: "Active"})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_CreateUsingMapPlusAliasMap(t *testing.T) {
	e := enumimpl.New.BasicByte.CreateUsingMapPlusAliasMap("test",
		map[byte]string{0: "Invalid", 1: "Active"},
		map[string]byte{"active": 1})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_Create(t *testing.T) {
	e := enumimpl.New.BasicByte.Create("test", []byte{0, 1}, []string{"Invalid", "Active"}, 0, 1)

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_UsingTypeSlice(t *testing.T) {
	e := enumimpl.New.BasicByte.UsingTypeSlice("test", []string{"Invalid", "Active"})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_DefaultWithAliasMap(t *testing.T) {
	e := enumimpl.New.BasicByte.DefaultWithAliasMap(testByteVal, []string{"Invalid", "Active"}, map[string]byte{"act": 1})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_DefaultAllCases(t *testing.T) {
	e := enumimpl.New.BasicByte.DefaultAllCases(testByteVal, []string{"Invalid", "Active"})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_DefaultWithAliasMapAllCases(t *testing.T) {
	e := enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(testByteVal, []string{"Invalid", "Active"}, map[string]byte{"act": 1})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_UsingFirstItemSliceAliasMap(t *testing.T) {
	e := enumimpl.New.BasicByte.UsingFirstItemSliceAliasMap(testByteVal, []string{"Invalid", "Active"}, nil)

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_UsingFirstItemSliceCaseOptions(t *testing.T) {
	e := enumimpl.New.BasicByte.UsingFirstItemSliceCaseOptions(true, testByteVal, []string{"Invalid", "Active"})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_UsingFirstItemSliceAllCases(t *testing.T) {
	e := enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(testByteVal, []string{"Invalid", "Active"})

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_CreateUsingSlicePlusAliasMapOptions(t *testing.T) {
	e := enumimpl.New.BasicByte.CreateUsingSlicePlusAliasMapOptions(false, testByteVal, []string{"Invalid", "Active"}, nil)

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov12_NewBasicByte_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	e := enumimpl.New.BasicByte.CreateUsingMapPlusAliasMapOptions(true, testByteVal, map[byte]string{0: "Invalid", 1: "Active"}, nil)

	if e == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicInt8Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NewBasicInt8_All(t *testing.T) {
	_ = enumimpl.New.BasicInt8.CreateUsingMap("test", map[int8]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicInt8.CreateUsingMapPlusAliasMap("test", map[int8]string{0: "Invalid"}, map[string]int8{"act": 1})
	_ = enumimpl.New.BasicInt8.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt8.DefaultWithAliasMap(testInt8Val, []string{"Invalid", "Active"}, map[string]int8{"act": 1})
	_ = enumimpl.New.BasicInt8.DefaultAllCases(testInt8Val, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt8.DefaultWithAliasMapAllCases(testInt8Val, []string{"Invalid", "Active"}, map[string]int8{"act": 1})
	_ = enumimpl.New.BasicInt8.UsingFirstItemSliceAliasMap(testInt8Val, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicInt8.CreateUsingSlicePlusAliasMapOptions(false, testInt8Val, []string{"Invalid"}, nil)
	_ = enumimpl.New.BasicInt8.CreateUsingMapPlusAliasMapOptions(true, testInt8Val, map[int8]string{0: "Invalid"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicInt16Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NewBasicInt16_All(t *testing.T) {
	_ = enumimpl.New.BasicInt16.CreateUsingMap("test", map[int16]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicInt16.CreateUsingMapPlusAliasMap("test", map[int16]string{0: "Invalid"}, map[string]int16{"act": 1})
	_ = enumimpl.New.BasicInt16.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt16.DefaultWithAliasMap(testInt16Val, []string{"Invalid", "Active"}, map[string]int16{"act": 1})
	_ = enumimpl.New.BasicInt16.DefaultAllCases(testInt16Val, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt16.DefaultWithAliasMapAllCases(testInt16Val, []string{"Invalid", "Active"}, map[string]int16{"act": 1})
	_ = enumimpl.New.BasicInt16.UsingFirstItemSliceAliasMap(testInt16Val, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicInt16.CreateUsingSlicePlusAliasMapOptions(false, testInt16Val, []string{"Invalid"}, nil)
	_ = enumimpl.New.BasicInt16.CreateUsingMapPlusAliasMapOptions(true, testInt16Val, map[int16]string{0: "Invalid"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicInt32Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NewBasicInt32_All(t *testing.T) {
	_ = enumimpl.New.BasicInt32.CreateUsingMap("test", map[int32]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicInt32.CreateUsingMapPlusAliasMap("test", map[int32]string{0: "Invalid"}, map[string]int32{"act": 1})
	_ = enumimpl.New.BasicInt32.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt32.DefaultWithAliasMap(testInt32Val, []string{"Invalid", "Active"}, map[string]int32{"act": 1})
	_ = enumimpl.New.BasicInt32.UsingFirstItemSliceAliasMap(testInt32Val, []string{"Invalid", "Active"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicUInt16Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NewBasicUInt16_All(t *testing.T) {
	_ = enumimpl.New.BasicUInt16.CreateUsingMap("test", map[uint16]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicUInt16.CreateUsingMapPlusAliasMap("test", map[uint16]string{0: "Invalid"}, map[string]uint16{"act": 1})
	_ = enumimpl.New.BasicUInt16.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicUInt16.DefaultWithAliasMap(testUInt16Val, []string{"Invalid", "Active"}, map[string]uint16{"act": 1})
	_ = enumimpl.New.BasicUInt16.UsingFirstItemSliceAliasMap(testUInt16Val, []string{"Invalid", "Active"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicStringCreator
// ══════════════════════════════════════════════════════════════════════════════

type testString string

const testStringVal testString = ""

func Test_Cov12_NewBasicString_All(t *testing.T) {
	_ = enumimpl.New.BasicString.Create("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicString.CreateDefault(testStringVal, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicString.CreateAliasMapOnly("test", []string{"Invalid", "Active"}, map[string]string{"act": "Active"})
	_ = enumimpl.New.BasicString.CreateUsingAliasMap("test", []string{"Invalid", "Active"}, nil, "Active", "Invalid")
	_ = enumimpl.New.BasicString.CreateUsingNamesSpread("test", "Invalid", "Active")
	_ = enumimpl.New.BasicString.CreateUsingNamesMinMax("test", []string{"Invalid", "Active"}, "Active", "Invalid")
	_ = enumimpl.New.BasicString.CreateUsingSlicePlusAliasMapOptions(true, testStringVal, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicString.CreateUsingMapPlusAliasMapOptions(false, testStringVal, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicString.UsingFirstItemSliceCaseOptions(true, testStringVal, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicString.UsingFirstItemSliceAllCases(testStringVal, []string{"Invalid", "Active"})
}

// ══════════════════════════════════════════════════════════════════════════════
// enumtype.Variant — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_EnumType_Variant_AllMethods(t *testing.T) {
	v := enumtype.Byte

	_ = v.TypeName()
	_ = v.ValueUInt16()
	_ = v.RangeNamesCsv()

	min, max := v.MinMaxAny()
	_, _ = min, max

	_ = v.MinValueString()
	_ = v.MaxValueString()
	_ = v.MaxInt()
	_ = v.MinInt()
	_ = v.RangesDynamicMap()
	_ = v.IntegerEnumRanges()
	_ = v.EnumType()
	_ = v.Value()
	_ = v.Name()
	_ = v.String()
	_ = v.NameValue()
	_ = v.ToNumberString()
	_ = v.ValueByte()
	_ = v.ValueInt()
	_ = v.ValueInt8()
	_ = v.ValueInt16()
	_ = v.ValueInt32()
	_ = v.ValueString()

	if !v.IsValid() || v.IsInvalid() {
		t.Fatal("expected valid")
	}

	if v.IsNameEqual("Invalid") {
		t.Fatal("expected false")
	}

	if !v.IsNameEqual("Byte") {
		t.Fatal("expected true")
	}

	if !v.IsAnyNamesOf("Invalid", "Byte") {
		t.Fatal("expected true")
	}

	if v.IsAnyNamesOf("Invalid") {
		t.Fatal("expected false")
	}
}

func Test_Cov12_EnumType_Variant_TypeChecks(t *testing.T) {
	if !enumtype.Boolean.IsBoolean() {
		t.Fatal("expected true")
	}

	if !enumtype.Byte.IsByte() {
		t.Fatal("expected true")
	}

	if !enumtype.UnsignedInteger16.IsUnsignedInteger16() {
		t.Fatal("expected true")
	}

	if !enumtype.UnsignedInteger32.IsUnsignedInteger32() {
		t.Fatal("expected true")
	}

	if !enumtype.UnsignedInteger64.IsUnsignedInteger64() {
		t.Fatal("expected true")
	}

	if !enumtype.Integer8.IsInteger8() {
		t.Fatal("expected true")
	}

	if !enumtype.Integer16.IsInteger16() {
		t.Fatal("expected true")
	}

	if !enumtype.Integer32.IsInteger32() {
		t.Fatal("expected true")
	}

	if !enumtype.Integer64.IsInteger64() {
		t.Fatal("expected true")
	}

	if !enumtype.Integer.IsInteger() {
		t.Fatal("expected true")
	}

	if !enumtype.String.IsString() {
		t.Fatal("expected true")
	}

	if !enumtype.Byte.IsNumber() {
		t.Fatal("expected true")
	}

	if !enumtype.Integer8.IsAnyInteger() {
		t.Fatal("expected true")
	}

	if !enumtype.Byte.IsAnyUnsignedNumber() {
		t.Fatal("expected true")
	}

	if enumtype.Invalid.IsValid() {
		t.Fatal("expected invalid")
	}

	if !enumtype.Invalid.IsInvalid() {
		t.Fatal("expected invalid")
	}
}

func Test_Cov12_EnumType_Variant_MarshalJSON(t *testing.T) {
	v := enumtype.Byte
	b, err := v.MarshalJSON()

	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_Cov12_EnumType_Variant_UnmarshalJSON(t *testing.T) {
	v := enumtype.Invalid

	// Valid
	err := v.UnmarshalJSON([]byte(`"Byte"`))
	if err != nil || v != enumtype.Byte {
		t.Fatal("expected Byte")
	}

	// Empty
	err = v.UnmarshalJSON([]byte(""))
	if err == nil {
		t.Fatal("expected error")
	}

	// Too short
	err = v.UnmarshalJSON([]byte(`""`))
	if err == nil {
		t.Fatal("expected error")
	}

	// Not found
	err = v.UnmarshalJSON([]byte(`"NotExist"`))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Cov12_EnumType_Variant_Format_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()

	v := enumtype.Byte
	v.Format("test")
}

func Test_Cov12_EnumType_Variant_RoundTrip(t *testing.T) {
	original := enumtype.Integer32
	b, err := json.Marshal(original)

	if err != nil {
		t.Fatal("marshal failed")
	}

	var result enumtype.Variant
	err = json.Unmarshal(b, &result)

	if err != nil || result != enumtype.Integer32 {
		t.Fatal("expected Integer32")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — isEqualSingle / isNotEqual via IsRawEqual (regardless type)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_DynamicMap_IsRawEqual_RegardlessType(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}

	// Regardless type: int 1 vs float64 should not match in string fmt
	if !dm.IsRawEqual(true, map[string]any{"a": 1}) {
		t.Fatal("expected equal regardless")
	}

	if dm.IsRawEqual(true, map[string]any{"a": 2}) {
		t.Fatal("expected not equal")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap LogShouldDiff* (covers fmt.Println paths - only call, no output check)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_DynamicMap_LogShouldDiffMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffMessage(false, "test", map[string]any{"a": 1})
}

func Test_Cov12_DynamicMap_LogShouldDiffLeftRightMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffLeftRightMessage(false, "test", map[string]any{"a": 1})
}

func Test_Cov12_DynamicMap_LogShouldDiffMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
}

func Test_Cov12_DynamicMap_LogShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffLeftRightMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
}

func Test_Cov12_DynamicMap_LogExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
}

// ══════════════════════════════════════════════════════════════════════════════
// toStringsSliceOfDiffMap — string value branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_DynamicMap_DiffJsonMessage_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "x"}
	msg := dm.DiffJsonMessage(false, map[string]any{"a": "y"})

	if msg == "" {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// NamesHashset empty
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov12_NamesHashset_Empty(t *testing.T) {
	e := enumimpl.New.BasicByte.Create("test", []byte{}, []string{}, 0, 0)
	h := e.NamesHashset()

	if len(h) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_Cov12_BasicString_GetIndexByName_EmptyEnum(t *testing.T) {
	e := enumimpl.New.BasicString.Create("test", []string{})
	idx := e.GetIndexByName("something")

	if idx >= 0 {
		t.Fatal("expected invalid for empty enum")
	}
}
