package coreoncetests

import (
	"encoding/json"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/smarty/assertions/should"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

// =============================================================================
// BoolOnce — Caching Behavior
// =============================================================================

func Test_BoolOnce_Value_CachesResult(t *testing.T) {
	callCount := 0
	once := coreonce.NewBoolOncePtr(func() bool {
		callCount++
		return true
	})

	convey.Convey("BoolOnce.Value caches — initializer runs exactly once", t, func() {
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		convey.So(r1, should.BeTrue)
		convey.So(r2, should.BeTrue)
		convey.So(r3, should.BeTrue)
		convey.So(callCount, should.Equal, 1)
	})
}

func Test_BoolOnce_Value_False(t *testing.T) {
	once := coreonce.NewBoolOncePtr(func() bool { return false })

	convey.Convey("BoolOnce.Value caches false correctly", t, func() {
		convey.So(once.Value(), should.BeFalse)
		convey.So(once.Value(), should.BeFalse)
	})
}

func Test_BoolOnce_Execute_SameAsValue(t *testing.T) {
	once := coreonce.NewBoolOncePtr(func() bool { return true })

	convey.Convey("BoolOnce.Execute returns same as Value", t, func() {
		convey.So(once.Execute(), should.Equal, once.Value())
	})
}

// =============================================================================
// BoolOnce — String & JSON
// =============================================================================

func Test_BoolOnce_String_True(t *testing.T) {
	once := coreonce.NewBoolOncePtr(func() bool { return true })

	convey.Convey("BoolOnce.String returns 'true'", t, func() {
		convey.So(once.String(), should.Equal, "true")
	})
}

func Test_BoolOnce_String_False(t *testing.T) {
	once := coreonce.NewBoolOncePtr(func() bool { return false })

	convey.Convey("BoolOnce.String returns 'false'", t, func() {
		convey.So(once.String(), should.Equal, "false")
	})
}

func Test_BoolOnce_MarshalJSON(t *testing.T) {
	once := coreonce.NewBoolOncePtr(func() bool { return true })

	convey.Convey("BoolOnce.MarshalJSON marshals true", t, func() {
		data, err := once.MarshalJSON()
		convey.So(err, should.BeNil)
		convey.So(string(data), should.Equal, "true")
	})
}

func Test_BoolOnce_UnmarshalJSON(t *testing.T) {
	once := coreonce.NewBoolOncePtr(func() bool { return true })

	convey.Convey("BoolOnce.UnmarshalJSON overrides with false", t, func() {
		err := once.UnmarshalJSON([]byte("false"))
		convey.So(err, should.BeNil)
		convey.So(once.Value(), should.BeFalse)
	})
}

func Test_BoolOnce_Serialize(t *testing.T) {
	once := coreonce.NewBoolOncePtr(func() bool { return false })

	convey.Convey("BoolOnce.Serialize returns JSON bytes", t, func() {
		data, err := once.Serialize()
		convey.So(err, should.BeNil)

		var result bool
		_ = json.Unmarshal(data, &result)
		convey.So(result, should.BeFalse)
	})
}
