package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/keymk"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage6 — appendStringsWithBaseAnyItems skip-empty branch
//
// Target: keymk/appendStringsWithBaseAnyItems.go:13-14
//   isSkipOnEmpty && item == "" → continue
//
// Exercise via CompileKeys with IsSkipEmptyEntry=true and a sub-key
// that has an empty string in its keyChains.
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_CompileKeys_SkipsEmptyKeyChain(t *testing.T) {
	// Arrange
	mainKey := keymk.NewKey.Default("root", "chain1")
	subKey := keymk.NewKey.Default("sub", "", "val")

	// Act
	result := mainKey.CompileKeys(subKey)

	// Assert
	convey.Convey("CompileKeys skips empty keyChain entries when IsSkipEmptyEntry is true", t, func() {
		convey.So(result, convey.ShouldNotContainSubstring, "--")
		convey.So(result, convey.ShouldContainSubstring, "sub")
		convey.So(result, convey.ShouldContainSubstring, "val")
	})
}
