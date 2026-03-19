package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/enums/versionindexes"
)

func Test_Cov_Index_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	idx := versionindexes.Major
	minor := versionindexes.Minor
	if idx.IsAnyEnumsEqual(&minor) {
		t.Error("expected false")
	}
}
