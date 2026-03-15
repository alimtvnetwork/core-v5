package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── CombineWithMsgTypeNoStack ──

func Test_Cov6_CombineWithMsgTypeNoStack(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack("test-type", "test msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack", actual)
}

// ── ConcatMessageWithErr ──

func Test_Cov6_ConcatMessageWithErr(t *testing.T) {
	result := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))
	nilResult := errcore.ConcatMessageWithErr("prefix", nil)
	actual := args.Map{
		"hasResult": result != nil,
		"nilResult": nilResult == nil,
	}
	expected := args.Map{"hasResult": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr", actual)
}

// ── ErrorToSplitLines / ErrorToSplitNonEmptyLines ──

func Test_Cov6_ErrorToSplitLines(t *testing.T) {
	result := errcore.ErrorToSplitLines(errors.New("a\nb\n"))
	nilResult := errcore.ErrorToSplitLines(nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 3, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines", actual)
}

func Test_Cov6_ErrorToSplitNonEmptyLines(t *testing.T) {
	result := errcore.ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))
	nilResult := errcore.ErrorToSplitNonEmptyLines(nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines", actual)
}

// ── ManyErrorToSingle / ManyErrorToSingleDirect ──

func Test_Cov6_ManyErrorToSingle(t *testing.T) {
	errs := []error{errors.New("a"), nil, errors.New("b")}
	result := errcore.ManyErrorToSingle(errs)
	nilResult := errcore.ManyErrorToSingle(nil)
	allNil := errcore.ManyErrorToSingle([]error{nil, nil})
	actual := args.Map{
		"hasErr":    result != nil,
		"nilResult": nilResult == nil,
		"allNil":    allNil == nil,
	}
	expected := args.Map{"hasErr": true, "nilResult": true, "allNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle", actual)
}

func Test_Cov6_ManyErrorToSingleDirect(t *testing.T) {
	result := errcore.ManyErrorToSingleDirect(errors.New("a"), nil, errors.New("b"))
	nilResult := errcore.ManyErrorToSingleDirect()
	actual := args.Map{"hasErr": result != nil, "nilResult": nilResult == nil}
	expected := args.Map{"hasErr": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect", actual)
}

// ── ToError / ToString / ToStringPtr / ToValueString ──

func Test_Cov6_ToError(t *testing.T) {
	result := errcore.ToError("hello")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToError", actual)
}

func Test_Cov6_ToString(t *testing.T) {
	err := errors.New("test")
	result := errcore.ToString(err)
	nilResult := errcore.ToString(nil)
	actual := args.Map{"result": result, "nilResult": nilResult}
	expected := args.Map{"result": "test", "nilResult": ""}
	expected.ShouldBeEqual(t, 0, "ToString", actual)
}

func Test_Cov6_ToStringPtr(t *testing.T) {
	err := errors.New("test")
	result := errcore.ToStringPtr(err)
	nilResult := errcore.ToStringPtr(nil)
	actual := args.Map{"notNil": result != nil, "nilResult": nilResult == nil}
	expected := args.Map{"notNil": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ToStringPtr", actual)
}

func Test_Cov6_ToValueString(t *testing.T) {
	err := errors.New("test")
	result := errcore.ToValueString(err)
	nilResult := errcore.ToValueString(nil)
	actual := args.Map{"result": result, "nilResult": nilResult}
	expected := args.Map{"result": "test", "nilResult": ""}
	expected.ShouldBeEqual(t, 0, "ToValueString", actual)
}

// ── RawErrCollection ──

func Test_Cov6_RawErrCollection(t *testing.T) {
	c := errcore.RawErrCollection{}
	c.Add(errors.New("a"))
	c.Add(errors.New("b"))
	actual := args.Map{
		"len":    c.Length(),
		"hasAny": c.HasAnyError(),
	}
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection", actual)
}

func Test_Cov6_RawErrCollection_CombinedError(t *testing.T) {
	c := errcore.RawErrCollection{}
	c.Add(errors.New("a"))
	result := c.CompiledError()
	empty := errcore.RawErrCollection{}
	emptyResult := empty.CompiledError()
	actual := args.Map{"hasErr": result != nil, "emptyNil": emptyResult == nil}
	expected := args.Map{"hasErr": true, "emptyNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection CombinedError", actual)
}

// ── SliceError / SliceErrorDefault ──

func Test_Cov6_SliceError(t *testing.T) {
	err := errcore.SliceError("|", []string{"a", "b"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceError", actual)
}

func Test_Cov6_SliceError_Empty(t *testing.T) {
	err := errcore.SliceError("|", []string{})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError Empty", actual)
}

func Test_Cov6_SliceErrorDefault(t *testing.T) {
	err := errcore.SliceErrorDefault([]string{"a", "b", "c"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault", actual)
}

// ── SliceErrorsToStrings ──

func Test_Cov6_SliceErrorsToStrings(t *testing.T) {
	errs := []error{errors.New("a"), nil, errors.New("b")}
	result := errcore.SliceErrorsToStrings(errs...)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings", actual)
}

// ── FmtDebug / FmtDebugIf ──

func Test_Cov6_FmtDebug(t *testing.T) {
	// FmtDebug returns void; just verify no panic
	errcore.FmtDebug("hello %s", "world")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug", actual)
}

func Test_Cov6_FmtDebugIf(t *testing.T) {
	errcore.FmtDebugIf(true, "hello %s", "world")
	errcore.FmtDebugIf(false, "hello")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf", actual)
}

// ── Expecting / ExpectingSimple / ExpectingRecord ──

func Test_Cov6_Expecting(t *testing.T) {
	result := errcore.Expecting("header", 42, "expected")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting", actual)
}

func Test_Cov6_ExpectingSimple(t *testing.T) {
	result := errcore.ExpectingSimple("header", 42, "expected")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple", actual)
}

func Test_Cov6_ExpectingSimpleNoType(t *testing.T) {
	result := errcore.ExpectingSimpleNoType("header", 42, "expected")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType", actual)
}

func Test_Cov6_ExpectingError(t *testing.T) {
	result := errcore.ExpectingErrorSimpleNoType("header", 42, "expected")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType", actual)
}

func Test_Cov6_ExpectingNotEqualSimpleNoType(t *testing.T) {
	result := errcore.ExpectingNotEqualSimpleNoType("header", 42, "expected")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType", actual)
}

func Test_Cov6_ExpectingRecord(t *testing.T) {
	rec := &errcore.ExpectingRecord{ExpectingTitle: "header", WasExpecting: "expected"}
	msg := rec.Message("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord", actual)
}

func Test_Cov6_ExpectingFuture(t *testing.T) {
	rec := errcore.ExpectingFuture("header", "expected")
	msg := rec.Message("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture", actual)
}

// ── Var helpers ──

func Test_Cov6_VarTwo(t *testing.T) {
	result := errcore.VarTwo(true, "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo", actual)
}

func Test_Cov6_VarTwoNoType(t *testing.T) {
	result := errcore.VarTwoNoType("a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType", actual)
}

func Test_Cov6_VarThree(t *testing.T) {
	result := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree", actual)
}

func Test_Cov6_VarThreeNoType(t *testing.T) {
	result := errcore.VarThreeNoType("a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType", actual)
}

func Test_Cov6_VarMap(t *testing.T) {
	result := errcore.VarMap(map[string]any{"a": 1})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap", actual)
}

func Test_Cov6_VarMapStrings(t *testing.T) {
	result := errcore.VarMapStrings(map[string]any{"a": "1"})
	actual := args.Map{"hasAny": len(result) > 0}
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "VarMapStrings", actual)
}

// ── MsgHeader ──

func Test_Cov6_MsgHeader(t *testing.T) {
	result := errcore.MsgHeader("header", "msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader", actual)
}

func Test_Cov6_MsgHeaderIf(t *testing.T) {
	result := errcore.MsgHeaderIf(true, "header", "msg")
	empty := errcore.MsgHeaderIf(false, "header", "msg")
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf", actual)
}

func Test_Cov6_MsgHeaderPlusEnding(t *testing.T) {
	result := errcore.MsgHeaderPlusEnding("header", "msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding", actual)
}

// ── Ref ──

func Test_Cov6_Ref(t *testing.T) {
	result := errcore.Ref("context")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref", actual)
}

func Test_Cov6_RefToError(t *testing.T) {
	result := errcore.RefToError("context")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RefToError", actual)
}

// ── GherkinsString ──

func Test_Cov6_GherkinsString(t *testing.T) {
	result := errcore.GherkinsString(0, "title", "given", "when", "then")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString", actual)
}

func Test_Cov6_GherkinsStringWithExpectation(t *testing.T) {
	result := errcore.GherkinsStringWithExpectation(0, "title", "given", "when", "then", "actual", "expected")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation", actual)
}

// ── SourceDestination ──

func Test_Cov6_SourceDestination(t *testing.T) {
	result := errcore.SourceDestination(false, "src", "dst")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination", actual)
}

func Test_Cov6_SourceDestinationErr(t *testing.T) {
	result := errcore.SourceDestinationErr(false, "src", "dst")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr", actual)
}

func Test_Cov6_SourceDestinationNoType(t *testing.T) {
	result := errcore.SourceDestinationNoType("src", "dst")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType", actual)
}

// ── MustBeEmpty / PanicOnIndexOutOfRange / PanicRangeNotMeet ──

func Test_Cov6_MustBeEmpty_NoErr(t *testing.T) {
	errcore.MustBeEmpty(nil) // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty no error", actual)
}

func Test_Cov6_MustBeEmpty_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "MustBeEmpty panic", actual)
	}()
	errcore.MustBeEmpty(errors.New("err"))
}

func Test_Cov6_PanicOnIndexOutOfRange(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange panic", actual)
	}()
	errcore.PanicOnIndexOutOfRange(5, 3)
}

// ── StringLinesToQuoteLines ──

func Test_Cov6_StringLinesToQuoteLines(t *testing.T) {
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines", actual)
}

func Test_Cov6_StringLinesToQuoteLinesToSingle(t *testing.T) {
	result := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle", actual)
}

func Test_Cov6_StringLinesToQuoteLinesWithTabs(t *testing.T) {
	result := errcore.StringLinesToQuoteLinesWithTabs([]string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesWithTabs", actual)
}

// ── MapMismatchError ──

func Test_Cov6_MapMismatchError(t *testing.T) {
	result := errcore.MapMismatchError(
		"ctx", "key", "expected", "actual",
	)
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapMismatchError", actual)
}

// ── MergeErrorsToStringDefault ──

func Test_Cov6_MergeErrorsToStringDefault(t *testing.T) {
	result := errcore.MergeErrorsToStringDefault(errors.New("a"), errors.New("b"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault", actual)
}

// ── MessageWithRef / MessageWithRefToError ──

func Test_Cov6_MessageWithRef(t *testing.T) {
	result := errcore.MessageWithRef("msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef", actual)
}

func Test_Cov6_MessageWithRefToError(t *testing.T) {
	result := errcore.MessageWithRefToError("msg", "ref")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError", actual)
}

// ── CountStateChangeTracker ──

func Test_Cov6_CountStateChangeTracker(t *testing.T) {
	tracker := errcore.CountStateChangeTracker{}
	tracker.SuccessCountChange()
	tracker.FailedCountChange()
	tracker.FailedCountChange()
	actual := args.Map{
		"success": tracker.SuccessCount,
		"failed":  tracker.FailedCount,
		"total":   tracker.Total(),
		"hasAny":  tracker.HasAny(),
	}
	expected := args.Map{"success": 1, "failed": 2, "total": 3, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker", actual)
}

// ── VarNameValues / VarNameValuesJoiner / VarNameValuesStrings ──

func Test_Cov6_VarNameValues(t *testing.T) {
	result := errcore.VarNameValues("a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues", actual)
}

func Test_Cov6_VarNameValuesJoiner(t *testing.T) {
	result := errcore.VarNameValuesJoiner(",", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner", actual)
}

func Test_Cov6_VarNameValuesStrings(t *testing.T) {
	result := errcore.VarNameValuesStrings("a", "1", "b", "2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings", actual)
}

// ── MessageNameValues ──

func Test_Cov6_MessageNameValues(t *testing.T) {
	result := errcore.MessageNameValues("msg", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues", actual)
}

// ── MessageVarTwo / MessageVarThree / MessageVarMap ──

func Test_Cov6_MessageVarTwo(t *testing.T) {
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo", actual)
}

func Test_Cov6_MessageVarThree(t *testing.T) {
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree", actual)
}

func Test_Cov6_MessageVarMap(t *testing.T) {
	result := errcore.MessageVarMap("msg", map[string]any{"a": 1})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap", actual)
}

// ── EnumRangeNotMeet / RangeNotMeet ──

func Test_Cov6_EnumRangeNotMeet(t *testing.T) {
	result := errcore.EnumRangeNotMeet("enumName", 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet", actual)
}

func Test_Cov6_RangeNotMeet(t *testing.T) {
	result := errcore.RangeNotMeet("name", 0, 10, 15)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet", actual)
}
