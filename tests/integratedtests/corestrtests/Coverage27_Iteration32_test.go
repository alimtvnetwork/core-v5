package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ═══════════════════════════════════════════════════════════════════════
// ValidValue — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_01_ValidValue_NewValidValue(t *testing.T) {
	v := corestr.NewValidValue("hello")
	if v.Value != "hello" || !v.IsValid {
		t.Error("expected hello, valid")
	}
}

func Test_C27_02_ValidValue_NewValidValueEmpty(t *testing.T) {
	v := corestr.NewValidValueEmpty()
	if v.Value != "" || !v.IsValid {
		t.Error("expected empty, valid")
	}
}

func Test_C27_03_ValidValue_InvalidValidValue(t *testing.T) {
	v := corestr.InvalidValidValue("err")
	if v.IsValid || v.Message != "err" {
		t.Error("expected invalid with err")
	}
}

func Test_C27_04_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	v := corestr.InvalidValidValueNoMessage()
	if v.IsValid || v.Message != "" {
		t.Error("expected invalid no message")
	}
}

func Test_C27_05_ValidValue_NewValidValueUsingAny(t *testing.T) {
	v := corestr.NewValidValueUsingAny(false, true, 42)
	if !v.IsValid || v.Value == "" {
		t.Error("expected valid with value")
	}
}

func Test_C27_06_ValidValue_NewValidValueUsingAnyAutoValid(t *testing.T) {
	v := corestr.NewValidValueUsingAnyAutoValid(false, 42)
	_ = v
}

func Test_C27_07_ValidValue_ValueBytesOnce(t *testing.T) {
	v := corestr.NewValidValue("hello")
	b := v.ValueBytesOnce()
	if string(b) != "hello" {
		t.Error("expected hello")
	}
	// second call should return same
	b2 := v.ValueBytesOnce()
	if string(b2) != "hello" {
		t.Error("expected cached hello")
	}
}

func Test_C27_08_ValidValue_ValueBytesOncePtr(t *testing.T) {
	v := corestr.NewValidValue("hi")
	b := v.ValueBytesOncePtr()
	if string(b) != "hi" {
		t.Error("expected hi")
	}
}

func Test_C27_09_ValidValue_IsEmpty(t *testing.T) {
	v := corestr.NewValidValueEmpty()
	if !v.IsEmpty() {
		t.Error("expected empty")
	}
	v2 := corestr.NewValidValue("x")
	if v2.IsEmpty() {
		t.Error("expected not empty")
	}
}

func Test_C27_10_ValidValue_IsWhitespace(t *testing.T) {
	v := corestr.NewValidValue("   ")
	if !v.IsWhitespace() {
		t.Error("expected whitespace")
	}
}

func Test_C27_11_ValidValue_Trim(t *testing.T) {
	v := corestr.NewValidValue("  hello  ")
	if v.Trim() != "hello" {
		t.Error("expected trimmed")
	}
}

func Test_C27_12_ValidValue_HasValidNonEmpty(t *testing.T) {
	v := corestr.NewValidValue("x")
	if !v.HasValidNonEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_13_ValidValue_HasValidNonWhitespace(t *testing.T) {
	v := corestr.NewValidValue("x")
	if !v.HasValidNonWhitespace() {
		t.Error("expected true")
	}
}

func Test_C27_14_ValidValue_ValueBool(t *testing.T) {
	v := corestr.NewValidValue("true")
	if !v.ValueBool() {
		t.Error("expected true")
	}
	v2 := corestr.NewValidValue("nope")
	if v2.ValueBool() {
		t.Error("expected false")
	}
	v3 := corestr.NewValidValue("")
	if v3.ValueBool() {
		t.Error("expected false for empty")
	}
}

func Test_C27_15_ValidValue_ValueInt(t *testing.T) {
	v := corestr.NewValidValue("42")
	if v.ValueInt(0) != 42 {
		t.Error("expected 42")
	}
	v2 := corestr.NewValidValue("bad")
	if v2.ValueInt(99) != 99 {
		t.Error("expected default 99")
	}
}

func Test_C27_16_ValidValue_ValueDefInt(t *testing.T) {
	v := corestr.NewValidValue("10")
	if v.ValueDefInt() != 10 {
		t.Error("expected 10")
	}
}

func Test_C27_17_ValidValue_ValueByte(t *testing.T) {
	v := corestr.NewValidValue("200")
	if v.ValueByte(0) != 200 {
		t.Error("expected 200")
	}
	v2 := corestr.NewValidValue("300")
	if v2.ValueByte(0) != 255 {
		t.Error("expected 255 for overflow")
	}
	v3 := corestr.NewValidValue("-1")
	if v3.ValueByte(0) != 0 {
		t.Error("expected 0 for negative")
	}
}

func Test_C27_18_ValidValue_ValueDefByte(t *testing.T) {
	v := corestr.NewValidValue("100")
	if v.ValueDefByte() != 100 {
		t.Error("expected 100")
	}
}

func Test_C27_19_ValidValue_ValueFloat64(t *testing.T) {
	v := corestr.NewValidValue("3.14")
	if v.ValueFloat64(0) != 3.14 {
		t.Error("expected 3.14")
	}
	v2 := corestr.NewValidValue("bad")
	if v2.ValueFloat64(1.0) != 1.0 {
		t.Error("expected default")
	}
}

func Test_C27_20_ValidValue_ValueDefFloat64(t *testing.T) {
	v := corestr.NewValidValue("2.5")
	if v.ValueDefFloat64() != 2.5 {
		t.Error("expected 2.5")
	}
}

func Test_C27_21_ValidValue_HasSafeNonEmpty(t *testing.T) {
	v := corestr.NewValidValue("x")
	if !v.HasSafeNonEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_22_ValidValue_Is(t *testing.T) {
	v := corestr.NewValidValue("hello")
	if !v.Is("hello") {
		t.Error("expected true")
	}
}

func Test_C27_23_ValidValue_IsAnyOf(t *testing.T) {
	v := corestr.NewValidValue("b")
	if !v.IsAnyOf("a", "b", "c") {
		t.Error("expected true")
	}
	if !v.IsAnyOf() {
		t.Error("expected true for empty values")
	}
	if v.IsAnyOf("x", "y") {
		t.Error("expected false")
	}
}

func Test_C27_24_ValidValue_IsContains(t *testing.T) {
	v := corestr.NewValidValue("hello world")
	if !v.IsContains("world") {
		t.Error("expected true")
	}
}

func Test_C27_25_ValidValue_IsAnyContains(t *testing.T) {
	v := corestr.NewValidValue("hello world")
	if !v.IsAnyContains("xyz", "world") {
		t.Error("expected true")
	}
	if !v.IsAnyContains() {
		t.Error("expected true for empty")
	}
	if v.IsAnyContains("xyz", "abc") {
		t.Error("expected false")
	}
}

func Test_C27_26_ValidValue_IsEqualNonSensitive(t *testing.T) {
	v := corestr.NewValidValue("Hello")
	if !v.IsEqualNonSensitive("hello") {
		t.Error("expected true")
	}
}

func Test_C27_27_ValidValue_IsRegexMatches(t *testing.T) {
	v := corestr.NewValidValue("abc123")
	re := regexp.MustCompile(`\d+`)
	if !v.IsRegexMatches(re) {
		t.Error("expected true")
	}
	if v.IsRegexMatches(nil) {
		t.Error("expected false for nil regex")
	}
}

func Test_C27_28_ValidValue_RegexFindString(t *testing.T) {
	v := corestr.NewValidValue("abc123")
	re := regexp.MustCompile(`\d+`)
	if v.RegexFindString(re) != "123" {
		t.Error("expected 123")
	}
	if v.RegexFindString(nil) != "" {
		t.Error("expected empty for nil")
	}
}

func Test_C27_29_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	v := corestr.NewValidValue("a1b2c3")
	re := regexp.MustCompile(`\d`)
	items, has := v.RegexFindAllStringsWithFlag(re, -1)
	if !has || len(items) != 3 {
		t.Error("expected 3")
	}
	_, has2 := v.RegexFindAllStringsWithFlag(nil, -1)
	if has2 {
		t.Error("expected false for nil")
	}
}

func Test_C27_30_ValidValue_RegexFindAllStrings(t *testing.T) {
	v := corestr.NewValidValue("a1b2")
	re := regexp.MustCompile(`\d`)
	items := v.RegexFindAllStrings(re, -1)
	if len(items) != 2 {
		t.Error("expected 2")
	}
	items2 := v.RegexFindAllStrings(nil, -1)
	if len(items2) != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_31_ValidValue_Split(t *testing.T) {
	v := corestr.NewValidValue("a,b,c")
	s := v.Split(",")
	if len(s) != 3 {
		t.Error("expected 3")
	}
}

func Test_C27_32_ValidValue_SplitNonEmpty(t *testing.T) {
	v := corestr.NewValidValue("a,,b")
	s := v.SplitNonEmpty(",")
	_ = s // just no panic
}

func Test_C27_33_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	v := corestr.NewValidValue("a , b , c")
	s := v.SplitTrimNonWhitespace(",")
	_ = s
}

func Test_C27_34_ValidValue_Clone(t *testing.T) {
	v := corestr.NewValidValue("hello")
	c := v.Clone()
	if c.Value != "hello" {
		t.Error("expected hello")
	}
}

func Test_C27_35_ValidValue_Clone_Nil(t *testing.T) {
	var v *corestr.ValidValue
	if v.Clone() != nil {
		t.Error("expected nil")
	}
}

func Test_C27_36_ValidValue_String(t *testing.T) {
	v := corestr.NewValidValue("hi")
	if v.String() != "hi" {
		t.Error("expected hi")
	}
}

func Test_C27_37_ValidValue_String_Nil(t *testing.T) {
	var v *corestr.ValidValue
	if v.String() != "" {
		t.Error("expected empty")
	}
}

func Test_C27_38_ValidValue_FullString(t *testing.T) {
	v := corestr.NewValidValue("hi")
	if v.FullString() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C27_39_ValidValue_FullString_Nil(t *testing.T) {
	var v *corestr.ValidValue
	if v.FullString() != "" {
		t.Error("expected empty")
	}
}

func Test_C27_40_ValidValue_Clear(t *testing.T) {
	v := corestr.NewValidValue("hi")
	v.Clear()
	if v.Value != "" || v.IsValid {
		t.Error("expected cleared")
	}
}

func Test_C27_41_ValidValue_Clear_Nil(t *testing.T) {
	var v *corestr.ValidValue
	v.Clear() // no panic
}

func Test_C27_42_ValidValue_Dispose(t *testing.T) {
	v := corestr.NewValidValue("hi")
	v.Dispose()
}

func Test_C27_43_ValidValue_Dispose_Nil(t *testing.T) {
	var v *corestr.ValidValue
	v.Dispose()
}

func Test_C27_44_ValidValue_Json(t *testing.T) {
	v := corestr.NewValidValue("hi")
	j := v.Json()
	_ = j
}

func Test_C27_45_ValidValue_JsonPtr(t *testing.T) {
	v := corestr.NewValidValue("hi")
	if v.JsonPtr() == nil {
		t.Error("expected non-nil")
	}
}

func Test_C27_46_ValidValue_Serialize(t *testing.T) {
	v := corestr.NewValidValue("hi")
	b, err := v.Serialize()
	if err != nil || len(b) == 0 {
		t.Error("expected bytes")
	}
}

func Test_C27_47_ValidValue_ParseInjectUsingJson(t *testing.T) {
	v := corestr.NewValidValue("hi")
	jp := v.JsonPtr()
	v2 := &corestr.ValidValue{}
	_, err := v2.ParseInjectUsingJson(jp)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

func Test_C27_48_ValidValue_Deserialize(t *testing.T) {
	v := corestr.NewValidValue("hi")
	var target corestr.ValidValue
	err := v.Deserialize(&target)
	if err != nil {
		t.Errorf("unexpected: %v", err)
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ValidValues
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_49_ValidValues_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	if !vv.IsEmpty() || vv.Length() != 0 {
		t.Error("expected empty")
	}
}

func Test_C27_50_ValidValues_NewValidValues(t *testing.T) {
	vv := corestr.NewValidValues(4)
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_51_ValidValues_NewValidValuesUsingValues(t *testing.T) {
	vv := corestr.NewValidValuesUsingValues(
		corestr.ValidValue{Value: "a", IsValid: true},
		corestr.ValidValue{Value: "b", IsValid: true},
	)
	if vv.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_52_ValidValues_NewValidValuesUsingValues_Empty(t *testing.T) {
	vv := corestr.NewValidValuesUsingValues()
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_53_ValidValues_Add(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("hello")
	if vv.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_54_ValidValues_AddFull(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddFull(false, "val", "msg")
	if vv.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_55_ValidValues_Count(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	if vv.Count() != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_56_ValidValues_HasAnyItem(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	if !vv.HasAnyItem() {
		t.Error("expected true")
	}
}

func Test_C27_57_ValidValues_LastIndex(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	vv.Add("b")
	if vv.LastIndex() != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_58_ValidValues_HasIndex(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	if !vv.HasIndex(0) {
		t.Error("expected true")
	}
	if vv.HasIndex(5) {
		t.Error("expected false")
	}
}

func Test_C27_59_ValidValues_SafeValueAt(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("hello")
	if vv.SafeValueAt(0) != "hello" {
		t.Error("expected hello")
	}
	if vv.SafeValueAt(99) != "" {
		t.Error("expected empty")
	}
}

func Test_C27_60_ValidValues_SafeValidValueAt(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("hello")
	if vv.SafeValidValueAt(0) != "hello" {
		t.Error("expected hello")
	}
}

func Test_C27_61_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	vv.Add("b")
	vals := vv.SafeValuesAtIndexes(0, 1)
	if len(vals) != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_62_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	vals := vv.SafeValidValuesAtIndexes(0)
	if len(vals) != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_63_ValidValues_Strings(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	s := vv.Strings()
	if len(s) != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_64_ValidValues_Strings_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	s := vv.Strings()
	if len(s) != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_65_ValidValues_FullStrings(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	s := vv.FullStrings()
	if len(s) != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_66_ValidValues_FullStrings_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	s := vv.FullStrings()
	if len(s) != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_67_ValidValues_String(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	if vv.String() == "" {
		t.Error("expected non-empty")
	}
}

func Test_C27_68_ValidValues_Find(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	vv.Add("b")
	found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, v.Value == "b", false
	})
	if len(found) != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_69_ValidValues_Find_Empty(t *testing.T) {
	vv := corestr.EmptyValidValues()
	found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
		return v, true, false
	})
	if len(found) != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_70_ValidValues_ConcatNew(t *testing.T) {
	vv1 := corestr.NewValidValues(4)
	vv1.Add("a")
	vv2 := corestr.NewValidValues(4)
	vv2.Add("b")
	result := vv1.ConcatNew(false, vv2)
	if result.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_71_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	result := vv.ConcatNew(true)
	if result.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_72_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	result := vv.ConcatNew(false)
	if result != vv {
		t.Error("expected same pointer")
	}
}

func Test_C27_73_ValidValues_AddValidValues(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Add("a")
	vv2 := corestr.NewValidValues(4)
	vv2.Add("b")
	vv.AddValidValues(vv2)
	if vv.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_74_ValidValues_AddValidValues_Nil(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddValidValues(nil)
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_75_ValidValues_Adds(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Adds(corestr.ValidValue{Value: "a"}, corestr.ValidValue{Value: "b"})
	if vv.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_76_ValidValues_Adds_Empty(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.Adds()
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_77_ValidValues_AddsPtr(t *testing.T) {
	vv := corestr.NewValidValues(4)
	v := corestr.NewValidValue("a")
	vv.AddsPtr(v)
	if vv.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_78_ValidValues_AddsPtr_Empty(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddsPtr()
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_79_ValidValues_AddHashsetMap(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddHashsetMap(map[string]bool{"a": true, "b": false})
	if vv.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_80_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddHashsetMap(nil)
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_81_ValidValues_AddHashset(t *testing.T) {
	hs := corestr.New.Hashset.Items("a", "b")
	vv := corestr.NewValidValues(4)
	vv.AddHashset(hs)
	if vv.Length() != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_82_ValidValues_AddHashset_Nil(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddHashset(nil)
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_83_ValidValues_Hashmap(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddFull(true, "k", "v")
	hm := vv.Hashmap()
	if hm.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_84_ValidValues_Map(t *testing.T) {
	vv := corestr.NewValidValues(4)
	vv.AddFull(true, "k", "v")
	m := vv.Map()
	if len(m) != 1 {
		t.Error("expected 1")
	}
}

func Test_C27_85_ValidValues_Length_Nil(t *testing.T) {
	var vv *corestr.ValidValues
	if vv.Length() != 0 {
		t.Error("expected 0")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_86_ValueStatus_InvalidNoMessage(t *testing.T) {
	vs := corestr.InvalidValueStatusNoMessage()
	if vs.ValueValid.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_87_ValueStatus_Invalid(t *testing.T) {
	vs := corestr.InvalidValueStatus("err")
	if vs.ValueValid.Message != "err" {
		t.Error("expected err")
	}
}

func Test_C27_88_ValueStatus_Clone(t *testing.T) {
	vs := corestr.InvalidValueStatus("err")
	c := vs.Clone()
	if c.ValueValid.Message != "err" {
		t.Error("expected err")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_89_LeftRight_NewLeftRight(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if lr.Left != "a" || lr.Right != "b" || !lr.IsValid {
		t.Error("expected a, b, valid")
	}
}

func Test_C27_90_LeftRight_InvalidLeftRight(t *testing.T) {
	lr := corestr.InvalidLeftRight("err")
	if lr.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_91_LeftRight_InvalidLeftRightNoMessage(t *testing.T) {
	lr := corestr.InvalidLeftRightNoMessage()
	if lr.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_92_LeftRight_LeftRightUsingSlice(t *testing.T) {
	lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
	if lr.Left != "a" || lr.Right != "b" {
		t.Error("expected a, b")
	}
}

func Test_C27_93_LeftRight_LeftRightUsingSlice_Empty(t *testing.T) {
	lr := corestr.LeftRightUsingSlice([]string{})
	if lr.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_94_LeftRight_LeftRightUsingSlice_Single(t *testing.T) {
	lr := corestr.LeftRightUsingSlice([]string{"a"})
	if lr.Left != "a" || lr.Right != "" {
		t.Error("expected a, empty")
	}
}

func Test_C27_95_LeftRight_LeftRightUsingSlicePtr(t *testing.T) {
	lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
	if lr.Left != "a" {
		t.Error("expected a")
	}
}

func Test_C27_96_LeftRight_LeftRightTrimmedUsingSlice(t *testing.T) {
	lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
	if lr.Left != "a" || lr.Right != "b" {
		t.Error("expected trimmed")
	}
}

func Test_C27_97_LeftRight_LeftRightTrimmedUsingSlice_Nil(t *testing.T) {
	lr := corestr.LeftRightTrimmedUsingSlice(nil)
	if lr.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_98_LeftRight_LeftRightTrimmedUsingSlice_Empty(t *testing.T) {
	lr := corestr.LeftRightTrimmedUsingSlice([]string{})
	if lr.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_99_LeftRight_LeftRightTrimmedUsingSlice_Single(t *testing.T) {
	lr := corestr.LeftRightTrimmedUsingSlice([]string{" a "})
	if lr.Left != "a" {
		t.Error("expected a")
	}
}

func Test_C27_100_LeftRight_LeftBytes(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if string(lr.LeftBytes()) != "a" {
		t.Error("expected a")
	}
}

func Test_C27_101_LeftRight_RightBytes(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if string(lr.RightBytes()) != "b" {
		t.Error("expected b")
	}
}

func Test_C27_102_LeftRight_LeftTrim(t *testing.T) {
	lr := corestr.NewLeftRight(" a ", "b")
	if lr.LeftTrim() != "a" {
		t.Error("expected a")
	}
}

func Test_C27_103_LeftRight_RightTrim(t *testing.T) {
	lr := corestr.NewLeftRight("a", " b ")
	if lr.RightTrim() != "b" {
		t.Error("expected b")
	}
}

func Test_C27_104_LeftRight_IsLeftEmpty(t *testing.T) {
	lr := corestr.NewLeftRight("", "b")
	if !lr.IsLeftEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_105_LeftRight_IsRightEmpty(t *testing.T) {
	lr := corestr.NewLeftRight("a", "")
	if !lr.IsRightEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_106_LeftRight_IsLeftWhitespace(t *testing.T) {
	lr := corestr.NewLeftRight("   ", "b")
	if !lr.IsLeftWhitespace() {
		t.Error("expected true")
	}
}

func Test_C27_107_LeftRight_IsRightWhitespace(t *testing.T) {
	lr := corestr.NewLeftRight("a", "   ")
	if !lr.IsRightWhitespace() {
		t.Error("expected true")
	}
}

func Test_C27_108_LeftRight_HasValidNonEmptyLeft(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.HasValidNonEmptyLeft() {
		t.Error("expected true")
	}
}

func Test_C27_109_LeftRight_HasValidNonEmptyRight(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.HasValidNonEmptyRight() {
		t.Error("expected true")
	}
}

func Test_C27_110_LeftRight_HasValidNonWhitespaceLeft(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.HasValidNonWhitespaceLeft() {
		t.Error("expected true")
	}
}

func Test_C27_111_LeftRight_HasValidNonWhitespaceRight(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.HasValidNonWhitespaceRight() {
		t.Error("expected true")
	}
}

func Test_C27_112_LeftRight_HasSafeNonEmpty(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.HasSafeNonEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_113_LeftRight_NonPtr(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	np := lr.NonPtr()
	if np.Left != "a" {
		t.Error("expected a")
	}
}

func Test_C27_114_LeftRight_Ptr(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if lr.Ptr() != lr {
		t.Error("expected same")
	}
}

func Test_C27_115_LeftRight_IsLeftRegexMatch(t *testing.T) {
	lr := corestr.NewLeftRight("abc123", "b")
	re := regexp.MustCompile(`\d+`)
	if !lr.IsLeftRegexMatch(re) {
		t.Error("expected true")
	}
	if lr.IsLeftRegexMatch(nil) {
		t.Error("expected false")
	}
}

func Test_C27_116_LeftRight_IsRightRegexMatch(t *testing.T) {
	lr := corestr.NewLeftRight("a", "abc123")
	re := regexp.MustCompile(`\d+`)
	if !lr.IsRightRegexMatch(re) {
		t.Error("expected true")
	}
	if lr.IsRightRegexMatch(nil) {
		t.Error("expected false")
	}
}

func Test_C27_117_LeftRight_IsLeft(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.IsLeft("a") {
		t.Error("expected true")
	}
}

func Test_C27_118_LeftRight_IsRight(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.IsRight("b") {
		t.Error("expected true")
	}
}

func Test_C27_119_LeftRight_Is(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if !lr.Is("a", "b") {
		t.Error("expected true")
	}
}

func Test_C27_120_LeftRight_IsEqual(t *testing.T) {
	lr1 := corestr.NewLeftRight("a", "b")
	lr2 := corestr.NewLeftRight("a", "b")
	if !lr1.IsEqual(lr2) {
		t.Error("expected true")
	}
}

func Test_C27_121_LeftRight_IsEqual_BothNil(t *testing.T) {
	var lr1, lr2 *corestr.LeftRight
	if !lr1.IsEqual(lr2) {
		t.Error("expected true")
	}
}

func Test_C27_122_LeftRight_IsEqual_OneNil(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	if lr.IsEqual(nil) {
		t.Error("expected false")
	}
}

func Test_C27_123_LeftRight_Clone(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	c := lr.Clone()
	if c.Left != "a" || c.Right != "b" {
		t.Error("expected a, b")
	}
}

func Test_C27_124_LeftRight_Clear(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	lr.Clear()
}

func Test_C27_125_LeftRight_Clear_Nil(t *testing.T) {
	var lr *corestr.LeftRight
	lr.Clear()
}

func Test_C27_126_LeftRight_Dispose(t *testing.T) {
	lr := corestr.NewLeftRight("a", "b")
	lr.Dispose()
}

func Test_C27_127_LeftRight_Dispose_Nil(t *testing.T) {
	var lr *corestr.LeftRight
	lr.Dispose()
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRightFromSplit
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_128_LeftRightFromSplit(t *testing.T) {
	lr := corestr.LeftRightFromSplit("key=value", "=")
	if lr.Left != "key" || lr.Right != "value" {
		t.Error("expected key, value")
	}
}

func Test_C27_129_LeftRightFromSplitTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
	if lr.Left != "key" || lr.Right != "value" {
		t.Error("expected trimmed key, value")
	}
}

func Test_C27_130_LeftRightFromSplitFull(t *testing.T) {
	lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
	if lr.Left != "a" || lr.Right != "b:c:d" {
		t.Error("expected a, b:c:d")
	}
}

func Test_C27_131_LeftRightFromSplitFullTrimmed(t *testing.T) {
	lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
	if lr.Left != "a" || lr.Right != "b : c" {
		t.Error("expected trimmed")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_132_LeftMiddleRight_New(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
		t.Error("expected a, b, c")
	}
}

func Test_C27_133_LeftMiddleRight_Invalid(t *testing.T) {
	lmr := corestr.InvalidLeftMiddleRight("err")
	if lmr.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_134_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	lmr := corestr.InvalidLeftMiddleRightNoMessage()
	if lmr.IsValid {
		t.Error("expected invalid")
	}
}

func Test_C27_135_LeftMiddleRight_Bytes(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	if string(lmr.LeftBytes()) != "a" || string(lmr.MiddleBytes()) != "b" || string(lmr.RightBytes()) != "c" {
		t.Error("expected a, b, c")
	}
}

func Test_C27_136_LeftMiddleRight_Trims(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
	if lmr.LeftTrim() != "a" || lmr.MiddleTrim() != "b" || lmr.RightTrim() != "c" {
		t.Error("expected trimmed")
	}
}

func Test_C27_137_LeftMiddleRight_IsEmpty(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("", "", "")
	if !lmr.IsLeftEmpty() || !lmr.IsMiddleEmpty() || !lmr.IsRightEmpty() {
		t.Error("expected all empty")
	}
}

func Test_C27_138_LeftMiddleRight_IsWhitespace(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("  ", "  ", "  ")
	if !lmr.IsLeftWhitespace() || !lmr.IsMiddleWhitespace() || !lmr.IsRightWhitespace() {
		t.Error("expected all whitespace")
	}
}

func Test_C27_139_LeftMiddleRight_HasValid(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	if !lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyMiddle() || !lmr.HasValidNonEmptyRight() {
		t.Error("expected all valid non-empty")
	}
	if !lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceMiddle() || !lmr.HasValidNonWhitespaceRight() {
		t.Error("expected all non-whitespace")
	}
}

func Test_C27_140_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	if !lmr.HasSafeNonEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_141_LeftMiddleRight_IsAll(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	if !lmr.IsAll("a", "b", "c") {
		t.Error("expected true")
	}
}

func Test_C27_142_LeftMiddleRight_Is(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	if !lmr.Is("a", "c") {
		t.Error("expected true")
	}
}

func Test_C27_143_LeftMiddleRight_Clone(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	c := lmr.Clone()
	if c.Left != "a" || c.Middle != "b" || c.Right != "c" {
		t.Error("expected a, b, c")
	}
}

func Test_C27_144_LeftMiddleRight_ToLeftRight(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lr := lmr.ToLeftRight()
	if lr.Left != "a" || lr.Right != "c" {
		t.Error("expected a, c")
	}
}

func Test_C27_145_LeftMiddleRight_Clear(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lmr.Clear()
}

func Test_C27_146_LeftMiddleRight_Clear_Nil(t *testing.T) {
	var lmr *corestr.LeftMiddleRight
	lmr.Clear()
}

func Test_C27_147_LeftMiddleRight_Dispose(t *testing.T) {
	lmr := corestr.NewLeftMiddleRight("a", "b", "c")
	lmr.Dispose()
}

func Test_C27_148_LeftMiddleRight_Dispose_Nil(t *testing.T) {
	var lmr *corestr.LeftMiddleRight
	lmr.Dispose()
}

// ═══════════════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_149_LeftMiddleRightFromSplit(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
	if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
		t.Error("expected a, b, c")
	}
}

func Test_C27_150_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
	if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
		t.Error("expected trimmed a, b, c")
	}
}

func Test_C27_151_LeftMiddleRightFromSplitN(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
	if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c:d:e" {
		t.Error("expected a, b, c:d:e")
	}
}

func Test_C27_152_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
	if lmr.Left != "a" || lmr.Middle != "b" {
		t.Error("expected trimmed")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_153_TextWithLineNumber_HasLineNumber(t *testing.T) {
	tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
	if !tl.HasLineNumber() {
		t.Error("expected true")
	}
}

func Test_C27_154_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
	if !tl.IsInvalidLineNumber() {
		t.Error("expected true")
	}
}

func Test_C27_155_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	var tl *corestr.TextWithLineNumber
	if !tl.IsInvalidLineNumber() {
		t.Error("expected true")
	}
}

func Test_C27_156_TextWithLineNumber_Length(t *testing.T) {
	tl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
	if tl.Length() != 5 {
		t.Error("expected 5")
	}
}

func Test_C27_157_TextWithLineNumber_Length_Nil(t *testing.T) {
	var tl *corestr.TextWithLineNumber
	if tl.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_158_TextWithLineNumber_IsEmpty(t *testing.T) {
	tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	if !tl.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_C27_159_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	var tl *corestr.TextWithLineNumber
	if !tl.IsEmpty() {
		t.Error("expected empty")
	}
}

func Test_C27_160_TextWithLineNumber_IsEmptyText(t *testing.T) {
	tl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
	if !tl.IsEmptyText() {
		t.Error("expected empty text")
	}
}

func Test_C27_161_TextWithLineNumber_IsEmptyText_Nil(t *testing.T) {
	var tl *corestr.TextWithLineNumber
	if !tl.IsEmptyText() {
		t.Error("expected true")
	}
}

func Test_C27_162_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
	if !tl.IsEmptyTextLineBoth() {
		t.Error("expected true")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CloneSlice, CloneSliceIf
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_163_CloneSlice_Basic(t *testing.T) {
	s := corestr.CloneSlice([]string{"a", "b"})
	if len(s) != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_164_CloneSlice_Empty(t *testing.T) {
	s := corestr.CloneSlice(nil)
	if len(s) != 0 {
		t.Error("expected 0")
	}
}

func Test_C27_165_CloneSliceIf_Clone(t *testing.T) {
	s := corestr.CloneSliceIf(true, "a", "b")
	if len(s) != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_166_CloneSliceIf_NoClone(t *testing.T) {
	s := corestr.CloneSliceIf(false, "a", "b")
	if len(s) != 2 {
		t.Error("expected 2")
	}
}

func Test_C27_167_CloneSliceIf_Empty(t *testing.T) {
	s := corestr.CloneSliceIf(true)
	if len(s) != 0 {
		t.Error("expected 0")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_168_AnyToString_WithFieldName(t *testing.T) {
	s := corestr.AnyToString(true, 42)
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C27_169_AnyToString_WithoutFieldName(t *testing.T) {
	s := corestr.AnyToString(false, 42)
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C27_170_AnyToString_Empty(t *testing.T) {
	s := corestr.AnyToString(false, "")
	if s != "" {
		t.Error("expected empty")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_171_AllIndividualStringsOfStringsLength(t *testing.T) {
	items := [][]string{{"a", "b"}, {"c"}}
	l := corestr.AllIndividualStringsOfStringsLength(&items)
	if l != 3 {
		t.Error("expected 3")
	}
}

func Test_C27_172_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	l := corestr.AllIndividualStringsOfStringsLength(nil)
	if l != 0 {
		t.Error("expected 0")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// AllIndividualsLengthOfSimpleSlices
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_173_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	s1 := corestr.New.SimpleSlice.Create([]string{"a", "b"})
	s2 := corestr.New.SimpleSlice.Create([]string{"c"})
	l := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)
	if l != 3 {
		t.Error("expected 3")
	}
}

func Test_C27_174_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	l := corestr.AllIndividualsLengthOfSimpleSlices()
	if l != 0 {
		t.Error("expected 0")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// Utils
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_175_Utils_WrapDoubleIfMissing(t *testing.T) {
	u := corestr.StringUtils
	if u.WrapDoubleIfMissing("hello") != `"hello"` {
		t.Error("expected wrapped")
	}
	if u.WrapDoubleIfMissing(`"hello"`) != `"hello"` {
		t.Error("expected already wrapped")
	}
	if u.WrapDoubleIfMissing("") != `""` {
		t.Error("expected empty wrapped")
	}
}

func Test_C27_176_Utils_WrapSingleIfMissing(t *testing.T) {
	u := corestr.StringUtils
	if u.WrapSingleIfMissing("hello") != "'hello'" {
		t.Error("expected wrapped")
	}
	if u.WrapSingleIfMissing("'hello'") != "'hello'" {
		t.Error("expected already wrapped")
	}
	if u.WrapSingleIfMissing("") != "''" {
		t.Error("expected empty wrapped")
	}
}

func Test_C27_177_Utils_WrapDouble(t *testing.T) {
	u := corestr.StringUtils
	if u.WrapDouble("hi") != `"hi"` {
		t.Error("expected wrapped")
	}
}

func Test_C27_178_Utils_WrapSingle(t *testing.T) {
	u := corestr.StringUtils
	if u.WrapSingle("hi") != "'hi'" {
		t.Error("expected wrapped")
	}
}

func Test_C27_179_Utils_WrapTilda(t *testing.T) {
	u := corestr.Utils
	if u.WrapTilda("hi") != "`hi`" {
		t.Error("expected wrapped")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValuePair — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_180_KeyValuePair_Basic(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	if kv.KeyName() != "k" || kv.VariableName() != "k" || kv.ValueString() != "v" {
		t.Error("expected k, v")
	}
}

func Test_C27_181_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	if !kv.IsVariableNameEqual("k") {
		t.Error("expected true")
	}
}

func Test_C27_182_KeyValuePair_IsValueEqual(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	if !kv.IsValueEqual("v") {
		t.Error("expected true")
	}
}

func Test_C27_183_KeyValuePair_ValueBool(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "true"}
	if !kv.ValueBool() {
		t.Error("expected true")
	}
	kv2 := corestr.KeyValuePair{Key: "k", Value: ""}
	if kv2.ValueBool() {
		t.Error("expected false")
	}
}

func Test_C27_184_KeyValuePair_ValueInt(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "42"}
	if kv.ValueInt(0) != 42 {
		t.Error("expected 42")
	}
}

func Test_C27_185_KeyValuePair_ValueDefInt(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "10"}
	if kv.ValueDefInt() != 10 {
		t.Error("expected 10")
	}
}

func Test_C27_186_KeyValuePair_ValueByte(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "100"}
	if kv.ValueByte(0) != 100 {
		t.Error("expected 100")
	}
}

func Test_C27_187_KeyValuePair_ValueDefByte(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "50"}
	if kv.ValueDefByte() != 50 {
		t.Error("expected 50")
	}
}

func Test_C27_188_KeyValuePair_ValueFloat64(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
	if kv.ValueFloat64(0) != 3.14 {
		t.Error("expected 3.14")
	}
}

func Test_C27_189_KeyValuePair_ValueDefFloat64(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "2.5"}
	if kv.ValueDefFloat64() != 2.5 {
		t.Error("expected 2.5")
	}
}

func Test_C27_190_KeyValuePair_ValueValid(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	vv := kv.ValueValid()
	if vv.Value != "v" || !vv.IsValid {
		t.Error("expected v, valid")
	}
}

func Test_C27_191_KeyValuePair_ValueValidOptions(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	vv := kv.ValueValidOptions(false, "msg")
	if vv.IsValid || vv.Message != "msg" {
		t.Error("expected false, msg")
	}
}

func Test_C27_192_KeyValuePair_IsKeyEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "", Value: "v"}
	if !kv.IsKeyEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_193_KeyValuePair_IsValueEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: ""}
	if !kv.IsValueEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_194_KeyValuePair_HasKey(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k"}
	if !kv.HasKey() {
		t.Error("expected true")
	}
}

func Test_C27_195_KeyValuePair_HasValue(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	if !kv.HasValue() {
		t.Error("expected true")
	}
}

func Test_C27_196_KeyValuePair_IsKeyValueEmpty(t *testing.T) {
	kv := corestr.KeyValuePair{}
	if !kv.IsKeyValueEmpty() {
		t.Error("expected true")
	}
}

func Test_C27_197_KeyValuePair_TrimKey(t *testing.T) {
	kv := corestr.KeyValuePair{Key: " k "}
	if kv.TrimKey() != "k" {
		t.Error("expected k")
	}
}

func Test_C27_198_KeyValuePair_TrimValue(t *testing.T) {
	kv := corestr.KeyValuePair{Value: " v "}
	if kv.TrimValue() != "v" {
		t.Error("expected v")
	}
}

func Test_C27_199_KeyValuePair_Is(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k", Value: "v"}
	if !kv.Is("k", "v") {
		t.Error("expected true")
	}
}

func Test_C27_200_KeyValuePair_IsKey(t *testing.T) {
	kv := corestr.KeyValuePair{Key: "k"}
	if !kv.IsKey("k") {
		t.Error("expected true")
	}
}
