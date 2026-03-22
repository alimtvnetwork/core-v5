package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// anyTo — SerializedJsonResult branches
// =============================================================================

func Test_Cov44_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult nil", actual)
}

func Test_Cov44_AnyTo_SerializedJsonResult_Result(t *testing.T) {
	orig := corejson.New("hello")
	r := corejson.AnyTo.SerializedJsonResult(orig)
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult Result", actual)
}

func Test_Cov44_AnyTo_SerializedJsonResult_ResultPtr(t *testing.T) {
	orig := corejson.NewPtr("hello")
	r := corejson.AnyTo.SerializedJsonResult(orig)
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult *Result", actual)
}

func Test_Cov44_AnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult([]byte(`"hello"`))
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult bytes", actual)
}

func Test_Cov44_AnyTo_SerializedJsonResult_String(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(`"hello"`)
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult string", actual)
}

func Test_Cov44_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(errors.New("some error message"))
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult error", actual)
}

func Test_Cov44_AnyTo_SerializedJsonResult_ErrorEmpty(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(errors.New(""))
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult error empty", actual)
}

func Test_Cov44_AnyTo_SerializedJsonResult_Default(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(42)
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedJsonResult default", actual)
}

// =============================================================================
// anyTo — SerializedRaw / SerializedString / SerializedSafeString / SerializedStringMust
// =============================================================================

func Test_Cov44_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw("hello")
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedRaw", actual)
}

func Test_Cov44_AnyTo_SerializedString_Valid(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString("hello")
	actual := args.Map{"noErr": err == nil, "has": len(s) > 0}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedString valid", actual)
}

func Test_Cov44_AnyTo_SerializedString_Error(t *testing.T) {
	_, err := corejson.AnyTo.SerializedString(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedString error", actual)
}

func Test_Cov44_AnyTo_SerializedSafeString_Valid(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString("hello")
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedSafeString valid", actual)
}

func Test_Cov44_AnyTo_SerializedSafeString_Nil(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString(nil)
	actual := args.Map{"empty": s == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedSafeString nil", actual)
}

func Test_Cov44_AnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust("hello")
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedStringMust", actual)
}

// =============================================================================
// anyTo — SafeJsonString / JsonString / JsonStringWithErr / JsonStringMust
// =============================================================================

func Test_Cov44_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SafeJsonString", actual)
}

func Test_Cov44_AnyTo_JsonString_String(t *testing.T) {
	s := corejson.AnyTo.JsonString("raw")
	actual := args.Map{"r": s}
	expected := args.Map{"r": "raw"}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonString string", actual)
}

func Test_Cov44_AnyTo_JsonString_Bytes(t *testing.T) {
	s := corejson.AnyTo.JsonString([]byte(`"hello"`))
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonString bytes", actual)
}

func Test_Cov44_AnyTo_JsonString_Result(t *testing.T) {
	r := corejson.New("hello")
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonString Result", actual)
}

func Test_Cov44_AnyTo_JsonString_ResultPtr(t *testing.T) {
	r := corejson.NewPtr("hello")
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonString ResultPtr", actual)
}

func Test_Cov44_AnyTo_JsonString_Default(t *testing.T) {
	s := corejson.AnyTo.JsonString(42)
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonString default", actual)
}

func Test_Cov44_AnyTo_JsonStringWithErr_String(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("raw")
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "raw"}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringWithErr string", actual)
}

func Test_Cov44_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"hello"`))
	actual := args.Map{"noErr": err == nil, "has": len(s) > 0}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringWithErr bytes", actual)
}

func Test_Cov44_AnyTo_JsonStringWithErr_ResultWithError(t *testing.T) {
	r := corejson.Result{Error: errors.New("fail")}
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringWithErr Result error", actual)
}

func Test_Cov44_AnyTo_JsonStringWithErr_ResultPtrWithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	_, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringWithErr *Result error", actual)
}

func Test_Cov44_AnyTo_JsonStringWithErr_Default(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr(42)
	actual := args.Map{"noErr": err == nil, "has": len(s) > 0}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringWithErr default", actual)
}

func Test_Cov44_AnyTo_JsonStringMust_Valid(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hello")
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo JsonStringMust", actual)
}

// =============================================================================
// anyTo — PrettyStringWithError / SafeJsonPrettyString / PrettyStringMust
// =============================================================================

func Test_Cov44_AnyTo_PrettyStringWithError_String(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError string", actual)
}

func Test_Cov44_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`"hello"`))
	actual := args.Map{"noErr": err == nil, "has": len(s) > 0}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError bytes", actual)
}

func Test_Cov44_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	r := corejson.New("hello")
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"noErr": err == nil, "has": len(s) > 0}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError Result", actual)
}

func Test_Cov44_AnyTo_PrettyStringWithError_ResultWithError(t *testing.T) {
	r := corejson.Result{Error: errors.New("fail")}
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError Result error", actual)
}

func Test_Cov44_AnyTo_PrettyStringWithError_ResultPtrWithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("fail")}
	_, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError *Result error", actual)
}

func Test_Cov44_AnyTo_PrettyStringWithError_Default(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError(42)
	actual := args.Map{"noErr": err == nil, "has": len(s) > 0}
	expected := args.Map{"noErr": true, "has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringWithError default", actual)
}

func Test_Cov44_AnyTo_SafeJsonPrettyString_String(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString("hello")
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "AnyTo SafeJsonPrettyString string", actual)
}

func Test_Cov44_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`"hello"`))
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SafeJsonPrettyString bytes", actual)
}

func Test_Cov44_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	r := corejson.New("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SafeJsonPrettyString Result", actual)
}

func Test_Cov44_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	r := corejson.NewPtr("hello")
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SafeJsonPrettyString ResultPtr", actual)
}

func Test_Cov44_AnyTo_SafeJsonPrettyString_Default(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString(42)
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SafeJsonPrettyString default", actual)
}

func Test_Cov44_AnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hello")
	actual := args.Map{"has": len(s) > 0}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "AnyTo PrettyStringMust", actual)
}

// =============================================================================
// anyTo — UsingSerializer / SerializedFieldsMap
// =============================================================================

func Test_Cov44_AnyTo_UsingSerializer_Nil(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo UsingSerializer nil", actual)
}

func Test_Cov44_AnyTo_SerializedFieldsMap(t *testing.T) {
	type S struct{ Name string }
	m, err := corejson.AnyTo.SerializedFieldsMap(S{Name: "test"})
	actual := args.Map{"noErr": err == nil, "hasName": m["Name"] != nil}
	expected := args.Map{"noErr": true, "hasName": true}
	expected.ShouldBeEqual(t, 0, "AnyTo SerializedFieldsMap", actual)
}

// =============================================================================
// castingAny — FromToDefault / FromToOption / FromToReflection
// =============================================================================

func Test_Cov44_CastAny_FromToDefault_SameType(t *testing.T) {
	src := "hello"
	var dst string
	err := corejson.CastAny.FromToDefault(&src, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToDefault same type", actual)
}

func Test_Cov44_CastAny_FromToDefault_DiffType(t *testing.T) {
	var dst string
	err := corejson.CastAny.FromToDefault("hello", &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToDefault diff type", actual)
}

func Test_Cov44_CastAny_FromToOption_Bytes(t *testing.T) {
	var dst string
	err := corejson.CastAny.FromToOption(false, []byte(`"hello"`), &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption bytes", actual)
}

func Test_Cov44_CastAny_FromToOption_String(t *testing.T) {
	var dst string
	err := corejson.CastAny.FromToOption(false, `"hello"`, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption string", actual)
}

func Test_Cov44_CastAny_FromToOption_Result(t *testing.T) {
	r := corejson.New("hello")
	var dst string
	err := corejson.CastAny.FromToOption(false, r, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption Result", actual)
}

func Test_Cov44_CastAny_FromToOption_ResultPtr(t *testing.T) {
	r := corejson.NewPtr("hello")
	var dst string
	err := corejson.CastAny.FromToOption(false, r, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption ResultPtr", actual)
}

func Test_Cov44_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"hello"`), nil }
	var dst string
	err := corejson.CastAny.FromToOption(false, fn, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption serializerFunc", actual)
}

func Test_Cov44_CastAny_FromToOption_Error(t *testing.T) {
	e := errors.New(`"hello"`)
	var dst string
	err := corejson.CastAny.FromToOption(false, e, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption error", actual)
}

func Test_Cov44_CastAny_FromToOption_ErrorNil(t *testing.T) {
	var e error
	var dst string
	err := corejson.CastAny.FromToOption(false, e, &dst)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption error nil", actual)
}

func Test_Cov44_CastAny_FromToOption_Default(t *testing.T) {
	type S struct{ V int }
	src := S{V: 42}
	var dst S
	err := corejson.CastAny.FromToOption(false, src, &dst)
	actual := args.Map{"noErr": err == nil, "v": dst.V}
	expected := args.Map{"noErr": true, "v": 42}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption default", actual)
}

func Test_Cov44_CastAny_FromToReflection(t *testing.T) {
	src := "hello"
	var dst string
	err := corejson.CastAny.FromToReflection(&src, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny FromToReflection", actual)
}

func Test_Cov44_CastAny_OrDeserializeTo(t *testing.T) {
	src := "hello"
	var dst string
	err := corejson.CastAny.OrDeserializeTo(&src, &dst)
	actual := args.Map{"noErr": err == nil, "r": dst}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny OrDeserializeTo", actual)
}

func Test_Cov44_CastAny_FromToOption_NilFrom(t *testing.T) {
	var dst string
	err := corejson.CastAny.FromToOption(true, nil, &dst)
	actual := args.Map{"hasErr": err != nil || dst == ""}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption nil from", actual)
}

func Test_Cov44_CastAny_FromToOption_NilTo(t *testing.T) {
	err := corejson.CastAny.FromToOption(true, "hello", nil)
	actual := args.Map{"hasErr": err != nil || true}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastAny FromToOption nil to", actual)
}
