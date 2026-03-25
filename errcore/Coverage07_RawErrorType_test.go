package errcore

import (
	"errors"
	"testing"
)

func TestRawErrorType_String(t *testing.T) {
	if InvalidType.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_Combine(t *testing.T) {
	s := InvalidType.Combine("msg", "ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_CombineWithAnother(t *testing.T) {
	r := InvalidType.CombineWithAnother(NotFound, "msg", "ref")
	if r.String() == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_TypesAttach(t *testing.T) {
	s := InvalidType.TypesAttach("msg", "hello", 42)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_TypesAttachErr(t *testing.T) {
	err := InvalidType.TypesAttachErr("msg", "hello")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_SrcDestination(t *testing.T) {
	s := InvalidType.SrcDestination("msg", "src", "sv", "dst", "dv")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_SrcDestinationErr(t *testing.T) {
	err := InvalidType.SrcDestinationErr("msg", "src", "sv", "dst", "dv")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_Error(t *testing.T) {
	err := InvalidType.Error("msg", "ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_ErrorSkip(t *testing.T) {
	err := InvalidType.ErrorSkip(0, "msg", "ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_Fmt_EmptyFormat(t *testing.T) {
	err := InvalidType.Fmt("")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_Fmt_WithFormat(t *testing.T) {
	err := InvalidType.Fmt("val=%d", 42)
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_FmtIf_False(t *testing.T) {
	err := InvalidType.FmtIf(false, "val=%d", 42)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrorType_FmtIf_True(t *testing.T) {
	err := InvalidType.FmtIf(true, "val=%d", 42)
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_MergeError_Nil(t *testing.T) {
	if InvalidType.MergeError(nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrorType_MergeError_WithErr(t *testing.T) {
	err := InvalidType.MergeError(errors.New("e"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	if InvalidType.MergeErrorWithMessage(nil, "msg") != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrorType_MergeErrorWithMessage_WithErr(t *testing.T) {
	err := InvalidType.MergeErrorWithMessage(errors.New("e"), "msg")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	if InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref") != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrorType_MergeErrorWithMessageRef_WithErr(t *testing.T) {
	err := InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	if InvalidType.MergeErrorWithRef(nil, "ref") != nil {
		t.Fatal("expected nil")
	}
}

func TestRawErrorType_MergeErrorWithRef_WithErr(t *testing.T) {
	err := InvalidType.MergeErrorWithRef(errors.New("e"), "ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_MsgCsvRef_Empty(t *testing.T) {
	s := InvalidType.MsgCsvRef("msg")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_MsgCsvRef_EmptyMsg(t *testing.T) {
	s := InvalidType.MsgCsvRef("", "ref")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_MsgCsvRef_WithItems(t *testing.T) {
	s := InvalidType.MsgCsvRef("msg", "r1", "r2")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_MsgCsvRefError(t *testing.T) {
	err := InvalidType.MsgCsvRefError("msg", "r1")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_ErrorRefOnly(t *testing.T) {
	err := InvalidType.ErrorRefOnly("ref")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_Expecting(t *testing.T) {
	err := InvalidType.Expecting("exp", "act")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_NoRef_EmptyMsg(t *testing.T) {
	s := InvalidType.NoRef("")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_NoRef_WithMsg(t *testing.T) {
	s := InvalidType.NoRef("msg")
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestRawErrorType_ErrorNoRefs(t *testing.T) {
	err := InvalidType.ErrorNoRefs("msg")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_ErrorNoRefs_Empty(t *testing.T) {
	err := InvalidType.ErrorNoRefs("")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_ErrorNoRefsSkip(t *testing.T) {
	err := InvalidType.ErrorNoRefsSkip(0, "msg")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestRawErrorType_HandleUsingPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	InvalidType.HandleUsingPanic("msg", "ref")
}

func TestGetSet_True(t *testing.T) {
	r := GetSet(true, InvalidType, NotFound)
	if r != InvalidType {
		t.Fatal("expected InvalidType")
	}
}

func TestGetSet_False(t *testing.T) {
	r := GetSet(false, InvalidType, NotFound)
	if r != NotFound {
		t.Fatal("expected NotFound")
	}
}

func TestGetSetVariant_True(t *testing.T) {
	r := GetSetVariant(true, "a", "b")
	if r != "a" {
		t.Fatal("expected a")
	}
}

func TestGetSetVariant_False(t *testing.T) {
	r := GetSetVariant(false, "a", "b")
	if r != "b" {
		t.Fatal("expected b")
	}
}

func TestMeaningfulError_Nil(t *testing.T) {
	if MeaningfulError(InvalidType, "fn", nil) != nil {
		t.Fatal("expected nil")
	}
}

func TestMeaningfulError_WithErr(t *testing.T) {
	err := MeaningfulError(InvalidType, "fn", errors.New("e"))
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestMeaningfulErrorHandle_Nil(t *testing.T) {
	MeaningfulErrorHandle(InvalidType, "fn", nil)
}

func TestMeaningfulErrorHandle_WithErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	MeaningfulErrorHandle(InvalidType, "fn", errors.New("e"))
}

func TestMeaningfulErrorWithData_Nil(t *testing.T) {
	if MeaningfulErrorWithData(InvalidType, "fn", nil, "data") != nil {
		t.Fatal("expected nil")
	}
}

func TestMeaningfulErrorWithData_WithErr(t *testing.T) {
	err := MeaningfulErrorWithData(InvalidType, "fn", errors.New("e"), "data")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestMeaningfulMessageError_Nil(t *testing.T) {
	if MeaningfulMessageError(InvalidType, "fn", nil, "msg") != nil {
		t.Fatal("expected nil")
	}
}

func TestMeaningfulMessageError_WithErr(t *testing.T) {
	err := MeaningfulMessageError(InvalidType, "fn", errors.New("e"), "msg")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestPathMeaningfulMessage_Empty(t *testing.T) {
	if PathMeaningfulMessage(InvalidType, "fn", "/path") != nil {
		t.Fatal("expected nil")
	}
}

func TestPathMeaningfulMessage_WithMsgs(t *testing.T) {
	err := PathMeaningfulMessage(InvalidType, "fn", "/path", "a", "b")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}

func TestPathMeaningfulError_Nil(t *testing.T) {
	if PathMeaningfulError(InvalidType, nil, "/path") != nil {
		t.Fatal("expected nil")
	}
}

func TestPathMeaningfulError_WithErr(t *testing.T) {
	err := PathMeaningfulError(InvalidType, errors.New("e"), "/path")
	if err == nil {
		t.Fatal("expected non-nil")
	}
}
