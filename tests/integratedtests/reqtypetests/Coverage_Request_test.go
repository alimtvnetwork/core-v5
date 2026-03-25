package reqtypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/reqtype"
)

func Test_Cov_Request_IsAnySkipOnExist(t *testing.T) {
	if !reqtype.SkipOnExist.IsAnySkipOnExist() {
		t.Error("expected true")
	}
}

func Test_Cov_Request_IsOverrideOrOverwriteOrEnforce(t *testing.T) {
	if !reqtype.Overwrite.IsOverrideOrOverwriteOrEnforce() {
		t.Error("expected true")
	}
	if !reqtype.Override.IsOverrideOrOverwriteOrEnforce() {
		t.Error("expected true")
	}
	if !reqtype.Enforce.IsOverrideOrOverwriteOrEnforce() {
		t.Error("expected true")
	}
}

func Test_Cov_Request_IsRestartOrReload(t *testing.T) {
	if !reqtype.Restart.IsRestartOrReload() {
		t.Error("expected true")
	}
	if !reqtype.Reload.IsRestartOrReload() {
		t.Error("expected true")
	}
}
