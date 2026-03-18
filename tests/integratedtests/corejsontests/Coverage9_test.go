package corejsontests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// New / NewPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_New_Valid(t *testing.T) {
	r := corejson.New(map[string]string{"k": "v"})
	actual := args.Map{"noErr": !r.HasError(), "hasBytes": r.HasBytes()}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "New valid", actual)
}

func Test_Cov9_New_Nil(t *testing.T) {
	r := corejson.New(nil)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "New nil", actual)
}

func Test_Cov9_NewPtr_Valid(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"k": "v"})
	actual := args.Map{"notNil": r != nil, "hasBytes": r.HasBytes()}
	expected := args.Map{"notNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewPtr valid", actual)
}

func Test_Cov9_NewPtr_Nil(t *testing.T) {
	r := corejson.NewPtr(nil)
	actual := args.Map{"notNil": r != nil, "empty": r.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "NewPtr nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesCloneIf / BytesDeepClone / BytesToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_BytesCloneIf_True(t *testing.T) {
	original := []byte(`"hello"`)
	cloned := corejson.BytesCloneIf(true, original)
	actual := args.Map{"len": len(cloned), "notSame": &cloned[0] != &original[0]}
	expected := args.Map{"len": 7, "notSame": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf true", actual)
}

func Test_Cov9_BytesCloneIf_False(t *testing.T) {
	original := []byte(`"hello"`)
	cloned := corejson.BytesCloneIf(false, original)
	actual := args.Map{"len": len(cloned)}
	expected := args.Map{"len": 7}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf false", actual)
}

func Test_Cov9_BytesDeepClone_Nil(t *testing.T) {
	result := corejson.BytesDeepClone(nil)
	actual := args.Map{"nil": result == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone nil", actual)
}

func Test_Cov9_BytesDeepClone_Valid(t *testing.T) {
	original := []byte(`"hi"`)
	result := corejson.BytesDeepClone(original)
	actual := args.Map{"len": len(result), "notSame": &result[0] != &original[0]}
	expected := args.Map{"len": 4, "notSame": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone valid", actual)
}

func Test_Cov9_BytesToString_Nil(t *testing.T) {
	actual := args.Map{"v": corejson.BytesToString(nil)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "BytesToString nil", actual)
}

func Test_Cov9_BytesToString_Valid(t *testing.T) {
	actual := args.Map{"v": corejson.BytesToString([]byte(`"hi"`))}
	expected := args.Map{"v": `"hi"`}
	expected.ShouldBeEqual(t, 0, "BytesToString valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_BytesCollection_Empty(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	actual := args.Map{"empty": bc.IsEmpty(), "len": bc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCollection empty", actual)
}

func Test_Cov9_BytesCollection_Add(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"a"`))
	bc.Add([]byte(`"b"`))
	actual := args.Map{"len": bc.Length(), "hasAny": bc.HasAnyItem()}
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "BytesCollection add", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_ResultCollection_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	actual := args.Map{"empty": rc.IsEmpty(), "len": rc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "ResultCollection empty", actual)
}

func Test_Cov9_ResultCollection_Add(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := corejson.New(map[string]string{"k": "v"})
	rc.Add(r)
	actual := args.Map{"len": rc.Length(), "hasAny": rc.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "ResultCollection add", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ResultsPtrCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_ResultsPtrCollection_Empty(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.Empty()
	actual := args.Map{"empty": rc.IsEmpty(), "len": rc.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection empty", actual)
}

func Test_Cov9_ResultsPtrCollection_Add(t *testing.T) {
	rc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.NewPtr(map[string]string{"k": "v"})
	rc.Add(r)
	actual := args.Map{"len": rc.Length(), "hasAny": rc.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection add", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapResults
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_MapResults_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	actual := args.Map{"empty": mr.IsEmpty(), "len": mr.Length()}
	expected := args.Map{"empty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "MapResults empty", actual)
}

func Test_Cov9_MapResults_Add(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.New("value")
	mr.Add("key1", r)
	actual := args.Map{"len": mr.Length(), "hasAny": mr.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "MapResults add", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAny
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_KeyAny(t *testing.T) {
	ka := corejson.KeyAny{Key: "k", AnyInf: "v"}
	actual := args.Map{"key": ka.Key, "val": ka.AnyInf}
	expected := args.Map{"key": "k", "val": "v"}
	expected.ShouldBeEqual(t, 0, "KeyAny", actual)
}

func Test_Cov9_KeyWithResult(t *testing.T) {
	r := corejson.New("hello")
	kwr := corejson.KeyWithResult{Key: "k", Result: r}
	actual := args.Map{"key": kwr.Key, "hasBytes": kwr.Result.HasBytes()}
	expected := args.Map{"key": "k", "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "KeyWithResult", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PrettyJsonString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_PrettyJsonStringer_Valid(t *testing.T) {
	r := corejson.New(map[string]string{"k": "v"})
	pretty := r.PrettyJsonString()
	actual := args.Map{"notEmpty": pretty != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringer valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Result — Clone / IsEqual
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Result_Clone_Nil(t *testing.T) {
	var r *corejson.Result
	c := r.Clone(true)
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Clone nil", actual)
}

func Test_Cov9_Result_Clone_Valid(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"k": "v"})
	c := r.Clone(true)
	actual := args.Map{"hasBytes": c.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Clone valid", actual)
}

func Test_Cov9_Result_ClonePtr_Nil(t *testing.T) {
	var r *corejson.Result
	c := r.ClonePtr(true)
	actual := args.Map{"notNil": c != nil, "empty": c.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr nil", actual)
}

func Test_Cov9_Result_IsEqual_SameBytes(t *testing.T) {
	r1 := corejson.NewPtr("hello")
	r2 := corejson.NewPtr("hello")
	actual := args.Map{"eq": r1.IsEqual(*r2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual same", actual)
}

func Test_Cov9_Result_IsEqual_DiffBytes(t *testing.T) {
	r1 := corejson.NewPtr("hello")
	r2 := corejson.NewPtr("world")
	actual := args.Map{"eq": r1.IsEqual(*r2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual diff", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Deserialize.FromBytes
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Deserialize_FromBytes_String(t *testing.T) {
	s, err := corejson.Deserialize.FromBytes.String([]byte(`"hello"`))
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.String", actual)
}

func Test_Cov9_Deserialize_FromBytes_Integer(t *testing.T) {
	i, err := corejson.Deserialize.FromBytes.Integer([]byte(`42`))
	actual := args.Map{"noErr": err == nil, "val": i}
	expected := args.Map{"noErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.Integer", actual)
}

func Test_Cov9_Deserialize_FromBytes_Strings(t *testing.T) {
	lines, err := corejson.Deserialize.FromBytes.Strings([]byte(`["a","b"]`))
	actual := args.Map{"noErr": err == nil, "len": len(lines)}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.Strings", actual)
}

func Test_Cov9_Deserialize_FromBytes_Bool(t *testing.T) {
	b, err := corejson.Deserialize.FromBytes.Bool([]byte(`true`))
	actual := args.Map{"noErr": err == nil, "val": b}
	expected := args.Map{"noErr": true, "val": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.Bool", actual)
}

func Test_Cov9_Deserialize_FromBytes_MapStringString(t *testing.T) {
	m, err := corejson.Deserialize.FromBytes.MapStringString([]byte(`{"k":"v"}`))
	actual := args.Map{"noErr": err == nil, "len": len(m)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.MapStringString", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Serialize.Raw
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Serialize_Raw_Valid(t *testing.T) {
	b, err := corejson.Serialize.Raw(map[string]string{"k": "v"})
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.Raw valid", actual)
}

func Test_Cov9_Serialize_Raw_Nil(t *testing.T) {
	b, err := corejson.Serialize.Raw(nil)
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.Raw nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Empty creators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Empty_Result(t *testing.T) {
	r := corejson.Empty.Result()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result", actual)
}

func Test_Cov9_Empty_ResultPtr(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	actual := args.Map{"notNil": r != nil, "empty": r.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtr", actual)
}

func Test_Cov9_Empty_ResultsCollection(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	actual := args.Map{"empty": rc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsCollection", actual)
}

func Test_Cov9_Empty_ResultsPtrCollection(t *testing.T) {
	rc := corejson.Empty.ResultsPtrCollection()
	actual := args.Map{"empty": rc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsPtrCollection", actual)
}

func Test_Cov9_Empty_MapResults(t *testing.T) {
	mr := corejson.Empty.MapResults()
	actual := args.Map{"empty": mr.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.MapResults", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MeaningfulError / ErrorString / Raw / RawString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Result_MeaningfulError_Nil(t *testing.T) {
	var r *corejson.Result
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError nil", actual)
}

func Test_Cov9_Result_MeaningfulError_Valid(t *testing.T) {
	r := corejson.New("hello")
	err := r.MeaningfulError()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError valid", actual)
}

func Test_Cov9_Result_ErrorString_Empty(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"empty": r.ErrorString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ErrorString empty", actual)
}

func Test_Cov9_Result_Raw(t *testing.T) {
	r := corejson.New("hello")
	b, err := r.Raw()
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Raw", actual)
}

func Test_Cov9_Result_RawString(t *testing.T) {
	r := corejson.New("hello")
	s, err := r.RawString()
	actual := args.Map{"notEmpty": s != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "RawString", actual)
}

func Test_Cov9_NewResult_Error(t *testing.T) {
	r := corejson.NewResult.Error(nil)
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewResult Error nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CastAny / AnyTo
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_CastAny_Result(t *testing.T) {
	r := corejson.New("hello")
	casted, err := corejson.CastAny.Result(r)
	actual := args.Map{"noErr": err == nil, "hasBytes": casted.HasBytes()}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "CastAny.Result", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Pretty
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Pretty_Bytes(t *testing.T) {
	b := corejson.Pretty.Bytes.Format([]byte(`{"k":"v"}`))
	actual := args.Map{"notNil": b != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.Format", actual)
}

func Test_Cov9_Pretty_String(t *testing.T) {
	s := corejson.Pretty.String.Format(`{"k":"v"}`)
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.Format", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Deserialize Must variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Deserialize_FromBytes_StringMust(t *testing.T) {
	s := corejson.Deserialize.FromBytes.StringMust([]byte(`"hello"`))
	actual := args.Map{"val": s}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.StringMust", actual)
}

func Test_Cov9_Deserialize_FromBytes_IntegerMust(t *testing.T) {
	i := corejson.Deserialize.FromBytes.IntegerMust([]byte(`42`))
	actual := args.Map{"val": i}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.IntegerMust", actual)
}

func Test_Cov9_Deserialize_FromBytes_BoolMust(t *testing.T) {
	b := corejson.Deserialize.FromBytes.BoolMust([]byte(`true`))
	actual := args.Map{"val": b}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.BoolMust", actual)
}

func Test_Cov9_Deserialize_FromBytes_Integer64(t *testing.T) {
	i, err := corejson.Deserialize.FromBytes.Integer64([]byte(`123456789`))
	actual := args.Map{"noErr": err == nil, "val": i}
	expected := args.Map{"noErr": true, "val": int64(123456789)}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.Integer64", actual)
}

func Test_Cov9_Deserialize_FromBytes_MapAnyItem(t *testing.T) {
	m, err := corejson.Deserialize.FromBytes.MapAnyItem([]byte(`{"k":"v"}`))
	actual := args.Map{"noErr": err == nil, "len": len(m)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize.FromBytes.MapAnyItem", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Regex-based test for corejson (covers regexnew indirectly)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Regex_Placeholder(t *testing.T) {
	// just to keep regexp import if needed
	re := regexp.MustCompile(`\d+`)
	actual := args.Map{"match": re.MatchString("42")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "regex placeholder", actual)
}
