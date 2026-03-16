package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ── ToError ──

func Test_Cov9_ToError_NonEmpty(t *testing.T) {
	err := errcore.ToError("some error")
	actual := args.Map{"notNil": err != nil, "msg": err.Error()}
	expected := args.Map{"notNil": true, "msg": "some error"}
	expected.ShouldBeEqual(t, 0, "ToError non-empty", actual)
}

func Test_Cov9_ToError_Empty(t *testing.T) {
	err := errcore.ToError("")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError empty", actual)
}

// ── ToString ──

func Test_Cov9_ToString_Nil(t *testing.T) {
	result := errcore.ToString(nil)
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToString nil", actual)
}

func Test_Cov9_ToString_WithError(t *testing.T) {
	result := errcore.ToString(errors.New("hello"))
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ToString with error", actual)
}

// ── ToStringPtr ──

func Test_Cov9_ToStringPtr_Nil(t *testing.T) {
	result := errcore.ToStringPtr(nil)
	actual := args.Map{"notNil": result != nil, "val": *result}
	expected := args.Map{"notNil": true, "val": ""}
	expected.ShouldBeEqual(t, 0, "ToStringPtr nil", actual)
}

func Test_Cov9_ToStringPtr_WithError(t *testing.T) {
	result := errcore.ToStringPtr(errors.New("err msg"))
	actual := args.Map{"val": *result}
	expected := args.Map{"val": "err msg"}
	expected.ShouldBeEqual(t, 0, "ToStringPtr with error", actual)
}

// ── ToValueString ──

func Test_Cov9_ToValueString(t *testing.T) {
	result := errcore.ToValueString("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToValueString", actual)
}

// ── Ref ──

func Test_Cov9_Ref_Nil(t *testing.T) {
	result := errcore.Ref(nil)
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Ref nil", actual)
}

func Test_Cov9_Ref_WithValue(t *testing.T) {
	result := errcore.Ref("some-ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref with value", actual)
}

// ── RefToError ──

func Test_Cov9_RefToError_Nil(t *testing.T) {
	err := errcore.RefToError(nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError nil", actual)
}

func Test_Cov9_RefToError_WithValue(t *testing.T) {
	err := errcore.RefToError("ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RefToError with value", actual)
}

// ── SliceError ──

func Test_Cov9_SliceError_Empty(t *testing.T) {
	err := errcore.SliceError(",", []string{})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError empty", actual)
}

func Test_Cov9_SliceError_NonEmpty(t *testing.T) {
	err := errcore.SliceError(",", []string{"a", "b"})
	actual := args.Map{"msg": err.Error()}
	expected := args.Map{"msg": "a,b"}
	expected.ShouldBeEqual(t, 0, "SliceError non-empty", actual)
}

// ── SliceErrorDefault ──

func Test_Cov9_SliceErrorDefault_Empty(t *testing.T) {
	err := errcore.SliceErrorDefault([]string{})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault empty", actual)
}

// ── SliceToError ──

func Test_Cov9_SliceToError_Empty(t *testing.T) {
	err := errcore.SliceToError([]string{})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError empty", actual)
}

func Test_Cov9_SliceToError_NonEmpty(t *testing.T) {
	err := errcore.SliceToError([]string{"x", "y"})
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError non-empty", actual)
}

// ── SliceToErrorPtr ──

func Test_Cov9_SliceToErrorPtr_Empty(t *testing.T) {
	err := errcore.SliceToErrorPtr([]string{})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr empty", actual)
}

// ── SliceErrorsToStrings ──

func Test_Cov9_SliceErrorsToStrings_Nil(t *testing.T) {
	result := errcore.SliceErrorsToStrings(nil...)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings nil", actual)
}

func Test_Cov9_SliceErrorsToStrings_WithNils(t *testing.T) {
	result := errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b"))
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings with nils", actual)
}

// ── MergeErrors ──

func Test_Cov9_MergeErrors_AllNil(t *testing.T) {
	err := errcore.MergeErrors(nil, nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors all nil", actual)
}

func Test_Cov9_MergeErrors_Mixed(t *testing.T) {
	err := errcore.MergeErrors(errors.New("e1"), nil, errors.New("e2"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors mixed", actual)
}

// ── ManyErrorToSingle ──

func Test_Cov9_ManyErrorToSingle_Empty(t *testing.T) {
	err := errcore.ManyErrorToSingle([]error{})
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle empty", actual)
}

// ── ManyErrorToSingleDirect ──

func Test_Cov9_ManyErrorToSingleDirect_Empty(t *testing.T) {
	err := errcore.ManyErrorToSingleDirect()
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect empty", actual)
}

// ── MergeErrorsToString ──

func Test_Cov9_MergeErrorsToString_Nil(t *testing.T) {
	result := errcore.MergeErrorsToString(",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString nil", actual)
}

func Test_Cov9_MergeErrorsToString_WithErrors(t *testing.T) {
	result := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString with errors", actual)
}

// ── MergeErrorsToStringDefault ──

func Test_Cov9_MergeErrorsToStringDefault_Nil(t *testing.T) {
	result := errcore.MergeErrorsToStringDefault()
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault nil", actual)
}

// ── ErrorToSplitLines ──

func Test_Cov9_ErrorToSplitLines_Nil(t *testing.T) {
	result := errcore.ErrorToSplitLines(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines nil", actual)
}

func Test_Cov9_ErrorToSplitLines_WithLines(t *testing.T) {
	result := errcore.ErrorToSplitLines(errors.New("line1\nline2"))
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "line1", "second": "line2"}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitLines with lines", actual)
}

// ── ErrorToSplitNonEmptyLines ──

func Test_Cov9_ErrorToSplitNonEmptyLines_Nil(t *testing.T) {
	result := errcore.ErrorToSplitNonEmptyLines(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ErrorToSplitNonEmptyLines nil", actual)
}

// ── VarTwo ──

func Test_Cov9_VarTwo_NoType(t *testing.T) {
	result := errcore.VarTwo(false, "a", 1, "b", 2)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(a, b) = (1, 2)"}
	expected.ShouldBeEqual(t, 0, "VarTwo no type", actual)
}

func Test_Cov9_VarTwo_WithType(t *testing.T) {
	result := errcore.VarTwo(true, "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo with type", actual)
}

// ── VarTwoNoType ──

func Test_Cov9_VarTwoNoType(t *testing.T) {
	result := errcore.VarTwoNoType("x", 10, "y", 20)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(x, y) = (10, 20)"}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType", actual)
}

// ── VarThree ──

func Test_Cov9_VarThree_NoType(t *testing.T) {
	result := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(a, b, c) = (1, 2, 3)"}
	expected.ShouldBeEqual(t, 0, "VarThree no type", actual)
}

func Test_Cov9_VarThree_WithType(t *testing.T) {
	result := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree with type", actual)
}

// ── VarThreeNoType ──

func Test_Cov9_VarThreeNoType(t *testing.T) {
	result := errcore.VarThreeNoType("x", 1, "y", 2, "z", 3)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "(x, y, z) = (1, 2, 3)"}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType", actual)
}

// ── MessageVarTwo ──

func Test_Cov9_MessageVarTwo(t *testing.T) {
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo", actual)
}

// ── MessageVarThree ──

func Test_Cov9_MessageVarThree(t *testing.T) {
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree", actual)
}

// ── MessageWithRef ──

func Test_Cov9_MessageWithRef(t *testing.T) {
	result := errcore.MessageWithRef("msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef", actual)
}

// ── MessageWithRefToError ──

func Test_Cov9_MessageWithRefToError(t *testing.T) {
	err := errcore.MessageWithRefToError("msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError", actual)
}

// ── ErrorWithRef ──

func Test_Cov9_ErrorWithRef_NilErr(t *testing.T) {
	result := errcore.ErrorWithRef(nil, "ref")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef nil err", actual)
}

func Test_Cov9_ErrorWithRef_NilRef(t *testing.T) {
	result := errcore.ErrorWithRef(errors.New("e"), nil)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "e"}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef nil ref", actual)
}

func Test_Cov9_ErrorWithRef_EmptyRef(t *testing.T) {
	result := errcore.ErrorWithRef(errors.New("e"), "")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "e"}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef empty ref", actual)
}

func Test_Cov9_ErrorWithRef_Both(t *testing.T) {
	result := errcore.ErrorWithRef(errors.New("e"), "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef both", actual)
}

// ── Combine ──

func Test_Cov9_Combine(t *testing.T) {
	result := errcore.Combine("generic", "other", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine", actual)
}

// ── ConcatMessageWithErr ──

func Test_Cov9_ConcatMessageWithErr_NilErr(t *testing.T) {
	err := errcore.ConcatMessageWithErr("msg", nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr nil", actual)
}

func Test_Cov9_ConcatMessageWithErr_WithErr(t *testing.T) {
	err := errcore.ConcatMessageWithErr("prefix", errors.New("inner"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ConcatMessageWithErr with err", actual)
}

// ── MustBeEmpty ──

func Test_Cov9_MustBeEmpty_Nil(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.MustBeEmpty(nil)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": false}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty nil -- no panic", actual)
}

func Test_Cov9_MustBeEmpty_WithError(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "MustBeEmpty with error -- panic", actual)
}

// ── VarMap ──

func Test_Cov9_VarMap_Empty(t *testing.T) {
	result := errcore.VarMap(map[string]any{})
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarMap empty", actual)
}

func Test_Cov9_VarMap_NonEmpty(t *testing.T) {
	result := errcore.VarMap(map[string]any{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap non-empty", actual)
}

// ── VarMapStrings ──

func Test_Cov9_VarMapStrings_Empty(t *testing.T) {
	result := errcore.VarMapStrings(map[string]any{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "VarMapStrings empty", actual)
}

// ── MessageVarMap ──

func Test_Cov9_MessageVarMap_Empty(t *testing.T) {
	result := errcore.MessageVarMap("msg", map[string]any{})
	actual := args.Map{"result": result}
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageVarMap empty", actual)
}

func Test_Cov9_MessageVarMap_NonEmpty(t *testing.T) {
	result := errcore.MessageVarMap("msg", map[string]any{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap non-empty", actual)
}

// ── Expecting ──

func Test_Cov9_Expecting(t *testing.T) {
	result := errcore.Expecting("title", "expected", "actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting", actual)
}

// ── ExpectingSimple ──

func Test_Cov9_ExpectingSimple(t *testing.T) {
	result := errcore.ExpectingSimple("title", "expected", "actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple", actual)
}

// ── ExpectingSimpleNoType ──

func Test_Cov9_ExpectingSimpleNoType(t *testing.T) {
	result := errcore.ExpectingSimpleNoType("title", "expected", "actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType", actual)
}

// ── ExpectingErrorSimpleNoType ──

func Test_Cov9_ExpectingErrorSimpleNoType(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("title", "expected", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType", actual)
}

// ── ExpectingErrorSimpleNoTypeNewLineEnds ──

func Test_Cov9_ExpectingErrorSimpleNoTypeNewLineEnds(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoTypeNewLineEnds("title", "expected", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoTypeNewLineEnds", actual)
}

// ── ShouldBe ──

func Test_Cov9_ShouldBe_StrEqMsg(t *testing.T) {
	result := errcore.ShouldBe.StrEqMsg("actual", "expected")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg", actual)
}

func Test_Cov9_ShouldBe_StrEqErr(t *testing.T) {
	err := errcore.ShouldBe.StrEqErr("actual", "expected")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr", actual)
}

func Test_Cov9_ShouldBe_AnyEqMsg(t *testing.T) {
	result := errcore.ShouldBe.AnyEqMsg(1, 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg", actual)
}

func Test_Cov9_ShouldBe_AnyEqErr(t *testing.T) {
	err := errcore.ShouldBe.AnyEqErr(1, 2)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr", actual)
}

func Test_Cov9_ShouldBe_JsonEqMsg(t *testing.T) {
	result := errcore.ShouldBe.JsonEqMsg("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg", actual)
}

func Test_Cov9_ShouldBe_JsonEqErr(t *testing.T) {
	err := errcore.ShouldBe.JsonEqErr("a", "b")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr", actual)
}

// ── Expected ──

func Test_Cov9_Expected_But(t *testing.T) {
	err := errcore.Expected.But("title", "exp", "act")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.But", actual)
}

func Test_Cov9_Expected_ButFoundAsMsg(t *testing.T) {
	result := errcore.Expected.ButFoundAsMsg("title", "exp", "act")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg", actual)
}

func Test_Cov9_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	result := errcore.Expected.ButFoundWithTypeAsMsg("title", "exp", "act")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg", actual)
}

func Test_Cov9_Expected_ButUsingType(t *testing.T) {
	err := errcore.Expected.ButUsingType("title", "exp", "act")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType", actual)
}

// ── RawErrorType ──

func Test_Cov9_RawErrorType_String(t *testing.T) {
	result := errcore.InvalidRequestType.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.String", actual)
}

func Test_Cov9_RawErrorType_Combine(t *testing.T) {
	result := errcore.InvalidRequestType.Combine("other", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Combine", actual)
}

func Test_Cov9_RawErrorType_CombineWithAnother(t *testing.T) {
	result := errcore.InvalidRequestType.CombineWithAnother(errcore.NotFound, "other", "ref")
	actual := args.Map{"notEmpty": string(result) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.CombineWithAnother", actual)
}

func Test_Cov9_RawErrorType_MergeError_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeError(nil)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError nil", actual)
}

func Test_Cov9_RawErrorType_MergeError_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeError(errors.New("inner"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError with err", actual)
}

func Test_Cov9_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessage(nil, "msg")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage nil", actual)
}

func Test_Cov9_RawErrorType_MergeErrorWithMessage_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessage(errors.New("inner"), "msg")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage with err", actual)
}

func Test_Cov9_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithRef(nil, "ref")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef nil", actual)
}

func Test_Cov9_RawErrorType_MergeErrorWithRef_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithRef(errors.New("inner"), "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef with err", actual)
}

func Test_Cov9_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(nil, "msg", "ref")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef nil", actual)
}

func Test_Cov9_RawErrorType_MergeErrorWithMessageRef_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(errors.New("inner"), "msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef with err", actual)
}

func Test_Cov9_RawErrorType_FmtIf_False(t *testing.T) {
	err := errcore.InvalidRequestType.FmtIf(false, "x=%d", 1)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf false", actual)
}

func Test_Cov9_RawErrorType_FmtIf_True(t *testing.T) {
	err := errcore.InvalidRequestType.FmtIf(true, "x=%d", 1)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf true", actual)
}

func Test_Cov9_RawErrorType_SrcDestination(t *testing.T) {
	result := errcore.InvalidRequestType.SrcDestination("msg", "src", "sv", "dst", "dv")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestination", actual)
}

func Test_Cov9_RawErrorType_SrcDestinationErr(t *testing.T) {
	err := errcore.InvalidRequestType.SrcDestinationErr("msg", "src", "sv", "dst", "dv")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestinationErr", actual)
}

func Test_Cov9_RawErrorType_TypesAttach(t *testing.T) {
	result := errcore.InvalidRequestType.TypesAttach("msg", "type1")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttach", actual)
}

func Test_Cov9_RawErrorType_TypesAttachErr(t *testing.T) {
	err := errcore.InvalidRequestType.TypesAttachErr("msg", "type1")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttachErr", actual)
}

func Test_Cov9_GetSet_True(t *testing.T) {
	result := errcore.GetSet(true, errcore.InvalidRequestType, errcore.NotFound)
	actual := args.Map{"result": string(result)}
	expected := args.Map{"result": string(errcore.InvalidRequestType)}
	expected.ShouldBeEqual(t, 0, "GetSet true", actual)
}

func Test_Cov9_GetSet_False(t *testing.T) {
	result := errcore.GetSet(false, errcore.InvalidRequestType, errcore.NotFound)
	actual := args.Map{"result": string(result)}
	expected := args.Map{"result": string(errcore.NotFound)}
	expected.ShouldBeEqual(t, 0, "GetSet false", actual)
}

func Test_Cov9_GetSetVariant_True(t *testing.T) {
	result := errcore.GetSetVariant(true, "yes", "no")
	actual := args.Map{"result": string(result)}
	expected := args.Map{"result": "yes"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant true", actual)
}

func Test_Cov9_GetSetVariant_False(t *testing.T) {
	result := errcore.GetSetVariant(false, "yes", "no")
	actual := args.Map{"result": string(result)}
	expected := args.Map{"result": "no"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant false", actual)
}

// ── CombineWithMsgTypeNoStack ──

func Test_Cov9_CombineWithMsgTypeNoStack_EmptyOther(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack empty other", actual)
}

func Test_Cov9_CombineWithMsgTypeNoStack_WithOther(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidRequestType, "other", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack with other", actual)
}

// ── WasExpectingErrorF ──

func Test_Cov9_WasExpectingErrorF(t *testing.T) {
	err := errcore.WasExpectingErrorF("exp", "act", "title: %d", 1)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "WasExpectingErrorF", actual)
}
