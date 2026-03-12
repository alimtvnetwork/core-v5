package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/issetter"
)

// ── New / NewMust / NewBool / NewBooleans ──

func Test_New_Valid_Cov2(t *testing.T) {
	v, err := issetter.New("True")
	if err != nil || v != issetter.True {
		t.Error("should parse True")
	}
}

func Test_New_Invalid_Cov2(t *testing.T) {
	_, err := issetter.New("invalid_name_xyz")
	if err == nil {
		t.Error("should error on invalid")
	}
}

func Test_NewMust_Valid_Cov2(t *testing.T) {
	v := issetter.NewMust("True")
	if v != issetter.True {
		t.Error("should be True")
	}
}

func Test_NewBool_True_Cov2(t *testing.T) {
	if issetter.NewBool(true) != issetter.True {
		t.Error("true should be True")
	}
}

func Test_NewBool_False_Cov2(t *testing.T) {
	if issetter.NewBool(false) != issetter.False {
		t.Error("false should be False")
	}
}

func Test_NewBooleans_AllTrue_Cov2(t *testing.T) {
	if issetter.NewBooleans(true, true) != issetter.True {
		t.Error("all true should be True")
	}
}

func Test_NewBooleans_OneFalse_Cov2(t *testing.T) {
	if issetter.NewBooleans(true, false) != issetter.False {
		t.Error("one false should be False")
	}
}

// ── CombinedBooleans ──

func Test_CombinedBooleans_Empty_Cov2(t *testing.T) {
	if issetter.CombinedBooleans() != issetter.True {
		t.Error("empty should be True")
	}
}

func Test_CombinedBooleans_AllTrue_Cov2(t *testing.T) {
	if issetter.CombinedBooleans(true, true, true) != issetter.True {
		t.Error("all true should be True")
	}
}

func Test_CombinedBooleans_OneFalse_Cov2(t *testing.T) {
	if issetter.CombinedBooleans(true, false) != issetter.False {
		t.Error("one false should be False")
	}
}

// ── GetSetByte ──

func Test_GetSetByte_True_Cov2(t *testing.T) {
	r := issetter.GetSetByte(true, byte(issetter.Set), byte(issetter.Unset))
	if r != issetter.Set {
		t.Error("true condition should return Set")
	}
}

func Test_GetSetByte_False_Cov2(t *testing.T) {
	r := issetter.GetSetByte(false, byte(issetter.Set), byte(issetter.Unset))
	if r != issetter.Unset {
		t.Error("false condition should return Unset")
	}
}

// ── GetSetUnset ──

func Test_GetSetUnset_True_Cov2(t *testing.T) {
	if issetter.GetSetUnset(true) != issetter.Set {
		t.Error("true should return Set")
	}
}

func Test_GetSetUnset_False_Cov2(t *testing.T) {
	if issetter.GetSetUnset(false) != issetter.Unset {
		t.Error("false should return Unset")
	}
}

// ── GetSetterByComparing ──

func Test_GetSetterByComparing_Match_Cov2(t *testing.T) {
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, 42, 1, 42, 100)
	if r != issetter.True {
		t.Error("should match 42")
	}
}

func Test_GetSetterByComparing_NoMatch_Cov2(t *testing.T) {
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, 42, 1, 2, 3)
	if r != issetter.False {
		t.Error("should not match")
	}
}

// ── IsCompareResult ──

func Test_IsCompareResult_Equal_Cov2(t *testing.T) {
	if !issetter.True.IsCompareResult(1, corecomparator.Equal) {
		t.Error("True(1) == 1")
	}
}

func Test_IsCompareResult_LeftGreater_Cov2(t *testing.T) {
	if !issetter.Set.IsCompareResult(1, corecomparator.LeftGreater) {
		t.Error("Set(4) > 1")
	}
}

func Test_IsCompareResult_LeftGreaterEqual_Cov2(t *testing.T) {
	if !issetter.True.IsCompareResult(1, corecomparator.LeftGreaterEqual) {
		t.Error("True(1) >= 1")
	}
}

func Test_IsCompareResult_LeftLess_Cov2(t *testing.T) {
	if !issetter.True.IsCompareResult(2, corecomparator.LeftLess) {
		t.Error("True(1) < 2")
	}
}

func Test_IsCompareResult_LeftLessEqual_Cov2(t *testing.T) {
	if !issetter.True.IsCompareResult(1, corecomparator.LeftLessEqual) {
		t.Error("True(1) <= 1")
	}
}

func Test_IsCompareResult_NotEqual_Cov2(t *testing.T) {
	if !issetter.True.IsCompareResult(2, corecomparator.NotEqual) {
		t.Error("True(1) != 2")
	}
}

// ── IsOutOfRange ──

func Test_IsOutOfRange_InRange_Cov2(t *testing.T) {
	if issetter.IsOutOfRange(1) {
		t.Error("1 should be in range")
	}
}

func Test_IsOutOfRange_OutOfRange_Cov2(t *testing.T) {
	if !issetter.IsOutOfRange(255) {
		t.Error("255 should be out of range")
	}
}

// ── Value methods not yet covered ──

func Test_Value_Is_Cov2(t *testing.T) {
	if !issetter.True.Is(issetter.True) {
		t.Error("True.Is(True) should be true")
	}
}

func Test_Value_IsEqual_Cov2(t *testing.T) {
	if !issetter.True.IsEqual(1) {
		t.Error("True.IsEqual(1)")
	}
}

func Test_Value_IsGreater_Cov2(t *testing.T) {
	if !issetter.Set.IsGreater(1) {
		t.Error("Set(4) > 1")
	}
}

func Test_Value_IsGreaterEqual_Cov2(t *testing.T) {
	if !issetter.True.IsGreaterEqual(1) {
		t.Error("True(1) >= 1")
	}
}

func Test_Value_IsLess_Cov2(t *testing.T) {
	if !issetter.True.IsLess(2) {
		t.Error("True(1) < 2")
	}
}

func Test_Value_IsLessEqual_Cov2(t *testing.T) {
	if !issetter.True.IsLessEqual(1) {
		t.Error("True(1) <= 1")
	}
}

func Test_Value_IsEqualInt_Cov2(t *testing.T) {
	if !issetter.True.IsEqualInt(1) {
		t.Error("True.IsEqualInt(1)")
	}
}

func Test_Value_IsGreaterInt_Cov2(t *testing.T) {
	if !issetter.Set.IsGreaterInt(1) {
		t.Error("Set(4) > 1")
	}
}

func Test_Value_IsGreaterEqualInt_Cov2(t *testing.T) {
	if !issetter.True.IsGreaterEqualInt(1) {
		t.Error("True >= 1")
	}
}

func Test_Value_IsLessInt_Cov2(t *testing.T) {
	if !issetter.True.IsLessInt(2) {
		t.Error("True < 2")
	}
}

func Test_Value_IsLessEqualInt_Cov2(t *testing.T) {
	if !issetter.True.IsLessEqualInt(1) {
		t.Error("True <= 1")
	}
}

func Test_Value_IsBetween_Cov2(t *testing.T) {
	if !issetter.True.IsBetween(0, 5) {
		t.Error("True(1) between 0-5")
	}
	if issetter.True.IsBetween(2, 5) {
		t.Error("True(1) not between 2-5")
	}
}

func Test_Value_IsBetweenInt_Cov2(t *testing.T) {
	if !issetter.True.IsBetweenInt(0, 5) {
		t.Error("True(1) between 0-5")
	}
}

func Test_Value_Add_Cov2(t *testing.T) {
	r := issetter.True.Add(1)
	if r != issetter.False {
		t.Error("True(1)+1 = False(2)")
	}
}

func Test_Value_IsNegative_Cov2(t *testing.T) {
	if !issetter.Uninitialized.IsNegative() {
		t.Error("Uninitialized should be negative")
	}
	if !issetter.False.IsNegative() {
		t.Error("False should be negative")
	}
	if !issetter.Unset.IsNegative() {
		t.Error("Unset should be negative")
	}
	if issetter.True.IsNegative() {
		t.Error("True should not be negative")
	}
}

func Test_Value_IsPositive_Cov2(t *testing.T) {
	if !issetter.True.IsPositive() {
		t.Error("True should be positive")
	}
	if !issetter.Set.IsPositive() {
		t.Error("Set should be positive")
	}
}

func Test_Value_GetErrorOnOutOfRange_InRange_Cov2(t *testing.T) {
	err := issetter.True.GetErrorOnOutOfRange(1, "test")
	if err != nil {
		t.Error("in range should be nil")
	}
}

func Test_Value_GetErrorOnOutOfRange_OutOfRange_Cov2(t *testing.T) {
	err := issetter.True.GetErrorOnOutOfRange(255, "test")
	if err == nil {
		t.Error("out of range should error")
	}
}

func Test_Value_YesNoMappedValue_Cov2(t *testing.T) {
	if issetter.True.YesNoMappedValue() != "yes" {
		t.Error("True should be yes")
	}
	if issetter.False.YesNoMappedValue() != "no" {
		t.Error("False should be no")
	}
	if issetter.Uninitialized.YesNoMappedValue() != "" {
		t.Error("Uninitialized should be empty")
	}
}

func Test_Value_YesNoLowercaseName_Cov2(t *testing.T) {
	if issetter.True.YesNoLowercaseName() != "yes" {
		t.Error("True should be yes")
	}
}

func Test_Value_YesNoName_Cov2(t *testing.T) {
	if issetter.True.YesNoName() != "Yes" {
		t.Error("True should be Yes")
	}
}

func Test_Value_TrueFalseName_Cov2(t *testing.T) {
	if issetter.True.TrueFalseName() != "True" {
		t.Error("True should be True")
	}
}

func Test_Value_OnOffLowercaseName_Cov2(t *testing.T) {
	if issetter.True.OnOffLowercaseName() != "on" {
		t.Error("True should be on")
	}
}

func Test_Value_OnOffName_Cov2(t *testing.T) {
	if issetter.True.OnOffName() != "On" {
		t.Error("True should be On")
	}
}

func Test_Value_TrueFalseLowercaseName_Cov2(t *testing.T) {
	if issetter.True.TrueFalseLowercaseName() != "true" {
		t.Error("True should be true")
	}
}

func Test_Value_SetUnsetLowercaseName_Cov2(t *testing.T) {
	if issetter.True.SetUnsetLowercaseName() != "set" {
		t.Error("True should be set")
	}
}

func Test_Value_Serialize_Cov2(t *testing.T) {
	data, err := issetter.True.Serialize()
	if err != nil || len(data) == 0 {
		t.Error("serialize should work")
	}
}

func Test_Value_TypeName_Cov2(t *testing.T) {
	if issetter.True.TypeName() == "" {
		t.Error("should not be empty")
	}
}

func Test_Value_IsAnyValuesEqual_Cov2(t *testing.T) {
	if !issetter.True.IsAnyValuesEqual(0, 1, 2) {
		t.Error("should find 1")
	}
	if issetter.True.IsAnyValuesEqual(0, 2, 3) {
		t.Error("should not find 1")
	}
}

func Test_Value_UnmarshallEnumToValue_Cov2(t *testing.T) {
	val, err := issetter.Uninitialized.UnmarshallEnumToValue([]byte(`"True"`))
	if err != nil || val != 1 {
		t.Error("should unmarshal True to 1")
	}
}

func Test_Value_Deserialize_Valid_Cov2(t *testing.T) {
	v, err := issetter.Uninitialized.Deserialize([]byte(`"True"`))
	if err != nil || v != issetter.True {
		t.Error("should deserialize True")
	}
}

func Test_Value_Deserialize_Invalid_Cov2(t *testing.T) {
	_, err := issetter.Uninitialized.Deserialize([]byte(`"INVALID_XYZ"`))
	if err == nil {
		t.Error("should error on invalid")
	}
}

func Test_Value_UnmarshalJSON_Nil_Cov2(t *testing.T) {
	var v issetter.Value
	err := v.UnmarshalJSON(nil)
	if err == nil {
		t.Error("nil data should error")
	}
}

func Test_Value_UnmarshalJSON_Invalid_Cov2(t *testing.T) {
	var v issetter.Value
	err := v.UnmarshalJSON([]byte(`"UNKNOWN_XYZ"`))
	if err == nil {
		t.Error("invalid value should error")
	}
}

func Test_Value_MaxByte_Cov2(t *testing.T) {
	if issetter.True.MaxByte() != issetter.Wildcard.ValueByte() {
		t.Error("MaxByte should be Wildcard")
	}
}

func Test_Value_MinByte_Cov2(t *testing.T) {
	if issetter.True.MinByte() != 0 {
		t.Error("MinByte should be 0")
	}
}

func Test_Value_ToPtr_Cov2(t *testing.T) {
	ptr := issetter.True.ToPtr()
	if ptr == nil || *ptr != issetter.True {
		t.Error("should return pointer to True")
	}
}

func Test_Value_ValueUInt16_Cov2(t *testing.T) {
	if issetter.True.ValueUInt16() != 1 {
		t.Error("True.ValueUInt16() should be 1")
	}
}

func Test_Value_IsNo_Cov2(t *testing.T) {
	if !issetter.False.IsNo() {
		t.Error("False.IsNo should be true")
	}
	if issetter.True.IsNo() {
		t.Error("True.IsNo should be false")
	}
}

func Test_Value_IsWildcardOrBool_False_Cov2(t *testing.T) {
	r := issetter.True.IsWildcardOrBool(false)
	if r {
		t.Error("True.IsWildcardOrBool(false) should be false since it falls through to isBool")
	}
}

func Test_Min_Cov2(t *testing.T) {
	if issetter.Min() != issetter.Uninitialized {
		t.Error("Min should be Uninitialized")
	}
}

func Test_Max_Cov2(t *testing.T) {
	if issetter.Max() != issetter.Wildcard {
		t.Error("Max should be Wildcard")
	}
}

func Test_MinByte_Cov2(t *testing.T) {
	if issetter.MinByte() != 0 {
		t.Error("MinByte should be 0")
	}
}

func Test_MaxByte_Func_Cov2(t *testing.T) {
	if issetter.MaxByte() != issetter.Set.Value() {
		t.Error("MaxByte should be Set value")
	}
}

func Test_RangeNamesCsv_Cov2(t *testing.T) {
	if issetter.RangeNamesCsv() == "" {
		t.Error("should not be empty")
	}
}

func Test_Value_OnlySupportedErr_Empty_Cov2(t *testing.T) {
	err := issetter.True.OnlySupportedErr()
	if err != nil {
		t.Error("empty names should return nil")
	}
}

func Test_Value_OnlySupportedMsgErr_Nil_Cov2(t *testing.T) {
	// all names supported should be nil
	names := []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}
	err := issetter.True.OnlySupportedMsgErr("prefix: ", names...)
	if err != nil {
		t.Error("all supported should be nil")
	}
}
