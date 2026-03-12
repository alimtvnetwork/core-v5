package resultstests

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
	panicked := results.ResultAny{Panicked: true}
	errored := results.ResultAny{Error: errors.New("err")}

	actual := args.Map{
		"safeSafe":      safe.IsSafe(),
		"panickedSafe":  panicked.IsSafe(),
		"erroredSafe":   errored.IsSafe(),
	}
	expected := args.Map{
		"safeSafe":      true,
		"panickedSafe":  false,
		"erroredSafe":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsSafe", actual)
}

func Test_Cov_Result_HasError(t *testing.T) {
	noErr := results.ResultAny{}
	hasErr := results.ResultAny{Error: errors.New("e")}
	actual := args.Map{"noErr": noErr.HasError(), "hasErr": hasErr.HasError()}
	expected := args.Map{"noErr": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "HasError", actual)
}

func Test_Cov_Result_HasPanicked(t *testing.T) {
	r := results.ResultAny{Panicked: true}
	actual := args.Map{"panicked": r.HasPanicked()}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HasPanicked", actual)
}

func Test_Cov_Result_IsResult(t *testing.T) {
	r := results.ResultAny{Value: "hello"}
	actual := args.Map{"match": r.IsResult("hello"), "noMatch": r.IsResult("world")}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsResult", actual)
}

func Test_Cov_Result_IsResultTypeOf(t *testing.T) {
	r := results.ResultAny{Value: "hello"}
	actual := args.Map{
		"matchString": r.IsResultTypeOf(""),
		"matchInt":    r.IsResultTypeOf(0),
	}
	expected := args.Map{
		"matchString": true,
		"matchInt":    false,
	}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf", actual)
}

func Test_Cov_Result_IsResultTypeOf_Nil(t *testing.T) {
	r := results.ResultAny{Value: nil}
	actual := args.Map{"nilExpected": r.IsResultTypeOf(nil)}
	expected := args.Map{"nilExpected": true}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf nil", actual)
}

func Test_Cov_Result_IsError(t *testing.T) {
	r := results.ResultAny{Error: errors.New("fail")}
	noErr := results.ResultAny{}
	actual := args.Map{
		"match":   r.IsError("fail"),
		"noMatch": r.IsError("other"),
		"nilErr":  noErr.IsError("fail"),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
		"nilErr":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsError", actual)
}

func Test_Cov_Result_ValueString(t *testing.T) {
	r := results.ResultAny{Value: 42}
	actual := args.Map{"val": r.ValueString()}
	expected := args.Map{"val": "42"}
	expected.ShouldBeEqual(t, 0, "ValueString", actual)
}

func Test_Cov_Result_ResultAt(t *testing.T) {
	r := results.ResultAny{AllResults: []any{"a", "b"}}
	actual := args.Map{
		"at0":    fmt.Sprintf("%v", r.ResultAt(0)),
		"at1":    fmt.Sprintf("%v", r.ResultAt(1)),
		"outLow": fmt.Sprintf("%v", r.ResultAt(-1)),
		"outHi":  fmt.Sprintf("%v", r.ResultAt(5)),
	}
	expected := args.Map{
		"at0":    "a",
		"at1":    "b",
		"outLow": "<nil>",
		"outHi":  "<nil>",
	}
	expected.ShouldBeEqual(t, 0, "ResultAt", actual)
}

func Test_Cov_Result_ToMap(t *testing.T) {
	r := results.ResultAny{Value: "v", ReturnCount: 1}
	m := r.ToMap()
	actual := args.Map{
		"hasValue": m.Has("value"),
		"hasPanic": m.Has("panicked"),
		"hasSafe":  m.Has("isSafe"),
	}
	expected := args.Map{
		"hasValue": true,
		"hasPanic": true,
		"hasSafe":  true,
	}
	expected.ShouldBeEqual(t, 0, "ToMap", actual)
}

func Test_Cov_Result_ToMapCompact(t *testing.T) {
	r := results.ResultAny{Value: "v"}
	m := r.ToMapCompact()
	actual := args.Map{"keys": m.Has("value") && m.Has("panicked")}
	expected := args.Map{"keys": true}
	expected.ShouldBeEqual(t, 0, "ToMapCompact", actual)
}

func Test_Cov_Result_String_Panicked(t *testing.T) {
	r := results.ResultAny{Panicked: true, PanicValue: "boom"}
	actual := args.Map{"contains": r.String() != ""}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "String panicked", actual)
}

func Test_Cov_Result_String_Error(t *testing.T) {
	r := results.ResultAny{Error: errors.New("fail"), ReturnCount: 1}
	actual := args.Map{"contains": r.String() != ""}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "String error", actual)
}

func Test_Cov_Result_String_Normal(t *testing.T) {
	r := results.ResultAny{Value: "ok", ReturnCount: 1}
	actual := args.Map{"contains": r.String() != ""}
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "String normal", actual)
}

// ── Results methods ──

func Test_Cov_Results_String_Panicked(t *testing.T) {
	r := results.ResultsAny{Result: results.ResultAny{Panicked: true, PanicValue: "p"}}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Results String panicked", actual)
}

func Test_Cov_Results_String_Error(t *testing.T) {
	r := results.ResultsAny{
		Result:  results.ResultAny{Value: "v", Error: errors.New("e")},
		Result2: "r2",
	}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Results String error", actual)
}

func Test_Cov_Results_String_Normal(t *testing.T) {
	r := results.ResultsAny{
		Result:  results.ResultAny{Value: "v"},
		Result2: "r2",
	}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Results String normal", actual)
}

func Test_Cov_Results_IsResult2(t *testing.T) {
	r := results.ResultsAny{Result2: "hello"}
	actual := args.Map{"match": r.IsResult2("hello"), "noMatch": r.IsResult2("x")}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsResult2", actual)
}

func Test_Cov_Results_Result2String(t *testing.T) {
	r := results.ResultsAny{Result2: 99}
	actual := args.Map{"val": r.Result2String()}
	expected := args.Map{"val": "99"}
	expected.ShouldBeEqual(t, 0, "Result2String", actual)
}

func Test_Cov_FromResultAny(t *testing.T) {
	ra := results.ResultAny{
		Value:       "v",
		Panicked:    false,
		AllResults:  []any{"v", "r2"},
		ReturnCount: 2,
	}
	r := results.FromResultAny[string, string](ra)
	actual := args.Map{"val": r.Value, "r2": r.Result2, "count": r.ReturnCount}
	expected := args.Map{"val": "v", "r2": "r2", "count": 2}
	expected.ShouldBeEqual(t, 0, "FromResultAny", actual)
}

func Test_Cov_FromResultAny_Empty(t *testing.T) {
	ra := results.ResultAny{AllResults: []any{}}
	r := results.FromResultAny[string, string](ra)
	actual := args.Map{"val": r.Value, "r2": r.Result2}
	expected := args.Map{"val": "", "r2": ""}
	expected.ShouldBeEqual(t, 0, "FromResultAny empty", actual)
}

// ── MethodName ──

type covTestStruct struct{}

func (s *covTestStruct) Hello() string { return "hi" }

func Test_Cov_MethodName(t *testing.T) {
	name := results.MethodName((*covTestStruct).Hello)
	nilName := results.MethodName(nil)
	nonFunc := results.MethodName("notafunc")

	actual := args.Map{"name": name, "nil": nilName, "nonFunc": nonFunc}
	expected := args.Map{"name": "Hello", "nil": "", "nonFunc": ""}
	expected.ShouldBeEqual(t, 0, "MethodName", actual)
}

// ── InvokeWithPanicRecovery ──

func Test_Cov_Invoke_NilFunc(t *testing.T) {
	r := results.InvokeWithPanicRecovery(nil, nil)
	actual := args.Map{"panicked": r.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Invoke nil func", actual)
}

func Test_Cov_Invoke_NotAFunc(t *testing.T) {
	r := results.InvokeWithPanicRecovery("notfunc", nil)
	actual := args.Map{"panicked": r.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Invoke not a func", actual)
}

func Test_Cov_Invoke_NilReceiver(t *testing.T) {
	r := results.InvokeWithPanicRecovery((*covTestStruct).Hello, nil)
	actual := args.Map{"panicked": r.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Invoke nil receiver panics", actual)
}

func Test_Cov_Invoke_ValidCall(t *testing.T) {
	s := &covTestStruct{}
	r := results.InvokeWithPanicRecovery((*covTestStruct).Hello, s)
	actual := args.Map{"value": fmt.Sprintf("%v", r.Value), "panicked": r.Panicked, "count": r.ReturnCount}
	expected := args.Map{"value": "hi", "panicked": false, "count": 1}
	expected.ShouldBeEqual(t, 0, "Invoke valid", actual)
}

type covErrStruct struct{}

func (s *covErrStruct) Fail() error { return errors.New("fail") }
func (s *covErrStruct) Ok() error   { return nil }

func Test_Cov_Invoke_ErrorReturn(t *testing.T) {
	s := &covErrStruct{}
	r := results.InvokeWithPanicRecovery((*covErrStruct).Fail, s)
	actual := args.Map{"hasError": r.HasError(), "panicked": r.Panicked}
	expected := args.Map{"hasError": true, "panicked": false}
	expected.ShouldBeEqual(t, 0, "Invoke error return", actual)
}

func Test_Cov_Invoke_NilErrorReturn(t *testing.T) {
	s := &covErrStruct{}
	r := results.InvokeWithPanicRecovery((*covErrStruct).Ok, s)
	actual := args.Map{"hasError": r.HasError(), "panicked": r.Panicked}
	expected := args.Map{"hasError": false, "panicked": false}
	expected.ShouldBeEqual(t, 0, "Invoke nil error return", actual)
}

type covVoidStruct struct{}

func (s *covVoidStruct) DoNothing() {}

func Test_Cov_Invoke_VoidReturn(t *testing.T) {
	s := &covVoidStruct{}
	r := results.InvokeWithPanicRecovery((*covVoidStruct).DoNothing, s)
	actual := args.Map{"count": r.ReturnCount, "panicked": r.Panicked}
	expected := args.Map{"count": 0, "panicked": false}
	expected.ShouldBeEqual(t, 0, "Invoke void", actual)
}

type covMultiStruct struct{}

func (s *covMultiStruct) TwoVals() (string, int) { return "a", 1 }

func Test_Cov_Invoke_MultiReturn(t *testing.T) {
	s := &covMultiStruct{}
	r := results.InvokeWithPanicRecovery((*covMultiStruct).TwoVals, s)
	actual := args.Map{
		"count":   r.ReturnCount,
		"val":     fmt.Sprintf("%v", r.Value),
		"allLen":  len(r.AllResults),
		"hasErr":  r.HasError(),
	}
	expected := args.Map{
		"count":   2,
		"val":     "a",
		"allLen":  2,
		"hasErr":  false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke multi return", actual)
}

// ── Invoke with nil args ──

type covArgStruct struct{}

func (s *covArgStruct) WithArg(v any) string {
	return fmt.Sprintf("%v", v)
}

func Test_Cov_Invoke_NilArg(t *testing.T) {
	s := &covArgStruct{}
	r := results.InvokeWithPanicRecovery((*covArgStruct).WithArg, s, nil)
	actual := args.Map{"panicked": r.Panicked, "count": r.ReturnCount}
	expected := args.Map{"panicked": false, "count": 1}
	expected.ShouldBeEqual(t, 0, "Invoke nil arg", actual)
}

// ── ExpectAnyError sentinel ──

func Test_Cov_ExpectAnyError(t *testing.T) {
	actual := args.Map{"notNil": results.ExpectAnyError != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectAnyError sentinel", actual)
}

// ── deriveCompareFields via ShouldMatchResult ──
// Indirect test: we test filterByFields by calling ToMap on empty result

func Test_Cov_FilterByFields_MissingKey(t *testing.T) {
	r := results.ResultAny{}
	m := r.ToMap()
	// Verify the map has expected keys
	actual := args.Map{
		"hasValue":    m.Has("value"),
		"hasPanicked": m.Has("panicked"),
	}
	expected := args.Map{
		"hasValue":    true,
		"hasPanicked": true,
	}
	expected.ShouldBeEqual(t, 0, "FilterByFields", actual)
}
