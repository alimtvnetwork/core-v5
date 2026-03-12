package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
)

// ==========================================
// DiffLeftRight
// ==========================================

func Test_DiffLeftRight_Types(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: 1}
	l, r := d.Types()
	if l == nil || r == nil {
		t.Error("types should not be nil")
	}
}

func Test_DiffLeftRight_IsSameTypeSame_True(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	if !d.IsSameTypeSame() {
		t.Error("same type should return true")
	}
}

func Test_DiffLeftRight_IsSameTypeSame_False(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: 1}
	if d.IsSameTypeSame() {
		t.Error("different types should return false")
	}
}

func Test_DiffLeftRight_IsSame_True(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}
	if !d.IsSame() {
		t.Error("same values should return true")
	}
}

func Test_DiffLeftRight_IsSame_False(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	if d.IsSame() {
		t.Error("different values should return false")
	}
}

func Test_DiffLeftRight_IsSameRegardlessOfType(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	if !d.IsSameRegardlessOfType() {
		t.Error("same value regardless of type should be true")
	}
}

func Test_DiffLeftRight_IsEqual_Regardless(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	if !d.IsEqual(true) {
		t.Error("should be equal regardless")
	}
}

func Test_DiffLeftRight_IsEqual_Strict(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}
	if !d.IsEqual(false) {
		t.Error("should be equal strict")
	}
}

func Test_DiffLeftRight_HasMismatch_Regardless(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	if !d.HasMismatch(true) {
		t.Error("should have mismatch")
	}
}

func Test_DiffLeftRight_HasMismatch_Strict(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	if !d.HasMismatch(false) {
		t.Error("should have mismatch")
	}
}

func Test_DiffLeftRight_IsNotEqual(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	if !d.IsNotEqual() {
		t.Error("should not be equal")
	}
}

func Test_DiffLeftRight_String(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	if d.String() == "" {
		t.Error("should return non-empty")
	}
}

func Test_DiffLeftRight_JsonString_Nil(t *testing.T) {
	var d *enumimpl.DiffLeftRight
	if d.JsonString() != "" {
		t.Error("nil should return empty")
	}
}

func Test_DiffLeftRight_SpecificFullString(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	l, r := d.SpecificFullString()
	if l == "" || r == "" {
		t.Error("should return non-empty")
	}
}

func Test_DiffLeftRight_DiffString_Same(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "a"}
	if d.DiffString() != "" {
		t.Error("same values should return empty diff")
	}
}

func Test_DiffLeftRight_DiffString_Different(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	if d.DiffString() == "" {
		t.Error("different values should return non-empty diff")
	}
}

// ==========================================
// DynamicMap
// ==========================================

func Test_DynamicMap_AddOrUpdate(t *testing.T) {
	dm := enumimpl.DynamicMap{"key1": "val1"}
	isNew := dm.AddOrUpdate("key2", "val2")
	if !isNew {
		t.Error("should be newly added")
	}
	isNew2 := dm.AddOrUpdate("key1", "updated")
	if isNew2 {
		t.Error("should be updated, not new")
	}
}

func Test_DynamicMap_AllKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	keys := dm.AllKeys()
	if len(keys) != 2 {
		t.Errorf("expected 2, got %d", len(keys))
	}
}

func Test_DynamicMap_AllKeys_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	keys := dm.AllKeys()
	if len(keys) != 0 {
		t.Errorf("expected 0, got %d", len(keys))
	}
}

func Test_DynamicMap_AllKeysSorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeysSorted()
	if len(keys) != 2 || keys[0] != "a" {
		t.Errorf("expected sorted [a b], got %v", keys)
	}
}

func Test_DynamicMap_AllValuesStrings(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	vals := dm.AllValuesStrings()
	if len(vals) != 1 {
		t.Errorf("expected 1, got %d", len(vals))
	}
}

func Test_DynamicMap_AllValuesStringsSorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "b", "c": "a"}
	vals := dm.AllValuesStringsSorted()
	if len(vals) != 2 {
		t.Errorf("expected 2, got %d", len(vals))
	}
}

func Test_DynamicMap_Length(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if dm.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_DynamicMap_Length_Nil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	if dm.Length() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_DynamicMap_Count(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if dm.Count() != 1 {
		t.Error("expected 1")
	}
}

func Test_DynamicMap_IsEmpty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if !dm.IsEmpty() {
		t.Error("should be empty")
	}
}

func Test_DynamicMap_HasAnyItem(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.HasAnyItem() {
		t.Error("should have items")
	}
}

func Test_DynamicMap_HasKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.HasKey("a") {
		t.Error("should have key 'a'")
	}
	if dm.HasKey("b") {
		t.Error("should not have key 'b'")
	}
}

func Test_DynamicMap_HasAllKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	if !dm.HasAllKeys("a", "b") {
		t.Error("should have all keys")
	}
	if dm.HasAllKeys("a", "c") {
		t.Error("should not have all keys")
	}
}

func Test_DynamicMap_HasAnyKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.HasAnyKeys("a", "b") {
		t.Error("should have any key")
	}
	if dm.HasAnyKeys("b", "c") {
		t.Error("should not have any key")
	}
}

func Test_DynamicMap_IsMissingKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.IsMissingKey("b") {
		t.Error("should be missing 'b'")
	}
}

func Test_DynamicMap_Raw(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	raw := dm.Raw()
	if len(raw) != 1 {
		t.Error("raw should have 1 item")
	}
}

func Test_DynamicMap_First(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	k, v := dm.First()
	if k == "" || v == nil {
		t.Error("should return first item")
	}
}

func Test_DynamicMap_First_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	k, v := dm.First()
	if k != "" || v != nil {
		t.Error("empty map should return empty first")
	}
}

func Test_DynamicMap_IsEqual_BothNil(t *testing.T) {
	var a, b *enumimpl.DynamicMap
	if !a.IsEqual(true, b) {
		t.Error("both nil should be equal")
	}
}

func Test_DynamicMap_IsEqual_OneNil(t *testing.T) {
	dm := &enumimpl.DynamicMap{"a": 1}
	if dm.IsEqual(true, nil) {
		t.Error("one nil should not be equal")
	}
}

func Test_DynamicMap_IsEqual_Same(t *testing.T) {
	dm := &enumimpl.DynamicMap{"a": 1}
	dm2 := &enumimpl.DynamicMap{"a": 1}
	if !dm.IsEqual(true, dm2) {
		t.Error("should be equal")
	}
}

func Test_DynamicMap_IsMismatch(t *testing.T) {
	dm := &enumimpl.DynamicMap{"a": 1}
	dm2 := &enumimpl.DynamicMap{"a": 2}
	if !dm.IsMismatch(false, dm2) {
		t.Error("should be mismatch")
	}
}

// ==========================================
// Format
// ==========================================

func Test_Format(t *testing.T) {
	result := enumimpl.Format(
		"MyEnum",
		"Invalid",
		"0",
		"Enum of {type-name} - {name} - {value}",
	)
	if result == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// NameWithValue
// ==========================================

func Test_NameWithValue(t *testing.T) {
	result := enumimpl.NameWithValue("test")
	if result == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// UnsupportedNames
// ==========================================

func Test_UnsupportedNames_AllSupported(t *testing.T) {
	result := enumimpl.UnsupportedNames([]string{"a", "b"}, "a", "b")
	if len(result) != 0 {
		t.Errorf("expected 0 unsupported, got %d", len(result))
	}
}

func Test_UnsupportedNames_SomeUnsupported(t *testing.T) {
	result := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")
	if len(result) != 2 {
		t.Errorf("expected 2 unsupported, got %d", len(result))
	}
}

// ==========================================
// OnlySupportedErr
// ==========================================

func Test_OnlySupportedErr_AllSupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b"}, "a", "b")
	if err != nil {
		t.Error("all supported should return nil")
	}
}

func Test_OnlySupportedErr_SomeUnsupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b", "c"}, "a")
	if err == nil {
		t.Error("unsupported should return error")
	}
}

func Test_OnlySupportedErr_EmptyAllNames(t *testing.T) {
	err := enumimpl.OnlySupportedErr(0, []string{})
	if err != nil {
		t.Error("empty allNames should return nil")
	}
}
