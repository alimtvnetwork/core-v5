package corejsontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/corejson"
)

// Test_Result_IsEqual_SameContent — IsEqual returns true for identical content
func Test_Result_IsEqual_SameContent(t *testing.T) {
	// Arrange
	data := map[string]string{"key": "value"}
	resultA := corejson.New(data)
	resultB := corejson.New(data)

	// Act
	isEqual := resultA.IsEqual(resultB)

	// Assert
	convey.Convey("IsEqual - same content should return true", t, func() {
		convey.So(isEqual, should.BeTrue)
	})
}

// Test_Result_IsEqual_DifferentContent — IsEqual returns false for different bytes
func Test_Result_IsEqual_DifferentContent(t *testing.T) {
	// Arrange
	resultA := corejson.New(map[string]string{"key": "a"})
	resultB := corejson.New(map[string]string{"key": "b"})

	// Act
	isEqual := resultA.IsEqual(resultB)

	// Assert
	convey.Convey("IsEqual - different content should return false", t, func() {
		convey.So(isEqual, should.BeFalse)
	})
}

// Test_Result_IsEqualPtr_BothNil — IsEqualPtr returns true when both nil
func Test_Result_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var a *corejson.Result
	var b *corejson.Result

	// Act
	isEqual := a.IsEqualPtr(b)

	// Assert
	convey.Convey("IsEqualPtr - both nil should return true", t, func() {
		convey.So(isEqual, should.BeTrue)
	})
}

// Test_Result_IsEqualPtr_OneNil — IsEqualPtr returns false when one is nil
func Test_Result_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	result := corejson.NewPtr(map[string]string{"k": "v"})
	var nilResult *corejson.Result

	// Act
	isEqual := result.IsEqualPtr(nilResult)

	// Assert
	convey.Convey("IsEqualPtr - one nil should return false", t, func() {
		convey.So(isEqual, should.BeFalse)
	})
}

// Test_Result_IsEqualPtr_SamePointer — IsEqualPtr returns true for same pointer
func Test_Result_IsEqualPtr_SamePointer(t *testing.T) {
	// Arrange
	result := corejson.NewPtr(map[string]string{"k": "v"})

	// Act
	isEqual := result.IsEqualPtr(result)

	// Assert
	convey.Convey("IsEqualPtr - same pointer should return true", t, func() {
		convey.So(isEqual, should.BeTrue)
	})
}
