package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
)

func Test_Cov_MinMaxByte_CreateRangeInt8(t *testing.T) {
	mmb := &corerange.MinMaxByte{Min: 0, Max: 10}
	r := mmb.CreateRangeInt8("0-10", "-")
	if r == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov_MinMaxByte_CreateRangeInt16(t *testing.T) {
	mmb := &corerange.MinMaxByte{Min: 0, Max: 10}
	r := mmb.CreateRangeInt16("0-10", "-")
	if r == nil {
		t.Error("expected non-nil")
	}
}

func Test_Cov_Within_StringRangeInt32(t *testing.T) {
	val, ok := corerange.Within.StringRangeInt32("100")
	if !ok || val != 100 {
		t.Error("expected 100")
	}
}

func Test_Cov_Within_StringRangeInt16(t *testing.T) {
	val, ok := corerange.Within.StringRangeInt16("100")
	if !ok || val != 100 {
		t.Error("expected 100")
	}
}

func Test_Cov_Within_StringRangeInt8(t *testing.T) {
	val, ok := corerange.Within.StringRangeInt8("50")
	if !ok || val != 50 {
		t.Error("expected 50")
	}
}

func Test_Cov_Within_StringRangeByte(t *testing.T) {
	val, ok := corerange.Within.StringRangeByte("200")
	if !ok || val != 200 {
		t.Error("expected 200")
	}
}

func Test_Cov_Within_StringRangeUint16(t *testing.T) {
	val, ok := corerange.Within.StringRangeUint16("1000")
	if !ok || val != 1000 {
		t.Error("expected 1000")
	}
}

func Test_Cov_Within_StringRangeUint32(t *testing.T) {
	val, ok := corerange.Within.StringRangeUint32("1000")
	if !ok || val != 1000 {
		t.Error("expected 1000")
	}
}

func Test_Cov_Within_StringRangeIntegerDefault(t *testing.T) {
	val, ok := corerange.Within.StringRangeIntegerDefault(0, 100, "50")
	if !ok || val != 50 {
		t.Error("expected 50")
	}
	// below min
	val2, ok2 := corerange.Within.StringRangeIntegerDefault(0, 100, "-5")
	if ok2 || val2 != 0 {
		t.Error("expected 0 for below min")
	}
	// above max
	val3, ok3 := corerange.Within.StringRangeIntegerDefault(0, 100, "200")
	if ok3 || val3 != 100 {
		t.Error("expected 100 for above max")
	}
}

func Test_Cov_Within_StringRangeFloat(t *testing.T) {
	val, ok := corerange.Within.StringRangeFloat(true, 0, 100, "50.5")
	if !ok || val != 50.5 {
		t.Errorf("expected 50.5 got %v", val)
	}
}

func Test_Cov_Within_StringRangeFloatDefault(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloatDefault("50.5")
	if !ok {
		t.Error("expected in range")
	}
}

func Test_Cov_Within_StringRangeFloat64(t *testing.T) {
	val, ok := corerange.Within.StringRangeFloat64(true, 0, 100, "50.5")
	if !ok || val != 50.5 {
		t.Error("expected 50.5")
	}
}

func Test_Cov_Within_StringRangeFloat64Default(t *testing.T) {
	_, ok := corerange.Within.StringRangeFloat64Default("50.5")
	if !ok {
		t.Error("expected in range")
	}
}

func Test_Cov_Within_RangeByteDefault(t *testing.T) {
	val, ok := corerange.Within.RangeByteDefault(100)
	if !ok || val != 100 {
		t.Error("expected 100")
	}
}

func Test_Cov_Within_RangeUint16Default(t *testing.T) {
	val, ok := corerange.Within.RangeUint16Default(1000)
	if !ok || val != 1000 {
		t.Error("expected 1000")
	}
}

func Test_Cov_Within_RangeFloat(t *testing.T) {
	val, ok := corerange.Within.RangeFloat(true, 0, 100, 50)
	if !ok || val != 50 {
		t.Error("expected 50")
	}
	// below min with boundary
	val2, ok2 := corerange.Within.RangeFloat(true, 10, 100, 5)
	if ok2 || val2 != 10 {
		t.Error("expected 10")
	}
	// above max with boundary
	val3, ok3 := corerange.Within.RangeFloat(true, 0, 100, 200)
	if ok3 || val3 != 100 {
		t.Error("expected 100")
	}
	// no boundary
	val4, ok4 := corerange.Within.RangeFloat(false, 0, 100, 200)
	if ok4 || val4 != 200 {
		t.Error("expected 200")
	}
}
