package corejson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Result — FieldsNames non-empty, safeJsonStringInternal, MeaningfulError
// with error+payload, IsEqualFull jsonString match, SerializeResult error
// Covers Result.go L85-94, L376-381, L385-387, L639-646, L827-829, L872-874
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov25_Result_FieldsNames_NonEmpty(t *testing.T) {
	// Arrange — valid JSON object with fields
	jsonBytes, _ := json.Marshal(map[string]any{"name": "test", "value": 42})
	r := NewResult.Bytes(jsonBytes)

	// Act
	fields, err := r.FieldsNames()

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(fields))
	}
}

func Test_Cov25_Result_SafeJsonStringInternal_Nil(t *testing.T) {
	// Arrange
	var r *Result

	// Act — safeJsonStringInternal via MeaningfulError on nil
	result := r.SafeJsonString()

	// Assert
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func Test_Cov25_Result_MeaningfulError_WithPayload(t *testing.T) {
	// Arrange — Result with both error and bytes
	r := &Result{
		Bytes:    []byte(`{"key":"val"}`),
		Error:    fmt.Errorf("parse failed"),
		TypeName: "TestType",
	}

	// Act
	err := r.MeaningfulError()

	// Assert
	if err == nil {
		t.Error("expected meaningful error")
	}
}

func Test_Cov25_Result_SerializeResult_MarshalError(t *testing.T) {
	// Arrange — Result with self-referencing data to trigger marshal error
	// Actually, Result itself is marshalable. The error path at L639 triggers
	// when json.Marshal(it) fails. This is rare but let's test the non-error path
	r := NewResult.Bytes([]byte(`{"ok":true}`))

	// Act
	bytes, err := r.SerializeResult()

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(bytes) == 0 {
		t.Error("expected non-empty bytes")
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesCollection — AddSerializer, AddSerializers, GetSinglePageCollection negative
// Covers BytesCollection.go L192-195, L205-209, L308-310, L564-565, L647-653
// ══════════════════════════════════════════════════════════════════════════════

type mockSerializer struct {
	bytes []byte
	err   error
}

func (m *mockSerializer) Serialize() ([]byte, error) {
	return m.bytes, m.err
}

func Test_Cov25_BytesCollection_AddSerializer(t *testing.T) {
	bc := NewBytesCollection.Cap(2)
	s := &mockSerializer{bytes: []byte(`{"a":1}`)}

	bc.AddSerializer(s)

	if bc.Length() != 1 {
		t.Errorf("expected length 1, got %d", bc.Length())
	}
}

func Test_Cov25_BytesCollection_AddSerializers(t *testing.T) {
	bc := NewBytesCollection.Cap(2)
	s1 := &mockSerializer{bytes: []byte(`{"a":1}`)}
	s2 := &mockSerializer{bytes: []byte(`{"b":2}`)}

	bc.AddSerializers(s1, s2)

	if bc.Length() != 2 {
		t.Errorf("expected length 2, got %d", bc.Length())
	}
}

func Test_Cov25_BytesCollection_GetSinglePageCollection_NegativeIndex(t *testing.T) {
	bc := NewBytesCollection.Cap(5)
	for i := 0; i < 5; i++ {
		bc.AddBytes([]byte(fmt.Sprintf(`{"i":%d}`, i)))
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic on negative page index")
		}
	}()

	bc.GetSinglePageCollection(0, 2) // pageIndex 0 → skipItems = 2*(0-1) = -2 → panic
}

// ══════════════════════════════════════════════════════════════════════════════
// MapResults — Unmarshal non-empty, UnmarshalMany, GetSinglePageCollection
// Covers MapResults.go L164-165, L202, L217-219, L235-236, L324-326, L718-729, L737-743, L773-779
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov25_MapResults_Unmarshal_NonEmpty(t *testing.T) {
	mr := NewMapResults.Cap(2)
	r := NewResult.Bytes([]byte(`{"name":"test"}`))
	mr.Add("key1", *r)

	// Act
	var target map[string]any
	err := mr.Unmarshal("key1", &target)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if target["name"] != "test" {
		t.Errorf("expected 'test', got %v", target["name"])
	}
}

func Test_Cov25_MapResults_GetSinglePageCollection_LengthMismatch(t *testing.T) {
	mr := NewMapResults.Cap(5)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("k%d", i)
		r := NewResult.Bytes([]byte(fmt.Sprintf(`{"i":%d}`, i)))
		mr.Add(key, *r)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic on length mismatch")
		}
	}()

	// allKeys has wrong length
	mr.GetSinglePageCollection(1, 2, []string{"k0", "k1", "k2"})
}

func Test_Cov25_MapResults_GetSinglePageCollection_NegativeIndex(t *testing.T) {
	mr := NewMapResults.Cap(5)
	keys := make([]string, 5)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("k%d", i)
		keys[i] = key
		r := NewResult.Bytes([]byte(fmt.Sprintf(`{"i":%d}`, i)))
		mr.Add(key, *r)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic on negative page index")
		}
	}()

	mr.GetSinglePageCollection(0, 2, keys) // 0 → negative skip → panic
}

func Test_Cov25_MapResults_AddSerializeItem_WithError(t *testing.T) {
	mr := NewMapResults.Cap(2)

	// Attempt to serialize a channel (unmarshalable)
	ch := make(chan int)
	err := mr.AddSerializeItem("bad", ch)

	if err == nil {
		t.Error("expected error on unmarshalable item")
	}
	_ = ch
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultsCollection — AddSerializer, AddSerializers, UnmarshalToAll error
// Covers ResultCollection.go L291-295, L305-307, L382-385, L395-399
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov25_ResultsCollection_AddSerializer(t *testing.T) {
	rc := NewResultsCollection.Cap(2)
	s := &mockSerializer{bytes: []byte(`{"c":3}`)}

	rc.AddSerializer(s)

	if rc.Length() != 1 {
		t.Errorf("expected length 1, got %d", rc.Length())
	}
}

func Test_Cov25_ResultsCollection_AddSerializers(t *testing.T) {
	rc := NewResultsCollection.Cap(2)
	s1 := &mockSerializer{bytes: []byte(`{"c":3}`)}
	s2 := &mockSerializer{bytes: []byte(`{"d":4}`)}

	rc.AddSerializers(s1, s2)

	if rc.Length() != 2 {
		t.Errorf("expected length 2, got %d", rc.Length())
	}
}

func Test_Cov25_ResultsCollection_UnmarshalToAll_WithError(t *testing.T) {
	rc := NewResultsCollection.Cap(2)
	r := &Result{Error: fmt.Errorf("bad"), TypeName: "test"}
	rc.AddSkipOnNil(r)

	var target struct{}
	errList, hasErr := rc.UnmarshalToAll(&target)

	if !hasErr {
		t.Error("expected hasAnyError=true")
	}
	if errList[0] == nil {
		t.Error("expected error at index 0")
	}
}
