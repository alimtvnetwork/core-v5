package coretestsresultstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/results"
)

// ── Result edge cases ──

func Test_Cov_Result_IsResultTypeOf_Nil(t *testing.T) {
	r := results.ResultAny{Value: nil}
	actual := args.Map{"nilExpected": r.IsResultTypeOf(nil)}
	expected := args.Map{"nilExpected": true}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf nil", actual)
}

func Test_Cov_Result_String_Panicked(t *testing.T) {
	r := results.ResultAny{Panicked: true, PanicValue: "boom"}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String panicked", actual)
}

func Test_Cov_Result_String_Error(t *testing.T) {
	r := results.ResultAny{Error: errors.New("fail"), ReturnCount: 1}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String error", actual)
}

func Test_Cov_Result_String_Normal(t *testing.T) {
	r := results.ResultAny{Value: "ok", ReturnCount: 1}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String normal", actual)
}

// ── ResultsAny methods ──

func Test_Cov_ResultsAny_String_Panicked(t *testing.T) {
	r := results.ResultsAny{Result: results.ResultAny{Panicked: true, PanicValue: "p"}}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsAny panicked", actual)
}

func Test_Cov_ResultsAny_String_Error(t *testing.T) {
	r := results.ResultsAny{
		Result:  results.ResultAny{Value: "v", Error: errors.New("e")},
		Result2: "r2",
	}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsAny error", actual)
}

func Test_Cov_ResultsAny_String_Normal(t *testing.T) {
	r := results.ResultsAny{
		Result:  results.ResultAny{Value: "v"},
		Result2: "r2",
	}
	actual := args.Map{"notEmpty": r.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsAny normal", actual)
}

func Test_Cov_ResultsAny_IsResult2(t *testing.T) {
	r := results.ResultsAny{Result2: "hello"}
	actual := args.Map{"match": r.IsResult2("hello"), "noMatch": r.IsResult2("x")}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "ResultsAny IsResult2", actual)
}

func Test_Cov_ResultsAny_Result2String(t *testing.T) {
	r := results.ResultsAny{Result2: 99}
	actual := args.Map{"val": r.Result2String()}
	expected := args.Map{"val": "99"}
	expected.ShouldBeEqual(t, 0, "ResultsAny Result2String", actual)
}

// ── FromResultAny edge ──

func Test_Cov_FromResultAny_Empty(t *testing.T) {
	ra := results.ResultAny{AllResults: []any{}}
	r := results.FromResultAny[string, string](ra)
	actual := args.Map{"val": r.Value, "val2": r.Result2}
	expected := args.Map{"val": "", "val2": ""}
	expected.ShouldBeEqual(t, 0, "FromResultAny empty", actual)
}

// ── InvokeWithPanicRecovery extended ──

type extCovTestStruct struct{}

func (s *extCovTestStruct) Hello() string { return "hi" }

func Test_Cov_Invoke_NilReceiver_Ext(t *testing.T) {
	r := results.InvokeWithPanicRecovery((*extCovTestStruct).Hello, nil)
	actual := args.Map{"panicked": r.Panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Invoke nil receiver panics", actual)
}

func Test_Cov_Invoke_ValidCall_Ext(t *testing.T) {
	s := &extCovTestStruct{}
	r := results.InvokeWithPanicRecovery((*extCovTestStruct).Hello, s)
	actual := args.Map{"value": fmt.Sprintf("%v", r.Value), "panicked": r.Panicked, "count": r.ReturnCount}
	expected := args.Map{"value": "hi", "panicked": false, "count": 1}
	expected.ShouldBeEqual(t, 0, "Invoke valid", actual)
}

type extCovErrStruct struct{}

func (s *extCovErrStruct) Fail() error { return errors.New("fail") }
func (s *extCovErrStruct) Ok() error   { return nil }

func Test_Cov_Invoke_ErrorReturn_Ext(t *testing.T) {
	s := &extCovErrStruct{}
	r := results.InvokeWithPanicRecovery((*extCovErrStruct).Fail, s)
	actual := args.Map{"hasError": r.HasError(), "panicked": r.Panicked}
	expected := args.Map{"hasError": true, "panicked": false}
	expected.ShouldBeEqual(t, 0, "Invoke error return", actual)
}

func Test_Cov_Invoke_NilErrorReturn_Ext(t *testing.T) {
	s := &extCovErrStruct{}
	r := results.InvokeWithPanicRecovery((*extCovErrStruct).Ok, s)
	actual := args.Map{"hasError": r.HasError(), "panicked": r.Panicked}
	expected := args.Map{"hasError": false, "panicked": false}
	expected.ShouldBeEqual(t, 0, "Invoke nil error return", actual)
}

type extCovVoidStruct struct{}

func (s *extCovVoidStruct) DoNothing() {}

func Test_Cov_Invoke_VoidReturn_Ext(t *testing.T) {
	s := &extCovVoidStruct{}
	r := results.InvokeWithPanicRecovery((*extCovVoidStruct).DoNothing, s)
	actual := args.Map{"count": r.ReturnCount, "panicked": r.Panicked}
	expected := args.Map{"count": 0, "panicked": false}
	expected.ShouldBeEqual(t, 0, "Invoke void", actual)
}

type extCovMultiStruct struct{}

func (s *extCovMultiStruct) TwoVals() (string, int) { return "a", 1 }

func Test_Cov_Invoke_MultiReturn_Ext(t *testing.T) {
	s := &extCovMultiStruct{}
	r := results.InvokeWithPanicRecovery((*extCovMultiStruct).TwoVals, s)
	actual := args.Map{
		"count":  r.ReturnCount,
		"val":    fmt.Sprintf("%v", r.Value),
		"allLen": len(r.AllResults),
		"hasErr": r.HasError(),
	}
	expected := args.Map{
		"count":  2,
		"val":    "a",
		"allLen": 2,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "Invoke multi return", actual)
}

// ── Invoke with nil args ──

type extCovArgStruct struct{}

func (s *extCovArgStruct) WithArg(v any) string {
	return fmt.Sprintf("%v", v)
}

func Test_Cov_Invoke_NilArg_Ext(t *testing.T) {
	s := &extCovArgStruct{}
	r := results.InvokeWithPanicRecovery((*extCovArgStruct).WithArg, s, nil)
	actual := args.Map{"panicked": r.Panicked, "count": r.ReturnCount}
	expected := args.Map{"panicked": false, "count": 1}
	expected.ShouldBeEqual(t, 0, "Invoke nil arg", actual)
}

// ── ExpectAnyError sentinel ──

func Test_Cov_ExpectAnyError_Ext(t *testing.T) {
	actual := args.Map{"notNil": results.ExpectAnyError != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ExpectAnyError sentinel", actual)
}

// ── FilterByFields via ToMap ──

func Test_Cov_FilterByFields_MissingKey_Ext(t *testing.T) {
	r := results.ResultAny{}
	m := r.ToMap()
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

// ── MethodName combined ──

func Test_Cov_MethodName_Combined_Ext(t *testing.T) {
	name := results.MethodName((*extCovTestStruct).Hello)
	nilName := results.MethodName(nil)
	nonFunc := results.MethodName("notafunc")

	actual := args.Map{"name": name, "nil": nilName, "nonFunc": nonFunc}
	expected := args.Map{"name": "Hello", "nil": "", "nonFunc": ""}
	expected.ShouldBeEqual(t, 0, "MethodName combined", actual)
}
