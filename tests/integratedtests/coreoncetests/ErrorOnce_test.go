package coreoncetests

import (
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/smarty/assertions/should"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

// =============================================================================
// ErrorOnce — Caching Behavior
// =============================================================================

func Test_ErrorOnce_Value_CachesResult(t *testing.T) {
	callCount := 0
	once := coreonce.NewErrorOncePtr(func() error {
		callCount++
		return errors.New("fail")
	})

	convey.Convey("ErrorOnce.Value caches — initializer runs exactly once", t, func() {
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		convey.So(r1.Error(), should.Equal, "fail")
		convey.So(r2.Error(), should.Equal, "fail")
		convey.So(r3.Error(), should.Equal, "fail")
		convey.So(callCount, should.Equal, 1)
	})
}

func Test_ErrorOnce_Value_NilError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.Value caches nil error", t, func() {
		convey.So(once.Value(), should.BeNil)
		convey.So(once.Value(), should.BeNil)
	})
}

func Test_ErrorOnce_Execute_SameAsValue(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("x") })

	convey.Convey("ErrorOnce.Execute returns same as Value", t, func() {
		convey.So(once.Execute().Error(), should.Equal, once.Value().Error())
	})
}

// =============================================================================
// ErrorOnce — State Queries
// =============================================================================

func Test_ErrorOnce_HasError_True(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("err") })

	convey.Convey("ErrorOnce.HasError returns true when error exists", t, func() {
		convey.So(once.HasError(), should.BeTrue)
	})
}

func Test_ErrorOnce_HasError_False(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.HasError returns false when nil", t, func() {
		convey.So(once.HasError(), should.BeFalse)
	})
}

func Test_ErrorOnce_IsNullOrEmpty_Nil(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.IsNullOrEmpty returns true for nil", t, func() {
		convey.So(once.IsNullOrEmpty(), should.BeTrue)
	})
}

func Test_ErrorOnce_IsNullOrEmpty_EmptyString(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("") })

	convey.Convey("ErrorOnce.IsNullOrEmpty returns true for empty error message", t, func() {
		convey.So(once.IsNullOrEmpty(), should.BeTrue)
	})
}

func Test_ErrorOnce_IsNullOrEmpty_WithMessage(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("msg") })

	convey.Convey("ErrorOnce.IsNullOrEmpty returns false when error has message", t, func() {
		convey.So(once.IsNullOrEmpty(), should.BeFalse)
	})
}

// =============================================================================
// ErrorOnce — Semantic Aliases
// =============================================================================

func Test_ErrorOnce_IsValid_NoError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.IsValid/IsSuccess return true when no error", t, func() {
		convey.So(once.IsValid(), should.BeTrue)
		convey.So(once.IsSuccess(), should.BeTrue)
		convey.So(once.IsEmpty(), should.BeTrue)
	})
}

func Test_ErrorOnce_IsInvalid_HasError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("x") })

	convey.Convey("ErrorOnce.IsInvalid/IsFailed return true when error exists", t, func() {
		convey.So(once.IsInvalid(), should.BeTrue)
		convey.So(once.IsFailed(), should.BeTrue)
		convey.So(once.HasAnyItem(), should.BeTrue)
		convey.So(once.IsDefined(), should.BeTrue)
	})
}

// =============================================================================
// ErrorOnce — Message
// =============================================================================

func Test_ErrorOnce_Message_WithError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("detail") })

	convey.Convey("ErrorOnce.Message returns error string", t, func() {
		convey.So(once.Message(), should.Equal, "detail")
	})
}

func Test_ErrorOnce_Message_NilError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.Message returns empty string for nil error", t, func() {
		convey.So(once.Message(), should.BeEmpty)
	})
}

func Test_ErrorOnce_IsMessageEqual(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("match") })

	convey.Convey("ErrorOnce.IsMessageEqual matches message", t, func() {
		convey.So(once.IsMessageEqual("match"), should.BeTrue)
		convey.So(once.IsMessageEqual("other"), should.BeFalse)
	})
}

func Test_ErrorOnce_IsMessageEqual_NilError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.IsMessageEqual returns false for nil error", t, func() {
		convey.So(once.IsMessageEqual("anything"), should.BeFalse)
	})
}

// =============================================================================
// ErrorOnce — ConcatNew
// =============================================================================

func Test_ErrorOnce_ConcatNewString_WithError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("base") })

	convey.Convey("ErrorOnce.ConcatNewString appends messages", t, func() {
		result := once.ConcatNewString("extra")
		convey.So(result, should.ContainSubstring, "base")
		convey.So(result, should.ContainSubstring, "extra")
	})
}

func Test_ErrorOnce_ConcatNewString_NilError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.ConcatNewString returns only additional messages when nil", t, func() {
		result := once.ConcatNewString("only")
		convey.So(result, should.Equal, "only")
	})
}

func Test_ErrorOnce_ConcatNew_ReturnsError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("err") })

	convey.Convey("ErrorOnce.ConcatNew returns error type", t, func() {
		result := once.ConcatNew("more")
		convey.So(result, should.NotBeNil)
		convey.So(result.Error(), should.ContainSubstring, "err")
		convey.So(result.Error(), should.ContainSubstring, "more")
	})
}

// =============================================================================
// ErrorOnce — JSON
// =============================================================================

func Test_ErrorOnce_MarshalJSON_WithError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return errors.New("marshal") })

	convey.Convey("ErrorOnce.MarshalJSON marshals error message", t, func() {
		data, err := once.MarshalJSON()
		convey.So(err, should.BeNil)
		convey.So(string(data), should.Equal, `"marshal"`)
	})
}

func Test_ErrorOnce_MarshalJSON_NilError(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.MarshalJSON marshals empty string for nil", t, func() {
		data, err := once.MarshalJSON()
		convey.So(err, should.BeNil)
		convey.So(string(data), should.Equal, `""`)
	})
}

func Test_ErrorOnce_UnmarshalJSON(t *testing.T) {
	once := coreonce.NewErrorOncePtr(func() error { return nil })

	convey.Convey("ErrorOnce.UnmarshalJSON sets error from JSON string", t, func() {
		err := once.UnmarshalJSON([]byte(`"unmarshaled"`))
		convey.So(err, should.BeNil)
		convey.So(once.Value().Error(), should.Equal, "unmarshaled")
	})
}
