package errcoretests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

func Test_RawErrorType_String(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_String", actual)
}

func Test_RawErrorType_Combine(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.Combine("other msg", "ref-value") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_Combine", actual)
}

func Test_RawErrorType_CombineWithAnother(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.CombineWithAnother(errcore.InvalidEmptyValueType, "msg", "ref") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_CombineWithAnother", actual)
}

func Test_RawErrorType_TypesAttach(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.TypeMismatchType.TypesAttach("msg", "string") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_TypesAttach", actual)
}

func Test_RawErrorType_TypesAttachErr(t *testing.T) {
	actual := args.Map{"hasErr": errcore.TypeMismatchType.TypesAttachErr("msg", "string") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_TypesAttachErr", actual)
}

func Test_RawErrorType_SrcDestination(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.SrcDestination("msg", "src", "srcVal", "dst", "dstVal") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_SrcDestination", actual)
}

func Test_RawErrorType_SrcDestinationErr(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.SrcDestinationErr("msg", "src", "srcVal", "dst", "dstVal") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_SrcDestinationErr", actual)
}

func Test_RawErrorType_Error(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Error("msg", "ref") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_Error", actual)
}

func Test_RawErrorType_ErrorSkip(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorSkip(0, "msg", "ref") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_ErrorSkip", actual)
}

func Test_RawErrorType_Fmt(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Fmt("value %d", 42) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_Fmt", actual)
}

func Test_RawErrorType_Fmt_Empty(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Fmt("") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_Fmt_Empty", actual)
}

func Test_RawErrorType_FmtIf_True(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.FmtIf(true, "value %d", 42) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_FmtIf_True", actual)
}

func Test_RawErrorType_FmtIf_False(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.FmtIf(false, "value %d", 42) != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "RawErrorType_FmtIf_False", actual)
}

func Test_RawErrorType_MergeError_Nil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeError(nil) != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeError_Nil", actual)
}

func Test_RawErrorType_MergeError_NonNil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeError(errors.New("inner")) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeError_NonNil", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithMessage(nil, "msg") != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeErrorWithMessage_Nil", actual)
}

func Test_RawErrorType_MergeErrorWithMessage_NonNil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithMessage(errors.New("inner"), "msg") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeErrorWithMessage_NonNil", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithMessageRef(nil, "msg", "ref") != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeErrorWithMessageRef_Nil", actual)
}

func Test_RawErrorType_MergeErrorWithMessageRef_NonNil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithMessageRef(errors.New("inner"), "msg", "ref") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeErrorWithMessageRef_NonNil", actual)
}

func Test_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithRef(nil, "ref") != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeErrorWithRef_Nil", actual)
}

func Test_RawErrorType_MergeErrorWithRef_NonNil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MergeErrorWithRef(errors.New("inner"), "ref") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MergeErrorWithRef_NonNil", actual)
}

func Test_RawErrorType_MsgCsvRef_WithItems(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.MsgCsvRef("msg", "a", "b") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MsgCsvRef_WithItems", actual)
}

func Test_RawErrorType_MsgCsvRef_NoItems(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.MsgCsvRef("msg") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MsgCsvRef_NoItems", actual)
}

func Test_RawErrorType_MsgCsvRef_EmptyMsg(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.MsgCsvRef("", "a") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MsgCsvRef_EmptyMsg", actual)
}

func Test_RawErrorType_MsgCsvRefError(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.MsgCsvRefError("msg", "a") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_MsgCsvRefError", actual)
}

func Test_RawErrorType_ErrorRefOnly(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorRefOnly("ref") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_ErrorRefOnly", actual)
}

func Test_RawErrorType_Expecting(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.Expecting("expected", "actual") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_Expecting", actual)
}

func Test_RawErrorType_NoRef_WithMsg(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.NoRef("other msg") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_NoRef_WithMsg", actual)
}

func Test_RawErrorType_NoRef_EmptyMsg(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.InvalidRequestType.NoRef("") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_NoRef_EmptyMsg", actual)
}

func Test_RawErrorType_ErrorNoRefs(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefs("msg") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_ErrorNoRefs", actual)
}

func Test_RawErrorType_ErrorNoRefs_Empty(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefs("") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_ErrorNoRefs_Empty", actual)
}

func Test_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefsSkip(0, "msg") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_ErrorNoRefsSkip", actual)
}

func Test_RawErrorType_ErrorNoRefsSkip_Empty(t *testing.T) {
	actual := args.Map{"hasErr": errcore.InvalidRequestType.ErrorNoRefsSkip(0, "") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType_ErrorNoRefsSkip_Empty", actual)
}

func Test_GetSet_True(t *testing.T) {
	actual := args.Map{"result": fmt.Sprintf("%v", errcore.GetSet(true, errcore.InvalidRequestType, errcore.InvalidEmptyValueType))}
	expected := args.Map{"result": fmt.Sprintf("%v", errcore.InvalidRequestType)}
	expected.ShouldBeEqual(t, 0, "GetSet_True", actual)
}

func Test_GetSet_False(t *testing.T) {
	actual := args.Map{"result": fmt.Sprintf("%v", errcore.GetSet(false, errcore.InvalidRequestType, errcore.InvalidEmptyValueType))}
	expected := args.Map{"result": fmt.Sprintf("%v", errcore.InvalidEmptyValueType)}
	expected.ShouldBeEqual(t, 0, "GetSet_False", actual)
}

func Test_GetSetVariant_True(t *testing.T) {
	actual := args.Map{"result": errcore.GetSetVariant(true, "trueVal", "falseVal")}
	expected := args.Map{"result": "trueVal"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant_True", actual)
}

func Test_GetSetVariant_False(t *testing.T) {
	actual := args.Map{"result": errcore.GetSetVariant(false, "trueVal", "falseVal")}
	expected := args.Map{"result": "falseVal"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant_False", actual)
}

func Test_HandleErr_NilError(t *testing.T) {
	errcore.HandleErr(nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr_NilError", actual)
}

func Test_SimpleHandleErr_NilError(t *testing.T) {
	errcore.SimpleHandleErr(nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr_NilError", actual)
}

func Test_MeaningFulError_EmptyMsg(t *testing.T) {
	actual := args.Map{"hasErr": errcore.MeaningFulError(errcore.InvalidRequestType, "funcName", "") != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MeaningFulError_EmptyMsg", actual)
}

func Test_MeaningFulError_WithMsg(t *testing.T) {
	actual := args.Map{"hasErr": errcore.MeaningFulError(errcore.InvalidRequestType, "funcName", "some error") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningFulError_WithMsg", actual)
}

func Test_PathMeaningfulMessage_NoMessages(t *testing.T) {
	actual := args.Map{"hasErr": errcore.PathMeaningfulMessage(errcore.InvalidRequestType, "fn", "loc") != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage_NoMessages", actual)
}

func Test_PathMeaningfulMessage_WithMessages(t *testing.T) {
	actual := args.Map{"hasErr": errcore.PathMeaningfulMessage(errcore.InvalidRequestType, "fn", "loc", "msg1", "msg2") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage_WithMessages", actual)
}

func Test_MergeErrorsToString_Nil(t *testing.T) {
	actual := args.Map{"isEmpty": errcore.MergeErrorsToString(",") == ""}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString_Nil", actual)
}

func Test_MergeErrorsToString_WithErrors(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b")) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString_WithErrors", actual)
}

func Test_MergeErrorsToStringDefault(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.MergeErrorsToStringDefault(errors.New("a")) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault", actual)
}

func Test_CountStateChangeTracker(t *testing.T) {
	tracker := errcore.CountStateChangeTracker{}
	initialChanged := tracker.HasChanged()
	tracker.IncrementOnChange()
	afterChanged := tracker.HasChanged()
	actual := args.Map{"initialChanged": initialChanged, "afterChanged": afterChanged}
	expected := args.Map{"initialChanged": false, "afterChanged": true}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker", actual)
}

func Test_MessageNameValues(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.MessageNameValues("msg", "name1", "val1", "name2", "val2") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues", actual)
}

func Test_VarNameValues(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.VarNameValues("name1", "val1", "name2", "val2") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues", actual)
}

func Test_SourceDestination(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.SourceDestination("src", "srcVal", "dst", "dstVal") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination", actual)
}

func Test_SourceDestinationErr(t *testing.T) {
	actual := args.Map{"hasErr": errcore.SourceDestinationErr("src", "srcVal", "dst", "dstVal") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr", actual)
}

func Test_SourceDestinationNoType(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.SourceDestinationNoType("src", "srcVal", "dst", "dstVal") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType", actual)
}

func Test_StringLinesToQuoteLines(t *testing.T) {
	actual := args.Map{"len": len(errcore.StringLinesToQuoteLines([]string{"a", "b"}))}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines", actual)
}

func Test_StringLinesToQuoteLinesToSingle(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"}) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle", actual)
}

func Test_LineDiff(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.LineDiff("actual", "expected") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiff", actual)
}

func Test_MustBeEmpty_Nil(t *testing.T) {
	actual := args.Map{"hasErr": errcore.MustBeEmpty(nil) != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty_Nil", actual)
}

func Test_MustBeEmpty_EmptySlice(t *testing.T) {
	actual := args.Map{"hasErr": errcore.MustBeEmpty([]string{}) != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty_EmptySlice", actual)
}

func Test_MustBeEmpty_NonEmpty(t *testing.T) {
	actual := args.Map{"hasErr": errcore.MustBeEmpty([]string{"a"}) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty_NonEmpty", actual)
}

func Test_GherkinsString(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.GherkinsString("given", "when", "then") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString", actual)
}

func Test_ExpectingFuture(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.ExpectingFuture("header", "expected", "actual") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture", actual)
}

func Test_ExpectingRecord(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.ExpectingRecord("header", 0, "expected", "actual") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord", actual)
}

func Test_RawErrCollection_AddNilAndNonNil(t *testing.T) {
	c := errcore.RawErrCollection{}
	c.Add(nil)
	hasErrAfterNil := c.HasError()
	c.Add(errors.New("err"))
	actual := args.Map{
		"hasErrAfterNil":    hasErrAfterNil,
		"hasErrAfterError":  c.HasError(),
		"lengthAtLeast1":    c.Length() >= 1,
	}
	expected := args.Map{
		"hasErrAfterNil":    false,
		"hasErrAfterError":  true,
		"lengthAtLeast1":    true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection_AddNilAndNonNil", actual)
}
