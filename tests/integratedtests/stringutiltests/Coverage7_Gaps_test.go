package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreutils/stringutil"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage7 — stringutil remaining gaps
//
// Target 1: IsEndsWith.go:37-39 — remainingLength < 0
//   Logically unreachable: line 25-27 already returns false when
//   endsWithLength > basePathLength, making remainingLength always >= 0.
//   Dead code.
//
// Target 2: replaceTemplate.go:316-341 — UsingNamerMapOptions both branches
// ══════════════════════════════════════════════════════════════════════════════

type testNamer struct {
	name string
}

func (n testNamer) Name() string {
	return n.name
}

func Test_Cov7_ReplaceTemplate_UsingNamerMapOptions_CurlyKeys(t *testing.T) {
	// Arrange
	namerMap := map[interface{ Name() string }]string{
		testNamer{name: "host"}: "example.com",
		testNamer{name: "port"}: "8080",
	}

	// Act
	result := stringutil.ReplaceTemplate.UsingNamerMapOptions(
		true,
		"https://{host}:{port}/api",
		namerMap,
	)

	// Assert
	convey.Convey("UsingNamerMapOptions with curly brace keys replaces correctly", t, func() {
		convey.So(result, convey.ShouldEqual, "https://example.com:8080/api")
	})
}

func Test_Cov7_ReplaceTemplate_UsingNamerMapOptions_DirectKeys(t *testing.T) {
	// Arrange
	namerMap := map[interface{ Name() string }]string{
		testNamer{name: "HOST"}: "example.com",
	}

	// Act
	result := stringutil.ReplaceTemplate.UsingNamerMapOptions(
		false,
		"https://HOST/api",
		namerMap,
	)

	// Assert
	convey.Convey("UsingNamerMapOptions with direct keys replaces correctly", t, func() {
		convey.So(result, convey.ShouldEqual, "https://example.com/api")
	})
}

// Coverage note: IsEndsWith.go:37-39 (remainingLength < 0) is dead code.
// The guard at line 25 (endsWithLength > basePathLength) makes this
// branch unreachable. Documented as accepted dead-code gap.
