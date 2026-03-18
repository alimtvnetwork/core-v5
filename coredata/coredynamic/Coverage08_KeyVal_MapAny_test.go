package coredynamic

import (
	"reflect"
	"testing"
	"time"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// === KeyVal ===

func TestKeyVal_Basic(t *testing.T) {
	kv := KeyVal{Key: "name", Value: "hello"}
	if kv.KeyString() != "name" { t.Fatal("expected name") }
	if kv.ValueString() != "hello" { t.Fatal("expected hello") }
	if kv.IsKeyNull() { t.Fatal("expected false") }
	if kv.IsValueNull() { t.Fatal("expected false") }
	if kv.IsKeyNullOrEmptyString() { t.Fatal("expected false") }
	_ = kv.KeyDynamic()
	_ = kv.ValueDynamic()
	_ = kv.KeyDynamicPtr()
	_ = kv.ValueDynamicPtr()
	_ = kv.ValueReflectValue()
	_ = kv.String()
}

func TestKeyVal_ValueCasts(t *testing.T) {
	kvi := KeyVal{Key: "k", Value: 42}
	if kvi.ValueInt() != 42 { t.Fatal("expected 42") }
	kvi2 := KeyVal{Key: "k", Value: "x"}
	if kvi2.ValueInt() != -1 { t.Fatal("expected -1") }
	kvb := KeyVal{Key: "k", Value: true}
	if !kvb.ValueBool() { t.Fatal("expected true") }
	kvu := KeyVal{Key: "k", Value: uint(5)}
	if kvu.ValueUInt() != 5 { t.Fatal("expected 5") }
	kv64 := KeyVal{Key: "k", Value: int64(99)}
	if kv64.ValueInt64() != 99 { t.Fatal("expected 99") }
	kvs := KeyVal{Key: "k", Value: []string{"a"}}
	if len(kvs.ValueStrings()) != 1 { t.Fatal("expected 1") }
	kvn := KeyVal{Key: "k", Value: "x"}
	if kvn.ValueBool() { t.Fatal("expected false") }
	if kvn.ValueUInt() != 0 { t.Fatal("expected 0") }
	if kvn.ValueInt64() != -1 { t.Fatal("expected -1") }
	if kvn.ValueStrings() != nil { t.Fatal("expected nil") }
}

func TestKeyVal_CastKeyVal(t *testing.T) {
	kv := KeyVal{Key: "name", Value: "hello"}
	var k, v string
	err := kv.CastKeyVal(&k, &v)
	if err != nil { t.Fatal("unexpected") }
	var nilKV *KeyVal
	err2 := nilKV.CastKeyVal(&k, &v)
	if err2 == nil { t.Fatal("expected error") }
}

func TestKeyVal_ReflectSet(t *testing.T) {
	kv := KeyVal{Key: "name", Value: "hello"}
	var k string
	err := kv.ReflectSetKey(&k)
	if err != nil { t.Fatal("unexpected") }
	var v string
	err2 := kv.ReflectSetTo(&v)
	if err2 != nil { t.Fatal("unexpected") }
	kv.ReflectSetToMust(&v)
	var nilKV *KeyVal
	if nilKV.ReflectSetKey(&k) == nil { t.Fatal("expected error") }
	if nilKV.ReflectSetTo(&v) == nil { t.Fatal("expected error") }
}

func TestKeyVal_KeyValueReflectSet(t *testing.T) {
	kv := KeyVal{Key: "name", Value: "hello"}
	var v string
	err := kv.KeyReflectSet(&v)
	if err != nil { t.Fatal("unexpected") }
	err2 := kv.ValueReflectSet(&v)
	if err2 != nil { t.Fatal("unexpected") }
	var nilKV *KeyVal
	if nilKV.KeyReflectSet(&v) == nil { t.Fatal("expected error") }
	if nilKV.ValueReflectSet(&v) == nil { t.Fatal("expected error") }
}

func TestKeyVal_NullErrors(t *testing.T) {
	kv := KeyVal{Key: "k", Value: "v"}
	if kv.ValueNullErr() != nil { t.Fatal("expected nil") }
	if kv.KeyNullErr() != nil { t.Fatal("expected nil") }
	var nilKV *KeyVal
	if nilKV.ValueNullErr() == nil { t.Fatal("expected error") }
	if nilKV.KeyNullErr() == nil { t.Fatal("expected error") }
	if nilKV.KeyString() != "" { t.Fatal("expected empty") }
	if nilKV.ValueString() != "" { t.Fatal("expected empty") }
}

func TestKeyVal_NullKey(t *testing.T) {
	kv := KeyVal{Key: nil, Value: "v"}
	if !kv.IsKeyNull() { t.Fatal("expected true") }
	if kv.KeyNullErr() == nil { t.Fatal("expected error") }
}

func TestKeyVal_NullValue(t *testing.T) {
	kv := KeyVal{Key: "k", Value: nil}
	if !kv.IsValueNull() { t.Fatal("expected true") }
	if kv.ValueNullErr() == nil { t.Fatal("expected error") }
}

func TestKeyVal_JSON(t *testing.T) {
	kv := KeyVal{Key: "k", Value: "v"}
	_ = kv.JsonModel()
	_ = kv.JsonModelAny()
	_ = kv.Json()
	_ = kv.JsonPtr()
	b, err := kv.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestKeyVal_ParseInjectUsingJson(t *testing.T) {
	kv := KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	kv2 := &KeyVal{}
	_, err := kv2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
	_ = kv2.JsonParseSelfInject(&jr)
}

// === KeyValCollection ===

func TestKeyValCollection_Basic(t *testing.T) {
	kvc := EmptyKeyValCollection()
	if !kvc.IsEmpty() { t.Fatal("expected empty") }
	if kvc.HasAnyItem() { t.Fatal("expected false") }
	kvc.Add(KeyVal{Key: "a", Value: 1})
	if kvc.Length() != 1 { t.Fatal("expected 1") }
	kvc.AddPtr(nil)
	kvc.AddPtr(&KeyVal{Key: "b", Value: 2})
	if kvc.Length() != 2 { t.Fatal("expected 2") }
	kvc.AddMany(KeyVal{Key: "c", Value: 3})
	kvc.AddMany()
	kvc.AddManyPtr(nil, &KeyVal{Key: "d", Value: 4})
	kvc.AddManyPtr()
}

func TestKeyValCollection_NilReceiver(t *testing.T) {
	var kvc *KeyValCollection
	if kvc.Length() != 0 { t.Fatal("expected 0") }
	if kvc.Items() != nil { t.Fatal("expected nil") }
	if kvc.String() != "" { t.Fatal("expected empty") }
}

func TestKeyValCollection_Items(t *testing.T) {
	kvc := NewKeyValCollection(2)
	kvc.Add(KeyVal{Key: "a", Value: 1})
	items := kvc.Items()
	if len(items) != 1 { t.Fatal("expected 1") }
}

func TestKeyValCollection_MapAnyItems(t *testing.T) {
	kvc := EmptyKeyValCollection()
	m := kvc.MapAnyItems()
	if m.Length() != 0 { t.Fatal("expected 0") }
	kvc.Add(KeyVal{Key: "a", Value: 1})
	m2 := kvc.MapAnyItems()
	if m2.Length() != 1 { t.Fatal("expected 1") }
}

func TestKeyValCollection_JsonMapResults(t *testing.T) {
	kvc := EmptyKeyValCollection()
	mr, err := kvc.JsonMapResults()
	if err != nil || mr.Length() != 0 { t.Fatal("unexpected") }
	kvc.Add(KeyVal{Key: "a", Value: 1})
	mr2, err2 := kvc.JsonMapResults()
	if err2 != nil || mr2.Length() != 1 { t.Fatal("unexpected") }
}

func TestKeyValCollection_JsonResults(t *testing.T) {
	kvc := EmptyKeyValCollection()
	_ = kvc.JsonResultsCollection()
	_ = kvc.JsonResultsPtrCollection()
	kvc.Add(KeyVal{Key: "a", Value: 1})
	_ = kvc.JsonResultsCollection()
	_ = kvc.JsonResultsPtrCollection()
}

func TestKeyValCollection_AllKeys(t *testing.T) {
	kvc := EmptyKeyValCollection()
	if len(kvc.AllKeys()) != 0 { t.Fatal("expected 0") }
	if len(kvc.AllKeysSorted()) != 0 { t.Fatal("expected 0") }
	if len(kvc.AllValues()) != 0 { t.Fatal("expected 0") }
	kvc.Add(KeyVal{Key: "b", Value: 2}).Add(KeyVal{Key: "a", Value: 1})
	keys := kvc.AllKeys()
	if len(keys) != 2 { t.Fatal("expected 2") }
	sorted := kvc.AllKeysSorted()
	if sorted[0] != "a" { t.Fatal("expected a first") }
	vals := kvc.AllValues()
	if len(vals) != 2 { t.Fatal("expected 2") }
}

func TestKeyValCollection_Paging(t *testing.T) {
	kvc := NewKeyValCollection(10)
	for i := 0; i < 10; i++ { kvc.Add(KeyVal{Key: "k", Value: i}) }
	if kvc.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if kvc.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := kvc.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
	single := kvc.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }
	small := NewKeyValCollection(1)
	small.Add(KeyVal{Key: "k", Value: 1})
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
}

func TestKeyValCollection_JSON(t *testing.T) {
	kvc := NewKeyValCollection(1)
	kvc.Add(KeyVal{Key: "k", Value: "v"})
	_ = kvc.JsonModel()
	_ = kvc.JsonModelAny()
	_ = kvc.Json()
	_ = kvc.JsonPtr()
	_ = kvc.String()
	s, err := kvc.JsonString()
	if err != nil || s == "" { t.Fatal("unexpected") }
	_ = kvc.JsonStringMust()
	b, err2 := kvc.Serialize()
	if err2 != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestKeyValCollection_Clone(t *testing.T) {
	kvc := NewKeyValCollection(2)
	kvc.Add(KeyVal{Key: "a", Value: 1})
	c := kvc.Clone()
	if c.Length() != 1 { t.Fatal("expected 1") }
	cp := kvc.ClonePtr()
	if cp == nil || cp.Length() != 1 { t.Fatal("expected 1") }
	var nilKVC *KeyValCollection
	if nilKVC.ClonePtr() != nil { t.Fatal("expected nil") }
	_ = kvc.NonPtr()
	_ = kvc.Ptr()
}

func TestKeyValCollection_ParseInjectUsingJson(t *testing.T) {
	kvc := NewKeyValCollection(1)
	kvc.Add(KeyVal{Key: "a", Value: 1})
	jr := kvc.Json()
	kvc2 := EmptyKeyValCollection()
	_, err := kvc2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
	_ = kvc2.JsonParseSelfInject(&jr)
}

// === MapAnyItems ===

func TestMapAnyItems_Basic(t *testing.T) {
	m := EmptyMapAnyItems()
	if m.Length() != 0 { t.Fatal("expected 0") }
	if !m.IsEmpty() { t.Fatal("expected empty") }
	if m.HasAnyItem() { t.Fatal("expected false") }
	m.Add("a", 1)
	if m.Length() != 1 { t.Fatal("expected 1") }
	if !m.HasKey("a") { t.Fatal("expected true") }
	if m.HasKey("z") { t.Fatal("expected false") }
}

func TestMapAnyItems_NilReceiver(t *testing.T) {
	var m *MapAnyItems
	if m.Length() != 0 { t.Fatal("expected 0") }
	if m.HasKey("x") { t.Fatal("expected false") }
}

func TestMapAnyItems_Constructors(t *testing.T) {
	_ = NewMapAnyItems(5)
	m := NewMapAnyItemsUsingItems(nil)
	if m.Length() != 0 { t.Fatal("expected 0") }
	m2 := NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	if m2.Length() != 1 { t.Fatal("expected 1") }
	_, _ = NewMapAnyItemsUsingAnyTypeMap(nil)
	_, _ = NewMapAnyItemsUsingAnyTypeMap(map[string]int{"a": 1})
}

func TestMapAnyItems_AddSet(t *testing.T) {
	m := NewMapAnyItems(5)
	m.Add("a", 1)
	m.Set("b", 2)
	m.AddKeyAny(corejson.KeyAny{Key: "c", AnyInf: 3})
	err := m.AddKeyAnyWithValidation(reflect.TypeOf(0), corejson.KeyAny{Key: "d", AnyInf: 4})
	if err != nil { t.Fatal("unexpected") }
	err2 := m.AddKeyAnyWithValidation(reflect.TypeOf(""), corejson.KeyAny{Key: "e", AnyInf: 5})
	if err2 == nil { t.Fatal("expected error") }
	err3 := m.AddWithValidation(reflect.TypeOf(0), "f", 6)
	if err3 != nil { t.Fatal("unexpected") }
	m.AddJsonResultPtr("g", nil)
	jr := corejson.NewResult.AnyPtr("x")
	m.AddJsonResultPtr("g", jr)
}

func TestMapAnyItems_Get(t *testing.T) {
	m := NewMapAnyItems(2)
	m.Add("a", 1)
	v := m.GetValue("a")
	if v != 1 { t.Fatal("expected 1") }
	if m.GetValue("z") != nil { t.Fatal("expected nil") }
	v2, has := m.Get("a")
	if !has || v2 != 1 { t.Fatal("unexpected") }
	_, has2 := m.Get("z")
	if has2 { t.Fatal("expected false") }
}

func TestMapAnyItems_GetFieldsMap(t *testing.T) {
	type s struct{ Name string }
	m := NewMapAnyItems(1)
	m.Add("x", s{Name: "hello"})
	fm, err, found := m.GetFieldsMap("x")
	if !found || err != nil || fm == nil { t.Fatal("unexpected") }
	_, _, found2 := m.GetFieldsMap("z")
	if found2 { t.Fatal("expected false") }
	sfm, found3 := m.GetSafeFieldsMap("x")
	if !found3 || sfm == nil { t.Fatal("unexpected") }
}

func TestMapAnyItems_Deserialize(t *testing.T) {
	m := NewMapAnyItems(1)
	m.Add("a", 42)
	var v int
	err := m.Deserialize("a", &v)
	if err != nil { t.Fatal("unexpected") }
	m.DeserializeMust("a", &v)
	err2 := m.Deserialize("z", &v)
	if err2 == nil { t.Fatal("expected error") }
}

func TestMapAnyItems_ReflectSetTo(t *testing.T) {
	m := NewMapAnyItems(1)
	m.Add("a", "hello")
	var s string
	err := m.ReflectSetTo("a", &s)
	if err != nil { t.Fatal("unexpected") }
	m.ReflectSetToMust("a", &s)
	err2 := m.ReflectSetTo("z", &s)
	if err2 == nil { t.Fatal("expected error") }
}

func TestMapAnyItems_AllKeys(t *testing.T) {
	m := EmptyMapAnyItems()
	if len(m.AllKeys()) != 0 { t.Fatal("expected 0") }
	if len(m.AllKeysSorted()) != 0 { t.Fatal("expected 0") }
	if len(m.AllValues()) != 0 { t.Fatal("expected 0") }
	m.Add("b", 2).Add("a", 1)
	if len(m.AllKeys()) != 2 { t.Fatal("expected 2") }
	if len(m.AllValues()) != 2 { t.Fatal("expected 2") }
}

func TestMapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	m := NewMapAnyItems(2)
	m.Add("a", 1).Add("b", 2)
	sub := m.GetNewMapUsingKeys(false, "a")
	if sub.Length() != 1 { t.Fatal("expected 1") }
	empty := m.GetNewMapUsingKeys(false)
	if empty.Length() != 0 { t.Fatal("expected 0") }
}

func TestMapAnyItems_Paging(t *testing.T) {
	m := NewMapAnyItems(10)
	for i := 0; i < 10; i++ { m.Add(corejson.Serialize.ToString(i), i) }
	if m.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if m.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
	paged := m.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }
	small := NewMapAnyItems(1)
	small.Add("a", 1)
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
}

func TestMapAnyItems_AddMapResult(t *testing.T) {
	m := NewMapAnyItems(2)
	m.AddMapResult(nil)
	m.AddMapResult(map[string]any{"a": 1})
	if m.Length() != 1 { t.Fatal("expected 1") }
	m.AddMapResultOption(false, map[string]any{"a": 2})
	m.AddMapResultOption(true, map[string]any{"a": 3})
	m.AddMapResultOption(true, nil)
	m.AddManyMapResultsUsingOption(true, nil)
	m.AddManyMapResultsUsingOption(true, map[string]any{"b": 1})
}

func TestMapAnyItems_JSON(t *testing.T) {
	m := NewMapAnyItems(1)
	m.Add("a", 1)
	s, err := m.JsonString()
	if err != nil || s == "" { t.Fatal("unexpected") }
	_ = m.JsonStringMust()
	_ = m.Json()
	_ = m.JsonPtr()
	_ = m.JsonModel()
	_ = m.JsonModelAny()
	_ = m.Strings()
	_ = m.String()
	_ = m.MapAnyItems()
}

func TestMapAnyItems_JsonResults(t *testing.T) {
	m := EmptyMapAnyItems()
	_, _ = m.JsonMapResults()
	_ = m.JsonResultsCollection()
	_ = m.JsonResultsPtrCollection()
	m.Add("a", 1)
	_, _ = m.JsonMapResults()
	_ = m.JsonResultsCollection()
	_ = m.JsonResultsPtrCollection()
}

func TestMapAnyItems_JsonResultOfKey(t *testing.T) {
	m := NewMapAnyItems(1)
	m.Add("a", 1)
	_ = m.JsonResultOfKey("a")
	_ = m.JsonResultOfKey("z")
	_ = m.JsonResultOfKeys("a", "z")
	_ = m.JsonResultOfKeys()
}

func TestMapAnyItems_Diff(t *testing.T) {
	m := NewMapAnyItems(2)
	m.Add("a", 1).Add("b", 2)
	m2 := NewMapAnyItems(2)
	m2.Add("a", 1).Add("b", 3)
	_ = m.Diff(false, m2)
	_ = m.DiffRaw(false, m2.Items)
	_ = m.IsRawEqual(false, m.Items)
	_ = m.HasAnyChanges(false, m2.Items)
	_ = m.HashmapDiffUsingRaw(false, m2.Items)
	_ = m.DiffJsonMessage(false, m2.Items)
	_ = m.ToStringsSliceOfDiffMap(m.DiffRaw(false, m2.Items))
	_ = m.ShouldDiffMessage(false, "test", m2.Items)
	_ = m.LogShouldDiffMessage(false, "test", m2.Items)
	_ = m.MapStringAnyDiff()
	_ = m.RawMapStringAnyDiff()
}

func TestMapAnyItems_IsEqual(t *testing.T) {
	m := NewMapAnyItems(1)
	m.Add("a", 1)
	m2 := NewMapAnyItems(1)
	m2.Add("a", 1)
	if !m.IsEqual(m2) { t.Fatal("expected equal") }
	if !m.IsEqualRaw(m2.Items) { t.Fatal("expected equal") }
	var nilM *MapAnyItems
	if !nilM.IsEqual(nil) { t.Fatal("expected equal") }
	if nilM.IsEqual(m) { t.Fatal("expected not equal") }
	if m.IsEqual(nil) { t.Fatal("expected not equal") }
}

func TestMapAnyItems_ClearDispose(t *testing.T) {
	m := NewMapAnyItems(2)
	m.Add("a", 1)
	m.Clear()
	if m.Length() != 0 { t.Fatal("expected 0") }
	m.Add("a", 1)
	m.DeepClear()
	time.Sleep(10 * time.Millisecond)
	m.Dispose()
	var nilM *MapAnyItems
	nilM.Clear()
	nilM.DeepClear()
	nilM.Dispose()
}

func TestMapAnyItems_ClonePtr(t *testing.T) {
	m := NewMapAnyItems(1)
	m.Add("a", 1)
	c, err := m.ClonePtr()
	if err != nil || c == nil { t.Fatal("unexpected") }
	var nilM *MapAnyItems
	_, err2 := nilM.ClonePtr()
	if err2 == nil { t.Fatal("expected error") }
}

func TestMapAnyItems_ParseInjectUsingJson(t *testing.T) {
	m := NewMapAnyItems(1)
	m.Add("a", 1)
	jr := m.Json()
	m2 := EmptyMapAnyItems()
	_, err := m2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
	_ = m2.JsonParseSelfInject(&jr)
}

func TestMapAnyItems_GetItemRef(t *testing.T) {
	m := NewMapAnyItems(2)
	s := "hello"
	m.Add("a", &s)
	var out *string
	err := m.GetItemRef("a", &out)
	if err != nil { t.Fatal("unexpected error") }
	err2 := m.GetItemRef("z", &out)
	if err2 == nil { t.Fatal("expected error") }
	err3 := m.GetItemRef("a", nil)
	if err3 == nil { t.Fatal("expected error") }
}

func TestMapAnyItems_GetUsingUnmarshallManyAt(t *testing.T) {
	m := NewMapAnyItems(2)
	m.Add("a", 1).Add("b", "hello")
	var i int
	var s string
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "a", AnyInf: &i},
		corejson.KeyAny{Key: "b", AnyInf: &s},
	)
	if err != nil { t.Fatal("unexpected") }
}

func TestMapAnyItems_GetManyItemsRefs(t *testing.T) {
	m := NewMapAnyItems(1)
	s := "hello"
	m.Add("a", &s)
	var out *string
	err := m.GetManyItemsRefs(corejson.KeyAny{Key: "a", AnyInf: &out})
	if err != nil { t.Fatal("unexpected") }
	err2 := m.GetManyItemsRefs()
	if err2 != nil { t.Fatal("unexpected") }
}

func TestMapAnyItems_NilRawMapDiff(t *testing.T) {
	var m *MapAnyItems
	d := m.RawMapStringAnyDiff()
	if d == nil { t.Fatal("expected non-nil") }
}

// === MapAnyItemDiff ===

func TestMapAnyItemDiff_Basic(t *testing.T) {
	var d *MapAnyItemDiff
	if d.Length() != 0 { t.Fatal("expected 0") }
	diff := MapAnyItemDiff(map[string]any{"a": 1})
	if diff.IsEmpty() { t.Fatal("expected non-empty") }
	if !diff.HasAnyItem() { t.Fatal("expected true") }
	if diff.LastIndex() != 0 { t.Fatal("expected 0") }
	_ = diff.AllKeysSorted()
	_ = diff.Raw()
	_ = diff.MapAnyItems()
	_ = diff.RawMapDiffer()
	_ = diff.Json()
	_ = diff.JsonPtr()
	_ = diff.PrettyJsonString()
}

func TestMapAnyItemDiff_NilRaw(t *testing.T) {
	var d *MapAnyItemDiff
	if d.Raw() == nil { t.Fatal("expected non-nil") }
}

func TestMapAnyItemDiff_Clear(t *testing.T) {
	var d *MapAnyItemDiff
	_ = d.Clear()
	diff := MapAnyItemDiff(map[string]any{"a": 1})
	_ = diff.Clear()
}

func TestMapAnyItemDiff_DiffMethods(t *testing.T) {
	diff := MapAnyItemDiff(map[string]any{"a": 1})
	right := map[string]any{"a": 2}
	_ = diff.DiffRaw(false, right)
	_ = diff.IsRawEqual(false, right)
	_ = diff.HasAnyChanges(false, right)
	_ = diff.HashmapDiffUsingRaw(false, right)
	_ = diff.DiffJsonMessage(false, right)
	_ = diff.ToStringsSliceOfDiffMap(map[string]any{"a": 1})
	_ = diff.ShouldDiffMessage(false, "test", right)
	_ = diff.LogShouldDiffMessage(false, "test", right)
}

func TestMapAnyItemDiff_LogPrettyJsonString(t *testing.T) {
	diff := MapAnyItemDiff(map[string]any{})
	diff.LogPrettyJsonString()
	diff2 := MapAnyItemDiff(map[string]any{"a": 1})
	diff2.LogPrettyJsonString()
}
