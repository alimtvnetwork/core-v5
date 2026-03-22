package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// deserializeFromBytesTo — String / Strings / StringMust / StringsMust
// =============================================================================

func Test_Cov42_BytesTo_String_Valid(t *testing.T) {
	s, err := corejson.Deserialize.BytesTo.String([]byte(`"hello"`))
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesTo String valid", actual)
}

func Test_Cov42_BytesTo_StringMust(t *testing.T) {
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesTo StringMust", actual)
}

func Test_Cov42_BytesTo_StringMust_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "BytesTo StringMust panics", actual)
	}()
	corejson.Deserialize.BytesTo.StringMust([]byte(`bad`))
}

func Test_Cov42_BytesTo_Strings(t *testing.T) {
	ss, err := corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
	actual := args.Map{"noErr": err == nil, "len": len(ss)}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "BytesTo Strings", actual)
}

func Test_Cov42_BytesTo_StringsMust(t *testing.T) {
	ss := corejson.Deserialize.BytesTo.StringsMust([]byte(`["a"]`))
	actual := args.Map{"len": len(ss)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesTo StringsMust", actual)
}

func Test_Cov42_BytesTo_StringsMust_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "BytesTo StringsMust panics", actual)
	}()
	corejson.Deserialize.BytesTo.StringsMust([]byte(`bad`))
}

// =============================================================================
// deserializeFromBytesTo — Integer / Integer64 / Integers
// =============================================================================

func Test_Cov42_BytesTo_Integer(t *testing.T) {
	v, err := corejson.Deserialize.BytesTo.Integer([]byte(`42`))
	actual := args.Map{"noErr": err == nil, "r": v}
	expected := args.Map{"noErr": true, "r": 42}
	expected.ShouldBeEqual(t, 0, "BytesTo Integer", actual)
}

func Test_Cov42_BytesTo_IntegerMust(t *testing.T) {
	v := corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))
	actual := args.Map{"r": v}
	expected := args.Map{"r": 42}
	expected.ShouldBeEqual(t, 0, "BytesTo IntegerMust", actual)
}

func Test_Cov42_BytesTo_Integer64(t *testing.T) {
	v, err := corejson.Deserialize.BytesTo.Integer64([]byte(`99`))
	actual := args.Map{"noErr": err == nil, "r": v}
	expected := args.Map{"noErr": true, "r": int64(99)}
	expected.ShouldBeEqual(t, 0, "BytesTo Integer64", actual)
}

func Test_Cov42_BytesTo_Integer64Must(t *testing.T) {
	v := corejson.Deserialize.BytesTo.Integer64Must([]byte(`99`))
	actual := args.Map{"r": v}
	expected := args.Map{"r": int64(99)}
	expected.ShouldBeEqual(t, 0, "BytesTo Integer64Must", actual)
}

func Test_Cov42_BytesTo_Integers(t *testing.T) {
	v, err := corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
	actual := args.Map{"noErr": err == nil, "len": len(v)}
	expected := args.Map{"noErr": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "BytesTo Integers", actual)
}

func Test_Cov42_BytesTo_IntegersMust(t *testing.T) {
	v := corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1]`))
	actual := args.Map{"len": len(v)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesTo IntegersMust", actual)
}

func Test_Cov42_BytesTo_IntegersMust_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "BytesTo IntegersMust panics", actual)
	}()
	corejson.Deserialize.BytesTo.IntegersMust([]byte(`bad`))
}

// =============================================================================
// deserializeFromBytesTo — Bool / MapAnyItem / MapStringString
// =============================================================================

func Test_Cov42_BytesTo_Bool(t *testing.T) {
	v, err := corejson.Deserialize.BytesTo.Bool([]byte(`true`))
	actual := args.Map{"noErr": err == nil, "r": v}
	expected := args.Map{"noErr": true, "r": true}
	expected.ShouldBeEqual(t, 0, "BytesTo Bool", actual)
}

func Test_Cov42_BytesTo_BoolMust(t *testing.T) {
	v := corejson.Deserialize.BytesTo.BoolMust([]byte(`false`))
	actual := args.Map{"r": v}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "BytesTo BoolMust", actual)
}

func Test_Cov42_BytesTo_MapAnyItem(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
	actual := args.Map{"noErr": err == nil, "hasA": m["a"] != nil}
	expected := args.Map{"noErr": true, "hasA": true}
	expected.ShouldBeEqual(t, 0, "BytesTo MapAnyItem", actual)
}

func Test_Cov42_BytesTo_MapAnyItemMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))
	actual := args.Map{"hasA": m["a"] != nil}
	expected := args.Map{"hasA": true}
	expected.ShouldBeEqual(t, 0, "BytesTo MapAnyItemMust", actual)
}

func Test_Cov42_BytesTo_MapStringString(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
	actual := args.Map{"noErr": err == nil, "a": m["a"]}
	expected := args.Map{"noErr": true, "a": "b"}
	expected.ShouldBeEqual(t, 0, "BytesTo MapStringString", actual)
}

func Test_Cov42_BytesTo_MapStringStringMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))
	actual := args.Map{"a": m["a"]}
	expected := args.Map{"a": "b"}
	expected.ShouldBeEqual(t, 0, "BytesTo MapStringStringMust", actual)
}

// =============================================================================
// deserializeFromBytesTo — Bytes / ResultCollection / ResultsPtrCollection / MapResults
// =============================================================================

func Test_Cov42_BytesTo_Bytes(t *testing.T) {
	b, err := corejson.Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "BytesTo Bytes", actual)
}

func Test_Cov42_BytesTo_BytesMust(t *testing.T) {
	b := corejson.Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "BytesTo BytesMust", actual)
}

func Test_Cov42_BytesTo_ResultCollection_Invalid(t *testing.T) {
	_, err := corejson.Deserialize.BytesTo.ResultCollection([]byte(`bad`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo ResultCollection invalid", actual)
}

func Test_Cov42_BytesTo_ResultsPtrCollection_Invalid(t *testing.T) {
	_, err := corejson.Deserialize.BytesTo.ResultsPtrCollection([]byte(`bad`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo ResultsPtrCollection invalid", actual)
}

func Test_Cov42_BytesTo_MapResults_Invalid(t *testing.T) {
	_, err := corejson.Deserialize.BytesTo.MapResults([]byte(`bad`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo MapResults invalid", actual)
}

func Test_Cov42_BytesTo_MapResultsMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust(map[string]corejson.Result{})
	mr := corejson.Deserialize.BytesTo.MapResultsMust(b)
	actual := args.Map{"notNil": mr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesTo MapResultsMust", actual)
}

// =============================================================================
// deserializeFromResultTo — String / Bool / Byte
// =============================================================================

func Test_Cov42_ResultTo_String(t *testing.T) {
	r := corejson.NewPtr("hello")
	s, err := corejson.Deserialize.ResultTo.String(r)
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "ResultTo String", actual)
}

func Test_Cov42_ResultTo_StringMust(t *testing.T) {
	r := corejson.NewPtr("hello")
	s := corejson.Deserialize.ResultTo.StringMust(r)
	actual := args.Map{"r": s}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "ResultTo StringMust", actual)
}

func Test_Cov42_ResultTo_StringsMust(t *testing.T) {
	r := corejson.NewPtr([]string{"a", "b"})
	ss := corejson.Deserialize.ResultTo.StringsMust(r)
	actual := args.Map{"len": len(ss)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ResultTo StringsMust", actual)
}

func Test_Cov42_ResultTo_Bool(t *testing.T) {
	r := corejson.NewPtr(true)
	v, err := corejson.Deserialize.ResultTo.Bool(r)
	actual := args.Map{"noErr": err == nil, "r": v}
	expected := args.Map{"noErr": true, "r": true}
	expected.ShouldBeEqual(t, 0, "ResultTo Bool", actual)
}

func Test_Cov42_ResultTo_BoolMust(t *testing.T) {
	r := corejson.NewPtr(true)
	v := corejson.Deserialize.ResultTo.BoolMust(r)
	actual := args.Map{"r": v}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "ResultTo BoolMust", actual)
}

func Test_Cov42_ResultTo_Byte(t *testing.T) {
	r := corejson.NewPtr(byte(65))
	v, err := corejson.Deserialize.ResultTo.Byte(r)
	actual := args.Map{"noErr": err == nil, "r": v}
	expected := args.Map{"noErr": true, "r": byte(65)}
	expected.ShouldBeEqual(t, 0, "ResultTo Byte", actual)
}

func Test_Cov42_ResultTo_ByteMust(t *testing.T) {
	r := corejson.NewPtr(byte(65))
	v := corejson.Deserialize.ResultTo.ByteMust(r)
	actual := args.Map{"r": v}
	expected := args.Map{"r": byte(65)}
	expected.ShouldBeEqual(t, 0, "ResultTo ByteMust", actual)
}

// =============================================================================
// deserializeFromResultTo — MapAnyItem / MapStringString
// =============================================================================

func Test_Cov42_ResultTo_MapAnyItem(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"a": 1})
	m, err := corejson.Deserialize.ResultTo.MapAnyItem(r)
	actual := args.Map{"noErr": err == nil, "hasA": m["a"] != nil}
	expected := args.Map{"noErr": true, "hasA": true}
	expected.ShouldBeEqual(t, 0, "ResultTo MapAnyItem", actual)
}

func Test_Cov42_ResultTo_MapAnyItemMust(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"a": 1})
	m := corejson.Deserialize.ResultTo.MapAnyItemMust(r)
	actual := args.Map{"hasA": m["a"] != nil}
	expected := args.Map{"hasA": true}
	expected.ShouldBeEqual(t, 0, "ResultTo MapAnyItemMust", actual)
}

func Test_Cov42_ResultTo_MapStringString(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"a": "b"})
	m, err := corejson.Deserialize.ResultTo.MapStringString(r)
	actual := args.Map{"noErr": err == nil, "a": m["a"]}
	expected := args.Map{"noErr": true, "a": "b"}
	expected.ShouldBeEqual(t, 0, "ResultTo MapStringString", actual)
}

func Test_Cov42_ResultTo_MapStringStringMust(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"a": "b"})
	m := corejson.Deserialize.ResultTo.MapStringStringMust(r)
	actual := args.Map{"a": m["a"]}
	expected := args.Map{"a": "b"}
	expected.ShouldBeEqual(t, 0, "ResultTo MapStringStringMust", actual)
}

// =============================================================================
// deserializeFromResultTo — ResultCollection / ResultsPtrCollection / MapResults
// =============================================================================

func Test_Cov42_ResultTo_ResultCollection_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`bad`)}
	_, err := corejson.Deserialize.ResultTo.ResultCollection(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultTo ResultCollection invalid", actual)
}

func Test_Cov42_ResultTo_ResultsPtrCollection_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`bad`)}
	_, err := corejson.Deserialize.ResultTo.ResultsPtrCollection(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultTo ResultsPtrCollection invalid", actual)
}

func Test_Cov42_ResultTo_MapResults_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`bad`)}
	_, err := corejson.Deserialize.ResultTo.MapResults(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultTo MapResults invalid", actual)
}

func Test_Cov42_ResultTo_MapResultsMust(t *testing.T) {
	r := corejson.NewPtr(map[string]corejson.Result{})
	mr := corejson.Deserialize.ResultTo.MapResultsMust(r)
	actual := args.Map{"notNil": mr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultTo MapResultsMust", actual)
}

// =============================================================================
// deserializeFromResultTo — Result / ResultPtr / Bytes
// =============================================================================

func Test_Cov42_ResultTo_Result_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`bad`)}
	_, err := corejson.Deserialize.ResultTo.Result(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultTo Result invalid", actual)
}

func Test_Cov42_ResultTo_ResultPtr_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`bad`)}
	_, err := corejson.Deserialize.ResultTo.ResultPtr(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultTo ResultPtr invalid", actual)
}

func Test_Cov42_ResultTo_Bytes_Invalid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`bad`)}
	_, err := corejson.Deserialize.ResultTo.Bytes(r)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ResultTo Bytes invalid", actual)
}
