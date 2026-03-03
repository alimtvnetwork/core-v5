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

// --- IsEmpty / HasAnyItem ---

func Test_IdentifiersWithGlobals_IsEmpty_True(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	// Assert
	convey.Convey("IsEmpty - empty collection should return true", t, func() {
		convey.So(ids.IsEmpty(), should.BeTrue)
		convey.So(ids.HasAnyItem(), should.BeFalse)
	})
}

func Test_IdentifiersWithGlobals_IsEmpty_False(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "item")

	// Assert
	convey.Convey("IsEmpty - non-empty should return false", t, func() {
		convey.So(ids.IsEmpty(), should.BeFalse)
		convey.So(ids.HasAnyItem(), should.BeTrue)
	})
}

// --- IndexOf ---

func Test_IdentifiersWithGlobals_IndexOf_Found(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a", "b", "c")

	// Act & Assert
	convey.Convey("IndexOf - existing id returns correct index", t, func() {
		convey.So(ids.IndexOf("a"), should.Equal, 0)
		convey.So(ids.IndexOf("b"), should.Equal, 1)
		convey.So(ids.IndexOf("c"), should.Equal, 2)
	})
}

func Test_IdentifiersWithGlobals_IndexOf_NotFound(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(false, "x")

	// Act
	index := ids.IndexOf("missing")

	// Assert
	convey.Convey("IndexOf - missing id returns -1", t, func() {
		convey.So(index, should.Equal, -1)
	})
}

func Test_IdentifiersWithGlobals_IndexOf_EmptyId(t *testing.T) {
	// Arrange
	ids := coreinstruction.NewIdentifiersWithGlobals(true, "a")

	// Act
	index := ids.IndexOf("")

	// Assert
	convey.Convey("IndexOf - empty string returns -1", t, func() {
		convey.So(index, should.Equal, -1)
	})
}

func Test_IdentifiersWithGlobals_IndexOf_EmptyCollection(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	// Act
	index := ids.IndexOf("any")

	// Assert
	convey.Convey("IndexOf - empty collection returns -1", t, func() {
		convey.So(index, should.Equal, -1)
	})
}

// --- Adds ---

func Test_IdentifiersWithGlobals_Adds_Multiple(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	// Act
	ids.Adds(true, "one", "two", "three")

	// Assert
	convey.Convey("Adds - batch add should add all items", t, func() {
		convey.So(ids.Length(), should.Equal, 3)
		convey.So(ids.GetById("one"), should.NotBeNil)
		convey.So(ids.GetById("two"), should.NotBeNil)
		convey.So(ids.GetById("three"), should.NotBeNil)
	})
}

func Test_IdentifiersWithGlobals_Adds_EmptyIds(t *testing.T) {
	// Arrange
	ids := coreinstruction.EmptyIdentifiersWithGlobals()

	// Act
	ids.Adds(true)

	// Assert
	convey.Convey("Adds - empty ids should not add anything", t, func() {
		convey.So(ids.Length(), should.Equal, 0)
	})
}

// --- NewIdentifiersWithGlobals edge cases ---

func Test_IdentifiersWithGlobals_New_NoIds(t *testing.T) {
	// Arrange & Act
	ids := coreinstruction.NewIdentifiersWithGlobals(true)

	// Assert
	convey.Convey("New - no ids should create empty collection", t, func() {
		convey.So(ids, should.NotBeNil)
		convey.So(ids.Length(), should.Equal, 0)
	})
}
