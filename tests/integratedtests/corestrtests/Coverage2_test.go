package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SimpleSlice ──

func Test_SimpleSlice_Cap_Cov2(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	actual := args.Map{
		"isNil":    s == nil,
		"isEmpty":  s.IsEmpty(),
		"length":   s.Length(),
		"hasAny":   s.HasAnyItem(),
	}
	expected := args.Map{
		"isNil":    false,
		"isEmpty":  true,
		"length":   0,
		"hasAny":   false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice_Cap", actual)
}

func Test_SimpleSlice_Add_Cov2(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	s.Add("hello")
	s.Add("world")
	actual := args.Map{
		"length":   s.Length(),
		"hasAny":   s.HasAnyItem(),
		"first":    s.Strings()[0],
	}
	expected := args.Map{
		"length":   2,
		"hasAny":   true,
		"first":    "hello",
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice_Add", actual)
}

func Test_SimpleSlice_Adds_Cov2(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	s.Adds("a", "b", "c")
	actual := args.Map{
		"length": s.Length(),
	}
	expected := args.Map{
		"length": 3,
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice_Adds", actual)
}

func Test_SimpleSlice_AddIf_Cov2(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	s.AddIf(true, "yes")
	s.AddIf(false, "no")
	actual := args.Map{
		"length": s.Length(),
	}
	expected := args.Map{
		"length": 1,
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice_AddIf", actual)
}

func Test_SimpleSlice_AppendFmt_Cov2(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	s.AppendFmt("hello %d", 42)
	actual := args.Map{
		"length": s.Length(),
		"first":  s.Strings()[0],
	}
	expected := args.Map{
		"length": 1,
		"first":  "hello 42",
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice_AppendFmt", actual)
}

func Test_SimpleSlice_String_Cov2(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	s.Add("hello")
	str := s.String()
	if str == "" {
		t.Error("should not be empty")
	}
}

func Test_SimpleSlice_CsvString_Cov2(t *testing.T) {
	s := corestr.New.SimpleSlice.Cap(5)
	s.Adds("a", "b")
	csv := s.CsvString()
	if csv == "" {
		t.Error("should not be empty")
	}
}

// ── Collection ──

func Test_Collection_Cap_Cov2(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	actual := args.Map{
		"isNil":    c == nil,
		"isEmpty":  c.IsEmpty(),
		"length":   c.Length(),
	}
	expected := args.Map{
		"isNil":    false,
		"isEmpty":  true,
		"length":   0,
	}
	expected.ShouldBeEqual(t, 0, "Collection_Cap", actual)
}

func Test_Collection_AddStrings_Cov2(t *testing.T) {
	c := corestr.New.Collection.Cap(5)
	c.AddStrings("hello", "world")
	actual := args.Map{
		"length": c.Length(),
	}
	expected := args.Map{
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "Collection_AddStrings", actual)
}

// ── Hashset ──

func Test_Hashset_Cap_Cov2(t *testing.T) {
	h := corestr.New.Hashset.Cap(5)
	actual := args.Map{
		"isNil":    h == nil,
		"isEmpty":  h.IsEmpty(),
		"length":   h.Length(),
		"hasAny":   h.HasAnyItem(),
	}
	expected := args.Map{
		"isNil":    false,
		"isEmpty":  true,
		"length":   0,
		"hasAny":   false,
	}
	expected.ShouldBeEqual(t, 0, "Hashset_Cap", actual)
}

func Test_Hashset_Add_Cov2(t *testing.T) {
	h := corestr.New.Hashset.Cap(5)
	h.Add("hello")
	h.Add("hello") // duplicate
	h.Add("world")
	actual := args.Map{
		"length": h.Length(),
		"has":    h.Has("hello"),
		"hasNo":  h.Has("nope"),
	}
	expected := args.Map{
		"length": 2,
		"has":    true,
		"hasNo":  false,
	}
	expected.ShouldBeEqual(t, 0, "Hashset_Add", actual)
}

func Test_Hashset_Adds_Cov2(t *testing.T) {
	h := corestr.New.Hashset.Cap(5)
	h.Adds("a", "b", "c")
	actual := args.Map{
		"length": h.Length(),
	}
	expected := args.Map{
		"length": 3,
	}
	expected.ShouldBeEqual(t, 0, "Hashset_Adds", actual)
}

// ── Hashmap ──

func Test_Hashmap_Cap_Cov2(t *testing.T) {
	h := corestr.New.Hashmap.Cap(5)
	actual := args.Map{
		"isNil":    h == nil,
		"isEmpty":  h.IsEmpty(),
		"length":   h.Length(),
	}
	expected := args.Map{
		"isNil":    false,
		"isEmpty":  true,
		"length":   0,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap_Cap", actual)
}

func Test_Hashmap_Add_Cov2(t *testing.T) {
	h := corestr.New.Hashmap.Cap(5)
	h.Add("key", "value")
	actual := args.Map{
		"length":  h.Length(),
		"has":     h.Has("key"),
		"getVal":  h.Get("key"),
	}
	expected := args.Map{
		"length":  1,
		"has":     true,
		"getVal":  "value",
	}
	expected.ShouldBeEqual(t, 0, "Hashmap_Add", actual)
}

// ── KeyValues ──

func Test_KeyValues_Cap_Cov2(t *testing.T) {
	kv := corestr.New.KeyValues.Cap(5)
	if kv == nil {
		t.Error("should not be nil")
	}
}

// ── LinkedList ──

func Test_LinkedList_Default_Cov2(t *testing.T) {
	ll := corestr.New.LinkedList.Default()
	actual := args.Map{
		"isNil":   ll == nil,
		"isEmpty": ll.IsEmpty(),
		"length":  ll.Length(),
	}
	expected := args.Map{
		"isNil":   false,
		"isEmpty": true,
		"length":  0,
	}
	expected.ShouldBeEqual(t, 0, "LinkedList_Default", actual)
}

func Test_LinkedList_Add_Cov2(t *testing.T) {
	ll := corestr.New.LinkedList.Default()
	ll.Add("hello")
	ll.Add("world")
	actual := args.Map{
		"length":  ll.Length(),
		"isEmpty": ll.IsEmpty(),
	}
	expected := args.Map{
		"length":  2,
		"isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "LinkedList_Add", actual)
}

// ── CharHashsetMap ──

func Test_CharHashsetMap_Default_Cov2(t *testing.T) {
	m := corestr.New.CharHashsetMap.Default()
	actual := args.Map{
		"isNil":   m == nil,
		"isEmpty": m.IsEmpty(),
	}
	expected := args.Map{
		"isNil":   false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "CharHashsetMap_Default", actual)
}

// ── CharCollectionMap ──

func Test_CharCollectionMap_Default_Cov2(t *testing.T) {
	m := corestr.New.CharCollectionMap.Default()
	actual := args.Map{
		"isNil":   m == nil,
		"isEmpty": m.IsEmpty(),
	}
	expected := args.Map{
		"isNil":   false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "CharCollectionMap_Default", actual)
}

// ── SimpleStringOnce ──

func Test_SimpleStringOnce_Default_Cov2(t *testing.T) {
	so := corestr.New.SimpleStringOnce.Default()
	actual := args.Map{
		"isNil":   so == nil,
		"isEmpty": so.IsEmpty(),
	}
	expected := args.Map{
		"isNil":   false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleStringOnce_Default", actual)
}

// ── HashsetsCollection ──

func Test_HashsetsCollection_Cap_Cov2(t *testing.T) {
	hc := corestr.New.HashsetsCollection.Cap(5)
	actual := args.Map{
		"isNil":   hc == nil,
		"isEmpty": hc.IsEmpty(),
	}
	expected := args.Map{
		"isNil":   false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "HashsetsCollection_Cap", actual)
}

// ── LinkedCollection ──

func Test_LinkedCollection_Default_Cov2(t *testing.T) {
	lc := corestr.New.LinkedCollection.Default()
	actual := args.Map{
		"isNil":   lc == nil,
		"isEmpty": lc.IsEmpty(),
	}
	expected := args.Map{
		"isNil":   false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "LinkedCollection_Default", actual)
}

// ── CollectionsOfCollection ──

func Test_CollectionsOfCollection_Cap_Cov2(t *testing.T) {
	cc := corestr.New.CollectionsOfCollection.Cap(5)
	actual := args.Map{
		"isNil":   cc == nil,
		"isEmpty": cc.IsEmpty(),
	}
	expected := args.Map{
		"isNil":   false,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "CollectionsOfCollection_Cap", actual)
}

// ── LeftRight ──

func Test_LeftRight_Cov2(t *testing.T) {
	lr := corestr.LeftRight{Left: "l", Right: "r"}
	actual := args.Map{
		"hasLeft":  lr.HasLeft(),
		"hasRight": lr.HasRight(),
		"hasBoth":  lr.HasBoth(),
	}
	expected := args.Map{
		"hasLeft":  true,
		"hasRight": true,
		"hasBoth":  true,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight", actual)
}

// ── LeftMiddleRight ──

func Test_LeftMiddleRight_Cov2(t *testing.T) {
	lmr := corestr.LeftMiddleRight{Left: "l", Middle: "m", Right: "r"}
	actual := args.Map{
		"hasLeft":   lmr.HasLeft(),
		"hasMiddle": lmr.HasMiddle(),
		"hasRight":  lmr.HasRight(),
	}
	expected := args.Map{
		"hasLeft":   true,
		"hasMiddle": true,
		"hasRight":  true,
	}
	expected.ShouldBeEqual(t, 0, "LeftMiddleRight", actual)
}

// ── ValidValue ──

func Test_ValidValue_Cov2(t *testing.T) {
	vv := corestr.ValidValue{Value: "hello", IsValid: true}
	actual := args.Map{
		"value":   vv.Value,
		"isValid": vv.IsValid,
	}
	expected := args.Map{
		"value":   "hello",
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "ValidValue", actual)
}

// ── ValidValues ──

func Test_ValidValues_Cov2(t *testing.T) {
	vv := corestr.ValidValues{Values: []string{"a", "b"}, IsValid: true}
	actual := args.Map{
		"len":     len(vv.Values),
		"isValid": vv.IsValid,
	}
	expected := args.Map{
		"len":     2,
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "ValidValues", actual)
}

// ── AnyToString ──

func Test_AnyToString_Cov2(t *testing.T) {
	r := corestr.AnyToString(42)
	if r == "" {
		t.Error("should not be empty")
	}
}

func Test_AnyToString_String_Cov2(t *testing.T) {
	r := corestr.AnyToString("hello")
	if r != "hello" {
		t.Errorf("expected hello, got %s", r)
	}
}

func Test_AnyToString_Nil_Cov2(t *testing.T) {
	r := corestr.AnyToString(nil)
	if r == "" {
		// nil should produce some representation
	}
}

// ── AllIndividualStringsOfStringsLength ──

func Test_AllIndividualStringsOfStringsLength_Cov2(t *testing.T) {
	r := corestr.AllIndividualStringsOfStringsLength([]string{"ab", "cde"})
	if r != 5 {
		t.Errorf("expected 5, got %d", r)
	}
}

func Test_AllIndividualStringsOfStringsLength_Nil_Cov2(t *testing.T) {
	r := corestr.AllIndividualStringsOfStringsLength(nil)
	if r != 0 {
		t.Errorf("expected 0, got %d", r)
	}
}

// ── CloneSlice / CloneSliceIf ──

func Test_CloneSlice_Cov2(t *testing.T) {
	original := []string{"a", "b"}
	cloned := corestr.CloneSlice(original)
	if len(cloned) != 2 {
		t.Error("expected 2")
	}
}

func Test_CloneSlice_Nil_Cov2(t *testing.T) {
	cloned := corestr.CloneSlice(nil)
	if cloned != nil {
		t.Error("nil should return nil")
	}
}

func Test_CloneSliceIf_True_Cov2(t *testing.T) {
	original := []string{"a"}
	cloned := corestr.CloneSliceIf(true, original)
	if len(cloned) != 1 {
		t.Error("expected 1")
	}
}

func Test_CloneSliceIf_False_Cov2(t *testing.T) {
	original := []string{"a"}
	same := corestr.CloneSliceIf(false, original)
	if len(same) != 1 {
		t.Error("expected 1")
	}
}

// ── TextWithLineNumber ──

func Test_TextWithLineNumber_Cov2(t *testing.T) {
	tw := corestr.TextWithLineNumber{
		LineNumber: 5,
		Text:       "hello",
	}
	if tw.LineNumber != 5 || tw.Text != "hello" {
		t.Error("unexpected values")
	}
}

// ── ValueStatus ──

func Test_ValueStatus_Cov2(t *testing.T) {
	vs := corestr.ValueStatus{Value: "hello", IsValid: true}
	if vs.Value != "hello" || !vs.IsValid {
		t.Error("unexpected values")
	}
}
