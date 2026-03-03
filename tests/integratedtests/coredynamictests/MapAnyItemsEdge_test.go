package coredynamictests

import (
	"testing"

	"github.com/smarty/assertions/should"
	convey "github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/coredata/coredynamic"
)

// =============================================================================
// MapAnyItems — IsEqual
// =============================================================================

func Test_MapAnyItems_IsEqual_BothNil(t *testing.T) {
	// Arrange
	var left *coredynamic.MapAnyItems
	var right *coredynamic.MapAnyItems

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - both nil should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_MapAnyItems_IsEqual_LeftNil(t *testing.T) {
	// Arrange
	var left *coredynamic.MapAnyItems
	right := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - left nil should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_MapAnyItems_IsEqual_RightNil(t *testing.T) {
	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})

	// Act
	result := left.IsEqual(nil)

	// Assert
	convey.Convey("IsEqual - right nil should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_MapAnyItems_IsEqual_SameContent(t *testing.T) {
	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "1", "b": "2"})
	right := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "1", "b": "2"})

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - same content should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_MapAnyItems_IsEqual_DifferentValues(t *testing.T) {
	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "1"})
	right := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "2"})

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - different values should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_MapAnyItems_IsEqual_DifferentKeys(t *testing.T) {
	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "1"})
	right := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"b": "1"})

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - different keys should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_MapAnyItems_IsEqual_DifferentLengths(t *testing.T) {
	// Arrange
	left := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "1"})
	right := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "1", "b": "2"})

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - different lengths should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_MapAnyItems_IsEqual_BothEmpty(t *testing.T) {
	// Arrange
	left := coredynamic.EmptyMapAnyItems()
	right := coredynamic.EmptyMapAnyItems()

	// Act
	result := left.IsEqual(right)

	// Assert
	convey.Convey("IsEqual - both empty should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// =============================================================================
// MapAnyItems — IsEqualRaw
// =============================================================================

func Test_MapAnyItems_IsEqualRaw_NilReceiver(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	result := m.IsEqualRaw(map[string]any{"k": "v"})

	// Assert
	convey.Convey("IsEqualRaw - nil receiver should return false", t, func() {
		convey.So(result, should.BeFalse)
	})
}

func Test_MapAnyItems_IsEqualRaw_NilReceiverNilMap(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	result := m.IsEqualRaw(nil)

	// Assert
	convey.Convey("IsEqualRaw - nil receiver nil map should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

func Test_MapAnyItems_IsEqualRaw_Match(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"x": "y"})

	// Act
	result := m.IsEqualRaw(map[string]any{"x": "y"})

	// Assert
	convey.Convey("IsEqualRaw - matching map should return true", t, func() {
		convey.So(result, should.BeTrue)
	})
}

// =============================================================================
// MapAnyItems — ClonePtr
// =============================================================================

func Test_MapAnyItems_ClonePtr_NilReceiver(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	clone, err := m.ClonePtr()

	// Assert
	convey.Convey("ClonePtr - nil receiver should return nil and error", t, func() {
		convey.So(clone, should.BeNil)
		convey.So(err, should.NotBeNil)
	})
}

func Test_MapAnyItems_ClonePtr_ValidData(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{
		"name": "alice",
		"age":  float64(30),
	})

	// Act
	clone, err := m.ClonePtr()

	// Assert
	convey.Convey("ClonePtr - valid data should clone successfully", t, func() {
		convey.So(err, should.BeNil)
		convey.So(clone, should.NotBeNil)
		convey.So(clone.Length(), should.Equal, 2)
		convey.So(clone.HasKey("name"), should.BeTrue)
		convey.So(clone.HasKey("age"), should.BeTrue)
	})
}

func Test_MapAnyItems_ClonePtr_Independence(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{
		"key": "original",
	})

	// Act
	clone, err := m.ClonePtr()
	clone.Add("new_key", "new_val")

	// Assert
	convey.Convey("ClonePtr - modifying clone should not affect original", t, func() {
		convey.So(err, should.BeNil)
		convey.So(m.HasKey("new_key"), should.BeFalse)
		convey.So(clone.HasKey("new_key"), should.BeTrue)
	})
}

func Test_MapAnyItems_ClonePtr_EmptyMap(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	clone, err := m.ClonePtr()

	// Assert
	convey.Convey("ClonePtr - empty map should clone to empty", t, func() {
		convey.So(err, should.BeNil)
		convey.So(clone, should.NotBeNil)
		convey.So(clone.Length(), should.Equal, 0)
	})
}

// =============================================================================
// MapAnyItems — Length / IsEmpty / HasAnyItem nil safety
// =============================================================================

func Test_MapAnyItems_Length_NilReceiver(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act & Assert
	convey.Convey("Length/IsEmpty/HasAnyItem - nil receiver safety", t, func() {
		convey.So(m.Length(), should.Equal, 0)
		convey.So(m.IsEmpty(), should.BeTrue)
		convey.So(m.HasAnyItem(), should.BeFalse)
	})
}

// =============================================================================
// MapAnyItems — HasKey
// =============================================================================

func Test_MapAnyItems_HasKey_NilReceiver(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act & Assert
	convey.Convey("HasKey - nil receiver should return false", t, func() {
		convey.So(m.HasKey("anything"), should.BeFalse)
	})
}

func Test_MapAnyItems_HasKey_Exists(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"key": "val"})

	// Act & Assert
	convey.Convey("HasKey - existing key should return true", t, func() {
		convey.So(m.HasKey("key"), should.BeTrue)
	})
}

func Test_MapAnyItems_HasKey_Missing(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"key": "val"})

	// Act & Assert
	convey.Convey("HasKey - missing key should return false", t, func() {
		convey.So(m.HasKey("nope"), should.BeFalse)
	})
}

// =============================================================================
// MapAnyItems — Add
// =============================================================================

func Test_MapAnyItems_Add_NewKey(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	isNew := m.Add("k", "v")

	// Assert
	convey.Convey("Add - new key should return true", t, func() {
		convey.So(isNew, should.BeTrue)
		convey.So(m.Length(), should.Equal, 1)
	})
}

func Test_MapAnyItems_Add_ExistingKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "old"})

	// Act
	isNew := m.Add("k", "new")

	// Assert
	convey.Convey("Add - existing key should return false and overwrite", t, func() {
		convey.So(isNew, should.BeFalse)
		convey.So(m.GetValue("k"), should.Equal, "new")
	})
}
