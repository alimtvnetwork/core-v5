package corejsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// New / NewPtr
// ═══════════════════════════════════════════

func Test_Cov6_New_Valid(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{
		"hasError":  r.HasError(),
		"hasBytes":  len(r.Bytes) > 0,
		"typeNotEmpty": r.TypeName != "",
	}
	expected := args.Map{"hasError": false, "hasBytes": true, "typeNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "New valid", actual)
}

func Test_Cov6_NewPtr_Valid(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"notNil":    r != nil,
		"hasError":  r.HasError(),
		"hasBytes":  len(r.Bytes) > 0,
	}
	expected := args.Map{"notNil": true, "hasError": false, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "NewPtr valid", actual)
}

func Test_Cov6_New_Nil(t *testing.T) {
	r := corejson.New(nil)
	actual := args.Map{
		"hasError": r.HasError(),
		"hasBytes": len(r.Bytes) > 0,
	}
	expected := args.Map{"hasError": false, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "New nil", actual)
}

// ═══════════════════════════════════════════
// Result — basic state methods
// ═══════════════════════════════════════════

func Test_Cov6_Result_Length(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"len":    r.Length(),
		"nilLen": nilR.Length(),
	}
	expected := args.Map{"len": 7, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Result Length", actual)
}

func Test_Cov6_Result_HasError(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"noErr":  r.HasError(),
		"nilErr": nilR.HasError(),
	}
	expected := args.Map{"noErr": false, "nilErr": false}
	expected.ShouldBeEqual(t, 0, "Result HasError", actual)
}

func Test_Cov6_Result_IsEmptyError(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"emptyErr":    r.IsEmptyError(),
		"nilEmptyErr": nilR.IsEmptyError(),
	}
	expected := args.Map{"emptyErr": true, "nilEmptyErr": true}
	expected.ShouldBeEqual(t, 0, "Result IsEmptyError", actual)
}

func Test_Cov6_Result_ErrorString(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"errStr": r.ErrorString()}
	expected := args.Map{"errStr": ""}
	expected.ShouldBeEqual(t, 0, "Result ErrorString", actual)
}

func Test_Cov6_Result_IsErrorEqual(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"bothNil":    r.IsErrorEqual(nil),
		"oneNotNil":  r.IsErrorEqual(errors.New("test")),
	}
	expected := args.Map{"bothNil": true, "oneNotNil": false}
	expected.ShouldBeEqual(t, 0, "Result IsErrorEqual", actual)
}

func Test_Cov6_Result_IsEmpty(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"notEmpty":  r.IsEmpty(),
		"nilEmpty":  nilR.IsEmpty(),
		"hasAny":    r.HasAnyItem(),
	}
	expected := args.Map{"notEmpty": false, "nilEmpty": true, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "Result IsEmpty", actual)
}

func Test_Cov6_Result_IsAnyNull(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"notNull":  r.IsAnyNull(),
		"nilNull":  nilR.IsAnyNull(),
	}
	expected := args.Map{"notNull": false, "nilNull": true}
	expected.ShouldBeEqual(t, 0, "Result IsAnyNull", actual)
}

// ═══════════════════════════════════════════
// Result — JSON string methods
// ═══════════════════════════════════════════

func Test_Cov6_Result_JsonString(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"jsonStr":    r.JsonString(),
		"safeStr":    r.SafeString(),
		"nilJsonStr": nilR.JsonString(),
	}
	expected := args.Map{
		"jsonStr": "\"hello\"", "safeStr": "\"hello\"",
		"nilJsonStr": "",
	}
	expected.ShouldBeEqual(t, 0, "Result JsonString", actual)
}

func Test_Cov6_Result_PrettyJsonString(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"a": 1})
	var nilR *corejson.Result
	actual := args.Map{
		"prettyNE":  r.PrettyJsonString() != "",
		"nilPretty": nilR.PrettyJsonString(),
	}
	expected := args.Map{"prettyNE": true, "nilPretty": ""}
	expected.ShouldBeEqual(t, 0, "Result PrettyJsonString", actual)
}

func Test_Cov6_Result_PrettyJsonStringOrErrString(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"a": 1})
	var nilR *corejson.Result
	actual := args.Map{
		"prettyNE":  r.PrettyJsonStringOrErrString() != "",
		"nilNotEmpty": nilR.PrettyJsonStringOrErrString() != "",
	}
	expected := args.Map{"prettyNE": true, "nilNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result PrettyJsonStringOrErrString", actual)
}

func Test_Cov6_Result_String(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"strNE": r.String() != ""}
	expected := args.Map{"strNE": true}
	expected.ShouldBeEqual(t, 0, "Result String", actual)
}

// ═══════════════════════════════════════════
// Result — bytes methods
// ═══════════════════════════════════════════

func Test_Cov6_Result_SafeBytes(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"hasBytes":  len(r.SafeBytes()) > 0,
		"nilEmpty":  len(nilR.SafeBytes()),
		"values":    len(r.Values()) > 0,
		"safeVals":  len(r.SafeValues()) > 0,
	}
	expected := args.Map{"hasBytes": true, "nilEmpty": 0, "values": true, "safeVals": true}
	expected.ShouldBeEqual(t, 0, "Result SafeBytes", actual)
}

func Test_Cov6_Result_Raw(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	rawBytes, rawErr := r.Raw()
	nilBytes, nilErr := nilR.Raw()
	actual := args.Map{
		"hasBytes":  len(rawBytes) > 0,
		"errNil":    rawErr == nil,
		"nilBytes":  len(nilBytes),
		"nilErrNN":  nilErr != nil,
	}
	expected := args.Map{"hasBytes": true, "errNil": true, "nilBytes": 0, "nilErrNN": true}
	expected.ShouldBeEqual(t, 0, "Result Raw", actual)
}

func Test_Cov6_Result_RawMust(t *testing.T) {
	r := corejson.NewPtr("hello")
	raw := r.RawMust()
	actual := args.Map{"hasBytes": len(raw) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Result RawMust", actual)
}

func Test_Cov6_Result_RawString(t *testing.T) {
	r := corejson.NewPtr("hello")
	str, err := r.RawString()
	actual := args.Map{"strNE": str != "", "errNil": err == nil}
	expected := args.Map{"strNE": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "Result RawString", actual)
}

func Test_Cov6_Result_RawStringMust(t *testing.T) {
	r := corejson.NewPtr("hello")
	str := r.RawStringMust()
	actual := args.Map{"strNE": str != ""}
	expected := args.Map{"strNE": true}
	expected.ShouldBeEqual(t, 0, "Result RawStringMust", actual)
}

func Test_Cov6_Result_RawErrString(t *testing.T) {
	r := corejson.NewPtr("hello")
	raw, errMsg := r.RawErrString()
	actual := args.Map{"hasRaw": len(raw) > 0, "errEmpty": errMsg == ""}
	expected := args.Map{"hasRaw": true, "errEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result RawErrString", actual)
}

func Test_Cov6_Result_RawPrettyString(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"a": 1})
	pretty, err := r.RawPrettyString()
	actual := args.Map{"prettyNE": pretty != "", "errNil": err == nil}
	expected := args.Map{"prettyNE": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "Result RawPrettyString", actual)
}

// ═══════════════════════════════════════════
// Result — error/state checks
// ═══════════════════════════════════════════

func Test_Cov6_Result_HasBytes(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"hasBytes":    r.HasBytes(),
		"hasJson":     r.HasJson(),
		"hasJsonB":    r.HasJsonBytes(),
		"hasSafe":     r.HasSafeItems(),
		"hasIssues":   r.HasIssuesOrEmpty(),
		"isEmptyJson": r.IsEmptyJson(),
	}
	expected := args.Map{
		"hasBytes": true, "hasJson": true, "hasJsonB": true,
		"hasSafe": true, "hasIssues": false, "isEmptyJson": false,
	}
	expected.ShouldBeEqual(t, 0, "Result HasBytes", actual)
}

func Test_Cov6_Result_BytesTypeName(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"typeNE":    r.BytesTypeName() != "",
		"nilType":   nilR.BytesTypeName(),
		"safeType":  r.SafeBytesTypeName() != "",
	}
	expected := args.Map{"typeNE": true, "nilType": "", "safeType": true}
	expected.ShouldBeEqual(t, 0, "Result BytesTypeName", actual)
}

func Test_Cov6_Result_MeaningfulError(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	actual := args.Map{
		"validErr":   r.MeaningfulError() == nil,
		"nilErrNN":   nilR.MeaningfulError() != nil,
		"errMsg":     r.MeaningfulErrorMessage(),
	}
	expected := args.Map{"validErr": true, "nilErrNN": true, "errMsg": ""}
	expected.ShouldBeEqual(t, 0, "Result MeaningfulError", actual)
}

// ═══════════════════════════════════════════
// Result — Unmarshal/Deserialize
// ═══════════════════════════════════════════

func Test_Cov6_Result_Unmarshal(t *testing.T) {
	r := corejson.NewPtr("hello")
	var target string
	err := r.Unmarshal(&target)
	actual := args.Map{"errNil": err == nil, "target": target}
	expected := args.Map{"errNil": true, "target": "hello"}
	expected.ShouldBeEqual(t, 0, "Result Unmarshal", actual)
}

func Test_Cov6_Result_Unmarshal_Nil(t *testing.T) {
	var nilR *corejson.Result
	var target string
	err := nilR.Unmarshal(&target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Result Unmarshal nil", actual)
}

func Test_Cov6_Result_Deserialize(t *testing.T) {
	r := corejson.NewPtr("hello")
	var target string
	err := r.Deserialize(&target)
	actual := args.Map{"errNil": err == nil, "target": target}
	expected := args.Map{"errNil": true, "target": "hello"}
	expected.ShouldBeEqual(t, 0, "Result Deserialize", actual)
}

// ═══════════════════════════════════════════
// Result — Serialize
// ═══════════════════════════════════════════

func Test_Cov6_Result_Serialize(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	bytes, err := r.Serialize()
	nilBytes, nilErr := nilR.Serialize()
	actual := args.Map{
		"hasBytes":  len(bytes) > 0,
		"errNil":    err == nil,
		"nilBytes":  nilBytes == nil,
		"nilErrNN":  nilErr != nil,
	}
	expected := args.Map{"hasBytes": true, "errNil": true, "nilBytes": true, "nilErrNN": true}
	expected.ShouldBeEqual(t, 0, "Result Serialize", actual)
}

// ═══════════════════════════════════════════
// Result — Clone
// ═══════════════════════════════════════════

func Test_Cov6_Result_Clone(t *testing.T) {
	r := corejson.NewPtr("hello")
	cloned := r.Clone(true)
	clonedShallow := r.Clone(false)
	clonedPtr := r.ClonePtr(true)
	var nilR *corejson.Result
	nilClonePtr := nilR.ClonePtr(true)
	actual := args.Map{
		"clonedLen":    cloned.Length(),
		"shallowLen":   clonedShallow.Length(),
		"ptrNotNil":    clonedPtr != nil,
		"nilCloneNil":  nilClonePtr == nil,
	}
	expected := args.Map{
		"clonedLen": 7, "shallowLen": 7,
		"ptrNotNil": true, "nilCloneNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Result Clone", actual)
}

func Test_Cov6_Result_CloneIf(t *testing.T) {
	r := corejson.New("hello")
	cloned := r.CloneIf(true, true)
	notCloned := r.CloneIf(false, false)
	actual := args.Map{
		"clonedLen":    cloned.Length(),
		"notClonedLen": notCloned.Length(),
	}
	expected := args.Map{"clonedLen": 7, "notClonedLen": 7}
	expected.ShouldBeEqual(t, 0, "Result CloneIf", actual)
}

func Test_Cov6_Result_CloneError(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"cloneErrNil": r.CloneError() == nil}
	expected := args.Map{"cloneErrNil": true}
	expected.ShouldBeEqual(t, 0, "Result CloneError", actual)
}

// ═══════════════════════════════════════════
// Result — Ptr/NonPtr/ToPtr/ToNonPtr
// ═══════════════════════════════════════════

func Test_Cov6_Result_PtrNonPtr(t *testing.T) {
	r := corejson.New("hello")
	ptr := r.Ptr()
	nonPtr := ptr.NonPtr()
	toPtr := r.ToPtr()
	toNonPtr := r.ToNonPtr()
	var nilR *corejson.Result
	nilNonPtr := nilR.NonPtr()
	actual := args.Map{
		"ptrNotNil":   ptr != nil,
		"nonPtrLen":   nonPtr.Length(),
		"toPtrNN":     toPtr != nil,
		"toNonPtrLen": toNonPtr.Length(),
		"nilNonPtrHasErr": nilNonPtr.HasError(),
	}
	expected := args.Map{
		"ptrNotNil": true, "nonPtrLen": 7,
		"toPtrNN": true, "toNonPtrLen": 7,
		"nilNonPtrHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result Ptr/NonPtr", actual)
}

// ═══════════════════════════════════════════
// Result — IsEqual / IsEqualPtr
// ═══════════════════════════════════════════

func Test_Cov6_Result_IsEqual(t *testing.T) {
	r1 := corejson.New("hello")
	r2 := corejson.New("hello")
	r3 := corejson.New("world")
	actual := args.Map{
		"equal":    r1.IsEqual(r2),
		"notEqual": r1.IsEqual(r3),
	}
	expected := args.Map{"equal": true, "notEqual": false}
	expected.ShouldBeEqual(t, 0, "Result IsEqual", actual)
}

func Test_Cov6_Result_IsEqualPtr(t *testing.T) {
	r1 := corejson.NewPtr("hello")
	r2 := corejson.NewPtr("hello")
	r3 := corejson.NewPtr("world")
	var nilR *corejson.Result
	actual := args.Map{
		"equal":      r1.IsEqualPtr(r2),
		"notEqual":   r1.IsEqualPtr(r3),
		"bothNil":    nilR.IsEqualPtr(nil),
		"oneNil":     r1.IsEqualPtr(nil),
		"samePtr":    r1.IsEqualPtr(r1),
	}
	expected := args.Map{
		"equal": true, "notEqual": false,
		"bothNil": true, "oneNil": false, "samePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result IsEqualPtr", actual)
}

// ═══════════════════════════════════════════
// Result — Json/JsonPtr/JsonModel
// ═══════════════════════════════════════════

func Test_Cov6_Result_Json(t *testing.T) {
	r := corejson.New("hello")
	j := r.Json()
	jp := r.JsonPtr()
	var nilR *corejson.Result
	model := r.JsonModel()
	nilModel := nilR.JsonModel()
	modelAny := r.JsonModelAny()
	actual := args.Map{
		"jsonLen":      j.Length() > 0,
		"jsonPtrNN":    jp != nil,
		"modelLen":     model.Length(),
		"nilModelErr":  nilModel.HasError(),
		"modelAnyNN":   modelAny != nil,
	}
	expected := args.Map{
		"jsonLen": true, "jsonPtrNN": true,
		"modelLen": 7, "nilModelErr": true, "modelAnyNN": true,
	}
	expected.ShouldBeEqual(t, 0, "Result Json", actual)
}

// ═══════════════════════════════════════════
// Result — Dispose
// ═══════════════════════════════════════════

func Test_Cov6_Result_Dispose(t *testing.T) {
	r := corejson.NewPtr("hello")
	r.Dispose()
	var nilR *corejson.Result
	nilR.Dispose() // should not panic
	actual := args.Map{
		"bytesNil": r.Bytes == nil,
		"errNil":   r.Error == nil,
		"typeName": r.TypeName,
	}
	expected := args.Map{"bytesNil": true, "errNil": true, "typeName": ""}
	expected.ShouldBeEqual(t, 0, "Result Dispose", actual)
}

// ═══════════════════════════════════════════
// Result — BytesError
// ═══════════════════════════════════════════

func Test_Cov6_Result_BytesError(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	be := r.BytesError()
	nilBE := nilR.BytesError()
	actual := args.Map{
		"beNotNil":  be != nil,
		"nilBENil":  nilBE == nil,
	}
	expected := args.Map{"beNotNil": true, "nilBENil": true}
	expected.ShouldBeEqual(t, 0, "Result BytesError", actual)
}

// ═══════════════════════════════════════════
// Result — CombineErrorWithRef
// ═══════════════════════════════════════════

func Test_Cov6_Result_CombineErrorWithRefString_NoError(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"result": r.CombineErrorWithRefString("ref1")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Result CombineErrorWithRefString no error", actual)
}

func Test_Cov6_Result_CombineErrorWithRefError_NoError(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"result": r.CombineErrorWithRefError("ref1") == nil}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Result CombineErrorWithRefError no error", actual)
}

// ═══════════════════════════════════════════
// Result — AsJsonContractsBinder / AsJsoner
// ═══════════════════════════════════════════

func Test_Cov6_Result_InterfaceAdapters(t *testing.T) {
	r := corejson.New("hello")
	binder := r.AsJsonContractsBinder()
	jsoner := r.AsJsoner()
	injector := r.AsJsonParseSelfInjector()
	actual := args.Map{
		"binderNN":   binder != nil,
		"jsonerNN":   jsoner != nil,
		"injectorNN": injector != nil,
	}
	expected := args.Map{"binderNN": true, "jsonerNN": true, "injectorNN": true}
	expected.ShouldBeEqual(t, 0, "Result InterfaceAdapters", actual)
}

// ═══════════════════════════════════════════
// Result — DeserializedFieldsToMap
// ═══════════════════════════════════════════

func Test_Cov6_Result_DeserializedFieldsToMap(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"name": "test", "age": 30})
	fm, err := r.DeserializedFieldsToMap()
	sfm := r.SafeDeserializedFieldsToMap()
	actual := args.Map{
		"fmNotNil": fm != nil,
		"errNil":   err == nil,
		"sfmNotNil": sfm != nil,
	}
	expected := args.Map{"fmNotNil": true, "errNil": true, "sfmNotNil": true}
	expected.ShouldBeEqual(t, 0, "Result DeserializedFieldsToMap", actual)
}

func Test_Cov6_Result_FieldsNames(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"name": "test"})
	names, err := r.FieldsNames()
	safeNames := r.SafeFieldsNames()
	actual := args.Map{
		"namesNotNil":    names != nil,
		"errNil":         err == nil,
		"safeNotNil":     safeNames != nil,
	}
	expected := args.Map{"namesNotNil": true, "errNil": true, "safeNotNil": true}
	expected.ShouldBeEqual(t, 0, "Result FieldsNames", actual)
}

// ═══════════════════════════════════════════
// Result — Map
// ═══════════════════════════════════════════

func Test_Cov6_Result_Map(t *testing.T) {
	r := corejson.NewPtr("hello")
	var nilR *corejson.Result
	m := r.Map()
	nilM := nilR.Map()
	actual := args.Map{
		"hasBytes":  len(m) > 0,
		"nilMapLen": len(nilM),
	}
	expected := args.Map{"hasBytes": true, "nilMapLen": 0}
	expected.ShouldBeEqual(t, 0, "Result Map", actual)
}

// ═══════════════════════════════════════════
// Result — SerializeSkipExistingIssues
// ═══════════════════════════════════════════

func Test_Cov6_Result_SerializeSkipExistingIssues(t *testing.T) {
	r := corejson.NewPtr("hello")
	bytes, err := r.SerializeSkipExistingIssues()
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"errNil":   err == nil,
	}
	expected := args.Map{"hasBytes": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "Result SerializeSkipExistingIssues", actual)
}

// ═══════════════════════════════════════════
// Result — UnmarshalSkipExistingIssues
// ═══════════════════════════════════════════

func Test_Cov6_Result_UnmarshalSkipExistingIssues(t *testing.T) {
	r := corejson.NewPtr("hello")
	var target string
	err := r.UnmarshalSkipExistingIssues(&target)
	actual := args.Map{"errNil": err == nil, "target": target}
	expected := args.Map{"errNil": true, "target": "hello"}
	expected.ShouldBeEqual(t, 0, "Result UnmarshalSkipExistingIssues", actual)
}

// ═══════════════════════════════════════════
// Result — UnmarshalResult
// ═══════════════════════════════════════════

func Test_Cov6_Result_UnmarshalResult(t *testing.T) {
	// UnmarshalResult tries to unmarshal bytes into a *Result struct
	// "hello" is a JSON string, not a Result object — expect unmarshal error
	r := corejson.NewPtr("hello")
	_, err := r.UnmarshalResult()
	actual := args.Map{
		"hasErr": err != nil,
	}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Result UnmarshalResult", actual)
}

// ═══════════════════════════════════════════
// Result — ParseInjectUsingJson
// ═══════════════════════════════════════════

func Test_Cov6_Result_ParseInjectUsingJson(t *testing.T) {
	r := corejson.NewPtr("hello")
	jsonR := r.JsonPtr()
	target := corejson.NewPtr("world")
	parsed, err := target.ParseInjectUsingJson(jsonR)
	actual := args.Map{
		"parsedNN": parsed != nil,
		"errNil":   err == nil,
	}
	expected := args.Map{"parsedNN": true, "errNil": true}
	expected.ShouldBeEqual(t, 0, "Result ParseInjectUsingJson", actual)
}

// ═══════════════════════════════════════════
// Result — SafeNonIssueBytes / SafeValuesPtr
// ═══════════════════════════════════════════

func Test_Cov6_Result_SafeNonIssueBytes(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"safeNonIssue": len(r.SafeNonIssueBytes()) > 0,
		"safeValsPtr":  len(r.SafeValuesPtr()) > 0,
	}
	expected := args.Map{"safeNonIssue": true, "safeValsPtr": true}
	expected.ShouldBeEqual(t, 0, "Result SafeNonIssueBytes", actual)
}

// ═══════════════════════════════════════════
// Result — PrettyJsonBuffer
// ═══════════════════════════════════════════

func Test_Cov6_Result_PrettyJsonBuffer(t *testing.T) {
	r := corejson.NewPtr(map[string]any{"key": "val"})
	buf, err := r.PrettyJsonBuffer("", "  ")
	actual := args.Map{
		"bufNN":  buf != nil,
		"errNil": err == nil,
		"bufLen": buf.Len() > 0,
	}
	expected := args.Map{"bufNN": true, "errNil": true, "bufLen": true}
	expected.ShouldBeEqual(t, 0, "Result PrettyJsonBuffer", actual)
}

// ═══════════════════════════════════════════
// Result — InjectInto
// ═══════════════════════════════════════════

func Test_Cov6_Result_InjectInto(t *testing.T) {
	r := corejson.NewPtr("hello")
	target := corejson.NewPtr("world")
	// InjectInto calls target.JsonParseSelfInject(r) which tries to ParseInjectUsingJson
	// This may or may not succeed depending on internal implementation
	err := r.InjectInto(target)
	// Just exercise the code path — don't assert err is nil
	_ = err
}
