package issettertests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/issetter"
)

// ── GetBool / GetSet ──

func Test_GetBool_Coverage(t *testing.T) {
	if issetter.GetBool(true) != issetter.True {
		t.Error("true should return True")
	}
	if issetter.GetBool(false) != issetter.False {
		t.Error("false should return False")
	}
}

func Test_GetSet_Coverage(t *testing.T) {
	result := issetter.GetSet(true, issetter.Set, issetter.Unset)
	if result != issetter.Set {
		t.Error("true condition should return Set")
	}

	result = issetter.GetSet(false, issetter.Set, issetter.Unset)
	if result != issetter.Unset {
		t.Error("false condition should return Unset")
	}
}

// ── Value constants and basic checks ──

func Test_Value_Constants(t *testing.T) {
	if issetter.Uninitialized != 0 {
		t.Error("Uninitialized should be 0")
	}
	if issetter.True != 1 {
		t.Error("True should be 1")
	}
	if issetter.False != 2 {
		t.Error("False should be 2")
	}
}

func Test_Value_BoolChecks(t *testing.T) {
	if !issetter.True.IsTrue() {
		t.Error("True.IsTrue")
	}
	if !issetter.False.IsFalse() {
		t.Error("False.IsFalse")
	}
	if !issetter.Set.IsSet() {
		t.Error("Set.IsSet")
	}
	if !issetter.Unset.IsUnset() {
		t.Error("Unset.IsUnset")
	}
	if !issetter.Wildcard.IsWildcard() {
		t.Error("Wildcard.IsWildcard")
	}
	if !issetter.Uninitialized.IsUninitialized() {
		t.Error("Uninitialized.IsUninitialized")
	}
}

func Test_Value_LogicalGroups(t *testing.T) {
	if !issetter.True.IsOn() {
		t.Error("True should be On")
	}
	if !issetter.Set.IsOn() {
		t.Error("Set should be On")
	}
	if !issetter.False.IsOff() {
		t.Error("False should be Off")
	}
	if !issetter.Unset.IsOff() {
		t.Error("Unset should be Off")
	}
	if !issetter.True.IsSuccess() {
		t.Error("True should be Success")
	}
	if !issetter.False.IsFailed() {
		t.Error("False should be Failed")
	}
	if !issetter.True.IsAccept() {
		t.Error("True should be Accept")
	}
	if !issetter.False.IsReject() {
		t.Error("False should be Reject")
	}
	if !issetter.Uninitialized.IsLater() {
		t.Error("Uninitialized should be Later")
	}
	if !issetter.Wildcard.IsAsk() {
		t.Error("Wildcard should be Ask")
	}
	if !issetter.Uninitialized.IsIndeterminate() {
		t.Error("Uninitialized should be Indeterminate")
	}
	if !issetter.Uninitialized.IsSkip() {
		t.Error("Uninitialized should be Skip")
	}
}

func Test_Value_DefinedUndefined(t *testing.T) {
	if !issetter.True.IsDefinedLogically() {
		t.Error("True should be defined")
	}
	if issetter.Uninitialized.IsDefinedLogically() {
		t.Error("Uninitialized should not be defined")
	}
	if !issetter.Uninitialized.IsUndefinedLogically() {
		t.Error("Uninitialized should be undefined")
	}
}

func Test_Value_Conversions(t *testing.T) {
	if issetter.True.Value() != 1 {
		t.Error("True.Value() should be 1")
	}
	if issetter.True.ValueByte() != 1 {
		t.Error("True.ValueByte() should be 1")
	}
	if issetter.True.ValueInt() != 1 {
		t.Error("True.ValueInt() should be 1")
	}
	if issetter.True.ValueInt8() != 1 {
		t.Error("True.ValueInt8() should be 1")
	}
	if issetter.True.ValueInt16() != 1 {
		t.Error("True.ValueInt16() should be 1")
	}
	if issetter.True.ValueInt32() != 1 {
		t.Error("True.ValueInt32() should be 1")
	}
	if issetter.True.ValueUInt16() != 1 {
		t.Error("True.ValueUInt16() should be 1")
	}
}

func Test_Value_String(t *testing.T) {
	if issetter.True.String() != "True" {
		t.Error("True.String() should be True")
	}
	if issetter.True.Name() == "" {
		t.Error("Name should not be empty")
	}
	if issetter.True.NameValue() == "" {
		t.Error("NameValue should not be empty")
	}
	if issetter.True.ValueString() == "" {
		t.Error("ValueString should not be empty")
	}
	if issetter.True.StringValue() == "" {
		t.Error("StringValue should not be empty")
	}
	if issetter.True.ToNumberString() == "" {
		t.Error("ToNumberString should not be empty")
	}
}

func Test_Value_IsNot(t *testing.T) {
	if !issetter.True.IsNot(issetter.False) {
		t.Error("True is not False")
	}
	if issetter.True.IsNot(issetter.True) {
		t.Error("True is True")
	}
}

func Test_Value_Init(t *testing.T) {
	if !issetter.True.IsInit() {
		t.Error("True should be Init")
	}
	if issetter.Uninitialized.IsInit() {
		t.Error("Uninitialized should not be Init")
	}
	if !issetter.True.IsInitBoolean() {
		t.Error("True should be InitBoolean")
	}
	if !issetter.False.IsInitBoolean() {
		t.Error("False should be InitBoolean")
	}
	if !issetter.True.IsInitBooleanWild() {
		t.Error("True should be InitBooleanWild")
	}
	if !issetter.Set.IsInitSet() {
		t.Error("Set should be InitSet")
	}
	if !issetter.Set.IsInitSetWild() {
		t.Error("Set should be InitSetWild")
	}
}

func Test_Value_ToBooleanSetUnset(t *testing.T) {
	if issetter.Set.ToBooleanValue() != issetter.True {
		t.Error("Set should convert to True")
	}
	if issetter.Unset.ToBooleanValue() != issetter.False {
		t.Error("Unset should convert to False")
	}
	if issetter.True.ToSetUnsetValue() != issetter.Set {
		t.Error("True should convert to Set")
	}
	if issetter.False.ToSetUnsetValue() != issetter.Unset {
		t.Error("False should convert to Unset")
	}
}

func Test_Value_BooleanOp(t *testing.T) {
	if !issetter.True.Boolean() {
		t.Error("True.Boolean should be true")
	}
	if issetter.False.Boolean() {
		t.Error("False.Boolean should be false")
	}
	if !issetter.True.IsYes() {
		t.Error("True.IsYes should be true")
	}
}

func Test_Value_WildcardApply_Cov(t *testing.T) {
	if !issetter.Wildcard.WildcardApply(true) {
		t.Error("Wildcard.WildcardApply(true) should return true")
	}
	if !issetter.True.WildcardApply(false) {
		t.Error("True.WildcardApply should return true")
	}
}

func Test_Value_OrBool(t *testing.T) {
	if !issetter.True.OrBool(false) {
		t.Error("True.OrBool(false) should be true")
	}
	if !issetter.Wildcard.OrBool(true) {
		t.Error("Wildcard.OrBool(true) should return true")
	}
}

func Test_Value_AndBool(t *testing.T) {
	if !issetter.True.AndBool(true) {
		t.Error("True.AndBool(true) should be true")
	}
	if issetter.True.AndBool(false) {
		t.Error("True.AndBool(false) should be false")
	}
}

func Test_Value_And(t *testing.T) {
	result := issetter.True.And(issetter.True)
	if result != issetter.True {
		t.Error("True.And(True) should be True")
	}
	result = issetter.Wildcard.And(issetter.False)
	if result != issetter.False {
		t.Error("Wildcard.And(False) should be False")
	}
}

func Test_Value_ToByteCondition(t *testing.T) {
	if issetter.True.ToByteCondition(1, 0, 255) != 1 {
		t.Error("True should return trueVal")
	}
	if issetter.False.ToByteCondition(1, 0, 255) != 0 {
		t.Error("False should return falseVal")
	}
	if issetter.Wildcard.ToByteCondition(1, 0, 255) != 255 {
		t.Error("Wildcard should return invalid")
	}
}

func Test_Value_ToByteConditionWithWildcard(t *testing.T) {
	if issetter.Wildcard.ToByteConditionWithWildcard(99, 1, 0, 255) != 99 {
		t.Error("Wildcard should return wildcard val")
	}
	if issetter.True.ToByteConditionWithWildcard(99, 1, 0, 255) != 1 {
		t.Error("True should return trueVal")
	}
}

func Test_Value_IsValid(t *testing.T) {
	if !issetter.True.IsValid() {
		t.Error("True should be valid")
	}
	if !issetter.Uninitialized.IsInvalid() {
		t.Error("Uninitialized should be invalid")
	}
}

func Test_Value_Format(t *testing.T) {
	result := issetter.True.Format("{name}={value}")
	if result == "" {
		t.Error("Format should not be empty")
	}
}

func Test_Value_EnumType(t *testing.T) {
	if issetter.True.EnumType() == nil {
		t.Error("EnumType should not be nil")
	}
}

func Test_Value_AllNameValues(t *testing.T) {
	names := issetter.True.AllNameValues()
	if len(names) == 0 {
		t.Error("AllNameValues should not be empty")
	}
}

func Test_Value_RangeNamesCsv(t *testing.T) {
	if issetter.True.RangeNamesCsv() == "" {
		t.Error("RangeNamesCsv should not be empty")
	}
}

func Test_Value_MinMaxAny(t *testing.T) {
	min, max := issetter.True.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny should not return nil")
	}
}

func Test_Value_MinMaxStrings(t *testing.T) {
	if issetter.True.MinValueString() == "" {
		t.Error("MinValueString should not be empty")
	}
	if issetter.True.MaxValueString() == "" {
		t.Error("MaxValueString should not be empty")
	}
}

func Test_Value_IntegerEnumRanges(t *testing.T) {
	ranges := issetter.True.IntegerEnumRanges()
	if len(ranges) == 0 {
		t.Error("IntegerEnumRanges should not be empty")
	}
}

func Test_Value_RangesDynamicMap(t *testing.T) {
	m := issetter.True.RangesDynamicMap()
	if len(m) == 0 {
		t.Error("RangesDynamicMap should not be empty")
	}
}

func Test_Value_IsNameEqual(t *testing.T) {
	if !issetter.True.IsNameEqual("True") {
		t.Error("True.IsNameEqual(True) should be true")
	}
}

func Test_Value_IsAnyNamesOf(t *testing.T) {
	if !issetter.True.IsAnyNamesOf("False", "True") {
		t.Error("should find True")
	}
}

func Test_Value_JSON(t *testing.T) {
	data, err := json.Marshal(issetter.True)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	if len(data) == 0 {
		t.Error("MarshalJSON should not be empty")
	}

	var v issetter.Value
	err = json.Unmarshal([]byte(`"True"`), &v)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}
	if v != issetter.True {
		t.Error("should unmarshal to True")
	}
}

func Test_Value_LazyEvaluateBool(t *testing.T) {
	v := issetter.Uninitialized
	called := v.LazyEvaluateBool(func() {})
	if !called {
		t.Error("should be called on Uninitialized")
	}

	called = v.LazyEvaluateBool(func() {})
	if called {
		t.Error("should not be called again")
	}
}

func Test_Value_LazyEvaluateSet(t *testing.T) {
	v := issetter.Uninitialized
	called := v.LazyEvaluateSet(func() {})
	if !called {
		t.Error("should be called on Uninitialized")
	}

	called = v.LazyEvaluateSet(func() {})
	if called {
		t.Error("should not be called again on Set")
	}
}

func Test_Value_GetSetBoolOnInvalid(t *testing.T) {
	v := issetter.Uninitialized
	result := v.GetSetBoolOnInvalid(true)
	if !result {
		t.Error("should return true after setting")
	}
}

func Test_Value_GetSetBoolOnInvalidFunc(t *testing.T) {
	v := issetter.Uninitialized
	result := v.GetSetBoolOnInvalidFunc(func() bool { return false })
	if result {
		t.Error("should return false after setting")
	}
}

func Test_Value_IsTrueOrSet(t *testing.T) {
	if !issetter.True.IsTrueOrSet() {
		t.Error("True.IsTrueOrSet should be true")
	}
	if !issetter.Set.IsTrueOrSet() {
		t.Error("Set.IsTrueOrSet should be true")
	}
	if issetter.False.IsTrueOrSet() {
		t.Error("False.IsTrueOrSet should be false")
	}
}

func Test_Value_HasInitialized(t *testing.T) {
	if !issetter.True.HasInitialized() {
		t.Error("True should be initialized")
	}
	if issetter.Uninitialized.HasInitialized() {
		t.Error("Uninitialized should not be initialized")
	}
}

func Test_Value_HasInitializedAndSet(t *testing.T) {
	if !issetter.Set.HasInitializedAndSet() {
		t.Error("Set should be initialized and set")
	}
}

func Test_Value_HasInitializedAndTrue(t *testing.T) {
	if !issetter.True.HasInitializedAndTrue() {
		t.Error("True should be initialized and true")
	}
}

func Test_Value_IsOnOffLogically(t *testing.T) {
	if !issetter.True.IsOnLogically() {
		t.Error("True should be on logically")
	}
	if !issetter.False.IsOffLogically() {
		t.Error("False should be off logically")
	}
}

func Test_Value_IsAcceptedRejected(t *testing.T) {
	if !issetter.True.IsAccepted() {
		t.Error("True should be accepted")
	}
	if !issetter.False.IsRejected() {
		t.Error("False should be rejected")
	}
}

func Test_Value_IsUnSetOrUninitialized(t *testing.T) {
	if !issetter.Uninitialized.IsUnSetOrUninitialized() {
		t.Error("Uninitialized should be unset or uninitialized")
	}
	if !issetter.Unset.IsUnSetOrUninitialized() {
		t.Error("Unset should be unset or uninitialized")
	}
	if issetter.True.IsUnSetOrUninitialized() {
		t.Error("True should not be unset or uninitialized")
	}
}

func Test_Value_WildcardValueApply(t *testing.T) {
	if !issetter.Wildcard.WildcardValueApply(issetter.True) {
		t.Error("Wildcard should pass through input")
	}
}

func Test_Value_OrValue(t *testing.T) {
	if !issetter.True.OrValue(issetter.False) {
		t.Error("True.OrValue(False) should be true")
	}
}

func Test_Value_IsWildcardOrBool(t *testing.T) {
	if !issetter.Wildcard.IsWildcardOrBool(false) {
		t.Error("Wildcard should always return true")
	}
}

func Test_Value_IsDefinedBoolean(t *testing.T) {
	if !issetter.True.IsDefinedBoolean() {
		t.Error("True should be defined boolean")
	}
	if issetter.Set.IsDefinedBoolean() {
		t.Error("Set should not be defined boolean")
	}
}

func Test_Value_IsValueEqual(t *testing.T) {
	if !issetter.True.IsValueEqual(1) {
		t.Error("True.IsValueEqual(1) should be true")
	}
}

func Test_Value_IsByteValueEqual(t *testing.T) {
	if !issetter.True.IsByteValueEqual(1) {
		t.Error("True.IsByteValueEqual(1) should be true")
	}
}

func Test_Value_OnlySupportedErr(t *testing.T) {
	names := []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}
	err := issetter.True.OnlySupportedErr(names...)
	if err != nil {
		t.Errorf("all supported should return nil, got: %v", err)
	}
}

func Test_Value_OnlySupportedMsgErr(t *testing.T) {
	err := issetter.True.OnlySupportedMsgErr("test: ", "NonExistent")
	if err == nil {
		t.Error("unsupported should return error")
	}
}

func Test_Value_MinMaxInt(t *testing.T) {
	if issetter.True.MaxInt() < issetter.True.MinInt() {
		t.Error("MaxInt should be >= MinInt")
	}
}

func Test_Value_Initialized(t *testing.T) {
	if !issetter.True.IsInitialized() {
		t.Error("True should be initialized")
	}
	if issetter.Uninitialized.IsInitialized() {
		t.Error("Uninitialized should not be initialized")
	}
}
