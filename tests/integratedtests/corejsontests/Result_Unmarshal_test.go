package corejsontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/corejson"
)

// Test_Result_Unmarshal_Valid — Unmarshal valid JSON into a struct
func Test_Result_Unmarshal_Valid(t *testing.T) {
	// Arrange
	type Example struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	input := Example{Name: "Alice", Age: 30}
	jsonResult := corejson.NewPtr(input)
	target := &Example{}

	// Act
	err := jsonResult.Unmarshal(target)

	// Assert
	convey.Convey("Unmarshal - valid JSON should deserialize correctly", t, func() {
		convey.So(err, should.BeNil)
		convey.So(target.Name, should.Equal, "Alice")
		convey.So(target.Age, should.Equal, 30)
	})
}

// Test_Result_Unmarshal_NilReceiver — Unmarshal on nil Result returns error
func Test_Result_Unmarshal_NilReceiver(t *testing.T) {
	// Arrange
	var nilResult *corejson.Result

	type Example struct {
		Name string
	}

	target := &Example{}

	// Act
	err := nilResult.Unmarshal(target)

	// Assert
	convey.Convey("Unmarshal - nil receiver should return error", t, func() {
		convey.So(err, should.NotBeNil)
		convey.So(err.Error(), should.ContainSubstring, "null")
	})
}

// Test_Result_Unmarshal_InvalidBytes — Unmarshal with corrupted bytes returns error
func Test_Result_Unmarshal_InvalidBytes(t *testing.T) {
	// Arrange
	result := corejson.NewResult.UsingBytesTypePtr(
		[]byte(`{invalid-json`),
		"TestType",
	)

	type Example struct {
		Name string
	}

	target := &Example{}

	// Act
	err := result.Unmarshal(target)

	// Assert
	convey.Convey("Unmarshal - invalid bytes should return unmarshal error", t, func() {
		convey.So(err, should.NotBeNil)
		convey.So(err.Error(), should.ContainSubstring, "unmarshal")
	})
}

// Test_Result_Unmarshal_ExistingError — Unmarshal on Result with pre-existing error returns error
func Test_Result_Unmarshal_ExistingError(t *testing.T) {
	// Arrange
	ch := make(chan int)
	result := corejson.NewPtr(ch) // will have marshalling error

	type Example struct {
		Name string
	}

	target := &Example{}

	// Act
	err := result.Unmarshal(target)

	// Assert
	convey.Convey("Unmarshal - result with existing error should propagate error", t, func() {
		convey.So(err, should.NotBeNil)
		convey.So(err.Error(), should.ContainSubstring, "unmarshal")
	})
}
