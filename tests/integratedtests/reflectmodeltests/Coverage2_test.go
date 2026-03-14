package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── FieldProcessor ──

func Test_Cov2_FieldProcessor_IsFieldType(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name: "A", Index: 0,
		FieldType: reflect.TypeOf(42),
	}
	var nilFP *reflectmodel.FieldProcessor
	actual := args.Map{
		"isInt":    fp.IsFieldType(reflect.TypeOf(42)),
		"isString": fp.IsFieldType(reflect.TypeOf("")),
		"nil":      nilFP.IsFieldType(reflect.TypeOf(42)),
	}
	expected := args.Map{"isInt": true, "isString": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor IsFieldType", actual)
}

func Test_Cov2_FieldProcessor_IsFieldKind(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		FieldType: reflect.TypeOf(42),
	}
	var nilFP *reflectmodel.FieldProcessor
	actual := args.Map{
		"isInt":  fp.IsFieldKind(reflect.Int),
		"isStr":  fp.IsFieldKind(reflect.String),
		"nil":    nilFP.IsFieldKind(reflect.Int),
	}
	expected := args.Map{"isInt": true, "isStr": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor IsFieldKind", actual)
}

// ── MethodProcessor extended ──

func Test_Cov2_MethodProcessor_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{
		"hasValid":   mp.HasValidFunc(),
		"funcName":   mp.GetFuncName(),
		"isInvalid":  mp.IsInvalid(),
		"funcNil":    mp.Func() == nil,
		"argsCount":  mp.ArgsCount(),
		"returnLen":  mp.ReturnLength(),
		"isPublic":   mp.IsPublicMethod(),
		"isPrivate":  mp.IsPrivateMethod(),
		"argsLength": mp.ArgsLength(),
		"typeNil":    mp.GetType() == nil,
	}
	expected := args.Map{
		"hasValid": false, "funcName": "", "isInvalid": true,
		"funcNil": true, "argsCount": -1, "returnLen": -1,
		"isPublic": false, "isPrivate": false, "argsLength": -1,
		"typeNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor Nil", actual)
}

func Test_Cov2_MethodProcessor_IsEqual(t *testing.T) {
	var mp1, mp2 *reflectmodel.MethodProcessor
	actual := args.Map{"nilNil": mp1.IsEqual(mp2)}
	expected := args.Map{"nilNil": true}
	expected.ShouldBeEqual(t, 0, "MethodProcessor IsEqual nil", actual)
}

func Test_Cov2_MethodProcessor_IsNotEqual(t *testing.T) {
	var mp1, mp2 *reflectmodel.MethodProcessor
	actual := args.Map{"notEqual": mp1.IsNotEqual(mp2)}
	expected := args.Map{"notEqual": false}
	expected.ShouldBeEqual(t, 0, "MethodProcessor IsNotEqual", actual)
}

func Test_Cov2_MethodProcessor_GetOutArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	result := mp.GetOutArgsTypes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor GetOutArgsTypes nil", actual)
}

func Test_Cov2_MethodProcessor_GetInArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	result := mp.GetInArgsTypes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor GetInArgsTypes nil", actual)
}

func Test_Cov2_MethodProcessor_GetInArgsTypesNames_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	result := mp.GetInArgsTypesNames()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor GetInArgsTypesNames nil", actual)
}

// ── rvUtils ──

func Test_Cov2_RvUtils_IsNull(t *testing.T) {
	actual := args.Map{
		"nil":    reflectmodel.Utils.IsNull(nil),
		"nonNil": reflectmodel.Utils.IsNull(42),
	}
	expected := args.Map{"nil": true, "nonNil": false}
	expected.ShouldBeEqual(t, 0, "RvUtils IsNull", actual)
}

func Test_Cov2_RvUtils_IndexToPosition(t *testing.T) {
	actual := args.Map{
		"first": reflectmodel.Utils.IndexToPosition(0),
		"third": reflectmodel.Utils.IndexToPosition(2),
	}
	expected := args.Map{"first": "1st", "third": "3rd"}
	expected.ShouldBeEqual(t, 0, "RvUtils IndexToPosition", actual)
}

func Test_Cov2_RvUtils_ArgsToReflectValues(t *testing.T) {
	result := reflectmodel.Utils.ArgsToReflectValues([]any{1, "a"})
	emptyResult := reflectmodel.Utils.ArgsToReflectValues(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "RvUtils ArgsToReflectValues", actual)
}

func Test_Cov2_RvUtils_ReflectValueToAnyValue(t *testing.T) {
	rv := reflect.ValueOf(42)
	result := reflectmodel.Utils.ReflectValueToAnyValue(rv)
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "RvUtils ReflectValueToAnyValue", actual)
}

func Test_Cov2_RvUtils_InterfacesToTypes(t *testing.T) {
	result := reflectmodel.Utils.InterfacesToTypes([]any{1, "a"})
	emptyResult := reflectmodel.Utils.InterfacesToTypes(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "RvUtils InterfacesToTypes", actual)
}

func Test_Cov2_RvUtils_InterfacesToTypesNamesWithValues(t *testing.T) {
	result := reflectmodel.Utils.InterfacesToTypesNamesWithValues([]any{1, "a"})
	emptyResult := reflectmodel.Utils.InterfacesToTypesNamesWithValues(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "RvUtils InterfacesToTypesNamesWithValues", actual)
}

func Test_Cov2_RvUtils_WithSpaces(t *testing.T) {
	result := reflectmodel.Utils.WithSpaces(2, "a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RvUtils WithSpaces", actual)
}

func Test_Cov2_RvUtils_PrependWithSpaces(t *testing.T) {
	result := reflectmodel.Utils.PrependWithSpaces(2, []string{"a"}, "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RvUtils PrependWithSpaces", actual)
}

func Test_Cov2_RvUtils_ReflectValuesToInterfaces(t *testing.T) {
	rvs := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf("a")}
	result := reflectmodel.Utils.ReflectValuesToInterfaces(rvs)
	emptyResult := reflectmodel.Utils.ReflectValuesToInterfaces(nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "RvUtils ReflectValuesToInterfaces", actual)
}

func Test_Cov2_RvUtils_IsReflectTypeMatchAny(t *testing.T) {
	ok, err := reflectmodel.Utils.IsReflectTypeMatchAny(42, 43)
	notOk, notErr := reflectmodel.Utils.IsReflectTypeMatchAny(42, "hello")
	actual := args.Map{"ok": ok, "noErr": err == nil, "notOk": notOk, "hasErr": notErr != nil}
	expected := args.Map{"ok": true, "noErr": true, "notOk": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "RvUtils IsReflectTypeMatchAny", actual)
}

func Test_Cov2_RvUtils_IsReflectTypeMatch(t *testing.T) {
	intType := reflect.TypeOf(42)
	strType := reflect.TypeOf("")
	ok, err := reflectmodel.Utils.IsReflectTypeMatch(intType, intType)
	notOk, notErr := reflectmodel.Utils.IsReflectTypeMatch(intType, strType)
	actual := args.Map{"ok": ok, "noErr": err == nil, "notOk": notOk, "hasErr": notErr != nil}
	expected := args.Map{"ok": true, "noErr": true, "notOk": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "RvUtils IsReflectTypeMatch", actual)
}
