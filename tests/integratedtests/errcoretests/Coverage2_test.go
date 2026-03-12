package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
)

// ==========================================
// RawErrorType methods
// ==========================================

func Test_RawErrorType_String(t *testing.T) {
	r := errcore.InvalidRequestType.String()
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_Combine(t *testing.T) {
	r := errcore.InvalidRequestType.Combine("other msg", "ref-value")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_CombineWithAnother(t *testing.T) {
	r := errcore.InvalidRequestType.CombineWithAnother(
		errcore.InvalidEmptyValueType, "msg", "ref",
	)
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_TypesAttach(t *testing.T) {
	r := errcore.TypeMismatchType.TypesAttach("msg", "string")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_TypesAttachErr(t *testing.T) {
	err := errcore.TypeMismatchType.TypesAttachErr("msg", "string")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_SrcDestination(t *testing.T) {
	r := errcore.InvalidRequestType.SrcDestination("msg", "src", "srcVal", "dst", "dstVal")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_SrcDestinationErr(t *testing.T) {
	err := errcore.InvalidRequestType.SrcDestinationErr("msg", "src", "srcVal", "dst", "dstVal")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_Error(t *testing.T) {
	err := errcore.InvalidRequestType.Error("msg", "ref")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_ErrorSkip(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorSkip(0, "msg", "ref")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_Fmt(t *testing.T) {
	err := errcore.InvalidRequestType.Fmt("value %d", 42)
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_Fmt_Empty(t *testing.T) {
	err := errcore.InvalidRequestType.Fmt("")
	if err == nil {
		t.Error("should return error even for empty format")
	}
}

func Test_RawErrorType_FmtIf_True(t *testing.T) {
	err := errcore.InvalidRequestType.FmtIf(true, "value %d", 42)
	if err == nil {
		t.Error("should return error when true")
	}
}

func Test_RawErrorType_FmtIf_False(t *testing.T) {
	err := errcore.InvalidRequestType.FmtIf(false, "value %d", 42)
	if err != nil {
		t.Error("should return nil when false")
	}
}

func Test_RawErrorType_MergeError_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeError(nil)
	if err != nil {
		t.Error("nil error should return nil")
	}
}

func Test_RawErrorType_MergeError_NonNil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeError(errors.New("inner"))
	if err == nil {
		t.Error("should return merged error")
	}
}

func Test_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessage(nil, "msg")
	if err != nil {
		t.Error("nil error should return nil")
	}
}

func Test_RawErrorType_MergeErrorWithMessage_NonNil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessage(errors.New("inner"), "msg")
	if err == nil {
		t.Error("should return merged error")
	}
}

func Test_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(nil, "msg", "ref")
	if err != nil {
		t.Error("nil error should return nil")
	}
}

func Test_RawErrorType_MergeErrorWithMessageRef_NonNil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(errors.New("inner"), "msg", "ref")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithRef(nil, "ref")
	if err != nil {
		t.Error("nil error should return nil")
	}
}

func Test_RawErrorType_MergeErrorWithRef_NonNil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithRef(errors.New("inner"), "ref")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_MsgCsvRef_WithItems(t *testing.T) {
	r := errcore.InvalidRequestType.MsgCsvRef("msg", "a", "b")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_MsgCsvRef_NoItems(t *testing.T) {
	r := errcore.InvalidRequestType.MsgCsvRef("msg")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_MsgCsvRef_EmptyMsg(t *testing.T) {
	r := errcore.InvalidRequestType.MsgCsvRef("", "a")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_MsgCsvRefError(t *testing.T) {
	err := errcore.InvalidRequestType.MsgCsvRefError("msg", "a")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_ErrorRefOnly(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorRefOnly("ref")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_Expecting(t *testing.T) {
	err := errcore.InvalidRequestType.Expecting("expected", "actual")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_NoRef_WithMsg(t *testing.T) {
	r := errcore.InvalidRequestType.NoRef("other msg")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_NoRef_EmptyMsg(t *testing.T) {
	r := errcore.InvalidRequestType.NoRef("")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_RawErrorType_ErrorNoRefs(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorNoRefs("msg")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_ErrorNoRefs_Empty(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorNoRefs("")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorNoRefsSkip(0, "msg")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_RawErrorType_ErrorNoRefsSkip_Empty(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorNoRefsSkip(0, "")
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// GetSet / GetSetVariant
// ==========================================

func Test_GetSet_True(t *testing.T) {
	r := errcore.GetSet(true, errcore.InvalidRequestType, errcore.InvalidEmptyValueType)
	if r != errcore.InvalidRequestType {
		t.Error("should return true value")
	}
}

func Test_GetSet_False(t *testing.T) {
	r := errcore.GetSet(false, errcore.InvalidRequestType, errcore.InvalidEmptyValueType)
	if r != errcore.InvalidEmptyValueType {
		t.Error("should return false value")
	}
}

func Test_GetSetVariant_True(t *testing.T) {
	r := errcore.GetSetVariant(true, "trueVal", "falseVal")
	if r != "trueVal" {
		t.Error("should return true value")
	}
}

func Test_GetSetVariant_False(t *testing.T) {
	r := errcore.GetSetVariant(false, "trueVal", "falseVal")
	if r != "falseVal" {
		t.Error("should return false value")
	}
}

// ==========================================
// HandleErr / SimpleHandleErr
// ==========================================

func Test_HandleErr_NilError(t *testing.T) {
	// Should not panic
	errcore.HandleErr(nil)
}

func Test_SimpleHandleErr_NilError(t *testing.T) {
	// Should not panic
	errcore.SimpleHandleErr(nil)
}

// ==========================================
// MeaningFulError
// ==========================================

func Test_MeaningFulError_EmptyMsg(t *testing.T) {
	err := errcore.MeaningFulError(errcore.InvalidRequestType, "funcName", "")
	if err != nil {
		t.Error("empty message should return nil")
	}
}

func Test_MeaningFulError_WithMsg(t *testing.T) {
	err := errcore.MeaningFulError(errcore.InvalidRequestType, "funcName", "some error")
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// PathMeaningfulMessage
// ==========================================

func Test_PathMeaningfulMessage_NoMessages(t *testing.T) {
	err := errcore.PathMeaningfulMessage(errcore.InvalidRequestType, "fn", "loc")
	if err != nil {
		t.Error("no messages should return nil")
	}
}

func Test_PathMeaningfulMessage_WithMessages(t *testing.T) {
	err := errcore.PathMeaningfulMessage(errcore.InvalidRequestType, "fn", "loc", "msg1", "msg2")
	if err == nil {
		t.Error("should return error")
	}
}

// ==========================================
// MergeErrorsToString / MergeErrorsToStringDefault
// ==========================================

func Test_MergeErrorsToString_Nil(t *testing.T) {
	r := errcore.MergeErrorsToString(",")
	if r != "" {
		t.Error("no errors should return empty")
	}
}

func Test_MergeErrorsToString_WithErrors(t *testing.T) {
	r := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_MergeErrorsToStringDefault(t *testing.T) {
	r := errcore.MergeErrorsToStringDefault(errors.New("a"))
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// CountStateChangeTracker
// ==========================================

func Test_CountStateChangeTracker(t *testing.T) {
	tracker := errcore.CountStateChangeTracker{}
	if tracker.HasChanged() {
		t.Error("initial should not be changed")
	}
	tracker.IncrementOnChange()
	if !tracker.HasChanged() {
		t.Error("should be changed after increment")
	}
}

// ==========================================
// MessageNameValues
// ==========================================

func Test_MessageNameValues(t *testing.T) {
	r := errcore.MessageNameValues("msg", "name1", "val1", "name2", "val2")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// VarNameValues
// ==========================================

func Test_VarNameValues(t *testing.T) {
	r := errcore.VarNameValues("name1", "val1", "name2", "val2")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// SourceDestination / SourceDestinationErr / SourceDestinationNoType
// ==========================================

func Test_SourceDestination(t *testing.T) {
	r := errcore.SourceDestination("src", "srcVal", "dst", "dstVal")
	if r == "" {
		t.Error("should return non-empty")
	}
}

func Test_SourceDestinationErr(t *testing.T) {
	err := errcore.SourceDestinationErr("src", "srcVal", "dst", "dstVal")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_SourceDestinationNoType(t *testing.T) {
	r := errcore.SourceDestinationNoType("src", "srcVal", "dst", "dstVal")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// StringLinesToQuoteLines / StringLinesToQuoteLinesWithTabs
// ==========================================

func Test_StringLinesToQuoteLines(t *testing.T) {
	r := errcore.StringLinesToQuoteLines([]string{"a", "b"})
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

func Test_StringLinesToQuoteLinesToSingle(t *testing.T) {
	r := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// LineDiff
// ==========================================

func Test_LineDiff(t *testing.T) {
	r := errcore.LineDiff("actual", "expected")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// MustBeEmpty
// ==========================================

func Test_MustBeEmpty_Nil(t *testing.T) {
	err := errcore.MustBeEmpty(nil)
	if err != nil {
		t.Error("nil should return nil")
	}
}

func Test_MustBeEmpty_EmptySlice(t *testing.T) {
	err := errcore.MustBeEmpty([]string{})
	if err != nil {
		t.Error("empty should return nil")
	}
}

func Test_MustBeEmpty_NonEmpty(t *testing.T) {
	err := errcore.MustBeEmpty([]string{"a"})
	if err == nil {
		t.Error("non-empty should return error")
	}
}

// ==========================================
// GherkinsString
// ==========================================

func Test_GherkinsString(t *testing.T) {
	r := errcore.GherkinsString("given", "when", "then")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// ExpectingFuture
// ==========================================

func Test_ExpectingFuture(t *testing.T) {
	r := errcore.ExpectingFuture("header", "expected", "actual")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// ExpectingRecord
// ==========================================

func Test_ExpectingRecord(t *testing.T) {
	r := errcore.ExpectingRecord("header", 0, "expected", "actual")
	if r == "" {
		t.Error("should return non-empty")
	}
}

// ==========================================
// RawErrCollection extended
// ==========================================

func Test_RawErrCollection_AddNilAndNonNil(t *testing.T) {
	c := errcore.RawErrCollection{}
	c.Add(nil)
	if c.HasError() {
		t.Error("nil should not count as error")
	}
	c.Add(errors.New("err"))
	if !c.HasError() {
		t.Error("should have error")
	}
	if c.Length() < 1 {
		t.Error("length should be at least 1")
	}
}
