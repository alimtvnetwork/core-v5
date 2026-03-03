package coreapitests

import (
	"testing"

	"github.com/smarty/assertions/should"
	convey "github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coreapi"
)

// --- IsPageSizeEmpty ---

func Test_PageRequest_IsPageSizeEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var req *coreapi.PageRequest

	// Act
	result := req.IsPageSizeEmpty()

	// Assert
	convey.Convey("IsPageSizeEmpty - nil receiver should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PageRequest_IsPageSizeEmpty_Zero(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageSize: 0, PageIndex: 1}

	// Act
	result := req.IsPageSizeEmpty()

	// Assert
	convey.Convey("IsPageSizeEmpty - zero should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PageRequest_IsPageSizeEmpty_Negative(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageSize: -1}

	// Act
	result := req.IsPageSizeEmpty()

	// Assert
	convey.Convey("IsPageSizeEmpty - negative should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PageRequest_IsPageSizeEmpty_Positive(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageSize: 10}

	// Act
	result := req.IsPageSizeEmpty()

	// Assert
	convey.Convey("IsPageSizeEmpty - positive should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- IsPageIndexEmpty ---

func Test_PageRequest_IsPageIndexEmpty_NilReceiver(t *testing.T) {
	// Arrange
	var req *coreapi.PageRequest

	// Act
	result := req.IsPageIndexEmpty()

	// Assert
	convey.Convey("IsPageIndexEmpty - nil receiver should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PageRequest_IsPageIndexEmpty_Zero(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageIndex: 0, PageSize: 10}

	// Act
	result := req.IsPageIndexEmpty()

	// Assert
	convey.Convey("IsPageIndexEmpty - zero should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_PageRequest_IsPageIndexEmpty_Positive(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageIndex: 2}

	// Act
	result := req.IsPageIndexEmpty()

	// Assert
	convey.Convey("IsPageIndexEmpty - positive should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

// --- HasPageSize ---

func Test_PageRequest_HasPageSize_NilReceiver(t *testing.T) {
	// Arrange
	var req *coreapi.PageRequest

	// Act
	result := req.HasPageSize()

	// Assert
	convey.Convey("HasPageSize - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PageRequest_HasPageSize_Positive(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageSize: 25}

	// Act
	result := req.HasPageSize()

	// Assert
	convey.Convey("HasPageSize - positive should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- HasPageIndex ---

func Test_PageRequest_HasPageIndex_NilReceiver(t *testing.T) {
	// Arrange
	var req *coreapi.PageRequest

	// Act
	result := req.HasPageIndex()

	// Assert
	convey.Convey("HasPageIndex - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_PageRequest_HasPageIndex_Positive(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageIndex: 3}

	// Act
	result := req.HasPageIndex()

	// Assert
	convey.Convey("HasPageIndex - positive should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// --- Clone ---

func Test_PageRequest_Clone_NilReceiver(t *testing.T) {
	// Arrange
	var req *coreapi.PageRequest

	// Act
	result := req.Clone()

	// Assert
	convey.Convey("Clone - nil receiver should return nil", t, func() {
		convey.So(result, should.BeNil)
	})
}

func Test_PageRequest_Clone_CopiesAllFields(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageSize: 20, PageIndex: 5}

	// Act
	result := req.Clone()

	// Assert
	convey.Convey("Clone - should copy all fields", t, func() {
		convey.So(result, should.NotBeNil)
		convey.So(result.PageSize, should.Equal, 20)
		convey.So(result.PageIndex, should.Equal, 5)
	})
}

func Test_PageRequest_Clone_Independence(t *testing.T) {
	// Arrange
	req := &coreapi.PageRequest{PageSize: 20, PageIndex: 5}

	// Act
	clone := req.Clone()
	clone.PageSize = 99
	clone.PageIndex = 99

	// Assert
	convey.Convey("Clone - modifying clone should not affect original", t, func() {
		convey.So(req.PageSize, should.Equal, 20)
		convey.So(req.PageIndex, should.Equal, 5)
	})
}
