package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ══════════════════════════════════════════════════════════════════════════════
// errcore Coverage — Batch 1: Utility functions, formatters, converters
// ══════════════════════════════════════════════════════════════════════════════

// --- CountStateChangeTracker ---

type mockLengthGetter struct {
	length int
}

func (m *mockLengthGetter) Length() int {
	return m.length
}

func Test_CovErr_01_CountStateChangeTracker(t *testing.T) {
	mg := &mockLengthGetter{length: 5}
	tracker := errcore.NewCountStateChangeTracker(mg)
	if !tracker.IsSameState() {
		t.Fatal("expected same state")
	}
	if !tracker.IsValid() {
		t.Fatal("expected valid")
	}
	if !tracker.IsSuccess() {
		t.Fatal("expected success")
	}
	if tracker.HasChanges() {
		t.Fatal("expected no changes")
	}
	if tracker.IsFailed() {
		t.Fatal("expected not failed")
	}
	if !tracker.IsSameStateUsingCount(5) {
		t.Fatal("expected same state using count")
	}
	if tracker.IsSameStateUsingCount(3) {
		t.Fatal("expected not same state")
	}
	// change length
	mg.length = 10
	if tracker.IsSameState() {
		t.Fatal("expected different state")
	}
	if !tracker.HasChanges() {
		t.Fatal("expected changes")
	}
	if !tracker.IsFailed() {
		t.Fatal("expected failed")
	}
}

// --- CombineWithMsgType ---

func Test_CovErr_02_CombineWithMsgTypeNoStack(t *testing.T) {
	r := errcore.CombineWithMsgTypeNoStack(errcore.OutOfRangeType, "extra", "ref")
	if r == "" {
		t.Fatal("expected non-empty")
	}
	// empty otherMsg
	r2 := errcore.CombineWithMsgTypeNoStack(errcore.OutOfRangeType, "", "ref")
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
	// nil reference
	r3 := errcore.CombineWithMsgTypeNoStack(errcore.OutOfRangeType, "msg", nil)
	if r3 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_03_CombineWithMsgTypeStackTrace(t *testing.T) {
	r := errcore.CombineWithMsgTypeStackTrace(errcore.OutOfRangeType, "msg", "ref")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

// --- ConcatMessageWithErr ---

func Test_CovErr_04_ConcatMessageWithErr(t *testing.T) {
	err := errcore.ConcatMessageWithErr("prefix", errors.New("original"))
	if err == nil {
		t.Fatal("expected error")
	}
	// nil err
	err2 := errcore.ConcatMessageWithErr("prefix", nil)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovErr_05_ConcatMessageWithErrWithStackTrace(t *testing.T) {
	err := errcore.ConcatMessageWithErrWithStackTrace("prefix", errors.New("orig"))
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.ConcatMessageWithErrWithStackTrace("prefix", nil)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

// --- Combine ---

func Test_CovErr_06_Combine(t *testing.T) {
	r := errcore.Combine("generic", "other", "ref")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

// --- EnumRangeNotMeet ---

func Test_CovErr_07_EnumRangeNotMeet(t *testing.T) {
	r := errcore.EnumRangeNotMeet(1, 10, []int{1, 2, 3})
	if r == "" {
		t.Fatal("expected non-empty")
	}
	// nil wholeRange
	r2 := errcore.EnumRangeNotMeet(1, 10, nil)
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

// --- ErrorToSplitLines ---

func Test_CovErr_08_ErrorToSplitLines(t *testing.T) {
	lines := errcore.ErrorToSplitLines(errors.New("a\nb"))
	if len(lines) != 2 {
		t.Fatal("expected 2")
	}
	lines2 := errcore.ErrorToSplitLines(nil)
	if len(lines2) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovErr_09_ErrorToSplitNonEmptyLines(t *testing.T) {
	lines := errcore.ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))
	if len(lines) < 2 {
		t.Fatal("expected at least 2")
	}
}

// --- ErrorWithRef ---

func Test_CovErr_10_ErrorWithRef(t *testing.T) {
	r := errcore.ErrorWithRef(errors.New("err"), "ref")
	if r == "" {
		t.Fatal("expected non-empty")
	}
	// nil err
	r2 := errcore.ErrorWithRef(nil, "ref")
	if r2 != "" {
		t.Fatal("expected empty")
	}
	// nil ref
	r3 := errcore.ErrorWithRef(errors.New("err"), nil)
	if r3 == "" {
		t.Fatal("expected non-empty")
	}
	// empty ref
	r4 := errcore.ErrorWithRef(errors.New("err"), "")
	if r4 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_11_ErrorWithRefToError(t *testing.T) {
	err := errcore.ErrorWithRefToError(errors.New("err"), "ref")
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.ErrorWithRefToError(nil, "ref")
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

// --- ErrorWithCompiledTraceRef ---

func Test_CovErr_12_ErrorWithCompiledTraceRef(t *testing.T) {
	// nil err
	r := errcore.ErrorWithCompiledTraceRef(nil, "trace", "ref")
	if r != "" {
		t.Fatal("expected empty")
	}
	// empty traces
	r2 := errcore.ErrorWithCompiledTraceRef(errors.New("err"), "", "ref")
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
	// nil reference
	r3 := errcore.ErrorWithCompiledTraceRef(errors.New("err"), "trace", nil)
	if r3 == "" {
		t.Fatal("expected non-empty")
	}
	// all present
	r4 := errcore.ErrorWithCompiledTraceRef(errors.New("err"), "trace", "ref")
	if r4 == "" {
		t.Fatal("expected non-empty")
	}
}

// --- MeaningfulError ---

func Test_CovErr_13_MeaningfulError(t *testing.T) {
	err := errcore.MeaningfulError(errcore.OutOfRangeType, "func", errors.New("orig"))
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.MeaningfulError(errcore.OutOfRangeType, "func", nil)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovErr_14_MeaningfulErrorWithData(t *testing.T) {
	err := errcore.MeaningfulErrorWithData(errcore.OutOfRangeType, "func", errors.New("orig"), "data")
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.MeaningfulErrorWithData(errcore.OutOfRangeType, "func", nil, "data")
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovErr_15_MeaningfulMessageError(t *testing.T) {
	err := errcore.MeaningfulMessageError(errcore.OutOfRangeType, "func", errors.New("orig"), " extra")
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.MeaningfulMessageError(errcore.OutOfRangeType, "func", nil, "extra")
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

// --- MeaningfulErrorHandle ---

func Test_CovErr_16_MeaningfulErrorHandle_NilSafe(t *testing.T) {
	// nil should not panic
	errcore.MeaningfulErrorHandle(errcore.OutOfRangeType, "func", nil)
}

// --- PathMeaningfulMessage ---

func Test_CovErr_17_PathMeaningfulMessage(t *testing.T) {
	err := errcore.PathMeaningfulMessage(errcore.OutOfRangeType, "func", "/path", "msg1", "msg2")
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.PathMeaningfulMessage(errcore.OutOfRangeType, "func", "/path")
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

// --- PathMeaningfulError ---

func Test_CovErr_18_PathMeaningfulError(t *testing.T) {
	err := errcore.PathMeaningfulError(errcore.OutOfRangeType, errors.New("orig"), "/path")
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.PathMeaningfulError(errcore.OutOfRangeType, nil, "/path")
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

// --- MergeErrors ---

func Test_CovErr_19_MergeErrors(t *testing.T) {
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.MergeErrors()
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovErr_20_MergeErrorsToString(t *testing.T) {
	r := errcore.MergeErrorsToString(", ", errors.New("a"), errors.New("b"))
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.MergeErrorsToString(", ")
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovErr_21_MergeErrorsToStringDefault(t *testing.T) {
	r := errcore.MergeErrorsToStringDefault(errors.New("a"))
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.MergeErrorsToStringDefault()
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

// --- ManyErrorToSingle ---

func Test_CovErr_22_ManyErrorToSingle(t *testing.T) {
	err := errcore.ManyErrorToSingle([]error{errors.New("a")})
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.ManyErrorToSingle(nil)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovErr_23_ManyErrorToSingleDirect(t *testing.T) {
	err := errcore.ManyErrorToSingleDirect(errors.New("a"), errors.New("b"))
	if err == nil {
		t.Fatal("expected error")
	}
}

// --- SliceError ---

func Test_CovErr_24_SliceError(t *testing.T) {
	err := errcore.SliceError(", ", []string{"a", "b"})
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.SliceError(", ", []string{})
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovErr_25_SliceErrorDefault(t *testing.T) {
	err := errcore.SliceErrorDefault([]string{"a"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovErr_26_SliceErrorsToStrings(t *testing.T) {
	ss := errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b"))
	if len(ss) != 2 {
		t.Fatal("expected 2")
	}
	ss2 := errcore.SliceErrorsToStrings()
	if len(ss2) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovErr_27_SliceToError_SliceToErrorPtr(t *testing.T) {
	err := errcore.SliceToError([]string{"a"})
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.SliceToError([]string{})
	if err2 != nil {
		t.Fatal("expected nil")
	}
	err3 := errcore.SliceToErrorPtr([]string{"a"})
	if err3 == nil {
		t.Fatal("expected error")
	}
	err4 := errcore.SliceToErrorPtr([]string{})
	if err4 != nil {
		t.Fatal("expected nil")
	}
}

// --- VarTwo / VarThree ---

func Test_CovErr_28_VarTwo(t *testing.T) {
	r := errcore.VarTwo(true, "a", 1, "b", 2)
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.VarTwo(false, "a", 1, "b", 2)
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_29_VarTwoNoType(t *testing.T) {
	r := errcore.VarTwoNoType("a", 1, "b", 2)
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_30_VarThree(t *testing.T) {
	r := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_31_VarThreeNoType(t *testing.T) {
	r := errcore.VarThreeNoType("a", 1, "b", 2, "c", 3)
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

// --- VarMap / VarMapStrings / VarNameValues ---

func Test_CovErr_32_VarMap(t *testing.T) {
	r := errcore.VarMap(map[string]any{"a": 1})
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.VarMap(map[string]any{})
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovErr_33_VarMapStrings(t *testing.T) {
	ss := errcore.VarMapStrings(map[string]any{"a": 1})
	if len(ss) != 1 {
		t.Fatal("expected 1")
	}
	ss2 := errcore.VarMapStrings(map[string]any{})
	if len(ss2) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovErr_34_VarNameValues(t *testing.T) {
	nv := namevalue.StringAny{Name: "a", Value: 1}
	r := errcore.VarNameValues(nv)
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.VarNameValues()
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovErr_35_VarNameValuesJoiner(t *testing.T) {
	nv := namevalue.StringAny{Name: "a", Value: 1}
	r := errcore.VarNameValuesJoiner(", ", nv)
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.VarNameValuesJoiner(", ")
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovErr_36_VarNameValuesStrings(t *testing.T) {
	nv := namevalue.StringAny{Name: "a", Value: 1}
	ss := errcore.VarNameValuesStrings(nv)
	if len(ss) != 1 {
		t.Fatal("expected 1")
	}
	ss2 := errcore.VarNameValuesStrings()
	if len(ss2) != 0 {
		t.Fatal("expected 0")
	}
}

// --- MessageVarTwo / MessageVarThree / MessageVarMap / MessageNameValues ---

func Test_CovErr_37_MessageVarTwo(t *testing.T) {
	r := errcore.MessageVarTwo("msg", "a", 1, "b", 2)
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_38_MessageVarThree(t *testing.T) {
	r := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_39_MessageVarMap(t *testing.T) {
	r := errcore.MessageVarMap("msg", map[string]any{"a": 1})
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.MessageVarMap("msg", map[string]any{})
	if r2 != "msg" {
		t.Fatal("expected just msg")
	}
}

func Test_CovErr_40_MessageNameValues(t *testing.T) {
	nv := namevalue.StringAny{Name: "a", Value: 1}
	r := errcore.MessageNameValues("msg", nv)
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.MessageNameValues("msg")
	if r2 != "msg" {
		t.Fatal("expected just msg")
	}
}

// --- MessageWithRef / MessageWithRefToError ---

func Test_CovErr_41_MessageWithRef(t *testing.T) {
	r := errcore.MessageWithRef("msg", "ref")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_42_MessageWithRefToError(t *testing.T) {
	err := errcore.MessageWithRefToError("msg", "ref")
	if err == nil {
		t.Fatal("expected error")
	}
}

// --- SourceDestination ---

func Test_CovErr_43_SourceDestination(t *testing.T) {
	r := errcore.SourceDestination(true, "src", "dst")
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.SourceDestination(false, "src", "dst")
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_44_SourceDestinationErr(t *testing.T) {
	err := errcore.SourceDestinationErr(false, "src", "dst")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_CovErr_45_SourceDestinationNoType(t *testing.T) {
	r := errcore.SourceDestinationNoType("src", "dst")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

// --- ToError / ToString / ToStringPtr / ToValueString ---

func Test_CovErr_46_ToError(t *testing.T) {
	err := errcore.ToError("msg")
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.ToError("")
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovErr_47_ToString(t *testing.T) {
	r := errcore.ToString(errors.New("msg"))
	if r != "msg" {
		t.Fatal("expected msg")
	}
	r2 := errcore.ToString(nil)
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovErr_48_ToStringPtr(t *testing.T) {
	r := errcore.ToStringPtr(errors.New("msg"))
	if *r != "msg" {
		t.Fatal("expected msg")
	}
	r2 := errcore.ToStringPtr(nil)
	if *r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovErr_49_ToValueString(t *testing.T) {
	r := errcore.ToValueString("hello")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_50_ToExitError(t *testing.T) {
	r := errcore.ToExitError(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
	r2 := errcore.ToExitError(errors.New("not exit error"))
	if r2 != nil {
		t.Fatal("expected nil for non-exit error")
	}
}

// --- StringLinesToQuoteLines ---

func Test_CovErr_51_StringLinesToQuoteLines(t *testing.T) {
	ss := errcore.StringLinesToQuoteLines([]string{"a", "b"})
	if len(ss) != 2 {
		t.Fatal("expected 2")
	}
	ss2 := errcore.StringLinesToQuoteLines([]string{})
	if len(ss2) != 0 {
		t.Fatal("expected 0")
	}
}

func Test_CovErr_52_StringLinesToQuoteLinesToSingle(t *testing.T) {
	r := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_53_LinesToDoubleQuoteLinesWithTabs(t *testing.T) {
	ss := errcore.LinesToDoubleQuoteLinesWithTabs(4, []string{"a"})
	if len(ss) != 1 {
		t.Fatal("expected 1")
	}
	ss2 := errcore.LinesToDoubleQuoteLinesWithTabs(0, []string{})
	if len(ss2) != 0 {
		t.Fatal("expected 0")
	}
}

// --- MsgHeader ---

func Test_CovErr_54_MsgHeader(t *testing.T) {
	r := errcore.MsgHeader("test")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_55_MsgHeaderIf(t *testing.T) {
	r := errcore.MsgHeaderIf(true, "test")
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.MsgHeaderIf(false, "test")
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_56_MsgHeaderPlusEnding(t *testing.T) {
	r := errcore.MsgHeaderPlusEnding("header", "message")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

// --- MustBeEmpty ---

func Test_CovErr_57_MustBeEmpty_NilSafe(t *testing.T) {
	errcore.MustBeEmpty(nil)
}

// --- HandleErr / HandleErrMessage ---

func Test_CovErr_58_HandleErr_NilSafe(t *testing.T) {
	errcore.HandleErr(nil)
}

func Test_CovErr_59_HandleErrMessage_NilSafe(t *testing.T) {
	errcore.HandleErrMessage("")
}

// --- SimpleHandleErr ---

func Test_CovErr_60_SimpleHandleErr_NilSafe(t *testing.T) {
	errcore.SimpleHandleErr(nil, "msg")
}

// --- SimpleHandleErrMany ---

func Test_CovErr_61_SimpleHandleErrMany_NilSafe(t *testing.T) {
	errcore.SimpleHandleErrMany("msg")
	errcore.SimpleHandleErrMany("msg", nil)
}

// --- PanicOnIndexOutOfRange ---

func Test_CovErr_62_PanicOnIndexOutOfRange_Valid(t *testing.T) {
	errcore.PanicOnIndexOutOfRange(5, []int{0, 1, 4})
}

// --- RangeNotMeet / PanicRangeNotMeet ---

func Test_CovErr_63_RangeNotMeet(t *testing.T) {
	r := errcore.RangeNotMeet("msg", 0, 10, nil)
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.RangeNotMeet("msg", 0, 10, "range")
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_64_PanicRangeNotMeet(t *testing.T) {
	r := errcore.PanicRangeNotMeet("msg", 0, 10, nil)
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.PanicRangeNotMeet("msg", 0, 10, "range")
	if r2 == "" {
		t.Fatal("expected non-empty")
	}
}

// --- Ref / RefToError ---

func Test_CovErr_65_Ref(t *testing.T) {
	r := errcore.Ref("ref")
	if r == "" {
		t.Fatal("expected non-empty")
	}
	r2 := errcore.Ref(nil)
	if r2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovErr_66_RefToError(t *testing.T) {
	err := errcore.RefToError("ref")
	if err == nil {
		t.Fatal("expected error")
	}
	err2 := errcore.RefToError(nil)
	if err2 != nil {
		t.Fatal("expected nil")
	}
}

// --- GherkinsString ---

func Test_CovErr_67_GherkinsString(t *testing.T) {
	r := errcore.GherkinsString(1, "feature", "given", "when", "then")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovErr_68_GherkinsStringWithExpectation(t *testing.T) {
	r := errcore.GherkinsStringWithExpectation(1, "feature", "given", "when", "then", "actual", "expect")
	if r == "" {
		t.Fatal("expected non-empty")
	}
}

// --- FmtDebug / FmtDebugIf / ValidPrint / FailedPrint / PrintError ---

func Test_CovErr_69_FmtDebug(t *testing.T) {
	errcore.FmtDebug("test %d", 1)
}

func Test_CovErr_70_FmtDebugIf(t *testing.T) {
	errcore.FmtDebugIf(false, "test %d", 1)
	errcore.FmtDebugIf(true, "test %d", 1)
}

func Test_CovErr_71_ValidPrint(t *testing.T) {
	errcore.ValidPrint(true, "val")
	errcore.ValidPrint(false, "val")
}

func Test_CovErr_72_FailedPrint(t *testing.T) {
	errcore.FailedPrint(true, "val")
	errcore.FailedPrint(false, "val")
}

func Test_CovErr_73_PrintError(t *testing.T) {
	errcore.PrintError(nil)
	errcore.PrintError(errors.New("err"))
}

func Test_CovErr_74_PrintErrorWithTestIndex(t *testing.T) {
	errcore.PrintErrorWithTestIndex(0, "header", nil)
	errcore.PrintErrorWithTestIndex(0, "header", errors.New("err"))
}
