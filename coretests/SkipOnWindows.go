package coretests

import (
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/osconsts"
)

// SkipOnWindows Skip on Windows
func SkipOnWindows(t *testing.T) {
	if osconsts.IsWindows {
		t.Skip(errcore.WindowsIgnoreType)
	}
}
