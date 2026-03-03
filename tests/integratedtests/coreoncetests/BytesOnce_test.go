package coreoncetests

import (
	"encoding/json"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

// =============================================================================
// BytesOnce — Caching Behavior
// =============================================================================

func Test_BytesOnce_Value_CachesResult(t *testing.T) {
	callCount := 0
	once := coreonce.NewBytesOncePtr(func() []byte {
		callCount++
		return []byte("hello")
	})

	convey.Convey("BytesOnce.Value caches — initializer runs exactly once", t, func() {
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		convey.So(string(r1), should.Equal, "hello")
		convey.So(string(r2), should.Equal, "hello")
		convey.So(string(r3), should.Equal, "hello")
		convey.So(callCount, should.Equal, 1)
	})
}

func Test_BytesOnce_Execute_SameAsValue(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte("data")
	})

	convey.Convey("BytesOnce.Execute returns same as Value", t, func() {
		convey.So(string(once.Execute()), should.Equal, string(once.Value()))
	})
}

func Test_BytesOnce_Value_NilBytes(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return nil
	})

	convey.Convey("BytesOnce.Value caches nil correctly", t, func() {
		convey.So(once.Value(), should.BeNil)
		convey.So(once.Value(), should.BeNil)
	})
}

func Test_BytesOnce_Value_EmptyBytes(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte{}
	})

	convey.Convey("BytesOnce.Value caches empty bytes", t, func() {
		convey.So(once.Value(), should.BeEmpty)
	})
}

// =============================================================================
// BytesOnce — String
// =============================================================================

func Test_BytesOnce_String(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte("test-string")
	})

	convey.Convey("BytesOnce.String returns string conversion", t, func() {
		convey.So(once.String(), should.Equal, "test-string")
	})
}

func Test_BytesOnce_String_Empty(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte{}
	})

	convey.Convey("BytesOnce.String returns empty for empty bytes", t, func() {
		convey.So(once.String(), should.BeEmpty)
	})
}

// =============================================================================
// BytesOnce — Length & IsEmpty
// =============================================================================

func Test_BytesOnce_Length(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte("abc")
	})

	convey.Convey("BytesOnce.Length returns correct length", t, func() {
		convey.So(once.Length(), should.Equal, 3)
	})
}

func Test_BytesOnce_Length_NilInitializer(t *testing.T) {
	once := &coreonce.BytesOnce{}

	convey.Convey("BytesOnce.Length returns 0 for nil initializer", t, func() {
		convey.So(once.Length(), should.Equal, 0)
	})
}

func Test_BytesOnce_IsEmpty_True(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte{}
	})

	convey.Convey("BytesOnce.IsEmpty returns true for empty bytes", t, func() {
		convey.So(once.IsEmpty(), should.BeTrue)
	})
}

func Test_BytesOnce_IsEmpty_False(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte("x")
	})

	convey.Convey("BytesOnce.IsEmpty returns false for non-empty", t, func() {
		convey.So(once.IsEmpty(), should.BeFalse)
	})
}

func Test_BytesOnce_IsEmpty_NilInitializer(t *testing.T) {
	once := &coreonce.BytesOnce{}

	convey.Convey("BytesOnce.IsEmpty returns true for nil initializer", t, func() {
		convey.So(once.IsEmpty(), should.BeTrue)
	})
}

// =============================================================================
// BytesOnce — JSON
// =============================================================================

func Test_BytesOnce_MarshalJSON(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte("hello")
	})

	convey.Convey("BytesOnce.MarshalJSON marshals bytes", t, func() {
		data, err := once.MarshalJSON()
		convey.So(err, should.BeNil)
		convey.So(len(data), should.BeGreaterThan, 0)
	})
}

func Test_BytesOnce_UnmarshalJSON(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte("original")
	})

	convey.Convey("BytesOnce.UnmarshalJSON overrides value", t, func() {
		input, _ := json.Marshal([]byte("replaced"))
		err := once.UnmarshalJSON(input)
		convey.So(err, should.BeNil)
		convey.So(string(once.Value()), should.Equal, "replaced")
	})
}

func Test_BytesOnce_Serialize(t *testing.T) {
	once := coreonce.NewBytesOncePtr(func() []byte {
		return []byte("serialize-me")
	})

	convey.Convey("BytesOnce.Serialize returns JSON bytes", t, func() {
		data, err := once.Serialize()
		convey.So(err, should.BeNil)
		convey.So(len(data), should.BeGreaterThan, 0)
	})
}

// =============================================================================
// BytesOnce — Constructor variants
// =============================================================================

func Test_BytesOnce_NewBytesOnce_Value(t *testing.T) {
	once := coreonce.NewBytesOnce(func() []byte {
		return []byte("val")
	})

	convey.Convey("NewBytesOnce (value) works correctly", t, func() {
		convey.So(string(once.Value()), should.Equal, "val")
	})
}
