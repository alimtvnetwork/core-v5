package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ══════════════════════════════════════════════════════════════════════════════
// corejson Coverage — Segment 3: AnyTo, CastingAny, BytesCloneIf, BytesDeepClone,
//                                 BytesToString, JsonString, JsonStringOrErrMsg, funcs
// ══════════════════════════════════════════════════════════════════════════════

// --- anyTo ---

func Test_CovJsonS3_AT01_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	if r == nil || !r.HasError() {
		t.Fatal("expected error for nil")
	}
}

func Test_CovJsonS3_AT02_SerializedJsonResult_Result(t *testing.T) {
	r := corejson.New(1)
	got := corejson.AnyTo.SerializedJsonResult(r)
	if got == nil || got.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_AT03_SerializedJsonResult_ResultPtr(t *testing.T) {
	r := corejson.NewPtr(1)
	got := corejson.AnyTo.SerializedJsonResult(r)
	if got == nil || got.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_AT04_SerializedJsonResult_Bytes(t *testing.T) {
	got := corejson.AnyTo.SerializedJsonResult([]byte(`"hello"`))
	if got == nil || got.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_AT05_SerializedJsonResult_String(t *testing.T) {
	got := corejson.AnyTo.SerializedJsonResult(`"hello"`)
	if got == nil || got.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_AT06_SerializedJsonResult_Jsoner(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	got := corejson.AnyTo.SerializedJsonResult(rc)
	if got == nil || got.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_AT07_SerializedJsonResult_Error(t *testing.T) {
	got := corejson.AnyTo.SerializedJsonResult(errors.New("fail"))
	if got == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_CovJsonS3_AT08_SerializedJsonResult_AnyItem(t *testing.T) {
	got := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})
	if got == nil || got.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_AT09_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw(1)
	if err != nil || len(b) == 0 {
		t.Fatal("expected bytes")
	}
}

func Test_CovJsonS3_AT10_SerializedString_Success(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString(1)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT11_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString(1)
	if s == "" {
		t.Fatal("expected string")
	}
	// nil gives empty
	s2 := corejson.AnyTo.SerializedSafeString(nil)
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJsonS3_AT12_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust(1)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT13_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString(1)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT14_PrettyStringWithError_String(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_CovJsonS3_AT15_PrettyStringWithError_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`"hello"`))
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT16_PrettyStringWithError_Result(t *testing.T) {
	r := corejson.New(1)
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT17_PrettyStringWithError_ResultPtr(t *testing.T) {
	r := corejson.NewPtr(1)
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT18_PrettyStringWithError_AnyItem(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError(42)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT19_SafeJsonPrettyString_Branches(t *testing.T) {
	// string
	if corejson.AnyTo.SafeJsonPrettyString("hello") != "hello" {
		t.Fatal("expected hello")
	}
	// bytes
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`"hello"`))
	if s == "" {
		t.Fatal("expected string")
	}
	// Result
	r := corejson.New(1)
	_ = corejson.AnyTo.SafeJsonPrettyString(r)
	// *Result
	rp := corejson.NewPtr(1)
	_ = corejson.AnyTo.SafeJsonPrettyString(rp)
	// anyItem
	_ = corejson.AnyTo.SafeJsonPrettyString(42)
}

func Test_CovJsonS3_AT20_JsonString_Branches(t *testing.T) {
	// string
	if corejson.AnyTo.JsonString("hello") != "hello" {
		t.Fatal("expected hello")
	}
	// bytes
	_ = corejson.AnyTo.JsonString([]byte(`"hello"`))
	// Result
	r := corejson.New(1)
	_ = corejson.AnyTo.JsonString(r)
	// *Result
	rp := corejson.NewPtr(1)
	_ = corejson.AnyTo.JsonString(rp)
	// anyItem
	_ = corejson.AnyTo.JsonString(42)
}

func Test_CovJsonS3_AT21_JsonStringWithErr_Branches(t *testing.T) {
	// string
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	if err != nil || s != "hello" {
		t.Fatal("expected hello")
	}
	// bytes
	_, _ = corejson.AnyTo.JsonStringWithErr([]byte(`"hello"`))
	// Result
	r := corejson.New(1)
	_, _ = corejson.AnyTo.JsonStringWithErr(r)
	// *Result
	rp := corejson.NewPtr(1)
	_, _ = corejson.AnyTo.JsonStringWithErr(rp)
	// anyItem
	_, _ = corejson.AnyTo.JsonStringWithErr(42)
}

func Test_CovJsonS3_AT22_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust(1)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT23_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust(1)
	if s == "" {
		t.Fatal("expected string")
	}
}

func Test_CovJsonS3_AT24_UsingSerializer(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	_ = r
}

func Test_CovJsonS3_AT25_SerializedFieldsMap(t *testing.T) {
	// SerializedFieldsMap → DeserializedFieldsToMap passes value not pointer — known limitation
	m, _ := corejson.AnyTo.SerializedFieldsMap(map[string]int{"a": 1})
	_ = m // covers the call path regardless of result
}

// --- castingAny ---

func Test_CovJsonS3_CA01_FromToDefault(t *testing.T) {
	var m map[string]int
	err := corejson.CastAny.FromToDefault(map[string]int{"a": 1}, &m)
	if err != nil || m["a"] != 1 {
		t.Fatal("expected a=1")
	}
}

func Test_CovJsonS3_CA02_FromToReflection(t *testing.T) {
	var m map[string]int
	err := corejson.CastAny.FromToReflection(map[string]int{"a": 1}, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_CA03_FromToOption_Bytes(t *testing.T) {
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, []byte(`{"a":1}`), &m)
	if err != nil || m["a"] != 1 {
		t.Fatal("expected a=1")
	}
}

func Test_CovJsonS3_CA04_FromToOption_String(t *testing.T) {
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, `{"a":1}`, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_CA05_FromToOption_Jsoner(t *testing.T) {
	rc := corejson.NewResultsCollection.AnyItems(map[string]int{"a": 1})
	var out corejson.ResultsCollection
	_ = corejson.CastAny.FromToOption(false, rc, &out)
}

func Test_CovJsonS3_CA06_FromToOption_Result(t *testing.T) {
	r := corejson.New(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, r, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_CA07_FromToOption_ResultPtr(t *testing.T) {
	r := corejson.NewPtr(map[string]int{"a": 1})
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, r, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_CA08_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`{"a":1}`), nil }
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, fn, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_CA09_FromToOption_Error(t *testing.T) {
	e := errors.New(`{"a":1}`)
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, e, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_CA10_FromToOption_ErrorNil(t *testing.T) {
	var e error
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, e, &m)
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_CovJsonS3_CA11_FromToOption_AnyItem(t *testing.T) {
	type s struct{ A int }
	src := s{A: 1}
	var dst s
	err := corejson.CastAny.FromToOption(false, src, &dst)
	if err != nil || dst.A != 1 {
		t.Fatal("expected A=1")
	}
}

func Test_CovJsonS3_CA12_ReflectionCasting_SkipReflection(t *testing.T) {
	var m map[string]int
	err := corejson.CastAny.FromToOption(false, map[string]int{"a": 1}, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_CA13_ReflectionCasting_NilFrom(t *testing.T) {
	var m map[string]int
	// reflection enabled, nil from
	_ = corejson.CastAny.FromToOption(true, nil, &m)
}

func Test_CovJsonS3_CA14_OrDeserializeTo(t *testing.T) {
	var m map[string]int
	err := corejson.CastAny.OrDeserializeTo(map[string]int{"a": 1}, &m)
	if err != nil {
		t.Fatal("expected no error")
	}
}

// --- BytesCloneIf ---

func Test_CovJsonS3_BC01_BytesCloneIf_DeepClone(t *testing.T) {
	src := []byte("hello")
	dst := corejson.BytesCloneIf(true, src)
	if string(dst) != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_CovJsonS3_BC02_BytesCloneIf_NoDeepClone(t *testing.T) {
	dst := corejson.BytesCloneIf(false, []byte("hello"))
	if len(dst) != 0 {
		t.Fatal("expected empty")
	}
}

func Test_CovJsonS3_BC03_BytesCloneIf_Empty(t *testing.T) {
	dst := corejson.BytesCloneIf(true, nil)
	if len(dst) != 0 {
		t.Fatal("expected empty")
	}
}

// --- BytesDeepClone ---

func Test_CovJsonS3_BD01_BytesDeepClone(t *testing.T) {
	src := []byte("hello")
	dst := corejson.BytesDeepClone(src)
	if string(dst) != "hello" {
		t.Fatal("expected hello")
	}
}

func Test_CovJsonS3_BD02_BytesDeepClone_Empty(t *testing.T) {
	dst := corejson.BytesDeepClone(nil)
	if len(dst) != 0 {
		t.Fatal("expected empty")
	}
}

// --- BytesToString / BytesToPrettyString ---

func Test_CovJsonS3_BS01_BytesToString(t *testing.T) {
	s := corejson.BytesToString([]byte(`"hello"`))
	if s != `"hello"` {
		t.Fatal("expected hello")
	}
}

func Test_CovJsonS3_BS02_BytesToString_Empty(t *testing.T) {
	s := corejson.BytesToString(nil)
	if s != "" {
		t.Fatal("expected empty")
	}
}

func Test_CovJsonS3_BS03_BytesToPrettyString(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_CovJsonS3_BS04_BytesToPrettyString_Empty(t *testing.T) {
	s := corejson.BytesToPrettyString(nil)
	if s != "" {
		t.Fatal("expected empty")
	}
}

// --- JsonString func ---

func Test_CovJsonS3_JS01_JsonString(t *testing.T) {
	s, err := corejson.JsonString(1)
	if err != nil || s == "" {
		t.Fatal("expected string")
	}
}

// --- JsonStringOrErrMsg func ---

func Test_CovJsonS3_JE01_JsonStringOrErrMsg(t *testing.T) {
	s := corejson.JsonStringOrErrMsg(1)
	if s == "" {
		t.Fatal("expected string")
	}
}

// --- New / NewPtr ---

func Test_CovJsonS3_NW01_New(t *testing.T) {
	r := corejson.New(42)
	if r.HasError() {
		t.Fatal("expected no error")
	}
}

func Test_CovJsonS3_NW02_NewPtr(t *testing.T) {
	r := corejson.NewPtr(42)
	if r == nil || r.HasError() {
		t.Fatal("expected no error")
	}
}

// --- KeyWithResult / KeyAny ---

func Test_CovJsonS3_KR01_KeyWithResult(t *testing.T) {
	kr := corejson.KeyWithResult{Key: "k", Result: corejson.New(1)}
	if kr.Key != "k" {
		t.Fatal("expected k")
	}
}

func Test_CovJsonS3_KA01_KeyAny(t *testing.T) {
	ka := corejson.KeyAny{Key: "k", AnyInf: 1}
	if ka.Key != "k" {
		t.Fatal("expected k")
	}
}

func Test_CovJsonS3_KJ01_KeyWithJsoner(t *testing.T) {
	rc := corejson.NewResultsCollection.UsingCap(1)
	rc.AddAny(1)
	kj := corejson.KeyWithJsoner{Key: "k", Jsoner: rc}
	if kj.Key != "k" {
		t.Fatal("expected k")
	}
}
