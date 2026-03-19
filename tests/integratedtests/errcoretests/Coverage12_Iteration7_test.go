package errcoretests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/namevalue"
)

// ==========================================================================
// shouldBe — all methods
// ==========================================================================

func Test_Cov12_ShouldBe_StrEqMsg(t *testing.T) {
	msg := errcore.ShouldBe.StrEqMsg("actual", "expect")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg", actual)
}

func Test_Cov12_ShouldBe_StrEqErr(t *testing.T) {
	err := errcore.ShouldBe.StrEqErr("actual", "expect")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr", actual)
}

func Test_Cov12_ShouldBe_AnyEqMsg(t *testing.T) {
	msg := errcore.ShouldBe.AnyEqMsg(1, 2)
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg", actual)
}

func Test_Cov12_ShouldBe_AnyEqErr(t *testing.T) {
	err := errcore.ShouldBe.AnyEqErr(1, 2)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr", actual)
}

func Test_Cov12_ShouldBe_JsonEqMsg(t *testing.T) {
	msg := errcore.ShouldBe.JsonEqMsg("a", "b")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg", actual)
}

func Test_Cov12_ShouldBe_JsonEqErr(t *testing.T) {
	err := errcore.ShouldBe.JsonEqErr("a", "b")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr", actual)
}

// ==========================================================================
// expected — all methods
// ==========================================================================

func Test_Cov12_Expected_But(t *testing.T) {
	err := errcore.Expected.But("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.But", actual)
}

func Test_Cov12_Expected_ButFoundAsMsg(t *testing.T) {
	msg := errcore.Expected.ButFoundAsMsg("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg", actual)
}

func Test_Cov12_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	msg := errcore.Expected.ButFoundWithTypeAsMsg("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg", actual)
}

func Test_Cov12_Expected_ButUsingType(t *testing.T) {
	err := errcore.Expected.ButUsingType("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType", actual)
}

func Test_Cov12_Expected_ReflectButFound(t *testing.T) {
	err := errcore.Expected.ReflectButFound(reflect.String, reflect.Int)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ReflectButFound", actual)
}

func Test_Cov12_Expected_PrimitiveButFound(t *testing.T) {
	err := errcore.Expected.PrimitiveButFound(reflect.Struct)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.PrimitiveButFound", actual)
}

func Test_Cov12_Expected_ValueHasNoElements(t *testing.T) {
	err := errcore.Expected.ValueHasNoElements(reflect.Slice)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ValueHasNoElements", actual)
}

// ==========================================================================
// CountStateChangeTracker — deeper coverage
// ==========================================================================

func Test_Cov12_CountStateChangeTracker(t *testing.T) {
	rec := &errcore.RawErrCollection{}
	rec.Add(errors.New("a"))
	tracker := errcore.NewCountStateChangeTracker(rec)
	actual := args.Map{
		"sameState":     tracker.IsSameState(),
		"isValid":       tracker.IsValid(),
		"isSuccess":     tracker.IsSuccess(),
		"hasChanges":    tracker.HasChanges(),
		"isFailed":      tracker.IsFailed(),
		"sameUsingCount": tracker.IsSameStateUsingCount(1),
	}
	expected := args.Map{
		"sameState":     true,
		"isValid":       true,
		"isSuccess":     true,
		"hasChanges":    false,
		"isFailed":      false,
		"sameUsingCount": true,
	}
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker no changes", actual)

	rec.Add(errors.New("b"))
	actual2 := args.Map{
		"sameState":  tracker.IsSameState(),
		"hasChanges": tracker.HasChanges(),
		"isFailed":   tracker.IsFailed(),
	}
	expected2 := args.Map{
		"sameState":  false,
		"hasChanges": true,
		"isFailed":   true,
	}
	expected2.ShouldBeEqual(t, 1, "CountStateChangeTracker with changes", actual2)
}

// ==========================================================================
// StackTracesCompiled
// ==========================================================================

func Test_Cov12_StackTracesCompiled(t *testing.T) {
	result := errcore.StackTracesCompiled([]string{"trace1", "trace2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled", actual)
}

// ==========================================================================
// stackTraceEnhance — all methods
// ==========================================================================

func Test_Cov12_StackEnhance_Error(t *testing.T) {
	err := errcore.StackEnhance.Error(errors.New("e"))
	errNil := errcore.StackEnhance.Error(nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error", actual)
}

func Test_Cov12_StackEnhance_ErrorSkip(t *testing.T) {
	err := errcore.StackEnhance.ErrorSkip(0, errors.New("e"))
	errNil := errcore.StackEnhance.ErrorSkip(0, nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.ErrorSkip", actual)
}

func Test_Cov12_StackEnhance_MsgToErrSkip(t *testing.T) {
	err := errcore.StackEnhance.MsgToErrSkip(0, "msg")
	errEmpty := errcore.StackEnhance.MsgToErrSkip(0, "")
	actual := args.Map{"notNil": err != nil, "nil": errEmpty == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgToErrSkip", actual)
}

func Test_Cov12_StackEnhance_FmtSkip(t *testing.T) {
	err := errcore.StackEnhance.FmtSkip(0, "hello %d", 1)
	errEmpty := errcore.StackEnhance.FmtSkip(0, "")
	actual := args.Map{"notNil": err != nil, "nil": errEmpty == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip", actual)
}

func Test_Cov12_StackEnhance_Msg(t *testing.T) {
	msg := errcore.StackEnhance.Msg("hello")
	msgEmpty := errcore.StackEnhance.Msg("")
	actual := args.Map{"notEmpty": msg != "", "empty": msgEmpty == ""}
	expected := args.Map{"notEmpty": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg", actual)
}

func Test_Cov12_StackEnhance_MsgErrorSkip(t *testing.T) {
	msg := errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("e"))
	msgNil := errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)
	actual := args.Map{"notEmpty": msg != "", "empty": msgNil == ""}
	expected := args.Map{"notEmpty": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip", actual)
}

func Test_Cov12_StackEnhance_MsgErrorToErrSkip(t *testing.T) {
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e"))
	errNil := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip", actual)
}

// ==========================================================================
// ExpectingFuture / ExpectingRecord
// ==========================================================================

func Test_Cov12_ExpectingFuture(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	actual := args.Map{
		"title": er.ExpectingTitle,
		"was":   er.WasExpecting,
	}
	expected := args.Map{
		"title": "title",
		"was":   "expect",
	}
	expected.ShouldBeEqual(t, 0, "ExpectingFuture", actual)
}

func Test_Cov12_ExpectingRecord_Message(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.Message("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Message", actual)
}

func Test_Cov12_ExpectingRecord_MessageSimple(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.MessageSimple("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimple", actual)
}

func Test_Cov12_ExpectingRecord_MessageSimpleNoType(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.MessageSimpleNoType("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimpleNoType", actual)
}

func Test_Cov12_ExpectingRecord_Error(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	err := er.Error("actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Error", actual)
}

func Test_Cov12_ExpectingRecord_ErrorSimple(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	err := er.ErrorSimple("actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimple", actual)
}

func Test_Cov12_ExpectingRecord_ErrorSimpleNoType(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	err := er.ErrorSimpleNoType("actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimpleNoType", actual)
}

// ==========================================================================
// Expecting / ExpectingSimple / ExpectingSimpleNoType / ExpectingError*
// ==========================================================================

func Test_Cov12_Expecting(t *testing.T) {
	msg := errcore.Expecting("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting", actual)
}

func Test_Cov12_ExpectingSimple(t *testing.T) {
	msg := errcore.ExpectingSimple("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple", actual)
}

func Test_Cov12_ExpectingSimpleNoType(t *testing.T) {
	msg := errcore.ExpectingSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType", actual)
}

func Test_Cov12_ExpectingErrorSimpleNoType(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType", actual)
}

func Test_Cov12_ExpectingNotEqualSimpleNoType(t *testing.T) {
	msg := errcore.ExpectingNotEqualSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType", actual)
}

func Test_Cov12_ExpectingError(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType", actual)
}

// ==========================================================================
// RawErrorType — remaining uncovered methods
// ==========================================================================

func Test_Cov12_RawErrorType_CombineWithAnother(t *testing.T) {
	result := errcore.InvalidType.CombineWithAnother(errcore.NotFound, "msg", "ref")
	actual := args.Map{"notEmpty": string(result) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.CombineWithAnother", actual)
}

func Test_Cov12_RawErrorType_TypesAttach(t *testing.T) {
	result := errcore.InvalidType.TypesAttach("msg", "str", 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttach", actual)
}

func Test_Cov12_RawErrorType_TypesAttachErr(t *testing.T) {
	err := errcore.InvalidType.TypesAttachErr("msg", "str", 42)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttachErr", actual)
}

func Test_Cov12_RawErrorType_SrcDestination(t *testing.T) {
	result := errcore.InvalidType.SrcDestination("msg", "src", "sv", "dst", "dv")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestination", actual)
}

func Test_Cov12_RawErrorType_SrcDestinationErr(t *testing.T) {
	err := errcore.InvalidType.SrcDestinationErr("msg", "src", "sv", "dst", "dv")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestinationErr", actual)
}

func Test_Cov12_RawErrorType_Error(t *testing.T) {
	err := errcore.InvalidType.Error("msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Error", actual)
}

func Test_Cov12_RawErrorType_ErrorSkip(t *testing.T) {
	err := errcore.InvalidType.ErrorSkip(0, "msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorSkip", actual)
}

func Test_Cov12_RawErrorType_Fmt(t *testing.T) {
	err := errcore.InvalidType.Fmt("hello %d", 42)
	errEmpty := errcore.InvalidType.Fmt("")
	actual := args.Map{"notNil": err != nil, "emptyNotNil": errEmpty != nil}
	expected := args.Map{"notNil": true, "emptyNotNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Fmt", actual)
}

func Test_Cov12_RawErrorType_FmtIf(t *testing.T) {
	err := errcore.InvalidType.FmtIf(true, "hello %d", 42)
	errNil := errcore.InvalidType.FmtIf(false, "hello %d", 42)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf", actual)
}

func Test_Cov12_RawErrorType_MergeError(t *testing.T) {
	err := errcore.InvalidType.MergeError(errors.New("e"))
	errNil := errcore.InvalidType.MergeError(nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError", actual)
}

func Test_Cov12_RawErrorType_MergeErrorWithMessage(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessage(errors.New("e"), "msg")
	errNil := errcore.InvalidType.MergeErrorWithMessage(nil, "msg")
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage", actual)
}

func Test_Cov12_RawErrorType_MergeErrorWithMessageRef(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref")
	errNil := errcore.InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref")
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef", actual)
}

func Test_Cov12_RawErrorType_MergeErrorWithRef(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithRef(errors.New("e"), "ref")
	errNil := errcore.InvalidType.MergeErrorWithRef(nil, "ref")
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef", actual)
}

func Test_Cov12_RawErrorType_MsgCsvRef(t *testing.T) {
	result := errcore.InvalidType.MsgCsvRef("msg", "a", "b")
	resultEmpty := errcore.InvalidType.MsgCsvRef("", "a")
	resultNoRef := errcore.InvalidType.MsgCsvRef("msg")
	actual := args.Map{
		"notEmpty":      result != "",
		"emptyMsg":      resultEmpty != "",
		"noRef":         resultNoRef != "",
	}
	expected := args.Map{
		"notEmpty":      true,
		"emptyMsg":      true,
		"noRef":         true,
	}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef", actual)
}

func Test_Cov12_RawErrorType_MsgCsvRefError(t *testing.T) {
	err := errcore.InvalidType.MsgCsvRefError("msg", "a", "b")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRefError", actual)
}

func Test_Cov12_RawErrorType_ErrorRefOnly(t *testing.T) {
	err := errcore.InvalidType.ErrorRefOnly("ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorRefOnly", actual)
}

func Test_Cov12_RawErrorType_Expecting(t *testing.T) {
	err := errcore.InvalidType.Expecting("expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Expecting", actual)
}

func Test_Cov12_RawErrorType_NoRef(t *testing.T) {
	result := errcore.InvalidType.NoRef("msg")
	resultEmpty := errcore.InvalidType.NoRef("")
	actual := args.Map{"notEmpty": result != "", "emptyNotEmpty": resultEmpty != ""}
	expected := args.Map{"notEmpty": true, "emptyNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.NoRef", actual)
}

func Test_Cov12_RawErrorType_ErrorNoRefs(t *testing.T) {
	err := errcore.InvalidType.ErrorNoRefs("msg")
	errEmpty := errcore.InvalidType.ErrorNoRefs("")
	actual := args.Map{"notNil": err != nil, "emptyNotNil": errEmpty != nil}
	expected := args.Map{"notNil": true, "emptyNotNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefs", actual)
}

func Test_Cov12_RawErrorType_HandleUsingPanic(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.InvalidType.HandleUsingPanic("msg", "ref")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.HandleUsingPanic", actual)
}

func Test_Cov12_GetSet(t *testing.T) {
	r1 := errcore.GetSet(true, errcore.InvalidType, errcore.NotFound)
	r2 := errcore.GetSet(false, errcore.InvalidType, errcore.NotFound)
	actual := args.Map{"r1": r1, "r2": r2}
	expected := args.Map{"r1": errcore.InvalidType, "r2": errcore.NotFound}
	expected.ShouldBeEqual(t, 0, "GetSet", actual)
}

func Test_Cov12_GetSetVariant(t *testing.T) {
	r1 := errcore.GetSetVariant(true, "a", "b")
	r2 := errcore.GetSetVariant(false, "a", "b")
	actual := args.Map{"r1": string(r1), "r2": string(r2)}
	expected := args.Map{"r1": "a", "r2": "b"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant", actual)
}

// ==========================================================================
// GherkinsString / GherkinsStringWithExpectation
// ==========================================================================

func Test_Cov12_GherkinsString(t *testing.T) {
	result := errcore.GherkinsString(0, "feature", "given", "when", "then")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString", actual)
}

func Test_Cov12_GherkinsStringWithExpectation(t *testing.T) {
	result := errcore.GherkinsStringWithExpectation(0, "feature", "given", "when", "then", "actual", "expect")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation", actual)
}

// ==========================================================================
// Message formatting functions
// ==========================================================================

func Test_Cov12_MessageNameValues(t *testing.T) {
	result := errcore.MessageNameValues("msg", namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues", actual)
}

func Test_Cov12_MessageVarTwo(t *testing.T) {
	result := errcore.MessageVarTwo("msg", "n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo", actual)
}

func Test_Cov12_MessageVarThree(t *testing.T) {
	result := errcore.MessageVarThree("msg", "n1", "v1", "n2", "v2", "n3", "v3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree", actual)
}

func Test_Cov12_MessageVarMap(t *testing.T) {
	result := errcore.MessageVarMap("msg", map[string]any{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap", actual)
}

func Test_Cov12_MessageWithRef(t *testing.T) {
	result := errcore.MessageWithRef("msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef", actual)
}

func Test_Cov12_MessageWithRefToError(t *testing.T) {
	err := errcore.MessageWithRefToError("msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError", actual)
}

func Test_Cov12_VarTwo(t *testing.T) {
	result := errcore.VarTwo(true, "n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo", actual)
}

func Test_Cov12_VarTwoNoType(t *testing.T) {
	result := errcore.VarTwo(false, "n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType", actual)
}

func Test_Cov12_VarThree(t *testing.T) {
	result := errcore.VarThree(true, "n1", "v1", "n2", "v2", "n3", "v3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree", actual)
}

func Test_Cov12_VarThreeNoType(t *testing.T) {
	result := errcore.VarThree(false, "n1", "v1", "n2", "v2", "n3", "v3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType", actual)
}

func Test_Cov12_VarNameValues(t *testing.T) {
	result := errcore.VarNameValues("n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues", actual)
}

func Test_Cov12_VarNameValuesJoiner(t *testing.T) {
	result := errcore.VarNameValuesJoiner(",", "n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner", actual)
}

func Test_Cov12_VarNameValuesStrings(t *testing.T) {
	result := errcore.VarNameValuesStrings("n1", "v1", "n2", "v2")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings", actual)
}

// ==========================================================================
// ErrorWith* and Handle* functions
// ==========================================================================

func Test_Cov12_ErrorWithRef(t *testing.T) {
	result := errcore.ErrorWithRef(errors.New("e"), "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef", actual)
}

func Test_Cov12_ErrorWithRefToError(t *testing.T) {
	err := errcore.ErrorWithRefToError(errors.New("e"), "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError", actual)
}

func Test_Cov12_ErrorWithCompiledTraceRef(t *testing.T) {
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "trace", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef", actual)
}

func Test_Cov12_ErrorWithCompiledTraceRefToError(t *testing.T) {
	err := errcore.ErrorWithCompiledTraceRefToError(errors.New("e"), "trace", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError", actual)
}

func Test_Cov12_ErrorWithTracesRefToError(t *testing.T) {
	err := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{"t"}, "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError", actual)
}

func Test_Cov12_HandleErr(t *testing.T) {
	errcore.HandleErr(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr nil", actual)
}

func Test_Cov12_HandleErr_Panic(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.HandleErr(errors.New("e"))
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErr panic", actual)
}

func Test_Cov12_HandleErrMessage(t *testing.T) {
	errcore.HandleErrMessage(nil, "msg") // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage nil", actual)
}

func Test_Cov12_SimpleHandleErr(t *testing.T) {
	errcore.SimpleHandleErr(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr nil", actual)
}

func Test_Cov12_SimpleHandleErrMany(t *testing.T) {
	errcore.SimpleHandleErrMany(nil, nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany nil", actual)
}

func Test_Cov12_HandleCompiledErrorGetter(t *testing.T) {
	errcore.HandleCompiledErrorGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorGetter nil", actual)
}

func Test_Cov12_HandleErrorGetter(t *testing.T) {
	errcore.HandleErrorGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorGetter nil", actual)
}

// ==========================================================================
// CombineWithMsgType variants
// ==========================================================================

func Test_Cov12_CombineWithMsgTypeNoStack(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack", actual)
}

func Test_Cov12_CombineWithMsgTypeStackTrace(t *testing.T) {
	result := errcore.CombineWithMsgTypeStackTrace(errcore.InvalidType, "msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeStackTrace", actual)
}

func Test_Cov12_Combine_Func(t *testing.T) {
	result := errcore.Combine("errType", "msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine", actual)
}

// ==========================================================================
// MeaningFulError / MeaningfulMessageError
// ==========================================================================

func Test_Cov12_MeaningFulError(t *testing.T) {
	err := errcore.MeaningFulError("msg", errors.New("e"))
	errNil := errcore.MeaningFulError("msg", nil)
	actual := args.Map{"notNil": err != nil, "nilNil": errNil == nil}
	expected := args.Map{"notNil": true, "nilNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningFulError", actual)
}

func Test_Cov12_MeaningFulErrorHandle(t *testing.T) {
	errcore.MeaningFulErrorHandle("msg", nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MeaningFulErrorHandle nil", actual)
}

func Test_Cov12_MeaningFulErrorWithData(t *testing.T) {
	err := errcore.MeaningFulErrorWithData("msg", errors.New("e"), "data")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningFulErrorWithData", actual)
}

func Test_Cov12_MeaningfulMessageError(t *testing.T) {
	err := errcore.MeaningfulMessageError("msg", errors.New("e"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError", actual)
}

// ==========================================================================
// MsgHeader / MsgHeaderIf / MsgHeaderPlusEnding
// ==========================================================================

func Test_Cov12_MsgHeader(t *testing.T) {
	result := errcore.MsgHeader("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader", actual)
}

func Test_Cov12_MsgHeaderIf(t *testing.T) {
	result := errcore.MsgHeaderIf(true, "msg")
	resultFalse := errcore.MsgHeaderIf(false, "msg")
	actual := args.Map{"notEmpty": result != "", "empty": resultFalse == ""}
	expected := args.Map{"notEmpty": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf", actual)
}

func Test_Cov12_MsgHeaderPlusEnding(t *testing.T) {
	result := errcore.MsgHeaderPlusEnding("msg", "ending")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding", actual)
}

// ==========================================================================
// StringLines functions
// ==========================================================================

func Test_Cov12_StringLinesToQuoteLines(t *testing.T) {
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines", actual)
}

func Test_Cov12_StringLinesToQuoteLinesWithTabs(t *testing.T) {
	result := errcore.StringLinesToQuoteLinesWithTabs([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesWithTabs", actual)
}

func Test_Cov12_StringLinesToQuoteLinesToSingle(t *testing.T) {
	result := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle", actual)
}

// ==========================================================================
// Print / FmtDebug / FmtDebugIf
// ==========================================================================

func Test_Cov12_PrintError(t *testing.T) {
	errcore.PrintError(nil)
	errcore.PrintError(errors.New("e"))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintError", actual)
}

func Test_Cov12_PrintErrorWithTestIndex(t *testing.T) {
	errcore.PrintErrorWithTestIndex(0, errors.New("e"))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex", actual)
}

func Test_Cov12_FmtDebug(t *testing.T) {
	result := errcore.FmtDebug("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug", actual)
}

func Test_Cov12_FmtDebugIf(t *testing.T) {
	result := errcore.FmtDebugIf(true, "msg")
	resultFalse := errcore.FmtDebugIf(false, "msg")
	actual := args.Map{"notEmpty": result != "", "empty": resultFalse == ""}
	expected := args.Map{"notEmpty": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf", actual)
}

// ==========================================================================
// GetActualAndExpect / GetSearchLine / PathMeaningFul
// ==========================================================================

func Test_Cov12_GetActualAndExpectProcessedMessage(t *testing.T) {
	result := errcore.GetActualAndExpectProcessedMessage("act", "exp")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectProcessedMessage", actual)
}

func Test_Cov12_GetActualAndExpectSortedMessage(t *testing.T) {
	result := errcore.GetActualAndExpectSortedMessage("act", "exp")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectSortedMessage", actual)
}

func Test_Cov12_GetSearchLineNumberExpectationMessage(t *testing.T) {
	result := errcore.GetSearchLineNumberExpectationMessage(1, "line")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage", actual)
}

func Test_Cov12_GetSearchTermExpectationMessage(t *testing.T) {
	result := errcore.GetSearchTermExpectationMessage("act", "exp")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage", actual)
}

func Test_Cov12_GetSearchTermExpectationSimpleMessage(t *testing.T) {
	result := errcore.GetSearchTermExpectationSimpleMessage("act", "exp")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationSimpleMessage", actual)
}

func Test_Cov12_PathMeaningFulMessage(t *testing.T) {
	result := errcore.PathMeaningFulMessage("msg", "path")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningFulMessage", actual)
}

func Test_Cov12_PathMeaningfulError(t *testing.T) {
	err := errcore.PathMeaningfulError("msg", "path")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError", actual)
}

// ==========================================================================
// Panic / Range functions
// ==========================================================================

func Test_Cov12_PanicOnIndexOutOfRange(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.PanicOnIndexOutOfRange(-1, 10)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange", actual)
}

func Test_Cov12_PanicOnIndexOutOfRange_Valid(t *testing.T) {
	errcore.PanicOnIndexOutOfRange(0, 10) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange valid", actual)
}

func Test_Cov12_RangeNotMeet(t *testing.T) {
	result := errcore.RangeNotMeet(0, 10)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet", actual)
}

func Test_Cov12_EnumRangeNotMeet(t *testing.T) {
	result := errcore.EnumRangeNotMeet(0, 10, "1,2,3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet", actual)
}

func Test_Cov12_PanicRangeNotMeet(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.PanicRangeNotMeet(0, 10)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicRangeNotMeet", actual)
}

// ==========================================================================
// ManyErrorToSingle / ManyErrorToSingleDirect
// ==========================================================================

func Test_Cov12_ManyErrorToSingle(t *testing.T) {
	err := errcore.ManyErrorToSingle(errors.New("a"), errors.New("b"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle", actual)
}

func Test_Cov12_ManyErrorToSingleDirect(t *testing.T) {
	err := errcore.ManyErrorToSingleDirect(errors.New("a"), errors.New("b"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect", actual)
}

// ==========================================================================
// SourceDestination / SourceDestinationErr / SourceDestinationNoType
// ==========================================================================

func Test_Cov12_SourceDestination(t *testing.T) {
	result := errcore.SourceDestination("src", "sv", "dst", "dv")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination", actual)
}

func Test_Cov12_SourceDestinationErr(t *testing.T) {
	err := errcore.SourceDestinationErr("src", "sv", "dst", "dv")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr", actual)
}

func Test_Cov12_SourceDestinationNoType(t *testing.T) {
	result := errcore.SourceDestinationNoType("src", "sv", "dst", "dv")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType", actual)
}

// ==========================================================================
// CompiledError
// ==========================================================================

func Test_Cov12_CompiledError(t *testing.T) {
	result := errcore.CompiledError("msg", errors.New("e"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompiledError", actual)
}

// ==========================================================================
// ToExitError
// ==========================================================================

func Test_Cov12_ToExitError(t *testing.T) {
	err := errcore.ToExitError(errors.New("e"), 1)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ToExitError", actual)
}

// ==========================================================================
// ExpectationMessageDef
// ==========================================================================

func Test_Cov12_ExpectationMessageDef(t *testing.T) {
	result := errcore.ExpectationMessageDef("title", "expect", "actual")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef", actual)
}

// ==========================================================================
// HandleCompiledErrorWithTracesGetter / HandleFullStringsWithTracesGetter
// ==========================================================================

func Test_Cov12_HandleCompiledErrorWithTracesGetter(t *testing.T) {
	errcore.HandleCompiledErrorWithTracesGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorWithTracesGetter nil", actual)
}

func Test_Cov12_HandleFullStringsWithTracesGetter(t *testing.T) {
	errcore.HandleFullStringsWithTracesGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleFullStringsWithTracesGetter nil", actual)
}

// ==========================================================================
// getReferenceMessage
// ==========================================================================

func Test_Cov12_GetReferenceMessage(t *testing.T) {
	result := errcore.GetReferenceMessage("ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetReferenceMessage", actual)
}
