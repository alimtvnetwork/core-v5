package corestr

import (
	"testing"
)

// ── CloneSlice ──

func TestCloneSlice_Empty(t *testing.T) {
	result := CloneSlice(nil)
	if len(result) != 0 {
		t.Fatal("expected empty")
	}
}

func TestCloneSlice_WithItems(t *testing.T) {
	result := CloneSlice([]string{"a", "b"})
	if len(result) != 2 || result[0] != "a" {
		t.Fatal("unexpected")
	}
}

// ── CloneSliceIf ──

func TestCloneSliceIf_Empty(t *testing.T) {
	result := CloneSliceIf(true)
	if len(result) != 0 {
		t.Fatal("expected empty")
	}
}

func TestCloneSliceIf_NoClone(t *testing.T) {
	result := CloneSliceIf(false, "a", "b")
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

func TestCloneSliceIf_Clone(t *testing.T) {
	result := CloneSliceIf(true, "a", "b")
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}

// ── AnyToString ──

func TestAnyToString_Empty(t *testing.T) {
	if AnyToString(false, "") != "" {
		t.Fatal("expected empty")
	}
}

func TestAnyToString_WithFieldName(t *testing.T) {
	s := AnyToString(true, "hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyToString_WithoutFieldName(t *testing.T) {
	s := AnyToString(false, "hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestAnyToString_Ptr(t *testing.T) {
	val := "hello"
	s := AnyToString(false, &val)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

// ── reflectInterfaceVal ──

func TestReflectInterfaceVal_Nil(t *testing.T) {
	if reflectInterfaceVal(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestReflectInterfaceVal_Value(t *testing.T) {
	v := reflectInterfaceVal(42)
	if v != 42 {
		t.Fatal("expected 42")
	}
}

func TestReflectInterfaceVal_Ptr(t *testing.T) {
	val := "hello"
	v := reflectInterfaceVal(&val)
	if v != "hello" {
		t.Fatal("expected hello")
	}
}

// ── AllIndividualStringsOfStringsLength ──

func TestAllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	if AllIndividualStringsOfStringsLength(nil) != 0 {
		t.Fatal("expected 0")
	}
}

func TestAllIndividualStringsOfStringsLength_Empty(t *testing.T) {
	items := [][]string{}
	if AllIndividualStringsOfStringsLength(&items) != 0 {
		t.Fatal("expected 0")
	}
}

func TestAllIndividualStringsOfStringsLength_WithItems(t *testing.T) {
	items := [][]string{{"a", "b"}, {"c"}}
	if AllIndividualStringsOfStringsLength(&items) != 3 {
		t.Fatal("expected 3")
	}
}

// ── AllIndividualsLengthOfSimpleSlices ──

func TestAllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	if AllIndividualsLengthOfSimpleSlices() != 0 {
		t.Fatal("expected 0")
	}
}

func TestAllIndividualsLengthOfSimpleSlices_WithItems(t *testing.T) {
	s1 := New.SimpleSlice.Lines("a", "b")
	s2 := New.SimpleSlice.Lines("c")
	if AllIndividualsLengthOfSimpleSlices(s1, s2) != 3 {
		t.Fatal("expected 3")
	}
}

// ── utils ──

func TestUtils_WrapDouble(t *testing.T) {
	if StringUtils.WrapDouble("a") != `"a"` {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapSingle(t *testing.T) {
	if StringUtils.WrapSingle("a") != "'a'" {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapTilda(t *testing.T) {
	if StringUtils.WrapTilda("a") != "`a`" {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapDoubleIfMissing_Empty(t *testing.T) {
	if StringUtils.WrapDoubleIfMissing("") != `""` {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	if StringUtils.WrapDoubleIfMissing(`"a"`) != `"a"` {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	if StringUtils.WrapDoubleIfMissing("a") != `"a"` {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapSingleIfMissing_Empty(t *testing.T) {
	if StringUtils.WrapSingleIfMissing("") != "''" {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	if StringUtils.WrapSingleIfMissing("'a'") != "'a'" {
		t.Fatal("unexpected")
	}
}

func TestUtils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	if StringUtils.WrapSingleIfMissing("a") != "'a'" {
		t.Fatal("unexpected")
	}
}

// ── Empty creators ──

func TestEmptyCreator_All(t *testing.T) {
	_ = Empty.Collection()
	_ = Empty.LinkedList()
	_ = Empty.SimpleSlice()
	_ = Empty.KeyAnyValuePair()
	_ = Empty.KeyValuePair()
	_ = Empty.KeyValueCollection()
	_ = Empty.LinkedCollections()
	_ = Empty.LeftRight()
	_ = Empty.SimpleStringOnce()
	_ = Empty.SimpleStringOncePtr()
	_ = Empty.Hashset()
	_ = Empty.HashsetsCollection()
	_ = Empty.Hashmap()
	_ = Empty.CharCollectionMap()
	_ = Empty.KeyValuesCollection()
	_ = Empty.CollectionsOfCollection()
	_ = Empty.CharHashsetMap()
}

// ── DataModels ──

func TestCharCollectionDataModel(t *testing.T) {
	dm := &CharCollectionDataModel{
		Items:                  map[byte]*Collection{},
		EachCollectionCapacity: 10,
	}
	ccm := NewCharCollectionMapUsingDataModel(dm)
	if ccm == nil {
		t.Fatal("expected non-nil")
	}
	dm2 := NewCharCollectionMapDataModelUsing(ccm)
	if dm2 == nil {
		t.Fatal("expected non-nil")
	}
}

func TestCharHashsetDataModel(t *testing.T) {
	dm := &CharHashsetDataModel{
		Items:               map[byte]*Hashset{},
		EachHashsetCapacity: 10,
	}
	chm := NewCharHashsetMapUsingDataModel(dm)
	if chm == nil {
		t.Fatal("expected non-nil")
	}
	dm2 := NewCharHashsetMapDataModelUsing(chm)
	if dm2 == nil {
		t.Fatal("expected non-nil")
	}
}

func TestHashmapDataModel(t *testing.T) {
	dm := &HashmapDataModel{Items: map[string]string{"a": "b"}}
	hm := NewHashmapUsingDataModel(dm)
	if hm == nil || hm.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	dm2 := NewHashmapsDataModelUsing(hm)
	if dm2 == nil {
		t.Fatal("expected non-nil")
	}
}

func TestHashsetDataModel(t *testing.T) {
	dm := &HashsetDataModel{Items: map[string]bool{"a": true}}
	hs := NewHashsetUsingDataModel(dm)
	if hs == nil || hs.IsEmpty() {
		t.Fatal("expected non-empty")
	}
	dm2 := NewHashsetsDataModelUsing(hs)
	if dm2 == nil {
		t.Fatal("expected non-nil")
	}
}

func TestHashsetsCollectionDataModel(t *testing.T) {
	dm := &HashsetsCollectionDataModel{Items: []*Hashset{}}
	hc := NewHashsetsCollectionUsingDataModel(dm)
	if hc == nil {
		t.Fatal("expected non-nil")
	}
	dm2 := NewHashsetsCollectionDataModelUsing(hc)
	if dm2 == nil {
		t.Fatal("expected non-nil")
	}
}

// ── SimpleStringOnceModel ──

func TestSimpleStringOnceModel(t *testing.T) {
	m := SimpleStringOnceModel{Value: "hello", IsInitialize: true}
	if m.Value != "hello" {
		t.Fatal("unexpected")
	}
}

// ── CollectionsOfCollectionModel ──

func TestCollectionsOfCollectionModel(t *testing.T) {
	m := CollectionsOfCollectionModel{Items: []*Collection{}}
	if m.Items == nil {
		t.Fatal("unexpected")
	}
}
