package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/enums/versionindexes"
)

func Test_Cov_Index_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	idx := versionindexes.Major
	if idx.IsAnyEnumsEqual(versionindexes.Minor) {
		t.Error("expected false")
	}
}
