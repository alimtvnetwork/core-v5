package corejsontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/corejson"
)

// Test_New_ValidStruct — New with a valid struct produces bytes and no error
func Test_New_ValidStruct(t *testing.T) {
	// Arrange
	input := struct {
		Name string
		Age  int
	}{Name: "Alice", Age: 30}

	// Act
	result := corejson.New(input)

	// Assert
	convey.Convey("New - valid struct should produce bytes and no error", t, func() {
		convey.So(result.HasError(), should.BeFalse)
		convey.So(result.IsEmpty(), should.BeFalse)
		convey.So(len(result.Bytes), should.BeGreaterThan, 0)
		convey.So(result.TypeName, should.NotBeEmpty)
	})
}

// Test_New_NilInput — New with nil input produces valid JSON "null"
func Test_New_NilInput(t *testing.T) {
	// Arrange / Act
	result := corejson.New(nil)

	// Assert
	convey.Convey("New - nil input should produce JSON null bytes without error", t, func() {
		convey.So(result.HasError(), should.BeFalse)
		convey.So(string(result.Bytes), should.Equal, "null")
	})
}

// Test_New_UnmarshalableType — New with a channel produces an error
func Test_New_UnmarshalableType(t *testing.T) {
	// Arrange
	ch := make(chan int)

	// Act
	result := corejson.New(ch)

	// Assert
	convey.Convey("New - channel input should produce marshalling error", t, func() {
		convey.So(result.HasError(), should.BeTrue)
		convey.So(result.Error.Error(), should.ContainSubstring, "marshal")
	})
}

// Test_NewPtr_ValidStruct — NewPtr with a valid struct produces pointer result
func Test_NewPtr_ValidStruct(t *testing.T) {
	// Arrange
	input := struct {
		Name string
		Age  int
	}{Name: "Bob", Age: 25}

	// Act
	result := corejson.NewPtr(input)

	// Assert
	convey.Convey("NewPtr - valid struct should produce non-nil result with bytes", t, func() {
		convey.So(result, should.NotBeNil)
		convey.So(result.HasError(), should.BeFalse)
		convey.So(result.IsEmpty(), should.BeFalse)
		convey.So(len(result.Bytes), should.BeGreaterThan, 0)
	})
}

// Test_NewPtr_NilInput — NewPtr with nil produces JSON "null"
func Test_NewPtr_NilInput(t *testing.T) {
	// Arrange / Act
	result := corejson.NewPtr(nil)

	// Assert
	convey.Convey("NewPtr - nil input should produce JSON null bytes without error", t, func() {
		convey.So(result, should.NotBeNil)
		convey.So(result.HasError(), should.BeFalse)
		convey.So(string(result.Bytes), should.Equal, "null")
	})
}

// Test_NewPtr_UnmarshalableType — NewPtr with a channel produces an error
func Test_NewPtr_UnmarshalableType(t *testing.T) {
	// Arrange
	ch := make(chan string)

	// Act
	result := corejson.NewPtr(ch)

	// Assert
	convey.Convey("NewPtr - channel input should produce marshalling error", t, func() {
		convey.So(result, should.NotBeNil)
		convey.So(result.HasError(), should.BeTrue)
		convey.So(result.Error.Error(), should.ContainSubstring, "marshal")
	})
}
