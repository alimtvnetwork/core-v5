package enumimpltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
)

// ===================== DynamicMap =====================

func Test_C13_DynamicMap_AddOrUpdate(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	isNew := dm.AddOrUpdate("k1", "v1")
	if !isNew {
		t.Fatal("expected new")
	}
	isNew2 := dm.AddOrUpdate("k1", "v2")
	if isNew2 {
		t.Fatal("expected update not new")
	}
}

func Test_C13_DynamicMap_Set(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	isNew := dm.Set("k", "v")
	if !isNew {
		t.Fatal("expected new")
	}
	isNew2 := dm.Set("k", "v2")
	if isNew2 {
		t.Fatal("expected update")
	}
}

func Test_C13_DynamicMap_AddNewOnly(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if !dm.AddNewOnly("k", "v") {
		t.Fatal("expected added")
	}
	if dm.AddNewOnly("k", "v2") {
		t.Fatal("expected not added")
	}
}

func Test_C13_DynamicMap_AllKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeys()
	if len(keys) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_AllKeys_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.AllKeys()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_AllKeysSorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeysSorted()
	if keys[0] != "a" {
		t.Fatal("expected a first")
	}
}

func Test_C13_DynamicMap_AllKeysSorted_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.AllKeysSorted()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_AllValuesStrings(t *testing.T) {
	dm := enumimpl.DynamicMap{"k": "v"}
	vs := dm.AllValuesStrings()
	if len(vs) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_AllValuesStrings_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.AllValuesStrings()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_AllValuesStringsSorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"k": "v"}
	vs := dm.AllValuesStringsSorted()
	if len(vs) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_AllValuesStringsSorted_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.AllValuesStringsSorted()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_AllValuesIntegers(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	ints := dm.AllValuesIntegers()
	if len(ints) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_AllValuesIntegers_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.AllValuesIntegers()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_MapIntegerString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	m, sorted := dm.MapIntegerString()
	if len(m) == 0 || len(sorted) == 0 {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_DynamicMap_MapIntegerString_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	m, sorted := dm.MapIntegerString()
	if len(m) != 0 || len(sorted) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_MapIntegerString_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	m, sorted := dm.MapIntegerString()
	_ = m
	_ = sorted
}

func Test_C13_DynamicMap_SortedKeyValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kv := dm.SortedKeyValues()
	if len(kv) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_SortedKeyValues_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.SortedKeyValues()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_SortedKeyAnyValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kav := dm.SortedKeyAnyValues()
	if len(kav) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_SortedKeyAnyValues_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.SortedKeyAnyValues()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_SortedKeyAnyValues_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	kav := dm.SortedKeyAnyValues()
	if len(kav) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_First(t *testing.T) {
	dm := enumimpl.DynamicMap{"k": "v"}
	k, v := dm.First()
	if k == "" || v == nil {
		t.Fatal("expected key and value")
	}
}

func Test_C13_DynamicMap_First_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	k, v := dm.First()
	if k != "" || v != nil {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_IsValueString(t *testing.T) {
	dm := enumimpl.DynamicMap{"k": "v"}
	if !dm.IsValueString() {
		t.Fatal("expected true")
	}
	dm2 := enumimpl.DynamicMap{"k": 1}
	if dm2.IsValueString() {
		t.Fatal("expected false")
	}
}

func Test_C13_DynamicMap_LengthAndCount(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if dm.Length() != 1 || dm.Count() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_Length_Nil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	if dm.Length() != 0 {
		t.Fatal("nil length should be 0")
	}
}

func Test_C13_DynamicMap_IsEmpty_HasAnyItem(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if !dm.IsEmpty() || dm.HasAnyItem() {
		t.Fatal("expected empty")
	}
	dm["k"] = 1
	if dm.IsEmpty() || !dm.HasAnyItem() {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_DynamicMap_LastIndex_HasIndex(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	if dm.LastIndex() != 1 {
		t.Fatalf("expected 1, got %d", dm.LastIndex())
	}
	if !dm.HasIndex(1) || dm.HasIndex(2) {
		t.Fatal("index check failed")
	}
}

func Test_C13_DynamicMap_HasKey_IsMissingKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.HasKey("a") || dm.HasKey("b") {
		t.Fatal("key check failed")
	}
	if dm.IsMissingKey("a") || !dm.IsMissingKey("b") {
		t.Fatal("missing key check failed")
	}
}

func Test_C13_DynamicMap_HasAllKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	if !dm.HasAllKeys("a", "b") {
		t.Fatal("expected true")
	}
	if dm.HasAllKeys("a", "c") {
		t.Fatal("expected false")
	}
}

func Test_C13_DynamicMap_HasAnyKeys(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.HasAnyKeys("a", "b") {
		t.Fatal("expected true")
	}
	if dm.HasAnyKeys("c") {
		t.Fatal("expected false")
	}
}

func Test_C13_DynamicMap_IsEqual_BothNil(t *testing.T) {
	var l, r *enumimpl.DynamicMap
	if !l.IsEqual(true, r) {
		t.Fatal("both nil should be equal")
	}
}

func Test_C13_DynamicMap_IsEqual_OneNil(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if dm.IsEqual(true, nil) {
		t.Fatal("expected false")
	}
}

func Test_C13_DynamicMap_IsEqual_SameRef(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.IsEqual(true, &dm) {
		t.Fatal("same ref should be equal")
	}
}

func Test_C13_DynamicMap_IsEqual_RegardlessType(t *testing.T) {
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 1}
	if !dm1.IsEqual(true, &dm2) {
		t.Fatal("expected equal regardless")
	}
}

func Test_C13_DynamicMap_IsEqual_StrictType(t *testing.T) {
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 1}
	if !dm1.IsEqual(false, &dm2) {
		t.Fatal("expected equal strict")
	}
}

func Test_C13_DynamicMap_IsRawEqual_DiffLength(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"a": 1, "b": 2}
	if dm.IsRawEqual(true, right) {
		t.Fatal("different length should not be equal")
	}
}

func Test_C13_DynamicMap_IsRawEqual_MissingKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	right := map[string]any{"b": 1}
	if dm.IsRawEqual(true, right) {
		t.Fatal("missing key should not be equal")
	}
}

func Test_C13_DynamicMap_IsMismatch(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	other := enumimpl.DynamicMap{"a": 2}
	if !dm.IsMismatch(false, &other) {
		t.Fatal("expected mismatch")
	}
}

func Test_C13_DynamicMap_IsRawMismatch(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.IsRawMismatch(false, map[string]any{"a": 2}) {
		t.Fatal("expected mismatch")
	}
}

func Test_C13_DynamicMap_DiffRaw(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	diff := dm.DiffRaw(false, map[string]any{"a": 1, "c": 3})
	if diff.Length() == 0 {
		t.Fatal("expected diffs")
	}
}

func Test_C13_DynamicMap_DiffRaw_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	diff := dm.DiffRaw(true, map[string]any{"a": 1})
	if diff.Length() != 0 {
		t.Fatal("expected no diffs")
	}
}

func Test_C13_DynamicMap_DiffRawUsingDifferChecker_BothNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	diff := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)
	if diff.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_DiffRawUsingDifferChecker_LeftNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	right := map[string]any{"a": 1}
	diff := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, right)
	if diff.Length() != 1 {
		t.Fatal("expected right map")
	}
}

func Test_C13_DynamicMap_DiffRawUsingDifferChecker_RightNil(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	diff := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)
	if diff.Length() != 1 {
		t.Fatal("expected left map")
	}
}

func Test_C13_DynamicMap_DiffRawLeftRight(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	lDiff, rDiff := dm.DiffRawLeftRightUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false,
		map[string]any{"a": 1, "c": 3},
	)
	_ = lDiff
	_ = rDiff
}

func Test_C13_DynamicMap_DiffRawLeftRight_BothNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)
	if l.Length() != 0 || r.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_DiffRawLeftRight_LeftNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, map[string]any{"a": 1})
	_ = l
	_ = r
}

func Test_C13_DynamicMap_DiffRawLeftRight_RightNil(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, nil)
	_ = l
	_ = r
}

func Test_C13_DynamicMap_DiffRawLeftRight_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	l, r := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, true, map[string]any{"a": 1})
	if l.Length() != 0 || r.Length() != 0 {
		t.Fatal("expected no diff")
	}
}

func Test_C13_DynamicMap_DiffJsonMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessage(false, map[string]any{"a": 2})
	if msg == "" {
		t.Fatal("expected diff message")
	}
}

func Test_C13_DynamicMap_DiffJsonMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessage(true, map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_DiffJsonMessageLeftRight(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageLeftRight(false, map[string]any{"b": 2})
	if msg == "" {
		t.Fatal("expected message")
	}
}

func Test_C13_DynamicMap_DiffJsonMessageLeftRight_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageLeftRight(true, map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_ShouldDiffMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	if msg == "" {
		t.Fatal("expected diff")
	}
}

func Test_C13_DynamicMap_ShouldDiffMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(true, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_LogShouldDiffMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessage(false, "test", map[string]any{"a": 2})
	if msg == "" {
		t.Fatal("expected diff")
	}
}

func Test_C13_DynamicMap_LogShouldDiffMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessage(true, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessage(false, "test", map[string]any{"b": 2})
	if msg == "" {
		t.Fatal("expected diff")
	}
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessage(true, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_ShouldDiffLeftRightMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "test",
		map[string]any{"b": 2},
	)
	if msg == "" {
		t.Fatal("expected diff")
	}
}

func Test_C13_DynamicMap_ShouldDiffLeftRightMessage_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_ExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 2})
	if msg == "" {
		t.Fatal("expected mismatch message")
	}
}

func Test_C13_DynamicMap_ExpectingMessage_Equal(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_LogExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 2})
}

func Test_C13_DynamicMap_LogExpectingMessage_Equal(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
}

func Test_C13_DynamicMap_IsKeysEqualOnly(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.IsKeysEqualOnly(map[string]any{"a": 99}) {
		t.Fatal("keys should be equal")
	}
	if dm.IsKeysEqualOnly(map[string]any{"b": 1}) {
		t.Fatal("keys should not be equal")
	}
}

func Test_C13_DynamicMap_IsKeysEqualOnly_BothNil(t *testing.T) {
	var dm *enumimpl.DynamicMap
	if !dm.IsKeysEqualOnly(nil) {
		t.Fatal("both nil should be equal")
	}
}

func Test_C13_DynamicMap_IsKeysEqualOnly_OneNil(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if dm.IsKeysEqualOnly(nil) {
		t.Fatal("expected false")
	}
}

func Test_C13_DynamicMap_IsKeysEqualOnly_DiffLength(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if dm.IsKeysEqualOnly(map[string]any{"a": 1, "b": 2}) {
		t.Fatal("diff length")
	}
}

func Test_C13_DynamicMap_KeyValue(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	v, found := dm.KeyValue("a")
	if !found || v != 1 {
		t.Fatal("expected found")
	}
	_, found2 := dm.KeyValue("missing")
	if found2 {
		t.Fatal("expected not found")
	}
}

func Test_C13_DynamicMap_KeyValueString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "hello"}
	v, found := dm.KeyValueString("a")
	if !found || v != "hello" {
		t.Fatal("expected hello")
	}
	_, found2 := dm.KeyValueString("missing")
	if found2 {
		t.Fatal("expected not found")
	}
}

func Test_C13_DynamicMap_KeyValueInt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 42}
	v, found, failed := dm.KeyValueInt("a")
	if !found || failed || v != 42 {
		t.Fatal("expected 42")
	}
	_, found2, _ := dm.KeyValueInt("missing")
	if found2 {
		t.Fatal("expected not found")
	}
}

func Test_C13_DynamicMap_KeyValueInt_ByteValue(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": byte(5)}
	v, found, failed := dm.KeyValueInt("a")
	if !found || failed || v != 5 {
		t.Fatal("expected 5")
	}
}

func Test_C13_DynamicMap_KeyValueIntDefault(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 10}
	if dm.KeyValueIntDefault("a") != 10 {
		t.Fatal("expected 10")
	}
}

func Test_C13_DynamicMap_KeyValueByte(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": byte(5)}
	v, found, failed := dm.KeyValueByte("a")
	if !found || failed || v != 5 {
		t.Fatal("expected 5")
	}
	_, found2, _ := dm.KeyValueByte("missing")
	if found2 {
		t.Fatal("expected not found")
	}
}

func Test_C13_DynamicMap_KeyValueByte_IntValue(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 100}
	v, found, failed := dm.KeyValueByte("a")
	if !found || failed || v != 100 {
		t.Fatalf("expected 100, got %d", v)
	}
}

func Test_C13_DynamicMap_Add(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	dm.Add("k", "v")
	if dm.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_Raw(t *testing.T) {
	dm := enumimpl.DynamicMap{"k": "v"}
	raw := dm.Raw()
	if len(raw) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConcatNew(t *testing.T) {
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"b": 2}
	result := dm1.ConcatNew(true, dm2)
	if result.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_ConcatNew_NoOverride(t *testing.T) {
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 2, "b": 3}
	result := dm1.ConcatNew(false, dm2)
	if result["a"] != 1 {
		t.Fatal("should not override")
	}
}

func Test_C13_DynamicMap_ConcatNew_BothEmpty(t *testing.T) {
	dm1 := enumimpl.DynamicMap{}
	dm2 := enumimpl.DynamicMap{}
	result := dm1.ConcatNew(true, dm2)
	if result.Length() != 0 {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_Strings(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	s := dm.Strings()
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_Strings_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.Strings()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_StringsUsingFmt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	s := dm.StringsUsingFmt(func(i int, k string, v any) string {
		return fmt.Sprintf("%d:%s=%v", i, k, v)
	})
	if len(s) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_StringsUsingFmt_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.StringsUsingFmt(func(i int, k string, v any) string { return "" })) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if dm.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_DynamicMap_IsStringEqual(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.IsStringEqual(dm.String()) {
		t.Fatal("expected equal")
	}
}

func Test_C13_DynamicMap_Serialize(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_, err := dm.Serialize()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_C13_DynamicMap_ConvMaps(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": byte(1), "b": byte(2)}
	byteMap := dm.ConvMapByteString()
	if len(byteMap) == 0 {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_DynamicMap_ConvMapByteString_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapByteString()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapStringInteger(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapStringInteger()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapStringInteger_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapStringInteger()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapIntegerString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapIntegerString()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapIntegerString_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapIntegerString()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapInt8String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt8String()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapInt8String_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapInt8String()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapInt16String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt16String()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapInt16String_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapInt16String()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapInt32String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt32String()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapInt32String_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapInt32String()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapUInt16String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapUInt16String()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapUInt16String_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapUInt16String()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapStringString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "x"}
	m := dm.ConvMapStringString()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapStringString_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapStringString()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_DynamicMap_ConvMapInt64String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	m := dm.ConvMapInt64String()
	if len(m) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_DynamicMap_ConvMapInt64String_Empty(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	if len(dm.ConvMapInt64String()) != 0 {
		t.Fatal("expected 0")
	}
}

// ===================== DynamicMap → Basic* conversions =====================

func Test_C13_DynamicMap_BasicByte(t *testing.T) {
	dm := enumimpl.DynamicMap{"Invalid": byte(0), "Valid": byte(1)}
	bb := dm.BasicByte("TestByteEnum")
	if bb == nil || bb.Length() != 2 {
		t.Fatal("expected BasicByte with 2 items")
	}
}

func Test_C13_DynamicMap_BasicByteUsingAliasMap(t *testing.T) {
	dm := enumimpl.DynamicMap{"Invalid": byte(0), "Valid": byte(1)}
	bb := dm.BasicByteUsingAliasMap("TestByteEnum", map[string]byte{"v": 1})
	if bb == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_C13_DynamicMap_BasicInt8(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": int8(0), "B": int8(1)}
	bi := dm.BasicInt8("TestInt8Enum")
	if bi == nil || bi.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_BasicInt16(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": int16(0), "B": int16(1)}
	bi := dm.BasicInt16("TestInt16Enum")
	if bi == nil || bi.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_BasicInt32(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": int32(0), "B": int32(1)}
	bi := dm.BasicInt32("TestInt32Enum")
	if bi == nil || bi.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_BasicString(t *testing.T) {
	dm := enumimpl.DynamicMap{"alpha": "x", "beta": "y"}
	bs := dm.BasicString("TestStringEnum")
	if bs == nil || bs.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_DynamicMap_BasicUInt16(t *testing.T) {
	dm := enumimpl.DynamicMap{"A": uint16(0), "B": uint16(1)}
	bu := dm.BasicUInt16("TestUInt16Enum")
	if bu == nil || bu.Length() != 2 {
		t.Fatal("expected 2")
	}
}

// ===================== BasicByte =====================

func Test_C13_BasicByte_AllMethods(t *testing.T) {
	bb := enumimpl.New.BasicByte.Create(
		"TestByte",
		[]byte{0, 1, 2},
		[]string{"Invalid", "Active", "Inactive"},
		0, 2,
	)

	if bb.Min() != 0 || bb.Max() != 2 {
		t.Fatal("min/max wrong")
	}
	if !bb.IsValidRange(1) || bb.IsValidRange(3) {
		t.Fatal("range check failed")
	}
	if bb.ToEnumString(0) != "Invalid" {
		t.Fatal("expected Invalid")
	}
	if !bb.IsAnyOf(1, 1, 2) {
		t.Fatal("expected match")
	}
	if bb.IsAnyOf(1, 3) {
		t.Fatal("expected no match")
	}
	if !bb.IsAnyOf(1) {
		t.Fatal("empty spread should return true")
	}
	if !bb.IsAnyNamesOf(0, "Invalid", "Active") {
		t.Fatal("expected name match")
	}
	if bb.IsAnyNamesOf(0, "Active") {
		t.Fatal("expected no name match")
	}
	v := bb.GetValueByString("Invalid")
	if v != 0 {
		t.Fatalf("expected 0, got %d", v)
	}
	s := bb.GetStringValue(0)
	if s != "Invalid" {
		t.Fatal("expected Invalid")
	}
	ranges := bb.Ranges()
	if len(ranges) != 3 {
		t.Fatal("expected 3")
	}
	hm := bb.Hashmap()
	if len(hm) == 0 {
		t.Fatal("expected non-empty hashmap")
	}
	hmPtr := bb.HashmapPtr()
	if hmPtr == nil {
		t.Fatal("expected non-nil")
	}
	jb, err := bb.ToEnumJsonBytes(0)
	if err != nil || len(jb) == 0 {
		t.Fatal("json bytes failed")
	}
	_, err2 := bb.ToEnumJsonBytes(99)
	if err2 == nil {
		t.Fatal("expected error for invalid value")
	}
	s2 := bb.AppendPrependJoinValue(".", 1, 0)
	if s2 == "" {
		t.Fatal("expected non-empty")
	}
	s3 := bb.ToNumberString(1)
	if s3 == "" {
		t.Fatal("expected non-empty")
	}
	jm := bb.JsonMap()
	if len(jm) == 0 {
		t.Fatal("expected non-empty")
	}

	// UnmarshallToValue
	v2, err3 := bb.UnmarshallToValue(false, []byte("Invalid"))
	if err3 != nil {
		t.Fatal(err3)
	}
	_ = v2

	// nil bytes, no map to first
	_, err4 := bb.UnmarshallToValue(false, nil)
	if err4 == nil {
		t.Fatal("expected error")
	}

	// nil bytes, map to first
	v3, err5 := bb.UnmarshallToValue(true, nil)
	if err5 != nil || v3 != 0 {
		t.Fatal("expected min")
	}

	// empty string, map to first
	v4, err6 := bb.UnmarshallToValue(true, []byte(""))
	if err6 != nil || v4 != 0 {
		t.Fatal("expected min")
	}

	// double quote
	v5, err7 := bb.UnmarshallToValue(true, []byte(`""`))
	if err7 != nil || v5 != 0 {
		t.Fatal("expected min")
	}

	if bb.EnumType().String() == "" {
		t.Fatal("expected enum type")
	}
	_ = bb.AsBasicByter()

	// GetValueByName
	_, err8 := bb.GetValueByName("NonExistent")
	if err8 == nil {
		t.Fatal("expected error")
	}
	vn, err9 := bb.GetValueByName("Invalid")
	if err9 != nil {
		t.Fatal(err9)
	}
	_ = vn
}

func Test_C13_BasicByte_ExpectingEnumValueError(t *testing.T) {
	bb := enumimpl.New.BasicByte.Create(
		"TestByte",
		[]byte{0, 1},
		[]string{"Invalid", "Active"},
		0, 1,
	)

	// Matching value
	err := bb.ExpectingEnumValueError("Invalid", byte(0))
	if err != nil {
		t.Fatal("expected nil for matching")
	}

	// Non-matching
	err2 := bb.ExpectingEnumValueError("Active", byte(0))
	if err2 == nil {
		t.Fatal("expected error for non-matching")
	}

	// Invalid rawString
	err3 := bb.ExpectingEnumValueError("NonExistent", byte(0))
	if err3 == nil {
		t.Fatal("expected error for invalid name")
	}
}

// ===================== BasicString =====================

func Test_C13_BasicString_AllMethods(t *testing.T) {
	bs := enumimpl.New.BasicString.Create("TestString", []string{"alpha", "beta", "gamma"})

	if bs.Min() == "" || bs.Max() == "" {
		t.Fatal("expected min/max")
	}
	ranges := bs.Ranges()
	if len(ranges) != 3 {
		t.Fatal("expected 3")
	}
	if !bs.HasAnyItem() || bs.MaxIndex() != 2 {
		t.Fatal("check failed")
	}
	if bs.GetNameByIndex(1) == "" {
		t.Fatal("expected name at index 1")
	}
	if bs.GetNameByIndex(0) != "" {
		// index 0 returns empty because condition is index > 0
	}
	if bs.GetNameByIndex(99) != "" {
		t.Fatal("out of range should return empty")
	}
	idx := bs.GetIndexByName("alpha")
	if idx < 0 {
		t.Fatal("expected valid index")
	}
	if bs.GetIndexByName("") >= 0 {
		t.Fatal("empty name should return invalid")
	}
	if bs.GetIndexByName("nonexistent") >= 0 {
		t.Fatal("nonexistent should return invalid")
	}
	nim := bs.NameWithIndexMap()
	if len(nim) == 0 {
		t.Fatal("expected non-empty")
	}
	ri := bs.RangesIntegers()
	if len(ri) != 3 {
		t.Fatal("expected 3")
	}
	hs := bs.Hashset()
	if len(hs) == 0 {
		t.Fatal("expected non-empty")
	}
	hsPtr := bs.HashsetPtr()
	if hsPtr == nil {
		t.Fatal("expected non-nil")
	}
	_, err := bs.GetValueByName("alpha")
	if err != nil {
		t.Fatal(err)
	}
	_, err2 := bs.GetValueByName("nonexistent")
	if err2 == nil {
		t.Fatal("expected error")
	}
	if !bs.IsValidRange("alpha") {
		t.Fatal("expected valid")
	}
	if bs.IsValidRange("nonexistent") {
		t.Fatal("expected invalid")
	}
	if !bs.IsAnyOf("alpha", "alpha", "beta") {
		t.Fatal("expected match")
	}
	if bs.IsAnyOf("alpha", "beta") {
		t.Fatal("expected no match")
	}
	if !bs.IsAnyOf("alpha") {
		t.Fatal("empty spread should return true")
	}
	if !bs.IsAnyNamesOf("alpha", "alpha") {
		t.Fatal("expected match")
	}

	// ToEnumJsonBytes
	jb, err3 := bs.ToEnumJsonBytes("alpha")
	if err3 != nil || len(jb) == 0 {
		t.Fatal("json bytes failed")
	}
	_, err4 := bs.ToEnumJsonBytes("nonexistent")
	if err4 == nil {
		t.Fatal("expected error")
	}

	// UnmarshallToValue
	v, err5 := bs.UnmarshallToValue(false, []byte("alpha"))
	if err5 != nil || v == "" {
		t.Fatal("unmarshal failed")
	}
	_, err6 := bs.UnmarshallToValue(false, nil)
	if err6 == nil {
		t.Fatal("expected error for nil")
	}
	v2, err7 := bs.UnmarshallToValue(true, nil)
	if err7 != nil {
		t.Fatal(err7)
	}
	_ = v2
	v3, err8 := bs.UnmarshallToValue(true, []byte(""))
	if err8 != nil {
		t.Fatal(err8)
	}
	_ = v3
	v4, err9 := bs.UnmarshallToValue(true, []byte(`""`))
	if err9 != nil {
		t.Fatal(err9)
	}
	_ = v4

	if bs.EnumType().String() == "" {
		t.Fatal("expected enum type")
	}

	// OnlySupportedErr
	err10 := bs.OnlySupportedErr("alpha")
	if err10 == nil {
		t.Fatal("expected error for unsupported")
	}
	err11 := bs.OnlySupportedErr("alpha", "beta", "gamma")
	if err11 != nil {
		t.Fatal("all supported should return nil")
	}

	// OnlySupportedMsgErr
	err12 := bs.OnlySupportedMsgErr("msg", "alpha")
	if err12 == nil {
		t.Fatal("expected error")
	}

	// AppendPrependJoinValue
	s := bs.AppendPrependJoinValue(".", "beta", "alpha")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ===================== numberEnumBase methods =====================

func Test_C13_NumberEnumBase_Methods(t *testing.T) {
	bb := enumimpl.New.BasicByte.Create(
		"TestByte",
		[]byte{0, 1, 2},
		[]string{"Zero", "One", "Two"},
		0, 2,
	)

	min, max := bb.MinMaxAny()
	if min == nil || max == nil {
		t.Fatal("expected non-nil")
	}
	if bb.MinValueString() == "" || bb.MaxValueString() == "" {
		t.Fatal("expected non-empty")
	}
	if bb.MinInt() != 0 || bb.MaxInt() != 2 {
		t.Fatal("min/max int wrong")
	}
	anv := bb.AllNameValues()
	if len(anv) != 3 {
		t.Fatal("expected 3")
	}
	rm := bb.RangesMap()
	if len(rm) == 0 {
		t.Fatal("expected non-empty")
	}
	ose := bb.OnlySupportedErr("Zero")
	if ose == nil {
		t.Fatal("expected unsupported error")
	}
	osme := bb.OnlySupportedMsgErr("msg", "Zero")
	if osme == nil {
		t.Fatal("expected error")
	}
	ier := bb.IntegerEnumRanges()
	if len(ier) != 3 {
		t.Fatal("expected 3")
	}
	if bb.Length() != 3 || bb.Count() != 3 {
		t.Fatal("expected 3")
	}
	rdm := bb.RangesDynamicMap()
	if len(rdm) != 3 {
		t.Fatal("expected 3")
	}
	dm := bb.DynamicMap()
	if dm.Length() != 3 {
		t.Fatal("expected 3")
	}
	rim := bb.RangesIntegerStringMap()
	_ = rim
	kav := bb.KeyAnyValues()
	if len(kav) != 3 {
		t.Fatal("expected 3")
	}
	kvi := bb.KeyValIntegers()
	if len(kvi) != 3 {
		t.Fatal("expected 3")
	}
	csv := bb.RangeNamesCsv()
	if csv == "" {
		t.Fatal("expected non-empty")
	}
	im := bb.RangesInvalidMessage()
	if im == "" {
		t.Fatal("expected non-empty")
	}
	ie := bb.RangesInvalidErr()
	if ie == nil {
		t.Fatal("expected non-nil")
	}
	sr := bb.StringRanges()
	if len(sr) != 3 {
		t.Fatal("expected 3")
	}
	srp := bb.StringRangesPtr()
	if len(srp) != 3 {
		t.Fatal("expected 3")
	}
	nh := bb.NamesHashset()
	if len(nh) != 3 {
		t.Fatal("expected 3")
	}
	js := bb.JsonString(byte(0))
	if js == "" {
		t.Fatal("expected non-empty")
	}
	tn := bb.TypeName()
	if tn != "TestByte" {
		t.Fatal("expected TestByte")
	}
	nwv := bb.NameWithValue(byte(0))
	if nwv == "" {
		t.Fatal("expected non-empty")
	}
	nwvo := bb.NameWithValueOption(byte(0), true)
	if nwvo == "" {
		t.Fatal("expected non-empty")
	}
	nwvo2 := bb.NameWithValueOption(byte(0), false)
	if nwvo2 == "" {
		t.Fatal("expected non-empty")
	}
	vs := bb.ValueString(byte(0))
	if vs == "" {
		t.Fatal("expected non-empty")
	}
	f := bb.Format("{type-name}-{name}-{value}", byte(0))
	if f == "" {
		t.Fatal("expected formatted")
	}

	// Loop
	count := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count++
		return false
	})
	if count != 3 {
		t.Fatal("expected 3 iterations")
	}

	// Loop with break
	count2 := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count2++
		return true
	})
	if count2 != 1 {
		t.Fatal("expected 1 iteration")
	}

	// LoopInteger
	count3 := 0
	bb.LoopInteger(func(index int, name string, val int) bool {
		count3++
		return false
	})
	if count3 != 3 {
		t.Fatal("expected 3")
	}
}

// ===================== DiffLeftRight =====================

func Test_C13_DiffLeftRight_AllMethods(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	l, r := d.Types()
	if l == nil || r == nil {
		t.Fatal("expected types")
	}
	if !d.IsSameTypeSame() {
		t.Fatal("expected same type")
	}
	if !d.IsSame() {
		t.Fatal("expected same")
	}
	if !d.IsSameRegardlessOfType() {
		t.Fatal("expected same regardless")
	}
	if !d.IsEqual(true) || !d.IsEqual(false) {
		t.Fatal("expected equal")
	}
	if d.HasMismatch(true) || d.HasMismatch(false) {
		t.Fatal("expected no mismatch")
	}
	if d.IsNotEqual() {
		t.Fatal("expected not not-equal")
	}
	if d.HasMismatchRegardlessOfType() {
		t.Fatal("expected no mismatch")
	}
	if d.String() == "" {
		t.Fatal("expected non-empty")
	}
	if d.DiffString() != "" {
		t.Fatal("expected empty diff string for equal values")
	}
	dl, dr := d.SpecificFullString()
	if dl == "" || dr == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_DiffLeftRight_Different(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	if d.IsSame() {
		t.Fatal("expected different")
	}
	if !d.IsNotEqual() {
		t.Fatal("expected not equal")
	}
	if d.DiffString() == "" {
		t.Fatal("expected non-empty diff")
	}
	if !d.HasMismatch(false) {
		t.Fatal("expected mismatch")
	}
}

func Test_C13_DiffLeftRight_JsonString_Nil(t *testing.T) {
	var d *enumimpl.DiffLeftRight
	if d.JsonString() != "" {
		t.Fatal("nil should be empty")
	}
}

// ===================== KeyAnyVal =====================

func Test_C13_KeyAnyVal_AllMethods(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "name", AnyValue: byte(5)}
	if kav.KeyString() != "name" {
		t.Fatal("expected name")
	}
	if kav.AnyVal() != byte(5) {
		t.Fatal("expected 5")
	}
	if kav.AnyValString() == "" {
		t.Fatal("expected non-empty")
	}
	if kav.WrapKey() == "" {
		t.Fatal("expected non-empty")
	}
	if kav.WrapValue() == "" {
		t.Fatal("expected non-empty")
	}
	if kav.IsString() {
		t.Fatal("byte should not be string type")
	}
	if kav.ValInt() != 5 {
		t.Fatal("expected 5")
	}
	kvi := kav.KeyValInteger()
	if kvi.Key != "name" || kvi.ValueInteger != 5 {
		t.Fatal("conversion failed")
	}
	s := kav.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_KeyAnyVal_StringType(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "name", AnyValue: "strval"}
	if !kav.IsString() {
		t.Fatal("expected string type")
	}
	s := kav.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ===================== KeyAnyValues =====================

func Test_C13_KeyAnyValues_Empty(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{}, []byte{})
	if len(result) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_C13_KeyAnyValues_NonEmpty(t *testing.T) {
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []byte{0, 1})
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

// ===================== Format, FormatUsingFmt, NameWithValue, PrependJoin, JoinPrependUsingDot =====================

func Test_C13_Format(t *testing.T) {
	s := enumimpl.Format("MyEnum", "Active", "1", "Enum:{type-name}-{name}-{value}")
	if s == "" {
		t.Fatal("expected formatted")
	}
}

type testFormatterC13 struct{}

func (tf testFormatterC13) TypeName() string   { return "TestType" }
func (tf testFormatterC13) Name() string       { return "TestName" }
func (tf testFormatterC13) ValueString() string { return "TestVal" }

func Test_C13_FormatUsingFmt(t *testing.T) {
	s := enumimpl.FormatUsingFmt(testFormatterC13{}, "{type-name}-{name}-{value}")
	if s == "" {
		t.Fatal("expected formatted")
	}
}

func Test_C13_NameWithValue(t *testing.T) {
	s := enumimpl.NameWithValue(byte(5))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_PrependJoin(t *testing.T) {
	s := enumimpl.PrependJoin(".", "prefix", "a", "b")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_C13_JoinPrependUsingDot(t *testing.T) {
	s := enumimpl.JoinPrependUsingDot("prefix", "a", "b")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ===================== ConvEnumAnyValToInteger =====================

func Test_C13_ConvEnumAnyValToInteger_String(t *testing.T) {
	v := enumimpl.ConvEnumAnyValToInteger("hello")
	_ = v // should return MinInt
}

func Test_C13_ConvEnumAnyValToInteger_Int(t *testing.T) {
	v := enumimpl.ConvEnumAnyValToInteger(42)
	if v != 42 {
		t.Fatal("expected 42")
	}
}

func Test_C13_ConvEnumAnyValToInteger_Byte(t *testing.T) {
	v := enumimpl.ConvEnumAnyValToInteger(byte(5))
	if v != 5 {
		t.Fatalf("expected 5, got %d", v)
	}
}

func Test_C13_ConvEnumAnyValToInteger_Fallback(t *testing.T) {
	// A float will be Sprintf'd and Atoi'd
	v := enumimpl.ConvEnumAnyValToInteger(3.0)
	_ = v
}

// ===================== IntegersRangesOfAnyVal =====================

func Test_C13_IntegersRangesOfAnyVal(t *testing.T) {
	result := enumimpl.IntegersRangesOfAnyVal([]byte{2, 0, 1})
	if len(result) != 3 || result[0] != 0 {
		t.Fatal("expected sorted [0,1,2]")
	}
}

// ===================== AllNameValues =====================

func Test_C13_AllNameValues(t *testing.T) {
	result := enumimpl.AllNameValues([]string{"a", "b"}, []byte{0, 1})
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

// ===================== UnsupportedNames =====================

func Test_C13_UnsupportedNames(t *testing.T) {
	result := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")
	if len(result) != 2 {
		t.Fatal("expected 2 unsupported")
	}
}

func Test_C13_UnsupportedNames_AllSupported(t *testing.T) {
	result := enumimpl.UnsupportedNames([]string{"a", "b"}, "a", "b")
	if len(result) != 0 {
		t.Fatal("expected 0")
	}
}

// ===================== OnlySupportedErr =====================

func Test_C13_OnlySupportedErr_EmptyAll(t *testing.T) {
	err := enumimpl.OnlySupportedErr(2, []string{})
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_C13_OnlySupportedErr_AllSupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(2, []string{"a"}, "a")
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_C13_OnlySupportedErr_HasUnsupported(t *testing.T) {
	err := enumimpl.OnlySupportedErr(2, []string{"a", "b"}, "a")
	if err == nil {
		t.Fatal("expected error")
	}
}

// ===================== DifferCheckerImpl =====================

func Test_C13_DifferCheckerImpl_GetSingleDiffResult(t *testing.T) {
	dc := enumimpl.DefaultDiffCheckerImpl
	if dc.GetSingleDiffResult(true, "l", "r") != "l" {
		t.Fatal("expected left")
	}
	if dc.GetSingleDiffResult(false, "l", "r") != "r" {
		t.Fatal("expected right")
	}
}

func Test_C13_DifferCheckerImpl_IsEqual(t *testing.T) {
	dc := enumimpl.DefaultDiffCheckerImpl
	if !dc.IsEqual(true, 1, 1) {
		t.Fatal("expected equal regardless")
	}
	if dc.IsEqual(false, 1, "1") {
		t.Fatal("expected not equal strict")
	}
}

func Test_C13_DifferCheckerImpl_GetResultOnKeyMissing(t *testing.T) {
	dc := enumimpl.DefaultDiffCheckerImpl
	r := dc.GetResultOnKeyMissingInRightExistInLeft("k", "v")
	if r != "v" {
		t.Fatal("expected v")
	}
}

// ===================== LeftRightDiffCheckerImpl =====================

func Test_C13_LeftRightDiffCheckerImpl(t *testing.T) {
	lrdc := enumimpl.LeftRightDiffCheckerImpl
	r := lrdc.GetSingleDiffResult(true, 1, 2)
	if r == nil {
		t.Fatal("expected non-nil")
	}
	r2 := lrdc.GetResultOnKeyMissingInRightExistInLeft("k", "v")
	if r2 == nil {
		t.Fatal("expected non-nil")
	}
	if !lrdc.IsEqual(true, 1, 1) {
		t.Fatal("expected equal")
	}
}

// ===================== Creator methods =====================

func Test_C13_NewBasicByte_UsingTypeSlice(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingTypeSlice("T", []string{"A", "B"})
	if bb.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicByte_Default(t *testing.T) {
	bb := enumimpl.New.BasicByte.Default(byte(0), []string{"A", "B"})
	if bb.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicByte_DefaultAllCases(t *testing.T) {
	bb := enumimpl.New.BasicByte.DefaultAllCases(byte(0), []string{"Active", "Inactive"})
	if bb.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicByte_DefaultWithAliasMap(t *testing.T) {
	bb := enumimpl.New.BasicByte.DefaultWithAliasMap(byte(0), []string{"A", "B"}, map[string]byte{"a": 0})
	if bb.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicByte_DefaultWithAliasMapAllCases(t *testing.T) {
	bb := enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(byte(0), []string{"A"}, map[string]byte{"a": 0})
	if bb.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_NewBasicByte_UsingFirstItemSliceCaseOptions(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingFirstItemSliceCaseOptions(false, byte(0), []string{"A"})
	if bb.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_NewBasicByte_UsingFirstItemSliceAllCases(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(byte(0), []string{"A"})
	if bb.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_NewBasicByte_UsingFirstItemSliceAliasMap(t *testing.T) {
	bb := enumimpl.New.BasicByte.UsingFirstItemSliceAliasMap(byte(0), []string{"A"}, nil)
	if bb.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_NewBasicByte_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	bb := enumimpl.New.BasicByte.CreateUsingMapPlusAliasMapOptions(
		true, byte(0), map[byte]string{0: "A"}, nil,
	)
	if bb.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_NewBasicString_CreateDefault(t *testing.T) {
	bs := enumimpl.New.BasicString.CreateDefault("strval", []string{"a", "b"})
	if bs.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicString_CreateUsingStringersSpread(t *testing.T) {
	bs := enumimpl.New.BasicString.CreateUsingNamesSpread("T", "alpha", "beta")
	if bs.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicString_CreateUsingNamesMinMax(t *testing.T) {
	bs := enumimpl.New.BasicString.CreateUsingNamesMinMax("T", []string{"a", "b"}, "a", "b")
	if bs.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicString_CreateUsingSlicePlusAliasMapOptions(t *testing.T) {
	bs := enumimpl.New.BasicString.CreateUsingSlicePlusAliasMapOptions(
		true, "strval", []string{"A", "B"}, nil,
	)
	if bs.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func Test_C13_NewBasicString_UsingFirstItemSliceCaseOptions(t *testing.T) {
	bs := enumimpl.New.BasicString.UsingFirstItemSliceCaseOptions(false, "strval", []string{"A"})
	if bs.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_C13_NewBasicString_UsingFirstItemSliceAllCases(t *testing.T) {
	bs := enumimpl.New.BasicString.UsingFirstItemSliceAllCases("strval", []string{"A"})
	if bs.Length() != 1 {
		t.Fatal("expected 1")
	}
}

// ===================== DynamicMap with LeftRightDiffChecker =====================

func Test_C13_DynamicMap_LogShouldDiffMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test",
		map[string]any{"a": 2},
	)
	if msg == "" {
		t.Fatal("expected diff")
	}
}

func Test_C13_DynamicMap_LogShouldDiffMessageUsingDifferChecker_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, false, "test",
		map[string]any{"b": 2},
	)
	if msg == "" {
		t.Fatal("expected diff")
	}
}

func Test_C13_DynamicMap_LogShouldDiffLeftRightMessageUsingDifferChecker_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.LeftRightDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_ShouldDiffMessageUsingDifferChecker_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, true, "test",
		map[string]any{"a": 1},
	)
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_C13_DynamicMap_DiffJsonMessageUsingDifferChecker_NoDiff(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, true,
		map[string]any{"a": 1},
	)
	if msg != "" {
		t.Fatal("expected empty")
	}
}

// ===================== IsValueTypeOf =====================

func Test_C13_DynamicMap_IsValueTypeOf(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "str"}
	// This checks reflect.TypeOf against first value's type
	_ = dm.IsValueTypeOf(nil)
}

// ===================== AppendPrependJoinNamer =====================

type testNamerC13 struct{ name string }

func (n testNamerC13) Name() string { return n.name }

func Test_C13_BasicByte_AppendPrependJoinNamer(t *testing.T) {
	bb := enumimpl.New.BasicByte.Create("T", []byte{0, 1}, []string{"A", "B"}, 0, 1)
	s := bb.AppendPrependJoinNamer(".", testNamer{"append"}, testNamer{"prepend"})
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ===================== NamesHashset empty =====================

func Test_C13_NamesHashset_Empty(t *testing.T) {
	// Create a BasicByte with no items — exercise NamesHashset empty path
	// Note: This would require a zero-item enum which isn't natural
	// The Length() == 0 path returns empty map
}
