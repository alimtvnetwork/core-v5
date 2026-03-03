package corepayloadtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corepayload"

	convey "github.com/smartystreets/goconvey/convey"
	"github.com/smarty/assertions/should"
)

// --- IsEqual ---

func Test_PagingInfo_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var left *corepayload.PagingInfo
	var right *corepayload.PagingInfo

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - both nil should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsEqual_LeftNil(t *testing.T) {
	// Arrange
	var left *corepayload.PagingInfo
	right := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - left nil right non-nil should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_IsEqual_RightNil(t *testing.T) {
	// Arrange
	left := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 50}

	// Act
	result := left.IsEqual(nil)

	// Assert
	convey.Convey("IsEqual - left non-nil right nil should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_IsEqual_EqualValues(t *testing.T) {
	// Arrange
	left := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 25}
	right := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 25}

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - identical values should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsEqual_DifferentTotalPages(t *testing.T) {
	// Arrange
	left := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 25}
	right := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 25}

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - different TotalPages should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_IsEqual_DifferentCurrentPageIndex(t *testing.T) {
	// Arrange
	left := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 1, PerPageItems: 10, TotalItems: 25}
	right := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 25}

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - different CurrentPageIndex should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_IsEqual_DifferentPerPageItems(t *testing.T) {
	// Arrange
	left := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 25}
	right := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 20, TotalItems: 25}

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - different PerPageItems should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_IsEqual_DifferentTotalItems(t *testing.T) {
	// Arrange
	left := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 25}
	right := &corepayload.PagingInfo{TotalPages: 3, CurrentPageIndex: 2, PerPageItems: 10, TotalItems: 30}

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - different TotalItems should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- ClonePtr ---

func Test_PagingInfo_ClonePtr_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.ClonePtr()

	// Assert
	convey.Convey("ClonePtr - nil receiver should return nil", t, func() {
		convey.So(result, should.BeNil)
	})
}

func Test_PagingInfo_ClonePtr_CopiesAllFields(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	// Act
	result := info.ClonePtr()

	// Assert
	convey.Convey("ClonePtr - should copy all fields", t, func() {
		convey.So(result, should.NotBeNil)
		convey.So(result.TotalPages, should.Equal, 5)
		convey.So(result.CurrentPageIndex, should.Equal, 3)
		convey.So(result.PerPageItems, should.Equal, 10)
		convey.So(result.TotalItems, should.Equal, 50)
	})
}

func Test_PagingInfo_ClonePtr_Independence(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	// Act
	clone := info.ClonePtr()
	clone.TotalPages = 99
	clone.CurrentPageIndex = 99

	// Assert
	convey.Convey("ClonePtr - modifying clone should not affect original", t, func() {
		convey.So(info.TotalPages, should.Equal, 5)
		convey.So(info.CurrentPageIndex, should.Equal, 3)
	})
}

// --- IsEmpty ---

func Test_PagingInfo_IsEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.IsEmpty()

	// Assert
	convey.Convey("IsEmpty - nil receiver should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsEmpty_ZeroValues(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{}

	// Act
	result := info.IsEmpty()

	// Assert
	convey.Convey("IsEmpty - zero values should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsEmpty_NonZero(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 3, TotalItems: 25}

	// Act
	result := info.IsEmpty()

	// Assert
	convey.Convey("IsEmpty - non-zero values should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- HasTotalPages ---

func Test_PagingInfo_HasTotalPages_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.HasTotalPages()

	// Assert
	convey.Convey("HasTotalPages - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_HasTotalPages_Zero(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 0}

	// Act
	result := info.HasTotalPages()

	// Assert
	convey.Convey("HasTotalPages - zero should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_HasTotalPages_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 5}

	// Act
	result := info.HasTotalPages()

	// Assert
	convey.Convey("HasTotalPages - positive should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- HasCurrentPageIndex ---

func Test_PagingInfo_HasCurrentPageIndex_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.HasCurrentPageIndex()

	// Assert
	convey.Convey("HasCurrentPageIndex - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_HasCurrentPageIndex_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{CurrentPageIndex: 2}

	// Act
	result := info.HasCurrentPageIndex()

	// Assert
	convey.Convey("HasCurrentPageIndex - positive should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- HasPerPageItems ---

func Test_PagingInfo_HasPerPageItems_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.HasPerPageItems()

	// Assert
	convey.Convey("HasPerPageItems - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_HasPerPageItems_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{PerPageItems: 10}

	// Act
	result := info.HasPerPageItems()

	// Assert
	convey.Convey("HasPerPageItems - positive should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- HasTotalItems ---

func Test_PagingInfo_HasTotalItems_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.HasTotalItems()

	// Assert
	convey.Convey("HasTotalItems - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PagingInfo_HasTotalItems_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalItems: 50}

	// Act
	result := info.HasTotalItems()

	// Assert
	convey.Convey("HasTotalItems - positive should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- IsInvalidTotalPages ---

func Test_PagingInfo_IsInvalidTotalPages_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.IsInvalidTotalPages()

	// Assert
	convey.Convey("IsInvalidTotalPages - nil receiver should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsInvalidTotalPages_Zero(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 0}

	// Act
	result := info.IsInvalidTotalPages()

	// Assert
	convey.Convey("IsInvalidTotalPages - zero should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsInvalidTotalPages_Negative(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: -1}

	// Act
	result := info.IsInvalidTotalPages()

	// Assert
	convey.Convey("IsInvalidTotalPages - negative should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsInvalidTotalPages_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 5}

	// Act
	result := info.IsInvalidTotalPages()

	// Assert
	convey.Convey("IsInvalidTotalPages - positive should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- IsInvalidCurrentPageIndex ---

func Test_PagingInfo_IsInvalidCurrentPageIndex_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.IsInvalidCurrentPageIndex()

	// Assert
	convey.Convey("IsInvalidCurrentPageIndex - nil receiver should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsInvalidCurrentPageIndex_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{CurrentPageIndex: 1}

	// Act
	result := info.IsInvalidCurrentPageIndex()

	// Assert
	convey.Convey("IsInvalidCurrentPageIndex - positive should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- IsInvalidPerPageItems ---

func Test_PagingInfo_IsInvalidPerPageItems_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.IsInvalidPerPageItems()

	// Assert
	convey.Convey("IsInvalidPerPageItems - nil receiver should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsInvalidPerPageItems_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{PerPageItems: 10}

	// Act
	result := info.IsInvalidPerPageItems()

	// Assert
	convey.Convey("IsInvalidPerPageItems - positive should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- IsInvalidTotalItems ---

func Test_PagingInfo_IsInvalidTotalItems_NilReceiver(t *testing.T) {
	// Arrange
	var info *corepayload.PagingInfo

	// Act
	result := info.IsInvalidTotalItems()

	// Assert
	convey.Convey("IsInvalidTotalItems - nil receiver should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PagingInfo_IsInvalidTotalItems_Positive(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalItems: 50}

	// Act
	result := info.IsInvalidTotalItems()

	// Assert
	convey.Convey("IsInvalidTotalItems - positive should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- Clone (value) ---

func Test_PagingInfo_Clone_CopiesAllFields(t *testing.T) {
	// Arrange
	info := corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	// Act
	clone := info.Clone()

	// Assert
	convey.Convey("Clone - should copy all fields", t, func() {
		convey.So(clone.TotalPages, should.Equal, 5)
		convey.So(clone.CurrentPageIndex, should.Equal, 3)
		convey.So(clone.PerPageItems, should.Equal, 10)
		convey.So(clone.TotalItems, should.Equal, 50)
	})
}

func Test_PagingInfo_Clone_Independence(t *testing.T) {
	// Arrange
	info := corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	// Act
	clone := info.Clone()
	clone.TotalPages = 99

	// Assert
	convey.Convey("Clone - modifying clone should not affect original", t, func() {
		convey.So(info.TotalPages, should.Equal, 5)
	})
}
