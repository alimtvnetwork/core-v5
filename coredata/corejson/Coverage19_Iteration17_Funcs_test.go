package corejson

import (
	"errors"
	"testing"
)

// Covers: BytesCloneIf, BytesToString, BytesToPrettyString,
// JsonString, JsonStringOrErrMsg, BytesDeepClone (already covered but edge cases)

func Test_I17_BytesCloneIf_DeepCloneTrue(t *testing.T) {
	src := []byte(`"hello"`)
	cloned := BytesCloneIf(true, src)

	if string(cloned) != string(src) {
		t.Fatal("expected deep clone content match")
	}

	// Verify it's a separate allocation
	cloned[0] = 'x'
	if src[0] == 'x' {
		t.Fatal("deep clone should not share memory")
	}
}

func Test_I17_BytesCloneIf_DeepCloneFalse(t *testing.T) {
	src := []byte(`"hello"`)
	cloned := BytesCloneIf(false, src)

	if len(cloned) != 0 {
		t.Fatal("expected empty when isDeepClone=false")
	}
}

func Test_I17_BytesCloneIf_EmptyInput(t *testing.T) {
	cloned := BytesCloneIf(true, []byte{})
	if len(cloned) != 0 {
		t.Fatal("expected empty for empty input")
	}

	cloned2 := BytesCloneIf(true, nil)
	if len(cloned2) != 0 {
		t.Fatal("expected empty for nil input")
	}
}

func Test_I17_BytesToString_Valid(t *testing.T) {
	s := BytesToString([]byte(`{"a":1}`))
	if s != `{"a":1}` {
		t.Fatalf("expected json string, got %s", s)
	}
}

func Test_I17_BytesToString_Empty(t *testing.T) {
	if BytesToString(nil) != "" {
		t.Fatal("expected empty for nil")
	}
	if BytesToString([]byte{}) != "" {
		t.Fatal("expected empty for empty")
	}
}

func Test_I17_BytesToPrettyString_Valid(t *testing.T) {
	s := BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" {
		t.Fatal("expected non-empty pretty string")
	}
}

func Test_I17_BytesToPrettyString_Empty(t *testing.T) {
	if BytesToPrettyString(nil) != "" {
		t.Fatal("expected empty for nil")
	}
}

func Test_I17_JsonString_Valid(t *testing.T) {
	s, err := JsonString("hello")
	if err != nil {
		t.Fatal("unexpected error")
	}
	if s != `"hello"` {
		t.Fatalf("expected quoted string, got %s", s)
	}
}

func Test_I17_JsonString_MarshalFail(t *testing.T) {
	_, err := JsonString(func() {})
	if err == nil {
		t.Fatal("expected error for unmarshalable type")
	}
}

func Test_I17_JsonStringOrErrMsg_Valid(t *testing.T) {
	s := JsonStringOrErrMsg("hello")
	if s != `"hello"` {
		t.Fatalf("unexpected result: %s", s)
	}
}

func Test_I17_JsonStringOrErrMsg_Error(t *testing.T) {
	s := JsonStringOrErrMsg(func() {})
	if s == "" {
		t.Fatal("expected error message string")
	}
}

// Covers: emptyCreator methods
func Test_I17_EmptyCreator_All(t *testing.T) {
	r := Empty.Result()
	if r.HasAnyItem() {
		t.Fatal("empty result should have no items")
	}

	rp := Empty.ResultPtr()
	if rp == nil {
		t.Fatal("expected non-nil ptr")
	}

	errResult := Empty.ResultWithErr("TestType", errors.New("test"))
	if errResult.Error == nil || errResult.TypeName != "TestType" {
		t.Fatal("expected error result")
	}

	errPtrResult := Empty.ResultPtrWithErr("TestType2", errors.New("test2"))
	if errPtrResult == nil || errPtrResult.Error == nil {
		t.Fatal("expected error ptr result")
	}

	bc := Empty.BytesCollection()
	if !bc.IsEmpty() {
		t.Fatal("expected empty bytes collection")
	}

	bcp := Empty.BytesCollectionPtr()
	if bcp == nil || !bcp.IsEmpty() {
		t.Fatal("expected empty bytes collection ptr")
	}

	rc := Empty.ResultsCollection()
	if rc == nil || !rc.IsEmpty() {
		t.Fatal("expected empty results collection")
	}

	rpc := Empty.ResultsPtrCollection()
	if rpc == nil || !rpc.IsEmpty() {
		t.Fatal("expected empty results ptr collection")
	}

	mr := Empty.MapResults()
	if mr == nil || !mr.IsEmpty() {
		t.Fatal("expected empty map results")
	}
}
