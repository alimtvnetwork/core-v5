package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
)

func Test_Value_Methods(t *testing.T) {
	// Arrange
	v := issetter.New(5)

	// Assert
	if v.IsUnSet() {
		t.Error("should be set")
	}
	if !v.IsSet() {
		t.Error("should be set")
	}
	if v.IsUnSetOrUninitialized() {
		t.Error("should not be unset or uninitialized")
	}
	if v.Value() != 5 {
		t.Errorf("expected 5, got %d", v.Value())
	}
}

func Test_NewBool_Verification(t *testing.T) {
	// Arrange
	v := issetter.NewBool(true)

	// Assert
	if v.IsUnSet() {
		t.Error("should be set")
	}
	if !v.BoolValue() {
		t.Error("should be true")
	}
}

func Test_NewMust_Verification(t *testing.T) {
	// Act
	v := issetter.NewMust(10)

	// Assert
	if v.IsUnSet() {
		t.Error("should be set")
	}
	if v.Value() != 10 {
		t.Error("should be 10")
	}
}

func Test_Max_Verification(t *testing.T) {
	// Act
	result := issetter.Max(issetter.New(5), issetter.New(10))

	// Assert
	if result.Value() != 10 {
		t.Errorf("expected 10, got %d", result.Value())
	}
}

func Test_Min_Verification(t *testing.T) {
	// Act
	result := issetter.Min(issetter.New(5), issetter.New(10))

	// Assert
	if result.Value() != 5 {
		t.Errorf("expected 5, got %d", result.Value())
	}
}

func Test_MaxByte_Verification(t *testing.T) {
	// Act
	result := issetter.MaxByte(issetter.New(5), issetter.New(10))

	// Assert
	if result.Value() != 10 {
		t.Errorf("expected 10, got %d", result.Value())
	}
}

func Test_MinByte_Verification(t *testing.T) {
	// Act
	result := issetter.MinByte(issetter.New(5), issetter.New(10))

	// Assert
	if result.Value() != 5 {
		t.Errorf("expected 5, got %d", result.Value())
	}
}

func Test_IsOutOfRange_Verification(t *testing.T) {
	// Act & Assert
	if !issetter.IsOutOfRange(issetter.New(0), issetter.New(5), issetter.New(10)) {
		t.Error("0 should be out of range [5,10]")
	}
	if issetter.IsOutOfRange(issetter.New(7), issetter.New(5), issetter.New(10)) {
		t.Error("7 should be in range [5,10]")
	}
}

func Test_IsCompareResult_Verification(t *testing.T) {
	// Act
	result := issetter.IsCompareResult(issetter.New(5), issetter.New(10))

	// Assert - left < right
	if result.Value() >= 0 {
		t.Error("5 < 10 should return negative")
	}
}

func Test_RangesNamesCsv_Verification(t *testing.T) {
	// Act
	csv := issetter.RangesNamesCsv()

	// Assert
	if csv == "" {
		t.Error("RangesNamesCsv should not be empty")
	}
}

func Test_IntegerEnumRanges_Verification(t *testing.T) {
	// Act
	ranges := issetter.IntegerEnumRanges()

	// Assert
	if len(ranges) == 0 {
		t.Error("IntegerEnumRanges should not be empty")
	}
}

func Test_GetBool_Verification(t *testing.T) {
	// Act
	v := issetter.GetBool(issetter.New(5))

	// Assert
	if !v {
		t.Error("set value should return true")
	}

	v2 := issetter.GetBool(issetter.Value{})
	if v2 {
		t.Error("unset value should return false")
	}
}

func Test_CombinedBooleans_Verification(t *testing.T) {
	// Act
	v1 := issetter.NewBool(true)
	v2 := issetter.NewBool(false)
	result := issetter.CombinedBooleans(v1, v2)

	// Assert
	if result.IsUnSet() {
		t.Error("combined should be set")
	}
}

func Test_NewBooleans_Verification(t *testing.T) {
	// Act
	v := issetter.NewBooleans(true, false)

	// Assert
	if v.IsUnSet() {
		t.Error("should be set")
	}
}

func Test_GetSet_Verification(t *testing.T) {
	// Act
	v := issetter.New(5)
	got := issetter.GetSet(v)

	// Assert
	if got.IsUnSet() {
		t.Error("should be set")
	}
}

func Test_GetSetByte_Verification(t *testing.T) {
	// Act
	v := issetter.New(5)
	got := issetter.GetSetByte(v)

	// Assert
	if got.IsUnSet() {
		t.Error("should be set")
	}
}

func Test_GetSetUnset_Verification(t *testing.T) {
	// Act - set
	v := issetter.New(5)
	got := issetter.GetSetUnset(v)
	if got.IsUnSet() {
		t.Error("set value should return set")
	}

	// Act - unset
	unset := issetter.Value{}
	got2 := issetter.GetSetUnset(unset)
	if got2.IsSet() {
		t.Error("unset value should return unset")
	}
}

func Test_GetSetterByComparing_Verification(t *testing.T) {
	// Act
	v := issetter.New(5)
	result := issetter.GetSetterByComparing(v, issetter.New(10))

	// Assert - should return comparison result
	if result.IsUnSet() {
		t.Error("should be set")
	}
}

func Test_JsonBytes_Verification(t *testing.T) {
	// Act
	v := issetter.New(5)
	bytes := v.JsonBytes()

	// Assert
	if len(bytes) == 0 {
		t.Error("json bytes should not be empty")
	}
}

func Test_GenerateDynamicRangesMap_Verification(t *testing.T) {
	// Call to ensure no panic
	v := issetter.New(5)
	_ = v.String()
}

func Test_ToHashset_Verification(t *testing.T) {
	// Act
	v := issetter.New(5)
	_ = v.Hashset()
}
