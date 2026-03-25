package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
)

func Test_Cov_MergeSlicesOfSlices(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{"b", "c"})
	if len(result) != 3 {
		t.Errorf("expected 3 got %d", len(result))
	}
	// with nil slice
	result2 := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{})
	if len(result2) != 1 {
		t.Errorf("expected 1 got %d", len(result2))
	}
	// empty
	result3 := stringslice.MergeSlicesOfSlices()
	if len(result3) != 0 {
		t.Error("expected 0")
	}
}

func Test_Cov_SplitTrimmedNonEmptyAll(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmptyAll("a , b , c", ",")
	if len(result) != 3 {
		t.Errorf("expected 3 got %d", len(result))
	}
}

func Test_Cov_SplitTrimmedNonEmpty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", -1)
	if len(result) != 3 {
		t.Errorf("expected 3 got %d", len(result))
	}
}

func Test_Cov_RegexTrimmedSplitNonEmptyAll(t *testing.T) {
	re := regexp.MustCompile(`[,;]`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "a , b ; c")
	if len(result) != 3 {
		t.Errorf("expected 3 got %d", len(result))
	}
}
