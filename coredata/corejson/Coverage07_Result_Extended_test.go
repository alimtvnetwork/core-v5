package corejson

import (
	"errors"
	"testing"
)

func TestResult_Map_Nil(t *testing.T) {
	var r *Result
	m := r.Map()
	if len(m) != 0 {
		t.Fatal("expected empty map")
	}
}

func TestResult_Map_WithData(t *testing.T) {
	r := &Result{
		Bytes:    []byte(`"hello"`),
		Error:    errors.New("test err"),
		TypeName: "string",
	}
	m := r.Map()
	if m["Bytes"] == "" {
		t.Fatal("expected bytes")
	}
	if m["Error"] == "" {
		t.Fatal("expected error")
	}
	if m["Type"] != "string" {
		t.Fatal("expected type")
	}
}

func TestResult_DeserializedFieldsToMap_Nil(t *testing.T) {
	var r *Result
	fm, err := r.DeserializedFieldsToMap()
	if err != nil || len(fm) != 0 {
		t.Fatal("expected empty")
	}
}

func TestResult_DeserializedFieldsToMap_Empty(t *testing.T) {
	r := &Result{Bytes: []byte{}}
	fm, err := r.DeserializedFieldsToMap()
	if err != nil || len(fm) != 0 {
		t.Fatal("expected empty")
	}
}

func TestResult_SafeDeserializedFieldsToMap(t *testing.T) {
	r := &Result{Bytes: []byte(`{"a":1}`)}
	_ = r.SafeDeserializedFieldsToMap()
}

func TestResult_FieldsNames_Empty(t *testing.T) {
	r := &Result{Bytes: []byte{}}
	names, err := r.FieldsNames()
	if err != nil || len(names) != 0 {
		t.Fatal("expected empty")
	}
}

func TestResult_SafeFieldsNames(t *testing.T) {
	r := &Result{Bytes: []byte(`{"a":1}`)}
	_ = r.SafeFieldsNames()
}

func TestResult_BytesTypeName_Nil(t *testing.T) {
	var r *Result
	if r.BytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}

func TestResult_BytesTypeName_Normal(t *testing.T) {
	r := &Result{TypeName: "test"}
	if r.BytesTypeName() != "test" {
		t.Fatal("expected test")
	}
}

func TestResult_SafeBytesTypeName_Empty(t *testing.T) {
	r := &Result{}
	if r.SafeBytesTypeName() != "" {
		t.Fatal("expected empty")
	}
}

func TestResult_SafeBytesTypeName_Normal(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`), TypeName: "test"}
	if r.SafeBytesTypeName() != "test" {
		t.Fatal("expected test")
	}
}

func TestResult_SafeString(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if r.SafeString() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestResult_JsonStringPtr_Nil(t *testing.T) {
	var r *Result
	s := r.JsonStringPtr()
	if s == nil || *s != "" {
		t.Fatal("expected empty string ptr")
	}
}

func TestResult_JsonStringPtr_Cached(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	s1 := r.JsonStringPtr()
	s2 := r.JsonStringPtr()
	if s1 != s2 {
		t.Fatal("expected same ptr")
	}
}

func TestResult_JsonStringPtr_NoBytes(t *testing.T) {
	r := &Result{}
	s := r.JsonStringPtr()
	if *s != "" {
		t.Fatal("expected empty")
	}
}

func TestResult_PrettyJsonBuffer_Empty(t *testing.T) {
	r := &Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() != 0 {
		t.Fatal("expected empty buffer")
	}
}

func TestResult_PrettyJsonBuffer_Normal(t *testing.T) {
	r := &Result{Bytes: []byte(`{"a":1}`)}
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() == 0 {
		t.Fatal("expected non-empty buffer")
	}
}

func TestResult_PrettyJsonString_Nil(t *testing.T) {
	var r *Result
	if r.PrettyJsonString() != "" {
		t.Fatal("expected empty")
	}
}

func TestResult_PrettyJsonString_Normal(t *testing.T) {
	r := &Result{Bytes: []byte(`{"a":1}`)}
	s := r.PrettyJsonString()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestResult_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *Result
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected error msg")
	}
}

func TestResult_PrettyJsonStringOrErrString_Error(t *testing.T) {
	r := &Result{Error: errors.New("oops")}
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected error string")
	}
}

func TestResult_PrettyJsonStringOrErrString_OK(t *testing.T) {
	r := &Result{Bytes: []byte(`{"a":1}`)}
	s := r.PrettyJsonStringOrErrString()
	if s == "" {
		t.Fatal("expected pretty json")
	}
}

func TestResult_Length_Nil(t *testing.T) {
	var r *Result
	if r.Length() != 0 {
		t.Fatal("expected 0")
	}
}

func TestResult_HasError(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	if !r.HasError() {
		t.Fatal("expected error")
	}
	r2 := &Result{}
	if r2.HasError() {
		t.Fatal("expected no error")
	}
}

func TestResult_ErrorString(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	if r.ErrorString() != "e" {
		t.Fatal("expected e")
	}
	r2 := &Result{}
	if r2.ErrorString() != "" {
		t.Fatal("expected empty")
	}
}

func TestResult_IsErrorEqual(t *testing.T) {
	e := errors.New("e")
	r := &Result{Error: e}
	if !r.IsErrorEqual(errors.New("e")) {
		t.Fatal("expected equal")
	}
	if r.IsErrorEqual(nil) {
		t.Fatal("expected not equal")
	}
	r2 := &Result{}
	if !r2.IsErrorEqual(nil) {
		t.Fatal("expected equal")
	}
	if r2.IsErrorEqual(e) {
		t.Fatal("expected not equal")
	}
}

func TestResult_String_Normal(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`), TypeName: "t"}
	s := r.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestResult_String_WithError(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "t"}
	s := r.String()
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestResult_SafeNonIssueBytes(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if len(r.SafeNonIssueBytes()) == 0 {
		t.Fatal("expected bytes")
	}
	r2 := &Result{Error: errors.New("e")}
	if len(r2.SafeNonIssueBytes()) != 0 {
		t.Fatal("expected empty")
	}
}

func TestResult_SafeBytes_Nil(t *testing.T) {
	var r *Result
	if len(r.SafeBytes()) != 0 {
		t.Fatal("expected empty")
	}
}

func TestResult_Values(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if len(r.Values()) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestResult_SafeValues_Nil(t *testing.T) {
	var r *Result
	if len(r.SafeValues()) != 0 {
		t.Fatal("expected empty")
	}
}

func TestResult_SafeValuesPtr(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	if len(r.SafeValuesPtr()) != 0 {
		t.Fatal("expected empty")
	}
}

func TestResult_Raw_Nil(t *testing.T) {
	var r *Result
	_, err := r.Raw()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestResult_Raw_Normal(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	b, err := r.Raw()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestResult_RawMust(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	b := r.RawMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestResult_RawString(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	s, err := r.RawString()
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestResult_RawStringMust(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	s := r.RawStringMust()
	if s == "" {
		t.Fatal("expected string")
	}
}

func TestResult_RawErrString(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	b, errStr := r.RawErrString()
	if len(b) == 0 || errStr != "" {
		t.Fatal("unexpected")
	}
}

func TestResult_RawPrettyString(t *testing.T) {
	r := &Result{Bytes: []byte(`{"a":1}`)}
	s, err := r.RawPrettyString()
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestResult_MeaningfulErrorMessage(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if r.MeaningfulErrorMessage() != "" {
		t.Fatal("expected empty")
	}
	r2 := &Result{Error: errors.New("e")}
	if r2.MeaningfulErrorMessage() == "" {
		t.Fatal("expected error msg")
	}
}

func TestResult_MeaningfulError_Nil(t *testing.T) {
	var r *Result
	if r.MeaningfulError() == nil {
		t.Fatal("expected error")
	}
}

func TestResult_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &Result{}
	if r.MeaningfulError() == nil {
		t.Fatal("expected error for empty")
	}
}

func TestResult_MeaningfulError_WithError(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`), Error: errors.New("e")}
	if r.MeaningfulError() == nil {
		t.Fatal("expected error")
	}
}

func TestResult_IsEmptyError(t *testing.T) {
	if !(&Result{}).IsEmptyError() {
		t.Fatal("expected true")
	}
	var r *Result
	if !r.IsEmptyError() {
		t.Fatal("expected true")
	}
}

func TestResult_HasSafeItems(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if !r.HasSafeItems() {
		t.Fatal("expected true")
	}
}

func TestResult_IsAnyNull(t *testing.T) {
	var r *Result
	if !r.IsAnyNull() {
		t.Fatal("expected true")
	}
	r2 := &Result{}
	if !r2.IsAnyNull() {
		t.Fatal("expected true")
	}
}

func TestResult_HasIssuesOrEmpty(t *testing.T) {
	var r *Result
	if !r.HasIssuesOrEmpty() {
		t.Fatal("expected true")
	}
}

func TestResult_HandleError_NoPanic(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	r.HandleError() // should not panic
}

func TestResult_MustBeSafe_NoPanic(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	r.MustBeSafe()
}

func TestResult_HasBytes(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if !r.HasBytes() {
		t.Fatal("expected true")
	}
}

func TestResult_HasJsonBytes(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if !r.HasJsonBytes() {
		t.Fatal("expected true")
	}
}

func TestResult_IsEmptyJsonBytes_Nil(t *testing.T) {
	var r *Result
	if !r.IsEmptyJsonBytes() {
		t.Fatal("expected true")
	}
}

func TestResult_IsEmptyJsonBytes_CurlyBraces(t *testing.T) {
	r := &Result{Bytes: []byte("{}")}
	if !r.IsEmptyJsonBytes() {
		t.Fatal("expected true for {}")
	}
}

func TestResult_IsEmpty(t *testing.T) {
	var r *Result
	if !r.IsEmpty() {
		t.Fatal("expected true")
	}
}

func TestResult_HasAnyItem(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	if !r.HasAnyItem() {
		t.Fatal("expected true")
	}
}

func TestResult_IsEmptyJson(t *testing.T) {
	r := &Result{Bytes: []byte("{}")}
	if !r.IsEmptyJson() {
		t.Fatal("expected true")
	}
}

func TestResult_HasJson(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if !r.HasJson() {
		t.Fatal("expected true")
	}
}

func TestResult_Deserialize(t *testing.T) {
	r := &Result{Bytes: []byte(`"hello"`)}
	var s string
	err := r.Deserialize(&s)
	if err != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func TestResult_DeserializeMust(t *testing.T) {
	r := &Result{Bytes: []byte(`"hello"`)}
	var s string
	r.DeserializeMust(&s)
	if s != "hello" {
		t.Fatal("unexpected")
	}
}

func TestResult_UnmarshalMust(t *testing.T) {
	r := &Result{Bytes: []byte(`42`)}
	var i int
	r.UnmarshalMust(&i)
	if i != 42 {
		t.Fatal("unexpected")
	}
}

func TestResult_Unmarshal_Nil(t *testing.T) {
	var r *Result
	err := r.Unmarshal(&struct{}{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestResult_Unmarshal_WithExistingError(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	err := r.Unmarshal(&struct{}{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestResult_Unmarshal_BadJSON(t *testing.T) {
	r := &Result{Bytes: []byte(`{invalid}`)}
	err := r.Unmarshal(&struct{}{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestResult_SerializeSkipExistingIssues_Issues(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	b, err := r.SerializeSkipExistingIssues()
	if b != nil || err != nil {
		t.Fatal("expected nil, nil")
	}
}

func TestResult_Serialize_Nil(t *testing.T) {
	var r *Result
	_, err := r.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestResult_Serialize_WithError(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	_, err := r.Serialize()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestResult_Serialize_Normal(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	b, err := r.Serialize()
	if err != nil || len(b) == 0 {
		t.Fatal("unexpected")
	}
}

func TestResult_SerializeMust(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	b := r.SerializeMust()
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestResult_UnmarshalSkipExistingIssues(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	err := r.UnmarshalSkipExistingIssues(&struct{}{})
	if err != nil {
		t.Fatal("expected nil for issues")
	}
	r2 := &Result{Bytes: []byte(`"hello"`)}
	var s string
	err2 := r2.UnmarshalSkipExistingIssues(&s)
	if err2 != nil || s != "hello" {
		t.Fatal("unexpected")
	}
}

func TestResult_UnmarshalResult(t *testing.T) {
	r := &Result{Bytes: []byte(`{"Bytes":"aGVsbG8="}`)}
	_, _ = r.UnmarshalResult()
}

func TestResult_JsonModel_Nil(t *testing.T) {
	var r *Result
	jm := r.JsonModel()
	if jm.Error == nil {
		t.Fatal("expected error")
	}
}

func TestResult_JsonModel_Normal(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	jm := r.JsonModel()
	if len(jm.Bytes) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestResult_JsonModelAny(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	_ = r.JsonModelAny()
}

func TestResult_Json(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	_ = r.Json()
}

func TestResult_JsonPtr(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	_ = r.JsonPtr()
}

func TestResult_ParseInjectUsingJson(t *testing.T) {
	r := &Result{}
	input := Serialize.Apply("hello")
	_, _ = r.ParseInjectUsingJson(input)
}

func TestResult_CloneError(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	if r.CloneError() == nil {
		t.Fatal("expected error")
	}
	r2 := &Result{}
	if r2.CloneError() != nil {
		t.Fatal("expected nil")
	}
}

func TestResult_Ptr_NonPtr_ToPtr_ToNonPtr(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	_ = r.Ptr()
	_ = r.ToPtr()
	_ = r.ToNonPtr()
	var rp *Result
	nr := rp.NonPtr()
	if nr.Error == nil {
		t.Fatal("expected error")
	}
	rp2 := &Result{Bytes: []byte(`"x"`)}
	nr2 := rp2.NonPtr()
	if len(nr2.Bytes) == 0 {
		t.Fatal("expected bytes")
	}
}

func TestResult_IsEqualPtr(t *testing.T) {
	r1 := &Result{Bytes: []byte(`"x"`), TypeName: "t"}
	r2 := &Result{Bytes: []byte(`"x"`), TypeName: "t"}
	if !r1.IsEqualPtr(r2) {
		t.Fatal("expected equal")
	}
	if !r1.IsEqualPtr(r1) {
		t.Fatal("expected equal same ptr")
	}
	var n *Result
	if !n.IsEqualPtr(nil) {
		t.Fatal("expected equal nil")
	}
	if n.IsEqualPtr(r1) {
		t.Fatal("expected not equal")
	}
	if r1.IsEqualPtr(nil) {
		t.Fatal("expected not equal")
	}
	r3 := &Result{Bytes: []byte(`"y"`), TypeName: "t"}
	if r1.IsEqualPtr(r3) {
		t.Fatal("expected not equal diff bytes")
	}
	r4 := &Result{Bytes: []byte(`"x"`), TypeName: "t2"}
	if r1.IsEqualPtr(r4) {
		t.Fatal("expected not equal diff type")
	}
	r5 := &Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "t"}
	if r1.IsEqualPtr(r5) {
		t.Fatal("expected not equal diff err")
	}
}

func TestResult_CombineErrorWithRefString(t *testing.T) {
	r := &Result{}
	if r.CombineErrorWithRefString("ref") != "" {
		t.Fatal("expected empty")
	}
	r2 := &Result{Error: errors.New("e")}
	if r2.CombineErrorWithRefString("ref") == "" {
		t.Fatal("expected non-empty")
	}
}

func TestResult_CombineErrorWithRefError(t *testing.T) {
	r := &Result{}
	if r.CombineErrorWithRefError("ref") != nil {
		t.Fatal("expected nil")
	}
	r2 := &Result{Error: errors.New("e")}
	if r2.CombineErrorWithRefError("ref") == nil {
		t.Fatal("expected error")
	}
}

func TestResult_IsEqual(t *testing.T) {
	r1 := Result{Bytes: []byte(`"x"`)}
	r2 := Result{Bytes: []byte(`"x"`)}
	if !r1.IsEqual(r2) {
		t.Fatal("expected equal")
	}
}

func TestResult_BytesError_Nil(t *testing.T) {
	var r *Result
	if r.BytesError() != nil {
		t.Fatal("expected nil")
	}
}

func TestResult_BytesError_Normal(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	be := r.BytesError()
	if be == nil {
		t.Fatal("expected non-nil")
	}
}

func TestResult_Dispose(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "t"}
	r.Dispose()
	if r.Bytes != nil || r.Error != nil || r.TypeName != "" {
		t.Fatal("expected disposed")
	}
	var rn *Result
	rn.Dispose() // should not panic
}

func TestResult_CloneIf(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`), TypeName: "t"}
	c := r.CloneIf(true, true)
	if len(c.Bytes) == 0 {
		t.Fatal("expected cloned")
	}
	c2 := r.CloneIf(false, false)
	if len(c2.Bytes) == 0 {
		t.Fatal("expected original")
	}
}

func TestResult_ClonePtr(t *testing.T) {
	var r *Result
	if r.ClonePtr(true) != nil {
		t.Fatal("expected nil")
	}
	r2 := &Result{Bytes: []byte(`"x"`), TypeName: "t"}
	c := r2.ClonePtr(true)
	if c == nil || len(c.Bytes) == 0 {
		t.Fatal("expected cloned")
	}
}

func TestResult_Clone(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`), TypeName: "t"}
	c := r.Clone(true)
	if len(c.Bytes) == 0 {
		t.Fatal("expected deep cloned")
	}
	c2 := r.Clone(false)
	if len(c2.Bytes) == 0 {
		t.Fatal("expected shallow cloned")
	}
	r3 := Result{TypeName: "t"}
	c3 := r3.Clone(true)
	if c3.TypeName != "t" {
		t.Fatal("expected type name preserved")
	}
}

func TestResult_AsJsonContractsBinder(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	_ = r.AsJsonContractsBinder()
}

func TestResult_AsJsoner(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	_ = r.AsJsoner()
}

func TestResult_JsonParseSelfInject(t *testing.T) {
	r := Result{}
	input := Serialize.Apply("hello")
	_ = r.JsonParseSelfInject(input)
}

func TestResult_AsJsonParseSelfInjector(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`)}
	_ = r.AsJsonParseSelfInjector()
}

func TestResult_InjectInto(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	r2 := Result{}
	_ = r.InjectInto(&r2)
}
