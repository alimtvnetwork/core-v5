package corecmptests

import (
	"github.com/alimtvnetwork/core/corecmp"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
)

// The "NotEqual" return at line 14 in Byte/Integer/Integer8/16/32/64
// is unreachable dead code (==, <, > cover all cases).
// We can't truly hit it, but we can exercise the equal/less/greater paths
// and cover the VersionSlice fallback.

func Test_I14_Byte_AllBranches(t *testing.T) {
	if corecmp.Byte(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.Byte(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.Byte(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer_AllBranches(t *testing.T) {
	if corecmp.Integer(5, 5) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.Integer(3, 5) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.Integer(5, 3) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer8_AllBranches(t *testing.T) {
	if corecmp.Integer8(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.Integer8(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.Integer8(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer16_AllBranches(t *testing.T) {
	if corecmp.Integer16(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.Integer16(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.Integer16(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer32_AllBranches(t *testing.T) {
	if corecmp.Integer32(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.Integer32(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.Integer32(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Integer64_AllBranches(t *testing.T) {
	if corecmp.Integer64(1, 1) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.Integer64(1, 2) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.Integer64(2, 1) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_Time_AllBranches(t *testing.T) {
	now := time.Now()
	past := now.Add(-time.Hour)
	future := now.Add(time.Hour)
	if corecmp.Time(now, now) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.Time(past, future) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.Time(future, past) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
}

func Test_I14_VersionSliceByte_AllBranches(t *testing.T) {
	if corecmp.VersionSliceByte(nil, nil) != corecomparator.Equal {
		t.Fatal("expected equal for nil/nil")
	}
	if corecmp.VersionSliceByte(nil, []byte{1}) != corecomparator.NotEqual {
		t.Fatal("expected not equal for nil/non-nil")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 2}) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1, 3}) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.VersionSliceByte([]byte{1, 3}, []byte{1, 2}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
	if corecmp.VersionSliceByte([]byte{1}, []byte{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected left less for shorter")
	}
	if corecmp.VersionSliceByte([]byte{1, 2}, []byte{1}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater for longer")
	}
}

func Test_I14_VersionSliceInteger_AllBranches(t *testing.T) {
	if corecmp.VersionSliceInteger(nil, nil) != corecomparator.Equal {
		t.Fatal("expected equal for nil/nil")
	}
	if corecmp.VersionSliceInteger(nil, []int{1}) != corecomparator.NotEqual {
		t.Fatal("expected not equal for nil/non-nil")
	}
	if corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 2}) != corecomparator.Equal {
		t.Fatal("expected equal")
	}
	if corecmp.VersionSliceInteger([]int{1, 2}, []int{1, 3}) != corecomparator.LeftLess {
		t.Fatal("expected left less")
	}
	if corecmp.VersionSliceInteger([]int{1, 3}, []int{1, 2}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater")
	}
	if corecmp.VersionSliceInteger([]int{1}, []int{1, 2}) != corecomparator.LeftLess {
		t.Fatal("expected left less for shorter")
	}
	if corecmp.VersionSliceInteger([]int{1, 2}, []int{1}) != corecomparator.LeftGreater {
		t.Fatal("expected left greater for longer")
	}
}
