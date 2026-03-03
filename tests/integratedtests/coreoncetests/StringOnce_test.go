package coreoncetests

import (
	"encoding/json"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/smarty/assertions/should"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

// =============================================================================
// StringOnce — Caching Behavior
// =============================================================================

func Test_StringOnce_Value_CachesResult(t *testing.T) {
	callCount := 0
	once := coreonce.NewStringOncePtr(func() string {
		callCount++
		return "hello"
	})

	convey.Convey("StringOnce.Value caches — initializer runs exactly once", t, func() {
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		convey.So(r1, should.Equal, "hello")
		convey.So(r2, should.Equal, "hello")
		convey.So(r3, should.Equal, "hello")
		convey.So(callCount, should.Equal, 1)
	})
}

func Test_StringOnce_Execute_SameAsValue(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "exec" })

	convey.Convey("StringOnce.Execute returns same as Value", t, func() {
		convey.So(once.Execute(), should.Equal, once.Value())
	})
}

func Test_StringOnce_ValuePtr_ReturnsPointer(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "ptr" })

	convey.Convey("StringOnce.ValuePtr returns pointer to value", t, func() {
		ptr := once.ValuePtr()
		convey.So(ptr, should.NotBeNil)
		convey.So(*ptr, should.Equal, "ptr")
	})
}

// =============================================================================
// StringOnce — String Operations
// =============================================================================

func Test_StringOnce_IsEmpty_True(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "" })

	convey.Convey("StringOnce.IsEmpty returns true for empty string", t, func() {
		convey.So(once.IsEmpty(), should.BeTrue)
	})
}

func Test_StringOnce_IsEmpty_False(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "data" })

	convey.Convey("StringOnce.IsEmpty returns false for non-empty string", t, func() {
		convey.So(once.IsEmpty(), should.BeFalse)
	})
}

func Test_StringOnce_IsEmptyOrWhitespace_True(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "   " })

	convey.Convey("StringOnce.IsEmptyOrWhitespace returns true for whitespace", t, func() {
		convey.So(once.IsEmptyOrWhitespace(), should.BeTrue)
	})
}

func Test_StringOnce_IsEqual(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "abc" })

	convey.Convey("StringOnce.IsEqual matches exact string", t, func() {
		convey.So(once.IsEqual("abc"), should.BeTrue)
		convey.So(once.IsEqual("xyz"), should.BeFalse)
	})
}

func Test_StringOnce_IsContains(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "hello world" })

	convey.Convey("StringOnce.IsContains checks substring", t, func() {
		convey.So(once.IsContains("world"), should.BeTrue)
		convey.So(once.IsContains("xyz"), should.BeFalse)
	})
}

func Test_StringOnce_HasPrefix(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "prefix-data" })

	convey.Convey("StringOnce.HasPrefix and IsStartsWith", t, func() {
		convey.So(once.HasPrefix("prefix"), should.BeTrue)
		convey.So(once.IsStartsWith("prefix"), should.BeTrue)
		convey.So(once.HasPrefix("data"), should.BeFalse)
	})
}

func Test_StringOnce_HasSuffix(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "data-suffix" })

	convey.Convey("StringOnce.HasSuffix and IsEndsWith", t, func() {
		convey.So(once.HasSuffix("suffix"), should.BeTrue)
		convey.So(once.IsEndsWith("suffix"), should.BeTrue)
		convey.So(once.HasSuffix("data"), should.BeFalse)
	})
}

func Test_StringOnce_SplitBy(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "a,b,c" })

	convey.Convey("StringOnce.SplitBy splits by delimiter", t, func() {
		parts := once.SplitBy(",")
		convey.So(len(parts), should.Equal, 3)
		convey.So(parts[0], should.Equal, "a")
		convey.So(parts[2], should.Equal, "c")
	})
}

func Test_StringOnce_SplitLeftRight(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "key=value" })

	convey.Convey("StringOnce.SplitLeftRight splits into two parts", t, func() {
		left, right := once.SplitLeftRight("=")
		convey.So(left, should.Equal, "key")
		convey.So(right, should.Equal, "value")
	})
}

func Test_StringOnce_SplitLeftRight_NoSplitter(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "nosplit" })

	convey.Convey("StringOnce.SplitLeftRight with no splitter returns full left, empty right", t, func() {
		left, right := once.SplitLeftRight("=")
		convey.So(left, should.Equal, "nosplit")
		convey.So(right, should.Equal, "")
	})
}

func Test_StringOnce_SplitLeftRightTrim(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return " key = value " })

	convey.Convey("StringOnce.SplitLeftRightTrim trims whitespace", t, func() {
		left, right := once.SplitLeftRightTrim("=")
		convey.So(left, should.Equal, "key")
		convey.So(right, should.Equal, "value")
	})
}

func Test_StringOnce_Bytes(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "bytes" })

	convey.Convey("StringOnce.Bytes returns byte slice", t, func() {
		convey.So(string(once.Bytes()), should.Equal, "bytes")
	})
}

func Test_StringOnce_Error(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "err msg" })

	convey.Convey("StringOnce.Error returns error with value as message", t, func() {
		convey.So(once.Error().Error(), should.Equal, "err msg")
	})
}

func Test_StringOnce_String(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "str" })

	convey.Convey("StringOnce.String returns value", t, func() {
		convey.So(once.String(), should.Equal, "str")
	})
}

// =============================================================================
// StringOnce — JSON
// =============================================================================

func Test_StringOnce_MarshalJSON(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "json" })

	convey.Convey("StringOnce.MarshalJSON marshals cached value", t, func() {
		data, err := once.MarshalJSON()
		convey.So(err, should.BeNil)
		convey.So(string(data), should.Equal, `"json"`)
	})
}

func Test_StringOnce_UnmarshalJSON(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "initial" })

	convey.Convey("StringOnce.UnmarshalJSON overrides value", t, func() {
		err := once.UnmarshalJSON([]byte(`"overridden"`))
		convey.So(err, should.BeNil)
		convey.So(once.Value(), should.Equal, "overridden")
	})
}

func Test_StringOnce_Serialize(t *testing.T) {
	once := coreonce.NewStringOncePtr(func() string { return "serial" })

	convey.Convey("StringOnce.Serialize returns JSON bytes", t, func() {
		data, err := once.Serialize()
		convey.So(err, should.BeNil)

		var result string
		_ = json.Unmarshal(data, &result)
		convey.So(result, should.Equal, "serial")
	})
}
