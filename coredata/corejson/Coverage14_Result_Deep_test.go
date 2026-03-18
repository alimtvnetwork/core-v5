package corejson

import (
	"errors"
	"testing"
)

// ── Result — nil/empty edge cases ──

func TestResult_Map_Nil(t *testing.T) {
	var r *Result
	m := r.Map()
	if len(m) != 0 { t.Fatal("expected empty") }
}

func TestResult_Map_WithAll(t *testing.T) {
	r := &Result{
		Bytes:    []byte(`"hello"`),
		Error:    errors.New("e"),
		TypeName: "string",
	}
	m := r.Map()
	if len(m) != 3 { t.Fatal("expected 3") }
}

func TestResult_Map_BytesOnly(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	m := r.Map()
	if _, ok := m["Bytes"]; !ok { t.Fatal("expected Bytes") }
}

func TestResult_DeserializedFieldsToMap_Nil(t *testing.T) {
	var r *Result
	m, err := r.DeserializedFieldsToMap()
	if err != nil { t.Fatal(err) }
	if len(m) != 0 { t.Fatal("expected empty") }
}

func TestResult_DeserializedFieldsToMap_Empty(t *testing.T) {
	r := &Result{}
	m, err := r.DeserializedFieldsToMap()
	if err != nil { t.Fatal(err) }
	if len(m) != 0 { t.Fatal("expected empty") }
}

func TestResult_SafeDeserializedFieldsToMap(t *testing.T) {
	r := &Result{}
	m := r.SafeDeserializedFieldsToMap()
	if len(m) != 0 { t.Fatal("expected empty") }
}

func TestResult_SafeFieldsNames(t *testing.T) {
	r := &Result{}
	names := r.SafeFieldsNames()
	if len(names) != 0 { t.Fatal("expected empty") }
}

func TestResult_BytesTypeName_Nil(t *testing.T) {
	var r *Result
	if r.BytesTypeName() != "" { t.Fatal("expected empty") }
}

func TestResult_BytesTypeName_Valid(t *testing.T) {
	r := &Result{TypeName: "test"}
	if r.BytesTypeName() != "test" { t.Fatal("unexpected") }
}

func TestResult_SafeBytesTypeName_Empty(t *testing.T) {
	r := &Result{}
	if r.SafeBytesTypeName() != "" { t.Fatal("expected empty") }
}

func TestResult_SafeBytesTypeName_Valid(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`), TypeName: "t"}
	if r.SafeBytesTypeName() != "t" { t.Fatal("expected t") }
}

func TestResult_SafeString(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	if r.SafeString() == "" { t.Fatal("expected non-empty") }
}

func TestResult_JsonStringPtr_Nil(t *testing.T) {
	var r *Result
	s := r.JsonStringPtr()
	if s == nil || *s != "" { t.Fatal("expected empty string ptr") }
}

func TestResult_JsonStringPtr_NoBytes(t *testing.T) {
	r := &Result{}
	s := r.JsonStringPtr()
	if *s != "" { t.Fatal("expected empty") }
}

func TestResult_JsonStringPtr_Cached(t *testing.T) {
	r := &Result{Bytes: []byte(`"x"`)}
	s1 := r.JsonStringPtr()
	s2 := r.JsonStringPtr()
	if s1 != s2 { t.Fatal("expected same pointer") }
}

func TestResult_PrettyJsonBuffer_Empty(t *testing.T) {
	r := &Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() != 0 { t.Fatal("unexpected") }
}

func TestResult_PrettyJsonBuffer_Valid(t *testing.T) {
	r := NewResult.AnyPtr(map[string]int{"a": 1})
	buf, err := r.PrettyJsonBuffer("", "  ")
	if err != nil || buf.Len() == 0 { t.Fatal("unexpected") }
}

func TestResult_PrettyJsonString_Nil(t *testing.T) {
	var r *Result
	if r.PrettyJsonString() != "" { t.Fatal("expected empty") }
}

func TestResult_PrettyJsonString_Valid(t *testing.T) {
	r := NewResult.AnyPtr(map[string]int{"a": 1})
	s := r.PrettyJsonString()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	var r *Result
	s := r.PrettyJsonStringOrErrString()
	if s == "" { t.Fatal("expected msg") }
}

func TestResult_PrettyJsonStringOrErrString_Err(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	s := r.PrettyJsonStringOrErrString()
	if s == "" { t.Fatal("expected err string") }
}

func TestResult_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	r := NewResult.AnyPtr(map[string]int{"a": 1})
	s := r.PrettyJsonStringOrErrString()
	if s == "" { t.Fatal("expected string") }
}

func TestResult_Length_Nil(t *testing.T) {
	var r *Result
	if r.Length() != 0 { t.Fatal("expected 0") }
}

func TestResult_Length_Valid(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if r.Length() == 0 { t.Fatal("expected > 0") }
}

func TestResult_HasError_False(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if r.HasError() { t.Fatal("unexpected error") }
}

func TestResult_ErrorString_NoError(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if r.ErrorString() != "" { t.Fatal("expected empty") }
}

func TestResult_ErrorString_WithError(t *testing.T) {
	r := &Result{Error: errors.New("fail")}
	if r.ErrorString() == "" { t.Fatal("expected msg") }
}

func TestResult_IsErrorEqual(t *testing.T) {
	e := errors.New("test")
	r := &Result{Error: e}
	if !r.IsErrorEqual(errors.New("test")) { t.Fatal("expected equal") }
	if r.IsErrorEqual(nil) { t.Fatal("expected not equal") }
	r2 := &Result{}
	if !r2.IsErrorEqual(nil) { t.Fatal("expected equal") }
	if r2.IsErrorEqual(e) { t.Fatal("expected not equal") }
}

func TestResult_String_Nil(t *testing.T) {
	r := Result{}
	_ = r.String()
}

func TestResult_String_WithError(t *testing.T) {
	r := Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "t"}
	s := r.String()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_String_NoError(t *testing.T) {
	r := NewResult.Any("hello")
	s := r.String()
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_SafeNonIssueBytes(t *testing.T) {
	r := NewResult.AnyPtr("x")
	b := r.SafeNonIssueBytes()
	if len(b) == 0 { t.Fatal("expected bytes") }
	r2 := &Result{Error: errors.New("e")}
	b2 := r2.SafeNonIssueBytes()
	if len(b2) != 0 { t.Fatal("expected empty") }
}

func TestResult_Values(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if r.Values() == nil { t.Fatal("expected bytes") }
}

func TestResult_SafeValues(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if len(r.SafeValues()) == 0 { t.Fatal("expected bytes") }
	var rNil *Result
	if len(rNil.SafeValues()) != 0 { t.Fatal("expected empty") }
}

func TestResult_SafeValuesPtr(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if len(r.SafeValuesPtr()) == 0 { t.Fatal("expected bytes") }
	r2 := &Result{Error: errors.New("e")}
	if len(r2.SafeValuesPtr()) != 0 { t.Fatal("expected empty") }
}

func TestResult_Raw_Nil(t *testing.T) {
	var r *Result
	_, err := r.Raw()
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Raw_Valid(t *testing.T) {
	r := NewResult.AnyPtr("x")
	b, err := r.Raw()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestResult_RawString(t *testing.T) {
	r := NewResult.AnyPtr("x")
	s, err := r.RawString()
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestResult_RawErrString(t *testing.T) {
	r := NewResult.AnyPtr("x")
	b, msg := r.RawErrString()
	if len(b) == 0 { t.Fatal("expected bytes") }
	_ = msg
}

func TestResult_RawPrettyString(t *testing.T) {
	r := NewResult.AnyPtr(map[string]int{"a": 1})
	s, err := r.RawPrettyString()
	if err != nil || s == "" { t.Fatal("unexpected") }
}

func TestResult_MeaningfulErrorMessage_Nil(t *testing.T) {
	var r *Result
	msg := r.MeaningfulErrorMessage()
	if msg == "" { t.Fatal("expected msg") }
}

func TestResult_MeaningfulErrorMessage_NoError(t *testing.T) {
	r := NewResult.AnyPtr("x")
	msg := r.MeaningfulErrorMessage()
	if msg != "" { t.Fatal("expected empty") }
}

func TestResult_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &Result{Bytes: []byte{}}
	err := r.MeaningfulError()
	if err == nil { t.Fatal("expected error") }
}

func TestResult_IsEmptyError(t *testing.T) {
	var r *Result
	if !r.IsEmptyError() { t.Fatal("expected true") }
	r2 := &Result{}
	if !r2.IsEmptyError() { t.Fatal("expected true") }
	r3 := &Result{Error: errors.New("e")}
	if r3.IsEmptyError() { t.Fatal("expected false") }
}

func TestResult_HasSafeItems(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if !r.HasSafeItems() { t.Fatal("expected true") }
	r2 := &Result{Error: errors.New("e")}
	if r2.HasSafeItems() { t.Fatal("expected false") }
}

func TestResult_IsAnyNull(t *testing.T) {
	var r *Result
	if !r.IsAnyNull() { t.Fatal("expected true") }
	r2 := &Result{}
	if !r2.IsAnyNull() { t.Fatal("expected true") }
	r3 := &Result{Bytes: []byte(`"x"`)}
	if r3.IsAnyNull() { t.Fatal("expected false") }
}

func TestResult_HasBytes(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if !r.HasBytes() { t.Fatal("expected true") }
}

func TestResult_HasJsonBytes(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if !r.HasJsonBytes() { t.Fatal("expected true") }
}

func TestResult_IsEmptyJsonBytes_EmptyObject(t *testing.T) {
	r := &Result{Bytes: []byte("{}")}
	if !r.IsEmptyJsonBytes() { t.Fatal("expected true for {}") }
}

func TestResult_IsEmpty(t *testing.T) {
	var r *Result
	if !r.IsEmpty() { t.Fatal("expected true") }
}

func TestResult_HasAnyItem(t *testing.T) {
	r := NewResult.Any("x")
	if !r.HasAnyItem() { t.Fatal("expected true") }
}

func TestResult_IsEmptyJson(t *testing.T) {
	r := &Result{Bytes: []byte("{}")}
	if !r.IsEmptyJson() { t.Fatal("expected true") }
}

func TestResult_HasJson(t *testing.T) {
	r := NewResult.AnyPtr("x")
	if !r.HasJson() { t.Fatal("expected true") }
}

func TestResult_InjectInto(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	var s string
	err := r.InjectInto(&s)
	if err != nil { t.Fatal(err) }
}

func TestResult_SerializeSkipExistingIssues_Issues(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	b, err := r.SerializeSkipExistingIssues()
	if b != nil || err != nil { t.Fatal("expected nil,nil") }
}

func TestResult_SerializeSkipExistingIssues_Valid(t *testing.T) {
	r := NewResult.AnyPtr("x")
	b, err := r.SerializeSkipExistingIssues()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestResult_Serialize_Nil(t *testing.T) {
	var r *Result
	_, err := r.Serialize()
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Serialize_WithError(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	_, err := r.Serialize()
	if err == nil { t.Fatal("expected error") }
}

func TestResult_Serialize_Valid(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	b, err := r.Serialize()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
}

func TestResult_UnmarshalSkipExistingIssues_Issues(t *testing.T) {
	r := &Result{Error: errors.New("e")}
	err := r.UnmarshalSkipExistingIssues(nil)
	if err != nil { t.Fatal("expected nil") }
}

func TestResult_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	r := NewResult.AnyPtr("hello")
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestResult_UnmarshalSkipExistingIssues_BadPayload(t *testing.T) {
	r := &Result{Bytes: []byte("not json")}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	if err == nil { t.Fatal("expected error") }
}

func TestResult_UnmarshalResult(t *testing.T) {
	r := NewResult.AnyPtr(NewResult.Any("hi"))
	r2, err := r.UnmarshalResult()
	_ = r2
	_ = err
}

func TestResult_JsonModel_Nil(t *testing.T) {
	var r *Result
	model := r.JsonModel()
	if model.Error == nil { t.Fatal("expected error") }
}

func TestResult_JsonModel_Valid(t *testing.T) {
	r := NewResult.AnyPtr("x")
	model := r.JsonModel()
	if model.Error != nil { t.Fatal("unexpected error") }
}

func TestResult_JsonModelAny(t *testing.T) {
	r := NewResult.AnyPtr("x")
	a := r.JsonModelAny()
	if a == nil { t.Fatal("expected non-nil") }
}

func TestResult_CloneError(t *testing.T) {
	r := &Result{Error: errors.New("test")}
	e := r.CloneError()
	if e == nil { t.Fatal("expected error") }
	r2 := &Result{}
	e2 := r2.CloneError()
	if e2 != nil { t.Fatal("expected nil") }
}

func TestResult_Ptr(t *testing.T) {
	r := NewResult.Any("x")
	p := r.Ptr()
	if p == nil { t.Fatal("expected non-nil") }
}

func TestResult_NonPtr_Nil(t *testing.T) {
	var r *Result
	np := r.NonPtr()
	if np.Error == nil { t.Fatal("expected error") }
}

func TestResult_NonPtr_Valid(t *testing.T) {
	r := NewResult.AnyPtr("x")
	np := r.NonPtr()
	if np.Error != nil { t.Fatal("unexpected error") }
}

func TestResult_ToPtr(t *testing.T) {
	r := NewResult.Any("x")
	if r.ToPtr() == nil { t.Fatal("expected non-nil") }
}

func TestResult_ToNonPtr(t *testing.T) {
	r := NewResult.Any("x")
	np := r.ToNonPtr()
	_ = np
}

func TestResult_IsEqualPtr_BothNil(t *testing.T) {
	var a, b *Result
	if !a.IsEqualPtr(b) { t.Fatal("expected true") }
}

func TestResult_IsEqualPtr_OneNil(t *testing.T) {
	a := NewResult.AnyPtr("x")
	if a.IsEqualPtr(nil) { t.Fatal("expected false") }
}

func TestResult_IsEqualPtr_Same(t *testing.T) {
	a := NewResult.AnyPtr("x")
	if !a.IsEqualPtr(a) { t.Fatal("expected true") }
}

func TestResult_IsEqualPtr_DiffLen(t *testing.T) {
	a := NewResult.AnyPtr("x")
	b := NewResult.AnyPtr("xx")
	if a.IsEqualPtr(b) { t.Fatal("expected false") }
}

func TestResult_IsEqualPtr_DiffError(t *testing.T) {
	a := &Result{Bytes: []byte(`"x"`), Error: errors.New("e1")}
	b := &Result{Bytes: []byte(`"x"`), Error: errors.New("e2")}
	if a.IsEqualPtr(b) { t.Fatal("expected false") }
}

func TestResult_IsEqualPtr_Equal(t *testing.T) {
	a := NewResult.AnyPtr("hello")
	b := NewResult.AnyPtr("hello")
	if !a.IsEqualPtr(b) { t.Fatal("expected true") }
}

func TestResult_CombineErrorWithRefString_NoError(t *testing.T) {
	r := &Result{}
	s := r.CombineErrorWithRefString("ref1")
	if s != "" { t.Fatal("expected empty") }
}

func TestResult_CombineErrorWithRefString_WithError(t *testing.T) {
	r := &Result{Error: errors.New("fail")}
	s := r.CombineErrorWithRefString("ref1", "ref2")
	if s == "" { t.Fatal("expected non-empty") }
}

func TestResult_CombineErrorWithRefError_NoError(t *testing.T) {
	r := &Result{}
	e := r.CombineErrorWithRefError("ref")
	if e != nil { t.Fatal("expected nil") }
}

func TestResult_CombineErrorWithRefError_WithError(t *testing.T) {
	r := &Result{Error: errors.New("fail")}
	e := r.CombineErrorWithRefError("ref")
	if e == nil { t.Fatal("expected error") }
}

func TestResult_IsEqual(t *testing.T) {
	a := NewResult.Any("hello")
	b := NewResult.Any("hello")
	if !a.IsEqual(b) { t.Fatal("expected true") }
	c := NewResult.Any("world")
	if a.IsEqual(c) { t.Fatal("expected false") }
}

func TestResult_BytesError_Nil(t *testing.T) {
	var r *Result
	if r.BytesError() != nil { t.Fatal("expected nil") }
}

func TestResult_BytesError_Valid(t *testing.T) {
	r := NewResult.AnyPtr("x")
	be := r.BytesError()
	if be == nil { t.Fatal("expected non-nil") }
}

func TestResult_Dispose(t *testing.T) {
	r := NewResult.AnyPtr("x")
	r.Dispose()
	if r.Bytes != nil { t.Fatal("expected nil") }
}

func TestResult_Dispose_Nil(t *testing.T) {
	var r *Result
	r.Dispose() // should not panic
}

func TestResult_CloneIf_Clone(t *testing.T) {
	r := NewResult.Any("x")
	c := r.CloneIf(true, true)
	if c.Length() == 0 { t.Fatal("expected bytes") }
}

func TestResult_CloneIf_NoClone(t *testing.T) {
	r := NewResult.Any("x")
	c := r.CloneIf(false, false)
	if c.Length() == 0 { t.Fatal("expected bytes") }
}

func TestResult_Clone_Deep(t *testing.T) {
	r := NewResult.Any("hello")
	c := r.Clone(true)
	if c.Length() == 0 { t.Fatal("expected bytes") }
}

func TestResult_Clone_Shallow(t *testing.T) {
	r := NewResult.Any("hello")
	c := r.Clone(false)
	if c.Length() == 0 { t.Fatal("expected bytes") }
}

func TestResult_Clone_Empty(t *testing.T) {
	r := Result{}
	c := r.Clone(true)
	_ = c
}

func TestResult_ClonePtr_Nil(t *testing.T) {
	var r *Result
	if r.ClonePtr(true) != nil { t.Fatal("expected nil") }
}

func TestResult_ClonePtr_Valid(t *testing.T) {
	r := NewResult.AnyPtr("x")
	cp := r.ClonePtr(true)
	if cp == nil { t.Fatal("expected non-nil") }
}

func TestResult_AsJsonContractsBinder(t *testing.T) {
	r := NewResult.Any("x")
	if r.AsJsonContractsBinder() == nil { t.Fatal("expected non-nil") }
}

func TestResult_AsJsoner(t *testing.T) {
	r := NewResult.Any("x")
	if r.AsJsoner() == nil { t.Fatal("expected non-nil") }
}

func TestResult_JsonParseSelfInject(t *testing.T) {
	r := NewResult.Any("x")
	err := r.JsonParseSelfInject(NewResult.AnyPtr("y"))
	_ = err
}

func TestResult_AsJsonParseSelfInjector(t *testing.T) {
	r := NewResult.Any("x")
	if r.AsJsonParseSelfInjector() == nil { t.Fatal("expected non-nil") }
}

// ── BytesCloneIf / BytesDeepClone ──

func TestBytesDeepClone_Empty(t *testing.T) {
	b := BytesDeepClone(nil)
	if len(b) != 0 { t.Fatal("expected empty") }
}

func TestBytesDeepClone_Valid(t *testing.T) {
	orig := []byte{1, 2, 3}
	c := BytesDeepClone(orig)
	if len(c) != 3 { t.Fatal("expected 3") }
	orig[0] = 99
	if c[0] == 99 { t.Fatal("expected deep clone") }
}

func TestBytesCloneIf_True(t *testing.T) {
	b := BytesCloneIf(true, []byte{1, 2})
	if len(b) != 2 { t.Fatal("expected 2") }
}

func TestBytesCloneIf_False(t *testing.T) {
	b := BytesCloneIf(false, []byte{1, 2})
	if len(b) != 0 { t.Fatal("expected 0") }
}

func TestBytesCloneIf_Empty(t *testing.T) {
	b := BytesCloneIf(true, nil)
	if len(b) != 0 { t.Fatal("expected 0") }
}

// ── BytesToString / BytesToPrettyString ──

func TestBytesToString_Empty(t *testing.T) {
	if BytesToString(nil) != "" { t.Fatal("expected empty") }
}

func TestBytesToString_Valid(t *testing.T) {
	s := BytesToString([]byte("hello"))
	if s != "hello" { t.Fatal("unexpected") }
}

func TestBytesToPrettyString_Empty(t *testing.T) {
	if BytesToPrettyString(nil) != "" { t.Fatal("expected empty") }
}

func TestBytesToPrettyString_Valid(t *testing.T) {
	s := BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" { t.Fatal("expected non-empty") }
}

// ── New / NewPtr ──

func TestNew_ValidItem(t *testing.T) {
	r := New("hello")
	if r.HasError() { t.Fatal("unexpected error") }
}

func TestNew_Unmarshalable(t *testing.T) {
	r := New(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

func TestNewPtr_ValidItem(t *testing.T) {
	r := NewPtr("hello")
	if r.HasError() { t.Fatal("unexpected error") }
}

func TestNewPtr_Unmarshalable(t *testing.T) {
	r := NewPtr(make(chan int))
	if !r.HasError() { t.Fatal("expected error") }
}

// ── Empty creator ──

func TestEmptyCreator_All(t *testing.T) {
	_ = Empty.Result()
	_ = Empty.ResultPtr()
	_ = Empty.ResultWithErr("t", errors.New("e"))
	_ = Empty.ResultPtrWithErr("t", errors.New("e"))
	_ = Empty.BytesCollection()
	_ = Empty.BytesCollectionPtr()
	_ = Empty.ResultsCollection()
	_ = Empty.ResultsPtrCollection()
	_ = Empty.MapResults()
}

// ── StaticJsonError ──

func TestStaticJsonError(t *testing.T) {
	if StaticJsonError == nil { t.Fatal("expected non-nil") }
}
