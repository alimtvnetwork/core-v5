package coregenerictests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/coregeneric"
)

// ==========================================
// Constructors
// ==========================================

func Test_GenericHashset_EmptyHashset(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	if !hs.IsEmpty() {
		t.Error("EmptyHashset should be empty")
	}
	if hs.Length() != 0 {
		t.Errorf("EmptyHashset length expected 0, got %d", hs.Length())
	}
}

func Test_GenericHashset_NewHashset_PreallocatesCapacity(t *testing.T) {
	hs := coregeneric.NewHashset[string](100)
	if !hs.IsEmpty() {
		t.Error("NewHashset with capacity should still be empty")
	}
}

func Test_GenericHashset_HashsetFrom_Slice(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 2, 1})
	if hs.Length() != 3 {
		t.Errorf("HashsetFrom with duplicates: expected 3, got %d", hs.Length())
	}
	if !hs.Has(1) || !hs.Has(2) || !hs.Has(3) {
		t.Error("HashsetFrom missing expected items")
	}
}

func Test_GenericHashset_HashsetFrom_EmptySlice(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{})
	if !hs.IsEmpty() {
		t.Error("HashsetFrom empty slice should be empty")
	}
}

func Test_GenericHashset_HashsetFromMap(t *testing.T) {
	m := map[string]bool{"a": true, "b": true}
	hs := coregeneric.HashsetFromMap(m)
	if hs.Length() != 2 {
		t.Errorf("HashsetFromMap expected 2, got %d", hs.Length())
	}
}

// ==========================================
// Add / AddBool
// ==========================================

func Test_GenericHashset_Add_ReturnsSelf(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	result := hs.Add(1)
	if result != hs {
		t.Error("Add should return same pointer")
	}
	if hs.Length() != 1 {
		t.Errorf("After Add: expected 1, got %d", hs.Length())
	}
}

func Test_GenericHashset_Add_Duplicate_NoIncrease(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	hs.Add(5).Add(5).Add(5)
	if hs.Length() != 1 {
		t.Errorf("Add duplicate: expected 1, got %d", hs.Length())
	}
}

func Test_GenericHashset_AddBool_FirstAdd_ReturnsFalse(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	existed := hs.AddBool("x")
	if existed {
		t.Error("AddBool first add should return false (did not exist)")
	}
}

func Test_GenericHashset_AddBool_SecondAdd_ReturnsTrue(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddBool("x")
	existed := hs.AddBool("x")
	if !existed {
		t.Error("AddBool second add should return true (already existed)")
	}
	if hs.Length() != 1 {
		t.Errorf("AddBool duplicate: expected 1, got %d", hs.Length())
	}
}

func Test_GenericHashset_AddBool_RepeatedAdds(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	r1 := hs.AddBool(42)
	r2 := hs.AddBool(42)
	r3 := hs.AddBool(42)
	r4 := hs.AddBool(42)
	if r1 {
		t.Error("First AddBool should return false")
	}
	if !r2 || !r3 || !r4 {
		t.Error("Subsequent AddBool calls should return true")
	}
	if hs.Length() != 1 {
		t.Errorf("Expected 1, got %d", hs.Length())
	}
}

// ==========================================
// Adds / AddSlice
// ==========================================

func Test_GenericHashset_Adds_Variadic(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	hs.Adds(1, 2, 3, 4, 5)
	if hs.Length() != 5 {
		t.Errorf("Adds: expected 5, got %d", hs.Length())
	}
}

func Test_GenericHashset_AddSlice(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSlice([]string{"a", "b", "c", "a"})
	if hs.Length() != 3 {
		t.Errorf("AddSlice with dup: expected 3, got %d", hs.Length())
	}
}

// ==========================================
// AddIf / AddIfMany
// ==========================================

func Test_GenericHashset_AddIf_True(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(true, 10)
	if !hs.Has(10) {
		t.Error("AddIf(true) should add item")
	}
}

func Test_GenericHashset_AddIf_False(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	hs.AddIf(false, 10)
	if hs.Has(10) {
		t.Error("AddIf(false) should not add item")
	}
}

func Test_GenericHashset_AddIfMany_True(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIfMany(true, "a", "b")
	if hs.Length() != 2 {
		t.Errorf("AddIfMany(true): expected 2, got %d", hs.Length())
	}
}

func Test_GenericHashset_AddIfMany_False(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIfMany(false, "a", "b")
	if hs.Length() != 0 {
		t.Errorf("AddIfMany(false): expected 0, got %d", hs.Length())
	}
}

// ==========================================
// AddHashsetItems / AddItemsMap
// ==========================================

func Test_GenericHashset_AddHashsetItems_Merge(t *testing.T) {
	hs1 := coregeneric.HashsetFrom([]int{1, 2})
	hs2 := coregeneric.HashsetFrom([]int{2, 3})
	hs1.AddHashsetItems(hs2)
	if hs1.Length() != 3 {
		t.Errorf("Merge: expected 3, got %d", hs1.Length())
	}
}

func Test_GenericHashset_AddHashsetItems_NilOther(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1})
	result := hs.AddHashsetItems(nil)
	if result != hs || hs.Length() != 1 {
		t.Error("AddHashsetItems(nil) should be no-op")
	}
}

func Test_GenericHashset_AddItemsMap_OnlyTrueValues(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{"yes": true, "no": false, "also": true})
	if hs.Length() != 2 {
		t.Errorf("AddItemsMap: expected 2 (only true), got %d", hs.Length())
	}
	if hs.Has("no") {
		t.Error("AddItemsMap should skip false-valued entries")
	}
}

// ==========================================
// Has / Contains / HasAll / HasAny
// ==========================================

func Test_GenericHashset_Has_Existing(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{10, 20, 30})
	if !hs.Has(20) {
		t.Error("Has should find existing item")
	}
}

func Test_GenericHashset_Has_Missing(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{10, 20, 30})
	if hs.Has(99) {
		t.Error("Has should not find missing item")
	}
}

func Test_GenericHashset_Contains_AliasForHas(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"x"})
	if hs.Contains("x") != hs.Has("x") {
		t.Error("Contains should match Has")
	}
}

func Test_GenericHashset_HasAll_AllPresent(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3, 4, 5})
	if !hs.HasAll(1, 3, 5) {
		t.Error("HasAll should return true when all present")
	}
}

func Test_GenericHashset_HasAll_OneMissing(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	if hs.HasAll(1, 2, 99) {
		t.Error("HasAll should return false when one missing")
	}
}

func Test_GenericHashset_HasAll_EmptyArgs(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1})
	if !hs.HasAll() {
		t.Error("HasAll with no args should return true")
	}
}

func Test_GenericHashset_HasAny_OnePresent(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	if !hs.HasAny(99, 2, 88) {
		t.Error("HasAny should return true when at least one present")
	}
}

func Test_GenericHashset_HasAny_NonePresent(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	if hs.HasAny(10, 20, 30) {
		t.Error("HasAny should return false when none present")
	}
}

func Test_GenericHashset_HasAny_EmptyArgs(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1})
	if hs.HasAny() {
		t.Error("HasAny with no args should return false")
	}
}

// ==========================================
// Remove
// ==========================================

func Test_GenericHashset_Remove_Existing(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(2)
	if !existed {
		t.Error("Remove should return true for existing item")
	}
	if hs.Length() != 2 {
		t.Errorf("After remove: expected 2, got %d", hs.Length())
	}
	if hs.Has(2) {
		t.Error("Removed item should not be found")
	}
}

func Test_GenericHashset_Remove_Missing(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	existed := hs.Remove(99)
	if existed {
		t.Error("Remove should return false for missing item")
	}
	if hs.Length() != 3 {
		t.Errorf("Length should be unchanged: expected 3, got %d", hs.Length())
	}
}

func Test_GenericHashset_Remove_UntilEmpty(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.Remove(1)
	hs.Remove(2)
	if !hs.IsEmpty() {
		t.Error("Should be empty after removing all items")
	}
}

// ==========================================
// IsEmpty / HasItems / Length
// ==========================================

func Test_GenericHashset_IsEmpty_NilReceiver(t *testing.T) {
	var hs *coregeneric.Hashset[int]
	if !hs.IsEmpty() {
		t.Error("nil Hashset should be empty")
	}
}

func Test_GenericHashset_HasItems_NonEmpty(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1})
	if !hs.HasItems() {
		t.Error("HasItems should return true for non-empty set")
	}
}

func Test_GenericHashset_HasItems_Empty(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	if hs.HasItems() {
		t.Error("HasItems should return false for empty set")
	}
}

func Test_GenericHashset_Length_NilReceiver(t *testing.T) {
	var hs *coregeneric.Hashset[int]
	if hs.Length() != 0 {
		t.Errorf("nil Hashset length expected 0, got %d", hs.Length())
	}
}

// ==========================================
// List / ListPtr / Map
// ==========================================

func Test_GenericHashset_List_ReturnsAllItems(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{10, 20, 30})
	list := hs.List()
	if len(list) != 3 {
		t.Errorf("List length: expected 3, got %d", len(list))
	}
}

func Test_GenericHashset_List_EmptySet(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	list := hs.List()
	if len(list) != 0 {
		t.Errorf("List of empty set should be empty, got %d", len(list))
	}
}

func Test_GenericHashset_ListPtr_NotNil(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	ptr := hs.ListPtr()
	if ptr == nil {
		t.Error("ListPtr should not return nil")
	}
	if len(*ptr) != 1 {
		t.Errorf("ListPtr: expected 1, got %d", len(*ptr))
	}
}

func Test_GenericHashset_Map_ReturnsUnderlyingMap(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2})
	m := hs.Map()
	if len(m) != 2 {
		t.Errorf("Map: expected 2, got %d", len(m))
	}
	if !m[1] || !m[2] {
		t.Error("Map should contain all items as true")
	}
}

// ==========================================
// Resize
// ==========================================

func Test_GenericHashset_Resize_LargerCapacity(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	result := hs.Resize(100)
	if result != hs {
		t.Error("Resize should return same pointer")
	}
	if hs.Length() != 3 {
		t.Errorf("Resize should preserve items: expected 3, got %d", hs.Length())
	}
	if !hs.Has(1) || !hs.Has(2) || !hs.Has(3) {
		t.Error("Resize should preserve all items")
	}
}

func Test_GenericHashset_Resize_SmallerThanLength_NoOp(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2, 3})
	result := hs.Resize(1)
	if result != hs {
		t.Error("Resize smaller should return same pointer")
	}
	if hs.Length() != 3 {
		t.Errorf("Resize smaller should not lose items: expected 3, got %d", hs.Length())
	}
}

func Test_GenericHashset_Resize_EqualToLength(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2})
	hs.Resize(2)
	if hs.Length() != 2 {
		t.Errorf("Resize equal: expected 2, got %d", hs.Length())
	}
}

// ==========================================
// IsEquals
// ==========================================

func Test_GenericHashset_IsEquals_BothNil(t *testing.T) {
	var a, b *coregeneric.Hashset[int]
	if !a.IsEquals(b) {
		t.Error("Two nil hashsets should be equal")
	}
}

func Test_GenericHashset_IsEquals_OneNil(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	var nilHs *coregeneric.Hashset[int]
	if hs.IsEquals(nilHs) {
		t.Error("Non-nil vs nil should not be equal")
	}
}

func Test_GenericHashset_IsEquals_SamePointer(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1, 2})
	if !hs.IsEquals(hs) {
		t.Error("Same pointer should be equal")
	}
}

func Test_GenericHashset_IsEquals_SameContent(t *testing.T) {
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{3, 2, 1})
	if !a.IsEquals(b) {
		t.Error("Same content should be equal")
	}
}

func Test_GenericHashset_IsEquals_DifferentContent(t *testing.T) {
	a := coregeneric.HashsetFrom([]int{1, 2, 3})
	b := coregeneric.HashsetFrom([]int{1, 2, 4})
	if a.IsEquals(b) {
		t.Error("Different content should not be equal")
	}
}

func Test_GenericHashset_IsEquals_DifferentLength(t *testing.T) {
	a := coregeneric.HashsetFrom([]int{1, 2})
	b := coregeneric.HashsetFrom([]int{1, 2, 3})
	if a.IsEquals(b) {
		t.Error("Different length should not be equal")
	}
}

func Test_GenericHashset_IsEquals_BothEmpty(t *testing.T) {
	a := coregeneric.EmptyHashset[string]()
	b := coregeneric.EmptyHashset[string]()
	if !a.IsEquals(b) {
		t.Error("Two empty hashsets should be equal")
	}
}

// ==========================================
// Collection conversion
// ==========================================

func Test_GenericHashset_Collection_Conversion(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{5, 10, 15})
	col := hs.Collection()
	if col.Length() != 3 {
		t.Errorf("Collection length: expected 3, got %d", col.Length())
	}
}

func Test_GenericHashset_Collection_EmptySet(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	col := hs.Collection()
	if col.Length() != 0 {
		t.Errorf("Collection of empty set: expected 0, got %d", col.Length())
	}
}

// ==========================================
// String
// ==========================================

func Test_GenericHashset_String_NotEmpty(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{1})
	s := hs.String()
	if s == "" {
		t.Error("String should not be empty for non-empty set")
	}
}

// ==========================================
// Nil receiver guards
// ==========================================

func Test_GenericHashset_NilReceiver_IsEmpty(t *testing.T) {
	var hs *coregeneric.Hashset[string]
	if !hs.IsEmpty() {
		t.Error("nil.IsEmpty() should return true")
	}
}

func Test_GenericHashset_NilReceiver_Length(t *testing.T) {
	var hs *coregeneric.Hashset[string]
	if hs.Length() != 0 {
		t.Error("nil.Length() should return 0")
	}
}

func Test_GenericHashset_NilReceiver_HasItems(t *testing.T) {
	var hs *coregeneric.Hashset[string]
	if hs.HasItems() {
		t.Error("nil.HasItems() should return false")
	}
}

// ==========================================
// Creator pattern (New.Hashset.X)
// ==========================================

func Test_GenericHashset_Creator_String_Items(t *testing.T) {
	hs := coregeneric.New.Hashset.String.Items("a", "b", "c")
	if hs.Length() != 3 {
		t.Errorf("Creator String.Items: expected 3, got %d", hs.Length())
	}
}

func Test_GenericHashset_Creator_Int_From(t *testing.T) {
	hs := coregeneric.New.Hashset.Int.From([]int{1, 2, 3, 1})
	if hs.Length() != 3 {
		t.Errorf("Creator Int.From: expected 3, got %d", hs.Length())
	}
}

func Test_GenericHashset_Creator_Empty(t *testing.T) {
	hs := coregeneric.New.Hashset.Float64.Empty()
	if !hs.IsEmpty() {
		t.Error("Creator Empty should produce empty set")
	}
}

func Test_GenericHashset_Creator_Cap(t *testing.T) {
	hs := coregeneric.New.Hashset.Bool.Cap(10)
	if !hs.IsEmpty() {
		t.Error("Creator Cap should produce empty set with capacity")
	}
}

func Test_GenericHashset_Creator_UsingMap(t *testing.T) {
	m := map[uint]bool{1: true, 2: true}
	hs := coregeneric.New.Hashset.Uint.UsingMap(m)
	if hs.Length() != 2 {
		t.Errorf("Creator UsingMap: expected 2, got %d", hs.Length())
	}
}
