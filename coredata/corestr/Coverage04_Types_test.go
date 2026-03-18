package corestr

import (
	"testing"
)

// ── ValidValue ──

func TestValidValue_NewFactories(t *testing.T) {
	v1 := NewValidValue("hello")
	if v1.Value != "hello" || !v1.IsValid {
		t.Fatal("unexpected")
	}
	v2 := NewValidValueEmpty()
	if v2.Value != "" || !v2.IsValid {
		t.Fatal("unexpected")
	}
	v3 := InvalidValidValue("msg")
	if v3.IsValid {
		t.Fatal("expected invalid")
	}
	v4 := InvalidValidValueNoMessage()
	if v4.IsValid {
		t.Fatal("expected invalid")
	}
	v5 := NewValidValueUsingAny(false, true, "hello")
	if v5.Value == "" {
		t.Fatal("expected non-empty")
	}
	v6 := NewValidValueUsingAnyAutoValid(false, "hello")
	_ = v6
}

func TestValidValue_Methods(t *testing.T) {
	v := NewValidValue("hello")
	if v.IsEmpty() || !v.HasValidNonEmpty() || !v.HasSafeNonEmpty() {
		t.Fatal("unexpected")
	}
	if v.IsWhitespace() || !v.HasValidNonWhitespace() {
		t.Fatal("unexpected")
	}
	if v.Trim() != "hello" {
		t.Fatal("unexpected")
	}
	if !v.Is("hello") || v.Is("world") {
		t.Fatal("unexpected")
	}
	if !v.IsAnyOf("hello", "world") || !v.IsAnyOf() {
		t.Fatal("unexpected")
	}
	if v.IsAnyOf("x") {
		t.Fatal("unexpected")
	}
	if !v.IsContains("hel") {
		t.Fatal("unexpected")
	}
	if !v.IsAnyContains("hel") || !v.IsAnyContains() {
		t.Fatal("unexpected")
	}
	if v.IsAnyContains("xyz") {
		t.Fatal("unexpected")
	}
	if !v.IsEqualNonSensitive("HELLO") {
		t.Fatal("unexpected")
	}
}

func TestValidValue_ValueConversions(t *testing.T) {
	v := &ValidValue{Value: "42", IsValid: true}
	if v.ValueBool() {
		// "42" is not a valid bool
	}
	v2 := &ValidValue{Value: "true", IsValid: true}
	if !v2.ValueBool() {
		t.Fatal("expected true")
	}
	if v.ValueInt(0) != 42 || v.ValueDefInt() != 42 {
		t.Fatal("unexpected")
	}
	if v.ValueByte(0) != 42 || v.ValueDefByte() != 42 {
		t.Fatal("unexpected")
	}
	vf := &ValidValue{Value: "3.14", IsValid: true}
	if vf.ValueFloat64(0) == 0 || vf.ValueDefFloat64() == 0 {
		t.Fatal("unexpected")
	}
	// error cases
	bad := &ValidValue{Value: "abc", IsValid: true}
	if bad.ValueInt(99) != 99 || bad.ValueDefInt() != 0 {
		t.Fatal("unexpected")
	}
	if bad.ValueByte(99) != 0 {
		// error path
	}
	if bad.ValueFloat64(1.0) != 1.0 {
		t.Fatal("unexpected")
	}
	// empty bool
	empty := &ValidValue{Value: "", IsValid: true}
	if empty.ValueBool() {
		t.Fatal("expected false")
	}
	// overflow byte
	big := &ValidValue{Value: "999", IsValid: true}
	_ = big.ValueByte(0)
	_ = big.ValueDefByte()
	// negative byte
	neg := &ValidValue{Value: "-1", IsValid: true}
	_ = neg.ValueByte(0)
}

func TestValidValue_BytesOnce(t *testing.T) {
	v := NewValidValue("hello")
	b1 := v.ValueBytesOnce()
	b2 := v.ValueBytesOnce() // cached
	if len(b1) != 5 || len(b2) != 5 {
		t.Fatal("unexpected")
	}
	_ = v.ValueBytesOncePtr()
}

func TestValidValue_Regex(t *testing.T) {
	v := NewValidValue("hello123")
	if v.IsRegexMatches(nil) {
		t.Fatal("expected false for nil regexp")
	}
	if v.RegexFindString(nil) != "" {
		t.Fatal("expected empty")
	}
	items, has := v.RegexFindAllStringsWithFlag(nil, -1)
	if has || len(items) != 0 {
		t.Fatal("unexpected")
	}
	if len(v.RegexFindAllStrings(nil, -1)) != 0 {
		t.Fatal("expected empty")
	}
}

func TestValidValue_Split(t *testing.T) {
	v := NewValidValue("a,b,c")
	if len(v.Split(",")) != 3 {
		t.Fatal("expected 3")
	}
	_ = v.SplitNonEmpty(",")
	_ = v.SplitTrimNonWhitespace(",")
}

func TestValidValue_Clone(t *testing.T) {
	v := NewValidValue("hello")
	c := v.Clone()
	if c.Value != "hello" {
		t.Fatal("unexpected")
	}
	var nilV *ValidValue
	if nilV.Clone() != nil {
		t.Fatal("expected nil")
	}
}

func TestValidValue_String(t *testing.T) {
	v := NewValidValue("hello")
	if v.String() != "hello" {
		t.Fatal("unexpected")
	}
	if v.FullString() == "" {
		t.Fatal("expected non-empty")
	}
	var nilV *ValidValue
	if nilV.String() != "" || nilV.FullString() != "" {
		t.Fatal("expected empty")
	}
}

func TestValidValue_ClearDispose(t *testing.T) {
	v := NewValidValue("hello")
	v.Clear()
	if v.Value != "" {
		t.Fatal("expected empty")
	}
	v2 := NewValidValue("world")
	v2.Dispose()
	var nilV *ValidValue
	nilV.Clear()
	nilV.Dispose()
}

func TestValidValue_Json(t *testing.T) {
	v := ValidValue{Value: "hello", IsValid: true}
	_ = v.Json()
	_ = v.JsonPtr()
	_, _ = v.Serialize()
}

// ── ValidValues ──

func TestValidValues_Basic(t *testing.T) {
	vv := EmptyValidValues()
	if !vv.IsEmpty() || vv.HasAnyItem() || vv.Length() != 0 {
		t.Fatal("expected empty")
	}
	vv.Add("a").AddFull(true, "b", "msg")
	if vv.Length() != 2 {
		t.Fatal("expected 2")
	}
}

func TestValidValues_NilReceiver(t *testing.T) {
	var vv *ValidValues
	if vv.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestValidValues_NewFactories(t *testing.T) {
	vv := NewValidValuesUsingValues(ValidValue{Value: "a"})
	if vv.Length() != 1 {
		t.Fatal("expected 1")
	}
	empty := NewValidValuesUsingValues()
	if empty.Length() != 0 {
		t.Fatal("expected 0")
	}
	cap := NewValidValues(10)
	if cap.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestValidValues_SafeValues(t *testing.T) {
	vv := EmptyValidValues()
	vv.Add("a").Add("b")
	if vv.SafeValueAt(0) != "a" || vv.SafeValueAt(99) != "" {
		t.Fatal("unexpected")
	}
	if vv.SafeValidValueAt(0) != "a" || vv.SafeValidValueAt(99) != "" {
		t.Fatal("unexpected")
	}
	vals := vv.SafeValuesAtIndexes(0, 1)
	if len(vals) != 2 {
		t.Fatal("expected 2")
	}
	validVals := vv.SafeValidValuesAtIndexes(0)
	if len(validVals) != 1 {
		t.Fatal("expected 1")
	}
}

func TestValidValues_Strings(t *testing.T) {
	vv := EmptyValidValues()
	vv.Add("a")
	if len(vv.Strings()) != 1 {
		t.Fatal("expected 1")
	}
	if len(vv.FullStrings()) != 1 {
		t.Fatal("expected 1")
	}
	if vv.String() == "" {
		t.Fatal("expected non-empty")
	}
}

// ── ValueStatus ──

func TestValueStatus(t *testing.T) {
	vs := InvalidValueStatusNoMessage()
	if vs.ValueValid.IsValid {
		t.Fatal("expected invalid")
	}
	vs2 := InvalidValueStatus("msg")
	c := vs2.Clone()
	if c.ValueValid.Message != "msg" {
		t.Fatal("unexpected")
	}
}

// ── TextWithLineNumber ──

func TestTextWithLineNumber(t *testing.T) {
	tw := &TextWithLineNumber{LineNumber: 1, Text: "hello"}
	if !tw.HasLineNumber() || tw.IsInvalidLineNumber() {
		t.Fatal("unexpected")
	}
	if tw.Length() != 5 || tw.IsEmpty() || tw.IsEmptyText() {
		t.Fatal("unexpected")
	}
	if tw.IsEmptyTextLineBoth() {
		t.Fatal("unexpected")
	}

	var nilTw *TextWithLineNumber
	if nilTw.HasLineNumber() || !nilTw.IsInvalidLineNumber() {
		t.Fatal("unexpected")
	}
	if nilTw.Length() != 0 || !nilTw.IsEmpty() || !nilTw.IsEmptyText() {
		t.Fatal("unexpected")
	}
}

// ── KeyValuePair ──

func TestKeyValuePair(t *testing.T) {
	kv := KeyValuePair{Key: "k", Value: "v"}
	if kv.KeyName() != "k" || kv.VariableName() != "k" {
		t.Fatal("unexpected")
	}
	if kv.ValueString() != "v" {
		t.Fatal("unexpected")
	}
	if !kv.IsVariableNameEqual("k") || !kv.IsValueEqual("v") {
		t.Fatal("unexpected")
	}
	if kv.Compile() == "" || kv.String() == "" {
		t.Fatal("expected non-empty")
	}
	if kv.IsKeyEmpty() || kv.IsValueEmpty() || kv.IsKeyValueEmpty() {
		t.Fatal("unexpected")
	}
	if !kv.HasKey() || !kv.HasValue() {
		t.Fatal("unexpected")
	}
	if kv.TrimKey() != "k" || kv.TrimValue() != "v" {
		t.Fatal("unexpected")
	}
	if !kv.Is("k", "v") || !kv.IsKey("k") || !kv.IsVal("v") {
		t.Fatal("unexpected")
	}
	if kv.IsKeyValueAnyEmpty() {
		t.Fatal("unexpected")
	}
	_ = kv.FormatString("%s=%s")
	_ = kv.Json()
	_ = kv.JsonPtr()
	_, _ = kv.Serialize()
	_ = kv.SerializeMust()
	_ = kv.ValueValid()
	_ = kv.ValueValidOptions(true, "")

	// value conversions
	kv2 := KeyValuePair{Key: "k", Value: "42"}
	if kv2.ValueBool() {
		// 42 is not bool
	}
	if kv2.ValueInt(0) != 42 || kv2.ValueDefInt() != 42 {
		t.Fatal("unexpected")
	}
	_ = kv2.ValueByte(0)
	_ = kv2.ValueDefByte()
	_ = kv2.ValueFloat64(0)
	_ = kv2.ValueDefFloat64()

	kv.Clear()
	kv.Dispose()
	var nilKv *KeyValuePair
	nilKv.Clear()
	nilKv.Dispose()
}

// ── LeftRight ──

func TestLeftRight_Factories(t *testing.T) {
	lr := NewLeftRight("a", "b")
	if lr.Left != "a" || lr.Right != "b" || !lr.IsValid {
		t.Fatal("unexpected")
	}
	inv := InvalidLeftRight("msg")
	if inv.IsValid {
		t.Fatal("expected invalid")
	}
	inv2 := InvalidLeftRightNoMessage()
	if inv2.IsValid {
		t.Fatal("expected invalid")
	}
}

func TestLeftRight_Methods(t *testing.T) {
	lr := NewLeftRight("hello", "world")
	if string(lr.LeftBytes()) != "hello" || string(lr.RightBytes()) != "world" {
		t.Fatal("unexpected")
	}
	if lr.LeftTrim() != "hello" || lr.RightTrim() != "world" {
		t.Fatal("unexpected")
	}
	if lr.IsLeftEmpty() || lr.IsRightEmpty() {
		t.Fatal("unexpected")
	}
	if lr.IsLeftWhitespace() || lr.IsRightWhitespace() {
		t.Fatal("unexpected")
	}
	if !lr.HasValidNonEmptyLeft() || !lr.HasValidNonEmptyRight() {
		t.Fatal("unexpected")
	}
	if !lr.HasValidNonWhitespaceLeft() || !lr.HasValidNonWhitespaceRight() {
		t.Fatal("unexpected")
	}
	if !lr.HasSafeNonEmpty() {
		t.Fatal("unexpected")
	}
	if !lr.IsLeft("hello") || !lr.IsRight("world") || !lr.Is("hello", "world") {
		t.Fatal("unexpected")
	}
	_ = lr.NonPtr()
	_ = lr.Ptr()
	if lr.IsLeftRegexMatch(nil) || lr.IsRightRegexMatch(nil) {
		t.Fatal("expected false for nil")
	}
	c := lr.Clone()
	if c.Left != "hello" {
		t.Fatal("unexpected")
	}
	if !lr.IsEqual(NewLeftRight("hello", "world")) {
		t.Fatal("expected equal")
	}
	lr.Clear()
	lr.Dispose()
	var nilLr *LeftRight
	nilLr.Clear()
	nilLr.Dispose()
}

func TestLeftRight_FromSlice(t *testing.T) {
	lr := LeftRightUsingSlice([]string{"a", "b"})
	if lr.Left != "a" || lr.Right != "b" {
		t.Fatal("unexpected")
	}
	lr2 := LeftRightUsingSlice([]string{"a"})
	if lr2.Left != "a" || lr2.Right != "" {
		t.Fatal("unexpected")
	}
	lr3 := LeftRightUsingSlice(nil)
	if lr3.IsValid {
		t.Fatal("expected invalid")
	}
	lr4 := LeftRightUsingSlicePtr([]string{"a", "b"})
	_ = lr4
	lr5 := LeftRightUsingSlicePtr(nil)
	_ = lr5
	lr6 := LeftRightTrimmedUsingSlice(nil)
	_ = lr6
	lr7 := LeftRightTrimmedUsingSlice([]string{" a "})
	_ = lr7
	lr8 := LeftRightTrimmedUsingSlice([]string{" a ", " b "})
	if lr8.Left != "a" || lr8.Right != "b" {
		t.Fatal("expected trimmed")
	}
}

func TestLeftRight_FromSplit(t *testing.T) {
	lr := LeftRightFromSplit("key=val", "=")
	if lr.Left != "key" || lr.Right != "val" {
		t.Fatal("unexpected")
	}
	lr2 := LeftRightFromSplitTrimmed(" key = val ", "=")
	if lr2.Left != "key" || lr2.Right != "val" {
		t.Fatal("expected trimmed")
	}
	lr3 := LeftRightFromSplitFull("a:b:c", ":")
	if lr3.Left != "a" || lr3.Right != "b:c" {
		t.Fatal("unexpected")
	}
	lr4 := LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
	_ = lr4
}

// ── LeftMiddleRight ──

func TestLeftMiddleRight(t *testing.T) {
	lmr := NewLeftMiddleRight("a", "b", "c")
	if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
		t.Fatal("unexpected")
	}
	if lmr.IsLeftEmpty() || lmr.IsMiddleEmpty() || lmr.IsRightEmpty() {
		t.Fatal("unexpected")
	}
	if lmr.IsLeftWhitespace() || lmr.IsMiddleWhitespace() || lmr.IsRightWhitespace() {
		t.Fatal("unexpected")
	}
	if string(lmr.LeftBytes()) != "a" || string(lmr.MiddleBytes()) != "b" || string(lmr.RightBytes()) != "c" {
		t.Fatal("unexpected")
	}
	if lmr.LeftTrim() != "a" || lmr.MiddleTrim() != "b" || lmr.RightTrim() != "c" {
		t.Fatal("unexpected")
	}
	if !lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyMiddle() || !lmr.HasValidNonEmptyRight() {
		t.Fatal("unexpected")
	}
	if !lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceMiddle() || !lmr.HasValidNonWhitespaceRight() {
		t.Fatal("unexpected")
	}
	if !lmr.HasSafeNonEmpty() || !lmr.IsAll("a", "b", "c") || !lmr.Is("a", "c") {
		t.Fatal("unexpected")
	}
	c := lmr.Clone()
	if c.Left != "a" {
		t.Fatal("unexpected")
	}
	lr := lmr.ToLeftRight()
	if lr.Left != "a" || lr.Right != "c" {
		t.Fatal("unexpected")
	}
	lmr.Clear()
	lmr.Dispose()
	var nilLmr *LeftMiddleRight
	nilLmr.Clear()
	nilLmr.Dispose()

	inv := InvalidLeftMiddleRight("msg")
	if inv.IsValid {
		t.Fatal("expected invalid")
	}
	inv2 := InvalidLeftMiddleRightNoMessage()
	_ = inv2
}

func TestLeftMiddleRight_FromSplit(t *testing.T) {
	lmr := LeftMiddleRightFromSplit("a.b.c", ".")
	if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
		t.Fatal("unexpected")
	}
	lmr2 := LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
	_ = lmr2
	lmr3 := LeftMiddleRightFromSplitN("a:b:c:d", ":")
	if lmr3.Left != "a" || lmr3.Middle != "b" || lmr3.Right != "c:d" {
		t.Fatal("unexpected")
	}
	lmr4 := LeftMiddleRightFromSplitNTrimmed(" a : b : c ", ":")
	_ = lmr4
}

// ── KeyAnyValuePair ──

func TestKeyAnyValuePair(t *testing.T) {
	kv := &KeyAnyValuePair{Key: "k", Value: 42}
	if kv.KeyName() != "k" || kv.VariableName() != "k" {
		t.Fatal("unexpected")
	}
	if kv.ValueAny() != 42 {
		t.Fatal("unexpected")
	}
	if !kv.IsVariableNameEqual("k") {
		t.Fatal("unexpected")
	}
	if kv.IsValueNull() || !kv.HasNonNull() || !kv.HasValue() {
		t.Fatal("unexpected")
	}
	if kv.IsValueEmptyString() || kv.IsValueWhitespace() {
		t.Fatal("unexpected")
	}
	s := kv.ValueString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	// call again for cache
	s2 := kv.ValueString()
	if s2 != s {
		t.Fatal("cache miss")
	}
	if kv.Compile() == "" || kv.String() == "" {
		t.Fatal("expected non-empty")
	}
	_ = kv.SerializeMust()
	_ = kv.AsJsonContractsBinder()
	_ = kv.AsJsoner()
	_ = kv.AsJsonParseSelfInjector()

	kv.Clear()
	kv.Dispose()
	var nilKv *KeyAnyValuePair
	nilKv.Clear()
	nilKv.Dispose()
}

// ── HashmapDiff ──

func TestHashmapDiff(t *testing.T) {
	hd := HashmapDiff(map[string]string{"a": "1"})
	if hd.IsEmpty() || !hd.HasAnyItem() {
		t.Fatal("unexpected")
	}
	if hd.Length() != 1 || hd.LastIndex() != 0 {
		t.Fatal("unexpected")
	}
	_ = hd.AllKeysSorted()
	_ = hd.MapAnyItems()
	_ = hd.Raw()
	_ = hd.RawMapStringAnyDiff()

	right := map[string]string{"a": "2"}
	if !hd.HasAnyChanges(right) {
		t.Fatal("expected changes")
	}
	if hd.IsRawEqual(right) {
		t.Fatal("expected not equal")
	}
	_ = hd.HashmapDiffUsingRaw(right)
	_ = hd.DiffRaw(right)
	_ = hd.DiffJsonMessage(right)
	_ = hd.ToStringsSliceOfDiffMap(map[string]string{"a": "diff"})
	_ = hd.ShouldDiffMessage("title", right)
	_ = hd.LogShouldDiffMessage("title", right)
	_, _ = hd.Serialize()
	_ = hd.Deserialize(&map[string]string{})

	// nil
	var nilHd *HashmapDiff
	if nilHd.Length() != 0 {
		t.Fatal("expected 0")
	}
	_ = nilHd.Raw()
	_ = nilHd.MapAnyItems()
}
