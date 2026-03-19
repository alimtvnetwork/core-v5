package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
)

func Test_QW_MinMaxByte_DifferenceAbsolute_NegDiff(t *testing.T) {
	// For byte type diff < 0 is impossible (unsigned), but we test the branch
	// by having Min > Max which wraps around
	mm := corerange.MinMaxByte{Min: 200, Max: 100}
	_ = mm.DifferenceAbsolute()
}

func Test_QW_Within_StringRangeUint32_Large(t *testing.T) {
	// Cover the branch where finalInt > MaxInt32
	// Use a very large number string
	val, inRange := corerange.IsRangeWithin.StringRangeUint32Default("999999999999")
	_, _ = val, inRange
}
