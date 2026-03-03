package coreinstructiontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	convey "github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// --- IsMatch ---

func Test_StringSearch_IsMatch_Equal_Match(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	result := ss.IsMatch("hello")

	// Assert
	convey.Convey("IsMatch - equal match should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_StringSearch_IsMatch_Equal_NoMatch(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	result := ss.IsMatch("world")

	// Assert
	convey.Convey("IsMatch - equal no match should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_StringSearch_IsMatch_Contains_Match(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Contains,
		Search:        "world",
	}

	// Act
	result := ss.IsMatch("hello world")

	// Assert
	convey.Convey("IsMatch - contains match should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_StringSearch_IsMatch_Contains_NoMatch(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Contains,
		Search:        "xyz",
	}

	// Act
	result := ss.IsMatch("hello world")

	// Assert
	convey.Convey("IsMatch - contains no match should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_StringSearch_IsMatch_NilReceiver(t *testing.T) {
	// Arrange
	var ss *coreinstruction.StringSearch

	// Act
	result := ss.IsMatch("anything")

	// Assert
	convey.Convey("IsMatch - nil receiver should return true (vacuous truth)", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- IsMatchFailed ---

func Test_StringSearch_IsMatchFailed_Match_ReturnsFalse(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	result := ss.IsMatchFailed("hello")

	// Assert
	convey.Convey("IsMatchFailed - match should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_StringSearch_IsMatchFailed_NoMatch_ReturnsTrue(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	result := ss.IsMatchFailed("world")

	// Assert
	convey.Convey("IsMatchFailed - no match should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_StringSearch_IsMatchFailed_NilReceiver(t *testing.T) {
	// Arrange
	var ss *coreinstruction.StringSearch

	// Act
	result := ss.IsMatchFailed("anything")

	// Assert
	convey.Convey("IsMatchFailed - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- IsAllMatch ---

func Test_StringSearch_IsAllMatch_AllContentsMatch(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Contains,
		Search:        "o",
	}

	// Act
	result := ss.IsAllMatch("hello", "world", "foo")

	// Assert
	convey.Convey("IsAllMatch - all contents containing search should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_StringSearch_IsAllMatch_OneContentFails(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Contains,
		Search:        "z",
	}

	// Act
	result := ss.IsAllMatch("hello", "buzz", "world")

	// Assert
	convey.Convey("IsAllMatch - one failing content should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_StringSearch_IsAllMatch_EmptyContents(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	result := ss.IsAllMatch()

	// Assert
	convey.Convey("IsAllMatch - empty contents should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- IsAnyMatchFailed ---

func Test_StringSearch_IsAnyMatchFailed_AllMatch_ReturnsFalse(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Contains,
		Search:        "o",
	}

	// Act
	result := ss.IsAnyMatchFailed("hello", "world", "foo")

	// Assert
	convey.Convey("IsAnyMatchFailed - all matching should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_StringSearch_IsAnyMatchFailed_OneFails_ReturnsTrue(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	result := ss.IsAnyMatchFailed("hello", "world")

	// Assert
	convey.Convey("IsAnyMatchFailed - one failing should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- IsEmpty / IsExist / Has ---

func Test_StringSearch_IsEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var ss *coreinstruction.StringSearch

	// Act & Assert
	convey.Convey("IsEmpty/IsExist/Has - nil receiver checks", t, func() {
		convey.So(ss.IsEmpty(), should.BeTrue)
		convey.So(ss.IsExist(), should.BeFalse)
		convey.So(ss.Has(), should.BeFalse)
	})
}

func Test_StringSearch_IsEmpty_NonNil(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "test",
	}

	// Act & Assert
	convey.Convey("IsEmpty/IsExist/Has - non-nil receiver checks", t, func() {
		convey.So(ss.IsEmpty(), should.BeFalse)
		convey.So(ss.IsExist(), should.BeTrue)
		convey.So(ss.Has(), should.BeTrue)
	})
}

// --- VerifyError ---

func Test_StringSearch_VerifyError_Match_NoError(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	err := ss.VerifyError("hello")

	// Assert
	convey.Convey("VerifyError - match should return nil", t, func() {
		convey.So(err, should.BeNil)
	})
}

func Test_StringSearch_VerifyError_NoMatch_ReturnsError(t *testing.T) {
	// Arrange
	ss := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act
	err := ss.VerifyError("world")

	// Assert
	convey.Convey("VerifyError - no match should return error", t, func() {
		convey.So(err, should.NotBeNil)
	})
}

func Test_StringSearch_VerifyError_NilReceiver(t *testing.T) {
	// Arrange
	var ss *coreinstruction.StringSearch

	// Act
	err := ss.VerifyError("anything")

	// Assert
	convey.Convey("VerifyError - nil receiver should return nil", t, func() {
		convey.So(err, should.BeNil)
	})
}
