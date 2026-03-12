package argstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Map basic operations
// ==========================================

func Test_Map_Length(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	if m.Length() != 2 {
		t.Errorf("expected 2, got %d", m.Length())
	}
}

func Test_Map_Has(t *testing.T) {
	m := args.Map{"a": 1}
	if !m.Has("a") {
		t.Error("should have 'a'")
	}
	if m.Has("b") {
		t.Error("should not have 'b'")
	}
}

func Test_Map_Has_Nil(t *testing.T) {
	var m args.Map
	if m.Has("a") {
		t.Error("nil map should return false")
	}
}

func Test_Map_HasDefined(t *testing.T) {
	m := args.Map{"a": "val", "b": nil}
	if !m.HasDefined("a") {
		t.Error("should be defined")
	}
}

func Test_Map_HasDefined_Nil(t *testing.T) {
	var m args.Map
	if m.HasDefined("a") {
		t.Error("nil map should return false")
	}
}

func Test_Map_IsKeyMissing(t *testing.T) {
	m := args.Map{"a": 1}
	if m.IsKeyMissing("a") {
		t.Error("'a' should not be missing")
	}
	if !m.IsKeyMissing("b") {
		t.Error("'b' should be missing")
	}
}

func Test_Map_IsKeyMissing_Nil(t *testing.T) {
	var m args.Map
	if m.IsKeyMissing("a") {
		t.Error("nil map should return false (per implementation)")
	}
}

func Test_Map_IsKeyInvalid(t *testing.T) {
	m := args.Map{"a": "val"}
	if m.IsKeyInvalid("a") {
		t.Error("'a' should not be invalid")
	}
}

func Test_Map_IsKeyInvalid_Nil(t *testing.T) {
	var m args.Map
	if m.IsKeyInvalid("a") {
		t.Error("nil map should return false (per implementation)")
	}
}

func Test_Map_HasDefinedAll(t *testing.T) {
	m := args.Map{"a": "v1", "b": "v2"}
	if !m.HasDefinedAll("a", "b") {
		t.Error("should have all defined")
	}
}

func Test_Map_HasDefinedAll_Nil(t *testing.T) {
	var m args.Map
	if m.HasDefinedAll("a") {
		t.Error("nil should return false")
	}
}

func Test_Map_HasDefinedAll_Empty(t *testing.T) {
	m := args.Map{"a": "v1"}
	if m.HasDefinedAll() {
		t.Error("no names should return false")
	}
}

// ==========================================
// Map Get operations
// ==========================================

func Test_Map_Get_Cov(t *testing.T) {
	m := args.Map{"a": "val"}
	item, isValid := m.Get("a")
	if !isValid || item != "val" {
		t.Error("should return valid item")
	}
}

func Test_Map_Get_Missing(t *testing.T) {
	m := args.Map{"a": "val"}
	_, isValid := m.Get("b")
	if isValid {
		t.Error("missing key should not be valid")
	}
}

func Test_Map_Get_Nil(t *testing.T) {
	var m args.Map
	_, isValid := m.Get("a")
	if isValid {
		t.Error("nil map should not be valid")
	}
}

func Test_Map_GetLowerCase(t *testing.T) {
	m := args.Map{"name": "val"}
	item, isValid := m.GetLowerCase("Name")
	if !isValid || item != "val" {
		t.Error("should find lowercase")
	}
}

func Test_Map_GetDirectLower(t *testing.T) {
	m := args.Map{"name": "val"}
	item := m.GetDirectLower("Name")
	if item != "val" {
		t.Error("should find lowercase")
	}
}

func Test_Map_GetDirectLower_Missing(t *testing.T) {
	m := args.Map{"name": "val"}
	item := m.GetDirectLower("Missing")
	if item != nil {
		t.Error("missing should return nil")
	}
}

// ==========================================
// Map semantic accessors
// ==========================================

func Test_Map_When(t *testing.T) {
	m := args.Map{"when": "condition"}
	if m.When() != "condition" {
		t.Error("should return when value")
	}
}

func Test_Map_Title(t *testing.T) {
	m := args.Map{"title": "test title"}
	if m.Title() != "test title" {
		t.Error("should return title value")
	}
}

func Test_Map_Expect(t *testing.T) {
	m := args.Map{"expect": "value"}
	if m.Expect() != "value" {
		t.Error("should return expect value")
	}
}

func Test_Map_Actual(t *testing.T) {
	m := args.Map{"actual": "value"}
	if m.Actual() != "value" {
		t.Error("should return actual value")
	}
}

func Test_Map_Arrange(t *testing.T) {
	m := args.Map{"arrange": "value"}
	if m.Arrange() != "value" {
		t.Error("should return arrange value")
	}
}

func Test_Map_SetActual_Cov(t *testing.T) {
	m := args.Map{}
	m.SetActual("hello")
	if m.Actual() != "hello" {
		t.Error("should set actual")
	}
}

// ==========================================
// Map numbered items
// ==========================================

func Test_Map_FirstItem(t *testing.T) {
	m := args.Map{"first": "val"}
	if m.FirstItem() != "val" {
		t.Error("should return first item")
	}
}

func Test_Map_SecondItem(t *testing.T) {
	m := args.Map{"second": "val"}
	if m.SecondItem() != "val" {
		t.Error("should return second item")
	}
}

func Test_Map_ThirdItem(t *testing.T) {
	m := args.Map{"third": "val"}
	if m.ThirdItem() != "val" {
		t.Error("should return third item")
	}
}

func Test_Map_FourthItem(t *testing.T) {
	m := args.Map{"fourth": "val"}
	if m.FourthItem() != "val" {
		t.Error("should return fourth item")
	}
}

func Test_Map_FifthItem(t *testing.T) {
	m := args.Map{"fifth": "val"}
	if m.FifthItem() != "val" {
		t.Error("should return fifth item")
	}
}

func Test_Map_SixthItem(t *testing.T) {
	m := args.Map{"sixth": "val"}
	if m.SixthItem() != "val" {
		t.Error("should return sixth item")
	}
}

func Test_Map_Seventh(t *testing.T) {
	m := args.Map{"seventh": "val"}
	if m.Seventh() != "val" {
		t.Error("should return seventh item")
	}
}

// ==========================================
// Map Expected
// ==========================================

func Test_Map_Expected(t *testing.T) {
	m := args.Map{"expected": "val"}
	if m.Expected() != "val" {
		t.Error("should return expected value")
	}
}

func Test_Map_Expected_Alias(t *testing.T) {
	m := args.Map{"expects": "val"}
	if m.Expected() != "val" {
		t.Error("should return expected from alias")
	}
}

func Test_Map_HasExpect(t *testing.T) {
	m := args.Map{"expected": "val"}
	if !m.HasExpect() {
		t.Error("should have expect")
	}
}

func Test_Map_HasFirst(t *testing.T) {
	m := args.Map{"first": "val"}
	if !m.HasFirst() {
		t.Error("should have first")
	}
}

// ==========================================
// Map Raw / Args / ValidArgs
// ==========================================

func Test_Map_Raw(t *testing.T) {
	m := args.Map{"a": 1}
	raw := m.Raw()
	if len(raw) != 1 {
		t.Error("raw should have 1 item")
	}
}

func Test_Map_Args_Cov(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	a := m.Args("a", "b")
	if len(a) != 2 {
		t.Errorf("expected 2 args, got %d", len(a))
	}
}

func Test_Map_GetByIndex_Cov(t *testing.T) {
	m := args.Map{"a": 1}
	v := m.GetByIndex(0)
	if v == nil {
		t.Error("should return value at index 0")
	}
}

func Test_Map_GetByIndex_OutOfBounds(t *testing.T) {
	m := args.Map{"a": 1}
	v := m.GetByIndex(10)
	if v != nil {
		t.Error("out of bounds should return nil")
	}
}

// ==========================================
// Map SortedKeys
// ==========================================

func Test_Map_SortedKeys(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys, err := m.SortedKeys()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(keys) != 2 || keys[0] != "a" {
		t.Errorf("expected sorted [a b], got %v", keys)
	}
}

func Test_Map_SortedKeys_Empty(t *testing.T) {
	m := args.Map{}
	keys, err := m.SortedKeys()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(keys) != 0 {
		t.Error("empty map should return empty keys")
	}
}

func Test_Map_SortedKeysMust(t *testing.T) {
	m := args.Map{"b": 2, "a": 1}
	keys := m.SortedKeysMust()
	if len(keys) != 2 {
		t.Errorf("expected 2, got %d", len(keys))
	}
}

// ==========================================
// Map ArgsCount
// ==========================================

func Test_Map_ArgsCount(t *testing.T) {
	m := args.Map{"a": 1, "b": 2}
	if m.ArgsCount() != 2 {
		t.Errorf("expected 2, got %d", m.ArgsCount())
	}
}

func Test_Map_ArgsCount_WithExpected(t *testing.T) {
	m := args.Map{"a": 1, "expected": "val"}
	c := m.ArgsCount()
	if c != 1 {
		t.Errorf("expected 1 (excluding expected), got %d", c)
	}
}

// ==========================================
// Map GetFirstOfNames
// ==========================================

func Test_Map_GetFirstOfNames(t *testing.T) {
	m := args.Map{"name": "val"}
	r := m.GetFirstOfNames("missing", "name")
	if r != "val" {
		t.Error("should return first found")
	}
}

func Test_Map_GetFirstOfNames_Empty(t *testing.T) {
	m := args.Map{"name": "val"}
	r := m.GetFirstOfNames()
	if r != nil {
		t.Error("empty names should return nil")
	}
}
