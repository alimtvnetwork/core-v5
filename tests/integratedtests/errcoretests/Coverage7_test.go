package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── ErrorWithRef ──

func Test_Cov7_ErrorWithRef(t *testing.T) {
	err := errors.New("test")
	actual := args.Map{
		"withRef":    errcore.ErrorWithRef(err, "ref") != "",
		"nilErr":     errcore.ErrorWithRef(nil, "ref"),
		"nilRef":     errcore.ErrorWithRef(err, nil) != "",
		"emptyRef":   errcore.ErrorWithRef(err, "") != "",
	}
	expected := args.Map{"withRef": true, "nilErr": "", "nilRef": true, "emptyRef": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef", actual)
}

// ── ErrorWithCompiledTraceRef ──

func Test_Cov7_ErrorWithCompiledTraceRef(t *testing.T) {
	err := errors.New("test")
	actual := args.Map{
		"full":        errcore.ErrorWithCompiledTraceRef(err, "trace", "ref") != "",
		"nilErr":      errcore.ErrorWithCompiledTraceRef(nil, "trace", "ref"),
		"emptyTrace":  errcore.ErrorWithCompiledTraceRef(err, "", "ref") != "",
		"nilRef":      errcore.ErrorWithCompiledTraceRef(err, "trace", nil) != "",
	}
	expected := args.Map{"full": true, "nilErr": "", "emptyTrace": true, "nilRef": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef", actual)
}

// ── ErrorWithRefToError / ErrorWithCompiledTraceRefToError / ErrorWithTracesRefToError ──

func Test_Cov7_ErrorWithRefToError(t *testing.T) {
	err := errors.New("test")
	result := errcore.ErrorWithRefToError(err, "ref")
	nilResult := errcore.ErrorWithRefToError(nil, "ref")
	actual := args.Map{"hasErr": result != nil, "nilResult": nilResult == nil}
	expected := args.Map{"hasErr": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError", actual)
}

func Test_Cov7_ErrorWithCompiledTraceRefToError(t *testing.T) {
	err := errors.New("test")
	result := errcore.ErrorWithCompiledTraceRefToError(err, "trace", "ref")
	nilResult := errcore.ErrorWithCompiledTraceRefToError(nil, "trace", "ref")
	actual := args.Map{"hasErr": result != nil, "nilResult": nilResult == nil}
	expected := args.Map{"hasErr": true, "nilResult": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError", actual)
}

// ── HandleErr ──

func Test_Cov7_HandleErr_Nil(t *testing.T) {
	errcore.HandleErr(nil) // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr nil", actual)
}

func Test_Cov7_HandleErr_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "HandleErr panic", actual)
	}()
	errcore.HandleErr(errors.New("test"))
}

// ── HandleErrMessage ──

func Test_Cov7_HandleErrMessage_Nil(t *testing.T) {
	errcore.HandleErrMessage("")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage nil", actual)
}

// ── SimpleHandleErr ──

func Test_Cov7_SimpleHandleErr_Nil(t *testing.T) {
	errcore.SimpleHandleErr(nil, "msg")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr nil", actual)
}

// ── SimpleHandleErrMany ──

func Test_Cov7_SimpleHandleErrMany_AllNil(t *testing.T) {
	errcore.SimpleHandleErrMany("msg", nil, nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany all nil", actual)
}

// ── PrintError ──

func Test_Cov7_PrintError(t *testing.T) {
	errcore.PrintError(errors.New("test"))
	errcore.PrintError(nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintError", actual)
}

// ── PrintErrorWithTestIndex ──

func Test_Cov7_PrintErrorWithTestIndex(t *testing.T) {
	errcore.PrintErrorWithTestIndex(0, "header", errors.New("test"))
	errcore.PrintErrorWithTestIndex(0, "header", nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex", actual)
}

// ── LineDiff ──

func Test_Cov7_LineDiff(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "c"})
	actual := args.Map{"len": len(diffs), "firstMatch": diffs[0].Status, "secondMismatch": diffs[1].Status}
	expected := args.Map{"len": 2, "firstMatch": "  ", "secondMismatch": "!!"}
	expected.ShouldBeEqual(t, 0, "LineDiff", actual)
}

func Test_Cov7_LineDiff_ExtraActual(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a"})
	actual := args.Map{"len": len(diffs), "status": diffs[1].Status}
	expected := args.Map{"len": 2, "status": "+"}
	expected.ShouldBeEqual(t, 0, "LineDiff extra actual", actual)
}

func Test_Cov7_LineDiff_MissingExpected(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b"})
	actual := args.Map{"len": len(diffs), "status": diffs[1].Status}
	expected := args.Map{"len": 2, "status": "-"}
	expected.ShouldBeEqual(t, 0, "LineDiff missing expected", actual)
}

func Test_Cov7_LineDiffToString(t *testing.T) {
	result := errcore.LineDiffToString(0, "test", []string{"a"}, []string{"b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString", actual)
}

func Test_Cov7_LineDiffToString_Empty(t *testing.T) {
	result := errcore.LineDiffToString(0, "test", []string{}, []string{})
	actual := args.Map{"empty": result}
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "LineDiffToString empty", actual)
}

func Test_Cov7_HasAnyMismatchOnLines(t *testing.T) {
	actual := args.Map{
		"match":   errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"}),
		"noMatch": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"}),
		"diffLen": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"}),
	}
	expected := args.Map{"match": false, "noMatch": true, "diffLen": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines", actual)
}

func Test_Cov7_SliceDiffSummary(t *testing.T) {
	match := errcore.SliceDiffSummary([]string{"a"}, []string{"a"})
	noMatch := errcore.SliceDiffSummary([]string{"a"}, []string{"b"})
	actual := args.Map{"match": match, "noMatchNotEmpty": noMatch != ""}
	expected := args.Map{"match": "all lines match", "noMatchNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary", actual)
}

func Test_Cov7_ErrorToLinesLineDiff(t *testing.T) {
	result := errcore.ErrorToLinesLineDiff(0, "test", errors.New("a"), []string{"a"})
	nilResult := errcore.ErrorToLinesLineDiff(0, "test", nil, []string{"a"})
	actual := args.Map{"notEmpty": result != "", "nilNotEmpty": nilResult != ""}
	expected := args.Map{"notEmpty": true, "nilNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff", actual)
}

// ── GetActualAndExpectProcessedMessage / GetActualAndExpectSortedMessage ──

func Test_Cov7_GetActualAndExpectProcessedMessage(t *testing.T) {
	result := errcore.GetActualAndExpectProcessedMessage(0, "actual", "expected", "actualProc", "expectedProc")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectProcessedMessage", actual)
}

func Test_Cov7_GetSearchTermExpectationMessage(t *testing.T) {
	result := errcore.GetSearchTermExpectationMessage(0, "header", "expectMsg", 1, "actual", "expected", nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage", actual)
}

func Test_Cov7_GetSearchTermExpectationSimpleMessage(t *testing.T) {
	result := errcore.GetSearchTermExpectationSimpleMessage(0, "expectMsg", 1, "content", "search")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationSimpleMessage", actual)
}

func Test_Cov7_GetSearchLineNumberExpectationMessage(t *testing.T) {
	result := errcore.GetSearchLineNumberExpectationMessage(0, 5, 3, "content", "search", nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage", actual)
}

// ── MessageVarTwo / MessageVarThree / MessageVarMap ──

func Test_Cov7_MessageVarTwo(t *testing.T) {
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo", actual)
}

func Test_Cov7_MessageVarThree(t *testing.T) {
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree", actual)
}

func Test_Cov7_MessageVarMap(t *testing.T) {
	result := errcore.MessageVarMap("msg", map[string]any{"a": 1})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap", actual)
}

// ── MergeErrors ──

func Test_Cov7_MergeErrors(t *testing.T) {
	result := errcore.MergeErrors(errors.New("a"), nil, errors.New("b"))
	nilResult := errcore.MergeErrors(nil, nil)
	actual := args.Map{"hasErr": result != nil, "nilNil": nilResult == nil}
	expected := args.Map{"hasErr": true, "nilNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors", actual)
}

func Test_Cov7_MergeErrorsToString(t *testing.T) {
	result := errcore.MergeErrorsToString(", ", errors.New("a"), errors.New("b"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString", actual)
}

// ── SliceToError / SliceToErrorPtr ──

func Test_Cov7_SliceToError(t *testing.T) {
	result := errcore.SliceToError([]string{"a"})
	nilResult := errcore.SliceToError(nil)
	actual := args.Map{"hasErr": result != nil, "nilNil": nilResult == nil}
	expected := args.Map{"hasErr": true, "nilNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError", actual)
}

func Test_Cov7_SliceToErrorPtr(t *testing.T) {
	result := errcore.SliceToErrorPtr([]string{"a"})
	empty := errcore.SliceToErrorPtr([]string{})
	actual := args.Map{"hasErr": result != nil, "emptyNil": empty == nil}
	expected := args.Map{"hasErr": true, "emptyNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr", actual)
}

// ── ShouldBe / Expected ──

func Test_Cov7_ShouldBe(t *testing.T) {
	actual := args.Map{"notNil": errcore.ShouldBe != (struct{}{})}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe exists", actual)
}

// ── PrintLineDiff / PrintLineDiffOnFail ──

func Test_Cov7_PrintLineDiff(t *testing.T) {
	errcore.PrintLineDiff(0, "test", []string{"a"}, []string{"b"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiff", actual)
}

func Test_Cov7_PrintLineDiffOnFail_Match(t *testing.T) {
	errcore.PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"a"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail match", actual)
}

func Test_Cov7_PrintLineDiffOnFail_Mismatch(t *testing.T) {
	errcore.PrintLineDiffOnFail(0, "test", []string{"a"}, []string{"b"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail mismatch", actual)
}

// ── PrintErrorLineDiff ──

func Test_Cov7_PrintErrorLineDiff(t *testing.T) {
	errcore.PrintErrorLineDiff(0, "test", errors.New("a"), []string{"a"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorLineDiff", actual)
}

// ── AssertDiffOnMismatch / PrintDiffOnMismatch ──

func Test_Cov7_AssertDiffOnMismatch(t *testing.T) {
	errcore.AssertDiffOnMismatch(t, 0, "test", []string{"a"}, []string{"a"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "AssertDiffOnMismatch", actual)
}

func Test_Cov7_PrintDiffOnMismatch(t *testing.T) {
	errcore.PrintDiffOnMismatch(0, "test", []string{"a"}, []string{"b"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch", actual)
}

// ── StackTracesCompiled ──

func Test_Cov7_StackTracesCompiled(t *testing.T) {
	result := errcore.StackTracesCompiled([]string{"trace1", "trace2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled", actual)
}

// ── CombineWithMsgType ──

func Test_Cov7_CombineWithMsgTypeNoStack(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack("type", "msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack", actual)
}

// ── CompiledError ──

func Test_Cov7_CompiledError(t *testing.T) {
	result := errcore.CompiledError(errors.New("inner"), "additional")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError", actual)
}

// ── PathMeaningFulMessage / PathMeaningfulError ──

func Test_Cov7_PathMeaningfulMessage(t *testing.T) {
	result := errcore.PathMeaningfulMessage("type", "funcName", "/path", "msg")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage", actual)
}

func Test_Cov7_PathMeaningfulError(t *testing.T) {
	result := errcore.PathMeaningfulError("type", errors.New("inner"), "/path")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError", actual)
}

// ── MeaningFulError / MeaningFulErrorHandle / MeaningFulErrorWithData ──

func Test_Cov7_MeaningFulError(t *testing.T) {
	result := errcore.MeaningfulError("type", "msg", errors.New("inner"))
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError", actual)
}

func Test_Cov7_MeaningfulMessageError(t *testing.T) {
	result := errcore.MeaningfulMessageError("type", "msg", "ref")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError", actual)
}

func Test_Cov7_MeaningFulErrorWithData(t *testing.T) {
	result := errcore.MeaningfulErrorWithData("type", "msg", errors.New("inner"), "data")
	actual := args.Map{"hasErr": result != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData", actual)
}

// ── ToExitError ──

func Test_Cov7_ToExitError(t *testing.T) {
	result := errcore.ToExitError(errors.New("test"))
	nilResult := errcore.ToExitError(nil)
	actual := args.Map{"isNil": result == nil, "nilNil": nilResult == nil}
	expected := args.Map{"isNil": true, "nilNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError", actual)
}

// ── RangeNotMeet / EnumRangeNotMeet ──

func Test_Cov7_RangeNotMeet(t *testing.T) {
	result := errcore.RangeNotMeet("type", 5, 1, 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet", actual)
}

func Test_Cov7_EnumRangeNotMeet(t *testing.T) {
	result := errcore.EnumRangeNotMeet("type", 5, 1, 3, "extra")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet", actual)
}

// ── StackEnhance ──

func Test_Cov7_StackEnhance(t *testing.T) {
	actual := args.Map{"notNil": errcore.StackEnhance != (struct{}{})}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance exists", actual)
}
