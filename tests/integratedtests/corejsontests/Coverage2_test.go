package corejsontests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result creation ──

func Test_Cov2_NewResult_Serialize(t *testing.T) {
	data := map[string]string{"key": "value"}
	result := corejson.NewResult.Serialize(data)
	actual := args.Map{
		"notNil":    result != nil,
		"hasBytes":  result.HasBytes(),
		"noErr":     result.IsEmptyError(),
		"hasLength": result.Length() > 0,
	}
	expected := args.Map{
		"notNil":    true,
		"hasBytes":  true,
		"noErr":     true,
		"hasLength": true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult.Serialize returns valid -- map input", actual)
}

func Test_Cov2_Result_JsonString(t *testing.T) {
	result := corejson.NewResult.Serialize(map[string]int{"a": 1})
	js := result.JsonString()
	actual := args.Map{"hasContent": len(js) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Result.JsonString returns non-empty -- valid result", actual)
}

func Test_Cov2_Result_Map(t *testing.T) {
	result := corejson.NewResult.Serialize("hello")
	m := result.Map()
	actual := args.Map{"hasBytes": len(m) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result.Map returns map with entries -- valid result", actual)
}

func Test_Cov2_Result_Map_Nil(t *testing.T) {
	var result *corejson.Result
	m := result.Map()
	actual := args.Map{"isEmpty": len(m) == 0}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result.Map returns empty -- nil result", actual)
}

func Test_Cov2_Result_String(t *testing.T) {
	result := corejson.NewResult.Serialize("hello")
	s := result.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Result.String returns non-empty -- valid result", actual)
}

func Test_Cov2_Result_HasError(t *testing.T) {
	result := corejson.NewResult.Serialize("hello")
	actual := args.Map{
		"hasError":   result.HasError(),
		"emptyError": result.IsEmptyError(),
	}
	expected := args.Map{
		"hasError":   false,
		"emptyError": true,
	}
	expected.ShouldBeEqual(t, 0, "Result has no error -- valid input", actual)
}

func Test_Cov2_Result_Clone(t *testing.T) {
	result := corejson.NewResult.Serialize("hello")
	cloned := result.Clone()
	actual := args.Map{
		"sameJson": cloned.JsonString() == result.JsonString(),
	}
	expected := args.Map{
		"sameJson": true,
	}
	expected.ShouldBeEqual(t, 0, "Result.Clone produces equal json -- valid", actual)
}

func Test_Cov2_Result_ClonePtr(t *testing.T) {
	result := corejson.NewResult.Serialize("hello")
	cloned := result.ClonePtr()
	actual := args.Map{
		"notNil":     cloned != nil,
		"notSamePtr": cloned != result,
	}
	expected := args.Map{
		"notNil":     true,
		"notSamePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result.ClonePtr returns different ptr -- valid", actual)
}

func Test_Cov2_Result_Nil_ClonePtr(t *testing.T) {
	var result *corejson.Result
	cloned := result.ClonePtr()
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Result.ClonePtr returns nil -- nil receiver", actual)
}

// ── Serialize/Deserialize roundtrip ──

func Test_Cov2_SerializeDeserialize_Roundtrip(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}
	original := testStruct{Name: "test", Age: 25}

	rawBytes, err := corejson.Serialize.Raw(original)
	actual1 := args.Map{"noErr": err == nil, "hasBytes": len(rawBytes) > 0}
	expected1 := args.Map{"noErr": true, "hasBytes": true}
	expected1.ShouldBeEqual(t, 0, "Serialize.Raw succeeds -- struct input", actual1)

	var deserialized testStruct
	err = corejson.Deserialize.UsingBytes(rawBytes, &deserialized)
	actual2 := args.Map{
		"noErr":    err == nil,
		"sameName": deserialized.Name == original.Name,
		"sameAge":  deserialized.Age == original.Age,
	}
	expected2 := args.Map{
		"noErr":    true,
		"sameName": true,
		"sameAge":  true,
	}
	expected2.ShouldBeEqual(t, 1, "Deserialize.UsingBytes roundtrip -- struct", actual2)
}

// ── AnyTo ──

func Test_Cov2_AnyTo_SerializedJsonResult(t *testing.T) {
	result := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})
	actual := args.Map{
		"notNil":  result != nil,
		"noErr":   result.IsEmptyError(),
		"hasData": result.HasBytes(),
	}
	expected := args.Map{
		"notNil":  true,
		"noErr":   true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns valid -- map input", actual)
}

func Test_Cov2_AnyTo_SerializedRaw(t *testing.T) {
	rawBytes, err := corejson.AnyTo.SerializedRaw(map[string]string{"k": "v"})
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(rawBytes) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedRaw returns bytes -- map input", actual)
}

func Test_Cov2_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString(map[string]int{"a": 1})
	actual := args.Map{
		"noErr":      err == nil,
		"hasContent": len(s) > 0,
	}
	expected := args.Map{
		"noErr":      true,
		"hasContent": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedString returns json string -- map input", actual)
}

func Test_Cov2_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString(map[string]int{"a": 1})
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonString returns non-empty -- map input", actual)
}

func Test_Cov2_AnyTo_SafeJsonPrettyString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString(map[string]int{"a": 1})
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonPrettyString returns non-empty -- map input", actual)
}

func Test_Cov2_AnyTo_JsonString(t *testing.T) {
	s := corejson.AnyTo.JsonString(map[string]int{"a": 1})
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonString returns non-empty -- map input", actual)
}

func Test_Cov2_AnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust(map[string]int{"a": 1})
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonStringMust returns non-empty -- map input", actual)
}

func Test_Cov2_AnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust(map[string]int{"a": 1})
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringMust returns non-empty -- map input", actual)
}

// ── Deserialize from bytes ──

func Test_Cov2_DeserializeFromBytes_String(t *testing.T) {
	b, _ := json.Marshal("hello")
	s, err := corejson.Deserialize.FromBytesTo.String(b)
	actual := args.Map{"noErr": err == nil, "val": s}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.String roundtrip -- hello", actual)
}

func Test_Cov2_DeserializeFromBytes_Integer(t *testing.T) {
	b, _ := json.Marshal(42)
	val, err := corejson.Deserialize.FromBytesTo.Integer(b)
	actual := args.Map{"noErr": err == nil, "val": val}
	expected := args.Map{"noErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Integer roundtrip -- 42", actual)
}

func Test_Cov2_DeserializeFromBytes_Integer64(t *testing.T) {
	b, _ := json.Marshal(int64(999))
	val, err := corejson.Deserialize.FromBytesTo.Integer64(b)
	actual := args.Map{"noErr": err == nil, "val": val}
	expected := args.Map{"noErr": true, "val": int64(999)}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Integer64 roundtrip -- 999", actual)
}

func Test_Cov2_DeserializeFromBytes_MapAnyItem(t *testing.T) {
	b, _ := json.Marshal(map[string]any{"k": "v"})
	m, err := corejson.Deserialize.FromBytesTo.MapAnyItem(b)
	actual := args.Map{
		"noErr":  err == nil,
		"hasKey": m["k"] == "v",
	}
	expected := args.Map{
		"noErr":  true,
		"hasKey": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.MapAnyItem roundtrip -- map", actual)
}

func Test_Cov2_DeserializeFromBytes_MapStringString(t *testing.T) {
	b, _ := json.Marshal(map[string]string{"k": "v"})
	m, err := corejson.Deserialize.FromBytesTo.MapStringString(b)
	actual := args.Map{
		"noErr":  err == nil,
		"hasKey": m["k"] == "v",
	}
	expected := args.Map{
		"noErr":  true,
		"hasKey": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.MapStringString roundtrip -- map", actual)
}

func Test_Cov2_DeserializeFromBytes_Bytes(t *testing.T) {
	original := []byte{1, 2, 3}
	b, _ := json.Marshal(original)
	result, err := corejson.Deserialize.FromBytesTo.Bytes(b)
	actual := args.Map{
		"noErr":  err == nil,
		"len":    len(result),
	}
	expected := args.Map{
		"noErr":  true,
		"len":    3,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Bytes roundtrip -- 3 bytes", actual)
}

func Test_Cov2_DeserializeFromBytes_Integers(t *testing.T) {
	b, _ := json.Marshal([]int{1, 2, 3})
	val, err := corejson.Deserialize.FromBytesTo.Integers(b)
	actual := args.Map{"noErr": err == nil, "len": len(val)}
	expected := args.Map{"noErr": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Integers roundtrip -- 3 ints", actual)
}

// ── Empty creators ──

func Test_Cov2_Empty_Result(t *testing.T) {
	r := corejson.Empty.Result()
	actual := args.Map{"isEmpty": r.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result returns empty -- no data", actual)
}

func Test_Cov2_Empty_ResultsCollection(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	actual := args.Map{"isEmpty": rc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsCollection returns empty -- no data", actual)
}

func Test_Cov2_Empty_ResultsPtrCollection(t *testing.T) {
	rpc := corejson.Empty.ResultsPtrCollection()
	actual := args.Map{"isEmpty": rpc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsPtrCollection returns empty -- no data", actual)
}

func Test_Cov2_Empty_BytesCollection(t *testing.T) {
	bc := corejson.Empty.BytesCollection()
	actual := args.Map{"isEmpty": bc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.BytesCollection returns empty -- no data", actual)
}

func Test_Cov2_Empty_MapResults(t *testing.T) {
	mr := corejson.Empty.MapResults()
	actual := args.Map{"notNil": mr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.MapResults returns non-nil -- empty map", actual)
}

// ── Deserialize methods ──

func Test_Cov2_Deserialize_UsingString(t *testing.T) {
	type s struct{ Name string }
	var result s
	err := corejson.Deserialize.UsingString(`{"Name":"test"}`, &result)
	actual := args.Map{"noErr": err == nil, "name": result.Name}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingString works -- json string", actual)
}

func Test_Cov2_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	type s struct{ Name string }
	var result s
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &result)
	actual := args.Map{"noErr": err == nil, "empty": result.Name == ""}
	expected := args.Map{"noErr": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Deserialize.UsingStringIgnoreEmpty skips empty -- empty string", actual)
}

func Test_Cov2_Deserialize_Apply(t *testing.T) {
	jsonResult := corejson.NewResult.Serialize(map[string]string{"Name": "test"})
	var result map[string]string
	err := corejson.Deserialize.Apply(jsonResult, &result)
	actual := args.Map{"noErr": err == nil, "name": result["Name"]}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "Deserialize.Apply works -- from Result", actual)
}

// ── CastAny ──

func Test_Cov2_CastAny_FromToOption_SameType(t *testing.T) {
	src := map[string]int{"a": 1}
	var dest map[string]int
	err := corejson.CastAny.FromToOption(src, &dest, false)
	actual := args.Map{
		"noErr": err == nil,
		"val":   dest["a"],
	}
	expected := args.Map{
		"noErr": true,
		"val":   1,
	}
	expected.ShouldBeEqual(t, 0, "CastAny.FromToOption works -- map to map", actual)
}

// ── Result HasBytes/IsEmpty variants ──

func Test_Cov2_Result_IsEmpty_NilResult(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{
		"isEmpty":    r.IsEmpty(),
		"hasBytes":   r.HasBytes(),
		"length":     r.Length(),
		"emptyError": r.IsEmptyError(),
	}
	expected := args.Map{
		"isEmpty":    true,
		"hasBytes":   false,
		"length":     0,
		"emptyError": true,
	}
	expected.ShouldBeEqual(t, 0, "Result nil receiver -- all safe methods", actual)
}

// ── ResultCollection ──

func Test_Cov2_ResultCollection_Basic(t *testing.T) {
	rc := corejson.NewResultsCollection.Cap(5)
	r1 := corejson.NewResult.Serialize("hello")
	rc.Add(*r1)
	actual := args.Map{
		"length":    rc.Length(),
		"hasAny":    rc.HasAnyItem(),
		"isEmpty":   rc.IsEmpty(),
		"lastIndex": rc.LastIndex(),
	}
	expected := args.Map{
		"length":    1,
		"hasAny":    true,
		"isEmpty":   false,
		"lastIndex": 0,
	}
	expected.ShouldBeEqual(t, 0, "ResultCollection basic ops -- single item", actual)
}

func Test_Cov2_ResultCollection_FirstLast(t *testing.T) {
	rc := corejson.NewResultsCollection.Cap(5)
	r1 := corejson.NewResult.Serialize("first")
	r2 := corejson.NewResult.Serialize("last")
	rc.Add(*r1)
	rc.Add(*r2)
	first := rc.First()
	last := rc.Last()
	actual := args.Map{
		"firstNotNil": first.HasBytes(),
		"lastNotNil":  last.HasBytes(),
	}
	expected := args.Map{
		"firstNotNil": true,
		"lastNotNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "ResultCollection First/Last work -- two items", actual)
}

func Test_Cov2_ResultCollection_Clone(t *testing.T) {
	rc := corejson.NewResultsCollection.Cap(5)
	r1 := corejson.NewResult.Serialize("hello")
	rc.Add(*r1)
	cloned := rc.Clone()
	actual := args.Map{"sameLen": cloned.Length() == rc.Length()}
	expected := args.Map{"sameLen": true}
	expected.ShouldBeEqual(t, 0, "ResultCollection.Clone returns same len -- single item", actual)
}

// ── ResultsPtrCollection ──

func Test_Cov2_ResultsPtrCollection_Basic(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(5)
	r1 := corejson.NewResult.Serialize("hello")
	rpc.Add(r1)
	actual := args.Map{
		"length":    rpc.Length(),
		"hasAny":    rpc.HasAnyItem(),
		"isEmpty":   rpc.IsEmpty(),
		"lastIndex": rpc.LastIndex(),
	}
	expected := args.Map{
		"length":    1,
		"hasAny":    true,
		"isEmpty":   false,
		"lastIndex": 0,
	}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection basic ops -- single item", actual)
}

func Test_Cov2_ResultsPtrCollection_FirstLastOrDefault(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(5)
	r1 := corejson.NewResult.Serialize("hello")
	rpc.Add(r1)
	actual := args.Map{
		"firstNotNil": rpc.FirstOrDefault() != nil,
		"lastNotNil":  rpc.LastOrDefault() != nil,
	}
	expected := args.Map{
		"firstNotNil": true,
		"lastNotNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection FirstOrDefault/LastOrDefault -- single item", actual)
}

func Test_Cov2_ResultsPtrCollection_Empty_FirstLastOrDefault(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(0)
	actual := args.Map{
		"firstNil": rpc.FirstOrDefault() == nil,
		"lastNil":  rpc.LastOrDefault() == nil,
	}
	expected := args.Map{
		"firstNil": true,
		"lastNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection FirstOrDefault returns nil -- empty", actual)
}

func Test_Cov2_ResultsPtrCollection_SkipTakeLimit(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(5)
	for i := 0; i < 3; i++ {
		rpc.Add(corejson.NewResult.Serialize(i))
	}
	actual := args.Map{
		"skipLen":  len(rpc.Skip(1)),
		"takeLen":  len(rpc.Take(2)),
		"limitLen": len(rpc.Limit(1)),
	}
	expected := args.Map{
		"skipLen":  2,
		"takeLen":  2,
		"limitLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection Skip/Take/Limit -- 3 items", actual)
}

func Test_Cov2_ResultsPtrCollection_AddSkipOnNil(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(5)
	rpc.AddSkipOnNil(nil)
	rpc.AddSkipOnNil(corejson.NewResult.Serialize("hello"))
	actual := args.Map{"length": rpc.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection.AddSkipOnNil skips nil -- 1 valid", actual)
}

func Test_Cov2_ResultsPtrCollection_HasError(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(5)
	rpc.Add(corejson.NewResult.Serialize("hello"))
	actual := args.Map{"hasError": rpc.HasError()}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection.HasError returns false -- no errors", actual)
}

func Test_Cov2_ResultsPtrCollection_GetStrings(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(5)
	rpc.Add(corejson.NewResult.Serialize("hello"))
	strs := rpc.GetStrings()
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection.GetStrings returns 1 -- single item", actual)
}

func Test_Cov2_ResultsPtrCollection_ClearDispose(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Cap(5)
	rpc.Add(corejson.NewResult.Serialize("hello"))
	rpc.Clear()
	actual := args.Map{"isEmpty": rpc.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "ResultsPtrCollection.Clear empties -- after clear", actual)
}
