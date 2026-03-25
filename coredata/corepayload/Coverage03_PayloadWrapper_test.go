package corepayload

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

func TestPayloadWrapper_BasicOps(t *testing.T) {
	pw := New.PayloadWrapper.Empty()
	if pw == nil {
		t.Fatal("expected non-nil")
	}
	if pw.IsNull() {
		t.Fatal("should not be null")
	}
	if pw.HasItems() {
		t.Fatal("should not have items")
	}
	if !pw.IsEmpty() {
		t.Fatal("should be empty")
	}
	if pw.HasAnyItem() {
		t.Fatal("should not have items")
	}
	if pw.Length() != 0 {
		t.Fatal("expected 0")
	}
	if pw.Count() != 0 {
		t.Fatal("expected 0")
	}

	var nilPW *PayloadWrapper
	if !nilPW.IsNull() {
		t.Fatal("nil should be null")
	}
	if nilPW.Length() != 0 {
		t.Fatal("nil length should be 0")
	}
}

func TestPayloadWrapper_WithPayloads(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes(
		"test", "1", "task", "cat", "entity",
		[]byte(`{"key":"value"}`),
	)
	if pw.IsEmpty() {
		t.Fatal("should not be empty")
	}
	if pw.Name != "test" {
		t.Fatal("name mismatch")
	}
	if pw.PayloadsString() == "" {
		t.Fatal("expected payloads string")
	}
	if pw.PayloadName() != "test" {
		t.Fatal("unexpected")
	}
	if pw.PayloadCategory() != "cat" {
		t.Fatal("unexpected")
	}
	if pw.PayloadTaskType() != "task" {
		t.Fatal("unexpected")
	}
	if pw.PayloadEntityType() != "entity" {
		t.Fatal("unexpected")
	}
}

func TestPayloadWrapper_IsChecks(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	if !pw.IsName("n") {
		t.Fatal("should match name")
	}
	if !pw.IsIdentifier("1") {
		t.Fatal("should match id")
	}
	if !pw.IsTaskTypeName("t") {
		t.Fatal("should match task")
	}
	if !pw.IsEntityType("e") {
		t.Fatal("should match entity")
	}
	if !pw.IsCategory("c") {
		t.Fatal("should match category")
	}
	if !pw.HasSingleRecord() {
		t.Fatal("should have single record")
	}
}

func TestPayloadWrapper_IdentifierInteger(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "42", "t", "c", "e", []byte(`"x"`))
	if pw.IdentifierInteger() != 42 {
		t.Fatal("expected 42")
	}
	if pw.IdentifierUnsignedInteger() != 42 {
		t.Fatal("expected 42")
	}
	if pw.IdString() != "42" {
		t.Fatal("expected 42")
	}
	if pw.IdInteger() != 42 {
		t.Fatal("expected 42")
	}
}

func TestPayloadWrapper_ErrorHandling(t *testing.T) {
	pw := New.PayloadWrapper.Empty()
	if pw.HasError() {
		t.Fatal("should not have error")
	}
	if !pw.IsEmptyError() {
		t.Fatal("error should be empty")
	}
	if pw.Error() != nil {
		t.Fatal("error should be nil")
	}
}

func TestPayloadWrapper_HasAttributes(t *testing.T) {
	pw := New.PayloadWrapper.Empty()
	if !pw.HasAttributes() {
		t.Fatal("should have attributes")
	}
	if pw.IsEmptyAttributes() {
		t.Fatal("should not be empty attributes")
	}
}

func TestPayloadWrapper_Json(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	s := pw.JsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
	_ = pw.JsonStringMust()
	_ = pw.String()
	_ = pw.PrettyJsonString()
	_ = pw.Json()
	_ = pw.JsonPtr()
	_ = pw.JsonModel()
	_ = pw.JsonModelAny()
}

func TestPayloadWrapper_MarshalUnmarshal(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"hello"`))
	b, err := pw.MarshalJSON()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}

	pw2 := New.PayloadWrapper.Empty()
	err2 := pw2.UnmarshalJSON(b)
	if err2 != nil {
		t.Fatal(err2)
	}
	if pw2.Name != "n" {
		t.Fatal("name mismatch after unmarshal")
	}
}

func TestPayloadWrapper_Serialize(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	b, err := pw.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
	_ = pw.SerializeMust()
}

func TestPayloadWrapper_Deserialize(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"hello"`))
	var out string
	err := pw.Deserialize(&out)
	if err != nil || out != "hello" {
		t.Fatal("unexpected")
	}

	pw.PayloadDeserialize(&out)
}

func TestPayloadWrapper_IsEqual(t *testing.T) {
	pw1 := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	pw2 := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	if !pw1.IsEqual(pw2) {
		t.Fatal("should be equal")
	}

	var nilPW *PayloadWrapper
	if !nilPW.IsEqual(nil) {
		t.Fatal("both nil should be equal")
	}
}

func TestPayloadWrapper_IsPayloadsEqual(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	if !pw.IsPayloadsEqual([]byte(`"x"`)) {
		t.Fatal("should be equal")
	}
}

func TestPayloadWrapper_Clone(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	c, err := pw.Clone(false)
	if err != nil {
		t.Fatal(err)
	}
	if c.Name != "n" {
		t.Fatal("clone mismatch")
	}

	c2, err2 := pw.Clone(true)
	if err2 != nil {
		t.Fatal(err2)
	}
	_ = c2

	var nilPW *PayloadWrapper
	cp, _ := nilPW.ClonePtr(false)
	if cp != nil {
		t.Fatal("expected nil")
	}
}

func TestPayloadWrapper_NonPtr_ToPtr(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	np := pw.NonPtr()
	_ = np

	var nilPW *PayloadWrapper
	np2 := nilPW.NonPtr()
	_ = np2

	p := np.ToPtr()
	_ = p
}

func TestPayloadWrapper_Clear_Dispose(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	pw.Clear()
	pw.Dispose()

	var nilPW *PayloadWrapper
	nilPW.Clear()
	nilPW.Dispose()
}

func TestPayloadWrapper_Value(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	v := pw.Value()
	if v == nil {
		t.Fatal("expected non-nil")
	}
}

func TestPayloadWrapper_DynamicPayloads(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	dp := pw.DynamicPayloads()
	if len(dp) == 0 {
		t.Fatal("expected bytes")
	}

	var nilPW *PayloadWrapper
	dp2 := nilPW.DynamicPayloads()
	if len(dp2) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func TestPayloadWrapper_SetDynamicPayloads(t *testing.T) {
	pw := New.PayloadWrapper.Empty()
	err := pw.SetDynamicPayloads([]byte(`"x"`))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPayloadWrapper_All(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	id, name, entity, cat, dp := pw.All()
	if id != "1" || name != "n" || entity != "e" || cat != "c" || len(dp) == 0 {
		t.Fatal("unexpected")
	}
}

func TestPayloadWrapper_AllSafe(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	id, name, entity, cat, dp := pw.AllSafe()
	if id != "1" || name != "n" || entity != "e" || cat != "c" || len(dp) == 0 {
		t.Fatal("unexpected")
	}

	var nilPW *PayloadWrapper
	_, _, _, _, dp2 := nilPW.AllSafe()
	if len(dp2) != 0 {
		t.Fatal("expected empty")
	}
}

func TestPayloadWrapper_HasSafeItems(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	if !pw.HasSafeItems() {
		t.Fatal("should have safe items")
	}
}

func TestPayloadWrapper_HasIssuesOrEmpty(t *testing.T) {
	pw := New.PayloadWrapper.Empty()
	if !pw.HasIssuesOrEmpty() {
		t.Fatal("empty should have issues")
	}
}

func TestPayloadWrapper_PayloadsJsonResult(t *testing.T) {
	pw := PayloadWrapper{Payloads: []byte(`"x"`)}
	r := pw.PayloadsJsonResult()
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func TestPayloadWrapper_PayloadsPrettyString(t *testing.T) {
	pw := PayloadWrapper{Payloads: []byte(`{"a":1}`)}
	s := pw.PayloadsPrettyString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestPayloadWrapper_ParseInjectUsingJson(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	jr := corejson.New(pw)
	target := New.PayloadWrapper.Empty()
	_, err := target.ParseInjectUsingJson(jr.Ptr())
	if err != nil {
		t.Fatal(err)
	}
}

func TestPayloadWrapper_BytesConverter(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	bc := pw.BytesConverter()
	if bc == nil {
		t.Fatal("expected non-nil")
	}
}

func TestPayloadWrapper_AsJsonContractsBinder(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"x"`))
	_ = pw.AsJsonContractsBinder()
}

func TestNewPayloadWrapperCreator(t *testing.T) {
	_ = New.PayloadWrapper.Empty()

	pw, err := New.PayloadWrapper.Create("n", "1", "t", "c", "hello")
	if err != nil || pw == nil {
		t.Fatal("unexpected")
	}

	pw2, err2 := New.PayloadWrapper.Record("n", "1", "t", "c", "hello")
	if err2 != nil || pw2 == nil {
		t.Fatal("unexpected")
	}

	pw3, err3 := New.PayloadWrapper.Records("n", "1", "t", "c", []string{"a", "b"})
	if err3 != nil || pw3 == nil {
		t.Fatal("unexpected")
	}

	pw4, err4 := New.PayloadWrapper.NameIdRecord("n", "1", "hello")
	if err4 != nil || pw4 == nil {
		t.Fatal("unexpected")
	}

	pw5, err5 := New.PayloadWrapper.NameIdCategory("n", "1", "c", "hello")
	if err5 != nil || pw5 == nil {
		t.Fatal("unexpected")
	}

	pw6, err6 := New.PayloadWrapper.NameIdTaskRecord("n", "1", "t", "hello")
	if err6 != nil || pw6 == nil {
		t.Fatal("unexpected")
	}

	pw7, err7 := New.PayloadWrapper.NameTaskNameRecord("1", "t", "hello")
	if err7 != nil || pw7 == nil {
		t.Fatal("unexpected")
	}
}

func TestNewPayloadWrapperCreator_Deserialize(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"hello"`))
	b, _ := pw.Serialize()
	pw2, err := New.PayloadWrapper.Deserialize(b)
	if err != nil || pw2 == nil {
		t.Fatal("unexpected")
	}
	if pw2.Name != "n" {
		t.Fatal("name mismatch")
	}

	_, err2 := New.PayloadWrapper.Deserialize([]byte("invalid"))
	if err2 == nil {
		t.Fatal("expected error")
	}
}

func TestNewPayloadWrapperCreator_DeserializeUsingJsonResult(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"hello"`))
	jr := pw.JsonPtr()
	pw2, err := New.PayloadWrapper.DeserializeUsingJsonResult(jr)
	if err != nil || pw2 == nil {
		t.Fatal("unexpected")
	}
}

func TestNewPayloadWrapperCreator_CastOrDeserializeFrom(t *testing.T) {
	pw := New.PayloadWrapper.UsingBytes("n", "1", "t", "c", "e", []byte(`"hello"`))
	pw2, err := New.PayloadWrapper.CastOrDeserializeFrom(pw)
	_ = pw2
	_ = err

	_, err2 := New.PayloadWrapper.CastOrDeserializeFrom(nil)
	if err2 == nil {
		t.Fatal("expected error for nil")
	}
}

func TestNewPayloadWrapperCreator_All(t *testing.T) {
	_ = New.PayloadWrapper.All("n", "1", "t", "c", "e", false, nil, []byte(`"x"`))
}
