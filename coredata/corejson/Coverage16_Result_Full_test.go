package corejson

import (
	"encoding/json"
	"errors"
	"testing"
)

// ── Result extended methods ──

func TestResult_BytesTypeName(t *testing.T) {
	r := NewResult.Any("hello")
	if r.BytesTypeName() == "" { t.Fatal("expected non-empty") }
}

func TestResult_BytesTypeName_Nil(t *testing.T) {
	var r *Result
	if r.BytesTypeName() != "" { t.Fatal("expected empty") }
}

func TestResult_SafeBytesTypeName(t *testing.T) {
	r := NewResult.Any("hello")
	if r.SafeBytesTypeName() == "" { t.Fatal("expected non-empty") }
}

func TestResult_SafeBytesTypeName_Empty(t *testing.T) {
	r := Empty.ResultPtr()
	if r.SafeBytesTypeName() != "" { t.Fatal("expected empty") }
}

func TestResult_SafeString(t *testing.T) {
	r := NewResult.Any("hello")
	if r.SafeString() == "" { t.Fatal("expected non-empty") }
}

func TestResult_JsonStringPtr_Nil(t *testing.T) {
	var r *Result
	s := r.JsonStringPtr()
	if s == nil || *s != "" { t.Fatal("expected empty string ptr") }
}

func TestResult_JsonStringPtr_Cached(t *testing.T) {
	r := NewResult.Any("hello")
	s1 := r.JsonStringPtr()
	s2 := r.JsonStringPtr()
	if s1 != s2 { t.Fatal("expected same ptr") }
}

func TestResult_PrettyJsonBuffer(t *testing.T) {
	r := NewResult.Any(map[string]string{"a": "1"})
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() == 0 { t.Fatal("unexpected") }
}

func TestResult_PrettyJsonBuffer_Empty(t *testing.T) {
	r := Empty.ResultPtr()
	buf, _ := r.PrettyJsonBuffer("", "  ")
	if buf.Len() != 0 { t.Fatal("expected empty") }
}

func TestResult_PrettyJsonString(t *testing.T) {
	r := NewResult.Any(map[string]string{"a": "1"})
	if r.PrettyJsonString() == "" { t.Fatal("expected non-empty") }
}

func TestResult_PrettyJsonString_Nil(t *testing.T) {
	var r *Result
	if r.PrettyJsonString() != "" { t.Fatal("expected empty") }
}

func TestResult_PrettyJsonStringOrErrString(t *testing.T) {
	r := NewResult.Any("hello")
	if r.PrettyJsonStringOrErrString() == "" { t.Fatal("expected non-empty") }
}

func TestResult_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *Result
	s := r.PrettyJsonStringOrErrString()
	if s == "" { t.Fatal("expected non-empty error msg") }
}

func TestResult_PrettyJsonStringOrErrString_Error(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	s := r.PrettyJsonStringOrErrString()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_Length(t *testing.T) {
	r := NewResult.Any("hello")
	if r.Length() == 0 { t.Fatal("expected > 0") }
}

func TestResult_Length_Nil(t *testing.T) {
	var r *Result
	if r.Length() != 0 { t.Fatal("expected 0") }
}

func TestResult_ErrorString(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	if r.ErrorString() != "test" { t.Fatal("unexpected") }
}

func TestResult_ErrorString_NoError(t *testing.T) {
	r := NewResult.Any("hello")
	if r.ErrorString() != "" { t.Fatal("expected empty") }
}

func TestResult_IsErrorEqual(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	if !r.IsErrorEqual(errors.New("test")) { t.Fatal("expected true") }
}

func TestResult_IsErrorEqual_BothNil(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.IsErrorEqual(nil) { t.Fatal("expected true") }
}

func TestResult_IsErrorEqual_OneNil(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	if r.IsErrorEqual(nil) { t.Fatal("expected false") }
}

func TestResult_String(t *testing.T) {
	r := NewResult.Any("hello")
	s := r.String()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_String_WithError(t *testing.T) {
	r := Result{Error: errors.New("test"), Bytes: []byte(`"x"`), TypeName: "test"}
	s := r.String()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_SafeNonIssueBytes(t *testing.T) {
	r := NewResult.Any("hello")
	if len(r.SafeNonIssueBytes()) == 0 { t.Fatal("expected non-empty") }
}

func TestResult_SafeNonIssueBytes_Issue(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	if len(r.SafeNonIssueBytes()) != 0 { t.Fatal("expected empty") }
}

func TestResult_SafeBytes(t *testing.T) {
	r := NewResult.Any("hello")
	if len(r.SafeBytes()) == 0 { t.Fatal("expected non-empty") }
}

func TestResult_SafeBytes_Nil(t *testing.T) {
	var r *Result
	if len(r.SafeBytes()) != 0 { t.Fatal("expected empty") }
}

func TestResult_Values(t *testing.T) {
	r := NewResult.Any("hello")
	if len(r.Values()) == 0 { t.Fatal("expected non-empty") }
}

func TestResult_SafeValues(t *testing.T) {
	r := NewResult.Any("hello")
	if len(r.SafeValues()) == 0 { t.Fatal("expected non-empty") }
}

func TestResult_SafeValuesPtr(t *testing.T) {
	r := NewResult.Any("hello")
	if len(r.SafeValuesPtr()) == 0 { t.Fatal("expected non-empty") }
}

func TestResult_Raw(t *testing.T) {
	r := NewResult.Any("hello")
	b, err := r.Raw()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestResult_Raw_Nil(t *testing.T) {
	var r *Result
	_, err := r.Raw()
	if err == nil { t.Fatal("expected error") }
}

func TestResult_RawString(t *testing.T) {
	r := NewResult.Any("hello")
	s, err := r.RawString()
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestResult_RawStringMust(t *testing.T) {
	r := NewResult.Any("hello")
	s := r.RawStringMust()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_RawErrString(t *testing.T) {
	r := NewResult.Any("hello")
	b, msg := r.RawErrString()
	if len(b) == 0 || msg != "" { t.Fatal("unexpected") }
}

func TestResult_RawPrettyString(t *testing.T) {
	r := NewResult.Any(map[string]string{"a": "1"})
	s, err := r.RawPrettyString()
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestResult_MeaningfulError_Nil(t *testing.T) {
	var r *Result
	if r.MeaningfulError() == nil { t.Fatal("expected error") }
}

func TestResult_MeaningfulError_NoBytes(t *testing.T) {
	r := &Result{}
	if r.MeaningfulError() == nil { t.Fatal("expected error") }
}

func TestResult_MeaningfulError_WithError(t *testing.T) {
	r := &Result{Error: errors.New("fail"), Bytes: []byte(`"x"`)}
	if r.MeaningfulError() == nil { t.Fatal("expected error") }
}

func TestResult_MeaningfulErrorMessage(t *testing.T) {
	r := NewResult.Any("hello")
	if r.MeaningfulErrorMessage() != "" { t.Fatal("expected empty") }
}

func TestResult_IsEmptyError(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.IsEmptyError() { t.Fatal("expected true") }
}

func TestResult_HasSafeItems(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.HasSafeItems() { t.Fatal("expected true") }
}

func TestResult_IsAnyNull(t *testing.T) {
	var r *Result
	if !r.IsAnyNull() { t.Fatal("expected true") }
}

func TestResult_HasIssuesOrEmpty(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	if !r.HasIssuesOrEmpty() { t.Fatal("expected true") }
}

func TestResult_HasBytes(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.HasBytes() { t.Fatal("expected true") }
}

func TestResult_HasJsonBytes(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.HasJsonBytes() { t.Fatal("expected true") }
}

func TestResult_IsEmptyJsonBytes(t *testing.T) {
	r := &Result{Bytes: []byte("{}")}
	if !r.IsEmptyJsonBytes() { t.Fatal("expected true for empty json") }
}

func TestResult_IsEmpty(t *testing.T) {
	r := Empty.ResultPtr()
	if !r.IsEmpty() { t.Fatal("expected true") }
}

func TestResult_HasAnyItem(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.HasAnyItem() { t.Fatal("expected true") }
}

func TestResult_HasJson(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.HasJson() { t.Fatal("expected true") }
}

func TestResult_InjectInto(t *testing.T) {
	r := NewResult.Any([]string{"a"})
	target := &Result{}
	err := r.InjectInto(target)
	// InjectInto calls JsonParseSelfInject which unmarshals
	_ = err
}

func TestResult_Deserialize(t *testing.T) {
	r := NewResult.Any([]string{"a", "b"})
	var target []string
	err := r.Deserialize(&target)
	if err != nil || len(target) != 2 { t.Fatal("unexpected") }
}

func TestResult_Unmarshal(t *testing.T) {
	r := NewResult.Any(map[string]string{"k": "v"})
	var target map[string]string
	err := r.Unmarshal(&target)
	if err != nil || target["k"] != "v" { t.Fatal("unexpected") }
}

func TestResult_Unmarshal_Nil(t *testing.T) {
	var r *Result
	var target string
	err := r.Unmarshal(&target)
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Unmarshal_WithExistingError(t *testing.T) {
	r := &Result{Error: errors.New("fail")}
	var target string
	err := r.Unmarshal(&target)
	if err == nil { t.Fatal("expected error") }
}

func TestResult_SerializeSkipExistingIssues(t *testing.T) {
	r := &Result{Error: errors.New("fail")}
	b, err := r.SerializeSkipExistingIssues()
	if b != nil || err != nil { t.Fatal("expected nil nil") }
}

func TestResult_Serialize(t *testing.T) {
	r := NewResult.Any("hello")
	b, err := r.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestResult_Serialize_Nil(t *testing.T) {
	var r *Result
	_, err := r.Serialize()
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Serialize_WithError(t *testing.T) {
	r := &Result{Error: errors.New("fail")}
	_, err := r.Serialize()
	if err == nil { t.Fatal("expected error") }
}

func TestResult_UnmarshalSkipExistingIssues(t *testing.T) {
	r := &Result{Error: errors.New("fail")}
	err := r.UnmarshalSkipExistingIssues(nil)
	if err != nil { t.Fatal("expected nil") }
}

func TestResult_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	r := NewResult.Any([]string{"a"})
	var target []string
	err := r.UnmarshalSkipExistingIssues(&target)
	if err != nil { t.Fatal("unexpected") }
}

func TestResult_UnmarshalResult(t *testing.T) {
	r := NewResult.Any(Result{TypeName: "test"})
	result, _ := r.UnmarshalResult()
	_ = result
}

func TestResult_JsonModel(t *testing.T) {
	r := NewResult.Any("hello")
	m := r.JsonModel()
	_ = m
}

func TestResult_JsonModel_Nil(t *testing.T) {
	var r *Result
	m := r.JsonModel()
	if m.Error == nil { t.Fatal("expected error") }
}

func TestResult_JsonModelAny(t *testing.T) {
	r := NewResult.Any("hello")
	_ = r.JsonModelAny()
}

func TestResult_HandleErrorWithMsg(t *testing.T) {
	defer func() { recover() }()
	r := &Result{Error: errors.New("fail")}
	r.HandleErrorWithMsg("test msg")
}

func TestResult_HandleError(t *testing.T) {
	defer func() { recover() }()
	r := &Result{Error: errors.New("fail")}
	r.HandleError()
}

func TestResult_MustBeSafe(t *testing.T) {
	r := NewResult.Any("hello")
	r.MustBeSafe() // should not panic
}

func TestResult_DeserializeMust(t *testing.T) {
	r := NewResult.Any([]string{"a"})
	var target []string
	r.DeserializeMust(&target)
	if len(target) != 1 { t.Fatal("expected 1") }
}

func TestResult_UnmarshalMust(t *testing.T) {
	r := NewResult.Any("hello")
	var target string
	r.UnmarshalMust(&target)
	if target != "hello" { t.Fatal("unexpected") }
}

func TestResult_SerializeMust(t *testing.T) {
	r := NewResult.Any("hello")
	b := r.SerializeMust()
	if len(b) == 0 { t.Fatal("expected non-empty") }
}

func TestResult_RawMust(t *testing.T) {
	r := NewResult.Any("hello")
	b := r.RawMust()
	if len(b) == 0 { t.Fatal("expected non-empty") }
}

func TestResult_ParseInjectUsingJsonMust(t *testing.T) {
	r := NewResult.Any(Result{TypeName: "test"})
	target := &Result{}
	_ = target.ParseInjectUsingJsonMust(&r)
}

func TestResult_CloneError(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	if r.CloneError() == nil { t.Fatal("expected error") }
}

func TestResult_CloneError_Nil(t *testing.T) {
	r := NewResult.Any("hello")
	if r.CloneError() != nil { t.Fatal("expected nil") }
}

func TestResult_Ptr(t *testing.T) {
	r := NewResult.Any("hello")
	p := r.Ptr()
	if p == nil { t.Fatal("expected non-nil") }
}

func TestResult_NonPtr(t *testing.T) {
	r := NewResult.Any("hello")
	n := r.NonPtr()
	_ = n
}

func TestResult_NonPtr_Nil(t *testing.T) {
	var r *Result
	n := r.NonPtr()
	if n.Error == nil { t.Fatal("expected error") }
}

func TestResult_ToPtr(t *testing.T) {
	r := NewResult.Any("hello")
	p := r.ToPtr()
	if p == nil { t.Fatal("expected non-nil") }
}

func TestResult_ToNonPtr(t *testing.T) {
	r := NewResult.Any("hello")
	n := r.ToNonPtr()
	_ = n
}

func TestResult_IsEqualPtr(t *testing.T) {
	r1 := NewResult.Any("hello")
	r2 := NewResult.Any("hello")
	if !r1.IsEqualPtr(&r2) { t.Fatal("expected true") }
}

func TestResult_IsEqualPtr_BothNil(t *testing.T) {
	var r1, r2 *Result
	if !r1.IsEqualPtr(r2) { t.Fatal("expected true") }
}

func TestResult_IsEqualPtr_SamePtr(t *testing.T) {
	r := NewResult.Any("hello")
	if !r.IsEqualPtr(&r) { t.Fatal("expected true") }
}

func TestResult_IsEqual(t *testing.T) {
	r1 := NewResult.Any("hello")
	r2 := NewResult.Any("hello")
	if !r1.IsEqual(r2) { t.Fatal("expected true") }
}

func TestResult_CombineErrorWithRefString(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	s := r.CombineErrorWithRefString("ref1")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_CombineErrorWithRefString_NoError(t *testing.T) {
	r := NewResult.Any("hello")
	s := r.CombineErrorWithRefString("ref1")
	if s != "" { t.Fatal("expected empty") }
}

func TestResult_CombineErrorWithRefError(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	err := r.CombineErrorWithRefError("ref1")
	if err == nil { t.Fatal("expected error") }
}

func TestResult_CombineErrorWithRefError_NoError(t *testing.T) {
	r := NewResult.Any("hello")
	err := r.CombineErrorWithRefError("ref1")
	if err != nil { t.Fatal("expected nil") }
}

func TestResult_BytesError(t *testing.T) {
	r := NewResult.Any("hello")
	be := r.BytesError()
	if be == nil { t.Fatal("expected non-nil") }
}

func TestResult_BytesError_Nil(t *testing.T) {
	var r *Result
	if r.BytesError() != nil { t.Fatal("expected nil") }
}

func TestResult_Dispose(t *testing.T) {
	r := NewResult.Any("hello")
	r.Dispose()
	if r.Bytes != nil { t.Fatal("expected nil") }
}

func TestResult_Dispose_Nil(t *testing.T) {
	var r *Result
	r.Dispose() // should not panic
}

func TestResult_CloneIf(t *testing.T) {
	r := NewResult.Any("hello")
	c := r.CloneIf(true, true)
	if c.Length() == 0 { t.Fatal("expected non-empty") }
}

func TestResult_CloneIf_NoClone(t *testing.T) {
	r := NewResult.Any("hello")
	c := r.CloneIf(false, false)
	if c.Length() == 0 { t.Fatal("expected non-empty") }
}

func TestResult_ClonePtr(t *testing.T) {
	r := NewResult.Any("hello")
	c := r.ClonePtr(true)
	if c == nil || c.Length() == 0 { t.Fatal("unexpected") }
}

func TestResult_ClonePtr_Nil(t *testing.T) {
	var r *Result
	if r.ClonePtr(true) != nil { t.Fatal("expected nil") }
}

func TestResult_Clone_DeepClone(t *testing.T) {
	r := NewResult.Any("hello")
	c := r.Clone(true)
	if c.Length() == 0 { t.Fatal("expected non-empty") }
}

func TestResult_Clone_ShallowClone(t *testing.T) {
	r := NewResult.Any("hello")
	c := r.Clone(false)
	if c.Length() == 0 { t.Fatal("expected non-empty") }
}

func TestResult_Clone_Empty(t *testing.T) {
	r := Result{}
	c := r.Clone(true)
	_ = c
}

func TestResult_AsJsonContractsBinder(t *testing.T) {
	r := NewResult.Any("hello")
	_ = r.AsJsonContractsBinder()
}

func TestResult_AsJsoner(t *testing.T) {
	r := NewResult.Any("hello")
	_ = r.AsJsoner()
}

func TestResult_JsonParseSelfInject(t *testing.T) {
	r := NewResult.Any("hello")
	err := r.JsonParseSelfInject(&r)
	_ = err
}

func TestResult_AsJsonParseSelfInjector(t *testing.T) {
	r := NewResult.Any("hello")
	_ = r.AsJsonParseSelfInjector()
}

func TestResult_FieldsNames(t *testing.T) {
	r := NewResult.Any(map[string]string{"a": "1"})
	names, _ := r.FieldsNames()
	_ = names
}

func TestResult_SafeFieldsNames(t *testing.T) {
	r := NewResult.Any(map[string]string{"a": "1"})
	names := r.SafeFieldsNames()
	_ = names
}

func TestResult_SafeDeserializedFieldsToMap(t *testing.T) {
	r := NewResult.Any(map[string]string{"a": "1"})
	m := r.SafeDeserializedFieldsToMap()
	_ = m
}

// ── Serializer extended ──

func TestSerialize_Raw(t *testing.T) {
	b, err := Serialize.Raw("hello")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestSerialize_Bytes(t *testing.T) {
	b, err := Serialize.Bytes("hello")
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestSerialize_String(t *testing.T) {
	s, err := Serialize.String("hello")
	if err != nil || s == "" { t.Fatal("unexpected") }
}

// ── Deserializer extended ──

func TestDeserialize_UsingBytes(t *testing.T) {
	b, _ := json.Marshal("hello")
	var target string
	err := Deserialize.UsingBytes(b, &target)
	if err != nil || target != "hello" { t.Fatal("unexpected") }
}

func TestDeserialize_UsingString(t *testing.T) {
	var target string
	err := Deserialize.UsingString(`"hello"`, &target)
	if err != nil || target != "hello" { t.Fatal("unexpected") }
}

// ── AnyTo ──

func TestAnyTo_JsonString(t *testing.T) {
	s := AnyTo.JsonString("hello")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestAnyTo_PrettyString(t *testing.T) {
	s := AnyTo.PrettyString(map[string]string{"a": "1"})
	if s == "" { t.Fatal("expected non-empty") }
}

// ── New / NewPtr ──

func TestNew_Basic(t *testing.T) {
	r := New("hello")
	if r.HasError() { t.Fatal("unexpected") }
}

func TestNewPtr_Basic(t *testing.T) {
	r := NewPtr("hello")
	if r.HasError() { t.Fatal("unexpected") }
}
