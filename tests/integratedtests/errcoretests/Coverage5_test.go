package errcoretests

import (
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── SliceToError / SliceToErrorPtr ──

func Test_Cov5_SliceToError_Empty(t *testing.T) {
	err := errcore.SliceToError(nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError nil -- nil", actual)
}

func Test_Cov5_SliceToError_NonEmpty(t *testing.T) {
	err := errcore.SliceToError([]string{"err1", "err2"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToError non-empty -- error", actual)
}

func Test_Cov5_SliceToErrorPtr_Nil(t *testing.T) {
	err := errcore.SliceToErrorPtr(nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr nil -- nil", actual)
}

func Test_Cov5_SliceToErrorPtr_NonEmpty(t *testing.T) {
	s := []string{"e1"}
	err := errcore.SliceToErrorPtr(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr non-empty -- error", actual)
}

// ── MergeErrors / MergeErrorsToString ──

func Test_Cov5_MergeErrors_BothNil(t *testing.T) {
	actual := args.Map{"isNil": errcore.MergeErrors(nil, nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors both nil -- nil", actual)
}

func Test_Cov5_MergeErrors_OneNil(t *testing.T) {
	e := errors.New("e")
	actual := args.Map{
		"first":  errcore.MergeErrors(e, nil) != nil,
		"second": errcore.MergeErrors(nil, e) != nil,
	}
	expected := args.Map{"first": true, "second": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors one nil -- non-nil", actual)
}

func Test_Cov5_MergeErrors_Both(t *testing.T) {
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors both -- merged", actual)
}

func Test_Cov5_MergeErrorsToString_Nil(t *testing.T) {
	actual := args.Map{"result": errcore.MergeErrorsToString(nil, nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString nil -- empty", actual)
}

func Test_Cov5_MergeErrorsToString_NonNil(t *testing.T) {
	result := errcore.MergeErrorsToString(errors.New("a"), errors.New("b"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString non-nil -- merged", actual)
}

func Test_Cov5_MergeErrorsToStringDefault_Nil(t *testing.T) {
	actual := args.Map{"result": errcore.MergeErrorsToStringDefault(nil, nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault nil -- empty", actual)
}

// ── ConcatMessageWithErr ──

func Test_Cov5_ConcatMessageWithErr_NilErr(t *testing.T) {
	err := errcore.ConcatMessageWithErr("msg", nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr nil err -- nil", actual)
}

func Test_Cov5_ConcatMessageWithErr_WithErr(t *testing.T) {
	err := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))
	actual := args.Map{"hasErr": err != nil, "containsPrefix": strings.Contains(err.Error(), "prefix")}
	expected := args.Map{"hasErr": true, "containsPrefix": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr with err -- prefixed", actual)
}

// ── ManyErrorToSingle ──

func Test_Cov5_ManyErrorToSingle_Empty(t *testing.T) {
	err := errcore.ManyErrorToSingle([]error{})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle empty -- nil", actual)
}

func Test_Cov5_ManyErrorToSingle_AllNil(t *testing.T) {
	err := errcore.ManyErrorToSingle([]error{nil, nil})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle all nil -- nil", actual)
}

func Test_Cov5_ManyErrorToSingle_WithErrors(t *testing.T) {
	err := errcore.ManyErrorToSingle([]error{errors.New("a"), errors.New("b")})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle with errors -- joined", actual)
}

// ── ToString / ToStringPtr ──

func Test_Cov5_ToString_Nil(t *testing.T) {
	actual := args.Map{"result": errcore.ToString(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToString nil -- empty", actual)
}

func Test_Cov5_ToString_NonNil(t *testing.T) {
	actual := args.Map{"result": errcore.ToString(errors.New("err"))}
	expected := args.Map{"result": "err"}
	expected.ShouldBeEqual(t, 0, "ToString non-nil -- err", actual)
}

func Test_Cov5_ToStringPtr_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.ToStringPtr(nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToStringPtr nil -- nil", actual)
}

func Test_Cov5_ToStringPtr_NonNil(t *testing.T) {
	result := errcore.ToStringPtr(errors.New("err"))
	actual := args.Map{"notNil": result != nil, "val": *result}
	expected := args.Map{"notNil": true, "val": "err"}
	expected.ShouldBeEqual(t, 0, "ToStringPtr non-nil -- err", actual)
}

// ── ToError ──

func Test_Cov5_ToError_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.ToError("") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError empty -- nil", actual)
}

func Test_Cov5_ToError_NonEmpty(t *testing.T) {
	err := errcore.ToError("msg")
	actual := args.Map{"hasErr": err != nil, "msg": err.Error()}
	expected := args.Map{"hasErr": true, "msg": "msg"}
	expected.ShouldBeEqual(t, 0, "ToError non-empty -- msg", actual)
}

// ── LineDiff ──

func Test_Cov5_LineDiff_Same(t *testing.T) {
	result := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})
	actual := args.Map{"isEmpty": result == ""}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiff same -- empty", actual)
}

func Test_Cov5_LineDiff_Different(t *testing.T) {
	result := errcore.LineDiff([]string{"a"}, []string{"b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiff different -- has diff", actual)
}

// ── GherkinsString ──

func Test_Cov5_GherkinsString(t *testing.T) {
	result := errcore.GherkinsString(0, "feature", "given", "when", "then")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString -- formatted", actual)
}

// ── StringLinesToQuoteLines ──

func Test_Cov5_StringLinesToQuoteLines_Empty(t *testing.T) {
	result := errcore.StringLinesToQuoteLines(nil)
	actual := args.Map{"isEmpty": result == ""}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines nil -- empty", actual)
}

func Test_Cov5_StringLinesToQuoteLines_NonEmpty(t *testing.T) {
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines non-empty -- formatted", actual)
}

// ── Ref / RefToError ──

func Test_Cov5_Ref(t *testing.T) {
	result := errcore.Ref("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref -- formatted", actual)
}

func Test_Cov5_RefToError(t *testing.T) {
	err := errcore.RefToError("a", "b")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RefToError -- error", actual)
}

// ── HandleErr ──

func Test_Cov5_HandleErr_Nil(t *testing.T) {
	// Should not panic
	errcore.HandleErr(nil)
	actual := args.Map{"passed": true}
	expected := args.Map{"passed": true}
	expected.ShouldBeEqual(t, 0, "HandleErr nil -- no panic", actual)
}

// ── MustBeEmpty ──

func Test_Cov5_MustBeEmpty_Empty(t *testing.T) {
	err := errcore.MustBeEmpty("test", "")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty empty -- nil", actual)
}

func Test_Cov5_MustBeEmpty_NonEmpty(t *testing.T) {
	err := errcore.MustBeEmpty("test", "value")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty non-empty -- error", actual)
}
