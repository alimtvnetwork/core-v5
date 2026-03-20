package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coreversion"
)

// Test_Cov5_VersionDisplayMajorMinorPatch_InvalidPatch tests the IsPatchInvalid branch.
func Test_Cov5_VersionDisplayMajorMinorPatch_InvalidPatch(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 1,
		VersionMinor: 2,
		VersionPatch: -1, // invalid patch
	}

	// Act
	actual := v.VersionDisplayMajorMinorPatch()

	// Assert — should fall back to MajorMinor display
	expected := "v1.2"
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"VersionDisplayMajorMinorPatch with invalid patch should fallback",
		actual, expected,
	)
}

// Test_Cov5_Major_Compare tests the Major() comparison method.
func Test_Cov5_Major_Compare(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 3,
		VersionMinor: 0,
	}

	// Act
	result := v.Major(3)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"Major(3) on version with major=3 should be Equal",
		result.IsEqual(), true,
	)
}

// Test_Cov5_HasDeductUsingNilNess_DeadCode documents that line 20-22
// in hasDeductUsingNilNess.go is unreachable dead code.
// After checking (nil,nil), (non-nil,nil), (nil,non-nil), the remaining
// `if left == nil || right == nil` can never be true.
// This is documented in spec/13-app-issues/golang/41-unreachable-branches-corejson-coreonce.md.
