package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
)

// ==========================================
// Additional method coverage
// ==========================================

func Test_Request_IsNone(t *testing.T) {
	if !reqtype.Invalid.IsNone() {
		t.Error("Invalid should be None")
	}
	if reqtype.Create.IsNone() {
		t.Error("Create should not be None")
	}
}

func Test_Request_IsStopEnableStart(t *testing.T) {
	if reqtype.Create.IsStopEnableStart() {
		t.Error("should return false")
	}
}

func Test_Request_IsStopDisable(t *testing.T) {
	if reqtype.Create.IsStopDisable() {
		t.Error("should return false")
	}
}

func Test_Request_IsUndefined(t *testing.T) {
	if !reqtype.Invalid.IsUndefined() {
		t.Error("Invalid should be undefined")
	}
}

func Test_Request_ValueUInt16(t *testing.T) {
	r := reqtype.Create.ValueUInt16()
	if r != 1 {
		t.Errorf("expected 1, got %d", r)
	}
}

func Test_Request_IntegerEnumRanges(t *testing.T) {
	r := reqtype.Create.IntegerEnumRanges()
	if len(r) == 0 {
		t.Error("should return non-empty")
	}
}

func Test_Request_MinMaxAny(t *testing.T) {
	min, max := reqtype.Create.MinMaxAny()
	if min == nil || max == nil {
		t.Error("should return non-nil")
	}
}

func Test_Request_MinMaxValueString(t *testing.T) {
	if reqtype.Create.MinValueString() == "" {
		t.Error("should return non-empty")
	}
	if reqtype.Create.MaxValueString() == "" {
		t.Error("should return non-empty")
	}
}

func Test_Request_MinMaxInt(t *testing.T) {
	_ = reqtype.Create.MinInt()
	_ = reqtype.Create.MaxInt()
}

func Test_Request_RangesDynamicMap(t *testing.T) {
	m := reqtype.Create.RangesDynamicMap()
	if len(m) == 0 {
		t.Error("should return non-empty map")
	}
}

func Test_Request_IsNotOverrideOrOverwriteOrEnforce(t *testing.T) {
	if !reqtype.Create.IsNotOverrideOrOverwriteOrEnforce() {
		t.Error("Create should not match override group")
	}
	if reqtype.Override.IsNotOverrideOrOverwriteOrEnforce() {
		t.Error("Override should match override group")
	}
}

func Test_Request_IsOverwrite(t *testing.T) {
	if !reqtype.Overwrite.IsOverwrite() {
		t.Error("Overwrite should match")
	}
}

func Test_Request_IsOverride(t *testing.T) {
	if !reqtype.Override.IsOverride() {
		t.Error("Override should match")
	}
}

func Test_Request_IsEnforce(t *testing.T) {
	if !reqtype.Enforce.IsEnforce() {
		t.Error("Enforce should match")
	}
}

func Test_Request_IsByteValueEqual(t *testing.T) {
	if !reqtype.Create.IsByteValueEqual(byte(reqtype.Create)) {
		t.Error("should be equal")
	}
}

func Test_Request_IsValueEqual(t *testing.T) {
	if !reqtype.Create.IsValueEqual(byte(reqtype.Create)) {
		t.Error("should be equal")
	}
}

func Test_Request_IsAnyValuesEqual(t *testing.T) {
	if !reqtype.Create.IsAnyValuesEqual(byte(reqtype.Create), byte(reqtype.Read)) {
		t.Error("should match one value")
	}
	if reqtype.Create.IsAnyValuesEqual(byte(reqtype.Read), byte(reqtype.Update)) {
		t.Error("should not match")
	}
}

func Test_Request_IsNameEqual(t *testing.T) {
	name := reqtype.Create.Name()
	if !reqtype.Create.IsNameEqual(name) {
		t.Error("should match own name")
	}
}

func Test_Request_IsAnyNamesOf(t *testing.T) {
	name := reqtype.Create.Name()
	if !reqtype.Create.IsAnyNamesOf(name, "Other") {
		t.Error("should match one name")
	}
	if reqtype.Create.IsAnyNamesOf("Other", "More") {
		t.Error("should not match")
	}
}

func Test_Request_ValueConversions(t *testing.T) {
	_ = reqtype.Create.ValueInt8()
	_ = reqtype.Create.ValueInt16()
	_ = reqtype.Create.ValueInt32()
}

func Test_Request_IsOnExistOrSkipOnNonExistLogically(t *testing.T) {
	if !reqtype.ExistCheck.IsOnExistOrSkipOnNonExistLogically() {
		t.Error("ExistCheck should match")
	}
}

func Test_Request_IsReadOrUpdateLogically(t *testing.T) {
	if !reqtype.Read.IsReadOrUpdateLogically() {
		t.Error("Read should match")
	}
}

func Test_Request_IsRestartOrReload(t *testing.T) {
	if !reqtype.Restart.IsRestartOrReload() {
		t.Error("Restart should match")
	}
	if !reqtype.Reload.IsRestartOrReload() {
		t.Error("Reload should match")
	}
}

func Test_Request_AllNameValues(t *testing.T) {
	names := reqtype.Create.AllNameValues()
	if len(names) == 0 {
		t.Error("should return non-empty")
	}
}

func Test_Request_OnlySupportedErr(t *testing.T) {
	// Passing all as supported should be nil
	ranges := reqtype.BasicEnumImpl.StringRanges()
	err := reqtype.Create.OnlySupportedErr(ranges...)
	if err != nil {
		t.Error("all supported should return nil")
	}
}

// ==========================================
// RangesOnlySupportedFor
// ==========================================

func Test_RangesOnlySupportedFor_Empty(t *testing.T) {
	err := reqtype.RangesOnlySupportedFor("msg")
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_RangesOnlySupportedFor_NonEmpty(t *testing.T) {
	err := reqtype.RangesOnlySupportedFor("msg", reqtype.Create, reqtype.Read)
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// RangesString / RangesStrings
// ==========================================

func Test_RangesString(t *testing.T) {
	r := reqtype.RangesString(", ", reqtype.Create, reqtype.Read)
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RangesStrings(t *testing.T) {
	r := reqtype.RangesStrings(reqtype.Create, reqtype.Read)
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

func Test_RangesStringDefaultJoiner(t *testing.T) {
	r := reqtype.RangesStringDefaultJoiner(reqtype.Create, reqtype.Read)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// Min / Max
// ==========================================

func Test_Request_Min(t *testing.T) {
	m := reqtype.Min()
	if m != reqtype.Invalid {
		t.Errorf("expected Invalid")
	}
}

func Test_Request_Max(t *testing.T) {
	m := reqtype.Max()
	if m == reqtype.Invalid {
		t.Error("max should not be Invalid")
	}
}

// ==========================================
// RangesNotMeet / RangesNotMeetError / RangesNotSupportedFor
// ==========================================

func Test_RangesNotMeet(t *testing.T) {
	r := reqtype.RangesNotMeet("msg", reqtype.Create, reqtype.Read)
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RangesNotMeetError(t *testing.T) {
	err := reqtype.RangesNotMeetError("msg", reqtype.Create, reqtype.Read)
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RangesNotSupportedFor_Empty(t *testing.T) {
	err := reqtype.RangesNotSupportedFor("msg")
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_RangesNotSupportedFor_NonEmpty(t *testing.T) {
	err := reqtype.RangesNotSupportedFor("msg", reqtype.Create)
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// RangesInvalidErr
// ==========================================

func Test_RangesInvalidErr(t *testing.T) {
	err := reqtype.RangesInvalidErr("msg", reqtype.Create)
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// RangesInBetween
// ==========================================

func Test_RangesInBetween(t *testing.T) {
	r := reqtype.RangesInBetween(reqtype.Create, reqtype.Read)
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// Start / End
// ==========================================

func Test_Request_Start(t *testing.T) {
	s := reqtype.StartRequest()
	if s != reqtype.Invalid {
		t.Errorf("expected Invalid for start")
	}
}

func Test_Request_End(t *testing.T) {
	e := reqtype.EndRequest()
	if e == reqtype.Invalid {
		t.Error("end should not be Invalid")
	}
}

// ==========================================
// ResultStatus
// ==========================================

func Test_ResultStatus(t *testing.T) {
	rs := reqtype.ResultStatus{}
	if rs.HasError {
		t.Error("default should not have error")
	}
}
