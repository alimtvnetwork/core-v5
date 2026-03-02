package coreinstructiontests

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coreinstruction"
)

// --- Length ---

func Test_IdentifiersWithGlobals_Length_Empty(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	// Act
	length := ids.Length()

	// Assert
	convey.Convey("Length - empty should return 0", t, func() {
		convey.So(length, should.Equal, 0)
	})
}

func Test_IdentifiersWithGlobals_Length_WithItems(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")

	// Act
	length := ids.Length()

	// Assert
	convey.Convey("Length - 3 items should return 3", t, func() {
		convey.So(length, should.Equal, 3)
	})
}

func Test_IdentifiersWithGlobals_Length_NilReceiver(t *testing.T) {
	// Arrange
	var nilIds *coreinstruction.IdentifiersWithGlobals

	// Act
	length := nilIds.Length()

	// Assert
	convey.Convey("Length - nil receiver should return 0", t, func() {
		convey.So(length, should.Equal, 0)
	})
}

// --- GetById ---

func Test_IdentifiersWithGlobals_GetById_Found(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha", "beta")

	// Act
	found := ids.GetById("beta")

	// Assert
	convey.Convey("GetById - existing id should return matching item", t, func() {
		convey.So(found, should.NotBeNil)
		convey.So(found.Id, should.Equal, "beta")
		convey.So(found.IsGlobal, should.BeTrue)
	})
}

func Test_IdentifiersWithGlobals_GetById_NotFound(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "alpha")

	// Act
	found := ids.GetById("missing")

	// Assert
	convey.Convey("GetById - missing id should return nil", t, func() {
		convey.So(found, should.BeNil)
	})
}

func Test_IdentifiersWithGlobals_GetById_EmptyId(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "alpha")

	// Act
	found := ids.GetById("")

	// Assert
	convey.Convey("GetById - empty string id should return nil", t, func() {
		convey.So(found, should.BeNil)
	})
}

// --- Clone ---

func Test_IdentifiersWithGlobals_Clone_Independence(t *testing.T) {
	// Arrange
	original := coreinstruction.NewIdentifiersWithGlobals(true, "x", "y")

	// Act
	cloned := original.Clone()
	cloned.Add(false, "z")

	// Assert
	convey.Convey("Clone - modifying clone should not affect original", t, func() {
		convey.So(original.Length(), should.Equal, 2)
		convey.So(cloned.Length(), should.Equal, 3)
	})
}

func Test_IdentifiersWithGlobals_Clone_Empty(t *testing.T) {
	// Arrange
	original := coreinstruction.EmptyIdentifiersWithGlobals()

	// Act
	cloned := original.Clone()

	// Assert
	convey.Convey("Clone - empty collection should clone to empty", t, func() {
		convey.So(cloned, should.NotBeNil)
		convey.So(cloned.Length(), should.Equal, 0)
	})
}

func Test_IdentifiersWithGlobals_Clone_PreservesValues(t *testing.T) {
	// Arrange
	original := coreinstruction.NewIdentifiersWithGlobals(false, "id-1", "id-2")

	// Act
	cloned := original.Clone()

	// Assert
	convey.Convey("Clone - cloned items should preserve id and isGlobal", t, func() {
		item := cloned.GetById("id-1")
		convey.So(item, should.NotBeNil)
		convey.So(item.Id, should.Equal, "id-1")
		convey.So(item.IsGlobal, should.BeFalse)
	})
}

// --- Add ---

func Test_IdentifiersWithGlobals_Add_Single(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	// Act
	ids.Add(true, "new-id")

	// Assert
	convey.Convey("Add - single item should increase length to 1", t, func() {
		convey.So(ids.Length(), should.Equal, 1)
		found := ids.GetById("new-id")
		convey.So(found, should.NotBeNil)
		convey.So(found.IsGlobal, should.BeTrue)
	})
}

func Test_IdentifiersWithGlobals_Add_EmptyId_Ignored(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	// Act
	ids.Add(true, "")

	// Assert
	convey.Convey("Add - empty string id should be ignored", t, func() {
		convey.So(ids.Length(), should.Equal, 0)
	})
}

func Test_IdentifiersWithGlobals_Add_Multiple(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "existing")

	// Act
	ids.Add(true, "second")
	ids.Add(false, "third")

	// Assert
	convey.Convey("Add - multiple adds should accumulate", t, func() {
		convey.So(ids.Length(), should.Equal, 3)
		second := ids.GetById("second")
		convey.So(second, should.NotBeNil)
		convey.So(second.IsGlobal, should.BeTrue)
		third := ids.GetById("third")
		convey.So(third, should.NotBeNil)
		convey.So(third.IsGlobal, should.BeFalse)
	})
}
