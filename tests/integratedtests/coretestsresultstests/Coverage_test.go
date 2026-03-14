package coretestsresultstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── Result methods ──

func Test_Cov_Result_IsSafe(t *testing.T) {
	safe := results.ResultAny{Value: "ok"}
	unsafe := results.ResultAny{Panicked: true}
	errResult := results.ResultAny{Error: errors.New("e")}
	actual := args.Map{
		"safe":   safe.IsSafe(),
		"unsafe": unsafe.IsSafe(),
		"errSafe": errResult.IsSafe(),
	}
	expected := args.Map{
		"safe":   true,
		"unsafe": false,
		"errSafe": false,
	}
	expected.ShouldBeEqual(t, 0, "IsSafe", actual)
}

func Test_Cov_Result_HasError(t *testing.T) {
	r := results.ResultAny{Error: errors.New("e")}
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "HasError", actual)
}

func Test_Cov_Result_HasPanicked(t *testing.T) {
	r := results.ResultAny{Panicked: true}
	actual := args.Map{"panicked": r.HasPanicked()}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HasPanicked", actual)
}

func Test_Cov_Result_IsResult(t *testing.T) {
	r := results.Result[int]{Value: 42}
	actual := args.Map{"match": r.IsResult(42), "noMatch": r.IsResult(99)}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsResult", actual)
}

func Test_Cov_Result_IsResultTypeOf(t *testing.T) {
	r := results.Result[int]{Value: 42}
	actual := args.Map{
		"intType":  r.IsResultTypeOf(0),
		"nilType":  r.IsResultTypeOf(nil),
	}
	expected := args.Map{
		"intType":  true,
		"nilType":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf", actual)
}

func Test_Cov_Result_IsError(t *testing.T) {
	r := results.ResultAny{Error: errors.New("test")}
	noErr := results.ResultAny{}
	actual := args.Map{
		"match":   r.IsError("test"),
		"noMatch": r.IsError("other"),
		"noErr":   noErr.IsError("test"),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
		"noErr":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsError", actual)
}

func Test_Cov_Result_ValueString(t *testing.T) {
	r := results.Result[int]{Value: 42}
	actual := args.Map{"val": r.ValueString()}
	expected := args.Map{"val": "42"}
	expected.ShouldBeEqual(t, 0, "ValueString", actual)
}

func Test_Cov_Result_ResultAt(t *testing.T) {
	r := results.ResultAny{AllResults: []any{"a", "b"}}
	actual := args.Map{
		"first":  r.ResultAt(0),
		"second": r.ResultAt(1),
		"outOfBounds": r.ResultAt(5) == nil,
		"negative":    r.ResultAt(-1) == nil,
	}
	expected := args.Map{
		"first":  "a",
		"second": "b",
		"outOfBounds": true,
		"negative":    true,
	}
	expected.ShouldBeEqual(t, 0, "ResultAt", actual)
}

func Test_Cov_Result_ToMap(t *testing.T) {
	r := results.Result[int]{Value: 42, ReturnCount: 1}
	m := r.ToMap()
	actual := args.Map{
		"value":       m["value"],
		"panicked":    m["panicked"],
		"returnCount": m["returnCount"],
	}
	expected := args.Map{
		"value":       "42",
		"panicked":    false,
		"returnCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "ToMap", actual)
}

func Test_Cov_Result_ToMapCompact(t *testing.T) {
	r := results.Result[int]{Value: 42}
	m := r.ToMapCompact()
	actual := args.Map{"value": m["value"], "panicked": m["panicked"]}
	expected := args.Map{"value": "42", "panicked": false}
	expected.ShouldBeEqual(t, 0, "ToMapCompact", actual)
}

func Test_Cov_Result_String(t *testing.T) {
	rNormal := results.Result[int]{Value: 42, ReturnCount: 1}
	rPanic := results.Result[int]{Panicked: true, PanicValue: "boom"}
	rErr := results.Result[int]{Value: 0, Error: errors.New("e"), ReturnCount: 1}
	actual := args.Map{
		"normalNotEmpty": rNormal.String() != "",
		"panicNotEmpty":  rPanic.String() != "",
		"errNotEmpty":    rErr.String() != "",
	}
	expected := args.Map{
		"normalNotEmpty": true,
		"panicNotEmpty":  true,
		"errNotEmpty":    true,
	}
	expected.ShouldBeEqual(t, 0, "String", actual)
}

// ── Results (two-value) ──

func Test_Cov_Results_String(t *testing.T) {
	rNormal := results.Results[int, bool]{Result: results.Result[int]{Value: 42}, Result2: true}
	rPanic := results.Results[int, bool]{Result: results.Result[int]{Panicked: true, PanicValue: "boom"}}
	rErr := results.Results[int, bool]{Result: results.Result[int]{Error: errors.New("e")}, Result2: false}
	actual := args.Map{
		"normalNotEmpty": rNormal.String() != "",
		"panicNotEmpty":  rPanic.String() != "",
		"errNotEmpty":    rErr.String() != "",
	}
	expected := args.Map{
		"normalNotEmpty": true,
		"panicNotEmpty":  true,
		"errNotEmpty":    true,
	}
	expected.ShouldBeEqual(t, 0, "Results_String", actual)
}

func Test_Cov_Results_IsResult2(t *testing.T) {
	r := results.Results[int, string]{Result2: "hello"}
	actual := args.Map{"match": r.IsResult2("hello"), "noMatch": r.IsResult2("other")}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsResult2", actual)
}

func Test_Cov_Results_Result2String(t *testing.T) {
	r := results.Results[int, string]{Result2: "hello"}
	actual := args.Map{"val": r.Result2String()}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Result2String", actual)
}

func Test_Cov_FromResultAny(t *testing.T) {
	raw := results.ResultAny{
		Value:       42,
		AllResults:  []any{42, "hello"},
		ReturnCount: 2,
	}
	r := results.FromResultAny[int, string](raw)
	actual := args.Map{"val1": r.Value, "val2": r.Result2, "count": r.ReturnCount}
	expected := args.Map{"val1": 42, "val2": "hello", "count": 2}
	expected.ShouldBeEqual(t, 0, "FromResultAny", actual)
}

func Test_Cov_FromResultAny_TypeMismatch(t *testing.T) {
	raw := results.ResultAny{AllResults: []any{"not-int", 42}}
	r := results.FromResultAny[int, string](raw)
	actual := args.Map{"val1": r.Value, "val2": r.Result2}
	expected := args.Map{"val1": 0, "val2": ""}
	expected.ShouldBeEqual(t, 0, "FromResultAny_TypeMismatch", actual)
}

// ── InvokeWithPanicRecovery ──

func Test_Cov_InvokeWithPanicRecovery_NilFunc(t *testing.T) {
	result := results.InvokeWithPanicRecovery(nil, nil)
	actual := args.Map{"panicked": result.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "NilFunc", actual)
}

func Test_Cov_InvokeWithPanicRecovery_NotFunc(t *testing.T) {
	result := results.InvokeWithPanicRecovery(42, nil)
	actual := args.Map{"panicked": result.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "NotFunc", actual)
}

func Test_Cov_InvokeWithPanicRecovery_SimpleFunc(t *testing.T) {
	fn := func(x int) int { return x * 2 }
	result := results.InvokeWithPanicRecovery(fn, nil, 5)
	// buildCallArgs creates a zero int for nil receiver + arg 5 = 2 args for 1-param func → panic
	actual := args.Map{"panicked": result.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "SimpleFunc", actual)
}

func Test_Cov_InvokeWithPanicRecovery_VoidFunc(t *testing.T) {
	called := false
	fn := func() { called = true }
	result := results.InvokeWithPanicRecovery(fn, nil)
	actual := args.Map{"panicked": result.Panicked, "called": called, "count": result.ReturnCount}
	expected := args.Map{"panicked": false, "called": true, "count": 0}
	expected.ShouldBeEqual(t, 0, "VoidFunc", actual)
}

func Test_Cov_InvokeWithPanicRecovery_PanicFunc(t *testing.T) {
	fn := func() { panic("boom") }
	result := results.InvokeWithPanicRecovery(fn, nil)
	actual := args.Map{"panicked": result.Panicked, "panicVal": fmt.Sprintf("%v", result.PanicValue)}
	expected := args.Map{"panicked": true, "panicVal": "boom"}
	expected.ShouldBeEqual(t, 0, "PanicFunc", actual)
}

// ── MethodName ──

func Test_Cov_MethodName_Nil(t *testing.T) {
	result := results.MethodName(nil)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "MethodName_Nil", actual)
}

func Test_Cov_MethodName_NotFunc(t *testing.T) {
	result := results.MethodName(42)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "MethodName_NotFunc", actual)
}

func Test_Cov_MethodName_Func(t *testing.T) {
	result := results.MethodName(fmt.Sprintf)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MethodName_Func", actual)
}

// ── ResultAssert (ShouldMatchResult) ──

func Test_Cov_ShouldMatchResult_Basic(t *testing.T) {
	r := results.Result[int]{Value: 42, Panicked: false, ReturnCount: 1}
	exp := results.ResultAny{Value: "42", Panicked: false, ReturnCount: 1}
	// Should not fail
	r.ShouldMatchResult(t, 0, "basic", exp)
}

func Test_Cov_ShouldMatchResult_ExplicitFields(t *testing.T) {
	r := results.Result[int]{Value: 42, Panicked: false}
	exp := results.ResultAny{Panicked: false}
	r.ShouldMatchResult(t, 0, "explicit", exp, "panicked")
}
