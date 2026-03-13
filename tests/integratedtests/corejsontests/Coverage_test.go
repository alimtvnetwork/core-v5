package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── New / NewPtr ──

func Test_New_Simple_Cov(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"hasError": r.Error != nil, "isEmpty": r.IsEmpty(), "hasBytes": len(r.Bytes) > 0, "typeName": r.TypeName != ""}
	expected := args.Map{"hasError": false, "isEmpty": false, "hasBytes": true, "typeName": true}
	expected.ShouldBeEqual(t, 0, "New_Simple", actual)
}

func Test_NewPtr_Simple_Cov(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"isNil": r == nil, "hasError": r.Error != nil}
	expected := args.Map{"isNil": false, "hasError": false}
	expected.ShouldBeEqual(t, 0, "NewPtr_Simple", actual)
}

func Test_New_Struct_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	actual := args.Map{"hasError": r.Error != nil, "hasBytes": len(r.Bytes) > 0, "stringNotEmpty": r.String() != "", "jsonStringNotNil": r.JsonString() != ""}
	expected := args.Map{"hasError": false, "hasBytes": true, "stringNotEmpty": true, "jsonStringNotNil": true}
	expected.ShouldBeEqual(t, 0, "New_Struct", actual)
}

func Test_New_Nil_Cov(t *testing.T) {
	actual := args.Map{"hasError": corejson.New(nil).Error != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "New_Nil", actual)
}

// ── Result methods ──

func Test_Result_IsEmpty_Cov(t *testing.T) {
	empty := corejson.Result{}
	actual := args.Map{
		"isEmpty": empty.IsEmpty(), "isEmptyError": empty.IsEmptyError() != nil,
		"hasError": empty.HasError(), "hasNoError": !empty.HasError(), "isValid": empty.IsValid(),
	}
	expected := args.Map{
		"isEmpty": true, "isEmptyError": true,
		"hasError": false, "hasNoError": true, "isValid": false,
	}
	expected.ShouldBeEqual(t, 0, "Result_IsEmpty", actual)
}

func Test_Result_String_Cov(t *testing.T) {
	actual := args.Map{"notEmpty": corejson.New(42).String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result_String", actual)
}

func Test_Result_JsonString_Cov(t *testing.T) {
	actual := args.Map{"result": corejson.New(42).JsonString()}
	expected := args.Map{"result": "42"}
	expected.ShouldBeEqual(t, 0, "Result_JsonString", actual)
}

func Test_Result_SafeBytes_Cov(t *testing.T) {
	actual := args.Map{"hasBytes": len(corejson.New(42).SafeBytes()) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result_SafeBytes", actual)
}

func Test_Result_SafeBytes_Empty_Cov(t *testing.T) {
	r := corejson.NewPtr(nil)
	actual := args.Map{"notNil": r.SafeBytes() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Result_SafeBytes_Empty", actual)
}

func Test_Result_MustBytesJson_Cov(t *testing.T) {
	r := corejson.NewPtr(42)
	actual := args.Map{"hasBytes": len(r.SafeBytes()) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result_MustBytes", actual)
}

func Test_Result_PrettyJsonString_Cov(t *testing.T) {
	r := corejson.NewPtr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": r.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result_PrettyJsonString", actual)
}

func Test_Result_PrettyJsonBytes_Cov(t *testing.T) {
	r := corejson.NewPtr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": r.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result_PrettyJsonBytes", actual)
}

func Test_Result_Clone_Cov(t *testing.T) {
	c := corejson.New(42).Clone()
	actual := args.Map{"hasError": c.Error != nil, "hasBytes": len(c.Bytes) > 0}
	expected := args.Map{"hasError": false, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result_Clone", actual)
}

func Test_Result_ClonePtr_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.New(42).ClonePtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Result_ClonePtr", actual)
}

func Test_Result_IsEqual_Cov(t *testing.T) {
	r1 := corejson.New(42)
	r2 := corejson.New(42)
	r3 := corejson.New(43)
	actual := args.Map{"same": r1.IsEqual(&r2), "diff": r1.IsEqual(&r3), "nil": r1.IsEqual(nil)}
	expected := args.Map{"same": true, "diff": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "Result_IsEqual", actual)
}

func Test_Result_Unmarshal_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	var out testS
	err := r.Unmarshal(&out)
	actual := args.Map{"hasErr": err != nil, "val": out.A}
	expected := args.Map{"hasErr": false, "val": 42}
	expected.ShouldBeEqual(t, 0, "Result_Unmarshal", actual)
}

// ── Empty creator ──

func Test_EmptyResult_Cov(t *testing.T) {
	actual := args.Map{"isEmpty": corejson.Empty.Result().IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResult", actual)
}

func Test_EmptyResultPtr_Cov(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	actual := args.Map{"notNil": r != nil, "isEmpty": r.IsEmpty()}
	expected := args.Map{"notNil": true, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "EmptyResultPtr", actual)
}

// ── Serialize / Deserialize ──

func Test_Serialize_Default_Cov(t *testing.T) {
	actual := args.Map{"hasErr": corejson.Serialize.Default(42).HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Serialize_Default", actual)
}

func Test_Serialize_DefaultPtr_Cov(t *testing.T) {
	r := corejson.Serialize.DefaultPtr(42)
	actual := args.Map{"notNil": r != nil, "hasErr": r.HasError()}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "Serialize_DefaultPtr", actual)
}

func Test_Deserialize_FromResult_Cov(t *testing.T) {
	type testS struct{ A int }
	r := corejson.New(testS{A: 42})
	var out testS
	err := corejson.Deserialize.FromResult(r, &out)
	actual := args.Map{"hasErr": err != nil, "val": out.A}
	expected := args.Map{"hasErr": false, "val": 42}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromResult", actual)
}

func Test_Deserialize_FromBytes_Cov(t *testing.T) {
	type testS struct{ A int }
	var out testS
	err := corejson.Deserialize.FromBytes([]byte(`{"A":42}`), &out)
	actual := args.Map{"hasErr": err != nil, "val": out.A}
	expected := args.Map{"hasErr": false, "val": 42}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromBytes", actual)
}

// ── CastAny / AnyTo / NewResult ──

func Test_CastAny_ToString_Cov(t *testing.T) {
	actual := args.Map{"notEmpty": corejson.CastAny.ToString(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CastAny_ToString", actual)
}

func Test_AnyTo_Result_Cov(t *testing.T) {
	actual := args.Map{"hasErr": corejson.AnyTo.Result(42).HasError()}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyTo_Result", actual)
}

func Test_AnyTo_ResultPtr_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.AnyTo.ResultPtr(42) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_ResultPtr", actual)
}

func Test_NewResult_UsingBytes_Cov(t *testing.T) {
	actual := args.Map{"notEmpty": !corejson.NewResult.UsingBytes([]byte(`"hello"`)).IsEmpty()}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytes", actual)
}

func Test_NewResult_UsingBytesPtr_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewResult.UsingBytesPtr([]byte(`"hello"`)) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesPtr", actual)
}

// ── Collections ──

func Test_NewBytesCollection_Cap_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewBytesCollection.Cap(5) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewBytesCollection_Cap", actual)
}

func Test_NewResultsCollection_Cap_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewResultsCollection.Cap(5) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsCollection_Cap", actual)
}

func Test_NewResultsPtrCollection_Cap_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewResultsPtrCollection.Cap(5) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResultsPtrCollection_Cap", actual)
}

// ── Pretty ──

func Test_Pretty_FromBytes_Cov(t *testing.T) {
	actual := args.Map{"hasBytes": len(corejson.Pretty.FromBytes([]byte(`{"a":1}`))) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Pretty_FromBytes", actual)
}

func Test_Pretty_FromString_Cov(t *testing.T) {
	actual := args.Map{"notEmpty": corejson.Pretty.FromString(`{"a":1}`) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty_FromString", actual)
}

// ── NewMapResults ──

func Test_NewMapResults_Empty_Cov(t *testing.T) {
	actual := args.Map{"notNil": corejson.NewMapResults.Empty() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMapResults_Empty", actual)
}
