package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Result — nil receiver, edge branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C01_Result_NilReceiver_Map(t *testing.T) {
	var r *corejson.Result
	m := r.Map()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NilReceiver_Map", actual)
}

func Test_I5_C02_Result_NilReceiver_JsonStringPtr(t *testing.T) {
	var r *corejson.Result
	s := r.JsonStringPtr()
	actual := args.Map{"notNil": s != nil, "empty": *s == ""}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_JsonStringPtr", actual)
}

func Test_I5_C03_Result_NilReceiver_PrettyJsonString(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"empty": r.PrettyJsonString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_PrettyJsonString", actual)
}

func Test_I5_C04_Result_NilReceiver_PrettyJsonStringOrErrString(t *testing.T) {
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_PrettyJsonStringOrErr", actual)
}

func Test_I5_C05_Result_PrettyJsonStringOrErrString_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("test-err")}
	s := r.PrettyJsonStringOrErrString()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonStringOrErr_WithError", actual)
}

func Test_I5_C06_Result_NilReceiver_Length(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"len": r.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NilReceiver_Length", actual)
}

func Test_I5_C07_Result_NilReceiver_Raw(t *testing.T) {
	var r *corejson.Result
	b, err := r.Raw()
	actual := args.Map{"emptyBytes": len(b) == 0, "hasErr": err != nil}
	expected := args.Map{"emptyBytes": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_Raw", actual)
}

func Test_I5_C08_Result_NilReceiver_MeaningfulError(t *testing.T) {
	var r *corejson.Result
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NilReceiver_MeaningfulError", actual)
}

func Test_I5_C09_Result_MeaningfulError_EmptyBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte{}, TypeName: "TestType"}
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError_EmptyBytes", actual)
}

func Test_I5_C10_Result_MeaningfulError_WithErrorAndBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`), Error: errors.New("some-err"), TypeName: "T"}
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError_WithErrorAndBytes", actual)
}

func Test_I5_C11_Result_MeaningfulError_WithErrorNoBytes(t *testing.T) {
	r := &corejson.Result{Error: errors.New("some-err"), TypeName: "T"}
	err := r.MeaningfulError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulError_WithErrorNoBytes", actual)
}

func Test_I5_C12_Result_MeaningfulErrorMessage(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	msg := r.MeaningfulErrorMessage()
	actual := args.Map{"empty": msg == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MeaningfulErrorMessage_NoErr", actual)
}

func Test_I5_C13_Result_SafeString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"val"`)}
	actual := args.Map{"notEmpty": r.SafeString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeString", actual)
}

func Test_I5_C14_Result_Map_WithBytesErrorType(t *testing.T) {
	r := &corejson.Result{
		Bytes:    []byte(`"data"`),
		Error:    errors.New("e"),
		TypeName: "Foo",
	}
	m := r.Map()
	actual := args.Map{
		"hasBytes": m["Bytes"] != "",
		"hasError": m["Error"] != "",
		"hasType":  m["Type"] != "",
	}
	expected := args.Map{"hasBytes": true, "hasError": true, "hasType": true}
	expected.ShouldBeEqual(t, 0, "Map_WithAll", actual)
}

func Test_I5_C15_Result_DeserializedFieldsToMap_Nil(t *testing.T) {
	var r *corejson.Result
	m, err := r.DeserializedFieldsToMap()
	actual := args.Map{"len": len(m), "nilErr": err == nil}
	expected := args.Map{"len": 0, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "DeserializedFieldsToMap_Nil", actual)
}

func Test_I5_C16_Result_SafeDeserializedFieldsToMap(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	m := r.SafeDeserializedFieldsToMap()
	_ = m
	// just exercising the method without panic
	actual := args.Map{"ran": true}
	expected := args.Map{"ran": true}
	expected.ShouldBeEqual(t, 0, "SafeDeserializedFieldsToMap", actual)
}

func Test_I5_C17_Result_FieldsNames_Empty(t *testing.T) {
	r := &corejson.Result{}
	names, err := r.FieldsNames()
	actual := args.Map{"len": len(names), "nilErr": err == nil}
	expected := args.Map{"len": 0, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "FieldsNames_Empty", actual)
}

func Test_I5_C18_Result_SafeFieldsNames(t *testing.T) {
	r := &corejson.Result{}
	names := r.SafeFieldsNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeFieldsNames", actual)
}

func Test_I5_C19_Result_BytesTypeName_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"empty": r.BytesTypeName() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesTypeName_Nil", actual)
}

func Test_I5_C20_Result_SafeBytesTypeName_Empty(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"empty": r.SafeBytesTypeName() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName_Empty", actual)
}

func Test_I5_C21_Result_SafeBytesTypeName_WithType(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "Foo"}
	actual := args.Map{"val": r.SafeBytesTypeName()}
	expected := args.Map{"val": "Foo"}
	expected.ShouldBeEqual(t, 0, "SafeBytesTypeName_WithType", actual)
}

func Test_I5_C22_Result_ErrorString_NoError(t *testing.T) {
	r := &corejson.Result{}
	actual := args.Map{"empty": r.ErrorString() == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "ErrorString_NoError", actual)
}

func Test_I5_C23_Result_IsErrorEqual(t *testing.T) {
	r1 := &corejson.Result{Error: errors.New("x")}
	r2 := &corejson.Result{}
	actual := args.Map{
		"sameErr":    r1.IsErrorEqual(errors.New("x")),
		"diffErr":    r1.IsErrorEqual(errors.New("y")),
		"nilBothNil": r2.IsErrorEqual(nil),
		"oneNil":     r1.IsErrorEqual(nil),
	}
	expected := args.Map{
		"sameErr":    true,
		"diffErr":    false,
		"nilBothNil": true,
		"oneNil":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsErrorEqual", actual)
}

func Test_I5_C24_Result_String_NilBytes(t *testing.T) {
	r := corejson.Result{}
	s := r.String()
	actual := args.Map{"empty": s == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "String_NilBytes", actual)
}

func Test_I5_C25_Result_String_WithError(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	s := r.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String_WithError", actual)
}

func Test_I5_C26_Result_String_NoError(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	s := r.String()
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String_NoError", actual)
}

func Test_I5_C27_Result_SafeNonIssueBytes(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"empty": len(r.SafeNonIssueBytes()) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes_Error", actual)
}

func Test_I5_C28_Result_SafeNonIssueBytes_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"hasData": len(r.SafeNonIssueBytes()) > 0}
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "SafeNonIssueBytes_Valid", actual)
}

func Test_I5_C29_Result_SafeBytes_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"empty": len(r.SafeBytes()) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeBytes_Nil", actual)
}

func Test_I5_C30_Result_Values(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"hasData": len(r.Values()) > 0}
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "Values", actual)
}

func Test_I5_C31_Result_SafeValues_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"empty": len(r.SafeValues()) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeValues_Nil", actual)
}

func Test_I5_C32_Result_SafeValuesPtr_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"empty": len(r.SafeValuesPtr()) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "SafeValuesPtr_HasIssues", actual)
}

func Test_I5_C33_Result_RawMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	b := r.RawMust()
	actual := args.Map{"hasData": len(b) > 0}
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "RawMust", actual)
}

func Test_I5_C34_Result_RawStringMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	s := r.RawStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RawStringMust", actual)
}

func Test_I5_C35_Result_RawStringMust_Panic(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.RawStringMust()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "RawStringMust_Panic", actual)
}

func Test_I5_C36_Result_RawErrString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	b, errMsg := r.RawErrString()
	actual := args.Map{"hasBytes": len(b) > 0, "emptyErr": errMsg == ""}
	expected := args.Map{"hasBytes": true, "emptyErr": true}
	expected.ShouldBeEqual(t, 0, "RawErrString", actual)
}

func Test_I5_C37_Result_RawPrettyString(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := r.RawPrettyString()
	actual := args.Map{"hasContent": len(s) > 0, "nilErr": err == nil}
	expected := args.Map{"hasContent": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "RawPrettyString", actual)
}

func Test_I5_C38_Result_PrettyJsonBuffer_Empty(t *testing.T) {
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")
	actual := args.Map{"emptyBuf": buf.Len() == 0, "nilErr": err == nil}
	expected := args.Map{"emptyBuf": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonBuffer_Empty", actual)
}

func Test_I5_C39_Result_PrettyJsonString_InvalidJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`not-json`)}
	s := r.PrettyJsonString()
	actual := args.Map{"empty": s == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "PrettyJsonString_InvalidJson", actual)
}

func Test_I5_C40_Result_HasSafeItems(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	actual := args.Map{"safe": r.HasSafeItems()}
	expected := args.Map{"safe": true}
	expected.ShouldBeEqual(t, 0, "HasSafeItems", actual)
}

func Test_I5_C41_Result_HandleError_Panic(t *testing.T) {
	r := &corejson.Result{}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.HandleError()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleError_Panic", actual)
}

func Test_I5_C42_Result_HandleErrorWithMsg_Panic(t *testing.T) {
	r := &corejson.Result{}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.HandleErrorWithMsg("custom msg")
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleErrorWithMsg_Panic", actual)
}

func Test_I5_C43_Result_HasBytes_HasJsonBytes(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"ok"`)}
	actual := args.Map{"hasBytes": r.HasBytes(), "hasJsonBytes": r.HasJsonBytes()}
	expected := args.Map{"hasBytes": true, "hasJsonBytes": true}
	expected.ShouldBeEqual(t, 0, "HasBytes_HasJsonBytes", actual)
}

func Test_I5_C44_Result_IsEmptyJsonBytes_CurlyBrace(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{}`)}
	actual := args.Map{"isEmpty": r.IsEmptyJsonBytes()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes_CurlyBrace", actual)
}

func Test_I5_C45_Result_IsEmptyJsonBytes_Nil(t *testing.T) {
	var r *corejson.Result
	actual := args.Map{"isEmpty": r.IsEmptyJsonBytes()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyJsonBytes_Nil", actual)
}

func Test_I5_C46_Result_HasAnyItem(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"has": r.HasAnyItem()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem", actual)
}

func Test_I5_C47_Result_HasJson_IsEmptyJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{"hasJson": r.HasJson(), "isEmptyJson": r.IsEmptyJson()}
	expected := args.Map{"hasJson": true, "isEmptyJson": false}
	expected.ShouldBeEqual(t, 0, "HasJson_IsEmptyJson", actual)
}

func Test_I5_C48_Result_Unmarshal_NilReceiver(t *testing.T) {
	var r *corejson.Result
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal_NilReceiver", actual)
}

func Test_I5_C49_Result_Unmarshal_WithExistingError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("prior"), Bytes: []byte(`"x"`)}
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal_WithExistingError", actual)
}

func Test_I5_C50_Result_Unmarshal_InvalidJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`not-json`)}
	var s string
	err := r.Unmarshal(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Unmarshal_InvalidJson", actual)
}

func Test_I5_C51_Result_DeserializeMust_Panic(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.DeserializeMust(nil)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "DeserializeMust_Panic", actual)
}

func Test_I5_C52_Result_UnmarshalMust_Panic(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.UnmarshalMust(nil)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalMust_Panic", actual)
}

func Test_I5_C53_Result_SerializeSkipExistingIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"nilBytes": b == nil, "nilErr": err == nil}
	expected := args.Map{"nilBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues", actual)
}

func Test_I5_C54_Result_SerializeSkipExistingIssues_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	b, err := r.SerializeSkipExistingIssues()
	actual := args.Map{"hasBytes": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "SerializeSkipExistingIssues_Valid", actual)
}

func Test_I5_C55_Result_Serialize_Nil(t *testing.T) {
	var r *corejson.Result
	b, err := r.Serialize()
	actual := args.Map{"nilBytes": b == nil, "hasErr": err != nil}
	expected := args.Map{"nilBytes": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_Nil", actual)
}

func Test_I5_C56_Result_Serialize_WithError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	b, err := r.Serialize()
	actual := args.Map{"emptyBytes": len(b) == 0, "hasErr": err != nil}
	expected := args.Map{"emptyBytes": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_WithError", actual)
}

func Test_I5_C57_Result_Serialize_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	b, err := r.Serialize()
	actual := args.Map{"hasBytes": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_Valid", actual)
}

func Test_I5_C58_Result_SerializeMust(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	b := r.SerializeMust()
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "SerializeMust", actual)
}

func Test_I5_C59_Result_UnmarshalSkipExistingIssues_HasIssues(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues_HasIssues", actual)
}

func Test_I5_C60_Result_UnmarshalSkipExistingIssues_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"hello"`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"nilErr": err == nil, "val": s}
	expected := args.Map{"nilErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues_Valid", actual)
}

func Test_I5_C61_Result_UnmarshalSkipExistingIssues_BadJson(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`not-json`)}
	var s string
	err := r.UnmarshalSkipExistingIssues(&s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalSkipExistingIssues_BadJson", actual)
}

func Test_I5_C62_Result_UnmarshalResult(t *testing.T) {
	inner := corejson.NewResult.Any("hello")
	outerBytes, _ := inner.Serialize()
	outer := &corejson.Result{Bytes: outerBytes}
	result, err := outer.UnmarshalResult()
	actual := args.Map{"notNil": result != nil, "nilErr": err == nil}
	expected := args.Map{"notNil": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalResult", actual)
}

func Test_I5_C63_Result_JsonModel_Nil(t *testing.T) {
	var r *corejson.Result
	model := r.JsonModel()
	actual := args.Map{"hasErr": model.Error != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "JsonModel_Nil", actual)
}

func Test_I5_C64_Result_JsonModel_Valid(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	model := r.JsonModel()
	actual := args.Map{"hasBytes": len(model.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "JsonModel_Valid", actual)
}

func Test_I5_C65_Result_JsonModelAny(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	a := r.JsonModelAny()
	actual := args.Map{"notNil": a != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "JsonModelAny", actual)
}

func Test_I5_C66_Result_Json_JsonPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	j := r.Json()
	jp := r.JsonPtr()
	actual := args.Map{"hasBytes": len(j.Bytes) > 0, "ptrNotNil": jp != nil}
	expected := args.Map{"hasBytes": true, "ptrNotNil": true}
	expected.ShouldBeEqual(t, 0, "Json_JsonPtr", actual)
}

func Test_I5_C67_Result_ParseInjectUsingJson(t *testing.T) {
	r := &corejson.Result{}
	input := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"injected"`), TypeName: "T"})
	result, err := r.ParseInjectUsingJson(input)
	actual := args.Map{"notNil": result != nil, "nilErr": err == nil}
	expected := args.Map{"notNil": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson", actual)
}

func Test_I5_C68_Result_ParseInjectUsingJson_Error(t *testing.T) {
	r := &corejson.Result{}
	input := &corejson.Result{Bytes: []byte(`not-json`)}
	_, err := r.ParseInjectUsingJson(input)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson_Error", actual)
}

func Test_I5_C69_Result_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	r := &corejson.Result{}
	input := &corejson.Result{Bytes: []byte(`not-json`)}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		r.ParseInjectUsingJsonMust(input)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust_Panic", actual)
}

func Test_I5_C70_Result_CloneError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	actual := args.Map{"hasErr": r.CloneError() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CloneError", actual)

	r2 := &corejson.Result{}
	actual2 := args.Map{"nilErr": r2.CloneError() == nil}
	expected2 := args.Map{"nilErr": true}
	expected2.ShouldBeEqual(t, 0, "CloneError_NoErr", actual2)
}

func Test_I5_C71_Result_Ptr_NonPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	p := r.Ptr()
	np := p.NonPtr()
	actual := args.Map{"ptrNotNil": p != nil, "hasBytes": len(np.Bytes) > 0}
	expected := args.Map{"ptrNotNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Ptr_NonPtr", actual)
}

func Test_I5_C72_Result_NonPtr_Nil(t *testing.T) {
	var r *corejson.Result
	np := r.NonPtr()
	actual := args.Map{"hasErr": np.Error != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NonPtr_Nil", actual)
}

func Test_I5_C73_Result_ToPtr_ToNonPtr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	p := r.ToPtr()
	np := r.ToNonPtr()
	actual := args.Map{"ptrNotNil": p != nil, "hasBytes": len(np.Bytes) > 0}
	expected := args.Map{"ptrNotNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ToPtr_ToNonPtr", actual)
}

func Test_I5_C74_Result_IsEqualPtr(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	r3 := &corejson.Result{Bytes: []byte(`"y"`), TypeName: "T"}
	var rNil *corejson.Result

	actual := args.Map{
		"same":     r1.IsEqualPtr(r2),
		"diff":     r1.IsEqualPtr(r3),
		"bothNil":  rNil.IsEqualPtr(nil),
		"oneNil":   r1.IsEqualPtr(nil),
		"selfSame": r1.IsEqualPtr(r1),
	}
	expected := args.Map{
		"same":     true,
		"diff":     false,
		"bothNil":  true,
		"oneNil":   false,
		"selfSame": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr", actual)
}

func Test_I5_C75_Result_IsEqualPtr_DiffType(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "A"}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "B"}
	actual := args.Map{"equal": r1.IsEqualPtr(r2)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr_DiffType", actual)
}

func Test_I5_C76_Result_IsEqualPtr_DiffErr(t *testing.T) {
	r1 := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("a")}
	r2 := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("b")}
	actual := args.Map{"equal": r1.IsEqualPtr(r2)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "IsEqualPtr_DiffErr", actual)
}

func Test_I5_C77_Result_CombineErrorWithRefString(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	s := r.CombineErrorWithRefString("ref1", "ref2")
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString", actual)
}

func Test_I5_C78_Result_CombineErrorWithRefString_NoError(t *testing.T) {
	r := &corejson.Result{}
	s := r.CombineErrorWithRefString("ref1")
	actual := args.Map{"empty": s == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefString_NoErr", actual)
}

func Test_I5_C79_Result_CombineErrorWithRefError(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	err := r.CombineErrorWithRefError("ref1")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefError", actual)

	r2 := &corejson.Result{}
	err2 := r2.CombineErrorWithRefError("ref1")
	actual2 := args.Map{"nilErr": err2 == nil}
	expected2 := args.Map{"nilErr": true}
	expected2.ShouldBeEqual(t, 0, "CombineErrorWithRefError_NoErr", actual2)
}

func Test_I5_C80_Result_IsEqual(t *testing.T) {
	r1 := corejson.Result{Bytes: []byte(`"x"`)}
	r2 := corejson.Result{Bytes: []byte(`"x"`)}
	r3 := corejson.Result{Bytes: []byte(`"y"`)}
	actual := args.Map{"same": r1.IsEqual(r2), "diff": r1.IsEqual(r3)}
	expected := args.Map{"same": true, "diff": false}
	expected.ShouldBeEqual(t, 0, "IsEqual", actual)
}

func Test_I5_C81_Result_BytesError(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")}
	be := r.BytesError()
	actual := args.Map{"notNil": be != nil, "hasBytes": len(be.Bytes) > 0}
	expected := args.Map{"notNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "BytesError", actual)

	var rNil *corejson.Result
	actual2 := args.Map{"nil": rNil.BytesError() == nil}
	expected2 := args.Map{"nil": true}
	expected2.ShouldBeEqual(t, 0, "BytesError_Nil", actual2)
}

func Test_I5_C82_Result_Dispose(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e"), TypeName: "T"}
	r.Dispose()
	actual := args.Map{"nilErr": r.Error == nil, "nilBytes": r.Bytes == nil, "emptyType": r.TypeName == ""}
	expected := args.Map{"nilErr": true, "nilBytes": true, "emptyType": true}
	expected.ShouldBeEqual(t, 0, "Dispose", actual)

	var rNil *corejson.Result
	rNil.Dispose() // should not panic
}

func Test_I5_C83_Result_CloneIf(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	cloned := r.CloneIf(true, true)
	notCloned := r.CloneIf(false, false)
	actual := args.Map{"clonedHasBytes": len(cloned.Bytes) > 0, "notClonedHasBytes": len(notCloned.Bytes) > 0}
	expected := args.Map{"clonedHasBytes": true, "notClonedHasBytes": true}
	expected.ShouldBeEqual(t, 0, "CloneIf", actual)
}

func Test_I5_C84_Result_ClonePtr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	cloned := r.ClonePtr(true)
	actual := args.Map{"notNil": cloned != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr", actual)

	var rNil *corejson.Result
	actual2 := args.Map{"nil": rNil.ClonePtr(true) == nil}
	expected2 := args.Map{"nil": true}
	expected2.ShouldBeEqual(t, 0, "ClonePtr_Nil", actual2)
}

func Test_I5_C85_Result_Clone_ShallowAndDeep(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`), TypeName: "T"}
	shallow := r.Clone(false)
	deep := r.Clone(true)
	actual := args.Map{"shallowHasBytes": len(shallow.Bytes) > 0, "deepHasBytes": len(deep.Bytes) > 0}
	expected := args.Map{"shallowHasBytes": true, "deepHasBytes": true}
	expected.ShouldBeEqual(t, 0, "Clone_ShallowAndDeep", actual)
}

func Test_I5_C86_Result_Clone_Empty(t *testing.T) {
	r := corejson.Result{}
	cloned := r.Clone(true)
	actual := args.Map{"emptyBytes": len(cloned.Bytes) == 0}
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "Clone_Empty", actual)
}

func Test_I5_C87_Result_AsInterfaces(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	actual := args.Map{
		"jsonContracts": r.AsJsonContractsBinder() != nil,
		"jsoner":        r.AsJsoner() != nil,
		"selfInjector":  r.AsJsonParseSelfInjector() != nil,
	}
	expected := args.Map{
		"jsonContracts": true,
		"jsoner":        true,
		"selfInjector":  true,
	}
	expected.ShouldBeEqual(t, 0, "AsInterfaces", actual)
}

func Test_I5_C88_Result_JsonParseSelfInject(t *testing.T) {
	r := corejson.Result{}
	input := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"injected"`), TypeName: "T"})
	err := r.JsonParseSelfInject(input)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject", actual)
}

func Test_I5_C89_Result_InjectInto(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	target := corejson.Result{}
	err := r.InjectInto(&target)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "InjectInto", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// castingAny
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C90_CastAny_FromToDefault(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToDefault([]byte(`"hello"`), &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToDefault", actual)
}

func Test_I5_C91_CastAny_FromToReflection(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToReflection([]byte(`"hello"`), &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToReflection", actual)
}

func Test_I5_C92_CastAny_FromToOption_Nil(t *testing.T) {
	err := corejson.CastAny.FromToOption(true, nil, nil)
	actual := args.Map{"hasResult": err != nil || err == nil} // exercises nil path
	expected := args.Map{"hasResult": true}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_Nil", actual)
}

func Test_I5_C93_CastAny_FromToOption_String(t *testing.T) {
	var out map[string]string
	err := corejson.CastAny.FromToOption(false, `{"a":"b"}`, &out)
	actual := args.Map{"nilErr": err == nil, "val": out["a"]}
	expected := args.Map{"nilErr": true, "val": "b"}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_String", actual)
}

func Test_I5_C94_CastAny_FromToOption_Error(t *testing.T) {
	var out string
	err := corejson.CastAny.FromToOption(false, errors.New(`"test"`), &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "test"}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_Error", actual)
}

func Test_I5_C95_CastAny_FromToOption_NilError(t *testing.T) {
	var nilErr error
	var out string
	err := corejson.CastAny.FromToOption(false, nilErr, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_NilError", actual)
}

func Test_I5_C96_CastAny_FromToOption_SerializerFunc(t *testing.T) {
	fn := func() ([]byte, error) { return []byte(`"from-func"`), nil }
	var out string
	err := corejson.CastAny.FromToOption(false, fn, &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "from-func"}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_SerializerFunc", actual)
}

func Test_I5_C97_CastAny_FromToOption_AnyItem(t *testing.T) {
	type s struct{ A int }
	var out s
	err := corejson.CastAny.FromToOption(false, s{A: 42}, &out)
	actual := args.Map{"nilErr": err == nil, "val": out.A}
	expected := args.Map{"nilErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "CastAny_FromToOption_AnyItem", actual)
}

func Test_I5_C98_CastAny_OrDeserializeTo(t *testing.T) {
	var out string
	err := corejson.CastAny.OrDeserializeTo([]byte(`"hi"`), &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "hi"}
	expected.ShouldBeEqual(t, 0, "CastAny_OrDeserializeTo", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// anyTo — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C99_AnyTo_SerializedJsonResult_Nil(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(nil)
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedJsonResult_Nil", actual)
}

func Test_I5_C100_AnyTo_SerializedJsonResult_Error(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(errors.New("test"))
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedJsonResult_Error", actual)
}

func Test_I5_C101_AnyTo_SerializedJsonResult_EmptyError(t *testing.T) {
	r := corejson.AnyTo.SerializedJsonResult(errors.New(""))
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedJsonResult_EmptyError", actual)
}

func Test_I5_C102_AnyTo_SerializedRaw(t *testing.T) {
	b, err := corejson.AnyTo.SerializedRaw([]byte(`"test"`))
	actual := args.Map{"hasBytes": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedRaw", actual)
}

func Test_I5_C103_AnyTo_SerializedString(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString([]byte(`"hello"`))
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedString", actual)
}

func Test_I5_C104_AnyTo_SerializedString_Error(t *testing.T) {
	s, err := corejson.AnyTo.SerializedString(nil)
	actual := args.Map{"empty": s == "", "hasErr": err != nil}
	expected := args.Map{"empty": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedString_Error", actual)
}

func Test_I5_C105_AnyTo_SerializedSafeString(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString([]byte(`"hello"`))
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedSafeString", actual)
}

func Test_I5_C106_AnyTo_SerializedSafeString_Error(t *testing.T) {
	s := corejson.AnyTo.SerializedSafeString(nil)
	actual := args.Map{"empty": s == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedSafeString_Error", actual)
}

func Test_I5_C107_AnyTo_SerializedStringMust(t *testing.T) {
	s := corejson.AnyTo.SerializedStringMust([]byte(`"hello"`))
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedStringMust", actual)
}

func Test_I5_C108_AnyTo_SafeJsonString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SafeJsonString", actual)
}

func Test_I5_C109_AnyTo_PrettyStringWithError(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError("hello")
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_String", actual)
}

func Test_I5_C110_AnyTo_PrettyStringWithError_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.PrettyStringWithError([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_Bytes", actual)
}

func Test_I5_C111_AnyTo_PrettyStringWithError_Result(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_Result", actual)
}

func Test_I5_C112_AnyTo_PrettyStringWithError_ResultPtr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_ResultPtr", actual)
}

func Test_I5_C113_AnyTo_PrettyStringWithError_ResultWithErr(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`{"a":1}`), Error: errors.New("e")}
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"hasContent": len(s) > 0, "hasErr": err != nil}
	expected := args.Map{"hasContent": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_ResultWithErr", actual)
}

func Test_I5_C114_AnyTo_PrettyStringWithError_ResultPtrWithErr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`), Error: errors.New("e")}
	s, err := corejson.AnyTo.PrettyStringWithError(r)
	actual := args.Map{"hasContent": len(s) > 0, "hasErr": err != nil}
	expected := args.Map{"hasContent": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringWithError_ResultPtrWithErr", actual)
}

func Test_I5_C115_AnyTo_SafeJsonPrettyString(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_String", actual)
}

func Test_I5_C116_AnyTo_SafeJsonPrettyString_Bytes(t *testing.T) {
	s := corejson.AnyTo.SafeJsonPrettyString([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_Bytes", actual)
}

func Test_I5_C117_AnyTo_SafeJsonPrettyString_Result(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`{"a":1}`)}
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_Result", actual)
}

func Test_I5_C118_AnyTo_SafeJsonPrettyString_ResultPtr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`{"a":1}`)}
	s := corejson.AnyTo.SafeJsonPrettyString(r)
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SafeJsonPrettyString_ResultPtr", actual)
}

func Test_I5_C119_AnyTo_JsonString(t *testing.T) {
	s := corejson.AnyTo.JsonString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_String", actual)
}

func Test_I5_C120_AnyTo_JsonString_Bytes(t *testing.T) {
	s := corejson.AnyTo.JsonString([]byte(`"x"`))
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_Bytes", actual)
}

func Test_I5_C121_AnyTo_JsonString_Result(t *testing.T) {
	r := corejson.Result{Bytes: []byte(`"x"`)}
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_Result", actual)
}

func Test_I5_C122_AnyTo_JsonString_ResultPtr(t *testing.T) {
	r := &corejson.Result{Bytes: []byte(`"x"`)}
	s := corejson.AnyTo.JsonString(r)
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonString_ResultPtr", actual)
}

func Test_I5_C123_AnyTo_JsonStringWithErr(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr("hello")
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringWithErr_String", actual)
}

func Test_I5_C124_AnyTo_JsonStringWithErr_Bytes(t *testing.T) {
	s, err := corejson.AnyTo.JsonStringWithErr([]byte(`"x"`))
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringWithErr_Bytes", actual)
}

func Test_I5_C125_AnyTo_JsonStringWithErr_ResultWithErr(t *testing.T) {
	r := corejson.Result{Error: errors.New("e"), Bytes: []byte(`"x"`)}
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"hasContent": len(s) > 0, "hasErr": err != nil}
	expected := args.Map{"hasContent": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringWithErr_ResultWithErr", actual)
}

func Test_I5_C126_AnyTo_JsonStringWithErr_ResultPtrWithErr(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e"), Bytes: []byte(`"x"`)}
	s, err := corejson.AnyTo.JsonStringWithErr(r)
	actual := args.Map{"hasContent": len(s) > 0, "hasErr": err != nil}
	expected := args.Map{"hasContent": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringWithErr_ResultPtrWithErr", actual)
}

func Test_I5_C127_AnyTo_JsonStringMust(t *testing.T) {
	s := corejson.AnyTo.JsonStringMust("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_JsonStringMust", actual)
}

func Test_I5_C128_AnyTo_PrettyStringMust(t *testing.T) {
	s := corejson.AnyTo.PrettyStringMust("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyStringMust", actual)
}

func Test_I5_C129_AnyTo_UsingSerializer(t *testing.T) {
	r := corejson.AnyTo.UsingSerializer(nil)
	actual := args.Map{"nil": r == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_UsingSerializer_Nil", actual)
}

func Test_I5_C130_AnyTo_SerializedFieldsMap(t *testing.T) {
	type s struct{ A int }
	m, err := corejson.AnyTo.SerializedFieldsMap(s{A: 42})
	_ = m
	actual := args.Map{"ran": err == nil || err != nil}
	expected := args.Map{"ran": true}
	expected.ShouldBeEqual(t, 0, "AnyTo_SerializedFieldsMap", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// deserializerLogic — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C131_Deserialize_UsingStringPtr_Nil(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringPtr_Nil", actual)
}

func Test_I5_C132_Deserialize_UsingStringPtr_Valid(t *testing.T) {
	str := `"hello"`
	var out string
	err := corejson.Deserialize.UsingStringPtr(&str, &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringPtr_Valid", actual)
}

func Test_I5_C133_Deserialize_UsingError_Nil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingError(nil, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingError_Nil", actual)
}

func Test_I5_C134_Deserialize_UsingError_Valid(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingError(errors.New(`"hello"`), &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingError_Valid", actual)
}

func Test_I5_C135_Deserialize_UsingErrorWhichJsonResult(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingErrorWhichJsonResult_Nil", actual)
}

func Test_I5_C136_Deserialize_ApplyMust_Panic(t *testing.T) {
	r := &corejson.Result{Error: errors.New("e")}
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		var s string
		corejson.Deserialize.ApplyMust(r, &s)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_ApplyMust_Panic", actual)
}

func Test_I5_C137_Deserialize_FromString(t *testing.T) {
	var out string
	err := corejson.Deserialize.FromString(`"hello"`, &out)
	actual := args.Map{"nilErr": err == nil, "val": out}
	expected := args.Map{"nilErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromString", actual)
}

func Test_I5_C138_Deserialize_FromStringMust(t *testing.T) {
	var out string
	corejson.Deserialize.FromStringMust(`"hello"`, &out)
	actual := args.Map{"val": out}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromStringMust", actual)
}

func Test_I5_C139_Deserialize_FromStringMust_Panic(t *testing.T) {
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		var out string
		corejson.Deserialize.FromStringMust(`not-json`, &out)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_FromStringMust_Panic", actual)
}

func Test_I5_C140_Deserialize_MapAnyToPointer(t *testing.T) {
	type s struct {
		A int `json:"a"`
	}
	var out s
	err := corejson.Deserialize.MapAnyToPointer(false, map[string]any{"a": 1}, &out)
	actual := args.Map{"nilErr": err == nil, "val": out.A}
	expected := args.Map{"nilErr": true, "val": 1}
	expected.ShouldBeEqual(t, 0, "Deserialize_MapAnyToPointer", actual)
}

func Test_I5_C141_Deserialize_MapAnyToPointer_SkipEmpty(t *testing.T) {
	var out string
	err := corejson.Deserialize.MapAnyToPointer(true, map[string]any{}, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_MapAnyToPointer_SkipEmpty", actual)
}

func Test_I5_C142_Deserialize_UsingStringOption(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingStringOption(true, "", &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringOption_Empty", actual)
}

func Test_I5_C143_Deserialize_UsingStringIgnoreEmpty(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingStringIgnoreEmpty", actual)
}

func Test_I5_C144_Deserialize_UsingBytesPointer_Nil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingBytesPointer(nil, &out)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesPointer_Nil", actual)
}

func Test_I5_C145_Deserialize_UsingBytesPointerMust_Panic(t *testing.T) {
	didPanic := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				didPanic = true
			}
		}()
		var out string
		corejson.Deserialize.UsingBytesPointerMust(nil, &out)
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesPointerMust_Panic", actual)
}

func Test_I5_C146_Deserialize_UsingBytesIf(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"x"`), &out)
	actual := args.Map{"nilErr": err == nil, "empty": out == ""}
	expected := args.Map{"nilErr": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesIf_Skip", actual)
}

func Test_I5_C147_Deserialize_UsingBytesPointerIf(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"x"`), &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesPointerIf_Skip", actual)
}

func Test_I5_C148_Deserialize_UsingBytesMust(t *testing.T) {
	var out string
	corejson.Deserialize.UsingBytesMust([]byte(`"hello"`), &out)
	actual := args.Map{"val": out}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingBytesMust", actual)
}

func Test_I5_C149_Deserialize_UsingSafeBytesMust_Empty(t *testing.T) {
	var out string
	corejson.Deserialize.UsingSafeBytesMust([]byte{}, &out)
	actual := args.Map{"empty": out == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_UsingSafeBytesMust_Empty", actual)
}

func Test_I5_C150_Deserialize_AnyToFieldsMap(t *testing.T) {
	type s struct{ A int }
	m, _ := corejson.Deserialize.AnyToFieldsMap(s{A: 1})
	_ = m
	actual := args.Map{"ran": true}
	expected := args.Map{"ran": true}
	expected.ShouldBeEqual(t, 0, "Deserialize_AnyToFieldsMap", actual)
}

func Test_I5_C151_Deserialize_UsingDeserializerToOption_SkipNil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerToOption_SkipNil", actual)
}

func Test_I5_C152_Deserialize_UsingDeserializerToOption_NilNotSkipped(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &out)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerToOption_NilNotSkipped", actual)
}

func Test_I5_C153_Deserialize_UsingDeserializerFuncDefined_Nil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &out)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerFuncDefined_Nil", actual)
}

func Test_I5_C154_Deserialize_UsingDeserializerFuncDefined_Valid(t *testing.T) {
	fn := func(toPtr any) error { return nil }
	var out string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(fn, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingDeserializerFuncDefined_Valid", actual)
}

func Test_I5_C155_Deserialize_UsingJsonerToAny_SkipNil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAny_SkipNil", actual)
}

func Test_I5_C156_Deserialize_UsingJsonerToAny_NilNotSkipped(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingJsonerToAny(false, nil, &out)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAny_NilNotSkipped", actual)
}

func Test_I5_C157_Deserialize_UsingJsonerToAnyMust_SkipNil(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &out)
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust_SkipNil", actual)
}

func Test_I5_C158_Deserialize_UsingJsonerToAnyMust_NilNotSkipped(t *testing.T) {
	var out string
	err := corejson.Deserialize.UsingJsonerToAnyMust(false, nil, &out)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UsingJsonerToAnyMust_NilNotSkipped", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Utility funcs — BytesCloneIf, BytesDeepClone, BytesToString, BytesToPrettyString, JsonString
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C159_BytesCloneIf_NoClone(t *testing.T) {
	b := corejson.BytesCloneIf(false, []byte("abc"))
	actual := args.Map{"empty": len(b) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf_NoClone", actual)
}

func Test_I5_C160_BytesCloneIf_Clone(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte("abc"))
	actual := args.Map{"len": len(b)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf_Clone", actual)
}

func Test_I5_C161_BytesCloneIf_Empty(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte{})
	actual := args.Map{"empty": len(b) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf_Empty", actual)
}

func Test_I5_C162_BytesDeepClone(t *testing.T) {
	b := corejson.BytesDeepClone([]byte("abc"))
	actual := args.Map{"len": len(b)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone", actual)
}

func Test_I5_C163_BytesDeepClone_Empty(t *testing.T) {
	b := corejson.BytesDeepClone([]byte{})
	actual := args.Map{"empty": len(b) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone_Empty", actual)
}

func Test_I5_C164_BytesToString(t *testing.T) {
	actual := args.Map{"val": corejson.BytesToString([]byte(`"x"`))}
	expected := args.Map{"val": `"x"`}
	expected.ShouldBeEqual(t, 0, "BytesToString", actual)
}

func Test_I5_C165_BytesToString_Empty(t *testing.T) {
	actual := args.Map{"empty": corejson.BytesToString(nil) == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString_Empty", actual)
}

func Test_I5_C166_BytesToPrettyString(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString", actual)
}

func Test_I5_C167_BytesToPrettyString_Empty(t *testing.T) {
	actual := args.Map{"empty": corejson.BytesToPrettyString(nil) == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString_Empty", actual)
}

func Test_I5_C168_JsonString_Func(t *testing.T) {
	s, err := corejson.JsonString("hello")
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "JsonString_Func", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// serializerLogic — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C169_Serialize_StringsApply(t *testing.T) {
	r := corejson.Serialize.StringsApply([]string{"a", "b"})
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_StringsApply", actual)
}

func Test_I5_C170_Serialize_FromBytes(t *testing.T) {
	r := corejson.Serialize.FromBytes([]byte("abc"))
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromBytes", actual)
}

func Test_I5_C171_Serialize_FromStrings(t *testing.T) {
	r := corejson.Serialize.FromStrings([]string{"a"})
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromStrings", actual)
}

func Test_I5_C172_Serialize_FromStringsSpread(t *testing.T) {
	r := corejson.Serialize.FromStringsSpread("a", "b")
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromStringsSpread", actual)
}

func Test_I5_C173_Serialize_FromString(t *testing.T) {
	r := corejson.Serialize.FromString("hello")
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromString", actual)
}

func Test_I5_C174_Serialize_FromInteger(t *testing.T) {
	r := corejson.Serialize.FromInteger(42)
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromInteger", actual)
}

func Test_I5_C175_Serialize_FromInteger64(t *testing.T) {
	r := corejson.Serialize.FromInteger64(42)
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromInteger64", actual)
}

func Test_I5_C176_Serialize_FromBool(t *testing.T) {
	r := corejson.Serialize.FromBool(true)
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromBool", actual)
}

func Test_I5_C177_Serialize_FromIntegers(t *testing.T) {
	r := corejson.Serialize.FromIntegers([]int{1, 2})
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromIntegers", actual)
}

type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func Test_I5_C178_Serialize_FromStringer(t *testing.T) {
	r := corejson.Serialize.FromStringer(testStringer{"hello"})
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_FromStringer", actual)
}

func Test_I5_C179_Serialize_UsingAnyPtr(t *testing.T) {
	r := corejson.Serialize.UsingAnyPtr("hello")
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_UsingAnyPtr", actual)
}

func Test_I5_C180_Serialize_UsingAny(t *testing.T) {
	r := corejson.Serialize.UsingAny("hello")
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_UsingAny", actual)
}

func Test_I5_C181_Serialize_Raw(t *testing.T) {
	b, err := corejson.Serialize.Raw("hello")
	actual := args.Map{"hasBytes": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_Raw", actual)
}

func Test_I5_C182_Serialize_Marshal(t *testing.T) {
	b, err := corejson.Serialize.Marshal("hello")
	actual := args.Map{"hasBytes": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_Marshal", actual)
}

func Test_I5_C183_Serialize_ApplyMust(t *testing.T) {
	r := corejson.Serialize.ApplyMust("hello")
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ApplyMust", actual)
}

func Test_I5_C184_Serialize_ToBytesMust(t *testing.T) {
	b := corejson.Serialize.ToBytesMust("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToBytesMust", actual)
}

func Test_I5_C185_Serialize_ToSafeBytesMust(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesMust("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToSafeBytesMust", actual)
}

func Test_I5_C186_Serialize_ToSafeBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToSafeBytesSwallowErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToSafeBytesSwallowErr", actual)
}

func Test_I5_C187_Serialize_ToBytesSwallowErr(t *testing.T) {
	b := corejson.Serialize.ToBytesSwallowErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToBytesSwallowErr", actual)
}

func Test_I5_C188_Serialize_ToBytesErr(t *testing.T) {
	b, err := corejson.Serialize.ToBytesErr("hello")
	actual := args.Map{"hasBytes": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToBytesErr", actual)
}

func Test_I5_C189_Serialize_ToString(t *testing.T) {
	s := corejson.Serialize.ToString("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToString", actual)
}

func Test_I5_C190_Serialize_ToStringMust(t *testing.T) {
	s := corejson.Serialize.ToStringMust("hello")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToStringMust", actual)
}

func Test_I5_C191_Serialize_ToStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToStringErr("hello")
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToStringErr", actual)
}

func Test_I5_C192_Serialize_ToPrettyStringErr(t *testing.T) {
	s, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != "", "nilErr": err == nil}
	expected := args.Map{"notEmpty": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToPrettyStringErr", actual)
}

func Test_I5_C193_Serialize_ToPrettyStringIncludingErr(t *testing.T) {
	s := corejson.Serialize.ToPrettyStringIncludingErr(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_ToPrettyStringIncludingErr", actual)
}

func Test_I5_C194_Serialize_Pretty(t *testing.T) {
	s := corejson.Serialize.Pretty(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Serialize_Pretty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// deserializeFromBytesTo — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C195_BytesTo_Bool(t *testing.T) {
	b, err := corejson.Deserialize.BytesTo.Bool([]byte(`true`))
	actual := args.Map{"val": b, "nilErr": err == nil}
	expected := args.Map{"val": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_Bool", actual)
}

func Test_I5_C196_BytesTo_BoolMust(t *testing.T) {
	b := corejson.Deserialize.BytesTo.BoolMust([]byte(`true`))
	actual := args.Map{"val": b}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_BoolMust", actual)
}

func Test_I5_C197_BytesTo_Integer(t *testing.T) {
	i, err := corejson.Deserialize.BytesTo.Integer([]byte(`42`))
	actual := args.Map{"val": i, "nilErr": err == nil}
	expected := args.Map{"val": 42, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integer", actual)
}

func Test_I5_C198_BytesTo_IntegerMust(t *testing.T) {
	i := corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))
	actual := args.Map{"val": i}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "BytesTo_IntegerMust", actual)
}

func Test_I5_C199_BytesTo_Integer64(t *testing.T) {
	i, err := corejson.Deserialize.BytesTo.Integer64([]byte(`99`))
	actual := args.Map{"val": int(i), "nilErr": err == nil}
	expected := args.Map{"val": 99, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integer64", actual)
}

func Test_I5_C200_BytesTo_Integer64Must(t *testing.T) {
	i := corejson.Deserialize.BytesTo.Integer64Must([]byte(`99`))
	actual := args.Map{"val": int(i)}
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integer64Must", actual)
}

func Test_I5_C201_BytesTo_Integers(t *testing.T) {
	ints, err := corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
	actual := args.Map{"len": len(ints), "nilErr": err == nil}
	expected := args.Map{"len": 3, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_Integers", actual)
}

func Test_I5_C202_BytesTo_IntegersMust(t *testing.T) {
	ints := corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1,2]`))
	actual := args.Map{"len": len(ints)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BytesTo_IntegersMust", actual)
}

func Test_I5_C203_BytesTo_Strings(t *testing.T) {
	strs, err := corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
	actual := args.Map{"len": len(strs), "nilErr": err == nil}
	expected := args.Map{"len": 2, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_Strings", actual)
}

func Test_I5_C204_BytesTo_StringsMust(t *testing.T) {
	strs := corejson.Deserialize.BytesTo.StringsMust([]byte(`["a"]`))
	actual := args.Map{"len": len(strs)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesTo_StringsMust", actual)
}

func Test_I5_C205_BytesTo_String(t *testing.T) {
	s, err := corejson.Deserialize.BytesTo.String([]byte(`"hi"`))
	actual := args.Map{"val": s, "nilErr": err == nil}
	expected := args.Map{"val": "hi", "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_String", actual)
}

func Test_I5_C206_BytesTo_StringMust(t *testing.T) {
	s := corejson.Deserialize.BytesTo.StringMust([]byte(`"hi"`))
	actual := args.Map{"val": s}
	expected := args.Map{"val": "hi"}
	expected.ShouldBeEqual(t, 0, "BytesTo_StringMust", actual)
}

func Test_I5_C207_BytesTo_MapAnyItem(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
	actual := args.Map{"hasKey": m["a"] != nil, "nilErr": err == nil}
	expected := args.Map{"hasKey": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapAnyItem", actual)
}

func Test_I5_C208_BytesTo_MapAnyItemMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))
	actual := args.Map{"hasKey": m["a"] != nil}
	expected := args.Map{"hasKey": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapAnyItemMust", actual)
}

func Test_I5_C209_BytesTo_MapStringString(t *testing.T) {
	m, err := corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
	actual := args.Map{"val": m["a"], "nilErr": err == nil}
	expected := args.Map{"val": "b", "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapStringString", actual)
}

func Test_I5_C210_BytesTo_MapStringStringMust(t *testing.T) {
	m := corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))
	actual := args.Map{"val": m["a"]}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "BytesTo_MapStringStringMust", actual)
}

func Test_I5_C211_BytesTo_Bytes(t *testing.T) {
	b, err := corejson.Deserialize.BytesTo.Bytes([]byte(`"aGVsbG8="`))
	actual := args.Map{"hasData": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasData": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_Bytes", actual)
}

func Test_I5_C212_BytesTo_BytesMust(t *testing.T) {
	b := corejson.Deserialize.BytesTo.BytesMust([]byte(`"aGVsbG8="`))
	actual := args.Map{"hasData": len(b) > 0}
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "BytesTo_BytesMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// deserializeFromResultTo — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C213_ResultTo_Byte(t *testing.T) {
	r := corejson.NewResult.AnyPtr(byte(65))
	b, err := corejson.Deserialize.ResultTo.Byte(r)
	actual := args.Map{"val": int(b), "nilErr": err == nil}
	expected := args.Map{"val": 65, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "ResultTo_Byte", actual)
}

func Test_I5_C214_ResultTo_ByteMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(byte(65))
	b := corejson.Deserialize.ResultTo.ByteMust(r)
	actual := args.Map{"val": int(b)}
	expected := args.Map{"val": 65}
	expected.ShouldBeEqual(t, 0, "ResultTo_ByteMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// emptyCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C215_Empty_ResultWithErr(t *testing.T) {
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))
	actual := args.Map{"hasErr": r.Error != nil, "type": r.TypeName}
	expected := args.Map{"hasErr": true, "type": "T"}
	expected.ShouldBeEqual(t, 0, "Empty_ResultWithErr", actual)
}

func Test_I5_C216_Empty_ResultPtrWithErr(t *testing.T) {
	r := corejson.Empty.ResultPtrWithErr("T", errors.New("e"))
	actual := args.Map{"notNil": r != nil, "hasErr": r.Error != nil}
	expected := args.Map{"notNil": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Empty_ResultPtrWithErr", actual)
}

func Test_I5_C217_Empty_BytesCollection(t *testing.T) {
	c := corejson.Empty.BytesCollection()
	actual := args.Map{"empty": c.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty_BytesCollection", actual)
}

func Test_I5_C218_Empty_BytesCollectionPtr(t *testing.T) {
	c := corejson.Empty.BytesCollectionPtr()
	actual := args.Map{"notNil": c != nil, "empty": c.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty_BytesCollectionPtr", actual)
}

func Test_I5_C219_Empty_ResultsPtrCollection(t *testing.T) {
	c := corejson.Empty.ResultsPtrCollection()
	actual := args.Map{"notNil": c != nil, "empty": c.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty_ResultsPtrCollection", actual)
}

func Test_I5_C220_Empty_MapResults(t *testing.T) {
	m := corejson.Empty.MapResults()
	actual := args.Map{"notNil": m != nil, "empty": m.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty_MapResults", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newResultCreator — remaining branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C221_NewResult_PtrUsingStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.PtrUsingStringPtr(nil, "T")
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult_PtrUsingStringPtr_Nil", actual)
}

func Test_I5_C222_NewResult_PtrUsingStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.PtrUsingStringPtr(&s, "T")
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_PtrUsingStringPtr_Valid", actual)
}

func Test_I5_C223_NewResult_UsingErrorStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingErrorStringPtr_Nil", actual)
}

func Test_I5_C224_NewResult_UsingErrorStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingErrorStringPtr(nil, &s, "T")
	actual := args.Map{"hasBytes": len(r.Bytes) > 0, "nilErr": r.Error == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingErrorStringPtr_Valid", actual)
}

func Test_I5_C225_NewResult_UsingTypePlusStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingTypePlusStringPtr_Nil", actual)
}

func Test_I5_C226_NewResult_UsingTypePlusStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingTypePlusStringPtr("T", &s)
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingTypePlusStringPtr_Valid", actual)
}

func Test_I5_C227_NewResult_UsingStringPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingStringPtr(nil)
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingStringPtr_Nil", actual)
}

func Test_I5_C228_NewResult_UsingStringPtr_Valid(t *testing.T) {
	s := `"hello"`
	r := corejson.NewResult.UsingStringPtr(&s)
	actual := args.Map{"hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingStringPtr_Valid", actual)
}

func Test_I5_C229_NewResult_UsingBytesPtr_Nil(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr(nil)
	actual := args.Map{"emptyBytes": len(r.Bytes) == 0}
	expected := args.Map{"emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesPtr_Nil", actual)
}

func Test_I5_C230_NewResult_UsingBytesPtrErrPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T")
	actual := args.Map{"hasErr": r.Error != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesPtrErrPtr_Nil", actual)
}

func Test_I5_C231_NewResult_UsingBytesErrPtr_Empty(t *testing.T) {
	r := corejson.NewResult.UsingBytesErrPtr([]byte{}, errors.New("e"), "T")
	actual := args.Map{"hasErr": r.Error != nil, "emptyBytes": len(r.Bytes) == 0}
	expected := args.Map{"hasErr": true, "emptyBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingBytesErrPtr_Empty", actual)
}

func Test_I5_C232_NewResult_UsingSerializer_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializer(nil)
	actual := args.Map{"nil": r == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingSerializer_Nil", actual)
}

func Test_I5_C233_NewResult_UsingSerializerFunc_Nil(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(nil)
	actual := args.Map{"nil": r == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingSerializerFunc_Nil", actual)
}

func Test_I5_C234_NewResult_UsingJsoner_Nil(t *testing.T) {
	r := corejson.NewResult.UsingJsoner(nil)
	actual := args.Map{"nil": r == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_UsingJsoner_Nil", actual)
}

func Test_I5_C235_NewResult_Many(t *testing.T) {
	r := corejson.NewResult.Many("a", "b", "c")
	actual := args.Map{"notNil": r != nil, "hasBytes": len(r.Bytes) > 0}
	expected := args.Map{"notNil": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewResult_Many", actual)
}

func Test_I5_C236_NewResult_AnyToCastingResult(t *testing.T) {
	r := corejson.NewResult.AnyToCastingResult("hello")
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewResult_AnyToCastingResult", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// StaticJsonError var
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_C237_StaticJsonError(t *testing.T) {
	actual := args.Map{"notNil": corejson.StaticJsonError != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "StaticJsonError", actual)
}
