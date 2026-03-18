package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ── Result.go uncovered branches ──

func Test_Gap_Result_HandleErrorWithMsg_NoPanic(t *testing.T) {
	r := corejson.NewResult.Any("x")
	r.HandleErrorWithMsg("prefix msg")
}

func Test_Gap_Result_ParseInjectUsingJsonMust_NoPanic(t *testing.T) {
	inner := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"hello"`), TypeName: "test"})
	target := &corejson.Result{}
	_ = target.ParseInjectUsingJsonMust(inner.Ptr())
}

func Test_Gap_Result_Unmarshal_BadPayload(t *testing.T) {
	// Valid result with bytes that don't unmarshal to target type
	r := corejson.NewResult.AnyPtr("hello")
	var out int
	err := r.Unmarshal(&out)
	if err == nil {
		t.Fatal("expected unmarshal error for type mismatch")
	}
}

func Test_Gap_Result_UnmarshalSkipExistingIssues_BadPayload(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	var out int
	err := r.UnmarshalSkipExistingIssues(&out)
	if err == nil {
		t.Fatal("expected error for bad payload")
	}
}

func Test_Gap_Result_MeaningfulError_EmptyBytesWithError(t *testing.T) {
	// Has error AND empty bytes
	r := &corejson.Result{
		Bytes:    []byte{},
		Error:    errors.New("some err"),
		TypeName: "TestType",
	}
	err := r.MeaningfulError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Gap_Result_MeaningfulError_HasErrorAndPayload(t *testing.T) {
	// Has error AND has payload
	r := &corejson.Result{
		Bytes:    []byte(`"payload"`),
		Error:    errors.New("some err"),
		TypeName: "TestType",
	}
	err := r.MeaningfulError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Gap_Result_String_WithNilBytes(t *testing.T) {
	r := corejson.Result{}
	s := r.String()
	if s != "" {
		t.Fatal("expected empty for nil bytes")
	}
}

// ── deserializerLogic uncovered methods ──

func Test_Gap_Deserialize_UsingSerializerFuncTo(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var s string
	err := corejson.Deserialize.UsingSerializerFuncTo(fn, &s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func Test_Gap_Deserialize_UsingSerializerTo_Nil(t *testing.T) {
	var s string
	// nil serializer returns nil from NewResult.UsingSerializer
	// which causes Deserialize on nil result
	r := corejson.Deserialize.UsingSerializerTo(nil, &s)
	_ = r
}

// ── deserializeFromBytesTo uncovered ──

func Test_Gap_BytesTo_ResultCollection(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc, err := corejson.Deserialize.BytesTo.ResultCollection(b)
	if err != nil {
		t.Fatal(err)
	}
	if rc.Length() == 0 {
		t.Fatal("expected items")
	}
}

func Test_Gap_BytesTo_ResultCollectionMust(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc := corejson.Deserialize.BytesTo.ResultCollectionMust(b)
	if rc.Length() == 0 {
		t.Fatal("expected items")
	}
}

func Test_Gap_BytesTo_ResultsPtrCollection(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc, err := corejson.Deserialize.BytesTo.ResultsPtrCollection(b)
	if err != nil {
		t.Fatal(err)
	}
	_ = rc
}

func Test_Gap_BytesTo_ResultsPtrCollectionMust(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc := corejson.Deserialize.BytesTo.ResultsPtrCollectionMust(b)
	_ = rc
}

func Test_Gap_BytesTo_MapResults(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	b, _ := corejson.Serialize.Raw(m)
	mr, err := corejson.Deserialize.BytesTo.MapResults(b)
	if err != nil {
		t.Fatal(err)
	}
	_ = mr
}

func Test_Gap_BytesTo_MapResultsMust(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	b, _ := corejson.Serialize.Raw(m)
	mr := corejson.Deserialize.BytesTo.MapResultsMust(b)
	_ = mr
}

// ── deserializeFromResultTo uncovered ──

func Test_Gap_ResultTo_ResultCollection(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	rc, err := corejson.Deserialize.ResultTo.ResultCollection(jr)
	if err != nil {
		t.Fatal(err)
	}
	if rc.Length() == 0 {
		t.Fatal("expected items")
	}
}

func Test_Gap_ResultTo_ResultCollectionMust(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	rc := corejson.Deserialize.ResultTo.ResultCollectionMust(jr)
	if rc.Length() == 0 {
		t.Fatal("expected items")
	}
}

func Test_Gap_ResultTo_ResultsPtrCollection(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	_, err := corejson.Deserialize.ResultTo.ResultsPtrCollection(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Gap_ResultTo_ResultsPtrCollectionMust(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	_ = corejson.Deserialize.ResultTo.ResultsPtrCollectionMust(jr)
}

func Test_Gap_ResultTo_Result(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_, err := corejson.Deserialize.ResultTo.Result(jr)
	_ = err
}

func Test_Gap_ResultTo_ResultMust(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_ = corejson.Deserialize.ResultTo.ResultMust(jr)
}

func Test_Gap_ResultTo_ResultPtr(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_, err := corejson.Deserialize.ResultTo.ResultPtr(jr)
	_ = err
}

func Test_Gap_ResultTo_ResultPtrMust(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_ = corejson.Deserialize.ResultTo.ResultPtrMust(jr)
}

func Test_Gap_ResultTo_MapResults(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	jr := corejson.NewResult.AnyPtr(m)
	_, err := corejson.Deserialize.ResultTo.MapResults(jr)
	_ = err
}

func Test_Gap_ResultTo_MapResultsMust(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	jr := corejson.NewResult.AnyPtr(m)
	_ = corejson.Deserialize.ResultTo.MapResultsMust(jr)
}

func Test_Gap_ResultTo_Bytes(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_, err := corejson.Deserialize.ResultTo.Bytes(jr)
	_ = err
}

func Test_Gap_ResultTo_BytesMust(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	jr := corejson.NewResult.AnyPtr(inner)
	_ = corejson.Deserialize.ResultTo.BytesMust(jr)
}

// ── newResultCreator uncovered ──

func Test_Gap_NewResult_UsingBytesError_NonNil(t *testing.T) {
	// import coredata for BytesError is needed - use creator with valid data
	r := corejson.NewResult.Any("hello")
	be := r.BytesError()
	if be == nil {
		t.Fatal("expected non-nil BytesError")
	}
	r2 := corejson.NewResult.UsingBytesError(be)
	if r2.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_Gap_NewResult_UsingBytesError_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesError(nil)
	if r.Bytes != nil {
		t.Fatal("expected nil bytes")
	}
}

func Test_Gap_NewResult_DeserializeUsingResult_WithIssues(t *testing.T) {
	errResult := corejson.NewResult.ErrorPtr(errors.New("bad"))
	r := corejson.NewResult.DeserializeUsingResult(errResult)
	if r == nil || !r.HasError() {
		t.Fatal("expected error result")
	}
}

func Test_Gap_NewResult_FromStringer(t *testing.T) {
	// errors.New implements fmt.Stringer via Error()
	// But we can use a simple stringer
	r := corejson.Serialize.FromStringer(errors.New("hello"))
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

// ── castingAny uncovered ──

func Test_Gap_CastAny_FromToOption_NilFrom(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(true, nil, &out)
	// nil from should return an error (not applicable) and fall through
	_ = err
}

func Test_Gap_CastAny_FromToOption_NilTo(t *testing.T) {
	err := corejson.CastAny.FromToOption(true, "hello", nil)
	_ = err
}

func Test_Gap_CastAny_FromToOption_Error(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(false, errors.New(`"hello"`), &out)
	if err != nil {
		t.Fatal("expected nil for error-to-string deserialization")
	}
	if out != "hello" {
		t.Fatal("expected hello, got", out)
	}
}

func Test_Gap_CastAny_FromToOption_ErrorInvalidJson(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(false, errors.New("not json"), &out)
	if err == nil {
		t.Fatal("expected error for invalid json in error message")
	}
}

func Test_Gap_CastAny_FromToOption_NilError(t *testing.T) {
	var nilErr error
	var out string
	err := corejson.CastAny.FromToOption(false, nilErr, &out)
	// nil error case goes to fallback serialization
	_ = err
}

// ── Deserialize.Result / ResultPtr / ResultMust / ResultPtrMust (on deserializerLogic struct) ──

func Test_Gap_Deserialize_Result(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	b, _ := r.Serialize()
	_, err := corejson.Deserialize.Result(b)
	_ = err
}

func Test_Gap_Deserialize_ResultPtr(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	b, _ := r.Serialize()
	_, err := corejson.Deserialize.ResultPtr(b)
	_ = err
}

// ── newResultsCollectionCreator uncovered ──

func Test_Gap_NewResultsCollection_DeserializeUsingResult(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	jr := corejson.NewResult.AnyPtr(c)
	rc, err := corejson.NewResultsCollection.DeserializeUsingResult(jr)
	if err != nil {
		t.Fatal(err)
	}
	if rc.Length() == 0 {
		t.Fatal("expected items")
	}
}

func Test_Gap_NewResultsCollection_DeserializeUsingResult_Issues(t *testing.T) {
	errResult := corejson.NewResult.ErrorPtr(errors.New("bad"))
	_, err := corejson.NewResultsCollection.DeserializeUsingResult(errResult)
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Gap_NewResultsCollection_UnmarshalUsingBytes(t *testing.T) {
	c := corejson.NewResultsCollection.Empty()
	c.Add(corejson.NewResult.Any("x"))
	b, _ := corejson.Serialize.Raw(c)
	rc, err := corejson.NewResultsCollection.UnmarshalUsingBytes(b)
	if err != nil {
		t.Fatal(err)
	}
	if rc.Length() == 0 {
		t.Fatal("expected items")
	}
}

// ── newResultsPtrCollectionCreator uncovered ──

func Test_Gap_NewResultsPtrCollection_DeserializeUsingResult(t *testing.T) {
	c := corejson.NewResultsPtrCollection.Default()
	c.Add(corejson.NewResult.AnyPtr("x"))
	jr := corejson.NewResult.AnyPtr(c)
	_, err := corejson.NewResultsPtrCollection.DeserializeUsingResult(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Gap_NewResultsPtrCollection_Jsoners(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Jsoners()
}

func Test_Gap_NewResultsPtrCollection_JsonersPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.JsonersPlusCap(true, 5)
}

// ── newMapResultsCreator uncovered ──

func Test_Gap_NewMapResults_DeserializeUsingResult(t *testing.T) {
	m := corejson.NewMapResults.Empty()
	m.Add("k", corejson.NewResult.Any("v"))
	jr := corejson.NewResult.AnyPtr(m)
	_, err := corejson.NewMapResults.DeserializeUsingResult(jr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Gap_NewMapResults_DeserializeUsingResult_Issues(t *testing.T) {
	errResult := corejson.NewResult.ErrorPtr(errors.New("bad"))
	_, err := corejson.NewMapResults.DeserializeUsingResult(errResult)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ── newBytesCollectionCreator uncovered ──

func Test_Gap_NewBytesCollection_DeserializeUsingResult(t *testing.T) {
	c := corejson.NewBytesCollection.Empty()
	c.Add([]byte(`"x"`))
	jr := corejson.NewResult.AnyPtr(c)
	_, err := corejson.NewBytesCollection.DeserializeUsingResult(jr)
	_ = err
}

func Test_Gap_NewBytesCollection_Jsoners(t *testing.T) {
	_ = corejson.NewBytesCollection.Jsoners()
}

// ── Serializer uncovered: FromStringer ──

func Test_Gap_Serialize_Apply_MarshalError(t *testing.T) {
	ch := make(chan int)
	r := corejson.Serialize.Apply(ch)
	if !r.HasError() {
		t.Fatal("expected error for channel")
	}
}
