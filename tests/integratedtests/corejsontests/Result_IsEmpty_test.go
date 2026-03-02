package corejsontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/corejson"
)

// Test_Result_IsEmpty_EmptyBytes — IsEmpty returns true for empty bytes
func Test_Result_IsEmpty_EmptyBytes(t *testing.T) {
	// Arrange
	result := corejson.NewResult.UsingBytes([]byte{})

	// Act
	isEmpty := result.IsEmpty()

	// Assert
	convey.Convey("IsEmpty - empty bytes should return true", t, func() {
		convey.So(isEmpty, should.BeTrue)
	})
}

// Test_Result_IsEmpty_NilReceiver — IsEmpty returns true for nil receiver
func Test_Result_IsEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var nilResult *corejson.Result

	// Act
	isEmpty := nilResult.IsEmpty()

	// Assert
	convey.Convey("IsEmpty - nil receiver should return true", t, func() {
		convey.So(isEmpty, should.BeTrue)
	})
}

// Test_Result_IsEmpty_ValidBytes — IsEmpty returns false for valid bytes
func Test_Result_IsEmpty_ValidBytes(t *testing.T) {
	// Arrange
	result := corejson.New(map[string]string{"key": "value"})

	// Act
	isEmpty := result.IsEmpty()

	// Assert
	convey.Convey("IsEmpty - valid bytes should return false", t, func() {
		convey.So(isEmpty, should.BeFalse)
	})
}
