package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// BytesCloneIf
// =============================================================================

func Test_Cov40_BytesCloneIf_NotDeep(t *testing.T) {
	b := []byte("hello")
	c := corejson.BytesCloneIf(false, b)
	actual := args.Map{"len": len(c)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf not deep", actual)
}

func Test_Cov40_BytesCloneIf_DeepEmpty(t *testing.T) {
	c := corejson.BytesCloneIf(true, []byte{})
	actual := args.Map{"len": len(c)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf deep empty", actual)
}

func Test_Cov40_BytesCloneIf_DeepValid(t *testing.T) {
	b := []byte("hello")
	c := corejson.BytesCloneIf(true, b)
	actual := args.Map{"len": len(c), "same": &c[0] != &b[0]}
	expected := args.Map{"len": 5, "same": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf deep valid", actual)
}

// =============================================================================
// BytesDeepClone
// =============================================================================

func Test_Cov40_BytesDeepClone_Empty(t *testing.T) {
	c := corejson.BytesDeepClone([]byte{})
	actual := args.Map{"len": len(c)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone empty", actual)
}

func Test_Cov40_BytesDeepClone_Valid(t *testing.T) {
	b := []byte("hello")
	c := corejson.BytesDeepClone(b)
	actual := args.Map{"len": len(c)}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone valid", actual)
}

// =============================================================================
// BytesToString / BytesToPrettyString
// =============================================================================

func Test_Cov40_BytesToString_Empty(t *testing.T) {
	actual := args.Map{"r": corejson.BytesToString([]byte{})}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesToString empty", actual)
}

func Test_Cov40_BytesToString_Valid(t *testing.T) {
	actual := args.Map{"r": corejson.BytesToString([]byte("hello"))}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesToString valid", actual)
}

func Test_Cov40_BytesToPrettyString_Empty(t *testing.T) {
	actual := args.Map{"r": corejson.BytesToPrettyString([]byte{})}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString empty", actual)
}

func Test_Cov40_BytesToPrettyString_Valid(t *testing.T) {
	b := []byte(`{"a":1}`)
	s := corejson.BytesToPrettyString(b)
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "BytesToPrettyString valid", actual)
}

// =============================================================================
// JsonStringOrErrMsg
// =============================================================================

func Test_Cov40_JsonStringOrErrMsg_Valid(t *testing.T) {
	s := corejson.JsonStringOrErrMsg("hello")
	actual := args.Map{"r": s}
	expected := args.Map{"r": `"hello"`}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg valid", actual)
}

func Test_Cov40_JsonStringOrErrMsg_Unmarshalable(t *testing.T) {
	s := corejson.JsonStringOrErrMsg(make(chan int))
	actual := args.Map{"hasContent": len(s) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg unmarshalable", actual)
}

// =============================================================================
// New / NewPtr
// =============================================================================

func Test_Cov40_New_Valid(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{"noErr": !r.HasError(), "hasBytes": r.Length() > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "New valid", actual)
}

func Test_Cov40_New_Unmarshalable(t *testing.T) {
	r := corejson.New(make(chan int))
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New unmarshalable", actual)
}

func Test_Cov40_NewPtr_Valid(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{"notNil": r != nil, "noErr": !r.HasError()}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "NewPtr valid", actual)
}

func Test_Cov40_NewPtr_Unmarshalable(t *testing.T) {
	r := corejson.NewPtr(make(chan int))
	actual := args.Map{"hasErr": r.HasError()}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NewPtr unmarshalable", actual)
}

// =============================================================================
// emptyCreator
// =============================================================================

func Test_Cov40_Empty_Result(t *testing.T) {
	r := corejson.Empty.Result()
	actual := args.Map{"empty": r.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result", actual)
}

func Test_Cov40_Empty_ResultPtr(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	actual := args.Map{"notNil": r != nil, "empty": r.IsEmpty()}
	expected := args.Map{"notNil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtr", actual)
}

func Test_Cov40_Empty_ResultWithErr(t *testing.T) {
	r := corejson.Empty.ResultWithErr("int", nil)
	actual := args.Map{"typeName": r.TypeName}
	expected := args.Map{"typeName": "int"}
	expected.ShouldBeEqual(t, 0, "Empty.ResultWithErr", actual)
}

func Test_Cov40_Empty_ResultPtrWithErr(t *testing.T) {
	r := corejson.Empty.ResultPtrWithErr("int", nil)
	actual := args.Map{"notNil": r != nil, "typeName": r.TypeName}
	expected := args.Map{"notNil": true, "typeName": "int"}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtrWithErr", actual)
}

func Test_Cov40_Empty_BytesCollection(t *testing.T) {
	bc := corejson.Empty.BytesCollection()
	actual := args.Map{"len": bc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.BytesCollection", actual)
}

func Test_Cov40_Empty_BytesCollectionPtr(t *testing.T) {
	bc := corejson.Empty.BytesCollectionPtr()
	actual := args.Map{"notNil": bc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.BytesCollectionPtr", actual)
}

func Test_Cov40_Empty_ResultsCollection(t *testing.T) {
	rc := corejson.Empty.ResultsCollection()
	actual := args.Map{"notNil": rc != nil, "len": rc.Length()}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsCollection", actual)
}

func Test_Cov40_Empty_ResultsPtrCollection(t *testing.T) {
	rc := corejson.Empty.ResultsPtrCollection()
	actual := args.Map{"notNil": rc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.ResultsPtrCollection", actual)
}

func Test_Cov40_Empty_MapResults(t *testing.T) {
	mr := corejson.Empty.MapResults()
	actual := args.Map{"notNil": mr != nil, "len": mr.Length()}
	expected := args.Map{"notNil": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.MapResults", actual)
}

// =============================================================================
// KeyAny / KeyWithResult
// =============================================================================

func Test_Cov40_KeyAny(t *testing.T) {
	ka := corejson.KeyAny{Key: "x", AnyInf: 42}
	actual := args.Map{"key": ka.Key, "val": ka.AnyInf}
	expected := args.Map{"key": "x", "val": 42}
	expected.ShouldBeEqual(t, 0, "KeyAny", actual)
}

func Test_Cov40_KeyWithResult(t *testing.T) {
	kwr := corejson.KeyWithResult{Key: "x", Result: corejson.New("hello")}
	actual := args.Map{"key": kwr.Key, "noErr": !kwr.Result.HasError()}
	expected := args.Map{"key": "x", "noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyWithResult", actual)
}
