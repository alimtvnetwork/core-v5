package corestr

import (
	"testing"
)

// ── SimpleSlice comprehensive ──

func TestSimpleSlice_Add(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	ss.Add("b")
	if ss.Length() != 2 { t.Fatal("expected 2") }
}

func TestSimpleSlice_AddSplit(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AddSplit("a,b,c", ",")
	if ss.Length() != 3 { t.Fatal("expected 3") }
}

func TestSimpleSlice_AddIf(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AddIf(true, "a")
	ss.AddIf(false, "b")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_Adds(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.Adds("a", "b")
	if ss.Length() != 2 { t.Fatal("expected 2") }
}

func TestSimpleSlice_Append(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.Append("a", "b")
	if ss.Length() != 2 { t.Fatal("expected 2") }
}

func TestSimpleSlice_AppendFmt(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AppendFmt("hello %s", "world")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_AppendFmt_Empty(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AppendFmt("")
	if ss.Length() != 0 { t.Fatal("expected 0") }
}

func TestSimpleSlice_IsEmpty(t *testing.T) {
	ss := New.SimpleSlice.Cap(0)
	if !ss.IsEmpty() { t.Fatal("expected true") }
}

func TestSimpleSlice_Length(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_LastIndex(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a", "b")
	if ss.LastIndex() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_HasIndex(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if !ss.HasIndex(0) { t.Fatal("expected true") }
	if ss.HasIndex(5) { t.Fatal("expected false") }
}

func TestSimpleSlice_List(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if len(ss.List()) != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_Strings(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if len(ss.Strings()) != 1 { t.Fatal("expected 1") }
}

// ── AnyToString ──

func TestAnyToString(t *testing.T) {
	s := AnyToString(42)
	if s == "" { t.Fatal("expected non-empty") }
}

// ── AllIndividualsLengthOfSimpleSlices ──

func TestAllIndividualsLengthOfSimpleSlices(t *testing.T) {
	ss1 := New.SimpleSlice.SpreadStrings("a", "b")
	ss2 := New.SimpleSlice.SpreadStrings("c")
	l := AllIndividualsLengthOfSimpleSlices([]*SimpleSlice{ss1, ss2})
	if l != 3 { t.Fatal("expected 3") }
}

func TestAllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	l := AllIndividualsLengthOfSimpleSlices(nil)
	if l != 0 { t.Fatal("expected 0") }
}

// ── ValidValues ──

func TestValidValues(t *testing.T) {
	vv := ValidValues{
		Values: []ValidValue{
			NewValidValue("a"),
			NewInvalidValue("b"),
		},
	}
	if len(vv.Values) != 2 { t.Fatal("expected 2") }
}

// ── ValueStatus ──

func TestValueStatus(t *testing.T) {
	vs := ValueStatus{
		Value:    "test",
		HasValue: true,
	}
	if !vs.HasValue { t.Fatal("expected true") }
}

// ── TextWithLineNumber ──

func TestTextWithLineNumber(t *testing.T) {
	tw := TextWithLineNumber{
		Text:       "hello",
		LineNumber: 1,
	}
	if tw.Text != "hello" { t.Fatal("unexpected") }
}

// ── LeftRight ──

func TestLeftRight(t *testing.T) {
	lr := LeftRight{Left: "a", Right: "b"}
	if lr.Left != "a" { t.Fatal("unexpected") }
}

// ── LeftMiddleRight ──

func TestLeftMiddleRight(t *testing.T) {
	lmr := LeftMiddleRight{Left: "a", Middle: "b", Right: "c"}
	if lmr.Middle != "b" { t.Fatal("unexpected") }
}

// ── KeyValuePair ──

func TestKeyValuePair(t *testing.T) {
	kv := KeyValuePair{Key: "k", Value: "v"}
	if kv.Key != "k" { t.Fatal("unexpected") }
}

// ── KeyValueCollection ──

func TestKeyValueCollection_Add(t *testing.T) {
	kvc := New.KeyValues.Cap(5)
	kvc.Add("k", "v")
	if kvc.Length() != 1 { t.Fatal("expected 1") }
}

func TestKeyValueCollection_IsEmpty(t *testing.T) {
	kvc := New.KeyValues.Cap(0)
	if !kvc.IsEmpty() { t.Fatal("expected true") }
}

// ── HashsetsCollection ──

func TestHashsetsCollection_IsEmpty(t *testing.T) {
	hsc := New.HashsetsCollection.Cap(0)
	if !hsc.IsEmpty() { t.Fatal("expected true") }
}

func TestHashsetsCollection_Add(t *testing.T) {
	hsc := New.HashsetsCollection.Cap(5)
	hs := New.Hashset.Strings([]string{"a"})
	hsc.Add(hs)
	if hsc.Length() != 1 { t.Fatal("expected 1") }
}

// ── HashmapDataModel ──

func TestNewHashmapUsingDataModel(t *testing.T) {
	dm := &HashmapDataModel{Items: map[string]string{"a": "1"}}
	hm := NewHashmapUsingDataModel(dm)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmapsDataModelUsing(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	dm := NewHashmapsDataModelUsing(hm)
	if len(dm.Items) != 1 { t.Fatal("expected 1") }
}

// ── SimpleStringOnce ──

func TestSimpleStringOnce(t *testing.T) {
	so := New.SimpleStringOnce.Value("test")
	if so.IsEmpty() { t.Fatal("expected non-empty") }
}

// ── newCollectionCreator extended ──

func TestNewCollection_LenCap(t *testing.T) {
	c := New.Collection.LenCap(5, 10)
	if c.Length() != 5 { t.Fatal("expected 5") }
}

func TestNewCollection_LineUsingSep(t *testing.T) {
	c := New.Collection.LineUsingSep(",", "a,b,c")
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestNewCollection_LineDefault(t *testing.T) {
	c := New.Collection.LineDefault("a | b")
	if c.Length() < 1 { t.Fatal("expected >= 1") }
}

func TestNewCollection_StringsPlusCap(t *testing.T) {
	c := New.Collection.StringsPlusCap(5, []string{"a"})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollection_CapStrings(t *testing.T) {
	c := New.Collection.CapStrings(5, []string{"a"})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollection_CloneStrings(t *testing.T) {
	orig := []string{"a", "b"}
	c := New.Collection.CloneStrings(orig)
	orig[0] = "X"
	if c.First() != "a" { t.Fatal("expected deep clone") }
}

// ── newHashmapCreator extended ──

func TestNewHashmap_KeyAnyValues(t *testing.T) {
	hm := New.Hashmap.KeyAnyValues(KeyAnyValuePair{Key: "a", Value: 1})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmap_KeyValuesCollection(t *testing.T) {
	keys := New.Collection.Strings([]string{"a"})
	vals := New.Collection.Strings([]string{"1"})
	hm := New.Hashmap.KeyValuesCollection(keys, vals)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmap_KeyValuesStrings(t *testing.T) {
	hm := New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmap_MapWithCap(t *testing.T) {
	hm := New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

// ── newHashsetCreator extended ──

func TestNewHashset_StringsOption(t *testing.T) {
	hs := New.Hashset.StringsOption(10, false, "a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewHashset_Empty(t *testing.T) {
	hs := New.Hashset.Empty()
	if !hs.IsEmpty() { t.Fatal("expected true") }
}

// ── newCollectionsOfCollectionCreator extended ──

func TestNewCollectionsOfCollection_Empty(t *testing.T) {
	coc := New.CollectionsOfCollection.Empty()
	if !coc.IsEmpty() { t.Fatal("expected true") }
}

func TestNewCollectionsOfCollection_StringsOfStrings(t *testing.T) {
	coc := New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
	if coc.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCollectionsOfCollection_SpreadStrings(t *testing.T) {
	coc := New.CollectionsOfCollection.SpreadStrings(false, "a", "b")
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollectionsOfCollection_CloneStrings(t *testing.T) {
	coc := New.CollectionsOfCollection.CloneStrings([]string{"a"})
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollectionsOfCollection_StringsOptions(t *testing.T) {
	coc := New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"})
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

// ── CollectionsOfCollection JSON ──

func TestCollectionsOfCollection_MarshalJSON(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	b, err := coc.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_UnmarshalJSON(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	b, _ := coc.MarshalJSON()
	coc2 := &CollectionsOfCollection{}
	err := coc2.UnmarshalJSON(b)
	if err != nil { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_Json(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	r := coc.Json()
	if r.HasError() { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_ParseInjectUsingJson(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	jr := coc.Json()
	coc2 := &CollectionsOfCollection{}
	_, err := coc2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_JsonParseSelfInject(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	jr := coc.Json()
	coc2 := &CollectionsOfCollection{}
	err := coc2.JsonParseSelfInject(&jr)
	if err != nil { t.Fatal("unexpected") }
}

// ── CollectionsOfCollection AddAsyncFuncItems ──

func TestCollectionsOfCollection_AddEmpty(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	coc.Add(New.Collection.Cap(0)) // empty - should be skipped
	if coc.Length() != 0 { t.Fatal("expected 0") }
}
