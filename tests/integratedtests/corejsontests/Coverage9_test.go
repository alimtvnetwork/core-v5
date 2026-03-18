package corejsontests

import (
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
// KeyAny / KeyWithJsoner / KeyWithResult
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_KeyAny(t *testing.T) {
	ka := corejson.KeyAny{Key: "k", Value: "v"}
	actual := args.Map{"key": ka.Key, "val": ka.Value}
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
// SimpleJsonBinder / SimpleJsoner / JsonString / JsonStringer
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_JsonString_Nil(t *testing.T) {
	var js *corejson.JsonString
	actual := args.Map{"nil": js == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "JsonString nil", actual)
}

func Test_Cov9_JsonString_Valid(t *testing.T) {
	js := corejson.JsonString(`{"a":1}`)
	actual := args.Map{"notEmpty": string(js) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// PrettyJsonStringer
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
// Deserialize functions
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_DeserializeFromBytesTo_Valid(t *testing.T) {
	var s string
	err := corejson.DeserializeFromBytesTo([]byte(`"hello"`), &s)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytesTo valid", actual)
}

func Test_Cov9_DeserializeFromBytesTo_Invalid(t *testing.T) {
	var s string
	err := corejson.DeserializeFromBytesTo([]byte(`{bad`), &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytesTo invalid", actual)
}

func Test_Cov9_DeserializeFromResultTo_Valid(t *testing.T) {
	r := corejson.New("hello")
	var s string
	err := corejson.DeserializeFromResultTo(&r, &s)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "DeserializeFromResultTo valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Serialize (anyTo)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Serialize_Valid(t *testing.T) {
	b, err := corejson.Serialize(map[string]string{"k": "v"})
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize valid", actual)
}

func Test_Cov9_Serialize_Nil(t *testing.T) {
	b, err := corejson.Serialize(nil)
	actual := args.Map{"hasBytes": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Empty creators
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov9_Empty_Result(t *testing.T) {
	r := corejson.EmptyResult.Value()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult", actual)
}

func Test_Cov9_Empty_ResultPtr(t *testing.T) {
	r := corejson.EmptyResult.Ptr()
	actual := args.Map{"notNil": r != nil, "empty": r.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult Ptr", actual)
}

func Test_Cov9_Empty_ResultBytesCollection(t *testing.T) {
	bc := corejson.EmptyResult.BytesCollection()
	actual := args.Map{"empty": bc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult BytesCollection", actual)
}

func Test_Cov9_Empty_ResultCollection(t *testing.T) {
	rc := corejson.EmptyResult.ResultCollection()
	actual := args.Map{"empty": rc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult ResultCollection", actual)
}

func Test_Cov9_Empty_ResultsPtrCollection(t *testing.T) {
	rc := corejson.EmptyResult.ResultsPtrCollection()
	actual := args.Map{"empty": rc.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult ResultsPtrCollection", actual)
}

func Test_Cov9_Empty_MapResults(t *testing.T) {
	mr := corejson.EmptyResult.MapResults()
	actual := args.Map{"empty": mr.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult MapResults", actual)
}
