package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result ──

func Test_Cov7_Result_Basic(t *testing.T) {
	result := corejson.New(map[string]int{"a": 1})
	actual := args.Map{
		"hasErr":    result.HasError(),
		"isEmpty":   result.IsEmpty(),
		"hasBytes":  len(result.Bytes) > 0,
		"typeName":  result.TypeName != "",
		"jsonStr":   result.JsonString() != "",
		"prettyStr": result.PrettyJsonString() != "",
	}
	expected := args.Map{
		"hasErr": false, "isEmpty": false, "hasBytes": true,
		"typeName": true, "jsonStr": true, "prettyStr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns correct value -- basic", actual)
}

func Test_Cov7_Result_Ptr(t *testing.T) {
	result := corejson.NewPtr(map[string]int{"a": 1})
	actual := args.Map{"notNil": result != nil, "noErr": !result.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewPtr returns correct value -- with args", actual)
}

func Test_Cov7_Result_SafeBytes(t *testing.T) {
	result := corejson.New(map[string]int{"a": 1})
	actual := args.Map{"hasBytes": len(result.SafeBytes()) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "SafeBytes returns correct value -- with args", actual)
}

func Test_Cov7_Result_Unmarshal(t *testing.T) {
	result := corejson.New(map[string]int{"a": 1})
	var target map[string]int
	err := result.Unmarshal(&target)
	actual := args.Map{"noErr": err == nil, "a": target["a"]}
	expected := args.Map{"noErr": true, "a": 1}
	expected.ShouldBeEqual(t, 0, "Unmarshal returns correct value -- with args", actual)
}

func Test_Cov7_Result_Deserialize(t *testing.T) {
	result := corejson.New(map[string]int{"b": 2})
	var target map[string]int
	err := result.Deserialize(&target)
	actual := args.Map{"noErr": err == nil, "b": target["b"]}
	expected := args.Map{"noErr": true, "b": 2}
	expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- with args", actual)
}

func Test_Cov7_Result_MeaningfulError_NoError(t *testing.T) {
	result := corejson.New("hello")
	err := result.MeaningfulError()
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError returns nil -- nil", actual)
}

// ── Serialize ──

func Test_Cov7_Serialize_Raw(t *testing.T) {
	bytes, err := corejson.Serialize.Raw(map[string]int{"a": 1})
	actual := args.Map{"noErr": err == nil, "hasBytes": len(bytes) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize.Raw returns correct value -- with args", actual)
}

func Test_Cov7_Serialize_UsingAny(t *testing.T) {
	result := corejson.Serialize.UsingAny(map[string]int{"a": 1})
	actual := args.Map{"noErr": !result.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize.UsingAny returns correct value -- with args", actual)
}

// ── Deserialize ──

func Test_Cov7_Deserialize_UsingBytes(t *testing.T) {
	var target map[string]int
	err := corejson.Deserialize.UsingBytes([]byte(`{"a":1}`), &target)
	actual := args.Map{"noErr": err == nil, "a": target["a"]}
	expected := args.Map{"noErr": true, "a": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytes returns correct value -- with args", actual)
}

func Test_Cov7_Deserialize_UsingBytes_Invalid(t *testing.T) {
	var target map[string]int
	err := corejson.Deserialize.UsingBytes([]byte("invalid"), &target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingBytes returns error -- invalid", actual)
}

func Test_Cov7_Deserialize_UsingResult(t *testing.T) {
	result := corejson.New(map[string]int{"a": 1})
	var target map[string]int
	err := corejson.Deserialize.UsingResult(&result, &target)
	actual := args.Map{"noErr": err == nil, "a": target["a"]}
	expected := args.Map{"noErr": true, "a": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingResult returns correct value -- with args", actual)
}

func Test_Cov7_Deserialize_Apply(t *testing.T) {
	result := corejson.New(map[string]int{"a": 1})
	var target map[string]int
	err := corejson.Deserialize.Apply(&result, &target)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.Apply returns correct value -- with args", actual)
}

// ── BytesDeepClone ──

func Test_Cov7_BytesDeepClone(t *testing.T) {
	original := []byte("hello")
	cloned := corejson.BytesDeepClone(original)
	original[0] = 'X'
	actual := args.Map{"different": string(cloned) == "hello"}
	expected := args.Map{"different": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns correct value -- with args", actual)
}

func Test_Cov7_BytesDeepClone_Nil(t *testing.T) {
	cloned := corejson.BytesDeepClone(nil)
	actual := args.Map{"isEmpty": len(cloned) == 0}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns nil -- nil", actual)
}

// ── BytesToString ──

func Test_Cov7_BytesToString(t *testing.T) {
	result := corejson.BytesToString([]byte("hello"))
	actual := args.Map{"str": result}
	expected := args.Map{"str": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesToString returns correct value -- with args", actual)
}

func Test_Cov7_BytesToPrettyString(t *testing.T) {
	result := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString returns correct value -- with args", actual)
}

// ── Empty ──

func Test_Cov7_Empty_ResultPtr(t *testing.T) {
	result := corejson.Empty.ResultPtr()
	actual := args.Map{"notNil": result != nil, "isEmpty": result.IsEmpty()}
	expected := args.Map{"notNil": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtr returns empty -- with args", actual)
}

func Test_Cov7_Empty_Result(t *testing.T) {
	result := corejson.Empty.Result()
	actual := args.Map{"isEmpty": result.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result returns empty -- with args", actual)
}

// ── NewResult ──

func Test_Cov7_NewResult_UsingTypeBytesPtr(t *testing.T) {
	result := corejson.NewResult.UsingTypeBytesPtr("TestType", []byte(`{"a":1}`))
	actual := args.Map{"notNil": result != nil, "noErr": !result.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult.UsingTypeBytesPtr returns correct value -- with args", actual)
}

// ── BytesCloneIf ──

func Test_Cov7_BytesCloneIf_True(t *testing.T) {
	original := []byte("data")
	cloned := corejson.BytesCloneIf(true, original)
	original[0] = 'X'
	actual := args.Map{"different": string(cloned) == "data"}
	expected := args.Map{"different": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns non-empty -- true", actual)
}

func Test_Cov7_BytesCloneIf_False(t *testing.T) {
	cloned := corejson.BytesCloneIf(false, []byte("data"))
	actual := args.Map{"isEmpty": len(cloned) == 0}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns non-empty -- false", actual)
}

// ── AnyTo ──

func Test_Cov7_AnyTo_SerializedJsonResult(t *testing.T) {
	result := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})
	actual := args.Map{"noErr": !result.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- with args", actual)
}

func Test_Cov7_AnyTo_SerializedJsonResult_Bytes(t *testing.T) {
	result := corejson.AnyTo.SerializedJsonResult([]byte(`{"a":1}`))
	actual := args.Map{"noErr": !result.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- bytes", actual)
}

func Test_Cov7_AnyTo_SerializedJsonResult_String(t *testing.T) {
	result := corejson.AnyTo.SerializedJsonResult(`{"a":1}`)
	actual := args.Map{"noErr": !result.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns correct value -- string", actual)
}

// ── CastAny ──

func Test_Cov7_CastAny_FromToDefault(t *testing.T) {
	source := map[string]int{"a": 1}
	var target map[string]int
	err := corejson.CastAny.FromToDefault(source, &target)
	actual := args.Map{"noErr": err == nil, "a": target["a"]}
	expected := args.Map{"noErr": true, "a": 1}
	expected.ShouldBeEqual(t, 0, "CastAny.FromToDefault returns correct value -- with args", actual)
}

// ── BytesCollection ──

func Test_Cov7_BytesCollection_Empty(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	actual := args.Map{"isEmpty": bc.IsEmpty(), "len": bc.Length()}
	expected := args.Map{"isEmpty": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCollection returns empty -- empty", actual)
}

func Test_Cov7_BytesCollection_Add(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte("hello"))
	bc.Add([]byte("world"))
	actual := args.Map{"len": bc.Length(), "hasAny": bc.HasAnyItem()}
	expected := args.Map{"len": 2, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "BytesCollection returns correct value -- add", actual)
}

// ── ResultCollection ──

func Test_Cov7_ResultCollection_Empty(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	actual := args.Map{"isEmpty": rc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultCollection returns empty -- empty", actual)
}

func Test_Cov7_ResultCollection_Add(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	result := corejson.New("hello")
	rc.Add(result)
	actual := args.Map{"len": rc.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultCollection returns correct value -- add", actual)
}

// ── JsonString / JsonStringer ──

func Test_Cov7_JsonString(t *testing.T) {
	js, err := corejson.JsonString(`{"a":1}`)
	actual := args.Map{"str": js, "noErr": err == nil}
	expected := args.Map{"str": `"{\"a\":1}"`, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

// ── MapResults ──

func Test_Cov7_MapResults_Empty(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	actual := args.Map{"isEmpty": mr.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapResults returns empty -- empty", actual)
}

func Test_Cov7_MapResults_Add(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	result := corejson.New("hello")
	mr.Add("key1", result)
	actual := args.Map{"len": mr.Length(), "hasAny": mr.HasAnyItem()}
	expected := args.Map{"len": 1, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "MapResults returns correct value -- add", actual)
}
