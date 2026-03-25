package corestr

import (
	"testing"
)

// ── SimpleSlice comprehensive ──

func TestSimpleSlice_Add_C15(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	ss.Add("b")
	if ss.Length() != 2 { t.Fatal("expected 2") }
}

func TestSimpleSlice_AddSplit_C15(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AddSplit("a,b,c", ",")
	if ss.Length() != 3 { t.Fatal("expected 3") }
}

func TestSimpleSlice_AddIf_C15(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AddIf(true, "a")
	ss.AddIf(false, "b")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_Adds_C15(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.Adds("a", "b")
	if ss.Length() != 2 { t.Fatal("expected 2") }
}

func TestSimpleSlice_Append_C15(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.Append("a", "b")
	if ss.Length() != 2 { t.Fatal("expected 2") }
}

func TestSimpleSlice_AppendFmt_C15(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AppendFmt("hello %s", "world")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_AppendFmt_Empty_C15(t *testing.T) {
	ss := New.SimpleSlice.Cap(5)
	ss.AppendFmt("")
	if ss.Length() != 0 { t.Fatal("expected 0") }
}

func TestSimpleSlice_IsEmpty_C15(t *testing.T) {
	ss := New.SimpleSlice.Cap(0)
	if !ss.IsEmpty() { t.Fatal("expected true") }
}

func TestSimpleSlice_Length_C15(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if ss.Length() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_LastIndex_C15(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a", "b")
	if ss.LastIndex() != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_HasIndex_C15(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if !ss.HasIndex(0) { t.Fatal("expected true") }
	if ss.HasIndex(5) { t.Fatal("expected false") }
}

func TestSimpleSlice_List_C15(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if len(ss.List()) != 1 { t.Fatal("expected 1") }
}

func TestSimpleSlice_Strings_C15(t *testing.T) {
	ss := New.SimpleSlice.SpreadStrings("a")
	if len(ss.Strings()) != 1 { t.Fatal("expected 1") }
}

// ── AnyToString ──

func TestAnyToString_C15(t *testing.T) {
	s := AnyToString(42)
	if s == "" { t.Fatal("expected non-empty") }
}

// ── AllIndividualsLengthOfSimpleSlices ──

func TestAllIndividualsLengthOfSimpleSlices_C15(t *testing.T) {
	ss1 := New.SimpleSlice.SpreadStrings("a", "b")
	ss2 := New.SimpleSlice.SpreadStrings("c")
	l := AllIndividualsLengthOfSimpleSlices([]*SimpleSlice{ss1, ss2})
	if l != 3 { t.Fatal("expected 3") }
}

func TestAllIndividualsLengthOfSimpleSlices_Nil_C15(t *testing.T) {
	l := AllIndividualsLengthOfSimpleSlices(nil)
	if l != 0 { t.Fatal("expected 0") }
}

// ── ValidValues ──

func TestValidValues_C15(t *testing.T) {
	vv := ValidValues{
		Values: []ValidValue{
			NewValidValue("a"),
			NewInvalidValue("b"),
		},
	}
	if len(vv.Values) != 2 { t.Fatal("expected 2") }
}

// ── ValueStatus ──

func TestValueStatus_C15(t *testing.T) {
	vs := ValueStatus{
		Value:    "test",
		HasValue: true,
	}
	if !vs.HasValue { t.Fatal("expected true") }
}

// ── TextWithLineNumber ──

func TestTextWithLineNumber_C15(t *testing.T) {
	tw := TextWithLineNumber{
		Text:       "hello",
		LineNumber: 1,
	}
	if tw.Text != "hello" { t.Fatal("unexpected") }
}

// ── LeftRight ──

func TestLeftRight_C15(t *testing.T) {
	lr := LeftRight{Left: "a", Right: "b"}
	if lr.Left != "a" { t.Fatal("unexpected") }
}

// ── LeftMiddleRight ──

func TestLeftMiddleRight_C15(t *testing.T) {
	lmr := LeftMiddleRight{Left: "a", Middle: "b", Right: "c"}
	if lmr.Middle != "b" { t.Fatal("unexpected") }
}

// ── KeyValuePair ──

func TestKeyValuePair_C15(t *testing.T) {
	kv := KeyValuePair{Key: "k", Value: "v"}
	if kv.Key != "k" { t.Fatal("unexpected") }
}

// ── KeyValueCollection ──

func TestKeyValueCollection_Add_C15(t *testing.T) {
	kvc := New.KeyValues.Cap(5)
	kvc.Add("k", "v")
	if kvc.Length() != 1 { t.Fatal("expected 1") }
}

func TestKeyValueCollection_IsEmpty_C15(t *testing.T) {
	kvc := New.KeyValues.Cap(0)
	if !kvc.IsEmpty() { t.Fatal("expected true") }
}

// ── HashsetsCollection ──

func TestHashsetsCollection_IsEmpty_C15(t *testing.T) {
	hsc := New.HashsetsCollection.Cap(0)
	if !hsc.IsEmpty() { t.Fatal("expected true") }
}

func TestHashsetsCollection_Add_C15(t *testing.T) {
	hsc := New.HashsetsCollection.Cap(5)
	hs := New.Hashset.Strings([]string{"a"})
	hsc.Add(hs)
	if hsc.Length() != 1 { t.Fatal("expected 1") }
}

// ── HashmapDataModel ──

func TestNewHashmapUsingDataModel_C15(t *testing.T) {
	dm := &HashmapDataModel{Items: map[string]string{"a": "1"}}
	hm := NewHashmapUsingDataModel(dm)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmapsDataModelUsing_C15(t *testing.T) {
	hm := New.Hashmap.KeyValues(KeyValuePair{Key: "a", Value: "1"})
	dm := NewHashmapsDataModelUsing(hm)
	if len(dm.Items) != 1 { t.Fatal("expected 1") }
}

// ── SimpleStringOnce ──

func TestSimpleStringOnce_C15(t *testing.T) {
	so := New.SimpleStringOnce.Value("test")
	if so.IsEmpty() { t.Fatal("expected non-empty") }
}

// ── newCollectionCreator extended ──

func TestNewCollection_LenCap_C15(t *testing.T) {
	c := New.Collection.LenCap(5, 10)
	if c.Length() != 5 { t.Fatal("expected 5") }
}

func TestNewCollection_LineUsingSep_C15(t *testing.T) {
	c := New.Collection.LineUsingSep(",", "a,b,c")
	if c.Length() != 3 { t.Fatal("expected 3") }
}

func TestNewCollection_LineDefault_C15(t *testing.T) {
	c := New.Collection.LineDefault("a | b")
	if c.Length() < 1 { t.Fatal("expected >= 1") }
}

func TestNewCollection_StringsPlusCap_C15(t *testing.T) {
	c := New.Collection.StringsPlusCap(5, []string{"a"})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollection_CapStrings_C15(t *testing.T) {
	c := New.Collection.CapStrings(5, []string{"a"})
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollection_CloneStrings_C15(t *testing.T) {
	orig := []string{"a", "b"}
	c := New.Collection.CloneStrings(orig)
	orig[0] = "X"
	if c.First() != "a" { t.Fatal("expected deep clone") }
}

// ── newHashmapCreator extended ──

func TestNewHashmap_KeyAnyValues_C15(t *testing.T) {
	hm := New.Hashmap.KeyAnyValues(KeyAnyValuePair{Key: "a", Value: 1})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmap_KeyValuesCollection_C15(t *testing.T) {
	keys := New.Collection.Strings([]string{"a"})
	vals := New.Collection.Strings([]string{"1"})
	hm := New.Hashmap.KeyValuesCollection(keys, vals)
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmap_KeyValuesStrings_C15(t *testing.T) {
	hm := New.Hashmap.KeyValuesStrings([]string{"a"}, []string{"1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewHashmap_MapWithCap_C15(t *testing.T) {
	hm := New.Hashmap.MapWithCap(5, map[string]string{"a": "1"})
	if hm.Length() != 1 { t.Fatal("expected 1") }
}

// ── newHashsetCreator extended ──

func TestNewHashset_StringsOption_C15(t *testing.T) {
	hs := New.Hashset.StringsOption(10, false, "a", "b")
	if hs.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewHashset_Empty_C15(t *testing.T) {
	hs := New.Hashset.Empty()
	if !hs.IsEmpty() { t.Fatal("expected true") }
}

// ── newCollectionsOfCollectionCreator extended ──

func TestNewCollectionsOfCollection_Empty_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.Empty()
	if !coc.IsEmpty() { t.Fatal("expected true") }
}

func TestNewCollectionsOfCollection_StringsOfStrings_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.StringsOfStrings(false, []string{"a"}, []string{"b"})
	if coc.Length() != 2 { t.Fatal("expected 2") }
}

func TestNewCollectionsOfCollection_SpreadStrings_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.SpreadStrings(false, "a", "b")
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollectionsOfCollection_CloneStrings_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.CloneStrings([]string{"a"})
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

func TestNewCollectionsOfCollection_StringsOptions_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.StringsOptions(false, 5, []string{"a"})
	if coc.Length() != 1 { t.Fatal("expected 1") }
}

// ── CollectionsOfCollection JSON ──

func TestCollectionsOfCollection_MarshalJSON_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	b, err := coc.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_UnmarshalJSON_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	b, _ := coc.MarshalJSON()
	coc2 := &CollectionsOfCollection{}
	err := coc2.UnmarshalJSON(b)
	if err != nil { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_Json_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	r := coc.Json()
	if r.HasError() { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_ParseInjectUsingJson_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	jr := coc.Json()
	coc2 := &CollectionsOfCollection{}
	_, err := coc2.ParseInjectUsingJson(&jr)
	if err != nil { t.Fatal("unexpected") }
}

func TestCollectionsOfCollection_JsonParseSelfInject_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.Strings([]string{"a"})
	jr := coc.Json()
	coc2 := &CollectionsOfCollection{}
	err := coc2.JsonParseSelfInject(&jr)
	if err != nil { t.Fatal("unexpected") }
}

// ── CollectionsOfCollection AddAsyncFuncItems ──

func TestCollectionsOfCollection_AddEmpty_C15(t *testing.T) {
	coc := New.CollectionsOfCollection.Cap(5)
	coc.Add(New.Collection.Cap(0)) // empty - should be skipped
	if coc.Length() != 0 { t.Fatal("expected 0") }
}
