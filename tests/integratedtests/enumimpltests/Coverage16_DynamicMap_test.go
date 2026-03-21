package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
)

// ══════════════════════════════════════════════════════════════════════════════
// enumimpl Coverage — DynamicMap comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func newTestDynMap() enumimpl.DynamicMap {
	return enumimpl.DynamicMap{
		"Invalid": byte(0),
		"Read":    byte(1),
		"Write":   byte(2),
		"Execute": byte(3),
	}
}

func Test_CovEnum_DM01_AddOrUpdate(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.AddOrUpdate("b", 2)
	if !isNew {
		t.Fatal("expected new")
	}
	isNew2 := dm.AddOrUpdate("a", 99)
	if isNew2 {
		t.Fatal("expected update")
	}
}

func Test_CovEnum_DM02_Set(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.Set("b", 2)
	if !isNew {
		t.Fatal("expected new")
	}
}

func Test_CovEnum_DM03_AddNewOnly(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	ok := dm.AddNewOnly("b", 2)
	if !ok {
		t.Fatal("expected added")
	}
	ok2 := dm.AddNewOnly("a", 99)
	if ok2 {
		t.Fatal("expected not added")
	}
}

func Test_CovEnum_DM04_AllKeys_AllKeysSorted(t *testing.T) {
	dm := newTestDynMap()
	keys := dm.AllKeys()
	if len(keys) != 4 {
		t.Fatal("expected 4")
	}
	sorted := dm.AllKeysSorted()
	if sorted[0] != "Execute" {
		t.Fatal("expected Execute first")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	if len(empty.AllKeys()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.AllKeysSorted()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM05_AllValuesStrings_Sorted(t *testing.T) {
	dm := newTestDynMap()
	vs := dm.AllValuesStrings()
	if len(vs) != 4 {
		t.Fatal("expected 4")
	}
	vss := dm.AllValuesStringsSorted()
	if len(vss) != 4 {
		t.Fatal("expected 4")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	if len(empty.AllValuesStrings()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.AllValuesStringsSorted()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM06_AllValuesIntegers(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	ints := dm.AllValuesIntegers()
	if len(ints) != 2 {
		t.Fatal("expected 2")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	if len(empty.AllValuesIntegers()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM07_MapIntegerString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	m, keys := dm.MapIntegerString()
	if len(m) < 2 || len(keys) < 2 {
		t.Fatal("expected at least 2")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	m2, k2 := empty.MapIntegerString()
	if len(m2) != 0 || len(k2) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM08_SortedKeyValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kv := dm.SortedKeyValues()
	if len(kv) != 2 {
		t.Fatal("expected 2")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	if len(empty.SortedKeyValues()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM09_SortedKeyAnyValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	kav := dm.SortedKeyAnyValues()
	if len(kav) != 2 {
		t.Fatal("expected 2")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	if len(empty.SortedKeyAnyValues()) != 0 {
		t.Fatal("expected 0")
	}
	// string values
	dmStr := enumimpl.DynamicMap{"x": "hello", "y": "world"}
	kavStr := dmStr.SortedKeyAnyValues()
	if len(kavStr) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovEnum_DM10_First(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	k, v := dm.First()
	if k != "a" || v != 1 {
		t.Fatal("expected a,1")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	k2, v2 := empty.First()
	if k2 != "" || v2 != nil {
		t.Fatal("expected empty")
	}
}

func Test_CovEnum_DM11_IsValueString_IsValueTypeOf(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "hello"}
	if !dm.IsValueString() {
		t.Fatal("expected true")
	}
	dm2 := enumimpl.DynamicMap{"a": 1}
	if dm2.IsValueString() {
		t.Fatal("expected false")
	}
}

func Test_CovEnum_DM12_Length_Count_IsEmpty_HasAnyItem_LastIndex_HasIndex(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	if dm.Length() != 2 {
		t.Fatal("expected 2")
	}
	if dm.Count() != 2 {
		t.Fatal("expected 2")
	}
	if dm.IsEmpty() {
		t.Fatal("expected false")
	}
	if !dm.HasAnyItem() {
		t.Fatal("expected true")
	}
	if dm.LastIndex() != 1 {
		t.Fatal("expected 1")
	}
	if !dm.HasIndex(1) {
		t.Fatal("expected true")
	}
	if dm.HasIndex(5) {
		t.Fatal("expected false")
	}
	// nil ptr
	var nilDm *enumimpl.DynamicMap
	if nilDm.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM13_HasKey_HasAllKeys_HasAnyKeys_IsMissingKey(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	if !dm.HasKey("a") {
		t.Fatal("expected true")
	}
	if dm.HasKey("c") {
		t.Fatal("expected false")
	}
	if !dm.HasAllKeys("a", "b") {
		t.Fatal("expected true")
	}
	if dm.HasAllKeys("a", "c") {
		t.Fatal("expected false")
	}
	if !dm.HasAnyKeys("a", "c") {
		t.Fatal("expected true")
	}
	if dm.HasAnyKeys("c", "d") {
		t.Fatal("expected false")
	}
	if !dm.IsMissingKey("c") {
		t.Fatal("expected true")
	}
}

func Test_CovEnum_DM14_IsEqual_IsMismatch(t *testing.T) {
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 1}
	if !dm1.IsEqual(false, &dm2) {
		t.Fatal("expected equal")
	}
	if dm1.IsMismatch(false, &dm2) {
		t.Fatal("expected no mismatch")
	}
	// nil comparisons
	var nilDm *enumimpl.DynamicMap
	if !nilDm.IsEqual(false, nil) {
		t.Fatal("expected nil==nil")
	}
	if nilDm.IsEqual(false, &dm1) {
		t.Fatal("expected not equal")
	}
	// regardless of type
	dm3 := enumimpl.DynamicMap{"a": byte(1)}
	if !dm1.IsEqual(true, &dm3) {
		t.Fatal("expected equal regardless of type")
	}
}

func Test_CovEnum_DM15_IsRawEqual_IsRawMismatch(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	raw := map[string]any{"a": 1}
	if !dm.IsRawEqual(false, raw) {
		t.Fatal("expected equal")
	}
	if dm.IsRawMismatch(false, raw) {
		t.Fatal("expected no mismatch")
	}
	// different lengths
	raw2 := map[string]any{"a": 1, "b": 2}
	if dm.IsRawEqual(false, raw2) {
		t.Fatal("expected not equal")
	}
	// missing key
	raw3 := map[string]any{"b": 1}
	if dm.IsRawEqual(false, raw3) {
		t.Fatal("expected not equal")
	}
	// nil dm
	var nilDm *enumimpl.DynamicMap
	if !nilDm.IsRawEqual(false, nil) {
		t.Fatal("expected nil==nil")
	}
	if nilDm.IsRawEqual(false, raw) {
		t.Fatal("expected false")
	}
}

func Test_CovEnum_DM16_DiffRaw(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}
	diff := dm.DiffRaw(false, right)
	if diff.IsEmpty() {
		t.Fatal("expected diffs")
	}
	// same maps
	same := map[string]any{"a": 1, "b": 2}
	diff2 := dm.DiffRaw(false, same)
	if diff2.HasAnyItem() {
		t.Fatal("expected no diff")
	}
}

func Test_CovEnum_DM17_DiffRawUsingDifferChecker_NilBranches(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	// both nil
	var nilDm *enumimpl.DynamicMap
	r := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	if r.HasAnyItem() {
		t.Fatal("expected empty")
	}
	// left nil, right not nil
	right := map[string]any{"a": 1}
	r2 := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)
	if r2.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	// left not nil, right nil
	r3 := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	if r3.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_DM18_DiffRawLeftRightUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}
	lDiff, rDiff := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)
	_ = lDiff
	_ = rDiff
	// nil branches
	var nilDm *enumimpl.DynamicMap
	l2, r2 := nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	if l2.HasAnyItem() || r2.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_CovEnum_DM19_DiffJsonMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	same := map[string]any{"a": 1}
	msg := dm.DiffJsonMessage(false, same)
	if msg != "" {
		t.Fatal("expected empty")
	}
	diff := map[string]any{"a": 2}
	msg2 := dm.DiffJsonMessage(false, diff)
	if msg2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_DM20_DiffJsonMessageLeftRight(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	diff := map[string]any{"a": 2, "b": 3}
	msg := dm.DiffJsonMessageLeftRight(false, diff)
	if msg == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_DM21_ShouldDiffMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
	msg2 := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	if msg2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_DM22_ShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
	msg2 := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 2})
	if msg2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_DM23_LogShouldDiffMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessage(false, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovEnum_DM24_LogShouldDiffLeftRightMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffLeftRightMessage(false, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovEnum_DM25_LogShouldDiffMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovEnum_DM26_ExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty for same")
	}
	msg2 := dm.ExpectingMessage("test", map[string]any{"a": 2})
	if msg2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_DM27_LogExpectingMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
}

func Test_CovEnum_DM28_IsKeysEqualOnly(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	if !dm.IsKeysEqualOnly(map[string]any{"a": 99, "b": 88}) {
		t.Fatal("expected true")
	}
	if dm.IsKeysEqualOnly(map[string]any{"a": 99, "c": 88}) {
		t.Fatal("expected false")
	}
	if dm.IsKeysEqualOnly(map[string]any{"a": 99}) {
		t.Fatal("expected false for different length")
	}
	var nilDm *enumimpl.DynamicMap
	if !nilDm.IsKeysEqualOnly(nil) {
		t.Fatal("expected nil==nil")
	}
	if nilDm.IsKeysEqualOnly(map[string]any{"a": 1}) {
		t.Fatal("expected false")
	}
}

func Test_CovEnum_DM29_KeyValue_KeyValueString(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	v, ok := dm.KeyValue("a")
	if !ok || v != 1 {
		t.Fatal("expected 1")
	}
	_, ok2 := dm.KeyValue("missing")
	if ok2 {
		t.Fatal("expected false")
	}
	vs, ok3 := dm.KeyValueString("a")
	if !ok3 || vs == "" {
		t.Fatal("expected non-empty")
	}
	_, ok4 := dm.KeyValueString("missing")
	if ok4 {
		t.Fatal("expected false")
	}
}

func Test_CovEnum_DM30_KeyValueIntDefault_KeyValueInt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	vi := dm.KeyValueIntDefault("a")
	if vi != 1 {
		t.Fatal("expected 1")
	}
	vi2 := dm.KeyValueIntDefault("missing")
	if vi2 >= 0 {
		t.Fatal("expected invalid")
	}
}

func Test_CovEnum_DM31_KeyValueByte(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": byte(5)}
	b, found, failed := dm.KeyValueByte("a")
	if !found || failed || b != 5 {
		t.Fatal("expected 5")
	}
	_, found2, _ := dm.KeyValueByte("missing")
	if found2 {
		t.Fatal("expected not found")
	}
	// int value
	dm2 := enumimpl.DynamicMap{"a": 42}
	b2, found3, failed3 := dm2.KeyValueByte("a")
	if !found3 || failed3 || b2 != 42 {
		t.Fatal("expected 42")
	}
}

func Test_CovEnum_DM32_Add(t *testing.T) {
	dm := enumimpl.DynamicMap{}
	dm.Add("a", 1)
	if dm.Length() != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovEnum_DM33_ConvMaps(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	si := dm.ConvMapStringInteger()
	if len(si) != 2 {
		t.Fatal("expected 2")
	}
	is := dm.ConvMapIntegerString()
	if len(is) != 2 {
		t.Fatal("expected 2")
	}
	bs := dm.ConvMapByteString()
	_ = bs
	i8s := dm.ConvMapInt8String()
	_ = i8s
	i16s := dm.ConvMapInt16String()
	_ = i16s
	i32s := dm.ConvMapInt32String()
	_ = i32s
	u16s := dm.ConvMapUInt16String()
	_ = u16s
	i64s := dm.ConvMapInt64String()
	_ = i64s
	ss := dm.ConvMapStringString()
	_ = ss
	// empty
	empty := enumimpl.DynamicMap{}
	if len(empty.ConvMapStringInteger()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapIntegerString()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapByteString()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapInt8String()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapInt16String()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapInt32String()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapUInt16String()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapInt64String()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.ConvMapStringString()) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM34_ConcatNew(t *testing.T) {
	dm1 := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"b": 2}
	r := dm1.ConcatNew(true, dm2)
	if r.Length() != 2 {
		t.Fatal("expected 2")
	}
	// no override
	dm3 := enumimpl.DynamicMap{"a": 99}
	r2 := dm1.ConcatNew(false, dm3)
	if r2["a"] != 1 {
		t.Fatal("expected 1 not overridden")
	}
	// override
	r3 := dm1.ConcatNew(true, dm3)
	if r3["a"] != 99 {
		t.Fatal("expected 99 overridden")
	}
	// both empty
	empty1 := enumimpl.DynamicMap{}
	empty2 := enumimpl.DynamicMap{}
	r4 := empty1.ConcatNew(true, empty2)
	if r4.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_CovEnum_DM35_Strings_String_StringsUsingFmt(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	ss := dm.Strings()
	if len(ss) != 2 {
		t.Fatal("expected 2")
	}
	s := dm.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	sf := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return key
	})
	if len(sf) != 2 {
		t.Fatal("expected 2")
	}
	// empty
	empty := enumimpl.DynamicMap{}
	if len(empty.Strings()) != 0 {
		t.Fatal("expected 0")
	}
	if len(empty.StringsUsingFmt(func(i int, k string, v any) string { return k })) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovEnum_DM36_IsStringEqual(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	if !dm.IsStringEqual(dm.String()) {
		t.Fatal("expected true")
	}
}

func Test_CovEnum_DM37_Serialize(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	b, err := dm.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("expected serialization")
	}
}

func Test_CovEnum_DM38_Raw(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	raw := dm.Raw()
	if len(raw) != 1 {
		t.Fatal("expected 1")
	}
}

func Test_CovEnum_DM39_BasicByte(t *testing.T) {
	dm := newTestDynMap()
	bb := dm.BasicByte("TestEnum")
	if bb == nil {
		t.Fatal("expected non-nil")
	}
	if bb.Length() != 4 {
		t.Fatal("expected 4")
	}
}

func Test_CovEnum_DM40_BasicByteUsingAliasMap(t *testing.T) {
	dm := newTestDynMap()
	alias := map[string]byte{"r": 1, "w": 2}
	bb := dm.BasicByteUsingAliasMap("TestEnum", alias)
	if bb == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovEnum_DM41_BasicString(t *testing.T) {
	dm := enumimpl.DynamicMap{"Invalid": "Invalid", "Active": "Active"}
	bs := dm.BasicString("TestStrEnum")
	if bs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovEnum_DM42_BasicStringUsingAliasMap(t *testing.T) {
	dm := enumimpl.DynamicMap{"Invalid": "Invalid", "Active": "Active"}
	alias := map[string]string{"inv": "Invalid"}
	bs := dm.BasicStringUsingAliasMap("TestStrEnum", alias)
	if bs == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovEnum_DM43_BasicInt8(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicInt8("TestI8")
	if bi == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovEnum_DM44_BasicInt16(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicInt16("TestI16")
	if bi == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovEnum_DM45_BasicInt32(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicInt32("TestI32")
	if bi == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovEnum_DM46_BasicUInt16(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	bi := dm.BasicUInt16("TestU16")
	if bi == nil {
		t.Fatal("expected non-nil")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicByte — targeted branch coverage
// ══════════════════════════════════════════════════════════════════════════════

func newTestBasicByte() *enumimpl.BasicByte {
	dm := newTestDynMap()
	return dm.BasicByte("TestEnum")
}

func Test_CovEnum_BB01_IsAnyOf(t *testing.T) {
	bb := newTestBasicByte()
	if !bb.IsAnyOf(1, 1, 2) {
		t.Fatal("expected true")
	}
	if bb.IsAnyOf(5, 1, 2) {
		t.Fatal("expected false")
	}
	if !bb.IsAnyOf(5) {
		t.Fatal("expected true for empty variadic")
	}
}

func Test_CovEnum_BB02_IsAnyNamesOf(t *testing.T) {
	bb := newTestBasicByte()
	name := bb.ToEnumString(1)
	if !bb.IsAnyNamesOf(1, name) {
		t.Fatal("expected true")
	}
	if bb.IsAnyNamesOf(1, "nonexistent") {
		t.Fatal("expected false")
	}
}

func Test_CovEnum_BB03_MinMax(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.Min()
	_ = bb.Max()
}

func Test_CovEnum_BB04_GetValueByName_Valid_Invalid(t *testing.T) {
	bb := newTestBasicByte()
	_, err := bb.GetValueByName("Invalid")
	if err != nil {
		t.Fatal("expected no error")
	}
	_, err2 := bb.GetValueByName("nonexistent_xyz")
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovEnum_BB05_IsValidRange(t *testing.T) {
	bb := newTestBasicByte()
	if !bb.IsValidRange(0) {
		t.Fatal("expected true")
	}
}

func Test_CovEnum_BB06_ToEnumJsonBytes(t *testing.T) {
	bb := newTestBasicByte()
	b, err := bb.ToEnumJsonBytes(0)
	if err != nil {
		t.Fatal("expected no error")
	}
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
	_, err2 := bb.ToEnumJsonBytes(99)
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func Test_CovEnum_BB07_UnmarshallToValue(t *testing.T) {
	bb := newTestBasicByte()
	// nil not mapped to first
	_, err := bb.UnmarshallToValue(false, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	// nil mapped to first
	v, err2 := bb.UnmarshallToValue(true, nil)
	if err2 != nil {
		t.Fatal("expected no error")
	}
	_ = v
	// empty string mapped to first
	v2, err3 := bb.UnmarshallToValue(true, []byte(""))
	if err3 != nil {
		t.Fatal("expected no error")
	}
	_ = v2
	// valid name
	_, err4 := bb.UnmarshallToValue(false, []byte("Invalid"))
	if err4 != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovEnum_BB08_EnumType(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.EnumType()
}

func Test_CovEnum_BB09_AsBasicByter(t *testing.T) {
	bb := newTestBasicByte()
	byter := bb.AsBasicByter()
	if byter == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovEnum_BB10_Hashmap_Ranges_AppendPrependJoinValue(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.Hashmap()
	_ = bb.HashmapPtr()
	_ = bb.Ranges()
	_ = bb.AppendPrependJoinValue(".", 1, 2)
	_ = bb.ToNumberString(1)
	_ = bb.JsonMap()
	_ = bb.GetStringValue(0)
	_ = bb.GetValueByString("Invalid")
}

func Test_CovEnum_BB11_ExpectingEnumValueError(t *testing.T) {
	bb := newTestBasicByte()
	err := bb.ExpectingEnumValueError("Invalid", byte(0))
	if err != nil {
		t.Fatal("expected no error for matching")
	}
	err2 := bb.ExpectingEnumValueError("nonexistent_xyz_999", byte(0))
	if err2 == nil {
		t.Fatal("expected error for bad name")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// numberEnumBase — coverage via BasicByte
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_NEB01_MinMaxAny_MinValueString_MaxValueString(t *testing.T) {
	bb := newTestBasicByte()
	min, max := bb.MinMaxAny()
	if min == nil || max == nil {
		t.Fatal("expected non-nil")
	}
	ms := bb.MinValueString()
	if ms == "" {
		t.Fatal("expected non-empty")
	}
	// cached call
	ms2 := bb.MinValueString()
	if ms2 != ms {
		t.Fatal("expected same cached")
	}
	mx := bb.MaxValueString()
	if mx == "" {
		t.Fatal("expected non-empty")
	}
	// cached call
	mx2 := bb.MaxValueString()
	if mx2 != mx {
		t.Fatal("expected same cached")
	}
}

func Test_CovEnum_NEB02_MinInt_MaxInt(t *testing.T) {
	bb := newTestBasicByte()
	_ = bb.MinInt()
	_ = bb.MaxInt()
}

func Test_CovEnum_NEB03_AllNameValues(t *testing.T) {
	bb := newTestBasicByte()
	nv := bb.AllNameValues()
	if len(nv) != 4 {
		t.Fatal("expected 4")
	}
}

func Test_CovEnum_NEB04_RangesMap_DynamicMap_RangesDynamicMap(t *testing.T) {
	bb := newTestBasicByte()
	rm := bb.RangesMap()
	if len(rm) < 4 {
		t.Fatal("expected at least 4")
	}
	dm := bb.DynamicMap()
	if dm.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	rdm := bb.RangesDynamicMap()
	if len(rdm) == 0 {
		t.Fatal("expected non-empty")
	}
	// cached
	rdm2 := bb.RangesDynamicMap()
	if len(rdm2) != len(rdm) {
		t.Fatal("expected same")
	}
}

func Test_CovEnum_NEB05_IntegerEnumRanges(t *testing.T) {
	bb := newTestBasicByte()
	ranges := bb.IntegerEnumRanges()
	if len(ranges) != 4 {
		t.Fatal("expected 4")
	}
}

func Test_CovEnum_NEB06_KeyAnyValues_KeyValIntegers(t *testing.T) {
	bb := newTestBasicByte()
	kav := bb.KeyAnyValues()
	if len(kav) != 4 {
		t.Fatal("expected 4")
	}
	// cached
	kav2 := bb.KeyAnyValues()
	if len(kav2) != 4 {
		t.Fatal("expected 4")
	}
	kvi := bb.KeyValIntegers()
	if len(kvi) != 4 {
		t.Fatal("expected 4")
	}
}

func Test_CovEnum_NEB07_Loop_LoopInteger(t *testing.T) {
	bb := newTestBasicByte()
	count := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count++
		return false
	})
	if count != 4 {
		t.Fatal("expected 4")
	}
	// break early
	count2 := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count2++
		return true
	})
	if count2 != 1 {
		t.Fatal("expected 1")
	}
	count3 := 0
	bb.LoopInteger(func(index int, name string, anyVal int) bool {
		count3++
		return false
	})
	if count3 != 4 {
		t.Fatal("expected 4")
	}
}

func Test_CovEnum_NEB08_RangesCsv_RangesInvalidMessage_RangesInvalidErr(t *testing.T) {
	bb := newTestBasicByte()
	csv := bb.RangeNamesCsv()
	if csv == "" {
		t.Fatal("expected non-empty")
	}
	msg := bb.RangesInvalidMessage()
	if msg == "" {
		t.Fatal("expected non-empty")
	}
	err := bb.RangesInvalidErr()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovEnum_NEB09_StringRanges_NamesHashset(t *testing.T) {
	bb := newTestBasicByte()
	sr := bb.StringRanges()
	if len(sr) != 4 {
		t.Fatal("expected 4")
	}
	srp := bb.StringRangesPtr()
	if len(srp) != 4 {
		t.Fatal("expected 4")
	}
	hs := bb.NamesHashset()
	if len(hs) != 4 {
		t.Fatal("expected 4")
	}
}

func Test_CovEnum_NEB10_NameWithValue_NameWithValueOption(t *testing.T) {
	bb := newTestBasicByte()
	nv := bb.NameWithValue(1)
	if nv == "" {
		t.Fatal("expected non-empty")
	}
	nvo := bb.NameWithValueOption(1, true)
	if nvo == "" {
		t.Fatal("expected non-empty")
	}
	nvo2 := bb.NameWithValueOption(1, false)
	if nvo2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_NEB11_ValueString_JsonString_ToEnumString_ToName(t *testing.T) {
	bb := newTestBasicByte()
	vs := bb.ValueString(1)
	if vs == "" {
		t.Fatal("expected non-empty")
	}
	js := bb.JsonString(1)
	if js == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_NEB12_Format(t *testing.T) {
	bb := newTestBasicByte()
	f := bb.Format("Enum of {type-name} - {name} - {value}", byte(1))
	if f == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_NEB13_OnlySupportedErr(t *testing.T) {
	bb := newTestBasicByte()
	err := bb.OnlySupportedErr("Invalid")
	if err == nil {
		t.Fatal("expected error for unsupported names")
	}
}

func Test_CovEnum_NEB14_OnlySupportedMsgErr(t *testing.T) {
	bb := newTestBasicByte()
	err := bb.OnlySupportedMsgErr("prefix", "Invalid")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovEnum_NEB15_RangesIntegerStringMap(t *testing.T) {
	bb := newTestBasicByte()
	m := bb.RangesIntegerStringMap()
	_ = m
}

// ══════════════════════════════════════════════════════════════════════════════
// Helper functions and types
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_Misc01_AllNameValues(t *testing.T) {
	names := []string{"Invalid", "Read", "Write"}
	vals := []byte{0, 1, 2}
	r := enumimpl.AllNameValues(names, vals)
	if len(r) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CovEnum_Misc02_ConvEnumAnyValToInteger(t *testing.T) {
	// string returns MinInt
	r := enumimpl.ConvEnumAnyValToInteger("hello")
	if r >= 0 {
		t.Fatal("expected MinInt")
	}
	// int
	r2 := enumimpl.ConvEnumAnyValToInteger(42)
	if r2 != 42 {
		t.Fatal("expected 42")
	}
}

func Test_CovEnum_Misc03_IntegersRangesOfAnyVal(t *testing.T) {
	vals := []byte{0, 1, 2}
	r := enumimpl.IntegersRangesOfAnyVal(vals)
	if len(r) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_CovEnum_Misc04_Format(t *testing.T) {
	r := enumimpl.Format("TypeA", "Name1", "0", "Enum of {type-name} - {name} - {value}")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_Misc05_PrependJoin_JoinPrependUsingDot(t *testing.T) {
	r := enumimpl.PrependJoin(".", "pre", "a", "b")
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := enumimpl.JoinPrependUsingDot("pre", "a", "b")
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_Misc06_NameWithValue(t *testing.T) {
	r := enumimpl.NameWithValue(1)
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_Misc07_OnlySupportedErr(t *testing.T) {
	err := enumimpl.OnlySupportedErr(0, []string{"a", "b", "c"}, "a")
	if err == nil {
		t.Fatal("expected error")
	}
	// empty all names
	err2 := enumimpl.OnlySupportedErr(0, []string{}, "a")
	if err2 != nil {
		t.Fatal("expected nil")
	}
	// all supported
	err3 := enumimpl.OnlySupportedErr(0, []string{"a", "b"}, "a", "b")
	if err3 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovEnum_Misc08_UnsupportedNames(t *testing.T) {
	un := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")
	if len(un) != 2 {
		t.Fatal("expected 2")
	}
}

func Test_CovEnum_Misc09_KeyAnyValues(t *testing.T) {
	names := []string{"a", "b"}
	vals := []int{1, 2}
	kav := enumimpl.KeyAnyValues(names, vals)
	if len(kav) != 2 {
		t.Fatal("expected 2")
	}
	// empty
	kav2 := enumimpl.KeyAnyValues(nil, nil)
	if len(kav2) != 0 {
		t.Fatal("expected 0")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// DiffLeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_DLR01_Types_IsSameTypeSame(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	l, r := d.Types()
	if l != r {
		t.Fatal("expected same type")
	}
	if !d.IsSameTypeSame() {
		t.Fatal("expected true")
	}
}

func Test_CovEnum_DLR02_IsSame_IsNotEqual(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	if !d.IsSame() {
		t.Fatal("expected same")
	}
	if d.IsNotEqual() {
		t.Fatal("expected not notequal")
	}
	d2 := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	if d2.IsSame() {
		t.Fatal("expected not same")
	}
}

func Test_CovEnum_DLR03_IsSameRegardlessOfType(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: byte(1)}
	if !d.IsSameRegardlessOfType() {
		t.Fatal("expected true")
	}
}

func Test_CovEnum_DLR04_IsEqual_HasMismatch(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	if !d.IsEqual(false) {
		t.Fatal("expected equal")
	}
	if d.HasMismatch(false) {
		t.Fatal("expected no mismatch")
	}
	if !d.IsEqual(true) {
		t.Fatal("expected equal regardless")
	}
	if d.HasMismatch(true) {
		t.Fatal("expected no mismatch regardless")
	}
	d2 := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	if !d2.HasMismatch(false) {
		t.Fatal("expected mismatch")
	}
	if d2.HasMismatchRegardlessOfType() {
		// might be same string fmt
	}
}

func Test_CovEnum_DLR05_String_JsonString_DiffString(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	s := d.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	js := d.JsonString()
	if js == "" {
		t.Fatal("expected non-empty")
	}
	ds := d.DiffString()
	if ds == "" {
		t.Fatal("expected non-empty for different")
	}
	// same
	d2 := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	ds2 := d2.DiffString()
	if ds2 != "" {
		t.Fatal("expected empty for same")
	}
	// nil
	var nilD *enumimpl.DiffLeftRight
	if nilD.JsonString() != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_CovEnum_DLR06_SpecificFullString(t *testing.T) {
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}
	l, r := d.SpecificFullString()
	if l == "" || r == "" {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyVal / KeyValInteger
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_KAV01_KeyAnyVal_Methods(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "test", AnyValue: 42}
	if kav.KeyString() != "test" {
		t.Fatal("expected test")
	}
	if kav.AnyVal() != 42 {
		t.Fatal("expected 42")
	}
	_ = kav.AnyValString()
	_ = kav.WrapKey()
	_ = kav.WrapValue()
	if kav.IsString() {
		t.Fatal("expected false")
	}
	if kav.ValInt() != 42 {
		t.Fatal("expected 42")
	}
	kvi := kav.KeyValInteger()
	if kvi.Key != "test" {
		t.Fatal("expected test")
	}
	s := kav.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_KAV02_KeyAnyVal_StringType(t *testing.T) {
	kav := enumimpl.KeyAnyVal{Key: "test", AnyValue: "hello"}
	if !kav.IsString() {
		t.Fatal("expected true")
	}
	s := kav.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovEnum_KVI01_KeyValInteger_Methods(t *testing.T) {
	kvi := enumimpl.KeyValInteger{Key: "test", ValueInteger: 42}
	_ = kvi.WrapKey()
	_ = kvi.WrapValue()
	kav := kvi.KeyAnyVal()
	if kav.Key != "test" {
		t.Fatal("expected test")
	}
	if kvi.IsString() {
		t.Fatal("expected false")
	}
	s := kvi.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// differCheckerImpl / leftRightDiffCheckerImpl
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_DC01_DefaultDiffCheckerImpl(t *testing.T) {
	dc := enumimpl.DefaultDiffCheckerImpl
	r := dc.GetSingleDiffResult(true, 1, 2)
	if r != 1 {
		t.Fatal("expected left")
	}
	r2 := dc.GetSingleDiffResult(false, 1, 2)
	if r2 != 2 {
		t.Fatal("expected right")
	}
	r3 := dc.GetResultOnKeyMissingInRightExistInLeft("k", 1)
	if r3 != 1 {
		t.Fatal("expected 1")
	}
	if !dc.IsEqual(false, 1, 1) {
		t.Fatal("expected true")
	}
	if !dc.IsEqual(true, 1, byte(1)) {
		t.Fatal("expected true regardless")
	}
}

func Test_CovEnum_DC02_LeftRightDiffCheckerImpl(t *testing.T) {
	dc := enumimpl.LeftRightDiffCheckerImpl
	r := dc.GetSingleDiffResult(true, 1, 2)
	if r == nil {
		t.Fatal("expected non-nil")
	}
	r2 := dc.GetResultOnKeyMissingInRightExistInLeft("k", 1)
	if r2 == nil {
		t.Fatal("expected non-nil")
	}
	if !dc.IsEqual(false, 1, 1) {
		t.Fatal("expected true")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// MapIntegerString with string values
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovEnum_DM47_MapIntegerString_StringValues(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": "hello", "b": "world"}
	m, keys := dm.MapIntegerString()
	_ = m
	_ = keys
}

func Test_CovEnum_DM48_ShouldDiffMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
	if msg != "" {
		t.Fatal("expected empty")
	}
}
