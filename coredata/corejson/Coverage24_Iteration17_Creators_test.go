package corejson

import (
	"errors"
	"testing"
)

// Covers: newBytesCollectionCreator, newResultsCollectionCreator,
// newResultsPtrCollectionCreator, newMapResultsCreator

// --- newBytesCollectionCreator ---

func Test_I17_NewBytesCollection_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(BytesCollection{Items: [][]byte{[]byte(`"a"`)}})
	bc, err := NewBytesCollection.UnmarshalUsingBytes(b)
	if err != nil || bc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := NewBytesCollection.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingResult_Issues(t *testing.T) {
	r := &Result{Error: errors.New("bad")}
	_, err := NewBytesCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingResult_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(BytesCollection{Items: [][]byte{[]byte(`"a"`)}})
	r := &Result{Bytes: b}
	bc, err := NewBytesCollection.DeserializeUsingResult(r)
	if err != nil || bc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingResult_InvalidPayload(t *testing.T) {
	r := &Result{Bytes: []byte("invalid")}
	_, err := NewBytesCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewBytesCollection_Empty(t *testing.T) {
	bc := NewBytesCollection.Empty()
	if bc == nil || !bc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewBytesCollection_UsingCap(t *testing.T) {
	bc := NewBytesCollection.UsingCap(10)
	if bc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewBytesCollection_AnyItems(t *testing.T) {
	bc, err := NewBytesCollection.AnyItems("a", "b")
	if err != nil || bc.Length() != 2 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewBytesCollection_Serializers_Empty(t *testing.T) {
	bc := NewBytesCollection.Serializers()
	if !bc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewBytesCollection_Serializers_Valid(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	bc := NewBytesCollection.Serializers(s)
	if bc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// --- newResultsCollectionCreator ---

func Test_I17_NewResultsCollection_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(ResultsCollection{Items: []Result{}})
	rc, err := NewResultsCollection.UnmarshalUsingBytes(b)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := NewResultsCollection.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingResult_Issues(t *testing.T) {
	r := &Result{Error: errors.New("bad")}
	_, err := NewResultsCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingResult_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(ResultsCollection{Items: []Result{}})
	r := &Result{Bytes: b}
	rc, err := NewResultsCollection.DeserializeUsingResult(r)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingResult_InvalidPayload(t *testing.T) {
	r := &Result{Bytes: []byte("invalid")}
	_, err := NewResultsCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsCollection_Default(t *testing.T) {
	rc := NewResultsCollection.Default()
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingCap(t *testing.T) {
	rc := NewResultsCollection.UsingCap(5)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_AnyItems(t *testing.T) {
	rc := NewResultsCollection.AnyItems("a", "b")
	if rc.Length() != 2 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_AnyItemsPlusCap_Empty(t *testing.T) {
	rc := NewResultsCollection.AnyItemsPlusCap(5)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_AnyItemsPlusCap_Valid(t *testing.T) {
	rc := NewResultsCollection.AnyItemsPlusCap(5, "a")
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingJsonersOption_NilJsoners(t *testing.T) {
	rc := NewResultsCollection.UsingJsonersOption(true, 5, nil...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtrPlusCap_Nil(t *testing.T) {
	rc := NewResultsCollection.UsingResultsPtrPlusCap(5, nil...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtrPlusCap_Empty(t *testing.T) {
	results := []*Result{}
	rc := NewResultsCollection.UsingResultsPtrPlusCap(5, results...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtrPlusCap_Valid(t *testing.T) {
	r := NewPtr("a")
	rc := NewResultsCollection.UsingResultsPtrPlusCap(0, r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtr(t *testing.T) {
	r := NewPtr("a")
	rc := NewResultsCollection.UsingResultsPtr(r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPlusCap_Nil(t *testing.T) {
	rc := NewResultsCollection.UsingResultsPlusCap(5, nil...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPlusCap_Empty(t *testing.T) {
	results := []Result{}
	rc := NewResultsCollection.UsingResultsPlusCap(5, results...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPlusCap_Valid(t *testing.T) {
	r := New("a")
	rc := NewResultsCollection.UsingResultsPlusCap(0, r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingResults(t *testing.T) {
	r := New("a")
	rc := NewResultsCollection.UsingResults(r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_Serializers_Empty(t *testing.T) {
	rc := NewResultsCollection.Serializers()
	if !rc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResultsCollection_Serializers_Valid(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	rc := NewResultsCollection.Serializers(s)
	if rc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_NewResultsCollection_SerializerFunctions_Empty(t *testing.T) {
	rc := NewResultsCollection.SerializerFunctions()
	if !rc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResultsCollection_SerializerFunctions_Valid(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"x"`), nil }
	rc := NewResultsCollection.SerializerFunctions(fn)
	if rc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// --- newResultsPtrCollectionCreator ---

func Test_I17_NewResultsPtrCollection_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(ResultsPtrCollection{Items: []*Result{}})
	rc, err := NewResultsPtrCollection.UnmarshalUsingBytes(b)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := NewResultsPtrCollection.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingResult_Issues(t *testing.T) {
	r := &Result{Error: errors.New("bad")}
	_, err := NewResultsPtrCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingResult_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(ResultsPtrCollection{Items: []*Result{}})
	r := &Result{Bytes: b}
	rc, err := NewResultsPtrCollection.DeserializeUsingResult(r)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingResult_Invalid(t *testing.T) {
	r := &Result{Bytes: []byte("invalid")}
	_, err := NewResultsPtrCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsPtrCollection_Default(t *testing.T) {
	rc := NewResultsPtrCollection.Default()
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsPtrCollection_AnyItemsPlusCap_Empty(t *testing.T) {
	rc := NewResultsPtrCollection.AnyItemsPlusCap(0)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsPtrCollection_AnyItemsPlusCap_Valid(t *testing.T) {
	rc := NewResultsPtrCollection.AnyItemsPlusCap(0, "a", "b")
	if rc.Length() != 2 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_AnyItems(t *testing.T) {
	rc := NewResultsPtrCollection.AnyItems("a")
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_UsingResultsPlusCap_Empty(t *testing.T) {
	rc := NewResultsPtrCollection.UsingResultsPlusCap(0)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsPtrCollection_UsingResultsPlusCap_Valid(t *testing.T) {
	r := NewPtr("a")
	rc := NewResultsPtrCollection.UsingResultsPlusCap(0, r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_UsingResults(t *testing.T) {
	r := NewPtr("a")
	rc := NewResultsPtrCollection.UsingResults(r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_Serializers_Empty(t *testing.T) {
	rc := NewResultsPtrCollection.Serializers()
	if !rc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResultsPtrCollection_Serializers_Valid(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	rc := NewResultsPtrCollection.Serializers(s)
	if rc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// --- newMapResultsCreator ---

func Test_I17_NewMapResults_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(MapResults{Items: map[string]Result{}})
	mr, err := NewMapResults.UnmarshalUsingBytes(b)
	if err != nil || mr == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := NewMapResults.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewMapResults_DeserializeUsingResult_Issues(t *testing.T) {
	r := &Result{Error: errors.New("bad")}
	_, err := NewMapResults.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewMapResults_DeserializeUsingResult_Valid(t *testing.T) {
	b := Serialize.ToBytesMust(MapResults{Items: map[string]Result{}})
	r := &Result{Bytes: b}
	mr, err := NewMapResults.DeserializeUsingResult(r)
	if err != nil || mr == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_DeserializeUsingResult_Invalid(t *testing.T) {
	r := &Result{Bytes: []byte("invalid")}
	_, err := NewMapResults.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewMapResults_UsingCap(t *testing.T) {
	mr := NewMapResults.UsingCap(5)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyAnyItems_Empty(t *testing.T) {
	mr := NewMapResults.UsingKeyAnyItems(0)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyAnyItems_Valid(t *testing.T) {
	mr := NewMapResults.UsingKeyAnyItems(0,
		KeyAny{Key: "k", AnyItem: "v"})
	if mr.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewMapResults_UsingMapOptions_EmptyMap(t *testing.T) {
	mr := NewMapResults.UsingMapOptions(false, false, 5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapOptions_NoChange(t *testing.T) {
	items := map[string]Result{"k": New("v")}
	mr := NewMapResults.UsingMapOptions(false, false, 0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewMapResults_UsingMapOptions_WithClone(t *testing.T) {
	items := map[string]Result{"k": New("v")}
	mr := NewMapResults.UsingMapOptions(true, false, 0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCap_Empty(t *testing.T) {
	mr := NewMapResults.UsingMapPlusCap(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCap_Valid(t *testing.T) {
	items := map[string]Result{"k": New("v")}
	mr := NewMapResults.UsingMapPlusCap(0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapClone_Empty(t *testing.T) {
	mr := NewMapResults.UsingMapPlusCapClone(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapClone_Valid(t *testing.T) {
	items := map[string]Result{"k": New("v")}
	mr := NewMapResults.UsingMapPlusCapClone(0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapDeepClone_Empty(t *testing.T) {
	mr := NewMapResults.UsingMapPlusCapDeepClone(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapDeepClone_Valid(t *testing.T) {
	items := map[string]Result{"k": New("v")}
	mr := NewMapResults.UsingMapPlusCapDeepClone(0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMap_Empty(t *testing.T) {
	mr := NewMapResults.UsingMap(nil)
	if mr == nil || !mr.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewMapResults_UsingMap_Valid(t *testing.T) {
	items := map[string]Result{"k": New("v")}
	mr := NewMapResults.UsingMap(items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapAnyItemsPlusCap_Empty(t *testing.T) {
	mr := NewMapResults.UsingMapAnyItemsPlusCap(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapAnyItemsPlusCap_Valid(t *testing.T) {
	mr := NewMapResults.UsingMapAnyItemsPlusCap(0, map[string]any{"k": "v"})
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapAnyItems(t *testing.T) {
	mr := NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingKeyWithResultsPlusCap_Nil(t *testing.T) {
	mr := NewMapResults.UsingKeyWithResultsPlusCap(5, nil...)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyWithResults(t *testing.T) {
	kwr := KeyWithResult{Key: "k", Result: New("v")}
	mr := NewMapResults.UsingKeyWithResults(kwr)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingKeyJsonersPlusCap_Nil(t *testing.T) {
	mr := NewMapResults.UsingKeyJsonersPlusCap(5, nil...)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyJsoners(t *testing.T) {
	// Just test nil jsoners path
	mr := NewMapResults.UsingKeyJsoners(nil...)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}
