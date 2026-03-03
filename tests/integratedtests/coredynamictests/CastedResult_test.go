package coredynamictests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coredynamic"
)

// =============================================================================
// CastedResult — IsInvalid
// =============================================================================

func Test_CastedResult_IsInvalid_NilReceiver(t *testing.T) {
	convey.Convey("CastedResult.IsInvalid true on nil receiver", t, func() {
		var cr *coredynamic.CastedResult
		convey.So(cr.IsInvalid(), should.BeTrue)
	})
}

func Test_CastedResult_IsInvalid_False(t *testing.T) {
	convey.Convey("CastedResult.IsInvalid false when IsValid=true", t, func() {
		cr := &coredynamic.CastedResult{IsValid: true}
		convey.So(cr.IsInvalid(), should.BeFalse)
	})
}

func Test_CastedResult_IsInvalid_True(t *testing.T) {
	convey.Convey("CastedResult.IsInvalid true when IsValid=false", t, func() {
		cr := &coredynamic.CastedResult{IsValid: false}
		convey.So(cr.IsInvalid(), should.BeTrue)
	})
}

// =============================================================================
// CastedResult — IsNotNull
// =============================================================================

func Test_CastedResult_IsNotNull_NilReceiver(t *testing.T) {
	convey.Convey("CastedResult.IsNotNull false on nil receiver", t, func() {
		var cr *coredynamic.CastedResult
		convey.So(cr.IsNotNull(), should.BeFalse)
	})
}

func Test_CastedResult_IsNotNull_True(t *testing.T) {
	convey.Convey("CastedResult.IsNotNull true when IsNull=false", t, func() {
		cr := &coredynamic.CastedResult{IsNull: false}
		convey.So(cr.IsNotNull(), should.BeTrue)
	})
}

func Test_CastedResult_IsNotNull_False(t *testing.T) {
	convey.Convey("CastedResult.IsNotNull false when IsNull=true", t, func() {
		cr := &coredynamic.CastedResult{IsNull: true}
		convey.So(cr.IsNotNull(), should.BeFalse)
	})
}

// =============================================================================
// CastedResult — IsNotPointer
// =============================================================================

func Test_CastedResult_IsNotPointer_NilReceiver(t *testing.T) {
	convey.Convey("CastedResult.IsNotPointer false on nil receiver", t, func() {
		var cr *coredynamic.CastedResult
		convey.So(cr.IsNotPointer(), should.BeFalse)
	})
}

func Test_CastedResult_IsNotPointer_True(t *testing.T) {
	convey.Convey("CastedResult.IsNotPointer true when IsPointer=false", t, func() {
		cr := &coredynamic.CastedResult{IsPointer: false}
		convey.So(cr.IsNotPointer(), should.BeTrue)
	})
}

func Test_CastedResult_IsNotPointer_False(t *testing.T) {
	convey.Convey("CastedResult.IsNotPointer false when IsPointer=true", t, func() {
		cr := &coredynamic.CastedResult{IsPointer: true}
		convey.So(cr.IsNotPointer(), should.BeFalse)
	})
}

// =============================================================================
// CastedResult — IsNotMatchingAcceptedType
// =============================================================================

func Test_CastedResult_IsNotMatchingAcceptedType_NilReceiver(t *testing.T) {
	convey.Convey("CastedResult.IsNotMatchingAcceptedType false on nil receiver", t, func() {
		var cr *coredynamic.CastedResult
		convey.So(cr.IsNotMatchingAcceptedType(), should.BeFalse)
	})
}

func Test_CastedResult_IsNotMatchingAcceptedType_True(t *testing.T) {
	convey.Convey("CastedResult.IsNotMatchingAcceptedType true when not matching", t, func() {
		cr := &coredynamic.CastedResult{IsMatchingAcceptedType: false}
		convey.So(cr.IsNotMatchingAcceptedType(), should.BeTrue)
	})
}

func Test_CastedResult_IsNotMatchingAcceptedType_False(t *testing.T) {
	convey.Convey("CastedResult.IsNotMatchingAcceptedType false when matching", t, func() {
		cr := &coredynamic.CastedResult{IsMatchingAcceptedType: true}
		convey.So(cr.IsNotMatchingAcceptedType(), should.BeFalse)
	})
}

// =============================================================================
// CastedResult — IsSourceKind
// =============================================================================

func Test_CastedResult_IsSourceKind_NilReceiver(t *testing.T) {
	convey.Convey("CastedResult.IsSourceKind false on nil receiver", t, func() {
		var cr *coredynamic.CastedResult
		convey.So(cr.IsSourceKind(reflect.String), should.BeFalse)
	})
}

func Test_CastedResult_IsSourceKind_Match(t *testing.T) {
	convey.Convey("CastedResult.IsSourceKind true on kind match", t, func() {
		cr := &coredynamic.CastedResult{SourceKind: reflect.Int}
		convey.So(cr.IsSourceKind(reflect.Int), should.BeTrue)
	})
}

func Test_CastedResult_IsSourceKind_NoMatch(t *testing.T) {
	convey.Convey("CastedResult.IsSourceKind false on mismatch", t, func() {
		cr := &coredynamic.CastedResult{SourceKind: reflect.Int}
		convey.So(cr.IsSourceKind(reflect.String), should.BeFalse)
	})
}

// =============================================================================
// CastedResult — HasError
// =============================================================================

func Test_CastedResult_HasError_NilReceiver(t *testing.T) {
	convey.Convey("CastedResult.HasError false on nil receiver", t, func() {
		var cr *coredynamic.CastedResult
		convey.So(cr.HasError(), should.BeFalse)
	})
}

func Test_CastedResult_HasError_True(t *testing.T) {
	convey.Convey("CastedResult.HasError true when error present", t, func() {
		cr := &coredynamic.CastedResult{Error: errors.New("fail")}
		convey.So(cr.HasError(), should.BeTrue)
	})
}

func Test_CastedResult_HasError_False(t *testing.T) {
	convey.Convey("CastedResult.HasError false when no error", t, func() {
		cr := &coredynamic.CastedResult{}
		convey.So(cr.HasError(), should.BeFalse)
	})
}

// =============================================================================
// CastedResult — HasAnyIssues
// =============================================================================

func Test_CastedResult_HasAnyIssues_NilReceiver(t *testing.T) {
	convey.Convey("CastedResult.HasAnyIssues true on nil receiver", t, func() {
		var cr *coredynamic.CastedResult
		convey.So(cr.HasAnyIssues(), should.BeTrue)
	})
}

func Test_CastedResult_HasAnyIssues_Invalid(t *testing.T) {
	convey.Convey("CastedResult.HasAnyIssues true when invalid", t, func() {
		cr := &coredynamic.CastedResult{IsValid: false, IsMatchingAcceptedType: true}
		convey.So(cr.HasAnyIssues(), should.BeTrue)
	})
}

func Test_CastedResult_HasAnyIssues_Null(t *testing.T) {
	convey.Convey("CastedResult.HasAnyIssues true when null", t, func() {
		cr := &coredynamic.CastedResult{IsValid: true, IsNull: true, IsMatchingAcceptedType: true}
		convey.So(cr.HasAnyIssues(), should.BeTrue)
	})
}

func Test_CastedResult_HasAnyIssues_NotMatching(t *testing.T) {
	convey.Convey("CastedResult.HasAnyIssues true when type not matching", t, func() {
		cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: false}
		convey.So(cr.HasAnyIssues(), should.BeTrue)
	})
}

func Test_CastedResult_HasAnyIssues_AllGood(t *testing.T) {
	convey.Convey("CastedResult.HasAnyIssues false when all good", t, func() {
		cr := &coredynamic.CastedResult{
			IsValid:                true,
			IsNull:                 false,
			IsMatchingAcceptedType: true,
		}
		convey.So(cr.HasAnyIssues(), should.BeFalse)
	})
}

// =============================================================================
// CastedResult — SourceReflectType field
// =============================================================================

func Test_CastedResult_SourceReflectType(t *testing.T) {
	convey.Convey("CastedResult stores SourceReflectType correctly", t, func() {
		cr := &coredynamic.CastedResult{
			SourceReflectType: reflect.TypeOf(""),
			SourceKind:        reflect.String,
		}
		convey.So(cr.SourceReflectType, should.Equal, reflect.TypeOf(""))
		convey.So(cr.IsSourceKind(reflect.String), should.BeTrue)
	})
}

// =============================================================================
// CastedResult — Casted field
// =============================================================================

func Test_CastedResult_Casted_Value(t *testing.T) {
	convey.Convey("CastedResult.Casted stores casted value", t, func() {
		cr := &coredynamic.CastedResult{
			Casted:                 42,
			IsValid:                true,
			IsMatchingAcceptedType: true,
		}
		convey.So(cr.Casted, should.Equal, 42)
		convey.So(cr.HasAnyIssues(), should.BeFalse)
	})
}

func Test_CastedResult_IsSourcePointer(t *testing.T) {
	convey.Convey("CastedResult.IsSourcePointer field works", t, func() {
		cr := &coredynamic.CastedResult{IsSourcePointer: true}
		convey.So(cr.IsSourcePointer, should.BeTrue)
	})
}
