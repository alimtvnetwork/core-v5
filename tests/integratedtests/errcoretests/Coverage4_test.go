package errcoretests

import (
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ── RawErrorType methods ──

func Test_Cov4_RawErrorType_String(t *testing.T) {
	actual := args.Map{"result": errcore.InvalidRequestType.String()}
	expected := args.Map{"result": "Invalid : request, cannot process it."}
	expected.ShouldBeEqual(t, 0, "RawErrorType.String", actual)
}

func Test_Cov4_RawErrorType_Combine(t *testing.T) {
	result := errcore.InvalidRequestType.Combine("details", "ref")
	actual := args.Map{"notEmpty": result != "", "containsType": strings.Contains(result, "Invalid")}
	expected := args.Map{"notEmpty": true, "containsType": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Combine", actual)
}

func Test_Cov4_RawErrorType_CombineWithAnother(t *testing.T) {
	result := errcore.InvalidRequestType.CombineWithAnother(errcore.NotFound, "msg", "ref")
	actual := args.Map{"notEmpty": string(result) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.CombineWithAnother", actual)
}

func Test_Cov4_RawErrorType_TypesAttach(t *testing.T) {
	result := errcore.InvalidRequestType.TypesAttach("msg", "hello", 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttach", actual)
}

func Test_Cov4_RawErrorType_TypesAttachErr(t *testing.T) {
	err := errcore.InvalidRequestType.TypesAttachErr("msg", "hello")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttachErr", actual)
}

func Test_Cov4_RawErrorType_SrcDestination(t *testing.T) {
	result := errcore.InvalidRequestType.SrcDestination("msg", "src", "srcVal", "dst", "dstVal")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestination", actual)
}

func Test_Cov4_RawErrorType_SrcDestinationErr(t *testing.T) {
	err := errcore.InvalidRequestType.SrcDestinationErr("msg", "src", "srcVal", "dst", "dstVal")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestinationErr", actual)
}

func Test_Cov4_RawErrorType_Error(t *testing.T) {
	err := errcore.InvalidRequestType.Error("details", "ref")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Error", actual)
}

func Test_Cov4_RawErrorType_ErrorSkip(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorSkip(0, "details", "ref")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorSkip", actual)
}

func Test_Cov4_RawErrorType_Fmt(t *testing.T) {
	err := errcore.InvalidRequestType.Fmt("value: %d", 42)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Fmt", actual)
}

func Test_Cov4_RawErrorType_Fmt_Empty(t *testing.T) {
	err := errcore.InvalidRequestType.Fmt("")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Fmt empty", actual)
}

func Test_Cov4_RawErrorType_FmtIf_True(t *testing.T) {
	err := errcore.InvalidRequestType.FmtIf(true, "val: %d", 1)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf true", actual)
}

func Test_Cov4_RawErrorType_FmtIf_False(t *testing.T) {
	err := errcore.InvalidRequestType.FmtIf(false, "val: %d", 1)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf false", actual)
}

func Test_Cov4_RawErrorType_MergeError_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeError(nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError nil", actual)
}

func Test_Cov4_RawErrorType_MergeError_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeError(errors.New("inner"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError with err", actual)
}

func Test_Cov4_RawErrorType_MergeErrorWithMessage_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithMessage(nil, "msg") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage nil", actual)
}

func Test_Cov4_RawErrorType_MergeErrorWithMessage_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessage(errors.New("inner"), "msg")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage with err", actual)
}

func Test_Cov4_RawErrorType_MergeErrorWithMessageRef_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithMessageRef(nil, "msg", "ref") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef nil", actual)
}

func Test_Cov4_RawErrorType_MergeErrorWithMessageRef_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithMessageRef(errors.New("inner"), "msg", "ref")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef with err", actual)
}

func Test_Cov4_RawErrorType_MergeErrorWithRef_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.InvalidRequestType.MergeErrorWithRef(nil, "ref") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef nil", actual)
}

func Test_Cov4_RawErrorType_MergeErrorWithRef_WithErr(t *testing.T) {
	err := errcore.InvalidRequestType.MergeErrorWithRef(errors.New("inner"), "ref")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef with err", actual)
}

func Test_Cov4_RawErrorType_MsgCsvRef_NoItems(t *testing.T) {
	result := errcore.InvalidRequestType.MsgCsvRef("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef no items", actual)
}

func Test_Cov4_RawErrorType_MsgCsvRef_WithItems(t *testing.T) {
	result := errcore.InvalidRequestType.MsgCsvRef("msg", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef with items", actual)
}

func Test_Cov4_RawErrorType_MsgCsvRef_EmptyMsg(t *testing.T) {
	result := errcore.InvalidRequestType.MsgCsvRef("", "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef empty msg", actual)
}

func Test_Cov4_RawErrorType_MsgCsvRefError(t *testing.T) {
	err := errcore.InvalidRequestType.MsgCsvRefError("msg", "a")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRefError", actual)
}

func Test_Cov4_RawErrorType_ErrorRefOnly(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorRefOnly("ref")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorRefOnly", actual)
}

func Test_Cov4_RawErrorType_ErrorRefOnly_Nil(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorRefOnly(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorRefOnly nil", actual)
}

func Test_Cov4_RawErrorType_Expecting(t *testing.T) {
	err := errcore.InvalidRequestType.Expecting("expected", "actual")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Expecting", actual)
}

func Test_Cov4_RawErrorType_NoRef_EmptyMsg(t *testing.T) {
	result := errcore.InvalidRequestType.NoRef("")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.NoRef empty", actual)
}

func Test_Cov4_RawErrorType_NoRef_WithMsg(t *testing.T) {
	result := errcore.InvalidRequestType.NoRef("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.NoRef with msg", actual)
}

func Test_Cov4_RawErrorType_ErrorNoRefs(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorNoRefs("msg")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefs", actual)
}

func Test_Cov4_RawErrorType_ErrorNoRefs_EmptyMsg(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorNoRefs("")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefs empty", actual)
}

func Test_Cov4_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	err := errcore.InvalidRequestType.ErrorNoRefsSkip(0, "msg")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefsSkip", actual)
}

func Test_Cov4_RawErrorType_HandleUsingPanic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "RawErrorType.HandleUsingPanic", actual)
	}()
	errcore.InvalidRequestType.HandleUsingPanic("msg", "ref")
}

// ── GetSet / GetSetVariant ──

func Test_Cov4_GetSet_True(t *testing.T) {
	result := errcore.GetSet(true, errcore.InvalidRequestType, errcore.NotFound)
	actual := args.Map{"result": result.String()}
	expected := args.Map{"result": errcore.InvalidRequestType.String()}
	expected.ShouldBeEqual(t, 0, "GetSet true", actual)
}

func Test_Cov4_GetSet_False(t *testing.T) {
	result := errcore.GetSet(false, errcore.InvalidRequestType, errcore.NotFound)
	actual := args.Map{"result": result.String()}
	expected := args.Map{"result": errcore.NotFound.String()}
	expected.ShouldBeEqual(t, 0, "GetSet false", actual)
}

func Test_Cov4_GetSetVariant_True(t *testing.T) {
	result := errcore.GetSetVariant(true, "yes", "no")
	actual := args.Map{"result": result.String()}
	expected := args.Map{"result": "yes"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant true", actual)
}

func Test_Cov4_GetSetVariant_False(t *testing.T) {
	result := errcore.GetSetVariant(false, "yes", "no")
	actual := args.Map{"result": result.String()}
	expected := args.Map{"result": "no"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant false", actual)
}

// ── VarTwo / VarThree ──

func Test_Cov4_VarTwo_WithType(t *testing.T) {
	result := errcore.VarTwo(true, "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo with type", actual)
}

func Test_Cov4_VarTwo_NoType(t *testing.T) {
	result := errcore.VarTwo(false, "a", 1, "b", 2)
	actual := args.Map{"contains": strings.Contains(result, "a")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarTwo no type", actual)
}

func Test_Cov4_VarThree_WithType(t *testing.T) {
	result := errcore.VarThree(true, "a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree with type", actual)
}

func Test_Cov4_VarThree_NoType(t *testing.T) {
	result := errcore.VarThree(false, "a", 1, "b", 2, "c", 3)
	actual := args.Map{"contains": strings.Contains(result, "a")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarThree no type", actual)
}

func Test_Cov4_VarTwoNoType(t *testing.T) {
	result := errcore.VarTwoNoType("x", 10, "y", 20)
	actual := args.Map{"contains": strings.Contains(result, "x")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType", actual)
}

func Test_Cov4_VarThreeNoType(t *testing.T) {
	result := errcore.VarThreeNoType("x", 1, "y", 2, "z", 3)
	actual := args.Map{"contains": strings.Contains(result, "x")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType", actual)
}

// ── VarMap / MessageVarMap / MessageVarTwo / MessageVarThree ──

func Test_Cov4_VarMap_Empty(t *testing.T) {
	actual := args.Map{"result": errcore.VarMap(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarMap empty", actual)
}

func Test_Cov4_VarMap_NonEmpty(t *testing.T) {
	result := errcore.VarMap(map[string]any{"key": "val"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarMap non-empty", actual)
}

func Test_Cov4_MessageVarMap_Empty(t *testing.T) {
	actual := args.Map{"result": errcore.MessageVarMap("msg", nil)}
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageVarMap empty map", actual)
}

func Test_Cov4_MessageVarMap_NonEmpty(t *testing.T) {
	result := errcore.MessageVarMap("msg", map[string]any{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap non-empty", actual)
}

func Test_Cov4_MessageVarTwo(t *testing.T) {
	result := errcore.MessageVarTwo("msg", "a", 1, "b", 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo", actual)
}

func Test_Cov4_MessageVarThree(t *testing.T) {
	result := errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree", actual)
}

// ── VarNameValues / MessageNameValues ──

func Test_Cov4_VarNameValues_Empty(t *testing.T) {
	actual := args.Map{"result": errcore.VarNameValues()}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "VarNameValues empty", actual)
}

func Test_Cov4_VarNameValues_NonEmpty(t *testing.T) {
	nv := namevalue.StringAny{Name: "k", Value: "v"}
	actual := args.Map{"notEmpty": errcore.VarNameValues(nv) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues non-empty", actual)
}

func Test_Cov4_MessageNameValues_Empty(t *testing.T) {
	actual := args.Map{"result": errcore.MessageNameValues("msg")}
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MessageNameValues empty", actual)
}

func Test_Cov4_MessageNameValues_NonEmpty(t *testing.T) {
	nv := namevalue.StringAny{Name: "k", Value: "v"}
	result := errcore.MessageNameValues("msg", nv)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues non-empty", actual)
}

// ── Expecting functions ──

func Test_Cov4_Expecting(t *testing.T) {
	result := errcore.Expecting("title", "expected", "actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting", actual)
}

func Test_Cov4_ExpectingSimple(t *testing.T) {
	result := errcore.ExpectingSimple("title", "expected", "actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple", actual)
}

func Test_Cov4_ExpectingSimpleNoType(t *testing.T) {
	result := errcore.ExpectingSimpleNoType("title", "expected", "actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType", actual)
}

func Test_Cov4_ExpectingErrorSimpleNoType(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("title", "exp", "act")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType", actual)
}

func Test_Cov4_ExpectingErrorSimpleNoTypeNewLineEnds(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoTypeNewLineEnds("title", "exp", "act")
	actual := args.Map{"hasErr": err != nil, "endsNewLine": strings.HasSuffix(err.Error(), "\n")}
	expected := args.Map{"hasErr": true, "endsNewLine": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoTypeNewLineEnds", actual)
}

func Test_Cov4_WasExpectingErrorF(t *testing.T) {
	err := errcore.WasExpectingErrorF("exp", "act", "title %d", 1)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "WasExpectingErrorF", actual)
}

func Test_Cov4_ExpectingSimpleNoTypeError(t *testing.T) {
	err := errcore.ExpectingSimpleNoTypeError("title", "exp", "act")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoTypeError", actual)
}

func Test_Cov4_ExpectingNotEqualSimpleNoType(t *testing.T) {
	result := errcore.ExpectingNotEqualSimpleNoType("title", "exp", "act")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType", actual)
}

// ── ExpectingFuture / ExpectingRecord ──

func Test_Cov4_ExpectingFuture(t *testing.T) {
	rec := errcore.ExpectingFuture("title", "expected")
	actual := args.Map{
		"msg":            rec.Message("actual") != "",
		"msgSimple":      rec.MessageSimple("actual") != "",
		"msgSimpleNoType": rec.MessageSimpleNoType("actual") != "",
		"err":            rec.Error("actual") != nil,
		"errSimple":      rec.ErrorSimple("actual") != nil,
		"errSimpleNoType": rec.ErrorSimpleNoType("actual") != nil,
	}
	expected := args.Map{
		"msg":            true,
		"msgSimple":      true,
		"msgSimpleNoType": true,
		"err":            true,
		"errSimple":      true,
		"errSimpleNoType": true,
	}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture/ExpectingRecord", actual)
}

// ── expected struct methods ──

func Test_Cov4_Expected_But(t *testing.T) {
	err := errcore.Expected.But("title", "exp", "act")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.But", actual)
}

func Test_Cov4_Expected_ButFoundAsMsg(t *testing.T) {
	result := errcore.Expected.ButFoundAsMsg("title", "exp", "act")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg", actual)
}

func Test_Cov4_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	result := errcore.Expected.ButFoundWithTypeAsMsg("title", "exp", "act")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg", actual)
}

func Test_Cov4_Expected_ButUsingType(t *testing.T) {
	err := errcore.Expected.ButUsingType("title", "exp", "act")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType", actual)
}

func Test_Cov4_Expected_ReflectButFound(t *testing.T) {
	err := errcore.Expected.ReflectButFound(1, 2) // reflect.Kind values
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.ReflectButFound", actual)
}

func Test_Cov4_Expected_PrimitiveButFound(t *testing.T) {
	err := errcore.Expected.PrimitiveButFound(20) // reflect.Kind
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.PrimitiveButFound", actual)
}

func Test_Cov4_Expected_ValueHasNoElements(t *testing.T) {
	err := errcore.Expected.ValueHasNoElements(23) // reflect.Kind
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Expected.ValueHasNoElements", actual)
}

// ── shouldBe struct methods ──

func Test_Cov4_ShouldBe_StrEqMsg(t *testing.T) {
	result := errcore.ShouldBe.StrEqMsg("actual", "expecting")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg", actual)
}

func Test_Cov4_ShouldBe_StrEqErr(t *testing.T) {
	err := errcore.ShouldBe.StrEqErr("actual", "expecting")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr", actual)
}

func Test_Cov4_ShouldBe_AnyEqMsg(t *testing.T) {
	result := errcore.ShouldBe.AnyEqMsg(1, 2)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg", actual)
}

func Test_Cov4_ShouldBe_AnyEqErr(t *testing.T) {
	err := errcore.ShouldBe.AnyEqErr(1, 2)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr", actual)
}

func Test_Cov4_ShouldBe_JsonEqMsg(t *testing.T) {
	result := errcore.ShouldBe.JsonEqMsg("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg", actual)
}

func Test_Cov4_ShouldBe_JsonEqErr(t *testing.T) {
	err := errcore.ShouldBe.JsonEqErr("a", "b")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr", actual)
}

// ── Slice functions ──

func Test_Cov4_SliceToError_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.SliceToError(nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToError nil", actual)
}

func Test_Cov4_SliceToError_NonEmpty(t *testing.T) {
	err := errcore.SliceToError([]string{"e1", "e2"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToError non-empty", actual)
}

func Test_Cov4_SliceToErrorPtr_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.SliceToErrorPtr(nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr nil", actual)
}

func Test_Cov4_SliceToErrorPtr_NonEmpty(t *testing.T) {
	err := errcore.SliceToErrorPtr([]string{"e1"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceToErrorPtr non-empty", actual)
}

func Test_Cov4_SliceError_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.SliceError(",", nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SliceError empty", actual)
}

func Test_Cov4_SliceError_NonEmpty(t *testing.T) {
	err := errcore.SliceError(",", []string{"a", "b"})
	actual := args.Map{"result": err.Error()}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "SliceError non-empty", actual)
}

func Test_Cov4_SliceErrorDefault_NonEmpty(t *testing.T) {
	err := errcore.SliceErrorDefault([]string{"a"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceErrorDefault non-empty", actual)
}

func Test_Cov4_SliceErrorsToStrings_Nil(t *testing.T) {
	actual := args.Map{"len": len(errcore.SliceErrorsToStrings(nil...))}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings nil", actual)
}

func Test_Cov4_SliceErrorsToStrings_Mixed(t *testing.T) {
	result := errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b"))
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SliceErrorsToStrings mixed", actual)
}

// ── MergeErrors / MergeErrorsToString ──

func Test_Cov4_MergeErrors_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.MergeErrors(nil, nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors nil", actual)
}

func Test_Cov4_MergeErrors_NonNil(t *testing.T) {
	err := errcore.MergeErrors(errors.New("a"), errors.New("b"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MergeErrors non-nil", actual)
}

func Test_Cov4_MergeErrorsToString_Nil(t *testing.T) {
	actual := args.Map{"result": errcore.MergeErrorsToString(",", nil...)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString nil", actual)
}

func Test_Cov4_MergeErrorsToString_NonNil(t *testing.T) {
	result := errcore.MergeErrorsToString(",", errors.New("a"), errors.New("b"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToString non-nil", actual)
}

func Test_Cov4_MergeErrorsToStringDefault_Nil(t *testing.T) {
	actual := args.Map{"result": errcore.MergeErrorsToStringDefault(nil...)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault nil", actual)
}

func Test_Cov4_MergeErrorsToStringDefault_NonNil(t *testing.T) {
	result := errcore.MergeErrorsToStringDefault(errors.New("a"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MergeErrorsToStringDefault non-nil", actual)
}

// ── ToString / ToStringPtr / ToError / Ref ──

func Test_Cov4_ToString_Nil(t *testing.T) {
	actual := args.Map{"result": errcore.ToString(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToString nil", actual)
}

func Test_Cov4_ToString_NonNil(t *testing.T) {
	actual := args.Map{"result": errcore.ToString(errors.New("hello"))}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ToString non-nil", actual)
}

func Test_Cov4_ToStringPtr_Nil(t *testing.T) {
	result := errcore.ToStringPtr(nil)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ToStringPtr nil", actual)
}

func Test_Cov4_ToStringPtr_NonNil(t *testing.T) {
	result := errcore.ToStringPtr(errors.New("hi"))
	actual := args.Map{"result": *result}
	expected := args.Map{"result": "hi"}
	expected.ShouldBeEqual(t, 0, "ToStringPtr non-nil", actual)
}

func Test_Cov4_ToError_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.ToError("") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ToError empty", actual)
}

func Test_Cov4_ToError_NonEmpty(t *testing.T) {
	err := errcore.ToError("msg")
	actual := args.Map{"msg": err.Error()}
	expected := args.Map{"msg": "msg"}
	expected.ShouldBeEqual(t, 0, "ToError non-empty", actual)
}

func Test_Cov4_Ref_Nil(t *testing.T) {
	actual := args.Map{"result": errcore.Ref(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Ref nil", actual)
}

func Test_Cov4_Ref_NonNil(t *testing.T) {
	result := errcore.Ref("val")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Ref non-nil", actual)
}

// ── MessageWithRef ──

func Test_Cov4_MessageWithRef(t *testing.T) {
	result := errcore.MessageWithRef("msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef", actual)
}

// ── SourceDestination / SourceDestinationErr / SourceDestinationNoType ──

func Test_Cov4_SourceDestination_WithType(t *testing.T) {
	result := errcore.SourceDestination(true, "src", "dst")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination with type", actual)
}

func Test_Cov4_SourceDestination_NoType(t *testing.T) {
	result := errcore.SourceDestination(false, "src", "dst")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination no type", actual)
}

func Test_Cov4_SourceDestinationErr(t *testing.T) {
	err := errcore.SourceDestinationErr(false, "src", "dst")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr", actual)
}

func Test_Cov4_SourceDestinationNoType(t *testing.T) {
	result := errcore.SourceDestinationNoType("src", "dst")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType", actual)
}

// ── stackTraceEnhance ──

func Test_Cov4_StackEnhance_Error_Nil(t *testing.T) {
	actual := args.Map{"isNil": errcore.StackEnhance.Error(nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error nil", actual)
}

func Test_Cov4_StackEnhance_Error_NonNil(t *testing.T) {
	err := errcore.StackEnhance.Error(errors.New("test"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error non-nil", actual)
}

func Test_Cov4_StackEnhance_Msg_Empty(t *testing.T) {
	actual := args.Map{"result": errcore.StackEnhance.Msg("")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg empty", actual)
}

func Test_Cov4_StackEnhance_Msg_NonEmpty(t *testing.T) {
	result := errcore.StackEnhance.Msg("test")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg non-empty", actual)
}

func Test_Cov4_StackEnhance_MsgToErrSkip_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.StackEnhance.MsgToErrSkip(0, "") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgToErrSkip empty", actual)
}

func Test_Cov4_StackEnhance_FmtSkip_Empty(t *testing.T) {
	actual := args.Map{"isNil": errcore.StackEnhance.FmtSkip(0, "") == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip empty", actual)
}

func Test_Cov4_StackEnhance_FmtSkip_NonEmpty(t *testing.T) {
	err := errcore.StackEnhance.FmtSkip(0, "val %d", 1)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip non-empty", actual)
}

func Test_Cov4_StackEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	actual := args.Map{"result": errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip nil err", actual)
}

func Test_Cov4_StackEnhance_MsgErrorToErrSkip_NilErr(t *testing.T) {
	actual := args.Map{"isNil": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip nil", actual)
}

func Test_Cov4_StackEnhance_MsgErrorToErrSkip_WithErr(t *testing.T) {
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("inner"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip with err", actual)
}

// ── LineDiff ──

func Test_Cov4_LineDiff_Matching(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a", "b"})
	actual := args.Map{"len": len(diffs), "status0": diffs[0].Status}
	expected := args.Map{"len": 2, "status0": "  "}
	expected.ShouldBeEqual(t, 0, "LineDiff matching", actual)
}

func Test_Cov4_LineDiff_Mismatch(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a"}, []string{"b"})
	actual := args.Map{"status": diffs[0].Status}
	expected := args.Map{"status": "!!"}
	expected.ShouldBeEqual(t, 0, "LineDiff mismatch", actual)
}

func Test_Cov4_LineDiff_ExtraActual(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a", "b"}, []string{"a"})
	actual := args.Map{"status1": diffs[1].Status}
	expected := args.Map{"status1": "+"}
	expected.ShouldBeEqual(t, 0, "LineDiff extra actual", actual)
}

func Test_Cov4_LineDiff_MissingExpected(t *testing.T) {
	diffs := errcore.LineDiff([]string{"a"}, []string{"a", "b"})
	actual := args.Map{"status1": diffs[1].Status}
	expected := args.Map{"status1": "-"}
	expected.ShouldBeEqual(t, 0, "LineDiff missing expected", actual)
}

func Test_Cov4_LineDiffToString(t *testing.T) {
	result := errcore.LineDiffToString(0, "test", []string{"a"}, []string{"b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString", actual)
}

func Test_Cov4_LineDiffToString_Empty(t *testing.T) {
	result := errcore.LineDiffToString(0, "test", nil, nil)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LineDiffToString empty", actual)
}

func Test_Cov4_HasAnyMismatchOnLines_Same(t *testing.T) {
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines same", actual)
}

func Test_Cov4_HasAnyMismatchOnLines_Different(t *testing.T) {
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines different", actual)
}

func Test_Cov4_HasAnyMismatchOnLines_DiffLen(t *testing.T) {
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyMismatchOnLines diff len", actual)
}

func Test_Cov4_SliceDiffSummary_Match(t *testing.T) {
	actual := args.Map{"result": errcore.SliceDiffSummary([]string{"a"}, []string{"a"})}
	expected := args.Map{"result": "all lines match"}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary match", actual)
}

func Test_Cov4_SliceDiffSummary_Mismatch(t *testing.T) {
	result := errcore.SliceDiffSummary([]string{"a"}, []string{"b"})
	actual := args.Map{"contains": strings.Contains(result, "mismatch")}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "SliceDiffSummary mismatch", actual)
}

func Test_Cov4_ErrorToLinesLineDiff_NilErr(t *testing.T) {
	result := errcore.ErrorToLinesLineDiff(0, "test", nil, []string{"a"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorToLinesLineDiff nil err", actual)
}

// ── RawErrCollection ──

func Test_Cov4_RawErrCollection_BasicOps(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.AddError(errors.New("e2"))
	c.Add(nil) // should skip
	c.AddString("e3")
	c.AddString("") // should skip

	actual := args.Map{
		"len":          c.Length(),
		"isEmpty":      c.IsEmpty(),
		"hasError":     c.HasError(),
		"hasAnyError":  c.HasAnyError(),
		"hasAnyIssues": c.HasAnyIssues(),
		"isDefined":    c.IsDefined(),
		"isValid":      c.IsValid(),
		"isSuccess":    c.IsSuccess(),
		"isFailed":     c.IsFailed(),
		"isInvalid":    c.IsInvalid(),
		"isNull":       c.IsNull(),
		"isAnyNull":    c.IsAnyNull(),
		"isCollType":   c.IsCollectionType(),
	}
	expected := args.Map{
		"len":          3,
		"isEmpty":      false,
		"hasError":     true,
		"hasAnyError":  true,
		"hasAnyIssues": true,
		"isDefined":    true,
		"isValid":      false,
		"isSuccess":    false,
		"isFailed":     true,
		"isInvalid":    true,
		"isNull":       false,
		"isAnyNull":    false,
		"isCollType":   true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection basic ops", actual)
}

func Test_Cov4_RawErrCollection_StringOps(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.Add(errors.New("e2"))

	actual := args.Map{
		"stringNotEmpty":    c.String() != "",
		"errorStringEq":    c.ErrorString() == c.String(),
		"compileEq":        c.Compile() == c.String(),
		"fullStringEq":     c.FullString() == c.String(),
		"stringsLen":       len(c.Strings()),
		"splitLen":         len(c.FullStringSplitByNewLine()),
		"joinerNotEmpty":   c.StringUsingJoiner(",") != "",
		"joinerAddlNotEm":  c.StringUsingJoinerAdditional(",", "!") != "",
		"withAddlNotEmpty": c.StringWithAdditionalMessage("!") != "",
		"refCompStr":       c.ReferencesCompiledString() != "",
		"fullWithoutRef":   c.FullStringWithoutReferences() != "",
	}
	expected := args.Map{
		"stringNotEmpty":    true,
		"errorStringEq":    true,
		"compileEq":        true,
		"fullStringEq":     true,
		"stringsLen":       2,
		"splitLen":         2,
		"joinerNotEmpty":   true,
		"joinerAddlNotEm":  true,
		"withAddlNotEmpty": true,
		"refCompStr":       true,
		"fullWithoutRef":   true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection string ops", actual)
}

func Test_Cov4_RawErrCollection_CompiledErrors(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))

	actual := args.Map{
		"compiledErr":       c.CompiledError() != nil,
		"compiledJoiner":    c.CompiledErrorUsingJoiner(",") != nil,
		"compiledJoinerAdd": c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") != nil,
		"compiledStacks":    c.CompiledErrorWithStackTraces() != nil,
		"value":             c.Value() != nil,
	}
	expected := args.Map{
		"compiledErr":       true,
		"compiledJoiner":    true,
		"compiledJoinerAdd": true,
		"compiledStacks":    true,
		"value":             true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection compiled errors", actual)
}

func Test_Cov4_RawErrCollection_EmptyPaths(t *testing.T) {
	c := &errcore.RawErrCollection{}

	actual := args.Map{
		"stringEmpty":     c.String() == "",
		"compiledNil":     c.CompiledError() == nil,
		"compiledJoinNil": c.CompiledErrorUsingJoiner(",") == nil,
		"compiledAddNil":  c.CompiledErrorUsingJoinerAdditionalMessage(",", "!") == nil,
		"compiledStkNil":  c.CompiledErrorWithStackTraces() == nil,
		"stkStrEmpty":     c.CompiledStackTracesString() == "",
		"valueNil":        c.Value() == nil,
		"stringsLen":      len(c.Strings()),
		"joinerEmpty":     c.StringUsingJoiner(",") == "",
		"addlEmpty":       c.StringUsingJoinerAdditional(",", "!") == "",
		"withAddlEmpty":   c.StringWithAdditionalMessage("!") == "",
		"fullTracesIf":    c.FullStringWithTracesIf(false) == "",
	}
	expected := args.Map{
		"stringEmpty":     true,
		"compiledNil":     true,
		"compiledJoinNil": true,
		"compiledAddNil":  true,
		"compiledStkNil":  true,
		"stkStrEmpty":     true,
		"valueNil":        true,
		"stringsLen":      0,
		"joinerEmpty":     true,
		"addlEmpty":       true,
		"withAddlEmpty":   true,
		"fullTracesIf":    true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection empty paths", actual)
}

func Test_Cov4_RawErrCollection_AddOps(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Adds(errors.New("a"), nil, errors.New("b"))
	c.AddErrors(errors.New("c"))
	c.AddIf(true, "d")
	c.AddIf(false, "skip")
	c.ConditionalAddError(true, errors.New("e"))
	c.ConditionalAddError(false, errors.New("skip"))
	c.AddStringSliceAsErr("f", "", "g")
	c.AddFunc(func() error { return errors.New("h") })
	c.AddFunc(nil)
	c.AddFuncIf(true, func() error { return errors.New("i") })
	c.AddFuncIf(false, func() error { return errors.New("skip") })
	c.AddFuncIf(true, nil)

	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 9}
	expected.ShouldBeEqual(t, 0, "RawErrCollection add ops", actual)
}

func Test_Cov4_RawErrCollection_AddWithRef(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.AddWithRef(errors.New("e"), "ref")
	c.AddWithRef(nil, "ref")
	c.AddWithCompiledTraceRef(errors.New("e"), "trace", "ref")
	c.AddWithCompiledTraceRef(nil, "trace", "ref")
	c.AddWithTraceRef(errors.New("e"), []string{"t"}, "ref")
	c.AddWithTraceRef(nil, nil, "ref")

	actual := args.Map{"len": c.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "RawErrCollection add with ref", actual)
}

func Test_Cov4_RawErrCollection_ClearDispose(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.Clear()
	actual := args.Map{"isEmpty": c.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection clear", actual)
}

func Test_Cov4_RawErrCollection_Dispose(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Add(errors.New("e1"))
	c.Dispose()
	actual := args.Map{"isNull": c.IsNull()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection dispose", actual)
}

func Test_Cov4_RawErrCollection_ClearEmpty(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Clear()   // should not panic on empty
	c.Dispose() // should not panic on empty
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection clear/dispose empty", actual)
}

func Test_Cov4_RawErrCollection_ToRawErrCollection(t *testing.T) {
	c := errcore.RawErrCollection{}
	c.Add(errors.New("e"))
	ptr := c.ToRawErrCollection()
	actual := args.Map{"notNil": ptr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToRawErrCollection", actual)
}

func Test_Cov4_RawErrCollection_Serialize(t *testing.T) {
	c := &errcore.RawErrCollection{}
	emptyBytes, emptyErr := c.Serialize()
	c.Add(errors.New("e"))
	bytes, err := c.Serialize()
	mustBytes := c.SerializeMust()

	actual := args.Map{
		"emptyBytesNil": emptyBytes == nil,
		"emptyErrNil":   emptyErr == nil,
		"hasBytes":      len(bytes) > 0,
		"noErr":         err == nil,
		"mustHasBytes":  len(mustBytes) > 0,
	}
	expected := args.Map{
		"emptyBytesNil": true,
		"emptyErrNil":   true,
		"hasBytes":      true,
		"noErr":         true,
		"mustHasBytes":  true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrCollection serialize", actual)
}

func Test_Cov4_RawErrCollection_MarshalJSON_Empty(t *testing.T) {
	c := &errcore.RawErrCollection{}
	bytes, err := c.MarshalJSON()
	actual := args.Map{"bytesNil": bytes == nil, "errNil": err == nil}
	expected := args.Map{"bytesNil": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection MarshalJSON empty", actual)
}

func Test_Cov4_RawErrCollection_SerializeWithoutTraces_Empty(t *testing.T) {
	c := &errcore.RawErrCollection{}
	bytes, err := c.SerializeWithoutTraces()
	actual := args.Map{"bytesNil": bytes == nil, "errNil": err == nil}
	expected := args.Map{"bytesNil": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "SerializeWithoutTraces empty", actual)
}

func Test_Cov4_RawErrCollection_LogOps(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.Log()            // empty - should not log
	c.LogWithTraces()  // empty - should not log
	c.LogIf(false)     // false - should not log
	c.Add(errors.New("e"))
	c.Log()
	c.LogWithTraces()

	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection log ops", actual)
}

func Test_Cov4_RawErrCollection_HandleEmpty(t *testing.T) {
	c := &errcore.RawErrCollection{}
	c.HandleError()           // should not panic
	c.HandleErrorWithMsg("m") // should not panic
	c.HandleErrorWithRefs("m", "k", "v") // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "RawErrCollection handle empty", actual)
}

func Test_Cov4_RawErrCollection_IsErrorsCollected(t *testing.T) {
	c := &errcore.RawErrCollection{}
	changed := c.IsErrorsCollected(errors.New("e"))
	noChange := c.IsErrorsCollected(nil)
	actual := args.Map{"changed": changed, "noChange": noChange}
	expected := args.Map{"changed": true, "noChange": false}
	expected.ShouldBeEqual(t, 0, "IsErrorsCollected", actual)
}

func Test_Cov4_RawErrCollection_CountStateChangeTracker(t *testing.T) {
	c := errcore.RawErrCollection{}
	c.Add(errors.New("e"))
	tracker := c.CountStateChangeTracker()
	actual := args.Map{
		"isSameState": tracker.IsSameState(),
		"isValid":     tracker.IsValid(),
		"isSuccess":   tracker.IsSuccess(),
		"hasChanges":  tracker.HasChanges(),
		"isFailed":    tracker.IsFailed(),
		"isSameUsing": tracker.IsSameStateUsingCount(1),
	}
	expected := args.Map{
		"isSameState": true,
		"isValid":     true,
		"isSuccess":   true,
		"hasChanges":  false,
		"isFailed":    false,
		"isSameUsing": true,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker", actual)
}

// ── HandleErr (nil path) ──

func Test_Cov4_HandleErr_Nil(t *testing.T) {
	errcore.HandleErr(nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr nil", actual)
}

// ── MustBeEmpty (nil path) ──

func Test_Cov4_MustBeEmpty_Nil(t *testing.T) {
	errcore.MustBeEmpty(nil)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeEmpty nil", actual)
}
