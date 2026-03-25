package corejson

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata"
)

// Covers: newResultCreator methods

func Test_I17_NewResult_UnmarshalUsingBytes(t *testing.T) {
	b := Serialize.ToBytesMust(Result{Bytes: []byte(`"x"`)})
	r := NewResult.UnmarshalUsingBytes(b)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_DeserializeUsingBytes_Invalid(t *testing.T) {
	r := NewResult.DeserializeUsingBytes([]byte("invalid"))
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_DeserializeUsingResult_WithIssues(t *testing.T) {
	r := NewResult.DeserializeUsingResult(&Result{Error: errors.New("bad")})
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_DeserializeUsingResult_Valid(t *testing.T) {
	inner := NewPtr(Result{Bytes: []byte(`"x"`)})
	r := NewResult.DeserializeUsingResult(inner)
	_ = r
}

func Test_I17_NewResult_UsingBytes(t *testing.T) {
	r := NewResult.UsingBytes([]byte(`"x"`))
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_UsingBytesType(t *testing.T) {
	r := NewResult.UsingBytesType([]byte(`"x"`), "TestType")
	if r.TypeName != "TestType" {
		t.Fatal("unexpected type name")
	}
}

func Test_I17_NewResult_UsingBytesTypePtr(t *testing.T) {
	r := NewResult.UsingBytesTypePtr([]byte(`"x"`), "T")
	if r == nil || r.TypeName != "T" {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResult_UsingTypeBytesPtr(t *testing.T) {
	r := NewResult.UsingTypeBytesPtr("T", []byte(`"x"`))
	if r == nil || r.TypeName != "T" {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResult_UsingBytesPtr_Nil(t *testing.T) {
	r := NewResult.UsingBytesPtr(nil)
	if r == nil || len(r.Bytes) != 0 {
		t.Fatal("expected empty bytes")
	}
}

func Test_I17_NewResult_UsingBytesPtr_Valid(t *testing.T) {
	r := NewResult.UsingBytesPtr([]byte(`"x"`))
	if r == nil || len(r.Bytes) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I17_NewResult_UsingBytesPtrErrPtr_Nil(t *testing.T) {
	r := NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T")
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_UsingBytesPtrErrPtr_Valid(t *testing.T) {
	r := NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "T")
	if r.Error != nil {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_UsingBytesErrPtr_Empty(t *testing.T) {
	r := NewResult.UsingBytesErrPtr(nil, errors.New("e"), "T")
	if r.Error == nil || len(r.Bytes) != 0 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResult_UsingBytesErrPtr_Valid(t *testing.T) {
	r := NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "T")
	if len(r.Bytes) == 0 {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResult_PtrUsingStringPtr_Nil(t *testing.T) {
	r := NewResult.PtrUsingStringPtr(nil, "T")
	if r.Error == nil {
		t.Fatal("expected error for nil string ptr")
	}
}

func Test_I17_NewResult_PtrUsingStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := NewResult.PtrUsingStringPtr(&s, "T")
	if r.Error != nil {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_UsingErrorStringPtr_NilPtr(t *testing.T) {
	r := NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_UsingErrorStringPtr_NilErr(t *testing.T) {
	s := `"hello"`
	r := NewResult.UsingErrorStringPtr(nil, &s, "T")
	if r.Error != nil {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_UsingErrorStringPtr_BothNil(t *testing.T) {
	r := NewResult.UsingErrorStringPtr(nil, nil, "T")
	if r.Error == nil {
		t.Fatal("expected error for nil ptr")
	}
}

func Test_I17_NewResult_Ptr(t *testing.T) {
	r := NewResult.Ptr([]byte(`"x"`), nil, "T")
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_UsingJsonBytesTypeError(t *testing.T) {
	r := NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "T")
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_UsingJsonBytesError(t *testing.T) {
	r := NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_UsingTypePlusString(t *testing.T) {
	r := NewResult.UsingTypePlusString("T", `"x"`)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_UsingTypePlusStringPtr_Nil(t *testing.T) {
	r := NewResult.UsingTypePlusStringPtr("T", nil)
	if len(r.Bytes) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResult_UsingTypePlusStringPtr_Empty(t *testing.T) {
	s := ""
	r := NewResult.UsingTypePlusStringPtr("T", &s)
	if len(r.Bytes) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResult_UsingTypePlusStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := NewResult.UsingTypePlusStringPtr("T", &s)
	if len(r.Bytes) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I17_NewResult_UsingStringWithType(t *testing.T) {
	r := NewResult.UsingStringWithType(`"x"`, "T")
	if r.TypeName != "T" {
		t.Fatal("unexpected type")
	}
}

func Test_I17_NewResult_UsingString(t *testing.T) {
	r := NewResult.UsingString(`"x"`)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_UsingStringPtr_Nil(t *testing.T) {
	r := NewResult.UsingStringPtr(nil)
	if len(r.Bytes) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResult_UsingStringPtr_Empty(t *testing.T) {
	s := ""
	r := NewResult.UsingStringPtr(&s)
	if len(r.Bytes) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResult_UsingStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := NewResult.UsingStringPtr(&s)
	if len(r.Bytes) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I17_NewResult_CreatePtr(t *testing.T) {
	r := NewResult.CreatePtr([]byte(`"x"`), nil, "T")
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_NonPtr(t *testing.T) {
	r := NewResult.NonPtr([]byte(`"x"`), nil, "T")
	if r.TypeName != "T" {
		t.Fatal("unexpected type")
	}
}

func Test_I17_NewResult_Create(t *testing.T) {
	r := NewResult.Create([]byte(`"x"`), nil, "T")
	if r.TypeName != "T" {
		t.Fatal("unexpected type")
	}
}

func Test_I17_NewResult_PtrUsingBytesPtr_WithErr(t *testing.T) {
	r := NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "T")
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_PtrUsingBytesPtr_NilBytes(t *testing.T) {
	r := NewResult.PtrUsingBytesPtr(nil, nil, "T")
	if len(r.Bytes) != 0 {
		t.Fatal("expected empty bytes")
	}
}

func Test_I17_NewResult_PtrUsingBytesPtr_Valid(t *testing.T) {
	r := NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "T")
	if len(r.Bytes) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I17_NewResult_CastingAny(t *testing.T) {
	r := NewResult.CastingAny("hello")
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_Any(t *testing.T) {
	r := NewResult.Any("test")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_Any_Error(t *testing.T) {
	r := NewResult.Any(func() {})
	if !r.HasError() {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_AnyPtr(t *testing.T) {
	r := NewResult.AnyPtr("test")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_AnyPtr_Error(t *testing.T) {
	r := NewResult.AnyPtr(func() {})
	if !r.HasError() {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_UsingBytesError_Nil(t *testing.T) {
	r := NewResult.UsingBytesError(nil)
	if r.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResult_UsingBytesError_Valid(t *testing.T) {
	be := &coredata.BytesError{Bytes: []byte(`"x"`), Error: nil}
	r := NewResult.UsingBytesError(be)
	if len(r.Bytes) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_I17_NewResult_Error(t *testing.T) {
	r := NewResult.Error(errors.New("e"))
	if r.Error == nil {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_ErrorPtr(t *testing.T) {
	r := NewResult.ErrorPtr(errors.New("e"))
	if r == nil || r.Error == nil {
		t.Fatal("expected error ptr")
	}
}

func Test_I17_NewResult_Empty(t *testing.T) {
	r := NewResult.Empty()
	if r.HasAnyItem() {
		t.Fatal("expected empty")
	}
}

func Test_I17_NewResult_EmptyPtr(t *testing.T) {
	r := NewResult.EmptyPtr()
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_TypeName(t *testing.T) {
	r := NewResult.TypeName("T")
	if r.TypeName != "T" {
		t.Fatal("unexpected type name")
	}
}

func Test_I17_NewResult_TypeNameBytes(t *testing.T) {
	r := NewResult.TypeNameBytes("T")
	if r.TypeName != "T" {
		t.Fatal("unexpected type name")
	}
}

func Test_I17_NewResult_Many(t *testing.T) {
	r := NewResult.Many("a", "b")
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_Serialize(t *testing.T) {
	r := NewResult.Serialize("test")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_Serialize_Error(t *testing.T) {
	r := NewResult.Serialize(func() {})
	if !r.HasError() {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_Marshal(t *testing.T) {
	r := NewResult.Marshal("test")
	if r.HasError() {
		t.Fatal("unexpected error")
	}
}

func Test_I17_NewResult_Marshal_Error(t *testing.T) {
	r := NewResult.Marshal(func() {})
	if !r.HasError() {
		t.Fatal("expected error")
	}
}

func Test_I17_NewResult_UsingSerializer_Nil(t *testing.T) {
	r := NewResult.UsingSerializer(nil)
	if r != nil {
		t.Fatal("expected nil for nil serializer")
	}
}

func Test_I17_NewResult_UsingSerializer_Valid(t *testing.T) {
	s := testSerializer{data: []byte(`"x"`)}
	r := NewResult.UsingSerializer(s)
	if r == nil || r.HasError() {
		t.Fatal("unexpected")
	}
}

func Test_I17_NewResult_UsingSerializerFunc_Nil(t *testing.T) {
	r := NewResult.UsingSerializerFunc(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_I17_NewResult_UsingSerializerFunc_Valid(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"x"`), nil }
	r := NewResult.UsingSerializerFunc(fn)
	if r == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I17_NewResult_UsingJsoner_Nil(t *testing.T) {
	r := NewResult.UsingJsoner(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func Test_I17_NewResult_AnyToCastingResult(t *testing.T) {
	r := NewResult.AnyToCastingResult("hello")
	if r == nil {
		t.Fatal("expected non-nil")
	}
}
