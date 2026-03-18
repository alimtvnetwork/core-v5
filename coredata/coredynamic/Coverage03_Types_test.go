package coredynamic

import (
	"reflect"
	"testing"
)

// ── TypeStatus ──

func TestTypeStatus(t *testing.T) {
	ts := TypeSameStatus("hello", "world")
	if !ts.IsSame { t.Fatal("expected same type") }
	if ts.IsNotSame() || ts.IsNotEqualTypes() { t.Fatal("unexpected") }
	_ = ts.IsAnyPointer()
	_ = ts.IsBothPointer()
	_ = ts.LeftName()
	_ = ts.RightName()
	_ = ts.LeftFullName()
	_ = ts.RightFullName()
	_ = ts.IsSameRegardlessPointer()
	_ = ts.NotEqualSrcDestinationMessage()
	_ = ts.NotMatchMessage("l", "r")
	_ = ts.NotMatchErr("l", "r")
	_ = ts.ValidationError()
	_ = ts.NotEqualSrcDestinationErr()
	ts.MustBeSame()
	ts.SrcDestinationMustBeSame()

	ts2 := TypeSameStatus(nil, nil)
	_ = ts2.IsValid()
	_ = ts2.IsInvalid()
	_ = ts2.LeftName()
	_ = ts2.RightName()
	_ = ts2.LeftFullName()
	_ = ts2.RightFullName()

	var nilTs *TypeStatus
	if nilTs.IsValid() { t.Fatal("expected invalid") }
	if !nilTs.IsInvalid() { t.Fatal("expected invalid") }
}

func TestTypeStatus_IsEqual(t *testing.T) {
	ts1 := TypeSameStatus("a", "b")
	ts2 := TypeSameStatus("a", "b")
	if !ts1.IsEqual(&ts2) { t.Fatal("expected equal") }
	var nilTs *TypeStatus
	if !nilTs.IsEqual(nil) { t.Fatal("expected equal") }
}

func TestTypeStatus_NonPointer(t *testing.T) {
	s := "hello"
	ts := TypeSameStatus(&s, &s)
	_ = ts.NonPointerLeft()
	_ = ts.NonPointerRight()
}

// ── KeyVal ──

func TestKeyVal(t *testing.T) {
	kv := KeyVal{Key: "k", Value: "v"}
	_ = kv.KeyDynamic()
	_ = kv.ValueDynamic()
	_ = kv.KeyDynamicPtr()
	_ = kv.ValueDynamicPtr()
	if kv.IsKeyNull() || kv.IsValueNull() { t.Fatal("unexpected") }
	_ = kv.String()
	_ = kv.KeyString()
	_ = kv.ValueString()
	_ = kv.ValueReflectValue()
	_ = kv.ValueInt()
	_ = kv.ValueUInt()
	_ = kv.ValueStrings()
	_ = kv.ValueBool()
	_ = kv.ValueInt64()
	_ = kv.ValueNullErr()
	_ = kv.KeyNullErr()
	_ = kv.JsonModel()
	_ = kv.JsonModelAny()
	_ = kv.Json()
	_ = kv.JsonPtr()
	_, _ = kv.Serialize()

	var nilKv *KeyVal
	if nilKv.KeyString() != "" { t.Fatal("expected empty") }
	if nilKv.ValueString() != "" { t.Fatal("expected empty") }
	if nilKv.ValueNullErr() == nil { t.Fatal("expected error") }
	if nilKv.KeyNullErr() == nil { t.Fatal("expected error") }
}

// ── KeyValCollection ──

func TestKeyValCollection(t *testing.T) {
	kvc := EmptyKeyValCollection()
	if !kvc.IsEmpty() || kvc.HasAnyItem() { t.Fatal("expected empty") }
	kvc.Add(KeyVal{Key: "a", Value: 1})
	kvc.AddPtr(&KeyVal{Key: "b", Value: 2})
	kvc.AddPtr(nil)
	kvc.AddMany(KeyVal{Key: "c", Value: 3})
	kvc.AddMany()
	kvc.AddManyPtr(&KeyVal{Key: "d", Value: 4}, nil)
	kvc.AddManyPtr()
	if kvc.Length() != 4 { t.Fatal("expected 4") }
	_ = kvc.Items()
	_ = kvc.AllKeys()
	_ = kvc.AllKeysSorted()
	_ = kvc.AllValues()
	_ = kvc.String()
	_ = kvc.MapAnyItems()
	_, _ = kvc.JsonMapResults()
	_ = kvc.JsonResultsCollection()
	_ = kvc.JsonResultsPtrCollection()
	_ = kvc.GetPagesSize(2)
	_ = kvc.JsonModel()
	_ = kvc.JsonModelAny()
	_ = kvc.Json()
	_ = kvc.JsonPtr()
	_, _ = kvc.Serialize()
	_, _ = kvc.JsonString()
	_ = kvc.JsonStringMust()
	_ = kvc.Clone()
	_ = kvc.ClonePtr()
	_ = kvc.NonPtr()
	_ = kvc.Ptr()
	var nilKvc *KeyValCollection
	if nilKvc.Length() != 0 { t.Fatal("expected 0") }
	if nilKvc.Items() != nil { t.Fatal("expected nil") }
	if nilKvc.ClonePtr() != nil { t.Fatal("expected nil") }
	if nilKvc.String() != "" { t.Fatal("expected empty") }
}

// ── MapAnyItems ──

func TestMapAnyItems_Basic(t *testing.T) {
	m := EmptyMapAnyItems()
	if !m.IsEmpty() || m.HasAnyItem() { t.Fatal("expected empty") }
	m.Add("a", 1)
	m.Set("b", 2)
	if m.Length() != 2 { t.Fatal("expected 2") }
	if !m.HasKey("a") || m.HasKey("z") { t.Fatal("unexpected") }
	_ = m.GetValue("a")
	v, has := m.Get("a")
	if !has || v != 1 { t.Fatal("unexpected") }
	var nilM *MapAnyItems
	if nilM.Length() != 0 { t.Fatal("expected 0") }
	if nilM.HasKey("a") { t.Fatal("expected false") }
}

func TestMapAnyItems_AllKeys(t *testing.T) {
	m := NewMapAnyItemsUsingItems(map[string]any{"b": 2, "a": 1})
	_ = m.AllKeysSorted()
	_ = m.AllKeys()
	_ = m.AllValues()
}

// ── MapAnyItemDiff ──

func TestMapAnyItemDiff(t *testing.T) {
	d := MapAnyItemDiff(map[string]any{"a": 1})
	if d.IsEmpty() || !d.HasAnyItem() { t.Fatal("unexpected") }
	if d.Length() != 1 { t.Fatal("expected 1") }
	_ = d.AllKeysSorted()
	_ = d.Raw()
	_ = d.MapAnyItems()
	_ = d.RawMapDiffer()
	right := map[string]any{"a": 2}
	_ = d.IsRawEqual(true, right)
	_ = d.HasAnyChanges(true, right)
	_ = d.HashmapDiffUsingRaw(true, right)
	_ = d.DiffRaw(true, right)
	_ = d.DiffJsonMessage(true, right)
	_ = d.ToStringsSliceOfDiffMap(right)
	_ = d.ShouldDiffMessage(true, "title", right)
	_ = d.LogShouldDiffMessage(true, "title", right)
	_ = d.Json()
	_ = d.JsonPtr()
	_ = d.PrettyJsonString()
	_ = d.Clear()
	var nilD *MapAnyItemDiff
	if nilD.Length() != 0 { t.Fatal("expected 0") }
	_ = nilD.Raw()
}

// ── LeftRight ──

func TestLeftRight_Dynamic(t *testing.T) {
	lr := &LeftRight{Left: "a", Right: "b"}
	if lr.IsEmpty() || !lr.HasAnyItem() { t.Fatal("unexpected") }
	if !lr.HasLeft() || !lr.HasRight() { t.Fatal("unexpected") }
	if lr.IsLeftEmpty() || lr.IsRightEmpty() { t.Fatal("unexpected") }
	_ = lr.LeftToDynamic()
	_ = lr.RightToDynamic()
	_ = lr.DeserializeLeft()
	_ = lr.DeserializeRight()
	_ = lr.TypeStatus()
	var nilLr *LeftRight
	if !nilLr.IsEmpty() { t.Fatal("expected empty") }
	if nilLr.HasLeft() || nilLr.HasRight() { t.Fatal("unexpected") }
}

// ── CastedResult ──

func TestCastedResult(t *testing.T) {
	cr := CastTo(false, "hello", reflect.TypeOf(""))
	if cr.IsInvalid() || cr.IsNotNull() == false { t.Fatal("unexpected") }
	if cr.HasError() || cr.HasAnyIssues() { t.Fatal("unexpected") }
	_ = cr.IsSourceKind(reflect.String)
}

// ── SimpleRequest / SimpleResult ──

func TestSimpleRequest(t *testing.T) {
	sr := NewSimpleRequest("data", true, "")
	if sr.Request() != "data" || !sr.IsValid() { t.Fatal("unexpected") }
	_ = sr.Message()
	_ = sr.Value()
	_ = sr.InvalidError()
	sr2 := InvalidSimpleRequest("msg")
	_ = sr2.InvalidError()
	sr3 := InvalidSimpleRequestNoMessage()
	_ = sr3
	sr4 := NewSimpleRequestValid("data")
	_ = sr4.IsPointer()
	_ = sr4.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	_ = sr4.GetErrorOnTypeMismatch(reflect.TypeOf(42), true)
}

func TestSimpleResult(t *testing.T) {
	sr := NewSimpleResult("data", true, "")
	_ = sr.Result
	_ = sr.Message
	_ = sr.InvalidError()
	_ = sr.Clone()
	_ = sr.ClonePtr()
	var nilSr *SimpleResult
	if nilSr.ClonePtr() != nil { t.Fatal("expected nil") }
	sr2 := InvalidSimpleResult("msg")
	_ = sr2.InvalidError()
	sr3 := InvalidSimpleResultNoMessage()
	_ = sr3
	sr4 := NewSimpleResultValid("data")
	_ = sr4.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
}

// ── DynamicStatus ──

func TestDynamicStatus(t *testing.T) {
	ds := InvalidDynamicStatusNoMessage()
	_ = ds
	ds2 := InvalidDynamicStatus("msg")
	_ = ds2
}

// ── ValueStatus ──

func TestValueStatus_Dynamic(t *testing.T) {
	vs := InvalidValueStatusNoMessage()
	_ = vs
	vs2 := InvalidValueStatus("msg")
	_ = vs2
}
