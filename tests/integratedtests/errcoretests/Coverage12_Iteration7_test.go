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
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqMsg returns non-empty -- different strings", actual)
}

func Test_Cov12_ShouldBe_StrEqErr(t *testing.T) {
	err := errcore.ShouldBe.StrEqErr("actual", "expect")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.StrEqErr returns error -- different strings", actual)
}

func Test_Cov12_ShouldBe_AnyEqMsg(t *testing.T) {
	msg := errcore.ShouldBe.AnyEqMsg(1, 2)
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqMsg returns non-empty -- different values", actual)
}

func Test_Cov12_ShouldBe_AnyEqErr(t *testing.T) {
	err := errcore.ShouldBe.AnyEqErr(1, 2)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.AnyEqErr returns error -- different values", actual)
}

func Test_Cov12_ShouldBe_JsonEqMsg(t *testing.T) {
	msg := errcore.ShouldBe.JsonEqMsg("a", "b")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqMsg returns non-empty -- different json", actual)
}

func Test_Cov12_ShouldBe_JsonEqErr(t *testing.T) {
	err := errcore.ShouldBe.JsonEqErr("a", "b")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ShouldBe.JsonEqErr returns error -- different json", actual)
}

// ==========================================================================
// expected — all methods
// ==========================================================================

func Test_Cov12_Expected_But(t *testing.T) {
	err := errcore.Expected.But("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.But returns error -- with args", actual)
}

func Test_Cov12_Expected_ButFoundAsMsg(t *testing.T) {
	msg := errcore.Expected.ButFoundAsMsg("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundAsMsg returns non-empty -- with args", actual)
}

func Test_Cov12_Expected_ButFoundWithTypeAsMsg(t *testing.T) {
	msg := errcore.Expected.ButFoundWithTypeAsMsg("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButFoundWithTypeAsMsg returns non-empty -- with args", actual)
}

func Test_Cov12_Expected_ButUsingType(t *testing.T) {
	err := errcore.Expected.ButUsingType("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ButUsingType returns error -- with args", actual)
}

func Test_Cov12_Expected_ReflectButFound(t *testing.T) {
	err := errcore.Expected.ReflectButFound(reflect.String, reflect.Int)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ReflectButFound returns error -- different kinds", actual)
}

func Test_Cov12_Expected_PrimitiveButFound(t *testing.T) {
	err := errcore.Expected.PrimitiveButFound(reflect.Struct)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.PrimitiveButFound returns error -- non-primitive kind", actual)
}

func Test_Cov12_Expected_ValueHasNoElements(t *testing.T) {
	err := errcore.Expected.ValueHasNoElements(reflect.Slice)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Expected.ValueHasNoElements returns error -- with kind", actual)
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
	expected.ShouldBeEqual(t, 0, "CountStateChangeTracker returns same -- no changes", actual)

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
	expected2.ShouldBeEqual(t, 1, "CountStateChangeTracker returns changed -- length increased", actual2)
}

// ==========================================================================
// StackTracesCompiled
// ==========================================================================

func Test_Cov12_StackTracesCompiled(t *testing.T) {
	result := errcore.StackTracesCompiled([]string{"trace1", "trace2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StackTracesCompiled returns non-empty -- with traces", actual)
}

// ==========================================================================
// stackTraceEnhance — all methods
// ==========================================================================

func Test_Cov12_StackEnhance_Error(t *testing.T) {
	err := errcore.StackEnhance.Error(errors.New("e"))
	errNil := errcore.StackEnhance.Error(nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Error returns error -- with error", actual)
}

func Test_Cov12_StackEnhance_ErrorSkip(t *testing.T) {
	err := errcore.StackEnhance.ErrorSkip(0, errors.New("e"))
	errNil := errcore.StackEnhance.ErrorSkip(0, nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.ErrorSkip returns error -- with error", actual)
}

func Test_Cov12_StackEnhance_MsgToErrSkip(t *testing.T) {
	err := errcore.StackEnhance.MsgToErrSkip(0, "msg")
	errEmpty := errcore.StackEnhance.MsgToErrSkip(0, "")
	actual := args.Map{"notNil": err != nil, "nil": errEmpty == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgToErrSkip returns error -- with message", actual)
}

func Test_Cov12_StackEnhance_FmtSkip(t *testing.T) {
	err := errcore.StackEnhance.FmtSkip(0, "hello %d", 1)
	errEmpty := errcore.StackEnhance.FmtSkip(0, "")
	actual := args.Map{"notNil": err != nil, "nil": errEmpty == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.FmtSkip returns error -- with format", actual)
}

func Test_Cov12_StackEnhance_Msg(t *testing.T) {
	msg := errcore.StackEnhance.Msg("hello")
	msgEmpty := errcore.StackEnhance.Msg("")
	actual := args.Map{"notEmpty": msg != "", "empty": msgEmpty == ""}
	expected := args.Map{"notEmpty": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.Msg returns non-empty -- with message", actual)
}

func Test_Cov12_StackEnhance_MsgErrorSkip(t *testing.T) {
	msg := errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("e"))
	msgNil := errcore.StackEnhance.MsgErrorSkip(0, "msg", nil)
	actual := args.Map{"notEmpty": msg != "", "empty": msgNil == ""}
	expected := args.Map{"notEmpty": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorSkip returns non-empty -- with error", actual)
}

func Test_Cov12_StackEnhance_MsgErrorToErrSkip(t *testing.T) {
	err := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e"))
	errNil := errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "StackEnhance.MsgErrorToErrSkip returns error -- with error", actual)
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
	expected.ShouldBeEqual(t, 0, "ExpectingFuture returns record -- with title", actual)
}

func Test_Cov12_ExpectingRecord_Message(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.Message("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Message returns non-empty -- with actual", actual)
}

func Test_Cov12_ExpectingRecord_MessageSimple(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.MessageSimple("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimple returns non-empty -- with actual", actual)
}

func Test_Cov12_ExpectingRecord_MessageSimpleNoType(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	msg := er.MessageSimpleNoType("actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.MessageSimpleNoType returns non-empty -- with actual", actual)
}

func Test_Cov12_ExpectingRecord_Error(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	err := er.Error("actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.Error returns error -- with actual", actual)
}

func Test_Cov12_ExpectingRecord_ErrorSimple(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	err := er.ErrorSimple("actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimple returns error -- with actual", actual)
}

func Test_Cov12_ExpectingRecord_ErrorSimpleNoType(t *testing.T) {
	er := errcore.ExpectingFuture("title", "expect")
	err := er.ErrorSimpleNoType("actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingRecord.ErrorSimpleNoType returns error -- with actual", actual)
}

// ==========================================================================
// Expecting / ExpectingSimple / ExpectingSimpleNoType / ExpectingError*
// ==========================================================================

func Test_Cov12_Expecting(t *testing.T) {
	msg := errcore.Expecting("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Expecting returns formatted -- with args", actual)
}

func Test_Cov12_ExpectingSimple(t *testing.T) {
	msg := errcore.ExpectingSimple("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimple returns formatted -- with args", actual)
}

func Test_Cov12_ExpectingSimpleNoType(t *testing.T) {
	msg := errcore.ExpectingSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingSimpleNoType returns formatted -- with args", actual)
}

func Test_Cov12_ExpectingErrorSimpleNoType(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- with args", actual)
}

func Test_Cov12_ExpectingNotEqualSimpleNoType(t *testing.T) {
	msg := errcore.ExpectingNotEqualSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectingNotEqualSimpleNoType returns non-empty -- with args", actual)
}

func Test_Cov12_ExpectingError(t *testing.T) {
	err := errcore.ExpectingErrorSimpleNoType("title", "expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectingErrorSimpleNoType returns error -- with args", actual)
}

// ==========================================================================
// RawErrorType — remaining uncovered methods
// ==========================================================================

func Test_Cov12_RawErrorType_CombineWithAnother(t *testing.T) {
	result := errcore.InvalidType.CombineWithAnother(errcore.NotFound, "msg", "ref")
	actual := args.Map{"notEmpty": string(result) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.CombineWithAnother returns non-empty -- with another type", actual)
}

func Test_Cov12_RawErrorType_TypesAttach(t *testing.T) {
	result := errcore.InvalidType.TypesAttach("msg", "str", 42)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttach returns non-empty -- with types", actual)
}

func Test_Cov12_RawErrorType_TypesAttachErr(t *testing.T) {
	err := errcore.InvalidType.TypesAttachErr("msg", "str", 42)
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.TypesAttachErr returns error -- with types", actual)
}

func Test_Cov12_RawErrorType_SrcDestination(t *testing.T) {
	result := errcore.InvalidType.SrcDestination("msg", "src", "sv", "dst", "dv")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestination returns formatted -- with args", actual)
}

func Test_Cov12_RawErrorType_SrcDestinationErr(t *testing.T) {
	err := errcore.InvalidType.SrcDestinationErr("msg", "src", "sv", "dst", "dv")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.SrcDestinationErr returns error -- with args", actual)
}

func Test_Cov12_RawErrorType_Error(t *testing.T) {
	err := errcore.InvalidType.Error("msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Error returns error -- with msg and ref", actual)
}

func Test_Cov12_RawErrorType_ErrorSkip(t *testing.T) {
	err := errcore.InvalidType.ErrorSkip(0, "msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorSkip returns error -- with skip", actual)
}

func Test_Cov12_RawErrorType_Fmt(t *testing.T) {
	err := errcore.InvalidType.Fmt("hello %d", 42)
	errEmpty := errcore.InvalidType.Fmt("")
	actual := args.Map{"notNil": err != nil, "emptyNotNil": errEmpty != nil}
	expected := args.Map{"notNil": true, "emptyNotNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Fmt returns error -- with format", actual)
}

func Test_Cov12_RawErrorType_FmtIf(t *testing.T) {
	err := errcore.InvalidType.FmtIf(true, "hello %d", 42)
	errNil := errcore.InvalidType.FmtIf(false, "hello %d", 42)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.FmtIf returns correct value -- with condition", actual)
}

func Test_Cov12_RawErrorType_MergeError(t *testing.T) {
	err := errcore.InvalidType.MergeError(errors.New("e"))
	errNil := errcore.InvalidType.MergeError(nil)
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeError returns error -- with error", actual)
}

func Test_Cov12_RawErrorType_MergeErrorWithMessage(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessage(errors.New("e"), "msg")
	errNil := errcore.InvalidType.MergeErrorWithMessage(nil, "msg")
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessage returns error -- with error", actual)
}

func Test_Cov12_RawErrorType_MergeErrorWithMessageRef(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref")
	errNil := errcore.InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref")
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithMessageRef returns error -- with error", actual)
}

func Test_Cov12_RawErrorType_MergeErrorWithRef(t *testing.T) {
	err := errcore.InvalidType.MergeErrorWithRef(errors.New("e"), "ref")
	errNil := errcore.InvalidType.MergeErrorWithRef(nil, "ref")
	actual := args.Map{"notNil": err != nil, "nil": errNil == nil}
	expected := args.Map{"notNil": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MergeErrorWithRef returns error -- with error", actual)
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
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRef returns non-empty -- with items", actual)
}

func Test_Cov12_RawErrorType_MsgCsvRefError(t *testing.T) {
	err := errcore.InvalidType.MsgCsvRefError("msg", "a", "b")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.MsgCsvRefError returns error -- with items", actual)
}

func Test_Cov12_RawErrorType_ErrorRefOnly(t *testing.T) {
	err := errcore.InvalidType.ErrorRefOnly("ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorRefOnly returns error -- with ref", actual)
}

func Test_Cov12_RawErrorType_Expecting(t *testing.T) {
	err := errcore.InvalidType.Expecting("expect", "actual")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.Expecting returns error -- with args", actual)
}

func Test_Cov12_RawErrorType_NoRef(t *testing.T) {
	result := errcore.InvalidType.NoRef("msg")
	resultEmpty := errcore.InvalidType.NoRef("")
	actual := args.Map{"notEmpty": result != "", "emptyNotEmpty": resultEmpty != ""}
	expected := args.Map{"notEmpty": true, "emptyNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.NoRef returns non-empty -- with msg", actual)
}

func Test_Cov12_RawErrorType_ErrorNoRefs(t *testing.T) {
	err := errcore.InvalidType.ErrorNoRefs("msg")
	errEmpty := errcore.InvalidType.ErrorNoRefs("")
	actual := args.Map{"notNil": err != nil, "emptyNotNil": errEmpty != nil}
	expected := args.Map{"notNil": true, "emptyNotNil": true}
	expected.ShouldBeEqual(t, 0, "RawErrorType.ErrorNoRefs returns error -- with msg", actual)
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
	expected.ShouldBeEqual(t, 0, "RawErrorType.HandleUsingPanic panics -- with error", actual)
}

func Test_Cov12_GetSet(t *testing.T) {
	r1 := errcore.GetSet(true, errcore.InvalidType, errcore.NotFound)
	r2 := errcore.GetSet(false, errcore.InvalidType, errcore.NotFound)
	actual := args.Map{"r1": r1, "r2": r2}
	expected := args.Map{"r1": errcore.InvalidType, "r2": errcore.NotFound}
	expected.ShouldBeEqual(t, 0, "GetSet returns correct value -- with condition", actual)
}

func Test_Cov12_GetSetVariant(t *testing.T) {
	r1 := errcore.GetSetVariant(true, "a", "b")
	r2 := errcore.GetSetVariant(false, "a", "b")
	actual := args.Map{"r1": string(r1), "r2": string(r2)}
	expected := args.Map{"r1": "a", "r2": "b"}
	expected.ShouldBeEqual(t, 0, "GetSetVariant returns correct value -- with condition", actual)
}

// ==========================================================================
// GherkinsString / GherkinsStringWithExpectation
// ==========================================================================

func Test_Cov12_GherkinsString(t *testing.T) {
	result := errcore.GherkinsString(0, "feature", "given", "when", "then")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsString returns non-empty -- with args", actual)
}

func Test_Cov12_GherkinsStringWithExpectation(t *testing.T) {
	result := errcore.GherkinsStringWithExpectation(0, "feature", "given", "when", "then", "actual", "expect")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GherkinsStringWithExpectation returns non-empty -- with args", actual)
}

// ==========================================================================
// Message formatting functions
// ==========================================================================

func Test_Cov12_MessageNameValues(t *testing.T) {
	result := errcore.MessageNameValues("msg", namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageNameValues returns formatted -- with name-values", actual)
}

func Test_Cov12_MessageVarTwo(t *testing.T) {
	result := errcore.MessageVarTwo("msg", "n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarTwo returns formatted -- with args", actual)
}

func Test_Cov12_MessageVarThree(t *testing.T) {
	result := errcore.MessageVarThree("msg", "n1", "v1", "n2", "v2", "n3", "v3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarThree returns formatted -- with args", actual)
}

func Test_Cov12_MessageVarMap(t *testing.T) {
	result := errcore.MessageVarMap("msg", map[string]any{"k": "v"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageVarMap returns formatted -- with map", actual)
}

func Test_Cov12_MessageWithRef(t *testing.T) {
	result := errcore.MessageWithRef("msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRef returns non-empty -- with args", actual)
}

func Test_Cov12_MessageWithRefToError(t *testing.T) {
	err := errcore.MessageWithRefToError("msg", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MessageWithRefToError returns error -- with args", actual)
}

func Test_Cov12_VarTwo(t *testing.T) {
	result := errcore.VarTwo(true, "n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwo returns formatted -- with args", actual)
}

func Test_Cov12_VarTwoNoType(t *testing.T) {
	result := errcore.VarTwo(false, "n1", "v1", "n2", "v2")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarTwoNoType returns formatted -- with args", actual)
}

func Test_Cov12_VarThree(t *testing.T) {
	result := errcore.VarThree(true, "n1", "v1", "n2", "v2", "n3", "v3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThree returns formatted -- with args", actual)
}

func Test_Cov12_VarThreeNoType(t *testing.T) {
	result := errcore.VarThree(false, "n1", "v1", "n2", "v2", "n3", "v3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarThreeNoType returns formatted -- with args", actual)
}

func Test_Cov12_VarNameValues(t *testing.T) {
	result := errcore.VarNameValues(namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValues returns formatted -- with args", actual)
}

func Test_Cov12_VarNameValuesJoiner(t *testing.T) {
	result := errcore.VarNameValuesJoiner(",", namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VarNameValuesJoiner returns joined -- with name-values", actual)
}

func Test_Cov12_VarNameValuesStrings(t *testing.T) {
	result := errcore.VarNameValuesStrings(namevalue.StringAny{Name: "n1", Value: "v1"}, namevalue.StringAny{Name: "n2", Value: "v2"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "VarNameValuesStrings returns entries -- with name-values", actual)
}

// ==========================================================================
// ErrorWith* and Handle* functions
// ==========================================================================

func Test_Cov12_ErrorWithRef(t *testing.T) {
	result := errcore.ErrorWithRef(errors.New("e"), "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRef returns formatted -- with error and ref", actual)
}

func Test_Cov12_ErrorWithRefToError(t *testing.T) {
	err := errcore.ErrorWithRefToError(errors.New("e"), "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithRefToError returns error -- with error", actual)
}

func Test_Cov12_ErrorWithCompiledTraceRef(t *testing.T) {
	result := errcore.ErrorWithCompiledTraceRef(errors.New("e"), "trace", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRef returns non-empty -- with all args", actual)
}

func Test_Cov12_ErrorWithCompiledTraceRefToError(t *testing.T) {
	err := errcore.ErrorWithCompiledTraceRefToError(errors.New("e"), "trace", "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithCompiledTraceRefToError returns error -- with args", actual)
}

func Test_Cov12_ErrorWithTracesRefToError(t *testing.T) {
	err := errcore.ErrorWithTracesRefToError(errors.New("e"), []string{"t"}, "ref")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorWithTracesRefToError returns error -- with traces", actual)
}

func Test_Cov12_HandleErr(t *testing.T) {
	errcore.HandleErr(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErr completes safely -- nil error", actual)
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
	expected.ShouldBeEqual(t, 0, "HandleErr panics -- with error", actual)
}

func Test_Cov12_HandleErrMessage(t *testing.T) {
	errcore.HandleErrMessage("") // no panic for empty string
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrMessage completes safely -- empty message", actual)
}

func Test_Cov12_SimpleHandleErr(t *testing.T) {
	errcore.SimpleHandleErr(nil, "msg") // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErr completes safely -- nil error", actual)
}

func Test_Cov12_SimpleHandleErrMany(t *testing.T) {
	errcore.SimpleHandleErrMany("msg") // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SimpleHandleErrMany completes safely -- nil errors", actual)
}

func Test_Cov12_HandleCompiledErrorGetter(t *testing.T) {
	errcore.HandleCompiledErrorGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorGetter completes safely -- nil getter", actual)
}

func Test_Cov12_HandleErrorGetter(t *testing.T) {
	errcore.HandleErrorGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorGetter completes safely -- nil getter", actual)
}

// ==========================================================================
// CombineWithMsgType variants
// ==========================================================================

func Test_Cov12_CombineWithMsgTypeNoStack(t *testing.T) {
	result := errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeNoStack returns non-empty -- with args", actual)
}

func Test_Cov12_CombineWithMsgTypeStackTrace(t *testing.T) {
	result := errcore.CombineWithMsgTypeStackTrace(errcore.InvalidType, "msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CombineWithMsgTypeStackTrace returns non-empty -- with stack trace", actual)
}

func Test_Cov12_Combine_Func(t *testing.T) {
	result := errcore.Combine("errType", "msg", "ref")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Combine returns formatted -- with args", actual)
}

// ==========================================================================
// MeaningFulError / MeaningfulMessageError
// ==========================================================================

func Test_Cov12_MeaningfulError(t *testing.T) {
	err := errcore.MeaningfulError(errcore.InvalidType, "fn", errors.New("e"))
	errNil := errcore.MeaningfulError(errcore.InvalidType, "fn", nil)
	actual := args.Map{"notNil": err != nil, "nilNil": errNil == nil}
	expected := args.Map{"notNil": true, "nilNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns error -- with error", actual)
}

func Test_Cov12_MeaningfulErrorHandle(t *testing.T) {
	errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorHandle completes safely -- nil error", actual)
}

func Test_Cov12_MeaningfulErrorWithData(t *testing.T) {
	err := errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", errors.New("e"), "data")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorWithData returns error -- with error", actual)
}

func Test_Cov12_MeaningfulMessageError(t *testing.T) {
	err := errcore.MeaningfulMessageError(errcore.InvalidType, "fn", errors.New("e"), "msg")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulMessageError returns error -- with error", actual)
}

// ==========================================================================
// MsgHeader / MsgHeaderIf / MsgHeaderPlusEnding
// ==========================================================================

func Test_Cov12_MsgHeader(t *testing.T) {
	result := errcore.MsgHeader("msg")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeader returns non-empty -- with args", actual)
}

func Test_Cov12_MsgHeaderIf(t *testing.T) {
	result := errcore.MsgHeaderIf(true, "msg")
	resultFalse := errcore.MsgHeaderIf(false, "msg")
	actual := args.Map{"notEmpty": result != "", "empty": resultFalse == ""}
	expected := args.Map{"notEmpty": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderIf returns correct value -- with condition", actual)
}

func Test_Cov12_MsgHeaderPlusEnding(t *testing.T) {
	result := errcore.MsgHeaderPlusEnding("msg", "ending")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgHeaderPlusEnding returns non-empty -- with args", actual)
}

// ==========================================================================
// StringLines functions
// ==========================================================================

func Test_Cov12_StringLinesToQuoteLines(t *testing.T) {
	result := errcore.StringLinesToQuoteLines([]string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns formatted -- with input", actual)
}

func Test_Cov12_StringLinesToQuoteLines_Integrated(t *testing.T) {
	result := errcore.StringLinesToQuoteLines([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLines returns formatted -- integrated test", actual)
}

func Test_Cov12_StringLinesToQuoteLinesToSingle(t *testing.T) {
	result := errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringLinesToQuoteLinesToSingle returns non-empty -- with input", actual)
}

// ==========================================================================
// Print / FmtDebug / FmtDebugIf
// ==========================================================================

func Test_Cov12_PrintError(t *testing.T) {
	errcore.PrintError(nil)
	errcore.PrintError(errors.New("e"))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintError completes safely -- with error", actual)
}

func Test_Cov12_PrintErrorWithTestIndex(t *testing.T) {
	errcore.PrintErrorWithTestIndex(0, "test", errors.New("e"))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PrintErrorWithTestIndex completes safely -- with error", actual)
}

func Test_Cov12_FmtDebug(t *testing.T) {
	errcore.FmtDebug("msg")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebug completes safely -- with format", actual)
}

func Test_Cov12_FmtDebugIf(t *testing.T) {
	errcore.FmtDebugIf(true, "msg")
	errcore.FmtDebugIf(false, "msg")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FmtDebugIf completes safely -- with condition", actual)
}

// ==========================================================================
// GetActualAndExpect / GetSearchLine / PathMeaningful
// ==========================================================================

func Test_Cov12_GetActualAndExpectProcessedMessage(t *testing.T) {
	result := errcore.GetActualAndExpectProcessedMessage(1, "act", "exp", "act-processed", "exp-processed")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectProcessedMessage returns non-empty -- with args", actual)
}

func Test_Cov12_GetActualAndExpectSortedMessage(t *testing.T) {
	result := errcore.GetActualAndExpectProcessedMessage(2, []string{"b", "a"}, []string{"a", "b"}, []string{"a", "b"}, []string{"a", "b"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetActualAndExpectSortedMessage returns non-empty -- with args", actual)
}

func Test_Cov12_GetSearchLineNumberExpectationMessage(t *testing.T) {
	result := errcore.GetSearchLineNumberExpectationMessage(1, 10, 9, "line-content", "term", "extra")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchLineNumberExpectationMessage returns non-empty -- with args", actual)
}

func Test_Cov12_GetSearchTermExpectationMessage(t *testing.T) {
	result := errcore.GetSearchTermExpectationMessage(1, "header", "expectation", 0, "act", "exp", nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationMessage returns non-empty -- with args", actual)
}

func Test_Cov12_GetSearchTermExpectationSimpleMessage(t *testing.T) {
	result := errcore.GetSearchTermExpectationSimpleMessage(1, "expectation", 0, "act", "exp")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSearchTermExpectationSimpleMessage returns non-empty -- with args", actual)
}

func Test_Cov12_PathMeaningfulMessage(t *testing.T) {
	err := errcore.PathMeaningfulMessage(errcore.InvalidType, "fn", "path", "msg")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulMessage returns error -- with messages", actual)
}

func Test_Cov12_PathMeaningfulError(t *testing.T) {
	err := errcore.PathMeaningfulError(errcore.InvalidType, errors.New("boom"), "path")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "PathMeaningfulError returns error -- with error", actual)
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
		errcore.PanicOnIndexOutOfRange(-1, []int{10})
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange panics -- index out of range", actual)
}

func Test_Cov12_PanicOnIndexOutOfRange_Valid(t *testing.T) {
	errcore.PanicOnIndexOutOfRange(10, []int{0}) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "PanicOnIndexOutOfRange completes safely -- index in range", actual)
}

func Test_Cov12_RangeNotMeet(t *testing.T) {
	result := errcore.RangeNotMeet("test", 0, 10, nil)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNotMeet returns non-empty -- with range", actual)
}

func Test_Cov12_EnumRangeNotMeet(t *testing.T) {
	result := errcore.EnumRangeNotMeet(0, 10, "1,2,3")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "EnumRangeNotMeet returns non-empty -- with range", actual)
}

func Test_Cov12_PanicRangeNotMeet(t *testing.T) {
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		errcore.PanicRangeNotMeet("test", 0, 10, nil)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "PanicRangeNotMeet panics -- with range", actual)
}

// ==========================================================================
// ManyErrorToSingle / ManyErrorToSingleDirect
// ==========================================================================

func Test_Cov12_ManyErrorToSingle(t *testing.T) {
	err := errcore.ManyErrorToSingle([]error{errors.New("a"), errors.New("b")})
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingle returns error -- with errors", actual)
}

func Test_Cov12_ManyErrorToSingleDirect(t *testing.T) {
	err := errcore.ManyErrorToSingleDirect(errors.New("a"), errors.New("b"))
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ManyErrorToSingleDirect returns error -- with errors", actual)
}

// ==========================================================================
// SourceDestination / SourceDestinationErr / SourceDestinationNoType
// ==========================================================================

func Test_Cov12_SourceDestination(t *testing.T) {
	result := errcore.SourceDestination(true, "sv", "dv")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestination returns formatted -- with args", actual)
}

func Test_Cov12_SourceDestinationErr(t *testing.T) {
	err := errcore.SourceDestinationErr(true, "sv", "dv")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationErr returns error -- with args", actual)
}

func Test_Cov12_SourceDestinationNoType(t *testing.T) {
	result := errcore.SourceDestinationNoType("sv", "dv")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SourceDestinationNoType returns formatted -- with args", actual)
}

// ==========================================================================
// CompiledError
// ==========================================================================

func Test_Cov12_CompiledError(t *testing.T) {
	result := errcore.CompiledError(errors.New("main"), "additional")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompiledError returns error -- with message", actual)
}

// ==========================================================================
// ToExitError
// ==========================================================================

func Test_Cov12_ToExitError(t *testing.T) {
	err := errcore.ToExitError(errors.New("e"))
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true} // non-ExitError returns nil
	expected.ShouldBeEqual(t, 0, "ToExitError returns correct value -- with error type", actual)
}

// ==========================================================================
// ExpectationMessageDef
// ==========================================================================

func Test_Cov12_ExpectationMessageDef(t *testing.T) {
	def := errcore.ExpectationMessageDef{
		CaseIndex: 1,
		FuncName:  "TestFunc",
		When:      "when testing",
		Expected:  "expected-value",
	}
	result := def.ToString("actual-value")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ExpectationMessageDef returns non-empty -- with struct", actual)
}

// ==========================================================================
// HandleCompiledErrorWithTracesGetter / HandleFullStringsWithTracesGetter
// ==========================================================================

func Test_Cov12_HandleCompiledErrorWithTracesGetter(t *testing.T) {
	errcore.HandleCompiledErrorWithTracesGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleCompiledErrorWithTracesGetter completes safely -- nil getter", actual)
}

func Test_Cov12_HandleFullStringsWithTracesGetter(t *testing.T) {
	errcore.HandleFullStringsWithTracesGetter(nil) // no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleFullStringsWithTracesGetter completes safely -- nil getter", actual)
}

// ==========================================================================
// ReferenceStart / ReferenceEnd constants
// ==========================================================================

func Test_Cov12_ReferenceConstants(t *testing.T) {
	actual := args.Map{
		"startNotEmpty": errcore.ReferenceStart != "",
		"endNotEmpty":   errcore.ReferenceEnd != "",
	}
	expected := args.Map{
		"startNotEmpty": true,
		"endNotEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "ReferenceConstants returns non-empty -- defined constants", actual)
}
