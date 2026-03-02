package coreinstructiontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coreinstruction"
)

// --- Equal match ---

func Test_StringCompare_Equal_Match(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEqual("hello", "hello")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare Equal - identical strings should match", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}

// --- Equal no match ---

func Test_StringCompare_Equal_NoMatch(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEqual("hello", "world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare Equal - different strings should not match", t, func() {
		convey.So(isMatch, should.BeFalse)
	})
}

// --- Case sensitivity ---

func Test_StringCompare_CaseSensitive(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEqual("Hello", "hello")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare Equal - case-sensitive should not match different cases", t, func() {
		convey.So(isMatch, should.BeFalse)
	})
}

// --- Contains match ---

func Test_StringCompare_Contains_Match(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareContains(false, "world", "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare Contains - substring should match", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}

// --- Nil receiver ---

func Test_StringCompare_NilReceiver_IsMatch(t *testing.T) {
	// Arrange
	var sc *coreinstruction.StringCompare

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare - nil receiver IsMatch should return true (vacuous truth)", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}
