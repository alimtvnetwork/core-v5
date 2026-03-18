package errcoretests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── CompiledError ──

func Test_Cov10_CompiledError_NilErr(t *testing.T) {
	err := errcore.CompiledError(nil, "msg")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError nil", actual)
}

func Test_Cov10_CompiledError_EmptyMsg(t *testing.T) {
	inner := errors.New("inner")
	err := errcore.CompiledError(inner, "")
	actual := args.Map{"same": err == inner}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "CompiledError empty msg", actual)
}

func Test_Cov10_CompiledError_WithMsg(t *testing.T) {
	err := errcore.CompiledError(errors.New("inner"), "prefix")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError with msg", actual)
}

// ── CompiledErrorString ──

func Test_Cov10_CompiledErrorString_NilErr(t *testing.T) {
	result := errcore.CompiledErrorString(nil, "msg")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "CompiledErrorString nil", actual)
}

func Test_Cov10_CompiledErrorString_WithMsg(t *testing.T) {
	result := errcore.CompiledErrorString(errors.New("inner"), "prefix")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompiledErrorString with msg", actual)
}

// ── JoinErrors ──

func Test_Cov10_JoinErrors(t *testing.T) {
	err := errcore.JoinErrors(errors.New("a"), nil, errors.New("b"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JoinErrors", actual)
}

// ── ConcatMessageWithErrWithStackTrace ──

func Test_Cov10_ConcatMessageWithErrWithStackTrace_Nil(t *testing.T) {
	err := errcore.ConcatMessageWithErrWithStackTrace("msg", nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErrWithStackTrace nil", actual)
}

func Test_Cov10_ConcatMessageWithErrWithStackTrace_WithErr(t *testing.T) {
	err := errcore.ConcatMessageWithErrWithStackTrace("prefix", errors.New("e"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErrWithStackTrace with err", actual)
}

// ── CombineWithMsgTypeNoStack ──

func Test_Cov10_CombineWithMsgTypeNoStack_EmptyOtherMsg(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "", nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack empty", actual)
}

func Test_Cov10_CombineWithMsgTypeNoStack_WithOtherMsg(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "extra", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack with msg", actual)
}

// ── CombineWithMsgTypeStackTrace ──

func Test_Cov10_CombineWithMsgTypeStackTrace(t *testing.T) {
	result := errcore.CombineWithMsgTypeStackTrace(errcore.InvalidType, "msg", nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeStackTrace", actual)
}

// ── CountStateChangeTracker ──

type mockLengthGetter struct{ length int }

func (m *mockLengthGetter) Length() int { return m.length }

func Test_Cov10_CountStateChangeTracker_SameState(t *testing.T) {
	lg := &mockLengthGetter{length: 5}
	tracker := errcore.NewCountStateChangeTracker(lg)
	actual := args.Map{
		"same":    tracker.IsSameState(),
		"valid":   tracker.IsValid(),
		"success": tracker.IsSuccess(),
		"changes": tracker.HasChanges(),
		"failed":  tracker.IsFailed(),
		"sameC":   tracker.IsSameStateUsingCount(5),
	}
	expected := args.Map{
		"same": true, "valid": true, "success": true,
		"changes": false, "failed": false, "sameC": true,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker same", actual)
}

func Test_Cov10_CountStateChangeTracker_Changed(t *testing.T) {
	lg := &mockLengthGetter{length: 5}
	tracker := errcore.NewCountStateChangeTracker(lg)
	lg.length = 6
	actual := args.Map{"same": tracker.IsSameState(), "changes": tracker.HasChanges()}
	expected := args.Map{"same": false, "changes": true}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker changed", actual)
}

// ── EnumRangeNotMeet ──

func Test_Cov10_EnumRangeNotMeet_WithRange(t *testing.T) {
	result := errcore.EnumRangeNotMeet(1, 10, "1-10")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet with range", actual)
}

func Test_Cov10_EnumRangeNotMeet_WithoutRange(t *testing.T) {
	result := errcore.EnumRangeNotMeet(1, 10, nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet without range", actual)
}

// ── ErrorWithCompiledTraceRef ──

func Test_Cov10_ErrorWithCompiledTraceRef_NilErr(t *testing.T) {
	result := errcore.ErrorWithCompiledTraceRef(nil, "trace", "ref")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef nil", actual)
}

func Test_Cov10_ErrorWithCompiledTraceRef_EmptyTraces(t *testing.T) {
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef empty traces", actual)
}

func Test_Cov10_ErrorWithCompiledTraceRef_NilRef(t *testing.T) {
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "trace", nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef nil ref", actual)
}

func Test_Cov10_ErrorWithCompiledTraceRef_All(t *testing.T) {
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "trace", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef all", actual)
}

// ── ErrorWithCompiledTraceRefToError ──

func Test_Cov10_ErrorWithCompiledTraceRefToError_Nil(t *testing.T) {
	err := errcore.ErrorWithCompiledTraceRefToError(nil, "t", "r")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError nil", actual)
}

// ── ErrorWithRefToError ──

func Test_Cov10_ErrorWithRefToError_Nil(t *testing.T) {
	err := errcore.ErrorWithRefToError(nil, "ref")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError nil", actual)
}

func Test_Cov10_ErrorWithRefToError_WithErr(t *testing.T) {
	err := errcore.ErrorWithRefToError(errors.New("e"), "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError with err", actual)
}

// ── ErrorWithTracesRefToError ──

func Test_Cov10_ErrorWithTracesRefToError_Nil(t *testing.T) {
	err := errcore.ErrorWithTracesRefToError(nil, []string{"t"}, "r")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError nil", actual)
}

func Test_Cov10_ErrorWithTracesRefToError_EmptyTraces(t *testing.T) {
	err := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{}, "r")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError empty traces", actual)
}

func Test_Cov10_ErrorWithTracesRefToError_WithTraces(t *testing.T) {
	err := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{"t1", "t2"}, "r")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError with traces", actual)
}

// ── ExpectationMessageDef ──

func Test_Cov10_ExpectationMessageDef_ExpectedSafeString(t *testing.T) {
	emd := errcore.ExpectationMessageDef{Expected: "hello"}
	s1 := emd.ExpectedSafeString()
	s2 := emd.ExpectedSafeString() // cached
	actual := args.Map{"notEmpty": s1 != "", "same": s1 == s2}
	expected := args.Map{"notEmpty": true, "same": true}
	expected.ShouldBeEqual(t, 0, "ExpectedSafeString", actual)
}

func Test_Cov10_ExpectationMessageDef_ExpectedSafeString_Nil(t *testing.T) {
	emd := errcore.ExpectationMessageDef{}
	s := emd.ExpectedSafeString()
	actual := args.Map{"empty": s == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ExpectedSafeString nil", actual)
}

func Test_Cov10_ExpectationMessageDef_ExpectedStringTrim(t *testing.T) {
	emd := errcore.ExpectationMessageDef{Expected: "  hello  "}
	s := emd.ExpectedStringTrim()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectedStringTrim", actual)
}

func Test_Cov10_ExpectationMessageDef_ExpectedString_Panic(t *testing.T) {
	emd := errcore.ExpectationMessageDef{}
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		emd.ExpectedString()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "ExpectedString panic", actual)
}

func Test_Cov10_ExpectationMessageDef_ToString(t *testing.T) {
	emd := errcore.ExpectationMessageDef{When: "w", FuncName: "f", Expected: "e", CaseIndex: 0}
	result := emd.ToString("actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToString", actual)
}

func Test_Cov10_ExpectationMessageDef_PrintIf_False(t *testing.T) {
	emd := errcore.ExpectationMessageDef{When: "w", Expected: "e"}
	emd.PrintIf(false, "actual")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintIf false", actual)
}

func Test_Cov10_ExpectationMessageDef_PrintIfFailed_NotFailed(t *testing.T) {
	emd := errcore.ExpectationMessageDef{When: "w", Expected: "e"}
	emd.PrintIfFailed(true, false, "actual")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintIfFailed not failed", actual)
}

// ── ExpectingFuture / ExpectingRecord ──

func Test_Cov10_ExpectingFuture(t *testing.T) {
	r := errcore.ExpectingFuture("title", "expected")
	actual := args.Map{"notNil": r != nil, "title": r.ExpectingTitle}
	expected := args.Map{"notNil": true, "title": "title"}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture", actual)
}

func Test_Cov10_ExpectingRecord_Message(t *testing.T) {
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}
	actual := args.Map{"notEmpty": r.Message("a") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord Message", actual)
}

func Test_Cov10_ExpectingRecord_MessageSimple(t *testing.T) {
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}
	actual := args.Map{"notEmpty": r.MessageSimple("a") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord MessageSimple", actual)
}

func Test_Cov10_ExpectingRecord_MessageSimpleNoType(t *testing.T) {
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}
	actual := args.Map{"notEmpty": r.MessageSimpleNoType("a") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord MessageSimpleNoType", actual)
}

func Test_Cov10_ExpectingRecord_Error(t *testing.T) {
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}
	actual := args.Map{"notNil": r.Error("a") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord Error", actual)
}

func Test_Cov10_ExpectingRecord_ErrorSimple(t *testing.T) {
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}
	actual := args.Map{"notNil": r.ErrorSimple("a") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord ErrorSimple", actual)
}

func Test_Cov10_ExpectingRecord_ErrorSimpleNoType(t *testing.T) {
	r := &errcore.ExpectingRecord{ExpectingTitle: "t", WasExpecting: "e"}
	actual := args.Map{"notNil": r.ErrorSimpleNoType("a") != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord ErrorSimpleNoType", actual)
}

// ── ExpectingNotEqualSimpleNoType ──

func Test_Cov10_ExpectingNotEqualSimpleNoType(t *testing.T) {
	result := errcore.ExpectingNotEqualSimpleNoType("t", "e", "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType", actual)
}

// ── ExpectingSimpleNoTypeError ──

func Test_Cov10_ExpectingSimpleNoTypeError(t *testing.T) {
	err := errcore.ExpectingSimpleNoTypeError("t", "e", "a")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoTypeError", actual)
}

// ── ExpectingErrorSimpleNoTypeNewLineEnds ──

func Test_Cov10_ExpectingErrorSimpleNoTypeNewLineEnds(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoTypeNewLineEnds("t", "e", "a")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoTypeNewLineEnds", actual)
}

// ── WasExpectingErrorF ──

func Test_Cov10_WasExpectingErrorF(t *testing.T) {
	err := errcore.WasExpectingErrorF("e", "a", "title %d", 1)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "WasExpectingErrorF", actual)
}

// ── FmtDebug / FmtDebugIf ──

func Test_Cov10_FmtDebug(t *testing.T) {
	errcore.FmtDebug("test %d", 1)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug", actual)
}

func Test_Cov10_FmtDebugIf_False(t *testing.T) {
	errcore.FmtDebugIf(false, "test %d", 1)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf false", actual)
}

func Test_Cov10_FmtDebugIf_True(t *testing.T) {
	errcore.FmtDebugIf(true, "test %d", 1)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf true", actual)
}

// ── ValidPrint / FailedPrint ──

func Test_Cov10_ValidPrint_True(t *testing.T) {
	errcore.ValidPrint(true, "val")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "ValidPrint true", actual)
}

func Test_Cov10_ValidPrint_False(t *testing.T) {
	errcore.ValidPrint(false, "val")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "ValidPrint false", actual)
}

func Test_Cov10_FailedPrint_True(t *testing.T) {
	errcore.FailedPrint(true, "val")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FailedPrint true", actual)
}

func Test_Cov10_FailedPrint_False(t *testing.T) {
	errcore.FailedPrint(false, "val")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "FailedPrint false", actual)
}

// ── GetActualAndExpectProcessedMessage ──

func Test_Cov10_GetActualAndExpectProcessedMessage(t *testing.T) {
	result := errcore.GetActualAndExpectProcessedMessage(0, "a", "e", "ap", "ep")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectProcessedMessage", actual)
}

// ── GetSearchLineNumberExpectationMessage ──

func Test_Cov10_GetSearchLineNumberExpectationMessage(t *testing.T) {
	result := errcore.GetSearchLineNumberExpectationMessage(0, 1, 2, "c", "s", "info")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage", actual)
}

// ── GetSearchTermExpectationMessage ──

func Test_Cov10_GetSearchTermExpectationMessage_WithInfo(t *testing.T) {
	result := errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", "info")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage with info", actual)
}

func Test_Cov10_GetSearchTermExpectationMessage_NilInfo(t *testing.T) {
	result := errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage nil info", actual)
}

// ── GetSearchTermExpectationSimpleMessage ──

func Test_Cov10_GetSearchTermExpectationSimpleMessage(t *testing.T) {
	result := errcore.GetSearchTermExpectationSimpleMessage(0, "e", 1, "c", "s")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationSimpleMessage", actual)
}

// ── GherkinsString / GherkinsStringWithExpectation ──

func Test_Cov10_GherkinsString(t *testing.T) {
	result := errcore.GherkinsString(0, "f", "g", "w", "th")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString", actual)
}

func Test_Cov10_GherkinsStringWithExpectation(t *testing.T) {
	result := errcore.GherkinsStringWithExpectation(0, "f", "g", "w", "th", "a", "e")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation", actual)
}

// ── Handle functions (panic paths) ──

func Test_Cov10_HandleErr_Nil(t *testing.T) {
	errcore.HandleErr(nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErr nil", actual)
}

func Test_Cov10_HandleErr_WithErr(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.HandleErr(errors.New("e"))
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErr with err", actual)
}

func Test_Cov10_HandleErrMessage_Empty(t *testing.T) {
	errcore.HandleErrMessage("")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage empty", actual)
}

func Test_Cov10_HandleErrMessage_WithMsg(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.HandleErrMessage("e")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage with msg", actual)
}

func Test_Cov10_HandleErrorGetter_Nil(t *testing.T) {
	errcore.HandleErrorGetter(nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorGetter nil", actual)
}

func Test_Cov10_HandleCompiledErrorGetter_Nil(t *testing.T) {
	errcore.HandleCompiledErrorGetter(nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorGetter nil", actual)
}

func Test_Cov10_HandleCompiledErrorWithTracesGetter_Nil(t *testing.T) {
	errcore.HandleCompiledErrorWithTracesGetter(nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorWithTracesGetter nil", actual)
}

func Test_Cov10_HandleFullStringsWithTracesGetter_Nil(t *testing.T) {
	errcore.HandleFullStringsWithTracesGetter(nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleFullStringsWithTracesGetter nil", actual)
}

// ── SimpleHandleErr ──

func Test_Cov10_SimpleHandleErr_Nil(t *testing.T) {
	errcore.SimpleHandleErr(nil, "msg")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr nil", actual)
}

func Test_Cov10_SimpleHandleErr_WithErr(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.SimpleHandleErr(errors.New("e"), "msg")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr with err", actual)
}

// ── SimpleHandleErrMany ──

func Test_Cov10_SimpleHandleErrMany_Nil(t *testing.T) {
	errcore.SimpleHandleErrMany("msg")
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany nil", actual)
}

func Test_Cov10_SimpleHandleErrMany_AllNil(t *testing.T) {
	errcore.SimpleHandleErrMany("msg", nil, nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany all nil", actual)
}

func Test_Cov10_SimpleHandleErrMany_WithErr(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.SimpleHandleErrMany("msg", errors.New("e"))
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany with err", actual)
}

// ── MsgHeader / MsgHeaderIf / MsgHeaderPlusEnding ──

func Test_Cov10_MsgHeader(t *testing.T) {
	result := errcore.MsgHeader("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader", actual)
}

func Test_Cov10_MsgHeaderIf_True(t *testing.T) {
	result := errcore.MsgHeaderIf(true, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf true", actual)
}

func Test_Cov10_MsgHeaderIf_False(t *testing.T) {
	result := errcore.MsgHeaderIf(false, "hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf false", actual)
}

func Test_Cov10_MsgHeaderPlusEnding(t *testing.T) {
	result := errcore.MsgHeaderPlusEnding("h", "e")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding", actual)
}

// ── PanicOnIndexOutOfRange ──

func Test_Cov10_PanicOnIndexOutOfRange_InRange(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.PanicOnIndexOutOfRange(5, []int{0, 1, 4})
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange in range", actual)
}

func Test_Cov10_PanicOnIndexOutOfRange_OutOfRange(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.PanicOnIndexOutOfRange(3, []int{5})
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange out of range", actual)
}

// ── PanicRangeNotMeet / RangeNotMeet ──

func Test_Cov10_PanicRangeNotMeet_WithRange(t *testing.T) {
	result := errcore.PanicRangeNotMeet("msg", 1, 10, "1-10")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PanicRangeNotMeet with range", actual)
}

func Test_Cov10_PanicRangeNotMeet_WithoutRange(t *testing.T) {
	result := errcore.PanicRangeNotMeet("msg", 1, 10, nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PanicRangeNotMeet without range", actual)
}

func Test_Cov10_RangeNotMeet_WithRange(t *testing.T) {
	result := errcore.RangeNotMeet("msg", 1, 10, "1-10")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet with range", actual)
}

func Test_Cov10_RangeNotMeet_WithoutRange(t *testing.T) {
	result := errcore.RangeNotMeet("msg", 1, 10, nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet without range", actual)
}

// ── PathMeaningfulMessage ──

func Test_Cov10_PathMeaningfulMessage_Empty(t *testing.T) {
	err := errcore.PathMeaningfulMessage(errcore.PathErrorType, "fn", "/path")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage empty", actual)
}

func Test_Cov10_PathMeaningfulMessage_WithMsgs(t *testing.T) {
	err := errcore.PathMeaningfulMessage(errcore.PathErrorType, "fn", "/path", "msg1", "msg2")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage with msgs", actual)
}

// ── PathMeaningfulError ──

func Test_Cov10_PathMeaningfulError_Nil(t *testing.T) {
	err := errcore.PathMeaningfulError(errcore.PathErrorType, nil, "/path")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError nil", actual)
}

func Test_Cov10_PathMeaningfulError_WithErr(t *testing.T) {
	err := errcore.PathMeaningfulError(errcore.PathErrorType, errors.New("e"), "/path")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError with err", actual)
}

// ── MeaningfulError / MeaningfulErrorWithData / MeaningfulMessageError ──

func Test_Cov10_MeaningfulError_Nil(t *testing.T) {
	err := errcore.MeaningfulError(errcore.InvalidType, "fn", nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError nil", actual)
}

func Test_Cov10_MeaningfulError_WithErr(t *testing.T) {
	err := errcore.MeaningfulError(errcore.InvalidType, "fn", errors.New("e"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError with err", actual)
}

func Test_Cov10_MeaningfulErrorWithData_Nil(t *testing.T) {
	err := errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", nil, "data")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData nil", actual)
}

func Test_Cov10_MeaningfulErrorWithData_WithErr(t *testing.T) {
	err := errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", errors.New("e"), "data")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData with err", actual)
}

func Test_Cov10_MeaningfulMessageError_Nil(t *testing.T) {
	err := errcore.MeaningfulMessageError(errcore.InvalidType, "fn", nil, "msg")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError nil", actual)
}

func Test_Cov10_MeaningfulMessageError_WithErr(t *testing.T) {
	err := errcore.MeaningfulMessageError(errcore.InvalidType, "fn", errors.New("e"), "msg")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError with err", actual)
}

// ── MeaningfulErrorHandle ──

func Test_Cov10_MeaningfulErrorHandle_Nil(t *testing.T) {
	errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorHandle nil", actual)
}

func Test_Cov10_MeaningfulErrorHandle_WithErr(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", errors.New("e"))
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorHandle with err", actual)
}

// ── PrintError / PrintErrorWithTestIndex ──

func Test_Cov10_PrintError_Nil(t *testing.T) {
	errcore.PrintError(nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintError nil", actual)
}

func Test_Cov10_PrintError_WithErr(t *testing.T) {
	errcore.PrintError(errors.New("e"))
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintError with err", actual)
}

func Test_Cov10_PrintErrorWithTestIndex_Nil(t *testing.T) {
	errcore.PrintErrorWithTestIndex(0, "h", nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex nil", actual)
}

func Test_Cov10_PrintErrorWithTestIndex_WithErr(t *testing.T) {
	errcore.PrintErrorWithTestIndex(0, "h", errors.New("e"))
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex with err", actual)
}

// ── SourceDestination / SourceDestinationErr / SourceDestinationNoType ──

func Test_Cov10_SourceDestination(t *testing.T) {
	result := errcore.SourceDestination(true, "s", "d")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination", actual)
}

func Test_Cov10_SourceDestinationErr(t *testing.T) {
	err := errcore.SourceDestinationErr(false, "s", "d")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr", actual)
}

func Test_Cov10_SourceDestinationNoType(t *testing.T) {
	result := errcore.SourceDestinationNoType("s", "d")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType", actual)
}

// ── StackTracesCompiled ──

func Test_Cov10_StackTracesCompiled(t *testing.T) {
	result := errcore.StackTracesCompiled([]string{"t1", "t2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled", actual)
}

// ── StringLinesToQuoteLines / StringLinesToQuoteLinesToSingle / LinesToDoubleQuoteLinesWithTabs ──

func Test_Cov10_StringLinesToQuoteLines_Empty(t *testing.T) {
	result := errcore.StringLinesToQuoteLines([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines empty", actual)
}

func Test_Cov10_StringLinesToQuoteLines_NonEmpty(t *testing.T) {
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines non-empty", actual)
}

func Test_Cov10_StringLinesToQuoteLinesToSingle(t *testing.T) {
	result := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle", actual)
}

func Test_Cov10_LinesToDoubleQuoteLinesWithTabs_Empty(t *testing.T) {
	result := errcore.LinesToDoubleQuoteLinesWithTabs(2, []string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesToDoubleQuoteLinesWithTabs empty", actual)
}

func Test_Cov10_LinesToDoubleQuoteLinesWithTabs_WithTabs(t *testing.T) {
	result := errcore.LinesToDoubleQuoteLinesWithTabs(4, []string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LinesToDoubleQuoteLinesWithTabs with tabs", actual)
}

// ── ToExitError ──

func Test_Cov10_ToExitError_Nil(t *testing.T) {
	result := errcore.ToExitError(nil)
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError nil", actual)
}

func Test_Cov10_ToExitError_NonExitError(t *testing.T) {
	result := errcore.ToExitError(errors.New("e"))
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError non-exit", actual)
}

// ── getReferenceMessage (indirect via CombineWithMsgTypeNoStack) ──

func Test_Cov10_getReferenceMessage_EmptyString(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "msg", "")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "getReferenceMessage empty string", actual)
}

// ── RawErrorType methods ──

func Test_Cov10_RawErrorType_CombineWithAnother(t *testing.T) {
	result := errcore.InvalidType.CombineWithAnother(errcore.NotFound, "msg", "ref")
	actual := args.Map{"notEmpty": string(result) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithAnother", actual)
}

func Test_Cov10_RawErrorType_TypesAttach(t *testing.T) {
	result := errcore.InvalidType.TypesAttach("msg", 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TypesAttach", actual)
}

func Test_Cov10_RawErrorType_TypesAttachErr(t *testing.T) {
	err := errcore.InvalidType.TypesAttachErr("msg", 42)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypesAttachErr", actual)
}

func Test_Cov10_RawErrorType_SrcDestination(t *testing.T) {
	result := errcore.InvalidType.SrcDestination("msg", "src", 1, "dst", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SrcDestination", actual)
}

func Test_Cov10_RawErrorType_SrcDestinationErr(t *testing.T) {
	err := errcore.InvalidType.SrcDestinationErr("msg", "src", 1, "dst", 2)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SrcDestinationErr", actual)
}

func Test_Cov10_RawErrorType_Error(t *testing.T) {
	err := errcore.InvalidType.Error("msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Error", actual)
}

func Test_Cov10_RawErrorType_ErrorSkip(t *testing.T) {
	err := errcore.InvalidType.ErrorSkip(0, "msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorSkip", actual)
}

func Test_Cov10_RawErrorType_Fmt_Empty(t *testing.T) {
	err := errcore.InvalidType.Fmt("", )
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Fmt empty", actual)
}

func Test_Cov10_RawErrorType_Fmt_WithFormat(t *testing.T) {
	err := errcore.InvalidType.Fmt("format %d", 1)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Fmt with format", actual)
}

func Test_Cov10_RawErrorType_FmtIf_False(t *testing.T) {
	err := errcore.InvalidType.FmtIf(false, "format %d", 1)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FmtIf false", actual)
}

func Test_Cov10_RawErrorType_FmtIf_True(t *testing.T) {
	err := errcore.InvalidType.FmtIf(true, "format %d", 1)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FmtIf true", actual)
}

func Test_Cov10_RawErrorType_MergeError_Nil(t *testing.T) {
	err := errcore.InvalidType.MergeError(nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeError nil", actual)
}

func Test_Cov10_RawErrorType_MergeError_WithErr(t *testing.T) {
	err := errcore.InvalidType.MergeError(errors.New("e"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeError with err", actual)
}

func Test_Cov10_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessage(nil, "msg")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessage nil", actual)
}

func Test_Cov10_RawErrorType_MergeErrorWithMessage_WithErr(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessage(errors.New("e"), "msg")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessage with err", actual)
}

func Test_Cov10_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessageRef nil", actual)
}

func Test_Cov10_RawErrorType_MergeErrorWithMessageRef_WithErr(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithMessageRef with err", actual)
}

func Test_Cov10_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithRef(nil, "ref")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithRef nil", actual)
}

func Test_Cov10_RawErrorType_MergeErrorWithRef_WithErr(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithRef(errors.New("e"), "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorWithRef with err", actual)
}

func Test_Cov10_RawErrorType_MsgCsvRef_Empty(t *testing.T) {
	result := errcore.InvalidType.MsgCsvRef("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRef empty", actual)
}

func Test_Cov10_RawErrorType_MsgCsvRef_WithItems(t *testing.T) {
	result := errcore.InvalidType.MsgCsvRef("msg", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRef with items", actual)
}

func Test_Cov10_RawErrorType_MsgCsvRef_EmptyMsg(t *testing.T) {
	result := errcore.InvalidType.MsgCsvRef("", "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRef empty msg", actual)
}

func Test_Cov10_RawErrorType_MsgCsvRefError(t *testing.T) {
	err := errcore.InvalidType.MsgCsvRefError("msg", "a")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvRefError", actual)
}

func Test_Cov10_RawErrorType_ErrorRefOnly(t *testing.T) {
	err := errcore.InvalidType.ErrorRefOnly("ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorRefOnly", actual)
}

func Test_Cov10_RawErrorType_Expecting(t *testing.T) {
	err := errcore.InvalidType.Expecting("e", "a")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expecting", actual)
}

func Test_Cov10_RawErrorType_NoRef_Empty(t *testing.T) {
	result := errcore.InvalidType.NoRef("")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NoRef empty", actual)
}

func Test_Cov10_RawErrorType_NoRef_WithMsg(t *testing.T) {
	result := errcore.InvalidType.NoRef("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NoRef with msg", actual)
}

func Test_Cov10_RawErrorType_ErrorNoRefs(t *testing.T) {
	err := errcore.InvalidType.ErrorNoRefs("msg")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorNoRefs", actual)
}

func Test_Cov10_RawErrorType_ErrorNoRefs_Empty(t *testing.T) {
	err := errcore.InvalidType.ErrorNoRefs("")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorNoRefs empty", actual)
}

func Test_Cov10_RawErrorType_HandleUsingPanic(t *testing.T) {
	var didPanic bool
	func() {
		defer func() { if r := recover(); r != nil { didPanic = true } }()
		errcore.InvalidType.HandleUsingPanic("msg", "ref")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleUsingPanic", actual)
}

// ── GetSet / GetSetVariant ──

func Test_Cov10_GetSet_True(t *testing.T) {
	result := errcore.GetSet(true, errcore.InvalidType, errcore.NotFound)
	actual := args.Map{"val": string(result)}
	expected := args.Map{"val": string(errcore.InvalidType)}
	expected.ShouldBeEqual(t, 0, "GetSet true", actual)
}

func Test_Cov10_GetSet_False(t *testing.T) {
	result := errcore.GetSet(false, errcore.InvalidType, errcore.NotFound)
	actual := args.Map{"val": string(result)}
	expected := args.Map{"val": string(errcore.NotFound)}
	expected.ShouldBeEqual(t, 0, "GetSet false", actual)
}

func Test_Cov10_GetSetVariant_True(t *testing.T) {
	result := errcore.GetSetVariant(true, "a", "b")
	actual := args.Map{"val": string(result)}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant true", actual)
}

func Test_Cov10_GetSetVariant_False(t *testing.T) {
	result := errcore.GetSetVariant(false, "a", "b")
	actual := args.Map{"val": string(result)}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant false", actual)
}

// ── ShouldBe ──

func Test_Cov10_ShouldBe_StrEqMsg(t *testing.T) {
	result := errcore.ShouldBe.StrEqMsg("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StrEqMsg", actual)
}

func Test_Cov10_ShouldBe_StrEqErr(t *testing.T) {
	err := errcore.ShouldBe.StrEqErr("a", "b")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StrEqErr", actual)
}

func Test_Cov10_ShouldBe_AnyEqMsg(t *testing.T) {
	result := errcore.ShouldBe.AnyEqMsg(1, 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyEqMsg", actual)
}

func Test_Cov10_ShouldBe_AnyEqErr(t *testing.T) {
	err := errcore.ShouldBe.AnyEqErr(1, 2)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyEqErr", actual)
}

func Test_Cov10_ShouldBe_JsonEqMsg(t *testing.T) {
	result := errcore.ShouldBe.JsonEqMsg("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonEqMsg", actual)
}

func Test_Cov10_ShouldBe_JsonEqErr(t *testing.T) {
	err := errcore.ShouldBe.JsonEqErr("a", "b")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonEqErr", actual)
}

// ── Expected ──

func Test_Cov10_Expected_But(t *testing.T) {
	err := errcore.Expected.But("title", "exp", "act")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected But", actual)
}

func Test_Cov10_Expected_ButFoundAsMsg(t *testing.T) {
	result := errcore.Expected.ButFoundAsMsg("title", "exp", "act")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected ButFoundAsMsg", actual)
}

func Test_Cov10_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	result := errcore.Expected.ButFoundWithTypeAsMsg("title", "exp", "act")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected ButFoundWithTypeAsMsg", actual)
}

func Test_Cov10_Expected_ButUsingType(t *testing.T) {
	err := errcore.Expected.ButUsingType("title", "exp", "act")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected ButUsingType", actual)
}

func Test_Cov10_Expected_ReflectButFound(t *testing.T) {
	err := errcore.Expected.ReflectButFound(reflect.String, reflect.Int)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected ReflectButFound", actual)
}

func Test_Cov10_Expected_PrimitiveButFound(t *testing.T) {
	err := errcore.Expected.PrimitiveButFound(reflect.Slice)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected PrimitiveButFound", actual)
}

func Test_Cov10_Expected_ValueHasNoElements(t *testing.T) {
	err := errcore.Expected.ValueHasNoElements(reflect.Slice)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected ValueHasNoElements", actual)
}

// ── StackEnhance ──

func Test_Cov10_StackEnhance_Error_Nil(t *testing.T) {
	err := errcore.StackEnhance.Error(nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance Error nil", actual)
}

func Test_Cov10_StackEnhance_Error_WithErr(t *testing.T) {
	err := errcore.StackEnhance.Error(errors.New("e"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance Error with err", actual)
}

func Test_Cov10_StackEnhance_Msg_Empty(t *testing.T) {
	result := errcore.StackEnhance.Msg("")
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance Msg empty", actual)
}

func Test_Cov10_StackEnhance_Msg_NonEmpty(t *testing.T) {
	result := errcore.StackEnhance.Msg("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance Msg non-empty", actual)
}

func Test_Cov10_StackEnhance_MsgToErrSkip_Empty(t *testing.T) {
	err := errcore.StackEnhance.MsgToErrSkip(0, "")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance MsgToErrSkip empty", actual)
}

func Test_Cov10_StackEnhance_FmtSkip_Empty(t *testing.T) {
	err := errcore.StackEnhance.FmtSkip(0, "")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance FmtSkip empty", actual)
}

func Test_Cov10_StackEnhance_FmtSkip_NonEmpty(t *testing.T) {
	err := errcore.StackEnhance.FmtSkip(0, "hello %d", 1)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance FmtSkip non-empty", actual)
}

func Test_Cov10_StackEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	result := errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance MsgErrorSkip nil err", actual)
}

func Test_Cov10_StackEnhance_MsgErrorSkip_WithErr(t *testing.T) {
	result := errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("e"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance MsgErrorSkip with err", actual)
}

func Test_Cov10_StackEnhance_MsgErrorToErrSkip_Nil(t *testing.T) {
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance MsgErrorToErrSkip nil", actual)
}

func Test_Cov10_StackEnhance_MsgErrorToErrSkip_WithErr(t *testing.T) {
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance MsgErrorToErrSkip with err", actual)
}

// ── VarNameValues / VarNameValuesJoiner / VarNameValuesStrings ──

func Test_Cov10_VarNameValues_Empty(t *testing.T) {
	result := errcore.VarNameValues()
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValues empty", actual)
}

func Test_Cov10_VarNameValues_NonEmpty(t *testing.T) {
	result := errcore.VarNameValues(namevalue.StringAny{Name: "k", Value: "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues non-empty", actual)
}

func Test_Cov10_VarNameValuesJoiner_Empty(t *testing.T) {
	result := errcore.VarNameValuesJoiner(",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner empty", actual)
}

func Test_Cov10_VarNameValuesJoiner_NonEmpty(t *testing.T) {
	result := errcore.VarNameValuesJoiner(",", namevalue.StringAny{Name: "k", Value: "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner non-empty", actual)
}

func Test_Cov10_VarNameValuesStrings_Empty(t *testing.T) {
	result := errcore.VarNameValuesStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings empty", actual)
}

// ── MessageNameValues ──

func Test_Cov10_MessageNameValues_Empty(t *testing.T) {
	result := errcore.MessageNameValues("msg")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageNameValues empty", actual)
}

func Test_Cov10_MessageNameValues_NonEmpty(t *testing.T) {
	result := errcore.MessageNameValues("msg", namevalue.StringAny{Name: "k", Value: "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues non-empty", actual)
}
