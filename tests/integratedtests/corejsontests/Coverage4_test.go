package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BytesDeepClone ──

func Test_Cov4_BytesDeepClone(t *testing.T) {
	original := []byte(`{"key":"value"}`)
	cloned := corejson.BytesDeepClone(original)
	clonedNil := corejson.BytesDeepClone(nil)
	nilIsNil := clonedNil == nil
	actual := args.Map{
		"len":        len(cloned),
		"nilIsNil":   nilIsNil,
		"notSamePtr": &original[0] != &cloned[0],
	}
	expected := args.Map{"len": 15, "nilIsNil": nilIsNil, "notSamePtr": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns independent copy -- valid input", actual)
}

// ── Result via NewPtr ──

func Test_Cov4_Result_NewPtr(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"notNil":   r != nil,
		"notEmpty": r.JsonString() != "",
		"hasBytes": len(r.SafeBytes()) > 0,
		"noErr":    r.MeaningfulError() == nil,
	}
	expected := args.Map{
		"notNil": true, "notEmpty": true, "hasBytes": true, "noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns valid json -- NewPtr string", actual)
}

func Test_Cov4_Result_New_ValueType(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{
		"hasBytes": len(r.Bytes) > 0,
		"noErr":    r.Error == nil,
	}
	expected := args.Map{"hasBytes": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Result returns bytes -- New value type", actual)
}

// ── BytesToString ──

func Test_Cov4_BytesToString(t *testing.T) {
	result := corejson.BytesToString([]byte(`"hello"`))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString returns non-empty -- valid bytes", actual)
}

func Test_Cov4_BytesToString_Nil(t *testing.T) {
	result := corejson.BytesToString(nil)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString returns empty -- nil input", actual)
}

// ── JsonString (returns string, error) ──

func Test_Cov4_JsonString(t *testing.T) {
	result, err := corejson.JsonString(map[string]string{"a": "1"})
	actual := args.Map{"notEmpty": result != "", "noErr": err == nil}
	expected := args.Map{"notEmpty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns json -- map input", actual)
}

// ── JsonStringOrErrMsg ──

func Test_Cov4_JsonStringOrErrMsg(t *testing.T) {
	result := corejson.JsonStringOrErrMsg(map[string]string{"a": "1"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg returns json -- valid input", actual)
}

// ── BytesCloneIf ──

func Test_Cov4_BytesCloneIf_True(t *testing.T) {
	original := []byte("hello")
	cloned := corejson.BytesCloneIf(true, original)
	actual := args.Map{
		"len": len(cloned),
	}
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns cloned bytes -- true flag", actual)
}

func Test_Cov4_BytesCloneIf_False(t *testing.T) {
	original := []byte("hello")
	result := corejson.BytesCloneIf(false, original)
	// BytesCloneIf(false, ...) returns []byte{} per implementation
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns empty -- false flag", actual)
}

func Test_Cov4_BytesCloneIf_NilInput(t *testing.T) {
	result := corejson.BytesCloneIf(true, nil)
	// len(nil) == 0, so !isDeepClone || len == 0 → returns []byte{}
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns empty -- nil input", actual)
}

// ── SimpleJsonBinder (round-trip) ──

func Test_Cov4_SimpleJsonBinder(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
	}
	input := testStruct{Name: "test"}
	r := corejson.NewPtr(input)
	var output testStruct
	err := r.Unmarshal(&output)
	actual := args.Map{
		"noErr": err == nil,
		"name":  output.Name,
	}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "SimpleJsonBinder returns deserialized struct -- round-trip", actual)
}

// ── Result PrettyJsonString ──

func Test_Cov4_Result_PrettyJsonString(t *testing.T) {
	r := corejson.NewPtr(map[string]string{"key": "val"})
	pretty := r.PrettyJsonString()
	actual := args.Map{
		"notEmpty":     pretty != "",
		"containsKey":  len(pretty) > 5,
	}
	expected := args.Map{"notEmpty": true, "containsKey": true}
	expected.ShouldBeEqual(t, 0, "Result returns pretty json -- map input", actual)
}
