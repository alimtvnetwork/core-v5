package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
)

// Test_Cov3_JsonString_ZeroInfo exercises JsonString on a zero-value Info.
func Test_Cov3_JsonString_ZeroInfo(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{}

	// Act
	actual := info.JsonString()

	// Assert — zero-value Info is not nil, just verify it doesn't panic
	_ = actual
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
	if !hasPayloadsErr {
		t.Fatalf("MapWithPayloadAsAny with unmarshal-able payload should have error key, got keys: %v", result)
	}
}
