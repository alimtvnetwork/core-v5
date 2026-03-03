package coredynamictests

import (
	"testing"

	"github.com/smarty/assertions/should"
	convey "github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coredynamic"
)

// =============================================================================
// LeftRight — IsEmpty
// =============================================================================

func Test_LeftRight_IsEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act & Assert
	convey.Convey("IsEmpty - nil receiver should return true", t, func() {
		convey.So(lr.IsEmpty(), should.BeTrue)
	})
}

func Test_LeftRight_IsEmpty_BothNil(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: nil}

	// Act & Assert
	convey.Convey("IsEmpty - both nil should return true", t, func() {
		convey.So(lr.IsEmpty(), should.BeTrue)
	})
}

func Test_LeftRight_IsEmpty_HasLeft(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "hello", Right: nil}

	// Act & Assert
	convey.Convey("IsEmpty - has left should return false", t, func() {
		convey.So(lr.IsEmpty(), should.BeFalse)
	})
}

func Test_LeftRight_IsEmpty_HasRight(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: 42}

	// Act & Assert
	convey.Convey("IsEmpty - has right should return false", t, func() {
		convey.So(lr.IsEmpty(), should.BeFalse)
	})
}

func Test_LeftRight_IsEmpty_BothSet(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}

	// Act & Assert
	convey.Convey("IsEmpty - both set should return false", t, func() {
		convey.So(lr.IsEmpty(), should.BeFalse)
	})
}

// =============================================================================
// LeftRight — HasLeft / HasRight
// =============================================================================

func Test_LeftRight_HasLeft_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act & Assert
	convey.Convey("HasLeft - nil receiver should return false", t, func() {
		convey.So(lr.HasLeft(), should.BeFalse)
	})
}

func Test_LeftRight_HasLeft_Present(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "hello"}

	// Act & Assert
	convey.Convey("HasLeft - present should return true", t, func() {
		convey.So(lr.HasLeft(), should.BeTrue)
	})
}

func Test_LeftRight_HasRight_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act & Assert
	convey.Convey("HasRight - nil receiver should return false", t, func() {
		convey.So(lr.HasRight(), should.BeFalse)
	})
}

func Test_LeftRight_HasRight_Present(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: 42}

	// Act & Assert
	convey.Convey("HasRight - present should return true", t, func() {
		convey.So(lr.HasRight(), should.BeTrue)
	})
}

// =============================================================================
// LeftRight — IsLeftEmpty / IsRightEmpty
// =============================================================================

func Test_LeftRight_IsLeftEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act & Assert
	convey.Convey("IsLeftEmpty - nil receiver should return true", t, func() {
		convey.So(lr.IsLeftEmpty(), should.BeTrue)
	})
}

func Test_LeftRight_IsLeftEmpty_NilLeft(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: "x"}

	// Act & Assert
	convey.Convey("IsLeftEmpty - nil left should return true", t, func() {
		convey.So(lr.IsLeftEmpty(), should.BeTrue)
	})
}

func Test_LeftRight_IsLeftEmpty_NonNilLeft(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "x"}

	// Act & Assert
	convey.Convey("IsLeftEmpty - non-nil left should return false", t, func() {
		convey.So(lr.IsLeftEmpty(), should.BeFalse)
	})
}

func Test_LeftRight_IsRightEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act & Assert
	convey.Convey("IsRightEmpty - nil receiver should return true", t, func() {
		convey.So(lr.IsRightEmpty(), should.BeTrue)
	})
}

func Test_LeftRight_IsRightEmpty_NonNilRight(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Right: "y"}

	// Act & Assert
	convey.Convey("IsRightEmpty - non-nil right should return false", t, func() {
		convey.So(lr.IsRightEmpty(), should.BeFalse)
	})
}

// =============================================================================
// LeftRight — DeserializeLeft / DeserializeRight nil safety
// =============================================================================

func Test_LeftRight_DeserializeLeft_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	result := lr.DeserializeLeft()

	// Assert
	convey.Convey("DeserializeLeft - nil receiver should return nil", t, func() {
		convey.So(result, should.BeNil)
	})
}

func Test_LeftRight_DeserializeRight_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	result := lr.DeserializeRight()

	// Assert
	convey.Convey("DeserializeRight - nil receiver should return nil", t, func() {
		convey.So(result, should.BeNil)
	})
}

func Test_LeftRight_DeserializeLeft_ValidData(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: map[string]string{"key": "val"}}

	// Act
	result := lr.DeserializeLeft()

	// Assert
	convey.Convey("DeserializeLeft - valid data should return non-nil result", t, func() {
		convey.So(result, should.NotBeNil)
		convey.So(result.HasError(), should.BeFalse)
	})
}

// =============================================================================
// LeftRight — TypeStatus nil safety
// =============================================================================

func Test_LeftRight_TypeStatus_NilReceiver(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	status := lr.TypeStatus()

	// Assert
	convey.Convey("TypeStatus - nil receiver should not panic", t, func() {
		convey.So(status, should.NotBeNil)
	})
}
