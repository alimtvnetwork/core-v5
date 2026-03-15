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
	actual := args.Map{
		"len":        len(cloned),
		"nilIsNil":   clonedNil == nil,
		"notSamePtr": &original[0] != &cloned[0],
	}
	expected := args.Map{"len": 15, "nilIsNil": true, "notSamePtr": true}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone", actual)
}

// ── Result ──

func Test_Cov4_Result_New(t *testing.T) {
	r := corejson.New("hello")
	actual := args.Map{
		"notEmpty":   r.JsonString() != "",
		"hasBytes":   len(r.SafeBytes()) > 0,
		"noErr":      r.Err() == nil,
	}
	expected := args.Map{
		"notEmpty": true, "hasBytes": true, "noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result New", actual)
}

func Test_Cov4_Result_NewPtr(t *testing.T) {
	r := corejson.NewPtr("hello")
	actual := args.Map{
		"notNil":   r != nil,
		"notEmpty": r.JsonString() != "",
	}
	expected := args.Map{"notNil": true, "notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Result NewPtr", actual)
}

// ── BytesToString ──

func Test_Cov4_BytesToString(t *testing.T) {
	result := corejson.BytesToString([]byte(`"hello"`))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString", actual)
}

func Test_Cov4_BytesToString_Nil(t *testing.T) {
	result := corejson.BytesToString(nil)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString nil", actual)
}

// ── JsonString ──

func Test_Cov4_JsonString(t *testing.T) {
	result := corejson.JsonString(map[string]string{"a": "1"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString", actual)
}

// ── JsonStringOrErrMsg ──

func Test_Cov4_JsonStringOrErrMsg(t *testing.T) {
	result := corejson.JsonStringOrErrMsg(map[string]string{"a": "1"})
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg", actual)
}

// ── BytesCloneIf ──

func Test_Cov4_BytesCloneIf_True(t *testing.T) {
	original := []byte("hello")
	cloned := corejson.BytesCloneIf(true, original)
	actual := args.Map{
		"len":        len(cloned),
		"notSamePtr": &original[0] != &cloned[0],
	}
	expected := args.Map{"len": 5, "notSamePtr": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf true", actual)
}

func Test_Cov4_BytesCloneIf_False(t *testing.T) {
	original := []byte("hello")
	result := corejson.BytesCloneIf(false, original)
	actual := args.Map{"samePtr": &original[0] == &result[0]}
	expected := args.Map{"samePtr": true}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf false", actual)
}

// ── SimpleJsonBinder ──

func Test_Cov4_SimpleJsonBinder(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
	}
	input := testStruct{Name: "test"}
	result := corejson.New(input)
	var output testStruct
	err := result.Unmarshal(&output)
	actual := args.Map{
		"noErr": err == nil,
		"name":  output.Name,
	}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "SimpleJsonBinder round-trip", actual)
}
