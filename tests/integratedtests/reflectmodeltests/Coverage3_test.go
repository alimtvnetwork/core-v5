package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── ReflectValueKind ──

func Test_Cov3_ReflectValueKind_InvalidModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")
	actual := args.Map{
		"isValid":   rvk.IsValid,
		"isInvalid": rvk.IsInvalid(),
		"hasError":  rvk.HasError(),
		"isEmptyErr": rvk.IsEmptyError(),
		"errMsg":    rvk.Error.Error(),
	}
	expected := args.Map{
		"isValid": false, "isInvalid": true,
		"hasError": true, "isEmptyErr": false,
		"errMsg": "test error",
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind InvalidModel", actual)
}

func Test_Cov3_ReflectValueKind_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{
		"isInvalid":  rvk.IsInvalid(),
		"hasError":   rvk.HasError(),
		"isEmptyErr": rvk.IsEmptyError(),
		"actual":     rvk.ActualInstance() == nil,
		"pkgPath":    rvk.PkgPath(),
		"typeName":   rvk.TypeName(),
		"pointerRv":  rvk.PointerRv() == nil,
		"pointerInf": rvk.PointerInterface() == nil,
	}
	expected := args.Map{
		"isInvalid": true, "hasError": false,
		"isEmptyErr": true, "actual": true,
		"pkgPath": "", "typeName": "",
		"pointerRv": true, "pointerInf": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind nil receiver", actual)
}

func Test_Cov3_ReflectValueKind_Valid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf("hello"),
		Kind:            reflect.String,
		Error:           nil,
	}
	actual := args.Map{
		"isValid":    rvk.IsValid,
		"isInvalid":  rvk.IsInvalid(),
		"hasError":   rvk.HasError(),
		"actual":     rvk.ActualInstance(),
		"pkgPathEmp": rvk.PkgPath() == "",
		"typeName":   rvk.TypeName() != "",
	}
	expected := args.Map{
		"isValid": true, "isInvalid": false,
		"hasError": false, "actual": "hello",
		"pkgPathEmp": true, "typeName": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind valid", actual)
}

func Test_Cov3_ReflectValueKind_PointerRv_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
	}
	rv := rvk.PointerRv()
	actual := args.Map{"notNil": rv != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind PointerRv invalid", actual)
}

// ── ReflectValue ──

func Test_Cov3_ReflectValue_Fields(t *testing.T) {
	rv := reflectmodel.ReflectValue{
		TypeName:     "TestType",
		FieldsNames:  []string{"Field1", "Field2"},
		MethodsNames: []string{"Method1"},
		RawData:      "raw",
	}
	actual := args.Map{
		"typeName":    rv.TypeName,
		"fieldsLen":   len(rv.FieldsNames),
		"methodsLen":  len(rv.MethodsNames),
		"rawData":     rv.RawData,
	}
	expected := args.Map{
		"typeName": "TestType", "fieldsLen": 2,
		"methodsLen": 1, "rawData": "raw",
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue fields", actual)
}

// ── FieldProcessor ──

func Test_Cov3_FieldProcessor_Nil(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	actual := args.Map{
		"isFieldType": fp.IsFieldType(reflect.TypeOf("")),
		"isFieldKind": fp.IsFieldKind(reflect.String),
	}
	expected := args.Map{
		"isFieldType": false, "isFieldKind": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor nil", actual)
}

func Test_Cov3_FieldProcessor_Valid(t *testing.T) {
	strType := reflect.TypeOf("")
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: strType,
	}
	actual := args.Map{
		"isFieldType":  fp.IsFieldType(strType),
		"isFieldKind":  fp.IsFieldKind(reflect.String),
		"wrongKind":    fp.IsFieldKind(reflect.Int),
	}
	expected := args.Map{
		"isFieldType": true, "isFieldKind": true, "wrongKind": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor valid", actual)
}

// ── MethodProcessor ──

func Test_Cov3_MethodProcessor_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{
		"isInvalid":   mp.IsInvalid(),
		"hasValidFunc": mp.HasValidFunc(),
		"func":         mp.Func() == nil,
		"returnLen":    mp.ReturnLength(),
	}
	expected := args.Map{
		"isInvalid": true, "hasValidFunc": false,
		"func": true, "returnLen": -1,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor nil", actual)
}

func Test_Cov3_MethodProcessor_IsEqual_BothNil(t *testing.T) {
	var mp1, mp2 *reflectmodel.MethodProcessor
	actual := args.Map{
		"equal":    mp1.IsEqual(mp2),
		"notEqual": mp1.IsNotEqual(mp2),
	}
	expected := args.Map{"equal": true, "notEqual": false}
	expected.ShouldBeEqual(t, 0, "MethodProcessor IsEqual both nil", actual)
}

func Test_Cov3_MethodProcessor_IsEqual_OneNil(t *testing.T) {
	mp := &reflectmodel.MethodProcessor{Name: "Test"}
	actual := args.Map{"equal": mp.IsEqual(nil)}
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "MethodProcessor IsEqual one nil", actual)
}

func Test_Cov3_MethodProcessor_GetType_Invalid(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"isNil": mp.GetType() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MethodProcessor GetType nil", actual)
}

func Test_Cov3_MethodProcessor_GetInOutArgsTypes_Invalid(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{
		"inLen":       len(mp.GetInArgsTypes()),
		"outLen":      len(mp.GetOutArgsTypes()),
		"inNamesLen":  len(mp.GetInArgsTypesNames()),
	}
	expected := args.Map{"inLen": 0, "outLen": 0, "inNamesLen": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor GetInOutArgsTypes nil", actual)
}
