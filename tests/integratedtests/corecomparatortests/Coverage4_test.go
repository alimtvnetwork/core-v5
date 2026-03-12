package corecomparatortests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
)

// Cover remaining Compare methods not hit by existing tests

func Test_Compare_IsLess_Cov4(t *testing.T) {
	if !corecomparator.LeftLess.IsLess() {
		t.Error("LeftLess should IsLess")
	}
	if corecomparator.Equal.IsLess() {
		t.Error("Equal should not IsLess")
	}
}

func Test_Compare_IsLessEqual_Cov4(t *testing.T) {
	if !corecomparator.LeftLess.IsLessEqual() {
		t.Error("LeftLess should IsLessEqual")
	}
	if !corecomparator.Equal.IsLessEqual() {
		t.Error("Equal should IsLessEqual")
	}
	if corecomparator.LeftGreater.IsLessEqual() {
		t.Error("LeftGreater should not IsLessEqual")
	}
}

func Test_Compare_IsGreater_Cov4(t *testing.T) {
	if !corecomparator.LeftGreater.IsGreater() {
		t.Error("LeftGreater should IsGreater")
	}
	if corecomparator.Equal.IsGreater() {
		t.Error("Equal should not IsGreater")
	}
}

func Test_Compare_IsGreaterEqual_Cov4(t *testing.T) {
	if !corecomparator.LeftGreater.IsGreaterEqual() {
		t.Error("LeftGreater should IsGreaterEqual")
	}
	if !corecomparator.Equal.IsGreaterEqual() {
		t.Error("Equal should IsGreaterEqual")
	}
}

func Test_Compare_IsNameEqual_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsNameEqual("Equal") {
		t.Error("should match Equal name")
	}
	if corecomparator.Equal.IsNameEqual("NotEqual") {
		t.Error("should not match NotEqual")
	}
}

func Test_Compare_ToNumberString_Cov4(t *testing.T) {
	if corecomparator.Equal.ToNumberString() != "0" {
		t.Error("expected 0")
	}
}

func Test_Compare_IsDefined_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsDefined() {
		t.Error("Equal should be defined")
	}
	if corecomparator.Inconclusive.IsDefined() {
		t.Error("Inconclusive should not be defined")
	}
}

func Test_Compare_IsValid_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsValid() {
		t.Error("Equal should be valid")
	}
}

func Test_Compare_IsEqual_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsEqual() {
		t.Error("expected equal")
	}
}

func Test_Compare_IsNotEqual_Cov4(t *testing.T) {
	if !corecomparator.NotEqual.IsNotEqual() {
		t.Error("expected not equal")
	}
}

func Test_Compare_IsNotEqualLogically_Cov4(t *testing.T) {
	if corecomparator.Equal.IsNotEqualLogically() {
		t.Error("Equal should not be logically not-equal")
	}
	if !corecomparator.LeftGreater.IsNotEqualLogically() {
		t.Error("LeftGreater should be logically not-equal")
	}
}

func Test_Compare_IsLeftLess_Cov4(t *testing.T) {
	if !corecomparator.LeftLess.IsLeftLess() {
		t.Error("expected true")
	}
}

func Test_Compare_IsLeftLessEqual_Cov4(t *testing.T) {
	if !corecomparator.LeftLessEqual.IsLeftLessEqual() {
		t.Error("expected true")
	}
}

func Test_Compare_IsLeftLessEqualLogically_Cov4(t *testing.T) {
	if !corecomparator.LeftLess.IsLeftLessEqualLogically() {
		t.Error("expected true")
	}
	if !corecomparator.LeftLessEqual.IsLeftLessEqualLogically() {
		t.Error("expected true")
	}
	if !corecomparator.Equal.IsLeftLessEqualLogically() {
		t.Error("expected true")
	}
}

func Test_Compare_IsLeftGreaterEqualLogically_Cov4(t *testing.T) {
	if !corecomparator.LeftGreater.IsLeftGreaterEqualLogically() {
		t.Error("expected true")
	}
	if !corecomparator.LeftGreaterEqual.IsLeftGreaterEqualLogically() {
		t.Error("expected true")
	}
}

func Test_Compare_IsLeftGreaterOrGreaterEqualOrEqual_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsLeftGreaterOrGreaterEqualOrEqual() {
		t.Error("expected true")
	}
	if !corecomparator.LeftGreater.IsLeftGreaterOrGreaterEqualOrEqual() {
		t.Error("expected true")
	}
	if !corecomparator.LeftGreaterEqual.IsLeftGreaterOrGreaterEqualOrEqual() {
		t.Error("expected true")
	}
}

func Test_Compare_IsInconclusiveOrNotEqual_Cov4(t *testing.T) {
	if !corecomparator.Inconclusive.IsInconclusiveOrNotEqual() {
		t.Error("expected true")
	}
	if !corecomparator.NotEqual.IsInconclusiveOrNotEqual() {
		t.Error("expected true")
	}
	if corecomparator.Equal.IsInconclusiveOrNotEqual() {
		t.Error("expected false")
	}
}

func Test_Compare_IsDefinedProperly_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsDefinedProperly() {
		t.Error("expected true")
	}
}

func Test_Compare_IsAnyOf_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsAnyOf() {
		t.Error("empty should return true")
	}
	if !corecomparator.Equal.IsAnyOf(corecomparator.NotEqual, corecomparator.Equal) {
		t.Error("should find Equal")
	}
	if corecomparator.Equal.IsAnyOf(corecomparator.NotEqual) {
		t.Error("should not find Equal in [NotEqual]")
	}
}

func Test_Compare_NameValue_Cov4(t *testing.T) {
	r := corecomparator.Equal.NameValue()
	if r == "" {
		t.Error("should not be empty")
	}
}

func Test_Compare_CsvStrings_Cov4(t *testing.T) {
	r := corecomparator.Equal.CsvStrings()
	if len(r) != 0 {
		t.Error("empty args should return empty slice")
	}
	r = corecomparator.Equal.CsvStrings(corecomparator.Equal, corecomparator.NotEqual)
	if len(r) != 2 {
		t.Error("expected 2")
	}
}

func Test_Compare_CsvString_Cov4(t *testing.T) {
	r := corecomparator.Equal.CsvString()
	if r != "" {
		t.Error("empty args should return empty")
	}
	r = corecomparator.Equal.CsvString(corecomparator.Equal)
	if r == "" {
		t.Error("should not be empty")
	}
}

func Test_Compare_MarshalJSON_Cov4(t *testing.T) {
	data, err := json.Marshal(corecomparator.Equal)
	if err != nil {
		t.Error("should not error")
	}
	if string(data) == "" {
		t.Error("should not be empty")
	}
}

func Test_Compare_Value_Cov4(t *testing.T) {
	if corecomparator.Equal.Value() != 0 {
		t.Error("expected 0")
	}
}

func Test_Compare_ValueByte_Cov4(t *testing.T) {
	if corecomparator.Equal.ValueByte() != 0 {
		t.Error("expected 0")
	}
}

func Test_Compare_ValueInt_Cov4(t *testing.T) {
	if corecomparator.Equal.ValueInt() != 0 {
		t.Error("expected 0")
	}
}

func Test_Compare_OperatorSymbol_Cov4(t *testing.T) {
	if corecomparator.Equal.OperatorSymbol() != "=" {
		t.Error("expected =")
	}
}

func Test_Compare_OperatorShortForm_Cov4(t *testing.T) {
	if corecomparator.Equal.OperatorShortForm() != "eq" {
		t.Error("expected eq")
	}
}

func Test_Compare_NumberJsonString_Cov4(t *testing.T) {
	r := corecomparator.Equal.NumberJsonString()
	if r != "\"0\"" {
		t.Errorf("expected quoted 0, got %s", r)
	}
}

func Test_Compare_IsAnyNamesOf_Cov4(t *testing.T) {
	if !corecomparator.Equal.IsAnyNamesOf("NotEqual", "Equal") {
		t.Error("should find Equal")
	}
	if corecomparator.Equal.IsAnyNamesOf("NotEqual") {
		t.Error("should not find")
	}
}

func Test_Compare_IsCompareEqualLogically_Cov4(t *testing.T) {
	// it == expectedCompare
	if !corecomparator.Equal.IsCompareEqualLogically(corecomparator.Equal) {
		t.Error("expected true")
	}
	// expectedCompare == NotEqual
	if !corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.NotEqual) {
		t.Error("LeftGreater is logically not-equal")
	}
	// expectedCompare.IsLeftGreaterEqualLogically
	if !corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.LeftGreaterEqual) {
		t.Error("expected true")
	}
	// expectedCompare.IsLeftLessEqualLogically
	if !corecomparator.LeftLess.IsCompareEqualLogically(corecomparator.LeftLessEqual) {
		t.Error("expected true")
	}
	// fallthrough false
	if corecomparator.Inconclusive.IsCompareEqualLogically(corecomparator.LeftGreater) {
		t.Error("expected false")
	}
}

func Test_Compare_OnlySupportedErr_Cov4(t *testing.T) {
	// with message, supported
	err := corecomparator.Equal.OnlySupportedErr("test", corecomparator.Equal)
	if err != nil {
		t.Error("should be nil")
	}
	// with message, not supported
	err = corecomparator.LeftGreater.OnlySupportedErr("test", corecomparator.Equal)
	if err == nil {
		t.Error("should error")
	}
	// empty message delegates to OnlySupportedDirectErr
	err = corecomparator.Equal.OnlySupportedErr("", corecomparator.Equal)
	if err != nil {
		t.Error("should be nil")
	}
}

func Test_Compare_OnlySupportedDirectErr_Cov4(t *testing.T) {
	err := corecomparator.Equal.OnlySupportedDirectErr(corecomparator.Equal)
	if err != nil {
		t.Error("should be nil")
	}
	err = corecomparator.LeftGreater.OnlySupportedDirectErr(corecomparator.Equal)
	if err == nil {
		t.Error("should error")
	}
}

func Test_Min_Cov4(t *testing.T) {
	if corecomparator.Min() != corecomparator.Equal {
		t.Error("expected Equal")
	}
}

func Test_Max_Cov4(t *testing.T) {
	if corecomparator.Max() != corecomparator.NotEqual {
		t.Error("expected NotEqual")
	}
}

func Test_MinLength_Cov4(t *testing.T) {
	if corecomparator.MinLength(3, 5) != 3 {
		t.Error("expected 3")
	}
	if corecomparator.MinLength(5, 3) != 3 {
		t.Error("expected 3")
	}
}

func Test_Ranges_Cov4(t *testing.T) {
	r := corecomparator.Ranges()
	if len(r) == 0 {
		t.Error("should not be empty")
	}
}
