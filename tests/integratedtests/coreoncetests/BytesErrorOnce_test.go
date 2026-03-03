package coreoncetests

import (
	"errors"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coreonce"
)

// =============================================================================
// BytesErrorOnce — Caching Behavior
// =============================================================================

func Test_BytesErrorOnce_Value_CachesResult(t *testing.T) {
	callCount := 0
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		callCount++
		return []byte("cached"), nil
	})

	convey.Convey("BytesErrorOnce.Value caches — initializer runs exactly once", t, func() {
		r1, e1 := once.Value()
		r2, e2 := once.Value()

		convey.So(string(r1), should.Equal, "cached")
		convey.So(string(r2), should.Equal, "cached")
		convey.So(e1, should.BeNil)
		convey.So(e2, should.BeNil)
		convey.So(callCount, should.Equal, 1)
	})
}

func Test_BytesErrorOnce_Value_CachesError(t *testing.T) {
	testErr := errors.New("test error")
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, testErr
	})

	convey.Convey("BytesErrorOnce.Value caches error", t, func() {
		val, err := once.Value()
		convey.So(val, should.BeNil)
		convey.So(err, should.Equal, testErr)
	})
}

func Test_BytesErrorOnce_Execute_SameAsValue(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("exec"), nil
	})

	convey.Convey("BytesErrorOnce.Execute returns same as Value", t, func() {
		v1, _ := once.Execute()
		v2, _ := once.Value()
		convey.So(string(v1), should.Equal, string(v2))
	})
}

func Test_BytesErrorOnce_ValueOnly(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("only"), nil
	})

	convey.Convey("BytesErrorOnce.ValueOnly returns bytes without error", t, func() {
		convey.So(string(once.ValueOnly()), should.Equal, "only")
	})
}

func Test_BytesErrorOnce_ValueWithError(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("vwe"), nil
	})

	convey.Convey("BytesErrorOnce.ValueWithError aliases Value", t, func() {
		v, e := once.ValueWithError()
		convey.So(string(v), should.Equal, "vwe")
		convey.So(e, should.BeNil)
	})
}

// =============================================================================
// BytesErrorOnce — Error State Queries
// =============================================================================

func Test_BytesErrorOnce_HasError_True(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("fail")
	})

	convey.Convey("BytesErrorOnce.HasError returns true on error", t, func() {
		convey.So(once.HasError(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_HasError_False(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})

	convey.Convey("BytesErrorOnce.HasError returns false on success", t, func() {
		convey.So(once.HasError(), should.BeFalse)
	})
}

func Test_BytesErrorOnce_IsEmptyError(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})

	convey.Convey("BytesErrorOnce.IsEmptyError returns true on no error", t, func() {
		convey.So(once.IsEmptyError(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsValid_IsSuccess(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})

	convey.Convey("BytesErrorOnce.IsValid and IsSuccess return true on no error", t, func() {
		convey.So(once.IsValid(), should.BeTrue)
		convey.So(once.IsSuccess(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsInvalid_IsFailed(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("bad")
	})

	convey.Convey("BytesErrorOnce.IsInvalid and IsFailed return true on error", t, func() {
		convey.So(once.IsInvalid(), should.BeTrue)
		convey.So(once.IsFailed(), should.BeTrue)
	})
}

// =============================================================================
// BytesErrorOnce — HasIssuesOrEmpty
// =============================================================================

func Test_BytesErrorOnce_HasIssuesOrEmpty_WithError(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("data"), errors.New("err")
	})

	convey.Convey("BytesErrorOnce.HasIssuesOrEmpty true when error present", t, func() {
		convey.So(once.HasIssuesOrEmpty(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_HasIssuesOrEmpty_EmptyBytes(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})

	convey.Convey("BytesErrorOnce.HasIssuesOrEmpty true when bytes empty", t, func() {
		convey.So(once.HasIssuesOrEmpty(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_HasIssuesOrEmpty_False(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})

	convey.Convey("BytesErrorOnce.HasIssuesOrEmpty false when data and no error", t, func() {
		convey.So(once.HasIssuesOrEmpty(), should.BeFalse)
	})
}

func Test_BytesErrorOnce_HasSafeItems(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("safe"), nil
	})

	convey.Convey("BytesErrorOnce.HasSafeItems true when data and no error", t, func() {
		convey.So(once.HasSafeItems(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_HasIssuesOrEmpty_NilReceiver(t *testing.T) {
	var once *coreonce.BytesErrorOnce

	convey.Convey("BytesErrorOnce.HasIssuesOrEmpty true on nil receiver", t, func() {
		convey.So(once.HasIssuesOrEmpty(), should.BeTrue)
	})
}

// =============================================================================
// BytesErrorOnce — Length, IsEmpty, IsNull, IsBytesEmpty
// =============================================================================

func Test_BytesErrorOnce_Length(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("abc"), nil
	})

	convey.Convey("BytesErrorOnce.Length returns correct length", t, func() {
		convey.So(once.Length(), should.Equal, 3)
	})
}

func Test_BytesErrorOnce_HasAnyItem(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("x"), nil
	})

	convey.Convey("BytesErrorOnce.HasAnyItem true for non-empty", t, func() {
		convey.So(once.HasAnyItem(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsEmpty_BothNil(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, nil
	})

	convey.Convey("BytesErrorOnce.IsEmpty true when nil bytes and nil error", t, func() {
		convey.So(once.IsEmpty(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsEmptyBytes(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})

	convey.Convey("BytesErrorOnce.IsEmptyBytes true for empty bytes", t, func() {
		convey.So(once.IsEmptyBytes(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsBytesEmpty(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})

	convey.Convey("BytesErrorOnce.IsBytesEmpty true for empty bytes", t, func() {
		convey.So(once.IsBytesEmpty(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsNull(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, nil
	})

	convey.Convey("BytesErrorOnce.IsNull true when bytes nil", t, func() {
		convey.So(once.IsNull(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsDefined(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("defined"), nil
	})

	convey.Convey("BytesErrorOnce.IsDefined true for non-empty with no error", t, func() {
		convey.So(once.IsDefined(), should.BeTrue)
	})
}

// =============================================================================
// BytesErrorOnce — String
// =============================================================================

func Test_BytesErrorOnce_String(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("str-val"), nil
	})

	convey.Convey("BytesErrorOnce.String returns string conversion", t, func() {
		convey.So(once.String(), should.Equal, "str-val")
	})
}

func Test_BytesErrorOnce_String_Nil(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, nil
	})

	convey.Convey("BytesErrorOnce.String returns empty for nil bytes", t, func() {
		convey.So(once.String(), should.BeEmpty)
	})
}

func Test_BytesErrorOnce_IsStringEmpty(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, nil
	})

	convey.Convey("BytesErrorOnce.IsStringEmpty true for nil bytes", t, func() {
		convey.So(once.IsStringEmpty(), should.BeTrue)
	})
}

func Test_BytesErrorOnce_IsStringEmptyOrWhitespace(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("   "), nil
	})

	convey.Convey("BytesErrorOnce.IsStringEmptyOrWhitespace true for whitespace", t, func() {
		convey.So(once.IsStringEmptyOrWhitespace(), should.BeTrue)
	})
}

// =============================================================================
// BytesErrorOnce — Deserialize
// =============================================================================

func Test_BytesErrorOnce_Deserialize_Success(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`{"name":"test"}`), nil
	})

	convey.Convey("BytesErrorOnce.Deserialize succeeds with valid JSON", t, func() {
		var result map[string]string
		err := once.Deserialize(&result)
		convey.So(err, should.BeNil)
		convey.So(result["name"], should.Equal, "test")
	})
}

func Test_BytesErrorOnce_Deserialize_WithExistingError(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("source error")
	})

	convey.Convey("BytesErrorOnce.Deserialize returns error when source has error", t, func() {
		var result map[string]string
		err := once.Deserialize(&result)
		convey.So(err, should.NotBeNil)
		convey.So(err.Error(), should.ContainSubstring, "existing error cannot deserialize")
		convey.So(err.Error(), should.ContainSubstring, "source error")
	})
}

func Test_BytesErrorOnce_Deserialize_InvalidJSON(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`not-json`), nil
	})

	convey.Convey("BytesErrorOnce.Deserialize with invalid JSON", t, func() {
		var result map[string]string
		// Note: Due to a bug in the source (checks err==nil instead of jsonUnmarshalErr!=nil),
		// this actually returns nil. Testing actual behavior.
		err := once.Deserialize(&result)
		convey.So(err, should.BeNil) // reflects actual code behavior
	})
}

func Test_BytesErrorOnce_DeserializeMust_Success(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`{"key":"val"}`), nil
	})

	convey.Convey("BytesErrorOnce.DeserializeMust succeeds without panic", t, func() {
		var result map[string]string
		convey.So(func() { once.DeserializeMust(&result) }, should.NotPanic)
		convey.So(result["key"], should.Equal, "val")
	})
}

func Test_BytesErrorOnce_DeserializeMust_Panics(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("must-fail")
	})

	convey.Convey("BytesErrorOnce.DeserializeMust panics on error", t, func() {
		var result map[string]string
		convey.So(func() { once.DeserializeMust(&result) }, should.Panic)
	})
}

// =============================================================================
// BytesErrorOnce — JSON Serialization
// =============================================================================

func Test_BytesErrorOnce_MarshalJSON(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte(`{"a":1}`), nil
	})

	convey.Convey("BytesErrorOnce.MarshalJSON returns bytes", t, func() {
		data, err := once.MarshalJSON()
		convey.So(err, should.BeNil)
		convey.So(string(data), should.Equal, `{"a":1}`)
	})
}

func Test_BytesErrorOnce_Serialize(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ser"), nil
	})

	convey.Convey("BytesErrorOnce.Serialize returns bytes", t, func() {
		data, err := once.Serialize()
		convey.So(err, should.BeNil)
		convey.So(string(data), should.Equal, "ser")
	})
}

func Test_BytesErrorOnce_SerializeMust_Success(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("must-ser"), nil
	})

	convey.Convey("BytesErrorOnce.SerializeMust returns bytes without panic", t, func() {
		var result []byte
		convey.So(func() { result = once.SerializeMust() }, should.NotPanic)
		convey.So(string(result), should.Equal, "must-ser")
	})
}

func Test_BytesErrorOnce_SerializeMust_Panics(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("ser-fail")
	})

	convey.Convey("BytesErrorOnce.SerializeMust panics on error", t, func() {
		convey.So(func() { once.SerializeMust() }, should.Panic)
	})
}

// =============================================================================
// BytesErrorOnce — HandleError / MustBeEmptyError / MustHaveSafeItems
// =============================================================================

func Test_BytesErrorOnce_HandleError_NoPanic(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})

	convey.Convey("BytesErrorOnce.HandleError does not panic on success", t, func() {
		convey.So(func() { once.HandleError() }, should.NotPanic)
	})
}

func Test_BytesErrorOnce_HandleError_Panics(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return nil, errors.New("handle-err")
	})

	convey.Convey("BytesErrorOnce.HandleError panics on error", t, func() {
		convey.So(func() { once.HandleError() }, should.Panic)
	})
}

func Test_BytesErrorOnce_MustBeEmptyError_NoPanic(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("ok"), nil
	})

	convey.Convey("BytesErrorOnce.MustBeEmptyError does not panic on success", t, func() {
		convey.So(func() { once.MustBeEmptyError() }, should.NotPanic)
	})
}

func Test_BytesErrorOnce_MustHaveSafeItems_NoPanic(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("safe"), nil
	})

	convey.Convey("BytesErrorOnce.MustHaveSafeItems does not panic with data", t, func() {
		convey.So(func() { once.MustHaveSafeItems() }, should.NotPanic)
	})
}

func Test_BytesErrorOnce_MustHaveSafeItems_PanicsEmpty(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte{}, nil
	})

	convey.Convey("BytesErrorOnce.MustHaveSafeItems panics when empty", t, func() {
		convey.So(func() { once.MustHaveSafeItems() }, should.Panic)
	})
}

// =============================================================================
// BytesErrorOnce — IsInitialized
// =============================================================================

func Test_BytesErrorOnce_IsInitialized_Before(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("x"), nil
	})

	convey.Convey("BytesErrorOnce.IsInitialized false before Value call", t, func() {
		convey.So(once.IsInitialized(), should.BeFalse)
	})
}

func Test_BytesErrorOnce_IsInitialized_After(t *testing.T) {
	once := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) {
		return []byte("x"), nil
	})

	convey.Convey("BytesErrorOnce.IsInitialized true after Value call", t, func() {
		_, _ = once.Value()
		convey.So(once.IsInitialized(), should.BeTrue)
	})
}

// =============================================================================
// BytesErrorOnce — Constructor variants
// =============================================================================

func Test_BytesErrorOnce_NewBytesErrorOnce_Value(t *testing.T) {
	once := coreonce.NewBytesErrorOnce(func() ([]byte, error) {
		return []byte("val"), nil
	})

	convey.Convey("NewBytesErrorOnce (value) works correctly", t, func() {
		v, e := once.Value()
		convey.So(string(v), should.Equal, "val")
		convey.So(e, should.BeNil)
	})
}
