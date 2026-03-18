package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ══════════════════════════════════════════════════════════════════════════════
// LineDiff / LineDiffToString / PrintLineDiff / HasAnyMismatchOnLines
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_LineDiff_AllMatch(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})
	actual := args.Map{"len": len(diffs), "st0": diffs[0].Status, "st1": diffs[1].Status}
	expected := args.Map{"len": 2, "st0": "  ", "st1": "  "}
	expected.ShouldBeEqual(t, 0, "LineDiff all match", actual)
}

func Test_Cov11_LineDiff_Mismatch(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a", "x"}, []string{"a", "b"})
	actual := args.Map{"st1": diffs[1].Status}
	expected := args.Map{"st1": "!!"}
	expected.ShouldBeEqual(t, 0, "LineDiff mismatch", actual)
}

func Test_Cov11_LineDiff_ExtraActual(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a", "b", "c"}, []string{"a"})
	actual := args.Map{"st2": diffs[2].Status}
	expected := args.Map{"st2": "+"}
	expected.ShouldBeEqual(t, 0, "LineDiff extra actual", actual)
}

func Test_Cov11_LineDiff_MissingExpected(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b", "c"})
	actual := args.Map{"st2": diffs[2].Status}
	expected := args.Map{"st2": "-"}
	expected.ShouldBeEqual(t, 0, "LineDiff missing expected", actual)
}

func Test_Cov11_LineDiffToString_Empty(t *testing.T) {
	result := errcore.LineDiffToString(0, "h", []string{}, []string{})
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString empty", actual)
}

func Test_Cov11_LineDiffToString_WithDiffs(t *testing.T) {
	result := errcore.LineDiffToString(0, "h", []string{"a", "x"}, []string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString with diffs", actual)
}

func Test_Cov11_LineDiffToString_AllBranches(t *testing.T) {
	result := errcore.LineDiffToString(0, "h", []string{"a", "x", "extra"}, []string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString all branches", actual)
}

func Test_Cov11_PrintLineDiff(t *testing.T) {
	errcore.PrintLineDiff(0, "h", []string{"a"}, []string{"b"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiff", actual)
}

func Test_Cov11_PrintLineDiff_Empty(t *testing.T) {
	errcore.PrintLineDiff(0, "h", []string{}, []string{})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiff empty", actual)
}

func Test_Cov11_HasAnyMismatchOnLines_Match(t *testing.T) {
	actual := args.Map{"v": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"})}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines match", actual)
}

func Test_Cov11_HasAnyMismatchOnLines_DiffLen(t *testing.T) {
	actual := args.Map{"v": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines diff len", actual)
}

func Test_Cov11_HasAnyMismatchOnLines_DiffContent(t *testing.T) {
	actual := args.Map{"v": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"})}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines diff content", actual)
}

func Test_Cov11_PrintLineDiffOnFail_NoFail(t *testing.T) {
	errcore.PrintLineDiffOnFail(0, "h", []string{"a"}, []string{"a"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail no fail", actual)
}

func Test_Cov11_PrintLineDiffOnFail_Fail(t *testing.T) {
	errcore.PrintLineDiffOnFail(0, "h", []string{"a"}, []string{"b"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintLineDiffOnFail fail", actual)
}

func Test_Cov11_ErrorToLinesLineDiff_NilErr(t *testing.T) {
	result := errcore.ErrorToLinesLineDiff(0, "h", nil, []string{"a"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff nil err", actual)
}

func Test_Cov11_ErrorToLinesLineDiff_WithErr(t *testing.T) {
	result := errcore.ErrorToLinesLineDiff(0, "h", errors.New("line1\nline2"), []string{"line1"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff with err", actual)
}

func Test_Cov11_PrintErrorLineDiff(t *testing.T) {
	errcore.PrintErrorLineDiff(0, "h", errors.New("a"), []string{"b"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorLineDiff", actual)
}

func Test_Cov11_SliceDiffSummary_Match(t *testing.T) {
	result := errcore.SliceDiffSummary([]string{"a"}, []string{"a"})
	actual := args.Map{"v": result}
	expected := args.Map{"v": "all lines match"}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary match", actual)
}

func Test_Cov11_SliceDiffSummary_Mismatch(t *testing.T) {
	result := errcore.SliceDiffSummary([]string{"a"}, []string{"b"})
	actual := args.Map{"notEmpty": result != "all lines match"}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapMismatchError
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_MapMismatchError(t *testing.T) {
	result := errcore.MapMismatchError("TestFunc", 1, "title",
		[]string{`"k": "v"`}, []string{`"k": "v2"`})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapMismatchError", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AssertDiffOnMismatch / AssertErrorDiffOnMismatch
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_AssertDiffOnMismatch_Match(t *testing.T) {
	errcore.AssertDiffOnMismatch(t, 0, "t", []string{"a"}, []string{"a"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "AssertDiffOnMismatch match", actual)
}

func Test_Cov11_AssertErrorDiffOnMismatch_NilMatch(t *testing.T) {
	errcore.AssertErrorDiffOnMismatch(t, 0, "t", nil, []string{})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "AssertErrorDiffOnMismatch nil match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PrintDiffOnMismatch
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_PrintDiffOnMismatch_Match(t *testing.T) {
	errcore.PrintDiffOnMismatch(0, "t", []string{"a"}, []string{"a"})
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch match", actual)
}

func Test_Cov11_PrintDiffOnMismatch_Mismatch(t *testing.T) {
	errcore.PrintDiffOnMismatch(0, "t", []string{"a"}, []string{"b"}, "ctx1")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintDiffOnMismatch mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MergeErrors / MergeErrorsToString / MergeErrorsToStringDefault
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_MergeErrors_AllNil(t *testing.T) {
	err := errcore.MergeErrors(nil, nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors all nil", actual)
}

func Test_Cov11_MergeErrors_WithErr(t *testing.T) {
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors with err", actual)
}

func Test_Cov11_MergeErrorsToString_Nil(t *testing.T) {
	actual := args.Map{"v": errcore.MergeErrorsToString(",")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString nil", actual)
}

func Test_Cov11_MergeErrorsToString_WithErr(t *testing.T) {
	result := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString with err", actual)
}

func Test_Cov11_MergeErrorsToStringDefault_Nil(t *testing.T) {
	actual := args.Map{"v": errcore.MergeErrorsToStringDefault()}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault nil", actual)
}

func Test_Cov11_MergeErrorsToStringDefault_WithErr(t *testing.T) {
	result := errcore.MergeErrorsToStringDefault(errors.New("a"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault with err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceError / SliceErrorDefault / SliceErrorsToStrings / SliceToError / SliceToErrorPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_SliceError_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.SliceError(",", []string{}) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError empty", actual)
}

func Test_Cov11_SliceError_NonEmpty(t *testing.T) {
	err := errcore.SliceError(",", []string{"a", "b"})
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError non-empty", actual)
}

func Test_Cov11_SliceErrorDefault_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.SliceErrorDefault([]string{}) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault empty", actual)
}

func Test_Cov11_SliceErrorDefault_NonEmpty(t *testing.T) {
	err := errcore.SliceErrorDefault([]string{"a"})
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault non-empty", actual)
}

func Test_Cov11_SliceErrorsToStrings_Nil(t *testing.T) {
	result := errcore.SliceErrorsToStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings nil", actual)
}

func Test_Cov11_SliceErrorsToStrings_WithNils(t *testing.T) {
	result := errcore.SliceErrorsToStrings(nil, errors.New("a"), nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings with nils", actual)
}

func Test_Cov11_SliceToError_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.SliceToError([]string{}) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError empty", actual)
}

func Test_Cov11_SliceToError_NonEmpty(t *testing.T) {
	err := errcore.SliceToError([]string{"a", "b"})
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError non-empty", actual)
}

func Test_Cov11_SliceToErrorPtr_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.SliceToErrorPtr([]string{}) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr empty", actual)
}

func Test_Cov11_SliceToErrorPtr_NonEmpty(t *testing.T) {
	err := errcore.SliceToErrorPtr([]string{"a"})
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MustBeEmpty
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_MustBeEmpty_Nil(t *testing.T) {
	errcore.MustBeEmpty(nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty nil", actual)
}

func Test_Cov11_MustBeEmpty_Panic(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.MustBeEmpty(errors.New("e"))
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty panic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ErrorToSplitLines / ErrorToSplitNonEmptyLines
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_ErrorToSplitLines_Nil(t *testing.T) {
	result := errcore.ErrorToSplitLines(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines nil", actual)
}

func Test_Cov11_ErrorToSplitLines_Multi(t *testing.T) {
	result := errcore.ErrorToSplitLines(errors.New("a\nb\nc"))
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines multi", actual)
}

func Test_Cov11_ErrorToSplitNonEmptyLines_WithEmpty(t *testing.T) {
	result := errcore.ErrorToSplitNonEmptyLines(errors.New("a\n\nb"))
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines", actual)
}

func Test_Cov11_ErrorToSplitNonEmptyLines_Nil(t *testing.T) {
	result := errcore.ErrorToSplitNonEmptyLines(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Ref / RefToError / ToError / ToString / ToStringPtr / ToValueString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_Ref_Nil(t *testing.T) {
	actual := args.Map{"v": errcore.Ref(nil)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "Ref nil", actual)
}

func Test_Cov11_Ref_NonNil(t *testing.T) {
	actual := args.Map{"notEmpty": errcore.Ref("hello") != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref non-nil", actual)
}

func Test_Cov11_RefToError_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.RefToError(nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError nil", actual)
}

func Test_Cov11_RefToError_NonNil(t *testing.T) {
	err := errcore.RefToError("ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError non-nil", actual)
}

func Test_Cov11_ToError_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.ToError("") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError empty", actual)
}

func Test_Cov11_ToError_NonEmpty(t *testing.T) {
	err := errcore.ToError("e")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToError non-empty", actual)
}

func Test_Cov11_ToString_Nil(t *testing.T) {
	actual := args.Map{"v": errcore.ToString(nil)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ToString nil", actual)
}

func Test_Cov11_ToString_WithErr(t *testing.T) {
	actual := args.Map{"v": errcore.ToString(errors.New("e"))}
	expected := args.Map{"v": "e"}
	expected.ShouldBeEqual(t, 0, "ToString with err", actual)
}

func Test_Cov11_ToStringPtr_Nil(t *testing.T) {
	result := errcore.ToStringPtr(nil)
	actual := args.Map{"empty": *result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ToStringPtr nil", actual)
}

func Test_Cov11_ToStringPtr_WithErr(t *testing.T) {
	result := errcore.ToStringPtr(errors.New("e"))
	actual := args.Map{"v": *result}
	expected := args.Map{"v": "e"}
	expected.ShouldBeEqual(t, 0, "ToStringPtr with err", actual)
}

func Test_Cov11_ToValueString(t *testing.T) {
	result := errcore.ToValueString("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToValueString", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VarMap / VarMapStrings
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_VarMap_Empty(t *testing.T) {
	actual := args.Map{"v": errcore.VarMap(map[string]any{})}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "VarMap empty", actual)
}

func Test_Cov11_VarMap_NonEmpty(t *testing.T) {
	result := errcore.VarMap(map[string]any{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap non-empty", actual)
}

func Test_Cov11_VarMapStrings_Empty(t *testing.T) {
	result := errcore.VarMapStrings(map[string]any{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarMapStrings empty", actual)
}

func Test_Cov11_VarMapStrings_NonEmpty(t *testing.T) {
	result := errcore.VarMapStrings(map[string]any{"k": "v"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "VarMapStrings non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Combine
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_Combine(t *testing.T) {
	result := errcore.Combine("generic", "other", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ConcatMessageWithErr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_ConcatMessageWithErr_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.ConcatMessageWithErr("msg", nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr nil", actual)
}

func Test_Cov11_ConcatMessageWithErr_WithErr(t *testing.T) {
	err := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr with err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RawErrCollection — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_RawErrCollection_Basic(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{
		"isEmpty": rec.IsEmpty(), "isNull": rec.IsNull(), "isAnyNull": rec.IsAnyNull(),
		"length": rec.Length(), "hasError": rec.HasError(), "hasAny": rec.HasAnyError(),
		"isValid": rec.IsValid(), "isSuccess": rec.IsSuccess(), "isFailed": rec.IsFailed(),
		"isInvalid": rec.IsInvalid(), "isDefined": rec.IsDefined(), "hasIssues": rec.HasAnyIssues(),
		"isCollection": rec.IsCollectionType(),
	}
	expected := args.Map{
		"isEmpty": true, "isNull": true, "isAnyNull": true,
		"length": 0, "hasError": false, "hasAny": false,
		"isValid": true, "isSuccess": true, "isFailed": false,
		"isInvalid": false, "isDefined": false, "hasIssues": false,
		"isCollection": true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection basic", actual)
}

func Test_Cov11_RawErrCollection_Add(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(nil)
	rec.Add(errors.New("a"))
	rec.AddError(nil)
	rec.AddError(errors.New("b"))
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RawErrCollection Add", actual)
}

func Test_Cov11_RawErrCollection_AddMsg(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddMsg("hello")
	rec.AddMsg("")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddMsg", actual)
}

func Test_Cov11_RawErrCollection_AddMsgStackTrace(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddMsgStackTrace("")
	rec.AddMsgStackTrace("msg")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddMsgStackTrace", actual)
}

func Test_Cov11_RawErrCollection_AddStackTrace(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddStackTrace(nil)
	rec.AddStackTrace(errors.New("e"))
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddStackTrace", actual)
}

func Test_Cov11_RawErrCollection_AddMsgErrStackTrace(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddMsgErrStackTrace("msg", nil)
	rec.AddMsgErrStackTrace("msg", errors.New("e"))
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddMsgErrStackTrace", actual)
}

func Test_Cov11_RawErrCollection_AddMethodName(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddMethodName("")
	rec.AddMethodName("msg")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddMethodName", actual)
}

func Test_Cov11_RawErrCollection_AddMessages(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddMessages()
	rec.AddMessages("a", "b")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddMessages", actual)
}

func Test_Cov11_RawErrCollection_AddErrorWithMessage(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddErrorWithMessage(nil, "msg")
	rec.AddErrorWithMessage(errors.New("e"), "msg")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddErrorWithMessage", actual)
}

func Test_Cov11_RawErrCollection_AddIf(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddIf(false, "skip")
	rec.AddIf(true, "add")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddIf", actual)
}

func Test_Cov11_RawErrCollection_AddFunc(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddFunc(nil)
	rec.AddFunc(func() error { return nil })
	rec.AddFunc(func() error { return errors.New("e") })
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddFunc", actual)
}

func Test_Cov11_RawErrCollection_AddFuncIf(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddFuncIf(false, func() error { return errors.New("e") })
	rec.AddFuncIf(true, nil)
	rec.AddFuncIf(true, func() error { return errors.New("e") })
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddFuncIf", actual)
}

func Test_Cov11_RawErrCollection_AddErrorWithMessageRef(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddErrorWithMessageRef(nil, "msg", "ref")
	rec.AddErrorWithMessageRef(errors.New("e"), "msg", nil)
	rec.AddErrorWithMessageRef(errors.New("e"), "msg", "ref")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddErrorWithMessageRef", actual)
}

func Test_Cov11_RawErrCollection_AddFmt(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddFmt(nil, "fmt %d", 1)
	rec.AddFmt(errors.New("e"), "fmt %d", 1)
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddFmt", actual)
}

func Test_Cov11_RawErrCollection_Fmt(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Fmt("", )
	rec.Fmt("hello %d", 1)
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Fmt", actual)
}

func Test_Cov11_RawErrCollection_FmtIf(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.FmtIf(false, "skip %d", 1)
	rec.FmtIf(true, "add %d", 1)
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FmtIf", actual)
}

func Test_Cov11_RawErrCollection_References(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.References("msg", "r1")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "References", actual)
}

func Test_Cov11_RawErrCollection_Adds(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Adds()
	rec.Adds(nil, errors.New("a"), nil, errors.New("b"))
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Adds", actual)
}

func Test_Cov11_RawErrCollection_AddErrors(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddErrors(errors.New("a"))
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddErrors", actual)
}

func Test_Cov11_RawErrCollection_ConditionalAddError(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.ConditionalAddError(false, errors.New("skip"))
	rec.ConditionalAddError(true, errors.New("add"))
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ConditionalAddError", actual)
}

func Test_Cov11_RawErrCollection_AddString(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddString("")
	rec.AddString("msg")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddString", actual)
}

func Test_Cov11_RawErrCollection_AddStringSliceAsErr(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddStringSliceAsErr()
	rec.AddStringSliceAsErr("", "a", "", "b")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddStringSliceAsErr", actual)
}

func Test_Cov11_RawErrCollection_AddWithTraceRef(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddWithTraceRef(nil, []string{"t"}, "r")
	rec.AddWithTraceRef(errors.New("e"), []string{"t"}, "r")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddWithTraceRef", actual)
}

func Test_Cov11_RawErrCollection_AddWithCompiledTraceRef(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddWithCompiledTraceRef(nil, "trace", "r")
	rec.AddWithCompiledTraceRef(errors.New("e"), "trace", "r")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddWithCompiledTraceRef", actual)
}

func Test_Cov11_RawErrCollection_AddWithRef(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddWithRef(nil, "r")
	rec.AddWithRef(errors.New("e"), "r")
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddWithRef", actual)
}

func Test_Cov11_RawErrCollection_String(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"empty": rec.String() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "String empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notEmpty": rec.String() != ""}
	expected2 := args.Map{"notEmpty": true}
	expected2.ShouldBeEqual(t, 0, "String non-empty", actual2)
}

func Test_Cov11_RawErrCollection_Strings(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"len": len(rec.Strings())}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Strings empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"len": len(rec.Strings())}
	expected2 := args.Map{"len": 1}
	expected2.ShouldBeEqual(t, 0, "Strings non-empty", actual2)
}

func Test_Cov11_RawErrCollection_StringUsingJoiner(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"empty": rec.StringUsingJoiner(",") == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StringUsingJoiner empty", actual)

	rec.Add(errors.New("a"))
	rec.Add(errors.New("b"))
	actual2 := args.Map{"v": rec.StringUsingJoiner(",")}
	expected2 := args.Map{"v": "a,b"}
	expected2.ShouldBeEqual(t, 0, "StringUsingJoiner non-empty", actual2)
}

func Test_Cov11_RawErrCollection_StringUsingJoinerAdditional(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"empty": rec.StringUsingJoinerAdditional(",", "!") == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StringUsingJoinerAdditional empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"v": rec.StringUsingJoinerAdditional(",", "!")}
	expected2 := args.Map{"v": "a!"}
	expected2.ShouldBeEqual(t, 0, "StringUsingJoinerAdditional non-empty", actual2)
}

func Test_Cov11_RawErrCollection_CompiledError(t *testing.T) {
	rec := errcore.RawErrCollection{}
	actual := args.Map{"isNil": rec.CompiledError() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledError() != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "CompiledError non-empty", actual2)
}

func Test_Cov11_RawErrCollection_CompiledErrorUsingJoiner(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"isNil": rec.CompiledErrorUsingJoiner(",") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledErrorUsingJoiner empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorUsingJoiner(",") != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "CompiledErrorUsingJoiner non-empty", actual2)
}

func Test_Cov11_RawErrCollection_CompiledErrorUsingJoinerAdditionalMessage(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"isNil": rec.CompiledErrorUsingJoinerAdditionalMessage(",", "!") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CEUJAM empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorUsingJoinerAdditionalMessage(",", "!") != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "CEUJAM non-empty", actual2)
}

func Test_Cov11_RawErrCollection_CompiledErrorUsingStackTraces(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"isNil": rec.CompiledErrorUsingStackTraces(",", []string{"t"}) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CEUST empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorUsingStackTraces(",", []string{"t"}) != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "CEUST non-empty", actual2)
}

func Test_Cov11_RawErrCollection_StringWithAdditionalMessage(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"empty": rec.StringWithAdditionalMessage("!") == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SWAM empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"v": rec.StringWithAdditionalMessage("!")}
	expected2 := args.Map{"v": "a!"}
	expected2.ShouldBeEqual(t, 0, "SWAM non-empty", actual2)
}

func Test_Cov11_RawErrCollection_CompiledErrorWithStackTraces(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"isNil": rec.CompiledErrorWithStackTraces() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "CEWST empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notNil": rec.CompiledErrorWithStackTraces() != nil}
	expected2 := args.Map{"notNil": true}
	expected2.ShouldBeEqual(t, 0, "CEWST non-empty", actual2)
}

func Test_Cov11_RawErrCollection_CompiledStackTracesString(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"empty": rec.CompiledStackTracesString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "CSTS empty", actual)

	rec.Add(errors.New("a"))
	actual2 := args.Map{"notEmpty": rec.CompiledStackTracesString() != ""}
	expected2 := args.Map{"notEmpty": true}
	expected2.ShouldBeEqual(t, 0, "CSTS non-empty", actual2)
}

func Test_Cov11_RawErrCollection_CompiledJsonErrorWithStackTraces(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	err := rec.CompiledJsonErrorWithStackTraces()
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CJEWST", actual)
}

func Test_Cov11_RawErrCollection_CompiledJsonStringWithStackTraces(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	result := rec.CompiledJsonStringWithStackTraces()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CJSWST", actual)
}

func Test_Cov11_RawErrCollection_FullString(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	actual := args.Map{"v": rec.FullString() != ""}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "FullString", actual)
}

func Test_Cov11_RawErrCollection_FullStringWithTraces(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	actual := args.Map{"v": rec.FullStringWithTraces() != ""}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "FullStringWithTraces", actual)
}

func Test_Cov11_RawErrCollection_FullStringWithTracesIf(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	r1 := rec.FullStringWithTracesIf(true)
	r2 := rec.FullStringWithTracesIf(false)
	actual := args.Map{"t": r1 != "", "f": r2 != ""}
	expected := args.Map{"t": true, "f": true}
	expected.ShouldBeEqual(t, 0, "FullStringWithTracesIf", actual)
}

func Test_Cov11_RawErrCollection_ReferencesCompiledString(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	actual := args.Map{"v": rec.ReferencesCompiledString() != ""}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "ReferencesCompiledString", actual)
}

func Test_Cov11_RawErrCollection_FullStringSplitByNewLine(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	actual := args.Map{"len": len(rec.FullStringSplitByNewLine())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FullStringSplitByNewLine", actual)
}

func Test_Cov11_RawErrCollection_FullStringWithoutReferences(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	actual := args.Map{"v": rec.FullStringWithoutReferences() != ""}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "FullStringWithoutReferences", actual)
}

func Test_Cov11_RawErrCollection_ErrorString(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	actual := args.Map{"v": rec.ErrorString() != ""}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "ErrorString", actual)
}

func Test_Cov11_RawErrCollection_Compile(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	actual := args.Map{"v": rec.Compile() != ""}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "Compile", actual)
}

func Test_Cov11_RawErrCollection_Value(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"isNil": rec.Value() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Value empty", actual)
}

func Test_Cov11_RawErrCollection_Serialize(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	b, err := rec.Serialize()
	actual := args.Map{"nil": b == nil, "noErr": err == nil}
	expected := args.Map{"nil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize empty", actual)
}

func Test_Cov11_RawErrCollection_SerializeWithoutTraces(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	b, err := rec.SerializeWithoutTraces()
	actual := args.Map{"nil": b == nil, "noErr": err == nil}
	expected := args.Map{"nil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "SerializeWithoutTraces empty", actual)
}

func Test_Cov11_RawErrCollection_SerializeMust(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	b := rec.SerializeMust()
	actual := args.Map{"nil": b == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SerializeMust empty", actual)
}

func Test_Cov11_RawErrCollection_MarshalJSON(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	b, err := rec.MarshalJSON()
	actual := args.Map{"nil": b == nil, "noErr": err == nil}
	expected := args.Map{"nil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON empty", actual)
}

func Test_Cov11_RawErrCollection_UnmarshalJSON(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	err := rec.UnmarshalJSON([]byte(`[]`))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON", actual)
}

func Test_Cov11_RawErrCollection_Log(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Log() // empty, no-op
	rec.Add(errors.New("a"))
	rec.Log()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Log", actual)
}

func Test_Cov11_RawErrCollection_LogWithTraces(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.LogWithTraces() // empty, no-op
	rec.Add(errors.New("a"))
	rec.LogWithTraces()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LogWithTraces", actual)
}

func Test_Cov11_RawErrCollection_LogIf_False(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.LogIf(false)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LogIf false", actual)
}

func Test_Cov11_RawErrCollection_Clear(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Clear() // empty
	rec.Add(errors.New("a"))
	rec.Clear()
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear", actual)
}

func Test_Cov11_RawErrCollection_Dispose(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Dispose() // empty
	rec.Add(errors.New("a"))
	rec.Dispose()
	actual := args.Map{"isNull": rec.IsNull()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "Dispose", actual)
}

func Test_Cov11_RawErrCollection_MustBeSafe_Empty(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.MustBeSafe()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeSafe empty", actual)
}

func Test_Cov11_RawErrCollection_MustBeSafe_Panic(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.MustBeSafe()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MustBeSafe panic", actual)
}

func Test_Cov11_RawErrCollection_MustBeEmptyError(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.MustBeEmptyError() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmptyError", actual)
}

func Test_Cov11_RawErrCollection_HandleError_Empty(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.HandleError() // empty, no-op
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleError empty", actual)
}

func Test_Cov11_RawErrCollection_HandleError_Panic(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.HandleError()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleError panic", actual)
}

func Test_Cov11_RawErrCollection_HandleErrorWithRefs_Empty(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.HandleErrorWithRefs("msg", "k", "v") // empty, no-op
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorWithRefs empty", actual)
}

func Test_Cov11_RawErrCollection_HandleErrorWithRefs_Panic(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.HandleErrorWithRefs("msg", "k", "v")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorWithRefs panic", actual)
}

func Test_Cov11_RawErrCollection_HandleErrorWithMsg_Empty(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.HandleErrorWithMsg("msg") // empty, no-op
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorWithMsg empty", actual)
}

func Test_Cov11_RawErrCollection_HandleErrorWithMsg_Panic(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("e"))
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		rec.HandleErrorWithMsg("msg")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorWithMsg panic", actual)
}

func Test_Cov11_RawErrCollection_ReflectSetTo_Value(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	err := rec.ReflectSetTo(errcore.RawErrCollection{})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo value", actual)
}

func Test_Cov11_RawErrCollection_ReflectSetTo_NilPtr(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	var nilPtr *errcore.RawErrCollection
	err := rec.ReflectSetTo(nilPtr)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo nil ptr", actual)
}

func Test_Cov11_RawErrCollection_ReflectSetTo_ValidPtr(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	target := &errcore.RawErrCollection{}
	err := rec.ReflectSetTo(target)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo valid ptr", actual)
}

func Test_Cov11_RawErrCollection_ReflectSetTo_Other(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	err := rec.ReflectSetTo("unsupported")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo other", actual)
}

func Test_Cov11_RawErrCollection_CountStateChangeTracker(t *testing.T) {
	rec := errcore.RawErrCollection{}
	tracker := rec.CountStateChangeTracker()
	actual := args.Map{"same": tracker.IsSameState()}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker", actual)
}

func Test_Cov11_RawErrCollection_IsErrorsCollected_NoNew(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"v": rec.IsErrorsCollected(nil)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsErrorsCollected no new", actual)
}

func Test_Cov11_RawErrCollection_IsErrorsCollected_WithNew(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	actual := args.Map{"v": rec.IsErrorsCollected(errors.New("e"))}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsErrorsCollected with new", actual)
}

func Test_Cov11_RawErrCollection_ToRawErrCollection(t *testing.T) {
	rec := errcore.RawErrCollection{}
	ptr := rec.ToRawErrCollection()
	actual := args.Map{"notNil": ptr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToRawErrCollection", actual)
}

func Test_Cov11_RawErrCollection_AddErrorGetters(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddErrorGetters()
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddErrorGetters empty", actual)
}

func Test_Cov11_RawErrCollection_AddCompiledErrorGetters(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.AddCompiledErrorGetters()
	actual := args.Map{"len": rec.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddCompiledErrorGetters empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RawErrorType — String / Combine
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov11_RawErrorType_String(t *testing.T) {
	actual := args.Map{"v": errcore.InvalidType.String() != ""}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType String", actual)
}

func Test_Cov11_RawErrorType_Combine(t *testing.T) {
	result := errcore.InvalidType.Combine("msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType Combine", actual)
}

func Test_Cov11_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	err := errcore.InvalidType.ErrorNoRefsSkip(0, "msg")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorNoRefsSkip", actual)
}

func Test_Cov11_RawErrorType_ErrorNoRefsSkip_Empty(t *testing.T) {
	err := errcore.InvalidType.ErrorNoRefsSkip(0, "")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorNoRefsSkip empty", actual)
}
