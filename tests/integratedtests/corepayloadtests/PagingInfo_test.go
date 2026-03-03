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
