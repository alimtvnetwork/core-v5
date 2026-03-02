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

// --- StartsWith match ---

func Test_StringCompare_StartsWith_Match(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareStartsWith(false, "hello", "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare StartsWith - matching prefix should match", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}

func Test_StringCompare_StartsWith_NoMatch(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareStartsWith(false, "world", "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare StartsWith - non-prefix should not match", t, func() {
		convey.So(isMatch, should.BeFalse)
	})
}

func Test_StringCompare_StartsWith_IgnoreCase(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareStartsWith(true, "HELLO", "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare StartsWith - ignore case should match", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}

// --- EndsWith match ---

func Test_StringCompare_EndsWith_Match(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEndsWith(false, "world", "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare EndsWith - matching suffix should match", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}

func Test_StringCompare_EndsWith_NoMatch(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEndsWith(false, "hello", "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare EndsWith - non-suffix should not match", t, func() {
		convey.So(isMatch, should.BeFalse)
	})
}

func Test_StringCompare_EndsWith_IgnoreCase(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareEndsWith(true, "WORLD", "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare EndsWith - ignore case should match", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}

// --- Regex match ---

func Test_StringCompare_Regex_Match(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareRegex(`^hello\s\w+$`, "hello world")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare Regex - matching pattern should match", t, func() {
		convey.So(isMatch, should.BeTrue)
	})
}

func Test_StringCompare_Regex_NoMatch(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompareRegex(`^\d+$`, "hello")

	// Act
	isMatch := sc.IsMatch()

	// Assert
	convey.Convey("StringCompare Regex - non-matching pattern should not match", t, func() {
		convey.So(isMatch, should.BeFalse)
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
