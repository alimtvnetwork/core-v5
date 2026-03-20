package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretaskinfo"
)

// Test_Cov3_JsonString_NilInfo tests JsonString on a nil Info (IsNull branch).
func Test_Cov3_JsonString_NilInfo(t *testing.T) {
	// Arrange
	var info *coretaskinfo.Info

	// Act
	actual := info.JsonString()

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"nil Info.JsonString() should return empty string",
		actual, "",
	)
}

// Test_Cov3_MapWithPayloadAsAny_SerializeError tests the HasError branch in MapWithPayloadAsAny.
func Test_Cov3_MapWithPayloadAsAny_SerializeError(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}

	// Act — pass a channel which cannot be JSON-marshalled
	ch := make(chan int)
	result := info.MapWithPayloadAsAny(ch)

	// Assert — should have a serializing error field
	_, hasPayloadsErr := result["Payloads.SerializingErr"]
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"MapWithPayloadAsAny with unmarshal-able payload should have error",
		hasPayloadsErr, true,
	)
}
