package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretaskinfo"
)

// Test_Cov3_JsonString_NilInfo documents that the IsNull() branch in
// Info.JsonString() (line 25-27) is unreachable dead code.
// JsonString has a value receiver `func (it Info) JsonString()`, so `it`
// is always a valid value — `it.IsNull()` (which checks `it == nil`)
// will never be true because Go auto-addresses `&it` which is non-nil.
//
// We still exercise the zero-value path to get as close as possible.
func Test_Cov3_JsonString_ZeroInfo(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{}

	// Act
	actual := info.JsonString()

	// Assert — zero-value Info is not nil, IsNull returns false
	// so it will try to serialize. We just verify it doesn't panic.
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
	coretests.GetAssert.ShouldBeEqual(
		t, 0,
		"MapWithPayloadAsAny with unmarshal-able payload should have error",
		hasPayloadsErr, true,
	)
}
