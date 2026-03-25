package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
)

func Test_QW_RangesNotMeet_EmptyReqs(t *testing.T) {
	// Covers start() and end() with empty slice
	result := reqtype.RangesNotMeet("msg")
	if result != "" {
		t.Fatal("expected empty for no requests")
	}
}

func Test_QW_RangesString_EmptyReqs(t *testing.T) {
	result := reqtype.RangesString(",")
	if result != "" {
		t.Fatal("expected empty")
	}
}
