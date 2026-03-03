package coreoncetests

import (
	"encoding/json"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/smarty/assertions/should"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

// =============================================================================
// IntegerOnce — Caching Behavior
// =============================================================================

func Test_IntegerOnce_Value_CachesResult(t *testing.T) {
	callCount := 0
	once := coreonce.NewIntegerOncePtr(func() int {
		callCount++
		return 42
	})

	convey.Convey("IntegerOnce.Value caches — initializer runs exactly once", t, func() {
		r1 := once.Value()
		r2 := once.Value()

		convey.So(r1, should.Equal, 42)
		convey.So(r2, should.Equal, 42)
		convey.So(callCount, should.Equal, 1)
	})
}

func Test_IntegerOnce_Execute_SameAsValue(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 7 })

	convey.Convey("IntegerOnce.Execute returns same as Value", t, func() {
		convey.So(once.Execute(), should.Equal, once.Value())
	})
}

// =============================================================================
// IntegerOnce — Comparisons
// =============================================================================

func Test_IntegerOnce_IsZero(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 0 })

	convey.Convey("IntegerOnce.IsZero/IsEmpty true for 0", t, func() {
		convey.So(once.IsZero(), should.BeTrue)
		convey.So(once.IsEmpty(), should.BeTrue)
	})
}

func Test_IntegerOnce_IsAboveZero(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 5 })

	convey.Convey("IntegerOnce.IsAboveZero/IsPositive true for positive", t, func() {
		convey.So(once.IsAboveZero(), should.BeTrue)
		convey.So(once.IsPositive(), should.BeTrue)
		convey.So(once.IsValidIndex(), should.BeTrue)
	})
}

func Test_IntegerOnce_IsLessThanZero(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return -3 })

	convey.Convey("IntegerOnce.IsLessThanZero/IsNegative true for negative", t, func() {
		convey.So(once.IsLessThanZero(), should.BeTrue)
		convey.So(once.IsNegative(), should.BeTrue)
		convey.So(once.IsInvalidIndex(), should.BeTrue)
	})
}

func Test_IntegerOnce_IsAbove(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 10 })

	convey.Convey("IntegerOnce.IsAbove/IsAboveEqual", t, func() {
		convey.So(once.IsAbove(5), should.BeTrue)
		convey.So(once.IsAbove(10), should.BeFalse)
		convey.So(once.IsAboveEqual(10), should.BeTrue)
	})
}

func Test_IntegerOnce_IsLessThan(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 3 })

	convey.Convey("IntegerOnce.IsLessThan/IsLessThanEqual", t, func() {
		convey.So(once.IsLessThan(5), should.BeTrue)
		convey.So(once.IsLessThan(3), should.BeFalse)
		convey.So(once.IsLessThanEqual(3), should.BeTrue)
	})
}

// =============================================================================
// IntegerOnce — String & JSON
// =============================================================================

func Test_IntegerOnce_String(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 99 })

	convey.Convey("IntegerOnce.String returns decimal representation", t, func() {
		convey.So(once.String(), should.Equal, "99")
	})
}

func Test_IntegerOnce_MarshalJSON(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 42 })

	convey.Convey("IntegerOnce.MarshalJSON marshals integer", t, func() {
		data, err := once.MarshalJSON()
		convey.So(err, should.BeNil)
		convey.So(string(data), should.Equal, "42")
	})
}

func Test_IntegerOnce_UnmarshalJSON(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 0 })

	convey.Convey("IntegerOnce.UnmarshalJSON overrides value", t, func() {
		err := once.UnmarshalJSON([]byte("77"))
		convey.So(err, should.BeNil)
		convey.So(once.Value(), should.Equal, 77)
	})
}

func Test_IntegerOnce_Serialize(t *testing.T) {
	once := coreonce.NewIntegerOncePtr(func() int { return 15 })

	convey.Convey("IntegerOnce.Serialize returns JSON bytes", t, func() {
		data, err := once.Serialize()
		convey.So(err, should.BeNil)

		var result int
		_ = json.Unmarshal(data, &result)
		convey.So(result, should.Equal, 15)
	})
}
