package corestr

import (
	"sync"
	"testing"
)

// ── DiffRaw, Diff ──

func TestHashmap_DiffRaw(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	diff := hm.DiffRaw(map[string]string{"a": "2"})
	if len(diff) == 0 { t.Fatal("expected diff") }
}

func TestHashmap_Diff(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "2"})
	diff := hm1.Diff(hm2)
	if diff.IsEmpty() { t.Fatal("expected non-empty diff") }
}

// ── HasAllCollectionItems ──

func TestHashmap_HasAllCollectionItems(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	c := New.Collection.Strings([]string{"a", "b"})
	if !hm.HasAllCollectionItems(c) { t.Fatal("expected true") }
}

func TestHashmap_HasAllCollectionItems_Nil(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	if hm.HasAllCollectionItems(nil) { t.Fatal("expected false") }
}

func TestHashmap_HasAllCollectionItems_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	if hm.HasAllCollectionItems(New.Collection.Cap(0)) { t.Fatal("expected false") }
}

// ── HasAll ──

func TestHashmap_HasAll(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	if !hm.HasAll("a", "b") { t.Fatal("expected true") }
	if hm.HasAll("a", "z") { t.Fatal("expected false") }
}

// ── HasAnyItem ──

func TestHashmap_HasAnyItem(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if !hm.HasAnyItem() { t.Fatal("expected true") }
	empty := New.Hashmap.Cap(0)
	if empty.HasAnyItem() { t.Fatal("expected false") }
}

// ── HasAny ──

func TestHashmap_HasAny(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if !hm.HasAny("a", "z") { t.Fatal("expected true") }
	if hm.HasAny("x", "y") { t.Fatal("expected false") }
}

// ── HasWithLock ──

func TestHashmap_HasWithLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if !hm.HasWithLock("a") { t.Fatal("expected true") }
	if hm.HasWithLock("z") { t.Fatal("expected false") }
}

// ── GetKeysFilteredItems ──

func TestHashmap_GetKeysFilteredItems(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "abc", Value: "1"},
		KeyValuePair{Key: "def", Value: "2"},
	)
	result := hm.GetKeysFilteredItems(func(str string, i int) (string, bool, bool) {
		return str, true, false
	})
	if len(result) != 2 { t.Fatal("expected 2") }
}

func TestHashmap_GetKeysFilteredItems_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	result := hm.GetKeysFilteredItems(func(str string, i int) (string, bool, bool) {
		return str, true, false
	})
	if len(result) != 0 { t.Fatal("expected 0") }
}

func TestHashmap_GetKeysFilteredItems_Break(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	result := hm.GetKeysFilteredItems(func(str string, i int) (string, bool, bool) {
		return str, true, true // break after first
	})
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetKeysFilteredItems_Skip(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	result := hm.GetKeysFilteredItems(func(str string, i int) (string, bool, bool) {
		return str, false, false
	})
	if len(result) != 0 { t.Fatal("expected 0") }
}

// ── GetKeysFilteredCollection ──

func TestHashmap_GetKeysFilteredCollection(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := hm.GetKeysFilteredCollection(func(str string, i int) (string, bool, bool) {
		return str, true, false
	})
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashmap_GetKeysFilteredCollection_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	c := hm.GetKeysFilteredCollection(func(str string, i int) (string, bool, bool) {
		return str, true, false
	})
	if !c.IsEmpty() { t.Fatal("expected empty") }
}

func TestHashmap_GetKeysFilteredCollection_Break(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	c := hm.GetKeysFilteredCollection(func(str string, i int) (string, bool, bool) {
		return str, true, true
	})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

// ── Items, SafeItems ──

func TestHashmap_Items(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if len(hm.Items()) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_SafeItems(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if len(hm.SafeItems()) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_SafeItems_Nil(t *testing.T) {
	var hm *Hashmap
	if hm.SafeItems() != nil { t.Fatal("expected nil") }
}

// ── ItemsCopyLock ──

func TestHashmap_ItemsCopyLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	copied := hm.ItemsCopyLock()
	if len(*copied) != 1 { t.Fatal("expected 1") }
}

// ── ValuesCollection, ValuesHashset ──

func TestHashmap_ValuesCollection(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := hm.ValuesCollection()
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashmap_ValuesHashset(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hs := hm.ValuesHashset()
	if hs.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashmap_ValuesCollectionLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := hm.ValuesCollectionLock()
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashmap_ValuesHashsetLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hs := hm.ValuesHashsetLock()
	if hs.IsEmpty() { t.Fatal("expected non-empty") }
}

// ── ValuesList ──

func TestHashmap_ValuesList(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	vl := hm.ValuesList()
	if len(vl) != 1 { t.Fatal("expected 1") }
}

// ── KeysValuesCollection ──

func TestHashmap_KeysValuesCollection(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	keys, vals := hm.KeysValuesCollection()
	if keys.IsEmpty() || vals.IsEmpty() { t.Fatal("expected non-empty") }
}

// ── KeysValuesList ──

func TestHashmap_KeysValuesList(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	keys, vals := hm.KeysValuesList()
	if len(keys) != 1 || len(vals) != 1 { t.Fatal("expected 1") }
}

// ── KeysValuePairs ──

func TestHashmap_KeysValuePairs(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	pairs := hm.KeysValuePairs()
	if len(pairs) != 1 { t.Fatal("expected 1") }
}

// ── KeysValuePairsCollection ──

func TestHashmap_KeysValuePairsCollection(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := hm.KeysValuePairsCollection()
	if c == nil { t.Fatal("expected non-nil") }
}

// ── KeysValuesListLock ──

func TestHashmap_KeysValuesListLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	keys, vals := hm.KeysValuesListLock()
	if len(keys) != 1 || len(vals) != 1 { t.Fatal("expected 1") }
}

// ── AllKeys, Keys, KeysCollection, KeysLock ──

func TestHashmap_AllKeys(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if len(hm.AllKeys()) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AllKeys_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	if len(hm.AllKeys()) != 0 { t.Fatal("expected 0") }
}

func TestHashmap_Keys(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if len(hm.Keys()) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysCollection(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	c := hm.KeysCollection()
	if c.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashmap_KeysLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if len(hm.KeysLock()) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_KeysLock_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	if len(hm.KeysLock()) != 0 { t.Fatal("expected 0") }
}

// ── ValuesListCopyLock ──

func TestHashmap_ValuesListCopyLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	vl := hm.ValuesListCopyLock()
	if len(vl) != 1 { t.Fatal("expected 1") }
}

// ── KeysToLower, ValuesToLower ──

func TestHashmap_KeysToLower(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "ABC", Value: "1"})
	lower := hm.KeysToLower()
	if !lower.Has("abc") { t.Fatal("expected lowercase key") }
}

func TestHashmap_ValuesToLower(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "ABC", Value: "1"})
	lower := hm.ValuesToLower()
	if !lower.Has("abc") { t.Fatal("expected lowercase key") }
}

// ── Length, LengthLock ──

func TestHashmap_Length(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_Length_Nil(t *testing.T) {
	var hm *Hashmap
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_LengthLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if hm.LengthLock() != 1 { t.Fatal("expected 1") }
}

// ── IsEqual, IsEqualPtr, IsEqualPtrLock ──

func TestHashmap_IsEqual(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := *New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if !hm1.IsEqual(hm2) { t.Fatal("expected true") }
}

func TestHashmap_IsEqualPtr(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if !hm1.IsEqualPtr(hm2) { t.Fatal("expected true") }
}

func TestHashmap_IsEqualPtr_BothNil(t *testing.T) {
	var hm1, hm2 *Hashmap
	if !hm1.IsEqualPtr(hm2) { t.Fatal("expected true") }
}

func TestHashmap_IsEqualPtr_OneNil(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	if hm.IsEqualPtr(nil) { t.Fatal("expected false") }
}

func TestHashmap_IsEqualPtr_SamePtr(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if !hm.IsEqualPtr(hm) { t.Fatal("expected true") }
}

func TestHashmap_IsEqualPtr_BothEmpty(t *testing.T) {
	hm1 := New.Hashmap.Cap(0)
	hm2 := New.Hashmap.Cap(0)
	if !hm1.IsEqualPtr(hm2) { t.Fatal("expected true") }
}

func TestHashmap_IsEqualPtr_DiffLength(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"}, KeyValuePair{Key: "b", Value: "2"})
	if hm1.IsEqualPtr(hm2) { t.Fatal("expected false") }
}

func TestHashmap_IsEqualPtr_DiffValue(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "2"})
	if hm1.IsEqualPtr(hm2) { t.Fatal("expected false") }
}

func TestHashmap_IsEqualPtrLock(t *testing.T) {
	hm1 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm2 := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if !hm1.IsEqualPtrLock(hm2) { t.Fatal("expected true") }
}

// ── Remove, RemoveWithLock ──

func TestHashmap_Remove(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm.Remove("a")
	if hm.Has("a") { t.Fatal("expected removed") }
}

func TestHashmap_RemoveWithLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm.RemoveWithLock("a")
	if hm.Has("a") { t.Fatal("expected removed") }
}

// ── String, StringLock ──

func TestHashmap_String(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if hm.String() == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_String_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	s := hm.String()
	if s == "" { t.Fatal("expected non-empty (NoElements)") }
}

func TestHashmap_StringLock(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if hm.StringLock() == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_StringLock_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	s := hm.StringLock()
	if s == "" { t.Fatal("expected non-empty") }
}

// ── GetValuesExceptKeysInHashset ──

func TestHashmap_GetValuesExceptKeysInHashset(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	hs := New.Hashset.Strings([]string{"a"})
	result := hm.GetValuesExceptKeysInHashset(hs)
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetValuesExceptKeysInHashset_Nil(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	result := hm.GetValuesExceptKeysInHashset(nil)
	if len(result) != 1 { t.Fatal("expected 1") }
}

// ── GetValuesKeysExcept ──

func TestHashmap_GetValuesKeysExcept(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	result := hm.GetValuesKeysExcept([]string{"a"})
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetValuesKeysExcept_Nil(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	result := hm.GetValuesKeysExcept(nil)
	if len(result) != 1 { t.Fatal("expected 1") }
}

// ── GetAllExceptCollection ──

func TestHashmap_GetAllExceptCollection(t *testing.T) {
	hm := New.Hashmap.KeyValues(
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	c := New.Collection.Strings([]string{"a"})
	result := hm.GetAllExceptCollection(c)
	if len(result) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_GetAllExceptCollection_Nil(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	result := hm.GetAllExceptCollection(nil)
	if len(result) != 1 { t.Fatal("expected 1") }
}

// ── Join, JoinKeys ──

func TestHashmap_Join(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	r := hm.Join(",")
	if r == "" { t.Fatal("expected non-empty") }
}

func TestHashmap_JoinKeys(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	r := hm.JoinKeys(",")
	if r == "" { t.Fatal("expected non-empty") }
}

// ── JSON methods ──

func TestHashmap_JsonModel(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	if len(hm.JsonModel()) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_JsonModelAny(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	_ = hm.JsonModelAny()
}

func TestHashmap_MarshalJSON(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	b, err := hm.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashmap_UnmarshalJSON(t *testing.T) {
	hm := &Hashmap{}
	err := hm.UnmarshalJSON([]byte(`{"a":"1"}`))
	if err != nil || hm.Length() != 1 { t.Fatal("unexpected") }
}

func TestHashmap_UnmarshalJSON_Error(t *testing.T) {
	hm := &Hashmap{}
	err := hm.UnmarshalJSON([]byte(`invalid`))
	if err == nil { t.Fatal("expected error") }
}

func TestHashmap_Json(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	r := hm.Json()
	if r.HasError() { t.Fatal("unexpected error") }
}

func TestHashmap_JsonPtr(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	r := hm.JsonPtr()
	if r.HasError() { t.Fatal("unexpected error") }
}

// ── ParseInjectUsingJson ──

func TestHashmap_ParseInjectUsingJson(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	result, err := hm2.ParseInjectUsingJson(&jr)
	if err != nil || result.Length() != 1 { t.Fatal("unexpected") }
}

func TestHashmap_ParseInjectUsingJsonMust(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	result := hm2.ParseInjectUsingJsonMust(&jr)
	if result.Length() != 1 { t.Fatal("unexpected") }
}

func TestHashmap_JsonParseSelfInject(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	jr := hm.Json()
	hm2 := &Hashmap{items: map[string]string{}}
	err := hm2.JsonParseSelfInject(&jr)
	if err != nil { t.Fatal("unexpected") }
}

// ── ToError, ToDefaultError ──

func TestHashmap_ToError(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	err := hm.ToError(",")
	if err == nil { t.Fatal("expected error") }
}

func TestHashmap_ToDefaultError(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	err := hm.ToDefaultError()
	if err == nil { t.Fatal("expected error") }
}

// ── KeyValStringLines ──

func TestHashmap_KeyValStringLines(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	lines := hm.KeyValStringLines()
	if len(lines) != 1 { t.Fatal("expected 1") }
}

// ── Clear, Dispose ──

func TestHashmap_Clear(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm.Clear()
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_Clear_Nil(t *testing.T) {
	var hm *Hashmap
	r := hm.Clear()
	if r != nil { t.Fatal("expected nil") }
}

func TestHashmap_Dispose(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	hm.Dispose()
	if hm.items != nil { t.Fatal("expected nil") }
}

func TestHashmap_Dispose_Nil(t *testing.T) {
	var hm *Hashmap
	hm.Dispose() // should not panic
}

// ── ToStringsUsingCompiler ──

func TestHashmap_ToStringsUsingCompiler(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	lines := hm.ToStringsUsingCompiler(func(k, v string) string { return k + "=" + v })
	if len(lines) != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ToStringsUsingCompiler_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	lines := hm.ToStringsUsingCompiler(func(k, v string) string { return k })
	if len(lines) != 0 { t.Fatal("expected 0") }
}

// ── AsJsoner, AsJsonContractsBinder, AsJsonParseSelfInjector, AsJsonMarshaller ──

func TestHashmap_AsJsoner(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	_ = hm.AsJsoner()
}

func TestHashmap_AsJsonContractsBinder(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	_ = hm.AsJsonContractsBinder()
}

func TestHashmap_AsJsonParseSelfInjector(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	_ = hm.AsJsonParseSelfInjector()
}

func TestHashmap_AsJsonMarshaller(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	_ = hm.AsJsonMarshaller()
}

// ── ClonePtr, Clone ──

func TestHashmap_ClonePtr(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	cloned := hm.ClonePtr()
	if cloned.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_ClonePtr_Nil(t *testing.T) {
	var hm *Hashmap
	cloned := hm.ClonePtr()
	if cloned != nil { t.Fatal("expected nil") }
}

func TestHashmap_Clone(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	cloned := hm.Clone()
	if cloned.Length() != 1 { t.Fatal("expected 1") }
	// verify independence
	hm.Set("b", "2")
	if cloned.Has("b") { t.Fatal("expected independent clone") }
}

func TestHashmap_Clone_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(0)
	cloned := hm.Clone()
	if cloned.Length() != 0 { t.Fatal("expected 0") }
}

// ── Get, GetValue ──

func TestHashmap_Get(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	v, found := hm.Get("a")
	if !found || v != "1" { t.Fatal("unexpected") }
}

func TestHashmap_GetValue(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	v, found := hm.GetValue("a")
	if !found || v != "1" { t.Fatal("unexpected") }
}

// ── Serialize, Deserialize ──

func TestHashmap_Serialize(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	b, err := hm.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashmap_Deserialize(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	var target map[string]string
	err := hm.Deserialize(&target)
	if err != nil || len(target) != 1 { t.Fatal("unexpected") }
}

// ── AddsOrUpdatesAnyUsingFilter ──

func TestHashmap_AddsOrUpdatesAnyUsingFilter(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesAnyUsingFilter(
		func(pair KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, false
		},
		KeyAnyValuePair{Key: "a", Value: 1},
	)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddsOrUpdatesAnyUsingFilter_Nil(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesAnyUsingFilter(nil)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_AddsOrUpdatesAnyUsingFilter_Break(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesAnyUsingFilter(
		func(pair KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, true
		},
		KeyAnyValuePair{Key: "a", Value: 1},
		KeyAnyValuePair{Key: "b", Value: 2},
	)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddsOrUpdatesAnyUsingFilter_Skip(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesAnyUsingFilter(
		func(pair KeyAnyValuePair) (string, bool, bool) {
			return "", false, false
		},
		KeyAnyValuePair{Key: "a", Value: 1},
	)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddsOrUpdatesAnyUsingFilterLock ──

func TestHashmap_AddsOrUpdatesAnyUsingFilterLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesAnyUsingFilterLock(
		func(pair KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, false
		},
		KeyAnyValuePair{Key: "a", Value: 1},
	)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddsOrUpdatesAnyUsingFilterLock_Nil(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesAnyUsingFilterLock(nil)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_AddsOrUpdatesAnyUsingFilterLock_Break(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesAnyUsingFilterLock(
		func(pair KeyAnyValuePair) (string, bool, bool) {
			return pair.ValueString(), true, true
		},
		KeyAnyValuePair{Key: "a", Value: 1},
		KeyAnyValuePair{Key: "b", Value: 2},
	)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

// ── AddsOrUpdatesUsingFilter ──

func TestHashmap_AddsOrUpdatesUsingFilter(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesUsingFilter(
		func(pair KeyValuePair) (string, bool, bool) {
			return pair.Value, true, false
		},
		KeyValuePair{Key: "a", Value: "1"},
	)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddsOrUpdatesUsingFilter_Nil(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesUsingFilter(nil)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_AddsOrUpdatesUsingFilter_Break(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesUsingFilter(
		func(pair KeyValuePair) (string, bool, bool) {
			return pair.Value, true, true
		},
		KeyValuePair{Key: "a", Value: "1"},
		KeyValuePair{Key: "b", Value: "2"},
	)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmap_AddsOrUpdatesUsingFilter_Skip(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	hm.AddsOrUpdatesUsingFilter(
		func(pair KeyValuePair) (string, bool, bool) {
			return "", false, false
		},
		KeyValuePair{Key: "a", Value: "1"},
	)
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

// ── AddOrUpdateWithWgLock ──

func TestHashmap_AddOrUpdateWithWgLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hm.AddOrUpdateWithWgLock("k", "v", &wg)
	wg.Wait()
	if !hm.Has("k") { t.Fatal("expected found") }
}

// ── AddOrUpdateStringsPtrWgLock ──

func TestHashmap_AddOrUpdateStringsPtrWgLock(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hm.AddOrUpdateStringsPtrWgLock(&wg, []string{"a"}, []string{"1"})
	wg.Wait()
	if !hm.Has("a") { t.Fatal("expected found") }
}

func TestHashmap_AddOrUpdateStringsPtrWgLock_Empty(t *testing.T) {
	hm := New.Hashmap.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hm.AddOrUpdateStringsPtrWgLock(&wg, []string{}, []string{})
	// wg.Done not called for empty, so don't wait — method returns directly
	if hm.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmap_AddOrUpdateStringsPtrWgLock_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil { t.Fatal("expected panic") }
	}()
	hm := New.Hashmap.Cap(5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	hm.AddOrUpdateStringsPtrWgLock(&wg, []string{"a"}, []string{})
}

// ── HashmapDiff type ──

func TestHashmapDiff_Length(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	if d.Length() != 1 { t.Fatal("expected 1") }
}

func TestHashmapDiff_Length_Nil(t *testing.T) {
	var d *HashmapDiff
	if d.Length() != 0 { t.Fatal("expected 0") }
}

func TestHashmapDiff_IsEmpty(t *testing.T) {
	d := HashmapDiff(map[string]string{})
	if !d.IsEmpty() { t.Fatal("expected true") }
}

func TestHashmapDiff_HasAnyItem(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	if !d.HasAnyItem() { t.Fatal("expected true") }
}

func TestHashmapDiff_LastIndex(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	if d.LastIndex() != 0 { t.Fatal("expected 0") }
}

func TestHashmapDiff_Raw(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	r := d.Raw()
	if len(r) != 1 { t.Fatal("expected 1") }
}

func TestHashmapDiff_Raw_Nil(t *testing.T) {
	var d *HashmapDiff
	r := d.Raw()
	if len(r) != 0 { t.Fatal("expected 0") }
}

func TestHashmapDiff_MapAnyItems(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	r := d.MapAnyItems()
	if len(r) != 1 { t.Fatal("expected 1") }
}

func TestHashmapDiff_MapAnyItems_Nil(t *testing.T) {
	var d *HashmapDiff
	r := d.MapAnyItems()
	if len(r) != 0 { t.Fatal("expected 0") }
}

func TestHashmapDiff_AllKeysSorted(t *testing.T) {
	d := HashmapDiff(map[string]string{"b": "2", "a": "1"})
	keys := d.AllKeysSorted()
	if keys[0] != "a" { t.Fatal("expected sorted") }
}

func TestHashmapDiff_IsRawEqual(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	if !d.IsRawEqual(map[string]string{"a": "1"}) { t.Fatal("expected true") }
}

func TestHashmapDiff_HasAnyChanges(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	if !d.HasAnyChanges(map[string]string{"a": "2"}) { t.Fatal("expected true") }
}

func TestHashmapDiff_DiffRaw(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	diff := d.DiffRaw(map[string]string{"a": "2"})
	if len(diff) == 0 { t.Fatal("expected diff") }
}

func TestHashmapDiff_HashmapDiffUsingRaw(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	diff := d.HashmapDiffUsingRaw(map[string]string{"a": "2"})
	if diff.IsEmpty() { t.Fatal("expected non-empty") }
}

func TestHashmapDiff_HashmapDiffUsingRaw_NoDiff(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	diff := d.HashmapDiffUsingRaw(map[string]string{"a": "1"})
	if !diff.IsEmpty() { t.Fatal("expected empty") }
}

func TestHashmapDiff_DiffJsonMessage(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	msg := d.DiffJsonMessage(map[string]string{"a": "2"})
	_ = msg
}

func TestHashmapDiff_RawMapStringAnyDiff(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	r := d.RawMapStringAnyDiff()
	if len(r) != 1 { t.Fatal("expected 1") }
}

func TestHashmapDiff_ShouldDiffMessage(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	msg := d.ShouldDiffMessage("test", map[string]string{"a": "2"})
	_ = msg
}

func TestHashmapDiff_LogShouldDiffMessage(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	msg := d.LogShouldDiffMessage("test", map[string]string{"a": "2"})
	_ = msg
}

func TestHashmapDiff_ToStringsSliceOfDiffMap(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	sl := d.ToStringsSliceOfDiffMap(map[string]string{"a": "changed"})
	_ = sl
}

func TestHashmapDiff_Serialize(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	b, err := d.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestHashmapDiff_Deserialize(t *testing.T) {
	d := HashmapDiff(map[string]string{"a": "1"})
	var target map[string]string
	err := d.Deserialize(&target)
	if err != nil || len(target) != 1 { t.Fatal("unexpected") }
}
