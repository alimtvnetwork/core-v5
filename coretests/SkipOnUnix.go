package coretests

import (
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/osconsts"
)

// SkipOnUnix Skip on Unix
func SkipOnUnix(t *testing.T) {
	if osconsts.IsUnixGroup {
		t.Skip(errcore.UnixIgnoreType)
	}
}
