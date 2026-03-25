package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
)

func Test_QW_MergeSlicesOfSlices_Empty(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices()
	if len(result) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_QW_RegexTrimmedSplitNonEmptyAll_Empty(t *testing.T) {
	re := regexp.MustCompile(`\s+`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "")
	_ = result
}

func Test_QW_SplitTrimmedNonEmpty_Empty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty("", ",", -1)
	_ = result
}

func Test_QW_SplitTrimmedNonEmptyAll_Empty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmptyAll("", ",")
	_ = result
}
