package corecmp

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
)

// The "NotEqual" return at line 14 in Byte/Integer/Integer8/16/32/64
// is unreachable dead code (==, <, > cover all cases).
// We can't truly hit it, but we can exercise the equal/less/greater paths
// and cover the VersionSlice fallback.

func Test_I14_Byte_AllBranches(t *testing.T) {
	if Byte(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if Byte(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if Byte(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer_AllBranches(t *testing.T) {
	if Integer(5, 5) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if Integer(3, 5) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if Integer(5, 3) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer8_AllBranches(t *testing.T) {
	if Integer8(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if Integer8(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if Integer8(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer16_AllBranches(t *testing.T) {
	if Integer16(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if Integer16(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if Integer16(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer32_AllBranches(t *testing.T) {
	if Integer32(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if Integer32(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if Integer32(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer64_AllBranches(t *testing.T) {
	if Integer64(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if Integer64(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if Integer64(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Time_AllBranches(t *testing.T) {
	now := time.Now()
	past := now.Add(-time.Hour)
	future := now.Add(time.Hour)
	if Time(now, now) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if Time(past, future) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if Time(future, past) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_VersionSliceByte_AllBranches(t *testing.T) {
	if VersionSliceByte(nil, nil) != corecomparator.Equal {
		t.Fatal("expected equal for nil/nil")
	}
	if VersionSliceByte(nil, []byte{1}) != corecomparator.NotEqual {
		t.Fatal("expected not equal for nil/non-nil")
	}
	if VersionSliceByte([]byte{1, 2}, []byte{1, 2}) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if VersionSliceByte([]byte{1, 2}, []byte{1, 3}) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if VersionSliceByte([]byte{1, 3}, []byte{1, 2}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
	if VersionSliceByte([]byte{1}, []byte{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected left less for shorter")
	}
	if VersionSliceByte([]byte{1, 2}, []byte{1}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater for longer")
	}
}

func Test_I14_VersionSliceInteger_AllBranches(t *testing.T) {
	if VersionSliceInteger(nil, nil) != corecomparator.Equal {
		t.Fatal("expected equal for nil/nil")
	}
	if VersionSliceInteger(nil, []int{1}) != corecomparator.NotEqual {
		t.Fatal("expected not equal for nil/non-nil")
	}
	if VersionSliceInteger([]int{1, 2}, []int{1, 2}) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if VersionSliceInteger([]int{1, 2}, []int{1, 3}) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if VersionSliceInteger([]int{1, 3}, []int{1, 2}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
	if VersionSliceInteger([]int{1}, []int{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected left less for shorter")
	}
	if VersionSliceInteger([]int{1, 2}, []int{1}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater for longer")
	}
}
