package corejson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Result — FieldsNames non-empty, safeJsonStringInternal nil, MeaningfulError
// Covers Result.go L85-94, L376-381, L385-387, L639-646
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov25_Result_FieldsNames_NonEmpty(t *testing.T) {
	jsonBytes, _ := json.Marshal(map[string]any{"name": "test", "value": 42})
	r := NewResult.UsingBytes(jsonBytes)

	fields, err := r.FieldsNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(fields))
	}
}

func Test_Cov25_Result_SafeJsonStringInternal_Nil(t *testing.T) {
	var r *Result
	result := r.safeJsonStringInternal()
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func Test_Cov25_Result_MeaningfulError_WithPayload(t *testing.T) {
	r := &Result{
		Bytes:    []byte(`{"key":"val"}`),
		Error:    fmt.Errorf("parse failed"),
		TypeName: "TestType",
	}
	err := r.MeaningfulError()
	if err == nil {
		t.Error("expected meaningful error")
	}
}

func Test_Cov25_Result_Serialize(t *testing.T) {
	r := NewResult.UsingBytes([]byte(`{"ok":true}`))
	bytes, err := r.Serialize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(bytes) == 0 {
		t.Error("expected non-empty bytes")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Result — IsEqualFull jsonString cached match, IsEqual jsonString cached
// Covers Result.go L827-829, L872-874
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov25_Result_IsEqual_JsonStringCached(t *testing.T) {
	// Create two results sharing the same jsonString pointer
	sharedStr := `{"x":1}`
	r1 := Result{Bytes: []byte(sharedStr), jsonString: &sharedStr}
	r2 := Result{Bytes: []byte(sharedStr), jsonString: &sharedStr}

	// IsEqual should use the cached jsonString pointer equality
	result := r1.IsEqual(r2)
	if !result {
		t.Error("expected IsEqual=true for same jsonString pointer")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesCollection — AddSerializer, AddSerializers, GetSinglePageCollection
// Covers BytesCollection.go L192-195, L205-209, L647-653
// ══════════════════════════════════════════════════════════════════════════════

type mockSerializer struct {
	bytes []byte
	err   error
}

func (m *mockSerializer) Serialize() ([]byte, error) {
	return m.bytes, m.err
}

func Test_Cov25_BytesCollection_AddSerializer(t *testing.T) {
	bc := NewBytesCollection.UsingCap(2)
	s := &mockSerializer{bytes: []byte(`{"a":1}`)}
	bc.AddSerializer(s)
	if bc.Length() != 1 {
		t.Errorf("expected length 1, got %d", bc.Length())
	}
}

func Test_Cov25_BytesCollection_AddSerializers(t *testing.T) {
	bc := NewBytesCollection.UsingCap(2)
	s1 := &mockSerializer{bytes: []byte(`{"a":1}`)}
	s2 := &mockSerializer{bytes: []byte(`{"b":2}`)}
	bc.AddSerializers(s1, s2)
	if bc.Length() != 2 {
		t.Errorf("expected length 2, got %d", bc.Length())
	}
}

func Test_Cov25_BytesCollection_GetSinglePageCollection_NegativeIndex(t *testing.T) {
	bc := NewBytesCollection.UsingCap(5)
	for i := 0; i < 5; i++ {
		bc.Add([]byte(fmt.Sprintf(`{"i":%d}`, i)))
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic on negative page index")
		}
	}()

	bc.GetSinglePageCollection(0, 2) // pageIndex 0 → skip = 2*(0-1) = -2 → panic
}

// ══════════════════════════════════════════════════════════════════════════════
// MapResults — Unmarshal, AddAnySkipOnNil error, GetSinglePageCollection
// Covers MapResults.go L164-165, L202, L324-326, L718-729, L737-743
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov25_MapResults_Unmarshal(t *testing.T) {
	mr := NewMapResults.UsingCap(2)
	r := NewResult.UsingBytes([]byte(`{"name":"test"}`))
	mr.Add("key1", *r)

	var target map[string]any
	err := mr.Unmarshal("key1", &target)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if target["name"] != "test" {
		t.Errorf("expected 'test', got %v", target["name"])
	}
}

func Test_Cov25_MapResults_AddAnySkipOnNil_Error(t *testing.T) {
	mr := NewMapResults.UsingCap(2)
	// Channel is not marshalable
	ch := make(chan int)
	err := mr.AddAnySkipOnNil("bad", ch)
	if err == nil {
		t.Error("expected error on unmarshalable item")
	}
	close(ch)
}

func Test_Cov25_MapResults_GetSinglePageCollection_LengthMismatch(t *testing.T) {
	mr := NewMapResults.UsingCap(5)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("k%d", i)
		r := NewResult.UsingBytes([]byte(fmt.Sprintf(`{"i":%d}`, i)))
		mr.Add(key, *r)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic on length mismatch")
		}
	}()

	mr.GetSinglePageCollection(1, 2, []string{"k0", "k1", "k2"})
}

func Test_Cov25_MapResults_GetSinglePageCollection_NegativeIndex(t *testing.T) {
	mr := NewMapResults.UsingCap(5)
	keys := make([]string, 5)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("k%d", i)
		keys[i] = key
		r := NewResult.UsingBytes([]byte(fmt.Sprintf(`{"i":%d}`, i)))
		mr.Add(key, *r)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic on negative page index")
		}
	}()

	mr.GetSinglePageCollection(0, 2, keys)
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultsCollection — AddSerializer, AddSerializers, UnmarshalIntoSameIndex error
// Covers ResultCollection.go L291-295, L305-307, L382-385, L395-399
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov25_ResultsCollection_AddSerializer(t *testing.T) {
	rc := NewResultsCollection.UsingCap(2)
	s := &mockSerializer{bytes: []byte(`{"c":3}`)}
	rc.AddSerializer(s)
	if rc.Length() != 1 {
		t.Errorf("expected length 1, got %d", rc.Length())
	}
}

func Test_Cov25_ResultsCollection_AddSerializers(t *testing.T) {
	rc := NewResultsCollection.UsingCap(2)
	s1 := &mockSerializer{bytes: []byte(`{"c":3}`)}
	s2 := &mockSerializer{bytes: []byte(`{"d":4}`)}
	rc.AddSerializers(s1, s2)
	if rc.Length() != 2 {
		t.Errorf("expected length 2, got %d", rc.Length())
	}
}

func Test_Cov25_ResultsCollection_UnmarshalIntoSameIndex_WithError(t *testing.T) {
	rc := NewResultsCollection.UsingCap(2)
	r := &Result{Error: fmt.Errorf("bad"), TypeName: "test"}
	rc.AddSkipOnNil(r)

	var target struct{}
	errList, hasErr := rc.UnmarshalIntoSameIndex(&target)
	if !hasErr {
		t.Error("expected hasAnyError=true")
	}
	if errList[0] == nil {
		t.Error("expected error at index 0")
	}
}
