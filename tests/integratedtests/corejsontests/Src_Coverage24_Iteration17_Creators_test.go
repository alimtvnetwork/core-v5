package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"errors"
	"testing"
)

// Covers: newBytesCollectionCreator, newResultsCollectionCreator,
// newResultsPtrCollectionCreator, newMapResultsCreator

// --- newBytesCollectionCreator ---

func Test_I17_NewBytesCollection_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(BytesCollection{Items: [][]byte{[]byte(`"a"`)}})
	bc, err := corejson.NewBytesCollection.UnmarshalUsingBytes(b)
	if err != nil || bc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := corejson.NewBytesCollection.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingResult_Issues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("bad")}
	_, err := corejson.NewBytesCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingResult_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(BytesCollection{Items: [][]byte{[]byte(`"a"`)}})
	r := &corejson.Result{Bytes: b}
	bc, err := corejson.NewBytesCollection.DeserializeUsingResult(r)
	if err != nil || bc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewBytesCollection_DeserializeUsingResult_InvalidPayload(t *testing.T) {
	r := &corejson.Result{Bytes: []byte("invalid")}
	_, err := corejson.NewBytesCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewBytesCollection_Empty(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	if bc == nil || !bc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewBytesCollection_UsingCap(t *testing.T) {
	bc := corejson.NewBytesCollection.UsingCap(10)
	if bc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewBytesCollection_AnyItems(t *testing.T) {
	bc, err := corejson.NewBytesCollection.AnyItems("a", "b")
	if err != nil || bc.Length() != 2 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewBytesCollection_Serializers_Empty(t *testing.T) {
	bc := corejson.NewBytesCollection.Serializers()
	if !bc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewBytesCollection_Serializers_Valid(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	bc := corejson.NewBytesCollection.Serializers(s)
	if bc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// --- newResultsCollectionCreator ---

func Test_I17_NewResultsCollection_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsCollection{Items: []corejson.Result{}})
	rc, err := corejson.NewResultsCollection.UnmarshalUsingBytes(b)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := corejson.NewResultsCollection.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingResult_Issues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("bad")}
	_, err := corejson.NewResultsCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingResult_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsCollection{Items: []corejson.Result{}})
	r := &corejson.Result{Bytes: b}
	rc, err := corejson.NewResultsCollection.DeserializeUsingResult(r)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsCollection_DeserializeUsingResult_InvalidPayload(t *testing.T) {
	r := &corejson.Result{Bytes: []byte("invalid")}
	_, err := corejson.NewResultsCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsCollection_Default(t *testing.T) {
	rc := corejson.NewResultsCollection.Default()
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingCap(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(5)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_AnyItems(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItems("a", "b")
	if rc.Length() != 2 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_AnyItemsPlusCap_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItemsPlusCap(5)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_AnyItemsPlusCap_Valid(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItemsPlusCap(5, "a")
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingJsonersOption_NilJsoners(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingJsonersOption(true, 5, nil...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtrPlusCap_Nil(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingResultsPtrPlusCap(5, nil...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtrPlusCap_Empty(t *testing.T) {
	results := []*corejson.Result{}
	rc := corejson.NewResultsCollection.UsingResultsPtrPlusCap(5, results...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtrPlusCap_Valid(t *testing.T) {
	r := corejson.NewPtr("a")
	rc := corejson.NewResultsCollection.UsingResultsPtrPlusCap(0, r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPtr(t *testing.T) {
	r := corejson.NewPtr("a")
	rc := corejson.NewResultsCollection.UsingResultsPtr(r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPlusCap_Nil(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingResultsPlusCap(5, nil...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPlusCap_Empty(t *testing.T) {
	results := []corejson.Result{}
	rc := corejson.NewResultsCollection.UsingResultsPlusCap(5, results...)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsCollection_UsingResultsPlusCap_Valid(t *testing.T) {
	r := corejson.New("a")
	rc := corejson.NewResultsCollection.UsingResultsPlusCap(0, r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_UsingResults(t *testing.T) {
	r := corejson.New("a")
	rc := corejson.NewResultsCollection.UsingResults(r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsCollection_Serializers_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Serializers()
	if !rc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResultsCollection_Serializers_Valid(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	rc := corejson.NewResultsCollection.Serializers(s)
	if rc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

func Test_I17_NewResultsCollection_SerializerFunctions_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.SerializerFunctions()
	if !rc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResultsCollection_SerializerFunctions_Valid(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"x"`), nil }
	rc := corejson.NewResultsCollection.SerializerFunctions(fn)
	if rc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// --- newResultsPtrCollectionCreator ---

func Test_I17_NewResultsPtrCollection_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsPtrCollection{Items: []*corejson.Result{}})
	rc, err := corejson.NewResultsPtrCollection.UnmarshalUsingBytes(b)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := corejson.NewResultsPtrCollection.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingResult_Issues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("bad")}
	_, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingResult_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(ResultsPtrCollection{Items: []*corejson.Result{}})
	r := &corejson.Result{Bytes: b}
	rc, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(r)
	if err != nil || rc == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResultsPtrCollection_DeserializeUsingResult_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte("invalid")}
	_, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResultsPtrCollection_Default(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.Default()
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsPtrCollection_AnyItemsPlusCap_Empty(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.AnyItemsPlusCap(0)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsPtrCollection_AnyItemsPlusCap_Valid(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.AnyItemsPlusCap(0, "a", "b")
	if rc.Length() != 2 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_AnyItems(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.AnyItems("a")
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_UsingResultsPlusCap_Empty(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.UsingResultsPlusCap(0)
	if rc == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResultsPtrCollection_UsingResultsPlusCap_Valid(t *testing.T) {
	r := corejson.NewPtr("a")
	rc := corejson.NewResultsPtrCollection.UsingResultsPlusCap(0, r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_UsingResults(t *testing.T) {
	r := corejson.NewPtr("a")
	rc := corejson.NewResultsPtrCollection.UsingResults(r)
	if rc.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewResultsPtrCollection_Serializers_Empty(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.Serializers()
	if !rc.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResultsPtrCollection_Serializers_Valid(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	rc := corejson.NewResultsPtrCollection.Serializers(s)
	if rc.IsEmpty() {
		t.Fatal("expected non-empty")
	}
}

// --- newMapResultsCreator ---

func Test_I17_NewMapResults_UnmarshalUsingBytes_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(MapResults{Items: map[string]corejson.Result{}})
	mr, err := corejson.NewMapResults.UnmarshalUsingBytes(b)
	if err != nil || mr == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_DeserializeUsingBytes_Invalid(t *testing.T) {
	_, err := corejson.NewMapResults.DeserializeUsingBytes([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewMapResults_DeserializeUsingResult_Issues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("bad")}
	_, err := corejson.NewMapResults.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewMapResults_DeserializeUsingResult_Valid(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(MapResults{Items: map[string]corejson.Result{}})
	r := &corejson.Result{Bytes: b}
	mr, err := corejson.NewMapResults.DeserializeUsingResult(r)
	if err != nil || mr == nil {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_DeserializeUsingResult_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte("invalid")}
	_, err := corejson.NewMapResults.DeserializeUsingResult(r)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewMapResults_UsingCap(t *testing.T) {
	mr := corejson.NewMapResults.UsingCap(5)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyAnyItems_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyAnyItems(0)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyAnyItems_Valid(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyAnyItems(0,
		corejson.KeyAny{Key: "k", AnyInf: "v"})
	if mr.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewMapResults_UsingMapOptions_EmptyMap(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapOptions(false, false, 5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapOptions_NoChange(t *testing.T) {
	items := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapOptions(false, false, 0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewMapResults_UsingMapOptions_WithClone(t *testing.T) {
	items := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapOptions(true, false, 0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected length")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCap_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapPlusCap(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCap_Valid(t *testing.T) {
	items := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapPlusCap(0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapClone_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapPlusCapClone(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapClone_Valid(t *testing.T) {
	items := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapPlusCapClone(0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapDeepClone_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapPlusCapDeepClone(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapPlusCapDeepClone_Valid(t *testing.T) {
	items := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMapPlusCapDeepClone(0, items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMap_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMap(nil)
	if mr == nil || !mr.IsEmpty() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewMapResults_UsingMap_Valid(t *testing.T) {
	items := map[string]corejson.Result{"k": corejson.New("v")}
	mr := corejson.NewMapResults.UsingMap(items)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapAnyItemsPlusCap_Empty(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapAnyItemsPlusCap(5, nil)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingMapAnyItemsPlusCap_Valid(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapAnyItemsPlusCap(0, map[string]any{"k": "v"})
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingMapAnyItems(t *testing.T) {
	mr := corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingKeyWithResultsPlusCap_Nil(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyWithResultsPlusCap(5, nil...)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyWithResults(t *testing.T) {
	kwr := corejson.KeyWithResult{Key: "k", corejson.Result: corejson.New("v")}
	mr := corejson.NewMapResults.UsingKeyWithResults(kwr)
	if mr.Length() != 1 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewMapResults_UsingKeyJsonersPlusCap_Nil(t *testing.T) {
	mr := corejson.NewMapResults.UsingKeyJsonersPlusCap(5, nil...)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewMapResults_UsingKeyJsoners(t *testing.T) {
	// Just test nil jsoners path
	mr := corejson.NewMapResults.UsingKeyJsoners(nil...)
	if mr == nil {
		t.Fatal("expected non-nil")
	}
}
