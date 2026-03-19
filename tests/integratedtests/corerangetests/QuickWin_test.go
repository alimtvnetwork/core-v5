package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
)

func Test_QW_MinMaxByte_DifferenceAbsolute(t *testing.T) {
	// For byte type diff < 0 is impossible (unsigned), but we still call it
	mm := corerange.MinMaxByte{Min: 200, Max: 100}
	_ = mm.DifferenceAbsolute()
}

func Test_QW_Within_StringRangeUint32_LargeValue(t *testing.T) {
	// Cover the branch where finalInt > MaxInt32
	val, inRange := corerange.IsRangeWithin.StringRangeUint32(0, 4294967295, "3000000000")
	_, _ = val, inRange
}
